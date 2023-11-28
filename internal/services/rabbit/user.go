package rabbit

import (
	"context"
	"wash-payment/internal/app/conversions"
	"wash-payment/internal/transport/rabbit/entity"
)

func (s *rabbitService) UpsertUser(ctx context.Context, rabbitUser entity.User) error {
	user, err := conversions.UserFromRabbit(rabbitUser)
	if err != nil {
		return err
	}

	_, err = s.services.UserService.Upsert(ctx, user)
	if err != nil {
		return err
	}

	return nil
}
