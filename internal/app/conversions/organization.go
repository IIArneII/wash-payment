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

func OrganizationToRest(appOrganizationCreate entity.Organization) models.Organization {
	id := strfmt.UUID(appOrganizationCreate.ID.String())
	return models.Organization{
		ID:          &id,
		Name:        &appOrganizationCreate.Name,
		DisplayName: &appOrganizationCreate.DisplayName,
		Description: &appOrganizationCreate.Description,
		Balance:     &appOrganizationCreate.Balance,
	}
}
