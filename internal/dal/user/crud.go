package user

import (
	"context"
	"errors"
	"fmt"
	"wash-payment/internal/dal/dbmodels"

	"github.com/gocraft/dbr/v2"
	"github.com/lib/pq"
)

func (r *userRepo) Get(ctx context.Context, userID string) (dbmodels.User, error) {
	op := "failed to get user by ID: %w"

	var dbUser dbmodels.User
	err := r.db.NewSession(nil).
		Select("*").
		From(dbmodels.UsersTable).
		Where("id = ?", userID).
		LoadOneContext(ctx, &dbUser)

	if err == nil {
		return dbUser, nil
	}

	if errors.Is(err, dbr.ErrNotFound) {
		return dbmodels.User{}, dbmodels.ErrNotFound
	}

	return dbmodels.User{}, fmt.Errorf(op, err)
}

func (r *userRepo) GetList(ctx context.Context, filter dbmodels.BaseFilter) ([]dbmodels.User, error) {
	op := "failed to get users list: %w"

	var dbUsers []dbmodels.User
	_, err := r.db.NewSession(nil).
		Select("*").
		From(dbmodels.UsersTable).
		Limit(uint64(filter.Limit)).
		Offset(uint64(filter.Offset)).
		LoadContext(ctx, &dbUsers)

	if err != nil {
		return nil, fmt.Errorf(op, err)
	}

	return dbUsers, nil
}

func (r *userRepo) Create(ctx context.Context, userCreation dbmodels.User) (dbmodels.User, error) {
	op := "failed to create user: %w"

	tx, err := r.db.NewSession(nil).BeginTx(ctx, nil)
	if err != nil {
		return dbmodels.User{}, fmt.Errorf(op, err)
	}
	defer tx.RollbackUnlessCommitted()

	var dbUser dbmodels.User
	err = tx.InsertInto(dbmodels.UsersTable).
		Columns("id", "email", "name", "role", "organization_id").
		Record(userCreation).
		Returning("*").
		LoadContext(ctx, &dbUser)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == dbmodels.PQErrAlreadyExists {
			err = fmt.Errorf(op, dbmodels.ErrAlreadyExists)
		}

		return dbmodels.User{}, fmt.Errorf(op, err)
	}

	if err = tx.Commit(); err != nil {
		return dbmodels.User{}, fmt.Errorf(op, err)
	}

	return dbUser, nil
}
