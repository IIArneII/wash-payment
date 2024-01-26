// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"wash-payment/internal/app"
)

// DepositHandlerFunc turns a function with the right signature into a deposit handler
type DepositHandlerFunc func(DepositParams, *app.Auth) DepositResponder

// Handle executing the request and returning a response
func (fn DepositHandlerFunc) Handle(params DepositParams, principal *app.Auth) DepositResponder {
	return fn(params, principal)
}

// DepositHandler interface for that can handle valid deposit params
type DepositHandler interface {
	Handle(DepositParams, *app.Auth) DepositResponder
}

// NewDeposit creates a new http.Handler for the deposit operation
func NewDeposit(ctx *middleware.Context, handler DepositHandler) *Deposit {
	return &Deposit{Context: ctx, Handler: handler}
}

/*
	Deposit swagger:route POST /organizations/{id}/deposit Organizations deposit

# Top up balance

Increase the balance of the specified organization by the specified number of kopecks
*/
type Deposit struct {
	Context *middleware.Context
	Handler DepositHandler
}

func (o *Deposit) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDepositParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal *app.Auth
	if uprinc != nil {
		principal = uprinc.(*app.Auth) // this is really a app.Auth, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
