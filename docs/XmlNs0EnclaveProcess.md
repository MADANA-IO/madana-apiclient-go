# XmlNs0EnclaveProcess

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
**KubernetesEnclave** | Pointer to [**XmlNs0KubernetesEnclave**](xml_ns0_kubernetesEnclave.md) |  | [optional] 
**PortMapping** | Pointer to **map[string]interface{}** |  | [optional] 
**Ports** | Pointer to [**[]XmlNs0EnclavePort**](XmlNs0EnclavePort.md) |  | [optional] 
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

### NewXmlNs0EnclaveProcess

`func NewXmlNs0EnclaveProcess() *XmlNs0EnclaveProcess`

NewXmlNs0EnclaveProcess instantiates a new XmlNs0EnclaveProcess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewXmlNs0EnclaveProcessWithDefaults

`func NewXmlNs0EnclaveProcessWithDefaults() *XmlNs0EnclaveProcess`

NewXmlNs0EnclaveProcessWithDefaults instantiates a new XmlNs0EnclaveProcess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttestationServer

`func (o *XmlNs0EnclaveProcess) GetAttestationServer() string`

GetAttestationServer returns the AttestationServer field if non-nil, zero value otherwise.

### GetAttestationServerOk

`func (o *XmlNs0EnclaveProcess) GetAttestationServerOk() (*string, bool)`

GetAttestationServerOk returns a tuple with the AttestationServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttestationServer

`func (o *XmlNs0EnclaveProcess) SetAttestationServer(v string)`

SetAttestationServer sets AttestationServer field to given value.

### HasAttestationServer

`func (o *XmlNs0EnclaveProcess) HasAttestationServer() bool`

HasAttestationServer returns a boolean if a field has been set.

### GetConsoleOutput

`func (o *XmlNs0EnclaveProcess) GetConsoleOutput() string`

GetConsoleOutput returns the ConsoleOutput field if non-nil, zero value otherwise.

### GetConsoleOutputOk

`func (o *XmlNs0EnclaveProcess) GetConsoleOutputOk() (*string, bool)`

GetConsoleOutputOk returns a tuple with the ConsoleOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleOutput

`func (o *XmlNs0EnclaveProcess) SetConsoleOutput(v string)`

SetConsoleOutput sets ConsoleOutput field to given value.

### HasConsoleOutput

`func (o *XmlNs0EnclaveProcess) HasConsoleOutput() bool`

HasConsoleOutput returns a boolean if a field has been set.

### GetEnclaveIdent

`func (o *XmlNs0EnclaveProcess) GetEnclaveIdent() string`

GetEnclaveIdent returns the EnclaveIdent field if non-nil, zero value otherwise.

### GetEnclaveIdentOk

`func (o *XmlNs0EnclaveProcess) GetEnclaveIdentOk() (*string, bool)`

GetEnclaveIdentOk returns a tuple with the EnclaveIdent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnclaveIdent

`func (o *XmlNs0EnclaveProcess) SetEnclaveIdent(v string)`

SetEnclaveIdent sets EnclaveIdent field to given value.

### HasEnclaveIdent

`func (o *XmlNs0EnclaveProcess) HasEnclaveIdent() bool`

HasEnclaveIdent returns a boolean if a field has been set.

### GetEnclaveInputstream

`func (o *XmlNs0EnclaveProcess) GetEnclaveInputstream() XmlNs0InputStream`

GetEnclaveInputstream returns the EnclaveInputstream field if non-nil, zero value otherwise.

### GetEnclaveInputstreamOk

`func (o *XmlNs0EnclaveProcess) GetEnclaveInputstreamOk() (*XmlNs0InputStream, bool)`

GetEnclaveInputstreamOk returns a tuple with the EnclaveInputstream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnclaveInputstream

`func (o *XmlNs0EnclaveProcess) SetEnclaveInputstream(v XmlNs0InputStream)`

SetEnclaveInputstream sets EnclaveInputstream field to given value.

### HasEnclaveInputstream

`func (o *XmlNs0EnclaveProcess) HasEnclaveInputstream() bool`

HasEnclaveInputstream returns a boolean if a field has been set.

### GetEndingTime

`func (o *XmlNs0EnclaveProcess) GetEndingTime() string`

GetEndingTime returns the EndingTime field if non-nil, zero value otherwise.

### GetEndingTimeOk

`func (o *XmlNs0EnclaveProcess) GetEndingTimeOk() (*string, bool)`

GetEndingTimeOk returns a tuple with the EndingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndingTime

`func (o *XmlNs0EnclaveProcess) SetEndingTime(v string)`

SetEndingTime sets EndingTime field to given value.

### HasEndingTime

`func (o *XmlNs0EnclaveProcess) HasEndingTime() bool`

HasEndingTime returns a boolean if a field has been set.

### GetEnvironment

`func (o *XmlNs0EnclaveProcess) GetEnvironment() XmlNs0Environment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *XmlNs0EnclaveProcess) GetEnvironmentOk() (*XmlNs0Environment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *XmlNs0EnclaveProcess) SetEnvironment(v XmlNs0Environment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *XmlNs0EnclaveProcess) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetInternalAttesationServer

`func (o *XmlNs0EnclaveProcess) GetInternalAttesationServer() string`

GetInternalAttesationServer returns the InternalAttesationServer field if non-nil, zero value otherwise.

### GetInternalAttesationServerOk

`func (o *XmlNs0EnclaveProcess) GetInternalAttesationServerOk() (*string, bool)`

GetInternalAttesationServerOk returns a tuple with the InternalAttesationServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalAttesationServer

`func (o *XmlNs0EnclaveProcess) SetInternalAttesationServer(v string)`

SetInternalAttesationServer sets InternalAttesationServer field to given value.

### HasInternalAttesationServer

`func (o *XmlNs0EnclaveProcess) HasInternalAttesationServer() bool`

HasInternalAttesationServer returns a boolean if a field has been set.

### GetInternalIdent

`func (o *XmlNs0EnclaveProcess) GetInternalIdent() string`

GetInternalIdent returns the InternalIdent field if non-nil, zero value otherwise.

### GetInternalIdentOk

`func (o *XmlNs0EnclaveProcess) GetInternalIdentOk() (*string, bool)`

GetInternalIdentOk returns a tuple with the InternalIdent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalIdent

`func (o *XmlNs0EnclaveProcess) SetInternalIdent(v string)`

SetInternalIdent sets InternalIdent field to given value.

### HasInternalIdent

`func (o *XmlNs0EnclaveProcess) HasInternalIdent() bool`

HasInternalIdent returns a boolean if a field has been set.

### GetInternalRemoteControlServer

`func (o *XmlNs0EnclaveProcess) GetInternalRemoteControlServer() string`

GetInternalRemoteControlServer returns the InternalRemoteControlServer field if non-nil, zero value otherwise.

### GetInternalRemoteControlServerOk

`func (o *XmlNs0EnclaveProcess) GetInternalRemoteControlServerOk() (*string, bool)`

GetInternalRemoteControlServerOk returns a tuple with the InternalRemoteControlServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalRemoteControlServer

`func (o *XmlNs0EnclaveProcess) SetInternalRemoteControlServer(v string)`

SetInternalRemoteControlServer sets InternalRemoteControlServer field to given value.

### HasInternalRemoteControlServer

`func (o *XmlNs0EnclaveProcess) HasInternalRemoteControlServer() bool`

HasInternalRemoteControlServer returns a boolean if a field has been set.

### GetInternalWireguardServer

`func (o *XmlNs0EnclaveProcess) GetInternalWireguardServer() string`

GetInternalWireguardServer returns the InternalWireguardServer field if non-nil, zero value otherwise.

### GetInternalWireguardServerOk

`func (o *XmlNs0EnclaveProcess) GetInternalWireguardServerOk() (*string, bool)`

GetInternalWireguardServerOk returns a tuple with the InternalWireguardServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalWireguardServer

`func (o *XmlNs0EnclaveProcess) SetInternalWireguardServer(v string)`

SetInternalWireguardServer sets InternalWireguardServer field to given value.

### HasInternalWireguardServer

`func (o *XmlNs0EnclaveProcess) HasInternalWireguardServer() bool`

HasInternalWireguardServer returns a boolean if a field has been set.

### GetKubernetesEnclave

`func (o *XmlNs0EnclaveProcess) GetKubernetesEnclave() XmlNs0KubernetesEnclave`

GetKubernetesEnclave returns the KubernetesEnclave field if non-nil, zero value otherwise.

### GetKubernetesEnclaveOk

`func (o *XmlNs0EnclaveProcess) GetKubernetesEnclaveOk() (*XmlNs0KubernetesEnclave, bool)`

GetKubernetesEnclaveOk returns a tuple with the KubernetesEnclave field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesEnclave

`func (o *XmlNs0EnclaveProcess) SetKubernetesEnclave(v XmlNs0KubernetesEnclave)`

SetKubernetesEnclave sets KubernetesEnclave field to given value.

### HasKubernetesEnclave

`func (o *XmlNs0EnclaveProcess) HasKubernetesEnclave() bool`

HasKubernetesEnclave returns a boolean if a field has been set.

### GetPortMapping

`func (o *XmlNs0EnclaveProcess) GetPortMapping() map[string]interface{}`

GetPortMapping returns the PortMapping field if non-nil, zero value otherwise.

### GetPortMappingOk

`func (o *XmlNs0EnclaveProcess) GetPortMappingOk() (*map[string]interface{}, bool)`

GetPortMappingOk returns a tuple with the PortMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortMapping

`func (o *XmlNs0EnclaveProcess) SetPortMapping(v map[string]interface{})`

SetPortMapping sets PortMapping field to given value.

### HasPortMapping

`func (o *XmlNs0EnclaveProcess) HasPortMapping() bool`

HasPortMapping returns a boolean if a field has been set.

### GetPorts

`func (o *XmlNs0EnclaveProcess) GetPorts() []XmlNs0EnclavePort`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *XmlNs0EnclaveProcess) GetPortsOk() (*[]XmlNs0EnclavePort, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *XmlNs0EnclaveProcess) SetPorts(v []XmlNs0EnclavePort)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *XmlNs0EnclaveProcess) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetProcess

`func (o *XmlNs0EnclaveProcess) GetProcess() XmlNs0Process`

GetProcess returns the Process field if non-nil, zero value otherwise.

### GetProcessOk

`func (o *XmlNs0EnclaveProcess) GetProcessOk() (*XmlNs0Process, bool)`

GetProcessOk returns a tuple with the Process field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcess

`func (o *XmlNs0EnclaveProcess) SetProcess(v XmlNs0Process)`

SetProcess sets Process field to given value.

### HasProcess

`func (o *XmlNs0EnclaveProcess) HasProcess() bool`

HasProcess returns a boolean if a field has been set.

### GetPublicIdent

`func (o *XmlNs0EnclaveProcess) GetPublicIdent() string`

GetPublicIdent returns the PublicIdent field if non-nil, zero value otherwise.

### GetPublicIdentOk

`func (o *XmlNs0EnclaveProcess) GetPublicIdentOk() (*string, bool)`

GetPublicIdentOk returns a tuple with the PublicIdent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIdent

`func (o *XmlNs0EnclaveProcess) SetPublicIdent(v string)`

SetPublicIdent sets PublicIdent field to given value.

### HasPublicIdent

`func (o *XmlNs0EnclaveProcess) HasPublicIdent() bool`

HasPublicIdent returns a boolean if a field has been set.

### GetRemoteControlServer

`func (o *XmlNs0EnclaveProcess) GetRemoteControlServer() string`

GetRemoteControlServer returns the RemoteControlServer field if non-nil, zero value otherwise.

### GetRemoteControlServerOk

`func (o *XmlNs0EnclaveProcess) GetRemoteControlServerOk() (*string, bool)`

GetRemoteControlServerOk returns a tuple with the RemoteControlServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteControlServer

`func (o *XmlNs0EnclaveProcess) SetRemoteControlServer(v string)`

SetRemoteControlServer sets RemoteControlServer field to given value.

### HasRemoteControlServer

`func (o *XmlNs0EnclaveProcess) HasRemoteControlServer() bool`

HasRemoteControlServer returns a boolean if a field has been set.

### GetSignerIdent

`func (o *XmlNs0EnclaveProcess) GetSignerIdent() string`

GetSignerIdent returns the SignerIdent field if non-nil, zero value otherwise.

### GetSignerIdentOk

`func (o *XmlNs0EnclaveProcess) GetSignerIdentOk() (*string, bool)`

GetSignerIdentOk returns a tuple with the SignerIdent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignerIdent

`func (o *XmlNs0EnclaveProcess) SetSignerIdent(v string)`

SetSignerIdent sets SignerIdent field to given value.

### HasSignerIdent

`func (o *XmlNs0EnclaveProcess) HasSignerIdent() bool`

HasSignerIdent returns a boolean if a field has been set.

### GetStartupCMD

`func (o *XmlNs0EnclaveProcess) GetStartupCMD() string`

GetStartupCMD returns the StartupCMD field if non-nil, zero value otherwise.

### GetStartupCMDOk

`func (o *XmlNs0EnclaveProcess) GetStartupCMDOk() (*string, bool)`

GetStartupCMDOk returns a tuple with the StartupCMD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartupCMD

`func (o *XmlNs0EnclaveProcess) SetStartupCMD(v string)`

SetStartupCMD sets StartupCMD field to given value.

### HasStartupCMD

`func (o *XmlNs0EnclaveProcess) HasStartupCMD() bool`

HasStartupCMD returns a boolean if a field has been set.

### GetStartupTime

`func (o *XmlNs0EnclaveProcess) GetStartupTime() string`

GetStartupTime returns the StartupTime field if non-nil, zero value otherwise.

### GetStartupTimeOk

`func (o *XmlNs0EnclaveProcess) GetStartupTimeOk() (*string, bool)`

GetStartupTimeOk returns a tuple with the StartupTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartupTime

`func (o *XmlNs0EnclaveProcess) SetStartupTime(v string)`

SetStartupTime sets StartupTime field to given value.

### HasStartupTime

`func (o *XmlNs0EnclaveProcess) HasStartupTime() bool`

HasStartupTime returns a boolean if a field has been set.

### GetStatus

`func (o *XmlNs0EnclaveProcess) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *XmlNs0EnclaveProcess) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *XmlNs0EnclaveProcess) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *XmlNs0EnclaveProcess) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetWgInterface

`func (o *XmlNs0EnclaveProcess) GetWgInterface() XmlNs0WireguardInterface`

GetWgInterface returns the WgInterface field if non-nil, zero value otherwise.

### GetWgInterfaceOk

`func (o *XmlNs0EnclaveProcess) GetWgInterfaceOk() (*XmlNs0WireguardInterface, bool)`

GetWgInterfaceOk returns a tuple with the WgInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWgInterface

`func (o *XmlNs0EnclaveProcess) SetWgInterface(v XmlNs0WireguardInterface)`

SetWgInterface sets WgInterface field to given value.

### HasWgInterface

`func (o *XmlNs0EnclaveProcess) HasWgInterface() bool`

HasWgInterface returns a boolean if a field has been set.

### GetWireguardPublicKey

`func (o *XmlNs0EnclaveProcess) GetWireguardPublicKey() string`

GetWireguardPublicKey returns the WireguardPublicKey field if non-nil, zero value otherwise.

### GetWireguardPublicKeyOk

`func (o *XmlNs0EnclaveProcess) GetWireguardPublicKeyOk() (*string, bool)`

GetWireguardPublicKeyOk returns a tuple with the WireguardPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWireguardPublicKey

`func (o *XmlNs0EnclaveProcess) SetWireguardPublicKey(v string)`

SetWireguardPublicKey sets WireguardPublicKey field to given value.

### HasWireguardPublicKey

`func (o *XmlNs0EnclaveProcess) HasWireguardPublicKey() bool`

HasWireguardPublicKey returns a boolean if a field has been set.

### GetWireguardServer

`func (o *XmlNs0EnclaveProcess) GetWireguardServer() string`

GetWireguardServer returns the WireguardServer field if non-nil, zero value otherwise.

### GetWireguardServerOk

`func (o *XmlNs0EnclaveProcess) GetWireguardServerOk() (*string, bool)`

GetWireguardServerOk returns a tuple with the WireguardServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWireguardServer

`func (o *XmlNs0EnclaveProcess) SetWireguardServer(v string)`

SetWireguardServer sets WireguardServer field to given value.

### HasWireguardServer

`func (o *XmlNs0EnclaveProcess) HasWireguardServer() bool`

HasWireguardServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


