// Code generated by go-swagger; DO NOT EDIT.

package gateway_bundle_deployments

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

// NewUpdateGatewayBundleDeploymentUsingPUTParams creates a new UpdateGatewayBundleDeploymentUsingPUTParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateGatewayBundleDeploymentUsingPUTParams() *UpdateGatewayBundleDeploymentUsingPUTParams {
	return &UpdateGatewayBundleDeploymentUsingPUTParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateGatewayBundleDeploymentUsingPUTParamsWithTimeout creates a new UpdateGatewayBundleDeploymentUsingPUTParams object
// with the ability to set a timeout on a request.
func NewUpdateGatewayBundleDeploymentUsingPUTParamsWithTimeout(timeout time.Duration) *UpdateGatewayBundleDeploymentUsingPUTParams {
	return &UpdateGatewayBundleDeploymentUsingPUTParams{
		timeout: timeout,
	}
}

// NewUpdateGatewayBundleDeploymentUsingPUTParamsWithContext creates a new UpdateGatewayBundleDeploymentUsingPUTParams object
// with the ability to set a context for a request.
func NewUpdateGatewayBundleDeploymentUsingPUTParamsWithContext(ctx context.Context) *UpdateGatewayBundleDeploymentUsingPUTParams {
	return &UpdateGatewayBundleDeploymentUsingPUTParams{
		Context: ctx,
	}
}

// NewUpdateGatewayBundleDeploymentUsingPUTParamsWithHTTPClient creates a new UpdateGatewayBundleDeploymentUsingPUTParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateGatewayBundleDeploymentUsingPUTParamsWithHTTPClient(client *http.Client) *UpdateGatewayBundleDeploymentUsingPUTParams {
	return &UpdateGatewayBundleDeploymentUsingPUTParams{
		HTTPClient: client,
	}
}

/*
UpdateGatewayBundleDeploymentUsingPUTParams contains all the parameters to send to the API endpoint

	for the update gateway bundle deployment using p u t operation.

	Typically these are written to a http.Request.
*/
type UpdateGatewayBundleDeploymentUsingPUTParams struct {

	/* GatewayBundleDeployUpdateDto.

	   The updated Gateway bundle deployment
	*/
	GatewayBundleDeployUpdateDto *models.GatewayBundleDeployUpdateDto

	/* ProxyUUID.

	   The Proxy uuid.
	*/
	ProxyUUID string

	/* UUID.

	   The Gateway bundle uuid.
	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update gateway bundle deployment using p u t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateGatewayBundleDeploymentUsingPUTParams) WithDefaults() *UpdateGatewayBundleDeploymentUsingPUTParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update gateway bundle deployment using p u t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateGatewayBundleDeploymentUsingPUTParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update gateway bundle deployment using p u t params
func (o *UpdateGatewayBundleDeploymentUsingPUTParams) WithTimeout(timeout time.Duration) *UpdateGatewayBundleDeploymentUsingPUTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update gateway bundle deployment using p u t params
func (o *UpdateGatewayBundleDeploymentUsingPUTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update gateway bundle deployment using p u t params
func (o *UpdateGatewayBundleDeploymentUsingPUTParams) WithContext(ctx context.Context) *UpdateGatewayBundleDeploymentUsingPUTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update gateway bundle deployment using p u t params
func (o *UpdateGatewayBundleDeploymentUsingPUTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update gateway bundle deployment using p u t params
func (o *UpdateGatewayBundleDeploymentUsingPUTParams) WithHTTPClient(client *http.Client) *UpdateGatewayBundleDeploymentUsingPUTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update gateway bundle deployment using p u t params
func (o *UpdateGatewayBundleDeploymentUsingPUTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGatewayBundleDeployUpdateDto adds the gatewayBundleDeployUpdateDto to the update gateway bundle deployment using p u t params
func (o *UpdateGatewayBundleDeploymentUsingPUTParams) WithGatewayBundleDeployUpdateDto(gatewayBundleDeployUpdateDto *models.GatewayBundleDeployUpdateDto) *UpdateGatewayBundleDeploymentUsingPUTParams {
	o.SetGatewayBundleDeployUpdateDto(gatewayBundleDeployUpdateDto)
	return o
}

// SetGatewayBundleDeployUpdateDto adds the gatewayBundleDeployUpdateDto to the update gateway bundle deployment using p u t params
func (o *UpdateGatewayBundleDeploymentUsingPUTParams) SetGatewayBundleDeployUpdateDto(gatewayBundleDeployUpdateDto *models.GatewayBundleDeployUpdateDto) {
	o.GatewayBundleDeployUpdateDto = gatewayBundleDeployUpdateDto
}

// WithProxyUUID adds the proxyUUID to the update gateway bundle deployment using p u t params
func (o *UpdateGatewayBundleDeploymentUsingPUTParams) WithProxyUUID(proxyUUID string) *UpdateGatewayBundleDeploymentUsingPUTParams {
	o.SetProxyUUID(proxyUUID)
	return o
}

// SetProxyUUID adds the proxyUuid to the update gateway bundle deployment using p u t params
func (o *UpdateGatewayBundleDeploymentUsingPUTParams) SetProxyUUID(proxyUUID string) {
	o.ProxyUUID = proxyUUID
}

// WithUUID adds the uuid to the update gateway bundle deployment using p u t params
func (o *UpdateGatewayBundleDeploymentUsingPUTParams) WithUUID(uuid string) *UpdateGatewayBundleDeploymentUsingPUTParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the update gateway bundle deployment using p u t params
func (o *UpdateGatewayBundleDeploymentUsingPUTParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateGatewayBundleDeploymentUsingPUTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.GatewayBundleDeployUpdateDto != nil {
		if err := r.SetBodyParam(o.GatewayBundleDeployUpdateDto); err != nil {
			return err
		}
	}

	// path param proxyUuid
	if err := r.SetPathParam("proxyUuid", o.ProxyUUID); err != nil {
		return err
	}

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}