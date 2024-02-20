package conversions

import (
	"wash-payment/internal/app/entity"
	"wash-payment/internal/dal/dbmodels"

	uuid "github.com/satori/go.uuid"
)

func TransactionOperationFromDb(operation dbmodels.Operation) entity.Operation {
	switch operation {
	case dbmodels.DepositOperation:
		return entity.DepositOperation
	case dbmodels.DebitOperation:
		return entity.DebitOperation
	default:
		panic("Unknown db operation: " + operation)
	}
}

func TransactionServiceFromDb(service *dbmodels.Service) *entity.Service {
	if service == nil {
		return nil
	}
	switch *service {
	case dbmodels.BonusService:
		o := entity.BonusService
		return &o
	case dbmodels.SbpService:
		o := entity.SbpService
		return &o
	default:
		panic("Unknown db service: " + *service)
	}
}

func TransactionOperationToDb(operation entity.Operation) dbmodels.Operation {
	switch operation {
	case entity.DepositOperation:
		return dbmodels.DepositOperation
	case entity.DebitOperation:
		return dbmodels.DebitOperation
	default:
		panic("Unknown app operation: " + operation)
	}
}

func TransactionServiceToDb(operation *entity.Service) *dbmodels.Service {
	if operation == nil {
		return nil
	}
	switch *operation {
	case entity.BonusService:
		o := dbmodels.BonusService
		return &o
	case entity.SbpService:
		o := dbmodels.SbpService
		return &o
	default:
		panic("Unknown app service: " + *operation)
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
		Service:        TransactionServiceFromDb(transaction.Service),
		Operation:      TransactionOperationFromDb(transaction.Operation),
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

	return dbmodels.Transaction{
		ID:             transaction.ID,
		OrganizationID: transaction.OrganizationID,
		GroupID:        groupID,
		Amount:         transaction.Amount,
		CreatedAt:      transaction.CreatedAt,
		Service:        TransactionServiceToDb(transaction.Service),
		Operation:      TransactionOperationToDb(transaction.Operation),
		StationsСount:  transaction.StationsСount,
		UserID:         transaction.UserID,
	}
}
