package app

import (
	"context"
	"wash-payment/internal/app/entity"

	uuid "github.com/satori/go.uuid"
)

type (
	GroupService interface {
		Get(ctx context.Context, id uuid.UUID) (entity.Group, error)
		Upsert(ctx context.Context, group entity.Group) (entity.Group, error)
	}

	GroupRepo interface {
		Get(ctx context.Context, id uuid.UUID) (entity.Group, error)
		Create(ctx context.Context, group entity.Group) (entity.Group, error)
		Update(ctx context.Context, id uuid.UUID, groupUpdate entity.GroupUpdate) (entity.Group, error)
	}
)
