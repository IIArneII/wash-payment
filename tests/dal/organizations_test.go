package dal

import (
	"testing"
	"wash-payment/internal/app"
	"wash-payment/internal/app/entity"

	"github.com/Pallinder/go-randomdata"
	"github.com/powerman/check"
	uuid "github.com/satori/go.uuid"
)

func TestCreateOrganization(tt *testing.T) {
	t := check.T(tt)
	err := truncate()
	t.Nil(err)

	var organization1 = generateOrganization(10000, 1)
	var organization2 = generateOrganization(10000, 1)

	res1, err := repositories.OrganizationRepo.Create(ctx, organization1)
	t.Nil(err)
	t.DeepEqual(res1, organization1)

	res2, err := repositories.OrganizationRepo.Create(ctx, organization2)
	t.Nil(err)
	t.DeepEqual(res2, organization2)

	_, err = repositories.OrganizationRepo.Create(ctx, organization2)
	t.Err(err, app.ErrAlreadyExists)
}

func TestGetOrganization(tt *testing.T) {
	t := check.T(tt)
	err := truncate()
	t.Nil(err)

	var organization1 = generateOrganization(10000, 1)
	var organization2 = generateOrganization(10000, 1)

	_, err = repositories.OrganizationRepo.Create(ctx, organization1)
	t.Nil(err)

	resGet1, err := repositories.OrganizationRepo.Get(ctx, organization1.ID)
	t.Nil(err)
	t.DeepEqual(resGet1, organization1)

	_, err = repositories.OrganizationRepo.Get(ctx, organization2.ID)
	t.Err(err, app.ErrNotFound)
}

func TestUpdateOrganization(tt *testing.T) {
	t := check.T(tt)
	err := truncate()
	t.Nil(err)

	var organization1 = generateOrganization(10000, 1)
	var organization2 = generateOrganization(10000, 1)

	_, err = repositories.OrganizationRepo.Create(ctx, organization1)
	t.Nil(err)

	organization1.Name = randomdata.FirstName(randomdata.Male)
	organization1.Description = randomdata.RandStringRunes(50)
	organization1.DisplayName = uuid.NewV4().String()
	organization1.Version = int64(2)
	organizationUpdate := entity.OrganizationUpdate{
		Name:        &organization1.Name,
		DisplayName: &organization1.DisplayName,
		Description: &organization1.Description,
		Version:     &organization1.Version,
	}

	updatedOrg1, err := repositories.OrganizationRepo.Update(ctx, organization1.ID, organizationUpdate)
	t.Nil(err)
	t.DeepEqual(updatedOrg1, organization1)

	_, err = repositories.OrganizationRepo.Update(ctx, organization1.ID, entity.OrganizationUpdate{})
	t.Err(err, app.ErrEmptyUpdate)

	_, err = repositories.OrganizationRepo.Update(ctx, organization2.ID, organizationUpdate)
	t.Err(err, app.ErrNotFound)

	organization1.Version = int64(1)
	organizationUpdate = entity.OrganizationUpdate{
		Version: &organization1.Version,
	}
	_, err = repositories.OrganizationRepo.Update(ctx, organization1.ID, organizationUpdate)
	t.Err(err, app.ErrNotFound)

	organization1.Version = int64(3)
	organization1.Deleted = true
	organizationUpdate = entity.OrganizationUpdate{
		Deleted: &organization1.Deleted,
		Version: &organization1.Version,
	}
	updatedOrg1, err = repositories.OrganizationRepo.Update(ctx, organization1.ID, organizationUpdate)
	t.Nil(err)
	t.DeepEqual(updatedOrg1, organization1)
}

func TestListOrganization(tt *testing.T) {
	t := check.T(tt)
	err := truncate()
	t.Nil(err)

	var organization1 = generateOrganization(10000, 1)
	var organization2 = generateOrganization(10000, 1)
	organization1.Name = "a"
	organization2.Name = "b"

	_, err = repositories.OrganizationRepo.Create(ctx, organization1)
	t.Nil(err)

	_, err = repositories.OrganizationRepo.Create(ctx, organization2)
	t.Nil(err)

	filter := entity.OrganizationFilter{Filter: entity.Filter{
		Page:     1,
		PageSize: 10,
	}}
	list, err := repositories.OrganizationRepo.List(ctx, filter)
	t.Nil(err)
	t.Equal(list.Page, filter.Page)
	t.Equal(list.PageSize, filter.PageSize)
	t.Equal(list.TotalItems, 2)
	t.Equal(list.TotalPages, 1)
	t.DeepEqual(list.Items, []entity.Organization{organization1, organization2})

	filter = entity.OrganizationFilter{Filter: entity.Filter{
		Page:     10,
		PageSize: 10,
	}}
	list, err = repositories.OrganizationRepo.List(ctx, filter)
	t.Nil(err)
	t.Equal(list.Page, filter.Page)
	t.Equal(list.PageSize, filter.PageSize)
	t.Equal(list.TotalItems, 2)
	t.Equal(list.TotalPages, 1)
	t.DeepEqual(list.Items, []entity.Organization{})

	filter = entity.OrganizationFilter{Filter: entity.Filter{
		Page:     1,
		PageSize: 1,
	}}
	list, err = repositories.OrganizationRepo.List(ctx, filter)
	t.Nil(err)
	t.Equal(list.Page, filter.Page)
	t.Equal(list.PageSize, filter.PageSize)
	t.Equal(list.TotalItems, 2)
	t.Equal(list.TotalPages, 2)
	t.DeepEqual(list.Items, []entity.Organization{organization1})

	filter = entity.OrganizationFilter{Filter: entity.Filter{
		Page:     2,
		PageSize: 1,
	}}
	list, err = repositories.OrganizationRepo.List(ctx, filter)
	t.Nil(err)
	t.Equal(list.Page, filter.Page)
	t.Equal(list.PageSize, filter.PageSize)
	t.Equal(list.TotalItems, 2)
	t.Equal(list.TotalPages, 2)
	t.DeepEqual(list.Items, []entity.Organization{organization2})
}
