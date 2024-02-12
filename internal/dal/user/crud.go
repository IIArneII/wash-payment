package user

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
)

var columns = []string{"id", "email", "name", "role", "organization_id", "version"}

func (r *userRepo) Get(ctx context.Context, userID string) (entity.User, error) {
	op := "failed to get user by ID: %w"

	var dbUser dbmodels.User
	err := r.db.NewSession(nil).
		Select(columns...).
		From(dbmodels.UsersTable).
		Where(dbmodels.ByIDCondition, userID).
		LoadOneContext(ctx, &dbUser)

	if err != nil {
		if errors.Is(err, dbr.ErrNotFound) {
			err = app.ErrNotFound
		}
		return entity.User{}, fmt.Errorf(op, err)
	}

	return conversions.UserFromDB(dbUser), nil
}

func (r *userRepo) Create(ctx context.Context, user entity.User) (entity.User, error) {
	op := "failed to create user: %w"

	dbUser := conversions.UserToDB(user)
	var dbCreatedUser dbmodels.User
	err := r.db.NewSession(nil).
		InsertInto(dbmodels.UsersTable).
		Columns(columns...).
		Record(dbUser).
		Returning(columns...).
		LoadContext(ctx, &dbCreatedUser)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == dbmodels.PQErrAlreadyExists {
			err = app.ErrAlreadyExists
		}
		return entity.User{}, fmt.Errorf(op, err)
	}

	return conversions.UserFromDB(dbCreatedUser), nil
}

func (r *userRepo) Update(ctx context.Context, userID string, userUpdate entity.UserUpdate) (entity.User, error) {
	op := "failed to update user: %w"

	userUpdateDB := conversions.UserUpdateToDB(userUpdate)

	query := r.db.NewSession(nil).
		Update(dbmodels.UsersTable).
		Where(dbmodels.ByIDCondition, userID)

	if userUpdateDB.Role != nil {
		query.Set("role", userUpdateDB.Role)
	}
	if userUpdateDB.Name != nil {
		query.Set("name", userUpdateDB.Name)
	}
	if userUpdateDB.Email != nil {
		query.Set("email", userUpdateDB.Email)
	}
	if userUpdateDB.Version != nil {
		query.Set("version", userUpdateDB.Version)
		query.Where("version <= ?", userUpdateDB.Version)
	}

	result, err := query.ExecContext(ctx)
	if err != nil {
		if errors.Is(err, dbr.ErrColumnNotSpecified) {
			err = app.ErrEmptyUpdate
		}

		return entity.User{}, fmt.Errorf(op, err)
	}

	count, err := result.RowsAffected()
	if err != nil {
		return entity.User{}, fmt.Errorf(op, err)
	}
	if count == 0 {
		return entity.User{}, fmt.Errorf(op, app.ErrNotFound)
	}

	return r.Get(ctx, userID)
}
