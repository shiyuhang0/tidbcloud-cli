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

// V1beta1Branch Message for branch
//
// swagger:model v1beta1Branch
type V1beta1Branch struct {

	// Optional. The annotations of this branch. Only display in FULL view.
	Annotations map[string]string `json:"annotations,omitempty"`

	// Output only. The system-generated ID of the resource.
	// Read Only: true
	BranchID string `json:"branchId,omitempty"`

	// Output only. The cluster ID of this branch.
	// Read Only: true
	ClusterID string `json:"clusterId,omitempty"`

	// Output only. Create timestamp
	// Read Only: true
	// Format: date-time
	CreateTime strfmt.DateTime `json:"createTime,omitempty"`

	// Output only. The creator of the branch.
	// Read Only: true
	CreatedBy string `json:"createdBy,omitempty"`

	// Required. User-settable and human-readable display name for the branch.
	// Required: true
	DisplayName *string `json:"displayName"`

	// Optional. The endpoints of this branch.
	Endpoints *BranchEndpoints `json:"endpoints,omitempty"`

	// Output Only. The name of the resource.
	// Read Only: true
	Name string `json:"name,omitempty"`

	// OPTIONAL. The parent display name of this branch.
	ParentDisplayName string `json:"parentDisplayName,omitempty"`

	// OPTIONAL. The parent ID of this branch.
	ParentID string `json:"parentId,omitempty"`

	// Optional. The point in time on the parent branch the branch will be created from.
	// Format: date-time
	ParentTimestamp *strfmt.DateTime `json:"parentTimestamp,omitempty"`

	// Output only. The state of this branch.
	// Read Only: true
	State V1beta1BranchState `json:"state,omitempty"`

	// Output only. Update timestamp
	// Read Only: true
	// Format: date-time
	UpdateTime strfmt.DateTime `json:"updateTime,omitempty"`

	// OPTIONAL. Usage metrics of this branch. Only display in FULL view.
	Usage *BranchUsage `json:"usage,omitempty"`

	// Output only. User name prefix of this branch. For each TiDB Serverless branch,
	// TiDB Cloud generates a unique prefix to distinguish it from other branches.
	// Whenever you use or set a database user name, you must include the prefix in the user name.
	// Read Only: true
	UserPrefix *string `json:"userPrefix,omitempty"`
}

// Validate validates this v1beta1 branch
func (m *V1beta1Branch) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisplayName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndpoints(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParentTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1beta1Branch) validateCreateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.CreateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("createTime", "body", "date-time", m.CreateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1beta1Branch) validateDisplayName(formats strfmt.Registry) error {

	if err := validate.Required("displayName", "body", m.DisplayName); err != nil {
		return err
	}

	return nil
}

func (m *V1beta1Branch) validateEndpoints(formats strfmt.Registry) error {
	if swag.IsZero(m.Endpoints) { // not required
		return nil
	}

	if m.Endpoints != nil {
		if err := m.Endpoints.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("endpoints")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("endpoints")
			}
			return err
		}
	}

	return nil
}

func (m *V1beta1Branch) validateParentTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.ParentTimestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("parentTimestamp", "body", "date-time", m.ParentTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1beta1Branch) validateState(formats strfmt.Registry) error {
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

func (m *V1beta1Branch) validateUpdateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("updateTime", "body", "date-time", m.UpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1beta1Branch) validateUsage(formats strfmt.Registry) error {
	if swag.IsZero(m.Usage) { // not required
		return nil
	}

	if m.Usage != nil {
		if err := m.Usage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("usage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("usage")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1beta1 branch based on the context it is used
func (m *V1beta1Branch) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBranchID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateClusterID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreateTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreatedBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEndpoints(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdateTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUsage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUserPrefix(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1beta1Branch) contextValidateBranchID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "branchId", "body", string(m.BranchID)); err != nil {
		return err
	}

	return nil
}

func (m *V1beta1Branch) contextValidateClusterID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "clusterId", "body", string(m.ClusterID)); err != nil {
		return err
	}

	return nil
}

func (m *V1beta1Branch) contextValidateCreateTime(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "createTime", "body", strfmt.DateTime(m.CreateTime)); err != nil {
		return err
	}

	return nil
}

func (m *V1beta1Branch) contextValidateCreatedBy(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "createdBy", "body", string(m.CreatedBy)); err != nil {
		return err
	}

	return nil
}

func (m *V1beta1Branch) contextValidateEndpoints(ctx context.Context, formats strfmt.Registry) error {

	if m.Endpoints != nil {

		if swag.IsZero(m.Endpoints) { // not required
			return nil
		}

		if err := m.Endpoints.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("endpoints")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("endpoints")
			}
			return err
		}
	}

	return nil
}

func (m *V1beta1Branch) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "name", "body", string(m.Name)); err != nil {
		return err
	}

	return nil
}

func (m *V1beta1Branch) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

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

func (m *V1beta1Branch) contextValidateUpdateTime(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "updateTime", "body", strfmt.DateTime(m.UpdateTime)); err != nil {
		return err
	}

	return nil
}

func (m *V1beta1Branch) contextValidateUsage(ctx context.Context, formats strfmt.Registry) error {

	if m.Usage != nil {

		if swag.IsZero(m.Usage) { // not required
			return nil
		}

		if err := m.Usage.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("usage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("usage")
			}
			return err
		}
	}

	return nil
}

func (m *V1beta1Branch) contextValidateUserPrefix(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "userPrefix", "body", m.UserPrefix); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1beta1Branch) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1beta1Branch) UnmarshalBinary(b []byte) error {
	var res V1beta1Branch
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
