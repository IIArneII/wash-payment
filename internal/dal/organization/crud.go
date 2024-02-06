package organization

import (
	"context"
	"errors"
	"fmt"
	"wash-payment/internal/app"
	"wash-payment/internal/app/entity"
	"wash-payment/internal/dal/conversions"
	"wash-payment/internal/dal/dbmodels"

	"github.com/gocraft/dbr/v2"
	"github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
)

var columns = []string{"id", "name", "display_name", "description", "version", "balance", "deleted"}

func (r *organizationRepo) Get(ctx context.Context, organizationID uuid.UUID) (entity.Organization, error) {
	op := "failed to get organization by ID: %w"

	var dbOrganization dbmodels.Organization
	err := r.db.NewSession(nil).
		Select(columns...).
		From(dbmodels.OrganizationsTable).
		Where(dbmodels.ByIDCondition, organizationID).
		LoadOneContext(ctx, &dbOrganization)

	if err != nil {
		if errors.Is(err, dbr.ErrNotFound) {
			err = app.ErrNotFound
		}
		return entity.Organization{}, fmt.Errorf(op, err)
	}

	return conversions.OrganizationFromDB(dbOrganization), nil
}

func (r *organizationRepo) List(ctx context.Context, filter entity.OrganizationFilter) (entity.Page[entity.Organization], error) {
	op := "failed to get organizations list: %w"

	var count int
	err := r.db.NewSession(nil).
		Select(dbmodels.CountSelect).
		From(dbmodels.OrganizationsTable).
		LoadOneContext(ctx, &count)

	if err != nil {
		return entity.Page[entity.Organization]{}, fmt.Errorf(op, err)
	}

	var dbOrganizations []dbmodels.Organization
	_, err = r.db.NewSession(nil).
		Select(columns...).
		From(dbmodels.OrganizationsTable).
		OrderAsc("name").
		Offset(filter.Offset()).
		Limit(filter.Limit()).
		LoadContext(ctx, &dbOrganizations)

	if err != nil {
		return entity.Page[entity.Organization]{}, fmt.Errorf(op, err)
	}

	return entity.NewPage(conversions.OrganizationsFromDB(dbOrganizations), filter.Filter, count), nil
}

func (r *organizationRepo) Create(ctx context.Context, organization entity.Organization) (entity.Organization, error) {
	op := "failed to create organization: %w"

	dbOrganization := conversions.OrganizationToDB(organization)
	var dbCreatedOrganization dbmodels.Organization
	err := r.db.NewSession(nil).
		InsertInto(dbmodels.OrganizationsTable).
		Columns(columns...).
		Record(dbOrganization).
		Returning(columns...).
		LoadContext(ctx, &dbCreatedOrganization)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == dbmodels.PQErrAlreadyExists {
			err = app.ErrAlreadyExists
		}
		return entity.Organization{}, fmt.Errorf(op, err)
	}

	return conversions.OrganizationFromDB(dbCreatedOrganization), nil
}

func (r *organizationRepo) Update(ctx context.Context, organizationID uuid.UUID, organizationUpdate entity.OrganizationUpdate) (entity.Organization, error) {
	op := "failed to update organization: %w"

	organizationUpdateDB := conversions.OrganizationUpdateToDB(organizationUpdate)

	query := r.db.NewSession(nil).
		Update(dbmodels.OrganizationsTable).
		Where(dbmodels.ByIDCondition, organizationID)

	if organizationUpdateDB.Name != nil {
		query.Set("name", organizationUpdateDB.Name)
	}
	if organizationUpdateDB.DisplayName != nil {
		query.Set("display_name", organizationUpdateDB.DisplayName)
	}
	if organizationUpdateDB.Description != nil {
		query.Set("description", organizationUpdateDB.Description)
	}
	if organizationUpdateDB.Deleted != nil {
		query.Set("deleted", organizationUpdateDB.Deleted)
	}
	if organizationUpdateDB.Version != nil {
		query.Set("version", organizationUpdateDB.Version)
		query.Where("version <= ?", organizationUpdateDB.Version)
	}

	result, err := query.ExecContext(ctx)
	if err != nil {
		if errors.Is(err, dbr.ErrColumnNotSpecified) {
			err = app.ErrEmptyUpdate
		}

		return entity.Organization{}, fmt.Errorf(op, err)
	}

	count, err := result.RowsAffected()
	if err != nil {
		return entity.Organization{}, fmt.Errorf(op, err)
	}
	if count == 0 {
		return entity.Organization{}, fmt.Errorf(op, app.ErrNotFound)
	}

	return r.Get(ctx, organizationID)
}
