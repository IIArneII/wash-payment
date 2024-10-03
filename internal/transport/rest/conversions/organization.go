package conversions

import (
	"wash-payment/internal/app/entity"
	"wash-payment/internal/pkg/openapi/models"
	"wash-payment/internal/pkg/openapi/restapi/operations/organizations"

	"github.com/go-openapi/strfmt"
	uuid "github.com/satori/go.uuid"
)

func ServicePricesToRest(appServicePrices entity.ServicePrices) models.ServicePrices {
	return models.ServicePrices{
		Bonus: &appServicePrices.Bonus,
		Sbp:   &appServicePrices.Sbp,
	}
}

func ServicePricesFromRest(restServicePrices models.ServicePrices) entity.ServicePrices {
	return entity.ServicePrices{
		Bonus: *restServicePrices.Bonus,
		Sbp:   *restServicePrices.Sbp,
	}
}

func OrganizationToRest(appOrganization entity.Organization) models.Organization {
	id := strfmt.UUID(appOrganization.ID.String())
	sp := ServicePricesToRest(appOrganization.ServicePrices)
	return models.Organization{
		ID:            &id,
		Name:          &appOrganization.Name,
		DisplayName:   &appOrganization.DisplayName,
		Description:   &appOrganization.Description,
		Balance:       &appOrganization.Balance,
		ServicePrices: &sp,
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

func OrganizationsFilterFromRest(params organizations.ListParams) (entity.OrganizationFilter, error) {
	var ids []uuid.UUID
	if params.Ids != nil {
		ids = []uuid.UUID{}
		for _, v := range params.Ids {
			id, err := uuid.FromString(string(v))
			if err != nil {
				return entity.OrganizationFilter{}, err
			}
			ids = append(ids, id)
		}
	}

	return entity.OrganizationFilter{
		Filter: entity.NewFilter(int(*params.Page), int(*params.PageSize)),
		IDs:    ids,
	}, nil
}
