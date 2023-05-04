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

// UpdateAPITagsReader is a Reader for the UpdateAPITags structure.
type UpdateAPITagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateAPITagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateAPITagsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateAPITagsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateAPITagsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateAPITagsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateAPITagsOK creates a UpdateAPITagsOK with default headers values
func NewUpdateAPITagsOK() *UpdateAPITagsOK {
	return &UpdateAPITagsOK{}
}

/*
UpdateAPITagsOK describes a response with status code 200, with default header values.

Updated the API Tag name
*/
type UpdateAPITagsOK struct {
	Payload *models.TagResponse
}

// IsSuccess returns true when this update Api tags o k response has a 2xx status code
func (o *UpdateAPITagsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update Api tags o k response has a 3xx status code
func (o *UpdateAPITagsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update Api tags o k response has a 4xx status code
func (o *UpdateAPITagsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update Api tags o k response has a 5xx status code
func (o *UpdateAPITagsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update Api tags o k response a status code equal to that given
func (o *UpdateAPITagsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update Api tags o k response
func (o *UpdateAPITagsOK) Code() int {
	return 200
}

func (o *UpdateAPITagsOK) Error() string {
	return fmt.Sprintf("[PUT /api-management/1.0/tags/{tagUuid}][%d] updateApiTagsOK  %+v", 200, o.Payload)
}

func (o *UpdateAPITagsOK) String() string {
	return fmt.Sprintf("[PUT /api-management/1.0/tags/{tagUuid}][%d] updateApiTagsOK  %+v", 200, o.Payload)
}

func (o *UpdateAPITagsOK) GetPayload() *models.TagResponse {
	return o.Payload
}

func (o *UpdateAPITagsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TagResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAPITagsBadRequest creates a UpdateAPITagsBadRequest with default headers values
func NewUpdateAPITagsBadRequest() *UpdateAPITagsBadRequest {
	return &UpdateAPITagsBadRequest{}
}

/*
UpdateAPITagsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdateAPITagsBadRequest struct {
}

// IsSuccess returns true when this update Api tags bad request response has a 2xx status code
func (o *UpdateAPITagsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update Api tags bad request response has a 3xx status code
func (o *UpdateAPITagsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update Api tags bad request response has a 4xx status code
func (o *UpdateAPITagsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update Api tags bad request response has a 5xx status code
func (o *UpdateAPITagsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update Api tags bad request response a status code equal to that given
func (o *UpdateAPITagsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update Api tags bad request response
func (o *UpdateAPITagsBadRequest) Code() int {
	return 400
}

func (o *UpdateAPITagsBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api-management/1.0/tags/{tagUuid}][%d] updateApiTagsBadRequest ", 400)
}

func (o *UpdateAPITagsBadRequest) String() string {
	return fmt.Sprintf("[PUT /api-management/1.0/tags/{tagUuid}][%d] updateApiTagsBadRequest ", 400)
}

func (o *UpdateAPITagsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateAPITagsNotFound creates a UpdateAPITagsNotFound with default headers values
func NewUpdateAPITagsNotFound() *UpdateAPITagsNotFound {
	return &UpdateAPITagsNotFound{}
}

/*
UpdateAPITagsNotFound describes a response with status code 404, with default header values.

Resource Not Found
*/
type UpdateAPITagsNotFound struct {
}

// IsSuccess returns true when this update Api tags not found response has a 2xx status code
func (o *UpdateAPITagsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update Api tags not found response has a 3xx status code
func (o *UpdateAPITagsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update Api tags not found response has a 4xx status code
func (o *UpdateAPITagsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update Api tags not found response has a 5xx status code
func (o *UpdateAPITagsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update Api tags not found response a status code equal to that given
func (o *UpdateAPITagsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update Api tags not found response
func (o *UpdateAPITagsNotFound) Code() int {
	return 404
}

func (o *UpdateAPITagsNotFound) Error() string {
	return fmt.Sprintf("[PUT /api-management/1.0/tags/{tagUuid}][%d] updateApiTagsNotFound ", 404)
}

func (o *UpdateAPITagsNotFound) String() string {
	return fmt.Sprintf("[PUT /api-management/1.0/tags/{tagUuid}][%d] updateApiTagsNotFound ", 404)
}

func (o *UpdateAPITagsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateAPITagsInternalServerError creates a UpdateAPITagsInternalServerError with default headers values
func NewUpdateAPITagsInternalServerError() *UpdateAPITagsInternalServerError {
	return &UpdateAPITagsInternalServerError{}
}

/*
UpdateAPITagsInternalServerError describes a response with status code 500, with default header values.

Server Failure
*/
type UpdateAPITagsInternalServerError struct {
}

// IsSuccess returns true when this update Api tags internal server error response has a 2xx status code
func (o *UpdateAPITagsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update Api tags internal server error response has a 3xx status code
func (o *UpdateAPITagsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update Api tags internal server error response has a 4xx status code
func (o *UpdateAPITagsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update Api tags internal server error response has a 5xx status code
func (o *UpdateAPITagsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update Api tags internal server error response a status code equal to that given
func (o *UpdateAPITagsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update Api tags internal server error response
func (o *UpdateAPITagsInternalServerError) Code() int {
	return 500
}

func (o *UpdateAPITagsInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api-management/1.0/tags/{tagUuid}][%d] updateApiTagsInternalServerError ", 500)
}

func (o *UpdateAPITagsInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api-management/1.0/tags/{tagUuid}][%d] updateApiTagsInternalServerError ", 500)
}

func (o *UpdateAPITagsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}