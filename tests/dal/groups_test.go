package dal

import (
	"testing"
	"wash-payment/internal/app/entity"
	"wash-payment/internal/dal/dbmodels"

	"github.com/Pallinder/go-randomdata"
	"github.com/powerman/check"
	uuid "github.com/satori/go.uuid"
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
	var group1 = generateGroupForService(organization1.ID, 1)
	var group2 = generateGroupForService(organization1.ID, 1)

	//TEST Create with Service

	_, err := repositories.OrganizationRepo.Create(ctx, organization1)
	t.Nil(err)

	res1, err := service.GroupService.Upsert(ctx, group1, uuid.Nil, entity.GroupUpdate{})
	t.Nil(err)
	t.DeepEqual(res1, group1)

	res2, err := service.GroupService.Upsert(ctx, group2, uuid.Nil, entity.GroupUpdate{})
	t.Nil(err)
	t.DeepEqual(res2, group2)

	_, err = service.GroupService.Upsert(ctx, group2, uuid.Nil, entity.GroupUpdate{})
	t.Err(err, dbmodels.ErrAlreadyExists)

	//TEST Update with Service

	newName := randomdata.FirstName(randomdata.Male)
	newDescription := randomdata.RandStringRunes(50)
	newVersion := int64(2)
	groupUpdate := generateGroupUpdateForService(newVersion, newName, newDescription)

	_, err = service.GroupService.Upsert(ctx, entity.Group{}, group1.ID, groupUpdate)

	t.Nil(err)

	group1.Name = newName
	group1.Description = newDescription
	group1.Version = newVersion

	resGet1, err := repositories.GroupRepo.Get(ctx, group1.ID)
	t.Nil(err)
	t.DeepEqual(resGet1, group1)

	_, err = service.GroupService.Upsert(ctx, entity.Group{}, group1.ID, entity.GroupUpdate{})
	t.Err(err, dbmodels.ErrEmptyUpdate)

	_, err = service.GroupService.Upsert(ctx, entity.Group{}, group2.ID, groupUpdate)
	t.Err(err, dbmodels.ErrNotFound)

	newVersion = 1
	newName = randomdata.FullName(randomdata.RandomGender)
	groupUpdate = generateGroupUpdateForService(newVersion, newName, "")

	_, err = service.GroupService.Upsert(ctx, entity.Group{}, organization1.ID, groupUpdate)

	t.Err(err, dbmodels.ErrNotFound)
}
