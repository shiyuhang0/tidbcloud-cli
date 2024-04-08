// Code generated by go-swagger; DO NOT EDIT.

package export_service

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
)

// NewExportServiceCreateExportParams creates a new ExportServiceCreateExportParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExportServiceCreateExportParams() *ExportServiceCreateExportParams {
	return &ExportServiceCreateExportParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExportServiceCreateExportParamsWithTimeout creates a new ExportServiceCreateExportParams object
// with the ability to set a timeout on a request.
func NewExportServiceCreateExportParamsWithTimeout(timeout time.Duration) *ExportServiceCreateExportParams {
	return &ExportServiceCreateExportParams{
		timeout: timeout,
	}
}

// NewExportServiceCreateExportParamsWithContext creates a new ExportServiceCreateExportParams object
// with the ability to set a context for a request.
func NewExportServiceCreateExportParamsWithContext(ctx context.Context) *ExportServiceCreateExportParams {
	return &ExportServiceCreateExportParams{
		Context: ctx,
	}
}

// NewExportServiceCreateExportParamsWithHTTPClient creates a new ExportServiceCreateExportParams object
// with the ability to set a custom HTTPClient for a request.
func NewExportServiceCreateExportParamsWithHTTPClient(client *http.Client) *ExportServiceCreateExportParams {
	return &ExportServiceCreateExportParams{
		HTTPClient: client,
	}
}

/*
ExportServiceCreateExportParams contains all the parameters to send to the API endpoint

	for the export service create export operation.

	Typically these are written to a http.Request.
*/
type ExportServiceCreateExportParams struct {

	// Body.
	Body ExportServiceCreateExportBody

	/* ClusterID.

	   Required. The ID of the cluster.
	*/
	ClusterID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the export service create export params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExportServiceCreateExportParams) WithDefaults() *ExportServiceCreateExportParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the export service create export params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExportServiceCreateExportParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the export service create export params
func (o *ExportServiceCreateExportParams) WithTimeout(timeout time.Duration) *ExportServiceCreateExportParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the export service create export params
func (o *ExportServiceCreateExportParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the export service create export params
func (o *ExportServiceCreateExportParams) WithContext(ctx context.Context) *ExportServiceCreateExportParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the export service create export params
func (o *ExportServiceCreateExportParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the export service create export params
func (o *ExportServiceCreateExportParams) WithHTTPClient(client *http.Client) *ExportServiceCreateExportParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the export service create export params
func (o *ExportServiceCreateExportParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the export service create export params
func (o *ExportServiceCreateExportParams) WithBody(body ExportServiceCreateExportBody) *ExportServiceCreateExportParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the export service create export params
func (o *ExportServiceCreateExportParams) SetBody(body ExportServiceCreateExportBody) {
	o.Body = body
}

// WithClusterID adds the clusterID to the export service create export params
func (o *ExportServiceCreateExportParams) WithClusterID(clusterID string) *ExportServiceCreateExportParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the export service create export params
func (o *ExportServiceCreateExportParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WriteToRequest writes these params to a swagger request
func (o *ExportServiceCreateExportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param clusterId
	if err := r.SetPathParam("clusterId", o.ClusterID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
