package app

import (
	"context"
	"wash-payment/internal/app/entity"

	uuid "github.com/satori/go.uuid"
)

type (
	WashServerService interface {
		Get(ctx context.Context, id uuid.UUID) (entity.WashServer, error)
		Upsert(ctx context.Context, washServer entity.WashServer) (entity.WashServer, error)
	}

	WashServerRepo interface {
		Get(ctx context.Context, id uuid.UUID) (entity.WashServer, error)
		Create(ctx context.Context, washServer entity.WashServer) (entity.WashServer, error)
		Update(ctx context.Context, id uuid.UUID, washServerUpdate entity.WashServerUpdate) (entity.WashServer, error)
	}
)
