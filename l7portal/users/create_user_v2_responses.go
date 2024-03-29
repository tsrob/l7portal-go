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

// CreateUserV2Reader is a Reader for the CreateUserV2 structure.
type CreateUserV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateUserV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateUserV2Created()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateUserV2BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateUserV2InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateUserV2Created creates a CreateUserV2Created with default headers values
func NewCreateUserV2Created() *CreateUserV2Created {
	return &CreateUserV2Created{}
}

/*
CreateUserV2Created describes a response with status code 201, with default header values.

An object describing a single newly created User.
*/
type CreateUserV2Created struct {
	Payload *models.UserV2
}

// IsSuccess returns true when this create user v2 created response has a 2xx status code
func (o *CreateUserV2Created) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create user v2 created response has a 3xx status code
func (o *CreateUserV2Created) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create user v2 created response has a 4xx status code
func (o *CreateUserV2Created) IsClientError() bool {
	return false
}

// IsServerError returns true when this create user v2 created response has a 5xx status code
func (o *CreateUserV2Created) IsServerError() bool {
	return false
}

// IsCode returns true when this create user v2 created response a status code equal to that given
func (o *CreateUserV2Created) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create user v2 created response
func (o *CreateUserV2Created) Code() int {
	return 201
}

func (o *CreateUserV2Created) Error() string {
	return fmt.Sprintf("[POST /v2/users][%d] createUserV2Created  %+v", 201, o.Payload)
}

func (o *CreateUserV2Created) String() string {
	return fmt.Sprintf("[POST /v2/users][%d] createUserV2Created  %+v", 201, o.Payload)
}

func (o *CreateUserV2Created) GetPayload() *models.UserV2 {
	return o.Payload
}

func (o *CreateUserV2Created) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUserV2BadRequest creates a CreateUserV2BadRequest with default headers values
func NewCreateUserV2BadRequest() *CreateUserV2BadRequest {
	return &CreateUserV2BadRequest{}
}

/*
CreateUserV2BadRequest describes a response with status code 400, with default header values.

Bad Request due to a validation error.
*/
type CreateUserV2BadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this create user v2 bad request response has a 2xx status code
func (o *CreateUserV2BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create user v2 bad request response has a 3xx status code
func (o *CreateUserV2BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create user v2 bad request response has a 4xx status code
func (o *CreateUserV2BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create user v2 bad request response has a 5xx status code
func (o *CreateUserV2BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create user v2 bad request response a status code equal to that given
func (o *CreateUserV2BadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create user v2 bad request response
func (o *CreateUserV2BadRequest) Code() int {
	return 400
}

func (o *CreateUserV2BadRequest) Error() string {
	return fmt.Sprintf("[POST /v2/users][%d] createUserV2BadRequest  %+v", 400, o.Payload)
}

func (o *CreateUserV2BadRequest) String() string {
	return fmt.Sprintf("[POST /v2/users][%d] createUserV2BadRequest  %+v", 400, o.Payload)
}

func (o *CreateUserV2BadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateUserV2BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUserV2InternalServerError creates a CreateUserV2InternalServerError with default headers values
func NewCreateUserV2InternalServerError() *CreateUserV2InternalServerError {
	return &CreateUserV2InternalServerError{}
}

/*
CreateUserV2InternalServerError describes a response with status code 500, with default header values.

Internal Server Error. Certain fields are missing from the request.
*/
type CreateUserV2InternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this create user v2 internal server error response has a 2xx status code
func (o *CreateUserV2InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create user v2 internal server error response has a 3xx status code
func (o *CreateUserV2InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create user v2 internal server error response has a 4xx status code
func (o *CreateUserV2InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create user v2 internal server error response has a 5xx status code
func (o *CreateUserV2InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create user v2 internal server error response a status code equal to that given
func (o *CreateUserV2InternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create user v2 internal server error response
func (o *CreateUserV2InternalServerError) Code() int {
	return 500
}

func (o *CreateUserV2InternalServerError) Error() string {
	return fmt.Sprintf("[POST /v2/users][%d] createUserV2InternalServerError  %+v", 500, o.Payload)
}

func (o *CreateUserV2InternalServerError) String() string {
	return fmt.Sprintf("[POST /v2/users][%d] createUserV2InternalServerError  %+v", 500, o.Payload)
}

func (o *CreateUserV2InternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateUserV2InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
