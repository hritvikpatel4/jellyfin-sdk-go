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

// checks if the LibraryTypeOptionsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LibraryTypeOptionsDto{}

// LibraryTypeOptionsDto Library type options dto.
type LibraryTypeOptionsDto struct {
	// Gets or sets the type.
	Type NullableString `json:"Type,omitempty"`
	// Gets or sets the metadata fetchers.
	MetadataFetchers []LibraryOptionInfoDto `json:"MetadataFetchers,omitempty"`
	// Gets or sets the image fetchers.
	ImageFetchers []LibraryOptionInfoDto `json:"ImageFetchers,omitempty"`
	// Gets or sets the supported image types.
	SupportedImageTypes []ImageType `json:"SupportedImageTypes,omitempty"`
	// Gets or sets the default image options.
	DefaultImageOptions []ImageOption `json:"DefaultImageOptions,omitempty"`
}

// NewLibraryTypeOptionsDto instantiates a new LibraryTypeOptionsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLibraryTypeOptionsDto() *LibraryTypeOptionsDto {
	this := LibraryTypeOptionsDto{}
	return &this
}

// NewLibraryTypeOptionsDtoWithDefaults instantiates a new LibraryTypeOptionsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLibraryTypeOptionsDtoWithDefaults() *LibraryTypeOptionsDto {
	this := LibraryTypeOptionsDto{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LibraryTypeOptionsDto) GetType() string {
	if o == nil || IsNil(o.Type.Get()) {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LibraryTypeOptionsDto) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *LibraryTypeOptionsDto) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *LibraryTypeOptionsDto) SetType(v string) {
	o.Type.Set(&v)
}

// SetTypeNil sets the value for Type to be an explicit nil
func (o *LibraryTypeOptionsDto) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *LibraryTypeOptionsDto) UnsetType() {
	o.Type.Unset()
}

// GetMetadataFetchers returns the MetadataFetchers field value if set, zero value otherwise.
func (o *LibraryTypeOptionsDto) GetMetadataFetchers() []LibraryOptionInfoDto {
	if o == nil || IsNil(o.MetadataFetchers) {
		var ret []LibraryOptionInfoDto
		return ret
	}
	return o.MetadataFetchers
}

// GetMetadataFetchersOk returns a tuple with the MetadataFetchers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LibraryTypeOptionsDto) GetMetadataFetchersOk() ([]LibraryOptionInfoDto, bool) {
	if o == nil || IsNil(o.MetadataFetchers) {
		return nil, false
	}
	return o.MetadataFetchers, true
}

// HasMetadataFetchers returns a boolean if a field has been set.
func (o *LibraryTypeOptionsDto) HasMetadataFetchers() bool {
	if o != nil && !IsNil(o.MetadataFetchers) {
		return true
	}

	return false
}

// SetMetadataFetchers gets a reference to the given []LibraryOptionInfoDto and assigns it to the MetadataFetchers field.
func (o *LibraryTypeOptionsDto) SetMetadataFetchers(v []LibraryOptionInfoDto) {
	o.MetadataFetchers = v
}

// GetImageFetchers returns the ImageFetchers field value if set, zero value otherwise.
func (o *LibraryTypeOptionsDto) GetImageFetchers() []LibraryOptionInfoDto {
	if o == nil || IsNil(o.ImageFetchers) {
		var ret []LibraryOptionInfoDto
		return ret
	}
	return o.ImageFetchers
}

// GetImageFetchersOk returns a tuple with the ImageFetchers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LibraryTypeOptionsDto) GetImageFetchersOk() ([]LibraryOptionInfoDto, bool) {
	if o == nil || IsNil(o.ImageFetchers) {
		return nil, false
	}
	return o.ImageFetchers, true
}

// HasImageFetchers returns a boolean if a field has been set.
func (o *LibraryTypeOptionsDto) HasImageFetchers() bool {
	if o != nil && !IsNil(o.ImageFetchers) {
		return true
	}

	return false
}

// SetImageFetchers gets a reference to the given []LibraryOptionInfoDto and assigns it to the ImageFetchers field.
func (o *LibraryTypeOptionsDto) SetImageFetchers(v []LibraryOptionInfoDto) {
	o.ImageFetchers = v
}

// GetSupportedImageTypes returns the SupportedImageTypes field value if set, zero value otherwise.
func (o *LibraryTypeOptionsDto) GetSupportedImageTypes() []ImageType {
	if o == nil || IsNil(o.SupportedImageTypes) {
		var ret []ImageType
		return ret
	}
	return o.SupportedImageTypes
}

// GetSupportedImageTypesOk returns a tuple with the SupportedImageTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LibraryTypeOptionsDto) GetSupportedImageTypesOk() ([]ImageType, bool) {
	if o == nil || IsNil(o.SupportedImageTypes) {
		return nil, false
	}
	return o.SupportedImageTypes, true
}

// HasSupportedImageTypes returns a boolean if a field has been set.
func (o *LibraryTypeOptionsDto) HasSupportedImageTypes() bool {
	if o != nil && !IsNil(o.SupportedImageTypes) {
		return true
	}

	return false
}

// SetSupportedImageTypes gets a reference to the given []ImageType and assigns it to the SupportedImageTypes field.
func (o *LibraryTypeOptionsDto) SetSupportedImageTypes(v []ImageType) {
	o.SupportedImageTypes = v
}

// GetDefaultImageOptions returns the DefaultImageOptions field value if set, zero value otherwise.
func (o *LibraryTypeOptionsDto) GetDefaultImageOptions() []ImageOption {
	if o == nil || IsNil(o.DefaultImageOptions) {
		var ret []ImageOption
		return ret
	}
	return o.DefaultImageOptions
}

// GetDefaultImageOptionsOk returns a tuple with the DefaultImageOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LibraryTypeOptionsDto) GetDefaultImageOptionsOk() ([]ImageOption, bool) {
	if o == nil || IsNil(o.DefaultImageOptions) {
		return nil, false
	}
	return o.DefaultImageOptions, true
}

// HasDefaultImageOptions returns a boolean if a field has been set.
func (o *LibraryTypeOptionsDto) HasDefaultImageOptions() bool {
	if o != nil && !IsNil(o.DefaultImageOptions) {
		return true
	}

	return false
}

// SetDefaultImageOptions gets a reference to the given []ImageOption and assigns it to the DefaultImageOptions field.
func (o *LibraryTypeOptionsDto) SetDefaultImageOptions(v []ImageOption) {
	o.DefaultImageOptions = v
}

func (o LibraryTypeOptionsDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LibraryTypeOptionsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type.IsSet() {
		toSerialize["Type"] = o.Type.Get()
	}
	if !IsNil(o.MetadataFetchers) {
		toSerialize["MetadataFetchers"] = o.MetadataFetchers
	}
	if !IsNil(o.ImageFetchers) {
		toSerialize["ImageFetchers"] = o.ImageFetchers
	}
	if !IsNil(o.SupportedImageTypes) {
		toSerialize["SupportedImageTypes"] = o.SupportedImageTypes
	}
	if !IsNil(o.DefaultImageOptions) {
		toSerialize["DefaultImageOptions"] = o.DefaultImageOptions
	}
	return toSerialize, nil
}

type NullableLibraryTypeOptionsDto struct {
	value *LibraryTypeOptionsDto
	isSet bool
}

func (v NullableLibraryTypeOptionsDto) Get() *LibraryTypeOptionsDto {
	return v.value
}

func (v *NullableLibraryTypeOptionsDto) Set(val *LibraryTypeOptionsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableLibraryTypeOptionsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableLibraryTypeOptionsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLibraryTypeOptionsDto(val *LibraryTypeOptionsDto) *NullableLibraryTypeOptionsDto {
	return &NullableLibraryTypeOptionsDto{value: val, isSet: true}
}

func (v NullableLibraryTypeOptionsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLibraryTypeOptionsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
