// Code generated by go-swagger; DO NOT EDIT.

package apis

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

// NewGetAssociatedAPIOrgRLQParams creates a new GetAssociatedAPIOrgRLQParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAssociatedAPIOrgRLQParams() *GetAssociatedAPIOrgRLQParams {
	return &GetAssociatedAPIOrgRLQParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAssociatedAPIOrgRLQParamsWithTimeout creates a new GetAssociatedAPIOrgRLQParams object
// with the ability to set a timeout on a request.
func NewGetAssociatedAPIOrgRLQParamsWithTimeout(timeout time.Duration) *GetAssociatedAPIOrgRLQParams {
	return &GetAssociatedAPIOrgRLQParams{
		timeout: timeout,
	}
}

// NewGetAssociatedAPIOrgRLQParamsWithContext creates a new GetAssociatedAPIOrgRLQParams object
// with the ability to set a context for a request.
func NewGetAssociatedAPIOrgRLQParamsWithContext(ctx context.Context) *GetAssociatedAPIOrgRLQParams {
	return &GetAssociatedAPIOrgRLQParams{
		Context: ctx,
	}
}

// NewGetAssociatedAPIOrgRLQParamsWithHTTPClient creates a new GetAssociatedAPIOrgRLQParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAssociatedAPIOrgRLQParamsWithHTTPClient(client *http.Client) *GetAssociatedAPIOrgRLQParams {
	return &GetAssociatedAPIOrgRLQParams{
		HTTPClient: client,
	}
}

/*
GetAssociatedAPIOrgRLQParams contains all the parameters to send to the API endpoint

	for the get associated Api org r l q operation.

	Typically these are written to a http.Request.
*/
type GetAssociatedAPIOrgRLQParams struct {

	/* APIUUID.

	   apiUuid
	*/
	APIUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get associated Api org r l q params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAssociatedAPIOrgRLQParams) WithDefaults() *GetAssociatedAPIOrgRLQParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get associated Api org r l q params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAssociatedAPIOrgRLQParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get associated Api org r l q params
func (o *GetAssociatedAPIOrgRLQParams) WithTimeout(timeout time.Duration) *GetAssociatedAPIOrgRLQParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get associated Api org r l q params
func (o *GetAssociatedAPIOrgRLQParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get associated Api org r l q params
func (o *GetAssociatedAPIOrgRLQParams) WithContext(ctx context.Context) *GetAssociatedAPIOrgRLQParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get associated Api org r l q params
func (o *GetAssociatedAPIOrgRLQParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get associated Api org r l q params
func (o *GetAssociatedAPIOrgRLQParams) WithHTTPClient(client *http.Client) *GetAssociatedAPIOrgRLQParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get associated Api org r l q params
func (o *GetAssociatedAPIOrgRLQParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIUUID adds the aPIUUID to the get associated Api org r l q params
func (o *GetAssociatedAPIOrgRLQParams) WithAPIUUID(aPIUUID string) *GetAssociatedAPIOrgRLQParams {
	o.SetAPIUUID(aPIUUID)
	return o
}

// SetAPIUUID adds the apiUuid to the get associated Api org r l q params
func (o *GetAssociatedAPIOrgRLQParams) SetAPIUUID(aPIUUID string) {
	o.APIUUID = aPIUUID
}

// WriteToRequest writes these params to a swagger request
func (o *GetAssociatedAPIOrgRLQParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param apiUuid
	if err := r.SetPathParam("apiUuid", o.APIUUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}