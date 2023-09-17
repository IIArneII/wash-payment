package app

import (
	"context"
	"wash-payment/internal/dal/dbmodels"
	"wash-payment/internal/entity"
)

type (
	OrganizationService interface {
		Get(ctx context.Context, userID string) (entity.User, error)
		GetList(ctx context.Context, pagination entity.BaseFilter) ([]entity.User, error)
		Create(ctx context.Context, userCreation entity.UserCreation) (entity.User, error)
		Update(ctx context.Context, userModel entity.UserUpdate) error
		Delete(ctx context.Context, userID string) error
	}

	OrganizationRepo interface {
		Get(ctx context.Context, userID string) (dbmodels.User, error)
		GetList(ctx context.Context, filter dbmodels.BaseFilter) ([]dbmodels.User, error)
		Create(ctx context.Context, userCreation dbmodels.UserCreation) (dbmodels.User, error)
		Update(ctx context.Context, userModel dbmodels.UserUpdate) error
		Delete(ctx context.Context, userID string) error
	}
)
