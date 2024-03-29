// Code generated by go-swagger; DO NOT EDIT.

package api_eulas

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/tsrob/l7portal-go/models"
)

// GetAPIEulasAPIMgmtReader is a Reader for the GetAPIEulasAPIMgmt structure.
type GetAPIEulasAPIMgmtReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIEulasAPIMgmtReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPIEulasAPIMgmtOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAPIEulasAPIMgmtBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAPIEulasAPIMgmtInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAPIEulasAPIMgmtOK creates a GetAPIEulasAPIMgmtOK with default headers values
func NewGetAPIEulasAPIMgmtOK() *GetAPIEulasAPIMgmtOK {
	return &GetAPIEulasAPIMgmtOK{}
}

/*
GetAPIEulasAPIMgmtOK describes a response with status code 200, with default header values.

GetAPIEulasAPIMgmtOK get Api eulas Api mgmt o k
*/
type GetAPIEulasAPIMgmtOK struct {
	Payload *models.Eulas
}

// IsSuccess returns true when this get Api eulas Api mgmt o k response has a 2xx status code
func (o *GetAPIEulasAPIMgmtOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get Api eulas Api mgmt o k response has a 3xx status code
func (o *GetAPIEulasAPIMgmtOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Api eulas Api mgmt o k response has a 4xx status code
func (o *GetAPIEulasAPIMgmtOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Api eulas Api mgmt o k response has a 5xx status code
func (o *GetAPIEulasAPIMgmtOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get Api eulas Api mgmt o k response a status code equal to that given
func (o *GetAPIEulasAPIMgmtOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get Api eulas Api mgmt o k response
func (o *GetAPIEulasAPIMgmtOK) Code() int {
	return 200
}

func (o *GetAPIEulasAPIMgmtOK) Error() string {
	return fmt.Sprintf("[GET /api-management/1.0/eulas][%d] getApiEulasApiMgmtOK  %+v", 200, o.Payload)
}

func (o *GetAPIEulasAPIMgmtOK) String() string {
	return fmt.Sprintf("[GET /api-management/1.0/eulas][%d] getApiEulasApiMgmtOK  %+v", 200, o.Payload)
}

func (o *GetAPIEulasAPIMgmtOK) GetPayload() *models.Eulas {
	return o.Payload
}

func (o *GetAPIEulasAPIMgmtOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Eulas)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIEulasAPIMgmtBadRequest creates a GetAPIEulasAPIMgmtBadRequest with default headers values
func NewGetAPIEulasAPIMgmtBadRequest() *GetAPIEulasAPIMgmtBadRequest {
	return &GetAPIEulasAPIMgmtBadRequest{}
}

/*
GetAPIEulasAPIMgmtBadRequest describes a response with status code 400, with default header values.

Bad Request due to Invalid Parameters.
*/
type GetAPIEulasAPIMgmtBadRequest struct {
	Payload *models.LongerError
}

// IsSuccess returns true when this get Api eulas Api mgmt bad request response has a 2xx status code
func (o *GetAPIEulasAPIMgmtBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get Api eulas Api mgmt bad request response has a 3xx status code
func (o *GetAPIEulasAPIMgmtBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Api eulas Api mgmt bad request response has a 4xx status code
func (o *GetAPIEulasAPIMgmtBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get Api eulas Api mgmt bad request response has a 5xx status code
func (o *GetAPIEulasAPIMgmtBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get Api eulas Api mgmt bad request response a status code equal to that given
func (o *GetAPIEulasAPIMgmtBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get Api eulas Api mgmt bad request response
func (o *GetAPIEulasAPIMgmtBadRequest) Code() int {
	return 400
}

func (o *GetAPIEulasAPIMgmtBadRequest) Error() string {
	return fmt.Sprintf("[GET /api-management/1.0/eulas][%d] getApiEulasApiMgmtBadRequest  %+v", 400, o.Payload)
}

func (o *GetAPIEulasAPIMgmtBadRequest) String() string {
	return fmt.Sprintf("[GET /api-management/1.0/eulas][%d] getApiEulasApiMgmtBadRequest  %+v", 400, o.Payload)
}

func (o *GetAPIEulasAPIMgmtBadRequest) GetPayload() *models.LongerError {
	return o.Payload
}

func (o *GetAPIEulasAPIMgmtBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LongerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIEulasAPIMgmtInternalServerError creates a GetAPIEulasAPIMgmtInternalServerError with default headers values
func NewGetAPIEulasAPIMgmtInternalServerError() *GetAPIEulasAPIMgmtInternalServerError {
	return &GetAPIEulasAPIMgmtInternalServerError{}
}

/*
GetAPIEulasAPIMgmtInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. Certain fields are missing from the request.
*/
type GetAPIEulasAPIMgmtInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get Api eulas Api mgmt internal server error response has a 2xx status code
func (o *GetAPIEulasAPIMgmtInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get Api eulas Api mgmt internal server error response has a 3xx status code
func (o *GetAPIEulasAPIMgmtInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Api eulas Api mgmt internal server error response has a 4xx status code
func (o *GetAPIEulasAPIMgmtInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Api eulas Api mgmt internal server error response has a 5xx status code
func (o *GetAPIEulasAPIMgmtInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get Api eulas Api mgmt internal server error response a status code equal to that given
func (o *GetAPIEulasAPIMgmtInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get Api eulas Api mgmt internal server error response
func (o *GetAPIEulasAPIMgmtInternalServerError) Code() int {
	return 500
}

func (o *GetAPIEulasAPIMgmtInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api-management/1.0/eulas][%d] getApiEulasApiMgmtInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAPIEulasAPIMgmtInternalServerError) String() string {
	return fmt.Sprintf("[GET /api-management/1.0/eulas][%d] getApiEulasApiMgmtInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAPIEulasAPIMgmtInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAPIEulasAPIMgmtInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
