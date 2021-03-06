// Code generated by go-swagger; DO NOT EDIT.

package spoons

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetSpoonParams creates a new GetSpoonParams object
// with the default values initialized.
func NewGetSpoonParams() *GetSpoonParams {
	var ()
	return &GetSpoonParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSpoonParamsWithTimeout creates a new GetSpoonParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSpoonParamsWithTimeout(timeout time.Duration) *GetSpoonParams {
	var ()
	return &GetSpoonParams{

		timeout: timeout,
	}
}

// NewGetSpoonParamsWithContext creates a new GetSpoonParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSpoonParamsWithContext(ctx context.Context) *GetSpoonParams {
	var ()
	return &GetSpoonParams{

		Context: ctx,
	}
}

// NewGetSpoonParamsWithHTTPClient creates a new GetSpoonParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSpoonParamsWithHTTPClient(client *http.Client) *GetSpoonParams {
	var ()
	return &GetSpoonParams{
		HTTPClient: client,
	}
}

/*GetSpoonParams contains all the parameters to send to the API endpoint
for the get spoon operation typically these are written to a http.Request
*/
type GetSpoonParams struct {

	/*ID*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get spoon params
func (o *GetSpoonParams) WithTimeout(timeout time.Duration) *GetSpoonParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get spoon params
func (o *GetSpoonParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get spoon params
func (o *GetSpoonParams) WithContext(ctx context.Context) *GetSpoonParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get spoon params
func (o *GetSpoonParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get spoon params
func (o *GetSpoonParams) WithHTTPClient(client *http.Client) *GetSpoonParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get spoon params
func (o *GetSpoonParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get spoon params
func (o *GetSpoonParams) WithID(id int64) *GetSpoonParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get spoon params
func (o *GetSpoonParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetSpoonParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
