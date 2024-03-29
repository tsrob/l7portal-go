// Code generated by go-swagger; DO NOT EDIT.

package proxies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/tsrob/l7portal-go/models"
)

// UpdateProxyUsingPUTReader is a Reader for the UpdateProxyUsingPUT structure.
type UpdateProxyUsingPUTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateProxyUsingPUTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateProxyUsingPUTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateProxyUsingPUTBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateProxyUsingPUTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateProxyUsingPUTOK creates a UpdateProxyUsingPUTOK with default headers values
func NewUpdateProxyUsingPUTOK() *UpdateProxyUsingPUTOK {
	return &UpdateProxyUsingPUTOK{}
}

/*
UpdateProxyUsingPUTOK describes a response with status code 200, with default header values.

Success
*/
type UpdateProxyUsingPUTOK struct {
	Payload *models.Proxy
}

// IsSuccess returns true when this update proxy using p u t o k response has a 2xx status code
func (o *UpdateProxyUsingPUTOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update proxy using p u t o k response has a 3xx status code
func (o *UpdateProxyUsingPUTOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update proxy using p u t o k response has a 4xx status code
func (o *UpdateProxyUsingPUTOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update proxy using p u t o k response has a 5xx status code
func (o *UpdateProxyUsingPUTOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update proxy using p u t o k response a status code equal to that given
func (o *UpdateProxyUsingPUTOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update proxy using p u t o k response
func (o *UpdateProxyUsingPUTOK) Code() int {
	return 200
}

func (o *UpdateProxyUsingPUTOK) Error() string {
	return fmt.Sprintf("[PUT /deployments/1.0/proxies/{uuid}][%d] updateProxyUsingPUTOK  %+v", 200, o.Payload)
}

func (o *UpdateProxyUsingPUTOK) String() string {
	return fmt.Sprintf("[PUT /deployments/1.0/proxies/{uuid}][%d] updateProxyUsingPUTOK  %+v", 200, o.Payload)
}

func (o *UpdateProxyUsingPUTOK) GetPayload() *models.Proxy {
	return o.Payload
}

func (o *UpdateProxyUsingPUTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Proxy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateProxyUsingPUTBadRequest creates a UpdateProxyUsingPUTBadRequest with default headers values
func NewUpdateProxyUsingPUTBadRequest() *UpdateProxyUsingPUTBadRequest {
	return &UpdateProxyUsingPUTBadRequest{}
}

/*
UpdateProxyUsingPUTBadRequest describes a response with status code 400, with default header values.

Bad request due to a validation error.
*/
type UpdateProxyUsingPUTBadRequest struct {
}

// IsSuccess returns true when this update proxy using p u t bad request response has a 2xx status code
func (o *UpdateProxyUsingPUTBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update proxy using p u t bad request response has a 3xx status code
func (o *UpdateProxyUsingPUTBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update proxy using p u t bad request response has a 4xx status code
func (o *UpdateProxyUsingPUTBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update proxy using p u t bad request response has a 5xx status code
func (o *UpdateProxyUsingPUTBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update proxy using p u t bad request response a status code equal to that given
func (o *UpdateProxyUsingPUTBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update proxy using p u t bad request response
func (o *UpdateProxyUsingPUTBadRequest) Code() int {
	return 400
}

func (o *UpdateProxyUsingPUTBadRequest) Error() string {
	return fmt.Sprintf("[PUT /deployments/1.0/proxies/{uuid}][%d] updateProxyUsingPUTBadRequest ", 400)
}

func (o *UpdateProxyUsingPUTBadRequest) String() string {
	return fmt.Sprintf("[PUT /deployments/1.0/proxies/{uuid}][%d] updateProxyUsingPUTBadRequest ", 400)
}

func (o *UpdateProxyUsingPUTBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateProxyUsingPUTNotFound creates a UpdateProxyUsingPUTNotFound with default headers values
func NewUpdateProxyUsingPUTNotFound() *UpdateProxyUsingPUTNotFound {
	return &UpdateProxyUsingPUTNotFound{}
}

/*
UpdateProxyUsingPUTNotFound describes a response with status code 404, with default header values.

Entity not found
*/
type UpdateProxyUsingPUTNotFound struct {
}

// IsSuccess returns true when this update proxy using p u t not found response has a 2xx status code
func (o *UpdateProxyUsingPUTNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update proxy using p u t not found response has a 3xx status code
func (o *UpdateProxyUsingPUTNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update proxy using p u t not found response has a 4xx status code
func (o *UpdateProxyUsingPUTNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update proxy using p u t not found response has a 5xx status code
func (o *UpdateProxyUsingPUTNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update proxy using p u t not found response a status code equal to that given
func (o *UpdateProxyUsingPUTNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update proxy using p u t not found response
func (o *UpdateProxyUsingPUTNotFound) Code() int {
	return 404
}

func (o *UpdateProxyUsingPUTNotFound) Error() string {
	return fmt.Sprintf("[PUT /deployments/1.0/proxies/{uuid}][%d] updateProxyUsingPUTNotFound ", 404)
}

func (o *UpdateProxyUsingPUTNotFound) String() string {
	return fmt.Sprintf("[PUT /deployments/1.0/proxies/{uuid}][%d] updateProxyUsingPUTNotFound ", 404)
}

func (o *UpdateProxyUsingPUTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
