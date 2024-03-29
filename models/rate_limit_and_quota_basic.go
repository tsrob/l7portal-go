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

// RateLimitAndQuotaBasic rate limit and quota basic
//
// swagger:model RateLimitAndQuotaBasic
type RateLimitAndQuotaBasic struct {

	// Indicates whether the Rate Limit and Quota is applicable at ORGANIZATION or API level.
	// Enum: [ORGANIZATION API API_ORGANIZATION]
	AssignmentLevel string `json:"assignmentLevel,omitempty"`

	// Denotes if this Rate Limit and Quota is the default Rate Limit and Quota.
	DefaultRateQuota bool `json:"defaultRateQuota,omitempty"`

	// Description of the Rate Limit and Quota.
	// Max Length: 65000
	Description string `json:"description,omitempty"`

	// Name of the Rate Limit and Quota.
	// Max Length: 255
	Name string `json:"name,omitempty"`

	// quota
	Quota *RateLimitAndQuotaBasicQuota `json:"quota,omitempty"`

	// rate limit
	RateLimit *RateLimitAndQuotaBasicRateLimit `json:"rateLimit,omitempty"`

	// Unique key for this entity which conforms to the UUID standard according to RFC 4122.
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this rate limit and quota basic
func (m *RateLimitAndQuotaBasic) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAssignmentLevel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuota(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRateLimit(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var rateLimitAndQuotaBasicTypeAssignmentLevelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ORGANIZATION","API","API_ORGANIZATION"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		rateLimitAndQuotaBasicTypeAssignmentLevelPropEnum = append(rateLimitAndQuotaBasicTypeAssignmentLevelPropEnum, v)
	}
}

const (

	// RateLimitAndQuotaBasicAssignmentLevelORGANIZATION captures enum value "ORGANIZATION"
	RateLimitAndQuotaBasicAssignmentLevelORGANIZATION string = "ORGANIZATION"

	// RateLimitAndQuotaBasicAssignmentLevelAPI captures enum value "API"
	RateLimitAndQuotaBasicAssignmentLevelAPI string = "API"

	// RateLimitAndQuotaBasicAssignmentLevelAPIORGANIZATION captures enum value "API_ORGANIZATION"
	RateLimitAndQuotaBasicAssignmentLevelAPIORGANIZATION string = "API_ORGANIZATION"
)

// prop value enum
func (m *RateLimitAndQuotaBasic) validateAssignmentLevelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, rateLimitAndQuotaBasicTypeAssignmentLevelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RateLimitAndQuotaBasic) validateAssignmentLevel(formats strfmt.Registry) error {
	if swag.IsZero(m.AssignmentLevel) { // not required
		return nil
	}

	// value enum
	if err := m.validateAssignmentLevelEnum("assignmentLevel", "body", m.AssignmentLevel); err != nil {
		return err
	}

	return nil
}

func (m *RateLimitAndQuotaBasic) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 65000); err != nil {
		return err
	}

	return nil
}

func (m *RateLimitAndQuotaBasic) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MaxLength("name", "body", m.Name, 255); err != nil {
		return err
	}

	return nil
}

func (m *RateLimitAndQuotaBasic) validateQuota(formats strfmt.Registry) error {
	if swag.IsZero(m.Quota) { // not required
		return nil
	}

	if m.Quota != nil {
		if err := m.Quota.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("quota")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("quota")
			}
			return err
		}
	}

	return nil
}

func (m *RateLimitAndQuotaBasic) validateRateLimit(formats strfmt.Registry) error {
	if swag.IsZero(m.RateLimit) { // not required
		return nil
	}

	if m.RateLimit != nil {
		if err := m.RateLimit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rateLimit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("rateLimit")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this rate limit and quota basic based on the context it is used
func (m *RateLimitAndQuotaBasic) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateQuota(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRateLimit(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RateLimitAndQuotaBasic) contextValidateQuota(ctx context.Context, formats strfmt.Registry) error {

	if m.Quota != nil {
		if err := m.Quota.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("quota")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("quota")
			}
			return err
		}
	}

	return nil
}

func (m *RateLimitAndQuotaBasic) contextValidateRateLimit(ctx context.Context, formats strfmt.Registry) error {

	if m.RateLimit != nil {
		if err := m.RateLimit.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rateLimit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("rateLimit")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RateLimitAndQuotaBasic) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RateLimitAndQuotaBasic) UnmarshalBinary(b []byte) error {
	var res RateLimitAndQuotaBasic
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// RateLimitAndQuotaBasicQuota Provide Quota details
//
// swagger:model RateLimitAndQuotaBasicQuota
type RateLimitAndQuotaBasicQuota struct {

	// quota interval
	// Enum: [DAY HOUR MONTH]
	Interval string `json:"interval,omitempty"`

	// Maximum number of hits per defined interval
	Quota int64 `json:"quota,omitempty"`
}

// Validate validates this rate limit and quota basic quota
func (m *RateLimitAndQuotaBasicQuota) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInterval(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var rateLimitAndQuotaBasicQuotaTypeIntervalPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DAY","HOUR","MONTH"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		rateLimitAndQuotaBasicQuotaTypeIntervalPropEnum = append(rateLimitAndQuotaBasicQuotaTypeIntervalPropEnum, v)
	}
}

const (

	// RateLimitAndQuotaBasicQuotaIntervalDAY captures enum value "DAY"
	RateLimitAndQuotaBasicQuotaIntervalDAY string = "DAY"

	// RateLimitAndQuotaBasicQuotaIntervalHOUR captures enum value "HOUR"
	RateLimitAndQuotaBasicQuotaIntervalHOUR string = "HOUR"

	// RateLimitAndQuotaBasicQuotaIntervalMONTH captures enum value "MONTH"
	RateLimitAndQuotaBasicQuotaIntervalMONTH string = "MONTH"
)

// prop value enum
func (m *RateLimitAndQuotaBasicQuota) validateIntervalEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, rateLimitAndQuotaBasicQuotaTypeIntervalPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RateLimitAndQuotaBasicQuota) validateInterval(formats strfmt.Registry) error {
	if swag.IsZero(m.Interval) { // not required
		return nil
	}

	// value enum
	if err := m.validateIntervalEnum("quota"+"."+"interval", "body", m.Interval); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this rate limit and quota basic quota based on context it is used
func (m *RateLimitAndQuotaBasicQuota) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RateLimitAndQuotaBasicQuota) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RateLimitAndQuotaBasicQuota) UnmarshalBinary(b []byte) error {
	var res RateLimitAndQuotaBasicQuota
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// RateLimitAndQuotaBasicRateLimit Provide Rate Limit details
//
// swagger:model RateLimitAndQuotaBasicRateLimit
type RateLimitAndQuotaBasicRateLimit struct {

	// Maximum number of concurrent requests
	MaxConcurrency int64 `json:"maxConcurrency,omitempty"`

	// Maximum number of requests per second
	MaxRequestsPerSecond int64 `json:"maxRequestsPerSecond,omitempty"`

	// Spread over the window size in seconds
	WindowSizeInSeconds int64 `json:"windowSizeInSeconds,omitempty"`
}

// Validate validates this rate limit and quota basic rate limit
func (m *RateLimitAndQuotaBasicRateLimit) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this rate limit and quota basic rate limit based on context it is used
func (m *RateLimitAndQuotaBasicRateLimit) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RateLimitAndQuotaBasicRateLimit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RateLimitAndQuotaBasicRateLimit) UnmarshalBinary(b []byte) error {
	var res RateLimitAndQuotaBasicRateLimit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
