package conversions

import (
	"wash-payment/internal/app/entity"
	"wash-payment/internal/dal/dbmodels"
)

func GroupFromTransactionDB(transaction dbmodels.Transaction) *entity.Group {
	if !transaction.GroupID.Valid {
		return nil
	}
	var name string
	if transaction.GroupName != nil {
		name = *transaction.GroupName
	}
	var description string
	if transaction.GroupDescription != nil {
		description = *transaction.GroupDescription
	}
	var version int64
	if transaction.GroupVersion != nil {
		version = *transaction.GroupVersion
	}
	var deleted bool
	if transaction.GroupDeleted != nil {
		deleted = *transaction.GroupDeleted
	}

	return &entity.Group{
		ID:             transaction.GroupID.UUID,
		OrganizationID: transaction.GroupOrganizationID.UUID,
		Name:           name,
		Description:    description,
		Version:        version,
		Deleted:        deleted,
	}
}

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
