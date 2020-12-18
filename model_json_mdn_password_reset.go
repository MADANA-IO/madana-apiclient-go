/*
 * madana-api
 *
 * <h1>Using the madana-api</h1>        <p>This documentation contains a Quickstart Guide, relating client functionality and information about the available         endpoints and used datamodels.   </p>       <p> The madana-api and its implementations are still in heavy development. This means that there may be problems in our protocols, or there may be mistakes in our implementations. We take security vulnerabilities very seriously. If you discover a security issue, please bring it to our attention right away! If you find a vulnerability that may affect live deployments -- for example, by exposing a remote execution exploit -- please send your report privately to info@madana.io. Please DO NOT file a public issue. If the issue is a protocol weakness that cannot be immediately exploited or something not yet deployed, just discuss it openly   </p>   <br>   <p> Note: Not all functionality might be acessible without having accquired and api-license token. For more information visit <a href=\"https://www.madana.io\">www.madana.io</a> </p>       <br>
 *
 * API version: 0.5.0-master.16
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package madanaapiclient

import (
	"encoding/json"
)

// JsonMDNPasswordReset 
type JsonMDNPasswordReset struct {
	// 
	Password *string `json:"password,omitempty"`
	// 
	Token *string `json:"token,omitempty"`
	// 
	Mail *string `json:"mail,omitempty"`
}

// NewJsonMDNPasswordReset instantiates a new JsonMDNPasswordReset object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJsonMDNPasswordReset() *JsonMDNPasswordReset {
	this := JsonMDNPasswordReset{}
	return &this
}

// NewJsonMDNPasswordResetWithDefaults instantiates a new JsonMDNPasswordReset object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJsonMDNPasswordResetWithDefaults() *JsonMDNPasswordReset {
	this := JsonMDNPasswordReset{}
	return &this
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *JsonMDNPasswordReset) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonMDNPasswordReset) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *JsonMDNPasswordReset) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *JsonMDNPasswordReset) SetPassword(v string) {
	o.Password = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *JsonMDNPasswordReset) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonMDNPasswordReset) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *JsonMDNPasswordReset) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *JsonMDNPasswordReset) SetToken(v string) {
	o.Token = &v
}

// GetMail returns the Mail field value if set, zero value otherwise.
func (o *JsonMDNPasswordReset) GetMail() string {
	if o == nil || o.Mail == nil {
		var ret string
		return ret
	}
	return *o.Mail
}

// GetMailOk returns a tuple with the Mail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonMDNPasswordReset) GetMailOk() (*string, bool) {
	if o == nil || o.Mail == nil {
		return nil, false
	}
	return o.Mail, true
}

// HasMail returns a boolean if a field has been set.
func (o *JsonMDNPasswordReset) HasMail() bool {
	if o != nil && o.Mail != nil {
		return true
	}

	return false
}

// SetMail gets a reference to the given string and assigns it to the Mail field.
func (o *JsonMDNPasswordReset) SetMail(v string) {
	o.Mail = &v
}

func (o JsonMDNPasswordReset) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.Mail != nil {
		toSerialize["mail"] = o.Mail
	}
	return json.Marshal(toSerialize)
}

type NullableJsonMDNPasswordReset struct {
	value *JsonMDNPasswordReset
	isSet bool
}

func (v NullableJsonMDNPasswordReset) Get() *JsonMDNPasswordReset {
	return v.value
}

func (v *NullableJsonMDNPasswordReset) Set(val *JsonMDNPasswordReset) {
	v.value = val
	v.isSet = true
}

func (v NullableJsonMDNPasswordReset) IsSet() bool {
	return v.isSet
}

func (v *NullableJsonMDNPasswordReset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJsonMDNPasswordReset(val *JsonMDNPasswordReset) *NullableJsonMDNPasswordReset {
	return &NullableJsonMDNPasswordReset{value: val, isSet: true}
}

func (v NullableJsonMDNPasswordReset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJsonMDNPasswordReset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


