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

// ExportOptionsFileType  - TYPE_UNSPECIFIED: The type of the file is unknown.
//   - SQL: SQL type.
//   - CSV: CSV type.
//
// swagger:model ExportOptionsFileType
type ExportOptionsFileType string

func NewExportOptionsFileType(value ExportOptionsFileType) *ExportOptionsFileType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ExportOptionsFileType.
func (m ExportOptionsFileType) Pointer() *ExportOptionsFileType {
	return &m
}

const (

	// ExportOptionsFileTypeSQL captures enum value "SQL"
	ExportOptionsFileTypeSQL ExportOptionsFileType = "SQL"

	// ExportOptionsFileTypeCSV captures enum value "CSV"
	ExportOptionsFileTypeCSV ExportOptionsFileType = "CSV"
)

// for schema
var exportOptionsFileTypeEnum []interface{}

func init() {
	var res []ExportOptionsFileType
	if err := json.Unmarshal([]byte(`["SQL","CSV"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		exportOptionsFileTypeEnum = append(exportOptionsFileTypeEnum, v)
	}
}

func (m ExportOptionsFileType) validateExportOptionsFileTypeEnum(path, location string, value ExportOptionsFileType) error {
	if err := validate.EnumCase(path, location, value, exportOptionsFileTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this export options file type
func (m ExportOptionsFileType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateExportOptionsFileTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this export options file type based on context it is used
func (m ExportOptionsFileType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}