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

// AccountPlanFull account plan full
//
// swagger:model AccountPlanFull
type AccountPlanFull struct {
	AccountPlanRequest

	// Denotes this plan as the default plan. The default plan cannot be changed nor can there be more than one.
	DefaultPlan bool `json:"DefaultPlan,omitempty"`

	// The number of Organizations that have this Account Plan assigned to them.
	OrganizationUsage int64 `json:"OrganizationUsage,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *AccountPlanFull) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 AccountPlanRequest
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.AccountPlanRequest = aO0

	// AO1
	var dataAO1 struct {
		DefaultPlan bool `json:"DefaultPlan,omitempty"`

		OrganizationUsage int64 `json:"OrganizationUsage,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.DefaultPlan = dataAO1.DefaultPlan

	m.OrganizationUsage = dataAO1.OrganizationUsage

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m AccountPlanFull) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.AccountPlanRequest)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		DefaultPlan bool `json:"DefaultPlan,omitempty"`

		OrganizationUsage int64 `json:"OrganizationUsage,omitempty"`
	}

	dataAO1.DefaultPlan = m.DefaultPlan

	dataAO1.OrganizationUsage = m.OrganizationUsage

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this account plan full
func (m *AccountPlanFull) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with AccountPlanRequest
	if err := m.AccountPlanRequest.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this account plan full based on the context it is used
func (m *AccountPlanFull) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with AccountPlanRequest
	if err := m.AccountPlanRequest.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *AccountPlanFull) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountPlanFull) UnmarshalBinary(b []byte) error {
	var res AccountPlanFull
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}