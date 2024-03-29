// Code generated by go-swagger; DO NOT EDIT.

package api_permissions

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

// NewSaveUsingPUT1Params creates a new SaveUsingPUT1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSaveUsingPUT1Params() *SaveUsingPUT1Params {
	return &SaveUsingPUT1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewSaveUsingPUT1ParamsWithTimeout creates a new SaveUsingPUT1Params object
// with the ability to set a timeout on a request.
func NewSaveUsingPUT1ParamsWithTimeout(timeout time.Duration) *SaveUsingPUT1Params {
	return &SaveUsingPUT1Params{
		timeout: timeout,
	}
}

// NewSaveUsingPUT1ParamsWithContext creates a new SaveUsingPUT1Params object
// with the ability to set a context for a request.
func NewSaveUsingPUT1ParamsWithContext(ctx context.Context) *SaveUsingPUT1Params {
	return &SaveUsingPUT1Params{
		Context: ctx,
	}
}

// NewSaveUsingPUT1ParamsWithHTTPClient creates a new SaveUsingPUT1Params object
// with the ability to set a custom HTTPClient for a request.
func NewSaveUsingPUT1ParamsWithHTTPClient(client *http.Client) *SaveUsingPUT1Params {
	return &SaveUsingPUT1Params{
		HTTPClient: client,
	}
}

/*
SaveUsingPUT1Params contains all the parameters to send to the API endpoint

	for the save using p u t 1 operation.

	Typically these are written to a http.Request.
*/
type SaveUsingPUT1Params struct {

	/* APIUUID.

	   UUID of API
	*/
	APIUUID string

	/* UserUuids.

	   UUIDs of users to give permission to
	*/
	UserUuids []*models.ReferenceDto

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the save using p u t 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SaveUsingPUT1Params) WithDefaults() *SaveUsingPUT1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the save using p u t 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SaveUsingPUT1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the save using p u t 1 params
func (o *SaveUsingPUT1Params) WithTimeout(timeout time.Duration) *SaveUsingPUT1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the save using p u t 1 params
func (o *SaveUsingPUT1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the save using p u t 1 params
func (o *SaveUsingPUT1Params) WithContext(ctx context.Context) *SaveUsingPUT1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the save using p u t 1 params
func (o *SaveUsingPUT1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the save using p u t 1 params
func (o *SaveUsingPUT1Params) WithHTTPClient(client *http.Client) *SaveUsingPUT1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the save using p u t 1 params
func (o *SaveUsingPUT1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIUUID adds the aPIUUID to the save using p u t 1 params
func (o *SaveUsingPUT1Params) WithAPIUUID(aPIUUID string) *SaveUsingPUT1Params {
	o.SetAPIUUID(aPIUUID)
	return o
}

// SetAPIUUID adds the apiUuid to the save using p u t 1 params
func (o *SaveUsingPUT1Params) SetAPIUUID(aPIUUID string) {
	o.APIUUID = aPIUUID
}

// WithUserUuids adds the userUuids to the save using p u t 1 params
func (o *SaveUsingPUT1Params) WithUserUuids(userUuids []*models.ReferenceDto) *SaveUsingPUT1Params {
	o.SetUserUuids(userUuids)
	return o
}

// SetUserUuids adds the userUuids to the save using p u t 1 params
func (o *SaveUsingPUT1Params) SetUserUuids(userUuids []*models.ReferenceDto) {
	o.UserUuids = userUuids
}

// WriteToRequest writes these params to a swagger request
func (o *SaveUsingPUT1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param apiUuid
	if err := r.SetPathParam("apiUuid", o.APIUUID); err != nil {
		return err
	}
	if o.UserUuids != nil {
		if err := r.SetBodyParam(o.UserUuids); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
