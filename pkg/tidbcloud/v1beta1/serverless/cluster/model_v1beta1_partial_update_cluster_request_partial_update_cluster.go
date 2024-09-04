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

// checks if the V1beta1PartialUpdateClusterRequestPartialUpdateCluster type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1beta1PartialUpdateClusterRequestPartialUpdateCluster{}

// V1beta1PartialUpdateClusterRequestPartialUpdateCluster struct for V1beta1PartialUpdateClusterRequestPartialUpdateCluster
type V1beta1PartialUpdateClusterRequestPartialUpdateCluster struct {
	// Output_only. The unique ID of the cluster.
	ClusterId string `json:"clusterId"`
	// Required. User friendly display name of the cluster.
	DisplayName *string `json:"displayName,omitempty"`
	// Optional. The spending limit for the cluster.
	SpendingLimit *ClusterSpendingLimit `json:"spendingLimit,omitempty"`
	// Optional. Automated backup policy to set on the cluster.
	AutomatedBackupPolicy *V1beta1ClusterAutomatedBackupPolicy `json:"automatedBackupPolicy,omitempty"`
	// Optional. The endpoints for connecting to the cluster.
	Endpoints *V1beta1ClusterEndpoints `json:"endpoints,omitempty"`
	// Optional. The labels for the cluster. tidb.cloud/organization. The label for the cluster organization id. tidb.cloud/project. The label for the cluster project id.
	Labels *map[string]string `json:"labels,omitempty"`
}

type _V1beta1PartialUpdateClusterRequestPartialUpdateCluster V1beta1PartialUpdateClusterRequestPartialUpdateCluster

// NewV1beta1PartialUpdateClusterRequestPartialUpdateCluster instantiates a new V1beta1PartialUpdateClusterRequestPartialUpdateCluster object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1beta1PartialUpdateClusterRequestPartialUpdateCluster(clusterId string) *V1beta1PartialUpdateClusterRequestPartialUpdateCluster {
	this := V1beta1PartialUpdateClusterRequestPartialUpdateCluster{}
	this.ClusterId = clusterId
	return &this
}

// NewV1beta1PartialUpdateClusterRequestPartialUpdateClusterWithDefaults instantiates a new V1beta1PartialUpdateClusterRequestPartialUpdateCluster object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1beta1PartialUpdateClusterRequestPartialUpdateClusterWithDefaults() *V1beta1PartialUpdateClusterRequestPartialUpdateCluster {
	this := V1beta1PartialUpdateClusterRequestPartialUpdateCluster{}
	return &this
}

// GetClusterId returns the ClusterId field value
func (o *V1beta1PartialUpdateClusterRequestPartialUpdateCluster) GetClusterId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value
// and a boolean to check if the value has been set.
func (o *V1beta1PartialUpdateClusterRequestPartialUpdateCluster) GetClusterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusterId, true
}

// SetClusterId sets field value
func (o *V1beta1PartialUpdateClusterRequestPartialUpdateCluster) SetClusterId(v string) {
	o.ClusterId = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *V1beta1PartialUpdateClusterRequestPartialUpdateCluster) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1beta1PartialUpdateClusterRequestPartialUpdateCluster) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *V1beta1PartialUpdateClusterRequestPartialUpdateCluster) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *V1beta1PartialUpdateClusterRequestPartialUpdateCluster) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetSpendingLimit returns the SpendingLimit field value if set, zero value otherwise.
func (o *V1beta1PartialUpdateClusterRequestPartialUpdateCluster) GetSpendingLimit() ClusterSpendingLimit {
	if o == nil || IsNil(o.SpendingLimit) {
		var ret ClusterSpendingLimit
		return ret
	}
	return *o.SpendingLimit
}

// GetSpendingLimitOk returns a tuple with the SpendingLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1beta1PartialUpdateClusterRequestPartialUpdateCluster) GetSpendingLimitOk() (*ClusterSpendingLimit, bool) {
	if o == nil || IsNil(o.SpendingLimit) {
		return nil, false
	}
	return o.SpendingLimit, true
}

// HasSpendingLimit returns a boolean if a field has been set.
func (o *V1beta1PartialUpdateClusterRequestPartialUpdateCluster) HasSpendingLimit() bool {
	if o != nil && !IsNil(o.SpendingLimit) {
		return true
	}

	return false
}

// SetSpendingLimit gets a reference to the given ClusterSpendingLimit and assigns it to the SpendingLimit field.
func (o *V1beta1PartialUpdateClusterRequestPartialUpdateCluster) SetSpendingLimit(v ClusterSpendingLimit) {
	o.SpendingLimit = &v
}

// GetAutomatedBackupPolicy returns the AutomatedBackupPolicy field value if set, zero value otherwise.
func (o *V1beta1PartialUpdateClusterRequestPartialUpdateCluster) GetAutomatedBackupPolicy() V1beta1ClusterAutomatedBackupPolicy {
	if o == nil || IsNil(o.AutomatedBackupPolicy) {
		var ret V1beta1ClusterAutomatedBackupPolicy
		return ret
	}
	return *o.AutomatedBackupPolicy
}

// GetAutomatedBackupPolicyOk returns a tuple with the AutomatedBackupPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1beta1PartialUpdateClusterRequestPartialUpdateCluster) GetAutomatedBackupPolicyOk() (*V1beta1ClusterAutomatedBackupPolicy, bool) {
	if o == nil || IsNil(o.AutomatedBackupPolicy) {
		return nil, false
	}
	return o.AutomatedBackupPolicy, true
}

// HasAutomatedBackupPolicy returns a boolean if a field has been set.
func (o *V1beta1PartialUpdateClusterRequestPartialUpdateCluster) HasAutomatedBackupPolicy() bool {
	if o != nil && !IsNil(o.AutomatedBackupPolicy) {
		return true
	}

	return false
}

// SetAutomatedBackupPolicy gets a reference to the given V1beta1ClusterAutomatedBackupPolicy and assigns it to the AutomatedBackupPolicy field.
func (o *V1beta1PartialUpdateClusterRequestPartialUpdateCluster) SetAutomatedBackupPolicy(v V1beta1ClusterAutomatedBackupPolicy) {
	o.AutomatedBackupPolicy = &v
}

// GetEndpoints returns the Endpoints field value if set, zero value otherwise.
func (o *V1beta1PartialUpdateClusterRequestPartialUpdateCluster) GetEndpoints() V1beta1ClusterEndpoints {
	if o == nil || IsNil(o.Endpoints) {
		var ret V1beta1ClusterEndpoints
		return ret
	}
	return *o.Endpoints
}

// GetEndpointsOk returns a tuple with the Endpoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1beta1PartialUpdateClusterRequestPartialUpdateCluster) GetEndpointsOk() (*V1beta1ClusterEndpoints, bool) {
	if o == nil || IsNil(o.Endpoints) {
		return nil, false
	}
	return o.Endpoints, true
}

// HasEndpoints returns a boolean if a field has been set.
func (o *V1beta1PartialUpdateClusterRequestPartialUpdateCluster) HasEndpoints() bool {
	if o != nil && !IsNil(o.Endpoints) {
		return true
	}

	return false
}

// SetEndpoints gets a reference to the given V1beta1ClusterEndpoints and assigns it to the Endpoints field.
func (o *V1beta1PartialUpdateClusterRequestPartialUpdateCluster) SetEndpoints(v V1beta1ClusterEndpoints) {
	o.Endpoints = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *V1beta1PartialUpdateClusterRequestPartialUpdateCluster) GetLabels() map[string]string {
	if o == nil || IsNil(o.Labels) {
		var ret map[string]string
		return ret
	}
	return *o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1beta1PartialUpdateClusterRequestPartialUpdateCluster) GetLabelsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *V1beta1PartialUpdateClusterRequestPartialUpdateCluster) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given map[string]string and assigns it to the Labels field.
func (o *V1beta1PartialUpdateClusterRequestPartialUpdateCluster) SetLabels(v map[string]string) {
	o.Labels = &v
}

func (o V1beta1PartialUpdateClusterRequestPartialUpdateCluster) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1beta1PartialUpdateClusterRequestPartialUpdateCluster) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["clusterId"] = o.ClusterId
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.SpendingLimit) {
		toSerialize["spendingLimit"] = o.SpendingLimit
	}
	if !IsNil(o.AutomatedBackupPolicy) {
		toSerialize["automatedBackupPolicy"] = o.AutomatedBackupPolicy
	}
	if !IsNil(o.Endpoints) {
		toSerialize["endpoints"] = o.Endpoints
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	return toSerialize, nil
}

func (o *V1beta1PartialUpdateClusterRequestPartialUpdateCluster) UnmarshalJSON(data []byte) (err error) {
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

	varV1beta1PartialUpdateClusterRequestPartialUpdateCluster := _V1beta1PartialUpdateClusterRequestPartialUpdateCluster{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV1beta1PartialUpdateClusterRequestPartialUpdateCluster)

	if err != nil {
		return err
	}

	*o = V1beta1PartialUpdateClusterRequestPartialUpdateCluster(varV1beta1PartialUpdateClusterRequestPartialUpdateCluster)

	return err
}

type NullableV1beta1PartialUpdateClusterRequestPartialUpdateCluster struct {
	value *V1beta1PartialUpdateClusterRequestPartialUpdateCluster
	isSet bool
}

func (v NullableV1beta1PartialUpdateClusterRequestPartialUpdateCluster) Get() *V1beta1PartialUpdateClusterRequestPartialUpdateCluster {
	return v.value
}

func (v *NullableV1beta1PartialUpdateClusterRequestPartialUpdateCluster) Set(val *V1beta1PartialUpdateClusterRequestPartialUpdateCluster) {
	v.value = val
	v.isSet = true
}

func (v NullableV1beta1PartialUpdateClusterRequestPartialUpdateCluster) IsSet() bool {
	return v.isSet
}

func (v *NullableV1beta1PartialUpdateClusterRequestPartialUpdateCluster) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1beta1PartialUpdateClusterRequestPartialUpdateCluster(val *V1beta1PartialUpdateClusterRequestPartialUpdateCluster) *NullableV1beta1PartialUpdateClusterRequestPartialUpdateCluster {
	return &NullableV1beta1PartialUpdateClusterRequestPartialUpdateCluster{value: val, isSet: true}
}

func (v NullableV1beta1PartialUpdateClusterRequestPartialUpdateCluster) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1beta1PartialUpdateClusterRequestPartialUpdateCluster) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
