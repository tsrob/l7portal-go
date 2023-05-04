// Code generated by go-swagger; DO NOT EDIT.

package apis

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteAPIUsingDELETEReader is a Reader for the DeleteAPIUsingDELETE structure.
type DeleteAPIUsingDELETEReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAPIUsingDELETEReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteAPIUsingDELETENoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteAPIUsingDELETEBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteAPIUsingDELETENoContent creates a DeleteAPIUsingDELETENoContent with default headers values
func NewDeleteAPIUsingDELETENoContent() *DeleteAPIUsingDELETENoContent {
	return &DeleteAPIUsingDELETENoContent{}
}

/*
DeleteAPIUsingDELETENoContent describes a response with status code 204, with default header values.

No Content
*/
type DeleteAPIUsingDELETENoContent struct {
}

// IsSuccess returns true when this delete Api using d e l e t e no content response has a 2xx status code
func (o *DeleteAPIUsingDELETENoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete Api using d e l e t e no content response has a 3xx status code
func (o *DeleteAPIUsingDELETENoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete Api using d e l e t e no content response has a 4xx status code
func (o *DeleteAPIUsingDELETENoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete Api using d e l e t e no content response has a 5xx status code
func (o *DeleteAPIUsingDELETENoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete Api using d e l e t e no content response a status code equal to that given
func (o *DeleteAPIUsingDELETENoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete Api using d e l e t e no content response
func (o *DeleteAPIUsingDELETENoContent) Code() int {
	return 204
}

func (o *DeleteAPIUsingDELETENoContent) Error() string {
	return fmt.Sprintf("[DELETE /api-management/1.0/apis/{apiUuid}][%d] deleteApiUsingDELETENoContent ", 204)
}

func (o *DeleteAPIUsingDELETENoContent) String() string {
	return fmt.Sprintf("[DELETE /api-management/1.0/apis/{apiUuid}][%d] deleteApiUsingDELETENoContent ", 204)
}

func (o *DeleteAPIUsingDELETENoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAPIUsingDELETEBadRequest creates a DeleteAPIUsingDELETEBadRequest with default headers values
func NewDeleteAPIUsingDELETEBadRequest() *DeleteAPIUsingDELETEBadRequest {
	return &DeleteAPIUsingDELETEBadRequest{}
}

/*
DeleteAPIUsingDELETEBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteAPIUsingDELETEBadRequest struct {
}

// IsSuccess returns true when this delete Api using d e l e t e bad request response has a 2xx status code
func (o *DeleteAPIUsingDELETEBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete Api using d e l e t e bad request response has a 3xx status code
func (o *DeleteAPIUsingDELETEBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete Api using d e l e t e bad request response has a 4xx status code
func (o *DeleteAPIUsingDELETEBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete Api using d e l e t e bad request response has a 5xx status code
func (o *DeleteAPIUsingDELETEBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete Api using d e l e t e bad request response a status code equal to that given
func (o *DeleteAPIUsingDELETEBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete Api using d e l e t e bad request response
func (o *DeleteAPIUsingDELETEBadRequest) Code() int {
	return 400
}

func (o *DeleteAPIUsingDELETEBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api-management/1.0/apis/{apiUuid}][%d] deleteApiUsingDELETEBadRequest ", 400)
}

func (o *DeleteAPIUsingDELETEBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api-management/1.0/apis/{apiUuid}][%d] deleteApiUsingDELETEBadRequest ", 400)
}

func (o *DeleteAPIUsingDELETEBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
