// Code generated by go-swagger; DO NOT EDIT.

package gateway_bundle_deployments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/tsrob/l7portal-go/models"
)

// GetGatewayBundleDeploymentDetailsUsingGETReader is a Reader for the GetGatewayBundleDeploymentDetailsUsingGET structure.
type GetGatewayBundleDeploymentDetailsUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGatewayBundleDeploymentDetailsUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGatewayBundleDeploymentDetailsUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetGatewayBundleDeploymentDetailsUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetGatewayBundleDeploymentDetailsUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetGatewayBundleDeploymentDetailsUsingGETOK creates a GetGatewayBundleDeploymentDetailsUsingGETOK with default headers values
func NewGetGatewayBundleDeploymentDetailsUsingGETOK() *GetGatewayBundleDeploymentDetailsUsingGETOK {
	return &GetGatewayBundleDeploymentDetailsUsingGETOK{}
}

/*
GetGatewayBundleDeploymentDetailsUsingGETOK describes a response with status code 200, with default header values.

Success
*/
type GetGatewayBundleDeploymentDetailsUsingGETOK struct {
	Payload *models.GatewayBundleDeploymentBasicDto
}

// IsSuccess returns true when this get gateway bundle deployment details using g e t o k response has a 2xx status code
func (o *GetGatewayBundleDeploymentDetailsUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get gateway bundle deployment details using g e t o k response has a 3xx status code
func (o *GetGatewayBundleDeploymentDetailsUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gateway bundle deployment details using g e t o k response has a 4xx status code
func (o *GetGatewayBundleDeploymentDetailsUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get gateway bundle deployment details using g e t o k response has a 5xx status code
func (o *GetGatewayBundleDeploymentDetailsUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get gateway bundle deployment details using g e t o k response a status code equal to that given
func (o *GetGatewayBundleDeploymentDetailsUsingGETOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get gateway bundle deployment details using g e t o k response
func (o *GetGatewayBundleDeploymentDetailsUsingGETOK) Code() int {
	return 200
}

func (o *GetGatewayBundleDeploymentDetailsUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /deployments/0.1/gateway-bundles/{uuid}/proxies/{proxyUuid}][%d] getGatewayBundleDeploymentDetailsUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetGatewayBundleDeploymentDetailsUsingGETOK) String() string {
	return fmt.Sprintf("[GET /deployments/0.1/gateway-bundles/{uuid}/proxies/{proxyUuid}][%d] getGatewayBundleDeploymentDetailsUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetGatewayBundleDeploymentDetailsUsingGETOK) GetPayload() *models.GatewayBundleDeploymentBasicDto {
	return o.Payload
}

func (o *GetGatewayBundleDeploymentDetailsUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayBundleDeploymentBasicDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGatewayBundleDeploymentDetailsUsingGETBadRequest creates a GetGatewayBundleDeploymentDetailsUsingGETBadRequest with default headers values
func NewGetGatewayBundleDeploymentDetailsUsingGETBadRequest() *GetGatewayBundleDeploymentDetailsUsingGETBadRequest {
	return &GetGatewayBundleDeploymentDetailsUsingGETBadRequest{}
}

/*
GetGatewayBundleDeploymentDetailsUsingGETBadRequest describes a response with status code 400, with default header values.

Bad Request due to a validation error.
*/
type GetGatewayBundleDeploymentDetailsUsingGETBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get gateway bundle deployment details using g e t bad request response has a 2xx status code
func (o *GetGatewayBundleDeploymentDetailsUsingGETBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get gateway bundle deployment details using g e t bad request response has a 3xx status code
func (o *GetGatewayBundleDeploymentDetailsUsingGETBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gateway bundle deployment details using g e t bad request response has a 4xx status code
func (o *GetGatewayBundleDeploymentDetailsUsingGETBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get gateway bundle deployment details using g e t bad request response has a 5xx status code
func (o *GetGatewayBundleDeploymentDetailsUsingGETBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get gateway bundle deployment details using g e t bad request response a status code equal to that given
func (o *GetGatewayBundleDeploymentDetailsUsingGETBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get gateway bundle deployment details using g e t bad request response
func (o *GetGatewayBundleDeploymentDetailsUsingGETBadRequest) Code() int {
	return 400
}

func (o *GetGatewayBundleDeploymentDetailsUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /deployments/0.1/gateway-bundles/{uuid}/proxies/{proxyUuid}][%d] getGatewayBundleDeploymentDetailsUsingGETBadRequest  %+v", 400, o.Payload)
}

func (o *GetGatewayBundleDeploymentDetailsUsingGETBadRequest) String() string {
	return fmt.Sprintf("[GET /deployments/0.1/gateway-bundles/{uuid}/proxies/{proxyUuid}][%d] getGatewayBundleDeploymentDetailsUsingGETBadRequest  %+v", 400, o.Payload)
}

func (o *GetGatewayBundleDeploymentDetailsUsingGETBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetGatewayBundleDeploymentDetailsUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGatewayBundleDeploymentDetailsUsingGETNotFound creates a GetGatewayBundleDeploymentDetailsUsingGETNotFound with default headers values
func NewGetGatewayBundleDeploymentDetailsUsingGETNotFound() *GetGatewayBundleDeploymentDetailsUsingGETNotFound {
	return &GetGatewayBundleDeploymentDetailsUsingGETNotFound{}
}

/*
GetGatewayBundleDeploymentDetailsUsingGETNotFound describes a response with status code 404, with default header values.

Entity not found
*/
type GetGatewayBundleDeploymentDetailsUsingGETNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get gateway bundle deployment details using g e t not found response has a 2xx status code
func (o *GetGatewayBundleDeploymentDetailsUsingGETNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get gateway bundle deployment details using g e t not found response has a 3xx status code
func (o *GetGatewayBundleDeploymentDetailsUsingGETNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gateway bundle deployment details using g e t not found response has a 4xx status code
func (o *GetGatewayBundleDeploymentDetailsUsingGETNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get gateway bundle deployment details using g e t not found response has a 5xx status code
func (o *GetGatewayBundleDeploymentDetailsUsingGETNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get gateway bundle deployment details using g e t not found response a status code equal to that given
func (o *GetGatewayBundleDeploymentDetailsUsingGETNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get gateway bundle deployment details using g e t not found response
func (o *GetGatewayBundleDeploymentDetailsUsingGETNotFound) Code() int {
	return 404
}

func (o *GetGatewayBundleDeploymentDetailsUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /deployments/0.1/gateway-bundles/{uuid}/proxies/{proxyUuid}][%d] getGatewayBundleDeploymentDetailsUsingGETNotFound  %+v", 404, o.Payload)
}

func (o *GetGatewayBundleDeploymentDetailsUsingGETNotFound) String() string {
	return fmt.Sprintf("[GET /deployments/0.1/gateway-bundles/{uuid}/proxies/{proxyUuid}][%d] getGatewayBundleDeploymentDetailsUsingGETNotFound  %+v", 404, o.Payload)
}

func (o *GetGatewayBundleDeploymentDetailsUsingGETNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetGatewayBundleDeploymentDetailsUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}