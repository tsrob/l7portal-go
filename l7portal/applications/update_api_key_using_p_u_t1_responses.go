// Code generated by go-swagger; DO NOT EDIT.

package applications

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateAPIKeyUsingPUT1Reader is a Reader for the UpdateAPIKeyUsingPUT1 structure.
type UpdateAPIKeyUsingPUT1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateAPIKeyUsingPUT1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateAPIKeyUsingPUT1NoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateAPIKeyUsingPUT1NoContent creates a UpdateAPIKeyUsingPUT1NoContent with default headers values
func NewUpdateAPIKeyUsingPUT1NoContent() *UpdateAPIKeyUsingPUT1NoContent {
	return &UpdateAPIKeyUsingPUT1NoContent{}
}

/*
UpdateAPIKeyUsingPUT1NoContent describes a response with status code 204, with default header values.

No Content
*/
type UpdateAPIKeyUsingPUT1NoContent struct {
}

// IsSuccess returns true when this update Api key using p u t1 no content response has a 2xx status code
func (o *UpdateAPIKeyUsingPUT1NoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update Api key using p u t1 no content response has a 3xx status code
func (o *UpdateAPIKeyUsingPUT1NoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update Api key using p u t1 no content response has a 4xx status code
func (o *UpdateAPIKeyUsingPUT1NoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this update Api key using p u t1 no content response has a 5xx status code
func (o *UpdateAPIKeyUsingPUT1NoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this update Api key using p u t1 no content response a status code equal to that given
func (o *UpdateAPIKeyUsingPUT1NoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the update Api key using p u t1 no content response
func (o *UpdateAPIKeyUsingPUT1NoContent) Code() int {
	return 204
}

func (o *UpdateAPIKeyUsingPUT1NoContent) Error() string {
	return fmt.Sprintf("[PUT /api-management/1.0/applications/{appUuid}/api-keys/{apiKey}][%d] updateApiKeyUsingPUT1NoContent ", 204)
}

func (o *UpdateAPIKeyUsingPUT1NoContent) String() string {
	return fmt.Sprintf("[PUT /api-management/1.0/applications/{appUuid}/api-keys/{apiKey}][%d] updateApiKeyUsingPUT1NoContent ", 204)
}

func (o *UpdateAPIKeyUsingPUT1NoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
