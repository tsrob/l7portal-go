// Code generated by go-swagger; DO NOT EDIT.

package gateway_bundles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/tsrob/l7portal-go/models"
)

// GetGatewayBundleDetailsUsingGETReader is a Reader for the GetGatewayBundleDetailsUsingGET structure.
type GetGatewayBundleDetailsUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGatewayBundleDetailsUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGatewayBundleDetailsUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetGatewayBundleDetailsUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetGatewayBundleDetailsUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetGatewayBundleDetailsUsingGETOK creates a GetGatewayBundleDetailsUsingGETOK with default headers values
func NewGetGatewayBundleDetailsUsingGETOK() *GetGatewayBundleDetailsUsingGETOK {
	return &GetGatewayBundleDetailsUsingGETOK{}
}

/*
GetGatewayBundleDetailsUsingGETOK describes a response with status code 200, with default header values.

Success
*/
type GetGatewayBundleDetailsUsingGETOK struct {
	Payload *models.GatewayBundleDetailsDto
}

// IsSuccess returns true when this get gateway bundle details using g e t o k response has a 2xx status code
func (o *GetGatewayBundleDetailsUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get gateway bundle details using g e t o k response has a 3xx status code
func (o *GetGatewayBundleDetailsUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gateway bundle details using g e t o k response has a 4xx status code
func (o *GetGatewayBundleDetailsUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get gateway bundle details using g e t o k response has a 5xx status code
func (o *GetGatewayBundleDetailsUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get gateway bundle details using g e t o k response a status code equal to that given
func (o *GetGatewayBundleDetailsUsingGETOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get gateway bundle details using g e t o k response
func (o *GetGatewayBundleDetailsUsingGETOK) Code() int {
	return 200
}

func (o *GetGatewayBundleDetailsUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /policy-management/0.1/gateway-bundles/{uuid}][%d] getGatewayBundleDetailsUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetGatewayBundleDetailsUsingGETOK) String() string {
	return fmt.Sprintf("[GET /policy-management/0.1/gateway-bundles/{uuid}][%d] getGatewayBundleDetailsUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetGatewayBundleDetailsUsingGETOK) GetPayload() *models.GatewayBundleDetailsDto {
	return o.Payload
}

func (o *GetGatewayBundleDetailsUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayBundleDetailsDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGatewayBundleDetailsUsingGETBadRequest creates a GetGatewayBundleDetailsUsingGETBadRequest with default headers values
func NewGetGatewayBundleDetailsUsingGETBadRequest() *GetGatewayBundleDetailsUsingGETBadRequest {
	return &GetGatewayBundleDetailsUsingGETBadRequest{}
}

/*
GetGatewayBundleDetailsUsingGETBadRequest describes a response with status code 400, with default header values.

Bad Request due to a validation error.
*/
type GetGatewayBundleDetailsUsingGETBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get gateway bundle details using g e t bad request response has a 2xx status code
func (o *GetGatewayBundleDetailsUsingGETBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get gateway bundle details using g e t bad request response has a 3xx status code
func (o *GetGatewayBundleDetailsUsingGETBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gateway bundle details using g e t bad request response has a 4xx status code
func (o *GetGatewayBundleDetailsUsingGETBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get gateway bundle details using g e t bad request response has a 5xx status code
func (o *GetGatewayBundleDetailsUsingGETBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get gateway bundle details using g e t bad request response a status code equal to that given
func (o *GetGatewayBundleDetailsUsingGETBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get gateway bundle details using g e t bad request response
func (o *GetGatewayBundleDetailsUsingGETBadRequest) Code() int {
	return 400
}

func (o *GetGatewayBundleDetailsUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /policy-management/0.1/gateway-bundles/{uuid}][%d] getGatewayBundleDetailsUsingGETBadRequest  %+v", 400, o.Payload)
}

func (o *GetGatewayBundleDetailsUsingGETBadRequest) String() string {
	return fmt.Sprintf("[GET /policy-management/0.1/gateway-bundles/{uuid}][%d] getGatewayBundleDetailsUsingGETBadRequest  %+v", 400, o.Payload)
}

func (o *GetGatewayBundleDetailsUsingGETBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetGatewayBundleDetailsUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGatewayBundleDetailsUsingGETNotFound creates a GetGatewayBundleDetailsUsingGETNotFound with default headers values
func NewGetGatewayBundleDetailsUsingGETNotFound() *GetGatewayBundleDetailsUsingGETNotFound {
	return &GetGatewayBundleDetailsUsingGETNotFound{}
}

/*
GetGatewayBundleDetailsUsingGETNotFound describes a response with status code 404, with default header values.

Entity not found
*/
type GetGatewayBundleDetailsUsingGETNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get gateway bundle details using g e t not found response has a 2xx status code
func (o *GetGatewayBundleDetailsUsingGETNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get gateway bundle details using g e t not found response has a 3xx status code
func (o *GetGatewayBundleDetailsUsingGETNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gateway bundle details using g e t not found response has a 4xx status code
func (o *GetGatewayBundleDetailsUsingGETNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get gateway bundle details using g e t not found response has a 5xx status code
func (o *GetGatewayBundleDetailsUsingGETNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get gateway bundle details using g e t not found response a status code equal to that given
func (o *GetGatewayBundleDetailsUsingGETNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get gateway bundle details using g e t not found response
func (o *GetGatewayBundleDetailsUsingGETNotFound) Code() int {
	return 404
}

func (o *GetGatewayBundleDetailsUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /policy-management/0.1/gateway-bundles/{uuid}][%d] getGatewayBundleDetailsUsingGETNotFound  %+v", 404, o.Payload)
}

func (o *GetGatewayBundleDetailsUsingGETNotFound) String() string {
	return fmt.Sprintf("[GET /policy-management/0.1/gateway-bundles/{uuid}][%d] getGatewayBundleDetailsUsingGETNotFound  %+v", 404, o.Payload)
}

func (o *GetGatewayBundleDetailsUsingGETNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetGatewayBundleDetailsUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}