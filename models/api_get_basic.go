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

// APIGetBasic A representation of an API using Get
//
// swagger:model ApiGetBasic
type APIGetBasic struct {

	// Indicates whether the API is public or private. A private API is only available to organizations that have it in their account plan. public APIs are public and available to all organizations. You cannot change the AccessStatus if an application is using the API. If not specified, it sets the default AccessStatus to PRIVATE.
	// Enum: [PUBLIC PRIVATE]
	AccessStatus string `json:"accessStatus,omitempty"`

	// The uuid of the End-User License Agreement applied to the API. The EULA of an API currently used by an application/s cannot be change
	APIEulaUUID string `json:"apiEulaUuid,omitempty"`

	// The number of applications that consume this API.
	ApplicationUsage int64 `json:"applicationUsage,omitempty"`

	// The time that the api is created.
	CreateTs int64 `json:"createTs,omitempty"`

	// Description of the API. Developers see only the public description. Ensure that the public description includes any information that developers require to use the API, such as its authentication requirements.
	// Max Length: 65000
	Description string `json:"description,omitempty"`

	// The uuid of the managing organization that the API belongs to.The organization can be changed.
	ManagingOrgUUID string `json:"managingOrgUuid,omitempty"`

	// The time that the api is last modified.
	ModifyTs int64 `json:"modifyTs,omitempty"`

	// Name of API which must be unique.
	// Max Length: 255
	Name string `json:"name,omitempty"`

	// Specifies the filename of the Swagger or WADL file that documents the API. If SpecFilename is specified, SpecContent should be specified as well.
	SpecFilename string `json:"specFilename,omitempty"`

	// Indicates the type of the API.
	// Enum: [REST SOAP]
	SsgServiceType string `json:"ssgServiceType,omitempty"`

	// The relative API URI that is part the API proxy URL used by developers to send requests to the API. For Portal-published APIs, the * is appended by API Portal to include all endpoints under the API. For Gateway-published APIs, the value gets set as part of the sync, as defined by the Service Resolution Path in the service.
	SsgURL string `json:"ssgUrl,omitempty"`

	// List of API tags.
	Tags []string `json:"tags"`

	// Unique key for this entity which conforms to the UUID standard according to RFC 4122.
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this Api get basic
func (m *APIGetBasic) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccessStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSsgServiceType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var apiGetBasicTypeAccessStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PUBLIC","PRIVATE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		apiGetBasicTypeAccessStatusPropEnum = append(apiGetBasicTypeAccessStatusPropEnum, v)
	}
}

const (

	// APIGetBasicAccessStatusPUBLIC captures enum value "PUBLIC"
	APIGetBasicAccessStatusPUBLIC string = "PUBLIC"

	// APIGetBasicAccessStatusPRIVATE captures enum value "PRIVATE"
	APIGetBasicAccessStatusPRIVATE string = "PRIVATE"
)

// prop value enum
func (m *APIGetBasic) validateAccessStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, apiGetBasicTypeAccessStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *APIGetBasic) validateAccessStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.AccessStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateAccessStatusEnum("accessStatus", "body", m.AccessStatus); err != nil {
		return err
	}

	return nil
}

func (m *APIGetBasic) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 65000); err != nil {
		return err
	}

	return nil
}

func (m *APIGetBasic) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MaxLength("name", "body", m.Name, 255); err != nil {
		return err
	}

	return nil
}

var apiGetBasicTypeSsgServiceTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["REST","SOAP"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		apiGetBasicTypeSsgServiceTypePropEnum = append(apiGetBasicTypeSsgServiceTypePropEnum, v)
	}
}

const (

	// APIGetBasicSsgServiceTypeREST captures enum value "REST"
	APIGetBasicSsgServiceTypeREST string = "REST"

	// APIGetBasicSsgServiceTypeSOAP captures enum value "SOAP"
	APIGetBasicSsgServiceTypeSOAP string = "SOAP"
)

// prop value enum
func (m *APIGetBasic) validateSsgServiceTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, apiGetBasicTypeSsgServiceTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *APIGetBasic) validateSsgServiceType(formats strfmt.Registry) error {
	if swag.IsZero(m.SsgServiceType) { // not required
		return nil
	}

	// value enum
	if err := m.validateSsgServiceTypeEnum("ssgServiceType", "body", m.SsgServiceType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this Api get basic based on context it is used
func (m *APIGetBasic) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *APIGetBasic) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIGetBasic) UnmarshalBinary(b []byte) error {
	var res APIGetBasic
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}