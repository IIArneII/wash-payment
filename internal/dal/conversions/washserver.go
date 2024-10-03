package conversions

import (
	"wash-payment/internal/app/entity"
	"wash-payment/internal/dal/dbmodels"
)

func WashServerFromTransactionDB(transaction dbmodels.Transaction) *entity.WashServer {
	if !transaction.WashServerID.Valid {
		return nil
	}
	var name string
	if transaction.WashServerName != nil {
		name = *transaction.WashServerName
	}
	var description string
	if transaction.WashServerDescription != nil {
		description = *transaction.WashServerDescription
	}
	var version int
	if transaction.WashServerVersion != nil {
		version = *transaction.WashServerVersion
	}
	var deleted bool
	if transaction.WashServerDeleted != nil {
		deleted = *transaction.WashServerDeleted
	}
	var ownerID string
	if transaction.WashServerOwnerID != nil {
		ownerID = *transaction.WashServerOwnerID
	}

	return &entity.WashServer{
		ID:          transaction.WashServerID.UUID,
		GroupID:     transaction.WashServerGroupID.UUID,
		Name:        name,
		Description: description,
		OwnerID:     ownerID,
		Version:     version,
		Deleted:     deleted,
	}
}

func WashServerFromDB(gr dbmodels.WashServer) entity.WashServer {
	return entity.WashServer{
		ID:          gr.ID,
		Name:        gr.Name,
		Description: gr.Description,
		OwnerID:     gr.OwnerID,
		GroupID:     gr.GroupID,
		Version:     gr.Version,
		Deleted:     gr.Deleted,
	}
}

func WashServerToDB(gr entity.WashServer) dbmodels.WashServer {
	return dbmodels.WashServer{
		ID:          gr.ID,
		Name:        gr.Name,
		Description: gr.Description,
		OwnerID:     gr.OwnerID,
		GroupID:     gr.GroupID,
		Version:     gr.Version,
		Deleted:     gr.Deleted,
	}
}

func WashServerUpdateToDB(gr entity.WashServerUpdate) dbmodels.WashServerUpdate {
	washServerUpdate := dbmodels.WashServerUpdate{}

	if gr.Name != nil {
		washServerUpdate.Name = gr.Name
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
	if gr.OwnerID != nil {
		washServerUpdate.OwnerID = gr.OwnerID
	}
	if gr.GroupID != nil {
		washServerUpdate.GroupID.UUID = *gr.GroupID
		washServerUpdate.GroupID.Valid = true
	}

	return washServerUpdate
}
