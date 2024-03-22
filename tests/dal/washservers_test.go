package dal

import (
	"testing"
	"wash-payment/internal/app"
	"wash-payment/internal/app/entity"

	"github.com/Pallinder/go-randomdata"
	"github.com/powerman/check"
)

func TestCreateWashServer(tt *testing.T) {
	t := check.T(tt)
	err := truncate()
	t.Nil(err)

	organization1 := generateOrganization(10000, 1)
	group := generateGroup(organization1.ID, 1)
	washServer1 := generateWashServer(group.ID, 1)
	washServer2 := generateWashServer(group.ID, 1)

	_, err = repositories.OrganizationRepo.Create(ctx, organization1)
	t.Nil(err)

	_, err = repositories.GroupRepo.Create(ctx, group)
	t.Nil(err)

	dbWashServer1, err := repositories.WashServerRepo.Create(ctx, washServer1)
	t.Nil(err)
	t.Equal(dbWashServer1, washServer1)

	dbWashServer2, err := repositories.WashServerRepo.Create(ctx, washServer2)
	t.Nil(err)
	t.Equal(dbWashServer2, washServer2)

	_, err = repositories.WashServerRepo.Create(ctx, washServer2)
	t.Err(err, app.ErrAlreadyExists)
}

func TestGetWashServer(tt *testing.T) {
	t := check.T(tt)
	err := truncate()
	t.Nil(err)

	organization1 := generateOrganization(10000, 1)
	group := generateGroup(organization1.ID, 1)
	washServer1 := generateWashServer(group.ID, 1)
	washServer2 := generateWashServer(group.ID, 1)

	_, err = repositories.OrganizationRepo.Create(ctx, organization1)
	t.Nil(err)

	_, err = repositories.GroupRepo.Create(ctx, group)
	t.Nil(err)

	_, err = repositories.WashServerRepo.Create(ctx, washServer1)
	t.Nil(err)

	dbWashServer1, err := repositories.WashServerRepo.Get(ctx, washServer1.ID)
	t.Nil(err)
	t.DeepEqual(dbWashServer1, washServer1)

	_, err = repositories.WashServerRepo.Get(ctx, washServer2.ID)
	t.Err(err, app.ErrNotFound)
}

func TestUpdateWashServer(tt *testing.T) {
	t := check.T(tt)
	err := truncate()
	t.Nil(err)

	organization1 := generateOrganization(10000, 1)
	group1 := generateGroup(organization1.ID, 1)
	group2 := generateGroup(organization1.ID, 1)
	washServer1 := generateWashServer(group1.ID, 1)
	washServer2 := generateWashServer(group1.ID, 1)

	_, err = repositories.OrganizationRepo.Create(ctx, organization1)
	t.Nil(err)

	_, err = repositories.GroupRepo.Create(ctx, group1)
	t.Nil(err)

	_, err = repositories.GroupRepo.Create(ctx, group2)
	t.Nil(err)

	_, err = repositories.WashServerRepo.Create(ctx, washServer1)
	t.Nil(err)

	washServer1.Title = randomdata.FirstName(randomdata.Male)
	washServer1.Description = randomdata.RandStringRunes(50)
	washServer1.GroupID = group2.ID
	washServer1.Version = int64(2)
	washServer1Update := entity.WashServerUpdate{
		Title:       &washServer1.Title,
		Description: &washServer1.Description,
		GroupID:     &washServer1.GroupID,
		Version:     &washServer1.Version,
	}

	updatedWashServer1, err := repositories.WashServerRepo.Update(ctx, washServer1.ID, washServer1Update)
	t.Nil(err)
	t.DeepEqual(updatedWashServer1, washServer1)

	_, err = repositories.WashServerRepo.Update(ctx, washServer1.ID, entity.WashServerUpdate{})
	t.Err(err, app.ErrEmptyUpdate)

	_, err = repositories.WashServerRepo.Update(ctx, washServer2.ID, washServer1Update)
	t.Err(err, app.ErrNotFound)

	washServer1.Version = int64(1)
	washServer1Update = entity.WashServerUpdate{
		Version: &washServer1.Version,
	}

	_, err = repositories.WashServerRepo.Update(ctx, washServer1.ID, washServer1Update)
	t.Err(err, app.ErrNotFound)

	washServer1.Version = int64(3)
	washServer1.Deleted = true
	washServer1Update = entity.WashServerUpdate{
		Deleted: &washServer1.Deleted,
		Version: &washServer1.Version,
	}
	updatedWashServer1, err = repositories.WashServerRepo.Update(ctx, washServer1.ID, washServer1Update)
	t.Nil(err)
	t.DeepEqual(updatedWashServer1, washServer1)
}
