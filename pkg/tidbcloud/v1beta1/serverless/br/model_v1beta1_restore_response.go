/*
TiDB Cloud Serverless Open API

TiDB Cloud Serverless Open API

API version: v1beta1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package br

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the V1beta1RestoreResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1beta1RestoreResponse{}

// V1beta1RestoreResponse struct for V1beta1RestoreResponse
type V1beta1RestoreResponse struct {
	ClusterId string `json:"clusterId"`
}

type _V1beta1RestoreResponse V1beta1RestoreResponse

// NewV1beta1RestoreResponse instantiates a new V1beta1RestoreResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1beta1RestoreResponse(clusterId string) *V1beta1RestoreResponse {
	this := V1beta1RestoreResponse{}
	this.ClusterId = clusterId
	return &this
}

// NewV1beta1RestoreResponseWithDefaults instantiates a new V1beta1RestoreResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1beta1RestoreResponseWithDefaults() *V1beta1RestoreResponse {
	this := V1beta1RestoreResponse{}
	return &this
}

// GetClusterId returns the ClusterId field value
func (o *V1beta1RestoreResponse) GetClusterId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value
// and a boolean to check if the value has been set.
func (o *V1beta1RestoreResponse) GetClusterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusterId, true
}

// SetClusterId sets field value
func (o *V1beta1RestoreResponse) SetClusterId(v string) {
	o.ClusterId = v
}

func (o V1beta1RestoreResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1beta1RestoreResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["clusterId"] = o.ClusterId
	return toSerialize, nil
}

func (o *V1beta1RestoreResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"clusterId",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varV1beta1RestoreResponse := _V1beta1RestoreResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV1beta1RestoreResponse)

	if err != nil {
		return err
	}

	*o = V1beta1RestoreResponse(varV1beta1RestoreResponse)

	return err
}

type NullableV1beta1RestoreResponse struct {
	value *V1beta1RestoreResponse
	isSet bool
}

func (v NullableV1beta1RestoreResponse) Get() *V1beta1RestoreResponse {
	return v.value
}

func (v *NullableV1beta1RestoreResponse) Set(val *V1beta1RestoreResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableV1beta1RestoreResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableV1beta1RestoreResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1beta1RestoreResponse(val *V1beta1RestoreResponse) *NullableV1beta1RestoreResponse {
	return &NullableV1beta1RestoreResponse{value: val, isSet: true}
}

func (v NullableV1beta1RestoreResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1beta1RestoreResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
