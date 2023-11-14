package rabbit

import (
	"context"
	"wash-payment/internal/app/conversions"
	globalEntity "wash-payment/internal/app/entity"
	"wash-payment/internal/transport/rabbit/entity"

	uuid "github.com/satori/go.uuid"
)

func (s *rabbitService) UpsertGroup(ctx context.Context, group entity.Group) error {
	groupCreate, err := conversions.GroupFromRabbit(group)
	if err != nil {
		return err
	}

	_, err = s.services.GroupService.Upsert(ctx, groupCreate, uuid.Nil, globalEntity.GroupUpdate{})
	if err != nil {
		return err
	}

	return nil
}
