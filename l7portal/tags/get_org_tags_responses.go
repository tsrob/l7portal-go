// Code generated by go-swagger; DO NOT EDIT.

package tags

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/tsrob/l7portal-go/models"
)

// GetOrgTagsReader is a Reader for the GetOrgTags structure.
type GetOrgTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrgTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrgTagsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOrgTagsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOrgTagsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOrgTagsOK creates a GetOrgTagsOK with default headers values
func NewGetOrgTagsOK() *GetOrgTagsOK {
	return &GetOrgTagsOK{}
}

/*
GetOrgTagsOK describes a response with status code 200, with default header values.

An array containing all tags that can be associated with an Organization.
*/
type GetOrgTagsOK struct {
	Payload *models.TenantsTagGet
}

// IsSuccess returns true when this get org tags o k response has a 2xx status code
func (o *GetOrgTagsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get org tags o k response has a 3xx status code
func (o *GetOrgTagsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get org tags o k response has a 4xx status code
func (o *GetOrgTagsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get org tags o k response has a 5xx status code
func (o *GetOrgTagsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get org tags o k response a status code equal to that given
func (o *GetOrgTagsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get org tags o k response
func (o *GetOrgTagsOK) Code() int {
	return 200
}

func (o *GetOrgTagsOK) Error() string {
	return fmt.Sprintf("[GET /tenant-admin/1.0/tags][%d] getOrgTagsOK  %+v", 200, o.Payload)
}

func (o *GetOrgTagsOK) String() string {
	return fmt.Sprintf("[GET /tenant-admin/1.0/tags][%d] getOrgTagsOK  %+v", 200, o.Payload)
}

func (o *GetOrgTagsOK) GetPayload() *models.TenantsTagGet {
	return o.Payload
}

func (o *GetOrgTagsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TenantsTagGet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgTagsBadRequest creates a GetOrgTagsBadRequest with default headers values
func NewGetOrgTagsBadRequest() *GetOrgTagsBadRequest {
	return &GetOrgTagsBadRequest{}
}

/*
GetOrgTagsBadRequest describes a response with status code 400, with default header values.

Bad Request due to Invalid Parameters.
*/
type GetOrgTagsBadRequest struct {
	Payload *models.LongerError
}

// IsSuccess returns true when this get org tags bad request response has a 2xx status code
func (o *GetOrgTagsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get org tags bad request response has a 3xx status code
func (o *GetOrgTagsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get org tags bad request response has a 4xx status code
func (o *GetOrgTagsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get org tags bad request response has a 5xx status code
func (o *GetOrgTagsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get org tags bad request response a status code equal to that given
func (o *GetOrgTagsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get org tags bad request response
func (o *GetOrgTagsBadRequest) Code() int {
	return 400
}

func (o *GetOrgTagsBadRequest) Error() string {
	return fmt.Sprintf("[GET /tenant-admin/1.0/tags][%d] getOrgTagsBadRequest  %+v", 400, o.Payload)
}

func (o *GetOrgTagsBadRequest) String() string {
	return fmt.Sprintf("[GET /tenant-admin/1.0/tags][%d] getOrgTagsBadRequest  %+v", 400, o.Payload)
}

func (o *GetOrgTagsBadRequest) GetPayload() *models.LongerError {
	return o.Payload
}

func (o *GetOrgTagsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LongerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgTagsInternalServerError creates a GetOrgTagsInternalServerError with default headers values
func NewGetOrgTagsInternalServerError() *GetOrgTagsInternalServerError {
	return &GetOrgTagsInternalServerError{}
}

/*
GetOrgTagsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. Certain fields are missing from the request.
*/
type GetOrgTagsInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get org tags internal server error response has a 2xx status code
func (o *GetOrgTagsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get org tags internal server error response has a 3xx status code
func (o *GetOrgTagsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get org tags internal server error response has a 4xx status code
func (o *GetOrgTagsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get org tags internal server error response has a 5xx status code
func (o *GetOrgTagsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get org tags internal server error response a status code equal to that given
func (o *GetOrgTagsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get org tags internal server error response
func (o *GetOrgTagsInternalServerError) Code() int {
	return 500
}

func (o *GetOrgTagsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /tenant-admin/1.0/tags][%d] getOrgTagsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOrgTagsInternalServerError) String() string {
	return fmt.Sprintf("[GET /tenant-admin/1.0/tags][%d] getOrgTagsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOrgTagsInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetOrgTagsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
