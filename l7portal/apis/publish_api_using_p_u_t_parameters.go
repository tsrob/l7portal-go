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

// NewPublishAPIUsingPUTParams creates a new PublishAPIUsingPUTParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPublishAPIUsingPUTParams() *PublishAPIUsingPUTParams {
	return &PublishAPIUsingPUTParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPublishAPIUsingPUTParamsWithTimeout creates a new PublishAPIUsingPUTParams object
// with the ability to set a timeout on a request.
func NewPublishAPIUsingPUTParamsWithTimeout(timeout time.Duration) *PublishAPIUsingPUTParams {
	return &PublishAPIUsingPUTParams{
		timeout: timeout,
	}
}

// NewPublishAPIUsingPUTParamsWithContext creates a new PublishAPIUsingPUTParams object
// with the ability to set a context for a request.
func NewPublishAPIUsingPUTParamsWithContext(ctx context.Context) *PublishAPIUsingPUTParams {
	return &PublishAPIUsingPUTParams{
		Context: ctx,
	}
}

// NewPublishAPIUsingPUTParamsWithHTTPClient creates a new PublishAPIUsingPUTParams object
// with the ability to set a custom HTTPClient for a request.
func NewPublishAPIUsingPUTParamsWithHTTPClient(client *http.Client) *PublishAPIUsingPUTParams {
	return &PublishAPIUsingPUTParams{
		HTTPClient: client,
	}
}

/*
PublishAPIUsingPUTParams contains all the parameters to send to the API endpoint

	for the publish Api using p u t operation.

	Typically these are written to a http.Request.
*/
type PublishAPIUsingPUTParams struct {

	/* APIUUID.

	   uuid
	*/
	APIUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the publish Api using p u t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PublishAPIUsingPUTParams) WithDefaults() *PublishAPIUsingPUTParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the publish Api using p u t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PublishAPIUsingPUTParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the publish Api using p u t params
func (o *PublishAPIUsingPUTParams) WithTimeout(timeout time.Duration) *PublishAPIUsingPUTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the publish Api using p u t params
func (o *PublishAPIUsingPUTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the publish Api using p u t params
func (o *PublishAPIUsingPUTParams) WithContext(ctx context.Context) *PublishAPIUsingPUTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the publish Api using p u t params
func (o *PublishAPIUsingPUTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the publish Api using p u t params
func (o *PublishAPIUsingPUTParams) WithHTTPClient(client *http.Client) *PublishAPIUsingPUTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the publish Api using p u t params
func (o *PublishAPIUsingPUTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIUUID adds the aPIUUID to the publish Api using p u t params
func (o *PublishAPIUsingPUTParams) WithAPIUUID(aPIUUID string) *PublishAPIUsingPUTParams {
	o.SetAPIUUID(aPIUUID)
	return o
}

// SetAPIUUID adds the apiUuid to the publish Api using p u t params
func (o *PublishAPIUsingPUTParams) SetAPIUUID(aPIUUID string) {
	o.APIUUID = aPIUUID
}

// WriteToRequest writes these params to a swagger request
func (o *PublishAPIUsingPUTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
