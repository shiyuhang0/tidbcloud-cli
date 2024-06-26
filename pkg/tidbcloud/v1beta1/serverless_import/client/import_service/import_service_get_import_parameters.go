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
)

// NewImportServiceGetImportParams creates a new ImportServiceGetImportParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewImportServiceGetImportParams() *ImportServiceGetImportParams {
	return &ImportServiceGetImportParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewImportServiceGetImportParamsWithTimeout creates a new ImportServiceGetImportParams object
// with the ability to set a timeout on a request.
func NewImportServiceGetImportParamsWithTimeout(timeout time.Duration) *ImportServiceGetImportParams {
	return &ImportServiceGetImportParams{
		timeout: timeout,
	}
}

// NewImportServiceGetImportParamsWithContext creates a new ImportServiceGetImportParams object
// with the ability to set a context for a request.
func NewImportServiceGetImportParamsWithContext(ctx context.Context) *ImportServiceGetImportParams {
	return &ImportServiceGetImportParams{
		Context: ctx,
	}
}

// NewImportServiceGetImportParamsWithHTTPClient creates a new ImportServiceGetImportParams object
// with the ability to set a custom HTTPClient for a request.
func NewImportServiceGetImportParamsWithHTTPClient(client *http.Client) *ImportServiceGetImportParams {
	return &ImportServiceGetImportParams{
		HTTPClient: client,
	}
}

/*
ImportServiceGetImportParams contains all the parameters to send to the API endpoint

	for the import service get import operation.

	Typically these are written to a http.Request.
*/
type ImportServiceGetImportParams struct {

	// ClusterID.
	//
	// Format: uint64
	ClusterID string

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the import service get import params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImportServiceGetImportParams) WithDefaults() *ImportServiceGetImportParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the import service get import params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImportServiceGetImportParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the import service get import params
func (o *ImportServiceGetImportParams) WithTimeout(timeout time.Duration) *ImportServiceGetImportParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the import service get import params
func (o *ImportServiceGetImportParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the import service get import params
func (o *ImportServiceGetImportParams) WithContext(ctx context.Context) *ImportServiceGetImportParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the import service get import params
func (o *ImportServiceGetImportParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the import service get import params
func (o *ImportServiceGetImportParams) WithHTTPClient(client *http.Client) *ImportServiceGetImportParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the import service get import params
func (o *ImportServiceGetImportParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the import service get import params
func (o *ImportServiceGetImportParams) WithClusterID(clusterID string) *ImportServiceGetImportParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the import service get import params
func (o *ImportServiceGetImportParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithID adds the id to the import service get import params
func (o *ImportServiceGetImportParams) WithID(id string) *ImportServiceGetImportParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the import service get import params
func (o *ImportServiceGetImportParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ImportServiceGetImportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param clusterId
	if err := r.SetPathParam("clusterId", o.ClusterID); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
