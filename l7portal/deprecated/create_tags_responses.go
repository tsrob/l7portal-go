// Code generated by go-swagger; DO NOT EDIT.

package deprecated

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/tsrob/l7portal-go/models"
)

// CreateTagsReader is a Reader for the CreateTags structure.
type CreateTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateTagsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateTagsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateTagsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateTagsCreated creates a CreateTagsCreated with default headers values
func NewCreateTagsCreated() *CreateTagsCreated {
	return &CreateTagsCreated{}
}

/*
CreateTagsCreated describes a response with status code 201, with default header values.

Created
*/
type CreateTagsCreated struct {
	Payload models.TagsResponse
}

// IsSuccess returns true when this create tags created response has a 2xx status code
func (o *CreateTagsCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create tags created response has a 3xx status code
func (o *CreateTagsCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create tags created response has a 4xx status code
func (o *CreateTagsCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create tags created response has a 5xx status code
func (o *CreateTagsCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create tags created response a status code equal to that given
func (o *CreateTagsCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create tags created response
func (o *CreateTagsCreated) Code() int {
	return 201
}

func (o *CreateTagsCreated) Error() string {
	return fmt.Sprintf("[POST /tags][%d] createTagsCreated  %+v", 201, o.Payload)
}

func (o *CreateTagsCreated) String() string {
	return fmt.Sprintf("[POST /tags][%d] createTagsCreated  %+v", 201, o.Payload)
}

func (o *CreateTagsCreated) GetPayload() models.TagsResponse {
	return o.Payload
}

func (o *CreateTagsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTagsBadRequest creates a CreateTagsBadRequest with default headers values
func NewCreateTagsBadRequest() *CreateTagsBadRequest {
	return &CreateTagsBadRequest{}
}

/*
CreateTagsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateTagsBadRequest struct {
}

// IsSuccess returns true when this create tags bad request response has a 2xx status code
func (o *CreateTagsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create tags bad request response has a 3xx status code
func (o *CreateTagsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create tags bad request response has a 4xx status code
func (o *CreateTagsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create tags bad request response has a 5xx status code
func (o *CreateTagsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create tags bad request response a status code equal to that given
func (o *CreateTagsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create tags bad request response
func (o *CreateTagsBadRequest) Code() int {
	return 400
}

func (o *CreateTagsBadRequest) Error() string {
	return fmt.Sprintf("[POST /tags][%d] createTagsBadRequest ", 400)
}

func (o *CreateTagsBadRequest) String() string {
	return fmt.Sprintf("[POST /tags][%d] createTagsBadRequest ", 400)
}

func (o *CreateTagsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateTagsInternalServerError creates a CreateTagsInternalServerError with default headers values
func NewCreateTagsInternalServerError() *CreateTagsInternalServerError {
	return &CreateTagsInternalServerError{}
}

/*
CreateTagsInternalServerError describes a response with status code 500, with default header values.

Server Failure
*/
type CreateTagsInternalServerError struct {
}

// IsSuccess returns true when this create tags internal server error response has a 2xx status code
func (o *CreateTagsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create tags internal server error response has a 3xx status code
func (o *CreateTagsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create tags internal server error response has a 4xx status code
func (o *CreateTagsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create tags internal server error response has a 5xx status code
func (o *CreateTagsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create tags internal server error response a status code equal to that given
func (o *CreateTagsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create tags internal server error response
func (o *CreateTagsInternalServerError) Code() int {
	return 500
}

func (o *CreateTagsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /tags][%d] createTagsInternalServerError ", 500)
}

func (o *CreateTagsInternalServerError) String() string {
	return fmt.Sprintf("[POST /tags][%d] createTagsInternalServerError ", 500)
}

func (o *CreateTagsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
