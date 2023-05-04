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

// ThemeBasic A representation of a Theme.
//
// swagger:model ThemeBasic
type ThemeBasic struct {

	// color
	// Required: true
	Color *ThemeBasicColor `json:"color"`

	// display
	// Required: true
	Display *ThemeBasicDisplay `json:"display"`

	// Base64 string of portal favicon.
	// Required: true
	// Max Length: 716800
	Favicon *string `json:"favicon"`

	// font size
	// Required: true
	FontSize *ThemeBasicFontSize `json:"fontSize"`

	// label
	Label *ThemeBasicLabel `json:"label,omitempty"`

	// Base64 string of portal logo.
	// Required: true
	// Max Length: 716800
	Logo *string `json:"logo"`

	// site title
	// Required: true
	SiteTitle *string `json:"siteTitle"`

	// typography
	// Required: true
	Typography *ThemeBasicTypography `json:"typography"`
}

// Validate validates this theme basic
func (m *ThemeBasic) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateColor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisplay(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFavicon(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFontSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSiteTitle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTypography(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ThemeBasic) validateColor(formats strfmt.Registry) error {

	if err := validate.Required("color", "body", m.Color); err != nil {
		return err
	}

	if m.Color != nil {
		if err := m.Color.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("color")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("color")
			}
			return err
		}
	}

	return nil
}

func (m *ThemeBasic) validateDisplay(formats strfmt.Registry) error {

	if err := validate.Required("display", "body", m.Display); err != nil {
		return err
	}

	if m.Display != nil {
		if err := m.Display.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("display")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("display")
			}
			return err
		}
	}

	return nil
}

func (m *ThemeBasic) validateFavicon(formats strfmt.Registry) error {

	if err := validate.Required("favicon", "body", m.Favicon); err != nil {
		return err
	}

	if err := validate.MaxLength("favicon", "body", *m.Favicon, 716800); err != nil {
		return err
	}

	return nil
}

func (m *ThemeBasic) validateFontSize(formats strfmt.Registry) error {

	if err := validate.Required("fontSize", "body", m.FontSize); err != nil {
		return err
	}

	if m.FontSize != nil {
		if err := m.FontSize.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fontSize")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fontSize")
			}
			return err
		}
	}

	return nil
}

func (m *ThemeBasic) validateLabel(formats strfmt.Registry) error {
	if swag.IsZero(m.Label) { // not required
		return nil
	}

	if m.Label != nil {
		if err := m.Label.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("label")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("label")
			}
			return err
		}
	}

	return nil
}

func (m *ThemeBasic) validateLogo(formats strfmt.Registry) error {

	if err := validate.Required("logo", "body", m.Logo); err != nil {
		return err
	}

	if err := validate.MaxLength("logo", "body", *m.Logo, 716800); err != nil {
		return err
	}

	return nil
}

func (m *ThemeBasic) validateSiteTitle(formats strfmt.Registry) error {

	if err := validate.Required("siteTitle", "body", m.SiteTitle); err != nil {
		return err
	}

	return nil
}

func (m *ThemeBasic) validateTypography(formats strfmt.Registry) error {

	if err := validate.Required("typography", "body", m.Typography); err != nil {
		return err
	}

	if m.Typography != nil {
		if err := m.Typography.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("typography")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("typography")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this theme basic based on the context it is used
func (m *ThemeBasic) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateColor(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDisplay(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFontSize(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLabel(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTypography(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ThemeBasic) contextValidateColor(ctx context.Context, formats strfmt.Registry) error {

	if m.Color != nil {
		if err := m.Color.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("color")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("color")
			}
			return err
		}
	}

	return nil
}

func (m *ThemeBasic) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if m.Display != nil {
		if err := m.Display.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("display")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("display")
			}
			return err
		}
	}

	return nil
}

func (m *ThemeBasic) contextValidateFontSize(ctx context.Context, formats strfmt.Registry) error {

	if m.FontSize != nil {
		if err := m.FontSize.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fontSize")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fontSize")
			}
			return err
		}
	}

	return nil
}

func (m *ThemeBasic) contextValidateLabel(ctx context.Context, formats strfmt.Registry) error {

	if m.Label != nil {
		if err := m.Label.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("label")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("label")
			}
			return err
		}
	}

	return nil
}

func (m *ThemeBasic) contextValidateTypography(ctx context.Context, formats strfmt.Registry) error {

	if m.Typography != nil {
		if err := m.Typography.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("typography")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("typography")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ThemeBasic) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ThemeBasic) UnmarshalBinary(b []byte) error {
	var res ThemeBasic
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ThemeBasicColor theme basic color
//
// swagger:model ThemeBasicColor
type ThemeBasicColor struct {

	// background
	// Required: true
	Background *string `json:"background"`

	// body text
	// Required: true
	BodyText *string `json:"bodyText"`

	// border color
	// Required: true
	BorderColor *string `json:"borderColor"`

	// footer background
	// Required: true
	FooterBackground *string `json:"footerBackground"`

	// footer link
	// Required: true
	FooterLink *string `json:"footerLink"`

	// footer text
	// Required: true
	FooterText *string `json:"footerText"`

	// header background
	// Required: true
	HeaderBackground *string `json:"headerBackground"`

	// header link
	// Required: true
	HeaderLink *string `json:"headerLink"`

	// link
	// Required: true
	Link *string `json:"link"`

	// link hover
	// Required: true
	LinkHover *string `json:"linkHover"`

	// page title
	// Required: true
	PageTitle *string `json:"pageTitle"`

	// primary button background
	// Required: true
	PrimaryButtonBackground *string `json:"primaryButtonBackground"`

	// primary button hover
	// Required: true
	PrimaryButtonHover *string `json:"primaryButtonHover"`

	// primary button text
	// Required: true
	PrimaryButtonText *string `json:"primaryButtonText"`

	// site title
	// Required: true
	SiteTitle *string `json:"siteTitle"`

	// small text
	// Required: true
	SmallText *string `json:"smallText"`
}

// Validate validates this theme basic color
func (m *ThemeBasicColor) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBackground(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBodyText(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBorderColor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFooterBackground(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFooterLink(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFooterText(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHeaderBackground(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHeaderLink(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLink(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinkHover(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePageTitle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrimaryButtonBackground(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrimaryButtonHover(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrimaryButtonText(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSiteTitle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSmallText(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ThemeBasicColor) validateBackground(formats strfmt.Registry) error {

	if err := validate.Required("color"+"."+"background", "body", m.Background); err != nil {
		return err
	}

	return nil
}

func (m *ThemeBasicColor) validateBodyText(formats strfmt.Registry) error {

	if err := validate.Required("color"+"."+"bodyText", "body", m.BodyText); err != nil {
		return err
	}

	return nil
}

func (m *ThemeBasicColor) validateBorderColor(formats strfmt.Registry) error {

	if err := validate.Required("color"+"."+"borderColor", "body", m.BorderColor); err != nil {
		return err
	}

	return nil
}

func (m *ThemeBasicColor) validateFooterBackground(formats strfmt.Registry) error {

	if err := validate.Required("color"+"."+"footerBackground", "body", m.FooterBackground); err != nil {
		return err
	}

	return nil
}

func (m *ThemeBasicColor) validateFooterLink(formats strfmt.Registry) error {

	if err := validate.Required("color"+"."+"footerLink", "body", m.FooterLink); err != nil {
		return err
	}

	return nil
}

func (m *ThemeBasicColor) validateFooterText(formats strfmt.Registry) error {

	if err := validate.Required("color"+"."+"footerText", "body", m.FooterText); err != nil {
		return err
	}

	return nil
}

func (m *ThemeBasicColor) validateHeaderBackground(formats strfmt.Registry) error {

	if err := validate.Required("color"+"."+"headerBackground", "body", m.HeaderBackground); err != nil {
		return err
	}

	return nil
}

func (m *ThemeBasicColor) validateHeaderLink(formats strfmt.Registry) error {

	if err := validate.Required("color"+"."+"headerLink", "body", m.HeaderLink); err != nil {
		return err
	}

	return nil
}

func (m *ThemeBasicColor) validateLink(formats strfmt.Registry) error {

	if err := validate.Required("color"+"."+"link", "body", m.Link); err != nil {
		return err
	}

	return nil
}

func (m *ThemeBasicColor) validateLinkHover(formats strfmt.Registry) error {

	if err := validate.Required("color"+"."+"linkHover", "body", m.LinkHover); err != nil {
		return err
	}

	return nil
}

func (m *ThemeBasicColor) validatePageTitle(formats strfmt.Registry) error {

	if err := validate.Required("color"+"."+"pageTitle", "body", m.PageTitle); err != nil {
		return err
	}

	return nil
}

func (m *ThemeBasicColor) validatePrimaryButtonBackground(formats strfmt.Registry) error {

	if err := validate.Required("color"+"."+"primaryButtonBackground", "body", m.PrimaryButtonBackground); err != nil {
		return err
	}

	return nil
}

func (m *ThemeBasicColor) validatePrimaryButtonHover(formats strfmt.Registry) error {

	if err := validate.Required("color"+"."+"primaryButtonHover", "body", m.PrimaryButtonHover); err != nil {
		return err
	}

	return nil
}

func (m *ThemeBasicColor) validatePrimaryButtonText(formats strfmt.Registry) error {

	if err := validate.Required("color"+"."+"primaryButtonText", "body", m.PrimaryButtonText); err != nil {
		return err
	}

	return nil
}

func (m *ThemeBasicColor) validateSiteTitle(formats strfmt.Registry) error {

	if err := validate.Required("color"+"."+"siteTitle", "body", m.SiteTitle); err != nil {
		return err
	}

	return nil
}

func (m *ThemeBasicColor) validateSmallText(formats strfmt.Registry) error {

	if err := validate.Required("color"+"."+"smallText", "body", m.SmallText); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this theme basic color based on context it is used
func (m *ThemeBasicColor) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ThemeBasicColor) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ThemeBasicColor) UnmarshalBinary(b []byte) error {
	var res ThemeBasicColor
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ThemeBasicDisplay theme basic display
//
// swagger:model ThemeBasicDisplay
type ThemeBasicDisplay struct {

	// cookie consent
	CookieConsent *ThemeBasicDisplayCookieConsent `json:"cookieConsent,omitempty"`

	// copyright
	Copyright bool `json:"copyright,omitempty"`

	// version number
	VersionNumber bool `json:"versionNumber,omitempty"`
}

// Validate validates this theme basic display
func (m *ThemeBasicDisplay) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCookieConsent(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ThemeBasicDisplay) validateCookieConsent(formats strfmt.Registry) error {
	if swag.IsZero(m.CookieConsent) { // not required
		return nil
	}

	if m.CookieConsent != nil {
		if err := m.CookieConsent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("display" + "." + "cookieConsent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("display" + "." + "cookieConsent")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this theme basic display based on the context it is used
func (m *ThemeBasicDisplay) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCookieConsent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ThemeBasicDisplay) contextValidateCookieConsent(ctx context.Context, formats strfmt.Registry) error {

	if m.CookieConsent != nil {
		if err := m.CookieConsent.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("display" + "." + "cookieConsent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("display" + "." + "cookieConsent")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ThemeBasicDisplay) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ThemeBasicDisplay) UnmarshalBinary(b []byte) error {
	var res ThemeBasicDisplay
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ThemeBasicDisplayCookieConsent theme basic display cookie consent
//
// swagger:model ThemeBasicDisplayCookieConsent
type ThemeBasicDisplayCookieConsent struct {

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// position
	Position string `json:"position,omitempty"`

	// theme
	Theme string `json:"theme,omitempty"`
}

// Validate validates this theme basic display cookie consent
func (m *ThemeBasicDisplayCookieConsent) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this theme basic display cookie consent based on context it is used
func (m *ThemeBasicDisplayCookieConsent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ThemeBasicDisplayCookieConsent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ThemeBasicDisplayCookieConsent) UnmarshalBinary(b []byte) error {
	var res ThemeBasicDisplayCookieConsent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ThemeBasicFontSize theme basic font size
//
// swagger:model ThemeBasicFontSize
type ThemeBasicFontSize struct {

	// body text
	// Required: true
	BodyText *string `json:"bodyText"`

	// button text
	// Required: true
	ButtonText *string `json:"buttonText"`

	// footer text
	// Required: true
	FooterText *string `json:"footerText"`

	// page title
	// Required: true
	PageTitle *string `json:"pageTitle"`

	// site title
	// Required: true
	SiteTitle *string `json:"siteTitle"`

	// small text
	// Required: true
	SmallText *string `json:"smallText"`
}

// Validate validates this theme basic font size
func (m *ThemeBasicFontSize) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBodyText(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateButtonText(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFooterText(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePageTitle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSiteTitle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSmallText(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ThemeBasicFontSize) validateBodyText(formats strfmt.Registry) error {

	if err := validate.Required("fontSize"+"."+"bodyText", "body", m.BodyText); err != nil {
		return err
	}

	return nil
}

func (m *ThemeBasicFontSize) validateButtonText(formats strfmt.Registry) error {

	if err := validate.Required("fontSize"+"."+"buttonText", "body", m.ButtonText); err != nil {
		return err
	}

	return nil
}

func (m *ThemeBasicFontSize) validateFooterText(formats strfmt.Registry) error {

	if err := validate.Required("fontSize"+"."+"footerText", "body", m.FooterText); err != nil {
		return err
	}

	return nil
}

func (m *ThemeBasicFontSize) validatePageTitle(formats strfmt.Registry) error {

	if err := validate.Required("fontSize"+"."+"pageTitle", "body", m.PageTitle); err != nil {
		return err
	}

	return nil
}

func (m *ThemeBasicFontSize) validateSiteTitle(formats strfmt.Registry) error {

	if err := validate.Required("fontSize"+"."+"siteTitle", "body", m.SiteTitle); err != nil {
		return err
	}

	return nil
}

func (m *ThemeBasicFontSize) validateSmallText(formats strfmt.Registry) error {

	if err := validate.Required("fontSize"+"."+"smallText", "body", m.SmallText); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this theme basic font size based on context it is used
func (m *ThemeBasicFontSize) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ThemeBasicFontSize) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ThemeBasicFontSize) UnmarshalBinary(b []byte) error {
	var res ThemeBasicFontSize
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ThemeBasicLabel Custom menu label.
//
// swagger:model ThemeBasicLabel
type ThemeBasicLabel struct {

	// service publish nav apicatalog title
	ServicePublishNavApicatalogTitle string `json:"service.publish.nav.apicatalog.title,omitempty"`

	// service publish nav apiexplorer title
	ServicePublishNavApiexplorerTitle string `json:"service.publish.nav.apiexplorer.title,omitempty"`

	// service publish nav apps title
	ServicePublishNavAppsTitle string `json:"service.publish.nav.apps.title,omitempty"`

	// service publish nav manageapi title
	ServicePublishNavManageapiTitle string `json:"service.publish.nav.manageapi.title,omitempty"`
}

// Validate validates this theme basic label
func (m *ThemeBasicLabel) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this theme basic label based on context it is used
func (m *ThemeBasicLabel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ThemeBasicLabel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ThemeBasicLabel) UnmarshalBinary(b []byte) error {
	var res ThemeBasicLabel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ThemeBasicTypography theme basic typography
//
// swagger:model ThemeBasicTypography
type ThemeBasicTypography struct {

	// body text
	// Required: true
	BodyText *string `json:"bodyText"`

	// button text
	// Required: true
	ButtonText *string `json:"buttonText"`

	// footer text
	// Required: true
	FooterText *string `json:"footerText"`

	// page title
	// Required: true
	PageTitle *string `json:"pageTitle"`

	// site title
	// Required: true
	SiteTitle *string `json:"siteTitle"`

	// small text
	// Required: true
	SmallText *string `json:"smallText"`
}

// Validate validates this theme basic typography
func (m *ThemeBasicTypography) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBodyText(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateButtonText(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFooterText(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePageTitle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSiteTitle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSmallText(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ThemeBasicTypography) validateBodyText(formats strfmt.Registry) error {

	if err := validate.Required("typography"+"."+"bodyText", "body", m.BodyText); err != nil {
		return err
	}

	return nil
}

func (m *ThemeBasicTypography) validateButtonText(formats strfmt.Registry) error {

	if err := validate.Required("typography"+"."+"buttonText", "body", m.ButtonText); err != nil {
		return err
	}

	return nil
}

func (m *ThemeBasicTypography) validateFooterText(formats strfmt.Registry) error {

	if err := validate.Required("typography"+"."+"footerText", "body", m.FooterText); err != nil {
		return err
	}

	return nil
}

func (m *ThemeBasicTypography) validatePageTitle(formats strfmt.Registry) error {

	if err := validate.Required("typography"+"."+"pageTitle", "body", m.PageTitle); err != nil {
		return err
	}

	return nil
}

func (m *ThemeBasicTypography) validateSiteTitle(formats strfmt.Registry) error {

	if err := validate.Required("typography"+"."+"siteTitle", "body", m.SiteTitle); err != nil {
		return err
	}

	return nil
}

func (m *ThemeBasicTypography) validateSmallText(formats strfmt.Registry) error {

	if err := validate.Required("typography"+"."+"smallText", "body", m.SmallText); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this theme basic typography based on context it is used
func (m *ThemeBasicTypography) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ThemeBasicTypography) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ThemeBasicTypography) UnmarshalBinary(b []byte) error {
	var res ThemeBasicTypography
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
