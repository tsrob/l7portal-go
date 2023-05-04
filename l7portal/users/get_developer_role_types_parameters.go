// Code generated by go-swagger; DO NOT EDIT.

package users

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

// NewGetDeveloperRoleTypesParams creates a new GetDeveloperRoleTypesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDeveloperRoleTypesParams() *GetDeveloperRoleTypesParams {
	return &GetDeveloperRoleTypesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeveloperRoleTypesParamsWithTimeout creates a new GetDeveloperRoleTypesParams object
// with the ability to set a timeout on a request.
func NewGetDeveloperRoleTypesParamsWithTimeout(timeout time.Duration) *GetDeveloperRoleTypesParams {
	return &GetDeveloperRoleTypesParams{
		timeout: timeout,
	}
}

// NewGetDeveloperRoleTypesParamsWithContext creates a new GetDeveloperRoleTypesParams object
// with the ability to set a context for a request.
func NewGetDeveloperRoleTypesParamsWithContext(ctx context.Context) *GetDeveloperRoleTypesParams {
	return &GetDeveloperRoleTypesParams{
		Context: ctx,
	}
}

// NewGetDeveloperRoleTypesParamsWithHTTPClient creates a new GetDeveloperRoleTypesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDeveloperRoleTypesParamsWithHTTPClient(client *http.Client) *GetDeveloperRoleTypesParams {
	return &GetDeveloperRoleTypesParams{
		HTTPClient: client,
	}
}

/*
GetDeveloperRoleTypesParams contains all the parameters to send to the API endpoint

	for the get developer role types operation.

	Typically these are written to a http.Request.
*/
type GetDeveloperRoleTypesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get developer role types params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeveloperRoleTypesParams) WithDefaults() *GetDeveloperRoleTypesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get developer role types params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeveloperRoleTypesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get developer role types params
func (o *GetDeveloperRoleTypesParams) WithTimeout(timeout time.Duration) *GetDeveloperRoleTypesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get developer role types params
func (o *GetDeveloperRoleTypesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get developer role types params
func (o *GetDeveloperRoleTypesParams) WithContext(ctx context.Context) *GetDeveloperRoleTypesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get developer role types params
func (o *GetDeveloperRoleTypesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get developer role types params
func (o *GetDeveloperRoleTypesParams) WithHTTPClient(client *http.Client) *GetDeveloperRoleTypesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get developer role types params
func (o *GetDeveloperRoleTypesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeveloperRoleTypesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
