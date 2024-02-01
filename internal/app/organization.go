package app

import (
	"context"
	"wash-payment/internal/app/entity"
	"wash-payment/internal/dal/dbmodels"

	uuid "github.com/satori/go.uuid"
)

type (
	OrganizationService interface {
		Get(ctx context.Context, auth entity.Auth, organizationID uuid.UUID) (entity.Organization, error)
		List(ctx context.Context, auth entity.Auth, filter entity.OrganizationFilter) (entity.Page[entity.Organization], error)
		Transactions(ctx context.Context, auth entity.Auth, filter entity.TransactionFilter) (entity.Page[entity.Transaction], error)
		Upsert(ctx context.Context, organization entity.OrganizationCreate) (entity.Organization, error)
		Delete(ctx context.Context, organizationID uuid.UUID) error
		Deposit(ctx context.Context, auth entity.Auth, organizationID uuid.UUID, amount int64) error
		Withdrawal(ctx context.Context, organizationID uuid.UUID, amount int64, service_name string) error
	}

	OrganizationRepo interface {
		Get(ctx context.Context, organizationID uuid.UUID) (dbmodels.Organization, error)
		List(ctx context.Context, filter entity.OrganizationFilter) (entity.Page[entity.Organization], error)
		Create(ctx context.Context, organization dbmodels.Organization) (dbmodels.Organization, error)
		Update(ctx context.Context, organizationID uuid.UUID, organizationUpdate dbmodels.OrganizationUpdate) error
		Delete(ctx context.Context, organizationID uuid.UUID) error
	}
)
