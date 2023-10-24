package rabbit

import (
	"context"
	"wash-payment/internal/app/conversions"
	"wash-payment/internal/transport/rabbit/entity"
)

func (s *rabbitService) UpsertUser(ctx context.Context, user entity.User) error {
	userCreate, err := conversions.UserFromRabbit(user)
	if err != nil {
		return err
	}

	_, err = s.services.UserService.Create(ctx, userCreate)
	if err != nil {
		return err
	}

	return nil
}
