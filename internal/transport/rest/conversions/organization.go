package conversions

import (
	"wash-payment/internal/app/entity"
	"wash-payment/internal/pkg/openapi/models"

	"github.com/go-openapi/strfmt"
)

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

func OrganizationsToRest(appOrganizations entity.Page[entity.Organization]) *models.OrganizationPage {
	list := []*models.Organization{}
	for _, v := range appOrganizations.Items {
		org := OrganizationToRest(v)
		list = append(list, &org)
	}
	page := int64(appOrganizations.Page)
	pageSize := int64(appOrganizations.PageSize)
	totalPages := int64(appOrganizations.TotalPages)
	totalItems := int64(appOrganizations.TotalItems)
	return &models.OrganizationPage{
		Items:      list,
		Page:       &page,
		PageSize:   &pageSize,
		TotalPages: &totalPages,
		TotalItems: &totalItems,
	}
}
