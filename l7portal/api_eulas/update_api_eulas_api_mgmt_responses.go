// Code generated by go-swagger; DO NOT EDIT.

package api_eulas

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateAPIEulasAPIMgmtReader is a Reader for the UpdateAPIEulasAPIMgmt structure.
type UpdateAPIEulasAPIMgmtReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateAPIEulasAPIMgmtReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 400:
		result := NewUpdateAPIEulasAPIMgmtBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateAPIEulasAPIMgmtNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateAPIEulasAPIMgmtInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateAPIEulasAPIMgmtBadRequest creates a UpdateAPIEulasAPIMgmtBadRequest with default headers values
func NewUpdateAPIEulasAPIMgmtBadRequest() *UpdateAPIEulasAPIMgmtBadRequest {
	return &UpdateAPIEulasAPIMgmtBadRequest{}
}

/*
UpdateAPIEulasAPIMgmtBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdateAPIEulasAPIMgmtBadRequest struct {
}

// IsSuccess returns true when this update Api eulas Api mgmt bad request response has a 2xx status code
func (o *UpdateAPIEulasAPIMgmtBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update Api eulas Api mgmt bad request response has a 3xx status code
func (o *UpdateAPIEulasAPIMgmtBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update Api eulas Api mgmt bad request response has a 4xx status code
func (o *UpdateAPIEulasAPIMgmtBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update Api eulas Api mgmt bad request response has a 5xx status code
func (o *UpdateAPIEulasAPIMgmtBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update Api eulas Api mgmt bad request response a status code equal to that given
func (o *UpdateAPIEulasAPIMgmtBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update Api eulas Api mgmt bad request response
func (o *UpdateAPIEulasAPIMgmtBadRequest) Code() int {
	return 400
}

func (o *UpdateAPIEulasAPIMgmtBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api-management/1.0/eulas/{uuid}][%d] updateApiEulasApiMgmtBadRequest ", 400)
}

func (o *UpdateAPIEulasAPIMgmtBadRequest) String() string {
	return fmt.Sprintf("[PUT /api-management/1.0/eulas/{uuid}][%d] updateApiEulasApiMgmtBadRequest ", 400)
}

func (o *UpdateAPIEulasAPIMgmtBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateAPIEulasAPIMgmtNotFound creates a UpdateAPIEulasAPIMgmtNotFound with default headers values
func NewUpdateAPIEulasAPIMgmtNotFound() *UpdateAPIEulasAPIMgmtNotFound {
	return &UpdateAPIEulasAPIMgmtNotFound{}
}

/*
UpdateAPIEulasAPIMgmtNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdateAPIEulasAPIMgmtNotFound struct {
}

// IsSuccess returns true when this update Api eulas Api mgmt not found response has a 2xx status code
func (o *UpdateAPIEulasAPIMgmtNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update Api eulas Api mgmt not found response has a 3xx status code
func (o *UpdateAPIEulasAPIMgmtNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update Api eulas Api mgmt not found response has a 4xx status code
func (o *UpdateAPIEulasAPIMgmtNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update Api eulas Api mgmt not found response has a 5xx status code
func (o *UpdateAPIEulasAPIMgmtNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update Api eulas Api mgmt not found response a status code equal to that given
func (o *UpdateAPIEulasAPIMgmtNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update Api eulas Api mgmt not found response
func (o *UpdateAPIEulasAPIMgmtNotFound) Code() int {
	return 404
}

func (o *UpdateAPIEulasAPIMgmtNotFound) Error() string {
	return fmt.Sprintf("[PUT /api-management/1.0/eulas/{uuid}][%d] updateApiEulasApiMgmtNotFound ", 404)
}

func (o *UpdateAPIEulasAPIMgmtNotFound) String() string {
	return fmt.Sprintf("[PUT /api-management/1.0/eulas/{uuid}][%d] updateApiEulasApiMgmtNotFound ", 404)
}

func (o *UpdateAPIEulasAPIMgmtNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateAPIEulasAPIMgmtInternalServerError creates a UpdateAPIEulasAPIMgmtInternalServerError with default headers values
func NewUpdateAPIEulasAPIMgmtInternalServerError() *UpdateAPIEulasAPIMgmtInternalServerError {
	return &UpdateAPIEulasAPIMgmtInternalServerError{}
}

/*
UpdateAPIEulasAPIMgmtInternalServerError describes a response with status code 500, with default header values.

Server Failure
*/
type UpdateAPIEulasAPIMgmtInternalServerError struct {
}

// IsSuccess returns true when this update Api eulas Api mgmt internal server error response has a 2xx status code
func (o *UpdateAPIEulasAPIMgmtInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update Api eulas Api mgmt internal server error response has a 3xx status code
func (o *UpdateAPIEulasAPIMgmtInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update Api eulas Api mgmt internal server error response has a 4xx status code
func (o *UpdateAPIEulasAPIMgmtInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update Api eulas Api mgmt internal server error response has a 5xx status code
func (o *UpdateAPIEulasAPIMgmtInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update Api eulas Api mgmt internal server error response a status code equal to that given
func (o *UpdateAPIEulasAPIMgmtInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update Api eulas Api mgmt internal server error response
func (o *UpdateAPIEulasAPIMgmtInternalServerError) Code() int {
	return 500
}

func (o *UpdateAPIEulasAPIMgmtInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api-management/1.0/eulas/{uuid}][%d] updateApiEulasApiMgmtInternalServerError ", 500)
}

func (o *UpdateAPIEulasAPIMgmtInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api-management/1.0/eulas/{uuid}][%d] updateApiEulasApiMgmtInternalServerError ", 500)
}

func (o *UpdateAPIEulasAPIMgmtInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
