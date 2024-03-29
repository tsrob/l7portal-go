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

	"github.com/tsrob/l7portal-go/models"
)

// NewUpdateProxyOrganizationUsingPUTParams creates a new UpdateProxyOrganizationUsingPUTParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateProxyOrganizationUsingPUTParams() *UpdateProxyOrganizationUsingPUTParams {
	return &UpdateProxyOrganizationUsingPUTParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateProxyOrganizationUsingPUTParamsWithTimeout creates a new UpdateProxyOrganizationUsingPUTParams object
// with the ability to set a timeout on a request.
func NewUpdateProxyOrganizationUsingPUTParamsWithTimeout(timeout time.Duration) *UpdateProxyOrganizationUsingPUTParams {
	return &UpdateProxyOrganizationUsingPUTParams{
		timeout: timeout,
	}
}

// NewUpdateProxyOrganizationUsingPUTParamsWithContext creates a new UpdateProxyOrganizationUsingPUTParams object
// with the ability to set a context for a request.
func NewUpdateProxyOrganizationUsingPUTParamsWithContext(ctx context.Context) *UpdateProxyOrganizationUsingPUTParams {
	return &UpdateProxyOrganizationUsingPUTParams{
		Context: ctx,
	}
}

// NewUpdateProxyOrganizationUsingPUTParamsWithHTTPClient creates a new UpdateProxyOrganizationUsingPUTParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateProxyOrganizationUsingPUTParamsWithHTTPClient(client *http.Client) *UpdateProxyOrganizationUsingPUTParams {
	return &UpdateProxyOrganizationUsingPUTParams{
		HTTPClient: client,
	}
}

/*
UpdateProxyOrganizationUsingPUTParams contains all the parameters to send to the API endpoint

	for the update proxy organization using p u t operation.

	Typically these are written to a http.Request.
*/
type UpdateProxyOrganizationUsingPUTParams struct {

	/* Body.

	   The updated proxy organization assignment.
	*/
	Body []*models.ReferenceDto

	/* UUID.

	   The proxy UUID to update.
	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update proxy organization using p u t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateProxyOrganizationUsingPUTParams) WithDefaults() *UpdateProxyOrganizationUsingPUTParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update proxy organization using p u t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateProxyOrganizationUsingPUTParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update proxy organization using p u t params
func (o *UpdateProxyOrganizationUsingPUTParams) WithTimeout(timeout time.Duration) *UpdateProxyOrganizationUsingPUTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update proxy organization using p u t params
func (o *UpdateProxyOrganizationUsingPUTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update proxy organization using p u t params
func (o *UpdateProxyOrganizationUsingPUTParams) WithContext(ctx context.Context) *UpdateProxyOrganizationUsingPUTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update proxy organization using p u t params
func (o *UpdateProxyOrganizationUsingPUTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update proxy organization using p u t params
func (o *UpdateProxyOrganizationUsingPUTParams) WithHTTPClient(client *http.Client) *UpdateProxyOrganizationUsingPUTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update proxy organization using p u t params
func (o *UpdateProxyOrganizationUsingPUTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update proxy organization using p u t params
func (o *UpdateProxyOrganizationUsingPUTParams) WithBody(body []*models.ReferenceDto) *UpdateProxyOrganizationUsingPUTParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update proxy organization using p u t params
func (o *UpdateProxyOrganizationUsingPUTParams) SetBody(body []*models.ReferenceDto) {
	o.Body = body
}

// WithUUID adds the uuid to the update proxy organization using p u t params
func (o *UpdateProxyOrganizationUsingPUTParams) WithUUID(uuid string) *UpdateProxyOrganizationUsingPUTParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the update proxy organization using p u t params
func (o *UpdateProxyOrganizationUsingPUTParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateProxyOrganizationUsingPUTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
