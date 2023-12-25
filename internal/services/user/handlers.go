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

func (s *userService) Upsert(ctx context.Context, user entity.User) (entity.User, error) {
	_, err := s.userRepo.Get(ctx, user.ID)

	if err != nil {
		if errors.Is(err, dbmodels.ErrNotFound) {

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
		return entity.User{}, err
	} else {

		userUpdate := conversions.UserToUpdateUser(user)
		dbUserUpdate := conversions.UserUpdateToDB(userUpdate)

		err := s.userRepo.Update(ctx, user.ID, dbUserUpdate)
		if err != nil {
			if errors.Is(err, dbmodels.ErrNotFound) {
				err = app.ErrNotFound
			} else if errors.Is(err, dbmodels.ErrEmptyUpdate) {
				err = app.ErrBadRequest
			}

			return entity.User{}, err
		}

		return entity.User{}, nil
	}
}
