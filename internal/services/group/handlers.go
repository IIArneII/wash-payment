package group

import (
	"context"
	"errors"
	"wash-payment/internal/app"
	"wash-payment/internal/app/conversions"
	"wash-payment/internal/app/entity"
	"wash-payment/internal/dal/dbmodels"

	uuid "github.com/satori/go.uuid"
)

func (s *groupService) Get(ctx context.Context, groupID uuid.UUID) (entity.Group, error) {
	groupFromDB, err := s.groupRepo.Get(ctx, groupID)
	if err != nil {
		if errors.Is(err, dbmodels.ErrNotFound) {
			err = app.ErrNotFound
		}

		return entity.Group{}, err
	}

	return conversions.GroupFromDB(groupFromDB), nil
}

func (s *groupService) Upsert(ctx context.Context, group entity.Group, groupID uuid.UUID, groupUpdate entity.GroupUpdate) (entity.Group, error) {
	if groupID != uuid.Nil {
		dbGroupUpdate := conversions.GroupUpdateToDB(groupUpdate)

		groupFromDB, err := s.groupRepo.Get(ctx, groupID)
		if err != nil {
			if errors.Is(err, dbmodels.ErrNotFound) {
				err = app.ErrNotFound
			}

			return entity.Group{}, err
		}

		if groupFromDB.Version < *dbGroupUpdate.Version {
			err = s.groupRepo.Update(ctx, groupID, dbGroupUpdate)
			if err != nil {
				if errors.Is(err, dbmodels.ErrNotFound) {
					err = app.ErrNotFound
				} else if errors.Is(err, dbmodels.ErrEmptyUpdate) {
					err = app.ErrBadRequest
				}
				return entity.Group{}, err
			}
		}

		return entity.Group{}, nil
	} else {
		dbGroup := conversions.GroupToDB(group)

		newOrganization, err := s.groupRepo.Create(ctx, dbGroup)
		if err != nil {
			if errors.Is(err, dbmodels.ErrAlreadyExists) {
				err = app.ErrAlreadyExists
			}

			return entity.Group{}, err
		}

		return conversions.GroupFromDB(newOrganization), nil
	}
}

func (s *groupService) Delete(ctx context.Context, groupID uuid.UUID) error {
	err := s.groupRepo.Delete(ctx, groupID)
	if err != nil {
		if errors.Is(err, dbmodels.ErrNotFound) {
			err = app.ErrNotFound
		}

		return err
	}

	return nil
}
