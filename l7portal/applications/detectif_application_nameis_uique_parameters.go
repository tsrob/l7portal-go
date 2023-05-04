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

// NewDetectifApplicationNameisUiqueParams creates a new DetectifApplicationNameisUiqueParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDetectifApplicationNameisUiqueParams() *DetectifApplicationNameisUiqueParams {
	return &DetectifApplicationNameisUiqueParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDetectifApplicationNameisUiqueParamsWithTimeout creates a new DetectifApplicationNameisUiqueParams object
// with the ability to set a timeout on a request.
func NewDetectifApplicationNameisUiqueParamsWithTimeout(timeout time.Duration) *DetectifApplicationNameisUiqueParams {
	return &DetectifApplicationNameisUiqueParams{
		timeout: timeout,
	}
}

// NewDetectifApplicationNameisUiqueParamsWithContext creates a new DetectifApplicationNameisUiqueParams object
// with the ability to set a context for a request.
func NewDetectifApplicationNameisUiqueParamsWithContext(ctx context.Context) *DetectifApplicationNameisUiqueParams {
	return &DetectifApplicationNameisUiqueParams{
		Context: ctx,
	}
}

// NewDetectifApplicationNameisUiqueParamsWithHTTPClient creates a new DetectifApplicationNameisUiqueParams object
// with the ability to set a custom HTTPClient for a request.
func NewDetectifApplicationNameisUiqueParamsWithHTTPClient(client *http.Client) *DetectifApplicationNameisUiqueParams {
	return &DetectifApplicationNameisUiqueParams{
		HTTPClient: client,
	}
}

/*
DetectifApplicationNameisUiqueParams contains all the parameters to send to the API endpoint

	for the detectif application nameis uique operation.

	Typically these are written to a http.Request.
*/
type DetectifApplicationNameisUiqueParams struct {

	/* AppUUID.

	   The Application's UUID.
	*/
	AppUUID *string

	/* Name.

	   The name to detect if is unique.
	*/
	Name string

	/* OrganizationUUID.

	   The Organization's UUID.
	*/
	OrganizationUUID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the detectif application nameis uique params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DetectifApplicationNameisUiqueParams) WithDefaults() *DetectifApplicationNameisUiqueParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the detectif application nameis uique params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DetectifApplicationNameisUiqueParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the detectif application nameis uique params
func (o *DetectifApplicationNameisUiqueParams) WithTimeout(timeout time.Duration) *DetectifApplicationNameisUiqueParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the detectif application nameis uique params
func (o *DetectifApplicationNameisUiqueParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the detectif application nameis uique params
func (o *DetectifApplicationNameisUiqueParams) WithContext(ctx context.Context) *DetectifApplicationNameisUiqueParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the detectif application nameis uique params
func (o *DetectifApplicationNameisUiqueParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the detectif application nameis uique params
func (o *DetectifApplicationNameisUiqueParams) WithHTTPClient(client *http.Client) *DetectifApplicationNameisUiqueParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the detectif application nameis uique params
func (o *DetectifApplicationNameisUiqueParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppUUID adds the appUUID to the detectif application nameis uique params
func (o *DetectifApplicationNameisUiqueParams) WithAppUUID(appUUID *string) *DetectifApplicationNameisUiqueParams {
	o.SetAppUUID(appUUID)
	return o
}

// SetAppUUID adds the appUuid to the detectif application nameis uique params
func (o *DetectifApplicationNameisUiqueParams) SetAppUUID(appUUID *string) {
	o.AppUUID = appUUID
}

// WithName adds the name to the detectif application nameis uique params
func (o *DetectifApplicationNameisUiqueParams) WithName(name string) *DetectifApplicationNameisUiqueParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the detectif application nameis uique params
func (o *DetectifApplicationNameisUiqueParams) SetName(name string) {
	o.Name = name
}

// WithOrganizationUUID adds the organizationUUID to the detectif application nameis uique params
func (o *DetectifApplicationNameisUiqueParams) WithOrganizationUUID(organizationUUID *string) *DetectifApplicationNameisUiqueParams {
	o.SetOrganizationUUID(organizationUUID)
	return o
}

// SetOrganizationUUID adds the organizationUuid to the detectif application nameis uique params
func (o *DetectifApplicationNameisUiqueParams) SetOrganizationUUID(organizationUUID *string) {
	o.OrganizationUUID = organizationUUID
}

// WriteToRequest writes these params to a swagger request
func (o *DetectifApplicationNameisUiqueParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AppUUID != nil {

		// query param appUuid
		var qrAppUUID string

		if o.AppUUID != nil {
			qrAppUUID = *o.AppUUID
		}
		qAppUUID := qrAppUUID
		if qAppUUID != "" {

			if err := r.SetQueryParam("appUuid", qAppUUID); err != nil {
				return err
			}
		}
	}

	// query param name
	qrName := o.Name
	qName := qrName
	if qName != "" {

		if err := r.SetQueryParam("name", qName); err != nil {
			return err
		}
	}

	if o.OrganizationUUID != nil {

		// query param organizationUuid
		var qrOrganizationUUID string

		if o.OrganizationUUID != nil {
			qrOrganizationUUID = *o.OrganizationUUID
		}
		qOrganizationUUID := qrOrganizationUUID
		if qOrganizationUUID != "" {

			if err := r.SetQueryParam("organizationUuid", qOrganizationUUID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
