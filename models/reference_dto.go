// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ReferenceDto ReferenceDto
//
// swagger:model ReferenceDto
type ReferenceDto struct {

	// uuid
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this reference dto
func (m *ReferenceDto) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this reference dto based on context it is used
func (m *ReferenceDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ReferenceDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReferenceDto) UnmarshalBinary(b []byte) error {
	var res ReferenceDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
