/*
TiDB Cloud Serverless Open API

TiDB Cloud Serverless Open API

API version: v1beta1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package br

import (
	"encoding/json"
)

// checks if the V1beta1RestoreRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1beta1RestoreRequest{}

// V1beta1RestoreRequest struct for V1beta1RestoreRequest
type V1beta1RestoreRequest struct {
	Snapshot             *RestoreRequestSnapshot    `json:"snapshot,omitempty"`
	PointInTime          *RestoreRequestPointInTime `json:"pointInTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _V1beta1RestoreRequest V1beta1RestoreRequest

// NewV1beta1RestoreRequest instantiates a new V1beta1RestoreRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1beta1RestoreRequest() *V1beta1RestoreRequest {
	this := V1beta1RestoreRequest{}
	return &this
}

// NewV1beta1RestoreRequestWithDefaults instantiates a new V1beta1RestoreRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1beta1RestoreRequestWithDefaults() *V1beta1RestoreRequest {
	this := V1beta1RestoreRequest{}
	return &this
}

// GetSnapshot returns the Snapshot field value if set, zero value otherwise.
func (o *V1beta1RestoreRequest) GetSnapshot() RestoreRequestSnapshot {
	if o == nil || IsNil(o.Snapshot) {
		var ret RestoreRequestSnapshot
		return ret
	}
	return *o.Snapshot
}

// GetSnapshotOk returns a tuple with the Snapshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1beta1RestoreRequest) GetSnapshotOk() (*RestoreRequestSnapshot, bool) {
	if o == nil || IsNil(o.Snapshot) {
		return nil, false
	}
	return o.Snapshot, true
}

// HasSnapshot returns a boolean if a field has been set.
func (o *V1beta1RestoreRequest) HasSnapshot() bool {
	if o != nil && !IsNil(o.Snapshot) {
		return true
	}

	return false
}

// SetSnapshot gets a reference to the given RestoreRequestSnapshot and assigns it to the Snapshot field.
func (o *V1beta1RestoreRequest) SetSnapshot(v RestoreRequestSnapshot) {
	o.Snapshot = &v
}

// GetPointInTime returns the PointInTime field value if set, zero value otherwise.
func (o *V1beta1RestoreRequest) GetPointInTime() RestoreRequestPointInTime {
	if o == nil || IsNil(o.PointInTime) {
		var ret RestoreRequestPointInTime
		return ret
	}
	return *o.PointInTime
}

// GetPointInTimeOk returns a tuple with the PointInTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1beta1RestoreRequest) GetPointInTimeOk() (*RestoreRequestPointInTime, bool) {
	if o == nil || IsNil(o.PointInTime) {
		return nil, false
	}
	return o.PointInTime, true
}

// HasPointInTime returns a boolean if a field has been set.
func (o *V1beta1RestoreRequest) HasPointInTime() bool {
	if o != nil && !IsNil(o.PointInTime) {
		return true
	}

	return false
}

// SetPointInTime gets a reference to the given RestoreRequestPointInTime and assigns it to the PointInTime field.
func (o *V1beta1RestoreRequest) SetPointInTime(v RestoreRequestPointInTime) {
	o.PointInTime = &v
}

func (o V1beta1RestoreRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1beta1RestoreRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Snapshot) {
		toSerialize["snapshot"] = o.Snapshot
	}
	if !IsNil(o.PointInTime) {
		toSerialize["pointInTime"] = o.PointInTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *V1beta1RestoreRequest) UnmarshalJSON(data []byte) (err error) {
	varV1beta1RestoreRequest := _V1beta1RestoreRequest{}

	err = json.Unmarshal(data, &varV1beta1RestoreRequest)

	if err != nil {
		return err
	}

	*o = V1beta1RestoreRequest(varV1beta1RestoreRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "snapshot")
		delete(additionalProperties, "pointInTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableV1beta1RestoreRequest struct {
	value *V1beta1RestoreRequest
	isSet bool
}

func (v NullableV1beta1RestoreRequest) Get() *V1beta1RestoreRequest {
	return v.value
}

func (v *NullableV1beta1RestoreRequest) Set(val *V1beta1RestoreRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableV1beta1RestoreRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableV1beta1RestoreRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1beta1RestoreRequest(val *V1beta1RestoreRequest) *NullableV1beta1RestoreRequest {
	return &NullableV1beta1RestoreRequest{value: val, isSet: true}
}

func (v NullableV1beta1RestoreRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1beta1RestoreRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
