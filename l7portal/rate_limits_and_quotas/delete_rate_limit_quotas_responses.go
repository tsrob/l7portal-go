// Code generated by go-swagger; DO NOT EDIT.

package rate_limits_and_quotas

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteRateLimitQuotasReader is a Reader for the DeleteRateLimitQuotas structure.
type DeleteRateLimitQuotasReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRateLimitQuotasReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteRateLimitQuotasNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteRateLimitQuotasBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteRateLimitQuotasNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteRateLimitQuotasInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteRateLimitQuotasNoContent creates a DeleteRateLimitQuotasNoContent with default headers values
func NewDeleteRateLimitQuotasNoContent() *DeleteRateLimitQuotasNoContent {
	return &DeleteRateLimitQuotasNoContent{}
}

/*
DeleteRateLimitQuotasNoContent describes a response with status code 204, with default header values.

No Content
*/
type DeleteRateLimitQuotasNoContent struct {
}

// IsSuccess returns true when this delete rate limit quotas no content response has a 2xx status code
func (o *DeleteRateLimitQuotasNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete rate limit quotas no content response has a 3xx status code
func (o *DeleteRateLimitQuotasNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete rate limit quotas no content response has a 4xx status code
func (o *DeleteRateLimitQuotasNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete rate limit quotas no content response has a 5xx status code
func (o *DeleteRateLimitQuotasNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete rate limit quotas no content response a status code equal to that given
func (o *DeleteRateLimitQuotasNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete rate limit quotas no content response
func (o *DeleteRateLimitQuotasNoContent) Code() int {
	return 204
}

func (o *DeleteRateLimitQuotasNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api-management/1.0/rate-quotas/{uuid}][%d] deleteRateLimitQuotasNoContent ", 204)
}

func (o *DeleteRateLimitQuotasNoContent) String() string {
	return fmt.Sprintf("[DELETE /api-management/1.0/rate-quotas/{uuid}][%d] deleteRateLimitQuotasNoContent ", 204)
}

func (o *DeleteRateLimitQuotasNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRateLimitQuotasBadRequest creates a DeleteRateLimitQuotasBadRequest with default headers values
func NewDeleteRateLimitQuotasBadRequest() *DeleteRateLimitQuotasBadRequest {
	return &DeleteRateLimitQuotasBadRequest{}
}

/*
DeleteRateLimitQuotasBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteRateLimitQuotasBadRequest struct {
}

// IsSuccess returns true when this delete rate limit quotas bad request response has a 2xx status code
func (o *DeleteRateLimitQuotasBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete rate limit quotas bad request response has a 3xx status code
func (o *DeleteRateLimitQuotasBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete rate limit quotas bad request response has a 4xx status code
func (o *DeleteRateLimitQuotasBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete rate limit quotas bad request response has a 5xx status code
func (o *DeleteRateLimitQuotasBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete rate limit quotas bad request response a status code equal to that given
func (o *DeleteRateLimitQuotasBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete rate limit quotas bad request response
func (o *DeleteRateLimitQuotasBadRequest) Code() int {
	return 400
}

func (o *DeleteRateLimitQuotasBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api-management/1.0/rate-quotas/{uuid}][%d] deleteRateLimitQuotasBadRequest ", 400)
}

func (o *DeleteRateLimitQuotasBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api-management/1.0/rate-quotas/{uuid}][%d] deleteRateLimitQuotasBadRequest ", 400)
}

func (o *DeleteRateLimitQuotasBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRateLimitQuotasNotFound creates a DeleteRateLimitQuotasNotFound with default headers values
func NewDeleteRateLimitQuotasNotFound() *DeleteRateLimitQuotasNotFound {
	return &DeleteRateLimitQuotasNotFound{}
}

/*
DeleteRateLimitQuotasNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteRateLimitQuotasNotFound struct {
}

// IsSuccess returns true when this delete rate limit quotas not found response has a 2xx status code
func (o *DeleteRateLimitQuotasNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete rate limit quotas not found response has a 3xx status code
func (o *DeleteRateLimitQuotasNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete rate limit quotas not found response has a 4xx status code
func (o *DeleteRateLimitQuotasNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete rate limit quotas not found response has a 5xx status code
func (o *DeleteRateLimitQuotasNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete rate limit quotas not found response a status code equal to that given
func (o *DeleteRateLimitQuotasNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete rate limit quotas not found response
func (o *DeleteRateLimitQuotasNotFound) Code() int {
	return 404
}

func (o *DeleteRateLimitQuotasNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api-management/1.0/rate-quotas/{uuid}][%d] deleteRateLimitQuotasNotFound ", 404)
}

func (o *DeleteRateLimitQuotasNotFound) String() string {
	return fmt.Sprintf("[DELETE /api-management/1.0/rate-quotas/{uuid}][%d] deleteRateLimitQuotasNotFound ", 404)
}

func (o *DeleteRateLimitQuotasNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRateLimitQuotasInternalServerError creates a DeleteRateLimitQuotasInternalServerError with default headers values
func NewDeleteRateLimitQuotasInternalServerError() *DeleteRateLimitQuotasInternalServerError {
	return &DeleteRateLimitQuotasInternalServerError{}
}

/*
DeleteRateLimitQuotasInternalServerError describes a response with status code 500, with default header values.

Server Failure
*/
type DeleteRateLimitQuotasInternalServerError struct {
}

// IsSuccess returns true when this delete rate limit quotas internal server error response has a 2xx status code
func (o *DeleteRateLimitQuotasInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete rate limit quotas internal server error response has a 3xx status code
func (o *DeleteRateLimitQuotasInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete rate limit quotas internal server error response has a 4xx status code
func (o *DeleteRateLimitQuotasInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete rate limit quotas internal server error response has a 5xx status code
func (o *DeleteRateLimitQuotasInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete rate limit quotas internal server error response a status code equal to that given
func (o *DeleteRateLimitQuotasInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete rate limit quotas internal server error response
func (o *DeleteRateLimitQuotasInternalServerError) Code() int {
	return 500
}

func (o *DeleteRateLimitQuotasInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api-management/1.0/rate-quotas/{uuid}][%d] deleteRateLimitQuotasInternalServerError ", 500)
}

func (o *DeleteRateLimitQuotasInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api-management/1.0/rate-quotas/{uuid}][%d] deleteRateLimitQuotasInternalServerError ", 500)
}

func (o *DeleteRateLimitQuotasInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}