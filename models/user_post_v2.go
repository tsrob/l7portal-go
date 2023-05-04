// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UserPostV2 user post v2
//
// swagger:model UserPostV2
type UserPostV2 struct {

	// access
	// Required: true
	Access []*UserPostV2AccessItems0 `json:"access"`

	// Unique identifier of the authentication scheme to which the user belongs
	// Required: true
	AuthConfigUUID *string `json:"authConfigUuid"`

	// Email ID of user
	// Required: true
	Email *string `json:"email"`

	// First name of user
	// Required: true
	FirstName *string `json:"firstName"`

	// Last name of user
	// Required: true
	LastName *string `json:"lastName"`

	// Language the user created to support
	// Required: true
	// Enum: [en]
	Locale *string `json:"locale"`

	// Set account password
	// Required: true
	Password *string `json:"password"`

	// Status of the user
	// Required: true
	// Enum: [ENABLED]
	Status *string `json:"status"`

	// Login identifier. This value should be unique for every user
	// Required: true
	Username *string `json:"username"`
}

// Validate validates this user post v2
func (m *UserPostV2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccess(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthConfigUUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFirstName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocale(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserPostV2) validateAccess(formats strfmt.Registry) error {

	if err := validate.Required("access", "body", m.Access); err != nil {
		return err
	}

	for i := 0; i < len(m.Access); i++ {
		if swag.IsZero(m.Access[i]) { // not required
			continue
		}

		if m.Access[i] != nil {
			if err := m.Access[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("access" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("access" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *UserPostV2) validateAuthConfigUUID(formats strfmt.Registry) error {

	if err := validate.Required("authConfigUuid", "body", m.AuthConfigUUID); err != nil {
		return err
	}

	return nil
}

func (m *UserPostV2) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("email", "body", m.Email); err != nil {
		return err
	}

	return nil
}

func (m *UserPostV2) validateFirstName(formats strfmt.Registry) error {

	if err := validate.Required("firstName", "body", m.FirstName); err != nil {
		return err
	}

	return nil
}

func (m *UserPostV2) validateLastName(formats strfmt.Registry) error {

	if err := validate.Required("lastName", "body", m.LastName); err != nil {
		return err
	}

	return nil
}

var userPostV2TypeLocalePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["en"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		userPostV2TypeLocalePropEnum = append(userPostV2TypeLocalePropEnum, v)
	}
}

const (

	// UserPostV2LocaleEn captures enum value "en"
	UserPostV2LocaleEn string = "en"
)

// prop value enum
func (m *UserPostV2) validateLocaleEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, userPostV2TypeLocalePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UserPostV2) validateLocale(formats strfmt.Registry) error {

	if err := validate.Required("locale", "body", m.Locale); err != nil {
		return err
	}

	// value enum
	if err := m.validateLocaleEnum("locale", "body", *m.Locale); err != nil {
		return err
	}

	return nil
}

func (m *UserPostV2) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("password", "body", m.Password); err != nil {
		return err
	}

	return nil
}

var userPostV2TypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ENABLED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		userPostV2TypeStatusPropEnum = append(userPostV2TypeStatusPropEnum, v)
	}
}

const (

	// UserPostV2StatusENABLED captures enum value "ENABLED"
	UserPostV2StatusENABLED string = "ENABLED"
)

// prop value enum
func (m *UserPostV2) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, userPostV2TypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UserPostV2) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

func (m *UserPostV2) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("username", "body", m.Username); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this user post v2 based on the context it is used
func (m *UserPostV2) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAccess(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserPostV2) contextValidateAccess(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Access); i++ {

		if m.Access[i] != nil {
			if err := m.Access[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("access" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("access" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserPostV2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserPostV2) UnmarshalBinary(b []byte) error {
	var res UserPostV2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// UserPostV2AccessItems0 user post v2 access items0
//
// swagger:model UserPostV2AccessItems0
type UserPostV2AccessItems0 struct {

	// org Uuid
	OrgUUID string `json:"orgUuid,omitempty"`

	// role Uuid
	RoleUUID string `json:"roleUuid,omitempty"`
}

// Validate validates this user post v2 access items0
func (m *UserPostV2AccessItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this user post v2 access items0 based on context it is used
func (m *UserPostV2AccessItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserPostV2AccessItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserPostV2AccessItems0) UnmarshalBinary(b []byte) error {
	var res UserPostV2AccessItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
