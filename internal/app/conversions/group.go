package conversions

import (
	"wash-payment/internal/app/entity"
	"wash-payment/internal/dal/dbmodels"
	rabbitEntity "wash-payment/internal/transport/rabbit/entity"

	uuid "github.com/satori/go.uuid"
)

func GroupFromDB(dbOrganization dbmodels.Group) entity.Group {
	return entity.Group{
		ID:             dbOrganization.ID,
		OrganizationID: dbOrganization.OrganizationID,
		Name:           dbOrganization.Name,
		Description:    dbOrganization.Description,
		Version:        dbOrganization.Version,
		Deleted:        dbOrganization.Deleted,
	}
}

func GroupToDB(appGroup entity.Group) dbmodels.Group {
	return dbmodels.Group{
		ID:             appGroup.ID,
		OrganizationID: appGroup.OrganizationID,
		Name:           appGroup.Name,
		Description:    appGroup.Description,
		Version:        appGroup.Version,
		Deleted:        appGroup.Deleted,
	}
}

func GroupUpdateToDB(appGroupUpdate entity.GroupUpdate) dbmodels.GroupUpdate {
	userUpdate := dbmodels.GroupUpdate{}

	if appGroupUpdate.Name != nil {
		userUpdate.Name = appGroupUpdate.Name
	}
	if appGroupUpdate.Description != nil {
		userUpdate.Description = appGroupUpdate.Description
	}
	if appGroupUpdate.Version != nil {
		userUpdate.Version = appGroupUpdate.Version
	}

	return userUpdate
}

func GroupFromRabbit(rabbitGroup rabbitEntity.Group) (entity.Group, error) {
	id, err := uuid.FromString(rabbitGroup.ID)
	if err != nil {
		return entity.Group{}, err
	}

	orgId, err := uuid.FromString(rabbitGroup.OrganizationID)
	if err != nil {
		return entity.Group{}, err
	}

	return entity.Group{
		ID:             id,
		OrganizationID: orgId,
		Name:           rabbitGroup.Name,
		Description:    rabbitGroup.Description,
		Version:        rabbitGroup.Version,
		Deleted:        rabbitGroup.Deleted,
	}, nil
}

func GroupUpdateFromRabbit(rabbitGroup rabbitEntity.Group) entity.GroupUpdate {
	return entity.GroupUpdate{
		Name:        &rabbitGroup.Name,
		Description: &rabbitGroup.Description,
		Version:     &rabbitGroup.Version,
	}
}

func GroupToGroupUpdate(appGroup entity.Group) entity.GroupUpdate {
	return entity.GroupUpdate{
		Name:        &appGroup.Name,
		Description: &appGroup.Description,
		Version:     &appGroup.Version,
	}
}
