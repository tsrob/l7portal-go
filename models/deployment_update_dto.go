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

// DeploymentUpdateDto deployment update dto
// Example: {"message":"[{\"status\":\"DEPLOYED\",\"message\":\"\u003cl7:ApiPlans xmlns:l7=\\\"http://ns.l7tech.com/2012/04/api-management\\\"\u003e\\n\u003cl7:Method\u003ePUT\u003c/l7:Method\u003e\\n\u003cl7:Resource\u003e/1/api/plans\u003c/l7:Resource\u003e\\n\u003cl7:ResourceId/\u003e\\n\u003cl7:Timestamp\u003eTue Jun 29 17:34:30 PDT 2021\u003c/l7:Timestamp\u003e\\n\u003cl7:Status\u003e200\u003c/l7:Status\u003e\\n\u003cl7:Details\u003esuccess\u003c/l7:Details\u003e\\n\u003cl7:ApiPlan\u003e\\n\u003cl7:PlanId\u003e5b834a8e-a9b1-4fd4-9c3f-937facf8e222\u003c/l7:PlanId\u003e\\n\u003cl7:PlanName\u003eapiplan-1-test-me\u003c/l7:PlanName\u003e\\n\u003cl7:DefaultPlan\u003efalse\u003c/l7:DefaultPlan\u003e\\n\u003cl7:LastUpdate\u003e2021-06-29T17:34:30.875-07:00\u003c/l7:LastUpdate\u003e\\n\u003cl7:PlanPolicy/\u003e\\n\u003cl7:PlanDetails\u003e\\n\u003cl7:ThroughputQuota l7:enabled=\\\"false\\\"\u003e\\n\u003cl7:Quota\u003e1001\u003c/l7:Quota\u003e\\n\u003cl7:TimeUnit\u003e4\u003c/l7:TimeUnit\u003e\\n\u003cl7:CounterStrategy\u003e0\u003c/l7:CounterStrategy\u003e\\n\u003c/l7:ThroughputQuota\u003e\\n\u003cl7:RateLimit l7:enabled=\\\"false\\\"\u003e\\n\u003cl7:MaxRequestRate\u003e100\u003c/l7:MaxRequestRate\u003e\\n\u003cl7:WindowSizeInSeconds\u003e0\u003c/l7:WindowSizeInSeconds\u003e\\n\u003cl7:HardLimit\u003efalse\u003c/l7:HardLimit\u003e\\n\u003c/l7:RateLimit\u003e\\n\u003c/l7:PlanDetails\u003e\\n\u003c/l7:ApiPlan\u003e\\n\u003c/l7:ApiPlans\u003e\\n\",\"targetLocation\":\"https://localhost:8443/portalman/1/api/plans?removeOmitted=true\"}]","status":"DEPLOYED"}
//
// swagger:model DeploymentUpdateDto
type DeploymentUpdateDto struct {

	// The API deployment message which can be used to communicate details relating to the status of the deployment.
	Message string `json:"message,omitempty"`

	// The status of API deployment.
	// Required: true
	// Enum: [DEPLOYED PENDING_DEPLOYMENT ERROR PENDING_UNDEPLOYMENT]
	Status *string `json:"status"`
}

// Validate validates this deployment update dto
func (m *DeploymentUpdateDto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var deploymentUpdateDtoTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DEPLOYED","PENDING_DEPLOYMENT","ERROR","PENDING_UNDEPLOYMENT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		deploymentUpdateDtoTypeStatusPropEnum = append(deploymentUpdateDtoTypeStatusPropEnum, v)
	}
}

const (

	// DeploymentUpdateDtoStatusDEPLOYED captures enum value "DEPLOYED"
	DeploymentUpdateDtoStatusDEPLOYED string = "DEPLOYED"

	// DeploymentUpdateDtoStatusPENDINGDEPLOYMENT captures enum value "PENDING_DEPLOYMENT"
	DeploymentUpdateDtoStatusPENDINGDEPLOYMENT string = "PENDING_DEPLOYMENT"

	// DeploymentUpdateDtoStatusERROR captures enum value "ERROR"
	DeploymentUpdateDtoStatusERROR string = "ERROR"

	// DeploymentUpdateDtoStatusPENDINGUNDEPLOYMENT captures enum value "PENDING_UNDEPLOYMENT"
	DeploymentUpdateDtoStatusPENDINGUNDEPLOYMENT string = "PENDING_UNDEPLOYMENT"
)

// prop value enum
func (m *DeploymentUpdateDto) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, deploymentUpdateDtoTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DeploymentUpdateDto) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this deployment update dto based on context it is used
func (m *DeploymentUpdateDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeploymentUpdateDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeploymentUpdateDto) UnmarshalBinary(b []byte) error {
	var res DeploymentUpdateDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
