/*
 * madana-api
 *
 * <h1>Using the madana-api</h1>        <p>This documentation contains a Quickstart Guide, relating client functionality and information about the available         endpoints and used datamodels.   </p>       <p> The madana-api and its implementations are still in heavy development. This means that there may be problems in our protocols, or there may be mistakes in our implementations. We take security vulnerabilities very seriously. If you discover a security issue, please bring it to our attention right away! If you find a vulnerability that may affect live deployments -- for example, by exposing a remote execution exploit -- please send your report privately to info@madana.io. Please DO NOT file a public issue. If the issue is a protocol weakness that cannot be immediately exploited or something not yet deployed, just discuss it openly   </p>   <br>   <p> Note: Not all functionality might be acessible without having accquired and api-license token. For more information visit <a href=\"https://www.madana.io\">www.madana.io</a> </p>       <br>
 *
 * API version: 0.5.0-master.34
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package madanaapiclient

import (
	"encoding/json"
)

// XmlNs0MDNUserSetting 
type XmlNs0MDNUserSetting struct {
	// 
	Description *string `json:"description,omitempty"`
	// 
	Id *string `json:"id,omitempty"`
	// 
	Name *string `json:"name,omitempty"`
	// 
	Value *string `json:"value,omitempty"`
}

// NewXmlNs0MDNUserSetting instantiates a new XmlNs0MDNUserSetting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewXmlNs0MDNUserSetting() *XmlNs0MDNUserSetting {
	this := XmlNs0MDNUserSetting{}
	return &this
}

// NewXmlNs0MDNUserSettingWithDefaults instantiates a new XmlNs0MDNUserSetting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewXmlNs0MDNUserSettingWithDefaults() *XmlNs0MDNUserSetting {
	this := XmlNs0MDNUserSetting{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *XmlNs0MDNUserSetting) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0MDNUserSetting) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *XmlNs0MDNUserSetting) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *XmlNs0MDNUserSetting) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *XmlNs0MDNUserSetting) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0MDNUserSetting) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *XmlNs0MDNUserSetting) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *XmlNs0MDNUserSetting) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *XmlNs0MDNUserSetting) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0MDNUserSetting) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *XmlNs0MDNUserSetting) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *XmlNs0MDNUserSetting) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *XmlNs0MDNUserSetting) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0MDNUserSetting) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *XmlNs0MDNUserSetting) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *XmlNs0MDNUserSetting) SetValue(v string) {
	o.Value = &v
}

func (o XmlNs0MDNUserSetting) MarshalJSON() ([]byte, error) {
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
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableXmlNs0MDNUserSetting struct {
	value *XmlNs0MDNUserSetting
	isSet bool
}

func (v NullableXmlNs0MDNUserSetting) Get() *XmlNs0MDNUserSetting {
	return v.value
}

func (v *NullableXmlNs0MDNUserSetting) Set(val *XmlNs0MDNUserSetting) {
	v.value = val
	v.isSet = true
}

func (v NullableXmlNs0MDNUserSetting) IsSet() bool {
	return v.isSet
}

func (v *NullableXmlNs0MDNUserSetting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableXmlNs0MDNUserSetting(val *XmlNs0MDNUserSetting) *NullableXmlNs0MDNUserSetting {
	return &NullableXmlNs0MDNUserSetting{value: val, isSet: true}
}

func (v NullableXmlNs0MDNUserSetting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableXmlNs0MDNUserSetting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


