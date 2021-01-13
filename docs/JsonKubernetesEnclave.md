# JsonKubernetesEnclave

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PublicIdent** | Pointer to **string** |  | [optional] 
**RemoteControlServer** | Pointer to **string** |  | [optional] 
**Environment** | Pointer to [**JsonEnvironment**](json_Environment.md) |  | [optional] 
**EndingTime** | Pointer to **string** |  | [optional] 
**AttestationServer** | Pointer to **string** |  | [optional] 
**WgInterface** | Pointer to [**JsonWireguardInterface**](json_WireguardInterface.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**KubernetesEnclave** | Pointer to [**JsonKubernetesEnclave**](json_KubernetesEnclave.md) |  | [optional] 
**PortMapping** | Pointer to **map[string]string** |  | [optional] 
**EnclaveIdent** | Pointer to **string** |  | [optional] 
**StartupCMD** | Pointer to **string** |  | [optional] 
**InternalWireguardServer** | Pointer to **string** |  | [optional] 
**SignerIdent** | Pointer to **string** |  | [optional] 
**ConsoleOutput** | Pointer to **string** |  | [optional] 
**InternalRemoteControlServer** | Pointer to **string** |  | [optional] 
**InternalAttesationServer** | Pointer to **string** |  | [optional] 
**WireguardServer** | Pointer to **string** |  | [optional] 
**Process** | Pointer to [**JsonProcess**](json_Process.md) |  | [optional] 
**InternalIdent** | Pointer to **string** |  | [optional] 
**WireguardPublicKey** | Pointer to **string** |  | [optional] 
**Ports** | Pointer to [**[]JsonEnclavePort**](JsonEnclavePort.md) |  | [optional] 
**StartupTime** | Pointer to **string** |  | [optional] 
**EnclaveInputstream** | Pointer to **map[string]interface{}** |  | [optional] 
**EnclaveReplicaSetEvents** | Pointer to [**JsonV1EventList**](json_V1EventList.md) |  | [optional] 
**WireguardPort** | Pointer to **int32** |  | [optional] 
**IsUsingInitContainer** | Pointer to **bool** |  | [optional] 
**EnclaveDeploymentEvents** | Pointer to [**JsonV1EventList**](json_V1EventList.md) |  | [optional] 
**PodPhase** | Pointer to **string** |  | [optional] 
**EnclavePodEvents** | Pointer to [**JsonV1EventList**](json_V1EventList.md) |  | [optional] 
**AttestationPort** | Pointer to **int32** |  | [optional] 
**DebugInfo** | Pointer to **string** |  | [optional] 
**RemoteControlIP** | Pointer to **string** |  | [optional] 

## Methods

### NewJsonKubernetesEnclave

`func NewJsonKubernetesEnclave() *JsonKubernetesEnclave`

NewJsonKubernetesEnclave instantiates a new JsonKubernetesEnclave object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJsonKubernetesEnclaveWithDefaults

`func NewJsonKubernetesEnclaveWithDefaults() *JsonKubernetesEnclave`

NewJsonKubernetesEnclaveWithDefaults instantiates a new JsonKubernetesEnclave object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublicIdent

`func (o *JsonKubernetesEnclave) GetPublicIdent() string`

GetPublicIdent returns the PublicIdent field if non-nil, zero value otherwise.

### GetPublicIdentOk

`func (o *JsonKubernetesEnclave) GetPublicIdentOk() (*string, bool)`

GetPublicIdentOk returns a tuple with the PublicIdent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIdent

`func (o *JsonKubernetesEnclave) SetPublicIdent(v string)`

SetPublicIdent sets PublicIdent field to given value.

### HasPublicIdent

`func (o *JsonKubernetesEnclave) HasPublicIdent() bool`

HasPublicIdent returns a boolean if a field has been set.

### GetRemoteControlServer

`func (o *JsonKubernetesEnclave) GetRemoteControlServer() string`

GetRemoteControlServer returns the RemoteControlServer field if non-nil, zero value otherwise.

### GetRemoteControlServerOk

`func (o *JsonKubernetesEnclave) GetRemoteControlServerOk() (*string, bool)`

GetRemoteControlServerOk returns a tuple with the RemoteControlServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteControlServer

`func (o *JsonKubernetesEnclave) SetRemoteControlServer(v string)`

SetRemoteControlServer sets RemoteControlServer field to given value.

### HasRemoteControlServer

`func (o *JsonKubernetesEnclave) HasRemoteControlServer() bool`

HasRemoteControlServer returns a boolean if a field has been set.

### GetEnvironment

`func (o *JsonKubernetesEnclave) GetEnvironment() JsonEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *JsonKubernetesEnclave) GetEnvironmentOk() (*JsonEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *JsonKubernetesEnclave) SetEnvironment(v JsonEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *JsonKubernetesEnclave) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetEndingTime

`func (o *JsonKubernetesEnclave) GetEndingTime() string`

GetEndingTime returns the EndingTime field if non-nil, zero value otherwise.

### GetEndingTimeOk

`func (o *JsonKubernetesEnclave) GetEndingTimeOk() (*string, bool)`

GetEndingTimeOk returns a tuple with the EndingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndingTime

`func (o *JsonKubernetesEnclave) SetEndingTime(v string)`

SetEndingTime sets EndingTime field to given value.

### HasEndingTime

`func (o *JsonKubernetesEnclave) HasEndingTime() bool`

HasEndingTime returns a boolean if a field has been set.

### GetAttestationServer

`func (o *JsonKubernetesEnclave) GetAttestationServer() string`

GetAttestationServer returns the AttestationServer field if non-nil, zero value otherwise.

### GetAttestationServerOk

`func (o *JsonKubernetesEnclave) GetAttestationServerOk() (*string, bool)`

GetAttestationServerOk returns a tuple with the AttestationServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttestationServer

`func (o *JsonKubernetesEnclave) SetAttestationServer(v string)`

SetAttestationServer sets AttestationServer field to given value.

### HasAttestationServer

`func (o *JsonKubernetesEnclave) HasAttestationServer() bool`

HasAttestationServer returns a boolean if a field has been set.

### GetWgInterface

`func (o *JsonKubernetesEnclave) GetWgInterface() JsonWireguardInterface`

GetWgInterface returns the WgInterface field if non-nil, zero value otherwise.

### GetWgInterfaceOk

`func (o *JsonKubernetesEnclave) GetWgInterfaceOk() (*JsonWireguardInterface, bool)`

GetWgInterfaceOk returns a tuple with the WgInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWgInterface

`func (o *JsonKubernetesEnclave) SetWgInterface(v JsonWireguardInterface)`

SetWgInterface sets WgInterface field to given value.

### HasWgInterface

`func (o *JsonKubernetesEnclave) HasWgInterface() bool`

HasWgInterface returns a boolean if a field has been set.

### GetStatus

`func (o *JsonKubernetesEnclave) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *JsonKubernetesEnclave) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *JsonKubernetesEnclave) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *JsonKubernetesEnclave) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetKubernetesEnclave

`func (o *JsonKubernetesEnclave) GetKubernetesEnclave() JsonKubernetesEnclave`

GetKubernetesEnclave returns the KubernetesEnclave field if non-nil, zero value otherwise.

### GetKubernetesEnclaveOk

`func (o *JsonKubernetesEnclave) GetKubernetesEnclaveOk() (*JsonKubernetesEnclave, bool)`

GetKubernetesEnclaveOk returns a tuple with the KubernetesEnclave field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesEnclave

`func (o *JsonKubernetesEnclave) SetKubernetesEnclave(v JsonKubernetesEnclave)`

SetKubernetesEnclave sets KubernetesEnclave field to given value.

### HasKubernetesEnclave

`func (o *JsonKubernetesEnclave) HasKubernetesEnclave() bool`

HasKubernetesEnclave returns a boolean if a field has been set.

### GetPortMapping

`func (o *JsonKubernetesEnclave) GetPortMapping() map[string]string`

GetPortMapping returns the PortMapping field if non-nil, zero value otherwise.

### GetPortMappingOk

`func (o *JsonKubernetesEnclave) GetPortMappingOk() (*map[string]string, bool)`

GetPortMappingOk returns a tuple with the PortMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortMapping

`func (o *JsonKubernetesEnclave) SetPortMapping(v map[string]string)`

SetPortMapping sets PortMapping field to given value.

### HasPortMapping

`func (o *JsonKubernetesEnclave) HasPortMapping() bool`

HasPortMapping returns a boolean if a field has been set.

### GetEnclaveIdent

`func (o *JsonKubernetesEnclave) GetEnclaveIdent() string`

GetEnclaveIdent returns the EnclaveIdent field if non-nil, zero value otherwise.

### GetEnclaveIdentOk

`func (o *JsonKubernetesEnclave) GetEnclaveIdentOk() (*string, bool)`

GetEnclaveIdentOk returns a tuple with the EnclaveIdent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnclaveIdent

`func (o *JsonKubernetesEnclave) SetEnclaveIdent(v string)`

SetEnclaveIdent sets EnclaveIdent field to given value.

### HasEnclaveIdent

`func (o *JsonKubernetesEnclave) HasEnclaveIdent() bool`

HasEnclaveIdent returns a boolean if a field has been set.

### GetStartupCMD

`func (o *JsonKubernetesEnclave) GetStartupCMD() string`

GetStartupCMD returns the StartupCMD field if non-nil, zero value otherwise.

### GetStartupCMDOk

`func (o *JsonKubernetesEnclave) GetStartupCMDOk() (*string, bool)`

GetStartupCMDOk returns a tuple with the StartupCMD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartupCMD

`func (o *JsonKubernetesEnclave) SetStartupCMD(v string)`

SetStartupCMD sets StartupCMD field to given value.

### HasStartupCMD

`func (o *JsonKubernetesEnclave) HasStartupCMD() bool`

HasStartupCMD returns a boolean if a field has been set.

### GetInternalWireguardServer

`func (o *JsonKubernetesEnclave) GetInternalWireguardServer() string`

GetInternalWireguardServer returns the InternalWireguardServer field if non-nil, zero value otherwise.

### GetInternalWireguardServerOk

`func (o *JsonKubernetesEnclave) GetInternalWireguardServerOk() (*string, bool)`

GetInternalWireguardServerOk returns a tuple with the InternalWireguardServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalWireguardServer

`func (o *JsonKubernetesEnclave) SetInternalWireguardServer(v string)`

SetInternalWireguardServer sets InternalWireguardServer field to given value.

### HasInternalWireguardServer

`func (o *JsonKubernetesEnclave) HasInternalWireguardServer() bool`

HasInternalWireguardServer returns a boolean if a field has been set.

### GetSignerIdent

`func (o *JsonKubernetesEnclave) GetSignerIdent() string`

GetSignerIdent returns the SignerIdent field if non-nil, zero value otherwise.

### GetSignerIdentOk

`func (o *JsonKubernetesEnclave) GetSignerIdentOk() (*string, bool)`

GetSignerIdentOk returns a tuple with the SignerIdent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignerIdent

`func (o *JsonKubernetesEnclave) SetSignerIdent(v string)`

SetSignerIdent sets SignerIdent field to given value.

### HasSignerIdent

`func (o *JsonKubernetesEnclave) HasSignerIdent() bool`

HasSignerIdent returns a boolean if a field has been set.

### GetConsoleOutput

`func (o *JsonKubernetesEnclave) GetConsoleOutput() string`

GetConsoleOutput returns the ConsoleOutput field if non-nil, zero value otherwise.

### GetConsoleOutputOk

`func (o *JsonKubernetesEnclave) GetConsoleOutputOk() (*string, bool)`

GetConsoleOutputOk returns a tuple with the ConsoleOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleOutput

`func (o *JsonKubernetesEnclave) SetConsoleOutput(v string)`

SetConsoleOutput sets ConsoleOutput field to given value.

### HasConsoleOutput

`func (o *JsonKubernetesEnclave) HasConsoleOutput() bool`

HasConsoleOutput returns a boolean if a field has been set.

### GetInternalRemoteControlServer

`func (o *JsonKubernetesEnclave) GetInternalRemoteControlServer() string`

GetInternalRemoteControlServer returns the InternalRemoteControlServer field if non-nil, zero value otherwise.

### GetInternalRemoteControlServerOk

`func (o *JsonKubernetesEnclave) GetInternalRemoteControlServerOk() (*string, bool)`

GetInternalRemoteControlServerOk returns a tuple with the InternalRemoteControlServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalRemoteControlServer

`func (o *JsonKubernetesEnclave) SetInternalRemoteControlServer(v string)`

SetInternalRemoteControlServer sets InternalRemoteControlServer field to given value.

### HasInternalRemoteControlServer

`func (o *JsonKubernetesEnclave) HasInternalRemoteControlServer() bool`

HasInternalRemoteControlServer returns a boolean if a field has been set.

### GetInternalAttesationServer

`func (o *JsonKubernetesEnclave) GetInternalAttesationServer() string`

GetInternalAttesationServer returns the InternalAttesationServer field if non-nil, zero value otherwise.

### GetInternalAttesationServerOk

`func (o *JsonKubernetesEnclave) GetInternalAttesationServerOk() (*string, bool)`

GetInternalAttesationServerOk returns a tuple with the InternalAttesationServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalAttesationServer

`func (o *JsonKubernetesEnclave) SetInternalAttesationServer(v string)`

SetInternalAttesationServer sets InternalAttesationServer field to given value.

### HasInternalAttesationServer

`func (o *JsonKubernetesEnclave) HasInternalAttesationServer() bool`

HasInternalAttesationServer returns a boolean if a field has been set.

### GetWireguardServer

`func (o *JsonKubernetesEnclave) GetWireguardServer() string`

GetWireguardServer returns the WireguardServer field if non-nil, zero value otherwise.

### GetWireguardServerOk

`func (o *JsonKubernetesEnclave) GetWireguardServerOk() (*string, bool)`

GetWireguardServerOk returns a tuple with the WireguardServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWireguardServer

`func (o *JsonKubernetesEnclave) SetWireguardServer(v string)`

SetWireguardServer sets WireguardServer field to given value.

### HasWireguardServer

`func (o *JsonKubernetesEnclave) HasWireguardServer() bool`

HasWireguardServer returns a boolean if a field has been set.

### GetProcess

`func (o *JsonKubernetesEnclave) GetProcess() JsonProcess`

GetProcess returns the Process field if non-nil, zero value otherwise.

### GetProcessOk

`func (o *JsonKubernetesEnclave) GetProcessOk() (*JsonProcess, bool)`

GetProcessOk returns a tuple with the Process field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcess

`func (o *JsonKubernetesEnclave) SetProcess(v JsonProcess)`

SetProcess sets Process field to given value.

### HasProcess

`func (o *JsonKubernetesEnclave) HasProcess() bool`

HasProcess returns a boolean if a field has been set.

### GetInternalIdent

`func (o *JsonKubernetesEnclave) GetInternalIdent() string`

GetInternalIdent returns the InternalIdent field if non-nil, zero value otherwise.

### GetInternalIdentOk

`func (o *JsonKubernetesEnclave) GetInternalIdentOk() (*string, bool)`

GetInternalIdentOk returns a tuple with the InternalIdent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalIdent

`func (o *JsonKubernetesEnclave) SetInternalIdent(v string)`

SetInternalIdent sets InternalIdent field to given value.

### HasInternalIdent

`func (o *JsonKubernetesEnclave) HasInternalIdent() bool`

HasInternalIdent returns a boolean if a field has been set.

### GetWireguardPublicKey

`func (o *JsonKubernetesEnclave) GetWireguardPublicKey() string`

GetWireguardPublicKey returns the WireguardPublicKey field if non-nil, zero value otherwise.

### GetWireguardPublicKeyOk

`func (o *JsonKubernetesEnclave) GetWireguardPublicKeyOk() (*string, bool)`

GetWireguardPublicKeyOk returns a tuple with the WireguardPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWireguardPublicKey

`func (o *JsonKubernetesEnclave) SetWireguardPublicKey(v string)`

SetWireguardPublicKey sets WireguardPublicKey field to given value.

### HasWireguardPublicKey

`func (o *JsonKubernetesEnclave) HasWireguardPublicKey() bool`

HasWireguardPublicKey returns a boolean if a field has been set.

### GetPorts

`func (o *JsonKubernetesEnclave) GetPorts() []JsonEnclavePort`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *JsonKubernetesEnclave) GetPortsOk() (*[]JsonEnclavePort, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *JsonKubernetesEnclave) SetPorts(v []JsonEnclavePort)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *JsonKubernetesEnclave) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetStartupTime

`func (o *JsonKubernetesEnclave) GetStartupTime() string`

GetStartupTime returns the StartupTime field if non-nil, zero value otherwise.

### GetStartupTimeOk

`func (o *JsonKubernetesEnclave) GetStartupTimeOk() (*string, bool)`

GetStartupTimeOk returns a tuple with the StartupTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartupTime

`func (o *JsonKubernetesEnclave) SetStartupTime(v string)`

SetStartupTime sets StartupTime field to given value.

### HasStartupTime

`func (o *JsonKubernetesEnclave) HasStartupTime() bool`

HasStartupTime returns a boolean if a field has been set.

### GetEnclaveInputstream

`func (o *JsonKubernetesEnclave) GetEnclaveInputstream() map[string]interface{}`

GetEnclaveInputstream returns the EnclaveInputstream field if non-nil, zero value otherwise.

### GetEnclaveInputstreamOk

`func (o *JsonKubernetesEnclave) GetEnclaveInputstreamOk() (*map[string]interface{}, bool)`

GetEnclaveInputstreamOk returns a tuple with the EnclaveInputstream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnclaveInputstream

`func (o *JsonKubernetesEnclave) SetEnclaveInputstream(v map[string]interface{})`

SetEnclaveInputstream sets EnclaveInputstream field to given value.

### HasEnclaveInputstream

`func (o *JsonKubernetesEnclave) HasEnclaveInputstream() bool`

HasEnclaveInputstream returns a boolean if a field has been set.

### GetEnclaveReplicaSetEvents

`func (o *JsonKubernetesEnclave) GetEnclaveReplicaSetEvents() JsonV1EventList`

GetEnclaveReplicaSetEvents returns the EnclaveReplicaSetEvents field if non-nil, zero value otherwise.

### GetEnclaveReplicaSetEventsOk

`func (o *JsonKubernetesEnclave) GetEnclaveReplicaSetEventsOk() (*JsonV1EventList, bool)`

GetEnclaveReplicaSetEventsOk returns a tuple with the EnclaveReplicaSetEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnclaveReplicaSetEvents

`func (o *JsonKubernetesEnclave) SetEnclaveReplicaSetEvents(v JsonV1EventList)`

SetEnclaveReplicaSetEvents sets EnclaveReplicaSetEvents field to given value.

### HasEnclaveReplicaSetEvents

`func (o *JsonKubernetesEnclave) HasEnclaveReplicaSetEvents() bool`

HasEnclaveReplicaSetEvents returns a boolean if a field has been set.

### GetWireguardPort

`func (o *JsonKubernetesEnclave) GetWireguardPort() int32`

GetWireguardPort returns the WireguardPort field if non-nil, zero value otherwise.

### GetWireguardPortOk

`func (o *JsonKubernetesEnclave) GetWireguardPortOk() (*int32, bool)`

GetWireguardPortOk returns a tuple with the WireguardPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWireguardPort

`func (o *JsonKubernetesEnclave) SetWireguardPort(v int32)`

SetWireguardPort sets WireguardPort field to given value.

### HasWireguardPort

`func (o *JsonKubernetesEnclave) HasWireguardPort() bool`

HasWireguardPort returns a boolean if a field has been set.

### GetIsUsingInitContainer

`func (o *JsonKubernetesEnclave) GetIsUsingInitContainer() bool`

GetIsUsingInitContainer returns the IsUsingInitContainer field if non-nil, zero value otherwise.

### GetIsUsingInitContainerOk

`func (o *JsonKubernetesEnclave) GetIsUsingInitContainerOk() (*bool, bool)`

GetIsUsingInitContainerOk returns a tuple with the IsUsingInitContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUsingInitContainer

`func (o *JsonKubernetesEnclave) SetIsUsingInitContainer(v bool)`

SetIsUsingInitContainer sets IsUsingInitContainer field to given value.

### HasIsUsingInitContainer

`func (o *JsonKubernetesEnclave) HasIsUsingInitContainer() bool`

HasIsUsingInitContainer returns a boolean if a field has been set.

### GetEnclaveDeploymentEvents

`func (o *JsonKubernetesEnclave) GetEnclaveDeploymentEvents() JsonV1EventList`

GetEnclaveDeploymentEvents returns the EnclaveDeploymentEvents field if non-nil, zero value otherwise.

### GetEnclaveDeploymentEventsOk

`func (o *JsonKubernetesEnclave) GetEnclaveDeploymentEventsOk() (*JsonV1EventList, bool)`

GetEnclaveDeploymentEventsOk returns a tuple with the EnclaveDeploymentEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnclaveDeploymentEvents

`func (o *JsonKubernetesEnclave) SetEnclaveDeploymentEvents(v JsonV1EventList)`

SetEnclaveDeploymentEvents sets EnclaveDeploymentEvents field to given value.

### HasEnclaveDeploymentEvents

`func (o *JsonKubernetesEnclave) HasEnclaveDeploymentEvents() bool`

HasEnclaveDeploymentEvents returns a boolean if a field has been set.

### GetPodPhase

`func (o *JsonKubernetesEnclave) GetPodPhase() string`

GetPodPhase returns the PodPhase field if non-nil, zero value otherwise.

### GetPodPhaseOk

`func (o *JsonKubernetesEnclave) GetPodPhaseOk() (*string, bool)`

GetPodPhaseOk returns a tuple with the PodPhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodPhase

`func (o *JsonKubernetesEnclave) SetPodPhase(v string)`

SetPodPhase sets PodPhase field to given value.

### HasPodPhase

`func (o *JsonKubernetesEnclave) HasPodPhase() bool`

HasPodPhase returns a boolean if a field has been set.

### GetEnclavePodEvents

`func (o *JsonKubernetesEnclave) GetEnclavePodEvents() JsonV1EventList`

GetEnclavePodEvents returns the EnclavePodEvents field if non-nil, zero value otherwise.

### GetEnclavePodEventsOk

`func (o *JsonKubernetesEnclave) GetEnclavePodEventsOk() (*JsonV1EventList, bool)`

GetEnclavePodEventsOk returns a tuple with the EnclavePodEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnclavePodEvents

`func (o *JsonKubernetesEnclave) SetEnclavePodEvents(v JsonV1EventList)`

SetEnclavePodEvents sets EnclavePodEvents field to given value.

### HasEnclavePodEvents

`func (o *JsonKubernetesEnclave) HasEnclavePodEvents() bool`

HasEnclavePodEvents returns a boolean if a field has been set.

### GetAttestationPort

`func (o *JsonKubernetesEnclave) GetAttestationPort() int32`

GetAttestationPort returns the AttestationPort field if non-nil, zero value otherwise.

### GetAttestationPortOk

`func (o *JsonKubernetesEnclave) GetAttestationPortOk() (*int32, bool)`

GetAttestationPortOk returns a tuple with the AttestationPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttestationPort

`func (o *JsonKubernetesEnclave) SetAttestationPort(v int32)`

SetAttestationPort sets AttestationPort field to given value.

### HasAttestationPort

`func (o *JsonKubernetesEnclave) HasAttestationPort() bool`

HasAttestationPort returns a boolean if a field has been set.

### GetDebugInfo

`func (o *JsonKubernetesEnclave) GetDebugInfo() string`

GetDebugInfo returns the DebugInfo field if non-nil, zero value otherwise.

### GetDebugInfoOk

`func (o *JsonKubernetesEnclave) GetDebugInfoOk() (*string, bool)`

GetDebugInfoOk returns a tuple with the DebugInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugInfo

`func (o *JsonKubernetesEnclave) SetDebugInfo(v string)`

SetDebugInfo sets DebugInfo field to given value.

### HasDebugInfo

`func (o *JsonKubernetesEnclave) HasDebugInfo() bool`

HasDebugInfo returns a boolean if a field has been set.

### GetRemoteControlIP

`func (o *JsonKubernetesEnclave) GetRemoteControlIP() string`

GetRemoteControlIP returns the RemoteControlIP field if non-nil, zero value otherwise.

### GetRemoteControlIPOk

`func (o *JsonKubernetesEnclave) GetRemoteControlIPOk() (*string, bool)`

GetRemoteControlIPOk returns a tuple with the RemoteControlIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteControlIP

`func (o *JsonKubernetesEnclave) SetRemoteControlIP(v string)`

SetRemoteControlIP sets RemoteControlIP field to given value.

### HasRemoteControlIP

`func (o *JsonKubernetesEnclave) HasRemoteControlIP() bool`

HasRemoteControlIP returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


