/*
 * madana-api
 *
 * <h1>Using the madana-api</h1>        <p>This documentation contains a Quickstart Guide, relating client functionality and information about the available         endpoints and used datamodels.   </p>       <p> The madana-api and its implementations are still in heavy development. This means that there may be problems in our protocols, or there may be mistakes in our implementations. We take security vulnerabilities very seriously. If you discover a security issue, please bring it to our attention right away! If you find a vulnerability that may affect live deployments -- for example, by exposing a remote execution exploit -- please send your report privately to info@madana.io. Please DO NOT file a public issue. If the issue is a protocol weakness that cannot be immediately exploited or something not yet deployed, just discuss it openly   </p>   <br>   <p> Note: Not all functionality might be acessible without having accquired and api-license token. For more information visit <a href=\"https://www.madana.io\">www.madana.io</a> </p>       <br>
 *
 * API version: 0.5.0-master.56
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package madanaapiclient

import (
	"encoding/json"
)

// JsonV1EventSource 
type JsonV1EventSource struct {
	// 
	Component *string `json:"component,omitempty"`
	// 
	Host *string `json:"host,omitempty"`
}

// NewJsonV1EventSource instantiates a new JsonV1EventSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJsonV1EventSource() *JsonV1EventSource {
	this := JsonV1EventSource{}
	return &this
}

// NewJsonV1EventSourceWithDefaults instantiates a new JsonV1EventSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJsonV1EventSourceWithDefaults() *JsonV1EventSource {
	this := JsonV1EventSource{}
	return &this
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *JsonV1EventSource) GetComponent() string {
	if o == nil || o.Component == nil {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonV1EventSource) GetComponentOk() (*string, bool) {
	if o == nil || o.Component == nil {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *JsonV1EventSource) HasComponent() bool {
	if o != nil && o.Component != nil {
		return true
	}

	return false
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *JsonV1EventSource) SetComponent(v string) {
	o.Component = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *JsonV1EventSource) GetHost() string {
	if o == nil || o.Host == nil {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonV1EventSource) GetHostOk() (*string, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *JsonV1EventSource) HasHost() bool {
	if o != nil && o.Host != nil {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *JsonV1EventSource) SetHost(v string) {
	o.Host = &v
}

func (o JsonV1EventSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Component != nil {
		toSerialize["component"] = o.Component
	}
	if o.Host != nil {
		toSerialize["host"] = o.Host
	}
	return json.Marshal(toSerialize)
}

type NullableJsonV1EventSource struct {
	value *JsonV1EventSource
	isSet bool
}

func (v NullableJsonV1EventSource) Get() *JsonV1EventSource {
	return v.value
}

func (v *NullableJsonV1EventSource) Set(val *JsonV1EventSource) {
	v.value = val
	v.isSet = true
}

func (v NullableJsonV1EventSource) IsSet() bool {
	return v.isSet
}

func (v *NullableJsonV1EventSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJsonV1EventSource(val *JsonV1EventSource) *NullableJsonV1EventSource {
	return &NullableJsonV1EventSource{value: val, isSet: true}
}

func (v NullableJsonV1EventSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJsonV1EventSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


