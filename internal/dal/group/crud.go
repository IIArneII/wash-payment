package group

import (
	"context"
	"errors"
	"fmt"
	"wash-payment/internal/dal/dbmodels"

	"github.com/gocraft/dbr/v2"
	"github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
)

var columns = []string{"id", "organization_id", "name", "description", "version", "deleted"}

func (r *groupRepo) Get(ctx context.Context, groupID uuid.UUID) (dbmodels.Group, error) {
	op := "failed to get group by ID: %w"

	var dbGroup dbmodels.Group
	err := r.db.NewSession(nil).
		Select(columns...).
		From(dbmodels.GroupTable).
		Where(dbmodels.ByIDCondition, groupID).
		LoadOneContext(ctx, &dbGroup)

	if err == nil {
		return dbGroup, nil
	}

	if errors.Is(err, dbr.ErrNotFound) {
		return dbmodels.Group{}, dbmodels.ErrNotFound
	}

	return dbmodels.Group{}, fmt.Errorf(op, err)
}

func (r *groupRepo) Create(ctx context.Context, group dbmodels.Group) (dbmodels.Group, error) {
	op := "failed to create group: %w"

	var dbGroup dbmodels.Group
	err := r.db.NewSession(nil).
		InsertInto(dbmodels.GroupTable).
		Columns(columns...).
		Record(group).
		Returning(columns...).
		LoadContext(ctx, &dbGroup)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == dbmodels.PQErrAlreadyExists {
			err = fmt.Errorf(op, dbmodels.ErrAlreadyExists)
		}

		return dbmodels.Group{}, fmt.Errorf(op, err)
	}

	return dbGroup, nil
}

func (r *groupRepo) Update(ctx context.Context, groupID uuid.UUID, groupUpdate dbmodels.GroupUpdate) error {
	op := "failed to update group: %w"

	query := r.db.NewSession(nil).
		Update(dbmodels.GroupTable).
		Where(dbmodels.ByIDCondition, groupID)

	if groupUpdate.Name != nil {
		query.Set("name", groupUpdate.Name)
	}
	if groupUpdate.Description != nil {
		query.Set("description", groupUpdate.Description)
	}
	if groupUpdate.Version != nil {
		query.Set("version", groupUpdate.Version)
		query.Where("version <= ?", groupUpdate.Version)
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

func (r *groupRepo) Delete(ctx context.Context, groupID uuid.UUID) error {
	op := "failed to delete group: %w"

	result, err := r.db.NewSession(nil).
		Update(dbmodels.GroupTable).
		Where(dbmodels.ByIDCondition, groupID).
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
