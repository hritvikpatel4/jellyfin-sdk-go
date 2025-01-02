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

// checks if the ImageInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImageInfo{}

// ImageInfo Class ImageInfo.
type ImageInfo struct {
	// Gets or sets the type of the image.
	ImageType *ImageType `json:"ImageType,omitempty"`
	// Gets or sets the index of the image.
	ImageIndex NullableInt32 `json:"ImageIndex,omitempty"`
	// Gets or sets the image tag.
	ImageTag NullableString `json:"ImageTag,omitempty"`
	// Gets or sets the path.
	Path NullableString `json:"Path,omitempty"`
	// Gets or sets the blurhash.
	BlurHash NullableString `json:"BlurHash,omitempty"`
	// Gets or sets the height.
	Height NullableInt32 `json:"Height,omitempty"`
	// Gets or sets the width.
	Width NullableInt32 `json:"Width,omitempty"`
	// Gets or sets the size.
	Size *int64 `json:"Size,omitempty"`
}

// NewImageInfo instantiates a new ImageInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageInfo() *ImageInfo {
	this := ImageInfo{}
	return &this
}

// NewImageInfoWithDefaults instantiates a new ImageInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageInfoWithDefaults() *ImageInfo {
	this := ImageInfo{}
	return &this
}

// GetImageType returns the ImageType field value if set, zero value otherwise.
func (o *ImageInfo) GetImageType() ImageType {
	if o == nil || IsNil(o.ImageType) {
		var ret ImageType
		return ret
	}
	return *o.ImageType
}

// GetImageTypeOk returns a tuple with the ImageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageInfo) GetImageTypeOk() (*ImageType, bool) {
	if o == nil || IsNil(o.ImageType) {
		return nil, false
	}
	return o.ImageType, true
}

// HasImageType returns a boolean if a field has been set.
func (o *ImageInfo) HasImageType() bool {
	if o != nil && !IsNil(o.ImageType) {
		return true
	}

	return false
}

// SetImageType gets a reference to the given ImageType and assigns it to the ImageType field.
func (o *ImageInfo) SetImageType(v ImageType) {
	o.ImageType = &v
}

// GetImageIndex returns the ImageIndex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImageInfo) GetImageIndex() int32 {
	if o == nil || IsNil(o.ImageIndex.Get()) {
		var ret int32
		return ret
	}
	return *o.ImageIndex.Get()
}

// GetImageIndexOk returns a tuple with the ImageIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImageInfo) GetImageIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImageIndex.Get(), o.ImageIndex.IsSet()
}

// HasImageIndex returns a boolean if a field has been set.
func (o *ImageInfo) HasImageIndex() bool {
	if o != nil && o.ImageIndex.IsSet() {
		return true
	}

	return false
}

// SetImageIndex gets a reference to the given NullableInt32 and assigns it to the ImageIndex field.
func (o *ImageInfo) SetImageIndex(v int32) {
	o.ImageIndex.Set(&v)
}
// SetImageIndexNil sets the value for ImageIndex to be an explicit nil
func (o *ImageInfo) SetImageIndexNil() {
	o.ImageIndex.Set(nil)
}

// UnsetImageIndex ensures that no value is present for ImageIndex, not even an explicit nil
func (o *ImageInfo) UnsetImageIndex() {
	o.ImageIndex.Unset()
}

// GetImageTag returns the ImageTag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImageInfo) GetImageTag() string {
	if o == nil || IsNil(o.ImageTag.Get()) {
		var ret string
		return ret
	}
	return *o.ImageTag.Get()
}

// GetImageTagOk returns a tuple with the ImageTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImageInfo) GetImageTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImageTag.Get(), o.ImageTag.IsSet()
}

// HasImageTag returns a boolean if a field has been set.
func (o *ImageInfo) HasImageTag() bool {
	if o != nil && o.ImageTag.IsSet() {
		return true
	}

	return false
}

// SetImageTag gets a reference to the given NullableString and assigns it to the ImageTag field.
func (o *ImageInfo) SetImageTag(v string) {
	o.ImageTag.Set(&v)
}
// SetImageTagNil sets the value for ImageTag to be an explicit nil
func (o *ImageInfo) SetImageTagNil() {
	o.ImageTag.Set(nil)
}

// UnsetImageTag ensures that no value is present for ImageTag, not even an explicit nil
func (o *ImageInfo) UnsetImageTag() {
	o.ImageTag.Unset()
}

// GetPath returns the Path field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImageInfo) GetPath() string {
	if o == nil || IsNil(o.Path.Get()) {
		var ret string
		return ret
	}
	return *o.Path.Get()
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImageInfo) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Path.Get(), o.Path.IsSet()
}

// HasPath returns a boolean if a field has been set.
func (o *ImageInfo) HasPath() bool {
	if o != nil && o.Path.IsSet() {
		return true
	}

	return false
}

// SetPath gets a reference to the given NullableString and assigns it to the Path field.
func (o *ImageInfo) SetPath(v string) {
	o.Path.Set(&v)
}
// SetPathNil sets the value for Path to be an explicit nil
func (o *ImageInfo) SetPathNil() {
	o.Path.Set(nil)
}

// UnsetPath ensures that no value is present for Path, not even an explicit nil
func (o *ImageInfo) UnsetPath() {
	o.Path.Unset()
}

// GetBlurHash returns the BlurHash field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImageInfo) GetBlurHash() string {
	if o == nil || IsNil(o.BlurHash.Get()) {
		var ret string
		return ret
	}
	return *o.BlurHash.Get()
}

// GetBlurHashOk returns a tuple with the BlurHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImageInfo) GetBlurHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BlurHash.Get(), o.BlurHash.IsSet()
}

// HasBlurHash returns a boolean if a field has been set.
func (o *ImageInfo) HasBlurHash() bool {
	if o != nil && o.BlurHash.IsSet() {
		return true
	}

	return false
}

// SetBlurHash gets a reference to the given NullableString and assigns it to the BlurHash field.
func (o *ImageInfo) SetBlurHash(v string) {
	o.BlurHash.Set(&v)
}
// SetBlurHashNil sets the value for BlurHash to be an explicit nil
func (o *ImageInfo) SetBlurHashNil() {
	o.BlurHash.Set(nil)
}

// UnsetBlurHash ensures that no value is present for BlurHash, not even an explicit nil
func (o *ImageInfo) UnsetBlurHash() {
	o.BlurHash.Unset()
}

// GetHeight returns the Height field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImageInfo) GetHeight() int32 {
	if o == nil || IsNil(o.Height.Get()) {
		var ret int32
		return ret
	}
	return *o.Height.Get()
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImageInfo) GetHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Height.Get(), o.Height.IsSet()
}

// HasHeight returns a boolean if a field has been set.
func (o *ImageInfo) HasHeight() bool {
	if o != nil && o.Height.IsSet() {
		return true
	}

	return false
}

// SetHeight gets a reference to the given NullableInt32 and assigns it to the Height field.
func (o *ImageInfo) SetHeight(v int32) {
	o.Height.Set(&v)
}
// SetHeightNil sets the value for Height to be an explicit nil
func (o *ImageInfo) SetHeightNil() {
	o.Height.Set(nil)
}

// UnsetHeight ensures that no value is present for Height, not even an explicit nil
func (o *ImageInfo) UnsetHeight() {
	o.Height.Unset()
}

// GetWidth returns the Width field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImageInfo) GetWidth() int32 {
	if o == nil || IsNil(o.Width.Get()) {
		var ret int32
		return ret
	}
	return *o.Width.Get()
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImageInfo) GetWidthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Width.Get(), o.Width.IsSet()
}

// HasWidth returns a boolean if a field has been set.
func (o *ImageInfo) HasWidth() bool {
	if o != nil && o.Width.IsSet() {
		return true
	}

	return false
}

// SetWidth gets a reference to the given NullableInt32 and assigns it to the Width field.
func (o *ImageInfo) SetWidth(v int32) {
	o.Width.Set(&v)
}
// SetWidthNil sets the value for Width to be an explicit nil
func (o *ImageInfo) SetWidthNil() {
	o.Width.Set(nil)
}

// UnsetWidth ensures that no value is present for Width, not even an explicit nil
func (o *ImageInfo) UnsetWidth() {
	o.Width.Unset()
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *ImageInfo) GetSize() int64 {
	if o == nil || IsNil(o.Size) {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageInfo) GetSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *ImageInfo) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *ImageInfo) SetSize(v int64) {
	o.Size = &v
}

func (o ImageInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImageInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ImageType) {
		toSerialize["ImageType"] = o.ImageType
	}
	if o.ImageIndex.IsSet() {
		toSerialize["ImageIndex"] = o.ImageIndex.Get()
	}
	if o.ImageTag.IsSet() {
		toSerialize["ImageTag"] = o.ImageTag.Get()
	}
	if o.Path.IsSet() {
		toSerialize["Path"] = o.Path.Get()
	}
	if o.BlurHash.IsSet() {
		toSerialize["BlurHash"] = o.BlurHash.Get()
	}
	if o.Height.IsSet() {
		toSerialize["Height"] = o.Height.Get()
	}
	if o.Width.IsSet() {
		toSerialize["Width"] = o.Width.Get()
	}
	if !IsNil(o.Size) {
		toSerialize["Size"] = o.Size
	}
	return toSerialize, nil
}

type NullableImageInfo struct {
	value *ImageInfo
	isSet bool
}

func (v NullableImageInfo) Get() *ImageInfo {
	return v.value
}

func (v *NullableImageInfo) Set(val *ImageInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableImageInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableImageInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageInfo(val *ImageInfo) *NullableImageInfo {
	return &NullableImageInfo{value: val, isSet: true}
}

func (v NullableImageInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


