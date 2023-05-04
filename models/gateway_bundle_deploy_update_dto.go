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

// GatewayBundleDeployUpdateDto gateway bundle deploy update dto
//
// swagger:model GatewayBundleDeployUpdateDto
type GatewayBundleDeployUpdateDto struct {

	// The gateway bundle deployment message which can be used to communicate details relating to the status of the deployment.
	Message string `json:"message,omitempty"`

	// The status of gateway bundle deployment.
	// Required: true
	// Enum: [DEPLOYED PENDING_DEPLOYMENT ERROR PENDING_UNDEPLOYMENT]
	Status *string `json:"status"`
}

// Validate validates this gateway bundle deploy update dto
func (m *GatewayBundleDeployUpdateDto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var gatewayBundleDeployUpdateDtoTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DEPLOYED","PENDING_DEPLOYMENT","ERROR","PENDING_UNDEPLOYMENT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		gatewayBundleDeployUpdateDtoTypeStatusPropEnum = append(gatewayBundleDeployUpdateDtoTypeStatusPropEnum, v)
	}
}

const (

	// GatewayBundleDeployUpdateDtoStatusDEPLOYED captures enum value "DEPLOYED"
	GatewayBundleDeployUpdateDtoStatusDEPLOYED string = "DEPLOYED"

	// GatewayBundleDeployUpdateDtoStatusPENDINGDEPLOYMENT captures enum value "PENDING_DEPLOYMENT"
	GatewayBundleDeployUpdateDtoStatusPENDINGDEPLOYMENT string = "PENDING_DEPLOYMENT"

	// GatewayBundleDeployUpdateDtoStatusERROR captures enum value "ERROR"
	GatewayBundleDeployUpdateDtoStatusERROR string = "ERROR"

	// GatewayBundleDeployUpdateDtoStatusPENDINGUNDEPLOYMENT captures enum value "PENDING_UNDEPLOYMENT"
	GatewayBundleDeployUpdateDtoStatusPENDINGUNDEPLOYMENT string = "PENDING_UNDEPLOYMENT"
)

// prop value enum
func (m *GatewayBundleDeployUpdateDto) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, gatewayBundleDeployUpdateDtoTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *GatewayBundleDeployUpdateDto) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this gateway bundle deploy update dto based on context it is used
func (m *GatewayBundleDeployUpdateDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GatewayBundleDeployUpdateDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GatewayBundleDeployUpdateDto) UnmarshalBinary(b []byte) error {
	var res GatewayBundleDeployUpdateDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}