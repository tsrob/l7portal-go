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
)

// NewDeleteAPITagParams creates a new DeleteAPITagParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteAPITagParams() *DeleteAPITagParams {
	return &DeleteAPITagParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAPITagParamsWithTimeout creates a new DeleteAPITagParams object
// with the ability to set a timeout on a request.
func NewDeleteAPITagParamsWithTimeout(timeout time.Duration) *DeleteAPITagParams {
	return &DeleteAPITagParams{
		timeout: timeout,
	}
}

// NewDeleteAPITagParamsWithContext creates a new DeleteAPITagParams object
// with the ability to set a context for a request.
func NewDeleteAPITagParamsWithContext(ctx context.Context) *DeleteAPITagParams {
	return &DeleteAPITagParams{
		Context: ctx,
	}
}

// NewDeleteAPITagParamsWithHTTPClient creates a new DeleteAPITagParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteAPITagParamsWithHTTPClient(client *http.Client) *DeleteAPITagParams {
	return &DeleteAPITagParams{
		HTTPClient: client,
	}
}

/*
DeleteAPITagParams contains all the parameters to send to the API endpoint

	for the delete Api tag operation.

	Typically these are written to a http.Request.
*/
type DeleteAPITagParams struct {

	/* TagUUID.

	   tagUuid
	*/
	TagUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete Api tag params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAPITagParams) WithDefaults() *DeleteAPITagParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete Api tag params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAPITagParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete Api tag params
func (o *DeleteAPITagParams) WithTimeout(timeout time.Duration) *DeleteAPITagParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete Api tag params
func (o *DeleteAPITagParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete Api tag params
func (o *DeleteAPITagParams) WithContext(ctx context.Context) *DeleteAPITagParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete Api tag params
func (o *DeleteAPITagParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete Api tag params
func (o *DeleteAPITagParams) WithHTTPClient(client *http.Client) *DeleteAPITagParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete Api tag params
func (o *DeleteAPITagParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTagUUID adds the tagUUID to the delete Api tag params
func (o *DeleteAPITagParams) WithTagUUID(tagUUID string) *DeleteAPITagParams {
	o.SetTagUUID(tagUUID)
	return o
}

// SetTagUUID adds the tagUuid to the delete Api tag params
func (o *DeleteAPITagParams) SetTagUUID(tagUUID string) {
	o.TagUUID = tagUUID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAPITagParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param tagUuid
	if err := r.SetPathParam("tagUuid", o.TagUUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
