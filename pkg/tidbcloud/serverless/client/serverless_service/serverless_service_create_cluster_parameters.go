// Code generated by go-swagger; DO NOT EDIT.

package serverless_service

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

	"tidbcloud-cli/pkg/tidbcloud/serverless/models"
)

// NewServerlessServiceCreateClusterParams creates a new ServerlessServiceCreateClusterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewServerlessServiceCreateClusterParams() *ServerlessServiceCreateClusterParams {
	return &ServerlessServiceCreateClusterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewServerlessServiceCreateClusterParamsWithTimeout creates a new ServerlessServiceCreateClusterParams object
// with the ability to set a timeout on a request.
func NewServerlessServiceCreateClusterParamsWithTimeout(timeout time.Duration) *ServerlessServiceCreateClusterParams {
	return &ServerlessServiceCreateClusterParams{
		timeout: timeout,
	}
}

// NewServerlessServiceCreateClusterParamsWithContext creates a new ServerlessServiceCreateClusterParams object
// with the ability to set a context for a request.
func NewServerlessServiceCreateClusterParamsWithContext(ctx context.Context) *ServerlessServiceCreateClusterParams {
	return &ServerlessServiceCreateClusterParams{
		Context: ctx,
	}
}

// NewServerlessServiceCreateClusterParamsWithHTTPClient creates a new ServerlessServiceCreateClusterParams object
// with the ability to set a custom HTTPClient for a request.
func NewServerlessServiceCreateClusterParamsWithHTTPClient(client *http.Client) *ServerlessServiceCreateClusterParams {
	return &ServerlessServiceCreateClusterParams{
		HTTPClient: client,
	}
}

/*
ServerlessServiceCreateClusterParams contains all the parameters to send to the API endpoint

	for the serverless service create cluster operation.

	Typically these are written to a http.Request.
*/
type ServerlessServiceCreateClusterParams struct {

	/* Cluster.

	   Required. The cluster to be created.
	*/
	Cluster *models.TidbCloudOpenApiserverlessv1beta1Cluster

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the serverless service create cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ServerlessServiceCreateClusterParams) WithDefaults() *ServerlessServiceCreateClusterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the serverless service create cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ServerlessServiceCreateClusterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the serverless service create cluster params
func (o *ServerlessServiceCreateClusterParams) WithTimeout(timeout time.Duration) *ServerlessServiceCreateClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the serverless service create cluster params
func (o *ServerlessServiceCreateClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the serverless service create cluster params
func (o *ServerlessServiceCreateClusterParams) WithContext(ctx context.Context) *ServerlessServiceCreateClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the serverless service create cluster params
func (o *ServerlessServiceCreateClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the serverless service create cluster params
func (o *ServerlessServiceCreateClusterParams) WithHTTPClient(client *http.Client) *ServerlessServiceCreateClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the serverless service create cluster params
func (o *ServerlessServiceCreateClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCluster adds the cluster to the serverless service create cluster params
func (o *ServerlessServiceCreateClusterParams) WithCluster(cluster *models.TidbCloudOpenApiserverlessv1beta1Cluster) *ServerlessServiceCreateClusterParams {
	o.SetCluster(cluster)
	return o
}

// SetCluster adds the cluster to the serverless service create cluster params
func (o *ServerlessServiceCreateClusterParams) SetCluster(cluster *models.TidbCloudOpenApiserverlessv1beta1Cluster) {
	o.Cluster = cluster
}

// WriteToRequest writes these params to a swagger request
func (o *ServerlessServiceCreateClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Cluster != nil {
		if err := r.SetBodyParam(o.Cluster); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
