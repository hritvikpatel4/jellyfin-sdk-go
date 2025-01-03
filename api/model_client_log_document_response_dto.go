/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.10.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ClientLogDocumentResponseDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClientLogDocumentResponseDto{}

// ClientLogDocumentResponseDto Client log document response dto.
type ClientLogDocumentResponseDto struct {
	// Gets the resulting filename.
	FileName *string `json:"FileName,omitempty"`
}

// NewClientLogDocumentResponseDto instantiates a new ClientLogDocumentResponseDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientLogDocumentResponseDto() *ClientLogDocumentResponseDto {
	this := ClientLogDocumentResponseDto{}
	return &this
}

// NewClientLogDocumentResponseDtoWithDefaults instantiates a new ClientLogDocumentResponseDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientLogDocumentResponseDtoWithDefaults() *ClientLogDocumentResponseDto {
	this := ClientLogDocumentResponseDto{}
	return &this
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *ClientLogDocumentResponseDto) GetFileName() string {
	if o == nil || IsNil(o.FileName) {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientLogDocumentResponseDto) GetFileNameOk() (*string, bool) {
	if o == nil || IsNil(o.FileName) {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *ClientLogDocumentResponseDto) HasFileName() bool {
	if o != nil && !IsNil(o.FileName) {
		return true
	}

	return false
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *ClientLogDocumentResponseDto) SetFileName(v string) {
	o.FileName = &v
}

func (o ClientLogDocumentResponseDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientLogDocumentResponseDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FileName) {
		toSerialize["FileName"] = o.FileName
	}
	return toSerialize, nil
}

type NullableClientLogDocumentResponseDto struct {
	value *ClientLogDocumentResponseDto
	isSet bool
}

func (v NullableClientLogDocumentResponseDto) Get() *ClientLogDocumentResponseDto {
	return v.value
}

func (v *NullableClientLogDocumentResponseDto) Set(val *ClientLogDocumentResponseDto) {
	v.value = val
	v.isSet = true
}

func (v NullableClientLogDocumentResponseDto) IsSet() bool {
	return v.isSet
}

func (v *NullableClientLogDocumentResponseDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientLogDocumentResponseDto(val *ClientLogDocumentResponseDto) *NullableClientLogDocumentResponseDto {
	return &NullableClientLogDocumentResponseDto{value: val, isSet: true}
}

func (v NullableClientLogDocumentResponseDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientLogDocumentResponseDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
