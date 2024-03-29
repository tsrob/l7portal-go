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

// NewGetAPIKeyDeploymentUsingGET1Params creates a new GetAPIKeyDeploymentUsingGET1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAPIKeyDeploymentUsingGET1Params() *GetAPIKeyDeploymentUsingGET1Params {
	return &GetAPIKeyDeploymentUsingGET1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPIKeyDeploymentUsingGET1ParamsWithTimeout creates a new GetAPIKeyDeploymentUsingGET1Params object
// with the ability to set a timeout on a request.
func NewGetAPIKeyDeploymentUsingGET1ParamsWithTimeout(timeout time.Duration) *GetAPIKeyDeploymentUsingGET1Params {
	return &GetAPIKeyDeploymentUsingGET1Params{
		timeout: timeout,
	}
}

// NewGetAPIKeyDeploymentUsingGET1ParamsWithContext creates a new GetAPIKeyDeploymentUsingGET1Params object
// with the ability to set a context for a request.
func NewGetAPIKeyDeploymentUsingGET1ParamsWithContext(ctx context.Context) *GetAPIKeyDeploymentUsingGET1Params {
	return &GetAPIKeyDeploymentUsingGET1Params{
		Context: ctx,
	}
}

// NewGetAPIKeyDeploymentUsingGET1ParamsWithHTTPClient creates a new GetAPIKeyDeploymentUsingGET1Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetAPIKeyDeploymentUsingGET1ParamsWithHTTPClient(client *http.Client) *GetAPIKeyDeploymentUsingGET1Params {
	return &GetAPIKeyDeploymentUsingGET1Params{
		HTTPClient: client,
	}
}

/*
GetAPIKeyDeploymentUsingGET1Params contains all the parameters to send to the API endpoint

	for the get Api key deployment using g e t 1 operation.

	Typically these are written to a http.Request.
*/
type GetAPIKeyDeploymentUsingGET1Params struct {

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

// WithDefaults hydrates default values in the get Api key deployment using g e t 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIKeyDeploymentUsingGET1Params) WithDefaults() *GetAPIKeyDeploymentUsingGET1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get Api key deployment using g e t 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIKeyDeploymentUsingGET1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get Api key deployment using g e t 1 params
func (o *GetAPIKeyDeploymentUsingGET1Params) WithTimeout(timeout time.Duration) *GetAPIKeyDeploymentUsingGET1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get Api key deployment using g e t 1 params
func (o *GetAPIKeyDeploymentUsingGET1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get Api key deployment using g e t 1 params
func (o *GetAPIKeyDeploymentUsingGET1Params) WithContext(ctx context.Context) *GetAPIKeyDeploymentUsingGET1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get Api key deployment using g e t 1 params
func (o *GetAPIKeyDeploymentUsingGET1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get Api key deployment using g e t 1 params
func (o *GetAPIKeyDeploymentUsingGET1Params) WithHTTPClient(client *http.Client) *GetAPIKeyDeploymentUsingGET1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get Api key deployment using g e t 1 params
func (o *GetAPIKeyDeploymentUsingGET1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIKey adds the aPIKey to the get Api key deployment using g e t 1 params
func (o *GetAPIKeyDeploymentUsingGET1Params) WithAPIKey(aPIKey string) *GetAPIKeyDeploymentUsingGET1Params {
	o.SetAPIKey(aPIKey)
	return o
}

// SetAPIKey adds the apiKey to the get Api key deployment using g e t 1 params
func (o *GetAPIKeyDeploymentUsingGET1Params) SetAPIKey(aPIKey string) {
	o.APIKey = aPIKey
}

// WithProxyUUID adds the proxyUUID to the get Api key deployment using g e t 1 params
func (o *GetAPIKeyDeploymentUsingGET1Params) WithProxyUUID(proxyUUID string) *GetAPIKeyDeploymentUsingGET1Params {
	o.SetProxyUUID(proxyUUID)
	return o
}

// SetProxyUUID adds the proxyUuid to the get Api key deployment using g e t 1 params
func (o *GetAPIKeyDeploymentUsingGET1Params) SetProxyUUID(proxyUUID string) {
	o.ProxyUUID = proxyUUID
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPIKeyDeploymentUsingGET1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
