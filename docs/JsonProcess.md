# JsonProcess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OutputStream** | Pointer to **map[string]interface{}** |  | [optional] 
**Alive** | Pointer to **bool** |  | [optional] 
**InputStream** | Pointer to **map[string]interface{}** |  | [optional] 
**ErrorStream** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewJsonProcess

`func NewJsonProcess() *JsonProcess`

NewJsonProcess instantiates a new JsonProcess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJsonProcessWithDefaults

`func NewJsonProcessWithDefaults() *JsonProcess`

NewJsonProcessWithDefaults instantiates a new JsonProcess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOutputStream

`func (o *JsonProcess) GetOutputStream() map[string]interface{}`

GetOutputStream returns the OutputStream field if non-nil, zero value otherwise.

### GetOutputStreamOk

`func (o *JsonProcess) GetOutputStreamOk() (*map[string]interface{}, bool)`

GetOutputStreamOk returns a tuple with the OutputStream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputStream

`func (o *JsonProcess) SetOutputStream(v map[string]interface{})`

SetOutputStream sets OutputStream field to given value.

### HasOutputStream

`func (o *JsonProcess) HasOutputStream() bool`

HasOutputStream returns a boolean if a field has been set.

### GetAlive

`func (o *JsonProcess) GetAlive() bool`

GetAlive returns the Alive field if non-nil, zero value otherwise.

### GetAliveOk

`func (o *JsonProcess) GetAliveOk() (*bool, bool)`

GetAliveOk returns a tuple with the Alive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlive

`func (o *JsonProcess) SetAlive(v bool)`

SetAlive sets Alive field to given value.

### HasAlive

`func (o *JsonProcess) HasAlive() bool`

HasAlive returns a boolean if a field has been set.

### GetInputStream

`func (o *JsonProcess) GetInputStream() map[string]interface{}`

GetInputStream returns the InputStream field if non-nil, zero value otherwise.

### GetInputStreamOk

`func (o *JsonProcess) GetInputStreamOk() (*map[string]interface{}, bool)`

GetInputStreamOk returns a tuple with the InputStream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputStream

`func (o *JsonProcess) SetInputStream(v map[string]interface{})`

SetInputStream sets InputStream field to given value.

### HasInputStream

`func (o *JsonProcess) HasInputStream() bool`

HasInputStream returns a boolean if a field has been set.

### GetErrorStream

`func (o *JsonProcess) GetErrorStream() map[string]interface{}`

GetErrorStream returns the ErrorStream field if non-nil, zero value otherwise.

### GetErrorStreamOk

`func (o *JsonProcess) GetErrorStreamOk() (*map[string]interface{}, bool)`

GetErrorStreamOk returns a tuple with the ErrorStream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorStream

`func (o *JsonProcess) SetErrorStream(v map[string]interface{})`

SetErrorStream sets ErrorStream field to given value.

### HasErrorStream

`func (o *JsonProcess) HasErrorStream() bool`

HasErrorStream returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


