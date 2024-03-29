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

// APIDeployDto Api deploy dto
//
// swagger:model ApiDeployDto
type APIDeployDto struct {

	// The UUID of the API for the deployment.
	APIUUID string `json:"apiUuid,omitempty"`

	// The timestamp of last attempted API deployment.
	LastTimeDeployed int64 `json:"lastTimeDeployed,omitempty"`

	// The API deployment message which can be used to communicate details relating to the status of the deployment.
	Message string `json:"message,omitempty"`

	// The name of the proxy for the deployment.
	ProxyName string `json:"proxyName,omitempty"`

	// The UUID of the proxy for the deployment.
	ProxyUUID string `json:"proxyUuid,omitempty"`

	// The status of API deployment.
	// Enum: [DEPLOYED PENDING_DEPLOYMENT ERROR PENDING_UNDEPLOYMENT]
	Status string `json:"status,omitempty"`
}

// Validate validates this Api deploy dto
func (m *APIDeployDto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var apiDeployDtoTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DEPLOYED","PENDING_DEPLOYMENT","ERROR","PENDING_UNDEPLOYMENT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		apiDeployDtoTypeStatusPropEnum = append(apiDeployDtoTypeStatusPropEnum, v)
	}
}

const (

	// APIDeployDtoStatusDEPLOYED captures enum value "DEPLOYED"
	APIDeployDtoStatusDEPLOYED string = "DEPLOYED"

	// APIDeployDtoStatusPENDINGDEPLOYMENT captures enum value "PENDING_DEPLOYMENT"
	APIDeployDtoStatusPENDINGDEPLOYMENT string = "PENDING_DEPLOYMENT"

	// APIDeployDtoStatusERROR captures enum value "ERROR"
	APIDeployDtoStatusERROR string = "ERROR"

	// APIDeployDtoStatusPENDINGUNDEPLOYMENT captures enum value "PENDING_UNDEPLOYMENT"
	APIDeployDtoStatusPENDINGUNDEPLOYMENT string = "PENDING_UNDEPLOYMENT"
)

// prop value enum
func (m *APIDeployDto) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, apiDeployDtoTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *APIDeployDto) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this Api deploy dto based on context it is used
func (m *APIDeployDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *APIDeployDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIDeployDto) UnmarshalBinary(b []byte) error {
	var res APIDeployDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
