// Code generated by go-swagger; DO NOT EDIT.

package requests

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

// NewReviewRequestParams creates a new ReviewRequestParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReviewRequestParams() *ReviewRequestParams {
	return &ReviewRequestParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReviewRequestParamsWithTimeout creates a new ReviewRequestParams object
// with the ability to set a timeout on a request.
func NewReviewRequestParamsWithTimeout(timeout time.Duration) *ReviewRequestParams {
	return &ReviewRequestParams{
		timeout: timeout,
	}
}

// NewReviewRequestParamsWithContext creates a new ReviewRequestParams object
// with the ability to set a context for a request.
func NewReviewRequestParamsWithContext(ctx context.Context) *ReviewRequestParams {
	return &ReviewRequestParams{
		Context: ctx,
	}
}

// NewReviewRequestParamsWithHTTPClient creates a new ReviewRequestParams object
// with the ability to set a custom HTTPClient for a request.
func NewReviewRequestParamsWithHTTPClient(client *http.Client) *ReviewRequestParams {
	return &ReviewRequestParams{
		HTTPClient: client,
	}
}

/*
ReviewRequestParams contains all the parameters to send to the API endpoint

	for the review request operation.

	Typically these are written to a http.Request.
*/
type ReviewRequestParams struct {

	/* Body.

	   Provides the information and action needed to review the Request.
	*/
	Body *models.RequestReview

	/* UUID.

	   The UUID of the request to return.
	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the review request params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReviewRequestParams) WithDefaults() *ReviewRequestParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the review request params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReviewRequestParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the review request params
func (o *ReviewRequestParams) WithTimeout(timeout time.Duration) *ReviewRequestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the review request params
func (o *ReviewRequestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the review request params
func (o *ReviewRequestParams) WithContext(ctx context.Context) *ReviewRequestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the review request params
func (o *ReviewRequestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the review request params
func (o *ReviewRequestParams) WithHTTPClient(client *http.Client) *ReviewRequestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the review request params
func (o *ReviewRequestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the review request params
func (o *ReviewRequestParams) WithBody(body *models.RequestReview) *ReviewRequestParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the review request params
func (o *ReviewRequestParams) SetBody(body *models.RequestReview) {
	o.Body = body
}

// WithUUID adds the uuid to the review request params
func (o *ReviewRequestParams) WithUUID(uuid string) *ReviewRequestParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the review request params
func (o *ReviewRequestParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *ReviewRequestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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