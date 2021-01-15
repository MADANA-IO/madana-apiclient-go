/*
 * madana-api
 *
 * <h1>Using the madana-api</h1>        <p>This documentation contains a Quickstart Guide, relating client functionality and information about the available         endpoints and used datamodels.   </p>       <p> The madana-api and its implementations are still in heavy development. This means that there may be problems in our protocols, or there may be mistakes in our implementations. We take security vulnerabilities very seriously. If you discover a security issue, please bring it to our attention right away! If you find a vulnerability that may affect live deployments -- for example, by exposing a remote execution exploit -- please send your report privately to info@madana.io. Please DO NOT file a public issue. If the issue is a protocol weakness that cannot be immediately exploited or something not yet deployed, just discuss it openly   </p>   <br>   <p> Note: Not all functionality might be acessible without having accquired and api-license token. For more information visit <a href=\"https://www.madana.io\">www.madana.io</a> </p>       <br>
 *
 * API version: 0.5.0-master.27
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package madanaapiclient

import (
	"encoding/json"
)

// JsonMDNUser 
type JsonMDNUser struct {
	// 
	UserName *string `json:"userName,omitempty"`
	// 
	Activated *string `json:"activated,omitempty"`
	// 
	Image *string `json:"image,omitempty"`
	// 
	LastActive *string `json:"lastActive,omitempty"`
	// 
	Created *string `json:"created,omitempty"`
	// 
	LastName *string `json:"lastName,omitempty"`
	// 
	FirstName *string `json:"firstName,omitempty"`
	// 
	Mail *string `json:"mail,omitempty"`
	// 
	SocialAccounts *[]JsonMDNSocialUserObject `json:"socialAccounts,omitempty"`
	// 
	Settings *[]JsonMDNUserSetting `json:"settings,omitempty"`
	Credentials *JsonMDNUserCredentials `json:"credentials,omitempty"`
	// 
	Guid *string `json:"guid,omitempty"`
}

// NewJsonMDNUser instantiates a new JsonMDNUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJsonMDNUser() *JsonMDNUser {
	this := JsonMDNUser{}
	return &this
}

// NewJsonMDNUserWithDefaults instantiates a new JsonMDNUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJsonMDNUserWithDefaults() *JsonMDNUser {
	this := JsonMDNUser{}
	return &this
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *JsonMDNUser) GetUserName() string {
	if o == nil || o.UserName == nil {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonMDNUser) GetUserNameOk() (*string, bool) {
	if o == nil || o.UserName == nil {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *JsonMDNUser) HasUserName() bool {
	if o != nil && o.UserName != nil {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *JsonMDNUser) SetUserName(v string) {
	o.UserName = &v
}

// GetActivated returns the Activated field value if set, zero value otherwise.
func (o *JsonMDNUser) GetActivated() string {
	if o == nil || o.Activated == nil {
		var ret string
		return ret
	}
	return *o.Activated
}

// GetActivatedOk returns a tuple with the Activated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonMDNUser) GetActivatedOk() (*string, bool) {
	if o == nil || o.Activated == nil {
		return nil, false
	}
	return o.Activated, true
}

// HasActivated returns a boolean if a field has been set.
func (o *JsonMDNUser) HasActivated() bool {
	if o != nil && o.Activated != nil {
		return true
	}

	return false
}

// SetActivated gets a reference to the given string and assigns it to the Activated field.
func (o *JsonMDNUser) SetActivated(v string) {
	o.Activated = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *JsonMDNUser) GetImage() string {
	if o == nil || o.Image == nil {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonMDNUser) GetImageOk() (*string, bool) {
	if o == nil || o.Image == nil {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *JsonMDNUser) HasImage() bool {
	if o != nil && o.Image != nil {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *JsonMDNUser) SetImage(v string) {
	o.Image = &v
}

// GetLastActive returns the LastActive field value if set, zero value otherwise.
func (o *JsonMDNUser) GetLastActive() string {
	if o == nil || o.LastActive == nil {
		var ret string
		return ret
	}
	return *o.LastActive
}

// GetLastActiveOk returns a tuple with the LastActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonMDNUser) GetLastActiveOk() (*string, bool) {
	if o == nil || o.LastActive == nil {
		return nil, false
	}
	return o.LastActive, true
}

// HasLastActive returns a boolean if a field has been set.
func (o *JsonMDNUser) HasLastActive() bool {
	if o != nil && o.LastActive != nil {
		return true
	}

	return false
}

// SetLastActive gets a reference to the given string and assigns it to the LastActive field.
func (o *JsonMDNUser) SetLastActive(v string) {
	o.LastActive = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *JsonMDNUser) GetCreated() string {
	if o == nil || o.Created == nil {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonMDNUser) GetCreatedOk() (*string, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *JsonMDNUser) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *JsonMDNUser) SetCreated(v string) {
	o.Created = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *JsonMDNUser) GetLastName() string {
	if o == nil || o.LastName == nil {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonMDNUser) GetLastNameOk() (*string, bool) {
	if o == nil || o.LastName == nil {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *JsonMDNUser) HasLastName() bool {
	if o != nil && o.LastName != nil {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *JsonMDNUser) SetLastName(v string) {
	o.LastName = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *JsonMDNUser) GetFirstName() string {
	if o == nil || o.FirstName == nil {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonMDNUser) GetFirstNameOk() (*string, bool) {
	if o == nil || o.FirstName == nil {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *JsonMDNUser) HasFirstName() bool {
	if o != nil && o.FirstName != nil {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *JsonMDNUser) SetFirstName(v string) {
	o.FirstName = &v
}

// GetMail returns the Mail field value if set, zero value otherwise.
func (o *JsonMDNUser) GetMail() string {
	if o == nil || o.Mail == nil {
		var ret string
		return ret
	}
	return *o.Mail
}

// GetMailOk returns a tuple with the Mail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonMDNUser) GetMailOk() (*string, bool) {
	if o == nil || o.Mail == nil {
		return nil, false
	}
	return o.Mail, true
}

// HasMail returns a boolean if a field has been set.
func (o *JsonMDNUser) HasMail() bool {
	if o != nil && o.Mail != nil {
		return true
	}

	return false
}

// SetMail gets a reference to the given string and assigns it to the Mail field.
func (o *JsonMDNUser) SetMail(v string) {
	o.Mail = &v
}

// GetSocialAccounts returns the SocialAccounts field value if set, zero value otherwise.
func (o *JsonMDNUser) GetSocialAccounts() []JsonMDNSocialUserObject {
	if o == nil || o.SocialAccounts == nil {
		var ret []JsonMDNSocialUserObject
		return ret
	}
	return *o.SocialAccounts
}

// GetSocialAccountsOk returns a tuple with the SocialAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonMDNUser) GetSocialAccountsOk() (*[]JsonMDNSocialUserObject, bool) {
	if o == nil || o.SocialAccounts == nil {
		return nil, false
	}
	return o.SocialAccounts, true
}

// HasSocialAccounts returns a boolean if a field has been set.
func (o *JsonMDNUser) HasSocialAccounts() bool {
	if o != nil && o.SocialAccounts != nil {
		return true
	}

	return false
}

// SetSocialAccounts gets a reference to the given []JsonMDNSocialUserObject and assigns it to the SocialAccounts field.
func (o *JsonMDNUser) SetSocialAccounts(v []JsonMDNSocialUserObject) {
	o.SocialAccounts = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *JsonMDNUser) GetSettings() []JsonMDNUserSetting {
	if o == nil || o.Settings == nil {
		var ret []JsonMDNUserSetting
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonMDNUser) GetSettingsOk() (*[]JsonMDNUserSetting, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *JsonMDNUser) HasSettings() bool {
	if o != nil && o.Settings != nil {
		return true
	}

	return false
}

// SetSettings gets a reference to the given []JsonMDNUserSetting and assigns it to the Settings field.
func (o *JsonMDNUser) SetSettings(v []JsonMDNUserSetting) {
	o.Settings = &v
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *JsonMDNUser) GetCredentials() JsonMDNUserCredentials {
	if o == nil || o.Credentials == nil {
		var ret JsonMDNUserCredentials
		return ret
	}
	return *o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonMDNUser) GetCredentialsOk() (*JsonMDNUserCredentials, bool) {
	if o == nil || o.Credentials == nil {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *JsonMDNUser) HasCredentials() bool {
	if o != nil && o.Credentials != nil {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given JsonMDNUserCredentials and assigns it to the Credentials field.
func (o *JsonMDNUser) SetCredentials(v JsonMDNUserCredentials) {
	o.Credentials = &v
}

// GetGuid returns the Guid field value if set, zero value otherwise.
func (o *JsonMDNUser) GetGuid() string {
	if o == nil || o.Guid == nil {
		var ret string
		return ret
	}
	return *o.Guid
}

// GetGuidOk returns a tuple with the Guid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonMDNUser) GetGuidOk() (*string, bool) {
	if o == nil || o.Guid == nil {
		return nil, false
	}
	return o.Guid, true
}

// HasGuid returns a boolean if a field has been set.
func (o *JsonMDNUser) HasGuid() bool {
	if o != nil && o.Guid != nil {
		return true
	}

	return false
}

// SetGuid gets a reference to the given string and assigns it to the Guid field.
func (o *JsonMDNUser) SetGuid(v string) {
	o.Guid = &v
}

func (o JsonMDNUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UserName != nil {
		toSerialize["userName"] = o.UserName
	}
	if o.Activated != nil {
		toSerialize["activated"] = o.Activated
	}
	if o.Image != nil {
		toSerialize["image"] = o.Image
	}
	if o.LastActive != nil {
		toSerialize["lastActive"] = o.LastActive
	}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.LastName != nil {
		toSerialize["lastName"] = o.LastName
	}
	if o.FirstName != nil {
		toSerialize["firstName"] = o.FirstName
	}
	if o.Mail != nil {
		toSerialize["mail"] = o.Mail
	}
	if o.SocialAccounts != nil {
		toSerialize["socialAccounts"] = o.SocialAccounts
	}
	if o.Settings != nil {
		toSerialize["settings"] = o.Settings
	}
	if o.Credentials != nil {
		toSerialize["credentials"] = o.Credentials
	}
	if o.Guid != nil {
		toSerialize["guid"] = o.Guid
	}
	return json.Marshal(toSerialize)
}

type NullableJsonMDNUser struct {
	value *JsonMDNUser
	isSet bool
}

func (v NullableJsonMDNUser) Get() *JsonMDNUser {
	return v.value
}

func (v *NullableJsonMDNUser) Set(val *JsonMDNUser) {
	v.value = val
	v.isSet = true
}

func (v NullableJsonMDNUser) IsSet() bool {
	return v.isSet
}

func (v *NullableJsonMDNUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJsonMDNUser(val *JsonMDNUser) *NullableJsonMDNUser {
	return &NullableJsonMDNUser{value: val, isSet: true}
}

func (v NullableJsonMDNUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJsonMDNUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


