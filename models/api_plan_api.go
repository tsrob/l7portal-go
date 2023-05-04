// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// APIPlanAPI Api plan Api
//
// swagger:model ApiPlanApi
type APIPlanAPI struct {

	// The uuid of the API Plan.
	// Required: true
	APIPlanUUID *string `json:"ApiPlanUuid"`

	// The uuid of the API to associate with the specified API Plan.
	// Required: true
	APIUUID *string `json:"ApiUuid"`
}

// Validate validates this Api plan Api
func (m *APIPlanAPI) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAPIPlanUUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAPIUUID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIPlanAPI) validateAPIPlanUUID(formats strfmt.Registry) error {

	if err := validate.Required("ApiPlanUuid", "body", m.APIPlanUUID); err != nil {
		return err
	}

	return nil
}

func (m *APIPlanAPI) validateAPIUUID(formats strfmt.Registry) error {

	if err := validate.Required("ApiUuid", "body", m.APIUUID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this Api plan Api based on context it is used
func (m *APIPlanAPI) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *APIPlanAPI) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIPlanAPI) UnmarshalBinary(b []byte) error {
	var res APIPlanAPI
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}