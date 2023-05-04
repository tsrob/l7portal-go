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
)

// NewCreateAPIGroupParams creates a new CreateAPIGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateAPIGroupParams() *CreateAPIGroupParams {
	return &CreateAPIGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateAPIGroupParamsWithTimeout creates a new CreateAPIGroupParams object
// with the ability to set a timeout on a request.
func NewCreateAPIGroupParamsWithTimeout(timeout time.Duration) *CreateAPIGroupParams {
	return &CreateAPIGroupParams{
		timeout: timeout,
	}
}

// NewCreateAPIGroupParamsWithContext creates a new CreateAPIGroupParams object
// with the ability to set a context for a request.
func NewCreateAPIGroupParamsWithContext(ctx context.Context) *CreateAPIGroupParams {
	return &CreateAPIGroupParams{
		Context: ctx,
	}
}

// NewCreateAPIGroupParamsWithHTTPClient creates a new CreateAPIGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateAPIGroupParamsWithHTTPClient(client *http.Client) *CreateAPIGroupParams {
	return &CreateAPIGroupParams{
		HTTPClient: client,
	}
}

/*
CreateAPIGroupParams contains all the parameters to send to the API endpoint

	for the create Api group operation.

	Typically these are written to a http.Request.
*/
type CreateAPIGroupParams struct {

	/* Body.

	   Creates an API Group.
	*/
	Body interface{}

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create Api group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateAPIGroupParams) WithDefaults() *CreateAPIGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create Api group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateAPIGroupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create Api group params
func (o *CreateAPIGroupParams) WithTimeout(timeout time.Duration) *CreateAPIGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create Api group params
func (o *CreateAPIGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create Api group params
func (o *CreateAPIGroupParams) WithContext(ctx context.Context) *CreateAPIGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create Api group params
func (o *CreateAPIGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create Api group params
func (o *CreateAPIGroupParams) WithHTTPClient(client *http.Client) *CreateAPIGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create Api group params
func (o *CreateAPIGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create Api group params
func (o *CreateAPIGroupParams) WithBody(body interface{}) *CreateAPIGroupParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create Api group params
func (o *CreateAPIGroupParams) SetBody(body interface{}) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateAPIGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
