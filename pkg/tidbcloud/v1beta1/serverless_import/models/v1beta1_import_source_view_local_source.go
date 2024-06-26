// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1beta1ImportSourceViewLocalSource v1beta1 import source view local source
//
// swagger:model v1beta1ImportSourceViewLocalSource
type V1beta1ImportSourceViewLocalSource struct {

	// Optional. The file name of import source file.
	// Read Only: true
	FileName string `json:"fileName,omitempty"`

	// target database
	// Read Only: true
	TargetDatabase string `json:"targetDatabase,omitempty"`

	// target table
	// Read Only: true
	TargetTable string `json:"targetTable,omitempty"`

	// Optional. The upload id of import source file.
	// Read Only: true
	UploadID string `json:"uploadId,omitempty"`
}

// Validate validates this v1beta1 import source view local source
func (m *V1beta1ImportSourceViewLocalSource) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this v1beta1 import source view local source based on the context it is used
func (m *V1beta1ImportSourceViewLocalSource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFileName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTargetDatabase(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTargetTable(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUploadID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1beta1ImportSourceViewLocalSource) contextValidateFileName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "fileName", "body", string(m.FileName)); err != nil {
		return err
	}

	return nil
}

func (m *V1beta1ImportSourceViewLocalSource) contextValidateTargetDatabase(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "targetDatabase", "body", string(m.TargetDatabase)); err != nil {
		return err
	}

	return nil
}

func (m *V1beta1ImportSourceViewLocalSource) contextValidateTargetTable(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "targetTable", "body", string(m.TargetTable)); err != nil {
		return err
	}

	return nil
}

func (m *V1beta1ImportSourceViewLocalSource) contextValidateUploadID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "uploadId", "body", string(m.UploadID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1beta1ImportSourceViewLocalSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1beta1ImportSourceViewLocalSource) UnmarshalBinary(b []byte) error {
	var res V1beta1ImportSourceViewLocalSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
