// Code generated by go-swagger; DO NOT EDIT.

package tags

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/tsrob/l7portal-go/models"
)

// PUTTagOrganizationsReader is a Reader for the PUTTagOrganizations structure.
type PUTTagOrganizationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PUTTagOrganizationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPUTTagOrganizationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPUTTagOrganizationsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPUTTagOrganizationsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPUTTagOrganizationsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPUTTagOrganizationsOK creates a PUTTagOrganizationsOK with default headers values
func NewPUTTagOrganizationsOK() *PUTTagOrganizationsOK {
	return &PUTTagOrganizationsOK{}
}

/*
PUTTagOrganizationsOK describes a response with status code 200, with default header values.

OK
*/
type PUTTagOrganizationsOK struct {
	Payload []string
}

// IsSuccess returns true when this p u t tag organizations o k response has a 2xx status code
func (o *PUTTagOrganizationsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this p u t tag organizations o k response has a 3xx status code
func (o *PUTTagOrganizationsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this p u t tag organizations o k response has a 4xx status code
func (o *PUTTagOrganizationsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this p u t tag organizations o k response has a 5xx status code
func (o *PUTTagOrganizationsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this p u t tag organizations o k response a status code equal to that given
func (o *PUTTagOrganizationsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the p u t tag organizations o k response
func (o *PUTTagOrganizationsOK) Code() int {
	return 200
}

func (o *PUTTagOrganizationsOK) Error() string {
	return fmt.Sprintf("[PUT /tenant-admin/1.0/tags/{tagUuid}/organizations][%d] pUTTagOrganizationsOK  %+v", 200, o.Payload)
}

func (o *PUTTagOrganizationsOK) String() string {
	return fmt.Sprintf("[PUT /tenant-admin/1.0/tags/{tagUuid}/organizations][%d] pUTTagOrganizationsOK  %+v", 200, o.Payload)
}

func (o *PUTTagOrganizationsOK) GetPayload() []string {
	return o.Payload
}

func (o *PUTTagOrganizationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPUTTagOrganizationsBadRequest creates a PUTTagOrganizationsBadRequest with default headers values
func NewPUTTagOrganizationsBadRequest() *PUTTagOrganizationsBadRequest {
	return &PUTTagOrganizationsBadRequest{}
}

/*
PUTTagOrganizationsBadRequest describes a response with status code 400, with default header values.

Bad Request due to Invalid Uuid.
*/
type PUTTagOrganizationsBadRequest struct {
	Payload *models.LongerError
}

// IsSuccess returns true when this p u t tag organizations bad request response has a 2xx status code
func (o *PUTTagOrganizationsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this p u t tag organizations bad request response has a 3xx status code
func (o *PUTTagOrganizationsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this p u t tag organizations bad request response has a 4xx status code
func (o *PUTTagOrganizationsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this p u t tag organizations bad request response has a 5xx status code
func (o *PUTTagOrganizationsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this p u t tag organizations bad request response a status code equal to that given
func (o *PUTTagOrganizationsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the p u t tag organizations bad request response
func (o *PUTTagOrganizationsBadRequest) Code() int {
	return 400
}

func (o *PUTTagOrganizationsBadRequest) Error() string {
	return fmt.Sprintf("[PUT /tenant-admin/1.0/tags/{tagUuid}/organizations][%d] pUTTagOrganizationsBadRequest  %+v", 400, o.Payload)
}

func (o *PUTTagOrganizationsBadRequest) String() string {
	return fmt.Sprintf("[PUT /tenant-admin/1.0/tags/{tagUuid}/organizations][%d] pUTTagOrganizationsBadRequest  %+v", 400, o.Payload)
}

func (o *PUTTagOrganizationsBadRequest) GetPayload() *models.LongerError {
	return o.Payload
}

func (o *PUTTagOrganizationsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LongerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPUTTagOrganizationsNotFound creates a PUTTagOrganizationsNotFound with default headers values
func NewPUTTagOrganizationsNotFound() *PUTTagOrganizationsNotFound {
	return &PUTTagOrganizationsNotFound{}
}

/*
PUTTagOrganizationsNotFound describes a response with status code 404, with default header values.

Entity not Found
*/
type PUTTagOrganizationsNotFound struct {
	Payload *models.ErrorEntity
}

// IsSuccess returns true when this p u t tag organizations not found response has a 2xx status code
func (o *PUTTagOrganizationsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this p u t tag organizations not found response has a 3xx status code
func (o *PUTTagOrganizationsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this p u t tag organizations not found response has a 4xx status code
func (o *PUTTagOrganizationsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this p u t tag organizations not found response has a 5xx status code
func (o *PUTTagOrganizationsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this p u t tag organizations not found response a status code equal to that given
func (o *PUTTagOrganizationsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the p u t tag organizations not found response
func (o *PUTTagOrganizationsNotFound) Code() int {
	return 404
}

func (o *PUTTagOrganizationsNotFound) Error() string {
	return fmt.Sprintf("[PUT /tenant-admin/1.0/tags/{tagUuid}/organizations][%d] pUTTagOrganizationsNotFound  %+v", 404, o.Payload)
}

func (o *PUTTagOrganizationsNotFound) String() string {
	return fmt.Sprintf("[PUT /tenant-admin/1.0/tags/{tagUuid}/organizations][%d] pUTTagOrganizationsNotFound  %+v", 404, o.Payload)
}

func (o *PUTTagOrganizationsNotFound) GetPayload() *models.ErrorEntity {
	return o.Payload
}

func (o *PUTTagOrganizationsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPUTTagOrganizationsInternalServerError creates a PUTTagOrganizationsInternalServerError with default headers values
func NewPUTTagOrganizationsInternalServerError() *PUTTagOrganizationsInternalServerError {
	return &PUTTagOrganizationsInternalServerError{}
}

/*
PUTTagOrganizationsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. Certain fields are missing from the request.
*/
type PUTTagOrganizationsInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this p u t tag organizations internal server error response has a 2xx status code
func (o *PUTTagOrganizationsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this p u t tag organizations internal server error response has a 3xx status code
func (o *PUTTagOrganizationsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this p u t tag organizations internal server error response has a 4xx status code
func (o *PUTTagOrganizationsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this p u t tag organizations internal server error response has a 5xx status code
func (o *PUTTagOrganizationsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this p u t tag organizations internal server error response a status code equal to that given
func (o *PUTTagOrganizationsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the p u t tag organizations internal server error response
func (o *PUTTagOrganizationsInternalServerError) Code() int {
	return 500
}

func (o *PUTTagOrganizationsInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /tenant-admin/1.0/tags/{tagUuid}/organizations][%d] pUTTagOrganizationsInternalServerError  %+v", 500, o.Payload)
}

func (o *PUTTagOrganizationsInternalServerError) String() string {
	return fmt.Sprintf("[PUT /tenant-admin/1.0/tags/{tagUuid}/organizations][%d] pUTTagOrganizationsInternalServerError  %+v", 500, o.Payload)
}

func (o *PUTTagOrganizationsInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PUTTagOrganizationsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}