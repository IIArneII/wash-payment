package app

import (
	"context"
	"wash-payment/internal/app/entity"
	"wash-payment/internal/dal/dbmodels"
)

type (
	UserService interface {
		Get(ctx context.Context, userID string) (entity.User, error)
		Create(ctx context.Context, user entity.User) (entity.User, error)
		Update(ctx context.Context, userID string, userUpdate entity.UserUpdate) error
	}

	UserRepo interface {
		Get(ctx context.Context, userID string) (dbmodels.User, error)
		Create(ctx context.Context, user dbmodels.User) (dbmodels.User, error)
		Update(ctx context.Context, userID string, userUpdate dbmodels.UserUpdate) error
	}
)
