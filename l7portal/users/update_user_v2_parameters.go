// Code generated by go-swagger; DO NOT EDIT.

package users

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

// NewUpdateUserV2Params creates a new UpdateUserV2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateUserV2Params() *UpdateUserV2Params {
	return &UpdateUserV2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateUserV2ParamsWithTimeout creates a new UpdateUserV2Params object
// with the ability to set a timeout on a request.
func NewUpdateUserV2ParamsWithTimeout(timeout time.Duration) *UpdateUserV2Params {
	return &UpdateUserV2Params{
		timeout: timeout,
	}
}

// NewUpdateUserV2ParamsWithContext creates a new UpdateUserV2Params object
// with the ability to set a context for a request.
func NewUpdateUserV2ParamsWithContext(ctx context.Context) *UpdateUserV2Params {
	return &UpdateUserV2Params{
		Context: ctx,
	}
}

// NewUpdateUserV2ParamsWithHTTPClient creates a new UpdateUserV2Params object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateUserV2ParamsWithHTTPClient(client *http.Client) *UpdateUserV2Params {
	return &UpdateUserV2Params{
		HTTPClient: client,
	}
}

/*
UpdateUserV2Params contains all the parameters to send to the API endpoint

	for the update user v2 operation.

	Typically these are written to a http.Request.
*/
type UpdateUserV2Params struct {

	/* Body.

	   Provides the values needed to set up the User object.
	*/
	Body *models.UserPutV2

	/* UUID.

	   The UUID of the User to update.
	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update user v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateUserV2Params) WithDefaults() *UpdateUserV2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update user v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateUserV2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update user v2 params
func (o *UpdateUserV2Params) WithTimeout(timeout time.Duration) *UpdateUserV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update user v2 params
func (o *UpdateUserV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update user v2 params
func (o *UpdateUserV2Params) WithContext(ctx context.Context) *UpdateUserV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update user v2 params
func (o *UpdateUserV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update user v2 params
func (o *UpdateUserV2Params) WithHTTPClient(client *http.Client) *UpdateUserV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update user v2 params
func (o *UpdateUserV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update user v2 params
func (o *UpdateUserV2Params) WithBody(body *models.UserPutV2) *UpdateUserV2Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update user v2 params
func (o *UpdateUserV2Params) SetBody(body *models.UserPutV2) {
	o.Body = body
}

// WithUUID adds the uuid to the update user v2 params
func (o *UpdateUserV2Params) WithUUID(uuid string) *UpdateUserV2Params {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the update user v2 params
func (o *UpdateUserV2Params) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateUserV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}