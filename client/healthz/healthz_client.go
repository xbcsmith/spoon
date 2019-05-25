// Code generated by go-swagger; DO NOT EDIT.

package healthz

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new healthz API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for healthz API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetLiveness get liveness API
*/
func (a *Client) GetLiveness(params *GetLivenessParams, authInfo runtime.ClientAuthInfoWriter) (*GetLivenessOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLivenessParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_liveness",
		Method:             "GET",
		PathPattern:        "/healthz/liveness",
		ProducesMediaTypes: []string{"application/com.github.xbcsmith.spoon.v1+json"},
		ConsumesMediaTypes: []string{"application/com.github.xbcsmith.spoon.v1+json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetLivenessReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetLivenessOK), nil

}

/*
GetReadiness get readiness API
*/
func (a *Client) GetReadiness(params *GetReadinessParams, authInfo runtime.ClientAuthInfoWriter) (*GetReadinessOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetReadinessParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_readiness",
		Method:             "GET",
		PathPattern:        "/healthz/readiness",
		ProducesMediaTypes: []string{"application/com.github.xbcsmith.spoon.v1+json"},
		ConsumesMediaTypes: []string{"application/com.github.xbcsmith.spoon.v1+json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetReadinessReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetReadinessOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}