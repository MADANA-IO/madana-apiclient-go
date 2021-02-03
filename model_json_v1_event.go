/*
 * madana-api
 *
 * <h1>Using the madana-api</h1>        <p>This documentation contains a Quickstart Guide, relating client functionality and information about the available         endpoints and used datamodels.   </p>       <p> The madana-api and its implementations are still in heavy development. This means that there may be problems in our protocols, or there may be mistakes in our implementations. We take security vulnerabilities very seriously. If you discover a security issue, please bring it to our attention right away! If you find a vulnerability that may affect live deployments -- for example, by exposing a remote execution exploit -- please send your report privately to info@madana.io. Please DO NOT file a public issue. If the issue is a protocol weakness that cannot be immediately exploited or something not yet deployed, just discuss it openly   </p>   <br>   <p> Note: Not all functionality might be acessible without having accquired and api-license token. For more information visit <a href=\"https://www.madana.io\">www.madana.io</a> </p>       <br>
 *
 * API version: 0.5.0-master.47
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package madanaapiclient

import (
	"encoding/json"
)

// JsonV1Event 
type JsonV1Event struct {
	// 
	Type *string `json:"type,omitempty"`
	// 
	Reason *string `json:"reason,omitempty"`
	Metadata *JsonV1ObjectMeta `json:"metadata,omitempty"`
	// 
	ReportingInstance *string `json:"reportingInstance,omitempty"`
	// 
	ApiVersion *string `json:"apiVersion,omitempty"`
	// 
	FirstTimestamp *float32 `json:"firstTimestamp,omitempty"`
	// 
	Kind *string `json:"kind,omitempty"`
	Series *JsonV1EventSeries `json:"series,omitempty"`
	// 
	Message *string `json:"message,omitempty"`
	// 
	Count *float32 `json:"count,omitempty"`
	Related *JsonV1ObjectReference `json:"related,omitempty"`
	// 
	LastTimestamp *float32 `json:"lastTimestamp,omitempty"`
	InvolvedObject *JsonV1ObjectReference `json:"involvedObject,omitempty"`
	Source *JsonV1EventSource `json:"source,omitempty"`
	// 
	Action *string `json:"action,omitempty"`
	// 
	EventTime *float32 `json:"eventTime,omitempty"`
	// 
	ReportingComponent *string `json:"reportingComponent,omitempty"`
}

// NewJsonV1Event instantiates a new JsonV1Event object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJsonV1Event() *JsonV1Event {
	this := JsonV1Event{}
	return &this
}

// NewJsonV1EventWithDefaults instantiates a new JsonV1Event object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJsonV1EventWithDefaults() *JsonV1Event {
	this := JsonV1Event{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *JsonV1Event) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonV1Event) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *JsonV1Event) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *JsonV1Event) SetType(v string) {
	o.Type = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *JsonV1Event) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonV1Event) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *JsonV1Event) HasReason() bool {
	if o != nil && o.Reason != nil {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *JsonV1Event) SetReason(v string) {
	o.Reason = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *JsonV1Event) GetMetadata() JsonV1ObjectMeta {
	if o == nil || o.Metadata == nil {
		var ret JsonV1ObjectMeta
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonV1Event) GetMetadataOk() (*JsonV1ObjectMeta, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *JsonV1Event) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given JsonV1ObjectMeta and assigns it to the Metadata field.
func (o *JsonV1Event) SetMetadata(v JsonV1ObjectMeta) {
	o.Metadata = &v
}

// GetReportingInstance returns the ReportingInstance field value if set, zero value otherwise.
func (o *JsonV1Event) GetReportingInstance() string {
	if o == nil || o.ReportingInstance == nil {
		var ret string
		return ret
	}
	return *o.ReportingInstance
}

// GetReportingInstanceOk returns a tuple with the ReportingInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonV1Event) GetReportingInstanceOk() (*string, bool) {
	if o == nil || o.ReportingInstance == nil {
		return nil, false
	}
	return o.ReportingInstance, true
}

// HasReportingInstance returns a boolean if a field has been set.
func (o *JsonV1Event) HasReportingInstance() bool {
	if o != nil && o.ReportingInstance != nil {
		return true
	}

	return false
}

// SetReportingInstance gets a reference to the given string and assigns it to the ReportingInstance field.
func (o *JsonV1Event) SetReportingInstance(v string) {
	o.ReportingInstance = &v
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *JsonV1Event) GetApiVersion() string {
	if o == nil || o.ApiVersion == nil {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonV1Event) GetApiVersionOk() (*string, bool) {
	if o == nil || o.ApiVersion == nil {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *JsonV1Event) HasApiVersion() bool {
	if o != nil && o.ApiVersion != nil {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *JsonV1Event) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetFirstTimestamp returns the FirstTimestamp field value if set, zero value otherwise.
func (o *JsonV1Event) GetFirstTimestamp() float32 {
	if o == nil || o.FirstTimestamp == nil {
		var ret float32
		return ret
	}
	return *o.FirstTimestamp
}

// GetFirstTimestampOk returns a tuple with the FirstTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonV1Event) GetFirstTimestampOk() (*float32, bool) {
	if o == nil || o.FirstTimestamp == nil {
		return nil, false
	}
	return o.FirstTimestamp, true
}

// HasFirstTimestamp returns a boolean if a field has been set.
func (o *JsonV1Event) HasFirstTimestamp() bool {
	if o != nil && o.FirstTimestamp != nil {
		return true
	}

	return false
}

// SetFirstTimestamp gets a reference to the given float32 and assigns it to the FirstTimestamp field.
func (o *JsonV1Event) SetFirstTimestamp(v float32) {
	o.FirstTimestamp = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *JsonV1Event) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonV1Event) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *JsonV1Event) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *JsonV1Event) SetKind(v string) {
	o.Kind = &v
}

// GetSeries returns the Series field value if set, zero value otherwise.
func (o *JsonV1Event) GetSeries() JsonV1EventSeries {
	if o == nil || o.Series == nil {
		var ret JsonV1EventSeries
		return ret
	}
	return *o.Series
}

// GetSeriesOk returns a tuple with the Series field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonV1Event) GetSeriesOk() (*JsonV1EventSeries, bool) {
	if o == nil || o.Series == nil {
		return nil, false
	}
	return o.Series, true
}

// HasSeries returns a boolean if a field has been set.
func (o *JsonV1Event) HasSeries() bool {
	if o != nil && o.Series != nil {
		return true
	}

	return false
}

// SetSeries gets a reference to the given JsonV1EventSeries and assigns it to the Series field.
func (o *JsonV1Event) SetSeries(v JsonV1EventSeries) {
	o.Series = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *JsonV1Event) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonV1Event) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *JsonV1Event) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *JsonV1Event) SetMessage(v string) {
	o.Message = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *JsonV1Event) GetCount() float32 {
	if o == nil || o.Count == nil {
		var ret float32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonV1Event) GetCountOk() (*float32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *JsonV1Event) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given float32 and assigns it to the Count field.
func (o *JsonV1Event) SetCount(v float32) {
	o.Count = &v
}

// GetRelated returns the Related field value if set, zero value otherwise.
func (o *JsonV1Event) GetRelated() JsonV1ObjectReference {
	if o == nil || o.Related == nil {
		var ret JsonV1ObjectReference
		return ret
	}
	return *o.Related
}

// GetRelatedOk returns a tuple with the Related field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonV1Event) GetRelatedOk() (*JsonV1ObjectReference, bool) {
	if o == nil || o.Related == nil {
		return nil, false
	}
	return o.Related, true
}

// HasRelated returns a boolean if a field has been set.
func (o *JsonV1Event) HasRelated() bool {
	if o != nil && o.Related != nil {
		return true
	}

	return false
}

// SetRelated gets a reference to the given JsonV1ObjectReference and assigns it to the Related field.
func (o *JsonV1Event) SetRelated(v JsonV1ObjectReference) {
	o.Related = &v
}

// GetLastTimestamp returns the LastTimestamp field value if set, zero value otherwise.
func (o *JsonV1Event) GetLastTimestamp() float32 {
	if o == nil || o.LastTimestamp == nil {
		var ret float32
		return ret
	}
	return *o.LastTimestamp
}

// GetLastTimestampOk returns a tuple with the LastTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonV1Event) GetLastTimestampOk() (*float32, bool) {
	if o == nil || o.LastTimestamp == nil {
		return nil, false
	}
	return o.LastTimestamp, true
}

// HasLastTimestamp returns a boolean if a field has been set.
func (o *JsonV1Event) HasLastTimestamp() bool {
	if o != nil && o.LastTimestamp != nil {
		return true
	}

	return false
}

// SetLastTimestamp gets a reference to the given float32 and assigns it to the LastTimestamp field.
func (o *JsonV1Event) SetLastTimestamp(v float32) {
	o.LastTimestamp = &v
}

// GetInvolvedObject returns the InvolvedObject field value if set, zero value otherwise.
func (o *JsonV1Event) GetInvolvedObject() JsonV1ObjectReference {
	if o == nil || o.InvolvedObject == nil {
		var ret JsonV1ObjectReference
		return ret
	}
	return *o.InvolvedObject
}

// GetInvolvedObjectOk returns a tuple with the InvolvedObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonV1Event) GetInvolvedObjectOk() (*JsonV1ObjectReference, bool) {
	if o == nil || o.InvolvedObject == nil {
		return nil, false
	}
	return o.InvolvedObject, true
}

// HasInvolvedObject returns a boolean if a field has been set.
func (o *JsonV1Event) HasInvolvedObject() bool {
	if o != nil && o.InvolvedObject != nil {
		return true
	}

	return false
}

// SetInvolvedObject gets a reference to the given JsonV1ObjectReference and assigns it to the InvolvedObject field.
func (o *JsonV1Event) SetInvolvedObject(v JsonV1ObjectReference) {
	o.InvolvedObject = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *JsonV1Event) GetSource() JsonV1EventSource {
	if o == nil || o.Source == nil {
		var ret JsonV1EventSource
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonV1Event) GetSourceOk() (*JsonV1EventSource, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *JsonV1Event) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given JsonV1EventSource and assigns it to the Source field.
func (o *JsonV1Event) SetSource(v JsonV1EventSource) {
	o.Source = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *JsonV1Event) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonV1Event) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *JsonV1Event) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *JsonV1Event) SetAction(v string) {
	o.Action = &v
}

// GetEventTime returns the EventTime field value if set, zero value otherwise.
func (o *JsonV1Event) GetEventTime() float32 {
	if o == nil || o.EventTime == nil {
		var ret float32
		return ret
	}
	return *o.EventTime
}

// GetEventTimeOk returns a tuple with the EventTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonV1Event) GetEventTimeOk() (*float32, bool) {
	if o == nil || o.EventTime == nil {
		return nil, false
	}
	return o.EventTime, true
}

// HasEventTime returns a boolean if a field has been set.
func (o *JsonV1Event) HasEventTime() bool {
	if o != nil && o.EventTime != nil {
		return true
	}

	return false
}

// SetEventTime gets a reference to the given float32 and assigns it to the EventTime field.
func (o *JsonV1Event) SetEventTime(v float32) {
	o.EventTime = &v
}

// GetReportingComponent returns the ReportingComponent field value if set, zero value otherwise.
func (o *JsonV1Event) GetReportingComponent() string {
	if o == nil || o.ReportingComponent == nil {
		var ret string
		return ret
	}
	return *o.ReportingComponent
}

// GetReportingComponentOk returns a tuple with the ReportingComponent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonV1Event) GetReportingComponentOk() (*string, bool) {
	if o == nil || o.ReportingComponent == nil {
		return nil, false
	}
	return o.ReportingComponent, true
}

// HasReportingComponent returns a boolean if a field has been set.
func (o *JsonV1Event) HasReportingComponent() bool {
	if o != nil && o.ReportingComponent != nil {
		return true
	}

	return false
}

// SetReportingComponent gets a reference to the given string and assigns it to the ReportingComponent field.
func (o *JsonV1Event) SetReportingComponent(v string) {
	o.ReportingComponent = &v
}

func (o JsonV1Event) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.ReportingInstance != nil {
		toSerialize["reportingInstance"] = o.ReportingInstance
	}
	if o.ApiVersion != nil {
		toSerialize["apiVersion"] = o.ApiVersion
	}
	if o.FirstTimestamp != nil {
		toSerialize["firstTimestamp"] = o.FirstTimestamp
	}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.Series != nil {
		toSerialize["series"] = o.Series
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	if o.Related != nil {
		toSerialize["related"] = o.Related
	}
	if o.LastTimestamp != nil {
		toSerialize["lastTimestamp"] = o.LastTimestamp
	}
	if o.InvolvedObject != nil {
		toSerialize["involvedObject"] = o.InvolvedObject
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if o.Action != nil {
		toSerialize["action"] = o.Action
	}
	if o.EventTime != nil {
		toSerialize["eventTime"] = o.EventTime
	}
	if o.ReportingComponent != nil {
		toSerialize["reportingComponent"] = o.ReportingComponent
	}
	return json.Marshal(toSerialize)
}

type NullableJsonV1Event struct {
	value *JsonV1Event
	isSet bool
}

func (v NullableJsonV1Event) Get() *JsonV1Event {
	return v.value
}

func (v *NullableJsonV1Event) Set(val *JsonV1Event) {
	v.value = val
	v.isSet = true
}

func (v NullableJsonV1Event) IsSet() bool {
	return v.isSet
}

func (v *NullableJsonV1Event) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJsonV1Event(val *JsonV1Event) *NullableJsonV1Event {
	return &NullableJsonV1Event{value: val, isSet: true}
}

func (v NullableJsonV1Event) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJsonV1Event) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


