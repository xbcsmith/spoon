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

// AddSpoonReader is a Reader for the AddSpoon structure.
type AddSpoonReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddSpoonReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewAddSpoonCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewAddSpoonUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewAddSpoonNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewAddSpoonDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddSpoonCreated creates a AddSpoonCreated with default headers values
func NewAddSpoonCreated() *AddSpoonCreated {
	return &AddSpoonCreated{}
}

/*AddSpoonCreated handles this case with default header values.

Created
*/
type AddSpoonCreated struct {
	Payload *models.Item
}

func (o *AddSpoonCreated) Error() string {
	return fmt.Sprintf("[POST /v1/spoons][%d] addSpoonCreated  %+v", 201, o.Payload)
}

func (o *AddSpoonCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Item)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddSpoonUnauthorized creates a AddSpoonUnauthorized with default headers values
func NewAddSpoonUnauthorized() *AddSpoonUnauthorized {
	return &AddSpoonUnauthorized{}
}

/*AddSpoonUnauthorized handles this case with default header values.

unauthorized
*/
type AddSpoonUnauthorized struct {
	Payload *models.Error
}

func (o *AddSpoonUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/spoons][%d] addSpoonUnauthorized  %+v", 401, o.Payload)
}

func (o *AddSpoonUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddSpoonNotFound creates a AddSpoonNotFound with default headers values
func NewAddSpoonNotFound() *AddSpoonNotFound {
	return &AddSpoonNotFound{}
}

/*AddSpoonNotFound handles this case with default header values.

resource not found
*/
type AddSpoonNotFound struct {
	Payload *models.Error
}

func (o *AddSpoonNotFound) Error() string {
	return fmt.Sprintf("[POST /v1/spoons][%d] addSpoonNotFound  %+v", 404, o.Payload)
}

func (o *AddSpoonNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddSpoonDefault creates a AddSpoonDefault with default headers values
func NewAddSpoonDefault(code int) *AddSpoonDefault {
	return &AddSpoonDefault{
		_statusCode: code,
	}
}

/*AddSpoonDefault handles this case with default header values.

error
*/
type AddSpoonDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the add spoon default response
func (o *AddSpoonDefault) Code() int {
	return o._statusCode
}

func (o *AddSpoonDefault) Error() string {
	return fmt.Sprintf("[POST /v1/spoons][%d] add_spoon default  %+v", o._statusCode, o.Payload)
}

func (o *AddSpoonDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
