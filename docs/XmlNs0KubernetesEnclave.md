# XmlNs0KubernetesEnclave

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
**IsUsingInitContainer** | Pointer to **bool** |  | [optional] 

## Methods

### NewXmlNs0KubernetesEnclave

`func NewXmlNs0KubernetesEnclave() *XmlNs0KubernetesEnclave`

NewXmlNs0KubernetesEnclave instantiates a new XmlNs0KubernetesEnclave object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewXmlNs0KubernetesEnclaveWithDefaults

`func NewXmlNs0KubernetesEnclaveWithDefaults() *XmlNs0KubernetesEnclave`

NewXmlNs0KubernetesEnclaveWithDefaults instantiates a new XmlNs0KubernetesEnclave object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttestationServer

`func (o *XmlNs0KubernetesEnclave) GetAttestationServer() string`

GetAttestationServer returns the AttestationServer field if non-nil, zero value otherwise.

### GetAttestationServerOk

`func (o *XmlNs0KubernetesEnclave) GetAttestationServerOk() (*string, bool)`

GetAttestationServerOk returns a tuple with the AttestationServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttestationServer

`func (o *XmlNs0KubernetesEnclave) SetAttestationServer(v string)`

SetAttestationServer sets AttestationServer field to given value.

### HasAttestationServer

`func (o *XmlNs0KubernetesEnclave) HasAttestationServer() bool`

HasAttestationServer returns a boolean if a field has been set.

### GetConsoleOutput

`func (o *XmlNs0KubernetesEnclave) GetConsoleOutput() string`

GetConsoleOutput returns the ConsoleOutput field if non-nil, zero value otherwise.

### GetConsoleOutputOk

`func (o *XmlNs0KubernetesEnclave) GetConsoleOutputOk() (*string, bool)`

GetConsoleOutputOk returns a tuple with the ConsoleOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleOutput

`func (o *XmlNs0KubernetesEnclave) SetConsoleOutput(v string)`

SetConsoleOutput sets ConsoleOutput field to given value.

### HasConsoleOutput

`func (o *XmlNs0KubernetesEnclave) HasConsoleOutput() bool`

HasConsoleOutput returns a boolean if a field has been set.

### GetEnclaveIdent

`func (o *XmlNs0KubernetesEnclave) GetEnclaveIdent() string`

GetEnclaveIdent returns the EnclaveIdent field if non-nil, zero value otherwise.

### GetEnclaveIdentOk

`func (o *XmlNs0KubernetesEnclave) GetEnclaveIdentOk() (*string, bool)`

GetEnclaveIdentOk returns a tuple with the EnclaveIdent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnclaveIdent

`func (o *XmlNs0KubernetesEnclave) SetEnclaveIdent(v string)`

SetEnclaveIdent sets EnclaveIdent field to given value.

### HasEnclaveIdent

`func (o *XmlNs0KubernetesEnclave) HasEnclaveIdent() bool`

HasEnclaveIdent returns a boolean if a field has been set.

### GetEnclaveInputstream

`func (o *XmlNs0KubernetesEnclave) GetEnclaveInputstream() XmlNs0InputStream`

GetEnclaveInputstream returns the EnclaveInputstream field if non-nil, zero value otherwise.

### GetEnclaveInputstreamOk

`func (o *XmlNs0KubernetesEnclave) GetEnclaveInputstreamOk() (*XmlNs0InputStream, bool)`

GetEnclaveInputstreamOk returns a tuple with the EnclaveInputstream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnclaveInputstream

`func (o *XmlNs0KubernetesEnclave) SetEnclaveInputstream(v XmlNs0InputStream)`

SetEnclaveInputstream sets EnclaveInputstream field to given value.

### HasEnclaveInputstream

`func (o *XmlNs0KubernetesEnclave) HasEnclaveInputstream() bool`

HasEnclaveInputstream returns a boolean if a field has been set.

### GetEndingTime

`func (o *XmlNs0KubernetesEnclave) GetEndingTime() string`

GetEndingTime returns the EndingTime field if non-nil, zero value otherwise.

### GetEndingTimeOk

`func (o *XmlNs0KubernetesEnclave) GetEndingTimeOk() (*string, bool)`

GetEndingTimeOk returns a tuple with the EndingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndingTime

`func (o *XmlNs0KubernetesEnclave) SetEndingTime(v string)`

SetEndingTime sets EndingTime field to given value.

### HasEndingTime

`func (o *XmlNs0KubernetesEnclave) HasEndingTime() bool`

HasEndingTime returns a boolean if a field has been set.

### GetEnvironment

`func (o *XmlNs0KubernetesEnclave) GetEnvironment() XmlNs0Environment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *XmlNs0KubernetesEnclave) GetEnvironmentOk() (*XmlNs0Environment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *XmlNs0KubernetesEnclave) SetEnvironment(v XmlNs0Environment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *XmlNs0KubernetesEnclave) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetInternalAttesationServer

`func (o *XmlNs0KubernetesEnclave) GetInternalAttesationServer() string`

GetInternalAttesationServer returns the InternalAttesationServer field if non-nil, zero value otherwise.

### GetInternalAttesationServerOk

`func (o *XmlNs0KubernetesEnclave) GetInternalAttesationServerOk() (*string, bool)`

GetInternalAttesationServerOk returns a tuple with the InternalAttesationServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalAttesationServer

`func (o *XmlNs0KubernetesEnclave) SetInternalAttesationServer(v string)`

SetInternalAttesationServer sets InternalAttesationServer field to given value.

### HasInternalAttesationServer

`func (o *XmlNs0KubernetesEnclave) HasInternalAttesationServer() bool`

HasInternalAttesationServer returns a boolean if a field has been set.

### GetInternalIdent

`func (o *XmlNs0KubernetesEnclave) GetInternalIdent() string`

GetInternalIdent returns the InternalIdent field if non-nil, zero value otherwise.

### GetInternalIdentOk

`func (o *XmlNs0KubernetesEnclave) GetInternalIdentOk() (*string, bool)`

GetInternalIdentOk returns a tuple with the InternalIdent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalIdent

`func (o *XmlNs0KubernetesEnclave) SetInternalIdent(v string)`

SetInternalIdent sets InternalIdent field to given value.

### HasInternalIdent

`func (o *XmlNs0KubernetesEnclave) HasInternalIdent() bool`

HasInternalIdent returns a boolean if a field has been set.

### GetInternalRemoteControlServer

`func (o *XmlNs0KubernetesEnclave) GetInternalRemoteControlServer() string`

GetInternalRemoteControlServer returns the InternalRemoteControlServer field if non-nil, zero value otherwise.

### GetInternalRemoteControlServerOk

`func (o *XmlNs0KubernetesEnclave) GetInternalRemoteControlServerOk() (*string, bool)`

GetInternalRemoteControlServerOk returns a tuple with the InternalRemoteControlServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalRemoteControlServer

`func (o *XmlNs0KubernetesEnclave) SetInternalRemoteControlServer(v string)`

SetInternalRemoteControlServer sets InternalRemoteControlServer field to given value.

### HasInternalRemoteControlServer

`func (o *XmlNs0KubernetesEnclave) HasInternalRemoteControlServer() bool`

HasInternalRemoteControlServer returns a boolean if a field has been set.

### GetInternalWireguardServer

`func (o *XmlNs0KubernetesEnclave) GetInternalWireguardServer() string`

GetInternalWireguardServer returns the InternalWireguardServer field if non-nil, zero value otherwise.

### GetInternalWireguardServerOk

`func (o *XmlNs0KubernetesEnclave) GetInternalWireguardServerOk() (*string, bool)`

GetInternalWireguardServerOk returns a tuple with the InternalWireguardServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalWireguardServer

`func (o *XmlNs0KubernetesEnclave) SetInternalWireguardServer(v string)`

SetInternalWireguardServer sets InternalWireguardServer field to given value.

### HasInternalWireguardServer

`func (o *XmlNs0KubernetesEnclave) HasInternalWireguardServer() bool`

HasInternalWireguardServer returns a boolean if a field has been set.

### GetKubernetesEnclave

`func (o *XmlNs0KubernetesEnclave) GetKubernetesEnclave() XmlNs0KubernetesEnclave`

GetKubernetesEnclave returns the KubernetesEnclave field if non-nil, zero value otherwise.

### GetKubernetesEnclaveOk

`func (o *XmlNs0KubernetesEnclave) GetKubernetesEnclaveOk() (*XmlNs0KubernetesEnclave, bool)`

GetKubernetesEnclaveOk returns a tuple with the KubernetesEnclave field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesEnclave

`func (o *XmlNs0KubernetesEnclave) SetKubernetesEnclave(v XmlNs0KubernetesEnclave)`

SetKubernetesEnclave sets KubernetesEnclave field to given value.

### HasKubernetesEnclave

`func (o *XmlNs0KubernetesEnclave) HasKubernetesEnclave() bool`

HasKubernetesEnclave returns a boolean if a field has been set.

### GetPortMapping

`func (o *XmlNs0KubernetesEnclave) GetPortMapping() map[string]interface{}`

GetPortMapping returns the PortMapping field if non-nil, zero value otherwise.

### GetPortMappingOk

`func (o *XmlNs0KubernetesEnclave) GetPortMappingOk() (*map[string]interface{}, bool)`

GetPortMappingOk returns a tuple with the PortMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortMapping

`func (o *XmlNs0KubernetesEnclave) SetPortMapping(v map[string]interface{})`

SetPortMapping sets PortMapping field to given value.

### HasPortMapping

`func (o *XmlNs0KubernetesEnclave) HasPortMapping() bool`

HasPortMapping returns a boolean if a field has been set.

### GetPorts

`func (o *XmlNs0KubernetesEnclave) GetPorts() []XmlNs0EnclavePort`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *XmlNs0KubernetesEnclave) GetPortsOk() (*[]XmlNs0EnclavePort, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *XmlNs0KubernetesEnclave) SetPorts(v []XmlNs0EnclavePort)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *XmlNs0KubernetesEnclave) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetProcess

`func (o *XmlNs0KubernetesEnclave) GetProcess() XmlNs0Process`

GetProcess returns the Process field if non-nil, zero value otherwise.

### GetProcessOk

`func (o *XmlNs0KubernetesEnclave) GetProcessOk() (*XmlNs0Process, bool)`

GetProcessOk returns a tuple with the Process field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcess

`func (o *XmlNs0KubernetesEnclave) SetProcess(v XmlNs0Process)`

SetProcess sets Process field to given value.

### HasProcess

`func (o *XmlNs0KubernetesEnclave) HasProcess() bool`

HasProcess returns a boolean if a field has been set.

### GetPublicIdent

`func (o *XmlNs0KubernetesEnclave) GetPublicIdent() string`

GetPublicIdent returns the PublicIdent field if non-nil, zero value otherwise.

### GetPublicIdentOk

`func (o *XmlNs0KubernetesEnclave) GetPublicIdentOk() (*string, bool)`

GetPublicIdentOk returns a tuple with the PublicIdent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIdent

`func (o *XmlNs0KubernetesEnclave) SetPublicIdent(v string)`

SetPublicIdent sets PublicIdent field to given value.

### HasPublicIdent

`func (o *XmlNs0KubernetesEnclave) HasPublicIdent() bool`

HasPublicIdent returns a boolean if a field has been set.

### GetRemoteControlServer

`func (o *XmlNs0KubernetesEnclave) GetRemoteControlServer() string`

GetRemoteControlServer returns the RemoteControlServer field if non-nil, zero value otherwise.

### GetRemoteControlServerOk

`func (o *XmlNs0KubernetesEnclave) GetRemoteControlServerOk() (*string, bool)`

GetRemoteControlServerOk returns a tuple with the RemoteControlServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteControlServer

`func (o *XmlNs0KubernetesEnclave) SetRemoteControlServer(v string)`

SetRemoteControlServer sets RemoteControlServer field to given value.

### HasRemoteControlServer

`func (o *XmlNs0KubernetesEnclave) HasRemoteControlServer() bool`

HasRemoteControlServer returns a boolean if a field has been set.

### GetSignerIdent

`func (o *XmlNs0KubernetesEnclave) GetSignerIdent() string`

GetSignerIdent returns the SignerIdent field if non-nil, zero value otherwise.

### GetSignerIdentOk

`func (o *XmlNs0KubernetesEnclave) GetSignerIdentOk() (*string, bool)`

GetSignerIdentOk returns a tuple with the SignerIdent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignerIdent

`func (o *XmlNs0KubernetesEnclave) SetSignerIdent(v string)`

SetSignerIdent sets SignerIdent field to given value.

### HasSignerIdent

`func (o *XmlNs0KubernetesEnclave) HasSignerIdent() bool`

HasSignerIdent returns a boolean if a field has been set.

### GetStartupCMD

`func (o *XmlNs0KubernetesEnclave) GetStartupCMD() string`

GetStartupCMD returns the StartupCMD field if non-nil, zero value otherwise.

### GetStartupCMDOk

`func (o *XmlNs0KubernetesEnclave) GetStartupCMDOk() (*string, bool)`

GetStartupCMDOk returns a tuple with the StartupCMD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartupCMD

`func (o *XmlNs0KubernetesEnclave) SetStartupCMD(v string)`

SetStartupCMD sets StartupCMD field to given value.

### HasStartupCMD

`func (o *XmlNs0KubernetesEnclave) HasStartupCMD() bool`

HasStartupCMD returns a boolean if a field has been set.

### GetStartupTime

`func (o *XmlNs0KubernetesEnclave) GetStartupTime() string`

GetStartupTime returns the StartupTime field if non-nil, zero value otherwise.

### GetStartupTimeOk

`func (o *XmlNs0KubernetesEnclave) GetStartupTimeOk() (*string, bool)`

GetStartupTimeOk returns a tuple with the StartupTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartupTime

`func (o *XmlNs0KubernetesEnclave) SetStartupTime(v string)`

SetStartupTime sets StartupTime field to given value.

### HasStartupTime

`func (o *XmlNs0KubernetesEnclave) HasStartupTime() bool`

HasStartupTime returns a boolean if a field has been set.

### GetStatus

`func (o *XmlNs0KubernetesEnclave) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *XmlNs0KubernetesEnclave) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *XmlNs0KubernetesEnclave) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *XmlNs0KubernetesEnclave) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetWgInterface

`func (o *XmlNs0KubernetesEnclave) GetWgInterface() XmlNs0WireguardInterface`

GetWgInterface returns the WgInterface field if non-nil, zero value otherwise.

### GetWgInterfaceOk

`func (o *XmlNs0KubernetesEnclave) GetWgInterfaceOk() (*XmlNs0WireguardInterface, bool)`

GetWgInterfaceOk returns a tuple with the WgInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWgInterface

`func (o *XmlNs0KubernetesEnclave) SetWgInterface(v XmlNs0WireguardInterface)`

SetWgInterface sets WgInterface field to given value.

### HasWgInterface

`func (o *XmlNs0KubernetesEnclave) HasWgInterface() bool`

HasWgInterface returns a boolean if a field has been set.

### GetWireguardPublicKey

`func (o *XmlNs0KubernetesEnclave) GetWireguardPublicKey() string`

GetWireguardPublicKey returns the WireguardPublicKey field if non-nil, zero value otherwise.

### GetWireguardPublicKeyOk

`func (o *XmlNs0KubernetesEnclave) GetWireguardPublicKeyOk() (*string, bool)`

GetWireguardPublicKeyOk returns a tuple with the WireguardPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWireguardPublicKey

`func (o *XmlNs0KubernetesEnclave) SetWireguardPublicKey(v string)`

SetWireguardPublicKey sets WireguardPublicKey field to given value.

### HasWireguardPublicKey

`func (o *XmlNs0KubernetesEnclave) HasWireguardPublicKey() bool`

HasWireguardPublicKey returns a boolean if a field has been set.

### GetWireguardServer

`func (o *XmlNs0KubernetesEnclave) GetWireguardServer() string`

GetWireguardServer returns the WireguardServer field if non-nil, zero value otherwise.

### GetWireguardServerOk

`func (o *XmlNs0KubernetesEnclave) GetWireguardServerOk() (*string, bool)`

GetWireguardServerOk returns a tuple with the WireguardServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWireguardServer

`func (o *XmlNs0KubernetesEnclave) SetWireguardServer(v string)`

SetWireguardServer sets WireguardServer field to given value.

### HasWireguardServer

`func (o *XmlNs0KubernetesEnclave) HasWireguardServer() bool`

HasWireguardServer returns a boolean if a field has been set.

### GetIsUsingInitContainer

`func (o *XmlNs0KubernetesEnclave) GetIsUsingInitContainer() bool`

GetIsUsingInitContainer returns the IsUsingInitContainer field if non-nil, zero value otherwise.

### GetIsUsingInitContainerOk

`func (o *XmlNs0KubernetesEnclave) GetIsUsingInitContainerOk() (*bool, bool)`

GetIsUsingInitContainerOk returns a tuple with the IsUsingInitContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUsingInitContainer

`func (o *XmlNs0KubernetesEnclave) SetIsUsingInitContainer(v bool)`

SetIsUsingInitContainer sets IsUsingInitContainer field to given value.

### HasIsUsingInitContainer

`func (o *XmlNs0KubernetesEnclave) HasIsUsingInitContainer() bool`

HasIsUsingInitContainer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


