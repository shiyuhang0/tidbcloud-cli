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

// V1beta1ImportOptionsFileType  - TYPE_UNSPECIFIED: The type of the file is unknown.
//   - CSV: CSV type.
//
// swagger:model v1beta1ImportOptionsFileType
type V1beta1ImportOptionsFileType string

func NewV1beta1ImportOptionsFileType(value V1beta1ImportOptionsFileType) *V1beta1ImportOptionsFileType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated V1beta1ImportOptionsFileType.
func (m V1beta1ImportOptionsFileType) Pointer() *V1beta1ImportOptionsFileType {
	return &m
}

const (

	// V1beta1ImportOptionsFileTypeCSV captures enum value "CSV"
	V1beta1ImportOptionsFileTypeCSV V1beta1ImportOptionsFileType = "CSV"
)

// for schema
var v1beta1ImportOptionsFileTypeEnum []interface{}

func init() {
	var res []V1beta1ImportOptionsFileType
	if err := json.Unmarshal([]byte(`["CSV"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1beta1ImportOptionsFileTypeEnum = append(v1beta1ImportOptionsFileTypeEnum, v)
	}
}

func (m V1beta1ImportOptionsFileType) validateV1beta1ImportOptionsFileTypeEnum(path, location string, value V1beta1ImportOptionsFileType) error {
	if err := validate.EnumCase(path, location, value, v1beta1ImportOptionsFileTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this v1beta1 import options file type
func (m V1beta1ImportOptionsFileType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateV1beta1ImportOptionsFileTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this v1beta1 import options file type based on context it is used
func (m V1beta1ImportOptionsFileType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
