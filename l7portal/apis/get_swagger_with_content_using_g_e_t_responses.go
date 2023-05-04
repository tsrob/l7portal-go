// Code generated by go-swagger; DO NOT EDIT.

package apis

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/tsrob/l7portal-go/models"
)

// GetSwaggerWithContentUsingGETReader is a Reader for the GetSwaggerWithContentUsingGET structure.
type GetSwaggerWithContentUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSwaggerWithContentUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSwaggerWithContentUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetSwaggerWithContentUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSwaggerWithContentUsingGETInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSwaggerWithContentUsingGETOK creates a GetSwaggerWithContentUsingGETOK with default headers values
func NewGetSwaggerWithContentUsingGETOK() *GetSwaggerWithContentUsingGETOK {
	return &GetSwaggerWithContentUsingGETOK{}
}

/*
GetSwaggerWithContentUsingGETOK describes a response with status code 200, with default header values.

OK
*/
type GetSwaggerWithContentUsingGETOK struct {
	Payload *models.APIAsset
}

// IsSuccess returns true when this get swagger with content using g e t o k response has a 2xx status code
func (o *GetSwaggerWithContentUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get swagger with content using g e t o k response has a 3xx status code
func (o *GetSwaggerWithContentUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get swagger with content using g e t o k response has a 4xx status code
func (o *GetSwaggerWithContentUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get swagger with content using g e t o k response has a 5xx status code
func (o *GetSwaggerWithContentUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get swagger with content using g e t o k response a status code equal to that given
func (o *GetSwaggerWithContentUsingGETOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get swagger with content using g e t o k response
func (o *GetSwaggerWithContentUsingGETOK) Code() int {
	return 200
}

func (o *GetSwaggerWithContentUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /api-management/1.0/apis/{apiUuid}/assets/swagger][%d] getSwaggerWithContentUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetSwaggerWithContentUsingGETOK) String() string {
	return fmt.Sprintf("[GET /api-management/1.0/apis/{apiUuid}/assets/swagger][%d] getSwaggerWithContentUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetSwaggerWithContentUsingGETOK) GetPayload() *models.APIAsset {
	return o.Payload
}

func (o *GetSwaggerWithContentUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIAsset)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSwaggerWithContentUsingGETNotFound creates a GetSwaggerWithContentUsingGETNotFound with default headers values
func NewGetSwaggerWithContentUsingGETNotFound() *GetSwaggerWithContentUsingGETNotFound {
	return &GetSwaggerWithContentUsingGETNotFound{}
}

/*
GetSwaggerWithContentUsingGETNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetSwaggerWithContentUsingGETNotFound struct {
}

// IsSuccess returns true when this get swagger with content using g e t not found response has a 2xx status code
func (o *GetSwaggerWithContentUsingGETNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get swagger with content using g e t not found response has a 3xx status code
func (o *GetSwaggerWithContentUsingGETNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get swagger with content using g e t not found response has a 4xx status code
func (o *GetSwaggerWithContentUsingGETNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get swagger with content using g e t not found response has a 5xx status code
func (o *GetSwaggerWithContentUsingGETNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get swagger with content using g e t not found response a status code equal to that given
func (o *GetSwaggerWithContentUsingGETNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get swagger with content using g e t not found response
func (o *GetSwaggerWithContentUsingGETNotFound) Code() int {
	return 404
}

func (o *GetSwaggerWithContentUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /api-management/1.0/apis/{apiUuid}/assets/swagger][%d] getSwaggerWithContentUsingGETNotFound ", 404)
}

func (o *GetSwaggerWithContentUsingGETNotFound) String() string {
	return fmt.Sprintf("[GET /api-management/1.0/apis/{apiUuid}/assets/swagger][%d] getSwaggerWithContentUsingGETNotFound ", 404)
}

func (o *GetSwaggerWithContentUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSwaggerWithContentUsingGETInternalServerError creates a GetSwaggerWithContentUsingGETInternalServerError with default headers values
func NewGetSwaggerWithContentUsingGETInternalServerError() *GetSwaggerWithContentUsingGETInternalServerError {
	return &GetSwaggerWithContentUsingGETInternalServerError{}
}

/*
GetSwaggerWithContentUsingGETInternalServerError describes a response with status code 500, with default header values.

Server Failure
*/
type GetSwaggerWithContentUsingGETInternalServerError struct {
}

// IsSuccess returns true when this get swagger with content using g e t internal server error response has a 2xx status code
func (o *GetSwaggerWithContentUsingGETInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get swagger with content using g e t internal server error response has a 3xx status code
func (o *GetSwaggerWithContentUsingGETInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get swagger with content using g e t internal server error response has a 4xx status code
func (o *GetSwaggerWithContentUsingGETInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get swagger with content using g e t internal server error response has a 5xx status code
func (o *GetSwaggerWithContentUsingGETInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get swagger with content using g e t internal server error response a status code equal to that given
func (o *GetSwaggerWithContentUsingGETInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get swagger with content using g e t internal server error response
func (o *GetSwaggerWithContentUsingGETInternalServerError) Code() int {
	return 500
}

func (o *GetSwaggerWithContentUsingGETInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api-management/1.0/apis/{apiUuid}/assets/swagger][%d] getSwaggerWithContentUsingGETInternalServerError ", 500)
}

func (o *GetSwaggerWithContentUsingGETInternalServerError) String() string {
	return fmt.Sprintf("[GET /api-management/1.0/apis/{apiUuid}/assets/swagger][%d] getSwaggerWithContentUsingGETInternalServerError ", 500)
}

func (o *GetSwaggerWithContentUsingGETInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
