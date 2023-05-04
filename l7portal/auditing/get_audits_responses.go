// Code generated by go-swagger; DO NOT EDIT.

package auditing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/tsrob/l7portal-go/models"
)

// GetAuditsReader is a Reader for the GetAudits structure.
type GetAuditsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAuditsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAuditsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetAuditsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAuditsOK creates a GetAuditsOK with default headers values
func NewGetAuditsOK() *GetAuditsOK {
	return &GetAuditsOK{}
}

/*
GetAuditsOK describes a response with status code 200, with default header values.

GetAuditsOK get audits o k
*/
type GetAuditsOK struct {
	Payload *models.AuditsGet
}

// IsSuccess returns true when this get audits o k response has a 2xx status code
func (o *GetAuditsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get audits o k response has a 3xx status code
func (o *GetAuditsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get audits o k response has a 4xx status code
func (o *GetAuditsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get audits o k response has a 5xx status code
func (o *GetAuditsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get audits o k response a status code equal to that given
func (o *GetAuditsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get audits o k response
func (o *GetAuditsOK) Code() int {
	return 200
}

func (o *GetAuditsOK) Error() string {
	return fmt.Sprintf("[GET /tenant-admin/1.0/audits][%d] getAuditsOK  %+v", 200, o.Payload)
}

func (o *GetAuditsOK) String() string {
	return fmt.Sprintf("[GET /tenant-admin/1.0/audits][%d] getAuditsOK  %+v", 200, o.Payload)
}

func (o *GetAuditsOK) GetPayload() *models.AuditsGet {
	return o.Payload
}

func (o *GetAuditsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuditsGet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuditsInternalServerError creates a GetAuditsInternalServerError with default headers values
func NewGetAuditsInternalServerError() *GetAuditsInternalServerError {
	return &GetAuditsInternalServerError{}
}

/*
GetAuditsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. Certain fields are missing from the request.
*/
type GetAuditsInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get audits internal server error response has a 2xx status code
func (o *GetAuditsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get audits internal server error response has a 3xx status code
func (o *GetAuditsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get audits internal server error response has a 4xx status code
func (o *GetAuditsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get audits internal server error response has a 5xx status code
func (o *GetAuditsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get audits internal server error response a status code equal to that given
func (o *GetAuditsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get audits internal server error response
func (o *GetAuditsInternalServerError) Code() int {
	return 500
}

func (o *GetAuditsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /tenant-admin/1.0/audits][%d] getAuditsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAuditsInternalServerError) String() string {
	return fmt.Sprintf("[GET /tenant-admin/1.0/audits][%d] getAuditsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAuditsInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAuditsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
