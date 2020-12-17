/*
 * madana-api
 *
 * <h1>Using the madana-api</h1>        <p>This documentation contains a Quickstart Guide, relating client functionality and information about the available         endpoints and used datamodels.   </p>       <p> The madana-api and its implementations are still in heavy development. This means that there may be problems in our protocols, or there may be mistakes in our implementations. We take security vulnerabilities very seriously. If you discover a security issue, please bring it to our attention right away! If you find a vulnerability that may affect live deployments -- for example, by exposing a remote execution exploit -- please send your report privately to info@madana.io. Please DO NOT file a public issue. If the issue is a protocol weakness that cannot be immediately exploited or something not yet deployed, just discuss it openly   </p>   <br>   <p> Note: Not all functionality might be acessible without having accquired and api-license token. For more information visit <a href=\"https://www.madana.io\">www.madana.io</a> </p>       <br>
 *
 * API version: 0.5.0-master.8
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package madanaapiclient

import (
	"encoding/json"
)

// JsonSGXInfo 
type JsonSGXInfo struct {
	// 
	Version *string `json:"version,omitempty"`
	// 
	Status *string `json:"status,omitempty"`
}

// NewJsonSGXInfo instantiates a new JsonSGXInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJsonSGXInfo() *JsonSGXInfo {
	this := JsonSGXInfo{}
	return &this
}

// NewJsonSGXInfoWithDefaults instantiates a new JsonSGXInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJsonSGXInfoWithDefaults() *JsonSGXInfo {
	this := JsonSGXInfo{}
	return &this
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *JsonSGXInfo) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonSGXInfo) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *JsonSGXInfo) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *JsonSGXInfo) SetVersion(v string) {
	o.Version = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *JsonSGXInfo) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonSGXInfo) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *JsonSGXInfo) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *JsonSGXInfo) SetStatus(v string) {
	o.Status = &v
}

func (o JsonSGXInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableJsonSGXInfo struct {
	value *JsonSGXInfo
	isSet bool
}

func (v NullableJsonSGXInfo) Get() *JsonSGXInfo {
	return v.value
}

func (v *NullableJsonSGXInfo) Set(val *JsonSGXInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableJsonSGXInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableJsonSGXInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJsonSGXInfo(val *JsonSGXInfo) *NullableJsonSGXInfo {
	return &NullableJsonSGXInfo{value: val, isSet: true}
}

func (v NullableJsonSGXInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJsonSGXInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

