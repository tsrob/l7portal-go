// Code generated by go-swagger; DO NOT EDIT.

package deprecated

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

// NewGetApplicationParams creates a new GetApplicationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetApplicationParams() *GetApplicationParams {
	return &GetApplicationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetApplicationParamsWithTimeout creates a new GetApplicationParams object
// with the ability to set a timeout on a request.
func NewGetApplicationParamsWithTimeout(timeout time.Duration) *GetApplicationParams {
	return &GetApplicationParams{
		timeout: timeout,
	}
}

// NewGetApplicationParamsWithContext creates a new GetApplicationParams object
// with the ability to set a context for a request.
func NewGetApplicationParamsWithContext(ctx context.Context) *GetApplicationParams {
	return &GetApplicationParams{
		Context: ctx,
	}
}

// NewGetApplicationParamsWithHTTPClient creates a new GetApplicationParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetApplicationParamsWithHTTPClient(client *http.Client) *GetApplicationParams {
	return &GetApplicationParams{
		HTTPClient: client,
	}
}

/*
GetApplicationParams contains all the parameters to send to the API endpoint

	for the get application operation.

	Typically these are written to a http.Request.
*/
type GetApplicationParams struct {

	/* UUID.

	   The UUID of the application to return.
	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get application params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetApplicationParams) WithDefaults() *GetApplicationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get application params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetApplicationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get application params
func (o *GetApplicationParams) WithTimeout(timeout time.Duration) *GetApplicationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get application params
func (o *GetApplicationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get application params
func (o *GetApplicationParams) WithContext(ctx context.Context) *GetApplicationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get application params
func (o *GetApplicationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get application params
func (o *GetApplicationParams) WithHTTPClient(client *http.Client) *GetApplicationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get application params
func (o *GetApplicationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUUID adds the uuid to the get application params
func (o *GetApplicationParams) WithUUID(uuid string) *GetApplicationParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the get application params
func (o *GetApplicationParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *GetApplicationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
