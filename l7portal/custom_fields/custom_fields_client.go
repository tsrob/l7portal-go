// Code generated by go-swagger; DO NOT EDIT.

package custom_fields

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new custom fields API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for custom fields API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AddCustomField(params *AddCustomFieldParams, opts ...ClientOption) (*AddCustomFieldCreated, error)

	DeleteCustomField(params *DeleteCustomFieldParams, opts ...ClientOption) error

	GetCustomField(params *GetCustomFieldParams, opts ...ClientOption) (*GetCustomFieldOK, error)

	GetCustomFields(params *GetCustomFieldsParams, opts ...ClientOption) (*GetCustomFieldsOK, error)

	UpdateCustomField(params *UpdateCustomFieldParams, opts ...ClientOption) error

	SetTransport(transport runtime.ClientTransport)
}

/*
AddCustomField creates a custom field record

Create a custom field record
*/
func (a *Client) AddCustomField(params *AddCustomFieldParams, opts ...ClientOption) (*AddCustomFieldCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddCustomFieldParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "addCustomField",
		Method:             "POST",
		PathPattern:        "/CustomFields",
		ProducesMediaTypes: []string{"application/json;charset=UTF-8"},
		ConsumesMediaTypes: []string{"application/json;charset=UTF-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddCustomFieldReader{formats: a.formats},
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
	success, ok := result.(*AddCustomFieldCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for addCustomField: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteCustomField deletes an existing custom field record

Delete an existing CustomField record.
*/
func (a *Client) DeleteCustomField(params *DeleteCustomFieldParams, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteCustomFieldParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteCustomField",
		Method:             "DELETE",
		PathPattern:        "/CustomFields('{uuid}')",
		ProducesMediaTypes: []string{"application/json;charset=UTF-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteCustomFieldReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
GetCustomField returns a specific custom field

Returns a specific Custom Field.
*/
func (a *Client) GetCustomField(params *GetCustomFieldParams, opts ...ClientOption) (*GetCustomFieldOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCustomFieldParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getCustomField",
		Method:             "GET",
		PathPattern:        "/CustomFields('{uuid}')",
		ProducesMediaTypes: []string{"application/json;charset=UTF-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCustomFieldReader{formats: a.formats},
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
	success, ok := result.(*GetCustomFieldOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getCustomField: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetCustomFields lists all custom fields with their names description statuses input type and required optional state

List all Custom Fields with their names, description, statuses, input type and required/optional state. Input type contains the Single-Select Dropdown and Text. There must be at least one option for the Single-Select Dropdown input type.

	Custom Fields provide additional metadata for APIs and Applications. For each API or Application, you can assign the same or different values.
*/
func (a *Client) GetCustomFields(params *GetCustomFieldsParams, opts ...ClientOption) (*GetCustomFieldsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCustomFieldsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getCustomFields",
		Method:             "GET",
		PathPattern:        "/CustomFields",
		ProducesMediaTypes: []string{"application/json;charset=UTF-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCustomFieldsReader{formats: a.formats},
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
	success, ok := result.(*GetCustomFieldsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getCustomFields: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateCustomField updates an existing custom field record

Update a custom field record
*/
func (a *Client) UpdateCustomField(params *UpdateCustomFieldParams, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateCustomFieldParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateCustomField",
		Method:             "PUT",
		PathPattern:        "/CustomFields('{uuid}')",
		ProducesMediaTypes: []string{"application/json;charset=UTF-8"},
		ConsumesMediaTypes: []string{"application/json;charset=UTF-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateCustomFieldReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
