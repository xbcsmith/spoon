// Code generated by go-swagger; DO NOT EDIT.

package spoons

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/xbcsmith/spoon/models"
)

// UpdateSpoonReader is a Reader for the UpdateSpoon structure.
type UpdateSpoonReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateSpoonReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateSpoonOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewUpdateSpoonUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUpdateSpoonNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewUpdateSpoonDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateSpoonOK creates a UpdateSpoonOK with default headers values
func NewUpdateSpoonOK() *UpdateSpoonOK {
	return &UpdateSpoonOK{}
}

/*UpdateSpoonOK handles this case with default header values.

OK
*/
type UpdateSpoonOK struct {
	Payload *models.Item
}

func (o *UpdateSpoonOK) Error() string {
	return fmt.Sprintf("[PUT /v1/spoons/{id}][%d] updateSpoonOK  %+v", 200, o.Payload)
}

func (o *UpdateSpoonOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Item)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSpoonUnauthorized creates a UpdateSpoonUnauthorized with default headers values
func NewUpdateSpoonUnauthorized() *UpdateSpoonUnauthorized {
	return &UpdateSpoonUnauthorized{}
}

/*UpdateSpoonUnauthorized handles this case with default header values.

unauthorized
*/
type UpdateSpoonUnauthorized struct {
	Payload *models.Error
}

func (o *UpdateSpoonUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/spoons/{id}][%d] updateSpoonUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateSpoonUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSpoonNotFound creates a UpdateSpoonNotFound with default headers values
func NewUpdateSpoonNotFound() *UpdateSpoonNotFound {
	return &UpdateSpoonNotFound{}
}

/*UpdateSpoonNotFound handles this case with default header values.

resource not found
*/
type UpdateSpoonNotFound struct {
	Payload *models.Error
}

func (o *UpdateSpoonNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/spoons/{id}][%d] updateSpoonNotFound  %+v", 404, o.Payload)
}

func (o *UpdateSpoonNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSpoonDefault creates a UpdateSpoonDefault with default headers values
func NewUpdateSpoonDefault(code int) *UpdateSpoonDefault {
	return &UpdateSpoonDefault{
		_statusCode: code,
	}
}

/*UpdateSpoonDefault handles this case with default header values.

error
*/
type UpdateSpoonDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the update spoon default response
func (o *UpdateSpoonDefault) Code() int {
	return o._statusCode
}

func (o *UpdateSpoonDefault) Error() string {
	return fmt.Sprintf("[PUT /v1/spoons/{id}][%d] update_spoon default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateSpoonDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
