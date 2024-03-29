// Code generated by go-swagger; DO NOT EDIT.

package applications

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/tsrob/l7portal-go/models"
)

// DeleteAPIKeyUsingDELETEReader is a Reader for the DeleteAPIKeyUsingDELETE structure.
type DeleteAPIKeyUsingDELETEReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAPIKeyUsingDELETEReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteAPIKeyUsingDELETENoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteAPIKeyUsingDELETEBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteAPIKeyUsingDELETENotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteAPIKeyUsingDELETEInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteAPIKeyUsingDELETENoContent creates a DeleteAPIKeyUsingDELETENoContent with default headers values
func NewDeleteAPIKeyUsingDELETENoContent() *DeleteAPIKeyUsingDELETENoContent {
	return &DeleteAPIKeyUsingDELETENoContent{}
}

/*
DeleteAPIKeyUsingDELETENoContent describes a response with status code 204, with default header values.

No Content
*/
type DeleteAPIKeyUsingDELETENoContent struct {
}

// IsSuccess returns true when this delete Api key using d e l e t e no content response has a 2xx status code
func (o *DeleteAPIKeyUsingDELETENoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete Api key using d e l e t e no content response has a 3xx status code
func (o *DeleteAPIKeyUsingDELETENoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete Api key using d e l e t e no content response has a 4xx status code
func (o *DeleteAPIKeyUsingDELETENoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete Api key using d e l e t e no content response has a 5xx status code
func (o *DeleteAPIKeyUsingDELETENoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete Api key using d e l e t e no content response a status code equal to that given
func (o *DeleteAPIKeyUsingDELETENoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete Api key using d e l e t e no content response
func (o *DeleteAPIKeyUsingDELETENoContent) Code() int {
	return 204
}

func (o *DeleteAPIKeyUsingDELETENoContent) Error() string {
	return fmt.Sprintf("[DELETE /api-management/1.0/applications/{appUuid}/api-keys/{apiKey}][%d] deleteApiKeyUsingDELETENoContent ", 204)
}

func (o *DeleteAPIKeyUsingDELETENoContent) String() string {
	return fmt.Sprintf("[DELETE /api-management/1.0/applications/{appUuid}/api-keys/{apiKey}][%d] deleteApiKeyUsingDELETENoContent ", 204)
}

func (o *DeleteAPIKeyUsingDELETENoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAPIKeyUsingDELETEBadRequest creates a DeleteAPIKeyUsingDELETEBadRequest with default headers values
func NewDeleteAPIKeyUsingDELETEBadRequest() *DeleteAPIKeyUsingDELETEBadRequest {
	return &DeleteAPIKeyUsingDELETEBadRequest{}
}

/*
DeleteAPIKeyUsingDELETEBadRequest describes a response with status code 400, with default header values.

Bad Request due to Invalid Uuid.
*/
type DeleteAPIKeyUsingDELETEBadRequest struct {
	Payload *models.LongerError
}

// IsSuccess returns true when this delete Api key using d e l e t e bad request response has a 2xx status code
func (o *DeleteAPIKeyUsingDELETEBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete Api key using d e l e t e bad request response has a 3xx status code
func (o *DeleteAPIKeyUsingDELETEBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete Api key using d e l e t e bad request response has a 4xx status code
func (o *DeleteAPIKeyUsingDELETEBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete Api key using d e l e t e bad request response has a 5xx status code
func (o *DeleteAPIKeyUsingDELETEBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete Api key using d e l e t e bad request response a status code equal to that given
func (o *DeleteAPIKeyUsingDELETEBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete Api key using d e l e t e bad request response
func (o *DeleteAPIKeyUsingDELETEBadRequest) Code() int {
	return 400
}

func (o *DeleteAPIKeyUsingDELETEBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api-management/1.0/applications/{appUuid}/api-keys/{apiKey}][%d] deleteApiKeyUsingDELETEBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteAPIKeyUsingDELETEBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api-management/1.0/applications/{appUuid}/api-keys/{apiKey}][%d] deleteApiKeyUsingDELETEBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteAPIKeyUsingDELETEBadRequest) GetPayload() *models.LongerError {
	return o.Payload
}

func (o *DeleteAPIKeyUsingDELETEBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LongerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAPIKeyUsingDELETENotFound creates a DeleteAPIKeyUsingDELETENotFound with default headers values
func NewDeleteAPIKeyUsingDELETENotFound() *DeleteAPIKeyUsingDELETENotFound {
	return &DeleteAPIKeyUsingDELETENotFound{}
}

/*
DeleteAPIKeyUsingDELETENotFound describes a response with status code 404, with default header values.

Entity not Found
*/
type DeleteAPIKeyUsingDELETENotFound struct {
	Payload *models.ErrorEntity
}

// IsSuccess returns true when this delete Api key using d e l e t e not found response has a 2xx status code
func (o *DeleteAPIKeyUsingDELETENotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete Api key using d e l e t e not found response has a 3xx status code
func (o *DeleteAPIKeyUsingDELETENotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete Api key using d e l e t e not found response has a 4xx status code
func (o *DeleteAPIKeyUsingDELETENotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete Api key using d e l e t e not found response has a 5xx status code
func (o *DeleteAPIKeyUsingDELETENotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete Api key using d e l e t e not found response a status code equal to that given
func (o *DeleteAPIKeyUsingDELETENotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete Api key using d e l e t e not found response
func (o *DeleteAPIKeyUsingDELETENotFound) Code() int {
	return 404
}

func (o *DeleteAPIKeyUsingDELETENotFound) Error() string {
	return fmt.Sprintf("[DELETE /api-management/1.0/applications/{appUuid}/api-keys/{apiKey}][%d] deleteApiKeyUsingDELETENotFound  %+v", 404, o.Payload)
}

func (o *DeleteAPIKeyUsingDELETENotFound) String() string {
	return fmt.Sprintf("[DELETE /api-management/1.0/applications/{appUuid}/api-keys/{apiKey}][%d] deleteApiKeyUsingDELETENotFound  %+v", 404, o.Payload)
}

func (o *DeleteAPIKeyUsingDELETENotFound) GetPayload() *models.ErrorEntity {
	return o.Payload
}

func (o *DeleteAPIKeyUsingDELETENotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAPIKeyUsingDELETEInternalServerError creates a DeleteAPIKeyUsingDELETEInternalServerError with default headers values
func NewDeleteAPIKeyUsingDELETEInternalServerError() *DeleteAPIKeyUsingDELETEInternalServerError {
	return &DeleteAPIKeyUsingDELETEInternalServerError{}
}

/*
DeleteAPIKeyUsingDELETEInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. Certain fields are missing from the request.
*/
type DeleteAPIKeyUsingDELETEInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this delete Api key using d e l e t e internal server error response has a 2xx status code
func (o *DeleteAPIKeyUsingDELETEInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete Api key using d e l e t e internal server error response has a 3xx status code
func (o *DeleteAPIKeyUsingDELETEInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete Api key using d e l e t e internal server error response has a 4xx status code
func (o *DeleteAPIKeyUsingDELETEInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete Api key using d e l e t e internal server error response has a 5xx status code
func (o *DeleteAPIKeyUsingDELETEInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete Api key using d e l e t e internal server error response a status code equal to that given
func (o *DeleteAPIKeyUsingDELETEInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete Api key using d e l e t e internal server error response
func (o *DeleteAPIKeyUsingDELETEInternalServerError) Code() int {
	return 500
}

func (o *DeleteAPIKeyUsingDELETEInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api-management/1.0/applications/{appUuid}/api-keys/{apiKey}][%d] deleteApiKeyUsingDELETEInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteAPIKeyUsingDELETEInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api-management/1.0/applications/{appUuid}/api-keys/{apiKey}][%d] deleteApiKeyUsingDELETEInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteAPIKeyUsingDELETEInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteAPIKeyUsingDELETEInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
