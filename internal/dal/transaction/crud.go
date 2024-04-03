package transaction

import (
	"context"
	"errors"
	"fmt"
	"wash-payment/internal/app"
	"wash-payment/internal/app/entity"
	"wash-payment/internal/dal/conversions"
	"wash-payment/internal/dal/dbmodels"

	"github.com/gocraft/dbr/v2"
	"github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
)

var columns = []string{"id", "organization_id", "group_id", "amount", "operation", "created_at", "for_date", "service", "stations_count", "user_id", "wash_server_id"}
var selectColumns = []string{
	"t.id AS t_id",
	"t.organization_id AS t_organization_id",
	"t.amount AS t_amount",
	"t.operation AS t_operation",
	"t.created_at AS t_created_at",
	"t.for_date AS t_for_date",
	"t.service AS t_service",
	"t.stations_count AS t_stations_count",
	"g.id AS g_id",
	"g.organization_id AS g_organization_id",
	"g.name AS g_name",
	"g.description AS g_description",
	"g.version AS g_version",
	"g.deleted AS g_deleted",
	"ws.id AS ws_id",
	"ws.title AS ws_title",
	"ws.description AS ws_description",
	"ws.group_id AS ws_group_id",
	"ws.version AS ws_version",
	"ws.deleted AS ws_deleted",
	"u.id AS u_id",
	"u.email AS u_email",
	"u.name AS u_name",
	"u.role AS u_role",
	"u.organization_id AS u_organization_id",
	"u.version AS u_version",
}

func (r *transactionRepo) Get(ctx context.Context, transactionID uuid.UUID) (entity.Transaction, error) {
	op := "failed to get transaction by ID: %w"

	var dbTransaction dbmodels.Transaction
	err := r.db.NewSession(nil).
		Select(selectColumns...).
		From(dbr.I(dbmodels.TransactionsTable).As("t")).
		LeftJoin(dbr.I(dbmodels.GroupsTable).As("g"), "t.group_id = g.id").
		LeftJoin(dbr.I(dbmodels.WashServersTable).As("ws"), "t.wash_server_id = ws.id").
		LeftJoin(dbr.I(dbmodels.UsersTable).As("u"), "t.user_id = u.id").
		Where("t.id = ?", transactionID).
		LoadOneContext(ctx, &dbTransaction)

	if err != nil {
		if errors.Is(err, dbr.ErrNotFound) {
			err = app.ErrNotFound
		}

		return entity.Transaction{}, fmt.Errorf(op, err)
	}

	return conversions.TransactionFromDB(dbTransaction), nil
}

func (r *transactionRepo) List(ctx context.Context, filter entity.TransactionFilter) (entity.Page[entity.Transaction], error) {
	op := "failed to get transactions list: %w"

	var count int
	query := r.db.NewSession(nil).
		Select(dbmodels.CountSelect).
		From(dbr.I(dbmodels.TransactionsTable).As("t"))

	query = buildFilter(query, filter)

	err := query.LoadOneContext(ctx, &count)
	if err != nil {
		return entity.Page[entity.Transaction]{}, fmt.Errorf(op, err)
	}

	var dbTransaction []dbmodels.Transaction
	query = r.db.NewSession(nil).
		Select(selectColumns...).
		From(dbr.I(dbmodels.TransactionsTable).As("t")).
		LeftJoin(dbr.I(dbmodels.GroupsTable).As("g"), "t.group_id = g.id").
		LeftJoin(dbr.I(dbmodels.WashServersTable).As("ws"), "t.wash_server_id = ws.id").
		LeftJoin(dbr.I(dbmodels.UsersTable).As("u"), "t.user_id = u.id").
		OrderDesc("t.created_at").
		Paginate(uint64(filter.Page()), uint64(filter.PageSize()))

	query = buildFilter(query, filter)

	_, err = query.LoadContext(ctx, &dbTransaction)
	if err != nil {
		return entity.Page[entity.Transaction]{}, fmt.Errorf(op, err)
	}

	return entity.NewPage(conversions.TransactionsFromDB(dbTransaction), filter.Filter, count), nil
}

func (r *transactionRepo) Create(ctx context.Context, transaction entity.TransactionCreate) (entity.Transaction, error) {
	op := "failed to create transaction: %w"

	tx, err := r.db.NewSession(nil).BeginTx(ctx, nil)
	if err != nil {
		return entity.Transaction{}, fmt.Errorf(op, err)
	}
	defer tx.RollbackUnlessCommitted()

	currentBalance, err := getCurrentBalance(ctx, tx, transaction.OrganizationID)
	if err != nil {
		return entity.Transaction{}, fmt.Errorf(op, err)
	}

	newBalance := getNewBalance(currentBalance, transaction.Amount, transaction.Operation)
	if newBalance < 0 {
		return entity.Transaction{}, app.ErrInsufficientFunds
	}

	dbTransaction, err := createTransaction(ctx, tx, conversions.TransactionCreateToDB(transaction))
	if err != nil {
		return entity.Transaction{}, fmt.Errorf(op, err)
	}

	err = changeOrganizationBalance(ctx, tx, transaction.OrganizationID, newBalance)
	if err != nil {
		return entity.Transaction{}, fmt.Errorf(op, err)
	}

	err = tx.Commit()
	if err != nil {
		return entity.Transaction{}, fmt.Errorf(op, err)
	}

	return conversions.TransactionFromDB(dbTransaction), nil
}

func createTransaction(ctx context.Context, tx *dbr.Tx, transaction dbmodels.TransactionCreate) (dbmodels.Transaction, error) {
	op := "failed to create transaction record: %w"

	_, err := tx.
		InsertInto(dbmodels.TransactionsTable).
		Columns(columns...).
		Record(transaction).
		ExecContext(ctx)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == dbmodels.PQErrAlreadyExists {
			return dbmodels.Transaction{}, app.ErrAlreadyExists
		}

		return dbmodels.Transaction{}, fmt.Errorf(op, err)
	}

	var dbTransaction dbmodels.Transaction
	err = tx.
		Select(selectColumns...).
		From(dbr.I(dbmodels.TransactionsTable).As("t")).
		LeftJoin(dbr.I(dbmodels.GroupsTable).As("g"), "t.group_id = g.id").
		LeftJoin(dbr.I(dbmodels.WashServersTable).As("ws"), "t.wash_server_id = ws.id").
		LeftJoin(dbr.I(dbmodels.UsersTable).As("u"), "t.user_id = u.id").
		Where("t.id = ?", transaction.ID).
		LoadOneContext(ctx, &dbTransaction)

	if err != nil {
		if errors.Is(err, dbr.ErrNotFound) {
			return dbmodels.Transaction{}, app.ErrNotFound
		}

		return dbmodels.Transaction{}, fmt.Errorf(op, err)
	}

	return dbTransaction, nil
}

func changeOrganizationBalance(ctx context.Context, tx *dbr.Tx, organizationID uuid.UUID, balance int64) error {
	op := "failed to change organization balance: %w"

	result, err := tx.
		Update(dbmodels.OrganizationsTable).
		Where(dbmodels.ByIDCondition, organizationID).
		Set("balance", balance).
		ExecContext(ctx)

	if err != nil {
		return fmt.Errorf(op, err)
	}

	count, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf(op, err)
	}
	if count == 0 {
		return app.ErrNotFound
	}

	return nil
}

func getCurrentBalance(ctx context.Context, tx *dbr.Tx, organizationID uuid.UUID) (int64, error) {
	op := "failed to get current balance: %w"

	var balance int64
	err := tx.
		Select("balance").
		From(dbmodels.OrganizationsTable).
		Where(dbmodels.ByIDCondition, organizationID).
		LoadOneContext(ctx, &balance)

	if err != nil {
		if errors.Is(err, dbr.ErrNotFound) {
			return 0, app.ErrNotFound
		}

		return 0, fmt.Errorf(op, err)
	}

	return balance, nil
}

func getNewBalance(balance int64, amount int64, operation entity.Operation) int64 {
	switch operation {
	case entity.DepositOperation:
		return balance + amount
	case entity.DebitOperation:
		return balance - amount
	default:
		panic("Unknown transaction operation: " + operation)
	}
}

func buildFilter(query *dbr.SelectStmt, filter entity.TransactionFilter) *dbr.SelectStmt {
	query.Where("t.organization_id = ?", filter.OrganizationID)

	if filter.Operation != nil {
		query.Where("t.operation = ?", filter.Operation)
	}
	if filter.Service != nil {
		query.Where("t.service = ?", filter.Service)
	}
	if filter.GroupID != nil {
		query.Where("t.group_id = ?", filter.GroupID)
	}
	if filter.WashServerID != nil {
		query.Where("t.wash_server_id = ?", filter.WashServerID)
	}

	return query
}
