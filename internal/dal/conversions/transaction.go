package conversions

import (
	"wash-payment/internal/app/entity"
	"wash-payment/internal/dal/dbmodels"
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

func TransactionFromDB(transaction dbmodels.Transaction) entity.Transaction {
	return entity.Transaction{
		ID:             transaction.ID,
		OrganizationID: transaction.OrganizationID,
		Amount:         transaction.Amount,
		CreatedAt:      transaction.CreatedAt,
		Sevice:         transaction.Sevice,
		Operation:      TransactionOperationFromDb(transaction.Operation),
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
	return dbmodels.Transaction{
		ID:             transaction.ID,
		OrganizationID: transaction.OrganizationID,
		Amount:         transaction.Amount,
		CreatedAt:      transaction.CreatedAt,
		Sevice:         transaction.Sevice,
		Operation:      TransactionOperationToDb(transaction.Operation),
	}
}
