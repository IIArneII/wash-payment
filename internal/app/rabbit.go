package app

import (
	"context"
	"wash-payment/internal/transport/rabbit/entity"
)

type (
	RabbitService interface {
		UpsertOrganization(ctx context.Context, organization entity.Organization) error
		UpsertGroup(ctx context.Context, group entity.Group) error
		UpsertUser(ctx context.Context, user entity.User) error
		//NEW
		ProcessWithdrawal(ctx context.Context, organization entity.Organization, amount int64) error
	}
)
