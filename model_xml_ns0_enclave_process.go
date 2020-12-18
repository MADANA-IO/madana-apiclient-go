/*
 * madana-api
 *
 * <h1>Using the madana-api</h1>        <p>This documentation contains a Quickstart Guide, relating client functionality and information about the available         endpoints and used datamodels.   </p>       <p> The madana-api and its implementations are still in heavy development. This means that there may be problems in our protocols, or there may be mistakes in our implementations. We take security vulnerabilities very seriously. If you discover a security issue, please bring it to our attention right away! If you find a vulnerability that may affect live deployments -- for example, by exposing a remote execution exploit -- please send your report privately to info@madana.io. Please DO NOT file a public issue. If the issue is a protocol weakness that cannot be immediately exploited or something not yet deployed, just discuss it openly   </p>   <br>   <p> Note: Not all functionality might be acessible without having accquired and api-license token. For more information visit <a href=\"https://www.madana.io\">www.madana.io</a> </p>       <br>
 *
 * API version: 0.5.0-master.15
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package madanaapiclient

import (
	"encoding/json"
)

// XmlNs0EnclaveProcess 
type XmlNs0EnclaveProcess struct {
	XmlNs0EnclaveProcessAllOf
}

// NewXmlNs0EnclaveProcess instantiates a new XmlNs0EnclaveProcess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewXmlNs0EnclaveProcess() *XmlNs0EnclaveProcess {
	this := XmlNs0EnclaveProcess{}
	return &this
}

// NewXmlNs0EnclaveProcessWithDefaults instantiates a new XmlNs0EnclaveProcess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewXmlNs0EnclaveProcessWithDefaults() *XmlNs0EnclaveProcess {
	this := XmlNs0EnclaveProcess{}
	return &this
}

func (o XmlNs0EnclaveProcess) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedXmlNs0EnclaveProcessAllOf, errXmlNs0EnclaveProcessAllOf := json.Marshal(o.XmlNs0EnclaveProcessAllOf)
	if errXmlNs0EnclaveProcessAllOf != nil {
		return []byte{}, errXmlNs0EnclaveProcessAllOf
	}
	errXmlNs0EnclaveProcessAllOf = json.Unmarshal([]byte(serializedXmlNs0EnclaveProcessAllOf), &toSerialize)
	if errXmlNs0EnclaveProcessAllOf != nil {
		return []byte{}, errXmlNs0EnclaveProcessAllOf
	}
	return json.Marshal(toSerialize)
}

type NullableXmlNs0EnclaveProcess struct {
	value *XmlNs0EnclaveProcess
	isSet bool
}

func (v NullableXmlNs0EnclaveProcess) Get() *XmlNs0EnclaveProcess {
	return v.value
}

func (v *NullableXmlNs0EnclaveProcess) Set(val *XmlNs0EnclaveProcess) {
	v.value = val
	v.isSet = true
}

func (v NullableXmlNs0EnclaveProcess) IsSet() bool {
	return v.isSet
}

func (v *NullableXmlNs0EnclaveProcess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableXmlNs0EnclaveProcess(val *XmlNs0EnclaveProcess) *NullableXmlNs0EnclaveProcess {
	return &NullableXmlNs0EnclaveProcess{value: val, isSet: true}
}

func (v NullableXmlNs0EnclaveProcess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableXmlNs0EnclaveProcess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


