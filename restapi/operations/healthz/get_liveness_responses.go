// Code generated by go-swagger; DO NOT EDIT.

package healthz

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/xbcsmith/spoon/models"
)

// GetLivenessOKCode is the HTTP code returned for type GetLivenessOK
const GetLivenessOKCode int = 200

/*GetLivenessOK return the liveness of app

swagger:response getLivenessOK
*/
type GetLivenessOK struct {

	/*
	  In: Body
	*/
	Payload *models.Liveness `json:"body,omitempty"`
}

// NewGetLivenessOK creates GetLivenessOK with default headers values
func NewGetLivenessOK() *GetLivenessOK {

	return &GetLivenessOK{}
}

// WithPayload adds the payload to the get liveness o k response
func (o *GetLivenessOK) WithPayload(payload *models.Liveness) *GetLivenessOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get liveness o k response
func (o *GetLivenessOK) SetPayload(payload *models.Liveness) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetLivenessOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetLivenessUnauthorizedCode is the HTTP code returned for type GetLivenessUnauthorized
const GetLivenessUnauthorizedCode int = 401

/*GetLivenessUnauthorized unauthorized

swagger:response getLivenessUnauthorized
*/
type GetLivenessUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetLivenessUnauthorized creates GetLivenessUnauthorized with default headers values
func NewGetLivenessUnauthorized() *GetLivenessUnauthorized {

	return &GetLivenessUnauthorized{}
}

// WithPayload adds the payload to the get liveness unauthorized response
func (o *GetLivenessUnauthorized) WithPayload(payload *models.Error) *GetLivenessUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get liveness unauthorized response
func (o *GetLivenessUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetLivenessUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetLivenessNotFoundCode is the HTTP code returned for type GetLivenessNotFound
const GetLivenessNotFoundCode int = 404

/*GetLivenessNotFound resource not found

swagger:response getLivenessNotFound
*/
type GetLivenessNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetLivenessNotFound creates GetLivenessNotFound with default headers values
func NewGetLivenessNotFound() *GetLivenessNotFound {

	return &GetLivenessNotFound{}
}

// WithPayload adds the payload to the get liveness not found response
func (o *GetLivenessNotFound) WithPayload(payload *models.Error) *GetLivenessNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get liveness not found response
func (o *GetLivenessNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetLivenessNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetLivenessDefault generic error response

swagger:response getLivenessDefault
*/
type GetLivenessDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetLivenessDefault creates GetLivenessDefault with default headers values
func NewGetLivenessDefault(code int) *GetLivenessDefault {
	if code <= 0 {
		code = 500
	}

	return &GetLivenessDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get liveness default response
func (o *GetLivenessDefault) WithStatusCode(code int) *GetLivenessDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get liveness default response
func (o *GetLivenessDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get liveness default response
func (o *GetLivenessDefault) WithPayload(payload *models.Error) *GetLivenessDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get liveness default response
func (o *GetLivenessDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetLivenessDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}