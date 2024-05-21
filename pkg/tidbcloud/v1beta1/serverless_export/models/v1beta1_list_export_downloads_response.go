// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1beta1ListExportDownloadsResponse v1beta1 list export downloads response
//
// swagger:model v1beta1ListExportDownloadsResponse
type V1beta1ListExportDownloadsResponse struct {

	// The download urls of the export.
	Downloads []*V1beta1DownloadURL `json:"downloads"`

	// Token provided to retrieve the next page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Total number of backups.
	TotalSize int64 `json:"totalSize,omitempty"`
}

// Validate validates this v1beta1 list export downloads response
func (m *V1beta1ListExportDownloadsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDownloads(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1beta1ListExportDownloadsResponse) validateDownloads(formats strfmt.Registry) error {
	if swag.IsZero(m.Downloads) { // not required
		return nil
	}

	for i := 0; i < len(m.Downloads); i++ {
		if swag.IsZero(m.Downloads[i]) { // not required
			continue
		}

		if m.Downloads[i] != nil {
			if err := m.Downloads[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("downloads" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("downloads" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this v1beta1 list export downloads response based on the context it is used
func (m *V1beta1ListExportDownloadsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDownloads(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1beta1ListExportDownloadsResponse) contextValidateDownloads(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Downloads); i++ {

		if m.Downloads[i] != nil {

			if swag.IsZero(m.Downloads[i]) { // not required
				return nil
			}

			if err := m.Downloads[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("downloads" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("downloads" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1beta1ListExportDownloadsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1beta1ListExportDownloadsResponse) UnmarshalBinary(b []byte) error {
	var res V1beta1ListExportDownloadsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
