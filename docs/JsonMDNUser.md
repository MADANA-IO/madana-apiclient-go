# JsonMDNUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Activated** | Pointer to **string** |  | [optional] 
**LastActive** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **string** |  | [optional] 
**UserName** | Pointer to **string** |  | [optional] 
**Image** | Pointer to **string** |  | [optional] 
**Settings** | Pointer to [**[]JsonMDNUserSetting**](json_MDN_UserSetting.md) |  | [optional] 
**Credentials** | Pointer to [**JsonMDNUserCredentials**](json_MDN_UserCredentials.md) |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**FirstName** | Pointer to **string** |  | [optional] 
**SocialAccounts** | Pointer to [**[]JsonMDNSocialUserObject**](json_MDN_SocialUserObject.md) |  | [optional] 
**Mail** | Pointer to **string** |  | [optional] 
**Guid** | Pointer to **string** |  | [optional] 

## Methods

### NewJsonMDNUser

`func NewJsonMDNUser() *JsonMDNUser`

NewJsonMDNUser instantiates a new JsonMDNUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJsonMDNUserWithDefaults

`func NewJsonMDNUserWithDefaults() *JsonMDNUser`

NewJsonMDNUserWithDefaults instantiates a new JsonMDNUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivated

`func (o *JsonMDNUser) GetActivated() string`

GetActivated returns the Activated field if non-nil, zero value otherwise.

### GetActivatedOk

`func (o *JsonMDNUser) GetActivatedOk() (*string, bool)`

GetActivatedOk returns a tuple with the Activated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivated

`func (o *JsonMDNUser) SetActivated(v string)`

SetActivated sets Activated field to given value.

### HasActivated

`func (o *JsonMDNUser) HasActivated() bool`

HasActivated returns a boolean if a field has been set.

### GetLastActive

`func (o *JsonMDNUser) GetLastActive() string`

GetLastActive returns the LastActive field if non-nil, zero value otherwise.

### GetLastActiveOk

`func (o *JsonMDNUser) GetLastActiveOk() (*string, bool)`

GetLastActiveOk returns a tuple with the LastActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActive

`func (o *JsonMDNUser) SetLastActive(v string)`

SetLastActive sets LastActive field to given value.

### HasLastActive

`func (o *JsonMDNUser) HasLastActive() bool`

HasLastActive returns a boolean if a field has been set.

### GetCreated

`func (o *JsonMDNUser) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *JsonMDNUser) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *JsonMDNUser) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *JsonMDNUser) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUserName

`func (o *JsonMDNUser) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *JsonMDNUser) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *JsonMDNUser) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *JsonMDNUser) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetImage

`func (o *JsonMDNUser) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *JsonMDNUser) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *JsonMDNUser) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *JsonMDNUser) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetSettings

`func (o *JsonMDNUser) GetSettings() []JsonMDNUserSetting`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *JsonMDNUser) GetSettingsOk() (*[]JsonMDNUserSetting, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *JsonMDNUser) SetSettings(v []JsonMDNUserSetting)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *JsonMDNUser) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetCredentials

`func (o *JsonMDNUser) GetCredentials() JsonMDNUserCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *JsonMDNUser) GetCredentialsOk() (*JsonMDNUserCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *JsonMDNUser) SetCredentials(v JsonMDNUserCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *JsonMDNUser) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetLastName

`func (o *JsonMDNUser) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *JsonMDNUser) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *JsonMDNUser) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *JsonMDNUser) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetFirstName

`func (o *JsonMDNUser) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *JsonMDNUser) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *JsonMDNUser) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *JsonMDNUser) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetSocialAccounts

`func (o *JsonMDNUser) GetSocialAccounts() []JsonMDNSocialUserObject`

GetSocialAccounts returns the SocialAccounts field if non-nil, zero value otherwise.

### GetSocialAccountsOk

`func (o *JsonMDNUser) GetSocialAccountsOk() (*[]JsonMDNSocialUserObject, bool)`

GetSocialAccountsOk returns a tuple with the SocialAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialAccounts

`func (o *JsonMDNUser) SetSocialAccounts(v []JsonMDNSocialUserObject)`

SetSocialAccounts sets SocialAccounts field to given value.

### HasSocialAccounts

`func (o *JsonMDNUser) HasSocialAccounts() bool`

HasSocialAccounts returns a boolean if a field has been set.

### GetMail

`func (o *JsonMDNUser) GetMail() string`

GetMail returns the Mail field if non-nil, zero value otherwise.

### GetMailOk

`func (o *JsonMDNUser) GetMailOk() (*string, bool)`

GetMailOk returns a tuple with the Mail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMail

`func (o *JsonMDNUser) SetMail(v string)`

SetMail sets Mail field to given value.

### HasMail

`func (o *JsonMDNUser) HasMail() bool`

HasMail returns a boolean if a field has been set.

### GetGuid

`func (o *JsonMDNUser) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *JsonMDNUser) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *JsonMDNUser) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *JsonMDNUser) HasGuid() bool`

HasGuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


