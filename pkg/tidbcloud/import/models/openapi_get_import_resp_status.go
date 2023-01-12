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

// OpenapiGetImportRespStatus openapi get import resp status
//
// swagger:model openapiGetImportRespStatus
type OpenapiGetImportRespStatus string

func NewOpenapiGetImportRespStatus(value OpenapiGetImportRespStatus) *OpenapiGetImportRespStatus {
	return &value
}

// Pointer returns a pointer to a freshly-allocated OpenapiGetImportRespStatus.
func (m OpenapiGetImportRespStatus) Pointer() *OpenapiGetImportRespStatus {
	return &m
}

const (

	// OpenapiGetImportRespStatusPREPARING captures enum value "PREPARING"
	OpenapiGetImportRespStatusPREPARING OpenapiGetImportRespStatus = "PREPARING"

	// OpenapiGetImportRespStatusIMPORTING captures enum value "IMPORTING"
	OpenapiGetImportRespStatusIMPORTING OpenapiGetImportRespStatus = "IMPORTING"

	// OpenapiGetImportRespStatusCOMPLETED captures enum value "COMPLETED"
	OpenapiGetImportRespStatusCOMPLETED OpenapiGetImportRespStatus = "COMPLETED"

	// OpenapiGetImportRespStatusFAILED captures enum value "FAILED"
	OpenapiGetImportRespStatusFAILED OpenapiGetImportRespStatus = "FAILED"

	// OpenapiGetImportRespStatusCANCELING captures enum value "CANCELING"
	OpenapiGetImportRespStatusCANCELING OpenapiGetImportRespStatus = "CANCELING"

	// OpenapiGetImportRespStatusCANCELED captures enum value "CANCELED"
	OpenapiGetImportRespStatusCANCELED OpenapiGetImportRespStatus = "CANCELED"
)

// for schema
var openapiGetImportRespStatusEnum []interface{}

func init() {
	var res []OpenapiGetImportRespStatus
	if err := json.Unmarshal([]byte(`["PREPARING","IMPORTING","COMPLETED","FAILED","CANCELING","CANCELED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		openapiGetImportRespStatusEnum = append(openapiGetImportRespStatusEnum, v)
	}
}

func (m OpenapiGetImportRespStatus) validateOpenapiGetImportRespStatusEnum(path, location string, value OpenapiGetImportRespStatus) error {
	if err := validate.EnumCase(path, location, value, openapiGetImportRespStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this openapi get import resp status
func (m OpenapiGetImportRespStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateOpenapiGetImportRespStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this openapi get import resp status based on context it is used
func (m OpenapiGetImportRespStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}