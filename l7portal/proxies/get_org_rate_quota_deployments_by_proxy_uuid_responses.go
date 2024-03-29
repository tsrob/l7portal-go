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

// GetOrgRateQuotaDeploymentsByProxyUUIDReader is a Reader for the GetOrgRateQuotaDeploymentsByProxyUUID structure.
type GetOrgRateQuotaDeploymentsByProxyUUIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrgRateQuotaDeploymentsByProxyUUIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrgRateQuotaDeploymentsByProxyUUIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOrgRateQuotaDeploymentsByProxyUUIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOrgRateQuotaDeploymentsByProxyUUIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOrgRateQuotaDeploymentsByProxyUUIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOrgRateQuotaDeploymentsByProxyUUIDOK creates a GetOrgRateQuotaDeploymentsByProxyUUIDOK with default headers values
func NewGetOrgRateQuotaDeploymentsByProxyUUIDOK() *GetOrgRateQuotaDeploymentsByProxyUUIDOK {
	return &GetOrgRateQuotaDeploymentsByProxyUUIDOK{}
}

/*
GetOrgRateQuotaDeploymentsByProxyUUIDOK describes a response with status code 200, with default header values.

Success
*/
type GetOrgRateQuotaDeploymentsByProxyUUIDOK struct {
	Payload *models.EntityDeploymentDto
}

// IsSuccess returns true when this get org rate quota deployments by proxy Uuid o k response has a 2xx status code
func (o *GetOrgRateQuotaDeploymentsByProxyUUIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get org rate quota deployments by proxy Uuid o k response has a 3xx status code
func (o *GetOrgRateQuotaDeploymentsByProxyUUIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get org rate quota deployments by proxy Uuid o k response has a 4xx status code
func (o *GetOrgRateQuotaDeploymentsByProxyUUIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get org rate quota deployments by proxy Uuid o k response has a 5xx status code
func (o *GetOrgRateQuotaDeploymentsByProxyUUIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get org rate quota deployments by proxy Uuid o k response a status code equal to that given
func (o *GetOrgRateQuotaDeploymentsByProxyUUIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get org rate quota deployments by proxy Uuid o k response
func (o *GetOrgRateQuotaDeploymentsByProxyUUIDOK) Code() int {
	return 200
}

func (o *GetOrgRateQuotaDeploymentsByProxyUUIDOK) Error() string {
	return fmt.Sprintf("[GET /deployments/1.0/proxies/{uuid}/org-rate-quotas][%d] getOrgRateQuotaDeploymentsByProxyUuidOK  %+v", 200, o.Payload)
}

func (o *GetOrgRateQuotaDeploymentsByProxyUUIDOK) String() string {
	return fmt.Sprintf("[GET /deployments/1.0/proxies/{uuid}/org-rate-quotas][%d] getOrgRateQuotaDeploymentsByProxyUuidOK  %+v", 200, o.Payload)
}

func (o *GetOrgRateQuotaDeploymentsByProxyUUIDOK) GetPayload() *models.EntityDeploymentDto {
	return o.Payload
}

func (o *GetOrgRateQuotaDeploymentsByProxyUUIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EntityDeploymentDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgRateQuotaDeploymentsByProxyUUIDBadRequest creates a GetOrgRateQuotaDeploymentsByProxyUUIDBadRequest with default headers values
func NewGetOrgRateQuotaDeploymentsByProxyUUIDBadRequest() *GetOrgRateQuotaDeploymentsByProxyUUIDBadRequest {
	return &GetOrgRateQuotaDeploymentsByProxyUUIDBadRequest{}
}

/*
GetOrgRateQuotaDeploymentsByProxyUUIDBadRequest describes a response with status code 400, with default header values.

Bad Request due to a validation error.
*/
type GetOrgRateQuotaDeploymentsByProxyUUIDBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get org rate quota deployments by proxy Uuid bad request response has a 2xx status code
func (o *GetOrgRateQuotaDeploymentsByProxyUUIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get org rate quota deployments by proxy Uuid bad request response has a 3xx status code
func (o *GetOrgRateQuotaDeploymentsByProxyUUIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get org rate quota deployments by proxy Uuid bad request response has a 4xx status code
func (o *GetOrgRateQuotaDeploymentsByProxyUUIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get org rate quota deployments by proxy Uuid bad request response has a 5xx status code
func (o *GetOrgRateQuotaDeploymentsByProxyUUIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get org rate quota deployments by proxy Uuid bad request response a status code equal to that given
func (o *GetOrgRateQuotaDeploymentsByProxyUUIDBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get org rate quota deployments by proxy Uuid bad request response
func (o *GetOrgRateQuotaDeploymentsByProxyUUIDBadRequest) Code() int {
	return 400
}

func (o *GetOrgRateQuotaDeploymentsByProxyUUIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /deployments/1.0/proxies/{uuid}/org-rate-quotas][%d] getOrgRateQuotaDeploymentsByProxyUuidBadRequest  %+v", 400, o.Payload)
}

func (o *GetOrgRateQuotaDeploymentsByProxyUUIDBadRequest) String() string {
	return fmt.Sprintf("[GET /deployments/1.0/proxies/{uuid}/org-rate-quotas][%d] getOrgRateQuotaDeploymentsByProxyUuidBadRequest  %+v", 400, o.Payload)
}

func (o *GetOrgRateQuotaDeploymentsByProxyUUIDBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetOrgRateQuotaDeploymentsByProxyUUIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgRateQuotaDeploymentsByProxyUUIDNotFound creates a GetOrgRateQuotaDeploymentsByProxyUUIDNotFound with default headers values
func NewGetOrgRateQuotaDeploymentsByProxyUUIDNotFound() *GetOrgRateQuotaDeploymentsByProxyUUIDNotFound {
	return &GetOrgRateQuotaDeploymentsByProxyUUIDNotFound{}
}

/*
GetOrgRateQuotaDeploymentsByProxyUUIDNotFound describes a response with status code 404, with default header values.

Entity not found
*/
type GetOrgRateQuotaDeploymentsByProxyUUIDNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get org rate quota deployments by proxy Uuid not found response has a 2xx status code
func (o *GetOrgRateQuotaDeploymentsByProxyUUIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get org rate quota deployments by proxy Uuid not found response has a 3xx status code
func (o *GetOrgRateQuotaDeploymentsByProxyUUIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get org rate quota deployments by proxy Uuid not found response has a 4xx status code
func (o *GetOrgRateQuotaDeploymentsByProxyUUIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get org rate quota deployments by proxy Uuid not found response has a 5xx status code
func (o *GetOrgRateQuotaDeploymentsByProxyUUIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get org rate quota deployments by proxy Uuid not found response a status code equal to that given
func (o *GetOrgRateQuotaDeploymentsByProxyUUIDNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get org rate quota deployments by proxy Uuid not found response
func (o *GetOrgRateQuotaDeploymentsByProxyUUIDNotFound) Code() int {
	return 404
}

func (o *GetOrgRateQuotaDeploymentsByProxyUUIDNotFound) Error() string {
	return fmt.Sprintf("[GET /deployments/1.0/proxies/{uuid}/org-rate-quotas][%d] getOrgRateQuotaDeploymentsByProxyUuidNotFound  %+v", 404, o.Payload)
}

func (o *GetOrgRateQuotaDeploymentsByProxyUUIDNotFound) String() string {
	return fmt.Sprintf("[GET /deployments/1.0/proxies/{uuid}/org-rate-quotas][%d] getOrgRateQuotaDeploymentsByProxyUuidNotFound  %+v", 404, o.Payload)
}

func (o *GetOrgRateQuotaDeploymentsByProxyUUIDNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetOrgRateQuotaDeploymentsByProxyUUIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgRateQuotaDeploymentsByProxyUUIDInternalServerError creates a GetOrgRateQuotaDeploymentsByProxyUUIDInternalServerError with default headers values
func NewGetOrgRateQuotaDeploymentsByProxyUUIDInternalServerError() *GetOrgRateQuotaDeploymentsByProxyUUIDInternalServerError {
	return &GetOrgRateQuotaDeploymentsByProxyUUIDInternalServerError{}
}

/*
GetOrgRateQuotaDeploymentsByProxyUUIDInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. Certain fields are missing from the request.
*/
type GetOrgRateQuotaDeploymentsByProxyUUIDInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get org rate quota deployments by proxy Uuid internal server error response has a 2xx status code
func (o *GetOrgRateQuotaDeploymentsByProxyUUIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get org rate quota deployments by proxy Uuid internal server error response has a 3xx status code
func (o *GetOrgRateQuotaDeploymentsByProxyUUIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get org rate quota deployments by proxy Uuid internal server error response has a 4xx status code
func (o *GetOrgRateQuotaDeploymentsByProxyUUIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get org rate quota deployments by proxy Uuid internal server error response has a 5xx status code
func (o *GetOrgRateQuotaDeploymentsByProxyUUIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get org rate quota deployments by proxy Uuid internal server error response a status code equal to that given
func (o *GetOrgRateQuotaDeploymentsByProxyUUIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get org rate quota deployments by proxy Uuid internal server error response
func (o *GetOrgRateQuotaDeploymentsByProxyUUIDInternalServerError) Code() int {
	return 500
}

func (o *GetOrgRateQuotaDeploymentsByProxyUUIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /deployments/1.0/proxies/{uuid}/org-rate-quotas][%d] getOrgRateQuotaDeploymentsByProxyUuidInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOrgRateQuotaDeploymentsByProxyUUIDInternalServerError) String() string {
	return fmt.Sprintf("[GET /deployments/1.0/proxies/{uuid}/org-rate-quotas][%d] getOrgRateQuotaDeploymentsByProxyUuidInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOrgRateQuotaDeploymentsByProxyUUIDInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetOrgRateQuotaDeploymentsByProxyUUIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
