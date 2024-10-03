package user

import (
	"context"
	"errors"
	"wash-payment/internal/app"
	"wash-payment/internal/app/entity"
)

func (s *userService) Get(ctx context.Context, id string) (entity.User, error) {
	user, err := s.userRepo.Get(ctx, id)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (s *userService) Upsert(ctx context.Context, user entity.User) (entity.User, error) {
	dbUser, err := s.userRepo.Get(ctx, user.ID)
	if err != nil {
		if errors.Is(err, app.ErrNotFound) {
			newUser, err := s.userRepo.Create(ctx, user)
			if err != nil {
				return entity.User{}, err
			}
			return newUser, nil
		}
		return entity.User{}, err
	} else {
		if dbUser.Version >= user.Version {
			return entity.User{}, app.ErrOldVersion
		}

		userUpdate := userToUpdate(user)
		updatedUser, err := s.userRepo.Update(ctx, user.ID, userUpdate)
		if err != nil {
			return entity.User{}, err
		}

		return updatedUser, nil
	}
}

func userToUpdate(user entity.User) entity.UserUpdate {
	return entity.UserUpdate{
		Name:    &user.Name,
		Email:   &user.Email,
		Version: &user.Version,
		Role:    &user.Role,
	}
}
