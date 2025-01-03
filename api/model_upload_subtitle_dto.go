/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.10.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the UploadSubtitleDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UploadSubtitleDto{}

// UploadSubtitleDto Upload subtitles dto.
type UploadSubtitleDto struct {
	// Gets or sets the subtitle language.
	Language string `json:"Language"`
	// Gets or sets the subtitle format.
	Format string `json:"Format"`
	// Gets or sets a value indicating whether the subtitle is forced.
	IsForced bool `json:"IsForced"`
	// Gets or sets a value indicating whether the subtitle is for hearing impaired.
	IsHearingImpaired bool `json:"IsHearingImpaired"`
	// Gets or sets the subtitle data.
	Data string `json:"Data"`
}

type _UploadSubtitleDto UploadSubtitleDto

// NewUploadSubtitleDto instantiates a new UploadSubtitleDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUploadSubtitleDto(language string, format string, isForced bool, isHearingImpaired bool, data string) *UploadSubtitleDto {
	this := UploadSubtitleDto{}
	this.Language = language
	this.Format = format
	this.IsForced = isForced
	this.IsHearingImpaired = isHearingImpaired
	this.Data = data
	return &this
}

// NewUploadSubtitleDtoWithDefaults instantiates a new UploadSubtitleDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUploadSubtitleDtoWithDefaults() *UploadSubtitleDto {
	this := UploadSubtitleDto{}
	return &this
}

// GetLanguage returns the Language field value
func (o *UploadSubtitleDto) GetLanguage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Language
}

// GetLanguageOk returns a tuple with the Language field value
// and a boolean to check if the value has been set.
func (o *UploadSubtitleDto) GetLanguageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Language, true
}

// SetLanguage sets field value
func (o *UploadSubtitleDto) SetLanguage(v string) {
	o.Language = v
}

// GetFormat returns the Format field value
func (o *UploadSubtitleDto) GetFormat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Format
}

// GetFormatOk returns a tuple with the Format field value
// and a boolean to check if the value has been set.
func (o *UploadSubtitleDto) GetFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Format, true
}

// SetFormat sets field value
func (o *UploadSubtitleDto) SetFormat(v string) {
	o.Format = v
}

// GetIsForced returns the IsForced field value
func (o *UploadSubtitleDto) GetIsForced() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsForced
}

// GetIsForcedOk returns a tuple with the IsForced field value
// and a boolean to check if the value has been set.
func (o *UploadSubtitleDto) GetIsForcedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsForced, true
}

// SetIsForced sets field value
func (o *UploadSubtitleDto) SetIsForced(v bool) {
	o.IsForced = v
}

// GetIsHearingImpaired returns the IsHearingImpaired field value
func (o *UploadSubtitleDto) GetIsHearingImpaired() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsHearingImpaired
}

// GetIsHearingImpairedOk returns a tuple with the IsHearingImpaired field value
// and a boolean to check if the value has been set.
func (o *UploadSubtitleDto) GetIsHearingImpairedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsHearingImpaired, true
}

// SetIsHearingImpaired sets field value
func (o *UploadSubtitleDto) SetIsHearingImpaired(v bool) {
	o.IsHearingImpaired = v
}

// GetData returns the Data field value
func (o *UploadSubtitleDto) GetData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *UploadSubtitleDto) GetDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *UploadSubtitleDto) SetData(v string) {
	o.Data = v
}

func (o UploadSubtitleDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UploadSubtitleDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Language"] = o.Language
	toSerialize["Format"] = o.Format
	toSerialize["IsForced"] = o.IsForced
	toSerialize["IsHearingImpaired"] = o.IsHearingImpaired
	toSerialize["Data"] = o.Data
	return toSerialize, nil
}

func (o *UploadSubtitleDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"Language",
		"Format",
		"IsForced",
		"IsHearingImpaired",
		"Data",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varUploadSubtitleDto := _UploadSubtitleDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUploadSubtitleDto)

	if err != nil {
		return err
	}

	*o = UploadSubtitleDto(varUploadSubtitleDto)

	return err
}

type NullableUploadSubtitleDto struct {
	value *UploadSubtitleDto
	isSet bool
}

func (v NullableUploadSubtitleDto) Get() *UploadSubtitleDto {
	return v.value
}

func (v *NullableUploadSubtitleDto) Set(val *UploadSubtitleDto) {
	v.value = val
	v.isSet = true
}

func (v NullableUploadSubtitleDto) IsSet() bool {
	return v.isSet
}

func (v *NullableUploadSubtitleDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUploadSubtitleDto(val *UploadSubtitleDto) *NullableUploadSubtitleDto {
	return &NullableUploadSubtitleDto{value: val, isSet: true}
}

func (v NullableUploadSubtitleDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUploadSubtitleDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
