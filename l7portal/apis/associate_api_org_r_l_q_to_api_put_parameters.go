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

	"github.com/tsrob/l7portal-go/models"
)

// NewAssociateAPIOrgRLQToAPIPutParams creates a new AssociateAPIOrgRLQToAPIPutParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAssociateAPIOrgRLQToAPIPutParams() *AssociateAPIOrgRLQToAPIPutParams {
	return &AssociateAPIOrgRLQToAPIPutParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAssociateAPIOrgRLQToAPIPutParamsWithTimeout creates a new AssociateAPIOrgRLQToAPIPutParams object
// with the ability to set a timeout on a request.
func NewAssociateAPIOrgRLQToAPIPutParamsWithTimeout(timeout time.Duration) *AssociateAPIOrgRLQToAPIPutParams {
	return &AssociateAPIOrgRLQToAPIPutParams{
		timeout: timeout,
	}
}

// NewAssociateAPIOrgRLQToAPIPutParamsWithContext creates a new AssociateAPIOrgRLQToAPIPutParams object
// with the ability to set a context for a request.
func NewAssociateAPIOrgRLQToAPIPutParamsWithContext(ctx context.Context) *AssociateAPIOrgRLQToAPIPutParams {
	return &AssociateAPIOrgRLQToAPIPutParams{
		Context: ctx,
	}
}

// NewAssociateAPIOrgRLQToAPIPutParamsWithHTTPClient creates a new AssociateAPIOrgRLQToAPIPutParams object
// with the ability to set a custom HTTPClient for a request.
func NewAssociateAPIOrgRLQToAPIPutParamsWithHTTPClient(client *http.Client) *AssociateAPIOrgRLQToAPIPutParams {
	return &AssociateAPIOrgRLQToAPIPutParams{
		HTTPClient: client,
	}
}

/*
AssociateAPIOrgRLQToAPIPutParams contains all the parameters to send to the API endpoint

	for the associate Api org r l q to Api put operation.

	Typically these are written to a http.Request.
*/
type AssociateAPIOrgRLQToAPIPutParams struct {

	/* APIUUID.

	   apiUuid
	*/
	APIUUID string

	/* Body.

	   An Array of Organization and Rate Limit Quota(Api per Org) combination.
	*/
	Body *models.APIOrgRLQPut

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the associate Api org r l q to Api put params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AssociateAPIOrgRLQToAPIPutParams) WithDefaults() *AssociateAPIOrgRLQToAPIPutParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the associate Api org r l q to Api put params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AssociateAPIOrgRLQToAPIPutParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the associate Api org r l q to Api put params
func (o *AssociateAPIOrgRLQToAPIPutParams) WithTimeout(timeout time.Duration) *AssociateAPIOrgRLQToAPIPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the associate Api org r l q to Api put params
func (o *AssociateAPIOrgRLQToAPIPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the associate Api org r l q to Api put params
func (o *AssociateAPIOrgRLQToAPIPutParams) WithContext(ctx context.Context) *AssociateAPIOrgRLQToAPIPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the associate Api org r l q to Api put params
func (o *AssociateAPIOrgRLQToAPIPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the associate Api org r l q to Api put params
func (o *AssociateAPIOrgRLQToAPIPutParams) WithHTTPClient(client *http.Client) *AssociateAPIOrgRLQToAPIPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the associate Api org r l q to Api put params
func (o *AssociateAPIOrgRLQToAPIPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIUUID adds the aPIUUID to the associate Api org r l q to Api put params
func (o *AssociateAPIOrgRLQToAPIPutParams) WithAPIUUID(aPIUUID string) *AssociateAPIOrgRLQToAPIPutParams {
	o.SetAPIUUID(aPIUUID)
	return o
}

// SetAPIUUID adds the apiUuid to the associate Api org r l q to Api put params
func (o *AssociateAPIOrgRLQToAPIPutParams) SetAPIUUID(aPIUUID string) {
	o.APIUUID = aPIUUID
}

// WithBody adds the body to the associate Api org r l q to Api put params
func (o *AssociateAPIOrgRLQToAPIPutParams) WithBody(body *models.APIOrgRLQPut) *AssociateAPIOrgRLQToAPIPutParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the associate Api org r l q to Api put params
func (o *AssociateAPIOrgRLQToAPIPutParams) SetBody(body *models.APIOrgRLQPut) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AssociateAPIOrgRLQToAPIPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param apiUuid
	if err := r.SetPathParam("apiUuid", o.APIUUID); err != nil {
		return err
	}
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}