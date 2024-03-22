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

func (r *transactionRepo) Get(ctx context.Context, transactionID uuid.UUID) (entity.Transaction, error) {
	op := "failed to get transaction by ID: %w"

	var dbTransaction dbmodels.Transaction
	err := r.db.NewSession(nil).
		Select(columns...).
		From(dbmodels.TransactionsTable).
		Where(dbmodels.ByIDCondition, transactionID).
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
	err := r.db.NewSession(nil).
		Select(dbmodels.CountSelect).
		From(dbmodels.TransactionsTable).
		Where("organization_id = ?", filter.OrganizationID).
		LoadOneContext(ctx, &count)

	if err != nil {
		return entity.Page[entity.Transaction]{}, fmt.Errorf(op, err)
	}

	var dbTransaction []dbmodels.Transaction
	_, err = r.db.NewSession(nil).
		Select(columns...).
		From(dbmodels.TransactionsTable).
		Where("organization_id = ?", filter.OrganizationID).
		OrderDesc("created_at").
		Paginate(uint64(filter.Page), uint64(filter.PageSize)).
		LoadContext(ctx, &dbTransaction)

	if err != nil {
		return entity.Page[entity.Transaction]{}, fmt.Errorf(op, err)
	}

	return entity.NewPage(conversions.TransactionsFromDB(dbTransaction), filter.Filter, count), nil
}

func (r *transactionRepo) Create(ctx context.Context, transaction entity.Transaction) (entity.Transaction, error) {
	op := "failed to create transaction: %w"

	tx, err := r.db.NewSession(nil).BeginTx(ctx, nil)
	if err != nil {
		return entity.Transaction{}, fmt.Errorf(op, err)
	}
	defer tx.RollbackUnlessCommitted()

	currentBalance, err := getСurrentBalance(ctx, tx, transaction.OrganizationID)
	if err != nil {
		return entity.Transaction{}, err
	}

	newBalance := getNewBalance(currentBalance, transaction.Amount, transaction.Operation)
	if newBalance < 0 {
		return entity.Transaction{}, app.ErrInsufficientFunds
	}

	dbTransaction, err := createTransaction(ctx, tx, conversions.TransactionToDB(transaction))
	if err != nil {
		return entity.Transaction{}, err
	}

	err = changeOrganizationBalance(ctx, tx, transaction.OrganizationID, newBalance)
	if err != nil {
		return entity.Transaction{}, err
	}

	err = tx.Commit()
	if err != nil {
		return entity.Transaction{}, fmt.Errorf(op, err)
	}

	return conversions.TransactionFromDB(dbTransaction), nil
}

func createTransaction(ctx context.Context, tx *dbr.Tx, transaction dbmodels.Transaction) (dbmodels.Transaction, error) {
	op := "failed to create transaction: %w"

	var dbTransaction dbmodels.Transaction
	err := tx.
		InsertInto(dbmodels.TransactionsTable).
		Columns(columns...).
		Record(transaction).
		Returning(columns...).
		LoadContext(ctx, &dbTransaction)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == dbmodels.PQErrAlreadyExists {
			return dbmodels.Transaction{}, app.ErrAlreadyExists
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

func getСurrentBalance(ctx context.Context, tx *dbr.Tx, organizationID uuid.UUID) (int64, error) {
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
