// Code generated by go-swagger; DO NOT EDIT.

package proxies

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

// NewUpdateDeploymentTypeUsingPUTParams creates a new UpdateDeploymentTypeUsingPUTParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateDeploymentTypeUsingPUTParams() *UpdateDeploymentTypeUsingPUTParams {
	return &UpdateDeploymentTypeUsingPUTParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateDeploymentTypeUsingPUTParamsWithTimeout creates a new UpdateDeploymentTypeUsingPUTParams object
// with the ability to set a timeout on a request.
func NewUpdateDeploymentTypeUsingPUTParamsWithTimeout(timeout time.Duration) *UpdateDeploymentTypeUsingPUTParams {
	return &UpdateDeploymentTypeUsingPUTParams{
		timeout: timeout,
	}
}

// NewUpdateDeploymentTypeUsingPUTParamsWithContext creates a new UpdateDeploymentTypeUsingPUTParams object
// with the ability to set a context for a request.
func NewUpdateDeploymentTypeUsingPUTParamsWithContext(ctx context.Context) *UpdateDeploymentTypeUsingPUTParams {
	return &UpdateDeploymentTypeUsingPUTParams{
		Context: ctx,
	}
}

// NewUpdateDeploymentTypeUsingPUTParamsWithHTTPClient creates a new UpdateDeploymentTypeUsingPUTParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateDeploymentTypeUsingPUTParamsWithHTTPClient(client *http.Client) *UpdateDeploymentTypeUsingPUTParams {
	return &UpdateDeploymentTypeUsingPUTParams{
		HTTPClient: client,
	}
}

/*
UpdateDeploymentTypeUsingPUTParams contains all the parameters to send to the API endpoint

	for the update deployment type using p u t operation.

	Typically these are written to a http.Request.
*/
type UpdateDeploymentTypeUsingPUTParams struct {

	/* Body.

	   The updated proxy deployment type.
	*/
	Body *models.EntityDeploymentTypeDto

	/* Entity.

	   The proxy deployment type to update.
	*/
	Entity string

	/* UUID.

	   The proxy UUID to update.
	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update deployment type using p u t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDeploymentTypeUsingPUTParams) WithDefaults() *UpdateDeploymentTypeUsingPUTParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update deployment type using p u t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDeploymentTypeUsingPUTParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update deployment type using p u t params
func (o *UpdateDeploymentTypeUsingPUTParams) WithTimeout(timeout time.Duration) *UpdateDeploymentTypeUsingPUTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update deployment type using p u t params
func (o *UpdateDeploymentTypeUsingPUTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update deployment type using p u t params
func (o *UpdateDeploymentTypeUsingPUTParams) WithContext(ctx context.Context) *UpdateDeploymentTypeUsingPUTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update deployment type using p u t params
func (o *UpdateDeploymentTypeUsingPUTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update deployment type using p u t params
func (o *UpdateDeploymentTypeUsingPUTParams) WithHTTPClient(client *http.Client) *UpdateDeploymentTypeUsingPUTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update deployment type using p u t params
func (o *UpdateDeploymentTypeUsingPUTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update deployment type using p u t params
func (o *UpdateDeploymentTypeUsingPUTParams) WithBody(body *models.EntityDeploymentTypeDto) *UpdateDeploymentTypeUsingPUTParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update deployment type using p u t params
func (o *UpdateDeploymentTypeUsingPUTParams) SetBody(body *models.EntityDeploymentTypeDto) {
	o.Body = body
}

// WithEntity adds the entity to the update deployment type using p u t params
func (o *UpdateDeploymentTypeUsingPUTParams) WithEntity(entity string) *UpdateDeploymentTypeUsingPUTParams {
	o.SetEntity(entity)
	return o
}

// SetEntity adds the entity to the update deployment type using p u t params
func (o *UpdateDeploymentTypeUsingPUTParams) SetEntity(entity string) {
	o.Entity = entity
}

// WithUUID adds the uuid to the update deployment type using p u t params
func (o *UpdateDeploymentTypeUsingPUTParams) WithUUID(uuid string) *UpdateDeploymentTypeUsingPUTParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the update deployment type using p u t params
func (o *UpdateDeploymentTypeUsingPUTParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateDeploymentTypeUsingPUTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param entity
	if err := r.SetPathParam("entity", o.Entity); err != nil {
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
