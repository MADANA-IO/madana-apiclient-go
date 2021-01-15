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

// XmlNs0RunConfig 
type XmlNs0RunConfig struct {
	// 
	Args *[]string `json:"args,omitempty"`
	// 
	DiskConfig *[]XmlNs0DiskConfig `json:"disk_config,omitempty"`
	// 
	Environment *map[string]interface{} `json:"environment,omitempty"`
	// 
	Run *string `json:"run,omitempty"`
}

// NewXmlNs0RunConfig instantiates a new XmlNs0RunConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewXmlNs0RunConfig() *XmlNs0RunConfig {
	this := XmlNs0RunConfig{}
	return &this
}

// NewXmlNs0RunConfigWithDefaults instantiates a new XmlNs0RunConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewXmlNs0RunConfigWithDefaults() *XmlNs0RunConfig {
	this := XmlNs0RunConfig{}
	return &this
}

// GetArgs returns the Args field value if set, zero value otherwise.
func (o *XmlNs0RunConfig) GetArgs() []string {
	if o == nil || o.Args == nil {
		var ret []string
		return ret
	}
	return *o.Args
}

// GetArgsOk returns a tuple with the Args field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0RunConfig) GetArgsOk() (*[]string, bool) {
	if o == nil || o.Args == nil {
		return nil, false
	}
	return o.Args, true
}

// HasArgs returns a boolean if a field has been set.
func (o *XmlNs0RunConfig) HasArgs() bool {
	if o != nil && o.Args != nil {
		return true
	}

	return false
}

// SetArgs gets a reference to the given []string and assigns it to the Args field.
func (o *XmlNs0RunConfig) SetArgs(v []string) {
	o.Args = &v
}

// GetDiskConfig returns the DiskConfig field value if set, zero value otherwise.
func (o *XmlNs0RunConfig) GetDiskConfig() []XmlNs0DiskConfig {
	if o == nil || o.DiskConfig == nil {
		var ret []XmlNs0DiskConfig
		return ret
	}
	return *o.DiskConfig
}

// GetDiskConfigOk returns a tuple with the DiskConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0RunConfig) GetDiskConfigOk() (*[]XmlNs0DiskConfig, bool) {
	if o == nil || o.DiskConfig == nil {
		return nil, false
	}
	return o.DiskConfig, true
}

// HasDiskConfig returns a boolean if a field has been set.
func (o *XmlNs0RunConfig) HasDiskConfig() bool {
	if o != nil && o.DiskConfig != nil {
		return true
	}

	return false
}

// SetDiskConfig gets a reference to the given []XmlNs0DiskConfig and assigns it to the DiskConfig field.
func (o *XmlNs0RunConfig) SetDiskConfig(v []XmlNs0DiskConfig) {
	o.DiskConfig = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *XmlNs0RunConfig) GetEnvironment() map[string]interface{} {
	if o == nil || o.Environment == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0RunConfig) GetEnvironmentOk() (*map[string]interface{}, bool) {
	if o == nil || o.Environment == nil {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *XmlNs0RunConfig) HasEnvironment() bool {
	if o != nil && o.Environment != nil {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given map[string]interface{} and assigns it to the Environment field.
func (o *XmlNs0RunConfig) SetEnvironment(v map[string]interface{}) {
	o.Environment = &v
}

// GetRun returns the Run field value if set, zero value otherwise.
func (o *XmlNs0RunConfig) GetRun() string {
	if o == nil || o.Run == nil {
		var ret string
		return ret
	}
	return *o.Run
}

// GetRunOk returns a tuple with the Run field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0RunConfig) GetRunOk() (*string, bool) {
	if o == nil || o.Run == nil {
		return nil, false
	}
	return o.Run, true
}

// HasRun returns a boolean if a field has been set.
func (o *XmlNs0RunConfig) HasRun() bool {
	if o != nil && o.Run != nil {
		return true
	}

	return false
}

// SetRun gets a reference to the given string and assigns it to the Run field.
func (o *XmlNs0RunConfig) SetRun(v string) {
	o.Run = &v
}

func (o XmlNs0RunConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Args != nil {
		toSerialize["args"] = o.Args
	}
	if o.DiskConfig != nil {
		toSerialize["disk_config"] = o.DiskConfig
	}
	if o.Environment != nil {
		toSerialize["environment"] = o.Environment
	}
	if o.Run != nil {
		toSerialize["run"] = o.Run
	}
	return json.Marshal(toSerialize)
}

type NullableXmlNs0RunConfig struct {
	value *XmlNs0RunConfig
	isSet bool
}

func (v NullableXmlNs0RunConfig) Get() *XmlNs0RunConfig {
	return v.value
}

func (v *NullableXmlNs0RunConfig) Set(val *XmlNs0RunConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableXmlNs0RunConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableXmlNs0RunConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableXmlNs0RunConfig(val *XmlNs0RunConfig) *NullableXmlNs0RunConfig {
	return &NullableXmlNs0RunConfig{value: val, isSet: true}
}

func (v NullableXmlNs0RunConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableXmlNs0RunConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


