package conversions

import (
	"time"
	"wash-payment/internal/app/entity"
	"wash-payment/internal/dal/dbmodels"

	uuid "github.com/satori/go.uuid"
)

func OperationFromDb(operation dbmodels.Operation) entity.Operation {
	switch operation {
	case dbmodels.DepositOperation:
		return entity.DepositOperation
	case dbmodels.DebitOperation:
		return entity.DebitOperation
	default:
		panic("Unknown db operation: " + operation)
	}
}

func OperationToDb(operation entity.Operation) dbmodels.Operation {
	switch operation {
	case entity.DepositOperation:
		return dbmodels.DepositOperation
	case entity.DebitOperation:
		return dbmodels.DebitOperation
	default:
		panic("Unknown app operation: " + operation)
	}
}

func TransactionFromDB(transaction dbmodels.Transaction) entity.Transaction {
	var groupID *uuid.UUID
	if transaction.GroupID.Valid {
		groupID = &transaction.GroupID.UUID
	}

	return entity.Transaction{
		ID:             transaction.ID,
		OrganizationID: transaction.OrganizationID,
		GroupID:        groupID,
		Amount:         transaction.Amount,
		CreatedAt:      transaction.CreatedAt,
		ForDate:        transaction.ForDate,
		Service:        ServiceFromDb(transaction.Service),
		Operation:      OperationFromDb(transaction.Operation),
		StationsСount:  transaction.StationsСount,
		UserID:         transaction.UserID,
	}
}

func TransactionsFromDB(transactions []dbmodels.Transaction) []entity.Transaction {
	txs := []entity.Transaction{}
	for _, v := range transactions {
		txs = append(txs, TransactionFromDB(v))
	}
	return txs
}

func TransactionToDB(transaction entity.Transaction) dbmodels.Transaction {
	var groupID uuid.NullUUID
	if transaction.GroupID != nil {
		groupID.UUID = *transaction.GroupID
		groupID.Valid = true
	}
	var forDate *time.Time = nil
	if transaction.ForDate != nil {
		fd := transaction.ForDate.Truncate(24 * time.Hour)
		forDate = &fd
	}

	return dbmodels.Transaction{
		ID:             transaction.ID,
		OrganizationID: transaction.OrganizationID,
		GroupID:        groupID,
		Amount:         transaction.Amount,
		CreatedAt:      transaction.CreatedAt,
		ForDate:        forDate,
		Service:        ServiceToDb(transaction.Service),
		Operation:      OperationToDb(transaction.Operation),
		StationsСount:  transaction.StationsСount,
		UserID:         transaction.UserID,
	}
}
