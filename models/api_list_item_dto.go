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

// APIListItemDto ApiListItemDto
//
// swagger:model ApiListItemDto
type APIListItemDto struct {

	// access status
	// Enum: [PUBLIC PRIVATE]
	AccessStatus string `json:"accessStatus,omitempty"`

	// api eula Uuid
	APIEulaUUID string `json:"apiEulaUuid,omitempty"`

	// application usage
	ApplicationUsage int64 `json:"applicationUsage,omitempty"`

	// available orgs
	AvailableOrgs []string `json:"availableOrgs"`

	// create ts
	CreateTs int64 `json:"createTs,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// modify ts
	ModifyTs int64 `json:"modifyTs,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// portal status
	// Enum: [NEW ENABLED DISABLED DEPRECATED DELETED INCOMPLETE PENDING_BUNDLE PENDING_DELETE]
	PortalStatus string `json:"portalStatus,omitempty"`

	// spec filename
	SpecFilename string `json:"specFilename,omitempty"`

	// ssg service type
	// Enum: [ApiServiceType[name: REST] ApiServiceType[name: SOAP]]
	SsgServiceType string `json:"ssgServiceType,omitempty"`

	// type
	// Enum: [BASIC ADVANCED]
	Type string `json:"type,omitempty"`

	// uuid
	UUID string `json:"uuid,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this Api list item dto
func (m *APIListItemDto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccessStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePortalStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSsgServiceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var apiListItemDtoTypeAccessStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PUBLIC","PRIVATE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		apiListItemDtoTypeAccessStatusPropEnum = append(apiListItemDtoTypeAccessStatusPropEnum, v)
	}
}

const (

	// APIListItemDtoAccessStatusPUBLIC captures enum value "PUBLIC"
	APIListItemDtoAccessStatusPUBLIC string = "PUBLIC"

	// APIListItemDtoAccessStatusPRIVATE captures enum value "PRIVATE"
	APIListItemDtoAccessStatusPRIVATE string = "PRIVATE"
)

// prop value enum
func (m *APIListItemDto) validateAccessStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, apiListItemDtoTypeAccessStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *APIListItemDto) validateAccessStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.AccessStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateAccessStatusEnum("accessStatus", "body", m.AccessStatus); err != nil {
		return err
	}

	return nil
}

var apiListItemDtoTypePortalStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NEW","ENABLED","DISABLED","DEPRECATED","DELETED","INCOMPLETE","PENDING_BUNDLE","PENDING_DELETE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		apiListItemDtoTypePortalStatusPropEnum = append(apiListItemDtoTypePortalStatusPropEnum, v)
	}
}

const (

	// APIListItemDtoPortalStatusNEW captures enum value "NEW"
	APIListItemDtoPortalStatusNEW string = "NEW"

	// APIListItemDtoPortalStatusENABLED captures enum value "ENABLED"
	APIListItemDtoPortalStatusENABLED string = "ENABLED"

	// APIListItemDtoPortalStatusDISABLED captures enum value "DISABLED"
	APIListItemDtoPortalStatusDISABLED string = "DISABLED"

	// APIListItemDtoPortalStatusDEPRECATED captures enum value "DEPRECATED"
	APIListItemDtoPortalStatusDEPRECATED string = "DEPRECATED"

	// APIListItemDtoPortalStatusDELETED captures enum value "DELETED"
	APIListItemDtoPortalStatusDELETED string = "DELETED"

	// APIListItemDtoPortalStatusINCOMPLETE captures enum value "INCOMPLETE"
	APIListItemDtoPortalStatusINCOMPLETE string = "INCOMPLETE"

	// APIListItemDtoPortalStatusPENDINGBUNDLE captures enum value "PENDING_BUNDLE"
	APIListItemDtoPortalStatusPENDINGBUNDLE string = "PENDING_BUNDLE"

	// APIListItemDtoPortalStatusPENDINGDELETE captures enum value "PENDING_DELETE"
	APIListItemDtoPortalStatusPENDINGDELETE string = "PENDING_DELETE"
)

// prop value enum
func (m *APIListItemDto) validatePortalStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, apiListItemDtoTypePortalStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *APIListItemDto) validatePortalStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.PortalStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validatePortalStatusEnum("portalStatus", "body", m.PortalStatus); err != nil {
		return err
	}

	return nil
}

var apiListItemDtoTypeSsgServiceTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ApiServiceType[name: REST]","ApiServiceType[name: SOAP]"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		apiListItemDtoTypeSsgServiceTypePropEnum = append(apiListItemDtoTypeSsgServiceTypePropEnum, v)
	}
}

const (

	// APIListItemDtoSsgServiceTypeAPIServiceTypeNameREST captures enum value "ApiServiceType[name: REST]"
	APIListItemDtoSsgServiceTypeAPIServiceTypeNameREST string = "ApiServiceType[name: REST]"

	// APIListItemDtoSsgServiceTypeAPIServiceTypeNameSOAP captures enum value "ApiServiceType[name: SOAP]"
	APIListItemDtoSsgServiceTypeAPIServiceTypeNameSOAP string = "ApiServiceType[name: SOAP]"
)

// prop value enum
func (m *APIListItemDto) validateSsgServiceTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, apiListItemDtoTypeSsgServiceTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *APIListItemDto) validateSsgServiceType(formats strfmt.Registry) error {
	if swag.IsZero(m.SsgServiceType) { // not required
		return nil
	}

	// value enum
	if err := m.validateSsgServiceTypeEnum("ssgServiceType", "body", m.SsgServiceType); err != nil {
		return err
	}

	return nil
}

var apiListItemDtoTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["BASIC","ADVANCED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		apiListItemDtoTypeTypePropEnum = append(apiListItemDtoTypeTypePropEnum, v)
	}
}

const (

	// APIListItemDtoTypeBASIC captures enum value "BASIC"
	APIListItemDtoTypeBASIC string = "BASIC"

	// APIListItemDtoTypeADVANCED captures enum value "ADVANCED"
	APIListItemDtoTypeADVANCED string = "ADVANCED"
)

// prop value enum
func (m *APIListItemDto) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, apiListItemDtoTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *APIListItemDto) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this Api list item dto based on context it is used
func (m *APIListItemDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *APIListItemDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIListItemDto) UnmarshalBinary(b []byte) error {
	var res APIListItemDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
