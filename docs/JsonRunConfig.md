# JsonRunConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Run** | Pointer to **string** |  | [optional] 
**Args** | Pointer to **[]string** |  | [optional] 
**DiskConfig** | Pointer to [**[]JsonDiskConfig**](JsonDiskConfig.md) |  | [optional] 
**Environment** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewJsonRunConfig

`func NewJsonRunConfig() *JsonRunConfig`

NewJsonRunConfig instantiates a new JsonRunConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJsonRunConfigWithDefaults

`func NewJsonRunConfigWithDefaults() *JsonRunConfig`

NewJsonRunConfigWithDefaults instantiates a new JsonRunConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRun

`func (o *JsonRunConfig) GetRun() string`

GetRun returns the Run field if non-nil, zero value otherwise.

### GetRunOk

`func (o *JsonRunConfig) GetRunOk() (*string, bool)`

GetRunOk returns a tuple with the Run field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRun

`func (o *JsonRunConfig) SetRun(v string)`

SetRun sets Run field to given value.

### HasRun

`func (o *JsonRunConfig) HasRun() bool`

HasRun returns a boolean if a field has been set.

### GetArgs

`func (o *JsonRunConfig) GetArgs() []string`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *JsonRunConfig) GetArgsOk() (*[]string, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *JsonRunConfig) SetArgs(v []string)`

SetArgs sets Args field to given value.

### HasArgs

`func (o *JsonRunConfig) HasArgs() bool`

HasArgs returns a boolean if a field has been set.

### GetDiskConfig

`func (o *JsonRunConfig) GetDiskConfig() []JsonDiskConfig`

GetDiskConfig returns the DiskConfig field if non-nil, zero value otherwise.

### GetDiskConfigOk

`func (o *JsonRunConfig) GetDiskConfigOk() (*[]JsonDiskConfig, bool)`

GetDiskConfigOk returns a tuple with the DiskConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskConfig

`func (o *JsonRunConfig) SetDiskConfig(v []JsonDiskConfig)`

SetDiskConfig sets DiskConfig field to given value.

### HasDiskConfig

`func (o *JsonRunConfig) HasDiskConfig() bool`

HasDiskConfig returns a boolean if a field has been set.

### GetEnvironment

`func (o *JsonRunConfig) GetEnvironment() map[string]string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *JsonRunConfig) GetEnvironmentOk() (*map[string]string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *JsonRunConfig) SetEnvironment(v map[string]string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *JsonRunConfig) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


