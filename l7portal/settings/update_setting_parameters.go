// Code generated by go-swagger; DO NOT EDIT.

package settings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/tsrob/l7portal-go/models"
)

// NewUpdateSettingParams creates a new UpdateSettingParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateSettingParams() *UpdateSettingParams {
	return &UpdateSettingParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateSettingParamsWithTimeout creates a new UpdateSettingParams object
// with the ability to set a timeout on a request.
func NewUpdateSettingParamsWithTimeout(timeout time.Duration) *UpdateSettingParams {
	return &UpdateSettingParams{
		timeout: timeout,
	}
}

// NewUpdateSettingParamsWithContext creates a new UpdateSettingParams object
// with the ability to set a context for a request.
func NewUpdateSettingParamsWithContext(ctx context.Context) *UpdateSettingParams {
	return &UpdateSettingParams{
		Context: ctx,
	}
}

// NewUpdateSettingParamsWithHTTPClient creates a new UpdateSettingParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateSettingParamsWithHTTPClient(client *http.Client) *UpdateSettingParams {
	return &UpdateSettingParams{
		HTTPClient: client,
	}
}

/*
UpdateSettingParams contains all the parameters to send to the API endpoint

	for the update setting operation.

	Typically these are written to a http.Request.
*/
type UpdateSettingParams struct {

	/* Body.

	   Provides the values needed to set up the Setting object.
	*/
	Body *models.Setting

	/* Input.

	   Name of the setting
	*/
	Input string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update setting params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateSettingParams) WithDefaults() *UpdateSettingParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update setting params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateSettingParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update setting params
func (o *UpdateSettingParams) WithTimeout(timeout time.Duration) *UpdateSettingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update setting params
func (o *UpdateSettingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update setting params
func (o *UpdateSettingParams) WithContext(ctx context.Context) *UpdateSettingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update setting params
func (o *UpdateSettingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update setting params
func (o *UpdateSettingParams) WithHTTPClient(client *http.Client) *UpdateSettingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update setting params
func (o *UpdateSettingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update setting params
func (o *UpdateSettingParams) WithBody(body *models.Setting) *UpdateSettingParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update setting params
func (o *UpdateSettingParams) SetBody(body *models.Setting) {
	o.Body = body
}

// WithInput adds the input to the update setting params
func (o *UpdateSettingParams) WithInput(input string) *UpdateSettingParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the update setting params
func (o *UpdateSettingParams) SetInput(input string) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateSettingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param input
	if err := r.SetPathParam("input", o.Input); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
