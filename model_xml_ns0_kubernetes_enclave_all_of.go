/*
 * madana-api
 *
 * <h1>Using the madana-api</h1>        <p>This documentation contains a Quickstart Guide, relating client functionality and information about the available         endpoints and used datamodels.   </p>       <p> The madana-api and its implementations are still in heavy development. This means that there may be problems in our protocols, or there may be mistakes in our implementations. We take security vulnerabilities very seriously. If you discover a security issue, please bring it to our attention right away! If you find a vulnerability that may affect live deployments -- for example, by exposing a remote execution exploit -- please send your report privately to info@madana.io. Please DO NOT file a public issue. If the issue is a protocol weakness that cannot be immediately exploited or something not yet deployed, just discuss it openly   </p>   <br>   <p> Note: Not all functionality might be acessible without having accquired and api-license token. For more information visit <a href=\"https://www.madana.io\">www.madana.io</a> </p>       <br>
 *
 * API version: 0.5.0-master.28
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package madanaapiclient

import (
	"encoding/json"
)

// XmlNs0KubernetesEnclaveAllOf struct for XmlNs0KubernetesEnclaveAllOf
type XmlNs0KubernetesEnclaveAllOf struct {
	// 
	IsUsingInitContainer *bool `json:"isUsingInitContainer,omitempty"`
}

// NewXmlNs0KubernetesEnclaveAllOf instantiates a new XmlNs0KubernetesEnclaveAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewXmlNs0KubernetesEnclaveAllOf() *XmlNs0KubernetesEnclaveAllOf {
	this := XmlNs0KubernetesEnclaveAllOf{}
	return &this
}

// NewXmlNs0KubernetesEnclaveAllOfWithDefaults instantiates a new XmlNs0KubernetesEnclaveAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewXmlNs0KubernetesEnclaveAllOfWithDefaults() *XmlNs0KubernetesEnclaveAllOf {
	this := XmlNs0KubernetesEnclaveAllOf{}
	return &this
}

// GetIsUsingInitContainer returns the IsUsingInitContainer field value if set, zero value otherwise.
func (o *XmlNs0KubernetesEnclaveAllOf) GetIsUsingInitContainer() bool {
	if o == nil || o.IsUsingInitContainer == nil {
		var ret bool
		return ret
	}
	return *o.IsUsingInitContainer
}

// GetIsUsingInitContainerOk returns a tuple with the IsUsingInitContainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0KubernetesEnclaveAllOf) GetIsUsingInitContainerOk() (*bool, bool) {
	if o == nil || o.IsUsingInitContainer == nil {
		return nil, false
	}
	return o.IsUsingInitContainer, true
}

// HasIsUsingInitContainer returns a boolean if a field has been set.
func (o *XmlNs0KubernetesEnclaveAllOf) HasIsUsingInitContainer() bool {
	if o != nil && o.IsUsingInitContainer != nil {
		return true
	}

	return false
}

// SetIsUsingInitContainer gets a reference to the given bool and assigns it to the IsUsingInitContainer field.
func (o *XmlNs0KubernetesEnclaveAllOf) SetIsUsingInitContainer(v bool) {
	o.IsUsingInitContainer = &v
}

func (o XmlNs0KubernetesEnclaveAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IsUsingInitContainer != nil {
		toSerialize["isUsingInitContainer"] = o.IsUsingInitContainer
	}
	return json.Marshal(toSerialize)
}

type NullableXmlNs0KubernetesEnclaveAllOf struct {
	value *XmlNs0KubernetesEnclaveAllOf
	isSet bool
}

func (v NullableXmlNs0KubernetesEnclaveAllOf) Get() *XmlNs0KubernetesEnclaveAllOf {
	return v.value
}

func (v *NullableXmlNs0KubernetesEnclaveAllOf) Set(val *XmlNs0KubernetesEnclaveAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableXmlNs0KubernetesEnclaveAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableXmlNs0KubernetesEnclaveAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableXmlNs0KubernetesEnclaveAllOf(val *XmlNs0KubernetesEnclaveAllOf) *NullableXmlNs0KubernetesEnclaveAllOf {
	return &NullableXmlNs0KubernetesEnclaveAllOf{value: val, isSet: true}
}

func (v NullableXmlNs0KubernetesEnclaveAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableXmlNs0KubernetesEnclaveAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


