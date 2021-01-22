/*
 * madana-api
 *
 * <h1>Using the madana-api</h1>        <p>This documentation contains a Quickstart Guide, relating client functionality and information about the available         endpoints and used datamodels.   </p>       <p> The madana-api and its implementations are still in heavy development. This means that there may be problems in our protocols, or there may be mistakes in our implementations. We take security vulnerabilities very seriously. If you discover a security issue, please bring it to our attention right away! If you find a vulnerability that may affect live deployments -- for example, by exposing a remote execution exploit -- please send your report privately to info@madana.io. Please DO NOT file a public issue. If the issue is a protocol weakness that cannot be immediately exploited or something not yet deployed, just discuss it openly   </p>   <br>   <p> Note: Not all functionality might be acessible without having accquired and api-license token. For more information visit <a href=\"https://www.madana.io\">www.madana.io</a> </p>       <br>
 *
 * API version: 0.5.0-master.43
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package madanaapiclient

import (
	"encoding/json"
)

// JsonMDNSocialUserObject 
type JsonMDNSocialUserObject struct {
	// 
	Image *string `json:"image,omitempty"`
	// 
	Platform *string `json:"platform,omitempty"`
	// 
	Ident *string `json:"ident,omitempty"`
}

// NewJsonMDNSocialUserObject instantiates a new JsonMDNSocialUserObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJsonMDNSocialUserObject() *JsonMDNSocialUserObject {
	this := JsonMDNSocialUserObject{}
	return &this
}

// NewJsonMDNSocialUserObjectWithDefaults instantiates a new JsonMDNSocialUserObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJsonMDNSocialUserObjectWithDefaults() *JsonMDNSocialUserObject {
	this := JsonMDNSocialUserObject{}
	return &this
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *JsonMDNSocialUserObject) GetImage() string {
	if o == nil || o.Image == nil {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonMDNSocialUserObject) GetImageOk() (*string, bool) {
	if o == nil || o.Image == nil {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *JsonMDNSocialUserObject) HasImage() bool {
	if o != nil && o.Image != nil {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *JsonMDNSocialUserObject) SetImage(v string) {
	o.Image = &v
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
func (o *JsonMDNSocialUserObject) GetPlatform() string {
	if o == nil || o.Platform == nil {
		var ret string
		return ret
	}
	return *o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonMDNSocialUserObject) GetPlatformOk() (*string, bool) {
	if o == nil || o.Platform == nil {
		return nil, false
	}
	return o.Platform, true
}

// HasPlatform returns a boolean if a field has been set.
func (o *JsonMDNSocialUserObject) HasPlatform() bool {
	if o != nil && o.Platform != nil {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given string and assigns it to the Platform field.
func (o *JsonMDNSocialUserObject) SetPlatform(v string) {
	o.Platform = &v
}

// GetIdent returns the Ident field value if set, zero value otherwise.
func (o *JsonMDNSocialUserObject) GetIdent() string {
	if o == nil || o.Ident == nil {
		var ret string
		return ret
	}
	return *o.Ident
}

// GetIdentOk returns a tuple with the Ident field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonMDNSocialUserObject) GetIdentOk() (*string, bool) {
	if o == nil || o.Ident == nil {
		return nil, false
	}
	return o.Ident, true
}

// HasIdent returns a boolean if a field has been set.
func (o *JsonMDNSocialUserObject) HasIdent() bool {
	if o != nil && o.Ident != nil {
		return true
	}

	return false
}

// SetIdent gets a reference to the given string and assigns it to the Ident field.
func (o *JsonMDNSocialUserObject) SetIdent(v string) {
	o.Ident = &v
}

func (o JsonMDNSocialUserObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Image != nil {
		toSerialize["image"] = o.Image
	}
	if o.Platform != nil {
		toSerialize["platform"] = o.Platform
	}
	if o.Ident != nil {
		toSerialize["ident"] = o.Ident
	}
	return json.Marshal(toSerialize)
}

type NullableJsonMDNSocialUserObject struct {
	value *JsonMDNSocialUserObject
	isSet bool
}

func (v NullableJsonMDNSocialUserObject) Get() *JsonMDNSocialUserObject {
	return v.value
}

func (v *NullableJsonMDNSocialUserObject) Set(val *JsonMDNSocialUserObject) {
	v.value = val
	v.isSet = true
}

func (v NullableJsonMDNSocialUserObject) IsSet() bool {
	return v.isSet
}

func (v *NullableJsonMDNSocialUserObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJsonMDNSocialUserObject(val *JsonMDNSocialUserObject) *NullableJsonMDNSocialUserObject {
	return &NullableJsonMDNSocialUserObject{value: val, isSet: true}
}

func (v NullableJsonMDNSocialUserObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJsonMDNSocialUserObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


