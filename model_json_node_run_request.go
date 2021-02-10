/*
 * madana-api
 *
 * <h1>Using the madana-api</h1>        <p>This documentation contains a Quickstart Guide, relating client functionality and information about the available         endpoints and used datamodels.   </p>       <p> The madana-api and its implementations are still in heavy development. This means that there may be problems in our protocols, or there may be mistakes in our implementations. We take security vulnerabilities very seriously. If you discover a security issue, please bring it to our attention right away! If you find a vulnerability that may affect live deployments -- for example, by exposing a remote execution exploit -- please send your report privately to info@madana.io. Please DO NOT file a public issue. If the issue is a protocol weakness that cannot be immediately exploited or something not yet deployed, just discuss it openly   </p>   <br>   <p> Note: Not all functionality might be acessible without having accquired and api-license token. For more information visit <a href=\"https://www.madana.io\">www.madana.io</a> </p>       <br>
 *
 * API version: 0.5.0-master.55
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package madanaapiclient

import (
	"encoding/json"
)

// JsonNodeRunRequest 
type JsonNodeRunRequest struct {
	// 
	CpuCount *string `json:"cpuCount,omitempty"`
	// 
	Subdomain *string `json:"subdomain,omitempty"`
}

// NewJsonNodeRunRequest instantiates a new JsonNodeRunRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJsonNodeRunRequest() *JsonNodeRunRequest {
	this := JsonNodeRunRequest{}
	return &this
}

// NewJsonNodeRunRequestWithDefaults instantiates a new JsonNodeRunRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJsonNodeRunRequestWithDefaults() *JsonNodeRunRequest {
	this := JsonNodeRunRequest{}
	return &this
}

// GetCpuCount returns the CpuCount field value if set, zero value otherwise.
func (o *JsonNodeRunRequest) GetCpuCount() string {
	if o == nil || o.CpuCount == nil {
		var ret string
		return ret
	}
	return *o.CpuCount
}

// GetCpuCountOk returns a tuple with the CpuCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonNodeRunRequest) GetCpuCountOk() (*string, bool) {
	if o == nil || o.CpuCount == nil {
		return nil, false
	}
	return o.CpuCount, true
}

// HasCpuCount returns a boolean if a field has been set.
func (o *JsonNodeRunRequest) HasCpuCount() bool {
	if o != nil && o.CpuCount != nil {
		return true
	}

	return false
}

// SetCpuCount gets a reference to the given string and assigns it to the CpuCount field.
func (o *JsonNodeRunRequest) SetCpuCount(v string) {
	o.CpuCount = &v
}

// GetSubdomain returns the Subdomain field value if set, zero value otherwise.
func (o *JsonNodeRunRequest) GetSubdomain() string {
	if o == nil || o.Subdomain == nil {
		var ret string
		return ret
	}
	return *o.Subdomain
}

// GetSubdomainOk returns a tuple with the Subdomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonNodeRunRequest) GetSubdomainOk() (*string, bool) {
	if o == nil || o.Subdomain == nil {
		return nil, false
	}
	return o.Subdomain, true
}

// HasSubdomain returns a boolean if a field has been set.
func (o *JsonNodeRunRequest) HasSubdomain() bool {
	if o != nil && o.Subdomain != nil {
		return true
	}

	return false
}

// SetSubdomain gets a reference to the given string and assigns it to the Subdomain field.
func (o *JsonNodeRunRequest) SetSubdomain(v string) {
	o.Subdomain = &v
}

func (o JsonNodeRunRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CpuCount != nil {
		toSerialize["cpuCount"] = o.CpuCount
	}
	if o.Subdomain != nil {
		toSerialize["subdomain"] = o.Subdomain
	}
	return json.Marshal(toSerialize)
}

type NullableJsonNodeRunRequest struct {
	value *JsonNodeRunRequest
	isSet bool
}

func (v NullableJsonNodeRunRequest) Get() *JsonNodeRunRequest {
	return v.value
}

func (v *NullableJsonNodeRunRequest) Set(val *JsonNodeRunRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableJsonNodeRunRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableJsonNodeRunRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJsonNodeRunRequest(val *JsonNodeRunRequest) *NullableJsonNodeRunRequest {
	return &NullableJsonNodeRunRequest{value: val, isSet: true}
}

func (v NullableJsonNodeRunRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJsonNodeRunRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


