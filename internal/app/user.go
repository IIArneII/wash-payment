package app

import (
	"context"
	"wash-payment/internal/dal/dbmodels"
	"wash-payment/internal/entity"
)

type (
	UserService interface {
		Get(ctx context.Context, auth Auth, userID string) (entity.User, error)
		GetList(ctx context.Context, auth Auth, pagination entity.BaseFilter) ([]entity.User, error)
		Create(ctx context.Context, auth Auth, userCreation entity.UserCreation) (entity.User, error)
		Update(ctx context.Context, auth Auth, userModel entity.UserUpdate) error
		Delete(ctx context.Context, auth Auth, userID string) error
	}

	UserRepo interface {
		Get(ctx context.Context, userID string) (dbmodels.User, error)
		GetList(ctx context.Context, filter dbmodels.BaseFilter) ([]dbmodels.User, error)
		Create(ctx context.Context, userCreation dbmodels.UserCreation) (dbmodels.User, error)
		Update(ctx context.Context, userModel dbmodels.UserUpdate) error
		Delete(ctx context.Context, userID string) error
	}
)
