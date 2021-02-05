/*
 * madana-api
 *
 * <h1>Using the madana-api</h1>        <p>This documentation contains a Quickstart Guide, relating client functionality and information about the available         endpoints and used datamodels.   </p>       <p> The madana-api and its implementations are still in heavy development. This means that there may be problems in our protocols, or there may be mistakes in our implementations. We take security vulnerabilities very seriously. If you discover a security issue, please bring it to our attention right away! If you find a vulnerability that may affect live deployments -- for example, by exposing a remote execution exploit -- please send your report privately to info@madana.io. Please DO NOT file a public issue. If the issue is a protocol weakness that cannot be immediately exploited or something not yet deployed, just discuss it openly   </p>   <br>   <p> Note: Not all functionality might be acessible without having accquired and api-license token. For more information visit <a href=\"https://www.madana.io\">www.madana.io</a> </p>       <br>
 *
 * API version: 0.5.0-master.49
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package madanaapiclient

import (
	"encoding/json"
)

// XmlNs0IPFSSystemInfoAllOf struct for XmlNs0IPFSSystemInfoAllOf
type XmlNs0IPFSSystemInfoAllOf struct {
	// 
	AgentVersion *string `json:"agentVersion,omitempty"`
	// 
	Id *string `json:"id,omitempty"`
	// 
	ProtocolVersion *string `json:"protocolVersion,omitempty"`
	// 
	PublicKey *string `json:"publicKey,omitempty"`
	// 
	SwarmConnection *string `json:"swarmConnection,omitempty"`
}

// NewXmlNs0IPFSSystemInfoAllOf instantiates a new XmlNs0IPFSSystemInfoAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewXmlNs0IPFSSystemInfoAllOf() *XmlNs0IPFSSystemInfoAllOf {
	this := XmlNs0IPFSSystemInfoAllOf{}
	return &this
}

// NewXmlNs0IPFSSystemInfoAllOfWithDefaults instantiates a new XmlNs0IPFSSystemInfoAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewXmlNs0IPFSSystemInfoAllOfWithDefaults() *XmlNs0IPFSSystemInfoAllOf {
	this := XmlNs0IPFSSystemInfoAllOf{}
	return &this
}

// GetAgentVersion returns the AgentVersion field value if set, zero value otherwise.
func (o *XmlNs0IPFSSystemInfoAllOf) GetAgentVersion() string {
	if o == nil || o.AgentVersion == nil {
		var ret string
		return ret
	}
	return *o.AgentVersion
}

// GetAgentVersionOk returns a tuple with the AgentVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0IPFSSystemInfoAllOf) GetAgentVersionOk() (*string, bool) {
	if o == nil || o.AgentVersion == nil {
		return nil, false
	}
	return o.AgentVersion, true
}

// HasAgentVersion returns a boolean if a field has been set.
func (o *XmlNs0IPFSSystemInfoAllOf) HasAgentVersion() bool {
	if o != nil && o.AgentVersion != nil {
		return true
	}

	return false
}

// SetAgentVersion gets a reference to the given string and assigns it to the AgentVersion field.
func (o *XmlNs0IPFSSystemInfoAllOf) SetAgentVersion(v string) {
	o.AgentVersion = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *XmlNs0IPFSSystemInfoAllOf) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0IPFSSystemInfoAllOf) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *XmlNs0IPFSSystemInfoAllOf) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *XmlNs0IPFSSystemInfoAllOf) SetId(v string) {
	o.Id = &v
}

// GetProtocolVersion returns the ProtocolVersion field value if set, zero value otherwise.
func (o *XmlNs0IPFSSystemInfoAllOf) GetProtocolVersion() string {
	if o == nil || o.ProtocolVersion == nil {
		var ret string
		return ret
	}
	return *o.ProtocolVersion
}

// GetProtocolVersionOk returns a tuple with the ProtocolVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0IPFSSystemInfoAllOf) GetProtocolVersionOk() (*string, bool) {
	if o == nil || o.ProtocolVersion == nil {
		return nil, false
	}
	return o.ProtocolVersion, true
}

// HasProtocolVersion returns a boolean if a field has been set.
func (o *XmlNs0IPFSSystemInfoAllOf) HasProtocolVersion() bool {
	if o != nil && o.ProtocolVersion != nil {
		return true
	}

	return false
}

// SetProtocolVersion gets a reference to the given string and assigns it to the ProtocolVersion field.
func (o *XmlNs0IPFSSystemInfoAllOf) SetProtocolVersion(v string) {
	o.ProtocolVersion = &v
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise.
func (o *XmlNs0IPFSSystemInfoAllOf) GetPublicKey() string {
	if o == nil || o.PublicKey == nil {
		var ret string
		return ret
	}
	return *o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0IPFSSystemInfoAllOf) GetPublicKeyOk() (*string, bool) {
	if o == nil || o.PublicKey == nil {
		return nil, false
	}
	return o.PublicKey, true
}

// HasPublicKey returns a boolean if a field has been set.
func (o *XmlNs0IPFSSystemInfoAllOf) HasPublicKey() bool {
	if o != nil && o.PublicKey != nil {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given string and assigns it to the PublicKey field.
func (o *XmlNs0IPFSSystemInfoAllOf) SetPublicKey(v string) {
	o.PublicKey = &v
}

// GetSwarmConnection returns the SwarmConnection field value if set, zero value otherwise.
func (o *XmlNs0IPFSSystemInfoAllOf) GetSwarmConnection() string {
	if o == nil || o.SwarmConnection == nil {
		var ret string
		return ret
	}
	return *o.SwarmConnection
}

// GetSwarmConnectionOk returns a tuple with the SwarmConnection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0IPFSSystemInfoAllOf) GetSwarmConnectionOk() (*string, bool) {
	if o == nil || o.SwarmConnection == nil {
		return nil, false
	}
	return o.SwarmConnection, true
}

// HasSwarmConnection returns a boolean if a field has been set.
func (o *XmlNs0IPFSSystemInfoAllOf) HasSwarmConnection() bool {
	if o != nil && o.SwarmConnection != nil {
		return true
	}

	return false
}

// SetSwarmConnection gets a reference to the given string and assigns it to the SwarmConnection field.
func (o *XmlNs0IPFSSystemInfoAllOf) SetSwarmConnection(v string) {
	o.SwarmConnection = &v
}

func (o XmlNs0IPFSSystemInfoAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AgentVersion != nil {
		toSerialize["agentVersion"] = o.AgentVersion
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ProtocolVersion != nil {
		toSerialize["protocolVersion"] = o.ProtocolVersion
	}
	if o.PublicKey != nil {
		toSerialize["publicKey"] = o.PublicKey
	}
	if o.SwarmConnection != nil {
		toSerialize["swarmConnection"] = o.SwarmConnection
	}
	return json.Marshal(toSerialize)
}

type NullableXmlNs0IPFSSystemInfoAllOf struct {
	value *XmlNs0IPFSSystemInfoAllOf
	isSet bool
}

func (v NullableXmlNs0IPFSSystemInfoAllOf) Get() *XmlNs0IPFSSystemInfoAllOf {
	return v.value
}

func (v *NullableXmlNs0IPFSSystemInfoAllOf) Set(val *XmlNs0IPFSSystemInfoAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableXmlNs0IPFSSystemInfoAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableXmlNs0IPFSSystemInfoAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableXmlNs0IPFSSystemInfoAllOf(val *XmlNs0IPFSSystemInfoAllOf) *NullableXmlNs0IPFSSystemInfoAllOf {
	return &NullableXmlNs0IPFSSystemInfoAllOf{value: val, isSet: true}
}

func (v NullableXmlNs0IPFSSystemInfoAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableXmlNs0IPFSSystemInfoAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


