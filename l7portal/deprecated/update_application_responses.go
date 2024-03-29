// Code generated by go-swagger; DO NOT EDIT.

package deprecated

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/tsrob/l7portal-go/models"
)

// UpdateApplicationReader is a Reader for the UpdateApplication structure.
type UpdateApplicationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateApplicationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateApplicationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateApplicationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateApplicationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateApplicationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateApplicationOK creates a UpdateApplicationOK with default headers values
func NewUpdateApplicationOK() *UpdateApplicationOK {
	return &UpdateApplicationOK{}
}

/*
UpdateApplicationOK describes a response with status code 200, with default header values.

An object describing updating a single Application.
*/
type UpdateApplicationOK struct {
	Payload *models.Application
}

// IsSuccess returns true when this update application o k response has a 2xx status code
func (o *UpdateApplicationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update application o k response has a 3xx status code
func (o *UpdateApplicationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update application o k response has a 4xx status code
func (o *UpdateApplicationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update application o k response has a 5xx status code
func (o *UpdateApplicationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update application o k response a status code equal to that given
func (o *UpdateApplicationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update application o k response
func (o *UpdateApplicationOK) Code() int {
	return 200
}

func (o *UpdateApplicationOK) Error() string {
	return fmt.Sprintf("[PUT /Applications('{uuid}')][%d] updateApplicationOK  %+v", 200, o.Payload)
}

func (o *UpdateApplicationOK) String() string {
	return fmt.Sprintf("[PUT /Applications('{uuid}')][%d] updateApplicationOK  %+v", 200, o.Payload)
}

func (o *UpdateApplicationOK) GetPayload() *models.Application {
	return o.Payload
}

func (o *UpdateApplicationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Application)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateApplicationBadRequest creates a UpdateApplicationBadRequest with default headers values
func NewUpdateApplicationBadRequest() *UpdateApplicationBadRequest {
	return &UpdateApplicationBadRequest{}
}

/*
UpdateApplicationBadRequest describes a response with status code 400, with default header values.

Bad Request due to a validation error.
*/
type UpdateApplicationBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this update application bad request response has a 2xx status code
func (o *UpdateApplicationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update application bad request response has a 3xx status code
func (o *UpdateApplicationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update application bad request response has a 4xx status code
func (o *UpdateApplicationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update application bad request response has a 5xx status code
func (o *UpdateApplicationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update application bad request response a status code equal to that given
func (o *UpdateApplicationBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update application bad request response
func (o *UpdateApplicationBadRequest) Code() int {
	return 400
}

func (o *UpdateApplicationBadRequest) Error() string {
	return fmt.Sprintf("[PUT /Applications('{uuid}')][%d] updateApplicationBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateApplicationBadRequest) String() string {
	return fmt.Sprintf("[PUT /Applications('{uuid}')][%d] updateApplicationBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateApplicationBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateApplicationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateApplicationNotFound creates a UpdateApplicationNotFound with default headers values
func NewUpdateApplicationNotFound() *UpdateApplicationNotFound {
	return &UpdateApplicationNotFound{}
}

/*
UpdateApplicationNotFound describes a response with status code 404, with default header values.

Entity not found
*/
type UpdateApplicationNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this update application not found response has a 2xx status code
func (o *UpdateApplicationNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update application not found response has a 3xx status code
func (o *UpdateApplicationNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update application not found response has a 4xx status code
func (o *UpdateApplicationNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update application not found response has a 5xx status code
func (o *UpdateApplicationNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update application not found response a status code equal to that given
func (o *UpdateApplicationNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update application not found response
func (o *UpdateApplicationNotFound) Code() int {
	return 404
}

func (o *UpdateApplicationNotFound) Error() string {
	return fmt.Sprintf("[PUT /Applications('{uuid}')][%d] updateApplicationNotFound  %+v", 404, o.Payload)
}

func (o *UpdateApplicationNotFound) String() string {
	return fmt.Sprintf("[PUT /Applications('{uuid}')][%d] updateApplicationNotFound  %+v", 404, o.Payload)
}

func (o *UpdateApplicationNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateApplicationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateApplicationInternalServerError creates a UpdateApplicationInternalServerError with default headers values
func NewUpdateApplicationInternalServerError() *UpdateApplicationInternalServerError {
	return &UpdateApplicationInternalServerError{}
}

/*
UpdateApplicationInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. Certain fields are missing from the request.
*/
type UpdateApplicationInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this update application internal server error response has a 2xx status code
func (o *UpdateApplicationInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update application internal server error response has a 3xx status code
func (o *UpdateApplicationInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update application internal server error response has a 4xx status code
func (o *UpdateApplicationInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update application internal server error response has a 5xx status code
func (o *UpdateApplicationInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update application internal server error response a status code equal to that given
func (o *UpdateApplicationInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update application internal server error response
func (o *UpdateApplicationInternalServerError) Code() int {
	return 500
}

func (o *UpdateApplicationInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /Applications('{uuid}')][%d] updateApplicationInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateApplicationInternalServerError) String() string {
	return fmt.Sprintf("[PUT /Applications('{uuid}')][%d] updateApplicationInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateApplicationInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateApplicationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
