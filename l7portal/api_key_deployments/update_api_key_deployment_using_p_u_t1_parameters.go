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

	"github.com/tsrob/l7portal-go/models"
)

// NewUpdateAPIKeyDeploymentUsingPUT1Params creates a new UpdateAPIKeyDeploymentUsingPUT1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateAPIKeyDeploymentUsingPUT1Params() *UpdateAPIKeyDeploymentUsingPUT1Params {
	return &UpdateAPIKeyDeploymentUsingPUT1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateAPIKeyDeploymentUsingPUT1ParamsWithTimeout creates a new UpdateAPIKeyDeploymentUsingPUT1Params object
// with the ability to set a timeout on a request.
func NewUpdateAPIKeyDeploymentUsingPUT1ParamsWithTimeout(timeout time.Duration) *UpdateAPIKeyDeploymentUsingPUT1Params {
	return &UpdateAPIKeyDeploymentUsingPUT1Params{
		timeout: timeout,
	}
}

// NewUpdateAPIKeyDeploymentUsingPUT1ParamsWithContext creates a new UpdateAPIKeyDeploymentUsingPUT1Params object
// with the ability to set a context for a request.
func NewUpdateAPIKeyDeploymentUsingPUT1ParamsWithContext(ctx context.Context) *UpdateAPIKeyDeploymentUsingPUT1Params {
	return &UpdateAPIKeyDeploymentUsingPUT1Params{
		Context: ctx,
	}
}

// NewUpdateAPIKeyDeploymentUsingPUT1ParamsWithHTTPClient creates a new UpdateAPIKeyDeploymentUsingPUT1Params object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateAPIKeyDeploymentUsingPUT1ParamsWithHTTPClient(client *http.Client) *UpdateAPIKeyDeploymentUsingPUT1Params {
	return &UpdateAPIKeyDeploymentUsingPUT1Params{
		HTTPClient: client,
	}
}

/*
UpdateAPIKeyDeploymentUsingPUT1Params contains all the parameters to send to the API endpoint

	for the update Api key deployment using p u t 1 operation.

	Typically these are written to a http.Request.
*/
type UpdateAPIKeyDeploymentUsingPUT1Params struct {

	/* APIKey.

	   The API Key.
	*/
	APIKey string

	/* DeploymentUpdateDto.

	   The updated API deployment.
	*/
	DeploymentUpdateDto *models.DeploymentUpdateDto

	/* ProxyUUID.

	   The proxy UUID.
	*/
	ProxyUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update Api key deployment using p u t 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateAPIKeyDeploymentUsingPUT1Params) WithDefaults() *UpdateAPIKeyDeploymentUsingPUT1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update Api key deployment using p u t 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateAPIKeyDeploymentUsingPUT1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update Api key deployment using p u t 1 params
func (o *UpdateAPIKeyDeploymentUsingPUT1Params) WithTimeout(timeout time.Duration) *UpdateAPIKeyDeploymentUsingPUT1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update Api key deployment using p u t 1 params
func (o *UpdateAPIKeyDeploymentUsingPUT1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update Api key deployment using p u t 1 params
func (o *UpdateAPIKeyDeploymentUsingPUT1Params) WithContext(ctx context.Context) *UpdateAPIKeyDeploymentUsingPUT1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update Api key deployment using p u t 1 params
func (o *UpdateAPIKeyDeploymentUsingPUT1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update Api key deployment using p u t 1 params
func (o *UpdateAPIKeyDeploymentUsingPUT1Params) WithHTTPClient(client *http.Client) *UpdateAPIKeyDeploymentUsingPUT1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update Api key deployment using p u t 1 params
func (o *UpdateAPIKeyDeploymentUsingPUT1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIKey adds the aPIKey to the update Api key deployment using p u t 1 params
func (o *UpdateAPIKeyDeploymentUsingPUT1Params) WithAPIKey(aPIKey string) *UpdateAPIKeyDeploymentUsingPUT1Params {
	o.SetAPIKey(aPIKey)
	return o
}

// SetAPIKey adds the apiKey to the update Api key deployment using p u t 1 params
func (o *UpdateAPIKeyDeploymentUsingPUT1Params) SetAPIKey(aPIKey string) {
	o.APIKey = aPIKey
}

// WithDeploymentUpdateDto adds the deploymentUpdateDto to the update Api key deployment using p u t 1 params
func (o *UpdateAPIKeyDeploymentUsingPUT1Params) WithDeploymentUpdateDto(deploymentUpdateDto *models.DeploymentUpdateDto) *UpdateAPIKeyDeploymentUsingPUT1Params {
	o.SetDeploymentUpdateDto(deploymentUpdateDto)
	return o
}

// SetDeploymentUpdateDto adds the deploymentUpdateDto to the update Api key deployment using p u t 1 params
func (o *UpdateAPIKeyDeploymentUsingPUT1Params) SetDeploymentUpdateDto(deploymentUpdateDto *models.DeploymentUpdateDto) {
	o.DeploymentUpdateDto = deploymentUpdateDto
}

// WithProxyUUID adds the proxyUUID to the update Api key deployment using p u t 1 params
func (o *UpdateAPIKeyDeploymentUsingPUT1Params) WithProxyUUID(proxyUUID string) *UpdateAPIKeyDeploymentUsingPUT1Params {
	o.SetProxyUUID(proxyUUID)
	return o
}

// SetProxyUUID adds the proxyUuid to the update Api key deployment using p u t 1 params
func (o *UpdateAPIKeyDeploymentUsingPUT1Params) SetProxyUUID(proxyUUID string) {
	o.ProxyUUID = proxyUUID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateAPIKeyDeploymentUsingPUT1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param apiKey
	if err := r.SetPathParam("apiKey", o.APIKey); err != nil {
		return err
	}
	if o.DeploymentUpdateDto != nil {
		if err := r.SetBodyParam(o.DeploymentUpdateDto); err != nil {
			return err
		}
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
