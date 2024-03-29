// Code generated by go-swagger; DO NOT EDIT.

package role_mappings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/tsrob/l7portal-go/models"
)

// GetRoleMappingsReader is a Reader for the GetRoleMappings structure.
type GetRoleMappingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRoleMappingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRoleMappingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRoleMappingsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRoleMappingsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRoleMappingsOK creates a GetRoleMappingsOK with default headers values
func NewGetRoleMappingsOK() *GetRoleMappingsOK {
	return &GetRoleMappingsOK{}
}

/*
GetRoleMappingsOK describes a response with status code 200, with default header values.

An object with total User-Org-Role association count and the list of all the User-Org-Role associations.
*/
type GetRoleMappingsOK struct {
	Payload *models.RoleMappingsGet
}

// IsSuccess returns true when this get role mappings o k response has a 2xx status code
func (o *GetRoleMappingsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get role mappings o k response has a 3xx status code
func (o *GetRoleMappingsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get role mappings o k response has a 4xx status code
func (o *GetRoleMappingsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get role mappings o k response has a 5xx status code
func (o *GetRoleMappingsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get role mappings o k response a status code equal to that given
func (o *GetRoleMappingsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get role mappings o k response
func (o *GetRoleMappingsOK) Code() int {
	return 200
}

func (o *GetRoleMappingsOK) Error() string {
	return fmt.Sprintf("[GET /v2/role-mappings][%d] getRoleMappingsOK  %+v", 200, o.Payload)
}

func (o *GetRoleMappingsOK) String() string {
	return fmt.Sprintf("[GET /v2/role-mappings][%d] getRoleMappingsOK  %+v", 200, o.Payload)
}

func (o *GetRoleMappingsOK) GetPayload() *models.RoleMappingsGet {
	return o.Payload
}

func (o *GetRoleMappingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RoleMappingsGet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoleMappingsBadRequest creates a GetRoleMappingsBadRequest with default headers values
func NewGetRoleMappingsBadRequest() *GetRoleMappingsBadRequest {
	return &GetRoleMappingsBadRequest{}
}

/*
GetRoleMappingsBadRequest describes a response with status code 400, with default header values.

Bad Request due to a validation error.
*/
type GetRoleMappingsBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get role mappings bad request response has a 2xx status code
func (o *GetRoleMappingsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get role mappings bad request response has a 3xx status code
func (o *GetRoleMappingsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get role mappings bad request response has a 4xx status code
func (o *GetRoleMappingsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get role mappings bad request response has a 5xx status code
func (o *GetRoleMappingsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get role mappings bad request response a status code equal to that given
func (o *GetRoleMappingsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get role mappings bad request response
func (o *GetRoleMappingsBadRequest) Code() int {
	return 400
}

func (o *GetRoleMappingsBadRequest) Error() string {
	return fmt.Sprintf("[GET /v2/role-mappings][%d] getRoleMappingsBadRequest  %+v", 400, o.Payload)
}

func (o *GetRoleMappingsBadRequest) String() string {
	return fmt.Sprintf("[GET /v2/role-mappings][%d] getRoleMappingsBadRequest  %+v", 400, o.Payload)
}

func (o *GetRoleMappingsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetRoleMappingsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoleMappingsInternalServerError creates a GetRoleMappingsInternalServerError with default headers values
func NewGetRoleMappingsInternalServerError() *GetRoleMappingsInternalServerError {
	return &GetRoleMappingsInternalServerError{}
}

/*
GetRoleMappingsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. Certain fields are missing from the request.
*/
type GetRoleMappingsInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get role mappings internal server error response has a 2xx status code
func (o *GetRoleMappingsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get role mappings internal server error response has a 3xx status code
func (o *GetRoleMappingsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get role mappings internal server error response has a 4xx status code
func (o *GetRoleMappingsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get role mappings internal server error response has a 5xx status code
func (o *GetRoleMappingsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get role mappings internal server error response a status code equal to that given
func (o *GetRoleMappingsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get role mappings internal server error response
func (o *GetRoleMappingsInternalServerError) Code() int {
	return 500
}

func (o *GetRoleMappingsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v2/role-mappings][%d] getRoleMappingsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRoleMappingsInternalServerError) String() string {
	return fmt.Sprintf("[GET /v2/role-mappings][%d] getRoleMappingsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRoleMappingsInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetRoleMappingsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
