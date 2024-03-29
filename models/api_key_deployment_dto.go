// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// APIKeyDeploymentDto Api key deployment dto
//
// swagger:model ApiKeyDeploymentDto
type APIKeyDeploymentDto struct {

	// api key
	APIKey string `json:"apiKey,omitempty"`

	// application Uuid
	ApplicationUUID string `json:"applicationUuid,omitempty"`

	// last time deployed
	LastTimeDeployed int64 `json:"lastTimeDeployed,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// proxy name
	ProxyName string `json:"proxyName,omitempty"`

	// proxy Uuid
	ProxyUUID string `json:"proxyUuid,omitempty"`

	// status
	Status string `json:"status,omitempty"`
}

// Validate validates this Api key deployment dto
func (m *APIKeyDeploymentDto) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this Api key deployment dto based on context it is used
func (m *APIKeyDeploymentDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *APIKeyDeploymentDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIKeyDeploymentDto) UnmarshalBinary(b []byte) error {
	var res APIKeyDeploymentDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
