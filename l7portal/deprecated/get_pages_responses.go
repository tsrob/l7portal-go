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

// GetPagesReader is a Reader for the GetPages structure.
type GetPagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetPagesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPagesOK creates a GetPagesOK with default headers values
func NewGetPagesOK() *GetPagesOK {
	return &GetPagesOK{}
}

/*
GetPagesOK describes a response with status code 200, with default header values.

GetPagesOK get pages o k
*/
type GetPagesOK struct {
	Payload models.CustomPageGet
}

// IsSuccess returns true when this get pages o k response has a 2xx status code
func (o *GetPagesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get pages o k response has a 3xx status code
func (o *GetPagesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get pages o k response has a 4xx status code
func (o *GetPagesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get pages o k response has a 5xx status code
func (o *GetPagesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get pages o k response a status code equal to that given
func (o *GetPagesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get pages o k response
func (o *GetPagesOK) Code() int {
	return 200
}

func (o *GetPagesOK) Error() string {
	return fmt.Sprintf("[GET /custom/1.0/pages][%d] getPagesOK  %+v", 200, o.Payload)
}

func (o *GetPagesOK) String() string {
	return fmt.Sprintf("[GET /custom/1.0/pages][%d] getPagesOK  %+v", 200, o.Payload)
}

func (o *GetPagesOK) GetPayload() models.CustomPageGet {
	return o.Payload
}

func (o *GetPagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPagesInternalServerError creates a GetPagesInternalServerError with default headers values
func NewGetPagesInternalServerError() *GetPagesInternalServerError {
	return &GetPagesInternalServerError{}
}

/*
GetPagesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. Certain fields are missing from the request.
*/
type GetPagesInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get pages internal server error response has a 2xx status code
func (o *GetPagesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get pages internal server error response has a 3xx status code
func (o *GetPagesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get pages internal server error response has a 4xx status code
func (o *GetPagesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get pages internal server error response has a 5xx status code
func (o *GetPagesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get pages internal server error response a status code equal to that given
func (o *GetPagesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get pages internal server error response
func (o *GetPagesInternalServerError) Code() int {
	return 500
}

func (o *GetPagesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /custom/1.0/pages][%d] getPagesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetPagesInternalServerError) String() string {
	return fmt.Sprintf("[GET /custom/1.0/pages][%d] getPagesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetPagesInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetPagesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}