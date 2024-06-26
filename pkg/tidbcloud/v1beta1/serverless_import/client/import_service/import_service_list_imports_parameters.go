// Code generated by go-swagger; DO NOT EDIT.

package import_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewImportServiceListImportsParams creates a new ImportServiceListImportsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewImportServiceListImportsParams() *ImportServiceListImportsParams {
	return &ImportServiceListImportsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewImportServiceListImportsParamsWithTimeout creates a new ImportServiceListImportsParams object
// with the ability to set a timeout on a request.
func NewImportServiceListImportsParamsWithTimeout(timeout time.Duration) *ImportServiceListImportsParams {
	return &ImportServiceListImportsParams{
		timeout: timeout,
	}
}

// NewImportServiceListImportsParamsWithContext creates a new ImportServiceListImportsParams object
// with the ability to set a context for a request.
func NewImportServiceListImportsParamsWithContext(ctx context.Context) *ImportServiceListImportsParams {
	return &ImportServiceListImportsParams{
		Context: ctx,
	}
}

// NewImportServiceListImportsParamsWithHTTPClient creates a new ImportServiceListImportsParams object
// with the ability to set a custom HTTPClient for a request.
func NewImportServiceListImportsParamsWithHTTPClient(client *http.Client) *ImportServiceListImportsParams {
	return &ImportServiceListImportsParams{
		HTTPClient: client,
	}
}

/*
ImportServiceListImportsParams contains all the parameters to send to the API endpoint

	for the import service list imports operation.

	Typically these are written to a http.Request.
*/
type ImportServiceListImportsParams struct {

	// ClusterID.
	//
	// Format: uint64
	ClusterID string

	// PageSize.
	//
	// Format: int32
	PageSize *int32

	// PageToken.
	PageToken *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the import service list imports params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImportServiceListImportsParams) WithDefaults() *ImportServiceListImportsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the import service list imports params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImportServiceListImportsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the import service list imports params
func (o *ImportServiceListImportsParams) WithTimeout(timeout time.Duration) *ImportServiceListImportsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the import service list imports params
func (o *ImportServiceListImportsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the import service list imports params
func (o *ImportServiceListImportsParams) WithContext(ctx context.Context) *ImportServiceListImportsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the import service list imports params
func (o *ImportServiceListImportsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the import service list imports params
func (o *ImportServiceListImportsParams) WithHTTPClient(client *http.Client) *ImportServiceListImportsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the import service list imports params
func (o *ImportServiceListImportsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the import service list imports params
func (o *ImportServiceListImportsParams) WithClusterID(clusterID string) *ImportServiceListImportsParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the import service list imports params
func (o *ImportServiceListImportsParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithPageSize adds the pageSize to the import service list imports params
func (o *ImportServiceListImportsParams) WithPageSize(pageSize *int32) *ImportServiceListImportsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the import service list imports params
func (o *ImportServiceListImportsParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithPageToken adds the pageToken to the import service list imports params
func (o *ImportServiceListImportsParams) WithPageToken(pageToken *string) *ImportServiceListImportsParams {
	o.SetPageToken(pageToken)
	return o
}

// SetPageToken adds the pageToken to the import service list imports params
func (o *ImportServiceListImportsParams) SetPageToken(pageToken *string) {
	o.PageToken = pageToken
}

// WriteToRequest writes these params to a swagger request
func (o *ImportServiceListImportsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param clusterId
	if err := r.SetPathParam("clusterId", o.ClusterID); err != nil {
		return err
	}

	if o.PageSize != nil {

		// query param pageSize
		var qrPageSize int32

		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt32(qrPageSize)
		if qPageSize != "" {

			if err := r.SetQueryParam("pageSize", qPageSize); err != nil {
				return err
			}
		}
	}

	if o.PageToken != nil {

		// query param pageToken
		var qrPageToken string

		if o.PageToken != nil {
			qrPageToken = *o.PageToken
		}
		qPageToken := qrPageToken
		if qPageToken != "" {

			if err := r.SetQueryParam("pageToken", qPageToken); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
