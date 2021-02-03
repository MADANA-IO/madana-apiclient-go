/*
 * madana-api
 *
 * <h1>Using the madana-api</h1>        <p>This documentation contains a Quickstart Guide, relating client functionality and information about the available         endpoints and used datamodels.   </p>       <p> The madana-api and its implementations are still in heavy development. This means that there may be problems in our protocols, or there may be mistakes in our implementations. We take security vulnerabilities very seriously. If you discover a security issue, please bring it to our attention right away! If you find a vulnerability that may affect live deployments -- for example, by exposing a remote execution exploit -- please send your report privately to info@madana.io. Please DO NOT file a public issue. If the issue is a protocol weakness that cannot be immediately exploited or something not yet deployed, just discuss it openly   </p>   <br>   <p> Note: Not all functionality might be acessible without having accquired and api-license token. For more information visit <a href=\"https://www.madana.io\">www.madana.io</a> </p>       <br>
 *
 * API version: 0.5.0-master.47
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package madanaapiclient

import (
	"encoding/json"
)

// JsonEnclaveRunRequest 
type JsonEnclaveRunRequest struct {
	// 
	EnvironmentUUID *string `json:"environmentUUID,omitempty"`
	// 
	Ports *[]JsonEnclavePort `json:"ports,omitempty"`
	// 
	WireguardPublicKey *string `json:"wireguardPublicKey,omitempty"`
	// 
	EnclaveExecutionType *string `json:"enclaveExecutionType,omitempty"`
	// 
	UsingDefaultRunConfig *bool `json:"usingDefaultRunConfig,omitempty"`
}

// NewJsonEnclaveRunRequest instantiates a new JsonEnclaveRunRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJsonEnclaveRunRequest() *JsonEnclaveRunRequest {
	this := JsonEnclaveRunRequest{}
	return &this
}

// NewJsonEnclaveRunRequestWithDefaults instantiates a new JsonEnclaveRunRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJsonEnclaveRunRequestWithDefaults() *JsonEnclaveRunRequest {
	this := JsonEnclaveRunRequest{}
	return &this
}

// GetEnvironmentUUID returns the EnvironmentUUID field value if set, zero value otherwise.
func (o *JsonEnclaveRunRequest) GetEnvironmentUUID() string {
	if o == nil || o.EnvironmentUUID == nil {
		var ret string
		return ret
	}
	return *o.EnvironmentUUID
}

// GetEnvironmentUUIDOk returns a tuple with the EnvironmentUUID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonEnclaveRunRequest) GetEnvironmentUUIDOk() (*string, bool) {
	if o == nil || o.EnvironmentUUID == nil {
		return nil, false
	}
	return o.EnvironmentUUID, true
}

// HasEnvironmentUUID returns a boolean if a field has been set.
func (o *JsonEnclaveRunRequest) HasEnvironmentUUID() bool {
	if o != nil && o.EnvironmentUUID != nil {
		return true
	}

	return false
}

// SetEnvironmentUUID gets a reference to the given string and assigns it to the EnvironmentUUID field.
func (o *JsonEnclaveRunRequest) SetEnvironmentUUID(v string) {
	o.EnvironmentUUID = &v
}

// GetPorts returns the Ports field value if set, zero value otherwise.
func (o *JsonEnclaveRunRequest) GetPorts() []JsonEnclavePort {
	if o == nil || o.Ports == nil {
		var ret []JsonEnclavePort
		return ret
	}
	return *o.Ports
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonEnclaveRunRequest) GetPortsOk() (*[]JsonEnclavePort, bool) {
	if o == nil || o.Ports == nil {
		return nil, false
	}
	return o.Ports, true
}

// HasPorts returns a boolean if a field has been set.
func (o *JsonEnclaveRunRequest) HasPorts() bool {
	if o != nil && o.Ports != nil {
		return true
	}

	return false
}

// SetPorts gets a reference to the given []JsonEnclavePort and assigns it to the Ports field.
func (o *JsonEnclaveRunRequest) SetPorts(v []JsonEnclavePort) {
	o.Ports = &v
}

// GetWireguardPublicKey returns the WireguardPublicKey field value if set, zero value otherwise.
func (o *JsonEnclaveRunRequest) GetWireguardPublicKey() string {
	if o == nil || o.WireguardPublicKey == nil {
		var ret string
		return ret
	}
	return *o.WireguardPublicKey
}

// GetWireguardPublicKeyOk returns a tuple with the WireguardPublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonEnclaveRunRequest) GetWireguardPublicKeyOk() (*string, bool) {
	if o == nil || o.WireguardPublicKey == nil {
		return nil, false
	}
	return o.WireguardPublicKey, true
}

// HasWireguardPublicKey returns a boolean if a field has been set.
func (o *JsonEnclaveRunRequest) HasWireguardPublicKey() bool {
	if o != nil && o.WireguardPublicKey != nil {
		return true
	}

	return false
}

// SetWireguardPublicKey gets a reference to the given string and assigns it to the WireguardPublicKey field.
func (o *JsonEnclaveRunRequest) SetWireguardPublicKey(v string) {
	o.WireguardPublicKey = &v
}

// GetEnclaveExecutionType returns the EnclaveExecutionType field value if set, zero value otherwise.
func (o *JsonEnclaveRunRequest) GetEnclaveExecutionType() string {
	if o == nil || o.EnclaveExecutionType == nil {
		var ret string
		return ret
	}
	return *o.EnclaveExecutionType
}

// GetEnclaveExecutionTypeOk returns a tuple with the EnclaveExecutionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonEnclaveRunRequest) GetEnclaveExecutionTypeOk() (*string, bool) {
	if o == nil || o.EnclaveExecutionType == nil {
		return nil, false
	}
	return o.EnclaveExecutionType, true
}

// HasEnclaveExecutionType returns a boolean if a field has been set.
func (o *JsonEnclaveRunRequest) HasEnclaveExecutionType() bool {
	if o != nil && o.EnclaveExecutionType != nil {
		return true
	}

	return false
}

// SetEnclaveExecutionType gets a reference to the given string and assigns it to the EnclaveExecutionType field.
func (o *JsonEnclaveRunRequest) SetEnclaveExecutionType(v string) {
	o.EnclaveExecutionType = &v
}

// GetUsingDefaultRunConfig returns the UsingDefaultRunConfig field value if set, zero value otherwise.
func (o *JsonEnclaveRunRequest) GetUsingDefaultRunConfig() bool {
	if o == nil || o.UsingDefaultRunConfig == nil {
		var ret bool
		return ret
	}
	return *o.UsingDefaultRunConfig
}

// GetUsingDefaultRunConfigOk returns a tuple with the UsingDefaultRunConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonEnclaveRunRequest) GetUsingDefaultRunConfigOk() (*bool, bool) {
	if o == nil || o.UsingDefaultRunConfig == nil {
		return nil, false
	}
	return o.UsingDefaultRunConfig, true
}

// HasUsingDefaultRunConfig returns a boolean if a field has been set.
func (o *JsonEnclaveRunRequest) HasUsingDefaultRunConfig() bool {
	if o != nil && o.UsingDefaultRunConfig != nil {
		return true
	}

	return false
}

// SetUsingDefaultRunConfig gets a reference to the given bool and assigns it to the UsingDefaultRunConfig field.
func (o *JsonEnclaveRunRequest) SetUsingDefaultRunConfig(v bool) {
	o.UsingDefaultRunConfig = &v
}

func (o JsonEnclaveRunRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EnvironmentUUID != nil {
		toSerialize["environmentUUID"] = o.EnvironmentUUID
	}
	if o.Ports != nil {
		toSerialize["ports"] = o.Ports
	}
	if o.WireguardPublicKey != nil {
		toSerialize["wireguardPublicKey"] = o.WireguardPublicKey
	}
	if o.EnclaveExecutionType != nil {
		toSerialize["enclaveExecutionType"] = o.EnclaveExecutionType
	}
	if o.UsingDefaultRunConfig != nil {
		toSerialize["usingDefaultRunConfig"] = o.UsingDefaultRunConfig
	}
	return json.Marshal(toSerialize)
}

type NullableJsonEnclaveRunRequest struct {
	value *JsonEnclaveRunRequest
	isSet bool
}

func (v NullableJsonEnclaveRunRequest) Get() *JsonEnclaveRunRequest {
	return v.value
}

func (v *NullableJsonEnclaveRunRequest) Set(val *JsonEnclaveRunRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableJsonEnclaveRunRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableJsonEnclaveRunRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJsonEnclaveRunRequest(val *JsonEnclaveRunRequest) *NullableJsonEnclaveRunRequest {
	return &NullableJsonEnclaveRunRequest{value: val, isSet: true}
}

func (v NullableJsonEnclaveRunRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJsonEnclaveRunRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


