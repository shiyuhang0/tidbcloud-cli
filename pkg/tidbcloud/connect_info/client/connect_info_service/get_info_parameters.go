// Code generated by go-swagger; DO NOT EDIT.

package connect_info_service

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

// NewGetInfoParams creates a new GetInfoParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetInfoParams() *GetInfoParams {
	return &GetInfoParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetInfoParamsWithTimeout creates a new GetInfoParams object
// with the ability to set a timeout on a request.
func NewGetInfoParamsWithTimeout(timeout time.Duration) *GetInfoParams {
	return &GetInfoParams{
		timeout: timeout,
	}
}

// NewGetInfoParamsWithContext creates a new GetInfoParams object
// with the ability to set a context for a request.
func NewGetInfoParamsWithContext(ctx context.Context) *GetInfoParams {
	return &GetInfoParams{
		Context: ctx,
	}
}

// NewGetInfoParamsWithHTTPClient creates a new GetInfoParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetInfoParamsWithHTTPClient(client *http.Client) *GetInfoParams {
	return &GetInfoParams{
		HTTPClient: client,
	}
}

/*
GetInfoParams contains all the parameters to send to the API endpoint

	for the get info operation.

	Typically these are written to a http.Request.
*/
type GetInfoParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get info params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetInfoParams) WithDefaults() *GetInfoParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get info params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetInfoParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get info params
func (o *GetInfoParams) WithTimeout(timeout time.Duration) *GetInfoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get info params
func (o *GetInfoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get info params
func (o *GetInfoParams) WithContext(ctx context.Context) *GetInfoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get info params
func (o *GetInfoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get info params
func (o *GetInfoParams) WithHTTPClient(client *http.Client) *GetInfoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get info params
func (o *GetInfoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetInfoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
