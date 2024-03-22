package conversions

import (
	"wash-payment/internal/app/entity"
	"wash-payment/internal/pkg/openapi/models"

	"github.com/go-openapi/strfmt"
)

func operationToRest(operation entity.Operation) *models.Operation {
	switch operation {
	case entity.DepositOperation:
		o := models.OperationDeposit
		return &o
	case entity.DebitOperation:
		o := models.OperationDebit
		return &o
	default:
		panic("Unknown operation: " + operation)
	}
}

func serviceToRest(service entity.Service) models.Service {
	switch service {
	case entity.PaymentService:
		return models.ServicePayment
	case entity.BonusService:
		return models.ServiceBonus
	case entity.SbpService:
		return models.ServiceSbp
	default:
		panic("Unknown service: " + service)
	}
}

func GroupToRest(group entity.Group) models.Group {
	id := strfmt.UUID(group.ID.String())

	return models.Group{
		ID:      &id,
		Name:    &group.Name,
		Deleted: &group.Deleted,
	}
}

func WashServerToRest(group entity.WashServer) models.WashServer {
	id := strfmt.UUID(group.ID.String())

	return models.WashServer{
		ID:      &id,
		Title:   &group.Title,
		Deleted: &group.Deleted,
	}
}

func TransactionToRest(transaction entity.Transaction) models.Transaction {
	id := strfmt.UUID(transaction.ID.String())
	organizationID := strfmt.UUID(transaction.OrganizationID.String())
	createAt := strfmt.DateTime(transaction.CreatedAt)
	service := serviceToRest(transaction.Service)

	var stationsСount *int64 = nil
	if transaction.StationsСount != nil {
		sc := int64(*transaction.StationsСount)
		stationsСount = &sc
	}

	var group *models.Group = nil
	if transaction.Group != nil {
		g := GroupToRest(*transaction.Group)
		group = &g
	}

	var washServer *models.WashServer = nil
	if transaction.WashServer != nil {
		w := WashServerToRest(*transaction.WashServer)
		washServer = &w
	}

	return models.Transaction{
		ID:             &id,
		Operation:      operationToRest(transaction.Operation),
		OrganizationID: &organizationID,
		CreatedAt:      &createAt,
		ForDate:        (*strfmt.Date)(transaction.ForDate),
		Amount:         &transaction.Amount,
		Sevice:         &service,
		UserID:         transaction.UserID,
		StationsСount:  stationsСount,
		Group:          group,
		WashServer:     washServer,
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
