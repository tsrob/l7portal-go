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

// GetApisOfApplicationByUUIDReader is a Reader for the GetApisOfApplicationByUUID structure.
type GetApisOfApplicationByUUIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetApisOfApplicationByUUIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApisOfApplicationByUUIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetApisOfApplicationByUUIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetApisOfApplicationByUUIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetApisOfApplicationByUUIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetApisOfApplicationByUUIDOK creates a GetApisOfApplicationByUUIDOK with default headers values
func NewGetApisOfApplicationByUUIDOK() *GetApisOfApplicationByUUIDOK {
	return &GetApisOfApplicationByUUIDOK{}
}

/*
GetApisOfApplicationByUUIDOK describes a response with status code 200, with default header values.

OK
*/
type GetApisOfApplicationByUUIDOK struct {
	Payload []*models.APIReference
}

// IsSuccess returns true when this get apis of application by Uuid o k response has a 2xx status code
func (o *GetApisOfApplicationByUUIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get apis of application by Uuid o k response has a 3xx status code
func (o *GetApisOfApplicationByUUIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get apis of application by Uuid o k response has a 4xx status code
func (o *GetApisOfApplicationByUUIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get apis of application by Uuid o k response has a 5xx status code
func (o *GetApisOfApplicationByUUIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get apis of application by Uuid o k response a status code equal to that given
func (o *GetApisOfApplicationByUUIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get apis of application by Uuid o k response
func (o *GetApisOfApplicationByUUIDOK) Code() int {
	return 200
}

func (o *GetApisOfApplicationByUUIDOK) Error() string {
	return fmt.Sprintf("[GET /api-management/1.0/applications/{appUuid}/apis][%d] getApisOfApplicationByUuidOK  %+v", 200, o.Payload)
}

func (o *GetApisOfApplicationByUUIDOK) String() string {
	return fmt.Sprintf("[GET /api-management/1.0/applications/{appUuid}/apis][%d] getApisOfApplicationByUuidOK  %+v", 200, o.Payload)
}

func (o *GetApisOfApplicationByUUIDOK) GetPayload() []*models.APIReference {
	return o.Payload
}

func (o *GetApisOfApplicationByUUIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApisOfApplicationByUUIDBadRequest creates a GetApisOfApplicationByUUIDBadRequest with default headers values
func NewGetApisOfApplicationByUUIDBadRequest() *GetApisOfApplicationByUUIDBadRequest {
	return &GetApisOfApplicationByUUIDBadRequest{}
}

/*
GetApisOfApplicationByUUIDBadRequest describes a response with status code 400, with default header values.

Bad Request due to Invalid Uuid.
*/
type GetApisOfApplicationByUUIDBadRequest struct {
	Payload *models.LongerError
}

// IsSuccess returns true when this get apis of application by Uuid bad request response has a 2xx status code
func (o *GetApisOfApplicationByUUIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get apis of application by Uuid bad request response has a 3xx status code
func (o *GetApisOfApplicationByUUIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get apis of application by Uuid bad request response has a 4xx status code
func (o *GetApisOfApplicationByUUIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get apis of application by Uuid bad request response has a 5xx status code
func (o *GetApisOfApplicationByUUIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get apis of application by Uuid bad request response a status code equal to that given
func (o *GetApisOfApplicationByUUIDBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get apis of application by Uuid bad request response
func (o *GetApisOfApplicationByUUIDBadRequest) Code() int {
	return 400
}

func (o *GetApisOfApplicationByUUIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /api-management/1.0/applications/{appUuid}/apis][%d] getApisOfApplicationByUuidBadRequest  %+v", 400, o.Payload)
}

func (o *GetApisOfApplicationByUUIDBadRequest) String() string {
	return fmt.Sprintf("[GET /api-management/1.0/applications/{appUuid}/apis][%d] getApisOfApplicationByUuidBadRequest  %+v", 400, o.Payload)
}

func (o *GetApisOfApplicationByUUIDBadRequest) GetPayload() *models.LongerError {
	return o.Payload
}

func (o *GetApisOfApplicationByUUIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LongerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApisOfApplicationByUUIDNotFound creates a GetApisOfApplicationByUUIDNotFound with default headers values
func NewGetApisOfApplicationByUUIDNotFound() *GetApisOfApplicationByUUIDNotFound {
	return &GetApisOfApplicationByUUIDNotFound{}
}

/*
GetApisOfApplicationByUUIDNotFound describes a response with status code 404, with default header values.

Entity not Found
*/
type GetApisOfApplicationByUUIDNotFound struct {
	Payload *models.ErrorEntity
}

// IsSuccess returns true when this get apis of application by Uuid not found response has a 2xx status code
func (o *GetApisOfApplicationByUUIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get apis of application by Uuid not found response has a 3xx status code
func (o *GetApisOfApplicationByUUIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get apis of application by Uuid not found response has a 4xx status code
func (o *GetApisOfApplicationByUUIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get apis of application by Uuid not found response has a 5xx status code
func (o *GetApisOfApplicationByUUIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get apis of application by Uuid not found response a status code equal to that given
func (o *GetApisOfApplicationByUUIDNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get apis of application by Uuid not found response
func (o *GetApisOfApplicationByUUIDNotFound) Code() int {
	return 404
}

func (o *GetApisOfApplicationByUUIDNotFound) Error() string {
	return fmt.Sprintf("[GET /api-management/1.0/applications/{appUuid}/apis][%d] getApisOfApplicationByUuidNotFound  %+v", 404, o.Payload)
}

func (o *GetApisOfApplicationByUUIDNotFound) String() string {
	return fmt.Sprintf("[GET /api-management/1.0/applications/{appUuid}/apis][%d] getApisOfApplicationByUuidNotFound  %+v", 404, o.Payload)
}

func (o *GetApisOfApplicationByUUIDNotFound) GetPayload() *models.ErrorEntity {
	return o.Payload
}

func (o *GetApisOfApplicationByUUIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApisOfApplicationByUUIDInternalServerError creates a GetApisOfApplicationByUUIDInternalServerError with default headers values
func NewGetApisOfApplicationByUUIDInternalServerError() *GetApisOfApplicationByUUIDInternalServerError {
	return &GetApisOfApplicationByUUIDInternalServerError{}
}

/*
GetApisOfApplicationByUUIDInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. Certain fields are missing from the request.
*/
type GetApisOfApplicationByUUIDInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get apis of application by Uuid internal server error response has a 2xx status code
func (o *GetApisOfApplicationByUUIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get apis of application by Uuid internal server error response has a 3xx status code
func (o *GetApisOfApplicationByUUIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get apis of application by Uuid internal server error response has a 4xx status code
func (o *GetApisOfApplicationByUUIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get apis of application by Uuid internal server error response has a 5xx status code
func (o *GetApisOfApplicationByUUIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get apis of application by Uuid internal server error response a status code equal to that given
func (o *GetApisOfApplicationByUUIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get apis of application by Uuid internal server error response
func (o *GetApisOfApplicationByUUIDInternalServerError) Code() int {
	return 500
}

func (o *GetApisOfApplicationByUUIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api-management/1.0/applications/{appUuid}/apis][%d] getApisOfApplicationByUuidInternalServerError  %+v", 500, o.Payload)
}

func (o *GetApisOfApplicationByUUIDInternalServerError) String() string {
	return fmt.Sprintf("[GET /api-management/1.0/applications/{appUuid}/apis][%d] getApisOfApplicationByUuidInternalServerError  %+v", 500, o.Payload)
}

func (o *GetApisOfApplicationByUUIDInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetApisOfApplicationByUUIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
