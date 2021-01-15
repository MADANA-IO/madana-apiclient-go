# JsonV1EventSeries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** |  | [optional] 
**LastObservedTime** | Pointer to **float32** |  | [optional] 
**Count** | Pointer to **float32** |  | [optional] 

## Methods

### NewJsonV1EventSeries

`func NewJsonV1EventSeries() *JsonV1EventSeries`

NewJsonV1EventSeries instantiates a new JsonV1EventSeries object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJsonV1EventSeriesWithDefaults

`func NewJsonV1EventSeriesWithDefaults() *JsonV1EventSeries`

NewJsonV1EventSeriesWithDefaults instantiates a new JsonV1EventSeries object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *JsonV1EventSeries) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *JsonV1EventSeries) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *JsonV1EventSeries) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *JsonV1EventSeries) HasState() bool`

HasState returns a boolean if a field has been set.

### GetLastObservedTime

`func (o *JsonV1EventSeries) GetLastObservedTime() float32`

GetLastObservedTime returns the LastObservedTime field if non-nil, zero value otherwise.

### GetLastObservedTimeOk

`func (o *JsonV1EventSeries) GetLastObservedTimeOk() (*float32, bool)`

GetLastObservedTimeOk returns a tuple with the LastObservedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastObservedTime

`func (o *JsonV1EventSeries) SetLastObservedTime(v float32)`

SetLastObservedTime sets LastObservedTime field to given value.

### HasLastObservedTime

`func (o *JsonV1EventSeries) HasLastObservedTime() bool`

HasLastObservedTime returns a boolean if a field has been set.

### GetCount

`func (o *JsonV1EventSeries) GetCount() float32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *JsonV1EventSeries) GetCountOk() (*float32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *JsonV1EventSeries) SetCount(v float32)`

SetCount sets Count field to given value.

### HasCount

`func (o *JsonV1EventSeries) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


