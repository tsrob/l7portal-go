// Code generated by go-swagger; DO NOT EDIT.

package documents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new documents API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for documents API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteMarkDownAsset(params *DeleteMarkDownAssetParams, opts ...ClientOption) (*DeleteMarkDownAssetOK, error)

	DeleteMarkDownAssets(params *DeleteMarkDownAssetsParams, opts ...ClientOption) (*DeleteMarkDownAssetsOK, error)

	GetMarkDownAssetByNavTitle(params *GetMarkDownAssetByNavTitleParams, opts ...ClientOption) (*GetMarkDownAssetByNavTitleOK, error)

	GetMarkDownAssetDocs(params *GetMarkDownAssetDocsParams, opts ...ClientOption) (*GetMarkDownAssetDocsOK, error)

	PostMarkDownAsset(params *PostMarkDownAssetParams, opts ...ClientOption) (*PostMarkDownAssetCreated, error)

	PutMarkDownAsset(params *PutMarkDownAssetParams, opts ...ClientOption) (*PutMarkDownAssetOK, error)

	PutMarkDownDocs(params *PutMarkDownDocsParams, opts ...ClientOption) (*PutMarkDownDocsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
DeleteMarkDownAsset deletes a specific document for the associated entity

Deletes a specific document associated with a navigation title and locale. Use this endpoint to delete the individual document associated with the 'navtitle' and 'locale' combination. You can delete parent documents that do not have children. Attempts to delete parent documents with children results in errors. To delete parent documents with children, use the query parameter 'forceDelete=true'.
*/
func (a *Client) DeleteMarkDownAsset(params *DeleteMarkDownAssetParams, opts ...ClientOption) (*DeleteMarkDownAssetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteMarkDownAssetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteMarkDownAsset",
		Method:             "DELETE",
		PathPattern:        "/document-service/1.0/docs/{type}/{typeUuid}/{navtitle}",
		ProducesMediaTypes: []string{"application/json;charset=UTF-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteMarkDownAssetReader{formats: a.formats},
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
	success, ok := result.(*DeleteMarkDownAssetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteMarkDownAsset: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteMarkDownAssets deletes all document for the associated entity

Deletes all document associated with an entity. Use this endpoint to delete all documents associated with an entity. You can delete locale-specific documents by passing the optional query parameter 'locale'.
*/
func (a *Client) DeleteMarkDownAssets(params *DeleteMarkDownAssetsParams, opts ...ClientOption) (*DeleteMarkDownAssetsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteMarkDownAssetsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteMarkDownAssets",
		Method:             "DELETE",
		PathPattern:        "/document-service/1.0/docs/{type}/{typeUuid}",
		ProducesMediaTypes: []string{"application/json;charset=UTF-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteMarkDownAssetsReader{formats: a.formats},
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
	success, ok := result.(*DeleteMarkDownAssetsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteMarkDownAssets: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetMarkDownAssetByNavTitle returns a specific document for the associated entity

Returns a specific document. Use this endpoint to return the specific document associated with the 'navtitle' and 'locale' combination.
*/
func (a *Client) GetMarkDownAssetByNavTitle(params *GetMarkDownAssetByNavTitleParams, opts ...ClientOption) (*GetMarkDownAssetByNavTitleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMarkDownAssetByNavTitleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getMarkDownAssetByNavTitle",
		Method:             "GET",
		PathPattern:        "/document-service/1.0/docs/{type}/{typeUuid}/{navtitle}",
		ProducesMediaTypes: []string{"application/json;charset=UTF-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetMarkDownAssetByNavTitleReader{formats: a.formats},
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
	success, ok := result.(*GetMarkDownAssetByNavTitleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getMarkDownAssetByNavTitle: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetMarkDownAssetDocs returns a list of available documents for a specific entity

Returns the  metadata (excluding the markdown) for each of the documents that is associated with entity. Use the endpoint to provide a navigation panel of your documents.
*/
func (a *Client) GetMarkDownAssetDocs(params *GetMarkDownAssetDocsParams, opts ...ClientOption) (*GetMarkDownAssetDocsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMarkDownAssetDocsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getMarkDownAssetDocs",
		Method:             "GET",
		PathPattern:        "/document-service/1.0/docs/{type}/{typeUuid}",
		ProducesMediaTypes: []string{"application/json;charset=UTF-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetMarkDownAssetDocsReader{formats: a.formats},
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
	success, ok := result.(*GetMarkDownAssetDocsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getMarkDownAssetDocs: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostMarkDownAsset adds a document for a specific entity

Adds a document for a specific entity. Use this endpoint to add documents associated with the entity.
*/
func (a *Client) PostMarkDownAsset(params *PostMarkDownAssetParams, opts ...ClientOption) (*PostMarkDownAssetCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostMarkDownAssetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "postMarkDownAsset",
		Method:             "POST",
		PathPattern:        "/document-service/1.0/docs/{type}/{typeUuid}/{navtitle}",
		ProducesMediaTypes: []string{"application/json;charset=UTF-8"},
		ConsumesMediaTypes: []string{"application/json;charset=UTF-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostMarkDownAssetReader{formats: a.formats},
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
	success, ok := result.(*PostMarkDownAssetCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for postMarkDownAsset: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PutMarkDownAsset updates a specific document for the associated entity

Updates a specific document associated with a navigation title and locale. Requires the latest modified timestamp to update document. If the modified timestamp is not the latest, this indicates that changes have been made to the document since your document was last loaded, and an error is thrown. To update documents that have an older modified timestamp, override using the query parameter 'forceUpdate=true'.
*/
func (a *Client) PutMarkDownAsset(params *PutMarkDownAssetParams, opts ...ClientOption) (*PutMarkDownAssetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutMarkDownAssetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "putMarkDownAsset",
		Method:             "PUT",
		PathPattern:        "/document-service/1.0/docs/{type}/{typeUuid}/{navtitle}",
		ProducesMediaTypes: []string{"application/json;charset=UTF-8"},
		ConsumesMediaTypes: []string{"application/json;charset=UTF-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutMarkDownAssetReader{formats: a.formats},
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
	success, ok := result.(*PutMarkDownAssetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for putMarkDownAsset: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PutMarkDownDocs updates the set of documents for a specific entity

Updates the set of documents' metadata (excluding the markdown) for a specific entity. The latest modified timestamp in the request must match the timestamp in the database. If the modified timestamp in the request is not the latest, this indicates that changes have been made to the document since your document was last loaded, and an error will be thrown. To update the document's metadata that has an older modified timestamp, override using the query parameter 'forceUpdate=true'.
*/
func (a *Client) PutMarkDownDocs(params *PutMarkDownDocsParams, opts ...ClientOption) (*PutMarkDownDocsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutMarkDownDocsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "putMarkDownDocs",
		Method:             "PUT",
		PathPattern:        "/document-service/1.0/docs/{type}/{typeUuid}",
		ProducesMediaTypes: []string{"application/json;charset=UTF-8"},
		ConsumesMediaTypes: []string{"application/json;charset=UTF-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutMarkDownDocsReader{formats: a.formats},
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
	success, ok := result.(*PutMarkDownDocsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for putMarkDownDocs: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
