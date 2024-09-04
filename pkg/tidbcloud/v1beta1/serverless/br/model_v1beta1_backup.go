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
	"time"
)

// checks if the V1beta1Backup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1beta1Backup{}

// V1beta1Backup Message for backup resource.
type V1beta1Backup struct {
	// Output_only. The unique name of the backup.
	Name *string `json:"name,omitempty"`
	// Output_only. The unique ID of the backup.
	BackupId *string `json:"backupId,omitempty"`
	// Required. The cluster ID that backup belong to.
	ClusterId string `json:"clusterId"`
	// Output_only. Timestamp when the backup was created.
	CreateTime *time.Time `json:"createTime,omitempty"`
}

type _V1beta1Backup V1beta1Backup

// NewV1beta1Backup instantiates a new V1beta1Backup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1beta1Backup(clusterId string) *V1beta1Backup {
	this := V1beta1Backup{}
	this.ClusterId = clusterId
	return &this
}

// NewV1beta1BackupWithDefaults instantiates a new V1beta1Backup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1beta1BackupWithDefaults() *V1beta1Backup {
	this := V1beta1Backup{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V1beta1Backup) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1beta1Backup) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V1beta1Backup) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V1beta1Backup) SetName(v string) {
	o.Name = &v
}

// GetBackupId returns the BackupId field value if set, zero value otherwise.
func (o *V1beta1Backup) GetBackupId() string {
	if o == nil || IsNil(o.BackupId) {
		var ret string
		return ret
	}
	return *o.BackupId
}

// GetBackupIdOk returns a tuple with the BackupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1beta1Backup) GetBackupIdOk() (*string, bool) {
	if o == nil || IsNil(o.BackupId) {
		return nil, false
	}
	return o.BackupId, true
}

// HasBackupId returns a boolean if a field has been set.
func (o *V1beta1Backup) HasBackupId() bool {
	if o != nil && !IsNil(o.BackupId) {
		return true
	}

	return false
}

// SetBackupId gets a reference to the given string and assigns it to the BackupId field.
func (o *V1beta1Backup) SetBackupId(v string) {
	o.BackupId = &v
}

// GetClusterId returns the ClusterId field value
func (o *V1beta1Backup) GetClusterId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value
// and a boolean to check if the value has been set.
func (o *V1beta1Backup) GetClusterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusterId, true
}

// SetClusterId sets field value
func (o *V1beta1Backup) SetClusterId(v string) {
	o.ClusterId = v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *V1beta1Backup) GetCreateTime() time.Time {
	if o == nil || IsNil(o.CreateTime) {
		var ret time.Time
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1beta1Backup) GetCreateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *V1beta1Backup) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given time.Time and assigns it to the CreateTime field.
func (o *V1beta1Backup) SetCreateTime(v time.Time) {
	o.CreateTime = &v
}

func (o V1beta1Backup) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1beta1Backup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.BackupId) {
		toSerialize["backupId"] = o.BackupId
	}
	toSerialize["clusterId"] = o.ClusterId
	if !IsNil(o.CreateTime) {
		toSerialize["createTime"] = o.CreateTime
	}
	return toSerialize, nil
}

func (o *V1beta1Backup) UnmarshalJSON(data []byte) (err error) {
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

	varV1beta1Backup := _V1beta1Backup{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV1beta1Backup)

	if err != nil {
		return err
	}

	*o = V1beta1Backup(varV1beta1Backup)

	return err
}

type NullableV1beta1Backup struct {
	value *V1beta1Backup
	isSet bool
}

func (v NullableV1beta1Backup) Get() *V1beta1Backup {
	return v.value
}

func (v *NullableV1beta1Backup) Set(val *V1beta1Backup) {
	v.value = val
	v.isSet = true
}

func (v NullableV1beta1Backup) IsSet() bool {
	return v.isSet
}

func (v *NullableV1beta1Backup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1beta1Backup(val *V1beta1Backup) *NullableV1beta1Backup {
	return &NullableV1beta1Backup{value: val, isSet: true}
}

func (v NullableV1beta1Backup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1beta1Backup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
