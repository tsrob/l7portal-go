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

// CustomField custom field
//
// swagger:model CustomField
type CustomField struct {

	// description
	// Max Length: 65000
	Description string `json:"Description,omitempty"`

	// The Entity Type that the Custom Field belongs to which is API or APPLICATION.
	// Required: true
	EntityType *string `json:"EntityType"`

	// Name of Custom Field which must be unique. Only supports alphanumerical characters and spaces.
	// Required: true
	// Max Length: 50
	Name *string `json:"Name"`

	// A list of valid values for the Custom Field when Type is set to SINGLE_SELECT. At least one option should be supplied. The maximum length of the value is 255 characters
	// Required: true
	OptionsList interface{} `json:"OptionsList"`

	// Indicates if the Custom Field must have a value or not if it has been added to an API or Application.
	// Required: true
	Required *bool `json:"Required"`

	// Indicates if the Custom Field is active. A disabled Custom Field cannot be applied to an API or Application.
	// Required: true
	// Enum: [ENABLED DISABLED]
	Status *string `json:"Status"`

	// The type of values that are allowed for the Custom Field. SINGLE_SELECT indicates that the value must be one specified in the OptionsList property. TEXT has a maximum of 5000 characters.
	// Required: true
	// Enum: [SINGLE_SELECT TEXT]
	Type *string `json:"Type"`

	// Unique key for this entity which conforms to the UUID standard according to RFC 4122.
	// Required: true
	UUID *string `json:"Uuid"`
}

// Validate validates this custom field
func (m *CustomField) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOptionsList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequired(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUUID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CustomField) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("Description", "body", m.Description, 65000); err != nil {
		return err
	}

	return nil
}

func (m *CustomField) validateEntityType(formats strfmt.Registry) error {

	if err := validate.Required("EntityType", "body", m.EntityType); err != nil {
		return err
	}

	return nil
}

func (m *CustomField) validateName(formats strfmt.Registry) error {

	if err := validate.Required("Name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MaxLength("Name", "body", *m.Name, 50); err != nil {
		return err
	}

	return nil
}

func (m *CustomField) validateOptionsList(formats strfmt.Registry) error {

	if m.OptionsList == nil {
		return errors.Required("OptionsList", "body", nil)
	}

	return nil
}

func (m *CustomField) validateRequired(formats strfmt.Registry) error {

	if err := validate.Required("Required", "body", m.Required); err != nil {
		return err
	}

	return nil
}

var customFieldTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ENABLED","DISABLED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		customFieldTypeStatusPropEnum = append(customFieldTypeStatusPropEnum, v)
	}
}

const (

	// CustomFieldStatusENABLED captures enum value "ENABLED"
	CustomFieldStatusENABLED string = "ENABLED"

	// CustomFieldStatusDISABLED captures enum value "DISABLED"
	CustomFieldStatusDISABLED string = "DISABLED"
)

// prop value enum
func (m *CustomField) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, customFieldTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CustomField) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("Status", "body", m.Status); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("Status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

var customFieldTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SINGLE_SELECT","TEXT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		customFieldTypeTypePropEnum = append(customFieldTypeTypePropEnum, v)
	}
}

const (

	// CustomFieldTypeSINGLESELECT captures enum value "SINGLE_SELECT"
	CustomFieldTypeSINGLESELECT string = "SINGLE_SELECT"

	// CustomFieldTypeTEXT captures enum value "TEXT"
	CustomFieldTypeTEXT string = "TEXT"
)

// prop value enum
func (m *CustomField) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, customFieldTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CustomField) validateType(formats strfmt.Registry) error {

	if err := validate.Required("Type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("Type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

func (m *CustomField) validateUUID(formats strfmt.Registry) error {

	if err := validate.Required("Uuid", "body", m.UUID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this custom field based on context it is used
func (m *CustomField) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CustomField) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomField) UnmarshalBinary(b []byte) error {
	var res CustomField
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}