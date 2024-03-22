package washserver

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

var columns = []string{"id", "title", "description", "group_id", "version", "deleted"}

func (r *washServerRepo) Get(ctx context.Context, id uuid.UUID) (entity.WashServer, error) {
	op := "failed to get wash server by ID: %w"

	var dbWashServer dbmodels.WashServer
	err := r.db.NewSession(nil).
		Select(columns...).
		From(dbmodels.WashServersTable).
		Where(dbmodels.ByIDCondition, id).
		LoadOneContext(ctx, &dbWashServer)

	if err != nil {
		if errors.Is(err, dbr.ErrNotFound) {
			err = app.ErrNotFound
		}
		return entity.WashServer{}, fmt.Errorf(op, err)
	}

	return conversions.WashServerFromDB(dbWashServer), nil
}

func (r *washServerRepo) Create(ctx context.Context, washServer entity.WashServer) (entity.WashServer, error) {
	op := "failed to create wash server: %w"

	dbWashServer := conversions.WashServerToDB(washServer)
	var dbCreatedWashServer dbmodels.WashServer
	err := r.db.NewSession(nil).
		InsertInto(dbmodels.WashServersTable).
		Columns(columns...).
		Record(dbWashServer).
		Returning(columns...).
		LoadContext(ctx, &dbCreatedWashServer)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == dbmodels.PQErrAlreadyExists {
			err = app.ErrAlreadyExists
		}
		return entity.WashServer{}, fmt.Errorf(op, err)
	}

	return conversions.WashServerFromDB(dbCreatedWashServer), nil
}

func (r *washServerRepo) Update(ctx context.Context, id uuid.UUID, washServerUpdate entity.WashServerUpdate) (entity.WashServer, error) {
	op := "failed to update wash server: %w"

	WashServerUpdateDB := conversions.WashServerUpdateToDB(washServerUpdate)

	query := r.db.NewSession(nil).
		Update(dbmodels.WashServersTable).
		Where(dbmodels.ByIDCondition, id)

	if WashServerUpdateDB.Title != nil {
		query.Set("title", WashServerUpdateDB.Title)
	}
	if WashServerUpdateDB.Description != nil {
		query.Set("description", WashServerUpdateDB.Description)
	}
	if WashServerUpdateDB.GroupID.Valid {
		query.Set("group_id", WashServerUpdateDB.GroupID.UUID)
	}
	if WashServerUpdateDB.Deleted != nil {
		query.Set("deleted", WashServerUpdateDB.Deleted)
	}
	if WashServerUpdateDB.Version != nil {
		query.Set("version", WashServerUpdateDB.Version)
		query.Where("version <= ?", WashServerUpdateDB.Version)
	}

	result, err := query.ExecContext(ctx)
	if err != nil {
		if errors.Is(err, dbr.ErrColumnNotSpecified) {
			err = app.ErrEmptyUpdate
		}
		return entity.WashServer{}, fmt.Errorf(op, err)
	}

	count, err := result.RowsAffected()
	if err != nil {
		return entity.WashServer{}, fmt.Errorf(op, err)
	}
	if count == 0 {
		return entity.WashServer{}, fmt.Errorf(op, app.ErrNotFound)
	}

	return r.Get(ctx, id)
}
