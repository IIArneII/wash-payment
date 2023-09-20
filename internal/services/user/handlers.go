package user

import (
	"context"
	"errors"
	"wash-payment/internal/app"
	"wash-payment/internal/conversions"
	"wash-payment/internal/dal/dbmodels"
	"wash-payment/internal/entity"
)

func (s *userService) Get(ctx context.Context, auth app.Auth, userID string) (entity.User, error) {
	return entity.User{}, nil
}

func (s *userService) GetAuth(ctx context.Context, userID string) (entity.User, error) {
	userFromDB, err := s.userRepo.Get(ctx, userID)
	if err != nil {
		if errors.Is(err, dbmodels.ErrNotFound) {
			err = app.ErrNotFound
		}
		return entity.User{}, err
	}

	return conversions.UserFromDb(userFromDB), nil
}

func (s *userService) GetList(ctx context.Context, auth app.Auth, filter entity.BaseFilter) ([]entity.User, error) {
	return make([]entity.User, 0), nil
}

func (s *userService) Create(ctx context.Context, auth app.Auth, userCreation entity.UserCreation) (entity.User, error) {
	return entity.User{}, nil
}

func (s *userService) Update(ctx context.Context, auth app.Auth, userModel entity.UserUpdate) error {
	return nil
}

func (s *userService) Delete(ctx context.Context, auth app.Auth, userID string) error {
	return nil
}
