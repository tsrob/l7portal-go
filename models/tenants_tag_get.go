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
)

// TenantsTagGet tenants tag get
//
// swagger:model TenantsTagGet
type TenantsTagGet struct {
	ListBasic

	// results
	Results []*TenantTagGet `json:"results"`

	LinksBasic
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *TenantsTagGet) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 ListBasic
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ListBasic = aO0

	// AO1
	var dataAO1 struct {
		Results []*TenantTagGet `json:"results"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Results = dataAO1.Results

	// AO2
	var aO2 LinksBasic
	if err := swag.ReadJSON(raw, &aO2); err != nil {
		return err
	}
	m.LinksBasic = aO2

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m TenantsTagGet) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 3)

	aO0, err := swag.WriteJSON(m.ListBasic)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		Results []*TenantTagGet `json:"results"`
	}

	dataAO1.Results = m.Results

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	aO2, err := swag.WriteJSON(m.LinksBasic)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO2)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this tenants tag get
func (m *TenantsTagGet) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ListBasic
	if err := m.ListBasic.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResults(formats); err != nil {
		res = append(res, err)
	}

	// validation for a type composition with LinksBasic
	if err := m.LinksBasic.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TenantsTagGet) validateResults(formats strfmt.Registry) error {

	if swag.IsZero(m.Results) { // not required
		return nil
	}

	for i := 0; i < len(m.Results); i++ {
		if swag.IsZero(m.Results[i]) { // not required
			continue
		}

		if m.Results[i] != nil {
			if err := m.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this tenants tag get based on the context it is used
func (m *TenantsTagGet) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ListBasic
	if err := m.ListBasic.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResults(ctx, formats); err != nil {
		res = append(res, err)
	}

	// validation for a type composition with LinksBasic
	if err := m.LinksBasic.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TenantsTagGet) contextValidateResults(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Results); i++ {

		if m.Results[i] != nil {
			if err := m.Results[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TenantsTagGet) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TenantsTagGet) UnmarshalBinary(b []byte) error {
	var res TenantsTagGet
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
