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

// JsonV1ObjectReference 
type JsonV1ObjectReference struct {
	// 
	Namespace *string `json:"namespace,omitempty"`
	// 
	Kind *string `json:"kind,omitempty"`
	// 
	Uid *string `json:"uid,omitempty"`
	// 
	ApiVersion *string `json:"apiVersion,omitempty"`
	// 
	ResourceVersion *string `json:"resourceVersion,omitempty"`
	// 
	FieldPath *string `json:"fieldPath,omitempty"`
	// 
	Name *string `json:"name,omitempty"`
}

// NewJsonV1ObjectReference instantiates a new JsonV1ObjectReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJsonV1ObjectReference() *JsonV1ObjectReference {
	this := JsonV1ObjectReference{}
	return &this
}

// NewJsonV1ObjectReferenceWithDefaults instantiates a new JsonV1ObjectReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJsonV1ObjectReferenceWithDefaults() *JsonV1ObjectReference {
	this := JsonV1ObjectReference{}
	return &this
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *JsonV1ObjectReference) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonV1ObjectReference) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *JsonV1ObjectReference) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *JsonV1ObjectReference) SetNamespace(v string) {
	o.Namespace = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *JsonV1ObjectReference) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonV1ObjectReference) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *JsonV1ObjectReference) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *JsonV1ObjectReference) SetKind(v string) {
	o.Kind = &v
}

// GetUid returns the Uid field value if set, zero value otherwise.
func (o *JsonV1ObjectReference) GetUid() string {
	if o == nil || o.Uid == nil {
		var ret string
		return ret
	}
	return *o.Uid
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonV1ObjectReference) GetUidOk() (*string, bool) {
	if o == nil || o.Uid == nil {
		return nil, false
	}
	return o.Uid, true
}

// HasUid returns a boolean if a field has been set.
func (o *JsonV1ObjectReference) HasUid() bool {
	if o != nil && o.Uid != nil {
		return true
	}

	return false
}

// SetUid gets a reference to the given string and assigns it to the Uid field.
func (o *JsonV1ObjectReference) SetUid(v string) {
	o.Uid = &v
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *JsonV1ObjectReference) GetApiVersion() string {
	if o == nil || o.ApiVersion == nil {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonV1ObjectReference) GetApiVersionOk() (*string, bool) {
	if o == nil || o.ApiVersion == nil {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *JsonV1ObjectReference) HasApiVersion() bool {
	if o != nil && o.ApiVersion != nil {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *JsonV1ObjectReference) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetResourceVersion returns the ResourceVersion field value if set, zero value otherwise.
func (o *JsonV1ObjectReference) GetResourceVersion() string {
	if o == nil || o.ResourceVersion == nil {
		var ret string
		return ret
	}
	return *o.ResourceVersion
}

// GetResourceVersionOk returns a tuple with the ResourceVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonV1ObjectReference) GetResourceVersionOk() (*string, bool) {
	if o == nil || o.ResourceVersion == nil {
		return nil, false
	}
	return o.ResourceVersion, true
}

// HasResourceVersion returns a boolean if a field has been set.
func (o *JsonV1ObjectReference) HasResourceVersion() bool {
	if o != nil && o.ResourceVersion != nil {
		return true
	}

	return false
}

// SetResourceVersion gets a reference to the given string and assigns it to the ResourceVersion field.
func (o *JsonV1ObjectReference) SetResourceVersion(v string) {
	o.ResourceVersion = &v
}

// GetFieldPath returns the FieldPath field value if set, zero value otherwise.
func (o *JsonV1ObjectReference) GetFieldPath() string {
	if o == nil || o.FieldPath == nil {
		var ret string
		return ret
	}
	return *o.FieldPath
}

// GetFieldPathOk returns a tuple with the FieldPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonV1ObjectReference) GetFieldPathOk() (*string, bool) {
	if o == nil || o.FieldPath == nil {
		return nil, false
	}
	return o.FieldPath, true
}

// HasFieldPath returns a boolean if a field has been set.
func (o *JsonV1ObjectReference) HasFieldPath() bool {
	if o != nil && o.FieldPath != nil {
		return true
	}

	return false
}

// SetFieldPath gets a reference to the given string and assigns it to the FieldPath field.
func (o *JsonV1ObjectReference) SetFieldPath(v string) {
	o.FieldPath = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *JsonV1ObjectReference) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonV1ObjectReference) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *JsonV1ObjectReference) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *JsonV1ObjectReference) SetName(v string) {
	o.Name = &v
}

func (o JsonV1ObjectReference) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.Uid != nil {
		toSerialize["uid"] = o.Uid
	}
	if o.ApiVersion != nil {
		toSerialize["apiVersion"] = o.ApiVersion
	}
	if o.ResourceVersion != nil {
		toSerialize["resourceVersion"] = o.ResourceVersion
	}
	if o.FieldPath != nil {
		toSerialize["fieldPath"] = o.FieldPath
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableJsonV1ObjectReference struct {
	value *JsonV1ObjectReference
	isSet bool
}

func (v NullableJsonV1ObjectReference) Get() *JsonV1ObjectReference {
	return v.value
}

func (v *NullableJsonV1ObjectReference) Set(val *JsonV1ObjectReference) {
	v.value = val
	v.isSet = true
}

func (v NullableJsonV1ObjectReference) IsSet() bool {
	return v.isSet
}

func (v *NullableJsonV1ObjectReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJsonV1ObjectReference(val *JsonV1ObjectReference) *NullableJsonV1ObjectReference {
	return &NullableJsonV1ObjectReference{value: val, isSet: true}
}

func (v NullableJsonV1ObjectReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJsonV1ObjectReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


