// Code generated by go-swagger; DO NOT EDIT.

package applications

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

// NewUpdateAPIKeyUsingPUT1Params creates a new UpdateAPIKeyUsingPUT1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateAPIKeyUsingPUT1Params() *UpdateAPIKeyUsingPUT1Params {
	return &UpdateAPIKeyUsingPUT1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateAPIKeyUsingPUT1ParamsWithTimeout creates a new UpdateAPIKeyUsingPUT1Params object
// with the ability to set a timeout on a request.
func NewUpdateAPIKeyUsingPUT1ParamsWithTimeout(timeout time.Duration) *UpdateAPIKeyUsingPUT1Params {
	return &UpdateAPIKeyUsingPUT1Params{
		timeout: timeout,
	}
}

// NewUpdateAPIKeyUsingPUT1ParamsWithContext creates a new UpdateAPIKeyUsingPUT1Params object
// with the ability to set a context for a request.
func NewUpdateAPIKeyUsingPUT1ParamsWithContext(ctx context.Context) *UpdateAPIKeyUsingPUT1Params {
	return &UpdateAPIKeyUsingPUT1Params{
		Context: ctx,
	}
}

// NewUpdateAPIKeyUsingPUT1ParamsWithHTTPClient creates a new UpdateAPIKeyUsingPUT1Params object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateAPIKeyUsingPUT1ParamsWithHTTPClient(client *http.Client) *UpdateAPIKeyUsingPUT1Params {
	return &UpdateAPIKeyUsingPUT1Params{
		HTTPClient: client,
	}
}

/*
UpdateAPIKeyUsingPUT1Params contains all the parameters to send to the API endpoint

	for the update Api key using p u t 1 operation.

	Typically these are written to a http.Request.
*/
type UpdateAPIKeyUsingPUT1Params struct {

	/* APIKey.

	   The API key to update.
	*/
	APIKey string

	/* APIKeyDto.

	   The API key to update
	*/
	APIKeyDto *models.APIKeyDto

	/* AppUUID.

	   The Application UUID the API key is in.
	*/
	AppUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update Api key using p u t 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateAPIKeyUsingPUT1Params) WithDefaults() *UpdateAPIKeyUsingPUT1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update Api key using p u t 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateAPIKeyUsingPUT1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update Api key using p u t 1 params
func (o *UpdateAPIKeyUsingPUT1Params) WithTimeout(timeout time.Duration) *UpdateAPIKeyUsingPUT1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update Api key using p u t 1 params
func (o *UpdateAPIKeyUsingPUT1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update Api key using p u t 1 params
func (o *UpdateAPIKeyUsingPUT1Params) WithContext(ctx context.Context) *UpdateAPIKeyUsingPUT1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update Api key using p u t 1 params
func (o *UpdateAPIKeyUsingPUT1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update Api key using p u t 1 params
func (o *UpdateAPIKeyUsingPUT1Params) WithHTTPClient(client *http.Client) *UpdateAPIKeyUsingPUT1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update Api key using p u t 1 params
func (o *UpdateAPIKeyUsingPUT1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIKey adds the aPIKey to the update Api key using p u t 1 params
func (o *UpdateAPIKeyUsingPUT1Params) WithAPIKey(aPIKey string) *UpdateAPIKeyUsingPUT1Params {
	o.SetAPIKey(aPIKey)
	return o
}

// SetAPIKey adds the apiKey to the update Api key using p u t 1 params
func (o *UpdateAPIKeyUsingPUT1Params) SetAPIKey(aPIKey string) {
	o.APIKey = aPIKey
}

// WithAPIKeyDto adds the aPIKeyDto to the update Api key using p u t 1 params
func (o *UpdateAPIKeyUsingPUT1Params) WithAPIKeyDto(aPIKeyDto *models.APIKeyDto) *UpdateAPIKeyUsingPUT1Params {
	o.SetAPIKeyDto(aPIKeyDto)
	return o
}

// SetAPIKeyDto adds the apiKeyDto to the update Api key using p u t 1 params
func (o *UpdateAPIKeyUsingPUT1Params) SetAPIKeyDto(aPIKeyDto *models.APIKeyDto) {
	o.APIKeyDto = aPIKeyDto
}

// WithAppUUID adds the appUUID to the update Api key using p u t 1 params
func (o *UpdateAPIKeyUsingPUT1Params) WithAppUUID(appUUID string) *UpdateAPIKeyUsingPUT1Params {
	o.SetAppUUID(appUUID)
	return o
}

// SetAppUUID adds the appUuid to the update Api key using p u t 1 params
func (o *UpdateAPIKeyUsingPUT1Params) SetAppUUID(appUUID string) {
	o.AppUUID = appUUID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateAPIKeyUsingPUT1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param apiKey
	if err := r.SetPathParam("apiKey", o.APIKey); err != nil {
		return err
	}
	if o.APIKeyDto != nil {
		if err := r.SetBodyParam(o.APIKeyDto); err != nil {
			return err
		}
	}

	// path param appUuid
	if err := r.SetPathParam("appUuid", o.AppUUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
