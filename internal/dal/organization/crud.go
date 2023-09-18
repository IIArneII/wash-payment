package organization

import (
	"context"
	"wash-payment/internal/dal/dbmodels"
)

func (r *organizationRepo) Get(ctx context.Context, organizationID string) (dbmodels.Organization, error) {
	return dbmodels.Organization{}, nil
}

func (r *organizationRepo) GetList(ctx context.Context, filter dbmodels.OrganizationFilter) ([]dbmodels.Organization, error) {
	return make([]dbmodels.Organization, 1), nil
}

func (r *organizationRepo) Create(ctx context.Context, userCreation dbmodels.OrganizationCreation) (dbmodels.Organization, error) {
	return dbmodels.Organization{}, nil
}

func (r *organizationRepo) Update(ctx context.Context, userModel dbmodels.OrganizationUpdate) error {
	return nil
}

func (r *organizationRepo) Delete(ctx context.Context, organizationID string) error {
	return nil
}
