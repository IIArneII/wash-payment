package rabbit

import (
	"context"
	"wash-payment/internal/app/conversions"
	"wash-payment/internal/transport/rabbit/entity"
)

func (s *rabbitService) UpsertGroup(ctx context.Context, group entity.Group) error {
	groupCreate, err := conversions.GroupFromRabbit(group)
	if err != nil {
		return err
	}

	_, err = s.services.GroupService.Create(ctx, groupCreate)
	if err != nil {
		return err
	}

	return nil
}
