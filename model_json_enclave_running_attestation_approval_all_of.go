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

// JsonEnclaveRunningAttestationApprovalAllOf struct for JsonEnclaveRunningAttestationApprovalAllOf
type JsonEnclaveRunningAttestationApprovalAllOf struct {
	// 
	Approved *string `json:"approved,omitempty"`
}

// NewJsonEnclaveRunningAttestationApprovalAllOf instantiates a new JsonEnclaveRunningAttestationApprovalAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJsonEnclaveRunningAttestationApprovalAllOf() *JsonEnclaveRunningAttestationApprovalAllOf {
	this := JsonEnclaveRunningAttestationApprovalAllOf{}
	return &this
}

// NewJsonEnclaveRunningAttestationApprovalAllOfWithDefaults instantiates a new JsonEnclaveRunningAttestationApprovalAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJsonEnclaveRunningAttestationApprovalAllOfWithDefaults() *JsonEnclaveRunningAttestationApprovalAllOf {
	this := JsonEnclaveRunningAttestationApprovalAllOf{}
	return &this
}

// GetApproved returns the Approved field value if set, zero value otherwise.
func (o *JsonEnclaveRunningAttestationApprovalAllOf) GetApproved() string {
	if o == nil || o.Approved == nil {
		var ret string
		return ret
	}
	return *o.Approved
}

// GetApprovedOk returns a tuple with the Approved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonEnclaveRunningAttestationApprovalAllOf) GetApprovedOk() (*string, bool) {
	if o == nil || o.Approved == nil {
		return nil, false
	}
	return o.Approved, true
}

// HasApproved returns a boolean if a field has been set.
func (o *JsonEnclaveRunningAttestationApprovalAllOf) HasApproved() bool {
	if o != nil && o.Approved != nil {
		return true
	}

	return false
}

// SetApproved gets a reference to the given string and assigns it to the Approved field.
func (o *JsonEnclaveRunningAttestationApprovalAllOf) SetApproved(v string) {
	o.Approved = &v
}

func (o JsonEnclaveRunningAttestationApprovalAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Approved != nil {
		toSerialize["approved"] = o.Approved
	}
	return json.Marshal(toSerialize)
}

type NullableJsonEnclaveRunningAttestationApprovalAllOf struct {
	value *JsonEnclaveRunningAttestationApprovalAllOf
	isSet bool
}

func (v NullableJsonEnclaveRunningAttestationApprovalAllOf) Get() *JsonEnclaveRunningAttestationApprovalAllOf {
	return v.value
}

func (v *NullableJsonEnclaveRunningAttestationApprovalAllOf) Set(val *JsonEnclaveRunningAttestationApprovalAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableJsonEnclaveRunningAttestationApprovalAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableJsonEnclaveRunningAttestationApprovalAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJsonEnclaveRunningAttestationApprovalAllOf(val *JsonEnclaveRunningAttestationApprovalAllOf) *NullableJsonEnclaveRunningAttestationApprovalAllOf {
	return &NullableJsonEnclaveRunningAttestationApprovalAllOf{value: val, isSet: true}
}

func (v NullableJsonEnclaveRunningAttestationApprovalAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJsonEnclaveRunningAttestationApprovalAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


