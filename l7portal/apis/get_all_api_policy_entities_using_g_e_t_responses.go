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

// GetAllAPIPolicyEntitiesUsingGETReader is a Reader for the GetAllAPIPolicyEntitiesUsingGET structure.
type GetAllAPIPolicyEntitiesUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllAPIPolicyEntitiesUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAllAPIPolicyEntitiesUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetAllAPIPolicyEntitiesUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAllAPIPolicyEntitiesUsingGETInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAllAPIPolicyEntitiesUsingGETOK creates a GetAllAPIPolicyEntitiesUsingGETOK with default headers values
func NewGetAllAPIPolicyEntitiesUsingGETOK() *GetAllAPIPolicyEntitiesUsingGETOK {
	return &GetAllAPIPolicyEntitiesUsingGETOK{}
}

/*
GetAllAPIPolicyEntitiesUsingGETOK describes a response with status code 200, with default header values.

OK
*/
type GetAllAPIPolicyEntitiesUsingGETOK struct {
	Payload []*models.APIPolicyEntityDto
}

// IsSuccess returns true when this get all Api policy entities using g e t o k response has a 2xx status code
func (o *GetAllAPIPolicyEntitiesUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get all Api policy entities using g e t o k response has a 3xx status code
func (o *GetAllAPIPolicyEntitiesUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all Api policy entities using g e t o k response has a 4xx status code
func (o *GetAllAPIPolicyEntitiesUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get all Api policy entities using g e t o k response has a 5xx status code
func (o *GetAllAPIPolicyEntitiesUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get all Api policy entities using g e t o k response a status code equal to that given
func (o *GetAllAPIPolicyEntitiesUsingGETOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get all Api policy entities using g e t o k response
func (o *GetAllAPIPolicyEntitiesUsingGETOK) Code() int {
	return 200
}

func (o *GetAllAPIPolicyEntitiesUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /api-management/1.0/apis/{apiUuid}/policy-entities][%d] getAllApiPolicyEntitiesUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetAllAPIPolicyEntitiesUsingGETOK) String() string {
	return fmt.Sprintf("[GET /api-management/1.0/apis/{apiUuid}/policy-entities][%d] getAllApiPolicyEntitiesUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetAllAPIPolicyEntitiesUsingGETOK) GetPayload() []*models.APIPolicyEntityDto {
	return o.Payload
}

func (o *GetAllAPIPolicyEntitiesUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllAPIPolicyEntitiesUsingGETNotFound creates a GetAllAPIPolicyEntitiesUsingGETNotFound with default headers values
func NewGetAllAPIPolicyEntitiesUsingGETNotFound() *GetAllAPIPolicyEntitiesUsingGETNotFound {
	return &GetAllAPIPolicyEntitiesUsingGETNotFound{}
}

/*
GetAllAPIPolicyEntitiesUsingGETNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetAllAPIPolicyEntitiesUsingGETNotFound struct {
}

// IsSuccess returns true when this get all Api policy entities using g e t not found response has a 2xx status code
func (o *GetAllAPIPolicyEntitiesUsingGETNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get all Api policy entities using g e t not found response has a 3xx status code
func (o *GetAllAPIPolicyEntitiesUsingGETNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all Api policy entities using g e t not found response has a 4xx status code
func (o *GetAllAPIPolicyEntitiesUsingGETNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get all Api policy entities using g e t not found response has a 5xx status code
func (o *GetAllAPIPolicyEntitiesUsingGETNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get all Api policy entities using g e t not found response a status code equal to that given
func (o *GetAllAPIPolicyEntitiesUsingGETNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get all Api policy entities using g e t not found response
func (o *GetAllAPIPolicyEntitiesUsingGETNotFound) Code() int {
	return 404
}

func (o *GetAllAPIPolicyEntitiesUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /api-management/1.0/apis/{apiUuid}/policy-entities][%d] getAllApiPolicyEntitiesUsingGETNotFound ", 404)
}

func (o *GetAllAPIPolicyEntitiesUsingGETNotFound) String() string {
	return fmt.Sprintf("[GET /api-management/1.0/apis/{apiUuid}/policy-entities][%d] getAllApiPolicyEntitiesUsingGETNotFound ", 404)
}

func (o *GetAllAPIPolicyEntitiesUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllAPIPolicyEntitiesUsingGETInternalServerError creates a GetAllAPIPolicyEntitiesUsingGETInternalServerError with default headers values
func NewGetAllAPIPolicyEntitiesUsingGETInternalServerError() *GetAllAPIPolicyEntitiesUsingGETInternalServerError {
	return &GetAllAPIPolicyEntitiesUsingGETInternalServerError{}
}

/*
GetAllAPIPolicyEntitiesUsingGETInternalServerError describes a response with status code 500, with default header values.

Server Failure
*/
type GetAllAPIPolicyEntitiesUsingGETInternalServerError struct {
}

// IsSuccess returns true when this get all Api policy entities using g e t internal server error response has a 2xx status code
func (o *GetAllAPIPolicyEntitiesUsingGETInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get all Api policy entities using g e t internal server error response has a 3xx status code
func (o *GetAllAPIPolicyEntitiesUsingGETInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all Api policy entities using g e t internal server error response has a 4xx status code
func (o *GetAllAPIPolicyEntitiesUsingGETInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get all Api policy entities using g e t internal server error response has a 5xx status code
func (o *GetAllAPIPolicyEntitiesUsingGETInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get all Api policy entities using g e t internal server error response a status code equal to that given
func (o *GetAllAPIPolicyEntitiesUsingGETInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get all Api policy entities using g e t internal server error response
func (o *GetAllAPIPolicyEntitiesUsingGETInternalServerError) Code() int {
	return 500
}

func (o *GetAllAPIPolicyEntitiesUsingGETInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api-management/1.0/apis/{apiUuid}/policy-entities][%d] getAllApiPolicyEntitiesUsingGETInternalServerError ", 500)
}

func (o *GetAllAPIPolicyEntitiesUsingGETInternalServerError) String() string {
	return fmt.Sprintf("[GET /api-management/1.0/apis/{apiUuid}/policy-entities][%d] getAllApiPolicyEntitiesUsingGETInternalServerError ", 500)
}

func (o *GetAllAPIPolicyEntitiesUsingGETInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
