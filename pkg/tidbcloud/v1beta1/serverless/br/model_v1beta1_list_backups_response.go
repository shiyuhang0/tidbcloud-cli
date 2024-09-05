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

// checks if the V1beta1ListBackupsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1beta1ListBackupsResponse{}

// V1beta1ListBackupsResponse struct for V1beta1ListBackupsResponse
type V1beta1ListBackupsResponse struct {
	// A list of clusters.
	Backups []V1beta1Backup `json:"backups,omitempty"`
	// Token provided to retrieve the next page of results.
	NextPageToken *string `json:"nextPageToken,omitempty"`
	// Total number of backups.
	TotalSize            *int64 `json:"totalSize,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _V1beta1ListBackupsResponse V1beta1ListBackupsResponse

// NewV1beta1ListBackupsResponse instantiates a new V1beta1ListBackupsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1beta1ListBackupsResponse() *V1beta1ListBackupsResponse {
	this := V1beta1ListBackupsResponse{}
	return &this
}

// NewV1beta1ListBackupsResponseWithDefaults instantiates a new V1beta1ListBackupsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1beta1ListBackupsResponseWithDefaults() *V1beta1ListBackupsResponse {
	this := V1beta1ListBackupsResponse{}
	return &this
}

// GetBackups returns the Backups field value if set, zero value otherwise.
func (o *V1beta1ListBackupsResponse) GetBackups() []V1beta1Backup {
	if o == nil || IsNil(o.Backups) {
		var ret []V1beta1Backup
		return ret
	}
	return o.Backups
}

// GetBackupsOk returns a tuple with the Backups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1beta1ListBackupsResponse) GetBackupsOk() ([]V1beta1Backup, bool) {
	if o == nil || IsNil(o.Backups) {
		return nil, false
	}
	return o.Backups, true
}

// HasBackups returns a boolean if a field has been set.
func (o *V1beta1ListBackupsResponse) HasBackups() bool {
	if o != nil && !IsNil(o.Backups) {
		return true
	}

	return false
}

// SetBackups gets a reference to the given []V1beta1Backup and assigns it to the Backups field.
func (o *V1beta1ListBackupsResponse) SetBackups(v []V1beta1Backup) {
	o.Backups = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *V1beta1ListBackupsResponse) GetNextPageToken() string {
	if o == nil || IsNil(o.NextPageToken) {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1beta1ListBackupsResponse) GetNextPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextPageToken) {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *V1beta1ListBackupsResponse) HasNextPageToken() bool {
	if o != nil && !IsNil(o.NextPageToken) {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *V1beta1ListBackupsResponse) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

// GetTotalSize returns the TotalSize field value if set, zero value otherwise.
func (o *V1beta1ListBackupsResponse) GetTotalSize() int64 {
	if o == nil || IsNil(o.TotalSize) {
		var ret int64
		return ret
	}
	return *o.TotalSize
}

// GetTotalSizeOk returns a tuple with the TotalSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1beta1ListBackupsResponse) GetTotalSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalSize) {
		return nil, false
	}
	return o.TotalSize, true
}

// HasTotalSize returns a boolean if a field has been set.
func (o *V1beta1ListBackupsResponse) HasTotalSize() bool {
	if o != nil && !IsNil(o.TotalSize) {
		return true
	}

	return false
}

// SetTotalSize gets a reference to the given int64 and assigns it to the TotalSize field.
func (o *V1beta1ListBackupsResponse) SetTotalSize(v int64) {
	o.TotalSize = &v
}

func (o V1beta1ListBackupsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1beta1ListBackupsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Backups) {
		toSerialize["backups"] = o.Backups
	}
	if !IsNil(o.NextPageToken) {
		toSerialize["nextPageToken"] = o.NextPageToken
	}
	if !IsNil(o.TotalSize) {
		toSerialize["totalSize"] = o.TotalSize
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *V1beta1ListBackupsResponse) UnmarshalJSON(data []byte) (err error) {
	varV1beta1ListBackupsResponse := _V1beta1ListBackupsResponse{}

	err = json.Unmarshal(data, &varV1beta1ListBackupsResponse)

	if err != nil {
		return err
	}

	*o = V1beta1ListBackupsResponse(varV1beta1ListBackupsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "backups")
		delete(additionalProperties, "nextPageToken")
		delete(additionalProperties, "totalSize")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableV1beta1ListBackupsResponse struct {
	value *V1beta1ListBackupsResponse
	isSet bool
}

func (v NullableV1beta1ListBackupsResponse) Get() *V1beta1ListBackupsResponse {
	return v.value
}

func (v *NullableV1beta1ListBackupsResponse) Set(val *V1beta1ListBackupsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableV1beta1ListBackupsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableV1beta1ListBackupsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1beta1ListBackupsResponse(val *V1beta1ListBackupsResponse) *NullableV1beta1ListBackupsResponse {
	return &NullableV1beta1ListBackupsResponse{value: val, isSet: true}
}

func (v NullableV1beta1ListBackupsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1beta1ListBackupsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
