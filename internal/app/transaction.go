package app

import (
	"context"
	"wash-payment/internal/app/entity"

	uuid "github.com/satori/go.uuid"
)

type (
	TransactionService interface {
		List(ctx context.Context, auth entity.Auth, filter entity.TransactionFilter) (entity.Page[entity.Transaction], error)
		Deposit(ctx context.Context, auth entity.Auth, organizationID uuid.UUID, amount int64) error
		Withdrawal(ctx context.Context, withdrawal entity.Withdrawal) error
	}

	TransactionRepo interface {
		Get(ctx context.Context, transactionID uuid.UUID) (entity.Transaction, error)
		List(ctx context.Context, filter entity.TransactionFilter) (entity.Page[entity.Transaction], error)
		Create(ctx context.Context, transaction entity.TransactionCreate) (entity.Transaction, error)
	}
)
