package conversions

import (
	"wash-payment/internal/app/entity"
	"wash-payment/internal/dal/dbmodels"
)

func WashServerFromDB(gr dbmodels.WashServer) entity.WashServer {
	return entity.WashServer{
		ID:          gr.ID,
		Title:       gr.Title,
		Description: gr.Description,
		GroupID:     gr.GroupID,
		Version:     gr.Version,
		Deleted:     gr.Deleted,
	}
}

func WashServerToDB(gr entity.WashServer) dbmodels.WashServer {
	return dbmodels.WashServer{
		ID:          gr.ID,
		Title:       gr.Title,
		Description: gr.Description,
		GroupID:     gr.GroupID,
		Version:     gr.Version,
		Deleted:     gr.Deleted,
	}
}

func WashServerUpdateToDB(gr entity.WashServerUpdate) dbmodels.WashServerUpdate {
	washServerUpdate := dbmodels.WashServerUpdate{}

	if gr.Title != nil {
		washServerUpdate.Title = gr.Title
	}
	if gr.Description != nil {
		washServerUpdate.Description = gr.Description
	}
	if gr.Version != nil {
		washServerUpdate.Version = gr.Version
	}
	if gr.Deleted != nil {
		washServerUpdate.Deleted = gr.Deleted
	}
	if gr.GroupID != nil {
		washServerUpdate.GroupID.UUID = *gr.GroupID
		washServerUpdate.GroupID.Valid = true
	}

	return washServerUpdate
}
