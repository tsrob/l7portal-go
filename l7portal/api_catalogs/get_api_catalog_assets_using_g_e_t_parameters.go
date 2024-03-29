// Code generated by go-swagger; DO NOT EDIT.

package api_catalogs

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

// NewGetAPICatalogAssetsUsingGETParams creates a new GetAPICatalogAssetsUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAPICatalogAssetsUsingGETParams() *GetAPICatalogAssetsUsingGETParams {
	return &GetAPICatalogAssetsUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPICatalogAssetsUsingGETParamsWithTimeout creates a new GetAPICatalogAssetsUsingGETParams object
// with the ability to set a timeout on a request.
func NewGetAPICatalogAssetsUsingGETParamsWithTimeout(timeout time.Duration) *GetAPICatalogAssetsUsingGETParams {
	return &GetAPICatalogAssetsUsingGETParams{
		timeout: timeout,
	}
}

// NewGetAPICatalogAssetsUsingGETParamsWithContext creates a new GetAPICatalogAssetsUsingGETParams object
// with the ability to set a context for a request.
func NewGetAPICatalogAssetsUsingGETParamsWithContext(ctx context.Context) *GetAPICatalogAssetsUsingGETParams {
	return &GetAPICatalogAssetsUsingGETParams{
		Context: ctx,
	}
}

// NewGetAPICatalogAssetsUsingGETParamsWithHTTPClient creates a new GetAPICatalogAssetsUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAPICatalogAssetsUsingGETParamsWithHTTPClient(client *http.Client) *GetAPICatalogAssetsUsingGETParams {
	return &GetAPICatalogAssetsUsingGETParams{
		HTTPClient: client,
	}
}

/*
GetAPICatalogAssetsUsingGETParams contains all the parameters to send to the API endpoint

	for the get Api catalog assets using g e t operation.

	Typically these are written to a http.Request.
*/
type GetAPICatalogAssetsUsingGETParams struct {

	/* APIUUID.

	   apiUuid
	*/
	APIUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get Api catalog assets using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPICatalogAssetsUsingGETParams) WithDefaults() *GetAPICatalogAssetsUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get Api catalog assets using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPICatalogAssetsUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get Api catalog assets using g e t params
func (o *GetAPICatalogAssetsUsingGETParams) WithTimeout(timeout time.Duration) *GetAPICatalogAssetsUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get Api catalog assets using g e t params
func (o *GetAPICatalogAssetsUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get Api catalog assets using g e t params
func (o *GetAPICatalogAssetsUsingGETParams) WithContext(ctx context.Context) *GetAPICatalogAssetsUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get Api catalog assets using g e t params
func (o *GetAPICatalogAssetsUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get Api catalog assets using g e t params
func (o *GetAPICatalogAssetsUsingGETParams) WithHTTPClient(client *http.Client) *GetAPICatalogAssetsUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get Api catalog assets using g e t params
func (o *GetAPICatalogAssetsUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIUUID adds the aPIUUID to the get Api catalog assets using g e t params
func (o *GetAPICatalogAssetsUsingGETParams) WithAPIUUID(aPIUUID string) *GetAPICatalogAssetsUsingGETParams {
	o.SetAPIUUID(aPIUUID)
	return o
}

// SetAPIUUID adds the apiUuid to the get Api catalog assets using g e t params
func (o *GetAPICatalogAssetsUsingGETParams) SetAPIUUID(aPIUUID string) {
	o.APIUUID = aPIUUID
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPICatalogAssetsUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
