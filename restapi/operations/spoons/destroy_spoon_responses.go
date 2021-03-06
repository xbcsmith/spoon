// Code generated by go-swagger; DO NOT EDIT.

package spoons

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/xbcsmith/spoon/models"
)

// DestroySpoonNoContentCode is the HTTP code returned for type DestroySpoonNoContent
const DestroySpoonNoContentCode int = 204

/*DestroySpoonNoContent Deleted

swagger:response destroySpoonNoContent
*/
type DestroySpoonNoContent struct {
}

// NewDestroySpoonNoContent creates DestroySpoonNoContent with default headers values
func NewDestroySpoonNoContent() *DestroySpoonNoContent {

	return &DestroySpoonNoContent{}
}

// WriteResponse to the client
func (o *DestroySpoonNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DestroySpoonUnauthorizedCode is the HTTP code returned for type DestroySpoonUnauthorized
const DestroySpoonUnauthorizedCode int = 401

/*DestroySpoonUnauthorized unauthorized

swagger:response destroySpoonUnauthorized
*/
type DestroySpoonUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDestroySpoonUnauthorized creates DestroySpoonUnauthorized with default headers values
func NewDestroySpoonUnauthorized() *DestroySpoonUnauthorized {

	return &DestroySpoonUnauthorized{}
}

// WithPayload adds the payload to the destroy spoon unauthorized response
func (o *DestroySpoonUnauthorized) WithPayload(payload *models.Error) *DestroySpoonUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the destroy spoon unauthorized response
func (o *DestroySpoonUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DestroySpoonUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DestroySpoonNotFoundCode is the HTTP code returned for type DestroySpoonNotFound
const DestroySpoonNotFoundCode int = 404

/*DestroySpoonNotFound resource not found

swagger:response destroySpoonNotFound
*/
type DestroySpoonNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDestroySpoonNotFound creates DestroySpoonNotFound with default headers values
func NewDestroySpoonNotFound() *DestroySpoonNotFound {

	return &DestroySpoonNotFound{}
}

// WithPayload adds the payload to the destroy spoon not found response
func (o *DestroySpoonNotFound) WithPayload(payload *models.Error) *DestroySpoonNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the destroy spoon not found response
func (o *DestroySpoonNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DestroySpoonNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*DestroySpoonDefault error

swagger:response destroySpoonDefault
*/
type DestroySpoonDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDestroySpoonDefault creates DestroySpoonDefault with default headers values
func NewDestroySpoonDefault(code int) *DestroySpoonDefault {
	if code <= 0 {
		code = 500
	}

	return &DestroySpoonDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the destroy spoon default response
func (o *DestroySpoonDefault) WithStatusCode(code int) *DestroySpoonDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the destroy spoon default response
func (o *DestroySpoonDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the destroy spoon default response
func (o *DestroySpoonDefault) WithPayload(payload *models.Error) *DestroySpoonDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the destroy spoon default response
func (o *DestroySpoonDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DestroySpoonDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
