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

// XmlNs0NetworkInterface 
type XmlNs0NetworkInterface struct {
	XmlNs0NetworkInterfaceAllOf
}

// NewXmlNs0NetworkInterface instantiates a new XmlNs0NetworkInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewXmlNs0NetworkInterface() *XmlNs0NetworkInterface {
	this := XmlNs0NetworkInterface{}
	return &this
}

// NewXmlNs0NetworkInterfaceWithDefaults instantiates a new XmlNs0NetworkInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewXmlNs0NetworkInterfaceWithDefaults() *XmlNs0NetworkInterface {
	this := XmlNs0NetworkInterface{}
	return &this
}

func (o XmlNs0NetworkInterface) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedXmlNs0NetworkInterfaceAllOf, errXmlNs0NetworkInterfaceAllOf := json.Marshal(o.XmlNs0NetworkInterfaceAllOf)
	if errXmlNs0NetworkInterfaceAllOf != nil {
		return []byte{}, errXmlNs0NetworkInterfaceAllOf
	}
	errXmlNs0NetworkInterfaceAllOf = json.Unmarshal([]byte(serializedXmlNs0NetworkInterfaceAllOf), &toSerialize)
	if errXmlNs0NetworkInterfaceAllOf != nil {
		return []byte{}, errXmlNs0NetworkInterfaceAllOf
	}
	return json.Marshal(toSerialize)
}

type NullableXmlNs0NetworkInterface struct {
	value *XmlNs0NetworkInterface
	isSet bool
}

func (v NullableXmlNs0NetworkInterface) Get() *XmlNs0NetworkInterface {
	return v.value
}

func (v *NullableXmlNs0NetworkInterface) Set(val *XmlNs0NetworkInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableXmlNs0NetworkInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableXmlNs0NetworkInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableXmlNs0NetworkInterface(val *XmlNs0NetworkInterface) *NullableXmlNs0NetworkInterface {
	return &NullableXmlNs0NetworkInterface{value: val, isSet: true}
}

func (v NullableXmlNs0NetworkInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableXmlNs0NetworkInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


