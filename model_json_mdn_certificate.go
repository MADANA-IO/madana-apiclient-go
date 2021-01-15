/*
 * madana-api
 *
 * <h1>Using the madana-api</h1>        <p>This documentation contains a Quickstart Guide, relating client functionality and information about the available         endpoints and used datamodels.   </p>       <p> The madana-api and its implementations are still in heavy development. This means that there may be problems in our protocols, or there may be mistakes in our implementations. We take security vulnerabilities very seriously. If you discover a security issue, please bring it to our attention right away! If you find a vulnerability that may affect live deployments -- for example, by exposing a remote execution exploit -- please send your report privately to info@madana.io. Please DO NOT file a public issue. If the issue is a protocol weakness that cannot be immediately exploited or something not yet deployed, just discuss it openly   </p>   <br>   <p> Note: Not all functionality might be acessible without having accquired and api-license token. For more information visit <a href=\"https://www.madana.io\">www.madana.io</a> </p>       <br>
 *
 * API version: 0.5.0-master.31
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package madanaapiclient

import (
	"encoding/json"
)

// JsonMDNCertificate 
type JsonMDNCertificate struct {
	// 
	Pem *string `json:"pem,omitempty"`
}

// NewJsonMDNCertificate instantiates a new JsonMDNCertificate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJsonMDNCertificate() *JsonMDNCertificate {
	this := JsonMDNCertificate{}
	return &this
}

// NewJsonMDNCertificateWithDefaults instantiates a new JsonMDNCertificate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJsonMDNCertificateWithDefaults() *JsonMDNCertificate {
	this := JsonMDNCertificate{}
	return &this
}

// GetPem returns the Pem field value if set, zero value otherwise.
func (o *JsonMDNCertificate) GetPem() string {
	if o == nil || o.Pem == nil {
		var ret string
		return ret
	}
	return *o.Pem
}

// GetPemOk returns a tuple with the Pem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonMDNCertificate) GetPemOk() (*string, bool) {
	if o == nil || o.Pem == nil {
		return nil, false
	}
	return o.Pem, true
}

// HasPem returns a boolean if a field has been set.
func (o *JsonMDNCertificate) HasPem() bool {
	if o != nil && o.Pem != nil {
		return true
	}

	return false
}

// SetPem gets a reference to the given string and assigns it to the Pem field.
func (o *JsonMDNCertificate) SetPem(v string) {
	o.Pem = &v
}

func (o JsonMDNCertificate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Pem != nil {
		toSerialize["pem"] = o.Pem
	}
	return json.Marshal(toSerialize)
}

type NullableJsonMDNCertificate struct {
	value *JsonMDNCertificate
	isSet bool
}

func (v NullableJsonMDNCertificate) Get() *JsonMDNCertificate {
	return v.value
}

func (v *NullableJsonMDNCertificate) Set(val *JsonMDNCertificate) {
	v.value = val
	v.isSet = true
}

func (v NullableJsonMDNCertificate) IsSet() bool {
	return v.isSet
}

func (v *NullableJsonMDNCertificate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJsonMDNCertificate(val *JsonMDNCertificate) *NullableJsonMDNCertificate {
	return &NullableJsonMDNCertificate{value: val, isSet: true}
}

func (v NullableJsonMDNCertificate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJsonMDNCertificate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


