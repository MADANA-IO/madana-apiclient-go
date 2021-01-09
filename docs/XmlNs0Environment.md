# XmlNs0Environment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **[]string** |  | [optional] 
**DefaultRunConfiguration** | Pointer to [**XmlNs0RunConfig**](xml_ns0_runConfig.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**IpfsHash** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Packages** | Pointer to **[]string** |  | [optional] 
**Published** | Pointer to **bool** |  | [optional] 
**RootHashOffset** | Pointer to **string** |  | [optional] 
**Roothash** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **string** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 

## Methods

### NewXmlNs0Environment

`func NewXmlNs0Environment() *XmlNs0Environment`

NewXmlNs0Environment instantiates a new XmlNs0Environment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewXmlNs0EnvironmentWithDefaults

`func NewXmlNs0EnvironmentWithDefaults() *XmlNs0Environment`

NewXmlNs0EnvironmentWithDefaults instantiates a new XmlNs0Environment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *XmlNs0Environment) GetContent() []string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *XmlNs0Environment) GetContentOk() (*[]string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *XmlNs0Environment) SetContent(v []string)`

SetContent sets Content field to given value.

### HasContent

`func (o *XmlNs0Environment) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetDefaultRunConfiguration

`func (o *XmlNs0Environment) GetDefaultRunConfiguration() XmlNs0RunConfig`

GetDefaultRunConfiguration returns the DefaultRunConfiguration field if non-nil, zero value otherwise.

### GetDefaultRunConfigurationOk

`func (o *XmlNs0Environment) GetDefaultRunConfigurationOk() (*XmlNs0RunConfig, bool)`

GetDefaultRunConfigurationOk returns a tuple with the DefaultRunConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRunConfiguration

`func (o *XmlNs0Environment) SetDefaultRunConfiguration(v XmlNs0RunConfig)`

SetDefaultRunConfiguration sets DefaultRunConfiguration field to given value.

### HasDefaultRunConfiguration

`func (o *XmlNs0Environment) HasDefaultRunConfiguration() bool`

HasDefaultRunConfiguration returns a boolean if a field has been set.

### GetDescription

`func (o *XmlNs0Environment) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *XmlNs0Environment) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *XmlNs0Environment) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *XmlNs0Environment) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIpfsHash

`func (o *XmlNs0Environment) GetIpfsHash() string`

GetIpfsHash returns the IpfsHash field if non-nil, zero value otherwise.

### GetIpfsHashOk

`func (o *XmlNs0Environment) GetIpfsHashOk() (*string, bool)`

GetIpfsHashOk returns a tuple with the IpfsHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpfsHash

`func (o *XmlNs0Environment) SetIpfsHash(v string)`

SetIpfsHash sets IpfsHash field to given value.

### HasIpfsHash

`func (o *XmlNs0Environment) HasIpfsHash() bool`

HasIpfsHash returns a boolean if a field has been set.

### GetName

`func (o *XmlNs0Environment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *XmlNs0Environment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *XmlNs0Environment) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *XmlNs0Environment) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPackages

`func (o *XmlNs0Environment) GetPackages() []string`

GetPackages returns the Packages field if non-nil, zero value otherwise.

### GetPackagesOk

`func (o *XmlNs0Environment) GetPackagesOk() (*[]string, bool)`

GetPackagesOk returns a tuple with the Packages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackages

`func (o *XmlNs0Environment) SetPackages(v []string)`

SetPackages sets Packages field to given value.

### HasPackages

`func (o *XmlNs0Environment) HasPackages() bool`

HasPackages returns a boolean if a field has been set.

### GetPublished

`func (o *XmlNs0Environment) GetPublished() bool`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *XmlNs0Environment) GetPublishedOk() (*bool, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *XmlNs0Environment) SetPublished(v bool)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *XmlNs0Environment) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetRootHashOffset

`func (o *XmlNs0Environment) GetRootHashOffset() string`

GetRootHashOffset returns the RootHashOffset field if non-nil, zero value otherwise.

### GetRootHashOffsetOk

`func (o *XmlNs0Environment) GetRootHashOffsetOk() (*string, bool)`

GetRootHashOffsetOk returns a tuple with the RootHashOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootHashOffset

`func (o *XmlNs0Environment) SetRootHashOffset(v string)`

SetRootHashOffset sets RootHashOffset field to given value.

### HasRootHashOffset

`func (o *XmlNs0Environment) HasRootHashOffset() bool`

HasRootHashOffset returns a boolean if a field has been set.

### GetRoothash

`func (o *XmlNs0Environment) GetRoothash() string`

GetRoothash returns the Roothash field if non-nil, zero value otherwise.

### GetRoothashOk

`func (o *XmlNs0Environment) GetRoothashOk() (*string, bool)`

GetRoothashOk returns a tuple with the Roothash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoothash

`func (o *XmlNs0Environment) SetRoothash(v string)`

SetRoothash sets Roothash field to given value.

### HasRoothash

`func (o *XmlNs0Environment) HasRoothash() bool`

HasRoothash returns a boolean if a field has been set.

### GetSize

`func (o *XmlNs0Environment) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *XmlNs0Environment) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *XmlNs0Environment) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *XmlNs0Environment) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetUuid

`func (o *XmlNs0Environment) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *XmlNs0Environment) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *XmlNs0Environment) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *XmlNs0Environment) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


