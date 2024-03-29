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

// DocumentServicePut document service put
//
// swagger:model DocumentServicePut
type DocumentServicePut struct {
}

// Validate validates this document service put
func (m *DocumentServicePut) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this document service put based on the context it is used
func (m *DocumentServicePut) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *DocumentServicePut) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DocumentServicePut) UnmarshalBinary(b []byte) error {
	var res DocumentServicePut
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
