// Code generated by go-swagger; DO NOT EDIT.

package settings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/tsrob/l7portal-go/models"
)

// UpdateSettingReader is a Reader for the UpdateSetting structure.
type UpdateSettingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateSettingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 400:
		result := NewUpdateSettingBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateSettingNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateSettingInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateSettingBadRequest creates a UpdateSettingBadRequest with default headers values
func NewUpdateSettingBadRequest() *UpdateSettingBadRequest {
	return &UpdateSettingBadRequest{}
}

/*
UpdateSettingBadRequest describes a response with status code 400, with default header values.

Bad Request due to a validation error.
*/
type UpdateSettingBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this update setting bad request response has a 2xx status code
func (o *UpdateSettingBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update setting bad request response has a 3xx status code
func (o *UpdateSettingBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update setting bad request response has a 4xx status code
func (o *UpdateSettingBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update setting bad request response has a 5xx status code
func (o *UpdateSettingBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update setting bad request response a status code equal to that given
func (o *UpdateSettingBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update setting bad request response
func (o *UpdateSettingBadRequest) Code() int {
	return 400
}

func (o *UpdateSettingBadRequest) Error() string {
	return fmt.Sprintf("[PUT /Settings('{input}')][%d] updateSettingBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateSettingBadRequest) String() string {
	return fmt.Sprintf("[PUT /Settings('{input}')][%d] updateSettingBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateSettingBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateSettingBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSettingNotFound creates a UpdateSettingNotFound with default headers values
func NewUpdateSettingNotFound() *UpdateSettingNotFound {
	return &UpdateSettingNotFound{}
}

/*
UpdateSettingNotFound describes a response with status code 404, with default header values.

Entity not found
*/
type UpdateSettingNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this update setting not found response has a 2xx status code
func (o *UpdateSettingNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update setting not found response has a 3xx status code
func (o *UpdateSettingNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update setting not found response has a 4xx status code
func (o *UpdateSettingNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update setting not found response has a 5xx status code
func (o *UpdateSettingNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update setting not found response a status code equal to that given
func (o *UpdateSettingNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update setting not found response
func (o *UpdateSettingNotFound) Code() int {
	return 404
}

func (o *UpdateSettingNotFound) Error() string {
	return fmt.Sprintf("[PUT /Settings('{input}')][%d] updateSettingNotFound  %+v", 404, o.Payload)
}

func (o *UpdateSettingNotFound) String() string {
	return fmt.Sprintf("[PUT /Settings('{input}')][%d] updateSettingNotFound  %+v", 404, o.Payload)
}

func (o *UpdateSettingNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateSettingNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSettingInternalServerError creates a UpdateSettingInternalServerError with default headers values
func NewUpdateSettingInternalServerError() *UpdateSettingInternalServerError {
	return &UpdateSettingInternalServerError{}
}

/*
UpdateSettingInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. Certain fields are missing from the request.
*/
type UpdateSettingInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this update setting internal server error response has a 2xx status code
func (o *UpdateSettingInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update setting internal server error response has a 3xx status code
func (o *UpdateSettingInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update setting internal server error response has a 4xx status code
func (o *UpdateSettingInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update setting internal server error response has a 5xx status code
func (o *UpdateSettingInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update setting internal server error response a status code equal to that given
func (o *UpdateSettingInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update setting internal server error response
func (o *UpdateSettingInternalServerError) Code() int {
	return 500
}

func (o *UpdateSettingInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /Settings('{input}')][%d] updateSettingInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateSettingInternalServerError) String() string {
	return fmt.Sprintf("[PUT /Settings('{input}')][%d] updateSettingInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateSettingInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateSettingInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}