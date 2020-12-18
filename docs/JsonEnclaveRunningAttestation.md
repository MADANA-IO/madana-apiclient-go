# JsonEnclaveRunningAttestation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeInfo** | Pointer to [**JsonNodeInfo**](json_NodeInfo.md) |  | [optional] 
**EnclaveProcess** | Pointer to [**JsonEnclaveProcess**](json_EnclaveProcess.md) |  | [optional] 

## Methods

### NewJsonEnclaveRunningAttestation

`func NewJsonEnclaveRunningAttestation() *JsonEnclaveRunningAttestation`

NewJsonEnclaveRunningAttestation instantiates a new JsonEnclaveRunningAttestation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJsonEnclaveRunningAttestationWithDefaults

`func NewJsonEnclaveRunningAttestationWithDefaults() *JsonEnclaveRunningAttestation`

NewJsonEnclaveRunningAttestationWithDefaults instantiates a new JsonEnclaveRunningAttestation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeInfo

`func (o *JsonEnclaveRunningAttestation) GetNodeInfo() JsonNodeInfo`

GetNodeInfo returns the NodeInfo field if non-nil, zero value otherwise.

### GetNodeInfoOk

`func (o *JsonEnclaveRunningAttestation) GetNodeInfoOk() (*JsonNodeInfo, bool)`

GetNodeInfoOk returns a tuple with the NodeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeInfo

`func (o *JsonEnclaveRunningAttestation) SetNodeInfo(v JsonNodeInfo)`

SetNodeInfo sets NodeInfo field to given value.

### HasNodeInfo

`func (o *JsonEnclaveRunningAttestation) HasNodeInfo() bool`

HasNodeInfo returns a boolean if a field has been set.

### GetEnclaveProcess

`func (o *JsonEnclaveRunningAttestation) GetEnclaveProcess() JsonEnclaveProcess`

GetEnclaveProcess returns the EnclaveProcess field if non-nil, zero value otherwise.

### GetEnclaveProcessOk

`func (o *JsonEnclaveRunningAttestation) GetEnclaveProcessOk() (*JsonEnclaveProcess, bool)`

GetEnclaveProcessOk returns a tuple with the EnclaveProcess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnclaveProcess

`func (o *JsonEnclaveRunningAttestation) SetEnclaveProcess(v JsonEnclaveProcess)`

SetEnclaveProcess sets EnclaveProcess field to given value.

### HasEnclaveProcess

`func (o *JsonEnclaveRunningAttestation) HasEnclaveProcess() bool`

HasEnclaveProcess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


