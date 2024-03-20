package app

import (
	"context"
	"wash-payment/internal/app/entity"

	uuid "github.com/satori/go.uuid"
)

type (
	OrganizationService interface {
		List(ctx context.Context, auth entity.Auth, filter entity.OrganizationFilter) (entity.Page[entity.Organization], error)
		Get(ctx context.Context, auth entity.Auth, organizationID uuid.UUID) (entity.Organization, error)
		Upsert(ctx context.Context, organization entity.Organization) (entity.Organization, error)
		SetServicePrices(ctx context.Context, auth entity.Auth, organizationID uuid.UUID, servicePrices entity.ServicePrices) error
	}

	OrganizationRepo interface {
		List(ctx context.Context, filter entity.OrganizationFilter) (entity.Page[entity.Organization], error)
		Get(ctx context.Context, organizationID uuid.UUID) (entity.Organization, error)
		Create(ctx context.Context, organization entity.Organization) (entity.Organization, error)
		Update(ctx context.Context, organizationID uuid.UUID, organizationUpdate entity.OrganizationUpdate) (entity.Organization, error)
	}
)
