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

// GetGatewayBundlesUsingGETReader is a Reader for the GetGatewayBundlesUsingGET structure.
type GetGatewayBundlesUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGatewayBundlesUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGatewayBundlesUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetGatewayBundlesUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetGatewayBundlesUsingGETOK creates a GetGatewayBundlesUsingGETOK with default headers values
func NewGetGatewayBundlesUsingGETOK() *GetGatewayBundlesUsingGETOK {
	return &GetGatewayBundlesUsingGETOK{}
}

/*
GetGatewayBundlesUsingGETOK describes a response with status code 200, with default header values.

Success
*/
type GetGatewayBundlesUsingGETOK struct {
	Payload *models.GatewayBundleDto
}

// IsSuccess returns true when this get gateway bundles using g e t o k response has a 2xx status code
func (o *GetGatewayBundlesUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get gateway bundles using g e t o k response has a 3xx status code
func (o *GetGatewayBundlesUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gateway bundles using g e t o k response has a 4xx status code
func (o *GetGatewayBundlesUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get gateway bundles using g e t o k response has a 5xx status code
func (o *GetGatewayBundlesUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get gateway bundles using g e t o k response a status code equal to that given
func (o *GetGatewayBundlesUsingGETOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get gateway bundles using g e t o k response
func (o *GetGatewayBundlesUsingGETOK) Code() int {
	return 200
}

func (o *GetGatewayBundlesUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /policy-management/0.1/gateway-bundles][%d] getGatewayBundlesUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetGatewayBundlesUsingGETOK) String() string {
	return fmt.Sprintf("[GET /policy-management/0.1/gateway-bundles][%d] getGatewayBundlesUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetGatewayBundlesUsingGETOK) GetPayload() *models.GatewayBundleDto {
	return o.Payload
}

func (o *GetGatewayBundlesUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayBundleDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGatewayBundlesUsingGETBadRequest creates a GetGatewayBundlesUsingGETBadRequest with default headers values
func NewGetGatewayBundlesUsingGETBadRequest() *GetGatewayBundlesUsingGETBadRequest {
	return &GetGatewayBundlesUsingGETBadRequest{}
}

/*
GetGatewayBundlesUsingGETBadRequest describes a response with status code 400, with default header values.

Bad Request due to a validation error.
*/
type GetGatewayBundlesUsingGETBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get gateway bundles using g e t bad request response has a 2xx status code
func (o *GetGatewayBundlesUsingGETBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get gateway bundles using g e t bad request response has a 3xx status code
func (o *GetGatewayBundlesUsingGETBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gateway bundles using g e t bad request response has a 4xx status code
func (o *GetGatewayBundlesUsingGETBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get gateway bundles using g e t bad request response has a 5xx status code
func (o *GetGatewayBundlesUsingGETBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get gateway bundles using g e t bad request response a status code equal to that given
func (o *GetGatewayBundlesUsingGETBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get gateway bundles using g e t bad request response
func (o *GetGatewayBundlesUsingGETBadRequest) Code() int {
	return 400
}

func (o *GetGatewayBundlesUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /policy-management/0.1/gateway-bundles][%d] getGatewayBundlesUsingGETBadRequest  %+v", 400, o.Payload)
}

func (o *GetGatewayBundlesUsingGETBadRequest) String() string {
	return fmt.Sprintf("[GET /policy-management/0.1/gateway-bundles][%d] getGatewayBundlesUsingGETBadRequest  %+v", 400, o.Payload)
}

func (o *GetGatewayBundlesUsingGETBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetGatewayBundlesUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
