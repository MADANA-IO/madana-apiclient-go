# XmlNs0EnclaveProcessAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttestationServer** | Pointer to **string** |  | [optional] 
**ConsoleOutput** | Pointer to **string** |  | [optional] 
**EnclaveIdent** | Pointer to **string** |  | [optional] 
**EnclaveInputstream** | Pointer to [**XmlNs0InputStream**](xml_ns0_inputStream.md) |  | [optional] 
**EndingTime** | Pointer to **string** |  | [optional] 
**Environment** | Pointer to [**XmlNs0Environment**](xml_ns0_environment.md) |  | [optional] 
**InternalAttesationServer** | Pointer to **string** |  | [optional] 
**InternalIdent** | Pointer to **string** |  | [optional] 
**InternalRemoteControlServer** | Pointer to **string** |  | [optional] 
**InternalWireguardServer** | Pointer to **string** |  | [optional] 
**PortMapping** | Pointer to **map[string]interface{}** |  | [optional] 
**Ports** | Pointer to [**[]XmlNs0EnclavePort**](xml_ns0_enclavePort.md) |  | [optional] 
**Process** | Pointer to [**XmlNs0Process**](xml_ns0_process.md) |  | [optional] 
**PublicIdent** | Pointer to **string** |  | [optional] 
**RemoteControlServer** | Pointer to **string** |  | [optional] 
**SignerIdent** | Pointer to **string** |  | [optional] 
**StartupCMD** | Pointer to **string** |  | [optional] 
**StartupTime** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**WgInterface** | Pointer to [**XmlNs0WireguardInterface**](xml_ns0_wireguardInterface.md) |  | [optional] 
**WireguardPublicKey** | Pointer to **string** |  | [optional] 
**WireguardServer** | Pointer to **string** |  | [optional] 

## Methods

### NewXmlNs0EnclaveProcessAllOf

`func NewXmlNs0EnclaveProcessAllOf() *XmlNs0EnclaveProcessAllOf`

NewXmlNs0EnclaveProcessAllOf instantiates a new XmlNs0EnclaveProcessAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewXmlNs0EnclaveProcessAllOfWithDefaults

`func NewXmlNs0EnclaveProcessAllOfWithDefaults() *XmlNs0EnclaveProcessAllOf`

NewXmlNs0EnclaveProcessAllOfWithDefaults instantiates a new XmlNs0EnclaveProcessAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttestationServer

`func (o *XmlNs0EnclaveProcessAllOf) GetAttestationServer() string`

GetAttestationServer returns the AttestationServer field if non-nil, zero value otherwise.

### GetAttestationServerOk

`func (o *XmlNs0EnclaveProcessAllOf) GetAttestationServerOk() (*string, bool)`

GetAttestationServerOk returns a tuple with the AttestationServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttestationServer

`func (o *XmlNs0EnclaveProcessAllOf) SetAttestationServer(v string)`

SetAttestationServer sets AttestationServer field to given value.

### HasAttestationServer

`func (o *XmlNs0EnclaveProcessAllOf) HasAttestationServer() bool`

HasAttestationServer returns a boolean if a field has been set.

### GetConsoleOutput

`func (o *XmlNs0EnclaveProcessAllOf) GetConsoleOutput() string`

GetConsoleOutput returns the ConsoleOutput field if non-nil, zero value otherwise.

### GetConsoleOutputOk

`func (o *XmlNs0EnclaveProcessAllOf) GetConsoleOutputOk() (*string, bool)`

GetConsoleOutputOk returns a tuple with the ConsoleOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleOutput

`func (o *XmlNs0EnclaveProcessAllOf) SetConsoleOutput(v string)`

SetConsoleOutput sets ConsoleOutput field to given value.

### HasConsoleOutput

`func (o *XmlNs0EnclaveProcessAllOf) HasConsoleOutput() bool`

HasConsoleOutput returns a boolean if a field has been set.

### GetEnclaveIdent

`func (o *XmlNs0EnclaveProcessAllOf) GetEnclaveIdent() string`

GetEnclaveIdent returns the EnclaveIdent field if non-nil, zero value otherwise.

### GetEnclaveIdentOk

`func (o *XmlNs0EnclaveProcessAllOf) GetEnclaveIdentOk() (*string, bool)`

GetEnclaveIdentOk returns a tuple with the EnclaveIdent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnclaveIdent

`func (o *XmlNs0EnclaveProcessAllOf) SetEnclaveIdent(v string)`

SetEnclaveIdent sets EnclaveIdent field to given value.

### HasEnclaveIdent

`func (o *XmlNs0EnclaveProcessAllOf) HasEnclaveIdent() bool`

HasEnclaveIdent returns a boolean if a field has been set.

### GetEnclaveInputstream

`func (o *XmlNs0EnclaveProcessAllOf) GetEnclaveInputstream() XmlNs0InputStream`

GetEnclaveInputstream returns the EnclaveInputstream field if non-nil, zero value otherwise.

### GetEnclaveInputstreamOk

`func (o *XmlNs0EnclaveProcessAllOf) GetEnclaveInputstreamOk() (*XmlNs0InputStream, bool)`

GetEnclaveInputstreamOk returns a tuple with the EnclaveInputstream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnclaveInputstream

`func (o *XmlNs0EnclaveProcessAllOf) SetEnclaveInputstream(v XmlNs0InputStream)`

SetEnclaveInputstream sets EnclaveInputstream field to given value.

### HasEnclaveInputstream

`func (o *XmlNs0EnclaveProcessAllOf) HasEnclaveInputstream() bool`

HasEnclaveInputstream returns a boolean if a field has been set.

### GetEndingTime

`func (o *XmlNs0EnclaveProcessAllOf) GetEndingTime() string`

GetEndingTime returns the EndingTime field if non-nil, zero value otherwise.

### GetEndingTimeOk

`func (o *XmlNs0EnclaveProcessAllOf) GetEndingTimeOk() (*string, bool)`

GetEndingTimeOk returns a tuple with the EndingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndingTime

`func (o *XmlNs0EnclaveProcessAllOf) SetEndingTime(v string)`

SetEndingTime sets EndingTime field to given value.

### HasEndingTime

`func (o *XmlNs0EnclaveProcessAllOf) HasEndingTime() bool`

HasEndingTime returns a boolean if a field has been set.

### GetEnvironment

`func (o *XmlNs0EnclaveProcessAllOf) GetEnvironment() XmlNs0Environment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *XmlNs0EnclaveProcessAllOf) GetEnvironmentOk() (*XmlNs0Environment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *XmlNs0EnclaveProcessAllOf) SetEnvironment(v XmlNs0Environment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *XmlNs0EnclaveProcessAllOf) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetInternalAttesationServer

`func (o *XmlNs0EnclaveProcessAllOf) GetInternalAttesationServer() string`

GetInternalAttesationServer returns the InternalAttesationServer field if non-nil, zero value otherwise.

### GetInternalAttesationServerOk

`func (o *XmlNs0EnclaveProcessAllOf) GetInternalAttesationServerOk() (*string, bool)`

GetInternalAttesationServerOk returns a tuple with the InternalAttesationServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalAttesationServer

`func (o *XmlNs0EnclaveProcessAllOf) SetInternalAttesationServer(v string)`

SetInternalAttesationServer sets InternalAttesationServer field to given value.

### HasInternalAttesationServer

`func (o *XmlNs0EnclaveProcessAllOf) HasInternalAttesationServer() bool`

HasInternalAttesationServer returns a boolean if a field has been set.

### GetInternalIdent

`func (o *XmlNs0EnclaveProcessAllOf) GetInternalIdent() string`

GetInternalIdent returns the InternalIdent field if non-nil, zero value otherwise.

### GetInternalIdentOk

`func (o *XmlNs0EnclaveProcessAllOf) GetInternalIdentOk() (*string, bool)`

GetInternalIdentOk returns a tuple with the InternalIdent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalIdent

`func (o *XmlNs0EnclaveProcessAllOf) SetInternalIdent(v string)`

SetInternalIdent sets InternalIdent field to given value.

### HasInternalIdent

`func (o *XmlNs0EnclaveProcessAllOf) HasInternalIdent() bool`

HasInternalIdent returns a boolean if a field has been set.

### GetInternalRemoteControlServer

`func (o *XmlNs0EnclaveProcessAllOf) GetInternalRemoteControlServer() string`

GetInternalRemoteControlServer returns the InternalRemoteControlServer field if non-nil, zero value otherwise.

### GetInternalRemoteControlServerOk

`func (o *XmlNs0EnclaveProcessAllOf) GetInternalRemoteControlServerOk() (*string, bool)`

GetInternalRemoteControlServerOk returns a tuple with the InternalRemoteControlServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalRemoteControlServer

`func (o *XmlNs0EnclaveProcessAllOf) SetInternalRemoteControlServer(v string)`

SetInternalRemoteControlServer sets InternalRemoteControlServer field to given value.

### HasInternalRemoteControlServer

`func (o *XmlNs0EnclaveProcessAllOf) HasInternalRemoteControlServer() bool`

HasInternalRemoteControlServer returns a boolean if a field has been set.

### GetInternalWireguardServer

`func (o *XmlNs0EnclaveProcessAllOf) GetInternalWireguardServer() string`

GetInternalWireguardServer returns the InternalWireguardServer field if non-nil, zero value otherwise.

### GetInternalWireguardServerOk

`func (o *XmlNs0EnclaveProcessAllOf) GetInternalWireguardServerOk() (*string, bool)`

GetInternalWireguardServerOk returns a tuple with the InternalWireguardServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalWireguardServer

`func (o *XmlNs0EnclaveProcessAllOf) SetInternalWireguardServer(v string)`

SetInternalWireguardServer sets InternalWireguardServer field to given value.

### HasInternalWireguardServer

`func (o *XmlNs0EnclaveProcessAllOf) HasInternalWireguardServer() bool`

HasInternalWireguardServer returns a boolean if a field has been set.

### GetPortMapping

`func (o *XmlNs0EnclaveProcessAllOf) GetPortMapping() map[string]interface{}`

GetPortMapping returns the PortMapping field if non-nil, zero value otherwise.

### GetPortMappingOk

`func (o *XmlNs0EnclaveProcessAllOf) GetPortMappingOk() (*map[string]interface{}, bool)`

GetPortMappingOk returns a tuple with the PortMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortMapping

`func (o *XmlNs0EnclaveProcessAllOf) SetPortMapping(v map[string]interface{})`

SetPortMapping sets PortMapping field to given value.

### HasPortMapping

`func (o *XmlNs0EnclaveProcessAllOf) HasPortMapping() bool`

HasPortMapping returns a boolean if a field has been set.

### GetPorts

`func (o *XmlNs0EnclaveProcessAllOf) GetPorts() []XmlNs0EnclavePort`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *XmlNs0EnclaveProcessAllOf) GetPortsOk() (*[]XmlNs0EnclavePort, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *XmlNs0EnclaveProcessAllOf) SetPorts(v []XmlNs0EnclavePort)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *XmlNs0EnclaveProcessAllOf) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetProcess

`func (o *XmlNs0EnclaveProcessAllOf) GetProcess() XmlNs0Process`

GetProcess returns the Process field if non-nil, zero value otherwise.

### GetProcessOk

`func (o *XmlNs0EnclaveProcessAllOf) GetProcessOk() (*XmlNs0Process, bool)`

GetProcessOk returns a tuple with the Process field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcess

`func (o *XmlNs0EnclaveProcessAllOf) SetProcess(v XmlNs0Process)`

SetProcess sets Process field to given value.

### HasProcess

`func (o *XmlNs0EnclaveProcessAllOf) HasProcess() bool`

HasProcess returns a boolean if a field has been set.

### GetPublicIdent

`func (o *XmlNs0EnclaveProcessAllOf) GetPublicIdent() string`

GetPublicIdent returns the PublicIdent field if non-nil, zero value otherwise.

### GetPublicIdentOk

`func (o *XmlNs0EnclaveProcessAllOf) GetPublicIdentOk() (*string, bool)`

GetPublicIdentOk returns a tuple with the PublicIdent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIdent

`func (o *XmlNs0EnclaveProcessAllOf) SetPublicIdent(v string)`

SetPublicIdent sets PublicIdent field to given value.

### HasPublicIdent

`func (o *XmlNs0EnclaveProcessAllOf) HasPublicIdent() bool`

HasPublicIdent returns a boolean if a field has been set.

### GetRemoteControlServer

`func (o *XmlNs0EnclaveProcessAllOf) GetRemoteControlServer() string`

GetRemoteControlServer returns the RemoteControlServer field if non-nil, zero value otherwise.

### GetRemoteControlServerOk

`func (o *XmlNs0EnclaveProcessAllOf) GetRemoteControlServerOk() (*string, bool)`

GetRemoteControlServerOk returns a tuple with the RemoteControlServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteControlServer

`func (o *XmlNs0EnclaveProcessAllOf) SetRemoteControlServer(v string)`

SetRemoteControlServer sets RemoteControlServer field to given value.

### HasRemoteControlServer

`func (o *XmlNs0EnclaveProcessAllOf) HasRemoteControlServer() bool`

HasRemoteControlServer returns a boolean if a field has been set.

### GetSignerIdent

`func (o *XmlNs0EnclaveProcessAllOf) GetSignerIdent() string`

GetSignerIdent returns the SignerIdent field if non-nil, zero value otherwise.

### GetSignerIdentOk

`func (o *XmlNs0EnclaveProcessAllOf) GetSignerIdentOk() (*string, bool)`

GetSignerIdentOk returns a tuple with the SignerIdent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignerIdent

`func (o *XmlNs0EnclaveProcessAllOf) SetSignerIdent(v string)`

SetSignerIdent sets SignerIdent field to given value.

### HasSignerIdent

`func (o *XmlNs0EnclaveProcessAllOf) HasSignerIdent() bool`

HasSignerIdent returns a boolean if a field has been set.

### GetStartupCMD

`func (o *XmlNs0EnclaveProcessAllOf) GetStartupCMD() string`

GetStartupCMD returns the StartupCMD field if non-nil, zero value otherwise.

### GetStartupCMDOk

`func (o *XmlNs0EnclaveProcessAllOf) GetStartupCMDOk() (*string, bool)`

GetStartupCMDOk returns a tuple with the StartupCMD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartupCMD

`func (o *XmlNs0EnclaveProcessAllOf) SetStartupCMD(v string)`

SetStartupCMD sets StartupCMD field to given value.

### HasStartupCMD

`func (o *XmlNs0EnclaveProcessAllOf) HasStartupCMD() bool`

HasStartupCMD returns a boolean if a field has been set.

### GetStartupTime

`func (o *XmlNs0EnclaveProcessAllOf) GetStartupTime() string`

GetStartupTime returns the StartupTime field if non-nil, zero value otherwise.

### GetStartupTimeOk

`func (o *XmlNs0EnclaveProcessAllOf) GetStartupTimeOk() (*string, bool)`

GetStartupTimeOk returns a tuple with the StartupTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartupTime

`func (o *XmlNs0EnclaveProcessAllOf) SetStartupTime(v string)`

SetStartupTime sets StartupTime field to given value.

### HasStartupTime

`func (o *XmlNs0EnclaveProcessAllOf) HasStartupTime() bool`

HasStartupTime returns a boolean if a field has been set.

### GetStatus

`func (o *XmlNs0EnclaveProcessAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *XmlNs0EnclaveProcessAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *XmlNs0EnclaveProcessAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *XmlNs0EnclaveProcessAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetWgInterface

`func (o *XmlNs0EnclaveProcessAllOf) GetWgInterface() XmlNs0WireguardInterface`

GetWgInterface returns the WgInterface field if non-nil, zero value otherwise.

### GetWgInterfaceOk

`func (o *XmlNs0EnclaveProcessAllOf) GetWgInterfaceOk() (*XmlNs0WireguardInterface, bool)`

GetWgInterfaceOk returns a tuple with the WgInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWgInterface

`func (o *XmlNs0EnclaveProcessAllOf) SetWgInterface(v XmlNs0WireguardInterface)`

SetWgInterface sets WgInterface field to given value.

### HasWgInterface

`func (o *XmlNs0EnclaveProcessAllOf) HasWgInterface() bool`

HasWgInterface returns a boolean if a field has been set.

### GetWireguardPublicKey

`func (o *XmlNs0EnclaveProcessAllOf) GetWireguardPublicKey() string`

GetWireguardPublicKey returns the WireguardPublicKey field if non-nil, zero value otherwise.

### GetWireguardPublicKeyOk

`func (o *XmlNs0EnclaveProcessAllOf) GetWireguardPublicKeyOk() (*string, bool)`

GetWireguardPublicKeyOk returns a tuple with the WireguardPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWireguardPublicKey

`func (o *XmlNs0EnclaveProcessAllOf) SetWireguardPublicKey(v string)`

SetWireguardPublicKey sets WireguardPublicKey field to given value.

### HasWireguardPublicKey

`func (o *XmlNs0EnclaveProcessAllOf) HasWireguardPublicKey() bool`

HasWireguardPublicKey returns a boolean if a field has been set.

### GetWireguardServer

`func (o *XmlNs0EnclaveProcessAllOf) GetWireguardServer() string`

GetWireguardServer returns the WireguardServer field if non-nil, zero value otherwise.

### GetWireguardServerOk

`func (o *XmlNs0EnclaveProcessAllOf) GetWireguardServerOk() (*string, bool)`

GetWireguardServerOk returns a tuple with the WireguardServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWireguardServer

`func (o *XmlNs0EnclaveProcessAllOf) SetWireguardServer(v string)`

SetWireguardServer sets WireguardServer field to given value.

### HasWireguardServer

`func (o *XmlNs0EnclaveProcessAllOf) HasWireguardServer() bool`

HasWireguardServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


