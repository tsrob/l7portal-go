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
)

// NewGetProxyUsingGETParams creates a new GetProxyUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetProxyUsingGETParams() *GetProxyUsingGETParams {
	return &GetProxyUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetProxyUsingGETParamsWithTimeout creates a new GetProxyUsingGETParams object
// with the ability to set a timeout on a request.
func NewGetProxyUsingGETParamsWithTimeout(timeout time.Duration) *GetProxyUsingGETParams {
	return &GetProxyUsingGETParams{
		timeout: timeout,
	}
}

// NewGetProxyUsingGETParamsWithContext creates a new GetProxyUsingGETParams object
// with the ability to set a context for a request.
func NewGetProxyUsingGETParamsWithContext(ctx context.Context) *GetProxyUsingGETParams {
	return &GetProxyUsingGETParams{
		Context: ctx,
	}
}

// NewGetProxyUsingGETParamsWithHTTPClient creates a new GetProxyUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetProxyUsingGETParamsWithHTTPClient(client *http.Client) *GetProxyUsingGETParams {
	return &GetProxyUsingGETParams{
		HTTPClient: client,
	}
}

/*
GetProxyUsingGETParams contains all the parameters to send to the API endpoint

	for the get proxy using g e t operation.

	Typically these are written to a http.Request.
*/
type GetProxyUsingGETParams struct {

	/* UUID.

	   The proxy UUID to get.
	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get proxy using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProxyUsingGETParams) WithDefaults() *GetProxyUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get proxy using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProxyUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get proxy using g e t params
func (o *GetProxyUsingGETParams) WithTimeout(timeout time.Duration) *GetProxyUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get proxy using g e t params
func (o *GetProxyUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get proxy using g e t params
func (o *GetProxyUsingGETParams) WithContext(ctx context.Context) *GetProxyUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get proxy using g e t params
func (o *GetProxyUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get proxy using g e t params
func (o *GetProxyUsingGETParams) WithHTTPClient(client *http.Client) *GetProxyUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get proxy using g e t params
func (o *GetProxyUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUUID adds the uuid to the get proxy using g e t params
func (o *GetProxyUsingGETParams) WithUUID(uuid string) *GetProxyUsingGETParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the get proxy using g e t params
func (o *GetProxyUsingGETParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *GetProxyUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
