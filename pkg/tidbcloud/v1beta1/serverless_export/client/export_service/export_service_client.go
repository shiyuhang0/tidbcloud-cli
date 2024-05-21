// Code generated by go-swagger; DO NOT EDIT.

package export_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new export service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new export service API client with basic auth credentials.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - user: user for basic authentication header.
// - password: password for basic authentication header.
func NewClientWithBasicAuth(host, basePath, scheme, user, password string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BasicAuth(user, password)
	return &Client{transport: transport, formats: strfmt.Default}
}

// New creates a new export service API client with a bearer token for authentication.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - bearerToken: bearer token for Bearer authentication header.
func NewClientWithBearerToken(host, basePath, scheme, bearerToken string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BearerToken(bearerToken)
	return &Client{transport: transport, formats: strfmt.Default}
}

/*
Client for export service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ExportServiceCancelExport(params *ExportServiceCancelExportParams, opts ...ClientOption) (*ExportServiceCancelExportOK, error)

	ExportServiceCreateExport(params *ExportServiceCreateExportParams, opts ...ClientOption) (*ExportServiceCreateExportOK, error)

	ExportServiceDeleteExport(params *ExportServiceDeleteExportParams, opts ...ClientOption) (*ExportServiceDeleteExportOK, error)

	ExportServiceDownloadExport(params *ExportServiceDownloadExportParams, opts ...ClientOption) (*ExportServiceDownloadExportOK, error)

	ExportServiceDownloadExportFiles(params *ExportServiceDownloadExportFilesParams, opts ...ClientOption) (*ExportServiceDownloadExportFilesOK, error)

	ExportServiceGetExport(params *ExportServiceGetExportParams, opts ...ClientOption) (*ExportServiceGetExportOK, error)

	ExportServiceListExportFiles(params *ExportServiceListExportFilesParams, opts ...ClientOption) (*ExportServiceListExportFilesOK, error)

	ExportServiceListExports(params *ExportServiceListExportsParams, opts ...ClientOption) (*ExportServiceListExportsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
ExportServiceCancelExport cancels a specific export job
*/
func (a *Client) ExportServiceCancelExport(params *ExportServiceCancelExportParams, opts ...ClientOption) (*ExportServiceCancelExportOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExportServiceCancelExportParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ExportService_CancelExport",
		Method:             "POST",
		PathPattern:        "/v1beta1/clusters/{clusterId}/exports/{exportId}:cancel",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExportServiceCancelExportReader{formats: a.formats},
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
	success, ok := result.(*ExportServiceCancelExportOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ExportServiceCancelExportDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ExportServiceCreateExport creates an export job
*/
func (a *Client) ExportServiceCreateExport(params *ExportServiceCreateExportParams, opts ...ClientOption) (*ExportServiceCreateExportOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExportServiceCreateExportParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ExportService_CreateExport",
		Method:             "POST",
		PathPattern:        "/v1beta1/clusters/{clusterId}/exports",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExportServiceCreateExportReader{formats: a.formats},
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
	success, ok := result.(*ExportServiceCreateExportOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ExportServiceCreateExportDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ExportServiceDeleteExport deletes an export job
*/
func (a *Client) ExportServiceDeleteExport(params *ExportServiceDeleteExportParams, opts ...ClientOption) (*ExportServiceDeleteExportOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExportServiceDeleteExportParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ExportService_DeleteExport",
		Method:             "DELETE",
		PathPattern:        "/v1beta1/clusters/{clusterId}/exports/{exportId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExportServiceDeleteExportReader{formats: a.formats},
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
	success, ok := result.(*ExportServiceDeleteExportOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ExportServiceDeleteExportDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ExportServiceDownloadExport generates download url
*/
func (a *Client) ExportServiceDownloadExport(params *ExportServiceDownloadExportParams, opts ...ClientOption) (*ExportServiceDownloadExportOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExportServiceDownloadExportParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ExportService_DownloadExport",
		Method:             "POST",
		PathPattern:        "/v1beta1/clusters/{clusterId}/exports/{exportId}:download",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExportServiceDownloadExportReader{formats: a.formats},
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
	success, ok := result.(*ExportServiceDownloadExportOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ExportServiceDownloadExportDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ExportServiceDownloadExportFiles generates export files download url
*/
func (a *Client) ExportServiceDownloadExportFiles(params *ExportServiceDownloadExportFilesParams, opts ...ClientOption) (*ExportServiceDownloadExportFilesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExportServiceDownloadExportFilesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ExportService_DownloadExportFiles",
		Method:             "POST",
		PathPattern:        "/v1beta1/clusters/{clusterId}/exports/{exportId}/files:download",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExportServiceDownloadExportFilesReader{formats: a.formats},
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
	success, ok := result.(*ExportServiceDownloadExportFilesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ExportServiceDownloadExportFilesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ExportServiceGetExport retrieves details of an export job
*/
func (a *Client) ExportServiceGetExport(params *ExportServiceGetExportParams, opts ...ClientOption) (*ExportServiceGetExportOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExportServiceGetExportParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ExportService_GetExport",
		Method:             "GET",
		PathPattern:        "/v1beta1/clusters/{clusterId}/exports/{exportId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExportServiceGetExportReader{formats: a.formats},
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
	success, ok := result.(*ExportServiceGetExportOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ExportServiceGetExportDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ExportServiceListExportFiles lists export files
*/
func (a *Client) ExportServiceListExportFiles(params *ExportServiceListExportFilesParams, opts ...ClientOption) (*ExportServiceListExportFilesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExportServiceListExportFilesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ExportService_ListExportFiles",
		Method:             "GET",
		PathPattern:        "/v1beta1/clusters/{clusterId}/exports/{exportId}/files",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExportServiceListExportFilesReader{formats: a.formats},
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
	success, ok := result.(*ExportServiceListExportFilesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ExportServiceListExportFilesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ExportServiceListExports provides a list of export jobs
*/
func (a *Client) ExportServiceListExports(params *ExportServiceListExportsParams, opts ...ClientOption) (*ExportServiceListExportsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExportServiceListExportsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ExportService_ListExports",
		Method:             "GET",
		PathPattern:        "/v1beta1/clusters/{clusterId}/exports",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExportServiceListExportsReader{formats: a.formats},
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
	success, ok := result.(*ExportServiceListExportsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ExportServiceListExportsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
