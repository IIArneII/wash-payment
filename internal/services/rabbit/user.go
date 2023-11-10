package rabbit

import (
	"context"
	"wash-payment/internal/app/conversions"
	et "wash-payment/internal/app/entity"
	"wash-payment/internal/transport/rabbit/entity"
)

func (s *rabbitService) UpsertUser(ctx context.Context, user entity.User) error {
	userCreate, err := conversions.UserFromRabbit(user)
	if err != nil {
		return err
	}

	_, err = s.services.UserService.Upsert(ctx, userCreate, "", et.UserUpdate{})
	if err != nil {
		return err
	}

	return nil
}
