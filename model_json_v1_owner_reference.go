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

// JsonV1OwnerReference 
type JsonV1OwnerReference struct {
	// 
	Uid *string `json:"uid,omitempty"`
	// 
	Kind *string `json:"kind,omitempty"`
	// 
	Name *string `json:"name,omitempty"`
	// 
	BlockOwnerDeletion *bool `json:"blockOwnerDeletion,omitempty"`
	// 
	ApiVersion *string `json:"apiVersion,omitempty"`
	// 
	Controller *bool `json:"controller,omitempty"`
}

// NewJsonV1OwnerReference instantiates a new JsonV1OwnerReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJsonV1OwnerReference() *JsonV1OwnerReference {
	this := JsonV1OwnerReference{}
	return &this
}

// NewJsonV1OwnerReferenceWithDefaults instantiates a new JsonV1OwnerReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJsonV1OwnerReferenceWithDefaults() *JsonV1OwnerReference {
	this := JsonV1OwnerReference{}
	return &this
}

// GetUid returns the Uid field value if set, zero value otherwise.
func (o *JsonV1OwnerReference) GetUid() string {
	if o == nil || o.Uid == nil {
		var ret string
		return ret
	}
	return *o.Uid
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonV1OwnerReference) GetUidOk() (*string, bool) {
	if o == nil || o.Uid == nil {
		return nil, false
	}
	return o.Uid, true
}

// HasUid returns a boolean if a field has been set.
func (o *JsonV1OwnerReference) HasUid() bool {
	if o != nil && o.Uid != nil {
		return true
	}

	return false
}

// SetUid gets a reference to the given string and assigns it to the Uid field.
func (o *JsonV1OwnerReference) SetUid(v string) {
	o.Uid = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *JsonV1OwnerReference) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonV1OwnerReference) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *JsonV1OwnerReference) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *JsonV1OwnerReference) SetKind(v string) {
	o.Kind = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *JsonV1OwnerReference) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonV1OwnerReference) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *JsonV1OwnerReference) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *JsonV1OwnerReference) SetName(v string) {
	o.Name = &v
}

// GetBlockOwnerDeletion returns the BlockOwnerDeletion field value if set, zero value otherwise.
func (o *JsonV1OwnerReference) GetBlockOwnerDeletion() bool {
	if o == nil || o.BlockOwnerDeletion == nil {
		var ret bool
		return ret
	}
	return *o.BlockOwnerDeletion
}

// GetBlockOwnerDeletionOk returns a tuple with the BlockOwnerDeletion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonV1OwnerReference) GetBlockOwnerDeletionOk() (*bool, bool) {
	if o == nil || o.BlockOwnerDeletion == nil {
		return nil, false
	}
	return o.BlockOwnerDeletion, true
}

// HasBlockOwnerDeletion returns a boolean if a field has been set.
func (o *JsonV1OwnerReference) HasBlockOwnerDeletion() bool {
	if o != nil && o.BlockOwnerDeletion != nil {
		return true
	}

	return false
}

// SetBlockOwnerDeletion gets a reference to the given bool and assigns it to the BlockOwnerDeletion field.
func (o *JsonV1OwnerReference) SetBlockOwnerDeletion(v bool) {
	o.BlockOwnerDeletion = &v
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *JsonV1OwnerReference) GetApiVersion() string {
	if o == nil || o.ApiVersion == nil {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonV1OwnerReference) GetApiVersionOk() (*string, bool) {
	if o == nil || o.ApiVersion == nil {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *JsonV1OwnerReference) HasApiVersion() bool {
	if o != nil && o.ApiVersion != nil {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *JsonV1OwnerReference) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetController returns the Controller field value if set, zero value otherwise.
func (o *JsonV1OwnerReference) GetController() bool {
	if o == nil || o.Controller == nil {
		var ret bool
		return ret
	}
	return *o.Controller
}

// GetControllerOk returns a tuple with the Controller field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonV1OwnerReference) GetControllerOk() (*bool, bool) {
	if o == nil || o.Controller == nil {
		return nil, false
	}
	return o.Controller, true
}

// HasController returns a boolean if a field has been set.
func (o *JsonV1OwnerReference) HasController() bool {
	if o != nil && o.Controller != nil {
		return true
	}

	return false
}

// SetController gets a reference to the given bool and assigns it to the Controller field.
func (o *JsonV1OwnerReference) SetController(v bool) {
	o.Controller = &v
}

func (o JsonV1OwnerReference) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Uid != nil {
		toSerialize["uid"] = o.Uid
	}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.BlockOwnerDeletion != nil {
		toSerialize["blockOwnerDeletion"] = o.BlockOwnerDeletion
	}
	if o.ApiVersion != nil {
		toSerialize["apiVersion"] = o.ApiVersion
	}
	if o.Controller != nil {
		toSerialize["controller"] = o.Controller
	}
	return json.Marshal(toSerialize)
}

type NullableJsonV1OwnerReference struct {
	value *JsonV1OwnerReference
	isSet bool
}

func (v NullableJsonV1OwnerReference) Get() *JsonV1OwnerReference {
	return v.value
}

func (v *NullableJsonV1OwnerReference) Set(val *JsonV1OwnerReference) {
	v.value = val
	v.isSet = true
}

func (v NullableJsonV1OwnerReference) IsSet() bool {
	return v.isSet
}

func (v *NullableJsonV1OwnerReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJsonV1OwnerReference(val *JsonV1OwnerReference) *NullableJsonV1OwnerReference {
	return &NullableJsonV1OwnerReference{value: val, isSet: true}
}

func (v NullableJsonV1OwnerReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJsonV1OwnerReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


