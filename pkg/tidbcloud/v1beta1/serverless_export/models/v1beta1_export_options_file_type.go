// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// V1beta1ExportOptionsFileType  - TYPE_UNSPECIFIED: The type of the file is unknown.
//   - SQL: SQL type.
//   - CSV: CSV type.
//
// swagger:model v1beta1ExportOptionsFileType
type V1beta1ExportOptionsFileType string

func NewV1beta1ExportOptionsFileType(value V1beta1ExportOptionsFileType) *V1beta1ExportOptionsFileType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated V1beta1ExportOptionsFileType.
func (m V1beta1ExportOptionsFileType) Pointer() *V1beta1ExportOptionsFileType {
	return &m
}

const (

	// V1beta1ExportOptionsFileTypeSQL captures enum value "SQL"
	V1beta1ExportOptionsFileTypeSQL V1beta1ExportOptionsFileType = "SQL"

	// V1beta1ExportOptionsFileTypeCSV captures enum value "CSV"
	V1beta1ExportOptionsFileTypeCSV V1beta1ExportOptionsFileType = "CSV"
)

// for schema
var v1beta1ExportOptionsFileTypeEnum []interface{}

func init() {
	var res []V1beta1ExportOptionsFileType
	if err := json.Unmarshal([]byte(`["SQL","CSV"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1beta1ExportOptionsFileTypeEnum = append(v1beta1ExportOptionsFileTypeEnum, v)
	}
}

func (m V1beta1ExportOptionsFileType) validateV1beta1ExportOptionsFileTypeEnum(path, location string, value V1beta1ExportOptionsFileType) error {
	if err := validate.EnumCase(path, location, value, v1beta1ExportOptionsFileTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this v1beta1 export options file type
func (m V1beta1ExportOptionsFileType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateV1beta1ExportOptionsFileTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this v1beta1 export options file type based on context it is used
func (m V1beta1ExportOptionsFileType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}