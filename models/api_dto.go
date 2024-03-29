// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// APIDto ApiDto
// Example: {"accessStatus":"PUBLIC","apiEulaUuid":"c9406345-eb76-11e3-b0cd-000nosaj86a8","apiServiceType":"REST","authenticationParameters":"{\"ApiKeyType\":\"query\"}","authenticationType":"APIKEY","description":"sample public description","locationUrl":"http://petstore.swagger.io/v2","managingOrgUuid":"172594b6-18ba-4b0c-8d61-807db457e81d","name":"Swagger Petstore","portalStatus":"ENABLED","privateDescription":"sample private description","ssgUrl":"petstore-proxy","version":1}
//
// swagger:model ApiDto
type APIDto struct {

	// access status
	AccessStatus string `json:"accessStatus,omitempty"`

	// api eula Uuid
	APIEulaUUID string `json:"apiEulaUuid,omitempty"`

	// api service type
	APIServiceType string `json:"apiServiceType,omitempty"`

	// authentication parameters
	AuthenticationParameters string `json:"authenticationParameters,omitempty"`

	// authentication type
	AuthenticationType string `json:"authenticationType,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// links
	Links []*Link `json:"links"`

	// location Url
	LocationURL string `json:"locationUrl,omitempty"`

	// managing org Uuid
	ManagingOrgUUID string `json:"managingOrgUuid,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// portal status
	PortalStatus string `json:"portalStatus,omitempty"`

	// private description
	PrivateDescription string `json:"privateDescription,omitempty"`

	// published by portal
	PublishedByPortal bool `json:"publishedByPortal,omitempty"`

	// restricted
	Restricted bool `json:"restricted,omitempty"`

	// spec filename
	SpecFilename string `json:"specFilename,omitempty"`

	// spec filesize
	SpecFilesize int64 `json:"specFilesize,omitempty"`

	// ssg Url
	SsgURL string `json:"ssgUrl,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// uuid
	UUID string `json:"uuid,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this Api dto
func (m *APIDto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIDto) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	for i := 0; i < len(m.Links); i++ {
		if swag.IsZero(m.Links[i]) { // not required
			continue
		}

		if m.Links[i] != nil {
			if err := m.Links[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("links" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("links" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this Api dto based on the context it is used
func (m *APIDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIDto) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Links); i++ {

		if m.Links[i] != nil {
			if err := m.Links[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("links" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("links" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *APIDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIDto) UnmarshalBinary(b []byte) error {
	var res APIDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
