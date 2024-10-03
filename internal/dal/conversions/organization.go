package conversions

import (
	"wash-payment/internal/app/entity"
	"wash-payment/internal/dal/dbmodels"
)

func OrganizationFromDB(org dbmodels.Organization) entity.Organization {
	return entity.Organization{
		ID:          org.ID,
		Name:        org.Name,
		DisplayName: org.DisplayName,
		Description: org.Description,
		Version:     org.Version,
		Balance:     org.Balance,
		Deleted:     org.Deleted,
		ServicePrices: entity.ServicePrices{
			Bonus: org.ServicePricesBonus,
			Sbp:   org.ServicePricesSbp,
		},
	}
}

func OrganizationsFromDB(organizations []dbmodels.Organization) []entity.Organization {
	orgs := []entity.Organization{}
	for _, v := range organizations {
		orgs = append(orgs, OrganizationFromDB(v))
	}
	return orgs
}

func OrganizationToDB(org entity.Organization) dbmodels.Organization {
	return dbmodels.Organization{
		ID:          org.ID,
		Name:        org.Name,
		DisplayName: org.DisplayName,
		Description: org.Description,
		Version:     org.Version,
		Balance:     org.Balance,
		Deleted:     org.Deleted,
	}
}

func OrganizationUpdateToDB(org entity.OrganizationUpdate) dbmodels.OrganizationUpdate {
	userUpdate := dbmodels.OrganizationUpdate{}

	if org.Name != nil {
		userUpdate.Name = org.Name
	}
	if org.DisplayName != nil {
		userUpdate.DisplayName = org.DisplayName
	}
	if org.Description != nil {
		userUpdate.Description = org.Description
	}
	if org.Deleted != nil {
		userUpdate.Deleted = org.Deleted
	}
	if org.Version != nil {
		userUpdate.Version = org.Version
	}

	return userUpdate
}
