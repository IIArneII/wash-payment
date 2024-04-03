package conversions

import (
	"wash-payment/internal/app/entity"
	"wash-payment/internal/pkg/openapi/models"
	"wash-payment/internal/pkg/openapi/restapi/operations/organizations"

	"github.com/go-openapi/strfmt"
	uuid "github.com/satori/go.uuid"
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

func operationFromRest(operation models.Operation) entity.Operation {
	switch operation {
	case models.OperationDeposit:
		return entity.DepositOperation
	case models.OperationDebit:
		return entity.DebitOperation
	default:
		panic("Unknown operation: " + operation)
	}
}

func serviceFromRest(service models.Service) entity.Service {
	switch service {
	case models.ServicePayment:
		return entity.PaymentService
	case models.ServiceBonus:
		return entity.BonusService
	case models.ServiceSbp:
		return entity.SbpService
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

func UserToRest(user entity.User) models.User {
	return models.User{
		ID:   &user.ID,
		Name: &user.Name,
	}
}

func TransactionToRest(transaction entity.Transaction) models.Transaction {
	id := strfmt.UUID(transaction.ID.String())
	organizationID := strfmt.UUID(transaction.OrganizationID.String())
	createAt := strfmt.DateTime(transaction.CreatedAt)
	service := serviceToRest(transaction.Service)

	var stationsCount *int64 = nil
	if transaction.StationsCount != nil {
		sc := int64(*transaction.StationsCount)
		stationsCount = &sc
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

	var user *models.User = nil
	if transaction.User != nil {
		u := UserToRest(*transaction.User)
		user = &u
	}

	return models.Transaction{
		ID:             &id,
		Operation:      operationToRest(transaction.Operation),
		OrganizationID: &organizationID,
		CreatedAt:      &createAt,
		ForDate:        (*strfmt.Date)(transaction.ForDate),
		Amount:         &transaction.Amount,
		Sevice:         &service,
		StationsCount:  stationsCount,
		Group:          group,
		User:           user,
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

func TransactionsFilterFromRest(params organizations.TransactionsParams) (entity.TransactionFilter, error) {
	organizationID, err := uuid.FromString(params.ID.String())
	if err != nil {
		return entity.TransactionFilter{}, err
	}

	var groupID *uuid.UUID
	if params.GroupID != nil {
		id, err := uuid.FromString(params.GroupID.String())
		if err != nil {
			return entity.TransactionFilter{}, err
		}
		groupID = &id
	}

	var washID *uuid.UUID
	if params.WashServerID != nil {
		id, err := uuid.FromString(params.WashServerID.String())
		if err != nil {
			return entity.TransactionFilter{}, err
		}
		washID = &id
	}

	var service *entity.Service
	if params.Service != nil {
		s := serviceFromRest(models.Service(*params.Service))
		service = &s
	}

	var operation *entity.Operation
	if params.Operation != nil {
		o := operationFromRest(models.Operation(*params.Operation))
		operation = &o
	}

	return entity.TransactionFilter{
		Filter:         entity.NewFilter(int(*params.Page), int(*params.PageSize)),
		OrganizationID: organizationID,
		GroupID:        groupID,
		WashServerID:   washID,
		Service:        service,
		Operation:      operation,
	}, nil
}
