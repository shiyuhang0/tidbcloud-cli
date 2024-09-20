/*
TiDB Cloud Serverless Open API

TiDB Cloud Serverless Open API

API version: v1beta1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package imp

import (
	"encoding/json"
)

// checks if the CreationDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreationDetails{}

// CreationDetails struct for CreationDetails
type CreationDetails struct {
	// Optional. The options of the import.
	ImportOptions *ImportOptions `json:"importOptions,omitempty"`
	// Optional. The source of the import.
	Source               *ImportSource `json:"source,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreationDetails CreationDetails

// NewCreationDetails instantiates a new CreationDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreationDetails() *CreationDetails {
	this := CreationDetails{}
	return &this
}

// NewCreationDetailsWithDefaults instantiates a new CreationDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreationDetailsWithDefaults() *CreationDetails {
	this := CreationDetails{}
	return &this
}

// GetImportOptions returns the ImportOptions field value if set, zero value otherwise.
func (o *CreationDetails) GetImportOptions() ImportOptions {
	if o == nil || IsNil(o.ImportOptions) {
		var ret ImportOptions
		return ret
	}
	return *o.ImportOptions
}

// GetImportOptionsOk returns a tuple with the ImportOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreationDetails) GetImportOptionsOk() (*ImportOptions, bool) {
	if o == nil || IsNil(o.ImportOptions) {
		return nil, false
	}
	return o.ImportOptions, true
}

// HasImportOptions returns a boolean if a field has been set.
func (o *CreationDetails) HasImportOptions() bool {
	if o != nil && !IsNil(o.ImportOptions) {
		return true
	}

	return false
}

// SetImportOptions gets a reference to the given ImportOptions and assigns it to the ImportOptions field.
func (o *CreationDetails) SetImportOptions(v ImportOptions) {
	o.ImportOptions = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *CreationDetails) GetSource() ImportSource {
	if o == nil || IsNil(o.Source) {
		var ret ImportSource
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreationDetails) GetSourceOk() (*ImportSource, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *CreationDetails) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given ImportSource and assigns it to the Source field.
func (o *CreationDetails) SetSource(v ImportSource) {
	o.Source = &v
}

func (o CreationDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreationDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ImportOptions) {
		toSerialize["importOptions"] = o.ImportOptions
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreationDetails) UnmarshalJSON(data []byte) (err error) {
	varCreationDetails := _CreationDetails{}

	err = json.Unmarshal(data, &varCreationDetails)

	if err != nil {
		return err
	}

	*o = CreationDetails(varCreationDetails)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "importOptions")
		delete(additionalProperties, "source")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreationDetails struct {
	value *CreationDetails
	isSet bool
}

func (v NullableCreationDetails) Get() *CreationDetails {
	return v.value
}

func (v *NullableCreationDetails) Set(val *CreationDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableCreationDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableCreationDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreationDetails(val *CreationDetails) *NullableCreationDetails {
	return &NullableCreationDetails{value: val, isSet: true}
}

func (v NullableCreationDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreationDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}