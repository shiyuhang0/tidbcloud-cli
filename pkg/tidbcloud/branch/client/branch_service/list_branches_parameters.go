// Code generated by go-swagger; DO NOT EDIT.

package branch_service

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

// NewListBranchesParams creates a new ListBranchesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListBranchesParams() *ListBranchesParams {
	return &ListBranchesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListBranchesParamsWithTimeout creates a new ListBranchesParams object
// with the ability to set a timeout on a request.
func NewListBranchesParamsWithTimeout(timeout time.Duration) *ListBranchesParams {
	return &ListBranchesParams{
		timeout: timeout,
	}
}

// NewListBranchesParamsWithContext creates a new ListBranchesParams object
// with the ability to set a context for a request.
func NewListBranchesParamsWithContext(ctx context.Context) *ListBranchesParams {
	return &ListBranchesParams{
		Context: ctx,
	}
}

// NewListBranchesParamsWithHTTPClient creates a new ListBranchesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListBranchesParamsWithHTTPClient(client *http.Client) *ListBranchesParams {
	return &ListBranchesParams{
		HTTPClient: client,
	}
}

/*
ListBranchesParams contains all the parameters to send to the API endpoint

	for the list branches operation.

	Typically these are written to a http.Request.
*/
type ListBranchesParams struct {

	/* ClusterID.

	   The ID of the cluster.
	*/
	ClusterID string

	/* PageSize.

	   The size of a page.

	   Format: int64
	   Default: 10
	*/
	PageSize *int64

	/* PageToken.

	   The number of pages.

	   Format: int64
	   Default: 1
	*/
	PageToken *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list branches params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListBranchesParams) WithDefaults() *ListBranchesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list branches params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListBranchesParams) SetDefaults() {
	var (
		pageSizeDefault = int64(10)

		pageTokenDefault = int64(1)
	)

	val := ListBranchesParams{
		PageSize:  &pageSizeDefault,
		PageToken: &pageTokenDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the list branches params
func (o *ListBranchesParams) WithTimeout(timeout time.Duration) *ListBranchesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list branches params
func (o *ListBranchesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list branches params
func (o *ListBranchesParams) WithContext(ctx context.Context) *ListBranchesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list branches params
func (o *ListBranchesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list branches params
func (o *ListBranchesParams) WithHTTPClient(client *http.Client) *ListBranchesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list branches params
func (o *ListBranchesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the list branches params
func (o *ListBranchesParams) WithClusterID(clusterID string) *ListBranchesParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the list branches params
func (o *ListBranchesParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithPageSize adds the pageSize to the list branches params
func (o *ListBranchesParams) WithPageSize(pageSize *int64) *ListBranchesParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the list branches params
func (o *ListBranchesParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithPageToken adds the pageToken to the list branches params
func (o *ListBranchesParams) WithPageToken(pageToken *int64) *ListBranchesParams {
	o.SetPageToken(pageToken)
	return o
}

// SetPageToken adds the pageToken to the list branches params
func (o *ListBranchesParams) SetPageToken(pageToken *int64) {
	o.PageToken = pageToken
}

// WriteToRequest writes these params to a swagger request
func (o *ListBranchesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	if o.PageSize != nil {

		// query param page_size
		var qrPageSize int64

		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt64(qrPageSize)
		if qPageSize != "" {

			if err := r.SetQueryParam("page_size", qPageSize); err != nil {
				return err
			}
		}
	}

	if o.PageToken != nil {

		// query param page_token
		var qrPageToken int64

		if o.PageToken != nil {
			qrPageToken = *o.PageToken
		}
		qPageToken := swag.FormatInt64(qrPageToken)
		if qPageToken != "" {

			if err := r.SetQueryParam("page_token", qPageToken); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}