// Code generated by go-swagger; DO NOT EDIT.

package themes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/tsrob/l7portal-go/models"
)

// UpdateThemeReader is a Reader for the UpdateTheme structure.
type UpdateThemeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateThemeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 400:
		result := NewUpdateThemeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateThemeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateThemeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateThemeBadRequest creates a UpdateThemeBadRequest with default headers values
func NewUpdateThemeBadRequest() *UpdateThemeBadRequest {
	return &UpdateThemeBadRequest{}
}

/*
UpdateThemeBadRequest describes a response with status code 400, with default header values.

Bad Request due to a validation error.
*/
type UpdateThemeBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this update theme bad request response has a 2xx status code
func (o *UpdateThemeBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update theme bad request response has a 3xx status code
func (o *UpdateThemeBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update theme bad request response has a 4xx status code
func (o *UpdateThemeBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update theme bad request response has a 5xx status code
func (o *UpdateThemeBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update theme bad request response a status code equal to that given
func (o *UpdateThemeBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update theme bad request response
func (o *UpdateThemeBadRequest) Code() int {
	return 400
}

func (o *UpdateThemeBadRequest) Error() string {
	return fmt.Sprintf("[PUT /branding/1.0/themes][%d] updateThemeBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateThemeBadRequest) String() string {
	return fmt.Sprintf("[PUT /branding/1.0/themes][%d] updateThemeBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateThemeBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateThemeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateThemeNotFound creates a UpdateThemeNotFound with default headers values
func NewUpdateThemeNotFound() *UpdateThemeNotFound {
	return &UpdateThemeNotFound{}
}

/*
UpdateThemeNotFound describes a response with status code 404, with default header values.

Entity not found
*/
type UpdateThemeNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this update theme not found response has a 2xx status code
func (o *UpdateThemeNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update theme not found response has a 3xx status code
func (o *UpdateThemeNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update theme not found response has a 4xx status code
func (o *UpdateThemeNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update theme not found response has a 5xx status code
func (o *UpdateThemeNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update theme not found response a status code equal to that given
func (o *UpdateThemeNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update theme not found response
func (o *UpdateThemeNotFound) Code() int {
	return 404
}

func (o *UpdateThemeNotFound) Error() string {
	return fmt.Sprintf("[PUT /branding/1.0/themes][%d] updateThemeNotFound  %+v", 404, o.Payload)
}

func (o *UpdateThemeNotFound) String() string {
	return fmt.Sprintf("[PUT /branding/1.0/themes][%d] updateThemeNotFound  %+v", 404, o.Payload)
}

func (o *UpdateThemeNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateThemeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateThemeInternalServerError creates a UpdateThemeInternalServerError with default headers values
func NewUpdateThemeInternalServerError() *UpdateThemeInternalServerError {
	return &UpdateThemeInternalServerError{}
}

/*
UpdateThemeInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. Certain fields are missing from the request.
*/
type UpdateThemeInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this update theme internal server error response has a 2xx status code
func (o *UpdateThemeInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update theme internal server error response has a 3xx status code
func (o *UpdateThemeInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update theme internal server error response has a 4xx status code
func (o *UpdateThemeInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update theme internal server error response has a 5xx status code
func (o *UpdateThemeInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update theme internal server error response a status code equal to that given
func (o *UpdateThemeInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update theme internal server error response
func (o *UpdateThemeInternalServerError) Code() int {
	return 500
}

func (o *UpdateThemeInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /branding/1.0/themes][%d] updateThemeInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateThemeInternalServerError) String() string {
	return fmt.Sprintf("[PUT /branding/1.0/themes][%d] updateThemeInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateThemeInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateThemeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
