// Code generated by go-swagger; DO NOT EDIT.

package apis

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/tsrob/l7portal-go/models"
)

// GetApi10Reader is a Reader for the GetApi10 structure.
type GetApi10Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetApi10Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApi10OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetApi10BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetApi10NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetApi10InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetApi10OK creates a GetApi10OK with default headers values
func NewGetApi10OK() *GetApi10OK {
	return &GetApi10OK{}
}

/*
GetApi10OK describes a response with status code 200, with default header values.

GetApi10OK get api10 o k
*/
type GetApi10OK struct {
	Payload *models.APIGet
}

// IsSuccess returns true when this get api10 o k response has a 2xx status code
func (o *GetApi10OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get api10 o k response has a 3xx status code
func (o *GetApi10OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get api10 o k response has a 4xx status code
func (o *GetApi10OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get api10 o k response has a 5xx status code
func (o *GetApi10OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get api10 o k response a status code equal to that given
func (o *GetApi10OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get api10 o k response
func (o *GetApi10OK) Code() int {
	return 200
}

func (o *GetApi10OK) Error() string {
	return fmt.Sprintf("[GET /api-management/1.0/apis/{apiUuid}][%d] getApi10OK  %+v", 200, o.Payload)
}

func (o *GetApi10OK) String() string {
	return fmt.Sprintf("[GET /api-management/1.0/apis/{apiUuid}][%d] getApi10OK  %+v", 200, o.Payload)
}

func (o *GetApi10OK) GetPayload() *models.APIGet {
	return o.Payload
}

func (o *GetApi10OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIGet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApi10BadRequest creates a GetApi10BadRequest with default headers values
func NewGetApi10BadRequest() *GetApi10BadRequest {
	return &GetApi10BadRequest{}
}

/*
GetApi10BadRequest describes a response with status code 400, with default header values.

Bad Request due to Invalid Uuid.
*/
type GetApi10BadRequest struct {
	Payload *models.LongerError
}

// IsSuccess returns true when this get api10 bad request response has a 2xx status code
func (o *GetApi10BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get api10 bad request response has a 3xx status code
func (o *GetApi10BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get api10 bad request response has a 4xx status code
func (o *GetApi10BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get api10 bad request response has a 5xx status code
func (o *GetApi10BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get api10 bad request response a status code equal to that given
func (o *GetApi10BadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get api10 bad request response
func (o *GetApi10BadRequest) Code() int {
	return 400
}

func (o *GetApi10BadRequest) Error() string {
	return fmt.Sprintf("[GET /api-management/1.0/apis/{apiUuid}][%d] getApi10BadRequest  %+v", 400, o.Payload)
}

func (o *GetApi10BadRequest) String() string {
	return fmt.Sprintf("[GET /api-management/1.0/apis/{apiUuid}][%d] getApi10BadRequest  %+v", 400, o.Payload)
}

func (o *GetApi10BadRequest) GetPayload() *models.LongerError {
	return o.Payload
}

func (o *GetApi10BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LongerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApi10NotFound creates a GetApi10NotFound with default headers values
func NewGetApi10NotFound() *GetApi10NotFound {
	return &GetApi10NotFound{}
}

/*
GetApi10NotFound describes a response with status code 404, with default header values.

Entity not Found
*/
type GetApi10NotFound struct {
	Payload *models.ErrorEntity
}

// IsSuccess returns true when this get api10 not found response has a 2xx status code
func (o *GetApi10NotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get api10 not found response has a 3xx status code
func (o *GetApi10NotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get api10 not found response has a 4xx status code
func (o *GetApi10NotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get api10 not found response has a 5xx status code
func (o *GetApi10NotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get api10 not found response a status code equal to that given
func (o *GetApi10NotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get api10 not found response
func (o *GetApi10NotFound) Code() int {
	return 404
}

func (o *GetApi10NotFound) Error() string {
	return fmt.Sprintf("[GET /api-management/1.0/apis/{apiUuid}][%d] getApi10NotFound  %+v", 404, o.Payload)
}

func (o *GetApi10NotFound) String() string {
	return fmt.Sprintf("[GET /api-management/1.0/apis/{apiUuid}][%d] getApi10NotFound  %+v", 404, o.Payload)
}

func (o *GetApi10NotFound) GetPayload() *models.ErrorEntity {
	return o.Payload
}

func (o *GetApi10NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApi10InternalServerError creates a GetApi10InternalServerError with default headers values
func NewGetApi10InternalServerError() *GetApi10InternalServerError {
	return &GetApi10InternalServerError{}
}

/*
GetApi10InternalServerError describes a response with status code 500, with default header values.

Internal Server Error. Certain fields are missing from the request.
*/
type GetApi10InternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get api10 internal server error response has a 2xx status code
func (o *GetApi10InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get api10 internal server error response has a 3xx status code
func (o *GetApi10InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get api10 internal server error response has a 4xx status code
func (o *GetApi10InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get api10 internal server error response has a 5xx status code
func (o *GetApi10InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get api10 internal server error response a status code equal to that given
func (o *GetApi10InternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get api10 internal server error response
func (o *GetApi10InternalServerError) Code() int {
	return 500
}

func (o *GetApi10InternalServerError) Error() string {
	return fmt.Sprintf("[GET /api-management/1.0/apis/{apiUuid}][%d] getApi10InternalServerError  %+v", 500, o.Payload)
}

func (o *GetApi10InternalServerError) String() string {
	return fmt.Sprintf("[GET /api-management/1.0/apis/{apiUuid}][%d] getApi10InternalServerError  %+v", 500, o.Payload)
}

func (o *GetApi10InternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetApi10InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
