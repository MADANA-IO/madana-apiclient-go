# JsonV1EventList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]JsonV1Event**](JsonV1Event.md) |  | [optional] 
**Metadata** | Pointer to [**JsonV1ListMeta**](json_V1ListMeta.md) |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**ApiVersion** | Pointer to **string** |  | [optional] 

## Methods

### NewJsonV1EventList

`func NewJsonV1EventList() *JsonV1EventList`

NewJsonV1EventList instantiates a new JsonV1EventList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJsonV1EventListWithDefaults

`func NewJsonV1EventListWithDefaults() *JsonV1EventList`

NewJsonV1EventListWithDefaults instantiates a new JsonV1EventList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *JsonV1EventList) GetItems() []JsonV1Event`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *JsonV1EventList) GetItemsOk() (*[]JsonV1Event, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *JsonV1EventList) SetItems(v []JsonV1Event)`

SetItems sets Items field to given value.

### HasItems

`func (o *JsonV1EventList) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetMetadata

`func (o *JsonV1EventList) GetMetadata() JsonV1ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *JsonV1EventList) GetMetadataOk() (*JsonV1ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *JsonV1EventList) SetMetadata(v JsonV1ListMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *JsonV1EventList) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetKind

`func (o *JsonV1EventList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *JsonV1EventList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *JsonV1EventList) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *JsonV1EventList) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetApiVersion

`func (o *JsonV1EventList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *JsonV1EventList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *JsonV1EventList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *JsonV1EventList) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


