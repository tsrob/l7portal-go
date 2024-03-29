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

// UserResourceLinksBasic A HateoasLink consists of a link title and an href.  The link title describes what the href link refers to - for example - a link_title of 'self' will include an href to a link that can be used to reference the requested object, essentially it is a requested object's reference to itself
//
// swagger:model UserResourceLinksBasic
type UserResourceLinksBasic struct {

	// role mappings
	RoleMappings *UserResourceLinksBasicRoleMappings `json:"roleMappings,omitempty"`

	// self
	Self *UserResourceLinksBasicSelf `json:"self,omitempty"`
}

// Validate validates this user resource links basic
func (m *UserResourceLinksBasic) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRoleMappings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserResourceLinksBasic) validateRoleMappings(formats strfmt.Registry) error {
	if swag.IsZero(m.RoleMappings) { // not required
		return nil
	}

	if m.RoleMappings != nil {
		if err := m.RoleMappings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("roleMappings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("roleMappings")
			}
			return err
		}
	}

	return nil
}

func (m *UserResourceLinksBasic) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this user resource links basic based on the context it is used
func (m *UserResourceLinksBasic) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRoleMappings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserResourceLinksBasic) contextValidateRoleMappings(ctx context.Context, formats strfmt.Registry) error {

	if m.RoleMappings != nil {
		if err := m.RoleMappings.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("roleMappings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("roleMappings")
			}
			return err
		}
	}

	return nil
}

func (m *UserResourceLinksBasic) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserResourceLinksBasic) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserResourceLinksBasic) UnmarshalBinary(b []byte) error {
	var res UserResourceLinksBasic
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// UserResourceLinksBasicRoleMappings user resource links basic role mappings
//
// swagger:model UserResourceLinksBasicRoleMappings
type UserResourceLinksBasicRoleMappings struct {

	// The link that can be used to access the details of users access
	Href string `json:"href,omitempty"`
}

// Validate validates this user resource links basic role mappings
func (m *UserResourceLinksBasicRoleMappings) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this user resource links basic role mappings based on context it is used
func (m *UserResourceLinksBasicRoleMappings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserResourceLinksBasicRoleMappings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserResourceLinksBasicRoleMappings) UnmarshalBinary(b []byte) error {
	var res UserResourceLinksBasicRoleMappings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// UserResourceLinksBasicSelf user resource links basic self
//
// swagger:model UserResourceLinksBasicSelf
type UserResourceLinksBasicSelf struct {

	// The link that can be used to access the current user details
	Href string `json:"href,omitempty"`
}

// Validate validates this user resource links basic self
func (m *UserResourceLinksBasicSelf) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this user resource links basic self based on context it is used
func (m *UserResourceLinksBasicSelf) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserResourceLinksBasicSelf) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserResourceLinksBasicSelf) UnmarshalBinary(b []byte) error {
	var res UserResourceLinksBasicSelf
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
