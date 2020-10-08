# JsonEnclaveProcess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndingTime** | Pointer to **string** |  | [optional] 
**Process** | Pointer to [**JsonProcess**](json_Process.md) |  | [optional] 
**WireguardServer** | Pointer to **string** |  | [optional] 
**StartupTime** | Pointer to **string** |  | [optional] 
**EnclaveInputstream** | Pointer to **map[string]interface{}** |  | [optional] 
**RemoteControlServer** | Pointer to **string** |  | [optional] 
**InternalIdent** | Pointer to **string** |  | [optional] 
**StartupCMD** | Pointer to **string** |  | [optional] 
**Environment** | Pointer to [**JsonEnvironment**](json_Environment.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**InternalRemoteControlServer** | Pointer to **string** |  | [optional] 
**ConsoleOutput** | Pointer to **string** |  | [optional] 
**EnclaveIdent** | Pointer to **string** |  | [optional] 
**SignerIdent** | Pointer to **string** |  | [optional] 
**AttestationServer** | Pointer to **string** |  | [optional] 
**WgInterface** | Pointer to [**JsonWireguardInterface**](json_WireguardInterface.md) |  | [optional] 
**InternalAttesationServer** | Pointer to **string** |  | [optional] 
**WireguardPublicKey** | Pointer to **string** |  | [optional] 
**PublicIdent** | Pointer to **string** |  | [optional] 
**InternalWireguardServer** | Pointer to **string** |  | [optional] 

## Methods

### NewJsonEnclaveProcess

`func NewJsonEnclaveProcess() *JsonEnclaveProcess`

NewJsonEnclaveProcess instantiates a new JsonEnclaveProcess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJsonEnclaveProcessWithDefaults

`func NewJsonEnclaveProcessWithDefaults() *JsonEnclaveProcess`

NewJsonEnclaveProcessWithDefaults instantiates a new JsonEnclaveProcess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndingTime

`func (o *JsonEnclaveProcess) GetEndingTime() string`

GetEndingTime returns the EndingTime field if non-nil, zero value otherwise.

### GetEndingTimeOk

`func (o *JsonEnclaveProcess) GetEndingTimeOk() (*string, bool)`

GetEndingTimeOk returns a tuple with the EndingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndingTime

`func (o *JsonEnclaveProcess) SetEndingTime(v string)`

SetEndingTime sets EndingTime field to given value.

### HasEndingTime

`func (o *JsonEnclaveProcess) HasEndingTime() bool`

HasEndingTime returns a boolean if a field has been set.

### GetProcess

`func (o *JsonEnclaveProcess) GetProcess() JsonProcess`

GetProcess returns the Process field if non-nil, zero value otherwise.

### GetProcessOk

`func (o *JsonEnclaveProcess) GetProcessOk() (*JsonProcess, bool)`

GetProcessOk returns a tuple with the Process field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcess

`func (o *JsonEnclaveProcess) SetProcess(v JsonProcess)`

SetProcess sets Process field to given value.

### HasProcess

`func (o *JsonEnclaveProcess) HasProcess() bool`

HasProcess returns a boolean if a field has been set.

### GetWireguardServer

`func (o *JsonEnclaveProcess) GetWireguardServer() string`

GetWireguardServer returns the WireguardServer field if non-nil, zero value otherwise.

### GetWireguardServerOk

`func (o *JsonEnclaveProcess) GetWireguardServerOk() (*string, bool)`

GetWireguardServerOk returns a tuple with the WireguardServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWireguardServer

`func (o *JsonEnclaveProcess) SetWireguardServer(v string)`

SetWireguardServer sets WireguardServer field to given value.

### HasWireguardServer

`func (o *JsonEnclaveProcess) HasWireguardServer() bool`

HasWireguardServer returns a boolean if a field has been set.

### GetStartupTime

`func (o *JsonEnclaveProcess) GetStartupTime() string`

GetStartupTime returns the StartupTime field if non-nil, zero value otherwise.

### GetStartupTimeOk

`func (o *JsonEnclaveProcess) GetStartupTimeOk() (*string, bool)`

GetStartupTimeOk returns a tuple with the StartupTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartupTime

`func (o *JsonEnclaveProcess) SetStartupTime(v string)`

SetStartupTime sets StartupTime field to given value.

### HasStartupTime

`func (o *JsonEnclaveProcess) HasStartupTime() bool`

HasStartupTime returns a boolean if a field has been set.

### GetEnclaveInputstream

`func (o *JsonEnclaveProcess) GetEnclaveInputstream() map[string]interface{}`

GetEnclaveInputstream returns the EnclaveInputstream field if non-nil, zero value otherwise.

### GetEnclaveInputstreamOk

`func (o *JsonEnclaveProcess) GetEnclaveInputstreamOk() (*map[string]interface{}, bool)`

GetEnclaveInputstreamOk returns a tuple with the EnclaveInputstream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnclaveInputstream

`func (o *JsonEnclaveProcess) SetEnclaveInputstream(v map[string]interface{})`

SetEnclaveInputstream sets EnclaveInputstream field to given value.

### HasEnclaveInputstream

`func (o *JsonEnclaveProcess) HasEnclaveInputstream() bool`

HasEnclaveInputstream returns a boolean if a field has been set.

### GetRemoteControlServer

`func (o *JsonEnclaveProcess) GetRemoteControlServer() string`

GetRemoteControlServer returns the RemoteControlServer field if non-nil, zero value otherwise.

### GetRemoteControlServerOk

`func (o *JsonEnclaveProcess) GetRemoteControlServerOk() (*string, bool)`

GetRemoteControlServerOk returns a tuple with the RemoteControlServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteControlServer

`func (o *JsonEnclaveProcess) SetRemoteControlServer(v string)`

SetRemoteControlServer sets RemoteControlServer field to given value.

### HasRemoteControlServer

`func (o *JsonEnclaveProcess) HasRemoteControlServer() bool`

HasRemoteControlServer returns a boolean if a field has been set.

### GetInternalIdent

`func (o *JsonEnclaveProcess) GetInternalIdent() string`

GetInternalIdent returns the InternalIdent field if non-nil, zero value otherwise.

### GetInternalIdentOk

`func (o *JsonEnclaveProcess) GetInternalIdentOk() (*string, bool)`

GetInternalIdentOk returns a tuple with the InternalIdent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalIdent

`func (o *JsonEnclaveProcess) SetInternalIdent(v string)`

SetInternalIdent sets InternalIdent field to given value.

### HasInternalIdent

`func (o *JsonEnclaveProcess) HasInternalIdent() bool`

HasInternalIdent returns a boolean if a field has been set.

### GetStartupCMD

`func (o *JsonEnclaveProcess) GetStartupCMD() string`

GetStartupCMD returns the StartupCMD field if non-nil, zero value otherwise.

### GetStartupCMDOk

`func (o *JsonEnclaveProcess) GetStartupCMDOk() (*string, bool)`

GetStartupCMDOk returns a tuple with the StartupCMD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartupCMD

`func (o *JsonEnclaveProcess) SetStartupCMD(v string)`

SetStartupCMD sets StartupCMD field to given value.

### HasStartupCMD

`func (o *JsonEnclaveProcess) HasStartupCMD() bool`

HasStartupCMD returns a boolean if a field has been set.

### GetEnvironment

`func (o *JsonEnclaveProcess) GetEnvironment() JsonEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *JsonEnclaveProcess) GetEnvironmentOk() (*JsonEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *JsonEnclaveProcess) SetEnvironment(v JsonEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *JsonEnclaveProcess) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetStatus

`func (o *JsonEnclaveProcess) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *JsonEnclaveProcess) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *JsonEnclaveProcess) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *JsonEnclaveProcess) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetInternalRemoteControlServer

`func (o *JsonEnclaveProcess) GetInternalRemoteControlServer() string`

GetInternalRemoteControlServer returns the InternalRemoteControlServer field if non-nil, zero value otherwise.

### GetInternalRemoteControlServerOk

`func (o *JsonEnclaveProcess) GetInternalRemoteControlServerOk() (*string, bool)`

GetInternalRemoteControlServerOk returns a tuple with the InternalRemoteControlServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalRemoteControlServer

`func (o *JsonEnclaveProcess) SetInternalRemoteControlServer(v string)`

SetInternalRemoteControlServer sets InternalRemoteControlServer field to given value.

### HasInternalRemoteControlServer

`func (o *JsonEnclaveProcess) HasInternalRemoteControlServer() bool`

HasInternalRemoteControlServer returns a boolean if a field has been set.

### GetConsoleOutput

`func (o *JsonEnclaveProcess) GetConsoleOutput() string`

GetConsoleOutput returns the ConsoleOutput field if non-nil, zero value otherwise.

### GetConsoleOutputOk

`func (o *JsonEnclaveProcess) GetConsoleOutputOk() (*string, bool)`

GetConsoleOutputOk returns a tuple with the ConsoleOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleOutput

`func (o *JsonEnclaveProcess) SetConsoleOutput(v string)`

SetConsoleOutput sets ConsoleOutput field to given value.

### HasConsoleOutput

`func (o *JsonEnclaveProcess) HasConsoleOutput() bool`

HasConsoleOutput returns a boolean if a field has been set.

### GetEnclaveIdent

`func (o *JsonEnclaveProcess) GetEnclaveIdent() string`

GetEnclaveIdent returns the EnclaveIdent field if non-nil, zero value otherwise.

### GetEnclaveIdentOk

`func (o *JsonEnclaveProcess) GetEnclaveIdentOk() (*string, bool)`

GetEnclaveIdentOk returns a tuple with the EnclaveIdent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnclaveIdent

`func (o *JsonEnclaveProcess) SetEnclaveIdent(v string)`

SetEnclaveIdent sets EnclaveIdent field to given value.

### HasEnclaveIdent

`func (o *JsonEnclaveProcess) HasEnclaveIdent() bool`

HasEnclaveIdent returns a boolean if a field has been set.

### GetSignerIdent

`func (o *JsonEnclaveProcess) GetSignerIdent() string`

GetSignerIdent returns the SignerIdent field if non-nil, zero value otherwise.

### GetSignerIdentOk

`func (o *JsonEnclaveProcess) GetSignerIdentOk() (*string, bool)`

GetSignerIdentOk returns a tuple with the SignerIdent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignerIdent

`func (o *JsonEnclaveProcess) SetSignerIdent(v string)`

SetSignerIdent sets SignerIdent field to given value.

### HasSignerIdent

`func (o *JsonEnclaveProcess) HasSignerIdent() bool`

HasSignerIdent returns a boolean if a field has been set.

### GetAttestationServer

`func (o *JsonEnclaveProcess) GetAttestationServer() string`

GetAttestationServer returns the AttestationServer field if non-nil, zero value otherwise.

### GetAttestationServerOk

`func (o *JsonEnclaveProcess) GetAttestationServerOk() (*string, bool)`

GetAttestationServerOk returns a tuple with the AttestationServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttestationServer

`func (o *JsonEnclaveProcess) SetAttestationServer(v string)`

SetAttestationServer sets AttestationServer field to given value.

### HasAttestationServer

`func (o *JsonEnclaveProcess) HasAttestationServer() bool`

HasAttestationServer returns a boolean if a field has been set.

### GetWgInterface

`func (o *JsonEnclaveProcess) GetWgInterface() JsonWireguardInterface`

GetWgInterface returns the WgInterface field if non-nil, zero value otherwise.

### GetWgInterfaceOk

`func (o *JsonEnclaveProcess) GetWgInterfaceOk() (*JsonWireguardInterface, bool)`

GetWgInterfaceOk returns a tuple with the WgInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWgInterface

`func (o *JsonEnclaveProcess) SetWgInterface(v JsonWireguardInterface)`

SetWgInterface sets WgInterface field to given value.

### HasWgInterface

`func (o *JsonEnclaveProcess) HasWgInterface() bool`

HasWgInterface returns a boolean if a field has been set.

### GetInternalAttesationServer

`func (o *JsonEnclaveProcess) GetInternalAttesationServer() string`

GetInternalAttesationServer returns the InternalAttesationServer field if non-nil, zero value otherwise.

### GetInternalAttesationServerOk

`func (o *JsonEnclaveProcess) GetInternalAttesationServerOk() (*string, bool)`

GetInternalAttesationServerOk returns a tuple with the InternalAttesationServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalAttesationServer

`func (o *JsonEnclaveProcess) SetInternalAttesationServer(v string)`

SetInternalAttesationServer sets InternalAttesationServer field to given value.

### HasInternalAttesationServer

`func (o *JsonEnclaveProcess) HasInternalAttesationServer() bool`

HasInternalAttesationServer returns a boolean if a field has been set.

### GetWireguardPublicKey

`func (o *JsonEnclaveProcess) GetWireguardPublicKey() string`

GetWireguardPublicKey returns the WireguardPublicKey field if non-nil, zero value otherwise.

### GetWireguardPublicKeyOk

`func (o *JsonEnclaveProcess) GetWireguardPublicKeyOk() (*string, bool)`

GetWireguardPublicKeyOk returns a tuple with the WireguardPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWireguardPublicKey

`func (o *JsonEnclaveProcess) SetWireguardPublicKey(v string)`

SetWireguardPublicKey sets WireguardPublicKey field to given value.

### HasWireguardPublicKey

`func (o *JsonEnclaveProcess) HasWireguardPublicKey() bool`

HasWireguardPublicKey returns a boolean if a field has been set.

### GetPublicIdent

`func (o *JsonEnclaveProcess) GetPublicIdent() string`

GetPublicIdent returns the PublicIdent field if non-nil, zero value otherwise.

### GetPublicIdentOk

`func (o *JsonEnclaveProcess) GetPublicIdentOk() (*string, bool)`

GetPublicIdentOk returns a tuple with the PublicIdent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIdent

`func (o *JsonEnclaveProcess) SetPublicIdent(v string)`

SetPublicIdent sets PublicIdent field to given value.

### HasPublicIdent

`func (o *JsonEnclaveProcess) HasPublicIdent() bool`

HasPublicIdent returns a boolean if a field has been set.

### GetInternalWireguardServer

`func (o *JsonEnclaveProcess) GetInternalWireguardServer() string`

GetInternalWireguardServer returns the InternalWireguardServer field if non-nil, zero value otherwise.

### GetInternalWireguardServerOk

`func (o *JsonEnclaveProcess) GetInternalWireguardServerOk() (*string, bool)`

GetInternalWireguardServerOk returns a tuple with the InternalWireguardServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalWireguardServer

`func (o *JsonEnclaveProcess) SetInternalWireguardServer(v string)`

SetInternalWireguardServer sets InternalWireguardServer field to given value.

### HasInternalWireguardServer

`func (o *JsonEnclaveProcess) HasInternalWireguardServer() bool`

HasInternalWireguardServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


