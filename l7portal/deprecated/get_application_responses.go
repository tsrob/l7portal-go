// Code generated by go-swagger; DO NOT EDIT.

package deprecated

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/tsrob/l7portal-go/models"
)

// GetApplicationReader is a Reader for the GetApplication structure.
type GetApplicationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetApplicationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApplicationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetApplicationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetApplicationOK creates a GetApplicationOK with default headers values
func NewGetApplicationOK() *GetApplicationOK {
	return &GetApplicationOK{}
}

/*
GetApplicationOK describes a response with status code 200, with default header values.

An object describing a single Application.
*/
type GetApplicationOK struct {
	Payload *models.Application
}

// IsSuccess returns true when this get application o k response has a 2xx status code
func (o *GetApplicationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get application o k response has a 3xx status code
func (o *GetApplicationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get application o k response has a 4xx status code
func (o *GetApplicationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get application o k response has a 5xx status code
func (o *GetApplicationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get application o k response a status code equal to that given
func (o *GetApplicationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get application o k response
func (o *GetApplicationOK) Code() int {
	return 200
}

func (o *GetApplicationOK) Error() string {
	return fmt.Sprintf("[GET /Applications('{uuid}')][%d] getApplicationOK  %+v", 200, o.Payload)
}

func (o *GetApplicationOK) String() string {
	return fmt.Sprintf("[GET /Applications('{uuid}')][%d] getApplicationOK  %+v", 200, o.Payload)
}

func (o *GetApplicationOK) GetPayload() *models.Application {
	return o.Payload
}

func (o *GetApplicationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Application)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApplicationNotFound creates a GetApplicationNotFound with default headers values
func NewGetApplicationNotFound() *GetApplicationNotFound {
	return &GetApplicationNotFound{}
}

/*
GetApplicationNotFound describes a response with status code 404, with default header values.

Entity not found
*/
type GetApplicationNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get application not found response has a 2xx status code
func (o *GetApplicationNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get application not found response has a 3xx status code
func (o *GetApplicationNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get application not found response has a 4xx status code
func (o *GetApplicationNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get application not found response has a 5xx status code
func (o *GetApplicationNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get application not found response a status code equal to that given
func (o *GetApplicationNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get application not found response
func (o *GetApplicationNotFound) Code() int {
	return 404
}

func (o *GetApplicationNotFound) Error() string {
	return fmt.Sprintf("[GET /Applications('{uuid}')][%d] getApplicationNotFound  %+v", 404, o.Payload)
}

func (o *GetApplicationNotFound) String() string {
	return fmt.Sprintf("[GET /Applications('{uuid}')][%d] getApplicationNotFound  %+v", 404, o.Payload)
}

func (o *GetApplicationNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetApplicationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}