/*
TiDB Cloud Serverless Export Open API

TiDB Cloud Serverless Export Open API

API version: v1beta1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package export

import (
	"encoding/json"
	"fmt"
)

// checks if the AzureBlobTarget type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AzureBlobTarget{}

// AzureBlobTarget struct for AzureBlobTarget
type AzureBlobTarget struct {
	// The Azure Blob URI of the export target.
	AuthType ExportAzureBlobAuthTypeEnum `json:"authType"`
	// The sas token. This field is input-only.
	SasToken *string `json:"sasToken,omitempty"`
	// The Azure Blob URI of the export target. For example: azure://<account>.blob.core.windows.net/<container>/<path>.
	Uri                  string `json:"uri"`
	AdditionalProperties map[string]interface{}
}

type _AzureBlobTarget AzureBlobTarget

// NewAzureBlobTarget instantiates a new AzureBlobTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureBlobTarget(authType ExportAzureBlobAuthTypeEnum, uri string) *AzureBlobTarget {
	this := AzureBlobTarget{}
	this.AuthType = authType
	this.Uri = uri
	return &this
}

// NewAzureBlobTargetWithDefaults instantiates a new AzureBlobTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureBlobTargetWithDefaults() *AzureBlobTarget {
	this := AzureBlobTarget{}
	return &this
}

// GetAuthType returns the AuthType field value
func (o *AzureBlobTarget) GetAuthType() ExportAzureBlobAuthTypeEnum {
	if o == nil {
		var ret ExportAzureBlobAuthTypeEnum
		return ret
	}

	return o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value
// and a boolean to check if the value has been set.
func (o *AzureBlobTarget) GetAuthTypeOk() (*ExportAzureBlobAuthTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthType, true
}

// SetAuthType sets field value
func (o *AzureBlobTarget) SetAuthType(v ExportAzureBlobAuthTypeEnum) {
	o.AuthType = v
}

// GetSasToken returns the SasToken field value if set, zero value otherwise.
func (o *AzureBlobTarget) GetSasToken() string {
	if o == nil || IsNil(o.SasToken) {
		var ret string
		return ret
	}
	return *o.SasToken
}

// GetSasTokenOk returns a tuple with the SasToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureBlobTarget) GetSasTokenOk() (*string, bool) {
	if o == nil || IsNil(o.SasToken) {
		return nil, false
	}
	return o.SasToken, true
}

// HasSasToken returns a boolean if a field has been set.
func (o *AzureBlobTarget) HasSasToken() bool {
	if o != nil && !IsNil(o.SasToken) {
		return true
	}

	return false
}

// SetSasToken gets a reference to the given string and assigns it to the SasToken field.
func (o *AzureBlobTarget) SetSasToken(v string) {
	o.SasToken = &v
}

// GetUri returns the Uri field value
func (o *AzureBlobTarget) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *AzureBlobTarget) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *AzureBlobTarget) SetUri(v string) {
	o.Uri = v
}

func (o AzureBlobTarget) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AzureBlobTarget) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["authType"] = o.AuthType
	if !IsNil(o.SasToken) {
		toSerialize["sasToken"] = o.SasToken
	}
	toSerialize["uri"] = o.Uri

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AzureBlobTarget) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"authType",
		"uri",
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

	varAzureBlobTarget := _AzureBlobTarget{}

	err = json.Unmarshal(data, &varAzureBlobTarget)

	if err != nil {
		return err
	}

	*o = AzureBlobTarget(varAzureBlobTarget)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "authType")
		delete(additionalProperties, "sasToken")
		delete(additionalProperties, "uri")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAzureBlobTarget struct {
	value *AzureBlobTarget
	isSet bool
}

func (v NullableAzureBlobTarget) Get() *AzureBlobTarget {
	return v.value
}

func (v *NullableAzureBlobTarget) Set(val *AzureBlobTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureBlobTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureBlobTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureBlobTarget(val *AzureBlobTarget) *NullableAzureBlobTarget {
	return &NullableAzureBlobTarget{value: val, isSet: true}
}

func (v NullableAzureBlobTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureBlobTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
