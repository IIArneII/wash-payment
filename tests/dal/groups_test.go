package dal

import (
	"testing"
	"wash-payment/internal/app/conversions"
	"wash-payment/internal/app/entity"
	"wash-payment/internal/dal/dbmodels"

	"github.com/Pallinder/go-randomdata"
	"github.com/powerman/check"
)

func TestCreateGroup(tt *testing.T) {
	t := check.T(tt)

	var organization1 = generateOrganization(10000, 1)
	var group1 = generateGroup(organization1.ID, 1)
	var group2 = generateGroup(organization1.ID, 1)

	_, err := repositories.OrganizationRepo.Create(ctx, organization1)
	t.Nil(err)

	res1, err := repositories.GroupRepo.Create(ctx, group1)
	t.Nil(err)
	t.DeepEqual(res1, group1)

	res2, err := repositories.GroupRepo.Create(ctx, group2)
	t.Nil(err)
	t.DeepEqual(res2, group2)

	_, err = repositories.GroupRepo.Create(ctx, group2)
	t.Err(err, dbmodels.ErrAlreadyExists)
}

func TestGetGroup(tt *testing.T) {
	t := check.T(tt)

	var organization1 = generateOrganization(10000, 1)
	var group1 = generateGroup(organization1.ID, 1)
	var group2 = generateGroup(organization1.ID, 1)

	_, err := repositories.OrganizationRepo.Create(ctx, organization1)
	t.Nil(err)

	_, err = repositories.GroupRepo.Create(ctx, group1)
	t.Nil(err)

	resGet1, err := repositories.GroupRepo.Get(ctx, group1.ID)
	t.Nil(err)
	t.DeepEqual(resGet1, group1)

	_, err = repositories.GroupRepo.Get(ctx, group2.ID)
	t.Err(err, dbmodels.ErrNotFound)
}

func TestUpdateGroup(tt *testing.T) {
	t := check.T(tt)

	var organization1 = generateOrganization(10000, 1)
	var group1 = generateGroup(organization1.ID, 1)
	var group2 = generateGroup(organization1.ID, 1)

	_, err := repositories.OrganizationRepo.Create(ctx, organization1)
	t.Nil(err)

	_, err = repositories.GroupRepo.Create(ctx, group1)
	t.Nil(err)

	newName := randomdata.FirstName(randomdata.Male)
	newDescription := randomdata.RandStringRunes(50)
	newVersion := int64(2)
	groupUpdate := dbmodels.GroupUpdate{
		Name:        &newName,
		Description: &newDescription,
		Version:     &newVersion,
	}

	err = repositories.GroupRepo.Update(ctx, group1.ID, groupUpdate)
	t.Nil(err)

	group1.Name = newName
	group1.Description = newDescription
	group1.Version = newVersion

	resGet1, err := repositories.GroupRepo.Get(ctx, group1.ID)
	t.Nil(err)
	t.DeepEqual(resGet1, group1)

	err = repositories.GroupRepo.Update(ctx, group1.ID, dbmodels.GroupUpdate{})
	t.Err(err, dbmodels.ErrEmptyUpdate)

	err = repositories.GroupRepo.Update(ctx, group2.ID, groupUpdate)
	t.Err(err, dbmodels.ErrNotFound)

	newVersion = 1
	newName = randomdata.FullName(randomdata.RandomGender)
	groupUpdate = dbmodels.GroupUpdate{
		Name:    &newName,
		Version: &newVersion,
	}
	err = repositories.GroupRepo.Update(ctx, organization1.ID, groupUpdate)
	t.Err(err, dbmodels.ErrNotFound)
}

func TestDeleteGroup(tt *testing.T) {
	t := check.T(tt)

	var organization1 = generateOrganization(10000, 1)
	var group1 = generateGroup(organization1.ID, 1)

	_, err := repositories.OrganizationRepo.Create(ctx, organization1)
	t.Nil(err)

	_, err = repositories.GroupRepo.Create(ctx, group1)
	t.Nil(err)

	err = repositories.GroupRepo.Delete(ctx, group1.ID)
	t.Nil(err)

	_, err = repositories.GroupRepo.Get(ctx, group1.ID)
	t.Err(err, dbmodels.ErrNotFound)
}

func TestUpsertServiceGroup(tt *testing.T) {
	t := check.T(tt)

	var organization1 = generateOrganization(10000, 1)
	var group1 = generateGroup(organization1.ID, 1)
	var group2 = generateGroup(organization1.ID, 1)

	var groupService1 = conversions.GroupFromDB(group1)
	var groupService2 = conversions.GroupFromDB(group2)

	// TEST CREATE

	_, err := repositories.OrganizationRepo.Create(ctx, organization1)
	t.Nil(err)

	res1, err := service.GroupService.Upsert(ctx, groupService1)
	t.Nil(err)
	t.DeepEqual(res1, groupService1)

	res2, err := service.GroupService.Upsert(ctx, groupService2)
	t.Nil(err)
	t.DeepEqual(res2, groupService2)

	_, err = service.GroupService.Upsert(ctx, entity.Group{})
	t.Err(err, dbmodels.ErrNotFound)

	// TEST UPDATE

	newName := randomdata.FirstName(randomdata.Male)
	newDescription := randomdata.RandStringRunes(50)
	newVersion := int64(2)

	groupService1.Name = newName
	groupService1.Description = newDescription
	groupService1.Version = newVersion

	resUpdate, err := service.GroupService.Upsert(ctx, groupService1)

	t.Nil(err)
	t.DeepEqual(resUpdate, entity.Group{})

	newName = randomdata.FirstName(randomdata.Male)
	newVersion = int64(3)
	groupService1.Name = newName
	groupService1.Version = newVersion

	resUpdate2, err := service.GroupService.Upsert(ctx, groupService1)
	t.Nil(err)
	t.DeepEqual(resUpdate2, entity.Group{})

}
