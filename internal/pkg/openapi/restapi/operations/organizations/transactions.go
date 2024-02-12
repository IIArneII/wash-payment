// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"wash-payment/internal/app/entity"
)

// TransactionsHandlerFunc turns a function with the right signature into a transactions handler
type TransactionsHandlerFunc func(TransactionsParams, *entity.Auth) TransactionsResponder

// Handle executing the request and returning a response
func (fn TransactionsHandlerFunc) Handle(params TransactionsParams, principal *entity.Auth) TransactionsResponder {
	return fn(params, principal)
}

// TransactionsHandler interface for that can handle valid transactions params
type TransactionsHandler interface {
	Handle(TransactionsParams, *entity.Auth) TransactionsResponder
}

// NewTransactions creates a new http.Handler for the transactions operation
func NewTransactions(ctx *middleware.Context, handler TransactionsHandler) *Transactions {
	return &Transactions{Context: ctx, Handler: handler}
}

/*
	Transactions swagger:route GET /organizations/{id}/transactions Organizations transactions

# Get organization transactions

Get a list of transactions for the specified organization
*/
type Transactions struct {
	Context *middleware.Context
	Handler TransactionsHandler
}

func (o *Transactions) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewTransactionsParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal *entity.Auth
	if uprinc != nil {
		principal = uprinc.(*entity.Auth) // this is really a entity.Auth, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
