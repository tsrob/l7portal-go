// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new organizations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for organizations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AssociateOrgAPIRLQToAPIPatch(params *AssociateOrgAPIRLQToAPIPatchParams, opts ...ClientOption) (*AssociateOrgAPIRLQToAPIPatchNoContent, error)

	AssociateOrgAPIRLQToAPIPut(params *AssociateOrgAPIRLQToAPIPutParams, opts ...ClientOption) (*AssociateOrgAPIRLQToAPIPutNoContent, error)

	CreateOrganization(params *CreateOrganizationParams, opts ...ClientOption) (*CreateOrganizationCreated, error)

	DeleteOrganization(params *DeleteOrganizationParams, opts ...ClientOption) error

	GetAssociatedOrgAPIRLQ(params *GetAssociatedOrgAPIRLQParams, opts ...ClientOption) (*GetAssociatedOrgAPIRLQOK, error)

	GetOrganization(params *GetOrganizationParams, opts ...ClientOption) (*GetOrganizationOK, error)

	GetOrganizationTag(params *GetOrganizationTagParams, opts ...ClientOption) (*GetOrganizationTagOK, error)

	GetOrganizations(params *GetOrganizationsParams, opts ...ClientOption) (*GetOrganizationsOK, error)

	UpdateOrganization(params *UpdateOrganizationParams, opts ...ClientOption) error

	SetTransport(transport runtime.ClientTransport)
}

/*
AssociateOrgAPIRLQToAPIPatch associates API and rate quota Api per org entities to organization

Associates list of API and Rate Quota(Api per Org) combination to the specified Organization.
*/
func (a *Client) AssociateOrgAPIRLQToAPIPatch(params *AssociateOrgAPIRLQToAPIPatchParams, opts ...ClientOption) (*AssociateOrgAPIRLQToAPIPatchNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAssociateOrgAPIRLQToAPIPatchParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "associateOrgApiRLQToApiPatch",
		Method:             "PATCH",
		PathPattern:        "/tenant-admin/1.0/organizations/{orgUuid}/api-rate-quotas",
		ProducesMediaTypes: []string{"application/json;charset=UTF-8"},
		ConsumesMediaTypes: []string{"application/json;charset=UTF-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AssociateOrgAPIRLQToAPIPatchReader{formats: a.formats},
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
	success, ok := result.(*AssociateOrgAPIRLQToAPIPatchNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for associateOrgApiRLQToApiPatch: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
AssociateOrgAPIRLQToAPIPut associates API and rate quota Api per org entities to organization

Associates list of API and Rate Quota(Api per Org) combination to the specified Organization. An empty payload will disassociate all the existing API and Rate Quota(Api per Org).
*/
func (a *Client) AssociateOrgAPIRLQToAPIPut(params *AssociateOrgAPIRLQToAPIPutParams, opts ...ClientOption) (*AssociateOrgAPIRLQToAPIPutNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAssociateOrgAPIRLQToAPIPutParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "associateOrgApiRLQToApiPut",
		Method:             "PUT",
		PathPattern:        "/tenant-admin/1.0/organizations/{orgUuid}/api-rate-quotas",
		ProducesMediaTypes: []string{"application/json;charset=UTF-8"},
		ConsumesMediaTypes: []string{"application/json;charset=UTF-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AssociateOrgAPIRLQToAPIPutReader{formats: a.formats},
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
	success, ok := result.(*AssociateOrgAPIRLQToAPIPutNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for associateOrgApiRLQToApiPut: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CreateOrganization creates a new organization

Create a new organization
*/
func (a *Client) CreateOrganization(params *CreateOrganizationParams, opts ...ClientOption) (*CreateOrganizationCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateOrganizationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createOrganization",
		Method:             "POST",
		PathPattern:        "/Organizations",
		ProducesMediaTypes: []string{"application/json;charset=UTF-8"},
		ConsumesMediaTypes: []string{"application/json;charset=UTF-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateOrganizationReader{formats: a.formats},
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
	success, ok := result.(*CreateOrganizationCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createOrganization: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteOrganization deletes an existing organization

Delete an existing organization. If the organization is associated with any applications, users, private APIs, or account plans, then it cannot be deleted.
*/
func (a *Client) DeleteOrganization(params *DeleteOrganizationParams, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteOrganizationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteOrganization",
		Method:             "DELETE",
		PathPattern:        "/Organizations('{uuid}')",
		ProducesMediaTypes: []string{"application/json;charset=UTF-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteOrganizationReader{formats: a.formats},
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
GetAssociatedOrgAPIRLQ gets all a p is and rate quota Api per org associated with an organization

Fetches all APIs and Rate Quota(Api per Org) associated with an Organization.
*/
func (a *Client) GetAssociatedOrgAPIRLQ(params *GetAssociatedOrgAPIRLQParams, opts ...ClientOption) (*GetAssociatedOrgAPIRLQOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAssociatedOrgAPIRLQParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAssociatedOrgApiRLQ",
		Method:             "GET",
		PathPattern:        "/tenant-admin/1.0/organizations/{orgUuid}/api-rate-quotas",
		ProducesMediaTypes: []string{"application/json;charset=UTF-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAssociatedOrgAPIRLQReader{formats: a.formats},
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
	success, ok := result.(*GetAssociatedOrgAPIRLQOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAssociatedOrgApiRLQ: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetOrganization returns an organization

Returns an Organization with the specified UUID.
*/
func (a *Client) GetOrganization(params *GetOrganizationParams, opts ...ClientOption) (*GetOrganizationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOrganizationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getOrganization",
		Method:             "GET",
		PathPattern:        "/tenant-admin/1.0/organizations/{uuid}",
		ProducesMediaTypes: []string{"application/json;charset=UTF-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOrganizationReader{formats: a.formats},
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
	success, ok := result.(*GetOrganizationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getOrganization: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetOrganizationTag gets all organization tags that are associated with an organization

List all Organization Tags associated with the specified Organization
*/
func (a *Client) GetOrganizationTag(params *GetOrganizationTagParams, opts ...ClientOption) (*GetOrganizationTagOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOrganizationTagParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getOrganizationTag",
		Method:             "GET",
		PathPattern:        "/tenant-admin/1.0/organizations/{orgUuid}/tags",
		ProducesMediaTypes: []string{"application/json;charset=UTF-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOrganizationTagReader{formats: a.formats},
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
	success, ok := result.(*GetOrganizationTagOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getOrganizationTag: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	GetOrganizations returns a list of all organizations based on the filter parameters

	You use organizations to group and manage related developers. API Owners can also view the list of organizations. When you register a developer, you assign the developer to an organization. Before you register a developer from a new organization, add the organization to the API Portal.

You can assign a Rate Limit & Quota (formerly known as account plan) to each organization to limit how much the applications developed by an organization's developers can use the API.
*/
func (a *Client) GetOrganizations(params *GetOrganizationsParams, opts ...ClientOption) (*GetOrganizationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOrganizationsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getOrganizations",
		Method:             "GET",
		PathPattern:        "/tenant-admin/1.0/organizations",
		ProducesMediaTypes: []string{"application/json;charset=UTF-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOrganizationsReader{formats: a.formats},
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
	success, ok := result.(*GetOrganizationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getOrganizations: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateOrganization updates an existing organization

Update an existing organization.
*/
func (a *Client) UpdateOrganization(params *UpdateOrganizationParams, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateOrganizationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateOrganization",
		Method:             "PUT",
		PathPattern:        "/Organizations('{uuid}')",
		ProducesMediaTypes: []string{"application/json;charset=UTF-8"},
		ConsumesMediaTypes: []string{"application/json;charset=UTF-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateOrganizationReader{formats: a.formats},
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