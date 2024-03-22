package rabbit

import (
	"context"
	"wash-payment/internal/app/entity"
	rabbitEntity "wash-payment/internal/transport/rabbit/entity"

	uuid "github.com/satori/go.uuid"
)

func (s *rabbitService) UpsertWashServer(ctx context.Context, rabbitWashServer rabbitEntity.WashServer) error {
	washServer, err := washServerFromRabbit(rabbitWashServer)
	if err != nil {
		return err
	}

	_, err = s.services.WashServerService.Upsert(ctx, washServer)
	if err != nil {
		return err
	}

	return nil
}

func washServerFromRabbit(gr rabbitEntity.WashServer) (entity.WashServer, error) {
	id, err := uuid.FromString(gr.ID)
	if err != nil {
		return entity.WashServer{}, err
	}

	gId, err := uuid.FromString(gr.GroupID)
	if err != nil {
		return entity.WashServer{}, err
	}

	return entity.WashServer{
		ID:          id,
		Title:       gr.Title,
		GroupID:     gId,
		Description: gr.Description,
		Version:     gr.Version,
		Deleted:     gr.Deleted,
	}, nil
}
