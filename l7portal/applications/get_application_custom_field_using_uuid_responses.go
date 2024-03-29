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

// GetApplicationCustomFieldUsingUUIDReader is a Reader for the GetApplicationCustomFieldUsingUUID structure.
type GetApplicationCustomFieldUsingUUIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetApplicationCustomFieldUsingUUIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApplicationCustomFieldUsingUUIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetApplicationCustomFieldUsingUUIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetApplicationCustomFieldUsingUUIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetApplicationCustomFieldUsingUUIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetApplicationCustomFieldUsingUUIDOK creates a GetApplicationCustomFieldUsingUUIDOK with default headers values
func NewGetApplicationCustomFieldUsingUUIDOK() *GetApplicationCustomFieldUsingUUIDOK {
	return &GetApplicationCustomFieldUsingUUIDOK{}
}

/*
GetApplicationCustomFieldUsingUUIDOK describes a response with status code 200, with default header values.

Success
*/
type GetApplicationCustomFieldUsingUUIDOK struct {
	Payload []*models.CustomFieldValueDto
}

// IsSuccess returns true when this get application custom field using Uuid o k response has a 2xx status code
func (o *GetApplicationCustomFieldUsingUUIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get application custom field using Uuid o k response has a 3xx status code
func (o *GetApplicationCustomFieldUsingUUIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get application custom field using Uuid o k response has a 4xx status code
func (o *GetApplicationCustomFieldUsingUUIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get application custom field using Uuid o k response has a 5xx status code
func (o *GetApplicationCustomFieldUsingUUIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get application custom field using Uuid o k response a status code equal to that given
func (o *GetApplicationCustomFieldUsingUUIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get application custom field using Uuid o k response
func (o *GetApplicationCustomFieldUsingUUIDOK) Code() int {
	return 200
}

func (o *GetApplicationCustomFieldUsingUUIDOK) Error() string {
	return fmt.Sprintf("[GET /api-management/1.0/applications/{appUuid}/custom-fields][%d] getApplicationCustomFieldUsingUuidOK  %+v", 200, o.Payload)
}

func (o *GetApplicationCustomFieldUsingUUIDOK) String() string {
	return fmt.Sprintf("[GET /api-management/1.0/applications/{appUuid}/custom-fields][%d] getApplicationCustomFieldUsingUuidOK  %+v", 200, o.Payload)
}

func (o *GetApplicationCustomFieldUsingUUIDOK) GetPayload() []*models.CustomFieldValueDto {
	return o.Payload
}

func (o *GetApplicationCustomFieldUsingUUIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApplicationCustomFieldUsingUUIDBadRequest creates a GetApplicationCustomFieldUsingUUIDBadRequest with default headers values
func NewGetApplicationCustomFieldUsingUUIDBadRequest() *GetApplicationCustomFieldUsingUUIDBadRequest {
	return &GetApplicationCustomFieldUsingUUIDBadRequest{}
}

/*
GetApplicationCustomFieldUsingUUIDBadRequest describes a response with status code 400, with default header values.

Bad Request due to a validation error.
*/
type GetApplicationCustomFieldUsingUUIDBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get application custom field using Uuid bad request response has a 2xx status code
func (o *GetApplicationCustomFieldUsingUUIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get application custom field using Uuid bad request response has a 3xx status code
func (o *GetApplicationCustomFieldUsingUUIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get application custom field using Uuid bad request response has a 4xx status code
func (o *GetApplicationCustomFieldUsingUUIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get application custom field using Uuid bad request response has a 5xx status code
func (o *GetApplicationCustomFieldUsingUUIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get application custom field using Uuid bad request response a status code equal to that given
func (o *GetApplicationCustomFieldUsingUUIDBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get application custom field using Uuid bad request response
func (o *GetApplicationCustomFieldUsingUUIDBadRequest) Code() int {
	return 400
}

func (o *GetApplicationCustomFieldUsingUUIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /api-management/1.0/applications/{appUuid}/custom-fields][%d] getApplicationCustomFieldUsingUuidBadRequest  %+v", 400, o.Payload)
}

func (o *GetApplicationCustomFieldUsingUUIDBadRequest) String() string {
	return fmt.Sprintf("[GET /api-management/1.0/applications/{appUuid}/custom-fields][%d] getApplicationCustomFieldUsingUuidBadRequest  %+v", 400, o.Payload)
}

func (o *GetApplicationCustomFieldUsingUUIDBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetApplicationCustomFieldUsingUUIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApplicationCustomFieldUsingUUIDNotFound creates a GetApplicationCustomFieldUsingUUIDNotFound with default headers values
func NewGetApplicationCustomFieldUsingUUIDNotFound() *GetApplicationCustomFieldUsingUUIDNotFound {
	return &GetApplicationCustomFieldUsingUUIDNotFound{}
}

/*
GetApplicationCustomFieldUsingUUIDNotFound describes a response with status code 404, with default header values.

Entity not found
*/
type GetApplicationCustomFieldUsingUUIDNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get application custom field using Uuid not found response has a 2xx status code
func (o *GetApplicationCustomFieldUsingUUIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get application custom field using Uuid not found response has a 3xx status code
func (o *GetApplicationCustomFieldUsingUUIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get application custom field using Uuid not found response has a 4xx status code
func (o *GetApplicationCustomFieldUsingUUIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get application custom field using Uuid not found response has a 5xx status code
func (o *GetApplicationCustomFieldUsingUUIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get application custom field using Uuid not found response a status code equal to that given
func (o *GetApplicationCustomFieldUsingUUIDNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get application custom field using Uuid not found response
func (o *GetApplicationCustomFieldUsingUUIDNotFound) Code() int {
	return 404
}

func (o *GetApplicationCustomFieldUsingUUIDNotFound) Error() string {
	return fmt.Sprintf("[GET /api-management/1.0/applications/{appUuid}/custom-fields][%d] getApplicationCustomFieldUsingUuidNotFound  %+v", 404, o.Payload)
}

func (o *GetApplicationCustomFieldUsingUUIDNotFound) String() string {
	return fmt.Sprintf("[GET /api-management/1.0/applications/{appUuid}/custom-fields][%d] getApplicationCustomFieldUsingUuidNotFound  %+v", 404, o.Payload)
}

func (o *GetApplicationCustomFieldUsingUUIDNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetApplicationCustomFieldUsingUUIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApplicationCustomFieldUsingUUIDInternalServerError creates a GetApplicationCustomFieldUsingUUIDInternalServerError with default headers values
func NewGetApplicationCustomFieldUsingUUIDInternalServerError() *GetApplicationCustomFieldUsingUUIDInternalServerError {
	return &GetApplicationCustomFieldUsingUUIDInternalServerError{}
}

/*
GetApplicationCustomFieldUsingUUIDInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. Certain fields are missing from the request.
*/
type GetApplicationCustomFieldUsingUUIDInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get application custom field using Uuid internal server error response has a 2xx status code
func (o *GetApplicationCustomFieldUsingUUIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get application custom field using Uuid internal server error response has a 3xx status code
func (o *GetApplicationCustomFieldUsingUUIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get application custom field using Uuid internal server error response has a 4xx status code
func (o *GetApplicationCustomFieldUsingUUIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get application custom field using Uuid internal server error response has a 5xx status code
func (o *GetApplicationCustomFieldUsingUUIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get application custom field using Uuid internal server error response a status code equal to that given
func (o *GetApplicationCustomFieldUsingUUIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get application custom field using Uuid internal server error response
func (o *GetApplicationCustomFieldUsingUUIDInternalServerError) Code() int {
	return 500
}

func (o *GetApplicationCustomFieldUsingUUIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api-management/1.0/applications/{appUuid}/custom-fields][%d] getApplicationCustomFieldUsingUuidInternalServerError  %+v", 500, o.Payload)
}

func (o *GetApplicationCustomFieldUsingUUIDInternalServerError) String() string {
	return fmt.Sprintf("[GET /api-management/1.0/applications/{appUuid}/custom-fields][%d] getApplicationCustomFieldUsingUuidInternalServerError  %+v", 500, o.Payload)
}

func (o *GetApplicationCustomFieldUsingUUIDInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetApplicationCustomFieldUsingUUIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
