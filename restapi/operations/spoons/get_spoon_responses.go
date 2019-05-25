// Code generated by go-swagger; DO NOT EDIT.

package spoons

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/xbcsmith/spoon/models"
)

// GetSpoonOKCode is the HTTP code returned for type GetSpoonOK
const GetSpoonOKCode int = 200

/*GetSpoonOK get a spoon operation

swagger:response getSpoonOK
*/
type GetSpoonOK struct {

	/*
	  In: Body
	*/
	Payload *models.Item `json:"body,omitempty"`
}

// NewGetSpoonOK creates GetSpoonOK with default headers values
func NewGetSpoonOK() *GetSpoonOK {

	return &GetSpoonOK{}
}

// WithPayload adds the payload to the get spoon o k response
func (o *GetSpoonOK) WithPayload(payload *models.Item) *GetSpoonOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get spoon o k response
func (o *GetSpoonOK) SetPayload(payload *models.Item) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSpoonOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetSpoonUnauthorizedCode is the HTTP code returned for type GetSpoonUnauthorized
const GetSpoonUnauthorizedCode int = 401

/*GetSpoonUnauthorized unauthorized

swagger:response getSpoonUnauthorized
*/
type GetSpoonUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetSpoonUnauthorized creates GetSpoonUnauthorized with default headers values
func NewGetSpoonUnauthorized() *GetSpoonUnauthorized {

	return &GetSpoonUnauthorized{}
}

// WithPayload adds the payload to the get spoon unauthorized response
func (o *GetSpoonUnauthorized) WithPayload(payload *models.Error) *GetSpoonUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get spoon unauthorized response
func (o *GetSpoonUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSpoonUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetSpoonNotFoundCode is the HTTP code returned for type GetSpoonNotFound
const GetSpoonNotFoundCode int = 404

/*GetSpoonNotFound resource not found

swagger:response getSpoonNotFound
*/
type GetSpoonNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetSpoonNotFound creates GetSpoonNotFound with default headers values
func NewGetSpoonNotFound() *GetSpoonNotFound {

	return &GetSpoonNotFound{}
}

// WithPayload adds the payload to the get spoon not found response
func (o *GetSpoonNotFound) WithPayload(payload *models.Error) *GetSpoonNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get spoon not found response
func (o *GetSpoonNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSpoonNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetSpoonDefault generic error response

swagger:response getSpoonDefault
*/
type GetSpoonDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetSpoonDefault creates GetSpoonDefault with default headers values
func NewGetSpoonDefault(code int) *GetSpoonDefault {
	if code <= 0 {
		code = 500
	}

	return &GetSpoonDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get spoon default response
func (o *GetSpoonDefault) WithStatusCode(code int) *GetSpoonDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get spoon default response
func (o *GetSpoonDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get spoon default response
func (o *GetSpoonDefault) WithPayload(payload *models.Error) *GetSpoonDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get spoon default response
func (o *GetSpoonDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSpoonDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}