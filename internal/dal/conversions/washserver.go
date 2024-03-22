package conversions

import (
	"wash-payment/internal/app/entity"
	"wash-payment/internal/dal/dbmodels"
)

func WashServerFromTransactionDB(transaction dbmodels.Transaction) *entity.WashServer {
	if !transaction.WashServerID.Valid {
		return nil
	}
	var title string
	if transaction.WashServerTitle != nil {
		title = *transaction.WashServerTitle
	}
	var description string
	if transaction.WashServerDescription != nil {
		description = *transaction.WashServerDescription
	}
	var version int64
	if transaction.WashServerVersion != nil {
		version = *transaction.WashServerVersion
	}
	var deleted bool
	if transaction.WashServerDeleted != nil {
		deleted = *transaction.WashServerDeleted
	}

	return &entity.WashServer{
		ID:          transaction.WashServerID.UUID,
		GroupID:     transaction.WashServerGroupID.UUID,
		Title:       title,
		Description: description,
		Version:     version,
		Deleted:     deleted,
	}
}

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
