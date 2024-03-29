// Code generated by go-swagger; DO NOT EDIT.

package email_templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/tsrob/l7portal-go/models"
)

// GetAllEmailTemplateEntitiesReader is a Reader for the GetAllEmailTemplateEntities structure.
type GetAllEmailTemplateEntitiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllEmailTemplateEntitiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAllEmailTemplateEntitiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetAllEmailTemplateEntitiesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAllEmailTemplateEntitiesOK creates a GetAllEmailTemplateEntitiesOK with default headers values
func NewGetAllEmailTemplateEntitiesOK() *GetAllEmailTemplateEntitiesOK {
	return &GetAllEmailTemplateEntitiesOK{}
}

/*
GetAllEmailTemplateEntitiesOK describes a response with status code 200, with default header values.

GetAllEmailTemplateEntitiesOK get all email template entities o k
*/
type GetAllEmailTemplateEntitiesOK struct {
	Payload models.EmailTemplateList
}

// IsSuccess returns true when this get all email template entities o k response has a 2xx status code
func (o *GetAllEmailTemplateEntitiesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get all email template entities o k response has a 3xx status code
func (o *GetAllEmailTemplateEntitiesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all email template entities o k response has a 4xx status code
func (o *GetAllEmailTemplateEntitiesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get all email template entities o k response has a 5xx status code
func (o *GetAllEmailTemplateEntitiesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get all email template entities o k response a status code equal to that given
func (o *GetAllEmailTemplateEntitiesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get all email template entities o k response
func (o *GetAllEmailTemplateEntitiesOK) Code() int {
	return 200
}

func (o *GetAllEmailTemplateEntitiesOK) Error() string {
	return fmt.Sprintf("[GET /tenant-admin/1.0/email-templates][%d] getAllEmailTemplateEntitiesOK  %+v", 200, o.Payload)
}

func (o *GetAllEmailTemplateEntitiesOK) String() string {
	return fmt.Sprintf("[GET /tenant-admin/1.0/email-templates][%d] getAllEmailTemplateEntitiesOK  %+v", 200, o.Payload)
}

func (o *GetAllEmailTemplateEntitiesOK) GetPayload() models.EmailTemplateList {
	return o.Payload
}

func (o *GetAllEmailTemplateEntitiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllEmailTemplateEntitiesInternalServerError creates a GetAllEmailTemplateEntitiesInternalServerError with default headers values
func NewGetAllEmailTemplateEntitiesInternalServerError() *GetAllEmailTemplateEntitiesInternalServerError {
	return &GetAllEmailTemplateEntitiesInternalServerError{}
}

/*
GetAllEmailTemplateEntitiesInternalServerError describes a response with status code 500, with default header values.

Server Failure
*/
type GetAllEmailTemplateEntitiesInternalServerError struct {
}

// IsSuccess returns true when this get all email template entities internal server error response has a 2xx status code
func (o *GetAllEmailTemplateEntitiesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get all email template entities internal server error response has a 3xx status code
func (o *GetAllEmailTemplateEntitiesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all email template entities internal server error response has a 4xx status code
func (o *GetAllEmailTemplateEntitiesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get all email template entities internal server error response has a 5xx status code
func (o *GetAllEmailTemplateEntitiesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get all email template entities internal server error response a status code equal to that given
func (o *GetAllEmailTemplateEntitiesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get all email template entities internal server error response
func (o *GetAllEmailTemplateEntitiesInternalServerError) Code() int {
	return 500
}

func (o *GetAllEmailTemplateEntitiesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /tenant-admin/1.0/email-templates][%d] getAllEmailTemplateEntitiesInternalServerError ", 500)
}

func (o *GetAllEmailTemplateEntitiesInternalServerError) String() string {
	return fmt.Sprintf("[GET /tenant-admin/1.0/email-templates][%d] getAllEmailTemplateEntitiesInternalServerError ", 500)
}

func (o *GetAllEmailTemplateEntitiesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
