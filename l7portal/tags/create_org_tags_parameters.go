// Code generated by go-swagger; DO NOT EDIT.

package tags

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

// NewCreateOrgTagsParams creates a new CreateOrgTagsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateOrgTagsParams() *CreateOrgTagsParams {
	return &CreateOrgTagsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateOrgTagsParamsWithTimeout creates a new CreateOrgTagsParams object
// with the ability to set a timeout on a request.
func NewCreateOrgTagsParamsWithTimeout(timeout time.Duration) *CreateOrgTagsParams {
	return &CreateOrgTagsParams{
		timeout: timeout,
	}
}

// NewCreateOrgTagsParamsWithContext creates a new CreateOrgTagsParams object
// with the ability to set a context for a request.
func NewCreateOrgTagsParamsWithContext(ctx context.Context) *CreateOrgTagsParams {
	return &CreateOrgTagsParams{
		Context: ctx,
	}
}

// NewCreateOrgTagsParamsWithHTTPClient creates a new CreateOrgTagsParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateOrgTagsParamsWithHTTPClient(client *http.Client) *CreateOrgTagsParams {
	return &CreateOrgTagsParams{
		HTTPClient: client,
	}
}

/*
CreateOrgTagsParams contains all the parameters to send to the API endpoint

	for the create org tags operation.

	Typically these are written to a http.Request.
*/
type CreateOrgTagsParams struct {

	/* TagNames.

	   List of tag names
	*/
	TagNames models.TenantTagsPost

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create org tags params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateOrgTagsParams) WithDefaults() *CreateOrgTagsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create org tags params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateOrgTagsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create org tags params
func (o *CreateOrgTagsParams) WithTimeout(timeout time.Duration) *CreateOrgTagsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create org tags params
func (o *CreateOrgTagsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create org tags params
func (o *CreateOrgTagsParams) WithContext(ctx context.Context) *CreateOrgTagsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create org tags params
func (o *CreateOrgTagsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create org tags params
func (o *CreateOrgTagsParams) WithHTTPClient(client *http.Client) *CreateOrgTagsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create org tags params
func (o *CreateOrgTagsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTagNames adds the tagNames to the create org tags params
func (o *CreateOrgTagsParams) WithTagNames(tagNames models.TenantTagsPost) *CreateOrgTagsParams {
	o.SetTagNames(tagNames)
	return o
}

// SetTagNames adds the tagNames to the create org tags params
func (o *CreateOrgTagsParams) SetTagNames(tagNames models.TenantTagsPost) {
	o.TagNames = tagNames
}

// WriteToRequest writes these params to a swagger request
func (o *CreateOrgTagsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.TagNames != nil {
		if err := r.SetBodyParam(o.TagNames); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
