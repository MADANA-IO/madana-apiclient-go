/*
 * madana-api
 *
 * <h1>Using the madana-api</h1>        <p>This documentation contains a Quickstart Guide, relating client functionality and information about the available         endpoints and used datamodels.   </p>       <p> The madana-api and its implementations are still in heavy development. This means that there may be problems in our protocols, or there may be mistakes in our implementations. We take security vulnerabilities very seriously. If you discover a security issue, please bring it to our attention right away! If you find a vulnerability that may affect live deployments -- for example, by exposing a remote execution exploit -- please send your report privately to info@madana.io. Please DO NOT file a public issue. If the issue is a protocol weakness that cannot be immediately exploited or something not yet deployed, just discuss it openly   </p>   <br>   <p> Note: Not all functionality might be acessible without having accquired and api-license token. For more information visit <a href=\"https://www.madana.io\">www.madana.io</a> </p>       <br>
 *
 * API version: 0.5.0-master.37
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package madanaapiclient

import (
	"encoding/json"
)

// JsonMDNAUserObject 
type JsonMDNAUserObject struct {
	// 
	Created *string `json:"created,omitempty"`
	// 
	LastActive *string `json:"lastActive,omitempty"`
	// 
	Image *string `json:"image,omitempty"`
	// 
	UserName *string `json:"userName,omitempty"`
	// 
	Activated *string `json:"activated,omitempty"`
}

// NewJsonMDNAUserObject instantiates a new JsonMDNAUserObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJsonMDNAUserObject() *JsonMDNAUserObject {
	this := JsonMDNAUserObject{}
	return &this
}

// NewJsonMDNAUserObjectWithDefaults instantiates a new JsonMDNAUserObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJsonMDNAUserObjectWithDefaults() *JsonMDNAUserObject {
	this := JsonMDNAUserObject{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *JsonMDNAUserObject) GetCreated() string {
	if o == nil || o.Created == nil {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonMDNAUserObject) GetCreatedOk() (*string, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *JsonMDNAUserObject) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *JsonMDNAUserObject) SetCreated(v string) {
	o.Created = &v
}

// GetLastActive returns the LastActive field value if set, zero value otherwise.
func (o *JsonMDNAUserObject) GetLastActive() string {
	if o == nil || o.LastActive == nil {
		var ret string
		return ret
	}
	return *o.LastActive
}

// GetLastActiveOk returns a tuple with the LastActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonMDNAUserObject) GetLastActiveOk() (*string, bool) {
	if o == nil || o.LastActive == nil {
		return nil, false
	}
	return o.LastActive, true
}

// HasLastActive returns a boolean if a field has been set.
func (o *JsonMDNAUserObject) HasLastActive() bool {
	if o != nil && o.LastActive != nil {
		return true
	}

	return false
}

// SetLastActive gets a reference to the given string and assigns it to the LastActive field.
func (o *JsonMDNAUserObject) SetLastActive(v string) {
	o.LastActive = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *JsonMDNAUserObject) GetImage() string {
	if o == nil || o.Image == nil {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonMDNAUserObject) GetImageOk() (*string, bool) {
	if o == nil || o.Image == nil {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *JsonMDNAUserObject) HasImage() bool {
	if o != nil && o.Image != nil {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *JsonMDNAUserObject) SetImage(v string) {
	o.Image = &v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *JsonMDNAUserObject) GetUserName() string {
	if o == nil || o.UserName == nil {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonMDNAUserObject) GetUserNameOk() (*string, bool) {
	if o == nil || o.UserName == nil {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *JsonMDNAUserObject) HasUserName() bool {
	if o != nil && o.UserName != nil {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *JsonMDNAUserObject) SetUserName(v string) {
	o.UserName = &v
}

// GetActivated returns the Activated field value if set, zero value otherwise.
func (o *JsonMDNAUserObject) GetActivated() string {
	if o == nil || o.Activated == nil {
		var ret string
		return ret
	}
	return *o.Activated
}

// GetActivatedOk returns a tuple with the Activated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonMDNAUserObject) GetActivatedOk() (*string, bool) {
	if o == nil || o.Activated == nil {
		return nil, false
	}
	return o.Activated, true
}

// HasActivated returns a boolean if a field has been set.
func (o *JsonMDNAUserObject) HasActivated() bool {
	if o != nil && o.Activated != nil {
		return true
	}

	return false
}

// SetActivated gets a reference to the given string and assigns it to the Activated field.
func (o *JsonMDNAUserObject) SetActivated(v string) {
	o.Activated = &v
}

func (o JsonMDNAUserObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.LastActive != nil {
		toSerialize["lastActive"] = o.LastActive
	}
	if o.Image != nil {
		toSerialize["image"] = o.Image
	}
	if o.UserName != nil {
		toSerialize["userName"] = o.UserName
	}
	if o.Activated != nil {
		toSerialize["activated"] = o.Activated
	}
	return json.Marshal(toSerialize)
}

type NullableJsonMDNAUserObject struct {
	value *JsonMDNAUserObject
	isSet bool
}

func (v NullableJsonMDNAUserObject) Get() *JsonMDNAUserObject {
	return v.value
}

func (v *NullableJsonMDNAUserObject) Set(val *JsonMDNAUserObject) {
	v.value = val
	v.isSet = true
}

func (v NullableJsonMDNAUserObject) IsSet() bool {
	return v.isSet
}

func (v *NullableJsonMDNAUserObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJsonMDNAUserObject(val *JsonMDNAUserObject) *NullableJsonMDNAUserObject {
	return &NullableJsonMDNAUserObject{value: val, isSet: true}
}

func (v NullableJsonMDNAUserObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJsonMDNAUserObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


