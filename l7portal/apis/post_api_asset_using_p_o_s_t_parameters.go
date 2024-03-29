// Code generated by go-swagger; DO NOT EDIT.

package apis

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

// NewPostAPIAssetUsingPOSTParams creates a new PostAPIAssetUsingPOSTParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostAPIAssetUsingPOSTParams() *PostAPIAssetUsingPOSTParams {
	return &PostAPIAssetUsingPOSTParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostAPIAssetUsingPOSTParamsWithTimeout creates a new PostAPIAssetUsingPOSTParams object
// with the ability to set a timeout on a request.
func NewPostAPIAssetUsingPOSTParamsWithTimeout(timeout time.Duration) *PostAPIAssetUsingPOSTParams {
	return &PostAPIAssetUsingPOSTParams{
		timeout: timeout,
	}
}

// NewPostAPIAssetUsingPOSTParamsWithContext creates a new PostAPIAssetUsingPOSTParams object
// with the ability to set a context for a request.
func NewPostAPIAssetUsingPOSTParamsWithContext(ctx context.Context) *PostAPIAssetUsingPOSTParams {
	return &PostAPIAssetUsingPOSTParams{
		Context: ctx,
	}
}

// NewPostAPIAssetUsingPOSTParamsWithHTTPClient creates a new PostAPIAssetUsingPOSTParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostAPIAssetUsingPOSTParamsWithHTTPClient(client *http.Client) *PostAPIAssetUsingPOSTParams {
	return &PostAPIAssetUsingPOSTParams{
		HTTPClient: client,
	}
}

/*
PostAPIAssetUsingPOSTParams contains all the parameters to send to the API endpoint

	for the post Api asset using p o s t operation.

	Typically these are written to a http.Request.
*/
type PostAPIAssetUsingPOSTParams struct {

	/* APIUUID.

	   apiUuid
	*/
	APIUUID string

	/* Files.

	   The Asset to upload. A REST API can have an optionally specified WADL, or Swagger JSON file. A SOAP API must have a WSDL specified for it as well as any optional XSDs.
	*/
	Files runtime.NamedReadCloser

	/* FilesToDelete.

	   The UUID of file to delete. Multiple UUIDs can be provided in separate parts.
	*/
	FilesToDelete *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post Api asset using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAPIAssetUsingPOSTParams) WithDefaults() *PostAPIAssetUsingPOSTParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post Api asset using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAPIAssetUsingPOSTParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post Api asset using p o s t params
func (o *PostAPIAssetUsingPOSTParams) WithTimeout(timeout time.Duration) *PostAPIAssetUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post Api asset using p o s t params
func (o *PostAPIAssetUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post Api asset using p o s t params
func (o *PostAPIAssetUsingPOSTParams) WithContext(ctx context.Context) *PostAPIAssetUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post Api asset using p o s t params
func (o *PostAPIAssetUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post Api asset using p o s t params
func (o *PostAPIAssetUsingPOSTParams) WithHTTPClient(client *http.Client) *PostAPIAssetUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post Api asset using p o s t params
func (o *PostAPIAssetUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIUUID adds the aPIUUID to the post Api asset using p o s t params
func (o *PostAPIAssetUsingPOSTParams) WithAPIUUID(aPIUUID string) *PostAPIAssetUsingPOSTParams {
	o.SetAPIUUID(aPIUUID)
	return o
}

// SetAPIUUID adds the apiUuid to the post Api asset using p o s t params
func (o *PostAPIAssetUsingPOSTParams) SetAPIUUID(aPIUUID string) {
	o.APIUUID = aPIUUID
}

// WithFiles adds the files to the post Api asset using p o s t params
func (o *PostAPIAssetUsingPOSTParams) WithFiles(files runtime.NamedReadCloser) *PostAPIAssetUsingPOSTParams {
	o.SetFiles(files)
	return o
}

// SetFiles adds the files to the post Api asset using p o s t params
func (o *PostAPIAssetUsingPOSTParams) SetFiles(files runtime.NamedReadCloser) {
	o.Files = files
}

// WithFilesToDelete adds the filesToDelete to the post Api asset using p o s t params
func (o *PostAPIAssetUsingPOSTParams) WithFilesToDelete(filesToDelete *string) *PostAPIAssetUsingPOSTParams {
	o.SetFilesToDelete(filesToDelete)
	return o
}

// SetFilesToDelete adds the filesToDelete to the post Api asset using p o s t params
func (o *PostAPIAssetUsingPOSTParams) SetFilesToDelete(filesToDelete *string) {
	o.FilesToDelete = filesToDelete
}

// WriteToRequest writes these params to a swagger request
func (o *PostAPIAssetUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param apiUuid
	if err := r.SetPathParam("apiUuid", o.APIUUID); err != nil {
		return err
	}

	if o.Files != nil {

		if o.Files != nil {
			// form file param files
			if err := r.SetFileParam("files", o.Files); err != nil {
				return err
			}
		}
	}

	if o.FilesToDelete != nil {

		// form param filesToDelete
		var frFilesToDelete string
		if o.FilesToDelete != nil {
			frFilesToDelete = *o.FilesToDelete
		}
		fFilesToDelete := frFilesToDelete
		if fFilesToDelete != "" {
			if err := r.SetFormParam("filesToDelete", fFilesToDelete); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
