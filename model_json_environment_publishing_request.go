/*
 * madana-api
 *
 * <h1>Using the madana-api</h1>        <p>This documentation contains a Quickstart Guide, relating client functionality and information about the available         endpoints and used datamodels.   </p>       <p> The madana-api and its implementations are still in heavy development. This means that there may be problems in our protocols, or there may be mistakes in our implementations. We take security vulnerabilities very seriously. If you discover a security issue, please bring it to our attention right away! If you find a vulnerability that may affect live deployments -- for example, by exposing a remote execution exploit -- please send your report privately to info@madana.io. Please DO NOT file a public issue. If the issue is a protocol weakness that cannot be immediately exploited or something not yet deployed, just discuss it openly   </p>   <br>   <p> Note: Not all functionality might be acessible without having accquired and api-license token. For more information visit <a href=\"https://www.madana.io\">www.madana.io</a> </p>       <br>
 *
 * API version: 0.5.0-master.56
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package madanaapiclient

import (
	"encoding/json"
)

// JsonEnvironmentPublishingRequest 
type JsonEnvironmentPublishingRequest struct {
	// 
	IpfsPrimaryPeer *string `json:"ipfsPrimaryPeer,omitempty"`
	// 
	Packages *string `json:"packages,omitempty"`
	// 
	Content *string `json:"content,omitempty"`
	// 
	IsPublic *string `json:"isPublic,omitempty"`
	// 
	Description *string `json:"description,omitempty"`
	// 
	Name *string `json:"name,omitempty"`
	// 
	Uuid *string `json:"uuid,omitempty"`
	// 
	IpfsHash *string `json:"ipfsHash,omitempty"`
	// 
	Size *string `json:"size,omitempty"`
}

// NewJsonEnvironmentPublishingRequest instantiates a new JsonEnvironmentPublishingRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJsonEnvironmentPublishingRequest() *JsonEnvironmentPublishingRequest {
	this := JsonEnvironmentPublishingRequest{}
	return &this
}

// NewJsonEnvironmentPublishingRequestWithDefaults instantiates a new JsonEnvironmentPublishingRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJsonEnvironmentPublishingRequestWithDefaults() *JsonEnvironmentPublishingRequest {
	this := JsonEnvironmentPublishingRequest{}
	return &this
}

// GetIpfsPrimaryPeer returns the IpfsPrimaryPeer field value if set, zero value otherwise.
func (o *JsonEnvironmentPublishingRequest) GetIpfsPrimaryPeer() string {
	if o == nil || o.IpfsPrimaryPeer == nil {
		var ret string
		return ret
	}
	return *o.IpfsPrimaryPeer
}

// GetIpfsPrimaryPeerOk returns a tuple with the IpfsPrimaryPeer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonEnvironmentPublishingRequest) GetIpfsPrimaryPeerOk() (*string, bool) {
	if o == nil || o.IpfsPrimaryPeer == nil {
		return nil, false
	}
	return o.IpfsPrimaryPeer, true
}

// HasIpfsPrimaryPeer returns a boolean if a field has been set.
func (o *JsonEnvironmentPublishingRequest) HasIpfsPrimaryPeer() bool {
	if o != nil && o.IpfsPrimaryPeer != nil {
		return true
	}

	return false
}

// SetIpfsPrimaryPeer gets a reference to the given string and assigns it to the IpfsPrimaryPeer field.
func (o *JsonEnvironmentPublishingRequest) SetIpfsPrimaryPeer(v string) {
	o.IpfsPrimaryPeer = &v
}

// GetPackages returns the Packages field value if set, zero value otherwise.
func (o *JsonEnvironmentPublishingRequest) GetPackages() string {
	if o == nil || o.Packages == nil {
		var ret string
		return ret
	}
	return *o.Packages
}

// GetPackagesOk returns a tuple with the Packages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonEnvironmentPublishingRequest) GetPackagesOk() (*string, bool) {
	if o == nil || o.Packages == nil {
		return nil, false
	}
	return o.Packages, true
}

// HasPackages returns a boolean if a field has been set.
func (o *JsonEnvironmentPublishingRequest) HasPackages() bool {
	if o != nil && o.Packages != nil {
		return true
	}

	return false
}

// SetPackages gets a reference to the given string and assigns it to the Packages field.
func (o *JsonEnvironmentPublishingRequest) SetPackages(v string) {
	o.Packages = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *JsonEnvironmentPublishingRequest) GetContent() string {
	if o == nil || o.Content == nil {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonEnvironmentPublishingRequest) GetContentOk() (*string, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *JsonEnvironmentPublishingRequest) HasContent() bool {
	if o != nil && o.Content != nil {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *JsonEnvironmentPublishingRequest) SetContent(v string) {
	o.Content = &v
}

// GetIsPublic returns the IsPublic field value if set, zero value otherwise.
func (o *JsonEnvironmentPublishingRequest) GetIsPublic() string {
	if o == nil || o.IsPublic == nil {
		var ret string
		return ret
	}
	return *o.IsPublic
}

// GetIsPublicOk returns a tuple with the IsPublic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonEnvironmentPublishingRequest) GetIsPublicOk() (*string, bool) {
	if o == nil || o.IsPublic == nil {
		return nil, false
	}
	return o.IsPublic, true
}

// HasIsPublic returns a boolean if a field has been set.
func (o *JsonEnvironmentPublishingRequest) HasIsPublic() bool {
	if o != nil && o.IsPublic != nil {
		return true
	}

	return false
}

// SetIsPublic gets a reference to the given string and assigns it to the IsPublic field.
func (o *JsonEnvironmentPublishingRequest) SetIsPublic(v string) {
	o.IsPublic = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *JsonEnvironmentPublishingRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonEnvironmentPublishingRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *JsonEnvironmentPublishingRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *JsonEnvironmentPublishingRequest) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *JsonEnvironmentPublishingRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonEnvironmentPublishingRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *JsonEnvironmentPublishingRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *JsonEnvironmentPublishingRequest) SetName(v string) {
	o.Name = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *JsonEnvironmentPublishingRequest) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonEnvironmentPublishingRequest) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *JsonEnvironmentPublishingRequest) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *JsonEnvironmentPublishingRequest) SetUuid(v string) {
	o.Uuid = &v
}

// GetIpfsHash returns the IpfsHash field value if set, zero value otherwise.
func (o *JsonEnvironmentPublishingRequest) GetIpfsHash() string {
	if o == nil || o.IpfsHash == nil {
		var ret string
		return ret
	}
	return *o.IpfsHash
}

// GetIpfsHashOk returns a tuple with the IpfsHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonEnvironmentPublishingRequest) GetIpfsHashOk() (*string, bool) {
	if o == nil || o.IpfsHash == nil {
		return nil, false
	}
	return o.IpfsHash, true
}

// HasIpfsHash returns a boolean if a field has been set.
func (o *JsonEnvironmentPublishingRequest) HasIpfsHash() bool {
	if o != nil && o.IpfsHash != nil {
		return true
	}

	return false
}

// SetIpfsHash gets a reference to the given string and assigns it to the IpfsHash field.
func (o *JsonEnvironmentPublishingRequest) SetIpfsHash(v string) {
	o.IpfsHash = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *JsonEnvironmentPublishingRequest) GetSize() string {
	if o == nil || o.Size == nil {
		var ret string
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonEnvironmentPublishingRequest) GetSizeOk() (*string, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *JsonEnvironmentPublishingRequest) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given string and assigns it to the Size field.
func (o *JsonEnvironmentPublishingRequest) SetSize(v string) {
	o.Size = &v
}

func (o JsonEnvironmentPublishingRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IpfsPrimaryPeer != nil {
		toSerialize["ipfsPrimaryPeer"] = o.IpfsPrimaryPeer
	}
	if o.Packages != nil {
		toSerialize["packages"] = o.Packages
	}
	if o.Content != nil {
		toSerialize["content"] = o.Content
	}
	if o.IsPublic != nil {
		toSerialize["isPublic"] = o.IsPublic
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Uuid != nil {
		toSerialize["uuid"] = o.Uuid
	}
	if o.IpfsHash != nil {
		toSerialize["ipfsHash"] = o.IpfsHash
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	return json.Marshal(toSerialize)
}

type NullableJsonEnvironmentPublishingRequest struct {
	value *JsonEnvironmentPublishingRequest
	isSet bool
}

func (v NullableJsonEnvironmentPublishingRequest) Get() *JsonEnvironmentPublishingRequest {
	return v.value
}

func (v *NullableJsonEnvironmentPublishingRequest) Set(val *JsonEnvironmentPublishingRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableJsonEnvironmentPublishingRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableJsonEnvironmentPublishingRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJsonEnvironmentPublishingRequest(val *JsonEnvironmentPublishingRequest) *NullableJsonEnvironmentPublishingRequest {
	return &NullableJsonEnvironmentPublishingRequest{value: val, isSet: true}
}

func (v NullableJsonEnvironmentPublishingRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJsonEnvironmentPublishingRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


