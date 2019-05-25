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

// GetReadinessReader is a Reader for the GetReadiness structure.
type GetReadinessReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReadinessReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetReadinessOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetReadinessUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetReadinessNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetReadinessDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetReadinessOK creates a GetReadinessOK with default headers values
func NewGetReadinessOK() *GetReadinessOK {
	return &GetReadinessOK{}
}

/*GetReadinessOK handles this case with default header values.

return the readiness of app
*/
type GetReadinessOK struct {
	Payload *models.Readiness
}

func (o *GetReadinessOK) Error() string {
	return fmt.Sprintf("[GET /healthz/readiness][%d] getReadinessOK  %+v", 200, o.Payload)
}

func (o *GetReadinessOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Readiness)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReadinessUnauthorized creates a GetReadinessUnauthorized with default headers values
func NewGetReadinessUnauthorized() *GetReadinessUnauthorized {
	return &GetReadinessUnauthorized{}
}

/*GetReadinessUnauthorized handles this case with default header values.

unauthorized
*/
type GetReadinessUnauthorized struct {
	Payload *models.Error
}

func (o *GetReadinessUnauthorized) Error() string {
	return fmt.Sprintf("[GET /healthz/readiness][%d] getReadinessUnauthorized  %+v", 401, o.Payload)
}

func (o *GetReadinessUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReadinessNotFound creates a GetReadinessNotFound with default headers values
func NewGetReadinessNotFound() *GetReadinessNotFound {
	return &GetReadinessNotFound{}
}

/*GetReadinessNotFound handles this case with default header values.

resource not found
*/
type GetReadinessNotFound struct {
	Payload *models.Error
}

func (o *GetReadinessNotFound) Error() string {
	return fmt.Sprintf("[GET /healthz/readiness][%d] getReadinessNotFound  %+v", 404, o.Payload)
}

func (o *GetReadinessNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReadinessDefault creates a GetReadinessDefault with default headers values
func NewGetReadinessDefault(code int) *GetReadinessDefault {
	return &GetReadinessDefault{
		_statusCode: code,
	}
}

/*GetReadinessDefault handles this case with default header values.

generic error response
*/
type GetReadinessDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get readiness default response
func (o *GetReadinessDefault) Code() int {
	return o._statusCode
}

func (o *GetReadinessDefault) Error() string {
	return fmt.Sprintf("[GET /healthz/readiness][%d] get_readiness default  %+v", o._statusCode, o.Payload)
}

func (o *GetReadinessDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
