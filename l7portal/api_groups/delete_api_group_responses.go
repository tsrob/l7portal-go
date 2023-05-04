// Code generated by go-swagger; DO NOT EDIT.

package api_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/tsrob/l7portal-go/models"
)

// DeleteAPIGroupReader is a Reader for the DeleteAPIGroup structure.
type DeleteAPIGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAPIGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 400:
		result := NewDeleteAPIGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteAPIGroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteAPIGroupInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteAPIGroupBadRequest creates a DeleteAPIGroupBadRequest with default headers values
func NewDeleteAPIGroupBadRequest() *DeleteAPIGroupBadRequest {
	return &DeleteAPIGroupBadRequest{}
}

/*
DeleteAPIGroupBadRequest describes a response with status code 400, with default header values.

Bad Request due to a validation error.
*/
type DeleteAPIGroupBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this delete Api group bad request response has a 2xx status code
func (o *DeleteAPIGroupBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete Api group bad request response has a 3xx status code
func (o *DeleteAPIGroupBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete Api group bad request response has a 4xx status code
func (o *DeleteAPIGroupBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete Api group bad request response has a 5xx status code
func (o *DeleteAPIGroupBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete Api group bad request response a status code equal to that given
func (o *DeleteAPIGroupBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete Api group bad request response
func (o *DeleteAPIGroupBadRequest) Code() int {
	return 400
}

func (o *DeleteAPIGroupBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api-management/1.0/api-groups/{uuid}][%d] deleteApiGroupBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteAPIGroupBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api-management/1.0/api-groups/{uuid}][%d] deleteApiGroupBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteAPIGroupBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteAPIGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAPIGroupNotFound creates a DeleteAPIGroupNotFound with default headers values
func NewDeleteAPIGroupNotFound() *DeleteAPIGroupNotFound {
	return &DeleteAPIGroupNotFound{}
}

/*
DeleteAPIGroupNotFound describes a response with status code 404, with default header values.

Entity not found
*/
type DeleteAPIGroupNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this delete Api group not found response has a 2xx status code
func (o *DeleteAPIGroupNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete Api group not found response has a 3xx status code
func (o *DeleteAPIGroupNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete Api group not found response has a 4xx status code
func (o *DeleteAPIGroupNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete Api group not found response has a 5xx status code
func (o *DeleteAPIGroupNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete Api group not found response a status code equal to that given
func (o *DeleteAPIGroupNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete Api group not found response
func (o *DeleteAPIGroupNotFound) Code() int {
	return 404
}

func (o *DeleteAPIGroupNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api-management/1.0/api-groups/{uuid}][%d] deleteApiGroupNotFound  %+v", 404, o.Payload)
}

func (o *DeleteAPIGroupNotFound) String() string {
	return fmt.Sprintf("[DELETE /api-management/1.0/api-groups/{uuid}][%d] deleteApiGroupNotFound  %+v", 404, o.Payload)
}

func (o *DeleteAPIGroupNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteAPIGroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAPIGroupInternalServerError creates a DeleteAPIGroupInternalServerError with default headers values
func NewDeleteAPIGroupInternalServerError() *DeleteAPIGroupInternalServerError {
	return &DeleteAPIGroupInternalServerError{}
}

/*
DeleteAPIGroupInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. Certain fields are missing from the request.
*/
type DeleteAPIGroupInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this delete Api group internal server error response has a 2xx status code
func (o *DeleteAPIGroupInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete Api group internal server error response has a 3xx status code
func (o *DeleteAPIGroupInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete Api group internal server error response has a 4xx status code
func (o *DeleteAPIGroupInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete Api group internal server error response has a 5xx status code
func (o *DeleteAPIGroupInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete Api group internal server error response a status code equal to that given
func (o *DeleteAPIGroupInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete Api group internal server error response
func (o *DeleteAPIGroupInternalServerError) Code() int {
	return 500
}

func (o *DeleteAPIGroupInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api-management/1.0/api-groups/{uuid}][%d] deleteApiGroupInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteAPIGroupInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api-management/1.0/api-groups/{uuid}][%d] deleteApiGroupInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteAPIGroupInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteAPIGroupInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}