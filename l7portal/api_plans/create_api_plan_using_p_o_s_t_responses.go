// Code generated by go-swagger; DO NOT EDIT.

package api_plans

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/tsrob/l7portal-go/models"
)

// CreateAPIPlanUsingPOSTReader is a Reader for the CreateAPIPlanUsingPOST structure.
type CreateAPIPlanUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAPIPlanUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateAPIPlanUsingPOSTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateAPIPlanUsingPOSTBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateAPIPlanUsingPOSTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateAPIPlanUsingPOSTInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateAPIPlanUsingPOSTCreated creates a CreateAPIPlanUsingPOSTCreated with default headers values
func NewCreateAPIPlanUsingPOSTCreated() *CreateAPIPlanUsingPOSTCreated {
	return &CreateAPIPlanUsingPOSTCreated{}
}

/*
CreateAPIPlanUsingPOSTCreated describes a response with status code 201, with default header values.

CreateAPIPlanUsingPOSTCreated create Api plan using p o s t created
*/
type CreateAPIPlanUsingPOSTCreated struct {
	Payload *models.APIPlan10Create
}

// IsSuccess returns true when this create Api plan using p o s t created response has a 2xx status code
func (o *CreateAPIPlanUsingPOSTCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create Api plan using p o s t created response has a 3xx status code
func (o *CreateAPIPlanUsingPOSTCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create Api plan using p o s t created response has a 4xx status code
func (o *CreateAPIPlanUsingPOSTCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create Api plan using p o s t created response has a 5xx status code
func (o *CreateAPIPlanUsingPOSTCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create Api plan using p o s t created response a status code equal to that given
func (o *CreateAPIPlanUsingPOSTCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create Api plan using p o s t created response
func (o *CreateAPIPlanUsingPOSTCreated) Code() int {
	return 201
}

func (o *CreateAPIPlanUsingPOSTCreated) Error() string {
	return fmt.Sprintf("[POST /api-management/1.0/api-plans][%d] createApiPlanUsingPOSTCreated  %+v", 201, o.Payload)
}

func (o *CreateAPIPlanUsingPOSTCreated) String() string {
	return fmt.Sprintf("[POST /api-management/1.0/api-plans][%d] createApiPlanUsingPOSTCreated  %+v", 201, o.Payload)
}

func (o *CreateAPIPlanUsingPOSTCreated) GetPayload() *models.APIPlan10Create {
	return o.Payload
}

func (o *CreateAPIPlanUsingPOSTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIPlan10Create)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAPIPlanUsingPOSTBadRequest creates a CreateAPIPlanUsingPOSTBadRequest with default headers values
func NewCreateAPIPlanUsingPOSTBadRequest() *CreateAPIPlanUsingPOSTBadRequest {
	return &CreateAPIPlanUsingPOSTBadRequest{}
}

/*
CreateAPIPlanUsingPOSTBadRequest describes a response with status code 400, with default header values.

Bad Request due to Invalid Uuid.
*/
type CreateAPIPlanUsingPOSTBadRequest struct {
	Payload *models.LongerError
}

// IsSuccess returns true when this create Api plan using p o s t bad request response has a 2xx status code
func (o *CreateAPIPlanUsingPOSTBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create Api plan using p o s t bad request response has a 3xx status code
func (o *CreateAPIPlanUsingPOSTBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create Api plan using p o s t bad request response has a 4xx status code
func (o *CreateAPIPlanUsingPOSTBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create Api plan using p o s t bad request response has a 5xx status code
func (o *CreateAPIPlanUsingPOSTBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create Api plan using p o s t bad request response a status code equal to that given
func (o *CreateAPIPlanUsingPOSTBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create Api plan using p o s t bad request response
func (o *CreateAPIPlanUsingPOSTBadRequest) Code() int {
	return 400
}

func (o *CreateAPIPlanUsingPOSTBadRequest) Error() string {
	return fmt.Sprintf("[POST /api-management/1.0/api-plans][%d] createApiPlanUsingPOSTBadRequest  %+v", 400, o.Payload)
}

func (o *CreateAPIPlanUsingPOSTBadRequest) String() string {
	return fmt.Sprintf("[POST /api-management/1.0/api-plans][%d] createApiPlanUsingPOSTBadRequest  %+v", 400, o.Payload)
}

func (o *CreateAPIPlanUsingPOSTBadRequest) GetPayload() *models.LongerError {
	return o.Payload
}

func (o *CreateAPIPlanUsingPOSTBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LongerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAPIPlanUsingPOSTNotFound creates a CreateAPIPlanUsingPOSTNotFound with default headers values
func NewCreateAPIPlanUsingPOSTNotFound() *CreateAPIPlanUsingPOSTNotFound {
	return &CreateAPIPlanUsingPOSTNotFound{}
}

/*
CreateAPIPlanUsingPOSTNotFound describes a response with status code 404, with default header values.

Entity not found
*/
type CreateAPIPlanUsingPOSTNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this create Api plan using p o s t not found response has a 2xx status code
func (o *CreateAPIPlanUsingPOSTNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create Api plan using p o s t not found response has a 3xx status code
func (o *CreateAPIPlanUsingPOSTNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create Api plan using p o s t not found response has a 4xx status code
func (o *CreateAPIPlanUsingPOSTNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create Api plan using p o s t not found response has a 5xx status code
func (o *CreateAPIPlanUsingPOSTNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create Api plan using p o s t not found response a status code equal to that given
func (o *CreateAPIPlanUsingPOSTNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the create Api plan using p o s t not found response
func (o *CreateAPIPlanUsingPOSTNotFound) Code() int {
	return 404
}

func (o *CreateAPIPlanUsingPOSTNotFound) Error() string {
	return fmt.Sprintf("[POST /api-management/1.0/api-plans][%d] createApiPlanUsingPOSTNotFound  %+v", 404, o.Payload)
}

func (o *CreateAPIPlanUsingPOSTNotFound) String() string {
	return fmt.Sprintf("[POST /api-management/1.0/api-plans][%d] createApiPlanUsingPOSTNotFound  %+v", 404, o.Payload)
}

func (o *CreateAPIPlanUsingPOSTNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateAPIPlanUsingPOSTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAPIPlanUsingPOSTInternalServerError creates a CreateAPIPlanUsingPOSTInternalServerError with default headers values
func NewCreateAPIPlanUsingPOSTInternalServerError() *CreateAPIPlanUsingPOSTInternalServerError {
	return &CreateAPIPlanUsingPOSTInternalServerError{}
}

/*
CreateAPIPlanUsingPOSTInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. Certain fields are missing from the request.
*/
type CreateAPIPlanUsingPOSTInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this create Api plan using p o s t internal server error response has a 2xx status code
func (o *CreateAPIPlanUsingPOSTInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create Api plan using p o s t internal server error response has a 3xx status code
func (o *CreateAPIPlanUsingPOSTInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create Api plan using p o s t internal server error response has a 4xx status code
func (o *CreateAPIPlanUsingPOSTInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create Api plan using p o s t internal server error response has a 5xx status code
func (o *CreateAPIPlanUsingPOSTInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create Api plan using p o s t internal server error response a status code equal to that given
func (o *CreateAPIPlanUsingPOSTInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create Api plan using p o s t internal server error response
func (o *CreateAPIPlanUsingPOSTInternalServerError) Code() int {
	return 500
}

func (o *CreateAPIPlanUsingPOSTInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api-management/1.0/api-plans][%d] createApiPlanUsingPOSTInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateAPIPlanUsingPOSTInternalServerError) String() string {
	return fmt.Sprintf("[POST /api-management/1.0/api-plans][%d] createApiPlanUsingPOSTInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateAPIPlanUsingPOSTInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateAPIPlanUsingPOSTInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
