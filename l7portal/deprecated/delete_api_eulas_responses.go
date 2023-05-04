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

// DeleteAPIEulasReader is a Reader for the DeleteAPIEulas structure.
type DeleteAPIEulasReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAPIEulasReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 400:
		result := NewDeleteAPIEulasBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteAPIEulasNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteAPIEulasInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteAPIEulasBadRequest creates a DeleteAPIEulasBadRequest with default headers values
func NewDeleteAPIEulasBadRequest() *DeleteAPIEulasBadRequest {
	return &DeleteAPIEulasBadRequest{}
}

/*
DeleteAPIEulasBadRequest describes a response with status code 400, with default header values.

Bad Request due to a validation error.
*/
type DeleteAPIEulasBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this delete Api eulas bad request response has a 2xx status code
func (o *DeleteAPIEulasBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete Api eulas bad request response has a 3xx status code
func (o *DeleteAPIEulasBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete Api eulas bad request response has a 4xx status code
func (o *DeleteAPIEulasBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete Api eulas bad request response has a 5xx status code
func (o *DeleteAPIEulasBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete Api eulas bad request response a status code equal to that given
func (o *DeleteAPIEulasBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete Api eulas bad request response
func (o *DeleteAPIEulasBadRequest) Code() int {
	return 400
}

func (o *DeleteAPIEulasBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /ApiEulas('{uuid}')][%d] deleteApiEulasBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteAPIEulasBadRequest) String() string {
	return fmt.Sprintf("[DELETE /ApiEulas('{uuid}')][%d] deleteApiEulasBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteAPIEulasBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteAPIEulasBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAPIEulasNotFound creates a DeleteAPIEulasNotFound with default headers values
func NewDeleteAPIEulasNotFound() *DeleteAPIEulasNotFound {
	return &DeleteAPIEulasNotFound{}
}

/*
DeleteAPIEulasNotFound describes a response with status code 404, with default header values.

Entity not found
*/
type DeleteAPIEulasNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this delete Api eulas not found response has a 2xx status code
func (o *DeleteAPIEulasNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete Api eulas not found response has a 3xx status code
func (o *DeleteAPIEulasNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete Api eulas not found response has a 4xx status code
func (o *DeleteAPIEulasNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete Api eulas not found response has a 5xx status code
func (o *DeleteAPIEulasNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete Api eulas not found response a status code equal to that given
func (o *DeleteAPIEulasNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete Api eulas not found response
func (o *DeleteAPIEulasNotFound) Code() int {
	return 404
}

func (o *DeleteAPIEulasNotFound) Error() string {
	return fmt.Sprintf("[DELETE /ApiEulas('{uuid}')][%d] deleteApiEulasNotFound  %+v", 404, o.Payload)
}

func (o *DeleteAPIEulasNotFound) String() string {
	return fmt.Sprintf("[DELETE /ApiEulas('{uuid}')][%d] deleteApiEulasNotFound  %+v", 404, o.Payload)
}

func (o *DeleteAPIEulasNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteAPIEulasNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAPIEulasInternalServerError creates a DeleteAPIEulasInternalServerError with default headers values
func NewDeleteAPIEulasInternalServerError() *DeleteAPIEulasInternalServerError {
	return &DeleteAPIEulasInternalServerError{}
}

/*
DeleteAPIEulasInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. Certain fields are missing from the request.
*/
type DeleteAPIEulasInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this delete Api eulas internal server error response has a 2xx status code
func (o *DeleteAPIEulasInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete Api eulas internal server error response has a 3xx status code
func (o *DeleteAPIEulasInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete Api eulas internal server error response has a 4xx status code
func (o *DeleteAPIEulasInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete Api eulas internal server error response has a 5xx status code
func (o *DeleteAPIEulasInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete Api eulas internal server error response a status code equal to that given
func (o *DeleteAPIEulasInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete Api eulas internal server error response
func (o *DeleteAPIEulasInternalServerError) Code() int {
	return 500
}

func (o *DeleteAPIEulasInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /ApiEulas('{uuid}')][%d] deleteApiEulasInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteAPIEulasInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /ApiEulas('{uuid}')][%d] deleteApiEulasInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteAPIEulasInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteAPIEulasInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}