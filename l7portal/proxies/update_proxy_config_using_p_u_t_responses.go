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

// UpdateProxyConfigUsingPUTReader is a Reader for the UpdateProxyConfigUsingPUT structure.
type UpdateProxyConfigUsingPUTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateProxyConfigUsingPUTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateProxyConfigUsingPUTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewUpdateProxyConfigUsingPUTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateProxyConfigUsingPUTOK creates a UpdateProxyConfigUsingPUTOK with default headers values
func NewUpdateProxyConfigUsingPUTOK() *UpdateProxyConfigUsingPUTOK {
	return &UpdateProxyConfigUsingPUTOK{}
}

/*
UpdateProxyConfigUsingPUTOK describes a response with status code 200, with default header values.

Success
*/
type UpdateProxyConfigUsingPUTOK struct {
	Payload *models.ProxyConfigDto
}

// IsSuccess returns true when this update proxy config using p u t o k response has a 2xx status code
func (o *UpdateProxyConfigUsingPUTOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update proxy config using p u t o k response has a 3xx status code
func (o *UpdateProxyConfigUsingPUTOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update proxy config using p u t o k response has a 4xx status code
func (o *UpdateProxyConfigUsingPUTOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update proxy config using p u t o k response has a 5xx status code
func (o *UpdateProxyConfigUsingPUTOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update proxy config using p u t o k response a status code equal to that given
func (o *UpdateProxyConfigUsingPUTOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update proxy config using p u t o k response
func (o *UpdateProxyConfigUsingPUTOK) Code() int {
	return 200
}

func (o *UpdateProxyConfigUsingPUTOK) Error() string {
	return fmt.Sprintf("[PUT /deployments/0.1/proxies/{uuid}/config][%d] updateProxyConfigUsingPUTOK  %+v", 200, o.Payload)
}

func (o *UpdateProxyConfigUsingPUTOK) String() string {
	return fmt.Sprintf("[PUT /deployments/0.1/proxies/{uuid}/config][%d] updateProxyConfigUsingPUTOK  %+v", 200, o.Payload)
}

func (o *UpdateProxyConfigUsingPUTOK) GetPayload() *models.ProxyConfigDto {
	return o.Payload
}

func (o *UpdateProxyConfigUsingPUTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProxyConfigDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateProxyConfigUsingPUTNotFound creates a UpdateProxyConfigUsingPUTNotFound with default headers values
func NewUpdateProxyConfigUsingPUTNotFound() *UpdateProxyConfigUsingPUTNotFound {
	return &UpdateProxyConfigUsingPUTNotFound{}
}

/*
UpdateProxyConfigUsingPUTNotFound describes a response with status code 404, with default header values.

Entity not found
*/
type UpdateProxyConfigUsingPUTNotFound struct {
}

// IsSuccess returns true when this update proxy config using p u t not found response has a 2xx status code
func (o *UpdateProxyConfigUsingPUTNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update proxy config using p u t not found response has a 3xx status code
func (o *UpdateProxyConfigUsingPUTNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update proxy config using p u t not found response has a 4xx status code
func (o *UpdateProxyConfigUsingPUTNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update proxy config using p u t not found response has a 5xx status code
func (o *UpdateProxyConfigUsingPUTNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update proxy config using p u t not found response a status code equal to that given
func (o *UpdateProxyConfigUsingPUTNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update proxy config using p u t not found response
func (o *UpdateProxyConfigUsingPUTNotFound) Code() int {
	return 404
}

func (o *UpdateProxyConfigUsingPUTNotFound) Error() string {
	return fmt.Sprintf("[PUT /deployments/0.1/proxies/{uuid}/config][%d] updateProxyConfigUsingPUTNotFound ", 404)
}

func (o *UpdateProxyConfigUsingPUTNotFound) String() string {
	return fmt.Sprintf("[PUT /deployments/0.1/proxies/{uuid}/config][%d] updateProxyConfigUsingPUTNotFound ", 404)
}

func (o *UpdateProxyConfigUsingPUTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
