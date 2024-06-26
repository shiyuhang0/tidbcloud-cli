// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// APICreateSQLUserReq api create Sql user req
//
// swagger:model api.CreateSqlUserReq
type APICreateSQLUserReq struct {

	// available values [mysql_native_password] .
	AuthMethod string `json:"authMethod,omitempty"`

	// if autoPrefix is true ,username and  builtinRole will automatically add the serverless token prefix.
	AutoPrefix bool `json:"autoPrefix,omitempty"`

	// The builtinRole of the sql user,available values [role_admin,role_readonly,role_readwrite] . if cluster is serverless and autoPrefix is false, the builtinRole[role_readonly,role_readwrite] must be start with serverless token.
	BuiltinRole string `json:"builtinRole,omitempty"`

	// if cluster is serverless ,customRoles roles do not need to be prefixed.
	CustomRoles []string `json:"customRoles"`

	// password
	Password string `json:"password,omitempty"`

	// The username of the sql user, if cluster is serverless and autoPrefix is false, the userName must be start with serverless token.
	UserName string `json:"userName,omitempty"`
}

// Validate validates this api create Sql user req
func (m *APICreateSQLUserReq) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this api create Sql user req based on context it is used
func (m *APICreateSQLUserReq) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *APICreateSQLUserReq) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APICreateSQLUserReq) UnmarshalBinary(b []byte) error {
	var res APICreateSQLUserReq
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
