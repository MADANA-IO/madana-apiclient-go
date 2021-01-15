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

// JsonKubernetesEnclaveAllOf struct for JsonKubernetesEnclaveAllOf
type JsonKubernetesEnclaveAllOf struct {
	// 
	AttestationPort *int32 `json:"attestationPort,omitempty"`
	// 
	IsUsingInitContainer *bool `json:"isUsingInitContainer,omitempty"`
	EnclaveDeploymentEvents *JsonV1EventList `json:"enclaveDeploymentEvents,omitempty"`
	// 
	DebugInfo *string `json:"debugInfo,omitempty"`
	// 
	RemoteControlIP *string `json:"remoteControlIP,omitempty"`
	EnclavePodEvents *JsonV1EventList `json:"enclavePodEvents,omitempty"`
	EnclaveReplicaSetEvents *JsonV1EventList `json:"enclaveReplicaSetEvents,omitempty"`
	// 
	PodPhase *string `json:"podPhase,omitempty"`
	// 
	WireguardPort *int32 `json:"wireguardPort,omitempty"`
}

// NewJsonKubernetesEnclaveAllOf instantiates a new JsonKubernetesEnclaveAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJsonKubernetesEnclaveAllOf() *JsonKubernetesEnclaveAllOf {
	this := JsonKubernetesEnclaveAllOf{}
	return &this
}

// NewJsonKubernetesEnclaveAllOfWithDefaults instantiates a new JsonKubernetesEnclaveAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJsonKubernetesEnclaveAllOfWithDefaults() *JsonKubernetesEnclaveAllOf {
	this := JsonKubernetesEnclaveAllOf{}
	return &this
}

// GetAttestationPort returns the AttestationPort field value if set, zero value otherwise.
func (o *JsonKubernetesEnclaveAllOf) GetAttestationPort() int32 {
	if o == nil || o.AttestationPort == nil {
		var ret int32
		return ret
	}
	return *o.AttestationPort
}

// GetAttestationPortOk returns a tuple with the AttestationPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonKubernetesEnclaveAllOf) GetAttestationPortOk() (*int32, bool) {
	if o == nil || o.AttestationPort == nil {
		return nil, false
	}
	return o.AttestationPort, true
}

// HasAttestationPort returns a boolean if a field has been set.
func (o *JsonKubernetesEnclaveAllOf) HasAttestationPort() bool {
	if o != nil && o.AttestationPort != nil {
		return true
	}

	return false
}

// SetAttestationPort gets a reference to the given int32 and assigns it to the AttestationPort field.
func (o *JsonKubernetesEnclaveAllOf) SetAttestationPort(v int32) {
	o.AttestationPort = &v
}

// GetIsUsingInitContainer returns the IsUsingInitContainer field value if set, zero value otherwise.
func (o *JsonKubernetesEnclaveAllOf) GetIsUsingInitContainer() bool {
	if o == nil || o.IsUsingInitContainer == nil {
		var ret bool
		return ret
	}
	return *o.IsUsingInitContainer
}

// GetIsUsingInitContainerOk returns a tuple with the IsUsingInitContainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonKubernetesEnclaveAllOf) GetIsUsingInitContainerOk() (*bool, bool) {
	if o == nil || o.IsUsingInitContainer == nil {
		return nil, false
	}
	return o.IsUsingInitContainer, true
}

// HasIsUsingInitContainer returns a boolean if a field has been set.
func (o *JsonKubernetesEnclaveAllOf) HasIsUsingInitContainer() bool {
	if o != nil && o.IsUsingInitContainer != nil {
		return true
	}

	return false
}

// SetIsUsingInitContainer gets a reference to the given bool and assigns it to the IsUsingInitContainer field.
func (o *JsonKubernetesEnclaveAllOf) SetIsUsingInitContainer(v bool) {
	o.IsUsingInitContainer = &v
}

// GetEnclaveDeploymentEvents returns the EnclaveDeploymentEvents field value if set, zero value otherwise.
func (o *JsonKubernetesEnclaveAllOf) GetEnclaveDeploymentEvents() JsonV1EventList {
	if o == nil || o.EnclaveDeploymentEvents == nil {
		var ret JsonV1EventList
		return ret
	}
	return *o.EnclaveDeploymentEvents
}

// GetEnclaveDeploymentEventsOk returns a tuple with the EnclaveDeploymentEvents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonKubernetesEnclaveAllOf) GetEnclaveDeploymentEventsOk() (*JsonV1EventList, bool) {
	if o == nil || o.EnclaveDeploymentEvents == nil {
		return nil, false
	}
	return o.EnclaveDeploymentEvents, true
}

// HasEnclaveDeploymentEvents returns a boolean if a field has been set.
func (o *JsonKubernetesEnclaveAllOf) HasEnclaveDeploymentEvents() bool {
	if o != nil && o.EnclaveDeploymentEvents != nil {
		return true
	}

	return false
}

// SetEnclaveDeploymentEvents gets a reference to the given JsonV1EventList and assigns it to the EnclaveDeploymentEvents field.
func (o *JsonKubernetesEnclaveAllOf) SetEnclaveDeploymentEvents(v JsonV1EventList) {
	o.EnclaveDeploymentEvents = &v
}

// GetDebugInfo returns the DebugInfo field value if set, zero value otherwise.
func (o *JsonKubernetesEnclaveAllOf) GetDebugInfo() string {
	if o == nil || o.DebugInfo == nil {
		var ret string
		return ret
	}
	return *o.DebugInfo
}

// GetDebugInfoOk returns a tuple with the DebugInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonKubernetesEnclaveAllOf) GetDebugInfoOk() (*string, bool) {
	if o == nil || o.DebugInfo == nil {
		return nil, false
	}
	return o.DebugInfo, true
}

// HasDebugInfo returns a boolean if a field has been set.
func (o *JsonKubernetesEnclaveAllOf) HasDebugInfo() bool {
	if o != nil && o.DebugInfo != nil {
		return true
	}

	return false
}

// SetDebugInfo gets a reference to the given string and assigns it to the DebugInfo field.
func (o *JsonKubernetesEnclaveAllOf) SetDebugInfo(v string) {
	o.DebugInfo = &v
}

// GetRemoteControlIP returns the RemoteControlIP field value if set, zero value otherwise.
func (o *JsonKubernetesEnclaveAllOf) GetRemoteControlIP() string {
	if o == nil || o.RemoteControlIP == nil {
		var ret string
		return ret
	}
	return *o.RemoteControlIP
}

// GetRemoteControlIPOk returns a tuple with the RemoteControlIP field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonKubernetesEnclaveAllOf) GetRemoteControlIPOk() (*string, bool) {
	if o == nil || o.RemoteControlIP == nil {
		return nil, false
	}
	return o.RemoteControlIP, true
}

// HasRemoteControlIP returns a boolean if a field has been set.
func (o *JsonKubernetesEnclaveAllOf) HasRemoteControlIP() bool {
	if o != nil && o.RemoteControlIP != nil {
		return true
	}

	return false
}

// SetRemoteControlIP gets a reference to the given string and assigns it to the RemoteControlIP field.
func (o *JsonKubernetesEnclaveAllOf) SetRemoteControlIP(v string) {
	o.RemoteControlIP = &v
}

// GetEnclavePodEvents returns the EnclavePodEvents field value if set, zero value otherwise.
func (o *JsonKubernetesEnclaveAllOf) GetEnclavePodEvents() JsonV1EventList {
	if o == nil || o.EnclavePodEvents == nil {
		var ret JsonV1EventList
		return ret
	}
	return *o.EnclavePodEvents
}

// GetEnclavePodEventsOk returns a tuple with the EnclavePodEvents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonKubernetesEnclaveAllOf) GetEnclavePodEventsOk() (*JsonV1EventList, bool) {
	if o == nil || o.EnclavePodEvents == nil {
		return nil, false
	}
	return o.EnclavePodEvents, true
}

// HasEnclavePodEvents returns a boolean if a field has been set.
func (o *JsonKubernetesEnclaveAllOf) HasEnclavePodEvents() bool {
	if o != nil && o.EnclavePodEvents != nil {
		return true
	}

	return false
}

// SetEnclavePodEvents gets a reference to the given JsonV1EventList and assigns it to the EnclavePodEvents field.
func (o *JsonKubernetesEnclaveAllOf) SetEnclavePodEvents(v JsonV1EventList) {
	o.EnclavePodEvents = &v
}

// GetEnclaveReplicaSetEvents returns the EnclaveReplicaSetEvents field value if set, zero value otherwise.
func (o *JsonKubernetesEnclaveAllOf) GetEnclaveReplicaSetEvents() JsonV1EventList {
	if o == nil || o.EnclaveReplicaSetEvents == nil {
		var ret JsonV1EventList
		return ret
	}
	return *o.EnclaveReplicaSetEvents
}

// GetEnclaveReplicaSetEventsOk returns a tuple with the EnclaveReplicaSetEvents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonKubernetesEnclaveAllOf) GetEnclaveReplicaSetEventsOk() (*JsonV1EventList, bool) {
	if o == nil || o.EnclaveReplicaSetEvents == nil {
		return nil, false
	}
	return o.EnclaveReplicaSetEvents, true
}

// HasEnclaveReplicaSetEvents returns a boolean if a field has been set.
func (o *JsonKubernetesEnclaveAllOf) HasEnclaveReplicaSetEvents() bool {
	if o != nil && o.EnclaveReplicaSetEvents != nil {
		return true
	}

	return false
}

// SetEnclaveReplicaSetEvents gets a reference to the given JsonV1EventList and assigns it to the EnclaveReplicaSetEvents field.
func (o *JsonKubernetesEnclaveAllOf) SetEnclaveReplicaSetEvents(v JsonV1EventList) {
	o.EnclaveReplicaSetEvents = &v
}

// GetPodPhase returns the PodPhase field value if set, zero value otherwise.
func (o *JsonKubernetesEnclaveAllOf) GetPodPhase() string {
	if o == nil || o.PodPhase == nil {
		var ret string
		return ret
	}
	return *o.PodPhase
}

// GetPodPhaseOk returns a tuple with the PodPhase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonKubernetesEnclaveAllOf) GetPodPhaseOk() (*string, bool) {
	if o == nil || o.PodPhase == nil {
		return nil, false
	}
	return o.PodPhase, true
}

// HasPodPhase returns a boolean if a field has been set.
func (o *JsonKubernetesEnclaveAllOf) HasPodPhase() bool {
	if o != nil && o.PodPhase != nil {
		return true
	}

	return false
}

// SetPodPhase gets a reference to the given string and assigns it to the PodPhase field.
func (o *JsonKubernetesEnclaveAllOf) SetPodPhase(v string) {
	o.PodPhase = &v
}

// GetWireguardPort returns the WireguardPort field value if set, zero value otherwise.
func (o *JsonKubernetesEnclaveAllOf) GetWireguardPort() int32 {
	if o == nil || o.WireguardPort == nil {
		var ret int32
		return ret
	}
	return *o.WireguardPort
}

// GetWireguardPortOk returns a tuple with the WireguardPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonKubernetesEnclaveAllOf) GetWireguardPortOk() (*int32, bool) {
	if o == nil || o.WireguardPort == nil {
		return nil, false
	}
	return o.WireguardPort, true
}

// HasWireguardPort returns a boolean if a field has been set.
func (o *JsonKubernetesEnclaveAllOf) HasWireguardPort() bool {
	if o != nil && o.WireguardPort != nil {
		return true
	}

	return false
}

// SetWireguardPort gets a reference to the given int32 and assigns it to the WireguardPort field.
func (o *JsonKubernetesEnclaveAllOf) SetWireguardPort(v int32) {
	o.WireguardPort = &v
}

func (o JsonKubernetesEnclaveAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AttestationPort != nil {
		toSerialize["attestationPort"] = o.AttestationPort
	}
	if o.IsUsingInitContainer != nil {
		toSerialize["isUsingInitContainer"] = o.IsUsingInitContainer
	}
	if o.EnclaveDeploymentEvents != nil {
		toSerialize["enclaveDeploymentEvents"] = o.EnclaveDeploymentEvents
	}
	if o.DebugInfo != nil {
		toSerialize["debugInfo"] = o.DebugInfo
	}
	if o.RemoteControlIP != nil {
		toSerialize["remoteControlIP"] = o.RemoteControlIP
	}
	if o.EnclavePodEvents != nil {
		toSerialize["enclavePodEvents"] = o.EnclavePodEvents
	}
	if o.EnclaveReplicaSetEvents != nil {
		toSerialize["enclaveReplicaSetEvents"] = o.EnclaveReplicaSetEvents
	}
	if o.PodPhase != nil {
		toSerialize["podPhase"] = o.PodPhase
	}
	if o.WireguardPort != nil {
		toSerialize["wireguardPort"] = o.WireguardPort
	}
	return json.Marshal(toSerialize)
}

type NullableJsonKubernetesEnclaveAllOf struct {
	value *JsonKubernetesEnclaveAllOf
	isSet bool
}

func (v NullableJsonKubernetesEnclaveAllOf) Get() *JsonKubernetesEnclaveAllOf {
	return v.value
}

func (v *NullableJsonKubernetesEnclaveAllOf) Set(val *JsonKubernetesEnclaveAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableJsonKubernetesEnclaveAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableJsonKubernetesEnclaveAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJsonKubernetesEnclaveAllOf(val *JsonKubernetesEnclaveAllOf) *NullableJsonKubernetesEnclaveAllOf {
	return &NullableJsonKubernetesEnclaveAllOf{value: val, isSet: true}
}

func (v NullableJsonKubernetesEnclaveAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJsonKubernetesEnclaveAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


