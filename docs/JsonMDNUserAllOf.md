# JsonMDNUserAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SocialAccounts** | Pointer to [**[]JsonMDNSocialUserObject**](json_MDN_SocialUserObject.md) |  | [optional] 
**FirstName** | Pointer to **string** |  | [optional] 
**Guid** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**Credentials** | Pointer to [**JsonMDNUserCredentials**](json_MDN_UserCredentials.md) |  | [optional] 
**Settings** | Pointer to [**[]JsonMDNUserSetting**](json_MDN_UserSetting.md) |  | [optional] 
**Mail** | Pointer to **string** |  | [optional] 

## Methods

### NewJsonMDNUserAllOf

`func NewJsonMDNUserAllOf() *JsonMDNUserAllOf`

NewJsonMDNUserAllOf instantiates a new JsonMDNUserAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJsonMDNUserAllOfWithDefaults

`func NewJsonMDNUserAllOfWithDefaults() *JsonMDNUserAllOf`

NewJsonMDNUserAllOfWithDefaults instantiates a new JsonMDNUserAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSocialAccounts

`func (o *JsonMDNUserAllOf) GetSocialAccounts() []JsonMDNSocialUserObject`

GetSocialAccounts returns the SocialAccounts field if non-nil, zero value otherwise.

### GetSocialAccountsOk

`func (o *JsonMDNUserAllOf) GetSocialAccountsOk() (*[]JsonMDNSocialUserObject, bool)`

GetSocialAccountsOk returns a tuple with the SocialAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialAccounts

`func (o *JsonMDNUserAllOf) SetSocialAccounts(v []JsonMDNSocialUserObject)`

SetSocialAccounts sets SocialAccounts field to given value.

### HasSocialAccounts

`func (o *JsonMDNUserAllOf) HasSocialAccounts() bool`

HasSocialAccounts returns a boolean if a field has been set.

### GetFirstName

`func (o *JsonMDNUserAllOf) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *JsonMDNUserAllOf) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *JsonMDNUserAllOf) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *JsonMDNUserAllOf) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetGuid

`func (o *JsonMDNUserAllOf) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *JsonMDNUserAllOf) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *JsonMDNUserAllOf) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *JsonMDNUserAllOf) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetLastName

`func (o *JsonMDNUserAllOf) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *JsonMDNUserAllOf) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *JsonMDNUserAllOf) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *JsonMDNUserAllOf) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetCredentials

`func (o *JsonMDNUserAllOf) GetCredentials() JsonMDNUserCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *JsonMDNUserAllOf) GetCredentialsOk() (*JsonMDNUserCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *JsonMDNUserAllOf) SetCredentials(v JsonMDNUserCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *JsonMDNUserAllOf) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetSettings

`func (o *JsonMDNUserAllOf) GetSettings() []JsonMDNUserSetting`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *JsonMDNUserAllOf) GetSettingsOk() (*[]JsonMDNUserSetting, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *JsonMDNUserAllOf) SetSettings(v []JsonMDNUserSetting)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *JsonMDNUserAllOf) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetMail

`func (o *JsonMDNUserAllOf) GetMail() string`

GetMail returns the Mail field if non-nil, zero value otherwise.

### GetMailOk

`func (o *JsonMDNUserAllOf) GetMailOk() (*string, bool)`

GetMailOk returns a tuple with the Mail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMail

`func (o *JsonMDNUserAllOf) SetMail(v string)`

SetMail sets Mail field to given value.

### HasMail

`func (o *JsonMDNUserAllOf) HasMail() bool`

HasMail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


