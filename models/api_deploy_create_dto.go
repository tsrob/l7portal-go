// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// APIDeployCreateDto Api deploy create dto
//
// swagger:model ApiDeployCreateDto
type APIDeployCreateDto struct {

	// The UUID of the proxy for the deployment.
	ProxyUUID string `json:"proxyUuid,omitempty"`
}

// Validate validates this Api deploy create dto
func (m *APIDeployCreateDto) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this Api deploy create dto based on context it is used
func (m *APIDeployCreateDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *APIDeployCreateDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIDeployCreateDto) UnmarshalBinary(b []byte) error {
	var res APIDeployCreateDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
