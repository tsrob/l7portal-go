// Code generated by go-swagger; DO NOT EDIT.

package deployment_information

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new deployment information API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for deployment information API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetApiandAPIKeyDeploymentsStatus(params *GetApiandAPIKeyDeploymentsStatusParams, opts ...ClientOption) (*GetApiandAPIKeyDeploymentsStatusOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetApiandAPIKeyDeploymentsStatus returns a list of summary of the API and API key deployments status

Return a list of summary of the API and API Key deployments status
*/
func (a *Client) GetApiandAPIKeyDeploymentsStatus(params *GetApiandAPIKeyDeploymentsStatusParams, opts ...ClientOption) (*GetApiandAPIKeyDeploymentsStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetApiandAPIKeyDeploymentsStatusParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getApiandApiKeyDeploymentsStatus",
		Method:             "GET",
		PathPattern:        "/deployments/1.0/status",
		ProducesMediaTypes: []string{"application/json;charset=UTF-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetApiandAPIKeyDeploymentsStatusReader{formats: a.formats},
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
	success, ok := result.(*GetApiandAPIKeyDeploymentsStatusOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getApiandApiKeyDeploymentsStatus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
