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

// GetRoleMappingReader is a Reader for the GetRoleMapping structure.
type GetRoleMappingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRoleMappingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRoleMappingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetRoleMappingNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRoleMappingInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRoleMappingOK creates a GetRoleMappingOK with default headers values
func NewGetRoleMappingOK() *GetRoleMappingOK {
	return &GetRoleMappingOK{}
}

/*
GetRoleMappingOK describes a response with status code 200, with default header values.

An object describing a User-Org-Role association.
*/
type GetRoleMappingOK struct {
	Payload *models.RoleMapping
}

// IsSuccess returns true when this get role mapping o k response has a 2xx status code
func (o *GetRoleMappingOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get role mapping o k response has a 3xx status code
func (o *GetRoleMappingOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get role mapping o k response has a 4xx status code
func (o *GetRoleMappingOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get role mapping o k response has a 5xx status code
func (o *GetRoleMappingOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get role mapping o k response a status code equal to that given
func (o *GetRoleMappingOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get role mapping o k response
func (o *GetRoleMappingOK) Code() int {
	return 200
}

func (o *GetRoleMappingOK) Error() string {
	return fmt.Sprintf("[GET /v2/role-mappings/{roleMappingUuid}][%d] getRoleMappingOK  %+v", 200, o.Payload)
}

func (o *GetRoleMappingOK) String() string {
	return fmt.Sprintf("[GET /v2/role-mappings/{roleMappingUuid}][%d] getRoleMappingOK  %+v", 200, o.Payload)
}

func (o *GetRoleMappingOK) GetPayload() *models.RoleMapping {
	return o.Payload
}

func (o *GetRoleMappingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RoleMapping)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoleMappingNotFound creates a GetRoleMappingNotFound with default headers values
func NewGetRoleMappingNotFound() *GetRoleMappingNotFound {
	return &GetRoleMappingNotFound{}
}

/*
GetRoleMappingNotFound describes a response with status code 404, with default header values.

Entity not found
*/
type GetRoleMappingNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get role mapping not found response has a 2xx status code
func (o *GetRoleMappingNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get role mapping not found response has a 3xx status code
func (o *GetRoleMappingNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get role mapping not found response has a 4xx status code
func (o *GetRoleMappingNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get role mapping not found response has a 5xx status code
func (o *GetRoleMappingNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get role mapping not found response a status code equal to that given
func (o *GetRoleMappingNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get role mapping not found response
func (o *GetRoleMappingNotFound) Code() int {
	return 404
}

func (o *GetRoleMappingNotFound) Error() string {
	return fmt.Sprintf("[GET /v2/role-mappings/{roleMappingUuid}][%d] getRoleMappingNotFound  %+v", 404, o.Payload)
}

func (o *GetRoleMappingNotFound) String() string {
	return fmt.Sprintf("[GET /v2/role-mappings/{roleMappingUuid}][%d] getRoleMappingNotFound  %+v", 404, o.Payload)
}

func (o *GetRoleMappingNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetRoleMappingNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoleMappingInternalServerError creates a GetRoleMappingInternalServerError with default headers values
func NewGetRoleMappingInternalServerError() *GetRoleMappingInternalServerError {
	return &GetRoleMappingInternalServerError{}
}

/*
GetRoleMappingInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. Certain fields are missing from the request.
*/
type GetRoleMappingInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get role mapping internal server error response has a 2xx status code
func (o *GetRoleMappingInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get role mapping internal server error response has a 3xx status code
func (o *GetRoleMappingInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get role mapping internal server error response has a 4xx status code
func (o *GetRoleMappingInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get role mapping internal server error response has a 5xx status code
func (o *GetRoleMappingInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get role mapping internal server error response a status code equal to that given
func (o *GetRoleMappingInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get role mapping internal server error response
func (o *GetRoleMappingInternalServerError) Code() int {
	return 500
}

func (o *GetRoleMappingInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v2/role-mappings/{roleMappingUuid}][%d] getRoleMappingInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRoleMappingInternalServerError) String() string {
	return fmt.Sprintf("[GET /v2/role-mappings/{roleMappingUuid}][%d] getRoleMappingInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRoleMappingInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetRoleMappingInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
