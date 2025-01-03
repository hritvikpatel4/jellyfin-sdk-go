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

// PlayCommand Enum PlayCommand.
type PlayCommand string

// List of PlayCommand
const (
	PLAYCOMMAND_PLAY_NOW         PlayCommand = "PlayNow"
	PLAYCOMMAND_PLAY_NEXT        PlayCommand = "PlayNext"
	PLAYCOMMAND_PLAY_LAST        PlayCommand = "PlayLast"
	PLAYCOMMAND_PLAY_INSTANT_MIX PlayCommand = "PlayInstantMix"
	PLAYCOMMAND_PLAY_SHUFFLE     PlayCommand = "PlayShuffle"
)

// All allowed values of PlayCommand enum
var AllowedPlayCommandEnumValues = []PlayCommand{
	"PlayNow",
	"PlayNext",
	"PlayLast",
	"PlayInstantMix",
	"PlayShuffle",
}

func (v *PlayCommand) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PlayCommand(value)
	for _, existing := range AllowedPlayCommandEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PlayCommand", value)
}

// NewPlayCommandFromValue returns a pointer to a valid PlayCommand
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPlayCommandFromValue(v string) (*PlayCommand, error) {
	ev := PlayCommand(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PlayCommand: valid values are %v", v, AllowedPlayCommandEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PlayCommand) IsValid() bool {
	for _, existing := range AllowedPlayCommandEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PlayCommand value
func (v PlayCommand) Ptr() *PlayCommand {
	return &v
}

type NullablePlayCommand struct {
	value *PlayCommand
	isSet bool
}

func (v NullablePlayCommand) Get() *PlayCommand {
	return v.value
}

func (v *NullablePlayCommand) Set(val *PlayCommand) {
	v.value = val
	v.isSet = true
}

func (v NullablePlayCommand) IsSet() bool {
	return v.isSet
}

func (v *NullablePlayCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlayCommand(val *PlayCommand) *NullablePlayCommand {
	return &NullablePlayCommand{value: val, isSet: true}
}

func (v NullablePlayCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlayCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
