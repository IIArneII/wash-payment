package transaction

import (
	"context"
	"errors"
	"fmt"
	"wash-payment/internal/dal/dbmodels"

	"github.com/gocraft/dbr/v2"
	"github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
)

var columns = []string{"id", "organization_id", "amount", "operation", "created_at", "sevice"}

func (r *transactionRepo) Get(ctx context.Context, transactionID uuid.UUID) (dbmodels.Transaction, error) {
	op := "failed to get transaction by ID: %w"

	var dbTransaction dbmodels.Transaction
	err := r.db.NewSession(nil).
		Select(columns...).
		From(dbmodels.TransactionTable).
		Where("id = ?", transactionID).
		LoadOneContext(ctx, &dbTransaction)

	if err != nil {
		if errors.Is(err, dbr.ErrNotFound) {
			err = dbmodels.ErrNotFound
		}

		return dbmodels.Transaction{}, fmt.Errorf(op, err)
	}

	return dbTransaction, nil
}

func (r *transactionRepo) Create(ctx context.Context, transaction dbmodels.Transaction) (dbmodels.Transaction, error) {
	op := "failed to create transaction: %w"

	tx, err := r.db.NewSession(nil).BeginTx(ctx, nil)
	if err != nil {
		return dbmodels.Transaction{}, fmt.Errorf(op, err)
	}
	defer tx.RollbackUnlessCommitted()

	currentBalance, err := getСurrentBalance(ctx, tx, transaction.OrganizationID)
	if err != nil {
		return dbmodels.Transaction{}, err
	}

	newBalance := getNewBalance(currentBalance, transaction.Amount, transaction.Operation)
	if newBalance < 0 {
		return dbmodels.Transaction{}, dbmodels.ErrInsufficientFunds
	}

	dbTransaction, err := createTransaction(ctx, tx, transaction)
	if err != nil {
		return dbmodels.Transaction{}, err
	}

	err = changeOrganizationBalance(ctx, tx, transaction.OrganizationID, newBalance)
	if err != nil {
		return dbmodels.Transaction{}, err
	}

	err = tx.Commit()
	if err != nil {
		return dbmodels.Transaction{}, fmt.Errorf(op, err)
	}

	return dbTransaction, nil
}

func createTransaction(ctx context.Context, tx *dbr.Tx, transaction dbmodels.Transaction) (dbmodels.Transaction, error) {
	op := "failed to create transaction: %w"

	var dbTransaction dbmodels.Transaction
	err := tx.
		InsertInto(dbmodels.TransactionTable).
		Columns(columns...).
		Record(transaction).
		Returning(columns...).
		LoadContext(ctx, &dbTransaction)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == dbmodels.PQErrAlreadyExists {
			return dbmodels.Transaction{}, dbmodels.ErrAlreadyExists
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
		return dbmodels.ErrNotFound
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
			return 0, dbmodels.ErrNotFound
		}

		return 0, fmt.Errorf(op, err)
	}

	return balance, nil
}

func getNewBalance(balance int64, amount int64, operation dbmodels.Operation) int64 {
	switch operation {
	case dbmodels.DepositOperation:
		return balance + amount
	case dbmodels.DebitOperation:
		return balance - amount
	default:
		panic("Unknown transaction operation: " + operation)
	}
}
