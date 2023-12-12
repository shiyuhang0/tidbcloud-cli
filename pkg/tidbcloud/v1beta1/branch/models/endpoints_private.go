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

// EndpointsPrivate Message for Private Endpoint for this branch.
//
// swagger:model EndpointsPrivate
type EndpointsPrivate struct {

	// Message for AWS
	// Read Only: true
	Aws *PrivateAWS `json:"aws,omitempty"`

	// Message for GCP
	// Read Only: true
	Gcp *PrivateGCP `json:"gcp,omitempty"`

	// Output Only. Host Name of Public Endpoint.
	// Read Only: true
	Host string `json:"host,omitempty"`

	// Output Only. Port of Public Endpoint.
	// Read Only: true
	Port int32 `json:"port,omitempty"`
}

// Validate validates this endpoints private
func (m *EndpointsPrivate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAws(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGcp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EndpointsPrivate) validateAws(formats strfmt.Registry) error {
	if swag.IsZero(m.Aws) { // not required
		return nil
	}

	if m.Aws != nil {
		if err := m.Aws.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aws")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("aws")
			}
			return err
		}
	}

	return nil
}

func (m *EndpointsPrivate) validateGcp(formats strfmt.Registry) error {
	if swag.IsZero(m.Gcp) { // not required
		return nil
	}

	if m.Gcp != nil {
		if err := m.Gcp.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gcp")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("gcp")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this endpoints private based on the context it is used
func (m *EndpointsPrivate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAws(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGcp(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHost(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePort(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EndpointsPrivate) contextValidateAws(ctx context.Context, formats strfmt.Registry) error {

	if m.Aws != nil {

		if swag.IsZero(m.Aws) { // not required
			return nil
		}

		if err := m.Aws.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aws")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("aws")
			}
			return err
		}
	}

	return nil
}

func (m *EndpointsPrivate) contextValidateGcp(ctx context.Context, formats strfmt.Registry) error {

	if m.Gcp != nil {

		if swag.IsZero(m.Gcp) { // not required
			return nil
		}

		if err := m.Gcp.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gcp")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("gcp")
			}
			return err
		}
	}

	return nil
}

func (m *EndpointsPrivate) contextValidateHost(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "host", "body", string(m.Host)); err != nil {
		return err
	}

	return nil
}

func (m *EndpointsPrivate) contextValidatePort(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "port", "body", int32(m.Port)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EndpointsPrivate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EndpointsPrivate) UnmarshalBinary(b []byte) error {
	var res EndpointsPrivate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
