// Code generated by go-swagger; DO NOT EDIT.

package apis

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// SaveAPICustomFieldValueUsingPUTReader is a Reader for the SaveAPICustomFieldValueUsingPUT structure.
type SaveAPICustomFieldValueUsingPUTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SaveAPICustomFieldValueUsingPUTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewSaveAPICustomFieldValueUsingPUTNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSaveAPICustomFieldValueUsingPUTBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSaveAPICustomFieldValueUsingPUTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSaveAPICustomFieldValueUsingPUTInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSaveAPICustomFieldValueUsingPUTNoContent creates a SaveAPICustomFieldValueUsingPUTNoContent with default headers values
func NewSaveAPICustomFieldValueUsingPUTNoContent() *SaveAPICustomFieldValueUsingPUTNoContent {
	return &SaveAPICustomFieldValueUsingPUTNoContent{}
}

/*
SaveAPICustomFieldValueUsingPUTNoContent describes a response with status code 204, with default header values.

No Content
*/
type SaveAPICustomFieldValueUsingPUTNoContent struct {
}

// IsSuccess returns true when this save Api custom field value using p u t no content response has a 2xx status code
func (o *SaveAPICustomFieldValueUsingPUTNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this save Api custom field value using p u t no content response has a 3xx status code
func (o *SaveAPICustomFieldValueUsingPUTNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this save Api custom field value using p u t no content response has a 4xx status code
func (o *SaveAPICustomFieldValueUsingPUTNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this save Api custom field value using p u t no content response has a 5xx status code
func (o *SaveAPICustomFieldValueUsingPUTNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this save Api custom field value using p u t no content response a status code equal to that given
func (o *SaveAPICustomFieldValueUsingPUTNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the save Api custom field value using p u t no content response
func (o *SaveAPICustomFieldValueUsingPUTNoContent) Code() int {
	return 204
}

func (o *SaveAPICustomFieldValueUsingPUTNoContent) Error() string {
	return fmt.Sprintf("[PUT /api-management/1.0/apis/{apiUuid}/custom-fields][%d] saveApiCustomFieldValueUsingPUTNoContent ", 204)
}

func (o *SaveAPICustomFieldValueUsingPUTNoContent) String() string {
	return fmt.Sprintf("[PUT /api-management/1.0/apis/{apiUuid}/custom-fields][%d] saveApiCustomFieldValueUsingPUTNoContent ", 204)
}

func (o *SaveAPICustomFieldValueUsingPUTNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSaveAPICustomFieldValueUsingPUTBadRequest creates a SaveAPICustomFieldValueUsingPUTBadRequest with default headers values
func NewSaveAPICustomFieldValueUsingPUTBadRequest() *SaveAPICustomFieldValueUsingPUTBadRequest {
	return &SaveAPICustomFieldValueUsingPUTBadRequest{}
}

/*
SaveAPICustomFieldValueUsingPUTBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type SaveAPICustomFieldValueUsingPUTBadRequest struct {
}

// IsSuccess returns true when this save Api custom field value using p u t bad request response has a 2xx status code
func (o *SaveAPICustomFieldValueUsingPUTBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this save Api custom field value using p u t bad request response has a 3xx status code
func (o *SaveAPICustomFieldValueUsingPUTBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this save Api custom field value using p u t bad request response has a 4xx status code
func (o *SaveAPICustomFieldValueUsingPUTBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this save Api custom field value using p u t bad request response has a 5xx status code
func (o *SaveAPICustomFieldValueUsingPUTBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this save Api custom field value using p u t bad request response a status code equal to that given
func (o *SaveAPICustomFieldValueUsingPUTBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the save Api custom field value using p u t bad request response
func (o *SaveAPICustomFieldValueUsingPUTBadRequest) Code() int {
	return 400
}

func (o *SaveAPICustomFieldValueUsingPUTBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api-management/1.0/apis/{apiUuid}/custom-fields][%d] saveApiCustomFieldValueUsingPUTBadRequest ", 400)
}

func (o *SaveAPICustomFieldValueUsingPUTBadRequest) String() string {
	return fmt.Sprintf("[PUT /api-management/1.0/apis/{apiUuid}/custom-fields][%d] saveApiCustomFieldValueUsingPUTBadRequest ", 400)
}

func (o *SaveAPICustomFieldValueUsingPUTBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSaveAPICustomFieldValueUsingPUTNotFound creates a SaveAPICustomFieldValueUsingPUTNotFound with default headers values
func NewSaveAPICustomFieldValueUsingPUTNotFound() *SaveAPICustomFieldValueUsingPUTNotFound {
	return &SaveAPICustomFieldValueUsingPUTNotFound{}
}

/*
SaveAPICustomFieldValueUsingPUTNotFound describes a response with status code 404, with default header values.

API doesn't exist
*/
type SaveAPICustomFieldValueUsingPUTNotFound struct {
}

// IsSuccess returns true when this save Api custom field value using p u t not found response has a 2xx status code
func (o *SaveAPICustomFieldValueUsingPUTNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this save Api custom field value using p u t not found response has a 3xx status code
func (o *SaveAPICustomFieldValueUsingPUTNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this save Api custom field value using p u t not found response has a 4xx status code
func (o *SaveAPICustomFieldValueUsingPUTNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this save Api custom field value using p u t not found response has a 5xx status code
func (o *SaveAPICustomFieldValueUsingPUTNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this save Api custom field value using p u t not found response a status code equal to that given
func (o *SaveAPICustomFieldValueUsingPUTNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the save Api custom field value using p u t not found response
func (o *SaveAPICustomFieldValueUsingPUTNotFound) Code() int {
	return 404
}

func (o *SaveAPICustomFieldValueUsingPUTNotFound) Error() string {
	return fmt.Sprintf("[PUT /api-management/1.0/apis/{apiUuid}/custom-fields][%d] saveApiCustomFieldValueUsingPUTNotFound ", 404)
}

func (o *SaveAPICustomFieldValueUsingPUTNotFound) String() string {
	return fmt.Sprintf("[PUT /api-management/1.0/apis/{apiUuid}/custom-fields][%d] saveApiCustomFieldValueUsingPUTNotFound ", 404)
}

func (o *SaveAPICustomFieldValueUsingPUTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSaveAPICustomFieldValueUsingPUTInternalServerError creates a SaveAPICustomFieldValueUsingPUTInternalServerError with default headers values
func NewSaveAPICustomFieldValueUsingPUTInternalServerError() *SaveAPICustomFieldValueUsingPUTInternalServerError {
	return &SaveAPICustomFieldValueUsingPUTInternalServerError{}
}

/*
SaveAPICustomFieldValueUsingPUTInternalServerError describes a response with status code 500, with default header values.

Server Failure
*/
type SaveAPICustomFieldValueUsingPUTInternalServerError struct {
}

// IsSuccess returns true when this save Api custom field value using p u t internal server error response has a 2xx status code
func (o *SaveAPICustomFieldValueUsingPUTInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this save Api custom field value using p u t internal server error response has a 3xx status code
func (o *SaveAPICustomFieldValueUsingPUTInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this save Api custom field value using p u t internal server error response has a 4xx status code
func (o *SaveAPICustomFieldValueUsingPUTInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this save Api custom field value using p u t internal server error response has a 5xx status code
func (o *SaveAPICustomFieldValueUsingPUTInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this save Api custom field value using p u t internal server error response a status code equal to that given
func (o *SaveAPICustomFieldValueUsingPUTInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the save Api custom field value using p u t internal server error response
func (o *SaveAPICustomFieldValueUsingPUTInternalServerError) Code() int {
	return 500
}

func (o *SaveAPICustomFieldValueUsingPUTInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api-management/1.0/apis/{apiUuid}/custom-fields][%d] saveApiCustomFieldValueUsingPUTInternalServerError ", 500)
}

func (o *SaveAPICustomFieldValueUsingPUTInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api-management/1.0/apis/{apiUuid}/custom-fields][%d] saveApiCustomFieldValueUsingPUTInternalServerError ", 500)
}

func (o *SaveAPICustomFieldValueUsingPUTInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}