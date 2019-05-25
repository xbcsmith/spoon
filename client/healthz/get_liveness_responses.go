// Code generated by go-swagger; DO NOT EDIT.

package healthz

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/xbcsmith/spoon/models"
)

// GetLivenessReader is a Reader for the GetLiveness structure.
type GetLivenessReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLivenessReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetLivenessOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetLivenessUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetLivenessNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetLivenessDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetLivenessOK creates a GetLivenessOK with default headers values
func NewGetLivenessOK() *GetLivenessOK {
	return &GetLivenessOK{}
}

/*GetLivenessOK handles this case with default header values.

return the liveness of app
*/
type GetLivenessOK struct {
	Payload *models.Liveness
}

func (o *GetLivenessOK) Error() string {
	return fmt.Sprintf("[GET /healthz/liveness][%d] getLivenessOK  %+v", 200, o.Payload)
}

func (o *GetLivenessOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Liveness)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLivenessUnauthorized creates a GetLivenessUnauthorized with default headers values
func NewGetLivenessUnauthorized() *GetLivenessUnauthorized {
	return &GetLivenessUnauthorized{}
}

/*GetLivenessUnauthorized handles this case with default header values.

unauthorized
*/
type GetLivenessUnauthorized struct {
	Payload *models.Error
}

func (o *GetLivenessUnauthorized) Error() string {
	return fmt.Sprintf("[GET /healthz/liveness][%d] getLivenessUnauthorized  %+v", 401, o.Payload)
}

func (o *GetLivenessUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLivenessNotFound creates a GetLivenessNotFound with default headers values
func NewGetLivenessNotFound() *GetLivenessNotFound {
	return &GetLivenessNotFound{}
}

/*GetLivenessNotFound handles this case with default header values.

resource not found
*/
type GetLivenessNotFound struct {
	Payload *models.Error
}

func (o *GetLivenessNotFound) Error() string {
	return fmt.Sprintf("[GET /healthz/liveness][%d] getLivenessNotFound  %+v", 404, o.Payload)
}

func (o *GetLivenessNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLivenessDefault creates a GetLivenessDefault with default headers values
func NewGetLivenessDefault(code int) *GetLivenessDefault {
	return &GetLivenessDefault{
		_statusCode: code,
	}
}

/*GetLivenessDefault handles this case with default header values.

generic error response
*/
type GetLivenessDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get liveness default response
func (o *GetLivenessDefault) Code() int {
	return o._statusCode
}

func (o *GetLivenessDefault) Error() string {
	return fmt.Sprintf("[GET /healthz/liveness][%d] get_liveness default  %+v", o._statusCode, o.Payload)
}

func (o *GetLivenessDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}