// Code generated by go-swagger; DO NOT EDIT.

package spoons

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"

	models "github.com/xbcsmith/spoon/models"
)

// GetSpoonHandlerFunc turns a function with the right signature into a get spoon handler
type GetSpoonHandlerFunc func(GetSpoonParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn GetSpoonHandlerFunc) Handle(params GetSpoonParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// GetSpoonHandler interface for that can handle valid get spoon params
type GetSpoonHandler interface {
	Handle(GetSpoonParams, *models.Principal) middleware.Responder
}

// NewGetSpoon creates a new http.Handler for the get spoon operation
func NewGetSpoon(ctx *middleware.Context, handler GetSpoonHandler) *GetSpoon {
	return &GetSpoon{Context: ctx, Handler: handler}
}

/*GetSpoon swagger:route GET /v1/spoons/{id} spoons getSpoon

GetSpoon get spoon API

*/
type GetSpoon struct {
	Context *middleware.Context
	Handler GetSpoonHandler
}

func (o *GetSpoon) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetSpoonParams()

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
