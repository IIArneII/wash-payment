package dal

import (
	"testing"
	"wash-payment/internal/app"
	"wash-payment/internal/app/entity"

	"github.com/Pallinder/go-randomdata"
	"github.com/powerman/check"
)

func TestCreateUser(tt *testing.T) {
	t := check.T(tt)
	err := truncate()
	t.Nil(err)

	var user1 = generateUser(entity.AdminRole, nil, 1)
	var user2 = generateUser(entity.AdminRole, nil, 1)

	res1, err := repositories.UserRepo.Create(ctx, user1)
	t.Nil(err)
	t.DeepEqual(res1, user1)

	res2, err := repositories.UserRepo.Create(ctx, user2)
	t.Nil(err)
	t.DeepEqual(res2, user2)

	_, err = repositories.UserRepo.Create(ctx, user2)
	t.Err(err, app.ErrAlreadyExists)
}

func TestGetUser(tt *testing.T) {
	t := check.T(tt)
	err := truncate()
	t.Nil(err)

	var user1 = generateUser(entity.AdminRole, nil, 1)
	var user2 = generateUser(entity.AdminRole, nil, 1)

	_, err = repositories.UserRepo.Create(ctx, user1)
	t.Nil(err)

	resGet1, err := repositories.UserRepo.Get(ctx, user1.ID)
	t.Nil(err)
	t.DeepEqual(resGet1, user1)

	_, err = repositories.UserRepo.Get(ctx, user2.ID)
	t.Err(err, app.ErrNotFound)
}

func TestUpdateUser(tt *testing.T) {
	t := check.T(tt)
	err := truncate()
	t.Nil(err)

	var user1 = generateUser(entity.AdminRole, nil, 1)
	var user2 = generateUser(entity.AdminRole, nil, 1)

	_, err = repositories.UserRepo.Create(ctx, user1)
	t.Nil(err)

	user1.Role = entity.SystemManagerRole
	user1.Email = randomdata.Email()
	user1.Name = randomdata.FullName(randomdata.RandomGender)
	user1.Version = int64(2)
	updateUser := entity.UserUpdate{
		Role:    &user1.Role,
		Email:   &user1.Email,
		Name:    &user1.Name,
		Version: &user1.Version,
	}

	updatedUser1, err := repositories.UserRepo.Update(ctx, user1.ID, updateUser)
	t.Nil(err)
	t.DeepEqual(updatedUser1, user1)

	_, err = repositories.UserRepo.Update(ctx, user1.ID, entity.UserUpdate{})
	t.Err(err, app.ErrEmptyUpdate)

	_, err = repositories.UserRepo.Update(ctx, user2.ID, updateUser)
	t.Err(err, app.ErrNotFound)

	user1.Version = int64(1)
	updateUser = entity.UserUpdate{
		Version: &user1.Version,
	}
	_, err = repositories.UserRepo.Update(ctx, user1.ID, updateUser)
	t.Err(err, app.ErrNotFound)
}
