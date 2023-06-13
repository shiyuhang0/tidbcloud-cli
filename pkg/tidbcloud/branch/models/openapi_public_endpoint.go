// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OpenapiPublicEndpoint PublicEndpoint is the public endpoint of the branch.
//
// swagger:model openapiPublicEndpoint
type OpenapiPublicEndpoint struct {

	// The disabled of the public endpoint.
	Disabled bool `json:"disabled,omitempty"`

	// The host of the public endpoint.
	Host string `json:"host,omitempty"`

	// The port of the public endpoint.
	Port int32 `json:"port,omitempty"`
}

// Validate validates this openapi public endpoint
func (m *OpenapiPublicEndpoint) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this openapi public endpoint based on context it is used
func (m *OpenapiPublicEndpoint) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OpenapiPublicEndpoint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenapiPublicEndpoint) UnmarshalBinary(b []byte) error {
	var res OpenapiPublicEndpoint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
