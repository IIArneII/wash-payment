package washserver

import (
	"context"
	"errors"
	"wash-payment/internal/app"
	"wash-payment/internal/app/entity"

	uuid "github.com/satori/go.uuid"
)

func (s *washServerService) Get(ctx context.Context, id uuid.UUID) (entity.WashServer, error) {
	washServerFromDB, err := s.washServerRepo.Get(ctx, id)
	if err != nil {
		return entity.WashServer{}, err
	}
	if washServerFromDB.Deleted {
		return entity.WashServer{}, app.ErrNotFound
	}

	return washServerFromDB, nil
}

func (s *washServerService) Upsert(ctx context.Context, washServer entity.WashServer) (entity.WashServer, error) {
	if washServer.ID == uuid.Nil {
		return entity.WashServer{}, app.ErrNotFound
	}
	dbWashServer, err := s.washServerRepo.Get(ctx, washServer.ID)

	if err != nil {
		if errors.Is(err, app.ErrNotFound) {
			newWashServer, err := s.washServerRepo.Create(ctx, washServer)
			if err != nil {
				return entity.WashServer{}, err
			}

			return newWashServer, nil
		}
		return entity.WashServer{}, err

	} else {
		if dbWashServer.Version >= washServer.Version {
			return entity.WashServer{}, app.ErrOldVersion
		}
		washServerUpdate := washServerToUpdate(washServer)
		updatedGroup, err := s.washServerRepo.Update(ctx, washServer.ID, washServerUpdate)
		if err != nil {
			return entity.WashServer{}, err
		}
		return updatedGroup, nil
	}
}

func washServerToUpdate(ws entity.WashServer) entity.WashServerUpdate {
	return entity.WashServerUpdate{
		Title:       &ws.Title,
		Description: &ws.Description,
		Version:     &ws.Version,
		Deleted:     &ws.Deleted,
		GroupID:     &ws.GroupID,
	}
}
