// Code generated by go-swagger; DO NOT EDIT.

package api_groups

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

// NewUpdateAPIGroupAPIMappingParams creates a new UpdateAPIGroupAPIMappingParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateAPIGroupAPIMappingParams() *UpdateAPIGroupAPIMappingParams {
	return &UpdateAPIGroupAPIMappingParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateAPIGroupAPIMappingParamsWithTimeout creates a new UpdateAPIGroupAPIMappingParams object
// with the ability to set a timeout on a request.
func NewUpdateAPIGroupAPIMappingParamsWithTimeout(timeout time.Duration) *UpdateAPIGroupAPIMappingParams {
	return &UpdateAPIGroupAPIMappingParams{
		timeout: timeout,
	}
}

// NewUpdateAPIGroupAPIMappingParamsWithContext creates a new UpdateAPIGroupAPIMappingParams object
// with the ability to set a context for a request.
func NewUpdateAPIGroupAPIMappingParamsWithContext(ctx context.Context) *UpdateAPIGroupAPIMappingParams {
	return &UpdateAPIGroupAPIMappingParams{
		Context: ctx,
	}
}

// NewUpdateAPIGroupAPIMappingParamsWithHTTPClient creates a new UpdateAPIGroupAPIMappingParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateAPIGroupAPIMappingParamsWithHTTPClient(client *http.Client) *UpdateAPIGroupAPIMappingParams {
	return &UpdateAPIGroupAPIMappingParams{
		HTTPClient: client,
	}
}

/*
UpdateAPIGroupAPIMappingParams contains all the parameters to send to the API endpoint

	for the update Api group Api mapping operation.

	Typically these are written to a http.Request.
*/
type UpdateAPIGroupAPIMappingParams struct {

	/* Body.

	   List of API references to associate with the specified API Group.
	*/
	Body models.APIGroupApis

	/* UUID.

	   The UUID of the API Group to associate the provided APIs with.
	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update Api group Api mapping params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateAPIGroupAPIMappingParams) WithDefaults() *UpdateAPIGroupAPIMappingParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update Api group Api mapping params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateAPIGroupAPIMappingParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update Api group Api mapping params
func (o *UpdateAPIGroupAPIMappingParams) WithTimeout(timeout time.Duration) *UpdateAPIGroupAPIMappingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update Api group Api mapping params
func (o *UpdateAPIGroupAPIMappingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update Api group Api mapping params
func (o *UpdateAPIGroupAPIMappingParams) WithContext(ctx context.Context) *UpdateAPIGroupAPIMappingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update Api group Api mapping params
func (o *UpdateAPIGroupAPIMappingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update Api group Api mapping params
func (o *UpdateAPIGroupAPIMappingParams) WithHTTPClient(client *http.Client) *UpdateAPIGroupAPIMappingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update Api group Api mapping params
func (o *UpdateAPIGroupAPIMappingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update Api group Api mapping params
func (o *UpdateAPIGroupAPIMappingParams) WithBody(body models.APIGroupApis) *UpdateAPIGroupAPIMappingParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update Api group Api mapping params
func (o *UpdateAPIGroupAPIMappingParams) SetBody(body models.APIGroupApis) {
	o.Body = body
}

// WithUUID adds the uuid to the update Api group Api mapping params
func (o *UpdateAPIGroupAPIMappingParams) WithUUID(uuid string) *UpdateAPIGroupAPIMappingParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the update Api group Api mapping params
func (o *UpdateAPIGroupAPIMappingParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateAPIGroupAPIMappingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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