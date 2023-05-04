// Code generated by go-swagger; DO NOT EDIT.

package rate_limits_and_quotas

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

// NewDeleteRateLimitQuotasParams creates a new DeleteRateLimitQuotasParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteRateLimitQuotasParams() *DeleteRateLimitQuotasParams {
	return &DeleteRateLimitQuotasParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteRateLimitQuotasParamsWithTimeout creates a new DeleteRateLimitQuotasParams object
// with the ability to set a timeout on a request.
func NewDeleteRateLimitQuotasParamsWithTimeout(timeout time.Duration) *DeleteRateLimitQuotasParams {
	return &DeleteRateLimitQuotasParams{
		timeout: timeout,
	}
}

// NewDeleteRateLimitQuotasParamsWithContext creates a new DeleteRateLimitQuotasParams object
// with the ability to set a context for a request.
func NewDeleteRateLimitQuotasParamsWithContext(ctx context.Context) *DeleteRateLimitQuotasParams {
	return &DeleteRateLimitQuotasParams{
		Context: ctx,
	}
}

// NewDeleteRateLimitQuotasParamsWithHTTPClient creates a new DeleteRateLimitQuotasParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteRateLimitQuotasParamsWithHTTPClient(client *http.Client) *DeleteRateLimitQuotasParams {
	return &DeleteRateLimitQuotasParams{
		HTTPClient: client,
	}
}

/*
DeleteRateLimitQuotasParams contains all the parameters to send to the API endpoint

	for the delete rate limit quotas operation.

	Typically these are written to a http.Request.
*/
type DeleteRateLimitQuotasParams struct {

	/* UUID.

	   The UUID of the Rate Limit & Quota to delete.
	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete rate limit quotas params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteRateLimitQuotasParams) WithDefaults() *DeleteRateLimitQuotasParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete rate limit quotas params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteRateLimitQuotasParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete rate limit quotas params
func (o *DeleteRateLimitQuotasParams) WithTimeout(timeout time.Duration) *DeleteRateLimitQuotasParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete rate limit quotas params
func (o *DeleteRateLimitQuotasParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete rate limit quotas params
func (o *DeleteRateLimitQuotasParams) WithContext(ctx context.Context) *DeleteRateLimitQuotasParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete rate limit quotas params
func (o *DeleteRateLimitQuotasParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete rate limit quotas params
func (o *DeleteRateLimitQuotasParams) WithHTTPClient(client *http.Client) *DeleteRateLimitQuotasParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete rate limit quotas params
func (o *DeleteRateLimitQuotasParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUUID adds the uuid to the delete rate limit quotas params
func (o *DeleteRateLimitQuotasParams) WithUUID(uuid string) *DeleteRateLimitQuotasParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the delete rate limit quotas params
func (o *DeleteRateLimitQuotasParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteRateLimitQuotasParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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