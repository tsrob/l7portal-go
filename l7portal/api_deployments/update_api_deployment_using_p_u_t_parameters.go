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

	"github.com/tsrob/l7portal-go/models"
)

// NewUpdateAPIDeploymentUsingPUTParams creates a new UpdateAPIDeploymentUsingPUTParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateAPIDeploymentUsingPUTParams() *UpdateAPIDeploymentUsingPUTParams {
	return &UpdateAPIDeploymentUsingPUTParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateAPIDeploymentUsingPUTParamsWithTimeout creates a new UpdateAPIDeploymentUsingPUTParams object
// with the ability to set a timeout on a request.
func NewUpdateAPIDeploymentUsingPUTParamsWithTimeout(timeout time.Duration) *UpdateAPIDeploymentUsingPUTParams {
	return &UpdateAPIDeploymentUsingPUTParams{
		timeout: timeout,
	}
}

// NewUpdateAPIDeploymentUsingPUTParamsWithContext creates a new UpdateAPIDeploymentUsingPUTParams object
// with the ability to set a context for a request.
func NewUpdateAPIDeploymentUsingPUTParamsWithContext(ctx context.Context) *UpdateAPIDeploymentUsingPUTParams {
	return &UpdateAPIDeploymentUsingPUTParams{
		Context: ctx,
	}
}

// NewUpdateAPIDeploymentUsingPUTParamsWithHTTPClient creates a new UpdateAPIDeploymentUsingPUTParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateAPIDeploymentUsingPUTParamsWithHTTPClient(client *http.Client) *UpdateAPIDeploymentUsingPUTParams {
	return &UpdateAPIDeploymentUsingPUTParams{
		HTTPClient: client,
	}
}

/*
UpdateAPIDeploymentUsingPUTParams contains all the parameters to send to the API endpoint

	for the update Api deployment using p u t operation.

	Typically these are written to a http.Request.
*/
type UpdateAPIDeploymentUsingPUTParams struct {

	/* APIDeployUpdateDto.

	   The updated API deployment.
	*/
	APIDeployUpdateDto *models.APIDeployUpdateDto

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

// WithDefaults hydrates default values in the update Api deployment using p u t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateAPIDeploymentUsingPUTParams) WithDefaults() *UpdateAPIDeploymentUsingPUTParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update Api deployment using p u t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateAPIDeploymentUsingPUTParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update Api deployment using p u t params
func (o *UpdateAPIDeploymentUsingPUTParams) WithTimeout(timeout time.Duration) *UpdateAPIDeploymentUsingPUTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update Api deployment using p u t params
func (o *UpdateAPIDeploymentUsingPUTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update Api deployment using p u t params
func (o *UpdateAPIDeploymentUsingPUTParams) WithContext(ctx context.Context) *UpdateAPIDeploymentUsingPUTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update Api deployment using p u t params
func (o *UpdateAPIDeploymentUsingPUTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update Api deployment using p u t params
func (o *UpdateAPIDeploymentUsingPUTParams) WithHTTPClient(client *http.Client) *UpdateAPIDeploymentUsingPUTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update Api deployment using p u t params
func (o *UpdateAPIDeploymentUsingPUTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIDeployUpdateDto adds the aPIDeployUpdateDto to the update Api deployment using p u t params
func (o *UpdateAPIDeploymentUsingPUTParams) WithAPIDeployUpdateDto(aPIDeployUpdateDto *models.APIDeployUpdateDto) *UpdateAPIDeploymentUsingPUTParams {
	o.SetAPIDeployUpdateDto(aPIDeployUpdateDto)
	return o
}

// SetAPIDeployUpdateDto adds the apiDeployUpdateDto to the update Api deployment using p u t params
func (o *UpdateAPIDeploymentUsingPUTParams) SetAPIDeployUpdateDto(aPIDeployUpdateDto *models.APIDeployUpdateDto) {
	o.APIDeployUpdateDto = aPIDeployUpdateDto
}

// WithAPIUUID adds the aPIUUID to the update Api deployment using p u t params
func (o *UpdateAPIDeploymentUsingPUTParams) WithAPIUUID(aPIUUID string) *UpdateAPIDeploymentUsingPUTParams {
	o.SetAPIUUID(aPIUUID)
	return o
}

// SetAPIUUID adds the apiUuid to the update Api deployment using p u t params
func (o *UpdateAPIDeploymentUsingPUTParams) SetAPIUUID(aPIUUID string) {
	o.APIUUID = aPIUUID
}

// WithProxyUUID adds the proxyUUID to the update Api deployment using p u t params
func (o *UpdateAPIDeploymentUsingPUTParams) WithProxyUUID(proxyUUID string) *UpdateAPIDeploymentUsingPUTParams {
	o.SetProxyUUID(proxyUUID)
	return o
}

// SetProxyUUID adds the proxyUuid to the update Api deployment using p u t params
func (o *UpdateAPIDeploymentUsingPUTParams) SetProxyUUID(proxyUUID string) {
	o.ProxyUUID = proxyUUID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateAPIDeploymentUsingPUTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.APIDeployUpdateDto != nil {
		if err := r.SetBodyParam(o.APIDeployUpdateDto); err != nil {
			return err
		}
	}

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
