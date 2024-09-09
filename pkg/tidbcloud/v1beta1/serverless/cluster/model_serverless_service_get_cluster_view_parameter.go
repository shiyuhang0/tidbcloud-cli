/*
TiDB Cloud Serverless Open API

TiDB Cloud Serverless Open API

API version: v1beta1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cluster

import (
	"encoding/json"
	"fmt"
)

// ServerlessServiceGetClusterViewParameter the model 'ServerlessServiceGetClusterViewParameter'
type ServerlessServiceGetClusterViewParameter string

// List of ServerlessService_GetCluster_view_parameter
const (
	SERVERLESSSERVICEGETCLUSTERVIEWPARAMETER_BASIC ServerlessServiceGetClusterViewParameter = "BASIC"
	SERVERLESSSERVICEGETCLUSTERVIEWPARAMETER_FULL  ServerlessServiceGetClusterViewParameter = "FULL"
)

// All allowed values of ServerlessServiceGetClusterViewParameter enum
var AllowedServerlessServiceGetClusterViewParameterEnumValues = []ServerlessServiceGetClusterViewParameter{
	"BASIC",
	"FULL",
}

func (v *ServerlessServiceGetClusterViewParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ServerlessServiceGetClusterViewParameter(value)
	for _, existing := range AllowedServerlessServiceGetClusterViewParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ServerlessServiceGetClusterViewParameter", value)
}

// NewServerlessServiceGetClusterViewParameterFromValue returns a pointer to a valid ServerlessServiceGetClusterViewParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewServerlessServiceGetClusterViewParameterFromValue(v string) (*ServerlessServiceGetClusterViewParameter, error) {
	ev := ServerlessServiceGetClusterViewParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ServerlessServiceGetClusterViewParameter: valid values are %v", v, AllowedServerlessServiceGetClusterViewParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ServerlessServiceGetClusterViewParameter) IsValid() bool {
	for _, existing := range AllowedServerlessServiceGetClusterViewParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ServerlessService_GetCluster_view_parameter value
func (v ServerlessServiceGetClusterViewParameter) Ptr() *ServerlessServiceGetClusterViewParameter {
	return &v
}

type NullableServerlessServiceGetClusterViewParameter struct {
	value *ServerlessServiceGetClusterViewParameter
	isSet bool
}

func (v NullableServerlessServiceGetClusterViewParameter) Get() *ServerlessServiceGetClusterViewParameter {
	return v.value
}

func (v *NullableServerlessServiceGetClusterViewParameter) Set(val *ServerlessServiceGetClusterViewParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableServerlessServiceGetClusterViewParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableServerlessServiceGetClusterViewParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerlessServiceGetClusterViewParameter(val *ServerlessServiceGetClusterViewParameter) *NullableServerlessServiceGetClusterViewParameter {
	return &NullableServerlessServiceGetClusterViewParameter{value: val, isSet: true}
}

func (v NullableServerlessServiceGetClusterViewParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerlessServiceGetClusterViewParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
