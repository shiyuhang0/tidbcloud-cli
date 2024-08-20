/*
TiDB Cloud Serverless Export Open API

TiDB Cloud Serverless Export Open API

API version: v1beta1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the DownloadUrl type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DownloadUrl{}

// DownloadUrl struct for DownloadUrl
type DownloadUrl struct {
	// The name of the download file.
	Name *string `json:"name,omitempty"`
	// The download url.
	Url *string `json:"url,omitempty"`
	// The size in bytes of the the download file.
	Size *int64 `json:"size,omitempty"`
}

// NewDownloadUrl instantiates a new DownloadUrl object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDownloadUrl() *DownloadUrl {
	this := DownloadUrl{}
	return &this
}

// NewDownloadUrlWithDefaults instantiates a new DownloadUrl object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDownloadUrlWithDefaults() *DownloadUrl {
	this := DownloadUrl{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DownloadUrl) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DownloadUrl) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DownloadUrl) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DownloadUrl) SetName(v string) {
	o.Name = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *DownloadUrl) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DownloadUrl) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *DownloadUrl) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *DownloadUrl) SetUrl(v string) {
	o.Url = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *DownloadUrl) GetSize() int64 {
	if o == nil || IsNil(o.Size) {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DownloadUrl) GetSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *DownloadUrl) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *DownloadUrl) SetSize(v int64) {
	o.Size = &v
}

func (o DownloadUrl) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DownloadUrl) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	return toSerialize, nil
}

type NullableDownloadUrl struct {
	value *DownloadUrl
	isSet bool
}

func (v NullableDownloadUrl) Get() *DownloadUrl {
	return v.value
}

func (v *NullableDownloadUrl) Set(val *DownloadUrl) {
	v.value = val
	v.isSet = true
}

func (v NullableDownloadUrl) IsSet() bool {
	return v.isSet
}

func (v *NullableDownloadUrl) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDownloadUrl(val *DownloadUrl) *NullableDownloadUrl {
	return &NullableDownloadUrl{value: val, isSet: true}
}

func (v NullableDownloadUrl) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDownloadUrl) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

