package conversions

import (
	"wash-payment/internal/app/entity"
	"wash-payment/internal/pkg/openapi/models"

	"github.com/go-openapi/strfmt"
)

func TransactionToRest(transaction entity.Transaction) models.Transaction {
	id := strfmt.UUID(transaction.ID.String())
	op := (string)(transaction.Operation)
	createAt := strfmt.DateTime(transaction.CreatedAt)
	return models.Transaction{
		ID:        &id,
		Operation: &op,
		Sevice:    transaction.Sevice,
		CreatedAt: &createAt,
		Amount:    &transaction.Amount,
	}
}

func TransactionsToRest(transactions entity.Page[entity.Transaction]) *models.TransactionPage {
	txs := []*models.Transaction{}
	for _, v := range transactions.Items {
		tx := TransactionToRest(v)
		txs = append(txs, &tx)
	}
	page := int64(transactions.Page)
	pageSize := int64(transactions.PageSize)
	totalPages := int64(transactions.TotalPages)
	totalItems := int64(transactions.TotalItems)
	return &models.TransactionPage{
		Items:      txs,
		Page:       &page,
		PageSize:   &pageSize,
		TotalPages: &totalPages,
		TotalItems: &totalItems,
	}
}
