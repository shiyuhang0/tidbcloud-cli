/*
TiDB Cloud Serverless Export Open API

TiDB Cloud Serverless Export Open API

API version: v1beta1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package export

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// checks if the Export type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Export{}

// Export Message for export resource.
type Export struct {
	// Output_only. The unique ID of the export.
	ExportId *string `json:"exportId,omitempty"`
	// Output_only. The unique name of the export.
	Name *string `json:"name,omitempty"`
	// Required. The cluster ID that export belong to.
	ClusterId string `json:"clusterId"`
	// Output_only. The creator of the export.
	CreatedBy *string `json:"createdBy,omitempty"`
	// Output_only. The state of the export.
	State *ExportStateEnum `json:"state,omitempty"`
	// Optional. The options of the export.
	ExportOptions *ExportOptions `json:"exportOptions,omitempty"`
	// Optional. The target of the export.
	Target *ExportTarget `json:"target,omitempty"`
	// Optional. The failed reason of the export.
	Reason NullableString `json:"reason,omitempty"`
	// Optional. The display name of the export. Default: SNAPSHOT_{snapshot_time}.
	DisplayName *string `json:"displayName,omitempty"`
	// Output_only. Timestamp when the export was created.
	CreateTime *time.Time `json:"createTime,omitempty"`
	// Output_only. Timestamp when the export was updated.
	UpdateTime NullableTime `json:"updateTime,omitempty"`
	// Output_only. Timestamp when the export was completed.
	CompleteTime NullableTime `json:"completeTime,omitempty"`
	// Output_only. Snapshot time of the export.
	SnapshotTime NullableTime `json:"snapshotTime,omitempty"`
	// Output_only. Expire time of the export.
	ExpireTime NullableTime `json:"expireTime,omitempty"`
}

type _Export Export

// NewExport instantiates a new Export object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExport(clusterId string) *Export {
	this := Export{}
	this.ClusterId = clusterId
	return &this
}

// NewExportWithDefaults instantiates a new Export object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExportWithDefaults() *Export {
	this := Export{}
	return &this
}

// GetExportId returns the ExportId field value if set, zero value otherwise.
func (o *Export) GetExportId() string {
	if o == nil || IsNil(o.ExportId) {
		var ret string
		return ret
	}
	return *o.ExportId
}

// GetExportIdOk returns a tuple with the ExportId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Export) GetExportIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExportId) {
		return nil, false
	}
	return o.ExportId, true
}

// HasExportId returns a boolean if a field has been set.
func (o *Export) HasExportId() bool {
	if o != nil && !IsNil(o.ExportId) {
		return true
	}

	return false
}

// SetExportId gets a reference to the given string and assigns it to the ExportId field.
func (o *Export) SetExportId(v string) {
	o.ExportId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Export) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Export) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Export) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Export) SetName(v string) {
	o.Name = &v
}

// GetClusterId returns the ClusterId field value
func (o *Export) GetClusterId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value
// and a boolean to check if the value has been set.
func (o *Export) GetClusterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusterId, true
}

// SetClusterId sets field value
func (o *Export) SetClusterId(v string) {
	o.ClusterId = v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *Export) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy) {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Export) GetCreatedByOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *Export) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *Export) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *Export) GetState() ExportStateEnum {
	if o == nil || IsNil(o.State) {
		var ret ExportStateEnum
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Export) GetStateOk() (*ExportStateEnum, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *Export) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given ExportStateEnum and assigns it to the State field.
func (o *Export) SetState(v ExportStateEnum) {
	o.State = &v
}

// GetExportOptions returns the ExportOptions field value if set, zero value otherwise.
func (o *Export) GetExportOptions() ExportOptions {
	if o == nil || IsNil(o.ExportOptions) {
		var ret ExportOptions
		return ret
	}
	return *o.ExportOptions
}

// GetExportOptionsOk returns a tuple with the ExportOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Export) GetExportOptionsOk() (*ExportOptions, bool) {
	if o == nil || IsNil(o.ExportOptions) {
		return nil, false
	}
	return o.ExportOptions, true
}

// HasExportOptions returns a boolean if a field has been set.
func (o *Export) HasExportOptions() bool {
	if o != nil && !IsNil(o.ExportOptions) {
		return true
	}

	return false
}

// SetExportOptions gets a reference to the given ExportOptions and assigns it to the ExportOptions field.
func (o *Export) SetExportOptions(v ExportOptions) {
	o.ExportOptions = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *Export) GetTarget() ExportTarget {
	if o == nil || IsNil(o.Target) {
		var ret ExportTarget
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Export) GetTargetOk() (*ExportTarget, bool) {
	if o == nil || IsNil(o.Target) {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *Export) HasTarget() bool {
	if o != nil && !IsNil(o.Target) {
		return true
	}

	return false
}

// SetTarget gets a reference to the given ExportTarget and assigns it to the Target field.
func (o *Export) SetTarget(v ExportTarget) {
	o.Target = &v
}

// GetReason returns the Reason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Export) GetReason() string {
	if o == nil || IsNil(o.Reason.Get()) {
		var ret string
		return ret
	}
	return *o.Reason.Get()
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Export) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Reason.Get(), o.Reason.IsSet()
}

// HasReason returns a boolean if a field has been set.
func (o *Export) HasReason() bool {
	if o != nil && o.Reason.IsSet() {
		return true
	}

	return false
}

// SetReason gets a reference to the given NullableString and assigns it to the Reason field.
func (o *Export) SetReason(v string) {
	o.Reason.Set(&v)
}

// SetReasonNil sets the value for Reason to be an explicit nil
func (o *Export) SetReasonNil() {
	o.Reason.Set(nil)
}

// UnsetReason ensures that no value is present for Reason, not even an explicit nil
func (o *Export) UnsetReason() {
	o.Reason.Unset()
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *Export) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Export) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *Export) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *Export) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *Export) GetCreateTime() time.Time {
	if o == nil || IsNil(o.CreateTime) {
		var ret time.Time
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Export) GetCreateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *Export) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given time.Time and assigns it to the CreateTime field.
func (o *Export) SetCreateTime(v time.Time) {
	o.CreateTime = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Export) GetUpdateTime() time.Time {
	if o == nil || IsNil(o.UpdateTime.Get()) {
		var ret time.Time
		return ret
	}
	return *o.UpdateTime.Get()
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Export) GetUpdateTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdateTime.Get(), o.UpdateTime.IsSet()
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *Export) HasUpdateTime() bool {
	if o != nil && o.UpdateTime.IsSet() {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given NullableTime and assigns it to the UpdateTime field.
func (o *Export) SetUpdateTime(v time.Time) {
	o.UpdateTime.Set(&v)
}

// SetUpdateTimeNil sets the value for UpdateTime to be an explicit nil
func (o *Export) SetUpdateTimeNil() {
	o.UpdateTime.Set(nil)
}

// UnsetUpdateTime ensures that no value is present for UpdateTime, not even an explicit nil
func (o *Export) UnsetUpdateTime() {
	o.UpdateTime.Unset()
}

// GetCompleteTime returns the CompleteTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Export) GetCompleteTime() time.Time {
	if o == nil || IsNil(o.CompleteTime.Get()) {
		var ret time.Time
		return ret
	}
	return *o.CompleteTime.Get()
}

// GetCompleteTimeOk returns a tuple with the CompleteTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Export) GetCompleteTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CompleteTime.Get(), o.CompleteTime.IsSet()
}

// HasCompleteTime returns a boolean if a field has been set.
func (o *Export) HasCompleteTime() bool {
	if o != nil && o.CompleteTime.IsSet() {
		return true
	}

	return false
}

// SetCompleteTime gets a reference to the given NullableTime and assigns it to the CompleteTime field.
func (o *Export) SetCompleteTime(v time.Time) {
	o.CompleteTime.Set(&v)
}

// SetCompleteTimeNil sets the value for CompleteTime to be an explicit nil
func (o *Export) SetCompleteTimeNil() {
	o.CompleteTime.Set(nil)
}

// UnsetCompleteTime ensures that no value is present for CompleteTime, not even an explicit nil
func (o *Export) UnsetCompleteTime() {
	o.CompleteTime.Unset()
}

// GetSnapshotTime returns the SnapshotTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Export) GetSnapshotTime() time.Time {
	if o == nil || IsNil(o.SnapshotTime.Get()) {
		var ret time.Time
		return ret
	}
	return *o.SnapshotTime.Get()
}

// GetSnapshotTimeOk returns a tuple with the SnapshotTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Export) GetSnapshotTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.SnapshotTime.Get(), o.SnapshotTime.IsSet()
}

// HasSnapshotTime returns a boolean if a field has been set.
func (o *Export) HasSnapshotTime() bool {
	if o != nil && o.SnapshotTime.IsSet() {
		return true
	}

	return false
}

// SetSnapshotTime gets a reference to the given NullableTime and assigns it to the SnapshotTime field.
func (o *Export) SetSnapshotTime(v time.Time) {
	o.SnapshotTime.Set(&v)
}

// SetSnapshotTimeNil sets the value for SnapshotTime to be an explicit nil
func (o *Export) SetSnapshotTimeNil() {
	o.SnapshotTime.Set(nil)
}

// UnsetSnapshotTime ensures that no value is present for SnapshotTime, not even an explicit nil
func (o *Export) UnsetSnapshotTime() {
	o.SnapshotTime.Unset()
}

// GetExpireTime returns the ExpireTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Export) GetExpireTime() time.Time {
	if o == nil || IsNil(o.ExpireTime.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ExpireTime.Get()
}

// GetExpireTimeOk returns a tuple with the ExpireTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Export) GetExpireTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpireTime.Get(), o.ExpireTime.IsSet()
}

// HasExpireTime returns a boolean if a field has been set.
func (o *Export) HasExpireTime() bool {
	if o != nil && o.ExpireTime.IsSet() {
		return true
	}

	return false
}

// SetExpireTime gets a reference to the given NullableTime and assigns it to the ExpireTime field.
func (o *Export) SetExpireTime(v time.Time) {
	o.ExpireTime.Set(&v)
}

// SetExpireTimeNil sets the value for ExpireTime to be an explicit nil
func (o *Export) SetExpireTimeNil() {
	o.ExpireTime.Set(nil)
}

// UnsetExpireTime ensures that no value is present for ExpireTime, not even an explicit nil
func (o *Export) UnsetExpireTime() {
	o.ExpireTime.Unset()
}

func (o Export) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Export) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExportId) {
		toSerialize["exportId"] = o.ExportId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["clusterId"] = o.ClusterId
	if !IsNil(o.CreatedBy) {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.ExportOptions) {
		toSerialize["exportOptions"] = o.ExportOptions
	}
	if !IsNil(o.Target) {
		toSerialize["target"] = o.Target
	}
	if o.Reason.IsSet() {
		toSerialize["reason"] = o.Reason.Get()
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.CreateTime) {
		toSerialize["createTime"] = o.CreateTime
	}
	if o.UpdateTime.IsSet() {
		toSerialize["updateTime"] = o.UpdateTime.Get()
	}
	if o.CompleteTime.IsSet() {
		toSerialize["completeTime"] = o.CompleteTime.Get()
	}
	if o.SnapshotTime.IsSet() {
		toSerialize["snapshotTime"] = o.SnapshotTime.Get()
	}
	if o.ExpireTime.IsSet() {
		toSerialize["expireTime"] = o.ExpireTime.Get()
	}
	return toSerialize, nil
}

func (o *Export) UnmarshalJSON(data []byte) (err error) {
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

	varExport := _Export{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varExport)

	if err != nil {
		return err
	}

	*o = Export(varExport)

	return err
}

type NullableExport struct {
	value *Export
	isSet bool
}

func (v NullableExport) Get() *Export {
	return v.value
}

func (v *NullableExport) Set(val *Export) {
	v.value = val
	v.isSet = true
}

func (v NullableExport) IsSet() bool {
	return v.isSet
}

func (v *NullableExport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExport(val *Export) *NullableExport {
	return &NullableExport{value: val, isSet: true}
}

func (v NullableExport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
