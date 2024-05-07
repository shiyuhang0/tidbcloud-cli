// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TargetS3Target target s3 target
//
// swagger:model TargetS3Target
type TargetS3Target struct {

	// Optional. The access_key of the s3.
	AccessKey *S3TargetAccessKey `json:"accessKey,omitempty"`

	// Optional. The bucket URI of the s3.
	BucketURI string `json:"bucketUri,omitempty"`

	// Optional. The uri of the s3 bucket.
	URI string `json:"uri,omitempty"`
}

// Validate validates this target s3 target
func (m *TargetS3Target) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccessKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TargetS3Target) validateAccessKey(formats strfmt.Registry) error {
	if swag.IsZero(m.AccessKey) { // not required
		return nil
	}

	if m.AccessKey != nil {
		if err := m.AccessKey.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("accessKey")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("accessKey")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this target s3 target based on the context it is used
func (m *TargetS3Target) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAccessKey(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TargetS3Target) contextValidateAccessKey(ctx context.Context, formats strfmt.Registry) error {

	if m.AccessKey != nil {

		if swag.IsZero(m.AccessKey) { // not required
			return nil
		}

		if err := m.AccessKey.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("accessKey")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("accessKey")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TargetS3Target) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TargetS3Target) UnmarshalBinary(b []byte) error {
	var res TargetS3Target
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
