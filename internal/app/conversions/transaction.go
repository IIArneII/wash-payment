package conversions

import (
	"wash-payment/internal/app/entity"
	"wash-payment/internal/dal/dbmodels"
	"wash-payment/internal/pkg/openapi/models"

	"github.com/go-openapi/strfmt"
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

func TransactionToRest(transaction entity.Transaction) models.Transaction {
	id := strfmt.UUID(transaction.ID.String())
	op := (string)(transaction.Operation)
	createAt := strfmt.DateTime(transaction.CreatedAt)
	return models.Transaction{
		ID:        &id,
		Operation: &op,
		Sevice:    &transaction.Sevice,
		CreatedAt: &createAt,
		Amount:    &transaction.Amount,
	}
}

func TransactionsToRest(transactions entity.Page[entity.Transaction]) []*models.Transaction {
	txs := []*models.Transaction{}
	for _, v := range transactions.Items {
		tx := TransactionToRest(v)
		txs = append(txs, &tx)
	}
	return txs
}
