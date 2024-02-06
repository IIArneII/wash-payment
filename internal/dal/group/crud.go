package group

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

var columns = []string{"id", "organization_id", "name", "description", "version", "deleted"}

func (r *groupRepo) Get(ctx context.Context, id uuid.UUID) (entity.Group, error) {
	op := "failed to get group by ID: %w"

	var dbGroup dbmodels.Group
	err := r.db.NewSession(nil).
		Select(columns...).
		From(dbmodels.GroupTable).
		Where(dbmodels.ByIDCondition, id).
		LoadOneContext(ctx, &dbGroup)

	if err != nil {
		if errors.Is(err, dbr.ErrNotFound) {
			err = app.ErrNotFound
		}
		return entity.Group{}, fmt.Errorf(op, err)
	}

	return conversions.GroupFromDB(dbGroup), nil
}

func (r *groupRepo) Create(ctx context.Context, group entity.Group) (entity.Group, error) {
	op := "failed to create group: %w"

	dbGroup := conversions.GroupToDB(group)
	var dbCreatedGroup dbmodels.Group
	err := r.db.NewSession(nil).
		InsertInto(dbmodels.GroupTable).
		Columns(columns...).
		Record(dbGroup).
		Returning(columns...).
		LoadContext(ctx, &dbCreatedGroup)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == dbmodels.PQErrAlreadyExists {
			err = app.ErrAlreadyExists
		}
		return entity.Group{}, fmt.Errorf(op, err)
	}

	return conversions.GroupFromDB(dbCreatedGroup), nil
}

func (r *groupRepo) Update(ctx context.Context, id uuid.UUID, groupUpdate entity.GroupUpdate) (entity.Group, error) {
	op := "failed to update group: %w"

	groupUpdateDB := conversions.GroupUpdateToDB(groupUpdate)

	query := r.db.NewSession(nil).
		Update(dbmodels.GroupTable).
		Where(dbmodels.ByIDCondition, id)

	if groupUpdateDB.Name != nil {
		query.Set("name", groupUpdateDB.Name)
	}
	if groupUpdateDB.Description != nil {
		query.Set("description", groupUpdateDB.Description)
	}
	if groupUpdateDB.Deleted != nil {
		query.Set("deleted", groupUpdateDB.Deleted)
	}
	if groupUpdateDB.Version != nil {
		query.Set("version", groupUpdateDB.Version)
		query.Where("version <= ?", groupUpdateDB.Version)
	}

	result, err := query.ExecContext(ctx)
	if err != nil {
		if errors.Is(err, dbr.ErrColumnNotSpecified) {
			err = app.ErrEmptyUpdate
		}
		return entity.Group{}, fmt.Errorf(op, err)
	}

	count, err := result.RowsAffected()
	if err != nil {
		return entity.Group{}, fmt.Errorf(op, err)
	}
	if count == 0 {
		return entity.Group{}, fmt.Errorf(op, app.ErrNotFound)
	}

	return r.Get(ctx, id)
}
