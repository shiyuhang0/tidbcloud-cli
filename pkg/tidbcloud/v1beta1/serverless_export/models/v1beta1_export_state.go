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

// V1beta1ExportState Output Only. Export State.
//
//   - STATE_UNSPECIFIED: The state of the export is unknown.
//   - RUNNING: The export job is being created.
//   - SUCCEEDED: The export job is success.
//   - FAILED: The export job is failed.
//   - CANCELED: The export job is canceled.
//   - DELETED: The export job is deleted.
//
// swagger:model v1beta1ExportState
type V1beta1ExportState string

func NewV1beta1ExportState(value V1beta1ExportState) *V1beta1ExportState {
	return &value
}

// Pointer returns a pointer to a freshly-allocated V1beta1ExportState.
func (m V1beta1ExportState) Pointer() *V1beta1ExportState {
	return &m
}

const (

	// V1beta1ExportStateRUNNING captures enum value "RUNNING"
	V1beta1ExportStateRUNNING V1beta1ExportState = "RUNNING"

	// V1beta1ExportStateSUCCEEDED captures enum value "SUCCEEDED"
	V1beta1ExportStateSUCCEEDED V1beta1ExportState = "SUCCEEDED"

	// V1beta1ExportStateFAILED captures enum value "FAILED"
	V1beta1ExportStateFAILED V1beta1ExportState = "FAILED"

	// V1beta1ExportStateCANCELED captures enum value "CANCELED"
	V1beta1ExportStateCANCELED V1beta1ExportState = "CANCELED"

	// V1beta1ExportStateDELETED captures enum value "DELETED"
	V1beta1ExportStateDELETED V1beta1ExportState = "DELETED"
)

// for schema
var v1beta1ExportStateEnum []interface{}

func init() {
	var res []V1beta1ExportState
	if err := json.Unmarshal([]byte(`["RUNNING","SUCCEEDED","FAILED","CANCELED","DELETED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1beta1ExportStateEnum = append(v1beta1ExportStateEnum, v)
	}
}

func (m V1beta1ExportState) validateV1beta1ExportStateEnum(path, location string, value V1beta1ExportState) error {
	if err := validate.EnumCase(path, location, value, v1beta1ExportStateEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this v1beta1 export state
func (m V1beta1ExportState) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateV1beta1ExportStateEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this v1beta1 export state based on context it is used
func (m V1beta1ExportState) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
