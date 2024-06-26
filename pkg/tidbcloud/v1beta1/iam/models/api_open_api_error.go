// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// APIOpenAPIError api open Api error
//
// swagger:model api.OpenApiError
type APIOpenAPIError struct {

	// code
	Code string `json:"code,omitempty"`

	// error
	Error interface{} `json:"error,omitempty"`

	// msg prefix
	MsgPrefix string `json:"msgPrefix,omitempty"`

	// status
	Status int64 `json:"status,omitempty"`
}

// Validate validates this api open Api error
func (m *APIOpenAPIError) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this api open Api error based on context it is used
func (m *APIOpenAPIError) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *APIOpenAPIError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIOpenAPIError) UnmarshalBinary(b []byte) error {
	var res APIOpenAPIError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
