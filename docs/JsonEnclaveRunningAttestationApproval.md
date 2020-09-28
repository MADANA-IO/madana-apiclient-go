# JsonEnclaveRunningAttestationApproval

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeInfo** | Pointer to [**JsonNodeInfo**](json_NodeInfo.md) |  | [optional] 
**EnclaveProcess** | Pointer to [**JsonEnclaveProcess**](json_EnclaveProcess.md) |  | [optional] 
**Approved** | Pointer to **string** |  | [optional] 

## Methods

### NewJsonEnclaveRunningAttestationApproval

`func NewJsonEnclaveRunningAttestationApproval() *JsonEnclaveRunningAttestationApproval`

NewJsonEnclaveRunningAttestationApproval instantiates a new JsonEnclaveRunningAttestationApproval object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJsonEnclaveRunningAttestationApprovalWithDefaults

`func NewJsonEnclaveRunningAttestationApprovalWithDefaults() *JsonEnclaveRunningAttestationApproval`

NewJsonEnclaveRunningAttestationApprovalWithDefaults instantiates a new JsonEnclaveRunningAttestationApproval object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeInfo

`func (o *JsonEnclaveRunningAttestationApproval) GetNodeInfo() JsonNodeInfo`

GetNodeInfo returns the NodeInfo field if non-nil, zero value otherwise.

### GetNodeInfoOk

`func (o *JsonEnclaveRunningAttestationApproval) GetNodeInfoOk() (*JsonNodeInfo, bool)`

GetNodeInfoOk returns a tuple with the NodeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeInfo

`func (o *JsonEnclaveRunningAttestationApproval) SetNodeInfo(v JsonNodeInfo)`

SetNodeInfo sets NodeInfo field to given value.

### HasNodeInfo

`func (o *JsonEnclaveRunningAttestationApproval) HasNodeInfo() bool`

HasNodeInfo returns a boolean if a field has been set.

### GetEnclaveProcess

`func (o *JsonEnclaveRunningAttestationApproval) GetEnclaveProcess() JsonEnclaveProcess`

GetEnclaveProcess returns the EnclaveProcess field if non-nil, zero value otherwise.

### GetEnclaveProcessOk

`func (o *JsonEnclaveRunningAttestationApproval) GetEnclaveProcessOk() (*JsonEnclaveProcess, bool)`

GetEnclaveProcessOk returns a tuple with the EnclaveProcess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnclaveProcess

`func (o *JsonEnclaveRunningAttestationApproval) SetEnclaveProcess(v JsonEnclaveProcess)`

SetEnclaveProcess sets EnclaveProcess field to given value.

### HasEnclaveProcess

`func (o *JsonEnclaveRunningAttestationApproval) HasEnclaveProcess() bool`

HasEnclaveProcess returns a boolean if a field has been set.

### GetApproved

`func (o *JsonEnclaveRunningAttestationApproval) GetApproved() string`

GetApproved returns the Approved field if non-nil, zero value otherwise.

### GetApprovedOk

`func (o *JsonEnclaveRunningAttestationApproval) GetApprovedOk() (*string, bool)`

GetApprovedOk returns a tuple with the Approved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproved

`func (o *JsonEnclaveRunningAttestationApproval) SetApproved(v string)`

SetApproved sets Approved field to given value.

### HasApproved

`func (o *JsonEnclaveRunningAttestationApproval) HasApproved() bool`

HasApproved returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


