package app

import (
	"context"
	"wash-payment/internal/app/entity"
	"wash-payment/internal/dal/dbmodels"

	uuid "github.com/satori/go.uuid"
)

type (
	OrganizationService interface {
		Get(ctx context.Context, auth Auth, organizationID uuid.UUID) (entity.Organization, error)
		Upsert(ctx context.Context, organizationID uuid.UUID, organizationCreate entity.OrganizationCreate, organizationUpdate entity.OrganizationUpdate) (entity.Organization, error)
		Delete(ctx context.Context, organizationID uuid.UUID) error
		Deposit(ctx context.Context, auth Auth, organizationID uuid.UUID, amount int64) error
		Withdrawal(ctx context.Context, organizationID uuid.UUID, amount int64) error
	}

	OrganizationRepo interface {
		Get(ctx context.Context, organizationID uuid.UUID) (dbmodels.Organization, error)
		Create(ctx context.Context, organization dbmodels.Organization) (dbmodels.Organization, error)
		Update(ctx context.Context, organizationID uuid.UUID, organizationUpdate dbmodels.OrganizationUpdate) error
		Delete(ctx context.Context, organizationID uuid.UUID) error
	}
)
