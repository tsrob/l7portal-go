// Code generated by go-swagger; DO NOT EDIT.

package requests

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new requests API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for requests API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetRequest(params *GetRequestParams, opts ...ClientOption) (*GetRequestOK, error)

	GetRequests(params *GetRequestsParams, opts ...ClientOption) (*GetRequestsOK, error)

	ReviewAllRequest(params *ReviewAllRequestParams, opts ...ClientOption) (*ReviewAllRequestOK, *ReviewAllRequestNoContent, error)

	ReviewApplicationRequestDetail(params *ReviewApplicationRequestDetailParams, opts ...ClientOption) (*ReviewApplicationRequestDetailOK, error)

	ReviewRequest(params *ReviewRequestParams, opts ...ClientOption) (*ReviewRequestNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetRequest gets a specific request

Get a specific Request.
*/
func (a *Client) GetRequest(params *GetRequestParams, opts ...ClientOption) (*GetRequestOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRequestParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getRequest",
		Method:             "GET",
		PathPattern:        "/api-management/1.0/requests/{uuid}",
		ProducesMediaTypes: []string{"application/json;charset=UTF-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRequestReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetRequestOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getRequest: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetRequests returns a list of all developer requests

Returns a list of Developer Requests. There are two possible request types that developers can create. Registration Requests originate from individuals seeking to register an Organization and associated Admin User. Application Requests originate from developers who want a key for their application. Listing and viewing Developer Requests is restricted to those users with an Administrator or API Owner role.  Administrators can list and view both Registration and Application Requests while API owners can only list and view Application Requests
*/
func (a *Client) GetRequests(params *GetRequestsParams, opts ...ClientOption) (*GetRequestsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRequestsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getRequests",
		Method:             "GET",
		PathPattern:        "/api-management/1.0/requests",
		ProducesMediaTypes: []string{"application/json;charset=UTF-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRequestsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetRequestsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getRequests: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ReviewAllRequest reviews all requests for an application

Review all requests for a specific application. Support accept or reject all requests related to an application. Requests can be accepted by setting the Action field to 'ACCEPT' and by optionally setting the Reason field to provide an explanation for the approval. Requests can be rejected by setting the Action field to 'REJECT' and by further setting the Reason field to provide a required explanation as to why the request was rejected. If all the requests can be accepted (or rejected) without errors, the application status should be updated. If there are any requests that can't be accepted or rejected, the application status should not be updated.
*/
func (a *Client) ReviewAllRequest(params *ReviewAllRequestParams, opts ...ClientOption) (*ReviewAllRequestOK, *ReviewAllRequestNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReviewAllRequestParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "reviewAllRequest",
		Method:             "PUT",
		PathPattern:        "/api-management/1.0/requests/applications/{uuid}/review",
		ProducesMediaTypes: []string{"application/json;charset=UTF-8"},
		ConsumesMediaTypes: []string{"application/json;charset=UTF-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ReviewAllRequestReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ReviewAllRequestOK:
		return value, nil, nil
	case *ReviewAllRequestNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for requests: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ReviewApplicationRequestDetail displays a specific application request s detail information

This application request API displays the detail information of the request, which consists of three parts: request excerpt information, original application information, and update request information.
*/
func (a *Client) ReviewApplicationRequestDetail(params *ReviewApplicationRequestDetailParams, opts ...ClientOption) (*ReviewApplicationRequestDetailOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReviewApplicationRequestDetailParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "reviewApplicationRequestDetail",
		Method:             "GET",
		PathPattern:        "/api-management/0.1/application-requests/{uuid}",
		ProducesMediaTypes: []string{"application/json;charset=UTF-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ReviewApplicationRequestDetailReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ReviewApplicationRequestDetailOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for reviewApplicationRequestDetail: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ReviewRequest reviews a specific request

Review a specific Request. Support for updating a Request is provided to allow authorized users to accept or reject a Request. A Request can be accepted by setting the Action field to 'ACCEPT' and by optionally setting the Reason field to provide an explanation for the approval. A Request can be rejected by setting the Action field to 'REJECT' and by further setting the Reason field to provide a required explanation as to why the request was rejected. Updates to other fields will be ignored.
*/
func (a *Client) ReviewRequest(params *ReviewRequestParams, opts ...ClientOption) (*ReviewRequestNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReviewRequestParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "reviewRequest",
		Method:             "PUT",
		PathPattern:        "/api-management/1.0/requests/{uuid}/reviews",
		ProducesMediaTypes: []string{"application/json;charset=UTF-8"},
		ConsumesMediaTypes: []string{"application/json;charset=UTF-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ReviewRequestReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ReviewRequestNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for reviewRequest: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
