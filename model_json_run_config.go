/*
 * madana-api
 *
 * <h1>Using the madana-api</h1>        <p>This documentation contains a Quickstart Guide, relating client functionality and information about the available         endpoints and used datamodels.   </p>       <p> The madana-api and its implementations are still in heavy development. This means that there may be problems in our protocols, or there may be mistakes in our implementations. We take security vulnerabilities very seriously. If you discover a security issue, please bring it to our attention right away! If you find a vulnerability that may affect live deployments -- for example, by exposing a remote execution exploit -- please send your report privately to info@madana.io. Please DO NOT file a public issue. If the issue is a protocol weakness that cannot be immediately exploited or something not yet deployed, just discuss it openly   </p>   <br>   <p> Note: Not all functionality might be acessible without having accquired and api-license token. For more information visit <a href=\"https://www.madana.io\">www.madana.io</a> </p>       <br>
 *
 * API version: 0.5.0-master.34
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package madanaapiclient

import (
	"encoding/json"
)

// JsonRunConfig 
type JsonRunConfig struct {
	// 
	Environment *map[string]string `json:"environment,omitempty"`
	// 
	Run *string `json:"run,omitempty"`
	// 
	DiskConfig *[]JsonDiskConfig `json:"disk_config,omitempty"`
	// 
	Args *[]string `json:"args,omitempty"`
}

// NewJsonRunConfig instantiates a new JsonRunConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJsonRunConfig() *JsonRunConfig {
	this := JsonRunConfig{}
	return &this
}

// NewJsonRunConfigWithDefaults instantiates a new JsonRunConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJsonRunConfigWithDefaults() *JsonRunConfig {
	this := JsonRunConfig{}
	return &this
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *JsonRunConfig) GetEnvironment() map[string]string {
	if o == nil || o.Environment == nil {
		var ret map[string]string
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonRunConfig) GetEnvironmentOk() (*map[string]string, bool) {
	if o == nil || o.Environment == nil {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *JsonRunConfig) HasEnvironment() bool {
	if o != nil && o.Environment != nil {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given map[string]string and assigns it to the Environment field.
func (o *JsonRunConfig) SetEnvironment(v map[string]string) {
	o.Environment = &v
}

// GetRun returns the Run field value if set, zero value otherwise.
func (o *JsonRunConfig) GetRun() string {
	if o == nil || o.Run == nil {
		var ret string
		return ret
	}
	return *o.Run
}

// GetRunOk returns a tuple with the Run field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonRunConfig) GetRunOk() (*string, bool) {
	if o == nil || o.Run == nil {
		return nil, false
	}
	return o.Run, true
}

// HasRun returns a boolean if a field has been set.
func (o *JsonRunConfig) HasRun() bool {
	if o != nil && o.Run != nil {
		return true
	}

	return false
}

// SetRun gets a reference to the given string and assigns it to the Run field.
func (o *JsonRunConfig) SetRun(v string) {
	o.Run = &v
}

// GetDiskConfig returns the DiskConfig field value if set, zero value otherwise.
func (o *JsonRunConfig) GetDiskConfig() []JsonDiskConfig {
	if o == nil || o.DiskConfig == nil {
		var ret []JsonDiskConfig
		return ret
	}
	return *o.DiskConfig
}

// GetDiskConfigOk returns a tuple with the DiskConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonRunConfig) GetDiskConfigOk() (*[]JsonDiskConfig, bool) {
	if o == nil || o.DiskConfig == nil {
		return nil, false
	}
	return o.DiskConfig, true
}

// HasDiskConfig returns a boolean if a field has been set.
func (o *JsonRunConfig) HasDiskConfig() bool {
	if o != nil && o.DiskConfig != nil {
		return true
	}

	return false
}

// SetDiskConfig gets a reference to the given []JsonDiskConfig and assigns it to the DiskConfig field.
func (o *JsonRunConfig) SetDiskConfig(v []JsonDiskConfig) {
	o.DiskConfig = &v
}

// GetArgs returns the Args field value if set, zero value otherwise.
func (o *JsonRunConfig) GetArgs() []string {
	if o == nil || o.Args == nil {
		var ret []string
		return ret
	}
	return *o.Args
}

// GetArgsOk returns a tuple with the Args field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonRunConfig) GetArgsOk() (*[]string, bool) {
	if o == nil || o.Args == nil {
		return nil, false
	}
	return o.Args, true
}

// HasArgs returns a boolean if a field has been set.
func (o *JsonRunConfig) HasArgs() bool {
	if o != nil && o.Args != nil {
		return true
	}

	return false
}

// SetArgs gets a reference to the given []string and assigns it to the Args field.
func (o *JsonRunConfig) SetArgs(v []string) {
	o.Args = &v
}

func (o JsonRunConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Environment != nil {
		toSerialize["environment"] = o.Environment
	}
	if o.Run != nil {
		toSerialize["run"] = o.Run
	}
	if o.DiskConfig != nil {
		toSerialize["disk_config"] = o.DiskConfig
	}
	if o.Args != nil {
		toSerialize["args"] = o.Args
	}
	return json.Marshal(toSerialize)
}

type NullableJsonRunConfig struct {
	value *JsonRunConfig
	isSet bool
}

func (v NullableJsonRunConfig) Get() *JsonRunConfig {
	return v.value
}

func (v *NullableJsonRunConfig) Set(val *JsonRunConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableJsonRunConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableJsonRunConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJsonRunConfig(val *JsonRunConfig) *NullableJsonRunConfig {
	return &NullableJsonRunConfig{value: val, isSet: true}
}

func (v NullableJsonRunConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJsonRunConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


