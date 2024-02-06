package dal

import (
	"testing"
	"wash-payment/internal/app"
	"wash-payment/internal/app/entity"

	"github.com/Pallinder/go-randomdata"
	"github.com/powerman/check"
)

func TestCreateGroup(tt *testing.T) {
	t := check.T(tt)
	err := truncate()
	t.Nil(err)

	var organization1 = generateOrganization(10000, 1)
	var group1 = generateGroup(organization1.ID, 1)
	var group2 = generateGroup(organization1.ID, 1)

	_, err = repositories.OrganizationRepo.Create(ctx, organization1)
	t.Nil(err)

	res1, err := repositories.GroupRepo.Create(ctx, group1)
	t.Nil(err)
	t.DeepEqual(res1, group1)

	res2, err := repositories.GroupRepo.Create(ctx, group2)
	t.Nil(err)
	t.DeepEqual(res2, group2)

	_, err = repositories.GroupRepo.Create(ctx, group2)
	t.Err(err, app.ErrAlreadyExists)
}

func TestGetGroup(tt *testing.T) {
	t := check.T(tt)
	err := truncate()
	t.Nil(err)

	var organization1 = generateOrganization(10000, 1)
	var group1 = generateGroup(organization1.ID, 1)
	var group2 = generateGroup(organization1.ID, 1)

	_, err = repositories.OrganizationRepo.Create(ctx, organization1)
	t.Nil(err)

	_, err = repositories.GroupRepo.Create(ctx, group1)
	t.Nil(err)

	resGet1, err := repositories.GroupRepo.Get(ctx, group1.ID)
	t.Nil(err)
	t.DeepEqual(resGet1, group1)

	_, err = repositories.GroupRepo.Get(ctx, group2.ID)
	t.Err(err, app.ErrNotFound)
}

func TestUpdateGroup(tt *testing.T) {
	t := check.T(tt)
	err := truncate()
	t.Nil(err)

	var organization1 = generateOrganization(10000, 1)
	var group1 = generateGroup(organization1.ID, 1)
	var group2 = generateGroup(organization1.ID, 1)

	_, err = repositories.OrganizationRepo.Create(ctx, organization1)
	t.Nil(err)

	_, err = repositories.GroupRepo.Create(ctx, group1)
	t.Nil(err)

	group1.Name = randomdata.FirstName(randomdata.Male)
	group1.Description = randomdata.RandStringRunes(50)
	group1.Version = int64(2)
	groupUpdate := entity.GroupUpdate{
		Name:        &group1.Name,
		Description: &group1.Description,
		Version:     &group1.Version,
	}

	updatedGroup1, err := repositories.GroupRepo.Update(ctx, group1.ID, groupUpdate)
	t.Nil(err)
	t.DeepEqual(updatedGroup1, group1)

	_, err = repositories.GroupRepo.Update(ctx, group1.ID, entity.GroupUpdate{})
	t.Err(err, app.ErrEmptyUpdate)

	_, err = repositories.GroupRepo.Update(ctx, group2.ID, groupUpdate)
	t.Err(err, app.ErrNotFound)

	group1.Version = int64(1)
	groupUpdate = entity.GroupUpdate{
		Version: &group1.Version,
	}
	_, err = repositories.GroupRepo.Update(ctx, group1.ID, groupUpdate)
	t.Err(err, app.ErrNotFound)

	group1.Version = int64(3)
	group1.Deleted = true
	groupUpdate = entity.GroupUpdate{
		Deleted: &group1.Deleted,
		Version: &group1.Version,
	}
	updatedGroup1, err = repositories.GroupRepo.Update(ctx, group1.ID, groupUpdate)
	t.Nil(err)
	t.DeepEqual(updatedGroup1, group1)
}
