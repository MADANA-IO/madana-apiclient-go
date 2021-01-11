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

// XmlNs0DiskConfig 
type XmlNs0DiskConfig struct {
	// 
	Disk *string `json:"disk,omitempty"`
	// 
	Readonly *bool `json:"readonly,omitempty"`
	// 
	Roothash *string `json:"roothash,omitempty"`
	// 
	RoothashOffset *int32 `json:"roothash_offset,omitempty"`
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

// GetDisk returns the Disk field value if set, zero value otherwise.
func (o *XmlNs0DiskConfig) GetDisk() string {
	if o == nil || o.Disk == nil {
		var ret string
		return ret
	}
	return *o.Disk
}

// GetDiskOk returns a tuple with the Disk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0DiskConfig) GetDiskOk() (*string, bool) {
	if o == nil || o.Disk == nil {
		return nil, false
	}
	return o.Disk, true
}

// HasDisk returns a boolean if a field has been set.
func (o *XmlNs0DiskConfig) HasDisk() bool {
	if o != nil && o.Disk != nil {
		return true
	}

	return false
}

// SetDisk gets a reference to the given string and assigns it to the Disk field.
func (o *XmlNs0DiskConfig) SetDisk(v string) {
	o.Disk = &v
}

// GetReadonly returns the Readonly field value if set, zero value otherwise.
func (o *XmlNs0DiskConfig) GetReadonly() bool {
	if o == nil || o.Readonly == nil {
		var ret bool
		return ret
	}
	return *o.Readonly
}

// GetReadonlyOk returns a tuple with the Readonly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0DiskConfig) GetReadonlyOk() (*bool, bool) {
	if o == nil || o.Readonly == nil {
		return nil, false
	}
	return o.Readonly, true
}

// HasReadonly returns a boolean if a field has been set.
func (o *XmlNs0DiskConfig) HasReadonly() bool {
	if o != nil && o.Readonly != nil {
		return true
	}

	return false
}

// SetReadonly gets a reference to the given bool and assigns it to the Readonly field.
func (o *XmlNs0DiskConfig) SetReadonly(v bool) {
	o.Readonly = &v
}

// GetRoothash returns the Roothash field value if set, zero value otherwise.
func (o *XmlNs0DiskConfig) GetRoothash() string {
	if o == nil || o.Roothash == nil {
		var ret string
		return ret
	}
	return *o.Roothash
}

// GetRoothashOk returns a tuple with the Roothash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0DiskConfig) GetRoothashOk() (*string, bool) {
	if o == nil || o.Roothash == nil {
		return nil, false
	}
	return o.Roothash, true
}

// HasRoothash returns a boolean if a field has been set.
func (o *XmlNs0DiskConfig) HasRoothash() bool {
	if o != nil && o.Roothash != nil {
		return true
	}

	return false
}

// SetRoothash gets a reference to the given string and assigns it to the Roothash field.
func (o *XmlNs0DiskConfig) SetRoothash(v string) {
	o.Roothash = &v
}

// GetRoothashOffset returns the RoothashOffset field value if set, zero value otherwise.
func (o *XmlNs0DiskConfig) GetRoothashOffset() int32 {
	if o == nil || o.RoothashOffset == nil {
		var ret int32
		return ret
	}
	return *o.RoothashOffset
}

// GetRoothashOffsetOk returns a tuple with the RoothashOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0DiskConfig) GetRoothashOffsetOk() (*int32, bool) {
	if o == nil || o.RoothashOffset == nil {
		return nil, false
	}
	return o.RoothashOffset, true
}

// HasRoothashOffset returns a boolean if a field has been set.
func (o *XmlNs0DiskConfig) HasRoothashOffset() bool {
	if o != nil && o.RoothashOffset != nil {
		return true
	}

	return false
}

// SetRoothashOffset gets a reference to the given int32 and assigns it to the RoothashOffset field.
func (o *XmlNs0DiskConfig) SetRoothashOffset(v int32) {
	o.RoothashOffset = &v
}

func (o XmlNs0DiskConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Disk != nil {
		toSerialize["disk"] = o.Disk
	}
	if o.Readonly != nil {
		toSerialize["readonly"] = o.Readonly
	}
	if o.Roothash != nil {
		toSerialize["roothash"] = o.Roothash
	}
	if o.RoothashOffset != nil {
		toSerialize["roothash_offset"] = o.RoothashOffset
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


