package rabbit

import (
	"context"
	"wash-payment/internal/app/entity"
	rabbitEntity "wash-payment/internal/transport/rabbit/entity"

	uuid "github.com/satori/go.uuid"
)

func (s *rabbitService) UpsertGroup(ctx context.Context, rabbitGroup rabbitEntity.Group) error {
	group, err := groupFromRabbit(rabbitGroup)
	if err != nil {
		return err
	}

	_, err = s.services.GroupService.Upsert(ctx, group)
	if err != nil {
		return err
	}

	return nil
}

func groupFromRabbit(gr rabbitEntity.Group) (entity.Group, error) {
	id, err := uuid.FromString(gr.ID)
	if err != nil {
		return entity.Group{}, err
	}

	orgId, err := uuid.FromString(gr.OrganizationID)
	if err != nil {
		return entity.Group{}, err
	}

	return entity.Group{
		ID:             id,
		OrganizationID: orgId,
		Name:           gr.Name,
		Description:    gr.Description,
		Version:        gr.Version,
		Deleted:        gr.Deleted,
	}, nil
}
