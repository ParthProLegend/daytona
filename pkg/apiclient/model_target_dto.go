/*
Daytona Server API

Daytona Server API

API version: v0.0.0-dev
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apiclient

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the TargetDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TargetDTO{}

// TargetDTO struct for TargetDTO
type TargetDTO struct {
	Default          bool              `json:"default"`
	EnvVars          map[string]string `json:"envVars"`
	Id               string            `json:"id"`
	LastJob          *Job              `json:"lastJob,omitempty"`
	LastJobId        *string           `json:"lastJobId,omitempty"`
	Metadata         *TargetMetadata   `json:"metadata,omitempty"`
	Name             string            `json:"name"`
	ProviderMetadata *string           `json:"providerMetadata,omitempty"`
	State            ResourceState     `json:"state"`
	TargetConfig     TargetConfig      `json:"targetConfig"`
	TargetConfigId   string            `json:"targetConfigId"`
	Workspaces       []Workspace       `json:"workspaces"`
}

type _TargetDTO TargetDTO

// NewTargetDTO instantiates a new TargetDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTargetDTO(default_ bool, envVars map[string]string, id string, name string, state ResourceState, targetConfig TargetConfig, targetConfigId string, workspaces []Workspace) *TargetDTO {
	this := TargetDTO{}
	this.Default = default_
	this.EnvVars = envVars
	this.Id = id
	this.Name = name
	this.State = state
	this.TargetConfig = targetConfig
	this.TargetConfigId = targetConfigId
	this.Workspaces = workspaces
	return &this
}

// NewTargetDTOWithDefaults instantiates a new TargetDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTargetDTOWithDefaults() *TargetDTO {
	this := TargetDTO{}
	return &this
}

// GetDefault returns the Default field value
func (o *TargetDTO) GetDefault() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Default
}

// GetDefaultOk returns a tuple with the Default field value
// and a boolean to check if the value has been set.
func (o *TargetDTO) GetDefaultOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Default, true
}

// SetDefault sets field value
func (o *TargetDTO) SetDefault(v bool) {
	o.Default = v
}

// GetEnvVars returns the EnvVars field value
func (o *TargetDTO) GetEnvVars() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.EnvVars
}

// GetEnvVarsOk returns a tuple with the EnvVars field value
// and a boolean to check if the value has been set.
func (o *TargetDTO) GetEnvVarsOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvVars, true
}

// SetEnvVars sets field value
func (o *TargetDTO) SetEnvVars(v map[string]string) {
	o.EnvVars = v
}

// GetId returns the Id field value
func (o *TargetDTO) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TargetDTO) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TargetDTO) SetId(v string) {
	o.Id = v
}

// GetLastJob returns the LastJob field value if set, zero value otherwise.
func (o *TargetDTO) GetLastJob() Job {
	if o == nil || IsNil(o.LastJob) {
		var ret Job
		return ret
	}
	return *o.LastJob
}

// GetLastJobOk returns a tuple with the LastJob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetDTO) GetLastJobOk() (*Job, bool) {
	if o == nil || IsNil(o.LastJob) {
		return nil, false
	}
	return o.LastJob, true
}

// HasLastJob returns a boolean if a field has been set.
func (o *TargetDTO) HasLastJob() bool {
	if o != nil && !IsNil(o.LastJob) {
		return true
	}

	return false
}

// SetLastJob gets a reference to the given Job and assigns it to the LastJob field.
func (o *TargetDTO) SetLastJob(v Job) {
	o.LastJob = &v
}

// GetLastJobId returns the LastJobId field value if set, zero value otherwise.
func (o *TargetDTO) GetLastJobId() string {
	if o == nil || IsNil(o.LastJobId) {
		var ret string
		return ret
	}
	return *o.LastJobId
}

// GetLastJobIdOk returns a tuple with the LastJobId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetDTO) GetLastJobIdOk() (*string, bool) {
	if o == nil || IsNil(o.LastJobId) {
		return nil, false
	}
	return o.LastJobId, true
}

// HasLastJobId returns a boolean if a field has been set.
func (o *TargetDTO) HasLastJobId() bool {
	if o != nil && !IsNil(o.LastJobId) {
		return true
	}

	return false
}

// SetLastJobId gets a reference to the given string and assigns it to the LastJobId field.
func (o *TargetDTO) SetLastJobId(v string) {
	o.LastJobId = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *TargetDTO) GetMetadata() TargetMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret TargetMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetDTO) GetMetadataOk() (*TargetMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *TargetDTO) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given TargetMetadata and assigns it to the Metadata field.
func (o *TargetDTO) SetMetadata(v TargetMetadata) {
	o.Metadata = &v
}

// GetName returns the Name field value
func (o *TargetDTO) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TargetDTO) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TargetDTO) SetName(v string) {
	o.Name = v
}

// GetProviderMetadata returns the ProviderMetadata field value if set, zero value otherwise.
func (o *TargetDTO) GetProviderMetadata() string {
	if o == nil || IsNil(o.ProviderMetadata) {
		var ret string
		return ret
	}
	return *o.ProviderMetadata
}

// GetProviderMetadataOk returns a tuple with the ProviderMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetDTO) GetProviderMetadataOk() (*string, bool) {
	if o == nil || IsNil(o.ProviderMetadata) {
		return nil, false
	}
	return o.ProviderMetadata, true
}

// HasProviderMetadata returns a boolean if a field has been set.
func (o *TargetDTO) HasProviderMetadata() bool {
	if o != nil && !IsNil(o.ProviderMetadata) {
		return true
	}

	return false
}

// SetProviderMetadata gets a reference to the given string and assigns it to the ProviderMetadata field.
func (o *TargetDTO) SetProviderMetadata(v string) {
	o.ProviderMetadata = &v
}

// GetState returns the State field value
func (o *TargetDTO) GetState() ResourceState {
	if o == nil {
		var ret ResourceState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *TargetDTO) GetStateOk() (*ResourceState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *TargetDTO) SetState(v ResourceState) {
	o.State = v
}

// GetTargetConfig returns the TargetConfig field value
func (o *TargetDTO) GetTargetConfig() TargetConfig {
	if o == nil {
		var ret TargetConfig
		return ret
	}

	return o.TargetConfig
}

// GetTargetConfigOk returns a tuple with the TargetConfig field value
// and a boolean to check if the value has been set.
func (o *TargetDTO) GetTargetConfigOk() (*TargetConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetConfig, true
}

// SetTargetConfig sets field value
func (o *TargetDTO) SetTargetConfig(v TargetConfig) {
	o.TargetConfig = v
}

// GetTargetConfigId returns the TargetConfigId field value
func (o *TargetDTO) GetTargetConfigId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetConfigId
}

// GetTargetConfigIdOk returns a tuple with the TargetConfigId field value
// and a boolean to check if the value has been set.
func (o *TargetDTO) GetTargetConfigIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetConfigId, true
}

// SetTargetConfigId sets field value
func (o *TargetDTO) SetTargetConfigId(v string) {
	o.TargetConfigId = v
}

// GetWorkspaces returns the Workspaces field value
func (o *TargetDTO) GetWorkspaces() []Workspace {
	if o == nil {
		var ret []Workspace
		return ret
	}

	return o.Workspaces
}

// GetWorkspacesOk returns a tuple with the Workspaces field value
// and a boolean to check if the value has been set.
func (o *TargetDTO) GetWorkspacesOk() ([]Workspace, bool) {
	if o == nil {
		return nil, false
	}
	return o.Workspaces, true
}

// SetWorkspaces sets field value
func (o *TargetDTO) SetWorkspaces(v []Workspace) {
	o.Workspaces = v
}

func (o TargetDTO) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TargetDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["default"] = o.Default
	toSerialize["envVars"] = o.EnvVars
	toSerialize["id"] = o.Id
	if !IsNil(o.LastJob) {
		toSerialize["lastJob"] = o.LastJob
	}
	if !IsNil(o.LastJobId) {
		toSerialize["lastJobId"] = o.LastJobId
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.ProviderMetadata) {
		toSerialize["providerMetadata"] = o.ProviderMetadata
	}
	toSerialize["state"] = o.State
	toSerialize["targetConfig"] = o.TargetConfig
	toSerialize["targetConfigId"] = o.TargetConfigId
	toSerialize["workspaces"] = o.Workspaces
	return toSerialize, nil
}

func (o *TargetDTO) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"default",
		"envVars",
		"id",
		"name",
		"state",
		"targetConfig",
		"targetConfigId",
		"workspaces",
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

	varTargetDTO := _TargetDTO{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTargetDTO)

	if err != nil {
		return err
	}

	*o = TargetDTO(varTargetDTO)

	return err
}

type NullableTargetDTO struct {
	value *TargetDTO
	isSet bool
}

func (v NullableTargetDTO) Get() *TargetDTO {
	return v.value
}

func (v *NullableTargetDTO) Set(val *TargetDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableTargetDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableTargetDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTargetDTO(val *TargetDTO) *NullableTargetDTO {
	return &NullableTargetDTO{value: val, isSet: true}
}

func (v NullableTargetDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTargetDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
