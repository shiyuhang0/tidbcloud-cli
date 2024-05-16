// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// FilterTable filter table
//
// swagger:model FilterTable
type FilterTable struct {

	// Optional. The table filter expressions.
	Patterns []string `json:"patterns"`

	// Optional. The specify table of the export.
	Where string `json:"where,omitempty"`
}

// Validate validates this filter table
func (m *FilterTable) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this filter table based on context it is used
func (m *FilterTable) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FilterTable) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FilterTable) UnmarshalBinary(b []byte) error {
	var res FilterTable
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}