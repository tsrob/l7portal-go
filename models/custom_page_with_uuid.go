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

// CustomPageWithUUID custom page with Uuid
//
// swagger:model CustomPageWithUuid
type CustomPageWithUUID struct {

	// Unique key for this entity which conforms to the UUID standard according to RFC 4122.
	UUID string `json:"uuid,omitempty"`

	CustomPageBase
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *CustomPageWithUUID) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		UUID string `json:"uuid,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	m.UUID = dataAO0.UUID

	// AO1
	var aO1 CustomPageBase
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.CustomPageBase = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m CustomPageWithUUID) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	var dataAO0 struct {
		UUID string `json:"uuid,omitempty"`
	}

	dataAO0.UUID = m.UUID

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)

	aO1, err := swag.WriteJSON(m.CustomPageBase)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this custom page with Uuid
func (m *CustomPageWithUUID) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with CustomPageBase
	if err := m.CustomPageBase.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this custom page with Uuid based on the context it is used
func (m *CustomPageWithUUID) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with CustomPageBase
	if err := m.CustomPageBase.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *CustomPageWithUUID) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomPageWithUUID) UnmarshalBinary(b []byte) error {
	var res CustomPageWithUUID
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
