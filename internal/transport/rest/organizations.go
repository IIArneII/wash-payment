package rest

import (
	"fmt"
	"net/http"
	"wash-payment/internal/app"
	"wash-payment/internal/app/conversions"
	"wash-payment/internal/app/entity"
	"wash-payment/internal/pkg/openapi/restapi/operations"
	"wash-payment/internal/pkg/openapi/restapi/operations/organizations"

	uuid "github.com/satori/go.uuid"
)

func (svc *service) initOrganizationsHandlers(api *operations.WashPaymentAPI) {
	api.OrganizationsDepositHandler = organizations.DepositHandlerFunc(svc.deposit)
	api.OrganizationsGetHandler = organizations.GetHandlerFunc(svc.get)
}

func (svc *service) deposit(params organizations.DepositParams, profile *entity.Auth) organizations.DepositResponder {
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

func (svc *service) get(params organizations.GetParams, profile *entity.Auth) organizations.GetResponder {
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

func (svc *service) list(params organizations.ListParams, profile *entity.Auth) organizations.ListResponder {
	op := "List organizations: "
	resp := organizations.NewListDefault(http.StatusInternalServerError)

	org, err := svc.services.OrganizationService.List(params.HTTPRequest.Context(), *profile, entity.OrganizationFilter{})
	if err != nil {
		setAPIError(svc.l, op, err, resp)
		return resp
	}

	orgModels := conversions.OrganizationsToRest(org)
	return organizations.NewListOK().WithPayload(orgModels)
}

func (svc *service) transactions(params organizations.TransactionsParams, profile *entity.Auth) organizations.TransactionsResponder {
	op := "Transactions organizations: "
	resp := organizations.NewTransactionsDefault(http.StatusInternalServerError)

	txs, err := svc.services.OrganizationService.Transactions(params.HTTPRequest.Context(), *profile, entity.TransactionFilter{})
	if err != nil {
		setAPIError(svc.l, op, err, resp)
		return resp
	}

	txModels := conversions.TransactionsToRest(txs)
	return organizations.NewTransactionsOK().WithPayload(txModels)
}
