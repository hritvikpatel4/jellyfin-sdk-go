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

// checks if the LocalizationOption type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LocalizationOption{}

// LocalizationOption struct for LocalizationOption
type LocalizationOption struct {
	Name NullableString `json:"Name,omitempty"`
	Value NullableString `json:"Value,omitempty"`
}

// NewLocalizationOption instantiates a new LocalizationOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocalizationOption() *LocalizationOption {
	this := LocalizationOption{}
	return &this
}

// NewLocalizationOptionWithDefaults instantiates a new LocalizationOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocalizationOptionWithDefaults() *LocalizationOption {
	this := LocalizationOption{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LocalizationOption) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LocalizationOption) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *LocalizationOption) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *LocalizationOption) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *LocalizationOption) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *LocalizationOption) UnsetName() {
	o.Name.Unset()
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LocalizationOption) GetValue() string {
	if o == nil || IsNil(o.Value.Get()) {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LocalizationOption) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *LocalizationOption) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableString and assigns it to the Value field.
func (o *LocalizationOption) SetValue(v string) {
	o.Value.Set(&v)
}
// SetValueNil sets the value for Value to be an explicit nil
func (o *LocalizationOption) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *LocalizationOption) UnsetValue() {
	o.Value.Unset()
}

func (o LocalizationOption) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LocalizationOption) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["Name"] = o.Name.Get()
	}
	if o.Value.IsSet() {
		toSerialize["Value"] = o.Value.Get()
	}
	return toSerialize, nil
}

type NullableLocalizationOption struct {
	value *LocalizationOption
	isSet bool
}

func (v NullableLocalizationOption) Get() *LocalizationOption {
	return v.value
}

func (v *NullableLocalizationOption) Set(val *LocalizationOption) {
	v.value = val
	v.isSet = true
}

func (v NullableLocalizationOption) IsSet() bool {
	return v.isSet
}

func (v *NullableLocalizationOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocalizationOption(val *LocalizationOption) *NullableLocalizationOption {
	return &NullableLocalizationOption{value: val, isSet: true}
}

func (v NullableLocalizationOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocalizationOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

