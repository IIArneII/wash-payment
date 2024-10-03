package app

import (
	"context"
	"wash-payment/internal/app/entity"
)

type (
	UserService interface {
		Get(ctx context.Context, userID string) (entity.User, error)
		Upsert(ctx context.Context, user entity.User) (entity.User, error)
	}

	UserRepo interface {
		Get(ctx context.Context, userID string) (entity.User, error)
		Create(ctx context.Context, user entity.User) (entity.User, error)
		Update(ctx context.Context, userID string, userUpdate entity.UserUpdate) (entity.User, error)
	}
)
