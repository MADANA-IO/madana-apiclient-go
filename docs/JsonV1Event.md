# JsonV1Event

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | Pointer to **string** |  | [optional] 
**Action** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Series** | Pointer to [**JsonV1EventSeries**](json_V1EventSeries.md) |  | [optional] 
**ReportingInstance** | Pointer to **string** |  | [optional] 
**EventTime** | Pointer to **float32** |  | [optional] 
**LastTimestamp** | Pointer to **float32** |  | [optional] 
**Related** | Pointer to [**JsonV1ObjectReference**](json_V1ObjectReference.md) |  | [optional] 
**ReportingComponent** | Pointer to **string** |  | [optional] 
**Source** | Pointer to [**JsonV1EventSource**](json_V1EventSource.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Count** | Pointer to **float32** |  | [optional] 
**InvolvedObject** | Pointer to [**JsonV1ObjectReference**](json_V1ObjectReference.md) |  | [optional] 
**FirstTimestamp** | Pointer to **float32** |  | [optional] 
**Metadata** | Pointer to [**JsonV1ObjectMeta**](json_V1ObjectMeta.md) |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**ApiVersion** | Pointer to **string** |  | [optional] 

## Methods

### NewJsonV1Event

`func NewJsonV1Event() *JsonV1Event`

NewJsonV1Event instantiates a new JsonV1Event object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJsonV1EventWithDefaults

`func NewJsonV1EventWithDefaults() *JsonV1Event`

NewJsonV1EventWithDefaults instantiates a new JsonV1Event object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *JsonV1Event) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *JsonV1Event) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *JsonV1Event) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *JsonV1Event) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetAction

`func (o *JsonV1Event) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *JsonV1Event) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *JsonV1Event) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *JsonV1Event) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetType

`func (o *JsonV1Event) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *JsonV1Event) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *JsonV1Event) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *JsonV1Event) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSeries

`func (o *JsonV1Event) GetSeries() JsonV1EventSeries`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *JsonV1Event) GetSeriesOk() (*JsonV1EventSeries, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *JsonV1Event) SetSeries(v JsonV1EventSeries)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *JsonV1Event) HasSeries() bool`

HasSeries returns a boolean if a field has been set.

### GetReportingInstance

`func (o *JsonV1Event) GetReportingInstance() string`

GetReportingInstance returns the ReportingInstance field if non-nil, zero value otherwise.

### GetReportingInstanceOk

`func (o *JsonV1Event) GetReportingInstanceOk() (*string, bool)`

GetReportingInstanceOk returns a tuple with the ReportingInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingInstance

`func (o *JsonV1Event) SetReportingInstance(v string)`

SetReportingInstance sets ReportingInstance field to given value.

### HasReportingInstance

`func (o *JsonV1Event) HasReportingInstance() bool`

HasReportingInstance returns a boolean if a field has been set.

### GetEventTime

`func (o *JsonV1Event) GetEventTime() float32`

GetEventTime returns the EventTime field if non-nil, zero value otherwise.

### GetEventTimeOk

`func (o *JsonV1Event) GetEventTimeOk() (*float32, bool)`

GetEventTimeOk returns a tuple with the EventTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTime

`func (o *JsonV1Event) SetEventTime(v float32)`

SetEventTime sets EventTime field to given value.

### HasEventTime

`func (o *JsonV1Event) HasEventTime() bool`

HasEventTime returns a boolean if a field has been set.

### GetLastTimestamp

`func (o *JsonV1Event) GetLastTimestamp() float32`

GetLastTimestamp returns the LastTimestamp field if non-nil, zero value otherwise.

### GetLastTimestampOk

`func (o *JsonV1Event) GetLastTimestampOk() (*float32, bool)`

GetLastTimestampOk returns a tuple with the LastTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTimestamp

`func (o *JsonV1Event) SetLastTimestamp(v float32)`

SetLastTimestamp sets LastTimestamp field to given value.

### HasLastTimestamp

`func (o *JsonV1Event) HasLastTimestamp() bool`

HasLastTimestamp returns a boolean if a field has been set.

### GetRelated

`func (o *JsonV1Event) GetRelated() JsonV1ObjectReference`

GetRelated returns the Related field if non-nil, zero value otherwise.

### GetRelatedOk

`func (o *JsonV1Event) GetRelatedOk() (*JsonV1ObjectReference, bool)`

GetRelatedOk returns a tuple with the Related field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelated

`func (o *JsonV1Event) SetRelated(v JsonV1ObjectReference)`

SetRelated sets Related field to given value.

### HasRelated

`func (o *JsonV1Event) HasRelated() bool`

HasRelated returns a boolean if a field has been set.

### GetReportingComponent

`func (o *JsonV1Event) GetReportingComponent() string`

GetReportingComponent returns the ReportingComponent field if non-nil, zero value otherwise.

### GetReportingComponentOk

`func (o *JsonV1Event) GetReportingComponentOk() (*string, bool)`

GetReportingComponentOk returns a tuple with the ReportingComponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingComponent

`func (o *JsonV1Event) SetReportingComponent(v string)`

SetReportingComponent sets ReportingComponent field to given value.

### HasReportingComponent

`func (o *JsonV1Event) HasReportingComponent() bool`

HasReportingComponent returns a boolean if a field has been set.

### GetSource

`func (o *JsonV1Event) GetSource() JsonV1EventSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *JsonV1Event) GetSourceOk() (*JsonV1EventSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *JsonV1Event) SetSource(v JsonV1EventSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *JsonV1Event) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetMessage

`func (o *JsonV1Event) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *JsonV1Event) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *JsonV1Event) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *JsonV1Event) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetCount

`func (o *JsonV1Event) GetCount() float32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *JsonV1Event) GetCountOk() (*float32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *JsonV1Event) SetCount(v float32)`

SetCount sets Count field to given value.

### HasCount

`func (o *JsonV1Event) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetInvolvedObject

`func (o *JsonV1Event) GetInvolvedObject() JsonV1ObjectReference`

GetInvolvedObject returns the InvolvedObject field if non-nil, zero value otherwise.

### GetInvolvedObjectOk

`func (o *JsonV1Event) GetInvolvedObjectOk() (*JsonV1ObjectReference, bool)`

GetInvolvedObjectOk returns a tuple with the InvolvedObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvolvedObject

`func (o *JsonV1Event) SetInvolvedObject(v JsonV1ObjectReference)`

SetInvolvedObject sets InvolvedObject field to given value.

### HasInvolvedObject

`func (o *JsonV1Event) HasInvolvedObject() bool`

HasInvolvedObject returns a boolean if a field has been set.

### GetFirstTimestamp

`func (o *JsonV1Event) GetFirstTimestamp() float32`

GetFirstTimestamp returns the FirstTimestamp field if non-nil, zero value otherwise.

### GetFirstTimestampOk

`func (o *JsonV1Event) GetFirstTimestampOk() (*float32, bool)`

GetFirstTimestampOk returns a tuple with the FirstTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstTimestamp

`func (o *JsonV1Event) SetFirstTimestamp(v float32)`

SetFirstTimestamp sets FirstTimestamp field to given value.

### HasFirstTimestamp

`func (o *JsonV1Event) HasFirstTimestamp() bool`

HasFirstTimestamp returns a boolean if a field has been set.

### GetMetadata

`func (o *JsonV1Event) GetMetadata() JsonV1ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *JsonV1Event) GetMetadataOk() (*JsonV1ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *JsonV1Event) SetMetadata(v JsonV1ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *JsonV1Event) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetKind

`func (o *JsonV1Event) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *JsonV1Event) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *JsonV1Event) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *JsonV1Event) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetApiVersion

`func (o *JsonV1Event) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *JsonV1Event) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *JsonV1Event) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *JsonV1Event) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


