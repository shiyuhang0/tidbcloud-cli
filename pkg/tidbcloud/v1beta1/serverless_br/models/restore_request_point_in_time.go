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

// RestoreRequestPointInTime restore request point in time
//
// swagger:model RestoreRequestPointInTime
type RestoreRequestPointInTime struct {

	// backup time
	// Format: date-time
	BackupTime strfmt.DateTime `json:"backupTime,omitempty"`

	// cluster Id
	ClusterID string `json:"clusterId,omitempty"`
}

// Validate validates this restore request point in time
func (m *RestoreRequestPointInTime) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBackupTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestoreRequestPointInTime) validateBackupTime(formats strfmt.Registry) error {
	if swag.IsZero(m.BackupTime) { // not required
		return nil
	}

	if err := validate.FormatOf("backupTime", "body", "date-time", m.BackupTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this restore request point in time based on context it is used
func (m *RestoreRequestPointInTime) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RestoreRequestPointInTime) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestoreRequestPointInTime) UnmarshalBinary(b []byte) error {
	var res RestoreRequestPointInTime
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
