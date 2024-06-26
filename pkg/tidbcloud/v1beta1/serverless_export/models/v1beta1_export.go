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

// V1beta1Export Message for export resource.
//
// swagger:model v1beta1Export
type V1beta1Export struct {

	// Required. The cluster ID that export belong to.
	// Required: true
	ClusterID *string `json:"clusterId"`

	// Output_only. Timestamp when the export was completed.
	// Read Only: true
	// Format: date-time
	CompleteTime *strfmt.DateTime `json:"completeTime,omitempty"`

	// Output_only. Timestamp when the export was created.
	// Read Only: true
	// Format: date-time
	CreateTime strfmt.DateTime `json:"createTime,omitempty"`

	// Output_only. The creator of the export.
	// Read Only: true
	CreatedBy string `json:"createdBy,omitempty"`

	// OUTPUT_ONLY. Expire time of the export.
	// Read Only: true
	// Format: date-time
	ExpireTime *strfmt.DateTime `json:"expireTime,omitempty"`

	// Output_only. The unique ID of the export.
	// Read Only: true
	ExportID string `json:"exportId,omitempty"`

	// Optional. The options of the export.
	ExportOptions *V1beta1ExportOptions `json:"exportOptions,omitempty"`

	// Output_only. The unique name of the export.
	// Read Only: true
	Name string `json:"name,omitempty"`

	// Optional. The failed reason of the export.
	Reason *string `json:"reason,omitempty"`

	// OUTPUT_ONLY. Snapshot time of the export.
	// Read Only: true
	// Format: date-time
	SnapshotTime *strfmt.DateTime `json:"snapshotTime,omitempty"`

	// Output_only. The state of the export.
	// Read Only: true
	State V1beta1ExportState `json:"state,omitempty"`

	// Optional. The target of the export.
	Target *V1beta1Target `json:"target,omitempty"`

	// Output_only. Timestamp when the export was updated.
	// Read Only: true
	// Format: date-time
	UpdateTime *strfmt.DateTime `json:"updateTime,omitempty"`
}

// Validate validates this v1beta1 export
func (m *V1beta1Export) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCompleteTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpireTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExportOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSnapshotTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTarget(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1beta1Export) validateClusterID(formats strfmt.Registry) error {

	if err := validate.Required("clusterId", "body", m.ClusterID); err != nil {
		return err
	}

	return nil
}

func (m *V1beta1Export) validateCompleteTime(formats strfmt.Registry) error {
	if swag.IsZero(m.CompleteTime) { // not required
		return nil
	}

	if err := validate.FormatOf("completeTime", "body", "date-time", m.CompleteTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1beta1Export) validateCreateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.CreateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("createTime", "body", "date-time", m.CreateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1beta1Export) validateExpireTime(formats strfmt.Registry) error {
	if swag.IsZero(m.ExpireTime) { // not required
		return nil
	}

	if err := validate.FormatOf("expireTime", "body", "date-time", m.ExpireTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1beta1Export) validateExportOptions(formats strfmt.Registry) error {
	if swag.IsZero(m.ExportOptions) { // not required
		return nil
	}

	if m.ExportOptions != nil {
		if err := m.ExportOptions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("exportOptions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("exportOptions")
			}
			return err
		}
	}

	return nil
}

func (m *V1beta1Export) validateSnapshotTime(formats strfmt.Registry) error {
	if swag.IsZero(m.SnapshotTime) { // not required
		return nil
	}

	if err := validate.FormatOf("snapshotTime", "body", "date-time", m.SnapshotTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1beta1Export) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	if err := m.State.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("state")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("state")
		}
		return err
	}

	return nil
}

func (m *V1beta1Export) validateTarget(formats strfmt.Registry) error {
	if swag.IsZero(m.Target) { // not required
		return nil
	}

	if m.Target != nil {
		if err := m.Target.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("target")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("target")
			}
			return err
		}
	}

	return nil
}

func (m *V1beta1Export) validateUpdateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("updateTime", "body", "date-time", m.UpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this v1beta1 export based on the context it is used
func (m *V1beta1Export) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCompleteTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreateTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreatedBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExpireTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExportID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExportOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSnapshotTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTarget(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdateTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1beta1Export) contextValidateCompleteTime(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "completeTime", "body", m.CompleteTime); err != nil {
		return err
	}

	return nil
}

func (m *V1beta1Export) contextValidateCreateTime(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "createTime", "body", strfmt.DateTime(m.CreateTime)); err != nil {
		return err
	}

	return nil
}

func (m *V1beta1Export) contextValidateCreatedBy(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "createdBy", "body", string(m.CreatedBy)); err != nil {
		return err
	}

	return nil
}

func (m *V1beta1Export) contextValidateExpireTime(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "expireTime", "body", m.ExpireTime); err != nil {
		return err
	}

	return nil
}

func (m *V1beta1Export) contextValidateExportID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "exportId", "body", string(m.ExportID)); err != nil {
		return err
	}

	return nil
}

func (m *V1beta1Export) contextValidateExportOptions(ctx context.Context, formats strfmt.Registry) error {

	if m.ExportOptions != nil {

		if swag.IsZero(m.ExportOptions) { // not required
			return nil
		}

		if err := m.ExportOptions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("exportOptions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("exportOptions")
			}
			return err
		}
	}

	return nil
}

func (m *V1beta1Export) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "name", "body", string(m.Name)); err != nil {
		return err
	}

	return nil
}

func (m *V1beta1Export) contextValidateSnapshotTime(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "snapshotTime", "body", m.SnapshotTime); err != nil {
		return err
	}

	return nil
}

func (m *V1beta1Export) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	if err := m.State.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("state")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("state")
		}
		return err
	}

	return nil
}

func (m *V1beta1Export) contextValidateTarget(ctx context.Context, formats strfmt.Registry) error {

	if m.Target != nil {

		if swag.IsZero(m.Target) { // not required
			return nil
		}

		if err := m.Target.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("target")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("target")
			}
			return err
		}
	}

	return nil
}

func (m *V1beta1Export) contextValidateUpdateTime(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "updateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1beta1Export) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1beta1Export) UnmarshalBinary(b []byte) error {
	var res V1beta1Export
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
