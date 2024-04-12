// Code generated by go-swagger; DO NOT EDIT.

package import_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new import service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for import service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ImportServiceCancelImport(params *ImportServiceCancelImportParams, opts ...ClientOption) (*ImportServiceCancelImportOK, error)

	ImportServiceCancelUpload(params *ImportServiceCancelUploadParams, opts ...ClientOption) (*ImportServiceCancelUploadOK, error)

	ImportServiceCompleteUpload(params *ImportServiceCompleteUploadParams, opts ...ClientOption) (*ImportServiceCompleteUploadOK, error)

	ImportServiceCreateImport(params *ImportServiceCreateImportParams, opts ...ClientOption) (*ImportServiceCreateImportOK, error)

	ImportServiceGetImport(params *ImportServiceGetImportParams, opts ...ClientOption) (*ImportServiceGetImportOK, error)

	ImportServiceListImports(params *ImportServiceListImportsParams, opts ...ClientOption) (*ImportServiceListImportsOK, error)

	ImportServiceStartUpload(params *ImportServiceStartUploadParams, opts ...ClientOption) (*ImportServiceStartUploadOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
ImportServiceCancelImport import service cancel import API
*/
func (a *Client) ImportServiceCancelImport(params *ImportServiceCancelImportParams, opts ...ClientOption) (*ImportServiceCancelImportOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewImportServiceCancelImportParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ImportService_CancelImport",
		Method:             "POST",
		PathPattern:        "/v1beta1/clusters/{clusterId}/imports/{id}:cancel",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ImportServiceCancelImportReader{formats: a.formats},
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
	success, ok := result.(*ImportServiceCancelImportOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ImportServiceCancelImportDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ImportServiceCancelUpload import service cancel upload API
*/
func (a *Client) ImportServiceCancelUpload(params *ImportServiceCancelUploadParams, opts ...ClientOption) (*ImportServiceCancelUploadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewImportServiceCancelUploadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ImportService_CancelUpload",
		Method:             "POST",
		PathPattern:        "/v1beta1/clusters/{clusterId}/imports:cancelUpload",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ImportServiceCancelUploadReader{formats: a.formats},
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
	success, ok := result.(*ImportServiceCancelUploadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ImportServiceCancelUploadDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ImportServiceCompleteUpload import service complete upload API
*/
func (a *Client) ImportServiceCompleteUpload(params *ImportServiceCompleteUploadParams, opts ...ClientOption) (*ImportServiceCompleteUploadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewImportServiceCompleteUploadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ImportService_CompleteUpload",
		Method:             "POST",
		PathPattern:        "/v1beta1/clusters/{clusterId}/imports:completeUpload",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ImportServiceCompleteUploadReader{formats: a.formats},
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
	success, ok := result.(*ImportServiceCompleteUploadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ImportServiceCompleteUploadDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ImportServiceCreateImport import service create import API
*/
func (a *Client) ImportServiceCreateImport(params *ImportServiceCreateImportParams, opts ...ClientOption) (*ImportServiceCreateImportOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewImportServiceCreateImportParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ImportService_CreateImport",
		Method:             "POST",
		PathPattern:        "/v1beta1/clusters/{clusterId}/imports",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ImportServiceCreateImportReader{formats: a.formats},
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
	success, ok := result.(*ImportServiceCreateImportOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ImportServiceCreateImportDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ImportServiceGetImport import service get import API
*/
func (a *Client) ImportServiceGetImport(params *ImportServiceGetImportParams, opts ...ClientOption) (*ImportServiceGetImportOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewImportServiceGetImportParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ImportService_GetImport",
		Method:             "GET",
		PathPattern:        "/v1beta1/clusters/{clusterId}/imports/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ImportServiceGetImportReader{formats: a.formats},
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
	success, ok := result.(*ImportServiceGetImportOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ImportServiceGetImportDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ImportServiceListImports import service list imports API
*/
func (a *Client) ImportServiceListImports(params *ImportServiceListImportsParams, opts ...ClientOption) (*ImportServiceListImportsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewImportServiceListImportsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ImportService_ListImports",
		Method:             "GET",
		PathPattern:        "/v1beta1/clusters/{clusterId}/imports",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ImportServiceListImportsReader{formats: a.formats},
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
	success, ok := result.(*ImportServiceListImportsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ImportServiceListImportsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ImportServiceStartUpload generates upload url for importing data
*/
func (a *Client) ImportServiceStartUpload(params *ImportServiceStartUploadParams, opts ...ClientOption) (*ImportServiceStartUploadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewImportServiceStartUploadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ImportService_StartUpload",
		Method:             "POST",
		PathPattern:        "/v1beta1/clusters/{clusterId}/imports:startUpload",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ImportServiceStartUploadReader{formats: a.formats},
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
	success, ok := result.(*ImportServiceStartUploadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ImportServiceStartUploadDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
