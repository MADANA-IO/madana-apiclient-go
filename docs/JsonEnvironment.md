# JsonEnvironment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpfsHash** | Pointer to **string** |  | [optional] 
**Roothash** | Pointer to **string** |  | [optional] 
**Content** | Pointer to **[]string** |  | [optional] 
**RootHashOffset** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**Published** | Pointer to **bool** |  | [optional] 
**Packages** | Pointer to **[]string** |  | [optional] 
**DefaultRunConfiguration** | Pointer to [**JsonRunConfig**](json_RunConfig.md) |  | [optional] 

## Methods

### NewJsonEnvironment

`func NewJsonEnvironment() *JsonEnvironment`

NewJsonEnvironment instantiates a new JsonEnvironment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJsonEnvironmentWithDefaults

`func NewJsonEnvironmentWithDefaults() *JsonEnvironment`

NewJsonEnvironmentWithDefaults instantiates a new JsonEnvironment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpfsHash

`func (o *JsonEnvironment) GetIpfsHash() string`

GetIpfsHash returns the IpfsHash field if non-nil, zero value otherwise.

### GetIpfsHashOk

`func (o *JsonEnvironment) GetIpfsHashOk() (*string, bool)`

GetIpfsHashOk returns a tuple with the IpfsHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpfsHash

`func (o *JsonEnvironment) SetIpfsHash(v string)`

SetIpfsHash sets IpfsHash field to given value.

### HasIpfsHash

`func (o *JsonEnvironment) HasIpfsHash() bool`

HasIpfsHash returns a boolean if a field has been set.

### GetRoothash

`func (o *JsonEnvironment) GetRoothash() string`

GetRoothash returns the Roothash field if non-nil, zero value otherwise.

### GetRoothashOk

`func (o *JsonEnvironment) GetRoothashOk() (*string, bool)`

GetRoothashOk returns a tuple with the Roothash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoothash

`func (o *JsonEnvironment) SetRoothash(v string)`

SetRoothash sets Roothash field to given value.

### HasRoothash

`func (o *JsonEnvironment) HasRoothash() bool`

HasRoothash returns a boolean if a field has been set.

### GetContent

`func (o *JsonEnvironment) GetContent() []string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *JsonEnvironment) GetContentOk() (*[]string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *JsonEnvironment) SetContent(v []string)`

SetContent sets Content field to given value.

### HasContent

`func (o *JsonEnvironment) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetRootHashOffset

`func (o *JsonEnvironment) GetRootHashOffset() string`

GetRootHashOffset returns the RootHashOffset field if non-nil, zero value otherwise.

### GetRootHashOffsetOk

`func (o *JsonEnvironment) GetRootHashOffsetOk() (*string, bool)`

GetRootHashOffsetOk returns a tuple with the RootHashOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootHashOffset

`func (o *JsonEnvironment) SetRootHashOffset(v string)`

SetRootHashOffset sets RootHashOffset field to given value.

### HasRootHashOffset

`func (o *JsonEnvironment) HasRootHashOffset() bool`

HasRootHashOffset returns a boolean if a field has been set.

### GetSize

`func (o *JsonEnvironment) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *JsonEnvironment) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *JsonEnvironment) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *JsonEnvironment) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetDescription

`func (o *JsonEnvironment) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *JsonEnvironment) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *JsonEnvironment) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *JsonEnvironment) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *JsonEnvironment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JsonEnvironment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JsonEnvironment) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *JsonEnvironment) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUuid

`func (o *JsonEnvironment) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *JsonEnvironment) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *JsonEnvironment) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *JsonEnvironment) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetPublished

`func (o *JsonEnvironment) GetPublished() bool`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *JsonEnvironment) GetPublishedOk() (*bool, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *JsonEnvironment) SetPublished(v bool)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *JsonEnvironment) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetPackages

`func (o *JsonEnvironment) GetPackages() []string`

GetPackages returns the Packages field if non-nil, zero value otherwise.

### GetPackagesOk

`func (o *JsonEnvironment) GetPackagesOk() (*[]string, bool)`

GetPackagesOk returns a tuple with the Packages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackages

`func (o *JsonEnvironment) SetPackages(v []string)`

SetPackages sets Packages field to given value.

### HasPackages

`func (o *JsonEnvironment) HasPackages() bool`

HasPackages returns a boolean if a field has been set.

### GetDefaultRunConfiguration

`func (o *JsonEnvironment) GetDefaultRunConfiguration() JsonRunConfig`

GetDefaultRunConfiguration returns the DefaultRunConfiguration field if non-nil, zero value otherwise.

### GetDefaultRunConfigurationOk

`func (o *JsonEnvironment) GetDefaultRunConfigurationOk() (*JsonRunConfig, bool)`

GetDefaultRunConfigurationOk returns a tuple with the DefaultRunConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRunConfiguration

`func (o *JsonEnvironment) SetDefaultRunConfiguration(v JsonRunConfig)`

SetDefaultRunConfiguration sets DefaultRunConfiguration field to given value.

### HasDefaultRunConfiguration

`func (o *JsonEnvironment) HasDefaultRunConfiguration() bool`

HasDefaultRunConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


