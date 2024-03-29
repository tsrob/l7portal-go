// Code generated by go-swagger; DO NOT EDIT.

package requests

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/tsrob/l7portal-go/models"
)

// GetRequestReader is a Reader for the GetRequest structure.
type GetRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRequestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetRequestNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRequestOK creates a GetRequestOK with default headers values
func NewGetRequestOK() *GetRequestOK {
	return &GetRequestOK{}
}

/*
GetRequestOK describes a response with status code 200, with default header values.

An object describing a single Request.
*/
type GetRequestOK struct {
	Payload *models.Request
}

// IsSuccess returns true when this get request o k response has a 2xx status code
func (o *GetRequestOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get request o k response has a 3xx status code
func (o *GetRequestOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get request o k response has a 4xx status code
func (o *GetRequestOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get request o k response has a 5xx status code
func (o *GetRequestOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get request o k response a status code equal to that given
func (o *GetRequestOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get request o k response
func (o *GetRequestOK) Code() int {
	return 200
}

func (o *GetRequestOK) Error() string {
	return fmt.Sprintf("[GET /api-management/1.0/requests/{uuid}][%d] getRequestOK  %+v", 200, o.Payload)
}

func (o *GetRequestOK) String() string {
	return fmt.Sprintf("[GET /api-management/1.0/requests/{uuid}][%d] getRequestOK  %+v", 200, o.Payload)
}

func (o *GetRequestOK) GetPayload() *models.Request {
	return o.Payload
}

func (o *GetRequestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Request)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRequestNotFound creates a GetRequestNotFound with default headers values
func NewGetRequestNotFound() *GetRequestNotFound {
	return &GetRequestNotFound{}
}

/*
GetRequestNotFound describes a response with status code 404, with default header values.

Entity not found
*/
type GetRequestNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get request not found response has a 2xx status code
func (o *GetRequestNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get request not found response has a 3xx status code
func (o *GetRequestNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get request not found response has a 4xx status code
func (o *GetRequestNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get request not found response has a 5xx status code
func (o *GetRequestNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get request not found response a status code equal to that given
func (o *GetRequestNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get request not found response
func (o *GetRequestNotFound) Code() int {
	return 404
}

func (o *GetRequestNotFound) Error() string {
	return fmt.Sprintf("[GET /api-management/1.0/requests/{uuid}][%d] getRequestNotFound  %+v", 404, o.Payload)
}

func (o *GetRequestNotFound) String() string {
	return fmt.Sprintf("[GET /api-management/1.0/requests/{uuid}][%d] getRequestNotFound  %+v", 404, o.Payload)
}

func (o *GetRequestNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetRequestNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
