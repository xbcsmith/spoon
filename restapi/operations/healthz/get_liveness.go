// Code generated by go-swagger; DO NOT EDIT.

package healthz

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"

	models "github.com/xbcsmith/spoon/models"
)

// GetLivenessHandlerFunc turns a function with the right signature into a get liveness handler
type GetLivenessHandlerFunc func(GetLivenessParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn GetLivenessHandlerFunc) Handle(params GetLivenessParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// GetLivenessHandler interface for that can handle valid get liveness params
type GetLivenessHandler interface {
	Handle(GetLivenessParams, *models.Principal) middleware.Responder
}

// NewGetLiveness creates a new http.Handler for the get liveness operation
func NewGetLiveness(ctx *middleware.Context, handler GetLivenessHandler) *GetLiveness {
	return &GetLiveness{Context: ctx, Handler: handler}
}

/*GetLiveness swagger:route GET /healthz/liveness healthz getLiveness

GetLiveness get liveness API

*/
type GetLiveness struct {
	Context *middleware.Context
	Handler GetLivenessHandler
}

func (o *GetLiveness) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetLivenessParams()

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
