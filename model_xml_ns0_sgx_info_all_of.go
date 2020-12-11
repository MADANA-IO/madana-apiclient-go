/*
 * madana-api
 *
 * <h1>Using the madana-api</h1>        <p>This documentation contains a Quickstart Guide, relating client functionality and information about the available         endpoints and used datamodels.   </p>       <p> The madana-api and its implementations are still in heavy development. This means that there may be problems in our protocols, or there may be mistakes in our implementations. We take security vulnerabilities very seriously. If you discover a security issue, please bring it to our attention right away! If you find a vulnerability that may affect live deployments -- for example, by exposing a remote execution exploit -- please send your report privately to info@madana.io. Please DO NOT file a public issue. If the issue is a protocol weakness that cannot be immediately exploited or something not yet deployed, just discuss it openly   </p>   <br>   <p> Note: Not all functionality might be acessible without having accquired and api-license token. For more information visit <a href=\"https://www.madana.io\">www.madana.io</a> </p>       <br>
 *
 * API version: 0.5.0-master.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package madanaapiclient

import (
	"encoding/json"
)

// XmlNs0SGXInfoAllOf struct for XmlNs0SGXInfoAllOf
type XmlNs0SGXInfoAllOf struct {
	// 
	Status *string `json:"status,omitempty"`
	// 
	Version *string `json:"version,omitempty"`
}

// NewXmlNs0SGXInfoAllOf instantiates a new XmlNs0SGXInfoAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewXmlNs0SGXInfoAllOf() *XmlNs0SGXInfoAllOf {
	this := XmlNs0SGXInfoAllOf{}
	return &this
}

// NewXmlNs0SGXInfoAllOfWithDefaults instantiates a new XmlNs0SGXInfoAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewXmlNs0SGXInfoAllOfWithDefaults() *XmlNs0SGXInfoAllOf {
	this := XmlNs0SGXInfoAllOf{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *XmlNs0SGXInfoAllOf) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0SGXInfoAllOf) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *XmlNs0SGXInfoAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *XmlNs0SGXInfoAllOf) SetStatus(v string) {
	o.Status = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *XmlNs0SGXInfoAllOf) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0SGXInfoAllOf) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *XmlNs0SGXInfoAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *XmlNs0SGXInfoAllOf) SetVersion(v string) {
	o.Version = &v
}

func (o XmlNs0SGXInfoAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableXmlNs0SGXInfoAllOf struct {
	value *XmlNs0SGXInfoAllOf
	isSet bool
}

func (v NullableXmlNs0SGXInfoAllOf) Get() *XmlNs0SGXInfoAllOf {
	return v.value
}

func (v *NullableXmlNs0SGXInfoAllOf) Set(val *XmlNs0SGXInfoAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableXmlNs0SGXInfoAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableXmlNs0SGXInfoAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableXmlNs0SGXInfoAllOf(val *XmlNs0SGXInfoAllOf) *NullableXmlNs0SGXInfoAllOf {
	return &NullableXmlNs0SGXInfoAllOf{value: val, isSet: true}
}

func (v NullableXmlNs0SGXInfoAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableXmlNs0SGXInfoAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


