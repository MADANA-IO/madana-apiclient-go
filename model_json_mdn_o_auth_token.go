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

// JsonMDNOAuthToken 
type JsonMDNOAuthToken struct {
	// 
	Verifier *string `json:"verifier,omitempty"`
	// 
	Token *string `json:"token,omitempty"`
}

// NewJsonMDNOAuthToken instantiates a new JsonMDNOAuthToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJsonMDNOAuthToken() *JsonMDNOAuthToken {
	this := JsonMDNOAuthToken{}
	return &this
}

// NewJsonMDNOAuthTokenWithDefaults instantiates a new JsonMDNOAuthToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJsonMDNOAuthTokenWithDefaults() *JsonMDNOAuthToken {
	this := JsonMDNOAuthToken{}
	return &this
}

// GetVerifier returns the Verifier field value if set, zero value otherwise.
func (o *JsonMDNOAuthToken) GetVerifier() string {
	if o == nil || o.Verifier == nil {
		var ret string
		return ret
	}
	return *o.Verifier
}

// GetVerifierOk returns a tuple with the Verifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonMDNOAuthToken) GetVerifierOk() (*string, bool) {
	if o == nil || o.Verifier == nil {
		return nil, false
	}
	return o.Verifier, true
}

// HasVerifier returns a boolean if a field has been set.
func (o *JsonMDNOAuthToken) HasVerifier() bool {
	if o != nil && o.Verifier != nil {
		return true
	}

	return false
}

// SetVerifier gets a reference to the given string and assigns it to the Verifier field.
func (o *JsonMDNOAuthToken) SetVerifier(v string) {
	o.Verifier = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *JsonMDNOAuthToken) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonMDNOAuthToken) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *JsonMDNOAuthToken) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *JsonMDNOAuthToken) SetToken(v string) {
	o.Token = &v
}

func (o JsonMDNOAuthToken) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Verifier != nil {
		toSerialize["verifier"] = o.Verifier
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	return json.Marshal(toSerialize)
}

type NullableJsonMDNOAuthToken struct {
	value *JsonMDNOAuthToken
	isSet bool
}

func (v NullableJsonMDNOAuthToken) Get() *JsonMDNOAuthToken {
	return v.value
}

func (v *NullableJsonMDNOAuthToken) Set(val *JsonMDNOAuthToken) {
	v.value = val
	v.isSet = true
}

func (v NullableJsonMDNOAuthToken) IsSet() bool {
	return v.isSet
}

func (v *NullableJsonMDNOAuthToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJsonMDNOAuthToken(val *JsonMDNOAuthToken) *NullableJsonMDNOAuthToken {
	return &NullableJsonMDNOAuthToken{value: val, isSet: true}
}

func (v NullableJsonMDNOAuthToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJsonMDNOAuthToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


