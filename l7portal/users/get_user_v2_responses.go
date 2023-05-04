// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/tsrob/l7portal-go/models"
)

// GetUserV2Reader is a Reader for the GetUserV2 structure.
type GetUserV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUserV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetUserV2BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUserV2InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUserV2OK creates a GetUserV2OK with default headers values
func NewGetUserV2OK() *GetUserV2OK {
	return &GetUserV2OK{}
}

/*
GetUserV2OK describes a response with status code 200, with default header values.

An object describing a single User.
*/
type GetUserV2OK struct {
	Payload *models.UserV2
}

// IsSuccess returns true when this get user v2 o k response has a 2xx status code
func (o *GetUserV2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get user v2 o k response has a 3xx status code
func (o *GetUserV2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user v2 o k response has a 4xx status code
func (o *GetUserV2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get user v2 o k response has a 5xx status code
func (o *GetUserV2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get user v2 o k response a status code equal to that given
func (o *GetUserV2OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get user v2 o k response
func (o *GetUserV2OK) Code() int {
	return 200
}

func (o *GetUserV2OK) Error() string {
	return fmt.Sprintf("[GET /v2/users/{uuid}][%d] getUserV2OK  %+v", 200, o.Payload)
}

func (o *GetUserV2OK) String() string {
	return fmt.Sprintf("[GET /v2/users/{uuid}][%d] getUserV2OK  %+v", 200, o.Payload)
}

func (o *GetUserV2OK) GetPayload() *models.UserV2 {
	return o.Payload
}

func (o *GetUserV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserV2BadRequest creates a GetUserV2BadRequest with default headers values
func NewGetUserV2BadRequest() *GetUserV2BadRequest {
	return &GetUserV2BadRequest{}
}

/*
GetUserV2BadRequest describes a response with status code 400, with default header values.

Bad Request due to a validation error.
*/
type GetUserV2BadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get user v2 bad request response has a 2xx status code
func (o *GetUserV2BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user v2 bad request response has a 3xx status code
func (o *GetUserV2BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user v2 bad request response has a 4xx status code
func (o *GetUserV2BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user v2 bad request response has a 5xx status code
func (o *GetUserV2BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get user v2 bad request response a status code equal to that given
func (o *GetUserV2BadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get user v2 bad request response
func (o *GetUserV2BadRequest) Code() int {
	return 400
}

func (o *GetUserV2BadRequest) Error() string {
	return fmt.Sprintf("[GET /v2/users/{uuid}][%d] getUserV2BadRequest  %+v", 400, o.Payload)
}

func (o *GetUserV2BadRequest) String() string {
	return fmt.Sprintf("[GET /v2/users/{uuid}][%d] getUserV2BadRequest  %+v", 400, o.Payload)
}

func (o *GetUserV2BadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetUserV2BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserV2InternalServerError creates a GetUserV2InternalServerError with default headers values
func NewGetUserV2InternalServerError() *GetUserV2InternalServerError {
	return &GetUserV2InternalServerError{}
}

/*
GetUserV2InternalServerError describes a response with status code 500, with default header values.

Internal Server Error. Certain fields are missing from the request.
*/
type GetUserV2InternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get user v2 internal server error response has a 2xx status code
func (o *GetUserV2InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user v2 internal server error response has a 3xx status code
func (o *GetUserV2InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user v2 internal server error response has a 4xx status code
func (o *GetUserV2InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get user v2 internal server error response has a 5xx status code
func (o *GetUserV2InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get user v2 internal server error response a status code equal to that given
func (o *GetUserV2InternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get user v2 internal server error response
func (o *GetUserV2InternalServerError) Code() int {
	return 500
}

func (o *GetUserV2InternalServerError) Error() string {
	return fmt.Sprintf("[GET /v2/users/{uuid}][%d] getUserV2InternalServerError  %+v", 500, o.Payload)
}

func (o *GetUserV2InternalServerError) String() string {
	return fmt.Sprintf("[GET /v2/users/{uuid}][%d] getUserV2InternalServerError  %+v", 500, o.Payload)
}

func (o *GetUserV2InternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetUserV2InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
