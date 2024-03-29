// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// AppRequestType One of the application request types.
//
// swagger:model AppRequestType
type AppRequestType string

func NewAppRequestType(value AppRequestType) *AppRequestType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated AppRequestType.
func (m AppRequestType) Pointer() *AppRequestType {
	return &m
}

const (

	// AppRequestTypeAPPLICATION captures enum value "APPLICATION"
	AppRequestTypeAPPLICATION AppRequestType = "APPLICATION"

	// AppRequestTypeAPPLICATIONDETAILS captures enum value "APPLICATION_DETAILS"
	AppRequestTypeAPPLICATIONDETAILS AppRequestType = "APPLICATION_DETAILS"

	// AppRequestTypeAPPLICATIONAPIS captures enum value "APPLICATION_APIS"
	AppRequestTypeAPPLICATIONAPIS AppRequestType = "APPLICATION_APIS"

	// AppRequestTypeAPPLICATIONAPIGROUPS captures enum value "APPLICATION_API_GROUPS"
	AppRequestTypeAPPLICATIONAPIGROUPS AppRequestType = "APPLICATION_API_GROUPS"

	// AppRequestTypeAPPLICATIONCUSTOMFIELDS captures enum value "APPLICATION_CUSTOM_FIELDS"
	AppRequestTypeAPPLICATIONCUSTOMFIELDS AppRequestType = "APPLICATION_CUSTOM_FIELDS"

	// AppRequestTypeAPPLICATIONAPIKEYS captures enum value "APPLICATION_API_KEYS"
	AppRequestTypeAPPLICATIONAPIKEYS AppRequestType = "APPLICATION_API_KEYS"

	// AppRequestTypeAPPLICATIONAPIPLANS captures enum value "APPLICATION_API_PLANS"
	AppRequestTypeAPPLICATIONAPIPLANS AppRequestType = "APPLICATION_API_PLANS"
)

// for schema
var appRequestTypeEnum []interface{}

func init() {
	var res []AppRequestType
	if err := json.Unmarshal([]byte(`["APPLICATION","APPLICATION_DETAILS","APPLICATION_APIS","APPLICATION_API_GROUPS","APPLICATION_CUSTOM_FIELDS","APPLICATION_API_KEYS","APPLICATION_API_PLANS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		appRequestTypeEnum = append(appRequestTypeEnum, v)
	}
}

func (m AppRequestType) validateAppRequestTypeEnum(path, location string, value AppRequestType) error {
	if err := validate.EnumCase(path, location, value, appRequestTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this app request type
func (m AppRequestType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAppRequestTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this app request type based on context it is used
func (m AppRequestType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
