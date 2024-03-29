// Code generated by go-swagger; DO NOT EDIT.

package api_key_deployments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new api key deployments API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for api key deployments API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateAPIKeyDeploymentUsingPOST1(params *CreateAPIKeyDeploymentUsingPOST1Params, opts ...ClientOption) (*CreateAPIKeyDeploymentUsingPOST1Created, error)

	DeleteAPIKeyDeploymentUsingDELETE1(params *DeleteAPIKeyDeploymentUsingDELETE1Params, opts ...ClientOption) (*DeleteAPIKeyDeploymentUsingDELETE1NoContent, error)

	GetAPIKeyDeploymentUsingGET1(params *GetAPIKeyDeploymentUsingGET1Params, opts ...ClientOption) (*GetAPIKeyDeploymentUsingGET1OK, error)

	GetAPIKeyDeploymentsUsingGET1(params *GetAPIKeyDeploymentsUsingGET1Params, opts ...ClientOption) (*GetAPIKeyDeploymentsUsingGET1OK, error)

	UpdateAPIKeyDeploymentUsingPUT1(params *UpdateAPIKeyDeploymentUsingPUT1Params, opts ...ClientOption) (*UpdateAPIKeyDeploymentUsingPUT1OK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateAPIKeyDeploymentUsingPOST1 deploys the API key to the proxy

The API Key will be deployed to the specified proxy.
*/
func (a *Client) CreateAPIKeyDeploymentUsingPOST1(params *CreateAPIKeyDeploymentUsingPOST1Params, opts ...ClientOption) (*CreateAPIKeyDeploymentUsingPOST1Created, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateAPIKeyDeploymentUsingPOST1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "createApiKeyDeploymentUsingPOST_1",
		Method:             "POST",
		PathPattern:        "/deployments/1.0/api-keys/{apiKey}/proxies",
		ProducesMediaTypes: []string{"application/json;charset=UTF-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateAPIKeyDeploymentUsingPOST1Reader{formats: a.formats},
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
	success, ok := result.(*CreateAPIKeyDeploymentUsingPOST1Created)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createApiKeyDeploymentUsingPOST_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteAPIKeyDeploymentUsingDELETE1 undeploys the API key from the proxy

Un-deploys the specified API Key from the specified proxy.
*/
func (a *Client) DeleteAPIKeyDeploymentUsingDELETE1(params *DeleteAPIKeyDeploymentUsingDELETE1Params, opts ...ClientOption) (*DeleteAPIKeyDeploymentUsingDELETE1NoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAPIKeyDeploymentUsingDELETE1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteApiKeyDeploymentUsingDELETE_1",
		Method:             "DELETE",
		PathPattern:        "/deployments/1.0/api-keys/{apiKey}/proxies/{proxyUuid}",
		ProducesMediaTypes: []string{"application/json;charset=UTF-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteAPIKeyDeploymentUsingDELETE1Reader{formats: a.formats},
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
	success, ok := result.(*DeleteAPIKeyDeploymentUsingDELETE1NoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteApiKeyDeploymentUsingDELETE_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAPIKeyDeploymentUsingGET1 returns the API key s deployment details status message and where the API key is deployed for the proxy

Retrieves API Key's deployment details for the given proxy.
*/
func (a *Client) GetAPIKeyDeploymentUsingGET1(params *GetAPIKeyDeploymentUsingGET1Params, opts ...ClientOption) (*GetAPIKeyDeploymentUsingGET1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPIKeyDeploymentUsingGET1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "getApiKeyDeploymentUsingGET_1",
		Method:             "GET",
		PathPattern:        "/deployments/1.0/api-keys/{apiKey}/proxies/{proxyUuid}",
		ProducesMediaTypes: []string{"application/json;charset=UTF-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAPIKeyDeploymentUsingGET1Reader{formats: a.formats},
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
	success, ok := result.(*GetAPIKeyDeploymentUsingGET1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getApiKeyDeploymentUsingGET_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAPIKeyDeploymentsUsingGET1 returns the list of proxies and the API key s deployment details including status message and where the API key is deployed

Returns the list of proxies and the API key's deployment details, including status, message, and where the API key is deployed.
*/
func (a *Client) GetAPIKeyDeploymentsUsingGET1(params *GetAPIKeyDeploymentsUsingGET1Params, opts ...ClientOption) (*GetAPIKeyDeploymentsUsingGET1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPIKeyDeploymentsUsingGET1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "getApiKeyDeploymentsUsingGET_1",
		Method:             "GET",
		PathPattern:        "/deployments/1.0/api-keys/{apiKey}/proxies",
		ProducesMediaTypes: []string{"application/json;charset=UTF-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAPIKeyDeploymentsUsingGET1Reader{formats: a.formats},
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
	success, ok := result.(*GetAPIKeyDeploymentsUsingGET1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getApiKeyDeploymentsUsingGET_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateAPIKeyDeploymentUsingPUT1 updates the API key s deployment details status message and where the API key is deployed for the proxy

Updates the specified API key's deployment status and message for the specified proxy.
*/
func (a *Client) UpdateAPIKeyDeploymentUsingPUT1(params *UpdateAPIKeyDeploymentUsingPUT1Params, opts ...ClientOption) (*UpdateAPIKeyDeploymentUsingPUT1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateAPIKeyDeploymentUsingPUT1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateApiKeyDeploymentUsingPUT_1",
		Method:             "PUT",
		PathPattern:        "/deployments/1.0/api-keys/{apiKey}/proxies/{proxyUuid}",
		ProducesMediaTypes: []string{"application/json;charset=UTF-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateAPIKeyDeploymentUsingPUT1Reader{formats: a.formats},
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
	success, ok := result.(*UpdateAPIKeyDeploymentUsingPUT1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateApiKeyDeploymentUsingPUT_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
