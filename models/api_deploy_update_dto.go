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

// APIDeployUpdateDto Api deploy update dto
//
// swagger:model ApiDeployUpdateDto
type APIDeployUpdateDto struct {

	// The API deployment message which can be used to communicate details relating to the status of the deployment.
	Message string `json:"message,omitempty"`

	// The status of API deployment.
	// Required: true
	// Enum: [DEPLOYED PENDING_DEPLOYMENT ERROR PENDING_UNDEPLOYMENT]
	Status *string `json:"status"`
}

// Validate validates this Api deploy update dto
func (m *APIDeployUpdateDto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var apiDeployUpdateDtoTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DEPLOYED","PENDING_DEPLOYMENT","ERROR","PENDING_UNDEPLOYMENT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		apiDeployUpdateDtoTypeStatusPropEnum = append(apiDeployUpdateDtoTypeStatusPropEnum, v)
	}
}

const (

	// APIDeployUpdateDtoStatusDEPLOYED captures enum value "DEPLOYED"
	APIDeployUpdateDtoStatusDEPLOYED string = "DEPLOYED"

	// APIDeployUpdateDtoStatusPENDINGDEPLOYMENT captures enum value "PENDING_DEPLOYMENT"
	APIDeployUpdateDtoStatusPENDINGDEPLOYMENT string = "PENDING_DEPLOYMENT"

	// APIDeployUpdateDtoStatusERROR captures enum value "ERROR"
	APIDeployUpdateDtoStatusERROR string = "ERROR"

	// APIDeployUpdateDtoStatusPENDINGUNDEPLOYMENT captures enum value "PENDING_UNDEPLOYMENT"
	APIDeployUpdateDtoStatusPENDINGUNDEPLOYMENT string = "PENDING_UNDEPLOYMENT"
)

// prop value enum
func (m *APIDeployUpdateDto) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, apiDeployUpdateDtoTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *APIDeployUpdateDto) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this Api deploy update dto based on context it is used
func (m *APIDeployUpdateDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *APIDeployUpdateDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIDeployUpdateDto) UnmarshalBinary(b []byte) error {
	var res APIDeployUpdateDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
