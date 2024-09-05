/*
TiDB Cloud Serverless Export Open API

TiDB Cloud Serverless Export Open API

API version: v1beta1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package export

import (
	"encoding/json"
)

// checks if the ListExportFilesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListExportFilesResponse{}

// ListExportFilesResponse struct for ListExportFilesResponse
type ListExportFilesResponse struct {
	// The files of the export.
	Files []ExportFile `json:"files,omitempty"`
	// Token provided to retrieve the next page of results.
	NextPageToken *string `json:"nextPageToken,omitempty"`
	// Total number of export files.
	TotalSize            *int64 `json:"totalSize,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListExportFilesResponse ListExportFilesResponse

// NewListExportFilesResponse instantiates a new ListExportFilesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListExportFilesResponse() *ListExportFilesResponse {
	this := ListExportFilesResponse{}
	return &this
}

// NewListExportFilesResponseWithDefaults instantiates a new ListExportFilesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListExportFilesResponseWithDefaults() *ListExportFilesResponse {
	this := ListExportFilesResponse{}
	return &this
}

// GetFiles returns the Files field value if set, zero value otherwise.
func (o *ListExportFilesResponse) GetFiles() []ExportFile {
	if o == nil || IsNil(o.Files) {
		var ret []ExportFile
		return ret
	}
	return o.Files
}

// GetFilesOk returns a tuple with the Files field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListExportFilesResponse) GetFilesOk() ([]ExportFile, bool) {
	if o == nil || IsNil(o.Files) {
		return nil, false
	}
	return o.Files, true
}

// HasFiles returns a boolean if a field has been set.
func (o *ListExportFilesResponse) HasFiles() bool {
	if o != nil && !IsNil(o.Files) {
		return true
	}

	return false
}

// SetFiles gets a reference to the given []ExportFile and assigns it to the Files field.
func (o *ListExportFilesResponse) SetFiles(v []ExportFile) {
	o.Files = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *ListExportFilesResponse) GetNextPageToken() string {
	if o == nil || IsNil(o.NextPageToken) {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListExportFilesResponse) GetNextPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextPageToken) {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *ListExportFilesResponse) HasNextPageToken() bool {
	if o != nil && !IsNil(o.NextPageToken) {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *ListExportFilesResponse) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

// GetTotalSize returns the TotalSize field value if set, zero value otherwise.
func (o *ListExportFilesResponse) GetTotalSize() int64 {
	if o == nil || IsNil(o.TotalSize) {
		var ret int64
		return ret
	}
	return *o.TotalSize
}

// GetTotalSizeOk returns a tuple with the TotalSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListExportFilesResponse) GetTotalSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalSize) {
		return nil, false
	}
	return o.TotalSize, true
}

// HasTotalSize returns a boolean if a field has been set.
func (o *ListExportFilesResponse) HasTotalSize() bool {
	if o != nil && !IsNil(o.TotalSize) {
		return true
	}

	return false
}

// SetTotalSize gets a reference to the given int64 and assigns it to the TotalSize field.
func (o *ListExportFilesResponse) SetTotalSize(v int64) {
	o.TotalSize = &v
}

func (o ListExportFilesResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListExportFilesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Files) {
		toSerialize["files"] = o.Files
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

func (o *ListExportFilesResponse) UnmarshalJSON(data []byte) (err error) {
	varListExportFilesResponse := _ListExportFilesResponse{}

	err = json.Unmarshal(data, &varListExportFilesResponse)

	if err != nil {
		return err
	}

	*o = ListExportFilesResponse(varListExportFilesResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "files")
		delete(additionalProperties, "nextPageToken")
		delete(additionalProperties, "totalSize")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListExportFilesResponse struct {
	value *ListExportFilesResponse
	isSet bool
}

func (v NullableListExportFilesResponse) Get() *ListExportFilesResponse {
	return v.value
}

func (v *NullableListExportFilesResponse) Set(val *ListExportFilesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListExportFilesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListExportFilesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListExportFilesResponse(val *ListExportFilesResponse) *NullableListExportFilesResponse {
	return &NullableListExportFilesResponse{value: val, isSet: true}
}

func (v NullableListExportFilesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListExportFilesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
