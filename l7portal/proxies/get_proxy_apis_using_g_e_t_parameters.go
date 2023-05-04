// Code generated by go-swagger; DO NOT EDIT.

package proxies

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
	"github.com/go-openapi/swag"
)

// NewGetProxyApisUsingGETParams creates a new GetProxyApisUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetProxyApisUsingGETParams() *GetProxyApisUsingGETParams {
	return &GetProxyApisUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetProxyApisUsingGETParamsWithTimeout creates a new GetProxyApisUsingGETParams object
// with the ability to set a timeout on a request.
func NewGetProxyApisUsingGETParamsWithTimeout(timeout time.Duration) *GetProxyApisUsingGETParams {
	return &GetProxyApisUsingGETParams{
		timeout: timeout,
	}
}

// NewGetProxyApisUsingGETParamsWithContext creates a new GetProxyApisUsingGETParams object
// with the ability to set a context for a request.
func NewGetProxyApisUsingGETParamsWithContext(ctx context.Context) *GetProxyApisUsingGETParams {
	return &GetProxyApisUsingGETParams{
		Context: ctx,
	}
}

// NewGetProxyApisUsingGETParamsWithHTTPClient creates a new GetProxyApisUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetProxyApisUsingGETParamsWithHTTPClient(client *http.Client) *GetProxyApisUsingGETParams {
	return &GetProxyApisUsingGETParams{
		HTTPClient: client,
	}
}

/*
GetProxyApisUsingGETParams contains all the parameters to send to the API endpoint

	for the get proxy apis using g e t operation.

	Typically these are written to a http.Request.
*/
type GetProxyApisUsingGETParams struct {

	/* APIServiceType.

	   The type of API.
	*/
	APIServiceType *string

	/* End.

	   API deployments deployed before this date and time.

	   Format: yyyy-MM-dd'T'HH:mm:ssZ
	*/
	End *string

	/* EntityUUID.

	   The UUID of the API.
	*/
	EntityUUID *string

	/* PortalStatus.

	   The status of the API. APIs that are deployed to proxies can be DISABLED or DEPRECATED. PENDING_DELETE is a transition status when the API is being deleted and needs to be undeployed from all proxies.
	*/
	PortalStatus *string

	/* PublishedByPortal.

	   Whether the API is portal published or gateway published.
	*/
	PublishedByPortal *bool

	/* Start.

	   API deployments deployed after this date and time.

	   Format: yyyy-MM-dd'T'HH:mm:ssZ
	*/
	Start *string

	/* Status.

	   The status of the API deployment.
	*/
	Status *string

	/* UUID.

	   The proxy UUID to get.
	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get proxy apis using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProxyApisUsingGETParams) WithDefaults() *GetProxyApisUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get proxy apis using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProxyApisUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get proxy apis using g e t params
func (o *GetProxyApisUsingGETParams) WithTimeout(timeout time.Duration) *GetProxyApisUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get proxy apis using g e t params
func (o *GetProxyApisUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get proxy apis using g e t params
func (o *GetProxyApisUsingGETParams) WithContext(ctx context.Context) *GetProxyApisUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get proxy apis using g e t params
func (o *GetProxyApisUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get proxy apis using g e t params
func (o *GetProxyApisUsingGETParams) WithHTTPClient(client *http.Client) *GetProxyApisUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get proxy apis using g e t params
func (o *GetProxyApisUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIServiceType adds the aPIServiceType to the get proxy apis using g e t params
func (o *GetProxyApisUsingGETParams) WithAPIServiceType(aPIServiceType *string) *GetProxyApisUsingGETParams {
	o.SetAPIServiceType(aPIServiceType)
	return o
}

// SetAPIServiceType adds the apiServiceType to the get proxy apis using g e t params
func (o *GetProxyApisUsingGETParams) SetAPIServiceType(aPIServiceType *string) {
	o.APIServiceType = aPIServiceType
}

// WithEnd adds the end to the get proxy apis using g e t params
func (o *GetProxyApisUsingGETParams) WithEnd(end *string) *GetProxyApisUsingGETParams {
	o.SetEnd(end)
	return o
}

// SetEnd adds the end to the get proxy apis using g e t params
func (o *GetProxyApisUsingGETParams) SetEnd(end *string) {
	o.End = end
}

// WithEntityUUID adds the entityUUID to the get proxy apis using g e t params
func (o *GetProxyApisUsingGETParams) WithEntityUUID(entityUUID *string) *GetProxyApisUsingGETParams {
	o.SetEntityUUID(entityUUID)
	return o
}

// SetEntityUUID adds the entityUuid to the get proxy apis using g e t params
func (o *GetProxyApisUsingGETParams) SetEntityUUID(entityUUID *string) {
	o.EntityUUID = entityUUID
}

// WithPortalStatus adds the portalStatus to the get proxy apis using g e t params
func (o *GetProxyApisUsingGETParams) WithPortalStatus(portalStatus *string) *GetProxyApisUsingGETParams {
	o.SetPortalStatus(portalStatus)
	return o
}

// SetPortalStatus adds the portalStatus to the get proxy apis using g e t params
func (o *GetProxyApisUsingGETParams) SetPortalStatus(portalStatus *string) {
	o.PortalStatus = portalStatus
}

// WithPublishedByPortal adds the publishedByPortal to the get proxy apis using g e t params
func (o *GetProxyApisUsingGETParams) WithPublishedByPortal(publishedByPortal *bool) *GetProxyApisUsingGETParams {
	o.SetPublishedByPortal(publishedByPortal)
	return o
}

// SetPublishedByPortal adds the publishedByPortal to the get proxy apis using g e t params
func (o *GetProxyApisUsingGETParams) SetPublishedByPortal(publishedByPortal *bool) {
	o.PublishedByPortal = publishedByPortal
}

// WithStart adds the start to the get proxy apis using g e t params
func (o *GetProxyApisUsingGETParams) WithStart(start *string) *GetProxyApisUsingGETParams {
	o.SetStart(start)
	return o
}

// SetStart adds the start to the get proxy apis using g e t params
func (o *GetProxyApisUsingGETParams) SetStart(start *string) {
	o.Start = start
}

// WithStatus adds the status to the get proxy apis using g e t params
func (o *GetProxyApisUsingGETParams) WithStatus(status *string) *GetProxyApisUsingGETParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the get proxy apis using g e t params
func (o *GetProxyApisUsingGETParams) SetStatus(status *string) {
	o.Status = status
}

// WithUUID adds the uuid to the get proxy apis using g e t params
func (o *GetProxyApisUsingGETParams) WithUUID(uuid string) *GetProxyApisUsingGETParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the get proxy apis using g e t params
func (o *GetProxyApisUsingGETParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *GetProxyApisUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.APIServiceType != nil {

		// query param apiServiceType
		var qrAPIServiceType string

		if o.APIServiceType != nil {
			qrAPIServiceType = *o.APIServiceType
		}
		qAPIServiceType := qrAPIServiceType
		if qAPIServiceType != "" {

			if err := r.SetQueryParam("apiServiceType", qAPIServiceType); err != nil {
				return err
			}
		}
	}

	if o.End != nil {

		// query param end
		var qrEnd string

		if o.End != nil {
			qrEnd = *o.End
		}
		qEnd := qrEnd
		if qEnd != "" {

			if err := r.SetQueryParam("end", qEnd); err != nil {
				return err
			}
		}
	}

	if o.EntityUUID != nil {

		// query param entityUuid
		var qrEntityUUID string

		if o.EntityUUID != nil {
			qrEntityUUID = *o.EntityUUID
		}
		qEntityUUID := qrEntityUUID
		if qEntityUUID != "" {

			if err := r.SetQueryParam("entityUuid", qEntityUUID); err != nil {
				return err
			}
		}
	}

	if o.PortalStatus != nil {

		// query param portalStatus
		var qrPortalStatus string

		if o.PortalStatus != nil {
			qrPortalStatus = *o.PortalStatus
		}
		qPortalStatus := qrPortalStatus
		if qPortalStatus != "" {

			if err := r.SetQueryParam("portalStatus", qPortalStatus); err != nil {
				return err
			}
		}
	}

	if o.PublishedByPortal != nil {

		// query param publishedByPortal
		var qrPublishedByPortal bool

		if o.PublishedByPortal != nil {
			qrPublishedByPortal = *o.PublishedByPortal
		}
		qPublishedByPortal := swag.FormatBool(qrPublishedByPortal)
		if qPublishedByPortal != "" {

			if err := r.SetQueryParam("publishedByPortal", qPublishedByPortal); err != nil {
				return err
			}
		}
	}

	if o.Start != nil {

		// query param start
		var qrStart string

		if o.Start != nil {
			qrStart = *o.Start
		}
		qStart := qrStart
		if qStart != "" {

			if err := r.SetQueryParam("start", qStart); err != nil {
				return err
			}
		}
	}

	if o.Status != nil {

		// query param status
		var qrStatus string

		if o.Status != nil {
			qrStatus = *o.Status
		}
		qStatus := qrStatus
		if qStatus != "" {

			if err := r.SetQueryParam("status", qStatus); err != nil {
				return err
			}
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
