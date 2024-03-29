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
)

// NewGetAPIKeyUsingGETParams creates a new GetAPIKeyUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAPIKeyUsingGETParams() *GetAPIKeyUsingGETParams {
	return &GetAPIKeyUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPIKeyUsingGETParamsWithTimeout creates a new GetAPIKeyUsingGETParams object
// with the ability to set a timeout on a request.
func NewGetAPIKeyUsingGETParamsWithTimeout(timeout time.Duration) *GetAPIKeyUsingGETParams {
	return &GetAPIKeyUsingGETParams{
		timeout: timeout,
	}
}

// NewGetAPIKeyUsingGETParamsWithContext creates a new GetAPIKeyUsingGETParams object
// with the ability to set a context for a request.
func NewGetAPIKeyUsingGETParamsWithContext(ctx context.Context) *GetAPIKeyUsingGETParams {
	return &GetAPIKeyUsingGETParams{
		Context: ctx,
	}
}

// NewGetAPIKeyUsingGETParamsWithHTTPClient creates a new GetAPIKeyUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAPIKeyUsingGETParamsWithHTTPClient(client *http.Client) *GetAPIKeyUsingGETParams {
	return &GetAPIKeyUsingGETParams{
		HTTPClient: client,
	}
}

/*
GetAPIKeyUsingGETParams contains all the parameters to send to the API endpoint

	for the get Api key using g e t operation.

	Typically these are written to a http.Request.
*/
type GetAPIKeyUsingGETParams struct {

	/* APIKey.

	   The API key to get.
	*/
	APIKey string

	/* AppUUID.

	   The Application UUID to get the API key from.
	*/
	AppUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get Api key using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIKeyUsingGETParams) WithDefaults() *GetAPIKeyUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get Api key using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIKeyUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get Api key using g e t params
func (o *GetAPIKeyUsingGETParams) WithTimeout(timeout time.Duration) *GetAPIKeyUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get Api key using g e t params
func (o *GetAPIKeyUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get Api key using g e t params
func (o *GetAPIKeyUsingGETParams) WithContext(ctx context.Context) *GetAPIKeyUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get Api key using g e t params
func (o *GetAPIKeyUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get Api key using g e t params
func (o *GetAPIKeyUsingGETParams) WithHTTPClient(client *http.Client) *GetAPIKeyUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get Api key using g e t params
func (o *GetAPIKeyUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIKey adds the aPIKey to the get Api key using g e t params
func (o *GetAPIKeyUsingGETParams) WithAPIKey(aPIKey string) *GetAPIKeyUsingGETParams {
	o.SetAPIKey(aPIKey)
	return o
}

// SetAPIKey adds the apiKey to the get Api key using g e t params
func (o *GetAPIKeyUsingGETParams) SetAPIKey(aPIKey string) {
	o.APIKey = aPIKey
}

// WithAppUUID adds the appUUID to the get Api key using g e t params
func (o *GetAPIKeyUsingGETParams) WithAppUUID(appUUID string) *GetAPIKeyUsingGETParams {
	o.SetAppUUID(appUUID)
	return o
}

// SetAppUUID adds the appUuid to the get Api key using g e t params
func (o *GetAPIKeyUsingGETParams) SetAppUUID(appUUID string) {
	o.AppUUID = appUUID
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPIKeyUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param apiKey
	if err := r.SetPathParam("apiKey", o.APIKey); err != nil {
		return err
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
