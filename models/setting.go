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

// Setting setting
// Example: {"Name":"GOOGLE_ANALYTICS_ID","Uuid":"UUID_OF_SETTING","Value":"VALUE_OF_SETTING"}
//
// swagger:model Setting
type Setting struct {

	// Setting Name.
	//
	//  Supported names at this time:
	//  GOOGLE_ANALYTICS_ID: google analytics account id;
	//  ACTIVATE_MSSO: enable/disable single signon. Values (true, false);
	//  APPLICATION_REQUEST_WORKFLOW: enable/disable new applications approval workflow,  Values (ENABLED, DISABLED);
	//  EDIT_APPLICATION_REQUEST_WORKFLOW: enable/disable existing applications approval workflow,  Values (ENABLED, DISABLED);
	//  DELETE_APPLICATION_REQUEST_WORKFLOW: delete existing applications approval workflow,  Values (ENABLED, DISABLED);
	//  APP_EDIT_WF_EXCLUDE_PROPS: workflow excluded fields, Sample ({"Fields":["Name","Description","Status"],"CustomFields":["MyCustomField","AnotherCustomField"]});
	// Required: true
	// Max Length: 80
	// Enum: [GOOGLE_ANALYTICS_ID ACTIVATE_MSSO EDIT_APPLICATION_REQUEST_WORKFLOW DELETE_APPLICATION_REQUEST_WORKFLOW APP_EDIT_WF_EXCLUDE_PROPS APPLICATION_REQUEST_WORKFLOW]
	Name *string `json:"Name"`

	// Unique key for this entity which conforms to the UUID standard according to RFC 4122.
	// Required: true
	UUID *string `json:"Uuid"`

	// Setting Value.
	// Required: true
	// Max Length: 65000
	Value *string `json:"Value"`
}

// Validate validates this setting
func (m *Setting) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var settingTypeNamePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["GOOGLE_ANALYTICS_ID","ACTIVATE_MSSO","EDIT_APPLICATION_REQUEST_WORKFLOW","DELETE_APPLICATION_REQUEST_WORKFLOW","APP_EDIT_WF_EXCLUDE_PROPS","APPLICATION_REQUEST_WORKFLOW"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		settingTypeNamePropEnum = append(settingTypeNamePropEnum, v)
	}
}

const (

	// SettingNameGOOGLEANALYTICSID captures enum value "GOOGLE_ANALYTICS_ID"
	SettingNameGOOGLEANALYTICSID string = "GOOGLE_ANALYTICS_ID"

	// SettingNameACTIVATEMSSO captures enum value "ACTIVATE_MSSO"
	SettingNameACTIVATEMSSO string = "ACTIVATE_MSSO"

	// SettingNameEDITAPPLICATIONREQUESTWORKFLOW captures enum value "EDIT_APPLICATION_REQUEST_WORKFLOW"
	SettingNameEDITAPPLICATIONREQUESTWORKFLOW string = "EDIT_APPLICATION_REQUEST_WORKFLOW"

	// SettingNameDELETEAPPLICATIONREQUESTWORKFLOW captures enum value "DELETE_APPLICATION_REQUEST_WORKFLOW"
	SettingNameDELETEAPPLICATIONREQUESTWORKFLOW string = "DELETE_APPLICATION_REQUEST_WORKFLOW"

	// SettingNameAPPEDITWFEXCLUDEPROPS captures enum value "APP_EDIT_WF_EXCLUDE_PROPS"
	SettingNameAPPEDITWFEXCLUDEPROPS string = "APP_EDIT_WF_EXCLUDE_PROPS"

	// SettingNameAPPLICATIONREQUESTWORKFLOW captures enum value "APPLICATION_REQUEST_WORKFLOW"
	SettingNameAPPLICATIONREQUESTWORKFLOW string = "APPLICATION_REQUEST_WORKFLOW"
)

// prop value enum
func (m *Setting) validateNameEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, settingTypeNamePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Setting) validateName(formats strfmt.Registry) error {

	if err := validate.Required("Name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MaxLength("Name", "body", *m.Name, 80); err != nil {
		return err
	}

	// value enum
	if err := m.validateNameEnum("Name", "body", *m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Setting) validateUUID(formats strfmt.Registry) error {

	if err := validate.Required("Uuid", "body", m.UUID); err != nil {
		return err
	}

	return nil
}

func (m *Setting) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("Value", "body", m.Value); err != nil {
		return err
	}

	if err := validate.MaxLength("Value", "body", *m.Value, 65000); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this setting based on context it is used
func (m *Setting) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Setting) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Setting) UnmarshalBinary(b []byte) error {
	var res Setting
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
