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

// AddOrRemoveApifromApplicationbyUUIDReader is a Reader for the AddOrRemoveApifromApplicationbyUUID structure.
type AddOrRemoveApifromApplicationbyUUIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddOrRemoveApifromApplicationbyUUIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewAddOrRemoveApifromApplicationbyUUIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddOrRemoveApifromApplicationbyUUIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAddOrRemoveApifromApplicationbyUUIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAddOrRemoveApifromApplicationbyUUIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAddOrRemoveApifromApplicationbyUUIDNoContent creates a AddOrRemoveApifromApplicationbyUUIDNoContent with default headers values
func NewAddOrRemoveApifromApplicationbyUUIDNoContent() *AddOrRemoveApifromApplicationbyUUIDNoContent {
	return &AddOrRemoveApifromApplicationbyUUIDNoContent{}
}

/*
AddOrRemoveApifromApplicationbyUUIDNoContent describes a response with status code 204, with default header values.

No Content
*/
type AddOrRemoveApifromApplicationbyUUIDNoContent struct {
}

// IsSuccess returns true when this add or remove apifrom applicationby Uuid no content response has a 2xx status code
func (o *AddOrRemoveApifromApplicationbyUUIDNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this add or remove apifrom applicationby Uuid no content response has a 3xx status code
func (o *AddOrRemoveApifromApplicationbyUUIDNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add or remove apifrom applicationby Uuid no content response has a 4xx status code
func (o *AddOrRemoveApifromApplicationbyUUIDNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this add or remove apifrom applicationby Uuid no content response has a 5xx status code
func (o *AddOrRemoveApifromApplicationbyUUIDNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this add or remove apifrom applicationby Uuid no content response a status code equal to that given
func (o *AddOrRemoveApifromApplicationbyUUIDNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the add or remove apifrom applicationby Uuid no content response
func (o *AddOrRemoveApifromApplicationbyUUIDNoContent) Code() int {
	return 204
}

func (o *AddOrRemoveApifromApplicationbyUUIDNoContent) Error() string {
	return fmt.Sprintf("[PATCH /api-management/1.0/applications/{appUuid}/apis][%d] addOrRemoveApifromApplicationbyUuidNoContent ", 204)
}

func (o *AddOrRemoveApifromApplicationbyUUIDNoContent) String() string {
	return fmt.Sprintf("[PATCH /api-management/1.0/applications/{appUuid}/apis][%d] addOrRemoveApifromApplicationbyUuidNoContent ", 204)
}

func (o *AddOrRemoveApifromApplicationbyUUIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddOrRemoveApifromApplicationbyUUIDBadRequest creates a AddOrRemoveApifromApplicationbyUUIDBadRequest with default headers values
func NewAddOrRemoveApifromApplicationbyUUIDBadRequest() *AddOrRemoveApifromApplicationbyUUIDBadRequest {
	return &AddOrRemoveApifromApplicationbyUUIDBadRequest{}
}

/*
AddOrRemoveApifromApplicationbyUUIDBadRequest describes a response with status code 400, with default header values.

Bad Request due to Invalid Uuid.
*/
type AddOrRemoveApifromApplicationbyUUIDBadRequest struct {
	Payload *models.LongerError
}

// IsSuccess returns true when this add or remove apifrom applicationby Uuid bad request response has a 2xx status code
func (o *AddOrRemoveApifromApplicationbyUUIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add or remove apifrom applicationby Uuid bad request response has a 3xx status code
func (o *AddOrRemoveApifromApplicationbyUUIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add or remove apifrom applicationby Uuid bad request response has a 4xx status code
func (o *AddOrRemoveApifromApplicationbyUUIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this add or remove apifrom applicationby Uuid bad request response has a 5xx status code
func (o *AddOrRemoveApifromApplicationbyUUIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this add or remove apifrom applicationby Uuid bad request response a status code equal to that given
func (o *AddOrRemoveApifromApplicationbyUUIDBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the add or remove apifrom applicationby Uuid bad request response
func (o *AddOrRemoveApifromApplicationbyUUIDBadRequest) Code() int {
	return 400
}

func (o *AddOrRemoveApifromApplicationbyUUIDBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api-management/1.0/applications/{appUuid}/apis][%d] addOrRemoveApifromApplicationbyUuidBadRequest  %+v", 400, o.Payload)
}

func (o *AddOrRemoveApifromApplicationbyUUIDBadRequest) String() string {
	return fmt.Sprintf("[PATCH /api-management/1.0/applications/{appUuid}/apis][%d] addOrRemoveApifromApplicationbyUuidBadRequest  %+v", 400, o.Payload)
}

func (o *AddOrRemoveApifromApplicationbyUUIDBadRequest) GetPayload() *models.LongerError {
	return o.Payload
}

func (o *AddOrRemoveApifromApplicationbyUUIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LongerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddOrRemoveApifromApplicationbyUUIDNotFound creates a AddOrRemoveApifromApplicationbyUUIDNotFound with default headers values
func NewAddOrRemoveApifromApplicationbyUUIDNotFound() *AddOrRemoveApifromApplicationbyUUIDNotFound {
	return &AddOrRemoveApifromApplicationbyUUIDNotFound{}
}

/*
AddOrRemoveApifromApplicationbyUUIDNotFound describes a response with status code 404, with default header values.

Entity not Found
*/
type AddOrRemoveApifromApplicationbyUUIDNotFound struct {
	Payload *models.ErrorEntity
}

// IsSuccess returns true when this add or remove apifrom applicationby Uuid not found response has a 2xx status code
func (o *AddOrRemoveApifromApplicationbyUUIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add or remove apifrom applicationby Uuid not found response has a 3xx status code
func (o *AddOrRemoveApifromApplicationbyUUIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add or remove apifrom applicationby Uuid not found response has a 4xx status code
func (o *AddOrRemoveApifromApplicationbyUUIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this add or remove apifrom applicationby Uuid not found response has a 5xx status code
func (o *AddOrRemoveApifromApplicationbyUUIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this add or remove apifrom applicationby Uuid not found response a status code equal to that given
func (o *AddOrRemoveApifromApplicationbyUUIDNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the add or remove apifrom applicationby Uuid not found response
func (o *AddOrRemoveApifromApplicationbyUUIDNotFound) Code() int {
	return 404
}

func (o *AddOrRemoveApifromApplicationbyUUIDNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api-management/1.0/applications/{appUuid}/apis][%d] addOrRemoveApifromApplicationbyUuidNotFound  %+v", 404, o.Payload)
}

func (o *AddOrRemoveApifromApplicationbyUUIDNotFound) String() string {
	return fmt.Sprintf("[PATCH /api-management/1.0/applications/{appUuid}/apis][%d] addOrRemoveApifromApplicationbyUuidNotFound  %+v", 404, o.Payload)
}

func (o *AddOrRemoveApifromApplicationbyUUIDNotFound) GetPayload() *models.ErrorEntity {
	return o.Payload
}

func (o *AddOrRemoveApifromApplicationbyUUIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddOrRemoveApifromApplicationbyUUIDInternalServerError creates a AddOrRemoveApifromApplicationbyUUIDInternalServerError with default headers values
func NewAddOrRemoveApifromApplicationbyUUIDInternalServerError() *AddOrRemoveApifromApplicationbyUUIDInternalServerError {
	return &AddOrRemoveApifromApplicationbyUUIDInternalServerError{}
}

/*
AddOrRemoveApifromApplicationbyUUIDInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. Certain fields are missing from the request.
*/
type AddOrRemoveApifromApplicationbyUUIDInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this add or remove apifrom applicationby Uuid internal server error response has a 2xx status code
func (o *AddOrRemoveApifromApplicationbyUUIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add or remove apifrom applicationby Uuid internal server error response has a 3xx status code
func (o *AddOrRemoveApifromApplicationbyUUIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add or remove apifrom applicationby Uuid internal server error response has a 4xx status code
func (o *AddOrRemoveApifromApplicationbyUUIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this add or remove apifrom applicationby Uuid internal server error response has a 5xx status code
func (o *AddOrRemoveApifromApplicationbyUUIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this add or remove apifrom applicationby Uuid internal server error response a status code equal to that given
func (o *AddOrRemoveApifromApplicationbyUUIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the add or remove apifrom applicationby Uuid internal server error response
func (o *AddOrRemoveApifromApplicationbyUUIDInternalServerError) Code() int {
	return 500
}

func (o *AddOrRemoveApifromApplicationbyUUIDInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api-management/1.0/applications/{appUuid}/apis][%d] addOrRemoveApifromApplicationbyUuidInternalServerError  %+v", 500, o.Payload)
}

func (o *AddOrRemoveApifromApplicationbyUUIDInternalServerError) String() string {
	return fmt.Sprintf("[PATCH /api-management/1.0/applications/{appUuid}/apis][%d] addOrRemoveApifromApplicationbyUuidInternalServerError  %+v", 500, o.Payload)
}

func (o *AddOrRemoveApifromApplicationbyUUIDInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *AddOrRemoveApifromApplicationbyUUIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
AddOrRemoveApifromApplicationbyUUIDBody add or remove apifrom applicationby UUID body
swagger:model AddOrRemoveApifromApplicationbyUUIDBody
*/
type AddOrRemoveApifromApplicationbyUUIDBody struct {

	// sources
	Sources []*models.APIReference `json:"sources"`
}

// Validate validates this add or remove apifrom applicationby UUID body
func (o *AddOrRemoveApifromApplicationbyUUIDBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSources(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddOrRemoveApifromApplicationbyUUIDBody) validateSources(formats strfmt.Registry) error {
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
					return ve.ValidateName("apis" + "." + "sources" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("apis" + "." + "sources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this add or remove apifrom applicationby UUID body based on the context it is used
func (o *AddOrRemoveApifromApplicationbyUUIDBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateSources(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddOrRemoveApifromApplicationbyUUIDBody) contextValidateSources(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Sources); i++ {

		if o.Sources[i] != nil {
			if err := o.Sources[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("apis" + "." + "sources" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("apis" + "." + "sources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddOrRemoveApifromApplicationbyUUIDBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddOrRemoveApifromApplicationbyUUIDBody) UnmarshalBinary(b []byte) error {
	var res AddOrRemoveApifromApplicationbyUUIDBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
