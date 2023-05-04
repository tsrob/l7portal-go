// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SpecContent spec content
// Example: {"swagger":"2.0"}
//
// swagger:model SpecContent
type SpecContent struct {

	// version of swagger file
	Swagger string `json:"swagger,omitempty"`
}

// Validate validates this spec content
func (m *SpecContent) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this spec content based on context it is used
func (m *SpecContent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SpecContent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SpecContent) UnmarshalBinary(b []byte) error {
	var res SpecContent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}