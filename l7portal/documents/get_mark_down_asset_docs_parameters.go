// Code generated by go-swagger; DO NOT EDIT.

package documents

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

// NewGetMarkDownAssetDocsParams creates a new GetMarkDownAssetDocsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetMarkDownAssetDocsParams() *GetMarkDownAssetDocsParams {
	return &GetMarkDownAssetDocsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetMarkDownAssetDocsParamsWithTimeout creates a new GetMarkDownAssetDocsParams object
// with the ability to set a timeout on a request.
func NewGetMarkDownAssetDocsParamsWithTimeout(timeout time.Duration) *GetMarkDownAssetDocsParams {
	return &GetMarkDownAssetDocsParams{
		timeout: timeout,
	}
}

// NewGetMarkDownAssetDocsParamsWithContext creates a new GetMarkDownAssetDocsParams object
// with the ability to set a context for a request.
func NewGetMarkDownAssetDocsParamsWithContext(ctx context.Context) *GetMarkDownAssetDocsParams {
	return &GetMarkDownAssetDocsParams{
		Context: ctx,
	}
}

// NewGetMarkDownAssetDocsParamsWithHTTPClient creates a new GetMarkDownAssetDocsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetMarkDownAssetDocsParamsWithHTTPClient(client *http.Client) *GetMarkDownAssetDocsParams {
	return &GetMarkDownAssetDocsParams{
		HTTPClient: client,
	}
}

/*
GetMarkDownAssetDocsParams contains all the parameters to send to the API endpoint

	for the get mark down asset docs operation.

	Typically these are written to a http.Request.
*/
type GetMarkDownAssetDocsParams struct {

	/* Locale.

	   The locale for the document. For example, en-US, fr-FR, and fr-CA.
	*/
	Locale string

	/* Type.

	   The entity type associated to the document. For example, api and application.
	*/
	Type string

	/* TypeUUID.

	   The UUID of the entity associated with the document.
	*/
	TypeUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get mark down asset docs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMarkDownAssetDocsParams) WithDefaults() *GetMarkDownAssetDocsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get mark down asset docs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMarkDownAssetDocsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get mark down asset docs params
func (o *GetMarkDownAssetDocsParams) WithTimeout(timeout time.Duration) *GetMarkDownAssetDocsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get mark down asset docs params
func (o *GetMarkDownAssetDocsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get mark down asset docs params
func (o *GetMarkDownAssetDocsParams) WithContext(ctx context.Context) *GetMarkDownAssetDocsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get mark down asset docs params
func (o *GetMarkDownAssetDocsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get mark down asset docs params
func (o *GetMarkDownAssetDocsParams) WithHTTPClient(client *http.Client) *GetMarkDownAssetDocsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get mark down asset docs params
func (o *GetMarkDownAssetDocsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLocale adds the locale to the get mark down asset docs params
func (o *GetMarkDownAssetDocsParams) WithLocale(locale string) *GetMarkDownAssetDocsParams {
	o.SetLocale(locale)
	return o
}

// SetLocale adds the locale to the get mark down asset docs params
func (o *GetMarkDownAssetDocsParams) SetLocale(locale string) {
	o.Locale = locale
}

// WithType adds the typeVar to the get mark down asset docs params
func (o *GetMarkDownAssetDocsParams) WithType(typeVar string) *GetMarkDownAssetDocsParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the get mark down asset docs params
func (o *GetMarkDownAssetDocsParams) SetType(typeVar string) {
	o.Type = typeVar
}

// WithTypeUUID adds the typeUUID to the get mark down asset docs params
func (o *GetMarkDownAssetDocsParams) WithTypeUUID(typeUUID string) *GetMarkDownAssetDocsParams {
	o.SetTypeUUID(typeUUID)
	return o
}

// SetTypeUUID adds the typeUuid to the get mark down asset docs params
func (o *GetMarkDownAssetDocsParams) SetTypeUUID(typeUUID string) {
	o.TypeUUID = typeUUID
}

// WriteToRequest writes these params to a swagger request
func (o *GetMarkDownAssetDocsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param locale
	qrLocale := o.Locale
	qLocale := qrLocale
	if qLocale != "" {

		if err := r.SetQueryParam("locale", qLocale); err != nil {
			return err
		}
	}

	// path param type
	if err := r.SetPathParam("type", o.Type); err != nil {
		return err
	}

	// path param typeUuid
	if err := r.SetPathParam("typeUuid", o.TypeUUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
