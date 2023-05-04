// Code generated by go-swagger; DO NOT EDIT.

package api_deployments

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

// NewDeleteAPIDeploymentUsingDELETE1Params creates a new DeleteAPIDeploymentUsingDELETE1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteAPIDeploymentUsingDELETE1Params() *DeleteAPIDeploymentUsingDELETE1Params {
	return &DeleteAPIDeploymentUsingDELETE1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAPIDeploymentUsingDELETE1ParamsWithTimeout creates a new DeleteAPIDeploymentUsingDELETE1Params object
// with the ability to set a timeout on a request.
func NewDeleteAPIDeploymentUsingDELETE1ParamsWithTimeout(timeout time.Duration) *DeleteAPIDeploymentUsingDELETE1Params {
	return &DeleteAPIDeploymentUsingDELETE1Params{
		timeout: timeout,
	}
}

// NewDeleteAPIDeploymentUsingDELETE1ParamsWithContext creates a new DeleteAPIDeploymentUsingDELETE1Params object
// with the ability to set a context for a request.
func NewDeleteAPIDeploymentUsingDELETE1ParamsWithContext(ctx context.Context) *DeleteAPIDeploymentUsingDELETE1Params {
	return &DeleteAPIDeploymentUsingDELETE1Params{
		Context: ctx,
	}
}

// NewDeleteAPIDeploymentUsingDELETE1ParamsWithHTTPClient creates a new DeleteAPIDeploymentUsingDELETE1Params object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteAPIDeploymentUsingDELETE1ParamsWithHTTPClient(client *http.Client) *DeleteAPIDeploymentUsingDELETE1Params {
	return &DeleteAPIDeploymentUsingDELETE1Params{
		HTTPClient: client,
	}
}

/*
DeleteAPIDeploymentUsingDELETE1Params contains all the parameters to send to the API endpoint

	for the delete Api deployment using d e l e t e 1 operation.

	Typically these are written to a http.Request.
*/
type DeleteAPIDeploymentUsingDELETE1Params struct {

	/* APIUUID.

	   The API UUID.
	*/
	APIUUID string

	/* ProxyUUID.

	   The proxy UUID.
	*/
	ProxyUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete Api deployment using d e l e t e 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAPIDeploymentUsingDELETE1Params) WithDefaults() *DeleteAPIDeploymentUsingDELETE1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete Api deployment using d e l e t e 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAPIDeploymentUsingDELETE1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete Api deployment using d e l e t e 1 params
func (o *DeleteAPIDeploymentUsingDELETE1Params) WithTimeout(timeout time.Duration) *DeleteAPIDeploymentUsingDELETE1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete Api deployment using d e l e t e 1 params
func (o *DeleteAPIDeploymentUsingDELETE1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete Api deployment using d e l e t e 1 params
func (o *DeleteAPIDeploymentUsingDELETE1Params) WithContext(ctx context.Context) *DeleteAPIDeploymentUsingDELETE1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete Api deployment using d e l e t e 1 params
func (o *DeleteAPIDeploymentUsingDELETE1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete Api deployment using d e l e t e 1 params
func (o *DeleteAPIDeploymentUsingDELETE1Params) WithHTTPClient(client *http.Client) *DeleteAPIDeploymentUsingDELETE1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete Api deployment using d e l e t e 1 params
func (o *DeleteAPIDeploymentUsingDELETE1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIUUID adds the aPIUUID to the delete Api deployment using d e l e t e 1 params
func (o *DeleteAPIDeploymentUsingDELETE1Params) WithAPIUUID(aPIUUID string) *DeleteAPIDeploymentUsingDELETE1Params {
	o.SetAPIUUID(aPIUUID)
	return o
}

// SetAPIUUID adds the apiUuid to the delete Api deployment using d e l e t e 1 params
func (o *DeleteAPIDeploymentUsingDELETE1Params) SetAPIUUID(aPIUUID string) {
	o.APIUUID = aPIUUID
}

// WithProxyUUID adds the proxyUUID to the delete Api deployment using d e l e t e 1 params
func (o *DeleteAPIDeploymentUsingDELETE1Params) WithProxyUUID(proxyUUID string) *DeleteAPIDeploymentUsingDELETE1Params {
	o.SetProxyUUID(proxyUUID)
	return o
}

// SetProxyUUID adds the proxyUuid to the delete Api deployment using d e l e t e 1 params
func (o *DeleteAPIDeploymentUsingDELETE1Params) SetProxyUUID(proxyUUID string) {
	o.ProxyUUID = proxyUUID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAPIDeploymentUsingDELETE1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param apiUuid
	if err := r.SetPathParam("apiUuid", o.APIUUID); err != nil {
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
