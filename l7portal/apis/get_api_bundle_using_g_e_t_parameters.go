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

// NewGetAPIBundleUsingGETParams creates a new GetAPIBundleUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAPIBundleUsingGETParams() *GetAPIBundleUsingGETParams {
	return &GetAPIBundleUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPIBundleUsingGETParamsWithTimeout creates a new GetAPIBundleUsingGETParams object
// with the ability to set a timeout on a request.
func NewGetAPIBundleUsingGETParamsWithTimeout(timeout time.Duration) *GetAPIBundleUsingGETParams {
	return &GetAPIBundleUsingGETParams{
		timeout: timeout,
	}
}

// NewGetAPIBundleUsingGETParamsWithContext creates a new GetAPIBundleUsingGETParams object
// with the ability to set a context for a request.
func NewGetAPIBundleUsingGETParamsWithContext(ctx context.Context) *GetAPIBundleUsingGETParams {
	return &GetAPIBundleUsingGETParams{
		Context: ctx,
	}
}

// NewGetAPIBundleUsingGETParamsWithHTTPClient creates a new GetAPIBundleUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAPIBundleUsingGETParamsWithHTTPClient(client *http.Client) *GetAPIBundleUsingGETParams {
	return &GetAPIBundleUsingGETParams{
		HTTPClient: client,
	}
}

/*
GetAPIBundleUsingGETParams contains all the parameters to send to the API endpoint

	for the get Api bundle using g e t operation.

	Typically these are written to a http.Request.
*/
type GetAPIBundleUsingGETParams struct {

	/* APIUUID.

	   The API UUID.
	*/
	APIUUID string

	/* Type.

	   The type of bundle to retrieve. Allowable values: ['delete', 'create']
	*/
	Type *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get Api bundle using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIBundleUsingGETParams) WithDefaults() *GetAPIBundleUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get Api bundle using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIBundleUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get Api bundle using g e t params
func (o *GetAPIBundleUsingGETParams) WithTimeout(timeout time.Duration) *GetAPIBundleUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get Api bundle using g e t params
func (o *GetAPIBundleUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get Api bundle using g e t params
func (o *GetAPIBundleUsingGETParams) WithContext(ctx context.Context) *GetAPIBundleUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get Api bundle using g e t params
func (o *GetAPIBundleUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get Api bundle using g e t params
func (o *GetAPIBundleUsingGETParams) WithHTTPClient(client *http.Client) *GetAPIBundleUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get Api bundle using g e t params
func (o *GetAPIBundleUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIUUID adds the aPIUUID to the get Api bundle using g e t params
func (o *GetAPIBundleUsingGETParams) WithAPIUUID(aPIUUID string) *GetAPIBundleUsingGETParams {
	o.SetAPIUUID(aPIUUID)
	return o
}

// SetAPIUUID adds the apiUuid to the get Api bundle using g e t params
func (o *GetAPIBundleUsingGETParams) SetAPIUUID(aPIUUID string) {
	o.APIUUID = aPIUUID
}

// WithType adds the typeVar to the get Api bundle using g e t params
func (o *GetAPIBundleUsingGETParams) WithType(typeVar *string) *GetAPIBundleUsingGETParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the get Api bundle using g e t params
func (o *GetAPIBundleUsingGETParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPIBundleUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param apiUuid
	if err := r.SetPathParam("apiUuid", o.APIUUID); err != nil {
		return err
	}

	if o.Type != nil {

		// query param type
		var qrType string

		if o.Type != nil {
			qrType = *o.Type
		}
		qType := qrType
		if qType != "" {

			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
