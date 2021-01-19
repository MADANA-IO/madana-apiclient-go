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

// JsonMDNUserAllOf struct for JsonMDNUserAllOf
type JsonMDNUserAllOf struct {
	// 
	LastName *string `json:"lastName,omitempty"`
	// 
	Mail *string `json:"mail,omitempty"`
	// 
	SocialAccounts *[]JsonMDNSocialUserObject `json:"socialAccounts,omitempty"`
	// 
	FirstName *string `json:"firstName,omitempty"`
	Credentials *JsonMDNUserCredentials `json:"credentials,omitempty"`
	// 
	Guid *string `json:"guid,omitempty"`
	// 
	Settings *[]JsonMDNUserSetting `json:"settings,omitempty"`
}

// NewJsonMDNUserAllOf instantiates a new JsonMDNUserAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJsonMDNUserAllOf() *JsonMDNUserAllOf {
	this := JsonMDNUserAllOf{}
	return &this
}

// NewJsonMDNUserAllOfWithDefaults instantiates a new JsonMDNUserAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJsonMDNUserAllOfWithDefaults() *JsonMDNUserAllOf {
	this := JsonMDNUserAllOf{}
	return &this
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *JsonMDNUserAllOf) GetLastName() string {
	if o == nil || o.LastName == nil {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonMDNUserAllOf) GetLastNameOk() (*string, bool) {
	if o == nil || o.LastName == nil {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *JsonMDNUserAllOf) HasLastName() bool {
	if o != nil && o.LastName != nil {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *JsonMDNUserAllOf) SetLastName(v string) {
	o.LastName = &v
}

// GetMail returns the Mail field value if set, zero value otherwise.
func (o *JsonMDNUserAllOf) GetMail() string {
	if o == nil || o.Mail == nil {
		var ret string
		return ret
	}
	return *o.Mail
}

// GetMailOk returns a tuple with the Mail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonMDNUserAllOf) GetMailOk() (*string, bool) {
	if o == nil || o.Mail == nil {
		return nil, false
	}
	return o.Mail, true
}

// HasMail returns a boolean if a field has been set.
func (o *JsonMDNUserAllOf) HasMail() bool {
	if o != nil && o.Mail != nil {
		return true
	}

	return false
}

// SetMail gets a reference to the given string and assigns it to the Mail field.
func (o *JsonMDNUserAllOf) SetMail(v string) {
	o.Mail = &v
}

// GetSocialAccounts returns the SocialAccounts field value if set, zero value otherwise.
func (o *JsonMDNUserAllOf) GetSocialAccounts() []JsonMDNSocialUserObject {
	if o == nil || o.SocialAccounts == nil {
		var ret []JsonMDNSocialUserObject
		return ret
	}
	return *o.SocialAccounts
}

// GetSocialAccountsOk returns a tuple with the SocialAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonMDNUserAllOf) GetSocialAccountsOk() (*[]JsonMDNSocialUserObject, bool) {
	if o == nil || o.SocialAccounts == nil {
		return nil, false
	}
	return o.SocialAccounts, true
}

// HasSocialAccounts returns a boolean if a field has been set.
func (o *JsonMDNUserAllOf) HasSocialAccounts() bool {
	if o != nil && o.SocialAccounts != nil {
		return true
	}

	return false
}

// SetSocialAccounts gets a reference to the given []JsonMDNSocialUserObject and assigns it to the SocialAccounts field.
func (o *JsonMDNUserAllOf) SetSocialAccounts(v []JsonMDNSocialUserObject) {
	o.SocialAccounts = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *JsonMDNUserAllOf) GetFirstName() string {
	if o == nil || o.FirstName == nil {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonMDNUserAllOf) GetFirstNameOk() (*string, bool) {
	if o == nil || o.FirstName == nil {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *JsonMDNUserAllOf) HasFirstName() bool {
	if o != nil && o.FirstName != nil {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *JsonMDNUserAllOf) SetFirstName(v string) {
	o.FirstName = &v
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *JsonMDNUserAllOf) GetCredentials() JsonMDNUserCredentials {
	if o == nil || o.Credentials == nil {
		var ret JsonMDNUserCredentials
		return ret
	}
	return *o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonMDNUserAllOf) GetCredentialsOk() (*JsonMDNUserCredentials, bool) {
	if o == nil || o.Credentials == nil {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *JsonMDNUserAllOf) HasCredentials() bool {
	if o != nil && o.Credentials != nil {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given JsonMDNUserCredentials and assigns it to the Credentials field.
func (o *JsonMDNUserAllOf) SetCredentials(v JsonMDNUserCredentials) {
	o.Credentials = &v
}

// GetGuid returns the Guid field value if set, zero value otherwise.
func (o *JsonMDNUserAllOf) GetGuid() string {
	if o == nil || o.Guid == nil {
		var ret string
		return ret
	}
	return *o.Guid
}

// GetGuidOk returns a tuple with the Guid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonMDNUserAllOf) GetGuidOk() (*string, bool) {
	if o == nil || o.Guid == nil {
		return nil, false
	}
	return o.Guid, true
}

// HasGuid returns a boolean if a field has been set.
func (o *JsonMDNUserAllOf) HasGuid() bool {
	if o != nil && o.Guid != nil {
		return true
	}

	return false
}

// SetGuid gets a reference to the given string and assigns it to the Guid field.
func (o *JsonMDNUserAllOf) SetGuid(v string) {
	o.Guid = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *JsonMDNUserAllOf) GetSettings() []JsonMDNUserSetting {
	if o == nil || o.Settings == nil {
		var ret []JsonMDNUserSetting
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonMDNUserAllOf) GetSettingsOk() (*[]JsonMDNUserSetting, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *JsonMDNUserAllOf) HasSettings() bool {
	if o != nil && o.Settings != nil {
		return true
	}

	return false
}

// SetSettings gets a reference to the given []JsonMDNUserSetting and assigns it to the Settings field.
func (o *JsonMDNUserAllOf) SetSettings(v []JsonMDNUserSetting) {
	o.Settings = &v
}

func (o JsonMDNUserAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LastName != nil {
		toSerialize["lastName"] = o.LastName
	}
	if o.Mail != nil {
		toSerialize["mail"] = o.Mail
	}
	if o.SocialAccounts != nil {
		toSerialize["socialAccounts"] = o.SocialAccounts
	}
	if o.FirstName != nil {
		toSerialize["firstName"] = o.FirstName
	}
	if o.Credentials != nil {
		toSerialize["credentials"] = o.Credentials
	}
	if o.Guid != nil {
		toSerialize["guid"] = o.Guid
	}
	if o.Settings != nil {
		toSerialize["settings"] = o.Settings
	}
	return json.Marshal(toSerialize)
}

type NullableJsonMDNUserAllOf struct {
	value *JsonMDNUserAllOf
	isSet bool
}

func (v NullableJsonMDNUserAllOf) Get() *JsonMDNUserAllOf {
	return v.value
}

func (v *NullableJsonMDNUserAllOf) Set(val *JsonMDNUserAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableJsonMDNUserAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableJsonMDNUserAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJsonMDNUserAllOf(val *JsonMDNUserAllOf) *NullableJsonMDNUserAllOf {
	return &NullableJsonMDNUserAllOf{value: val, isSet: true}
}

func (v NullableJsonMDNUserAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJsonMDNUserAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


