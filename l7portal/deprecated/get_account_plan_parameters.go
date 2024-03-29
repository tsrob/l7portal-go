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

// NewGetAccountPlanParams creates a new GetAccountPlanParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAccountPlanParams() *GetAccountPlanParams {
	return &GetAccountPlanParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAccountPlanParamsWithTimeout creates a new GetAccountPlanParams object
// with the ability to set a timeout on a request.
func NewGetAccountPlanParamsWithTimeout(timeout time.Duration) *GetAccountPlanParams {
	return &GetAccountPlanParams{
		timeout: timeout,
	}
}

// NewGetAccountPlanParamsWithContext creates a new GetAccountPlanParams object
// with the ability to set a context for a request.
func NewGetAccountPlanParamsWithContext(ctx context.Context) *GetAccountPlanParams {
	return &GetAccountPlanParams{
		Context: ctx,
	}
}

// NewGetAccountPlanParamsWithHTTPClient creates a new GetAccountPlanParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAccountPlanParamsWithHTTPClient(client *http.Client) *GetAccountPlanParams {
	return &GetAccountPlanParams{
		HTTPClient: client,
	}
}

/*
GetAccountPlanParams contains all the parameters to send to the API endpoint

	for the get account plan operation.

	Typically these are written to a http.Request.
*/
type GetAccountPlanParams struct {

	/* UUID.

	   The UUID of the Account Plan to return.
	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get account plan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAccountPlanParams) WithDefaults() *GetAccountPlanParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get account plan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAccountPlanParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get account plan params
func (o *GetAccountPlanParams) WithTimeout(timeout time.Duration) *GetAccountPlanParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get account plan params
func (o *GetAccountPlanParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get account plan params
func (o *GetAccountPlanParams) WithContext(ctx context.Context) *GetAccountPlanParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get account plan params
func (o *GetAccountPlanParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get account plan params
func (o *GetAccountPlanParams) WithHTTPClient(client *http.Client) *GetAccountPlanParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get account plan params
func (o *GetAccountPlanParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUUID adds the uuid to the get account plan params
func (o *GetAccountPlanParams) WithUUID(uuid string) *GetAccountPlanParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the get account plan params
func (o *GetAccountPlanParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *GetAccountPlanParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
