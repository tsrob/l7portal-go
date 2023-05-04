// Code generated by go-swagger; DO NOT EDIT.

package api_organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PATCHAPIOrganizationsReader is a Reader for the PATCHAPIOrganizations structure.
type PATCHAPIOrganizationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PATCHAPIOrganizationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPATCHAPIOrganizationsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPATCHAPIOrganizationsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPATCHAPIOrganizationsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPATCHAPIOrganizationsNoContent creates a PATCHAPIOrganizationsNoContent with default headers values
func NewPATCHAPIOrganizationsNoContent() *PATCHAPIOrganizationsNoContent {
	return &PATCHAPIOrganizationsNoContent{}
}

/*
PATCHAPIOrganizationsNoContent describes a response with status code 204, with default header values.

No Content
*/
type PATCHAPIOrganizationsNoContent struct {
}

// IsSuccess returns true when this p a t c h Api organizations no content response has a 2xx status code
func (o *PATCHAPIOrganizationsNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this p a t c h Api organizations no content response has a 3xx status code
func (o *PATCHAPIOrganizationsNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this p a t c h Api organizations no content response has a 4xx status code
func (o *PATCHAPIOrganizationsNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this p a t c h Api organizations no content response has a 5xx status code
func (o *PATCHAPIOrganizationsNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this p a t c h Api organizations no content response a status code equal to that given
func (o *PATCHAPIOrganizationsNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the p a t c h Api organizations no content response
func (o *PATCHAPIOrganizationsNoContent) Code() int {
	return 204
}

func (o *PATCHAPIOrganizationsNoContent) Error() string {
	return fmt.Sprintf("[PATCH /api-management/1.0/apis/{apiUuid}/organizations][%d] pATCHApiOrganizationsNoContent ", 204)
}

func (o *PATCHAPIOrganizationsNoContent) String() string {
	return fmt.Sprintf("[PATCH /api-management/1.0/apis/{apiUuid}/organizations][%d] pATCHApiOrganizationsNoContent ", 204)
}

func (o *PATCHAPIOrganizationsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPATCHAPIOrganizationsBadRequest creates a PATCHAPIOrganizationsBadRequest with default headers values
func NewPATCHAPIOrganizationsBadRequest() *PATCHAPIOrganizationsBadRequest {
	return &PATCHAPIOrganizationsBadRequest{}
}

/*
PATCHAPIOrganizationsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PATCHAPIOrganizationsBadRequest struct {
}

// IsSuccess returns true when this p a t c h Api organizations bad request response has a 2xx status code
func (o *PATCHAPIOrganizationsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this p a t c h Api organizations bad request response has a 3xx status code
func (o *PATCHAPIOrganizationsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this p a t c h Api organizations bad request response has a 4xx status code
func (o *PATCHAPIOrganizationsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this p a t c h Api organizations bad request response has a 5xx status code
func (o *PATCHAPIOrganizationsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this p a t c h Api organizations bad request response a status code equal to that given
func (o *PATCHAPIOrganizationsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the p a t c h Api organizations bad request response
func (o *PATCHAPIOrganizationsBadRequest) Code() int {
	return 400
}

func (o *PATCHAPIOrganizationsBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api-management/1.0/apis/{apiUuid}/organizations][%d] pATCHApiOrganizationsBadRequest ", 400)
}

func (o *PATCHAPIOrganizationsBadRequest) String() string {
	return fmt.Sprintf("[PATCH /api-management/1.0/apis/{apiUuid}/organizations][%d] pATCHApiOrganizationsBadRequest ", 400)
}

func (o *PATCHAPIOrganizationsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPATCHAPIOrganizationsInternalServerError creates a PATCHAPIOrganizationsInternalServerError with default headers values
func NewPATCHAPIOrganizationsInternalServerError() *PATCHAPIOrganizationsInternalServerError {
	return &PATCHAPIOrganizationsInternalServerError{}
}

/*
PATCHAPIOrganizationsInternalServerError describes a response with status code 500, with default header values.

Server Failure
*/
type PATCHAPIOrganizationsInternalServerError struct {
}

// IsSuccess returns true when this p a t c h Api organizations internal server error response has a 2xx status code
func (o *PATCHAPIOrganizationsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this p a t c h Api organizations internal server error response has a 3xx status code
func (o *PATCHAPIOrganizationsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this p a t c h Api organizations internal server error response has a 4xx status code
func (o *PATCHAPIOrganizationsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this p a t c h Api organizations internal server error response has a 5xx status code
func (o *PATCHAPIOrganizationsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this p a t c h Api organizations internal server error response a status code equal to that given
func (o *PATCHAPIOrganizationsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the p a t c h Api organizations internal server error response
func (o *PATCHAPIOrganizationsInternalServerError) Code() int {
	return 500
}

func (o *PATCHAPIOrganizationsInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api-management/1.0/apis/{apiUuid}/organizations][%d] pATCHApiOrganizationsInternalServerError ", 500)
}

func (o *PATCHAPIOrganizationsInternalServerError) String() string {
	return fmt.Sprintf("[PATCH /api-management/1.0/apis/{apiUuid}/organizations][%d] pATCHApiOrganizationsInternalServerError ", 500)
}

func (o *PATCHAPIOrganizationsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
