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

// TODO: Переписать тесты для Транзакций(Добавил поля новые), тесты для организаций(withDrowal и deposit)
func TestCreateOrganization(tt *testing.T) {
	t := check.T(tt)

	var organization1 = generateOrganization(10000, 1)
	var organization2 = generateOrganization(10000, 1)

	res1, err := repositories.OrganizationRepo.Create(ctx, organization1)
	t.Nil(err)
	t.DeepEqual(res1, organization1)

	res2, err := repositories.OrganizationRepo.Create(ctx, organization2)
	t.Nil(err)
	t.DeepEqual(res2, organization2)

	_, err = repositories.OrganizationRepo.Create(ctx, organization2)
	t.Err(err, dbmodels.ErrAlreadyExists)
}

func TestGetOrganization(tt *testing.T) {
	t := check.T(tt)

	var organization1 = generateOrganization(10000, 1)
	var organization2 = generateOrganization(10000, 1)

	_, err := repositories.OrganizationRepo.Create(ctx, organization1)
	t.Nil(err)

	resGet1, err := repositories.OrganizationRepo.Get(ctx, organization1.ID)
	t.Nil(err)
	t.DeepEqual(resGet1, organization1)

	_, err = repositories.OrganizationRepo.Get(ctx, organization2.ID)
	t.Err(err, dbmodels.ErrNotFound)
}

func TestUpdateOrganization(tt *testing.T) {
	t := check.T(tt)

	var organization1 = generateOrganization(10000, 1)
	var organization2 = generateOrganization(10000, 1)

	_, err := repositories.OrganizationRepo.Create(ctx, organization1)
	t.Nil(err)

	newName := randomdata.FirstName(randomdata.Male)
	newDescription := randomdata.RandStringRunes(50)
	newDisplayName := uuid.NewV4().String()
	newVersion := int64(2)
	organizationUpdate := dbmodels.OrganizationUpdate{
		Name:        &newName,
		DisplayName: &newDisplayName,
		Description: &newDescription,
		Version:     &newVersion,
	}

	err = repositories.OrganizationRepo.Update(ctx, organization1.ID, organizationUpdate)
	t.Nil(err)

	organization1.Name = newName
	organization1.Description = newDescription
	organization1.DisplayName = newDisplayName
	organization1.Version = newVersion

	resGet1, err := repositories.OrganizationRepo.Get(ctx, organization1.ID)
	t.Nil(err)
	t.DeepEqual(resGet1, organization1)

	err = repositories.OrganizationRepo.Update(ctx, organization1.ID, dbmodels.OrganizationUpdate{})
	t.Err(err, dbmodels.ErrEmptyUpdate)

	err = repositories.OrganizationRepo.Update(ctx, organization2.ID, organizationUpdate)
	t.Err(err, dbmodels.ErrNotFound)

	newVersion = 1
	newName = randomdata.FullName(randomdata.RandomGender)
	organizationUpdate = dbmodels.OrganizationUpdate{
		Name:    &newName,
		Version: &newVersion,
	}
	err = repositories.OrganizationRepo.Update(ctx, organization1.ID, organizationUpdate)
	t.Err(err, dbmodels.ErrNotFound)
}

func TestDeleteOrganization(tt *testing.T) {
	t := check.T(tt)

	var organization1 = generateOrganization(10000, 1)

	_, err := repositories.OrganizationRepo.Create(ctx, organization1)
	t.Nil(err)

	err = repositories.OrganizationRepo.Delete(ctx, organization1.ID)
	t.Nil(err)

	_, err = repositories.OrganizationRepo.Get(ctx, organization1.ID)
	t.Err(err, dbmodels.ErrNotFound)
}

func TestUpsertServiceOrgnization(tt *testing.T) {
	t := check.T(tt)

	var organization1 = generateOrganizationCreate(10000, 1)
	var organization2 = generateOrganizationCreate(10000, 1)
	//TEST CREATE
	var organizationTest = conversions.OrganizationFromDB(conversions.OrganizationCreateToDB(organization1))
	var organizationTest2 = conversions.OrganizationFromDB(conversions.OrganizationCreateToDB(organization2))

	res1, err := service.OrganizationService.Upsert(ctx, organization1)
	t.Nil(err)
	t.DeepEqual(res1, organizationTest)

	res2, err := service.OrganizationService.Upsert(ctx, organization2)
	t.Nil(err)
	t.DeepEqual(res2, organizationTest2)

	_, err = service.OrganizationService.Upsert(ctx, organization2)
	t.Nil(err)

	// TEST UPDATE

	newName := randomdata.FirstName(randomdata.Male)
	newDescription := randomdata.RandStringRunes(50)
	newDisplayName := uuid.NewV4().String()
	newVersion := int64(2)

	organization1.Name = newName
	organization1.Description = newDescription
	organization1.DisplayName = newDisplayName
	organization1.Version = newVersion

	_, err = repositories.OrganizationRepo.Get(ctx, organization1.ID)
	t.Nil(err)

	_, err = service.OrganizationService.Upsert(ctx, entity.OrganizationCreate{})
	t.Err(err, dbmodels.ErrNotFound)

	resUpdate, err := service.OrganizationService.Upsert(ctx, organization1)
	t.Nil(err)
	t.DeepEqual(resUpdate, entity.Organization{})

}
