package user

import (
	"context"
	"errors"
	"wash-payment/internal/app"
	"wash-payment/internal/app/conversions"
	"wash-payment/internal/app/entity"
	"wash-payment/internal/dal/dbmodels"
)

func (s *userService) Get(ctx context.Context, userID string) (entity.User, error) {
	userFromDB, err := s.userRepo.Get(ctx, userID)
	if err != nil {
		if errors.Is(err, dbmodels.ErrNotFound) {
			err = app.ErrNotFound
		}

		return entity.User{}, err
	}

	return conversions.UserFromDB(userFromDB), nil
}

func (s *userService) Upsert(ctx context.Context, user entity.User, userID string, userUpdate entity.UserUpdate) (entity.User, error) {
	if user.ID != "" {
		dbUserUpdate := conversions.UserUpdateToDB(userUpdate)

		userFromDB, err := s.userRepo.Get(ctx, userID)
		if err != nil {
			if errors.Is(err, dbmodels.ErrNotFound) {
				err = app.ErrNotFound
			}

			return entity.User{}, err
		}

		if userFromDB.Version < *dbUserUpdate.Version {
			err = s.userRepo.Update(ctx, userID, dbUserUpdate)
			if err != nil {
				if errors.Is(err, dbmodels.ErrNotFound) {
					err = app.ErrNotFound
				} else if errors.Is(err, dbmodels.ErrEmptyUpdate) {
					err = app.ErrBadRequest
				}

				return entity.User{}, err
			}
		}
		return entity.User{}, nil
	} else {
		dbUser := conversions.UserToDB(user)

		newUser, err := s.userRepo.Create(ctx, dbUser)
		if err != nil {
			if errors.Is(err, dbmodels.ErrAlreadyExists) {
				err = app.ErrAlreadyExists
			}

			return entity.User{}, err
		}

		return conversions.UserFromDB(newUser), nil
	}
}
