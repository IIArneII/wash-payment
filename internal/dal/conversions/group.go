package conversions

import (
	"wash-payment/internal/app/entity"
	"wash-payment/internal/dal/dbmodels"
)

func GroupFromDB(gr dbmodels.Group) entity.Group {
	return entity.Group{
		ID:             gr.ID,
		OrganizationID: gr.OrganizationID,
		Name:           gr.Name,
		Description:    gr.Description,
		Version:        gr.Version,
		Deleted:        gr.Deleted,
	}
}

func GroupToDB(gr entity.Group) dbmodels.Group {
	return dbmodels.Group{
		ID:             gr.ID,
		OrganizationID: gr.OrganizationID,
		Name:           gr.Name,
		Description:    gr.Description,
		Version:        gr.Version,
		Deleted:        gr.Deleted,
	}
}

func GroupUpdateToDB(gr entity.GroupUpdate) dbmodels.GroupUpdate {
	userUpdate := dbmodels.GroupUpdate{}

	if gr.Name != nil {
		userUpdate.Name = gr.Name
	}
	if gr.Description != nil {
		userUpdate.Description = gr.Description
	}
	if gr.Version != nil {
		userUpdate.Version = gr.Version
	}
	if gr.Deleted != nil {
		userUpdate.Deleted = gr.Deleted
	}

	return userUpdate
}
