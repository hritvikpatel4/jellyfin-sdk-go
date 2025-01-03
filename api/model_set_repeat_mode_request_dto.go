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

// checks if the SetRepeatModeRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetRepeatModeRequestDto{}

// SetRepeatModeRequestDto Class SetRepeatModeRequestDto.
type SetRepeatModeRequestDto struct {
	// Gets or sets the repeat mode.
	Mode *GroupRepeatMode `json:"Mode,omitempty"`
}

// NewSetRepeatModeRequestDto instantiates a new SetRepeatModeRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetRepeatModeRequestDto() *SetRepeatModeRequestDto {
	this := SetRepeatModeRequestDto{}
	return &this
}

// NewSetRepeatModeRequestDtoWithDefaults instantiates a new SetRepeatModeRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetRepeatModeRequestDtoWithDefaults() *SetRepeatModeRequestDto {
	this := SetRepeatModeRequestDto{}
	return &this
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *SetRepeatModeRequestDto) GetMode() GroupRepeatMode {
	if o == nil || IsNil(o.Mode) {
		var ret GroupRepeatMode
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetRepeatModeRequestDto) GetModeOk() (*GroupRepeatMode, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *SetRepeatModeRequestDto) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given GroupRepeatMode and assigns it to the Mode field.
func (o *SetRepeatModeRequestDto) SetMode(v GroupRepeatMode) {
	o.Mode = &v
}

func (o SetRepeatModeRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetRepeatModeRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Mode) {
		toSerialize["Mode"] = o.Mode
	}
	return toSerialize, nil
}

type NullableSetRepeatModeRequestDto struct {
	value *SetRepeatModeRequestDto
	isSet bool
}

func (v NullableSetRepeatModeRequestDto) Get() *SetRepeatModeRequestDto {
	return v.value
}

func (v *NullableSetRepeatModeRequestDto) Set(val *SetRepeatModeRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableSetRepeatModeRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableSetRepeatModeRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetRepeatModeRequestDto(val *SetRepeatModeRequestDto) *NullableSetRepeatModeRequestDto {
	return &NullableSetRepeatModeRequestDto{value: val, isSet: true}
}

func (v NullableSetRepeatModeRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetRepeatModeRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
