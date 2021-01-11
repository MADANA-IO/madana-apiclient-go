/*
 * madana-api
 *
 * <h1>Using the madana-api</h1>        <p>This documentation contains a Quickstart Guide, relating client functionality and information about the available         endpoints and used datamodels.   </p>       <p> The madana-api and its implementations are still in heavy development. This means that there may be problems in our protocols, or there may be mistakes in our implementations. We take security vulnerabilities very seriously. If you discover a security issue, please bring it to our attention right away! If you find a vulnerability that may affect live deployments -- for example, by exposing a remote execution exploit -- please send your report privately to info@madana.io. Please DO NOT file a public issue. If the issue is a protocol weakness that cannot be immediately exploited or something not yet deployed, just discuss it openly   </p>   <br>   <p> Note: Not all functionality might be acessible without having accquired and api-license token. For more information visit <a href=\"https://www.madana.io\">www.madana.io</a> </p>       <br>
 *
 * API version: 0.5.0-master.23
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package madanaapiclient

import (
	"encoding/json"
)

// JsonEnvironment 
type JsonEnvironment struct {
	// 
	Content *[]string `json:"content,omitempty"`
	// 
	RootHashOffset *string `json:"rootHashOffset,omitempty"`
	// 
	Published *bool `json:"published,omitempty"`
	// 
	Packages *[]string `json:"packages,omitempty"`
	// 
	Description *string `json:"description,omitempty"`
	// 
	Uuid *string `json:"uuid,omitempty"`
	// 
	Roothash *string `json:"roothash,omitempty"`
	// 
	IpfsHash *string `json:"ipfsHash,omitempty"`
	// 
	Name *string `json:"name,omitempty"`
	DefaultRunConfiguration *JsonRunConfig `json:"defaultRunConfiguration,omitempty"`
	// 
	Size *string `json:"size,omitempty"`
}

// NewJsonEnvironment instantiates a new JsonEnvironment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJsonEnvironment() *JsonEnvironment {
	this := JsonEnvironment{}
	return &this
}

// NewJsonEnvironmentWithDefaults instantiates a new JsonEnvironment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJsonEnvironmentWithDefaults() *JsonEnvironment {
	this := JsonEnvironment{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *JsonEnvironment) GetContent() []string {
	if o == nil || o.Content == nil {
		var ret []string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonEnvironment) GetContentOk() (*[]string, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *JsonEnvironment) HasContent() bool {
	if o != nil && o.Content != nil {
		return true
	}

	return false
}

// SetContent gets a reference to the given []string and assigns it to the Content field.
func (o *JsonEnvironment) SetContent(v []string) {
	o.Content = &v
}

// GetRootHashOffset returns the RootHashOffset field value if set, zero value otherwise.
func (o *JsonEnvironment) GetRootHashOffset() string {
	if o == nil || o.RootHashOffset == nil {
		var ret string
		return ret
	}
	return *o.RootHashOffset
}

// GetRootHashOffsetOk returns a tuple with the RootHashOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonEnvironment) GetRootHashOffsetOk() (*string, bool) {
	if o == nil || o.RootHashOffset == nil {
		return nil, false
	}
	return o.RootHashOffset, true
}

// HasRootHashOffset returns a boolean if a field has been set.
func (o *JsonEnvironment) HasRootHashOffset() bool {
	if o != nil && o.RootHashOffset != nil {
		return true
	}

	return false
}

// SetRootHashOffset gets a reference to the given string and assigns it to the RootHashOffset field.
func (o *JsonEnvironment) SetRootHashOffset(v string) {
	o.RootHashOffset = &v
}

// GetPublished returns the Published field value if set, zero value otherwise.
func (o *JsonEnvironment) GetPublished() bool {
	if o == nil || o.Published == nil {
		var ret bool
		return ret
	}
	return *o.Published
}

// GetPublishedOk returns a tuple with the Published field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonEnvironment) GetPublishedOk() (*bool, bool) {
	if o == nil || o.Published == nil {
		return nil, false
	}
	return o.Published, true
}

// HasPublished returns a boolean if a field has been set.
func (o *JsonEnvironment) HasPublished() bool {
	if o != nil && o.Published != nil {
		return true
	}

	return false
}

// SetPublished gets a reference to the given bool and assigns it to the Published field.
func (o *JsonEnvironment) SetPublished(v bool) {
	o.Published = &v
}

// GetPackages returns the Packages field value if set, zero value otherwise.
func (o *JsonEnvironment) GetPackages() []string {
	if o == nil || o.Packages == nil {
		var ret []string
		return ret
	}
	return *o.Packages
}

// GetPackagesOk returns a tuple with the Packages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonEnvironment) GetPackagesOk() (*[]string, bool) {
	if o == nil || o.Packages == nil {
		return nil, false
	}
	return o.Packages, true
}

// HasPackages returns a boolean if a field has been set.
func (o *JsonEnvironment) HasPackages() bool {
	if o != nil && o.Packages != nil {
		return true
	}

	return false
}

// SetPackages gets a reference to the given []string and assigns it to the Packages field.
func (o *JsonEnvironment) SetPackages(v []string) {
	o.Packages = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *JsonEnvironment) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonEnvironment) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *JsonEnvironment) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *JsonEnvironment) SetDescription(v string) {
	o.Description = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *JsonEnvironment) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonEnvironment) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *JsonEnvironment) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *JsonEnvironment) SetUuid(v string) {
	o.Uuid = &v
}

// GetRoothash returns the Roothash field value if set, zero value otherwise.
func (o *JsonEnvironment) GetRoothash() string {
	if o == nil || o.Roothash == nil {
		var ret string
		return ret
	}
	return *o.Roothash
}

// GetRoothashOk returns a tuple with the Roothash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonEnvironment) GetRoothashOk() (*string, bool) {
	if o == nil || o.Roothash == nil {
		return nil, false
	}
	return o.Roothash, true
}

// HasRoothash returns a boolean if a field has been set.
func (o *JsonEnvironment) HasRoothash() bool {
	if o != nil && o.Roothash != nil {
		return true
	}

	return false
}

// SetRoothash gets a reference to the given string and assigns it to the Roothash field.
func (o *JsonEnvironment) SetRoothash(v string) {
	o.Roothash = &v
}

// GetIpfsHash returns the IpfsHash field value if set, zero value otherwise.
func (o *JsonEnvironment) GetIpfsHash() string {
	if o == nil || o.IpfsHash == nil {
		var ret string
		return ret
	}
	return *o.IpfsHash
}

// GetIpfsHashOk returns a tuple with the IpfsHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonEnvironment) GetIpfsHashOk() (*string, bool) {
	if o == nil || o.IpfsHash == nil {
		return nil, false
	}
	return o.IpfsHash, true
}

// HasIpfsHash returns a boolean if a field has been set.
func (o *JsonEnvironment) HasIpfsHash() bool {
	if o != nil && o.IpfsHash != nil {
		return true
	}

	return false
}

// SetIpfsHash gets a reference to the given string and assigns it to the IpfsHash field.
func (o *JsonEnvironment) SetIpfsHash(v string) {
	o.IpfsHash = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *JsonEnvironment) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonEnvironment) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *JsonEnvironment) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *JsonEnvironment) SetName(v string) {
	o.Name = &v
}

// GetDefaultRunConfiguration returns the DefaultRunConfiguration field value if set, zero value otherwise.
func (o *JsonEnvironment) GetDefaultRunConfiguration() JsonRunConfig {
	if o == nil || o.DefaultRunConfiguration == nil {
		var ret JsonRunConfig
		return ret
	}
	return *o.DefaultRunConfiguration
}

// GetDefaultRunConfigurationOk returns a tuple with the DefaultRunConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonEnvironment) GetDefaultRunConfigurationOk() (*JsonRunConfig, bool) {
	if o == nil || o.DefaultRunConfiguration == nil {
		return nil, false
	}
	return o.DefaultRunConfiguration, true
}

// HasDefaultRunConfiguration returns a boolean if a field has been set.
func (o *JsonEnvironment) HasDefaultRunConfiguration() bool {
	if o != nil && o.DefaultRunConfiguration != nil {
		return true
	}

	return false
}

// SetDefaultRunConfiguration gets a reference to the given JsonRunConfig and assigns it to the DefaultRunConfiguration field.
func (o *JsonEnvironment) SetDefaultRunConfiguration(v JsonRunConfig) {
	o.DefaultRunConfiguration = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *JsonEnvironment) GetSize() string {
	if o == nil || o.Size == nil {
		var ret string
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonEnvironment) GetSizeOk() (*string, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *JsonEnvironment) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given string and assigns it to the Size field.
func (o *JsonEnvironment) SetSize(v string) {
	o.Size = &v
}

func (o JsonEnvironment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Content != nil {
		toSerialize["content"] = o.Content
	}
	if o.RootHashOffset != nil {
		toSerialize["rootHashOffset"] = o.RootHashOffset
	}
	if o.Published != nil {
		toSerialize["published"] = o.Published
	}
	if o.Packages != nil {
		toSerialize["packages"] = o.Packages
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Uuid != nil {
		toSerialize["uuid"] = o.Uuid
	}
	if o.Roothash != nil {
		toSerialize["roothash"] = o.Roothash
	}
	if o.IpfsHash != nil {
		toSerialize["ipfsHash"] = o.IpfsHash
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.DefaultRunConfiguration != nil {
		toSerialize["defaultRunConfiguration"] = o.DefaultRunConfiguration
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	return json.Marshal(toSerialize)
}

type NullableJsonEnvironment struct {
	value *JsonEnvironment
	isSet bool
}

func (v NullableJsonEnvironment) Get() *JsonEnvironment {
	return v.value
}

func (v *NullableJsonEnvironment) Set(val *JsonEnvironment) {
	v.value = val
	v.isSet = true
}

func (v NullableJsonEnvironment) IsSet() bool {
	return v.isSet
}

func (v *NullableJsonEnvironment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJsonEnvironment(val *JsonEnvironment) *NullableJsonEnvironment {
	return &NullableJsonEnvironment{value: val, isSet: true}
}

func (v NullableJsonEnvironment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJsonEnvironment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


