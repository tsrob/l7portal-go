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

// GetAPIEulaReader is a Reader for the GetAPIEula structure.
type GetAPIEulaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIEulaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPIEulaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetAPIEulaNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAPIEulaOK creates a GetAPIEulaOK with default headers values
func NewGetAPIEulaOK() *GetAPIEulaOK {
	return &GetAPIEulaOK{}
}

/*
GetAPIEulaOK describes a response with status code 200, with default header values.

GetAPIEulaOK get Api eula o k
*/
type GetAPIEulaOK struct {
	Payload *models.APIEula
}

// IsSuccess returns true when this get Api eula o k response has a 2xx status code
func (o *GetAPIEulaOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get Api eula o k response has a 3xx status code
func (o *GetAPIEulaOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Api eula o k response has a 4xx status code
func (o *GetAPIEulaOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Api eula o k response has a 5xx status code
func (o *GetAPIEulaOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get Api eula o k response a status code equal to that given
func (o *GetAPIEulaOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get Api eula o k response
func (o *GetAPIEulaOK) Code() int {
	return 200
}

func (o *GetAPIEulaOK) Error() string {
	return fmt.Sprintf("[GET /ApiEulas('{uuid}')][%d] getApiEulaOK  %+v", 200, o.Payload)
}

func (o *GetAPIEulaOK) String() string {
	return fmt.Sprintf("[GET /ApiEulas('{uuid}')][%d] getApiEulaOK  %+v", 200, o.Payload)
}

func (o *GetAPIEulaOK) GetPayload() *models.APIEula {
	return o.Payload
}

func (o *GetAPIEulaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIEula)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIEulaNotFound creates a GetAPIEulaNotFound with default headers values
func NewGetAPIEulaNotFound() *GetAPIEulaNotFound {
	return &GetAPIEulaNotFound{}
}

/*
GetAPIEulaNotFound describes a response with status code 404, with default header values.

Entity not found
*/
type GetAPIEulaNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get Api eula not found response has a 2xx status code
func (o *GetAPIEulaNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get Api eula not found response has a 3xx status code
func (o *GetAPIEulaNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Api eula not found response has a 4xx status code
func (o *GetAPIEulaNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get Api eula not found response has a 5xx status code
func (o *GetAPIEulaNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get Api eula not found response a status code equal to that given
func (o *GetAPIEulaNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get Api eula not found response
func (o *GetAPIEulaNotFound) Code() int {
	return 404
}

func (o *GetAPIEulaNotFound) Error() string {
	return fmt.Sprintf("[GET /ApiEulas('{uuid}')][%d] getApiEulaNotFound  %+v", 404, o.Payload)
}

func (o *GetAPIEulaNotFound) String() string {
	return fmt.Sprintf("[GET /ApiEulas('{uuid}')][%d] getApiEulaNotFound  %+v", 404, o.Payload)
}

func (o *GetAPIEulaNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAPIEulaNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
