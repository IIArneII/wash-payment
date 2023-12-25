package app

import (
	"context"
	"wash-payment/internal/app/entity"
	"wash-payment/internal/dal/dbmodels"

	uuid "github.com/satori/go.uuid"
)

type (
	GroupService interface {
		Get(ctx context.Context, groupID uuid.UUID) (entity.Group, error)
		Upsert(ctx context.Context, group entity.Group) (entity.Group, error)
		Delete(ctx context.Context, groupID uuid.UUID) error
	}

	GroupRepo interface {
		Get(ctx context.Context, groupID uuid.UUID) (dbmodels.Group, error)
		Create(ctx context.Context, group dbmodels.Group) (dbmodels.Group, error)
		Update(ctx context.Context, groupID uuid.UUID, groupUpdate dbmodels.GroupUpdate) error
		Delete(ctx context.Context, groupID uuid.UUID) error
	}
)
