// Code generated by go-swagger; DO NOT EDIT.

package api_key_deployments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/tsrob/l7portal-go/models"
)

// CreateAPIKeyDeploymentUsingPOST1Reader is a Reader for the CreateAPIKeyDeploymentUsingPOST1 structure.
type CreateAPIKeyDeploymentUsingPOST1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAPIKeyDeploymentUsingPOST1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateAPIKeyDeploymentUsingPOST1Created()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateAPIKeyDeploymentUsingPOST1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateAPIKeyDeploymentUsingPOST1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateAPIKeyDeploymentUsingPOST1Created creates a CreateAPIKeyDeploymentUsingPOST1Created with default headers values
func NewCreateAPIKeyDeploymentUsingPOST1Created() *CreateAPIKeyDeploymentUsingPOST1Created {
	return &CreateAPIKeyDeploymentUsingPOST1Created{}
}

/*
CreateAPIKeyDeploymentUsingPOST1Created describes a response with status code 201, with default header values.

Created
*/
type CreateAPIKeyDeploymentUsingPOST1Created struct {
	Payload *models.APIKeyDeploymentDto
}

// IsSuccess returns true when this create Api key deployment using p o s t1 created response has a 2xx status code
func (o *CreateAPIKeyDeploymentUsingPOST1Created) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create Api key deployment using p o s t1 created response has a 3xx status code
func (o *CreateAPIKeyDeploymentUsingPOST1Created) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create Api key deployment using p o s t1 created response has a 4xx status code
func (o *CreateAPIKeyDeploymentUsingPOST1Created) IsClientError() bool {
	return false
}

// IsServerError returns true when this create Api key deployment using p o s t1 created response has a 5xx status code
func (o *CreateAPIKeyDeploymentUsingPOST1Created) IsServerError() bool {
	return false
}

// IsCode returns true when this create Api key deployment using p o s t1 created response a status code equal to that given
func (o *CreateAPIKeyDeploymentUsingPOST1Created) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create Api key deployment using p o s t1 created response
func (o *CreateAPIKeyDeploymentUsingPOST1Created) Code() int {
	return 201
}

func (o *CreateAPIKeyDeploymentUsingPOST1Created) Error() string {
	return fmt.Sprintf("[POST /deployments/1.0/api-keys/{apiKey}/proxies][%d] createApiKeyDeploymentUsingPOST1Created  %+v", 201, o.Payload)
}

func (o *CreateAPIKeyDeploymentUsingPOST1Created) String() string {
	return fmt.Sprintf("[POST /deployments/1.0/api-keys/{apiKey}/proxies][%d] createApiKeyDeploymentUsingPOST1Created  %+v", 201, o.Payload)
}

func (o *CreateAPIKeyDeploymentUsingPOST1Created) GetPayload() *models.APIKeyDeploymentDto {
	return o.Payload
}

func (o *CreateAPIKeyDeploymentUsingPOST1Created) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIKeyDeploymentDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAPIKeyDeploymentUsingPOST1BadRequest creates a CreateAPIKeyDeploymentUsingPOST1BadRequest with default headers values
func NewCreateAPIKeyDeploymentUsingPOST1BadRequest() *CreateAPIKeyDeploymentUsingPOST1BadRequest {
	return &CreateAPIKeyDeploymentUsingPOST1BadRequest{}
}

/*
CreateAPIKeyDeploymentUsingPOST1BadRequest describes a response with status code 400, with default header values.

Bad request due to a validation error.
*/
type CreateAPIKeyDeploymentUsingPOST1BadRequest struct {
}

// IsSuccess returns true when this create Api key deployment using p o s t1 bad request response has a 2xx status code
func (o *CreateAPIKeyDeploymentUsingPOST1BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create Api key deployment using p o s t1 bad request response has a 3xx status code
func (o *CreateAPIKeyDeploymentUsingPOST1BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create Api key deployment using p o s t1 bad request response has a 4xx status code
func (o *CreateAPIKeyDeploymentUsingPOST1BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create Api key deployment using p o s t1 bad request response has a 5xx status code
func (o *CreateAPIKeyDeploymentUsingPOST1BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create Api key deployment using p o s t1 bad request response a status code equal to that given
func (o *CreateAPIKeyDeploymentUsingPOST1BadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create Api key deployment using p o s t1 bad request response
func (o *CreateAPIKeyDeploymentUsingPOST1BadRequest) Code() int {
	return 400
}

func (o *CreateAPIKeyDeploymentUsingPOST1BadRequest) Error() string {
	return fmt.Sprintf("[POST /deployments/1.0/api-keys/{apiKey}/proxies][%d] createApiKeyDeploymentUsingPOST1BadRequest ", 400)
}

func (o *CreateAPIKeyDeploymentUsingPOST1BadRequest) String() string {
	return fmt.Sprintf("[POST /deployments/1.0/api-keys/{apiKey}/proxies][%d] createApiKeyDeploymentUsingPOST1BadRequest ", 400)
}

func (o *CreateAPIKeyDeploymentUsingPOST1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateAPIKeyDeploymentUsingPOST1NotFound creates a CreateAPIKeyDeploymentUsingPOST1NotFound with default headers values
func NewCreateAPIKeyDeploymentUsingPOST1NotFound() *CreateAPIKeyDeploymentUsingPOST1NotFound {
	return &CreateAPIKeyDeploymentUsingPOST1NotFound{}
}

/*
CreateAPIKeyDeploymentUsingPOST1NotFound describes a response with status code 404, with default header values.

Entity not found
*/
type CreateAPIKeyDeploymentUsingPOST1NotFound struct {
}

// IsSuccess returns true when this create Api key deployment using p o s t1 not found response has a 2xx status code
func (o *CreateAPIKeyDeploymentUsingPOST1NotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create Api key deployment using p o s t1 not found response has a 3xx status code
func (o *CreateAPIKeyDeploymentUsingPOST1NotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create Api key deployment using p o s t1 not found response has a 4xx status code
func (o *CreateAPIKeyDeploymentUsingPOST1NotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create Api key deployment using p o s t1 not found response has a 5xx status code
func (o *CreateAPIKeyDeploymentUsingPOST1NotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create Api key deployment using p o s t1 not found response a status code equal to that given
func (o *CreateAPIKeyDeploymentUsingPOST1NotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the create Api key deployment using p o s t1 not found response
func (o *CreateAPIKeyDeploymentUsingPOST1NotFound) Code() int {
	return 404
}

func (o *CreateAPIKeyDeploymentUsingPOST1NotFound) Error() string {
	return fmt.Sprintf("[POST /deployments/1.0/api-keys/{apiKey}/proxies][%d] createApiKeyDeploymentUsingPOST1NotFound ", 404)
}

func (o *CreateAPIKeyDeploymentUsingPOST1NotFound) String() string {
	return fmt.Sprintf("[POST /deployments/1.0/api-keys/{apiKey}/proxies][%d] createApiKeyDeploymentUsingPOST1NotFound ", 404)
}

func (o *CreateAPIKeyDeploymentUsingPOST1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
