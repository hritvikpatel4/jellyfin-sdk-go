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

// checks if the TrickplayOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TrickplayOptions{}

// TrickplayOptions Class TrickplayOptions.
type TrickplayOptions struct {
	// Gets or sets a value indicating whether or not to use HW acceleration.
	EnableHwAcceleration *bool `json:"EnableHwAcceleration,omitempty"`
	// Gets or sets a value indicating whether or not to use HW accelerated MJPEG encoding.
	EnableHwEncoding *bool `json:"EnableHwEncoding,omitempty"`
	// Gets or sets a value indicating whether to only extract key frames.  Significantly faster, but is not compatible with all decoders and/or video files.
	EnableKeyFrameOnlyExtraction *bool `json:"EnableKeyFrameOnlyExtraction,omitempty"`
	// Gets or sets the behavior used by trickplay provider on library scan/update.
	ScanBehavior *TrickplayScanBehavior `json:"ScanBehavior,omitempty"`
	// Gets or sets the process priority for the ffmpeg process.
	ProcessPriority *ProcessPriorityClass `json:"ProcessPriority,omitempty"`
	// Gets or sets the interval, in ms, between each new trickplay image.
	Interval *int32 `json:"Interval,omitempty"`
	// Gets or sets the target width resolutions, in px, to generates preview images for.
	WidthResolutions []int32 `json:"WidthResolutions,omitempty"`
	// Gets or sets number of tile images to allow in X dimension.
	TileWidth *int32 `json:"TileWidth,omitempty"`
	// Gets or sets number of tile images to allow in Y dimension.
	TileHeight *int32 `json:"TileHeight,omitempty"`
	// Gets or sets the ffmpeg output quality level.
	Qscale *int32 `json:"Qscale,omitempty"`
	// Gets or sets the jpeg quality to use for image tiles.
	JpegQuality *int32 `json:"JpegQuality,omitempty"`
	// Gets or sets the number of threads to be used by ffmpeg.
	ProcessThreads *int32 `json:"ProcessThreads,omitempty"`
}

// NewTrickplayOptions instantiates a new TrickplayOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrickplayOptions() *TrickplayOptions {
	this := TrickplayOptions{}
	return &this
}

// NewTrickplayOptionsWithDefaults instantiates a new TrickplayOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrickplayOptionsWithDefaults() *TrickplayOptions {
	this := TrickplayOptions{}
	return &this
}

// GetEnableHwAcceleration returns the EnableHwAcceleration field value if set, zero value otherwise.
func (o *TrickplayOptions) GetEnableHwAcceleration() bool {
	if o == nil || IsNil(o.EnableHwAcceleration) {
		var ret bool
		return ret
	}
	return *o.EnableHwAcceleration
}

// GetEnableHwAccelerationOk returns a tuple with the EnableHwAcceleration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrickplayOptions) GetEnableHwAccelerationOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableHwAcceleration) {
		return nil, false
	}
	return o.EnableHwAcceleration, true
}

// HasEnableHwAcceleration returns a boolean if a field has been set.
func (o *TrickplayOptions) HasEnableHwAcceleration() bool {
	if o != nil && !IsNil(o.EnableHwAcceleration) {
		return true
	}

	return false
}

// SetEnableHwAcceleration gets a reference to the given bool and assigns it to the EnableHwAcceleration field.
func (o *TrickplayOptions) SetEnableHwAcceleration(v bool) {
	o.EnableHwAcceleration = &v
}

// GetEnableHwEncoding returns the EnableHwEncoding field value if set, zero value otherwise.
func (o *TrickplayOptions) GetEnableHwEncoding() bool {
	if o == nil || IsNil(o.EnableHwEncoding) {
		var ret bool
		return ret
	}
	return *o.EnableHwEncoding
}

// GetEnableHwEncodingOk returns a tuple with the EnableHwEncoding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrickplayOptions) GetEnableHwEncodingOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableHwEncoding) {
		return nil, false
	}
	return o.EnableHwEncoding, true
}

// HasEnableHwEncoding returns a boolean if a field has been set.
func (o *TrickplayOptions) HasEnableHwEncoding() bool {
	if o != nil && !IsNil(o.EnableHwEncoding) {
		return true
	}

	return false
}

// SetEnableHwEncoding gets a reference to the given bool and assigns it to the EnableHwEncoding field.
func (o *TrickplayOptions) SetEnableHwEncoding(v bool) {
	o.EnableHwEncoding = &v
}

// GetEnableKeyFrameOnlyExtraction returns the EnableKeyFrameOnlyExtraction field value if set, zero value otherwise.
func (o *TrickplayOptions) GetEnableKeyFrameOnlyExtraction() bool {
	if o == nil || IsNil(o.EnableKeyFrameOnlyExtraction) {
		var ret bool
		return ret
	}
	return *o.EnableKeyFrameOnlyExtraction
}

// GetEnableKeyFrameOnlyExtractionOk returns a tuple with the EnableKeyFrameOnlyExtraction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrickplayOptions) GetEnableKeyFrameOnlyExtractionOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableKeyFrameOnlyExtraction) {
		return nil, false
	}
	return o.EnableKeyFrameOnlyExtraction, true
}

// HasEnableKeyFrameOnlyExtraction returns a boolean if a field has been set.
func (o *TrickplayOptions) HasEnableKeyFrameOnlyExtraction() bool {
	if o != nil && !IsNil(o.EnableKeyFrameOnlyExtraction) {
		return true
	}

	return false
}

// SetEnableKeyFrameOnlyExtraction gets a reference to the given bool and assigns it to the EnableKeyFrameOnlyExtraction field.
func (o *TrickplayOptions) SetEnableKeyFrameOnlyExtraction(v bool) {
	o.EnableKeyFrameOnlyExtraction = &v
}

// GetScanBehavior returns the ScanBehavior field value if set, zero value otherwise.
func (o *TrickplayOptions) GetScanBehavior() TrickplayScanBehavior {
	if o == nil || IsNil(o.ScanBehavior) {
		var ret TrickplayScanBehavior
		return ret
	}
	return *o.ScanBehavior
}

// GetScanBehaviorOk returns a tuple with the ScanBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrickplayOptions) GetScanBehaviorOk() (*TrickplayScanBehavior, bool) {
	if o == nil || IsNil(o.ScanBehavior) {
		return nil, false
	}
	return o.ScanBehavior, true
}

// HasScanBehavior returns a boolean if a field has been set.
func (o *TrickplayOptions) HasScanBehavior() bool {
	if o != nil && !IsNil(o.ScanBehavior) {
		return true
	}

	return false
}

// SetScanBehavior gets a reference to the given TrickplayScanBehavior and assigns it to the ScanBehavior field.
func (o *TrickplayOptions) SetScanBehavior(v TrickplayScanBehavior) {
	o.ScanBehavior = &v
}

// GetProcessPriority returns the ProcessPriority field value if set, zero value otherwise.
func (o *TrickplayOptions) GetProcessPriority() ProcessPriorityClass {
	if o == nil || IsNil(o.ProcessPriority) {
		var ret ProcessPriorityClass
		return ret
	}
	return *o.ProcessPriority
}

// GetProcessPriorityOk returns a tuple with the ProcessPriority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrickplayOptions) GetProcessPriorityOk() (*ProcessPriorityClass, bool) {
	if o == nil || IsNil(o.ProcessPriority) {
		return nil, false
	}
	return o.ProcessPriority, true
}

// HasProcessPriority returns a boolean if a field has been set.
func (o *TrickplayOptions) HasProcessPriority() bool {
	if o != nil && !IsNil(o.ProcessPriority) {
		return true
	}

	return false
}

// SetProcessPriority gets a reference to the given ProcessPriorityClass and assigns it to the ProcessPriority field.
func (o *TrickplayOptions) SetProcessPriority(v ProcessPriorityClass) {
	o.ProcessPriority = &v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *TrickplayOptions) GetInterval() int32 {
	if o == nil || IsNil(o.Interval) {
		var ret int32
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrickplayOptions) GetIntervalOk() (*int32, bool) {
	if o == nil || IsNil(o.Interval) {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *TrickplayOptions) HasInterval() bool {
	if o != nil && !IsNil(o.Interval) {
		return true
	}

	return false
}

// SetInterval gets a reference to the given int32 and assigns it to the Interval field.
func (o *TrickplayOptions) SetInterval(v int32) {
	o.Interval = &v
}

// GetWidthResolutions returns the WidthResolutions field value if set, zero value otherwise.
func (o *TrickplayOptions) GetWidthResolutions() []int32 {
	if o == nil || IsNil(o.WidthResolutions) {
		var ret []int32
		return ret
	}
	return o.WidthResolutions
}

// GetWidthResolutionsOk returns a tuple with the WidthResolutions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrickplayOptions) GetWidthResolutionsOk() ([]int32, bool) {
	if o == nil || IsNil(o.WidthResolutions) {
		return nil, false
	}
	return o.WidthResolutions, true
}

// HasWidthResolutions returns a boolean if a field has been set.
func (o *TrickplayOptions) HasWidthResolutions() bool {
	if o != nil && !IsNil(o.WidthResolutions) {
		return true
	}

	return false
}

// SetWidthResolutions gets a reference to the given []int32 and assigns it to the WidthResolutions field.
func (o *TrickplayOptions) SetWidthResolutions(v []int32) {
	o.WidthResolutions = v
}

// GetTileWidth returns the TileWidth field value if set, zero value otherwise.
func (o *TrickplayOptions) GetTileWidth() int32 {
	if o == nil || IsNil(o.TileWidth) {
		var ret int32
		return ret
	}
	return *o.TileWidth
}

// GetTileWidthOk returns a tuple with the TileWidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrickplayOptions) GetTileWidthOk() (*int32, bool) {
	if o == nil || IsNil(o.TileWidth) {
		return nil, false
	}
	return o.TileWidth, true
}

// HasTileWidth returns a boolean if a field has been set.
func (o *TrickplayOptions) HasTileWidth() bool {
	if o != nil && !IsNil(o.TileWidth) {
		return true
	}

	return false
}

// SetTileWidth gets a reference to the given int32 and assigns it to the TileWidth field.
func (o *TrickplayOptions) SetTileWidth(v int32) {
	o.TileWidth = &v
}

// GetTileHeight returns the TileHeight field value if set, zero value otherwise.
func (o *TrickplayOptions) GetTileHeight() int32 {
	if o == nil || IsNil(o.TileHeight) {
		var ret int32
		return ret
	}
	return *o.TileHeight
}

// GetTileHeightOk returns a tuple with the TileHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrickplayOptions) GetTileHeightOk() (*int32, bool) {
	if o == nil || IsNil(o.TileHeight) {
		return nil, false
	}
	return o.TileHeight, true
}

// HasTileHeight returns a boolean if a field has been set.
func (o *TrickplayOptions) HasTileHeight() bool {
	if o != nil && !IsNil(o.TileHeight) {
		return true
	}

	return false
}

// SetTileHeight gets a reference to the given int32 and assigns it to the TileHeight field.
func (o *TrickplayOptions) SetTileHeight(v int32) {
	o.TileHeight = &v
}

// GetQscale returns the Qscale field value if set, zero value otherwise.
func (o *TrickplayOptions) GetQscale() int32 {
	if o == nil || IsNil(o.Qscale) {
		var ret int32
		return ret
	}
	return *o.Qscale
}

// GetQscaleOk returns a tuple with the Qscale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrickplayOptions) GetQscaleOk() (*int32, bool) {
	if o == nil || IsNil(o.Qscale) {
		return nil, false
	}
	return o.Qscale, true
}

// HasQscale returns a boolean if a field has been set.
func (o *TrickplayOptions) HasQscale() bool {
	if o != nil && !IsNil(o.Qscale) {
		return true
	}

	return false
}

// SetQscale gets a reference to the given int32 and assigns it to the Qscale field.
func (o *TrickplayOptions) SetQscale(v int32) {
	o.Qscale = &v
}

// GetJpegQuality returns the JpegQuality field value if set, zero value otherwise.
func (o *TrickplayOptions) GetJpegQuality() int32 {
	if o == nil || IsNil(o.JpegQuality) {
		var ret int32
		return ret
	}
	return *o.JpegQuality
}

// GetJpegQualityOk returns a tuple with the JpegQuality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrickplayOptions) GetJpegQualityOk() (*int32, bool) {
	if o == nil || IsNil(o.JpegQuality) {
		return nil, false
	}
	return o.JpegQuality, true
}

// HasJpegQuality returns a boolean if a field has been set.
func (o *TrickplayOptions) HasJpegQuality() bool {
	if o != nil && !IsNil(o.JpegQuality) {
		return true
	}

	return false
}

// SetJpegQuality gets a reference to the given int32 and assigns it to the JpegQuality field.
func (o *TrickplayOptions) SetJpegQuality(v int32) {
	o.JpegQuality = &v
}

// GetProcessThreads returns the ProcessThreads field value if set, zero value otherwise.
func (o *TrickplayOptions) GetProcessThreads() int32 {
	if o == nil || IsNil(o.ProcessThreads) {
		var ret int32
		return ret
	}
	return *o.ProcessThreads
}

// GetProcessThreadsOk returns a tuple with the ProcessThreads field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrickplayOptions) GetProcessThreadsOk() (*int32, bool) {
	if o == nil || IsNil(o.ProcessThreads) {
		return nil, false
	}
	return o.ProcessThreads, true
}

// HasProcessThreads returns a boolean if a field has been set.
func (o *TrickplayOptions) HasProcessThreads() bool {
	if o != nil && !IsNil(o.ProcessThreads) {
		return true
	}

	return false
}

// SetProcessThreads gets a reference to the given int32 and assigns it to the ProcessThreads field.
func (o *TrickplayOptions) SetProcessThreads(v int32) {
	o.ProcessThreads = &v
}

func (o TrickplayOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TrickplayOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EnableHwAcceleration) {
		toSerialize["EnableHwAcceleration"] = o.EnableHwAcceleration
	}
	if !IsNil(o.EnableHwEncoding) {
		toSerialize["EnableHwEncoding"] = o.EnableHwEncoding
	}
	if !IsNil(o.EnableKeyFrameOnlyExtraction) {
		toSerialize["EnableKeyFrameOnlyExtraction"] = o.EnableKeyFrameOnlyExtraction
	}
	if !IsNil(o.ScanBehavior) {
		toSerialize["ScanBehavior"] = o.ScanBehavior
	}
	if !IsNil(o.ProcessPriority) {
		toSerialize["ProcessPriority"] = o.ProcessPriority
	}
	if !IsNil(o.Interval) {
		toSerialize["Interval"] = o.Interval
	}
	if !IsNil(o.WidthResolutions) {
		toSerialize["WidthResolutions"] = o.WidthResolutions
	}
	if !IsNil(o.TileWidth) {
		toSerialize["TileWidth"] = o.TileWidth
	}
	if !IsNil(o.TileHeight) {
		toSerialize["TileHeight"] = o.TileHeight
	}
	if !IsNil(o.Qscale) {
		toSerialize["Qscale"] = o.Qscale
	}
	if !IsNil(o.JpegQuality) {
		toSerialize["JpegQuality"] = o.JpegQuality
	}
	if !IsNil(o.ProcessThreads) {
		toSerialize["ProcessThreads"] = o.ProcessThreads
	}
	return toSerialize, nil
}

type NullableTrickplayOptions struct {
	value *TrickplayOptions
	isSet bool
}

func (v NullableTrickplayOptions) Get() *TrickplayOptions {
	return v.value
}

func (v *NullableTrickplayOptions) Set(val *TrickplayOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableTrickplayOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableTrickplayOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrickplayOptions(val *TrickplayOptions) *NullableTrickplayOptions {
	return &NullableTrickplayOptions{value: val, isSet: true}
}

func (v NullableTrickplayOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrickplayOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

