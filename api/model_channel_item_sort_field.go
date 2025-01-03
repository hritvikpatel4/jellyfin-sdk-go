/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.10.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// ChannelItemSortField the model 'ChannelItemSortField'
type ChannelItemSortField string

// List of ChannelItemSortField
const (
	CHANNELITEMSORTFIELD_NAME                 ChannelItemSortField = "Name"
	CHANNELITEMSORTFIELD_COMMUNITY_RATING     ChannelItemSortField = "CommunityRating"
	CHANNELITEMSORTFIELD_PREMIERE_DATE        ChannelItemSortField = "PremiereDate"
	CHANNELITEMSORTFIELD_DATE_CREATED         ChannelItemSortField = "DateCreated"
	CHANNELITEMSORTFIELD_RUNTIME              ChannelItemSortField = "Runtime"
	CHANNELITEMSORTFIELD_PLAY_COUNT           ChannelItemSortField = "PlayCount"
	CHANNELITEMSORTFIELD_COMMUNITY_PLAY_COUNT ChannelItemSortField = "CommunityPlayCount"
)

// All allowed values of ChannelItemSortField enum
var AllowedChannelItemSortFieldEnumValues = []ChannelItemSortField{
	"Name",
	"CommunityRating",
	"PremiereDate",
	"DateCreated",
	"Runtime",
	"PlayCount",
	"CommunityPlayCount",
}

func (v *ChannelItemSortField) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ChannelItemSortField(value)
	for _, existing := range AllowedChannelItemSortFieldEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ChannelItemSortField", value)
}

// NewChannelItemSortFieldFromValue returns a pointer to a valid ChannelItemSortField
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewChannelItemSortFieldFromValue(v string) (*ChannelItemSortField, error) {
	ev := ChannelItemSortField(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ChannelItemSortField: valid values are %v", v, AllowedChannelItemSortFieldEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ChannelItemSortField) IsValid() bool {
	for _, existing := range AllowedChannelItemSortFieldEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ChannelItemSortField value
func (v ChannelItemSortField) Ptr() *ChannelItemSortField {
	return &v
}

type NullableChannelItemSortField struct {
	value *ChannelItemSortField
	isSet bool
}

func (v NullableChannelItemSortField) Get() *ChannelItemSortField {
	return v.value
}

func (v *NullableChannelItemSortField) Set(val *ChannelItemSortField) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelItemSortField) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelItemSortField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelItemSortField(val *ChannelItemSortField) *NullableChannelItemSortField {
	return &NullableChannelItemSortField{value: val, isSet: true}
}

func (v NullableChannelItemSortField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelItemSortField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
