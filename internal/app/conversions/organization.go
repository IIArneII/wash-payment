package conversions

import (
	"wash-payment/internal/app/entity"
	"wash-payment/internal/dal/dbmodels"
	"wash-payment/internal/pkg/openapi/models"
	rabbitEntity "wash-payment/internal/transport/rabbit/entity"

	"github.com/go-openapi/strfmt"
	uuid "github.com/satori/go.uuid"
)

func OrganizationFromDB(dbOrganization dbmodels.Organization) entity.Organization {
	return entity.Organization{
		ID:          dbOrganization.ID,
		Name:        dbOrganization.Name,
		DisplayName: dbOrganization.DisplayName,
		Description: dbOrganization.Description,
		Version:     dbOrganization.Version,
		Balance:     dbOrganization.Balance,
		Deleted:     dbOrganization.Deleted,
	}
}

func OrganizationsFromDB(organizations []dbmodels.Organization) []entity.Organization {
	orgs := []entity.Organization{}
	for _, v := range organizations {
		orgs = append(orgs, OrganizationFromDB(v))
	}
	return orgs
}

func OrganizationCreateToDB(appOrganizationCreate entity.OrganizationCreate) dbmodels.Organization {
	return dbmodels.Organization{
		ID:          appOrganizationCreate.ID,
		Name:        appOrganizationCreate.Name,
		DisplayName: appOrganizationCreate.DisplayName,
		Description: appOrganizationCreate.Description,
		Version:     appOrganizationCreate.Version,
		Deleted:     appOrganizationCreate.Deleted,
	}
}

func OrganizationUpdateToDB(appOrganizationUpdate entity.OrganizationUpdate) dbmodels.OrganizationUpdate {
	userUpdate := dbmodels.OrganizationUpdate{}

	if appOrganizationUpdate.Name != nil {
		userUpdate.Name = appOrganizationUpdate.Name
	}
	if appOrganizationUpdate.DisplayName != nil {
		userUpdate.DisplayName = appOrganizationUpdate.DisplayName
	}
	if appOrganizationUpdate.Description != nil {
		userUpdate.Description = appOrganizationUpdate.Description
	}
	if appOrganizationUpdate.Version != nil {
		userUpdate.Version = appOrganizationUpdate.Version
	}

	return userUpdate
}

func OrganizationCreateFromRabbit(rabbitOrganization rabbitEntity.Organization) (entity.OrganizationCreate, error) {
	id, err := uuid.FromString(rabbitOrganization.ID)
	if err != nil {
		return entity.OrganizationCreate{}, err
	}

	return entity.OrganizationCreate{
		ID:          id,
		Name:        rabbitOrganization.Name,
		DisplayName: rabbitOrganization.DisplayName,
		Description: rabbitOrganization.Description,
		Version:     rabbitOrganization.Version,
		Deleted:     rabbitOrganization.Deleted,
	}, nil
}

func OrganizationUpdateFromRabbit(rabbitOrganization rabbitEntity.Organization) entity.OrganizationUpdate {
	return entity.OrganizationUpdate{
		Name:        &rabbitOrganization.Name,
		DisplayName: &rabbitOrganization.DisplayName,
		Description: &rabbitOrganization.Description,
		Version:     &rabbitOrganization.Version,
	}
}

func OrganizationToRest(appOrganization entity.Organization) models.Organization {
	id := strfmt.UUID(appOrganization.ID.String())
	return models.Organization{
		ID:          &id,
		Name:        &appOrganization.Name,
		DisplayName: &appOrganization.DisplayName,
		Description: &appOrganization.Description,
		Balance:     &appOrganization.Balance,
	}
}

func OrganizationsToRest(appOrganizations entity.Page[entity.Organization]) []*models.Organization {
	list := []*models.Organization{}
	for _, v := range appOrganizations.Items {
		org := OrganizationToRest(v)
		list = append(list, &org)
	}
	return list
}

func OrganizationCreateToOrganizationUpdate(appOrganizationCreate entity.OrganizationCreate) entity.OrganizationUpdate {
	return entity.OrganizationUpdate{
		Name:        &appOrganizationCreate.Name,
		DisplayName: &appOrganizationCreate.DisplayName,
		Description: &appOrganizationCreate.Description,
		Version:     &appOrganizationCreate.Version,
	}
}
