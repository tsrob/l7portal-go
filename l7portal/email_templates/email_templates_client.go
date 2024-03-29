// Code generated by go-swagger; DO NOT EDIT.

package email_templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new email templates API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for email templates API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetAllEmailTemplateEntities(params *GetAllEmailTemplateEntitiesParams, opts ...ClientOption) (*GetAllEmailTemplateEntitiesOK, error)

	GetEmailTemplate(params *GetEmailTemplateParams, opts ...ClientOption) (*GetEmailTemplateOK, error)

	UpdateEmailTemplate(params *UpdateEmailTemplateParams, opts ...ClientOption) (*UpdateEmailTemplateNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetAllEmailTemplateEntities returns a list of all email templates

Return a list of all Email Templates.
*/
func (a *Client) GetAllEmailTemplateEntities(params *GetAllEmailTemplateEntitiesParams, opts ...ClientOption) (*GetAllEmailTemplateEntitiesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllEmailTemplateEntitiesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAllEmailTemplateEntities",
		Method:             "GET",
		PathPattern:        "/tenant-admin/1.0/email-templates",
		ProducesMediaTypes: []string{"application/json;charset=UTF-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAllEmailTemplateEntitiesReader{formats: a.formats},
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
	success, ok := result.(*GetAllEmailTemplateEntitiesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAllEmailTemplateEntities: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetEmailTemplate retrieves the type subject and body by the email template UUID

Retrieve the type, subject and body by the Email Template UUID.  The body can be HTML.
*/
func (a *Client) GetEmailTemplate(params *GetEmailTemplateParams, opts ...ClientOption) (*GetEmailTemplateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEmailTemplateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getEmailTemplate",
		Method:             "GET",
		PathPattern:        "/tenant-admin/1.0/email-templates/{uuid}",
		ProducesMediaTypes: []string{"application/json;charset=UTF-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetEmailTemplateReader{formats: a.formats},
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
	success, ok := result.(*GetEmailTemplateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getEmailTemplate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateEmailTemplate updates the email template associated with the given UUID

Updates the Email Template associated with the given UUID.
*/
func (a *Client) UpdateEmailTemplate(params *UpdateEmailTemplateParams, opts ...ClientOption) (*UpdateEmailTemplateNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateEmailTemplateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateEmailTemplate",
		Method:             "PUT",
		PathPattern:        "/tenant-admin/1.0/email-templates/{uuid}",
		ProducesMediaTypes: []string{"application/json;charset=UTF-8"},
		ConsumesMediaTypes: []string{"application/json;charset=UTF-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateEmailTemplateReader{formats: a.formats},
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
	success, ok := result.(*UpdateEmailTemplateNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateEmailTemplate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
