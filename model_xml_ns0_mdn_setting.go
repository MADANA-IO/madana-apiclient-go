/*
 * madana-api
 *
 * <h1>Using the madana-api</h1>        <p>This documentation contains a Quickstart Guide, relating client functionality and information about the available         endpoints and used datamodels.   </p>       <p> The madana-api and its implementations are still in heavy development. This means that there may be problems in our protocols, or there may be mistakes in our implementations. We take security vulnerabilities very seriously. If you discover a security issue, please bring it to our attention right away! If you find a vulnerability that may affect live deployments -- for example, by exposing a remote execution exploit -- please send your report privately to info@madana.io. Please DO NOT file a public issue. If the issue is a protocol weakness that cannot be immediately exploited or something not yet deployed, just discuss it openly   </p>   <br>   <p> Note: Not all functionality might be acessible without having accquired and api-license token. For more information visit <a href=\"https://www.madana.io\">www.madana.io</a> </p>       <br>
 *
 * API version: 0.5.0-master.52
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package madanaapiclient

import (
	"encoding/json"
)

// XmlNs0MDNSetting 
type XmlNs0MDNSetting struct {
	// 
	Description *string `json:"description,omitempty"`
	// 
	Id *string `json:"id,omitempty"`
	// 
	Name *string `json:"name,omitempty"`
}

// NewXmlNs0MDNSetting instantiates a new XmlNs0MDNSetting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewXmlNs0MDNSetting() *XmlNs0MDNSetting {
	this := XmlNs0MDNSetting{}
	return &this
}

// NewXmlNs0MDNSettingWithDefaults instantiates a new XmlNs0MDNSetting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewXmlNs0MDNSettingWithDefaults() *XmlNs0MDNSetting {
	this := XmlNs0MDNSetting{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *XmlNs0MDNSetting) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0MDNSetting) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *XmlNs0MDNSetting) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *XmlNs0MDNSetting) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *XmlNs0MDNSetting) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0MDNSetting) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *XmlNs0MDNSetting) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *XmlNs0MDNSetting) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *XmlNs0MDNSetting) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0MDNSetting) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *XmlNs0MDNSetting) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *XmlNs0MDNSetting) SetName(v string) {
	o.Name = &v
}

func (o XmlNs0MDNSetting) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableXmlNs0MDNSetting struct {
	value *XmlNs0MDNSetting
	isSet bool
}

func (v NullableXmlNs0MDNSetting) Get() *XmlNs0MDNSetting {
	return v.value
}

func (v *NullableXmlNs0MDNSetting) Set(val *XmlNs0MDNSetting) {
	v.value = val
	v.isSet = true
}

func (v NullableXmlNs0MDNSetting) IsSet() bool {
	return v.isSet
}

func (v *NullableXmlNs0MDNSetting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableXmlNs0MDNSetting(val *XmlNs0MDNSetting) *NullableXmlNs0MDNSetting {
	return &NullableXmlNs0MDNSetting{value: val, isSet: true}
}

func (v NullableXmlNs0MDNSetting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableXmlNs0MDNSetting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


