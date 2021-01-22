/*
 * madana-api
 *
 * <h1>Using the madana-api</h1>        <p>This documentation contains a Quickstart Guide, relating client functionality and information about the available         endpoints and used datamodels.   </p>       <p> The madana-api and its implementations are still in heavy development. This means that there may be problems in our protocols, or there may be mistakes in our implementations. We take security vulnerabilities very seriously. If you discover a security issue, please bring it to our attention right away! If you find a vulnerability that may affect live deployments -- for example, by exposing a remote execution exploit -- please send your report privately to info@madana.io. Please DO NOT file a public issue. If the issue is a protocol weakness that cannot be immediately exploited or something not yet deployed, just discuss it openly   </p>   <br>   <p> Note: Not all functionality might be acessible without having accquired and api-license token. For more information visit <a href=\"https://www.madana.io\">www.madana.io</a> </p>       <br>
 *
 * API version: 0.5.0-master.43
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package madanaapiclient

import (
	"encoding/json"
)

// XmlNs0MDNUserSettingAllOf struct for XmlNs0MDNUserSettingAllOf
type XmlNs0MDNUserSettingAllOf struct {
	// 
	Value *string `json:"value,omitempty"`
}

// NewXmlNs0MDNUserSettingAllOf instantiates a new XmlNs0MDNUserSettingAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewXmlNs0MDNUserSettingAllOf() *XmlNs0MDNUserSettingAllOf {
	this := XmlNs0MDNUserSettingAllOf{}
	return &this
}

// NewXmlNs0MDNUserSettingAllOfWithDefaults instantiates a new XmlNs0MDNUserSettingAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewXmlNs0MDNUserSettingAllOfWithDefaults() *XmlNs0MDNUserSettingAllOf {
	this := XmlNs0MDNUserSettingAllOf{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *XmlNs0MDNUserSettingAllOf) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0MDNUserSettingAllOf) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *XmlNs0MDNUserSettingAllOf) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *XmlNs0MDNUserSettingAllOf) SetValue(v string) {
	o.Value = &v
}

func (o XmlNs0MDNUserSettingAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableXmlNs0MDNUserSettingAllOf struct {
	value *XmlNs0MDNUserSettingAllOf
	isSet bool
}

func (v NullableXmlNs0MDNUserSettingAllOf) Get() *XmlNs0MDNUserSettingAllOf {
	return v.value
}

func (v *NullableXmlNs0MDNUserSettingAllOf) Set(val *XmlNs0MDNUserSettingAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableXmlNs0MDNUserSettingAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableXmlNs0MDNUserSettingAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableXmlNs0MDNUserSettingAllOf(val *XmlNs0MDNUserSettingAllOf) *NullableXmlNs0MDNUserSettingAllOf {
	return &NullableXmlNs0MDNUserSettingAllOf{value: val, isSet: true}
}

func (v NullableXmlNs0MDNUserSettingAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableXmlNs0MDNUserSettingAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


