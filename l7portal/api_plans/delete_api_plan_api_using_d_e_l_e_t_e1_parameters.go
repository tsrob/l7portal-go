// Code generated by go-swagger; DO NOT EDIT.

package api_plans

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

// NewDeleteAPIPlanAPIUsingDELETE1Params creates a new DeleteAPIPlanAPIUsingDELETE1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteAPIPlanAPIUsingDELETE1Params() *DeleteAPIPlanAPIUsingDELETE1Params {
	return &DeleteAPIPlanAPIUsingDELETE1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAPIPlanAPIUsingDELETE1ParamsWithTimeout creates a new DeleteAPIPlanAPIUsingDELETE1Params object
// with the ability to set a timeout on a request.
func NewDeleteAPIPlanAPIUsingDELETE1ParamsWithTimeout(timeout time.Duration) *DeleteAPIPlanAPIUsingDELETE1Params {
	return &DeleteAPIPlanAPIUsingDELETE1Params{
		timeout: timeout,
	}
}

// NewDeleteAPIPlanAPIUsingDELETE1ParamsWithContext creates a new DeleteAPIPlanAPIUsingDELETE1Params object
// with the ability to set a context for a request.
func NewDeleteAPIPlanAPIUsingDELETE1ParamsWithContext(ctx context.Context) *DeleteAPIPlanAPIUsingDELETE1Params {
	return &DeleteAPIPlanAPIUsingDELETE1Params{
		Context: ctx,
	}
}

// NewDeleteAPIPlanAPIUsingDELETE1ParamsWithHTTPClient creates a new DeleteAPIPlanAPIUsingDELETE1Params object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteAPIPlanAPIUsingDELETE1ParamsWithHTTPClient(client *http.Client) *DeleteAPIPlanAPIUsingDELETE1Params {
	return &DeleteAPIPlanAPIUsingDELETE1Params{
		HTTPClient: client,
	}
}

/*
DeleteAPIPlanAPIUsingDELETE1Params contains all the parameters to send to the API endpoint

	for the delete Api plan Api using d e l e t e 1 operation.

	Typically these are written to a http.Request.
*/
type DeleteAPIPlanAPIUsingDELETE1Params struct {

	/* APIUUID.

	   The UUID of the API
	*/
	APIUUID string

	/* PlanUUID.

	   The UUID of the API Plan
	*/
	PlanUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete Api plan Api using d e l e t e 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAPIPlanAPIUsingDELETE1Params) WithDefaults() *DeleteAPIPlanAPIUsingDELETE1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete Api plan Api using d e l e t e 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAPIPlanAPIUsingDELETE1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete Api plan Api using d e l e t e 1 params
func (o *DeleteAPIPlanAPIUsingDELETE1Params) WithTimeout(timeout time.Duration) *DeleteAPIPlanAPIUsingDELETE1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete Api plan Api using d e l e t e 1 params
func (o *DeleteAPIPlanAPIUsingDELETE1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete Api plan Api using d e l e t e 1 params
func (o *DeleteAPIPlanAPIUsingDELETE1Params) WithContext(ctx context.Context) *DeleteAPIPlanAPIUsingDELETE1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete Api plan Api using d e l e t e 1 params
func (o *DeleteAPIPlanAPIUsingDELETE1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete Api plan Api using d e l e t e 1 params
func (o *DeleteAPIPlanAPIUsingDELETE1Params) WithHTTPClient(client *http.Client) *DeleteAPIPlanAPIUsingDELETE1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete Api plan Api using d e l e t e 1 params
func (o *DeleteAPIPlanAPIUsingDELETE1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIUUID adds the aPIUUID to the delete Api plan Api using d e l e t e 1 params
func (o *DeleteAPIPlanAPIUsingDELETE1Params) WithAPIUUID(aPIUUID string) *DeleteAPIPlanAPIUsingDELETE1Params {
	o.SetAPIUUID(aPIUUID)
	return o
}

// SetAPIUUID adds the apiUuid to the delete Api plan Api using d e l e t e 1 params
func (o *DeleteAPIPlanAPIUsingDELETE1Params) SetAPIUUID(aPIUUID string) {
	o.APIUUID = aPIUUID
}

// WithPlanUUID adds the planUUID to the delete Api plan Api using d e l e t e 1 params
func (o *DeleteAPIPlanAPIUsingDELETE1Params) WithPlanUUID(planUUID string) *DeleteAPIPlanAPIUsingDELETE1Params {
	o.SetPlanUUID(planUUID)
	return o
}

// SetPlanUUID adds the planUuid to the delete Api plan Api using d e l e t e 1 params
func (o *DeleteAPIPlanAPIUsingDELETE1Params) SetPlanUUID(planUUID string) {
	o.PlanUUID = planUUID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAPIPlanAPIUsingDELETE1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param api-uuid
	if err := r.SetPathParam("api-uuid", o.APIUUID); err != nil {
		return err
	}

	// path param plan-uuid
	if err := r.SetPathParam("plan-uuid", o.PlanUUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}