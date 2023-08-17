// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ClusterSpendingLimit Message for Spend Limit for this cluster.
//
// swagger:model ClusterSpendingLimit
type ClusterSpendingLimit struct {

	// Required. Monthly Spend Limit
	Monthly int32 `json:"monthly,omitempty"`
}

// Validate validates this cluster spending limit
func (m *ClusterSpendingLimit) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cluster spending limit based on context it is used
func (m *ClusterSpendingLimit) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClusterSpendingLimit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterSpendingLimit) UnmarshalBinary(b []byte) error {
	var res ClusterSpendingLimit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}