// Code generated by go-swagger; DO NOT EDIT.

package spoons

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"

	models "github.com/xbcsmith/spoon/models"
)

// AddSpoonHandlerFunc turns a function with the right signature into a add spoon handler
type AddSpoonHandlerFunc func(AddSpoonParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn AddSpoonHandlerFunc) Handle(params AddSpoonParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// AddSpoonHandler interface for that can handle valid add spoon params
type AddSpoonHandler interface {
	Handle(AddSpoonParams, *models.Principal) middleware.Responder
}

// NewAddSpoon creates a new http.Handler for the add spoon operation
func NewAddSpoon(ctx *middleware.Context, handler AddSpoonHandler) *AddSpoon {
	return &AddSpoon{Context: ctx, Handler: handler}
}

/*AddSpoon swagger:route POST /v1/spoons spoons addSpoon

AddSpoon add spoon API

*/
type AddSpoon struct {
	Context *middleware.Context
	Handler AddSpoonHandler
}

func (o *AddSpoon) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAddSpoonParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
