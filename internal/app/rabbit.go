package app

import (
	"context"
	"wash-payment/internal/transport/rabbit/entity"
)

type (
	RabbitService interface {
		UpsertOrganization(ctx context.Context, organization entity.Organization) error
		UpsertGroup(ctx context.Context, group entity.Group) error
		UpsertWashServer(ctx context.Context, group entity.WashServer) error
		UpsertUser(ctx context.Context, user entity.User) error
		Withdrawal(ctx context.Context, payment entity.Withdrawal) error
	}
)
