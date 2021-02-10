# JsonEnclaveRunRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentUUID** | Pointer to **string** |  | [optional] 
**WireguardPublicKey** | Pointer to **string** |  | [optional] 
**Ports** | Pointer to [**[]JsonEnclavePort**](JsonEnclavePort.md) |  | [optional] 
**UsingDefaultRunConfig** | Pointer to **bool** |  | [optional] 
**EnclaveExecutionType** | Pointer to **string** |  | [optional] 

## Methods

### NewJsonEnclaveRunRequest

`func NewJsonEnclaveRunRequest() *JsonEnclaveRunRequest`

NewJsonEnclaveRunRequest instantiates a new JsonEnclaveRunRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJsonEnclaveRunRequestWithDefaults

`func NewJsonEnclaveRunRequestWithDefaults() *JsonEnclaveRunRequest`

NewJsonEnclaveRunRequestWithDefaults instantiates a new JsonEnclaveRunRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentUUID

`func (o *JsonEnclaveRunRequest) GetEnvironmentUUID() string`

GetEnvironmentUUID returns the EnvironmentUUID field if non-nil, zero value otherwise.

### GetEnvironmentUUIDOk

`func (o *JsonEnclaveRunRequest) GetEnvironmentUUIDOk() (*string, bool)`

GetEnvironmentUUIDOk returns a tuple with the EnvironmentUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentUUID

`func (o *JsonEnclaveRunRequest) SetEnvironmentUUID(v string)`

SetEnvironmentUUID sets EnvironmentUUID field to given value.

### HasEnvironmentUUID

`func (o *JsonEnclaveRunRequest) HasEnvironmentUUID() bool`

HasEnvironmentUUID returns a boolean if a field has been set.

### GetWireguardPublicKey

`func (o *JsonEnclaveRunRequest) GetWireguardPublicKey() string`

GetWireguardPublicKey returns the WireguardPublicKey field if non-nil, zero value otherwise.

### GetWireguardPublicKeyOk

`func (o *JsonEnclaveRunRequest) GetWireguardPublicKeyOk() (*string, bool)`

GetWireguardPublicKeyOk returns a tuple with the WireguardPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWireguardPublicKey

`func (o *JsonEnclaveRunRequest) SetWireguardPublicKey(v string)`

SetWireguardPublicKey sets WireguardPublicKey field to given value.

### HasWireguardPublicKey

`func (o *JsonEnclaveRunRequest) HasWireguardPublicKey() bool`

HasWireguardPublicKey returns a boolean if a field has been set.

### GetPorts

`func (o *JsonEnclaveRunRequest) GetPorts() []JsonEnclavePort`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *JsonEnclaveRunRequest) GetPortsOk() (*[]JsonEnclavePort, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *JsonEnclaveRunRequest) SetPorts(v []JsonEnclavePort)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *JsonEnclaveRunRequest) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetUsingDefaultRunConfig

`func (o *JsonEnclaveRunRequest) GetUsingDefaultRunConfig() bool`

GetUsingDefaultRunConfig returns the UsingDefaultRunConfig field if non-nil, zero value otherwise.

### GetUsingDefaultRunConfigOk

`func (o *JsonEnclaveRunRequest) GetUsingDefaultRunConfigOk() (*bool, bool)`

GetUsingDefaultRunConfigOk returns a tuple with the UsingDefaultRunConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsingDefaultRunConfig

`func (o *JsonEnclaveRunRequest) SetUsingDefaultRunConfig(v bool)`

SetUsingDefaultRunConfig sets UsingDefaultRunConfig field to given value.

### HasUsingDefaultRunConfig

`func (o *JsonEnclaveRunRequest) HasUsingDefaultRunConfig() bool`

HasUsingDefaultRunConfig returns a boolean if a field has been set.

### GetEnclaveExecutionType

`func (o *JsonEnclaveRunRequest) GetEnclaveExecutionType() string`

GetEnclaveExecutionType returns the EnclaveExecutionType field if non-nil, zero value otherwise.

### GetEnclaveExecutionTypeOk

`func (o *JsonEnclaveRunRequest) GetEnclaveExecutionTypeOk() (*string, bool)`

GetEnclaveExecutionTypeOk returns a tuple with the EnclaveExecutionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnclaveExecutionType

`func (o *JsonEnclaveRunRequest) SetEnclaveExecutionType(v string)`

SetEnclaveExecutionType sets EnclaveExecutionType field to given value.

### HasEnclaveExecutionType

`func (o *JsonEnclaveRunRequest) HasEnclaveExecutionType() bool`

HasEnclaveExecutionType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


