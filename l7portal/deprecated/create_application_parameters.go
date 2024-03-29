// Code generated by go-swagger; DO NOT EDIT.

package deprecated

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

// NewCreateApplicationParams creates a new CreateApplicationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateApplicationParams() *CreateApplicationParams {
	return &CreateApplicationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateApplicationParamsWithTimeout creates a new CreateApplicationParams object
// with the ability to set a timeout on a request.
func NewCreateApplicationParamsWithTimeout(timeout time.Duration) *CreateApplicationParams {
	return &CreateApplicationParams{
		timeout: timeout,
	}
}

// NewCreateApplicationParamsWithContext creates a new CreateApplicationParams object
// with the ability to set a context for a request.
func NewCreateApplicationParamsWithContext(ctx context.Context) *CreateApplicationParams {
	return &CreateApplicationParams{
		Context: ctx,
	}
}

// NewCreateApplicationParamsWithHTTPClient creates a new CreateApplicationParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateApplicationParamsWithHTTPClient(client *http.Client) *CreateApplicationParams {
	return &CreateApplicationParams{
		HTTPClient: client,
	}
}

/*
CreateApplicationParams contains all the parameters to send to the API endpoint

	for the create application operation.

	Typically these are written to a http.Request.
*/
type CreateApplicationParams struct {

	/* Body.

	   Creates an Application.
	*/
	Body *models.ApplicationCreate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create application params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateApplicationParams) WithDefaults() *CreateApplicationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create application params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateApplicationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create application params
func (o *CreateApplicationParams) WithTimeout(timeout time.Duration) *CreateApplicationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create application params
func (o *CreateApplicationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create application params
func (o *CreateApplicationParams) WithContext(ctx context.Context) *CreateApplicationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create application params
func (o *CreateApplicationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create application params
func (o *CreateApplicationParams) WithHTTPClient(client *http.Client) *CreateApplicationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create application params
func (o *CreateApplicationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create application params
func (o *CreateApplicationParams) WithBody(body *models.ApplicationCreate) *CreateApplicationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create application params
func (o *CreateApplicationParams) SetBody(body *models.ApplicationCreate) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateApplicationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
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
