/*
 * madana-api
 *
 * <h1>Using the madana-api</h1>        <p>This documentation contains a Quickstart Guide, relating client functionality and information about the available         endpoints and used datamodels.   </p>       <p> The madana-api and its implementations are still in heavy development. This means that there may be problems in our protocols, or there may be mistakes in our implementations. We take security vulnerabilities very seriously. If you discover a security issue, please bring it to our attention right away! If you find a vulnerability that may affect live deployments -- for example, by exposing a remote execution exploit -- please send your report privately to info@madana.io. Please DO NOT file a public issue. If the issue is a protocol weakness that cannot be immediately exploited or something not yet deployed, just discuss it openly   </p>   <br>   <p> Note: Not all functionality might be acessible without having accquired and api-license token. For more information visit <a href=\"https://www.madana.io\">www.madana.io</a> </p>       <br>
 *
 * API version: 0.5.0-master.56
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package madanaapiclient

import (
	"encoding/json"
)

// JsonEnclaveRunningAttestationApproval 
type JsonEnclaveRunningAttestationApproval struct {
	NodeInfo *JsonNodeInfo `json:"nodeInfo,omitempty"`
	EnclaveProcess *JsonEnclaveProcess `json:"enclaveProcess,omitempty"`
	// 
	Approved *string `json:"approved,omitempty"`
}

// NewJsonEnclaveRunningAttestationApproval instantiates a new JsonEnclaveRunningAttestationApproval object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJsonEnclaveRunningAttestationApproval() *JsonEnclaveRunningAttestationApproval {
	this := JsonEnclaveRunningAttestationApproval{}
	return &this
}

// NewJsonEnclaveRunningAttestationApprovalWithDefaults instantiates a new JsonEnclaveRunningAttestationApproval object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJsonEnclaveRunningAttestationApprovalWithDefaults() *JsonEnclaveRunningAttestationApproval {
	this := JsonEnclaveRunningAttestationApproval{}
	return &this
}

// GetNodeInfo returns the NodeInfo field value if set, zero value otherwise.
func (o *JsonEnclaveRunningAttestationApproval) GetNodeInfo() JsonNodeInfo {
	if o == nil || o.NodeInfo == nil {
		var ret JsonNodeInfo
		return ret
	}
	return *o.NodeInfo
}

// GetNodeInfoOk returns a tuple with the NodeInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonEnclaveRunningAttestationApproval) GetNodeInfoOk() (*JsonNodeInfo, bool) {
	if o == nil || o.NodeInfo == nil {
		return nil, false
	}
	return o.NodeInfo, true
}

// HasNodeInfo returns a boolean if a field has been set.
func (o *JsonEnclaveRunningAttestationApproval) HasNodeInfo() bool {
	if o != nil && o.NodeInfo != nil {
		return true
	}

	return false
}

// SetNodeInfo gets a reference to the given JsonNodeInfo and assigns it to the NodeInfo field.
func (o *JsonEnclaveRunningAttestationApproval) SetNodeInfo(v JsonNodeInfo) {
	o.NodeInfo = &v
}

// GetEnclaveProcess returns the EnclaveProcess field value if set, zero value otherwise.
func (o *JsonEnclaveRunningAttestationApproval) GetEnclaveProcess() JsonEnclaveProcess {
	if o == nil || o.EnclaveProcess == nil {
		var ret JsonEnclaveProcess
		return ret
	}
	return *o.EnclaveProcess
}

// GetEnclaveProcessOk returns a tuple with the EnclaveProcess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonEnclaveRunningAttestationApproval) GetEnclaveProcessOk() (*JsonEnclaveProcess, bool) {
	if o == nil || o.EnclaveProcess == nil {
		return nil, false
	}
	return o.EnclaveProcess, true
}

// HasEnclaveProcess returns a boolean if a field has been set.
func (o *JsonEnclaveRunningAttestationApproval) HasEnclaveProcess() bool {
	if o != nil && o.EnclaveProcess != nil {
		return true
	}

	return false
}

// SetEnclaveProcess gets a reference to the given JsonEnclaveProcess and assigns it to the EnclaveProcess field.
func (o *JsonEnclaveRunningAttestationApproval) SetEnclaveProcess(v JsonEnclaveProcess) {
	o.EnclaveProcess = &v
}

// GetApproved returns the Approved field value if set, zero value otherwise.
func (o *JsonEnclaveRunningAttestationApproval) GetApproved() string {
	if o == nil || o.Approved == nil {
		var ret string
		return ret
	}
	return *o.Approved
}

// GetApprovedOk returns a tuple with the Approved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonEnclaveRunningAttestationApproval) GetApprovedOk() (*string, bool) {
	if o == nil || o.Approved == nil {
		return nil, false
	}
	return o.Approved, true
}

// HasApproved returns a boolean if a field has been set.
func (o *JsonEnclaveRunningAttestationApproval) HasApproved() bool {
	if o != nil && o.Approved != nil {
		return true
	}

	return false
}

// SetApproved gets a reference to the given string and assigns it to the Approved field.
func (o *JsonEnclaveRunningAttestationApproval) SetApproved(v string) {
	o.Approved = &v
}

func (o JsonEnclaveRunningAttestationApproval) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NodeInfo != nil {
		toSerialize["nodeInfo"] = o.NodeInfo
	}
	if o.EnclaveProcess != nil {
		toSerialize["enclaveProcess"] = o.EnclaveProcess
	}
	if o.Approved != nil {
		toSerialize["approved"] = o.Approved
	}
	return json.Marshal(toSerialize)
}

type NullableJsonEnclaveRunningAttestationApproval struct {
	value *JsonEnclaveRunningAttestationApproval
	isSet bool
}

func (v NullableJsonEnclaveRunningAttestationApproval) Get() *JsonEnclaveRunningAttestationApproval {
	return v.value
}

func (v *NullableJsonEnclaveRunningAttestationApproval) Set(val *JsonEnclaveRunningAttestationApproval) {
	v.value = val
	v.isSet = true
}

func (v NullableJsonEnclaveRunningAttestationApproval) IsSet() bool {
	return v.isSet
}

func (v *NullableJsonEnclaveRunningAttestationApproval) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJsonEnclaveRunningAttestationApproval(val *JsonEnclaveRunningAttestationApproval) *NullableJsonEnclaveRunningAttestationApproval {
	return &NullableJsonEnclaveRunningAttestationApproval{value: val, isSet: true}
}

func (v NullableJsonEnclaveRunningAttestationApproval) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJsonEnclaveRunningAttestationApproval) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


