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

// CustomFieldPost custom field post
//
// swagger:model CustomFieldPost
type CustomFieldPost struct {
}

// Validate validates this custom field post
func (m *CustomFieldPost) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this custom field post based on the context it is used
func (m *CustomFieldPost) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *CustomFieldPost) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomFieldPost) UnmarshalBinary(b []byte) error {
	var res CustomFieldPost
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
