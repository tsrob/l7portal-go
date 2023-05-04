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

// Document document
//
// swagger:model Document
type Document struct {

	// The locale to which this document belongs.
	// Required: true
	// Max Length: 36
	Locale *string `json:"locale"`

	// The markdown content of the document.
	// Required: true
	Markdown *string `json:"markdown"`

	// This is a read-only field and is auto generated by the system. This is used to ensure optimistic locking.
	ModifyTs int64 `json:"modifyTs,omitempty"`

	// The URI of the document. This can be used by users to bookmark the document. Navtitle can only have alphanumeric characters, - and _.
	// Required: true
	// Max Length: 255
	Navtitle *string `json:"navtitle"`

	// The order in which the document appears on the tree.
	// Required: true
	Ordinal *int64 `json:"ordinal"`

	// The unique key associated with this document’s parent. For top-level documents, The unique key is an empty string.
	// Required: true
	// Max Length: 36
	ParentUUID *string `json:"parentUuid"`

	// Indicates the current state of the document. For example PUBLISHED.
	// Required: true
	// Enum: [PUBLISHED UPDATE_IN_PROGRESS]
	Status *string `json:"status"`

	// The title of the document.
	// Required: true
	// Max Length: 255
	Title *string `json:"title"`

	// The entity type associated with the document. For example, api, application, home or custom.
	// Required: true
	// Max Length: 25
	Type *string `json:"type"`

	// The entity type unique key associated with the document.
	// Required: true
	// Max Length: 36
	TypeUUID *string `json:"typeUuid"`

	// The unique key for this document which conforms to the UUID standard according to RFC 4122.
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this document
func (m *Document) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLocale(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMarkdown(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNavtitle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrdinal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParentUUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTypeUUID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Document) validateLocale(formats strfmt.Registry) error {

	if err := validate.Required("locale", "body", m.Locale); err != nil {
		return err
	}

	if err := validate.MaxLength("locale", "body", *m.Locale, 36); err != nil {
		return err
	}

	return nil
}

func (m *Document) validateMarkdown(formats strfmt.Registry) error {

	if err := validate.Required("markdown", "body", m.Markdown); err != nil {
		return err
	}

	return nil
}

func (m *Document) validateNavtitle(formats strfmt.Registry) error {

	if err := validate.Required("navtitle", "body", m.Navtitle); err != nil {
		return err
	}

	if err := validate.MaxLength("navtitle", "body", *m.Navtitle, 255); err != nil {
		return err
	}

	return nil
}

func (m *Document) validateOrdinal(formats strfmt.Registry) error {

	if err := validate.Required("ordinal", "body", m.Ordinal); err != nil {
		return err
	}

	return nil
}

func (m *Document) validateParentUUID(formats strfmt.Registry) error {

	if err := validate.Required("parentUuid", "body", m.ParentUUID); err != nil {
		return err
	}

	if err := validate.MaxLength("parentUuid", "body", *m.ParentUUID, 36); err != nil {
		return err
	}

	return nil
}

var documentTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PUBLISHED","UPDATE_IN_PROGRESS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		documentTypeStatusPropEnum = append(documentTypeStatusPropEnum, v)
	}
}

const (

	// DocumentStatusPUBLISHED captures enum value "PUBLISHED"
	DocumentStatusPUBLISHED string = "PUBLISHED"

	// DocumentStatusUPDATEINPROGRESS captures enum value "UPDATE_IN_PROGRESS"
	DocumentStatusUPDATEINPROGRESS string = "UPDATE_IN_PROGRESS"
)

// prop value enum
func (m *Document) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, documentTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Document) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

func (m *Document) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	if err := validate.MaxLength("title", "body", *m.Title, 255); err != nil {
		return err
	}

	return nil
}

func (m *Document) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	if err := validate.MaxLength("type", "body", *m.Type, 25); err != nil {
		return err
	}

	return nil
}

func (m *Document) validateTypeUUID(formats strfmt.Registry) error {

	if err := validate.Required("typeUuid", "body", m.TypeUUID); err != nil {
		return err
	}

	if err := validate.MaxLength("typeUuid", "body", *m.TypeUUID, 36); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this document based on context it is used
func (m *Document) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Document) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Document) UnmarshalBinary(b []byte) error {
	var res Document
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
