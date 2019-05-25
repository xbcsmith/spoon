// Code generated by go-swagger; DO NOT EDIT.

package spoons

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/xbcsmith/spoon/models"
)

// AddSpoonCreatedCode is the HTTP code returned for type AddSpoonCreated
const AddSpoonCreatedCode int = 201

/*AddSpoonCreated Created

swagger:response addSpoonCreated
*/
type AddSpoonCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Item `json:"body,omitempty"`
}

// NewAddSpoonCreated creates AddSpoonCreated with default headers values
func NewAddSpoonCreated() *AddSpoonCreated {

	return &AddSpoonCreated{}
}

// WithPayload adds the payload to the add spoon created response
func (o *AddSpoonCreated) WithPayload(payload *models.Item) *AddSpoonCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add spoon created response
func (o *AddSpoonCreated) SetPayload(payload *models.Item) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddSpoonCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddSpoonUnauthorizedCode is the HTTP code returned for type AddSpoonUnauthorized
const AddSpoonUnauthorizedCode int = 401

/*AddSpoonUnauthorized unauthorized

swagger:response addSpoonUnauthorized
*/
type AddSpoonUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAddSpoonUnauthorized creates AddSpoonUnauthorized with default headers values
func NewAddSpoonUnauthorized() *AddSpoonUnauthorized {

	return &AddSpoonUnauthorized{}
}

// WithPayload adds the payload to the add spoon unauthorized response
func (o *AddSpoonUnauthorized) WithPayload(payload *models.Error) *AddSpoonUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add spoon unauthorized response
func (o *AddSpoonUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddSpoonUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddSpoonNotFoundCode is the HTTP code returned for type AddSpoonNotFound
const AddSpoonNotFoundCode int = 404

/*AddSpoonNotFound resource not found

swagger:response addSpoonNotFound
*/
type AddSpoonNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAddSpoonNotFound creates AddSpoonNotFound with default headers values
func NewAddSpoonNotFound() *AddSpoonNotFound {

	return &AddSpoonNotFound{}
}

// WithPayload adds the payload to the add spoon not found response
func (o *AddSpoonNotFound) WithPayload(payload *models.Error) *AddSpoonNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add spoon not found response
func (o *AddSpoonNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddSpoonNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*AddSpoonDefault error

swagger:response addSpoonDefault
*/
type AddSpoonDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAddSpoonDefault creates AddSpoonDefault with default headers values
func NewAddSpoonDefault(code int) *AddSpoonDefault {
	if code <= 0 {
		code = 500
	}

	return &AddSpoonDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the add spoon default response
func (o *AddSpoonDefault) WithStatusCode(code int) *AddSpoonDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the add spoon default response
func (o *AddSpoonDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the add spoon default response
func (o *AddSpoonDefault) WithPayload(payload *models.Error) *AddSpoonDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add spoon default response
func (o *AddSpoonDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddSpoonDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}