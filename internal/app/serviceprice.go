package app

import (
	"context"
	"wash-payment/internal/app/entity"

	uuid "github.com/satori/go.uuid"
)

type (
	ServicePriceRepo interface {
		Get(ctx context.Context, organizationID uuid.UUID, service entity.Service) (entity.ServicePrice, error)
		Create(ctx context.Context, servicePrice entity.ServicePrice) (entity.ServicePrice, error)
		Update(ctx context.Context, organizationID uuid.UUID, service entity.Service, price int64) (entity.ServicePrice, error)
	}
)
