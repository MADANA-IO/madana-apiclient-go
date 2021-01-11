/*
 * madana-api
 *
 * <h1>Using the madana-api</h1>        <p>This documentation contains a Quickstart Guide, relating client functionality and information about the available         endpoints and used datamodels.   </p>       <p> The madana-api and its implementations are still in heavy development. This means that there may be problems in our protocols, or there may be mistakes in our implementations. We take security vulnerabilities very seriously. If you discover a security issue, please bring it to our attention right away! If you find a vulnerability that may affect live deployments -- for example, by exposing a remote execution exploit -- please send your report privately to info@madana.io. Please DO NOT file a public issue. If the issue is a protocol weakness that cannot be immediately exploited or something not yet deployed, just discuss it openly   </p>   <br>   <p> Note: Not all functionality might be acessible without having accquired and api-license token. For more information visit <a href=\"https://www.madana.io\">www.madana.io</a> </p>       <br>
 *
 * API version: 0.5.0-master.23
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package madanaapiclient

import (
	"encoding/json"
)

// JsonEnclavePort 
type JsonEnclavePort struct {
	// 
	Port *string `json:"port,omitempty"`
	// 
	Name *string `json:"name,omitempty"`
	// 
	Protocol *string `json:"protocol,omitempty"`
}

// NewJsonEnclavePort instantiates a new JsonEnclavePort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJsonEnclavePort() *JsonEnclavePort {
	this := JsonEnclavePort{}
	return &this
}

// NewJsonEnclavePortWithDefaults instantiates a new JsonEnclavePort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJsonEnclavePortWithDefaults() *JsonEnclavePort {
	this := JsonEnclavePort{}
	return &this
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *JsonEnclavePort) GetPort() string {
	if o == nil || o.Port == nil {
		var ret string
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonEnclavePort) GetPortOk() (*string, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *JsonEnclavePort) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given string and assigns it to the Port field.
func (o *JsonEnclavePort) SetPort(v string) {
	o.Port = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *JsonEnclavePort) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonEnclavePort) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *JsonEnclavePort) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *JsonEnclavePort) SetName(v string) {
	o.Name = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *JsonEnclavePort) GetProtocol() string {
	if o == nil || o.Protocol == nil {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonEnclavePort) GetProtocolOk() (*string, bool) {
	if o == nil || o.Protocol == nil {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *JsonEnclavePort) HasProtocol() bool {
	if o != nil && o.Protocol != nil {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *JsonEnclavePort) SetProtocol(v string) {
	o.Protocol = &v
}

func (o JsonEnclavePort) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Port != nil {
		toSerialize["port"] = o.Port
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Protocol != nil {
		toSerialize["protocol"] = o.Protocol
	}
	return json.Marshal(toSerialize)
}

type NullableJsonEnclavePort struct {
	value *JsonEnclavePort
	isSet bool
}

func (v NullableJsonEnclavePort) Get() *JsonEnclavePort {
	return v.value
}

func (v *NullableJsonEnclavePort) Set(val *JsonEnclavePort) {
	v.value = val
	v.isSet = true
}

func (v NullableJsonEnclavePort) IsSet() bool {
	return v.isSet
}

func (v *NullableJsonEnclavePort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJsonEnclavePort(val *JsonEnclavePort) *NullableJsonEnclavePort {
	return &NullableJsonEnclavePort{value: val, isSet: true}
}

func (v NullableJsonEnclavePort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJsonEnclavePort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


