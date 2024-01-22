package dal

import (
	"testing"
	"wash-payment/internal/app/conversions"
	"wash-payment/internal/app/entity"
	"wash-payment/internal/dal/dbmodels"

	"github.com/Pallinder/go-randomdata"
	"github.com/powerman/check"
	uuid "github.com/satori/go.uuid"
)

func TestCreateUser(tt *testing.T) {
	t := check.T(tt)

	var user1 = generateUser(dbmodels.AdminRole, uuid.NullUUID{}, 1)
	var user2 = generateUser(dbmodels.AdminRole, uuid.NullUUID{}, 1)

	res1, err := repositories.UserRepo.Create(ctx, user1)
	t.Nil(err)
	t.DeepEqual(res1, user1)

	res2, err := repositories.UserRepo.Create(ctx, user2)
	t.Nil(err)
	t.DeepEqual(res2, user2)

	_, err = repositories.UserRepo.Create(ctx, user2)
	t.Err(err, dbmodels.ErrAlreadyExists)
}

func TestGetUser(tt *testing.T) {
	t := check.T(tt)

	var user1 = generateUser(dbmodels.AdminRole, uuid.NullUUID{}, 1)
	var user2 = generateUser(dbmodels.AdminRole, uuid.NullUUID{}, 1)

	_, err := repositories.UserRepo.Create(ctx, user1)
	t.Nil(err)

	resGet1, err := repositories.UserRepo.Get(ctx, user1.ID)
	t.Nil(err)
	t.DeepEqual(resGet1, user1)

	_, err = repositories.UserRepo.Get(ctx, user2.ID)
	t.Err(err, dbmodels.ErrNotFound)
}

func TestUpdateUser(tt *testing.T) {
	t := check.T(tt)

	var user1 = generateUser(dbmodels.AdminRole, uuid.NullUUID{}, 1)
	var user2 = generateUser(dbmodels.AdminRole, uuid.NullUUID{}, 1)

	_, err := repositories.UserRepo.Create(ctx, user1)
	t.Nil(err)

	newRole := dbmodels.SystemManagerRole
	newVersion := int64(2)
	newEmail := randomdata.Email()
	newName := randomdata.FullName(randomdata.RandomGender)
	updateUser := dbmodels.UserUpdate{
		Role:    &newRole,
		Email:   &newEmail,
		Name:    &newName,
		Version: &newVersion,
	}

	err = repositories.UserRepo.Update(ctx, user1.ID, updateUser)
	t.Nil(err)

	user1.Role = dbmodels.SystemManagerRole
	user1.Email = newEmail
	user1.Name = newName
	user1.Version = newVersion

	resGet1, err := repositories.UserRepo.Get(ctx, user1.ID)
	t.Nil(err)
	t.DeepEqual(resGet1, user1)

	err = repositories.UserRepo.Update(ctx, user1.ID, dbmodels.UserUpdate{})
	t.Err(err, dbmodels.ErrEmptyUpdate)

	err = repositories.UserRepo.Update(ctx, user2.ID, updateUser)
	t.Err(err, dbmodels.ErrNotFound)

	newVersion = 1
	newName = randomdata.FullName(randomdata.RandomGender)
	updateUser = dbmodels.UserUpdate{
		Name:    &newName,
		Version: &newVersion,
	}
	err = repositories.UserRepo.Update(ctx, user1.ID, updateUser)
	t.Err(err, dbmodels.ErrNotFound)
}

func TestUpsertServiceUser(tt *testing.T) {
	t := check.T(tt)

	var user1 = generateUser(dbmodels.AdminRole, uuid.NullUUID{}, 1)
	var user2 = generateUser(dbmodels.AdminRole, uuid.NullUUID{}, 1)

	var userService1 = conversions.UserFromDB(user1)
	var userService2 = conversions.UserFromDB(user2)

	// TEST CREATE

	res1, err := service.UserService.Upsert(ctx, userService1)
	t.Nil(err)
	t.DeepEqual(res1, userService1)

	res2, err := service.UserService.Upsert(ctx, userService2)
	t.Nil(err)
	t.DeepEqual(res2, userService2)

	_, err = service.UserService.Upsert(ctx, entity.User{})
	t.Err(err, dbmodels.ErrNotFound)

	// TEST UPDATE

	newVersion := int64(2)
	newEmail := randomdata.Email()
	newName := randomdata.FullName(randomdata.RandomGender)

	userService1.Role = entity.SystemManagerRole
	userService1.Version = newVersion
	userService1.Email = newEmail
	userService1.Name = newName

	resUpdate, err := service.UserService.Upsert(ctx, userService1)

	t.Nil(err)
	t.DeepEqual(resUpdate, entity.User{})

	newName = randomdata.FirstName(randomdata.Male)
	newVersion = int64(3)

	userService1.Name = newName
	userService1.Version = newVersion

	resUpdate2, err := service.UserService.Upsert(ctx, userService1)
	t.Nil(err)
	t.DeepEqual(resUpdate2, entity.User{})
}
