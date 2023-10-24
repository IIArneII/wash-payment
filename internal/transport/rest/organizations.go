package rest

import (
	"fmt"
	"net/http"
	"wash-payment/internal/app"
	"wash-payment/internal/app/conversions"
	"wash-payment/internal/pkg/openapi/restapi/operations"
	"wash-payment/internal/pkg/openapi/restapi/operations/organizations"

	uuid "github.com/satori/go.uuid"
)

func (svc *service) initOrganizationsHandlers(api *operations.WashPaymentAPI) {
	api.OrganizationsDepositHandler = organizations.DepositHandlerFunc(svc.deposit)
}

func (svc *service) deposit(params organizations.DepositParams, profile *app.Auth) organizations.DepositResponder {
	op := "Top up the organization's balance: "
	resp := organizations.NewDepositDefault(http.StatusInternalServerError)

	id, err := uuid.FromString(params.ID.String())
	if err != nil {
		setAPIError(svc.l, op, fmt.Errorf("Wrong organization ID: %w", app.ErrBadRequest), resp)
		return resp
	}

	err = svc.services.OrganizationService.Deposit(params.HTTPRequest.Context(), *profile, id, *params.Body.Amount)
	if err != nil {
		setAPIError(svc.l, op, err, resp)
		return resp
	}

	return organizations.NewDepositNoContent()
}

func (svc *service) get(params organizations.GetParams, profile *app.Auth) organizations.GetResponder {
	op := "Get organization: "
	resp := organizations.NewGetDefault(http.StatusInternalServerError)

	id, err := uuid.FromString(params.ID.String())
	if err != nil {
		setAPIError(svc.l, op, fmt.Errorf("Wrong organization ID: %w", app.ErrBadRequest), resp)
		return resp
	}

	org, err := svc.services.OrganizationService.Get(params.HTTPRequest.Context(), *profile, id)
	if err != nil {
		setAPIError(svc.l, op, err, resp)
		return resp
	}

	orgModel := conversions.OrganizationToRest(org)
	return organizations.NewGetOK().WithPayload(&orgModel)
}
