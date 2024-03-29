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

// NewGetProxyOrganizationsUsingGETParams creates a new GetProxyOrganizationsUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetProxyOrganizationsUsingGETParams() *GetProxyOrganizationsUsingGETParams {
	return &GetProxyOrganizationsUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetProxyOrganizationsUsingGETParamsWithTimeout creates a new GetProxyOrganizationsUsingGETParams object
// with the ability to set a timeout on a request.
func NewGetProxyOrganizationsUsingGETParamsWithTimeout(timeout time.Duration) *GetProxyOrganizationsUsingGETParams {
	return &GetProxyOrganizationsUsingGETParams{
		timeout: timeout,
	}
}

// NewGetProxyOrganizationsUsingGETParamsWithContext creates a new GetProxyOrganizationsUsingGETParams object
// with the ability to set a context for a request.
func NewGetProxyOrganizationsUsingGETParamsWithContext(ctx context.Context) *GetProxyOrganizationsUsingGETParams {
	return &GetProxyOrganizationsUsingGETParams{
		Context: ctx,
	}
}

// NewGetProxyOrganizationsUsingGETParamsWithHTTPClient creates a new GetProxyOrganizationsUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetProxyOrganizationsUsingGETParamsWithHTTPClient(client *http.Client) *GetProxyOrganizationsUsingGETParams {
	return &GetProxyOrganizationsUsingGETParams{
		HTTPClient: client,
	}
}

/*
GetProxyOrganizationsUsingGETParams contains all the parameters to send to the API endpoint

	for the get proxy organizations using g e t operation.

	Typically these are written to a http.Request.
*/
type GetProxyOrganizationsUsingGETParams struct {

	/* UUID.

	   The proxy UUID to get.
	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get proxy organizations using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProxyOrganizationsUsingGETParams) WithDefaults() *GetProxyOrganizationsUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get proxy organizations using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProxyOrganizationsUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get proxy organizations using g e t params
func (o *GetProxyOrganizationsUsingGETParams) WithTimeout(timeout time.Duration) *GetProxyOrganizationsUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get proxy organizations using g e t params
func (o *GetProxyOrganizationsUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get proxy organizations using g e t params
func (o *GetProxyOrganizationsUsingGETParams) WithContext(ctx context.Context) *GetProxyOrganizationsUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get proxy organizations using g e t params
func (o *GetProxyOrganizationsUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get proxy organizations using g e t params
func (o *GetProxyOrganizationsUsingGETParams) WithHTTPClient(client *http.Client) *GetProxyOrganizationsUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get proxy organizations using g e t params
func (o *GetProxyOrganizationsUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUUID adds the uuid to the get proxy organizations using g e t params
func (o *GetProxyOrganizationsUsingGETParams) WithUUID(uuid string) *GetProxyOrganizationsUsingGETParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the get proxy organizations using g e t params
func (o *GetProxyOrganizationsUsingGETParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *GetProxyOrganizationsUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
