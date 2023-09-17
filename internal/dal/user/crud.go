package user

import (
	"context"
	"wash-payment/internal/dal/dbmodels"
)

func (r *userRepo) Get(ctx context.Context, userID string) (dbmodels.User, error) {
	return dbmodels.User{}, nil
}

func (r *userRepo) GetList(ctx context.Context, filter dbmodels.BaseFilter) ([]dbmodels.User, error) {
	return make([]dbmodels.User, 1), nil
}

func (r *userRepo) Create(ctx context.Context, userCreation dbmodels.UserCreation) (dbmodels.User, error) {
	return dbmodels.User{}, nil
}

func (r *userRepo) Update(ctx context.Context, userModel dbmodels.UserUpdate) error {
	return nil
}

func (r *userRepo) Delete(ctx context.Context, userID string) error {
	return nil
}
