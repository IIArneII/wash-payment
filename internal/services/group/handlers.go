package group

import (
	"context"
	"errors"
	"wash-payment/internal/app"
	"wash-payment/internal/app/entity"

	uuid "github.com/satori/go.uuid"
)

func (s *groupService) Get(ctx context.Context, groupID uuid.UUID) (entity.Group, error) {
	groupFromDB, err := s.groupRepo.Get(ctx, groupID)
	if err != nil {
		return entity.Group{}, err
	}
	if groupFromDB.Deleted {
		return entity.Group{}, app.ErrNotFound
	}

	return groupFromDB, nil
}

func (s *groupService) Upsert(ctx context.Context, group entity.Group) (entity.Group, error) {
	if group.ID == uuid.Nil {
		return entity.Group{}, app.ErrNotFound
	}
	dbGroup, err := s.groupRepo.Get(ctx, group.ID)

	if err != nil {
		if errors.Is(err, app.ErrNotFound) {
			newGroup, err := s.groupRepo.Create(ctx, group)
			if err != nil {
				return entity.Group{}, err
			}

			return newGroup, nil
		}
		return entity.Group{}, err

	} else {
		if dbGroup.Version >= group.Version {
			return entity.Group{}, app.ErrOldVersion
		}
		groupUpdate := groupToUpdate(group)
		updatedGroup, err := s.groupRepo.Update(ctx, group.ID, groupUpdate)
		if err != nil {
			return entity.Group{}, err
		}
		return updatedGroup, nil
	}
}

func groupToUpdate(gr entity.Group) entity.GroupUpdate {
	return entity.GroupUpdate{
		Name:        &gr.Name,
		Description: &gr.Description,
		Version:     &gr.Version,
		Deleted:     &gr.Deleted,
	}
}
