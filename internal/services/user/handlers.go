package user

import (
	"context"
	"wash-payment/internal/app"
	"wash-payment/internal/entity"
)

func (s *userService) Get(ctx context.Context, auth app.Auth, userID string) (entity.User, error) {
	return entity.User{}, nil
}

func (s *userService) GetList(ctx context.Context, auth app.Auth, pagination entity.BaseFilter) ([]entity.User, error) {
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
