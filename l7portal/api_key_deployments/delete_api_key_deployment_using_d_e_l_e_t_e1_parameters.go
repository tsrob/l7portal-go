// Code generated by go-swagger; DO NOT EDIT.

package api_key_deployments

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

// NewDeleteAPIKeyDeploymentUsingDELETE1Params creates a new DeleteAPIKeyDeploymentUsingDELETE1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteAPIKeyDeploymentUsingDELETE1Params() *DeleteAPIKeyDeploymentUsingDELETE1Params {
	return &DeleteAPIKeyDeploymentUsingDELETE1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAPIKeyDeploymentUsingDELETE1ParamsWithTimeout creates a new DeleteAPIKeyDeploymentUsingDELETE1Params object
// with the ability to set a timeout on a request.
func NewDeleteAPIKeyDeploymentUsingDELETE1ParamsWithTimeout(timeout time.Duration) *DeleteAPIKeyDeploymentUsingDELETE1Params {
	return &DeleteAPIKeyDeploymentUsingDELETE1Params{
		timeout: timeout,
	}
}

// NewDeleteAPIKeyDeploymentUsingDELETE1ParamsWithContext creates a new DeleteAPIKeyDeploymentUsingDELETE1Params object
// with the ability to set a context for a request.
func NewDeleteAPIKeyDeploymentUsingDELETE1ParamsWithContext(ctx context.Context) *DeleteAPIKeyDeploymentUsingDELETE1Params {
	return &DeleteAPIKeyDeploymentUsingDELETE1Params{
		Context: ctx,
	}
}

// NewDeleteAPIKeyDeploymentUsingDELETE1ParamsWithHTTPClient creates a new DeleteAPIKeyDeploymentUsingDELETE1Params object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteAPIKeyDeploymentUsingDELETE1ParamsWithHTTPClient(client *http.Client) *DeleteAPIKeyDeploymentUsingDELETE1Params {
	return &DeleteAPIKeyDeploymentUsingDELETE1Params{
		HTTPClient: client,
	}
}

/*
DeleteAPIKeyDeploymentUsingDELETE1Params contains all the parameters to send to the API endpoint

	for the delete Api key deployment using d e l e t e 1 operation.

	Typically these are written to a http.Request.
*/
type DeleteAPIKeyDeploymentUsingDELETE1Params struct {

	/* APIKey.

	   The API Key.
	*/
	APIKey string

	/* ProxyUUID.

	   The proxy UUID.
	*/
	ProxyUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete Api key deployment using d e l e t e 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAPIKeyDeploymentUsingDELETE1Params) WithDefaults() *DeleteAPIKeyDeploymentUsingDELETE1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete Api key deployment using d e l e t e 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAPIKeyDeploymentUsingDELETE1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete Api key deployment using d e l e t e 1 params
func (o *DeleteAPIKeyDeploymentUsingDELETE1Params) WithTimeout(timeout time.Duration) *DeleteAPIKeyDeploymentUsingDELETE1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete Api key deployment using d e l e t e 1 params
func (o *DeleteAPIKeyDeploymentUsingDELETE1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete Api key deployment using d e l e t e 1 params
func (o *DeleteAPIKeyDeploymentUsingDELETE1Params) WithContext(ctx context.Context) *DeleteAPIKeyDeploymentUsingDELETE1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete Api key deployment using d e l e t e 1 params
func (o *DeleteAPIKeyDeploymentUsingDELETE1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete Api key deployment using d e l e t e 1 params
func (o *DeleteAPIKeyDeploymentUsingDELETE1Params) WithHTTPClient(client *http.Client) *DeleteAPIKeyDeploymentUsingDELETE1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete Api key deployment using d e l e t e 1 params
func (o *DeleteAPIKeyDeploymentUsingDELETE1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIKey adds the aPIKey to the delete Api key deployment using d e l e t e 1 params
func (o *DeleteAPIKeyDeploymentUsingDELETE1Params) WithAPIKey(aPIKey string) *DeleteAPIKeyDeploymentUsingDELETE1Params {
	o.SetAPIKey(aPIKey)
	return o
}

// SetAPIKey adds the apiKey to the delete Api key deployment using d e l e t e 1 params
func (o *DeleteAPIKeyDeploymentUsingDELETE1Params) SetAPIKey(aPIKey string) {
	o.APIKey = aPIKey
}

// WithProxyUUID adds the proxyUUID to the delete Api key deployment using d e l e t e 1 params
func (o *DeleteAPIKeyDeploymentUsingDELETE1Params) WithProxyUUID(proxyUUID string) *DeleteAPIKeyDeploymentUsingDELETE1Params {
	o.SetProxyUUID(proxyUUID)
	return o
}

// SetProxyUUID adds the proxyUuid to the delete Api key deployment using d e l e t e 1 params
func (o *DeleteAPIKeyDeploymentUsingDELETE1Params) SetProxyUUID(proxyUUID string) {
	o.ProxyUUID = proxyUUID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAPIKeyDeploymentUsingDELETE1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param apiKey
	if err := r.SetPathParam("apiKey", o.APIKey); err != nil {
		return err
	}

	// path param proxyUuid
	if err := r.SetPathParam("proxyUuid", o.ProxyUUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
