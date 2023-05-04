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

// LinksApisBody links apis body
//
// swagger:model LinksApisBody
type LinksApisBody struct {
	LinkBasicBody

	// api usage
	APIUsage *LinksApisBodyAO1APIUsage `json:"apiUsage,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *LinksApisBody) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 LinkBasicBody
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.LinkBasicBody = aO0

	// AO1
	var dataAO1 struct {
		APIUsage *LinksApisBodyAO1APIUsage `json:"apiUsage,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.APIUsage = dataAO1.APIUsage

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m LinksApisBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.LinkBasicBody)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		APIUsage *LinksApisBodyAO1APIUsage `json:"apiUsage,omitempty"`
	}

	dataAO1.APIUsage = m.APIUsage

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this links apis body
func (m *LinksApisBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with LinkBasicBody
	if err := m.LinkBasicBody.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAPIUsage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LinksApisBody) validateAPIUsage(formats strfmt.Registry) error {

	if swag.IsZero(m.APIUsage) { // not required
		return nil
	}

	if m.APIUsage != nil {
		if err := m.APIUsage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("apiUsage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("apiUsage")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this links apis body based on the context it is used
func (m *LinksApisBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with LinkBasicBody
	if err := m.LinkBasicBody.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAPIUsage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LinksApisBody) contextValidateAPIUsage(ctx context.Context, formats strfmt.Registry) error {

	if m.APIUsage != nil {
		if err := m.APIUsage.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("apiUsage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("apiUsage")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LinksApisBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LinksApisBody) UnmarshalBinary(b []byte) error {
	var res LinksApisBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// LinksApisBodyAO1APIUsage links apis body a o1 API usage
//
// swagger:model LinksApisBodyAO1APIUsage
type LinksApisBodyAO1APIUsage struct {

	// href
	Href string `json:"href,omitempty"`
}

// Validate validates this links apis body a o1 API usage
func (m *LinksApisBodyAO1APIUsage) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this links apis body a o1 API usage based on context it is used
func (m *LinksApisBodyAO1APIUsage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LinksApisBodyAO1APIUsage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LinksApisBodyAO1APIUsage) UnmarshalBinary(b []byte) error {
	var res LinksApisBodyAO1APIUsage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}