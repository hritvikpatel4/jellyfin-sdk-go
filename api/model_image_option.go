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

// checks if the ImageOption type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImageOption{}

// ImageOption struct for ImageOption
type ImageOption struct {
	// Gets or sets the type.
	Type *ImageType `json:"Type,omitempty"`
	// Gets or sets the limit.
	Limit *int32 `json:"Limit,omitempty"`
	// Gets or sets the minimum width.
	MinWidth *int32 `json:"MinWidth,omitempty"`
}

// NewImageOption instantiates a new ImageOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageOption() *ImageOption {
	this := ImageOption{}
	return &this
}

// NewImageOptionWithDefaults instantiates a new ImageOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageOptionWithDefaults() *ImageOption {
	this := ImageOption{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ImageOption) GetType() ImageType {
	if o == nil || IsNil(o.Type) {
		var ret ImageType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageOption) GetTypeOk() (*ImageType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ImageOption) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given ImageType and assigns it to the Type field.
func (o *ImageOption) SetType(v ImageType) {
	o.Type = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *ImageOption) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageOption) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *ImageOption) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *ImageOption) SetLimit(v int32) {
	o.Limit = &v
}

// GetMinWidth returns the MinWidth field value if set, zero value otherwise.
func (o *ImageOption) GetMinWidth() int32 {
	if o == nil || IsNil(o.MinWidth) {
		var ret int32
		return ret
	}
	return *o.MinWidth
}

// GetMinWidthOk returns a tuple with the MinWidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageOption) GetMinWidthOk() (*int32, bool) {
	if o == nil || IsNil(o.MinWidth) {
		return nil, false
	}
	return o.MinWidth, true
}

// HasMinWidth returns a boolean if a field has been set.
func (o *ImageOption) HasMinWidth() bool {
	if o != nil && !IsNil(o.MinWidth) {
		return true
	}

	return false
}

// SetMinWidth gets a reference to the given int32 and assigns it to the MinWidth field.
func (o *ImageOption) SetMinWidth(v int32) {
	o.MinWidth = &v
}

func (o ImageOption) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImageOption) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["Type"] = o.Type
	}
	if !IsNil(o.Limit) {
		toSerialize["Limit"] = o.Limit
	}
	if !IsNil(o.MinWidth) {
		toSerialize["MinWidth"] = o.MinWidth
	}
	return toSerialize, nil
}

type NullableImageOption struct {
	value *ImageOption
	isSet bool
}

func (v NullableImageOption) Get() *ImageOption {
	return v.value
}

func (v *NullableImageOption) Set(val *ImageOption) {
	v.value = val
	v.isSet = true
}

func (v NullableImageOption) IsSet() bool {
	return v.isSet
}

func (v *NullableImageOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageOption(val *ImageOption) *NullableImageOption {
	return &NullableImageOption{value: val, isSet: true}
}

func (v NullableImageOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

