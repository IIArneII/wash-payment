package rabbit

import (
	"context"
	"wash-payment/internal/app/conversions"
	"wash-payment/internal/transport/rabbit/entity"
)

func (s *rabbitService) UpsertGroup(ctx context.Context, rabbitGroup entity.Group) error {

	group, err := conversions.GroupFromRabbit(rabbitGroup)
	if err != nil {
		return err
	}

	_, err = s.services.GroupService.Upsert(ctx, group)
	if err != nil {
		return err
	}

	return nil
}
