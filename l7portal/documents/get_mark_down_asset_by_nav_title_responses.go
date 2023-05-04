// Code generated by go-swagger; DO NOT EDIT.

package documents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/tsrob/l7portal-go/models"
)

// GetMarkDownAssetByNavTitleReader is a Reader for the GetMarkDownAssetByNavTitle structure.
type GetMarkDownAssetByNavTitleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMarkDownAssetByNavTitleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMarkDownAssetByNavTitleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetMarkDownAssetByNavTitleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetMarkDownAssetByNavTitleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetMarkDownAssetByNavTitleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetMarkDownAssetByNavTitleOK creates a GetMarkDownAssetByNavTitleOK with default headers values
func NewGetMarkDownAssetByNavTitleOK() *GetMarkDownAssetByNavTitleOK {
	return &GetMarkDownAssetByNavTitleOK{}
}

/*
GetMarkDownAssetByNavTitleOK describes a response with status code 200, with default header values.

OK
*/
type GetMarkDownAssetByNavTitleOK struct {
	Payload *models.DocumentServiceGet
}

// IsSuccess returns true when this get mark down asset by nav title o k response has a 2xx status code
func (o *GetMarkDownAssetByNavTitleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get mark down asset by nav title o k response has a 3xx status code
func (o *GetMarkDownAssetByNavTitleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get mark down asset by nav title o k response has a 4xx status code
func (o *GetMarkDownAssetByNavTitleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get mark down asset by nav title o k response has a 5xx status code
func (o *GetMarkDownAssetByNavTitleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get mark down asset by nav title o k response a status code equal to that given
func (o *GetMarkDownAssetByNavTitleOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get mark down asset by nav title o k response
func (o *GetMarkDownAssetByNavTitleOK) Code() int {
	return 200
}

func (o *GetMarkDownAssetByNavTitleOK) Error() string {
	return fmt.Sprintf("[GET /document-service/1.0/docs/{type}/{typeUuid}/{navtitle}][%d] getMarkDownAssetByNavTitleOK  %+v", 200, o.Payload)
}

func (o *GetMarkDownAssetByNavTitleOK) String() string {
	return fmt.Sprintf("[GET /document-service/1.0/docs/{type}/{typeUuid}/{navtitle}][%d] getMarkDownAssetByNavTitleOK  %+v", 200, o.Payload)
}

func (o *GetMarkDownAssetByNavTitleOK) GetPayload() *models.DocumentServiceGet {
	return o.Payload
}

func (o *GetMarkDownAssetByNavTitleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DocumentServiceGet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMarkDownAssetByNavTitleBadRequest creates a GetMarkDownAssetByNavTitleBadRequest with default headers values
func NewGetMarkDownAssetByNavTitleBadRequest() *GetMarkDownAssetByNavTitleBadRequest {
	return &GetMarkDownAssetByNavTitleBadRequest{}
}

/*
GetMarkDownAssetByNavTitleBadRequest describes a response with status code 400, with default header values.

Bad Request due to a validation error.
*/
type GetMarkDownAssetByNavTitleBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get mark down asset by nav title bad request response has a 2xx status code
func (o *GetMarkDownAssetByNavTitleBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get mark down asset by nav title bad request response has a 3xx status code
func (o *GetMarkDownAssetByNavTitleBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get mark down asset by nav title bad request response has a 4xx status code
func (o *GetMarkDownAssetByNavTitleBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get mark down asset by nav title bad request response has a 5xx status code
func (o *GetMarkDownAssetByNavTitleBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get mark down asset by nav title bad request response a status code equal to that given
func (o *GetMarkDownAssetByNavTitleBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get mark down asset by nav title bad request response
func (o *GetMarkDownAssetByNavTitleBadRequest) Code() int {
	return 400
}

func (o *GetMarkDownAssetByNavTitleBadRequest) Error() string {
	return fmt.Sprintf("[GET /document-service/1.0/docs/{type}/{typeUuid}/{navtitle}][%d] getMarkDownAssetByNavTitleBadRequest  %+v", 400, o.Payload)
}

func (o *GetMarkDownAssetByNavTitleBadRequest) String() string {
	return fmt.Sprintf("[GET /document-service/1.0/docs/{type}/{typeUuid}/{navtitle}][%d] getMarkDownAssetByNavTitleBadRequest  %+v", 400, o.Payload)
}

func (o *GetMarkDownAssetByNavTitleBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetMarkDownAssetByNavTitleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMarkDownAssetByNavTitleNotFound creates a GetMarkDownAssetByNavTitleNotFound with default headers values
func NewGetMarkDownAssetByNavTitleNotFound() *GetMarkDownAssetByNavTitleNotFound {
	return &GetMarkDownAssetByNavTitleNotFound{}
}

/*
GetMarkDownAssetByNavTitleNotFound describes a response with status code 404, with default header values.

Entity not found
*/
type GetMarkDownAssetByNavTitleNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get mark down asset by nav title not found response has a 2xx status code
func (o *GetMarkDownAssetByNavTitleNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get mark down asset by nav title not found response has a 3xx status code
func (o *GetMarkDownAssetByNavTitleNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get mark down asset by nav title not found response has a 4xx status code
func (o *GetMarkDownAssetByNavTitleNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get mark down asset by nav title not found response has a 5xx status code
func (o *GetMarkDownAssetByNavTitleNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get mark down asset by nav title not found response a status code equal to that given
func (o *GetMarkDownAssetByNavTitleNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get mark down asset by nav title not found response
func (o *GetMarkDownAssetByNavTitleNotFound) Code() int {
	return 404
}

func (o *GetMarkDownAssetByNavTitleNotFound) Error() string {
	return fmt.Sprintf("[GET /document-service/1.0/docs/{type}/{typeUuid}/{navtitle}][%d] getMarkDownAssetByNavTitleNotFound  %+v", 404, o.Payload)
}

func (o *GetMarkDownAssetByNavTitleNotFound) String() string {
	return fmt.Sprintf("[GET /document-service/1.0/docs/{type}/{typeUuid}/{navtitle}][%d] getMarkDownAssetByNavTitleNotFound  %+v", 404, o.Payload)
}

func (o *GetMarkDownAssetByNavTitleNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetMarkDownAssetByNavTitleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMarkDownAssetByNavTitleInternalServerError creates a GetMarkDownAssetByNavTitleInternalServerError with default headers values
func NewGetMarkDownAssetByNavTitleInternalServerError() *GetMarkDownAssetByNavTitleInternalServerError {
	return &GetMarkDownAssetByNavTitleInternalServerError{}
}

/*
GetMarkDownAssetByNavTitleInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. Certain fields are missing from the request.
*/
type GetMarkDownAssetByNavTitleInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get mark down asset by nav title internal server error response has a 2xx status code
func (o *GetMarkDownAssetByNavTitleInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get mark down asset by nav title internal server error response has a 3xx status code
func (o *GetMarkDownAssetByNavTitleInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get mark down asset by nav title internal server error response has a 4xx status code
func (o *GetMarkDownAssetByNavTitleInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get mark down asset by nav title internal server error response has a 5xx status code
func (o *GetMarkDownAssetByNavTitleInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get mark down asset by nav title internal server error response a status code equal to that given
func (o *GetMarkDownAssetByNavTitleInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get mark down asset by nav title internal server error response
func (o *GetMarkDownAssetByNavTitleInternalServerError) Code() int {
	return 500
}

func (o *GetMarkDownAssetByNavTitleInternalServerError) Error() string {
	return fmt.Sprintf("[GET /document-service/1.0/docs/{type}/{typeUuid}/{navtitle}][%d] getMarkDownAssetByNavTitleInternalServerError  %+v", 500, o.Payload)
}

func (o *GetMarkDownAssetByNavTitleInternalServerError) String() string {
	return fmt.Sprintf("[GET /document-service/1.0/docs/{type}/{typeUuid}/{navtitle}][%d] getMarkDownAssetByNavTitleInternalServerError  %+v", 500, o.Payload)
}

func (o *GetMarkDownAssetByNavTitleInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetMarkDownAssetByNavTitleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
