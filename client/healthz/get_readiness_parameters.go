// Code generated by go-swagger; DO NOT EDIT.

package healthz

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetReadinessParams creates a new GetReadinessParams object
// with the default values initialized.
func NewGetReadinessParams() *GetReadinessParams {

	return &GetReadinessParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetReadinessParamsWithTimeout creates a new GetReadinessParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetReadinessParamsWithTimeout(timeout time.Duration) *GetReadinessParams {

	return &GetReadinessParams{

		timeout: timeout,
	}
}

// NewGetReadinessParamsWithContext creates a new GetReadinessParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetReadinessParamsWithContext(ctx context.Context) *GetReadinessParams {

	return &GetReadinessParams{

		Context: ctx,
	}
}

// NewGetReadinessParamsWithHTTPClient creates a new GetReadinessParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetReadinessParamsWithHTTPClient(client *http.Client) *GetReadinessParams {

	return &GetReadinessParams{
		HTTPClient: client,
	}
}

/*GetReadinessParams contains all the parameters to send to the API endpoint
for the get readiness operation typically these are written to a http.Request
*/
type GetReadinessParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get readiness params
func (o *GetReadinessParams) WithTimeout(timeout time.Duration) *GetReadinessParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get readiness params
func (o *GetReadinessParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get readiness params
func (o *GetReadinessParams) WithContext(ctx context.Context) *GetReadinessParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get readiness params
func (o *GetReadinessParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get readiness params
func (o *GetReadinessParams) WithHTTPClient(client *http.Client) *GetReadinessParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get readiness params
func (o *GetReadinessParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetReadinessParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
