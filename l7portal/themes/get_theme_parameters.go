// Code generated by go-swagger; DO NOT EDIT.

package themes

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

// NewGetThemeParams creates a new GetThemeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetThemeParams() *GetThemeParams {
	return &GetThemeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetThemeParamsWithTimeout creates a new GetThemeParams object
// with the ability to set a timeout on a request.
func NewGetThemeParamsWithTimeout(timeout time.Duration) *GetThemeParams {
	return &GetThemeParams{
		timeout: timeout,
	}
}

// NewGetThemeParamsWithContext creates a new GetThemeParams object
// with the ability to set a context for a request.
func NewGetThemeParamsWithContext(ctx context.Context) *GetThemeParams {
	return &GetThemeParams{
		Context: ctx,
	}
}

// NewGetThemeParamsWithHTTPClient creates a new GetThemeParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetThemeParamsWithHTTPClient(client *http.Client) *GetThemeParams {
	return &GetThemeParams{
		HTTPClient: client,
	}
}

/*
GetThemeParams contains all the parameters to send to the API endpoint

	for the get theme operation.

	Typically these are written to a http.Request.
*/
type GetThemeParams struct {

	/* APIMOrgUUID.

	   The UUID of the Organization to return
	*/
	APIMOrgUUID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get theme params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetThemeParams) WithDefaults() *GetThemeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get theme params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetThemeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get theme params
func (o *GetThemeParams) WithTimeout(timeout time.Duration) *GetThemeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get theme params
func (o *GetThemeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get theme params
func (o *GetThemeParams) WithContext(ctx context.Context) *GetThemeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get theme params
func (o *GetThemeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get theme params
func (o *GetThemeParams) WithHTTPClient(client *http.Client) *GetThemeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get theme params
func (o *GetThemeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIMOrgUUID adds the aPIMOrgUUID to the get theme params
func (o *GetThemeParams) WithAPIMOrgUUID(aPIMOrgUUID *string) *GetThemeParams {
	o.SetAPIMOrgUUID(aPIMOrgUUID)
	return o
}

// SetAPIMOrgUUID adds the apiMOrgUuid to the get theme params
func (o *GetThemeParams) SetAPIMOrgUUID(aPIMOrgUUID *string) {
	o.APIMOrgUUID = aPIMOrgUUID
}

// WriteToRequest writes these params to a swagger request
func (o *GetThemeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.APIMOrgUUID != nil {

		// header param APIM-OrgUuid
		if err := r.SetHeaderParam("APIM-OrgUuid", *o.APIMOrgUUID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
