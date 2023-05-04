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

// CreateAPIEulaAPIMgmtReader is a Reader for the CreateAPIEulaAPIMgmt structure.
type CreateAPIEulaAPIMgmtReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAPIEulaAPIMgmtReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateAPIEulaAPIMgmtCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateAPIEulaAPIMgmtBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateAPIEulaAPIMgmtInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateAPIEulaAPIMgmtCreated creates a CreateAPIEulaAPIMgmtCreated with default headers values
func NewCreateAPIEulaAPIMgmtCreated() *CreateAPIEulaAPIMgmtCreated {
	return &CreateAPIEulaAPIMgmtCreated{}
}

/*
CreateAPIEulaAPIMgmtCreated describes a response with status code 201, with default header values.

CreateAPIEulaAPIMgmtCreated create Api eula Api mgmt created
*/
type CreateAPIEulaAPIMgmtCreated struct {
	Payload *models.Eula
}

// IsSuccess returns true when this create Api eula Api mgmt created response has a 2xx status code
func (o *CreateAPIEulaAPIMgmtCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create Api eula Api mgmt created response has a 3xx status code
func (o *CreateAPIEulaAPIMgmtCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create Api eula Api mgmt created response has a 4xx status code
func (o *CreateAPIEulaAPIMgmtCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create Api eula Api mgmt created response has a 5xx status code
func (o *CreateAPIEulaAPIMgmtCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create Api eula Api mgmt created response a status code equal to that given
func (o *CreateAPIEulaAPIMgmtCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create Api eula Api mgmt created response
func (o *CreateAPIEulaAPIMgmtCreated) Code() int {
	return 201
}

func (o *CreateAPIEulaAPIMgmtCreated) Error() string {
	return fmt.Sprintf("[POST /api-management/1.0/eulas][%d] createApiEulaApiMgmtCreated  %+v", 201, o.Payload)
}

func (o *CreateAPIEulaAPIMgmtCreated) String() string {
	return fmt.Sprintf("[POST /api-management/1.0/eulas][%d] createApiEulaApiMgmtCreated  %+v", 201, o.Payload)
}

func (o *CreateAPIEulaAPIMgmtCreated) GetPayload() *models.Eula {
	return o.Payload
}

func (o *CreateAPIEulaAPIMgmtCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Eula)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAPIEulaAPIMgmtBadRequest creates a CreateAPIEulaAPIMgmtBadRequest with default headers values
func NewCreateAPIEulaAPIMgmtBadRequest() *CreateAPIEulaAPIMgmtBadRequest {
	return &CreateAPIEulaAPIMgmtBadRequest{}
}

/*
CreateAPIEulaAPIMgmtBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateAPIEulaAPIMgmtBadRequest struct {
}

// IsSuccess returns true when this create Api eula Api mgmt bad request response has a 2xx status code
func (o *CreateAPIEulaAPIMgmtBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create Api eula Api mgmt bad request response has a 3xx status code
func (o *CreateAPIEulaAPIMgmtBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create Api eula Api mgmt bad request response has a 4xx status code
func (o *CreateAPIEulaAPIMgmtBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create Api eula Api mgmt bad request response has a 5xx status code
func (o *CreateAPIEulaAPIMgmtBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create Api eula Api mgmt bad request response a status code equal to that given
func (o *CreateAPIEulaAPIMgmtBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create Api eula Api mgmt bad request response
func (o *CreateAPIEulaAPIMgmtBadRequest) Code() int {
	return 400
}

func (o *CreateAPIEulaAPIMgmtBadRequest) Error() string {
	return fmt.Sprintf("[POST /api-management/1.0/eulas][%d] createApiEulaApiMgmtBadRequest ", 400)
}

func (o *CreateAPIEulaAPIMgmtBadRequest) String() string {
	return fmt.Sprintf("[POST /api-management/1.0/eulas][%d] createApiEulaApiMgmtBadRequest ", 400)
}

func (o *CreateAPIEulaAPIMgmtBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateAPIEulaAPIMgmtInternalServerError creates a CreateAPIEulaAPIMgmtInternalServerError with default headers values
func NewCreateAPIEulaAPIMgmtInternalServerError() *CreateAPIEulaAPIMgmtInternalServerError {
	return &CreateAPIEulaAPIMgmtInternalServerError{}
}

/*
CreateAPIEulaAPIMgmtInternalServerError describes a response with status code 500, with default header values.

Server Failure
*/
type CreateAPIEulaAPIMgmtInternalServerError struct {
}

// IsSuccess returns true when this create Api eula Api mgmt internal server error response has a 2xx status code
func (o *CreateAPIEulaAPIMgmtInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create Api eula Api mgmt internal server error response has a 3xx status code
func (o *CreateAPIEulaAPIMgmtInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create Api eula Api mgmt internal server error response has a 4xx status code
func (o *CreateAPIEulaAPIMgmtInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create Api eula Api mgmt internal server error response has a 5xx status code
func (o *CreateAPIEulaAPIMgmtInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create Api eula Api mgmt internal server error response a status code equal to that given
func (o *CreateAPIEulaAPIMgmtInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create Api eula Api mgmt internal server error response
func (o *CreateAPIEulaAPIMgmtInternalServerError) Code() int {
	return 500
}

func (o *CreateAPIEulaAPIMgmtInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api-management/1.0/eulas][%d] createApiEulaApiMgmtInternalServerError ", 500)
}

func (o *CreateAPIEulaAPIMgmtInternalServerError) String() string {
	return fmt.Sprintf("[POST /api-management/1.0/eulas][%d] createApiEulaApiMgmtInternalServerError ", 500)
}

func (o *CreateAPIEulaAPIMgmtInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
