// Code generated by go-swagger; DO NOT EDIT.

package rate_limits_and_quotas

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/tsrob/l7portal-go/models"
)

// GetRateLimitQuotasReader is a Reader for the GetRateLimitQuotas structure.
type GetRateLimitQuotasReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRateLimitQuotasReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRateLimitQuotasOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRateLimitQuotasBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRateLimitQuotasInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRateLimitQuotasOK creates a GetRateLimitQuotasOK with default headers values
func NewGetRateLimitQuotasOK() *GetRateLimitQuotasOK {
	return &GetRateLimitQuotasOK{}
}

/*
GetRateLimitQuotasOK describes a response with status code 200, with default header values.

GetRateLimitQuotasOK get rate limit quotas o k
*/
type GetRateLimitQuotasOK struct {
	Payload *models.RateLimitAndQuotasGet
}

// IsSuccess returns true when this get rate limit quotas o k response has a 2xx status code
func (o *GetRateLimitQuotasOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get rate limit quotas o k response has a 3xx status code
func (o *GetRateLimitQuotasOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get rate limit quotas o k response has a 4xx status code
func (o *GetRateLimitQuotasOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get rate limit quotas o k response has a 5xx status code
func (o *GetRateLimitQuotasOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get rate limit quotas o k response a status code equal to that given
func (o *GetRateLimitQuotasOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get rate limit quotas o k response
func (o *GetRateLimitQuotasOK) Code() int {
	return 200
}

func (o *GetRateLimitQuotasOK) Error() string {
	return fmt.Sprintf("[GET /api-management/1.0/rate-quotas][%d] getRateLimitQuotasOK  %+v", 200, o.Payload)
}

func (o *GetRateLimitQuotasOK) String() string {
	return fmt.Sprintf("[GET /api-management/1.0/rate-quotas][%d] getRateLimitQuotasOK  %+v", 200, o.Payload)
}

func (o *GetRateLimitQuotasOK) GetPayload() *models.RateLimitAndQuotasGet {
	return o.Payload
}

func (o *GetRateLimitQuotasOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RateLimitAndQuotasGet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRateLimitQuotasBadRequest creates a GetRateLimitQuotasBadRequest with default headers values
func NewGetRateLimitQuotasBadRequest() *GetRateLimitQuotasBadRequest {
	return &GetRateLimitQuotasBadRequest{}
}

/*
GetRateLimitQuotasBadRequest describes a response with status code 400, with default header values.

Bad Request due to Invalid Parameters.
*/
type GetRateLimitQuotasBadRequest struct {
	Payload *models.LongerError
}

// IsSuccess returns true when this get rate limit quotas bad request response has a 2xx status code
func (o *GetRateLimitQuotasBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get rate limit quotas bad request response has a 3xx status code
func (o *GetRateLimitQuotasBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get rate limit quotas bad request response has a 4xx status code
func (o *GetRateLimitQuotasBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get rate limit quotas bad request response has a 5xx status code
func (o *GetRateLimitQuotasBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get rate limit quotas bad request response a status code equal to that given
func (o *GetRateLimitQuotasBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get rate limit quotas bad request response
func (o *GetRateLimitQuotasBadRequest) Code() int {
	return 400
}

func (o *GetRateLimitQuotasBadRequest) Error() string {
	return fmt.Sprintf("[GET /api-management/1.0/rate-quotas][%d] getRateLimitQuotasBadRequest  %+v", 400, o.Payload)
}

func (o *GetRateLimitQuotasBadRequest) String() string {
	return fmt.Sprintf("[GET /api-management/1.0/rate-quotas][%d] getRateLimitQuotasBadRequest  %+v", 400, o.Payload)
}

func (o *GetRateLimitQuotasBadRequest) GetPayload() *models.LongerError {
	return o.Payload
}

func (o *GetRateLimitQuotasBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LongerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRateLimitQuotasInternalServerError creates a GetRateLimitQuotasInternalServerError with default headers values
func NewGetRateLimitQuotasInternalServerError() *GetRateLimitQuotasInternalServerError {
	return &GetRateLimitQuotasInternalServerError{}
}

/*
GetRateLimitQuotasInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. Certain fields are missing from the request.
*/
type GetRateLimitQuotasInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get rate limit quotas internal server error response has a 2xx status code
func (o *GetRateLimitQuotasInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get rate limit quotas internal server error response has a 3xx status code
func (o *GetRateLimitQuotasInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get rate limit quotas internal server error response has a 4xx status code
func (o *GetRateLimitQuotasInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get rate limit quotas internal server error response has a 5xx status code
func (o *GetRateLimitQuotasInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get rate limit quotas internal server error response a status code equal to that given
func (o *GetRateLimitQuotasInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get rate limit quotas internal server error response
func (o *GetRateLimitQuotasInternalServerError) Code() int {
	return 500
}

func (o *GetRateLimitQuotasInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api-management/1.0/rate-quotas][%d] getRateLimitQuotasInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRateLimitQuotasInternalServerError) String() string {
	return fmt.Sprintf("[GET /api-management/1.0/rate-quotas][%d] getRateLimitQuotasInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRateLimitQuotasInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetRateLimitQuotasInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
