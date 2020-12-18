/*
 * madana-api
 *
 * <h1>Using the madana-api</h1>        <p>This documentation contains a Quickstart Guide, relating client functionality and information about the available         endpoints and used datamodels.   </p>       <p> The madana-api and its implementations are still in heavy development. This means that there may be problems in our protocols, or there may be mistakes in our implementations. We take security vulnerabilities very seriously. If you discover a security issue, please bring it to our attention right away! If you find a vulnerability that may affect live deployments -- for example, by exposing a remote execution exploit -- please send your report privately to info@madana.io. Please DO NOT file a public issue. If the issue is a protocol weakness that cannot be immediately exploited or something not yet deployed, just discuss it openly   </p>   <br>   <p> Note: Not all functionality might be acessible without having accquired and api-license token. For more information visit <a href=\"https://www.madana.io\">www.madana.io</a> </p>       <br>
 *
 * API version: 0.5.0-master.16
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package madanaapiclient

import (
	"encoding/json"
)

// XmlNs0DiskConfig 
type XmlNs0DiskConfig struct {
	XmlNs0DiskConfigAllOf
}

// NewXmlNs0DiskConfig instantiates a new XmlNs0DiskConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewXmlNs0DiskConfig() *XmlNs0DiskConfig {
	this := XmlNs0DiskConfig{}
	return &this
}

// NewXmlNs0DiskConfigWithDefaults instantiates a new XmlNs0DiskConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewXmlNs0DiskConfigWithDefaults() *XmlNs0DiskConfig {
	this := XmlNs0DiskConfig{}
	return &this
}

func (o XmlNs0DiskConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedXmlNs0DiskConfigAllOf, errXmlNs0DiskConfigAllOf := json.Marshal(o.XmlNs0DiskConfigAllOf)
	if errXmlNs0DiskConfigAllOf != nil {
		return []byte{}, errXmlNs0DiskConfigAllOf
	}
	errXmlNs0DiskConfigAllOf = json.Unmarshal([]byte(serializedXmlNs0DiskConfigAllOf), &toSerialize)
	if errXmlNs0DiskConfigAllOf != nil {
		return []byte{}, errXmlNs0DiskConfigAllOf
	}
	return json.Marshal(toSerialize)
}

type NullableXmlNs0DiskConfig struct {
	value *XmlNs0DiskConfig
	isSet bool
}

func (v NullableXmlNs0DiskConfig) Get() *XmlNs0DiskConfig {
	return v.value
}

func (v *NullableXmlNs0DiskConfig) Set(val *XmlNs0DiskConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableXmlNs0DiskConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableXmlNs0DiskConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableXmlNs0DiskConfig(val *XmlNs0DiskConfig) *NullableXmlNs0DiskConfig {
	return &NullableXmlNs0DiskConfig{value: val, isSet: true}
}

func (v NullableXmlNs0DiskConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableXmlNs0DiskConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


