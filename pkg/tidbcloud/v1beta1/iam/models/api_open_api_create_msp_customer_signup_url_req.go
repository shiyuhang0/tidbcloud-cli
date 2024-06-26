// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// APIOpenAPICreateMspCustomerSignupURLReq api open Api create msp customer signup Url req
//
// swagger:model api.OpenApiCreateMspCustomerSignupUrlReq
type APIOpenAPICreateMspCustomerSignupURLReq struct {

	// The ID of the MSP.
	// Example: 123456
	OrgID string `json:"orgId,omitempty"`
}

// Validate validates this api open Api create msp customer signup Url req
func (m *APIOpenAPICreateMspCustomerSignupURLReq) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this api open Api create msp customer signup Url req based on context it is used
func (m *APIOpenAPICreateMspCustomerSignupURLReq) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *APIOpenAPICreateMspCustomerSignupURLReq) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIOpenAPICreateMspCustomerSignupURLReq) UnmarshalBinary(b []byte) error {
	var res APIOpenAPICreateMspCustomerSignupURLReq
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
