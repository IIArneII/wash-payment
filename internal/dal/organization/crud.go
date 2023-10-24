package organization

import (
	"context"
	"errors"
	"fmt"
	"wash-payment/internal/dal/dbmodels"

	"github.com/gocraft/dbr/v2"
	"github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
)

var columns = []string{"id", "name", "display_name", "description", "version", "balance", "deleted"}

func (r *organizationRepo) Get(ctx context.Context, organizationID uuid.UUID) (dbmodels.Organization, error) {
	op := "failed to get organization by ID: %w"

	var dbOrganization dbmodels.Organization
	err := r.db.NewSession(nil).
		Select(columns...).
		From(dbmodels.OrganizationsTable).
		Where(dbmodels.ByIDCondition, organizationID).
		LoadOneContext(ctx, &dbOrganization)

	if err == nil {
		return dbOrganization, nil
	}

	if errors.Is(err, dbr.ErrNotFound) {
		return dbmodels.Organization{}, dbmodels.ErrNotFound
	}

	return dbmodels.Organization{}, fmt.Errorf(op, err)
}

func (r *organizationRepo) Create(ctx context.Context, organization dbmodels.Organization) (dbmodels.Organization, error) {
	op := "failed to create organization: %w"

	var dbOrganization dbmodels.Organization
	err := r.db.NewSession(nil).
		InsertInto(dbmodels.OrganizationsTable).
		Columns(columns...).
		Record(organization).
		Returning(columns...).
		LoadContext(ctx, &dbOrganization)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == dbmodels.PQErrAlreadyExists {
			err = fmt.Errorf(op, dbmodels.ErrAlreadyExists)
		}

		return dbmodels.Organization{}, fmt.Errorf(op, err)
	}

	return dbOrganization, nil
}

func (r *organizationRepo) Update(ctx context.Context, organizationID uuid.UUID, organizationUpdate dbmodels.OrganizationUpdate) error {
	op := "failed to update organization: %w"

	query := r.db.NewSession(nil).
		Update(dbmodels.OrganizationsTable).
		Where(dbmodels.ByIDCondition, organizationID)

	if organizationUpdate.Name != nil {
		query.Set("name", organizationUpdate.Name)
	}
	if organizationUpdate.DisplayName != nil {
		query.Set("display_name", organizationUpdate.DisplayName)
	}
	if organizationUpdate.Description != nil {
		query.Set("description", organizationUpdate.Description)
	}
	if organizationUpdate.Version != nil {
		query.Set("version", organizationUpdate.Version)
		query.Where("version <= ?", organizationUpdate.Version)
	}

	result, err := query.ExecContext(ctx)
	if err != nil {
		if errors.Is(err, dbr.ErrColumnNotSpecified) {
			err = dbmodels.ErrEmptyUpdate
		}

		return fmt.Errorf(op, err)
	}

	count, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf(op, err)
	}
	if count == 0 {
		return dbmodels.ErrNotFound
	}

	return nil
}

func (r *organizationRepo) Delete(ctx context.Context, organizationID uuid.UUID) error {
	op := "failed to delete organization: %w"

	result, err := r.db.NewSession(nil).
		Update(dbmodels.OrganizationsTable).
		Where(dbmodels.ByIDCondition, organizationID).
		Set("deleted", true).
		ExecContext(ctx)

	if err != nil {
		return fmt.Errorf(op, err)
	}

	count, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf(op, err)
	}
	if count == 0 {
		return dbmodels.ErrNotFound
	}

	return nil
}
