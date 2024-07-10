// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1beta1ExportOptionsCSVFormat v1beta1 export options c s v format
//
// swagger:model v1beta1ExportOptionsCSVFormat
type V1beta1ExportOptionsCSVFormat struct {

	// Delimiter of string type variables in CSV files. Default is '"'.
	Delimiter *string `json:"delimiter,omitempty"`

	// Representation of null values in CSV files. Default is "\N".
	NullValue *string `json:"nullValue,omitempty"`

	// Separator of each value in CSV files. It is recommended to use '|+|' or other uncommon character combinations. Default is ','.
	Separator string `json:"separator,omitempty"`

	// Export CSV files of the tables without header. Default is false.
	SkipHeader bool `json:"skipHeader,omitempty"`
}

// Validate validates this v1beta1 export options c s v format
func (m *V1beta1ExportOptionsCSVFormat) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1beta1 export options c s v format based on context it is used
func (m *V1beta1ExportOptionsCSVFormat) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1beta1ExportOptionsCSVFormat) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1beta1ExportOptionsCSVFormat) UnmarshalBinary(b []byte) error {
	var res V1beta1ExportOptionsCSVFormat
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}