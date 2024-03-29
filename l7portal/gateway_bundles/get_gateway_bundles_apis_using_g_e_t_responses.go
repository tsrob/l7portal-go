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

// GetGatewayBundlesApisUsingGETReader is a Reader for the GetGatewayBundlesApisUsingGET structure.
type GetGatewayBundlesApisUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGatewayBundlesApisUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGatewayBundlesApisUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetGatewayBundlesApisUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetGatewayBundlesApisUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetGatewayBundlesApisUsingGETOK creates a GetGatewayBundlesApisUsingGETOK with default headers values
func NewGetGatewayBundlesApisUsingGETOK() *GetGatewayBundlesApisUsingGETOK {
	return &GetGatewayBundlesApisUsingGETOK{}
}

/*
GetGatewayBundlesApisUsingGETOK describes a response with status code 200, with default header values.

Success
*/
type GetGatewayBundlesApisUsingGETOK struct {
	Payload *models.GatewayBundleAPIListDto
}

// IsSuccess returns true when this get gateway bundles apis using g e t o k response has a 2xx status code
func (o *GetGatewayBundlesApisUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get gateway bundles apis using g e t o k response has a 3xx status code
func (o *GetGatewayBundlesApisUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gateway bundles apis using g e t o k response has a 4xx status code
func (o *GetGatewayBundlesApisUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get gateway bundles apis using g e t o k response has a 5xx status code
func (o *GetGatewayBundlesApisUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get gateway bundles apis using g e t o k response a status code equal to that given
func (o *GetGatewayBundlesApisUsingGETOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get gateway bundles apis using g e t o k response
func (o *GetGatewayBundlesApisUsingGETOK) Code() int {
	return 200
}

func (o *GetGatewayBundlesApisUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /policy-management/0.1/gateway-bundles/{uuid}/apis][%d] getGatewayBundlesApisUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetGatewayBundlesApisUsingGETOK) String() string {
	return fmt.Sprintf("[GET /policy-management/0.1/gateway-bundles/{uuid}/apis][%d] getGatewayBundlesApisUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetGatewayBundlesApisUsingGETOK) GetPayload() *models.GatewayBundleAPIListDto {
	return o.Payload
}

func (o *GetGatewayBundlesApisUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayBundleAPIListDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGatewayBundlesApisUsingGETBadRequest creates a GetGatewayBundlesApisUsingGETBadRequest with default headers values
func NewGetGatewayBundlesApisUsingGETBadRequest() *GetGatewayBundlesApisUsingGETBadRequest {
	return &GetGatewayBundlesApisUsingGETBadRequest{}
}

/*
GetGatewayBundlesApisUsingGETBadRequest describes a response with status code 400, with default header values.

Bad Request due to a validation error.
*/
type GetGatewayBundlesApisUsingGETBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get gateway bundles apis using g e t bad request response has a 2xx status code
func (o *GetGatewayBundlesApisUsingGETBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get gateway bundles apis using g e t bad request response has a 3xx status code
func (o *GetGatewayBundlesApisUsingGETBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gateway bundles apis using g e t bad request response has a 4xx status code
func (o *GetGatewayBundlesApisUsingGETBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get gateway bundles apis using g e t bad request response has a 5xx status code
func (o *GetGatewayBundlesApisUsingGETBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get gateway bundles apis using g e t bad request response a status code equal to that given
func (o *GetGatewayBundlesApisUsingGETBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get gateway bundles apis using g e t bad request response
func (o *GetGatewayBundlesApisUsingGETBadRequest) Code() int {
	return 400
}

func (o *GetGatewayBundlesApisUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /policy-management/0.1/gateway-bundles/{uuid}/apis][%d] getGatewayBundlesApisUsingGETBadRequest  %+v", 400, o.Payload)
}

func (o *GetGatewayBundlesApisUsingGETBadRequest) String() string {
	return fmt.Sprintf("[GET /policy-management/0.1/gateway-bundles/{uuid}/apis][%d] getGatewayBundlesApisUsingGETBadRequest  %+v", 400, o.Payload)
}

func (o *GetGatewayBundlesApisUsingGETBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetGatewayBundlesApisUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGatewayBundlesApisUsingGETNotFound creates a GetGatewayBundlesApisUsingGETNotFound with default headers values
func NewGetGatewayBundlesApisUsingGETNotFound() *GetGatewayBundlesApisUsingGETNotFound {
	return &GetGatewayBundlesApisUsingGETNotFound{}
}

/*
GetGatewayBundlesApisUsingGETNotFound describes a response with status code 404, with default header values.

Entity not found
*/
type GetGatewayBundlesApisUsingGETNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get gateway bundles apis using g e t not found response has a 2xx status code
func (o *GetGatewayBundlesApisUsingGETNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get gateway bundles apis using g e t not found response has a 3xx status code
func (o *GetGatewayBundlesApisUsingGETNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gateway bundles apis using g e t not found response has a 4xx status code
func (o *GetGatewayBundlesApisUsingGETNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get gateway bundles apis using g e t not found response has a 5xx status code
func (o *GetGatewayBundlesApisUsingGETNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get gateway bundles apis using g e t not found response a status code equal to that given
func (o *GetGatewayBundlesApisUsingGETNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get gateway bundles apis using g e t not found response
func (o *GetGatewayBundlesApisUsingGETNotFound) Code() int {
	return 404
}

func (o *GetGatewayBundlesApisUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /policy-management/0.1/gateway-bundles/{uuid}/apis][%d] getGatewayBundlesApisUsingGETNotFound  %+v", 404, o.Payload)
}

func (o *GetGatewayBundlesApisUsingGETNotFound) String() string {
	return fmt.Sprintf("[GET /policy-management/0.1/gateway-bundles/{uuid}/apis][%d] getGatewayBundlesApisUsingGETNotFound  %+v", 404, o.Payload)
}

func (o *GetGatewayBundlesApisUsingGETNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetGatewayBundlesApisUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
