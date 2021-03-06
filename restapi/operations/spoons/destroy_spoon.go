// Code generated by go-swagger; DO NOT EDIT.

package spoons

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"

	models "github.com/xbcsmith/spoon/models"
)

// DestroySpoonHandlerFunc turns a function with the right signature into a destroy spoon handler
type DestroySpoonHandlerFunc func(DestroySpoonParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn DestroySpoonHandlerFunc) Handle(params DestroySpoonParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// DestroySpoonHandler interface for that can handle valid destroy spoon params
type DestroySpoonHandler interface {
	Handle(DestroySpoonParams, *models.Principal) middleware.Responder
}

// NewDestroySpoon creates a new http.Handler for the destroy spoon operation
func NewDestroySpoon(ctx *middleware.Context, handler DestroySpoonHandler) *DestroySpoon {
	return &DestroySpoon{Context: ctx, Handler: handler}
}

/*DestroySpoon swagger:route DELETE /v1/spoons/{id} spoons destroySpoon

DestroySpoon destroy spoon API

*/
type DestroySpoon struct {
	Context *middleware.Context
	Handler DestroySpoonHandler
}

func (o *DestroySpoon) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDestroySpoonParams()

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
