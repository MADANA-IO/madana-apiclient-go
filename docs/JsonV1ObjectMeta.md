# JsonV1ObjectMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GenerateName** | Pointer to **string** |  | [optional] 
**SelfLink** | Pointer to **string** |  | [optional] 
**ClusterName** | Pointer to **string** |  | [optional] 
**OwnerReferences** | Pointer to [**[]JsonV1OwnerReference**](JsonV1OwnerReference.md) |  | [optional] 
**ManagedFields** | Pointer to [**[]JsonV1ManagedFieldsEntry**](JsonV1ManagedFieldsEntry.md) |  | [optional] 
**Generation** | Pointer to **float32** |  | [optional] 
**Annotations** | Pointer to **map[string]string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**Uid** | Pointer to **string** |  | [optional] 
**DeletionTimestamp** | Pointer to **float32** |  | [optional] 
**Labels** | Pointer to **map[string]string** |  | [optional] 
**CreationTimestamp** | Pointer to **float32** |  | [optional] 
**ResourceVersion** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Finalizers** | Pointer to **[]string** |  | [optional] 
**DeletionGracePeriodSeconds** | Pointer to **float32** |  | [optional] 

## Methods

### NewJsonV1ObjectMeta

`func NewJsonV1ObjectMeta() *JsonV1ObjectMeta`

NewJsonV1ObjectMeta instantiates a new JsonV1ObjectMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJsonV1ObjectMetaWithDefaults

`func NewJsonV1ObjectMetaWithDefaults() *JsonV1ObjectMeta`

NewJsonV1ObjectMetaWithDefaults instantiates a new JsonV1ObjectMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGenerateName

`func (o *JsonV1ObjectMeta) GetGenerateName() string`

GetGenerateName returns the GenerateName field if non-nil, zero value otherwise.

### GetGenerateNameOk

`func (o *JsonV1ObjectMeta) GetGenerateNameOk() (*string, bool)`

GetGenerateNameOk returns a tuple with the GenerateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateName

`func (o *JsonV1ObjectMeta) SetGenerateName(v string)`

SetGenerateName sets GenerateName field to given value.

### HasGenerateName

`func (o *JsonV1ObjectMeta) HasGenerateName() bool`

HasGenerateName returns a boolean if a field has been set.

### GetSelfLink

`func (o *JsonV1ObjectMeta) GetSelfLink() string`

GetSelfLink returns the SelfLink field if non-nil, zero value otherwise.

### GetSelfLinkOk

`func (o *JsonV1ObjectMeta) GetSelfLinkOk() (*string, bool)`

GetSelfLinkOk returns a tuple with the SelfLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfLink

`func (o *JsonV1ObjectMeta) SetSelfLink(v string)`

SetSelfLink sets SelfLink field to given value.

### HasSelfLink

`func (o *JsonV1ObjectMeta) HasSelfLink() bool`

HasSelfLink returns a boolean if a field has been set.

### GetClusterName

`func (o *JsonV1ObjectMeta) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *JsonV1ObjectMeta) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *JsonV1ObjectMeta) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *JsonV1ObjectMeta) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### GetOwnerReferences

`func (o *JsonV1ObjectMeta) GetOwnerReferences() []JsonV1OwnerReference`

GetOwnerReferences returns the OwnerReferences field if non-nil, zero value otherwise.

### GetOwnerReferencesOk

`func (o *JsonV1ObjectMeta) GetOwnerReferencesOk() (*[]JsonV1OwnerReference, bool)`

GetOwnerReferencesOk returns a tuple with the OwnerReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerReferences

`func (o *JsonV1ObjectMeta) SetOwnerReferences(v []JsonV1OwnerReference)`

SetOwnerReferences sets OwnerReferences field to given value.

### HasOwnerReferences

`func (o *JsonV1ObjectMeta) HasOwnerReferences() bool`

HasOwnerReferences returns a boolean if a field has been set.

### GetManagedFields

`func (o *JsonV1ObjectMeta) GetManagedFields() []JsonV1ManagedFieldsEntry`

GetManagedFields returns the ManagedFields field if non-nil, zero value otherwise.

### GetManagedFieldsOk

`func (o *JsonV1ObjectMeta) GetManagedFieldsOk() (*[]JsonV1ManagedFieldsEntry, bool)`

GetManagedFieldsOk returns a tuple with the ManagedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedFields

`func (o *JsonV1ObjectMeta) SetManagedFields(v []JsonV1ManagedFieldsEntry)`

SetManagedFields sets ManagedFields field to given value.

### HasManagedFields

`func (o *JsonV1ObjectMeta) HasManagedFields() bool`

HasManagedFields returns a boolean if a field has been set.

### GetGeneration

`func (o *JsonV1ObjectMeta) GetGeneration() float32`

GetGeneration returns the Generation field if non-nil, zero value otherwise.

### GetGenerationOk

`func (o *JsonV1ObjectMeta) GetGenerationOk() (*float32, bool)`

GetGenerationOk returns a tuple with the Generation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneration

`func (o *JsonV1ObjectMeta) SetGeneration(v float32)`

SetGeneration sets Generation field to given value.

### HasGeneration

`func (o *JsonV1ObjectMeta) HasGeneration() bool`

HasGeneration returns a boolean if a field has been set.

### GetAnnotations

`func (o *JsonV1ObjectMeta) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *JsonV1ObjectMeta) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *JsonV1ObjectMeta) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *JsonV1ObjectMeta) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetNamespace

`func (o *JsonV1ObjectMeta) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *JsonV1ObjectMeta) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *JsonV1ObjectMeta) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *JsonV1ObjectMeta) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetUid

`func (o *JsonV1ObjectMeta) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *JsonV1ObjectMeta) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *JsonV1ObjectMeta) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *JsonV1ObjectMeta) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetDeletionTimestamp

`func (o *JsonV1ObjectMeta) GetDeletionTimestamp() float32`

GetDeletionTimestamp returns the DeletionTimestamp field if non-nil, zero value otherwise.

### GetDeletionTimestampOk

`func (o *JsonV1ObjectMeta) GetDeletionTimestampOk() (*float32, bool)`

GetDeletionTimestampOk returns a tuple with the DeletionTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionTimestamp

`func (o *JsonV1ObjectMeta) SetDeletionTimestamp(v float32)`

SetDeletionTimestamp sets DeletionTimestamp field to given value.

### HasDeletionTimestamp

`func (o *JsonV1ObjectMeta) HasDeletionTimestamp() bool`

HasDeletionTimestamp returns a boolean if a field has been set.

### GetLabels

`func (o *JsonV1ObjectMeta) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *JsonV1ObjectMeta) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *JsonV1ObjectMeta) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *JsonV1ObjectMeta) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetCreationTimestamp

`func (o *JsonV1ObjectMeta) GetCreationTimestamp() float32`

GetCreationTimestamp returns the CreationTimestamp field if non-nil, zero value otherwise.

### GetCreationTimestampOk

`func (o *JsonV1ObjectMeta) GetCreationTimestampOk() (*float32, bool)`

GetCreationTimestampOk returns a tuple with the CreationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTimestamp

`func (o *JsonV1ObjectMeta) SetCreationTimestamp(v float32)`

SetCreationTimestamp sets CreationTimestamp field to given value.

### HasCreationTimestamp

`func (o *JsonV1ObjectMeta) HasCreationTimestamp() bool`

HasCreationTimestamp returns a boolean if a field has been set.

### GetResourceVersion

`func (o *JsonV1ObjectMeta) GetResourceVersion() string`

GetResourceVersion returns the ResourceVersion field if non-nil, zero value otherwise.

### GetResourceVersionOk

`func (o *JsonV1ObjectMeta) GetResourceVersionOk() (*string, bool)`

GetResourceVersionOk returns a tuple with the ResourceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceVersion

`func (o *JsonV1ObjectMeta) SetResourceVersion(v string)`

SetResourceVersion sets ResourceVersion field to given value.

### HasResourceVersion

`func (o *JsonV1ObjectMeta) HasResourceVersion() bool`

HasResourceVersion returns a boolean if a field has been set.

### GetName

`func (o *JsonV1ObjectMeta) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JsonV1ObjectMeta) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JsonV1ObjectMeta) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *JsonV1ObjectMeta) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFinalizers

`func (o *JsonV1ObjectMeta) GetFinalizers() []string`

GetFinalizers returns the Finalizers field if non-nil, zero value otherwise.

### GetFinalizersOk

`func (o *JsonV1ObjectMeta) GetFinalizersOk() (*[]string, bool)`

GetFinalizersOk returns a tuple with the Finalizers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalizers

`func (o *JsonV1ObjectMeta) SetFinalizers(v []string)`

SetFinalizers sets Finalizers field to given value.

### HasFinalizers

`func (o *JsonV1ObjectMeta) HasFinalizers() bool`

HasFinalizers returns a boolean if a field has been set.

### GetDeletionGracePeriodSeconds

`func (o *JsonV1ObjectMeta) GetDeletionGracePeriodSeconds() float32`

GetDeletionGracePeriodSeconds returns the DeletionGracePeriodSeconds field if non-nil, zero value otherwise.

### GetDeletionGracePeriodSecondsOk

`func (o *JsonV1ObjectMeta) GetDeletionGracePeriodSecondsOk() (*float32, bool)`

GetDeletionGracePeriodSecondsOk returns a tuple with the DeletionGracePeriodSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionGracePeriodSeconds

`func (o *JsonV1ObjectMeta) SetDeletionGracePeriodSeconds(v float32)`

SetDeletionGracePeriodSeconds sets DeletionGracePeriodSeconds field to given value.

### HasDeletionGracePeriodSeconds

`func (o *JsonV1ObjectMeta) HasDeletionGracePeriodSeconds() bool`

HasDeletionGracePeriodSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


