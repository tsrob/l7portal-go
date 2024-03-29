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

// NewDeleteAPIKeyUsingDELETEParams creates a new DeleteAPIKeyUsingDELETEParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteAPIKeyUsingDELETEParams() *DeleteAPIKeyUsingDELETEParams {
	return &DeleteAPIKeyUsingDELETEParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAPIKeyUsingDELETEParamsWithTimeout creates a new DeleteAPIKeyUsingDELETEParams object
// with the ability to set a timeout on a request.
func NewDeleteAPIKeyUsingDELETEParamsWithTimeout(timeout time.Duration) *DeleteAPIKeyUsingDELETEParams {
	return &DeleteAPIKeyUsingDELETEParams{
		timeout: timeout,
	}
}

// NewDeleteAPIKeyUsingDELETEParamsWithContext creates a new DeleteAPIKeyUsingDELETEParams object
// with the ability to set a context for a request.
func NewDeleteAPIKeyUsingDELETEParamsWithContext(ctx context.Context) *DeleteAPIKeyUsingDELETEParams {
	return &DeleteAPIKeyUsingDELETEParams{
		Context: ctx,
	}
}

// NewDeleteAPIKeyUsingDELETEParamsWithHTTPClient creates a new DeleteAPIKeyUsingDELETEParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteAPIKeyUsingDELETEParamsWithHTTPClient(client *http.Client) *DeleteAPIKeyUsingDELETEParams {
	return &DeleteAPIKeyUsingDELETEParams{
		HTTPClient: client,
	}
}

/*
DeleteAPIKeyUsingDELETEParams contains all the parameters to send to the API endpoint

	for the delete Api key using d e l e t e operation.

	Typically these are written to a http.Request.
*/
type DeleteAPIKeyUsingDELETEParams struct {

	/* APIKey.

	   The API key to delete.
	*/
	APIKey string

	/* AppUUID.

	   The Application UUID to delete the key from.
	*/
	AppUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete Api key using d e l e t e params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAPIKeyUsingDELETEParams) WithDefaults() *DeleteAPIKeyUsingDELETEParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete Api key using d e l e t e params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAPIKeyUsingDELETEParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete Api key using d e l e t e params
func (o *DeleteAPIKeyUsingDELETEParams) WithTimeout(timeout time.Duration) *DeleteAPIKeyUsingDELETEParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete Api key using d e l e t e params
func (o *DeleteAPIKeyUsingDELETEParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete Api key using d e l e t e params
func (o *DeleteAPIKeyUsingDELETEParams) WithContext(ctx context.Context) *DeleteAPIKeyUsingDELETEParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete Api key using d e l e t e params
func (o *DeleteAPIKeyUsingDELETEParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete Api key using d e l e t e params
func (o *DeleteAPIKeyUsingDELETEParams) WithHTTPClient(client *http.Client) *DeleteAPIKeyUsingDELETEParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete Api key using d e l e t e params
func (o *DeleteAPIKeyUsingDELETEParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIKey adds the aPIKey to the delete Api key using d e l e t e params
func (o *DeleteAPIKeyUsingDELETEParams) WithAPIKey(aPIKey string) *DeleteAPIKeyUsingDELETEParams {
	o.SetAPIKey(aPIKey)
	return o
}

// SetAPIKey adds the apiKey to the delete Api key using d e l e t e params
func (o *DeleteAPIKeyUsingDELETEParams) SetAPIKey(aPIKey string) {
	o.APIKey = aPIKey
}

// WithAppUUID adds the appUUID to the delete Api key using d e l e t e params
func (o *DeleteAPIKeyUsingDELETEParams) WithAppUUID(appUUID string) *DeleteAPIKeyUsingDELETEParams {
	o.SetAppUUID(appUUID)
	return o
}

// SetAppUUID adds the appUuid to the delete Api key using d e l e t e params
func (o *DeleteAPIKeyUsingDELETEParams) SetAppUUID(appUUID string) {
	o.AppUUID = appUUID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAPIKeyUsingDELETEParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
