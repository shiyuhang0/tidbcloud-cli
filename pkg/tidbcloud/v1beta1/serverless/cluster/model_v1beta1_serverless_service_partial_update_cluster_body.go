/*
TiDB Cloud Serverless Open API

TiDB Cloud Serverless Open API

API version: v1beta1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cluster

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the V1beta1ServerlessServicePartialUpdateClusterBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1beta1ServerlessServicePartialUpdateClusterBody{}

// V1beta1ServerlessServicePartialUpdateClusterBody Message for requesting a partial update on a TiDB Serverless cluster.
type V1beta1ServerlessServicePartialUpdateClusterBody struct {
	Cluster *RequiredTheClusterToBeUpdated `json:"cluster,omitempty"`
	// Required. The update mask indicating which fields are to be updated.
	UpdateMask string `json:"updateMask"`
}

type _V1beta1ServerlessServicePartialUpdateClusterBody V1beta1ServerlessServicePartialUpdateClusterBody

// NewV1beta1ServerlessServicePartialUpdateClusterBody instantiates a new V1beta1ServerlessServicePartialUpdateClusterBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1beta1ServerlessServicePartialUpdateClusterBody(updateMask string) *V1beta1ServerlessServicePartialUpdateClusterBody {
	this := V1beta1ServerlessServicePartialUpdateClusterBody{}
	this.UpdateMask = updateMask
	return &this
}

// NewV1beta1ServerlessServicePartialUpdateClusterBodyWithDefaults instantiates a new V1beta1ServerlessServicePartialUpdateClusterBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1beta1ServerlessServicePartialUpdateClusterBodyWithDefaults() *V1beta1ServerlessServicePartialUpdateClusterBody {
	this := V1beta1ServerlessServicePartialUpdateClusterBody{}
	return &this
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *V1beta1ServerlessServicePartialUpdateClusterBody) GetCluster() RequiredTheClusterToBeUpdated {
	if o == nil || IsNil(o.Cluster) {
		var ret RequiredTheClusterToBeUpdated
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1beta1ServerlessServicePartialUpdateClusterBody) GetClusterOk() (*RequiredTheClusterToBeUpdated, bool) {
	if o == nil || IsNil(o.Cluster) {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *V1beta1ServerlessServicePartialUpdateClusterBody) HasCluster() bool {
	if o != nil && !IsNil(o.Cluster) {
		return true
	}

	return false
}

// SetCluster gets a reference to the given RequiredTheClusterToBeUpdated and assigns it to the Cluster field.
func (o *V1beta1ServerlessServicePartialUpdateClusterBody) SetCluster(v RequiredTheClusterToBeUpdated) {
	o.Cluster = &v
}

// GetUpdateMask returns the UpdateMask field value
func (o *V1beta1ServerlessServicePartialUpdateClusterBody) GetUpdateMask() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdateMask
}

// GetUpdateMaskOk returns a tuple with the UpdateMask field value
// and a boolean to check if the value has been set.
func (o *V1beta1ServerlessServicePartialUpdateClusterBody) GetUpdateMaskOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdateMask, true
}

// SetUpdateMask sets field value
func (o *V1beta1ServerlessServicePartialUpdateClusterBody) SetUpdateMask(v string) {
	o.UpdateMask = v
}

func (o V1beta1ServerlessServicePartialUpdateClusterBody) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1beta1ServerlessServicePartialUpdateClusterBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cluster) {
		toSerialize["cluster"] = o.Cluster
	}
	toSerialize["updateMask"] = o.UpdateMask
	return toSerialize, nil
}

func (o *V1beta1ServerlessServicePartialUpdateClusterBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"updateMask",
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

	varV1beta1ServerlessServicePartialUpdateClusterBody := _V1beta1ServerlessServicePartialUpdateClusterBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV1beta1ServerlessServicePartialUpdateClusterBody)

	if err != nil {
		return err
	}

	*o = V1beta1ServerlessServicePartialUpdateClusterBody(varV1beta1ServerlessServicePartialUpdateClusterBody)

	return err
}

type NullableV1beta1ServerlessServicePartialUpdateClusterBody struct {
	value *V1beta1ServerlessServicePartialUpdateClusterBody
	isSet bool
}

func (v NullableV1beta1ServerlessServicePartialUpdateClusterBody) Get() *V1beta1ServerlessServicePartialUpdateClusterBody {
	return v.value
}

func (v *NullableV1beta1ServerlessServicePartialUpdateClusterBody) Set(val *V1beta1ServerlessServicePartialUpdateClusterBody) {
	v.value = val
	v.isSet = true
}

func (v NullableV1beta1ServerlessServicePartialUpdateClusterBody) IsSet() bool {
	return v.isSet
}

func (v *NullableV1beta1ServerlessServicePartialUpdateClusterBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1beta1ServerlessServicePartialUpdateClusterBody(val *V1beta1ServerlessServicePartialUpdateClusterBody) *NullableV1beta1ServerlessServicePartialUpdateClusterBody {
	return &NullableV1beta1ServerlessServicePartialUpdateClusterBody{value: val, isSet: true}
}

func (v NullableV1beta1ServerlessServicePartialUpdateClusterBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1beta1ServerlessServicePartialUpdateClusterBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}