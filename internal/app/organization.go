package app

import (
	"context"
	"wash-payment/internal/dal/dbmodels"
	"wash-payment/internal/entity"
)

type (
	OrganizationService interface {
		Get(ctx context.Context, organizationID string) (entity.Organization, error)
		GetList(ctx context.Context, filter entity.OrganizationFilter) ([]entity.User, error)
		Create(ctx context.Context, organizationCreation entity.OrganizationCreation) (entity.User, error)
		Update(ctx context.Context, organizationUpdate entity.OrganizationUpdate) error
		Delete(ctx context.Context, organizationID string) error
	}

	OrganizationRepo interface {
		Get(ctx context.Context, organizationID string) (dbmodels.Organization, error)
		GetList(ctx context.Context, filter dbmodels.OrganizationFilter) ([]dbmodels.Organization, error)
		Create(ctx context.Context, organizationCreation dbmodels.OrganizationCreation) (dbmodels.Organization, error)
		Update(ctx context.Context, organizationUpdate dbmodels.OrganizationUpdate) error
		Delete(ctx context.Context, organizationID string) error
	}
)
