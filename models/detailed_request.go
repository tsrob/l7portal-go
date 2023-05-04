// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DetailedRequest detailed request
//
// swagger:model DetailedRequest
type DetailedRequest struct {

	// application excerpt
	ApplicationExcerpt *ApplicationDto `json:"applicationExcerpt,omitempty"`

	// Original data before the request
	Original interface{} `json:"original,omitempty"`

	// request
	Request *Request `json:"request,omitempty"`

	// New data the request try to update to
	UpdateReqeust interface{} `json:"updateReqeust,omitempty"`
}

// Validate validates this detailed request
func (m *DetailedRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApplicationExcerpt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequest(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DetailedRequest) validateApplicationExcerpt(formats strfmt.Registry) error {
	if swag.IsZero(m.ApplicationExcerpt) { // not required
		return nil
	}

	if m.ApplicationExcerpt != nil {
		if err := m.ApplicationExcerpt.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("applicationExcerpt")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("applicationExcerpt")
			}
			return err
		}
	}

	return nil
}

func (m *DetailedRequest) validateRequest(formats strfmt.Registry) error {
	if swag.IsZero(m.Request) { // not required
		return nil
	}

	if m.Request != nil {
		if err := m.Request.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("request")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("request")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this detailed request based on the context it is used
func (m *DetailedRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateApplicationExcerpt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRequest(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DetailedRequest) contextValidateApplicationExcerpt(ctx context.Context, formats strfmt.Registry) error {

	if m.ApplicationExcerpt != nil {
		if err := m.ApplicationExcerpt.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("applicationExcerpt")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("applicationExcerpt")
			}
			return err
		}
	}

	return nil
}

func (m *DetailedRequest) contextValidateRequest(ctx context.Context, formats strfmt.Registry) error {

	if m.Request != nil {
		if err := m.Request.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("request")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("request")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DetailedRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DetailedRequest) UnmarshalBinary(b []byte) error {
	var res DetailedRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}