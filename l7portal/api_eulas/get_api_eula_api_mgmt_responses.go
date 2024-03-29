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

// GetAPIEulaAPIMgmtReader is a Reader for the GetAPIEulaAPIMgmt structure.
type GetAPIEulaAPIMgmtReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIEulaAPIMgmtReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPIEulaAPIMgmtOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetAPIEulaAPIMgmtNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAPIEulaAPIMgmtOK creates a GetAPIEulaAPIMgmtOK with default headers values
func NewGetAPIEulaAPIMgmtOK() *GetAPIEulaAPIMgmtOK {
	return &GetAPIEulaAPIMgmtOK{}
}

/*
GetAPIEulaAPIMgmtOK describes a response with status code 200, with default header values.

GetAPIEulaAPIMgmtOK get Api eula Api mgmt o k
*/
type GetAPIEulaAPIMgmtOK struct {
	Payload *models.Eula
}

// IsSuccess returns true when this get Api eula Api mgmt o k response has a 2xx status code
func (o *GetAPIEulaAPIMgmtOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get Api eula Api mgmt o k response has a 3xx status code
func (o *GetAPIEulaAPIMgmtOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Api eula Api mgmt o k response has a 4xx status code
func (o *GetAPIEulaAPIMgmtOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Api eula Api mgmt o k response has a 5xx status code
func (o *GetAPIEulaAPIMgmtOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get Api eula Api mgmt o k response a status code equal to that given
func (o *GetAPIEulaAPIMgmtOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get Api eula Api mgmt o k response
func (o *GetAPIEulaAPIMgmtOK) Code() int {
	return 200
}

func (o *GetAPIEulaAPIMgmtOK) Error() string {
	return fmt.Sprintf("[GET /api-management/1.0/eulas/{uuid}][%d] getApiEulaApiMgmtOK  %+v", 200, o.Payload)
}

func (o *GetAPIEulaAPIMgmtOK) String() string {
	return fmt.Sprintf("[GET /api-management/1.0/eulas/{uuid}][%d] getApiEulaApiMgmtOK  %+v", 200, o.Payload)
}

func (o *GetAPIEulaAPIMgmtOK) GetPayload() *models.Eula {
	return o.Payload
}

func (o *GetAPIEulaAPIMgmtOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Eula)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIEulaAPIMgmtNotFound creates a GetAPIEulaAPIMgmtNotFound with default headers values
func NewGetAPIEulaAPIMgmtNotFound() *GetAPIEulaAPIMgmtNotFound {
	return &GetAPIEulaAPIMgmtNotFound{}
}

/*
GetAPIEulaAPIMgmtNotFound describes a response with status code 404, with default header values.

Entity not found
*/
type GetAPIEulaAPIMgmtNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get Api eula Api mgmt not found response has a 2xx status code
func (o *GetAPIEulaAPIMgmtNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get Api eula Api mgmt not found response has a 3xx status code
func (o *GetAPIEulaAPIMgmtNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Api eula Api mgmt not found response has a 4xx status code
func (o *GetAPIEulaAPIMgmtNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get Api eula Api mgmt not found response has a 5xx status code
func (o *GetAPIEulaAPIMgmtNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get Api eula Api mgmt not found response a status code equal to that given
func (o *GetAPIEulaAPIMgmtNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get Api eula Api mgmt not found response
func (o *GetAPIEulaAPIMgmtNotFound) Code() int {
	return 404
}

func (o *GetAPIEulaAPIMgmtNotFound) Error() string {
	return fmt.Sprintf("[GET /api-management/1.0/eulas/{uuid}][%d] getApiEulaApiMgmtNotFound  %+v", 404, o.Payload)
}

func (o *GetAPIEulaAPIMgmtNotFound) String() string {
	return fmt.Sprintf("[GET /api-management/1.0/eulas/{uuid}][%d] getApiEulaApiMgmtNotFound  %+v", 404, o.Payload)
}

func (o *GetAPIEulaAPIMgmtNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAPIEulaAPIMgmtNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
