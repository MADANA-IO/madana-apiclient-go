/*
 * madana-api
 *
 * <h1>Using the madana-api</h1>        <p>This documentation contains a Quickstart Guide, relating client functionality and information about the available         endpoints and used datamodels.   </p>       <p> The madana-api and its implementations are still in heavy development. This means that there may be problems in our protocols, or there may be mistakes in our implementations. We take security vulnerabilities very seriously. If you discover a security issue, please bring it to our attention right away! If you find a vulnerability that may affect live deployments -- for example, by exposing a remote execution exploit -- please send your report privately to info@madana.io. Please DO NOT file a public issue. If the issue is a protocol weakness that cannot be immediately exploited or something not yet deployed, just discuss it openly   </p>   <br>   <p> Note: Not all functionality might be acessible without having accquired and api-license token. For more information visit <a href=\"https://www.madana.io\">www.madana.io</a> </p>       <br>
 *
 * API version: 0.5.0-master.52
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package madanaapiclient

import (
	"encoding/json"
)

// JsonProcess 
type JsonProcess struct {
	// 
	OutputStream *map[string]interface{} `json:"outputStream,omitempty"`
	// 
	Alive *bool `json:"alive,omitempty"`
	// 
	InputStream *map[string]interface{} `json:"inputStream,omitempty"`
	// 
	ErrorStream *map[string]interface{} `json:"errorStream,omitempty"`
}

// NewJsonProcess instantiates a new JsonProcess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJsonProcess() *JsonProcess {
	this := JsonProcess{}
	return &this
}

// NewJsonProcessWithDefaults instantiates a new JsonProcess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJsonProcessWithDefaults() *JsonProcess {
	this := JsonProcess{}
	return &this
}

// GetOutputStream returns the OutputStream field value if set, zero value otherwise.
func (o *JsonProcess) GetOutputStream() map[string]interface{} {
	if o == nil || o.OutputStream == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.OutputStream
}

// GetOutputStreamOk returns a tuple with the OutputStream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonProcess) GetOutputStreamOk() (*map[string]interface{}, bool) {
	if o == nil || o.OutputStream == nil {
		return nil, false
	}
	return o.OutputStream, true
}

// HasOutputStream returns a boolean if a field has been set.
func (o *JsonProcess) HasOutputStream() bool {
	if o != nil && o.OutputStream != nil {
		return true
	}

	return false
}

// SetOutputStream gets a reference to the given map[string]interface{} and assigns it to the OutputStream field.
func (o *JsonProcess) SetOutputStream(v map[string]interface{}) {
	o.OutputStream = &v
}

// GetAlive returns the Alive field value if set, zero value otherwise.
func (o *JsonProcess) GetAlive() bool {
	if o == nil || o.Alive == nil {
		var ret bool
		return ret
	}
	return *o.Alive
}

// GetAliveOk returns a tuple with the Alive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonProcess) GetAliveOk() (*bool, bool) {
	if o == nil || o.Alive == nil {
		return nil, false
	}
	return o.Alive, true
}

// HasAlive returns a boolean if a field has been set.
func (o *JsonProcess) HasAlive() bool {
	if o != nil && o.Alive != nil {
		return true
	}

	return false
}

// SetAlive gets a reference to the given bool and assigns it to the Alive field.
func (o *JsonProcess) SetAlive(v bool) {
	o.Alive = &v
}

// GetInputStream returns the InputStream field value if set, zero value otherwise.
func (o *JsonProcess) GetInputStream() map[string]interface{} {
	if o == nil || o.InputStream == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.InputStream
}

// GetInputStreamOk returns a tuple with the InputStream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonProcess) GetInputStreamOk() (*map[string]interface{}, bool) {
	if o == nil || o.InputStream == nil {
		return nil, false
	}
	return o.InputStream, true
}

// HasInputStream returns a boolean if a field has been set.
func (o *JsonProcess) HasInputStream() bool {
	if o != nil && o.InputStream != nil {
		return true
	}

	return false
}

// SetInputStream gets a reference to the given map[string]interface{} and assigns it to the InputStream field.
func (o *JsonProcess) SetInputStream(v map[string]interface{}) {
	o.InputStream = &v
}

// GetErrorStream returns the ErrorStream field value if set, zero value otherwise.
func (o *JsonProcess) GetErrorStream() map[string]interface{} {
	if o == nil || o.ErrorStream == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.ErrorStream
}

// GetErrorStreamOk returns a tuple with the ErrorStream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonProcess) GetErrorStreamOk() (*map[string]interface{}, bool) {
	if o == nil || o.ErrorStream == nil {
		return nil, false
	}
	return o.ErrorStream, true
}

// HasErrorStream returns a boolean if a field has been set.
func (o *JsonProcess) HasErrorStream() bool {
	if o != nil && o.ErrorStream != nil {
		return true
	}

	return false
}

// SetErrorStream gets a reference to the given map[string]interface{} and assigns it to the ErrorStream field.
func (o *JsonProcess) SetErrorStream(v map[string]interface{}) {
	o.ErrorStream = &v
}

func (o JsonProcess) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.OutputStream != nil {
		toSerialize["outputStream"] = o.OutputStream
	}
	if o.Alive != nil {
		toSerialize["alive"] = o.Alive
	}
	if o.InputStream != nil {
		toSerialize["inputStream"] = o.InputStream
	}
	if o.ErrorStream != nil {
		toSerialize["errorStream"] = o.ErrorStream
	}
	return json.Marshal(toSerialize)
}

type NullableJsonProcess struct {
	value *JsonProcess
	isSet bool
}

func (v NullableJsonProcess) Get() *JsonProcess {
	return v.value
}

func (v *NullableJsonProcess) Set(val *JsonProcess) {
	v.value = val
	v.isSet = true
}

func (v NullableJsonProcess) IsSet() bool {
	return v.isSet
}

func (v *NullableJsonProcess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJsonProcess(val *JsonProcess) *NullableJsonProcess {
	return &NullableJsonProcess{value: val, isSet: true}
}

func (v NullableJsonProcess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJsonProcess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


