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

// GetApis10Reader is a Reader for the GetApis10 structure.
type GetApis10Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetApis10Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApis10OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetApis10BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetApis10InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetApis10OK creates a GetApis10OK with default headers values
func NewGetApis10OK() *GetApis10OK {
	return &GetApis10OK{}
}

/*
GetApis10OK describes a response with status code 200, with default header values.

GetApis10OK get apis10 o k
*/
type GetApis10OK struct {
	Payload *models.ApisGet
}

// IsSuccess returns true when this get apis10 o k response has a 2xx status code
func (o *GetApis10OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get apis10 o k response has a 3xx status code
func (o *GetApis10OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get apis10 o k response has a 4xx status code
func (o *GetApis10OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get apis10 o k response has a 5xx status code
func (o *GetApis10OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get apis10 o k response a status code equal to that given
func (o *GetApis10OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get apis10 o k response
func (o *GetApis10OK) Code() int {
	return 200
}

func (o *GetApis10OK) Error() string {
	return fmt.Sprintf("[GET /api-management/1.0/apis][%d] getApis10OK  %+v", 200, o.Payload)
}

func (o *GetApis10OK) String() string {
	return fmt.Sprintf("[GET /api-management/1.0/apis][%d] getApis10OK  %+v", 200, o.Payload)
}

func (o *GetApis10OK) GetPayload() *models.ApisGet {
	return o.Payload
}

func (o *GetApis10OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApisGet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApis10BadRequest creates a GetApis10BadRequest with default headers values
func NewGetApis10BadRequest() *GetApis10BadRequest {
	return &GetApis10BadRequest{}
}

/*
GetApis10BadRequest describes a response with status code 400, with default header values.

Bad Request due to Invalid Parameters.
*/
type GetApis10BadRequest struct {
	Payload *models.LongerError
}

// IsSuccess returns true when this get apis10 bad request response has a 2xx status code
func (o *GetApis10BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get apis10 bad request response has a 3xx status code
func (o *GetApis10BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get apis10 bad request response has a 4xx status code
func (o *GetApis10BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get apis10 bad request response has a 5xx status code
func (o *GetApis10BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get apis10 bad request response a status code equal to that given
func (o *GetApis10BadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get apis10 bad request response
func (o *GetApis10BadRequest) Code() int {
	return 400
}

func (o *GetApis10BadRequest) Error() string {
	return fmt.Sprintf("[GET /api-management/1.0/apis][%d] getApis10BadRequest  %+v", 400, o.Payload)
}

func (o *GetApis10BadRequest) String() string {
	return fmt.Sprintf("[GET /api-management/1.0/apis][%d] getApis10BadRequest  %+v", 400, o.Payload)
}

func (o *GetApis10BadRequest) GetPayload() *models.LongerError {
	return o.Payload
}

func (o *GetApis10BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LongerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApis10InternalServerError creates a GetApis10InternalServerError with default headers values
func NewGetApis10InternalServerError() *GetApis10InternalServerError {
	return &GetApis10InternalServerError{}
}

/*
GetApis10InternalServerError describes a response with status code 500, with default header values.

Internal Server Error. Certain fields are missing from the request.
*/
type GetApis10InternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get apis10 internal server error response has a 2xx status code
func (o *GetApis10InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get apis10 internal server error response has a 3xx status code
func (o *GetApis10InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get apis10 internal server error response has a 4xx status code
func (o *GetApis10InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get apis10 internal server error response has a 5xx status code
func (o *GetApis10InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get apis10 internal server error response a status code equal to that given
func (o *GetApis10InternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get apis10 internal server error response
func (o *GetApis10InternalServerError) Code() int {
	return 500
}

func (o *GetApis10InternalServerError) Error() string {
	return fmt.Sprintf("[GET /api-management/1.0/apis][%d] getApis10InternalServerError  %+v", 500, o.Payload)
}

func (o *GetApis10InternalServerError) String() string {
	return fmt.Sprintf("[GET /api-management/1.0/apis][%d] getApis10InternalServerError  %+v", 500, o.Payload)
}

func (o *GetApis10InternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetApis10InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
