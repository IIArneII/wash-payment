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

func (s *userService) Create(ctx context.Context, user entity.User) (entity.User, error) {
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

func (s *userService) Update(ctx context.Context, userID string, userUpdate entity.UserUpdate) error {
	dbUserUpdate := conversions.UserUpdateToDB(userUpdate)

	err := s.userRepo.Update(ctx, userID, dbUserUpdate)
	if err != nil {
		if errors.Is(err, dbmodels.ErrNotFound) {
			err = app.ErrNotFound
		} else if errors.Is(err, dbmodels.ErrEmptyUpdate) {
			err = app.ErrBadRequest
		}

		return err
	}

	return nil
}
