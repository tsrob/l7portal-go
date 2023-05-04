// Code generated by go-swagger; DO NOT EDIT.

package organizations

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

// NewGetAssociatedOrgAPIRLQParams creates a new GetAssociatedOrgAPIRLQParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAssociatedOrgAPIRLQParams() *GetAssociatedOrgAPIRLQParams {
	return &GetAssociatedOrgAPIRLQParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAssociatedOrgAPIRLQParamsWithTimeout creates a new GetAssociatedOrgAPIRLQParams object
// with the ability to set a timeout on a request.
func NewGetAssociatedOrgAPIRLQParamsWithTimeout(timeout time.Duration) *GetAssociatedOrgAPIRLQParams {
	return &GetAssociatedOrgAPIRLQParams{
		timeout: timeout,
	}
}

// NewGetAssociatedOrgAPIRLQParamsWithContext creates a new GetAssociatedOrgAPIRLQParams object
// with the ability to set a context for a request.
func NewGetAssociatedOrgAPIRLQParamsWithContext(ctx context.Context) *GetAssociatedOrgAPIRLQParams {
	return &GetAssociatedOrgAPIRLQParams{
		Context: ctx,
	}
}

// NewGetAssociatedOrgAPIRLQParamsWithHTTPClient creates a new GetAssociatedOrgAPIRLQParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAssociatedOrgAPIRLQParamsWithHTTPClient(client *http.Client) *GetAssociatedOrgAPIRLQParams {
	return &GetAssociatedOrgAPIRLQParams{
		HTTPClient: client,
	}
}

/*
GetAssociatedOrgAPIRLQParams contains all the parameters to send to the API endpoint

	for the get associated org Api r l q operation.

	Typically these are written to a http.Request.
*/
type GetAssociatedOrgAPIRLQParams struct {

	/* OrgUUID.

	   Uuid of the Organization
	*/
	OrgUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get associated org Api r l q params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAssociatedOrgAPIRLQParams) WithDefaults() *GetAssociatedOrgAPIRLQParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get associated org Api r l q params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAssociatedOrgAPIRLQParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get associated org Api r l q params
func (o *GetAssociatedOrgAPIRLQParams) WithTimeout(timeout time.Duration) *GetAssociatedOrgAPIRLQParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get associated org Api r l q params
func (o *GetAssociatedOrgAPIRLQParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get associated org Api r l q params
func (o *GetAssociatedOrgAPIRLQParams) WithContext(ctx context.Context) *GetAssociatedOrgAPIRLQParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get associated org Api r l q params
func (o *GetAssociatedOrgAPIRLQParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get associated org Api r l q params
func (o *GetAssociatedOrgAPIRLQParams) WithHTTPClient(client *http.Client) *GetAssociatedOrgAPIRLQParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get associated org Api r l q params
func (o *GetAssociatedOrgAPIRLQParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrgUUID adds the orgUUID to the get associated org Api r l q params
func (o *GetAssociatedOrgAPIRLQParams) WithOrgUUID(orgUUID string) *GetAssociatedOrgAPIRLQParams {
	o.SetOrgUUID(orgUUID)
	return o
}

// SetOrgUUID adds the orgUuid to the get associated org Api r l q params
func (o *GetAssociatedOrgAPIRLQParams) SetOrgUUID(orgUUID string) {
	o.OrgUUID = orgUUID
}

// WriteToRequest writes these params to a swagger request
func (o *GetAssociatedOrgAPIRLQParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param orgUuid
	if err := r.SetPathParam("orgUuid", o.OrgUUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}