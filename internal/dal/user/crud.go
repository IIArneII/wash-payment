package user

import (
	"context"
	"errors"
	"fmt"
	"wash-payment/internal/dal/dbmodels"

	"github.com/gocraft/dbr/v2"
	"github.com/lib/pq"
)

var columns = []string{"id", "email", "name", "role", "organization_id", "version"}

func (r *userRepo) Get(ctx context.Context, userID string) (dbmodels.User, error) {
	op := "failed to get user by ID: %w"

	var dbUser dbmodels.User
	err := r.db.NewSession(nil).
		Select(columns...).
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

func (r *userRepo) Create(ctx context.Context, user dbmodels.User) (dbmodels.User, error) {
	op := "failed to create user: %w"

	var dbUser dbmodels.User
	err := r.db.NewSession(nil).
		InsertInto(dbmodels.UsersTable).
		Columns(columns...).
		Record(user).
		Returning(columns...).
		LoadContext(ctx, &dbUser)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == dbmodels.PQErrAlreadyExists {
			err = fmt.Errorf(op, dbmodels.ErrAlreadyExists)
		}

		return dbmodels.User{}, fmt.Errorf(op, err)
	}

	return dbUser, nil
}

func (r *userRepo) Update(ctx context.Context, userID string, userUpdate dbmodels.UserUpdate) error {
	op := "failed to update user: %w"

	query := r.db.NewSession(nil).
		Update(dbmodels.UsersTable).
		Where("id = ?", userID)

	if userUpdate.Role != nil {
		query.Set("role", userUpdate.Role)
	}
	if userUpdate.Name != nil {
		query.Set("name", userUpdate.Name)
	}
	if userUpdate.Email != nil {
		query.Set("email", userUpdate.Email)
	}
	if userUpdate.Version != nil {
		query.Set("version", userUpdate.Version)
		query.Where("version <= ?", userUpdate.Version)
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
