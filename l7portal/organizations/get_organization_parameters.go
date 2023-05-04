// Code generated by go-swagger; DO NOT EDIT.

package organizations

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

// NewGetOrganizationParams creates a new GetOrganizationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOrganizationParams() *GetOrganizationParams {
	return &GetOrganizationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrganizationParamsWithTimeout creates a new GetOrganizationParams object
// with the ability to set a timeout on a request.
func NewGetOrganizationParamsWithTimeout(timeout time.Duration) *GetOrganizationParams {
	return &GetOrganizationParams{
		timeout: timeout,
	}
}

// NewGetOrganizationParamsWithContext creates a new GetOrganizationParams object
// with the ability to set a context for a request.
func NewGetOrganizationParamsWithContext(ctx context.Context) *GetOrganizationParams {
	return &GetOrganizationParams{
		Context: ctx,
	}
}

// NewGetOrganizationParamsWithHTTPClient creates a new GetOrganizationParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOrganizationParamsWithHTTPClient(client *http.Client) *GetOrganizationParams {
	return &GetOrganizationParams{
		HTTPClient: client,
	}
}

/*
GetOrganizationParams contains all the parameters to send to the API endpoint

	for the get organization operation.

	Typically these are written to a http.Request.
*/
type GetOrganizationParams struct {

	/* UUID.

	   The UUID of the organization to return.
	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get organization params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationParams) WithDefaults() *GetOrganizationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get organization params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get organization params
func (o *GetOrganizationParams) WithTimeout(timeout time.Duration) *GetOrganizationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get organization params
func (o *GetOrganizationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get organization params
func (o *GetOrganizationParams) WithContext(ctx context.Context) *GetOrganizationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get organization params
func (o *GetOrganizationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get organization params
func (o *GetOrganizationParams) WithHTTPClient(client *http.Client) *GetOrganizationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get organization params
func (o *GetOrganizationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUUID adds the uuid to the get organization params
func (o *GetOrganizationParams) WithUUID(uuid string) *GetOrganizationParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the get organization params
func (o *GetOrganizationParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrganizationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
