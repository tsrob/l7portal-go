// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new users API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for users API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateUserV2(params *CreateUserV2Params, opts ...ClientOption) (*CreateUserV2Created, error)

	DeleteUserV2(params *DeleteUserV2Params, opts ...ClientOption) error

	GetDeveloperRoleTypes(params *GetDeveloperRoleTypesParams, opts ...ClientOption) (*GetDeveloperRoleTypesOK, error)

	GetLanguages(params *GetLanguagesParams, opts ...ClientOption) (*GetLanguagesOK, error)

	GetPublisherRoleTypes(params *GetPublisherRoleTypesParams, opts ...ClientOption) (*GetPublisherRoleTypesOK, error)

	GetUserV2(params *GetUserV2Params, opts ...ClientOption) (*GetUserV2OK, error)

	GetUsersV2(params *GetUsersV2Params, opts ...ClientOption) (*GetUsersV2OK, error)

	UpdateUserV2(params *UpdateUserV2Params, opts ...ClientOption) (*UpdateUserV2Created, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateUserV2 creates a new user

Create a new User
*/
func (a *Client) CreateUserV2(params *CreateUserV2Params, opts ...ClientOption) (*CreateUserV2Created, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateUserV2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "createUserV2",
		Method:             "POST",
		PathPattern:        "/v2/users",
		ProducesMediaTypes: []string{"application/json;charset=UTF-8"},
		ConsumesMediaTypes: []string{"application/json;charset=UTF-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateUserV2Reader{formats: a.formats},
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
	success, ok := result.(*CreateUserV2Created)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createUserV2: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteUserV2 deletes an existing user record

Delete an existing User record.
*/
func (a *Client) DeleteUserV2(params *DeleteUserV2Params, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteUserV2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteUserV2",
		Method:             "DELETE",
		PathPattern:        "/v2/users/{uuid}",
		ProducesMediaTypes: []string{"application/json;charset=UTF-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteUserV2Reader{formats: a.formats},
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
	GetDeveloperRoleTypes lists all developer role types

	List All Developer Role Types. The developer roles are Organization Administrator and Developer. Each developer organization needs at least one Organization Administrator. An Organization Administrator can be an application developer. The devorgadministrators or the Organization Administrator role can: test and Explore APIs, edit user accounts, reset user passwords, edit your user profile, add applications, edit applications. The developers or Developer role can: test and explore APIs, edit your user profile and edit applications.;

EXT_USER_DASHBOARD_URI: Updates the URI of the Dashboard seen by external users (e.g. "/admin/app/dashboard.html")
*/
func (a *Client) GetDeveloperRoleTypes(params *GetDeveloperRoleTypesParams, opts ...ClientOption) (*GetDeveloperRoleTypesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDeveloperRoleTypesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getDeveloperRoleTypes",
		Method:             "GET",
		PathPattern:        "/developerRoleTypes",
		ProducesMediaTypes: []string{"application/json;charset=UTF-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDeveloperRoleTypesReader{formats: a.formats},
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
	success, ok := result.(*GetDeveloperRoleTypesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getDeveloperRoleTypes: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetLanguages lists all languages

List All languages. Only English is the supported language type.
*/
func (a *Client) GetLanguages(params *GetLanguagesParams, opts ...ClientOption) (*GetLanguagesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLanguagesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getLanguages",
		Method:             "GET",
		PathPattern:        "/languages",
		ProducesMediaTypes: []string{"application/json;charset=UTF-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetLanguagesReader{formats: a.formats},
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
	success, ok := result.(*GetLanguagesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getLanguages: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetPublisherRoleTypes lists all publisher role types

List All Publisher Role Types. The portaladministrators or Administrator role can do everything in the API Portal: customize the API Portal, add and edit API EULAs, publish and edit APIs, test and explore APIs, add and edit account plans, add and edit developer organizations, add users, edit user accounts, edit your user profile, add applications and edit applications.  The API Portal needs at least one Administrator. The apiowners or API Owner role can add and edit API EULAs, publish and edit APIs, test and explore APIs, edit his profile, add applications and edit applications.
*/
func (a *Client) GetPublisherRoleTypes(params *GetPublisherRoleTypesParams, opts ...ClientOption) (*GetPublisherRoleTypesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPublisherRoleTypesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getPublisherRoleTypes",
		Method:             "GET",
		PathPattern:        "/publisherRoleTypes",
		ProducesMediaTypes: []string{"application/json;charset=UTF-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPublisherRoleTypesReader{formats: a.formats},
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
	success, ok := result.(*GetPublisherRoleTypesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getPublisherRoleTypes: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetUserV2 returns a single user

Returns a single User
*/
func (a *Client) GetUserV2(params *GetUserV2Params, opts ...ClientOption) (*GetUserV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUserV2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "getUserV2",
		Method:             "GET",
		PathPattern:        "/v2/users/{uuid}",
		ProducesMediaTypes: []string{"application/json;charset=UTF-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetUserV2Reader{formats: a.formats},
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
	success, ok := result.(*GetUserV2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getUserV2: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetUsersV2 returns a list of all users based on the request parameters

Returns a list of all Users.

	 All registered users of the API Portal have a role. The functionality available to users depends on their role. The API Portal has two role categories: publisher roles and developer roles. The publisher roles are Administrator and API Owner. The developer roles are Organization Administrator and Developer. Organization Administrators can perform the same account management tasks for their organization members. However, Organization Administrators can add users belongs to his Organization. Before developers can use APIs published on the API Portal, they need developer accounts.
	You can add user account to the API Portal. If the Notify User check box is set to true, the API Portal sends an email to the user. The email message contains an activation token in a URL.
	When users register for a developer account(Developer type > Org Admin role), the API Portal sends them an account activation email. The account is then pending. To finish the self-registration process, users click the account activation link in the email and complete the Account Setup form. Users must complete the Account Setup form before the account activation token expires.
	  When the User's Status is REGISTRATION_INIT, Administrators can resend the account activation email. This function helps when users cannot find, or mistakenly delete, the email. Pending accounts do not show the user's name.
	 Administrators can delete pending developer accounts. When you delete pending accounts, the API Portal deletes the user records and revokes the account activation tokens.
*/
func (a *Client) GetUsersV2(params *GetUsersV2Params, opts ...ClientOption) (*GetUsersV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUsersV2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "getUsersV2",
		Method:             "GET",
		PathPattern:        "/v2/users",
		ProducesMediaTypes: []string{"application/json;charset=UTF-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetUsersV2Reader{formats: a.formats},
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
	success, ok := result.(*GetUsersV2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getUsersV2: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateUserV2 updates an existing user record

Update an existing User record
*/
func (a *Client) UpdateUserV2(params *UpdateUserV2Params, opts ...ClientOption) (*UpdateUserV2Created, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateUserV2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateUserV2",
		Method:             "PUT",
		PathPattern:        "/v2/users/{uuid}",
		ProducesMediaTypes: []string{"application/json;charset=UTF-8"},
		ConsumesMediaTypes: []string{"application/json;charset=UTF-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateUserV2Reader{formats: a.formats},
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
	success, ok := result.(*UpdateUserV2Created)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateUserV2: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}