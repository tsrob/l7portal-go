// Code generated by go-swagger; DO NOT EDIT.

package tags

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteAPITagReader is a Reader for the DeleteAPITag structure.
type DeleteAPITagReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAPITagReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteAPITagNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteAPITagBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteAPITagNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteAPITagInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteAPITagNoContent creates a DeleteAPITagNoContent with default headers values
func NewDeleteAPITagNoContent() *DeleteAPITagNoContent {
	return &DeleteAPITagNoContent{}
}

/*
DeleteAPITagNoContent describes a response with status code 204, with default header values.

Deleted the Tag
*/
type DeleteAPITagNoContent struct {
}

// IsSuccess returns true when this delete Api tag no content response has a 2xx status code
func (o *DeleteAPITagNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete Api tag no content response has a 3xx status code
func (o *DeleteAPITagNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete Api tag no content response has a 4xx status code
func (o *DeleteAPITagNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete Api tag no content response has a 5xx status code
func (o *DeleteAPITagNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete Api tag no content response a status code equal to that given
func (o *DeleteAPITagNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete Api tag no content response
func (o *DeleteAPITagNoContent) Code() int {
	return 204
}

func (o *DeleteAPITagNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api-management/1.0/tags/{tagUuid}][%d] deleteApiTagNoContent ", 204)
}

func (o *DeleteAPITagNoContent) String() string {
	return fmt.Sprintf("[DELETE /api-management/1.0/tags/{tagUuid}][%d] deleteApiTagNoContent ", 204)
}

func (o *DeleteAPITagNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAPITagBadRequest creates a DeleteAPITagBadRequest with default headers values
func NewDeleteAPITagBadRequest() *DeleteAPITagBadRequest {
	return &DeleteAPITagBadRequest{}
}

/*
DeleteAPITagBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteAPITagBadRequest struct {
}

// IsSuccess returns true when this delete Api tag bad request response has a 2xx status code
func (o *DeleteAPITagBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete Api tag bad request response has a 3xx status code
func (o *DeleteAPITagBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete Api tag bad request response has a 4xx status code
func (o *DeleteAPITagBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete Api tag bad request response has a 5xx status code
func (o *DeleteAPITagBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete Api tag bad request response a status code equal to that given
func (o *DeleteAPITagBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete Api tag bad request response
func (o *DeleteAPITagBadRequest) Code() int {
	return 400
}

func (o *DeleteAPITagBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api-management/1.0/tags/{tagUuid}][%d] deleteApiTagBadRequest ", 400)
}

func (o *DeleteAPITagBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api-management/1.0/tags/{tagUuid}][%d] deleteApiTagBadRequest ", 400)
}

func (o *DeleteAPITagBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAPITagNotFound creates a DeleteAPITagNotFound with default headers values
func NewDeleteAPITagNotFound() *DeleteAPITagNotFound {
	return &DeleteAPITagNotFound{}
}

/*
DeleteAPITagNotFound describes a response with status code 404, with default header values.

Resource Not Found
*/
type DeleteAPITagNotFound struct {
}

// IsSuccess returns true when this delete Api tag not found response has a 2xx status code
func (o *DeleteAPITagNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete Api tag not found response has a 3xx status code
func (o *DeleteAPITagNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete Api tag not found response has a 4xx status code
func (o *DeleteAPITagNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete Api tag not found response has a 5xx status code
func (o *DeleteAPITagNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete Api tag not found response a status code equal to that given
func (o *DeleteAPITagNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete Api tag not found response
func (o *DeleteAPITagNotFound) Code() int {
	return 404
}

func (o *DeleteAPITagNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api-management/1.0/tags/{tagUuid}][%d] deleteApiTagNotFound ", 404)
}

func (o *DeleteAPITagNotFound) String() string {
	return fmt.Sprintf("[DELETE /api-management/1.0/tags/{tagUuid}][%d] deleteApiTagNotFound ", 404)
}

func (o *DeleteAPITagNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAPITagInternalServerError creates a DeleteAPITagInternalServerError with default headers values
func NewDeleteAPITagInternalServerError() *DeleteAPITagInternalServerError {
	return &DeleteAPITagInternalServerError{}
}

/*
DeleteAPITagInternalServerError describes a response with status code 500, with default header values.

Server Failure
*/
type DeleteAPITagInternalServerError struct {
}

// IsSuccess returns true when this delete Api tag internal server error response has a 2xx status code
func (o *DeleteAPITagInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete Api tag internal server error response has a 3xx status code
func (o *DeleteAPITagInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete Api tag internal server error response has a 4xx status code
func (o *DeleteAPITagInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete Api tag internal server error response has a 5xx status code
func (o *DeleteAPITagInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete Api tag internal server error response a status code equal to that given
func (o *DeleteAPITagInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete Api tag internal server error response
func (o *DeleteAPITagInternalServerError) Code() int {
	return 500
}

func (o *DeleteAPITagInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api-management/1.0/tags/{tagUuid}][%d] deleteApiTagInternalServerError ", 500)
}

func (o *DeleteAPITagInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api-management/1.0/tags/{tagUuid}][%d] deleteApiTagInternalServerError ", 500)
}

func (o *DeleteAPITagInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}