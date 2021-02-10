# JsonSignedData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Signature** | Pointer to **string** |  | [optional] 
**Fingerpint** | Pointer to **string** |  | [optional] 
**Data** | Pointer to **string** |  | [optional] 

## Methods

### NewJsonSignedData

`func NewJsonSignedData() *JsonSignedData`

NewJsonSignedData instantiates a new JsonSignedData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJsonSignedDataWithDefaults

`func NewJsonSignedDataWithDefaults() *JsonSignedData`

NewJsonSignedDataWithDefaults instantiates a new JsonSignedData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignature

`func (o *JsonSignedData) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *JsonSignedData) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *JsonSignedData) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *JsonSignedData) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### GetFingerpint

`func (o *JsonSignedData) GetFingerpint() string`

GetFingerpint returns the Fingerpint field if non-nil, zero value otherwise.

### GetFingerpintOk

`func (o *JsonSignedData) GetFingerpintOk() (*string, bool)`

GetFingerpintOk returns a tuple with the Fingerpint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerpint

`func (o *JsonSignedData) SetFingerpint(v string)`

SetFingerpint sets Fingerpint field to given value.

### HasFingerpint

`func (o *JsonSignedData) HasFingerpint() bool`

HasFingerpint returns a boolean if a field has been set.

### GetData

`func (o *JsonSignedData) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *JsonSignedData) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *JsonSignedData) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *JsonSignedData) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


