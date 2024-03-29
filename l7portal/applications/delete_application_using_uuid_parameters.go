// Code generated by go-swagger; DO NOT EDIT.

package applications

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

// NewDeleteApplicationUsingUUIDParams creates a new DeleteApplicationUsingUUIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteApplicationUsingUUIDParams() *DeleteApplicationUsingUUIDParams {
	return &DeleteApplicationUsingUUIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteApplicationUsingUUIDParamsWithTimeout creates a new DeleteApplicationUsingUUIDParams object
// with the ability to set a timeout on a request.
func NewDeleteApplicationUsingUUIDParamsWithTimeout(timeout time.Duration) *DeleteApplicationUsingUUIDParams {
	return &DeleteApplicationUsingUUIDParams{
		timeout: timeout,
	}
}

// NewDeleteApplicationUsingUUIDParamsWithContext creates a new DeleteApplicationUsingUUIDParams object
// with the ability to set a context for a request.
func NewDeleteApplicationUsingUUIDParamsWithContext(ctx context.Context) *DeleteApplicationUsingUUIDParams {
	return &DeleteApplicationUsingUUIDParams{
		Context: ctx,
	}
}

// NewDeleteApplicationUsingUUIDParamsWithHTTPClient creates a new DeleteApplicationUsingUUIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteApplicationUsingUUIDParamsWithHTTPClient(client *http.Client) *DeleteApplicationUsingUUIDParams {
	return &DeleteApplicationUsingUUIDParams{
		HTTPClient: client,
	}
}

/*
DeleteApplicationUsingUUIDParams contains all the parameters to send to the API endpoint

	for the delete application using Uuid operation.

	Typically these are written to a http.Request.
*/
type DeleteApplicationUsingUUIDParams struct {

	/* AppUUID.

	   The UUID for the Application to be delete.
	*/
	AppUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete application using Uuid params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteApplicationUsingUUIDParams) WithDefaults() *DeleteApplicationUsingUUIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete application using Uuid params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteApplicationUsingUUIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete application using Uuid params
func (o *DeleteApplicationUsingUUIDParams) WithTimeout(timeout time.Duration) *DeleteApplicationUsingUUIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete application using Uuid params
func (o *DeleteApplicationUsingUUIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete application using Uuid params
func (o *DeleteApplicationUsingUUIDParams) WithContext(ctx context.Context) *DeleteApplicationUsingUUIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete application using Uuid params
func (o *DeleteApplicationUsingUUIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete application using Uuid params
func (o *DeleteApplicationUsingUUIDParams) WithHTTPClient(client *http.Client) *DeleteApplicationUsingUUIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete application using Uuid params
func (o *DeleteApplicationUsingUUIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppUUID adds the appUUID to the delete application using Uuid params
func (o *DeleteApplicationUsingUUIDParams) WithAppUUID(appUUID string) *DeleteApplicationUsingUUIDParams {
	o.SetAppUUID(appUUID)
	return o
}

// SetAppUUID adds the appUuid to the delete application using Uuid params
func (o *DeleteApplicationUsingUUIDParams) SetAppUUID(appUUID string) {
	o.AppUUID = appUUID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteApplicationUsingUUIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param appUuid
	if err := r.SetPathParam("appUuid", o.AppUUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
