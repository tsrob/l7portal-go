// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PolicyTemplates policy templates
//
// swagger:model PolicyTemplates
type PolicyTemplates struct {

	// Unique key for this entity which conforms to the UUID standard according to RFC 4122.
	UUID string `json:"Uuid,omitempty"`

	// arguments
	Arguments []*PolicyTemplateArguments `json:"arguments"`

	// description
	Description string `json:"description,omitempty"`

	// The encapsulation assertion ID of the policy template in the gateway
	EncassID string `json:"encassId,omitempty"`

	// fragment details
	FragmentDetails *PolicyTemplateFragmentDetails `json:"fragmentDetails,omitempty"`

	// Name of Policy Template which must be unique.
	// Max Length: 50
	Name string `json:"name,omitempty"`

	// The policy GUID of the policy template in the gateway.
	PolicyGUID string `json:"policyGuid,omitempty"`
}

// Validate validates this policy templates
func (m *PolicyTemplates) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArguments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFragmentDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PolicyTemplates) validateArguments(formats strfmt.Registry) error {
	if swag.IsZero(m.Arguments) { // not required
		return nil
	}

	for i := 0; i < len(m.Arguments); i++ {
		if swag.IsZero(m.Arguments[i]) { // not required
			continue
		}

		if m.Arguments[i] != nil {
			if err := m.Arguments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("arguments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("arguments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PolicyTemplates) validateFragmentDetails(formats strfmt.Registry) error {
	if swag.IsZero(m.FragmentDetails) { // not required
		return nil
	}

	if m.FragmentDetails != nil {
		if err := m.FragmentDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fragmentDetails")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fragmentDetails")
			}
			return err
		}
	}

	return nil
}

func (m *PolicyTemplates) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MaxLength("name", "body", m.Name, 50); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this policy templates based on the context it is used
func (m *PolicyTemplates) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateArguments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFragmentDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PolicyTemplates) contextValidateArguments(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Arguments); i++ {

		if m.Arguments[i] != nil {
			if err := m.Arguments[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("arguments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("arguments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PolicyTemplates) contextValidateFragmentDetails(ctx context.Context, formats strfmt.Registry) error {

	if m.FragmentDetails != nil {
		if err := m.FragmentDetails.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fragmentDetails")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fragmentDetails")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PolicyTemplates) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PolicyTemplates) UnmarshalBinary(b []byte) error {
	var res PolicyTemplates
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
