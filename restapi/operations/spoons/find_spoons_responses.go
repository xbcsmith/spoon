// Code generated by go-swagger; DO NOT EDIT.

package spoons

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/xbcsmith/spoon/models"
)

// FindSpoonsOKCode is the HTTP code returned for type FindSpoonsOK
const FindSpoonsOKCode int = 200

/*FindSpoonsOK list the spoon operations

swagger:response findSpoonsOK
*/
type FindSpoonsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Item `json:"body,omitempty"`
}

// NewFindSpoonsOK creates FindSpoonsOK with default headers values
func NewFindSpoonsOK() *FindSpoonsOK {

	return &FindSpoonsOK{}
}

// WithPayload adds the payload to the find spoons o k response
func (o *FindSpoonsOK) WithPayload(payload []*models.Item) *FindSpoonsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the find spoons o k response
func (o *FindSpoonsOK) SetPayload(payload []*models.Item) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FindSpoonsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Item, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// FindSpoonsUnauthorizedCode is the HTTP code returned for type FindSpoonsUnauthorized
const FindSpoonsUnauthorizedCode int = 401

/*FindSpoonsUnauthorized unauthorized

swagger:response findSpoonsUnauthorized
*/
type FindSpoonsUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewFindSpoonsUnauthorized creates FindSpoonsUnauthorized with default headers values
func NewFindSpoonsUnauthorized() *FindSpoonsUnauthorized {

	return &FindSpoonsUnauthorized{}
}

// WithPayload adds the payload to the find spoons unauthorized response
func (o *FindSpoonsUnauthorized) WithPayload(payload *models.Error) *FindSpoonsUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the find spoons unauthorized response
func (o *FindSpoonsUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FindSpoonsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// FindSpoonsNotFoundCode is the HTTP code returned for type FindSpoonsNotFound
const FindSpoonsNotFoundCode int = 404

/*FindSpoonsNotFound resource not found

swagger:response findSpoonsNotFound
*/
type FindSpoonsNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewFindSpoonsNotFound creates FindSpoonsNotFound with default headers values
func NewFindSpoonsNotFound() *FindSpoonsNotFound {

	return &FindSpoonsNotFound{}
}

// WithPayload adds the payload to the find spoons not found response
func (o *FindSpoonsNotFound) WithPayload(payload *models.Error) *FindSpoonsNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the find spoons not found response
func (o *FindSpoonsNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FindSpoonsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*FindSpoonsDefault generic error response

swagger:response findSpoonsDefault
*/
type FindSpoonsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewFindSpoonsDefault creates FindSpoonsDefault with default headers values
func NewFindSpoonsDefault(code int) *FindSpoonsDefault {
	if code <= 0 {
		code = 500
	}

	return &FindSpoonsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the find spoons default response
func (o *FindSpoonsDefault) WithStatusCode(code int) *FindSpoonsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the find spoons default response
func (o *FindSpoonsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the find spoons default response
func (o *FindSpoonsDefault) WithPayload(payload *models.Error) *FindSpoonsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the find spoons default response
func (o *FindSpoonsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FindSpoonsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}