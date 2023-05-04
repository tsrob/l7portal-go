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

// NewUpdateAPITagsParams creates a new UpdateAPITagsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateAPITagsParams() *UpdateAPITagsParams {
	return &UpdateAPITagsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateAPITagsParamsWithTimeout creates a new UpdateAPITagsParams object
// with the ability to set a timeout on a request.
func NewUpdateAPITagsParamsWithTimeout(timeout time.Duration) *UpdateAPITagsParams {
	return &UpdateAPITagsParams{
		timeout: timeout,
	}
}

// NewUpdateAPITagsParamsWithContext creates a new UpdateAPITagsParams object
// with the ability to set a context for a request.
func NewUpdateAPITagsParamsWithContext(ctx context.Context) *UpdateAPITagsParams {
	return &UpdateAPITagsParams{
		Context: ctx,
	}
}

// NewUpdateAPITagsParamsWithHTTPClient creates a new UpdateAPITagsParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateAPITagsParamsWithHTTPClient(client *http.Client) *UpdateAPITagsParams {
	return &UpdateAPITagsParams{
		HTTPClient: client,
	}
}

/*
UpdateAPITagsParams contains all the parameters to send to the API endpoint

	for the update Api tags operation.

	Typically these are written to a http.Request.
*/
type UpdateAPITagsParams struct {

	/* Tag.

	   Tag name to be updated
	*/
	Tag models.TenantTagPut

	/* TagUUID.

	   tagUuid
	*/
	TagUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update Api tags params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateAPITagsParams) WithDefaults() *UpdateAPITagsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update Api tags params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateAPITagsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update Api tags params
func (o *UpdateAPITagsParams) WithTimeout(timeout time.Duration) *UpdateAPITagsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update Api tags params
func (o *UpdateAPITagsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update Api tags params
func (o *UpdateAPITagsParams) WithContext(ctx context.Context) *UpdateAPITagsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update Api tags params
func (o *UpdateAPITagsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update Api tags params
func (o *UpdateAPITagsParams) WithHTTPClient(client *http.Client) *UpdateAPITagsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update Api tags params
func (o *UpdateAPITagsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTag adds the tag to the update Api tags params
func (o *UpdateAPITagsParams) WithTag(tag models.TenantTagPut) *UpdateAPITagsParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the update Api tags params
func (o *UpdateAPITagsParams) SetTag(tag models.TenantTagPut) {
	o.Tag = tag
}

// WithTagUUID adds the tagUUID to the update Api tags params
func (o *UpdateAPITagsParams) WithTagUUID(tagUUID string) *UpdateAPITagsParams {
	o.SetTagUUID(tagUUID)
	return o
}

// SetTagUUID adds the tagUuid to the update Api tags params
func (o *UpdateAPITagsParams) SetTagUUID(tagUUID string) {
	o.TagUUID = tagUUID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateAPITagsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Tag != nil {
		if err := r.SetBodyParam(o.Tag); err != nil {
			return err
		}
	}

	// path param tagUuid
	if err := r.SetPathParam("tagUuid", o.TagUUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
