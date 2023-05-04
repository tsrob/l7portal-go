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

// NewGetWsdlWithContentUsingGETParams creates a new GetWsdlWithContentUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetWsdlWithContentUsingGETParams() *GetWsdlWithContentUsingGETParams {
	return &GetWsdlWithContentUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetWsdlWithContentUsingGETParamsWithTimeout creates a new GetWsdlWithContentUsingGETParams object
// with the ability to set a timeout on a request.
func NewGetWsdlWithContentUsingGETParamsWithTimeout(timeout time.Duration) *GetWsdlWithContentUsingGETParams {
	return &GetWsdlWithContentUsingGETParams{
		timeout: timeout,
	}
}

// NewGetWsdlWithContentUsingGETParamsWithContext creates a new GetWsdlWithContentUsingGETParams object
// with the ability to set a context for a request.
func NewGetWsdlWithContentUsingGETParamsWithContext(ctx context.Context) *GetWsdlWithContentUsingGETParams {
	return &GetWsdlWithContentUsingGETParams{
		Context: ctx,
	}
}

// NewGetWsdlWithContentUsingGETParamsWithHTTPClient creates a new GetWsdlWithContentUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetWsdlWithContentUsingGETParamsWithHTTPClient(client *http.Client) *GetWsdlWithContentUsingGETParams {
	return &GetWsdlWithContentUsingGETParams{
		HTTPClient: client,
	}
}

/*
GetWsdlWithContentUsingGETParams contains all the parameters to send to the API endpoint

	for the get wsdl with content using g e t operation.

	Typically these are written to a http.Request.
*/
type GetWsdlWithContentUsingGETParams struct {

	/* APIUUID.

	   apiUuid
	*/
	APIUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get wsdl with content using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWsdlWithContentUsingGETParams) WithDefaults() *GetWsdlWithContentUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get wsdl with content using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWsdlWithContentUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get wsdl with content using g e t params
func (o *GetWsdlWithContentUsingGETParams) WithTimeout(timeout time.Duration) *GetWsdlWithContentUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get wsdl with content using g e t params
func (o *GetWsdlWithContentUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get wsdl with content using g e t params
func (o *GetWsdlWithContentUsingGETParams) WithContext(ctx context.Context) *GetWsdlWithContentUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get wsdl with content using g e t params
func (o *GetWsdlWithContentUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get wsdl with content using g e t params
func (o *GetWsdlWithContentUsingGETParams) WithHTTPClient(client *http.Client) *GetWsdlWithContentUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get wsdl with content using g e t params
func (o *GetWsdlWithContentUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIUUID adds the aPIUUID to the get wsdl with content using g e t params
func (o *GetWsdlWithContentUsingGETParams) WithAPIUUID(aPIUUID string) *GetWsdlWithContentUsingGETParams {
	o.SetAPIUUID(aPIUUID)
	return o
}

// SetAPIUUID adds the apiUuid to the get wsdl with content using g e t params
func (o *GetWsdlWithContentUsingGETParams) SetAPIUUID(aPIUUID string) {
	o.APIUUID = aPIUUID
}

// WriteToRequest writes these params to a swagger request
func (o *GetWsdlWithContentUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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