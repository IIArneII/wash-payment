package app

import (
	"context"
	"wash-payment/internal/app/entity"
	"wash-payment/internal/dal/dbmodels"

	uuid "github.com/satori/go.uuid"
)

type (
	TransactionService interface {
		Transactions(ctx context.Context, auth entity.Auth, filter entity.TransactionFilter) (entity.Page[entity.Transaction], error)
	}

	TransactionRepo interface {
		Get(ctx context.Context, transactionID uuid.UUID) (dbmodels.Transaction, error)
		List(ctx context.Context, filter entity.TransactionFilter) (entity.Page[entity.Transaction], error)
		Create(ctx context.Context, transaction dbmodels.Transaction) (dbmodels.Transaction, error)
	}
)
