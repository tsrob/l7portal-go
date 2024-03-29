// Code generated by go-swagger; DO NOT EDIT.

package rate_limits_and_quotas

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

// NewCreateRateLimitQuotasParams creates a new CreateRateLimitQuotasParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateRateLimitQuotasParams() *CreateRateLimitQuotasParams {
	return &CreateRateLimitQuotasParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateRateLimitQuotasParamsWithTimeout creates a new CreateRateLimitQuotasParams object
// with the ability to set a timeout on a request.
func NewCreateRateLimitQuotasParamsWithTimeout(timeout time.Duration) *CreateRateLimitQuotasParams {
	return &CreateRateLimitQuotasParams{
		timeout: timeout,
	}
}

// NewCreateRateLimitQuotasParamsWithContext creates a new CreateRateLimitQuotasParams object
// with the ability to set a context for a request.
func NewCreateRateLimitQuotasParamsWithContext(ctx context.Context) *CreateRateLimitQuotasParams {
	return &CreateRateLimitQuotasParams{
		Context: ctx,
	}
}

// NewCreateRateLimitQuotasParamsWithHTTPClient creates a new CreateRateLimitQuotasParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateRateLimitQuotasParamsWithHTTPClient(client *http.Client) *CreateRateLimitQuotasParams {
	return &CreateRateLimitQuotasParams{
		HTTPClient: client,
	}
}

/*
CreateRateLimitQuotasParams contains all the parameters to send to the API endpoint

	for the create rate limit quotas operation.

	Typically these are written to a http.Request.
*/
type CreateRateLimitQuotasParams struct {

	/* RateLimitAndQuotaDto.

	   Create a new Rate Limit & Quota.
	*/
	RateLimitAndQuotaDto *models.RateLimitAndQuotaDto

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create rate limit quotas params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRateLimitQuotasParams) WithDefaults() *CreateRateLimitQuotasParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create rate limit quotas params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRateLimitQuotasParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create rate limit quotas params
func (o *CreateRateLimitQuotasParams) WithTimeout(timeout time.Duration) *CreateRateLimitQuotasParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create rate limit quotas params
func (o *CreateRateLimitQuotasParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create rate limit quotas params
func (o *CreateRateLimitQuotasParams) WithContext(ctx context.Context) *CreateRateLimitQuotasParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create rate limit quotas params
func (o *CreateRateLimitQuotasParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create rate limit quotas params
func (o *CreateRateLimitQuotasParams) WithHTTPClient(client *http.Client) *CreateRateLimitQuotasParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create rate limit quotas params
func (o *CreateRateLimitQuotasParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRateLimitAndQuotaDto adds the rateLimitAndQuotaDto to the create rate limit quotas params
func (o *CreateRateLimitQuotasParams) WithRateLimitAndQuotaDto(rateLimitAndQuotaDto *models.RateLimitAndQuotaDto) *CreateRateLimitQuotasParams {
	o.SetRateLimitAndQuotaDto(rateLimitAndQuotaDto)
	return o
}

// SetRateLimitAndQuotaDto adds the rateLimitAndQuotaDto to the create rate limit quotas params
func (o *CreateRateLimitQuotasParams) SetRateLimitAndQuotaDto(rateLimitAndQuotaDto *models.RateLimitAndQuotaDto) {
	o.RateLimitAndQuotaDto = rateLimitAndQuotaDto
}

// WriteToRequest writes these params to a swagger request
func (o *CreateRateLimitQuotasParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.RateLimitAndQuotaDto != nil {
		if err := r.SetBodyParam(o.RateLimitAndQuotaDto); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
