// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CustomPageBase custom page base
//
// swagger:model CustomPageBase
type CustomPageBase struct {

	// Name of the custom page.
	Name string `json:"name,omitempty"`

	// Key of the Custom page key to provide additional identification besides the name.
	PageKey string `json:"pageKey,omitempty"`

	// Unique type of custom page.
	// Enum: [HOME]
	PageType string `json:"pageType,omitempty"`

	// The URI to the custom page.
	URI string `json:"uri,omitempty"`
}

// Validate validates this custom page base
func (m *CustomPageBase) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePageType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var customPageBaseTypePageTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["HOME"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		customPageBaseTypePageTypePropEnum = append(customPageBaseTypePageTypePropEnum, v)
	}
}

const (

	// CustomPageBasePageTypeHOME captures enum value "HOME"
	CustomPageBasePageTypeHOME string = "HOME"
)

// prop value enum
func (m *CustomPageBase) validatePageTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, customPageBaseTypePageTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CustomPageBase) validatePageType(formats strfmt.Registry) error {
	if swag.IsZero(m.PageType) { // not required
		return nil
	}

	// value enum
	if err := m.validatePageTypeEnum("pageType", "body", m.PageType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this custom page base based on context it is used
func (m *CustomPageBase) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CustomPageBase) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomPageBase) UnmarshalBinary(b []byte) error {
	var res CustomPageBase
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}