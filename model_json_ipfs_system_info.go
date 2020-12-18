/*
 * madana-api
 *
 * <h1>Using the madana-api</h1>        <p>This documentation contains a Quickstart Guide, relating client functionality and information about the available         endpoints and used datamodels.   </p>       <p> The madana-api and its implementations are still in heavy development. This means that there may be problems in our protocols, or there may be mistakes in our implementations. We take security vulnerabilities very seriously. If you discover a security issue, please bring it to our attention right away! If you find a vulnerability that may affect live deployments -- for example, by exposing a remote execution exploit -- please send your report privately to info@madana.io. Please DO NOT file a public issue. If the issue is a protocol weakness that cannot be immediately exploited or something not yet deployed, just discuss it openly   </p>   <br>   <p> Note: Not all functionality might be acessible without having accquired and api-license token. For more information visit <a href=\"https://www.madana.io\">www.madana.io</a> </p>       <br>
 *
 * API version: 0.5.0-master.15
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package madanaapiclient

import (
	"encoding/json"
)

// JsonIPFSSystemInfo 
type JsonIPFSSystemInfo struct {
	// 
	Id *string `json:"id,omitempty"`
	// 
	PublicKey *string `json:"publicKey,omitempty"`
	// 
	SwarmConnection *string `json:"swarmConnection,omitempty"`
	// 
	ProtocolVersion *string `json:"protocolVersion,omitempty"`
	// 
	AgentVersion *string `json:"agentVersion,omitempty"`
}

// NewJsonIPFSSystemInfo instantiates a new JsonIPFSSystemInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJsonIPFSSystemInfo() *JsonIPFSSystemInfo {
	this := JsonIPFSSystemInfo{}
	return &this
}

// NewJsonIPFSSystemInfoWithDefaults instantiates a new JsonIPFSSystemInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJsonIPFSSystemInfoWithDefaults() *JsonIPFSSystemInfo {
	this := JsonIPFSSystemInfo{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *JsonIPFSSystemInfo) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonIPFSSystemInfo) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *JsonIPFSSystemInfo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *JsonIPFSSystemInfo) SetId(v string) {
	o.Id = &v
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise.
func (o *JsonIPFSSystemInfo) GetPublicKey() string {
	if o == nil || o.PublicKey == nil {
		var ret string
		return ret
	}
	return *o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonIPFSSystemInfo) GetPublicKeyOk() (*string, bool) {
	if o == nil || o.PublicKey == nil {
		return nil, false
	}
	return o.PublicKey, true
}

// HasPublicKey returns a boolean if a field has been set.
func (o *JsonIPFSSystemInfo) HasPublicKey() bool {
	if o != nil && o.PublicKey != nil {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given string and assigns it to the PublicKey field.
func (o *JsonIPFSSystemInfo) SetPublicKey(v string) {
	o.PublicKey = &v
}

// GetSwarmConnection returns the SwarmConnection field value if set, zero value otherwise.
func (o *JsonIPFSSystemInfo) GetSwarmConnection() string {
	if o == nil || o.SwarmConnection == nil {
		var ret string
		return ret
	}
	return *o.SwarmConnection
}

// GetSwarmConnectionOk returns a tuple with the SwarmConnection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonIPFSSystemInfo) GetSwarmConnectionOk() (*string, bool) {
	if o == nil || o.SwarmConnection == nil {
		return nil, false
	}
	return o.SwarmConnection, true
}

// HasSwarmConnection returns a boolean if a field has been set.
func (o *JsonIPFSSystemInfo) HasSwarmConnection() bool {
	if o != nil && o.SwarmConnection != nil {
		return true
	}

	return false
}

// SetSwarmConnection gets a reference to the given string and assigns it to the SwarmConnection field.
func (o *JsonIPFSSystemInfo) SetSwarmConnection(v string) {
	o.SwarmConnection = &v
}

// GetProtocolVersion returns the ProtocolVersion field value if set, zero value otherwise.
func (o *JsonIPFSSystemInfo) GetProtocolVersion() string {
	if o == nil || o.ProtocolVersion == nil {
		var ret string
		return ret
	}
	return *o.ProtocolVersion
}

// GetProtocolVersionOk returns a tuple with the ProtocolVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonIPFSSystemInfo) GetProtocolVersionOk() (*string, bool) {
	if o == nil || o.ProtocolVersion == nil {
		return nil, false
	}
	return o.ProtocolVersion, true
}

// HasProtocolVersion returns a boolean if a field has been set.
func (o *JsonIPFSSystemInfo) HasProtocolVersion() bool {
	if o != nil && o.ProtocolVersion != nil {
		return true
	}

	return false
}

// SetProtocolVersion gets a reference to the given string and assigns it to the ProtocolVersion field.
func (o *JsonIPFSSystemInfo) SetProtocolVersion(v string) {
	o.ProtocolVersion = &v
}

// GetAgentVersion returns the AgentVersion field value if set, zero value otherwise.
func (o *JsonIPFSSystemInfo) GetAgentVersion() string {
	if o == nil || o.AgentVersion == nil {
		var ret string
		return ret
	}
	return *o.AgentVersion
}

// GetAgentVersionOk returns a tuple with the AgentVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonIPFSSystemInfo) GetAgentVersionOk() (*string, bool) {
	if o == nil || o.AgentVersion == nil {
		return nil, false
	}
	return o.AgentVersion, true
}

// HasAgentVersion returns a boolean if a field has been set.
func (o *JsonIPFSSystemInfo) HasAgentVersion() bool {
	if o != nil && o.AgentVersion != nil {
		return true
	}

	return false
}

// SetAgentVersion gets a reference to the given string and assigns it to the AgentVersion field.
func (o *JsonIPFSSystemInfo) SetAgentVersion(v string) {
	o.AgentVersion = &v
}

func (o JsonIPFSSystemInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.PublicKey != nil {
		toSerialize["publicKey"] = o.PublicKey
	}
	if o.SwarmConnection != nil {
		toSerialize["swarmConnection"] = o.SwarmConnection
	}
	if o.ProtocolVersion != nil {
		toSerialize["protocolVersion"] = o.ProtocolVersion
	}
	if o.AgentVersion != nil {
		toSerialize["agentVersion"] = o.AgentVersion
	}
	return json.Marshal(toSerialize)
}

type NullableJsonIPFSSystemInfo struct {
	value *JsonIPFSSystemInfo
	isSet bool
}

func (v NullableJsonIPFSSystemInfo) Get() *JsonIPFSSystemInfo {
	return v.value
}

func (v *NullableJsonIPFSSystemInfo) Set(val *JsonIPFSSystemInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableJsonIPFSSystemInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableJsonIPFSSystemInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJsonIPFSSystemInfo(val *JsonIPFSSystemInfo) *NullableJsonIPFSSystemInfo {
	return &NullableJsonIPFSSystemInfo{value: val, isSet: true}
}

func (v NullableJsonIPFSSystemInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJsonIPFSSystemInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


