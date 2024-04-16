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
	ImportServiceAbortMultipartUpload(params *ImportServiceAbortMultipartUploadParams, opts ...ClientOption) (*ImportServiceAbortMultipartUploadOK, error)

	ImportServiceCompleteMultipartUpload(params *ImportServiceCompleteMultipartUploadParams, opts ...ClientOption) (*ImportServiceCompleteMultipartUploadOK, error)

	ImportServiceGenerateUploadURL(params *ImportServiceGenerateUploadURLParams, opts ...ClientOption) (*ImportServiceGenerateUploadURLOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
ImportServiceAbortMultipartUpload import service abort multipart upload API
*/
func (a *Client) ImportServiceAbortMultipartUpload(params *ImportServiceAbortMultipartUploadParams, opts ...ClientOption) (*ImportServiceAbortMultipartUploadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewImportServiceAbortMultipartUploadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ImportService_AbortMultipartUpload",
		Method:             "POST",
		PathPattern:        "/v1beta1/clusters/{clusterId}/upload_url/abort",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ImportServiceAbortMultipartUploadReader{formats: a.formats},
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
	success, ok := result.(*ImportServiceAbortMultipartUploadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ImportServiceAbortMultipartUploadDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ImportServiceCompleteMultipartUpload import service complete multipart upload API
*/
func (a *Client) ImportServiceCompleteMultipartUpload(params *ImportServiceCompleteMultipartUploadParams, opts ...ClientOption) (*ImportServiceCompleteMultipartUploadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewImportServiceCompleteMultipartUploadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ImportService_CompleteMultipartUpload",
		Method:             "POST",
		PathPattern:        "/v1beta1/clusters/{clusterId}/upload_url/complete",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ImportServiceCompleteMultipartUploadReader{formats: a.formats},
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
	success, ok := result.(*ImportServiceCompleteMultipartUploadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ImportServiceCompleteMultipartUploadDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ImportServiceGenerateUploadURL generates upload url for importing data
*/
func (a *Client) ImportServiceGenerateUploadURL(params *ImportServiceGenerateUploadURLParams, opts ...ClientOption) (*ImportServiceGenerateUploadURLOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewImportServiceGenerateUploadURLParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ImportService_GenerateUploadUrl",
		Method:             "GET",
		PathPattern:        "/v1beta1/clusters/{clusterId}/upload_url",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ImportServiceGenerateUploadURLReader{formats: a.formats},
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
	success, ok := result.(*ImportServiceGenerateUploadURLOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ImportServiceGenerateUploadURLDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}