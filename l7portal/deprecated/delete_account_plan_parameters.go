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

// NewDeleteAccountPlanParams creates a new DeleteAccountPlanParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteAccountPlanParams() *DeleteAccountPlanParams {
	return &DeleteAccountPlanParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAccountPlanParamsWithTimeout creates a new DeleteAccountPlanParams object
// with the ability to set a timeout on a request.
func NewDeleteAccountPlanParamsWithTimeout(timeout time.Duration) *DeleteAccountPlanParams {
	return &DeleteAccountPlanParams{
		timeout: timeout,
	}
}

// NewDeleteAccountPlanParamsWithContext creates a new DeleteAccountPlanParams object
// with the ability to set a context for a request.
func NewDeleteAccountPlanParamsWithContext(ctx context.Context) *DeleteAccountPlanParams {
	return &DeleteAccountPlanParams{
		Context: ctx,
	}
}

// NewDeleteAccountPlanParamsWithHTTPClient creates a new DeleteAccountPlanParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteAccountPlanParamsWithHTTPClient(client *http.Client) *DeleteAccountPlanParams {
	return &DeleteAccountPlanParams{
		HTTPClient: client,
	}
}

/*
DeleteAccountPlanParams contains all the parameters to send to the API endpoint

	for the delete account plan operation.

	Typically these are written to a http.Request.
*/
type DeleteAccountPlanParams struct {

	/* UUID.

	   The UUID of the account plan to be deleted.
	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete account plan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAccountPlanParams) WithDefaults() *DeleteAccountPlanParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete account plan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAccountPlanParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete account plan params
func (o *DeleteAccountPlanParams) WithTimeout(timeout time.Duration) *DeleteAccountPlanParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete account plan params
func (o *DeleteAccountPlanParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete account plan params
func (o *DeleteAccountPlanParams) WithContext(ctx context.Context) *DeleteAccountPlanParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete account plan params
func (o *DeleteAccountPlanParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete account plan params
func (o *DeleteAccountPlanParams) WithHTTPClient(client *http.Client) *DeleteAccountPlanParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete account plan params
func (o *DeleteAccountPlanParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUUID adds the uuid to the delete account plan params
func (o *DeleteAccountPlanParams) WithUUID(uuid string) *DeleteAccountPlanParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the delete account plan params
func (o *DeleteAccountPlanParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAccountPlanParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
