/*
 * madana-api
 *
 * <h1>Using the madana-api</h1>        <p>This documentation contains a Quickstart Guide, relating client functionality and information about the available         endpoints and used datamodels.   </p>       <p> The madana-api and its implementations are still in heavy development. This means that there may be problems in our protocols, or there may be mistakes in our implementations. We take security vulnerabilities very seriously. If you discover a security issue, please bring it to our attention right away! If you find a vulnerability that may affect live deployments -- for example, by exposing a remote execution exploit -- please send your report privately to info@madana.io. Please DO NOT file a public issue. If the issue is a protocol weakness that cannot be immediately exploited or something not yet deployed, just discuss it openly   </p>   <br>   <p> Note: Not all functionality might be acessible without having accquired and api-license token. For more information visit <a href=\"https://www.madana.io\">www.madana.io</a> </p>       <br>
 *
 * API version: 0.5.0-master.45
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package madanaapiclient

import (
	"encoding/json"
)

// JsonMDNData 
type JsonMDNData struct {
	// 
	Data *string `json:"data,omitempty"`
}

// NewJsonMDNData instantiates a new JsonMDNData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJsonMDNData() *JsonMDNData {
	this := JsonMDNData{}
	return &this
}

// NewJsonMDNDataWithDefaults instantiates a new JsonMDNData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJsonMDNDataWithDefaults() *JsonMDNData {
	this := JsonMDNData{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *JsonMDNData) GetData() string {
	if o == nil || o.Data == nil {
		var ret string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonMDNData) GetDataOk() (*string, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *JsonMDNData) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given string and assigns it to the Data field.
func (o *JsonMDNData) SetData(v string) {
	o.Data = &v
}

func (o JsonMDNData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableJsonMDNData struct {
	value *JsonMDNData
	isSet bool
}

func (v NullableJsonMDNData) Get() *JsonMDNData {
	return v.value
}

func (v *NullableJsonMDNData) Set(val *JsonMDNData) {
	v.value = val
	v.isSet = true
}

func (v NullableJsonMDNData) IsSet() bool {
	return v.isSet
}

func (v *NullableJsonMDNData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJsonMDNData(val *JsonMDNData) *NullableJsonMDNData {
	return &NullableJsonMDNData{value: val, isSet: true}
}

func (v NullableJsonMDNData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJsonMDNData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


