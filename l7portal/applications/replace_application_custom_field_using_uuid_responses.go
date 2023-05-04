// Code generated by go-swagger; DO NOT EDIT.

package applications

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/tsrob/l7portal-go/models"
)

// ReplaceApplicationCustomFieldUsingUUIDReader is a Reader for the ReplaceApplicationCustomFieldUsingUUID structure.
type ReplaceApplicationCustomFieldUsingUUIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceApplicationCustomFieldUsingUUIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewReplaceApplicationCustomFieldUsingUUIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewReplaceApplicationCustomFieldUsingUUIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewReplaceApplicationCustomFieldUsingUUIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewReplaceApplicationCustomFieldUsingUUIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReplaceApplicationCustomFieldUsingUUIDNoContent creates a ReplaceApplicationCustomFieldUsingUUIDNoContent with default headers values
func NewReplaceApplicationCustomFieldUsingUUIDNoContent() *ReplaceApplicationCustomFieldUsingUUIDNoContent {
	return &ReplaceApplicationCustomFieldUsingUUIDNoContent{}
}

/*
ReplaceApplicationCustomFieldUsingUUIDNoContent describes a response with status code 204, with default header values.

No Content
*/
type ReplaceApplicationCustomFieldUsingUUIDNoContent struct {
}

// IsSuccess returns true when this replace application custom field using Uuid no content response has a 2xx status code
func (o *ReplaceApplicationCustomFieldUsingUUIDNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this replace application custom field using Uuid no content response has a 3xx status code
func (o *ReplaceApplicationCustomFieldUsingUUIDNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this replace application custom field using Uuid no content response has a 4xx status code
func (o *ReplaceApplicationCustomFieldUsingUUIDNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this replace application custom field using Uuid no content response has a 5xx status code
func (o *ReplaceApplicationCustomFieldUsingUUIDNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this replace application custom field using Uuid no content response a status code equal to that given
func (o *ReplaceApplicationCustomFieldUsingUUIDNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the replace application custom field using Uuid no content response
func (o *ReplaceApplicationCustomFieldUsingUUIDNoContent) Code() int {
	return 204
}

func (o *ReplaceApplicationCustomFieldUsingUUIDNoContent) Error() string {
	return fmt.Sprintf("[PUT /api-management/1.0/applications/{appUuid}/custom-fields][%d] replaceApplicationCustomFieldUsingUuidNoContent ", 204)
}

func (o *ReplaceApplicationCustomFieldUsingUUIDNoContent) String() string {
	return fmt.Sprintf("[PUT /api-management/1.0/applications/{appUuid}/custom-fields][%d] replaceApplicationCustomFieldUsingUuidNoContent ", 204)
}

func (o *ReplaceApplicationCustomFieldUsingUUIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewReplaceApplicationCustomFieldUsingUUIDBadRequest creates a ReplaceApplicationCustomFieldUsingUUIDBadRequest with default headers values
func NewReplaceApplicationCustomFieldUsingUUIDBadRequest() *ReplaceApplicationCustomFieldUsingUUIDBadRequest {
	return &ReplaceApplicationCustomFieldUsingUUIDBadRequest{}
}

/*
ReplaceApplicationCustomFieldUsingUUIDBadRequest describes a response with status code 400, with default header values.

Bad Request due to a validation error.
*/
type ReplaceApplicationCustomFieldUsingUUIDBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this replace application custom field using Uuid bad request response has a 2xx status code
func (o *ReplaceApplicationCustomFieldUsingUUIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this replace application custom field using Uuid bad request response has a 3xx status code
func (o *ReplaceApplicationCustomFieldUsingUUIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this replace application custom field using Uuid bad request response has a 4xx status code
func (o *ReplaceApplicationCustomFieldUsingUUIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this replace application custom field using Uuid bad request response has a 5xx status code
func (o *ReplaceApplicationCustomFieldUsingUUIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this replace application custom field using Uuid bad request response a status code equal to that given
func (o *ReplaceApplicationCustomFieldUsingUUIDBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the replace application custom field using Uuid bad request response
func (o *ReplaceApplicationCustomFieldUsingUUIDBadRequest) Code() int {
	return 400
}

func (o *ReplaceApplicationCustomFieldUsingUUIDBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api-management/1.0/applications/{appUuid}/custom-fields][%d] replaceApplicationCustomFieldUsingUuidBadRequest  %+v", 400, o.Payload)
}

func (o *ReplaceApplicationCustomFieldUsingUUIDBadRequest) String() string {
	return fmt.Sprintf("[PUT /api-management/1.0/applications/{appUuid}/custom-fields][%d] replaceApplicationCustomFieldUsingUuidBadRequest  %+v", 400, o.Payload)
}

func (o *ReplaceApplicationCustomFieldUsingUUIDBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceApplicationCustomFieldUsingUUIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceApplicationCustomFieldUsingUUIDNotFound creates a ReplaceApplicationCustomFieldUsingUUIDNotFound with default headers values
func NewReplaceApplicationCustomFieldUsingUUIDNotFound() *ReplaceApplicationCustomFieldUsingUUIDNotFound {
	return &ReplaceApplicationCustomFieldUsingUUIDNotFound{}
}

/*
ReplaceApplicationCustomFieldUsingUUIDNotFound describes a response with status code 404, with default header values.

Entity not found
*/
type ReplaceApplicationCustomFieldUsingUUIDNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this replace application custom field using Uuid not found response has a 2xx status code
func (o *ReplaceApplicationCustomFieldUsingUUIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this replace application custom field using Uuid not found response has a 3xx status code
func (o *ReplaceApplicationCustomFieldUsingUUIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this replace application custom field using Uuid not found response has a 4xx status code
func (o *ReplaceApplicationCustomFieldUsingUUIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this replace application custom field using Uuid not found response has a 5xx status code
func (o *ReplaceApplicationCustomFieldUsingUUIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this replace application custom field using Uuid not found response a status code equal to that given
func (o *ReplaceApplicationCustomFieldUsingUUIDNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the replace application custom field using Uuid not found response
func (o *ReplaceApplicationCustomFieldUsingUUIDNotFound) Code() int {
	return 404
}

func (o *ReplaceApplicationCustomFieldUsingUUIDNotFound) Error() string {
	return fmt.Sprintf("[PUT /api-management/1.0/applications/{appUuid}/custom-fields][%d] replaceApplicationCustomFieldUsingUuidNotFound  %+v", 404, o.Payload)
}

func (o *ReplaceApplicationCustomFieldUsingUUIDNotFound) String() string {
	return fmt.Sprintf("[PUT /api-management/1.0/applications/{appUuid}/custom-fields][%d] replaceApplicationCustomFieldUsingUuidNotFound  %+v", 404, o.Payload)
}

func (o *ReplaceApplicationCustomFieldUsingUUIDNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceApplicationCustomFieldUsingUUIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceApplicationCustomFieldUsingUUIDInternalServerError creates a ReplaceApplicationCustomFieldUsingUUIDInternalServerError with default headers values
func NewReplaceApplicationCustomFieldUsingUUIDInternalServerError() *ReplaceApplicationCustomFieldUsingUUIDInternalServerError {
	return &ReplaceApplicationCustomFieldUsingUUIDInternalServerError{}
}

/*
ReplaceApplicationCustomFieldUsingUUIDInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. Certain fields are missing from the request.
*/
type ReplaceApplicationCustomFieldUsingUUIDInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this replace application custom field using Uuid internal server error response has a 2xx status code
func (o *ReplaceApplicationCustomFieldUsingUUIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this replace application custom field using Uuid internal server error response has a 3xx status code
func (o *ReplaceApplicationCustomFieldUsingUUIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this replace application custom field using Uuid internal server error response has a 4xx status code
func (o *ReplaceApplicationCustomFieldUsingUUIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this replace application custom field using Uuid internal server error response has a 5xx status code
func (o *ReplaceApplicationCustomFieldUsingUUIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this replace application custom field using Uuid internal server error response a status code equal to that given
func (o *ReplaceApplicationCustomFieldUsingUUIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the replace application custom field using Uuid internal server error response
func (o *ReplaceApplicationCustomFieldUsingUUIDInternalServerError) Code() int {
	return 500
}

func (o *ReplaceApplicationCustomFieldUsingUUIDInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api-management/1.0/applications/{appUuid}/custom-fields][%d] replaceApplicationCustomFieldUsingUuidInternalServerError  %+v", 500, o.Payload)
}

func (o *ReplaceApplicationCustomFieldUsingUUIDInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api-management/1.0/applications/{appUuid}/custom-fields][%d] replaceApplicationCustomFieldUsingUuidInternalServerError  %+v", 500, o.Payload)
}

func (o *ReplaceApplicationCustomFieldUsingUUIDInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceApplicationCustomFieldUsingUUIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
ReplaceApplicationCustomFieldUsingUUIDBody replace application custom field using UUID body
swagger:model ReplaceApplicationCustomFieldUsingUUIDBody
*/
type ReplaceApplicationCustomFieldUsingUUIDBody struct {

	// sources
	Sources []*models.CustomFieldValueDto `json:"sources"`
}

// Validate validates this replace application custom field using UUID body
func (o *ReplaceApplicationCustomFieldUsingUUIDBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSources(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ReplaceApplicationCustomFieldUsingUUIDBody) validateSources(formats strfmt.Registry) error {
	if swag.IsZero(o.Sources) { // not required
		return nil
	}

	for i := 0; i < len(o.Sources); i++ {
		if swag.IsZero(o.Sources[i]) { // not required
			continue
		}

		if o.Sources[i] != nil {
			if err := o.Sources[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("custom-fields" + "." + "sources" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("custom-fields" + "." + "sources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this replace application custom field using UUID body based on the context it is used
func (o *ReplaceApplicationCustomFieldUsingUUIDBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateSources(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ReplaceApplicationCustomFieldUsingUUIDBody) contextValidateSources(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Sources); i++ {

		if o.Sources[i] != nil {
			if err := o.Sources[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("custom-fields" + "." + "sources" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("custom-fields" + "." + "sources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ReplaceApplicationCustomFieldUsingUUIDBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ReplaceApplicationCustomFieldUsingUUIDBody) UnmarshalBinary(b []byte) error {
	var res ReplaceApplicationCustomFieldUsingUUIDBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}