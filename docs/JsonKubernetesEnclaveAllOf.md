# JsonKubernetesEnclaveAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DebugInfo** | Pointer to **string** |  | [optional] 
**AttestationPort** | Pointer to **int32** |  | [optional] 
**PodPhase** | Pointer to **string** |  | [optional] 
**IsUsingInitContainer** | Pointer to **bool** |  | [optional] 
**EnclavePodEvents** | Pointer to [**JsonV1EventList**](json_V1EventList.md) |  | [optional] 
**RemoteControlIP** | Pointer to **string** |  | [optional] 
**EnclaveReplicaSetEvents** | Pointer to [**JsonV1EventList**](json_V1EventList.md) |  | [optional] 
**EnclaveDeploymentEvents** | Pointer to [**JsonV1EventList**](json_V1EventList.md) |  | [optional] 
**WireguardPort** | Pointer to **int32** |  | [optional] 

## Methods

### NewJsonKubernetesEnclaveAllOf

`func NewJsonKubernetesEnclaveAllOf() *JsonKubernetesEnclaveAllOf`

NewJsonKubernetesEnclaveAllOf instantiates a new JsonKubernetesEnclaveAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJsonKubernetesEnclaveAllOfWithDefaults

`func NewJsonKubernetesEnclaveAllOfWithDefaults() *JsonKubernetesEnclaveAllOf`

NewJsonKubernetesEnclaveAllOfWithDefaults instantiates a new JsonKubernetesEnclaveAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDebugInfo

`func (o *JsonKubernetesEnclaveAllOf) GetDebugInfo() string`

GetDebugInfo returns the DebugInfo field if non-nil, zero value otherwise.

### GetDebugInfoOk

`func (o *JsonKubernetesEnclaveAllOf) GetDebugInfoOk() (*string, bool)`

GetDebugInfoOk returns a tuple with the DebugInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugInfo

`func (o *JsonKubernetesEnclaveAllOf) SetDebugInfo(v string)`

SetDebugInfo sets DebugInfo field to given value.

### HasDebugInfo

`func (o *JsonKubernetesEnclaveAllOf) HasDebugInfo() bool`

HasDebugInfo returns a boolean if a field has been set.

### GetAttestationPort

`func (o *JsonKubernetesEnclaveAllOf) GetAttestationPort() int32`

GetAttestationPort returns the AttestationPort field if non-nil, zero value otherwise.

### GetAttestationPortOk

`func (o *JsonKubernetesEnclaveAllOf) GetAttestationPortOk() (*int32, bool)`

GetAttestationPortOk returns a tuple with the AttestationPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttestationPort

`func (o *JsonKubernetesEnclaveAllOf) SetAttestationPort(v int32)`

SetAttestationPort sets AttestationPort field to given value.

### HasAttestationPort

`func (o *JsonKubernetesEnclaveAllOf) HasAttestationPort() bool`

HasAttestationPort returns a boolean if a field has been set.

### GetPodPhase

`func (o *JsonKubernetesEnclaveAllOf) GetPodPhase() string`

GetPodPhase returns the PodPhase field if non-nil, zero value otherwise.

### GetPodPhaseOk

`func (o *JsonKubernetesEnclaveAllOf) GetPodPhaseOk() (*string, bool)`

GetPodPhaseOk returns a tuple with the PodPhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodPhase

`func (o *JsonKubernetesEnclaveAllOf) SetPodPhase(v string)`

SetPodPhase sets PodPhase field to given value.

### HasPodPhase

`func (o *JsonKubernetesEnclaveAllOf) HasPodPhase() bool`

HasPodPhase returns a boolean if a field has been set.

### GetIsUsingInitContainer

`func (o *JsonKubernetesEnclaveAllOf) GetIsUsingInitContainer() bool`

GetIsUsingInitContainer returns the IsUsingInitContainer field if non-nil, zero value otherwise.

### GetIsUsingInitContainerOk

`func (o *JsonKubernetesEnclaveAllOf) GetIsUsingInitContainerOk() (*bool, bool)`

GetIsUsingInitContainerOk returns a tuple with the IsUsingInitContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUsingInitContainer

`func (o *JsonKubernetesEnclaveAllOf) SetIsUsingInitContainer(v bool)`

SetIsUsingInitContainer sets IsUsingInitContainer field to given value.

### HasIsUsingInitContainer

`func (o *JsonKubernetesEnclaveAllOf) HasIsUsingInitContainer() bool`

HasIsUsingInitContainer returns a boolean if a field has been set.

### GetEnclavePodEvents

`func (o *JsonKubernetesEnclaveAllOf) GetEnclavePodEvents() JsonV1EventList`

GetEnclavePodEvents returns the EnclavePodEvents field if non-nil, zero value otherwise.

### GetEnclavePodEventsOk

`func (o *JsonKubernetesEnclaveAllOf) GetEnclavePodEventsOk() (*JsonV1EventList, bool)`

GetEnclavePodEventsOk returns a tuple with the EnclavePodEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnclavePodEvents

`func (o *JsonKubernetesEnclaveAllOf) SetEnclavePodEvents(v JsonV1EventList)`

SetEnclavePodEvents sets EnclavePodEvents field to given value.

### HasEnclavePodEvents

`func (o *JsonKubernetesEnclaveAllOf) HasEnclavePodEvents() bool`

HasEnclavePodEvents returns a boolean if a field has been set.

### GetRemoteControlIP

`func (o *JsonKubernetesEnclaveAllOf) GetRemoteControlIP() string`

GetRemoteControlIP returns the RemoteControlIP field if non-nil, zero value otherwise.

### GetRemoteControlIPOk

`func (o *JsonKubernetesEnclaveAllOf) GetRemoteControlIPOk() (*string, bool)`

GetRemoteControlIPOk returns a tuple with the RemoteControlIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteControlIP

`func (o *JsonKubernetesEnclaveAllOf) SetRemoteControlIP(v string)`

SetRemoteControlIP sets RemoteControlIP field to given value.

### HasRemoteControlIP

`func (o *JsonKubernetesEnclaveAllOf) HasRemoteControlIP() bool`

HasRemoteControlIP returns a boolean if a field has been set.

### GetEnclaveReplicaSetEvents

`func (o *JsonKubernetesEnclaveAllOf) GetEnclaveReplicaSetEvents() JsonV1EventList`

GetEnclaveReplicaSetEvents returns the EnclaveReplicaSetEvents field if non-nil, zero value otherwise.

### GetEnclaveReplicaSetEventsOk

`func (o *JsonKubernetesEnclaveAllOf) GetEnclaveReplicaSetEventsOk() (*JsonV1EventList, bool)`

GetEnclaveReplicaSetEventsOk returns a tuple with the EnclaveReplicaSetEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnclaveReplicaSetEvents

`func (o *JsonKubernetesEnclaveAllOf) SetEnclaveReplicaSetEvents(v JsonV1EventList)`

SetEnclaveReplicaSetEvents sets EnclaveReplicaSetEvents field to given value.

### HasEnclaveReplicaSetEvents

`func (o *JsonKubernetesEnclaveAllOf) HasEnclaveReplicaSetEvents() bool`

HasEnclaveReplicaSetEvents returns a boolean if a field has been set.

### GetEnclaveDeploymentEvents

`func (o *JsonKubernetesEnclaveAllOf) GetEnclaveDeploymentEvents() JsonV1EventList`

GetEnclaveDeploymentEvents returns the EnclaveDeploymentEvents field if non-nil, zero value otherwise.

### GetEnclaveDeploymentEventsOk

`func (o *JsonKubernetesEnclaveAllOf) GetEnclaveDeploymentEventsOk() (*JsonV1EventList, bool)`

GetEnclaveDeploymentEventsOk returns a tuple with the EnclaveDeploymentEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnclaveDeploymentEvents

`func (o *JsonKubernetesEnclaveAllOf) SetEnclaveDeploymentEvents(v JsonV1EventList)`

SetEnclaveDeploymentEvents sets EnclaveDeploymentEvents field to given value.

### HasEnclaveDeploymentEvents

`func (o *JsonKubernetesEnclaveAllOf) HasEnclaveDeploymentEvents() bool`

HasEnclaveDeploymentEvents returns a boolean if a field has been set.

### GetWireguardPort

`func (o *JsonKubernetesEnclaveAllOf) GetWireguardPort() int32`

GetWireguardPort returns the WireguardPort field if non-nil, zero value otherwise.

### GetWireguardPortOk

`func (o *JsonKubernetesEnclaveAllOf) GetWireguardPortOk() (*int32, bool)`

GetWireguardPortOk returns a tuple with the WireguardPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWireguardPort

`func (o *JsonKubernetesEnclaveAllOf) SetWireguardPort(v int32)`

SetWireguardPort sets WireguardPort field to given value.

### HasWireguardPort

`func (o *JsonKubernetesEnclaveAllOf) HasWireguardPort() bool`

HasWireguardPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


