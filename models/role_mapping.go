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

// RoleMapping role mapping
//
// swagger:model RoleMapping
type RoleMapping struct {
	RoleMappingBasic

	// links
	Links *UserRoleMappingLink `json:"_links,omitempty"`

	// The uuid which identifies the role mapping
	UUID string `json:"uuid,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *RoleMapping) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 RoleMappingBasic
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.RoleMappingBasic = aO0

	// AO1
	var dataAO1 struct {
		Links *UserRoleMappingLink `json:"_links,omitempty"`

		UUID string `json:"uuid,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Links = dataAO1.Links

	m.UUID = dataAO1.UUID

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m RoleMapping) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.RoleMappingBasic)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		Links *UserRoleMappingLink `json:"_links,omitempty"`

		UUID string `json:"uuid,omitempty"`
	}

	dataAO1.Links = m.Links

	dataAO1.UUID = m.UUID

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this role mapping
func (m *RoleMapping) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with RoleMappingBasic
	if err := m.RoleMappingBasic.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RoleMapping) validateLinks(formats strfmt.Registry) error {

	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this role mapping based on the context it is used
func (m *RoleMapping) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with RoleMappingBasic
	if err := m.RoleMappingBasic.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RoleMapping) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RoleMapping) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RoleMapping) UnmarshalBinary(b []byte) error {
	var res RoleMapping
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}