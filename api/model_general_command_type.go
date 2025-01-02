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

// GeneralCommandType This exists simply to identify a set of known commands.
type GeneralCommandType string

// List of GeneralCommandType
const (
	GENERALCOMMANDTYPE_MOVE_UP GeneralCommandType = "MoveUp"
	GENERALCOMMANDTYPE_MOVE_DOWN GeneralCommandType = "MoveDown"
	GENERALCOMMANDTYPE_MOVE_LEFT GeneralCommandType = "MoveLeft"
	GENERALCOMMANDTYPE_MOVE_RIGHT GeneralCommandType = "MoveRight"
	GENERALCOMMANDTYPE_PAGE_UP GeneralCommandType = "PageUp"
	GENERALCOMMANDTYPE_PAGE_DOWN GeneralCommandType = "PageDown"
	GENERALCOMMANDTYPE_PREVIOUS_LETTER GeneralCommandType = "PreviousLetter"
	GENERALCOMMANDTYPE_NEXT_LETTER GeneralCommandType = "NextLetter"
	GENERALCOMMANDTYPE_TOGGLE_OSD GeneralCommandType = "ToggleOsd"
	GENERALCOMMANDTYPE_TOGGLE_CONTEXT_MENU GeneralCommandType = "ToggleContextMenu"
	GENERALCOMMANDTYPE_SELECT GeneralCommandType = "Select"
	GENERALCOMMANDTYPE_BACK GeneralCommandType = "Back"
	GENERALCOMMANDTYPE_TAKE_SCREENSHOT GeneralCommandType = "TakeScreenshot"
	GENERALCOMMANDTYPE_SEND_KEY GeneralCommandType = "SendKey"
	GENERALCOMMANDTYPE_SEND_STRING GeneralCommandType = "SendString"
	GENERALCOMMANDTYPE_GO_HOME GeneralCommandType = "GoHome"
	GENERALCOMMANDTYPE_GO_TO_SETTINGS GeneralCommandType = "GoToSettings"
	GENERALCOMMANDTYPE_VOLUME_UP GeneralCommandType = "VolumeUp"
	GENERALCOMMANDTYPE_VOLUME_DOWN GeneralCommandType = "VolumeDown"
	GENERALCOMMANDTYPE_MUTE GeneralCommandType = "Mute"
	GENERALCOMMANDTYPE_UNMUTE GeneralCommandType = "Unmute"
	GENERALCOMMANDTYPE_TOGGLE_MUTE GeneralCommandType = "ToggleMute"
	GENERALCOMMANDTYPE_SET_VOLUME GeneralCommandType = "SetVolume"
	GENERALCOMMANDTYPE_SET_AUDIO_STREAM_INDEX GeneralCommandType = "SetAudioStreamIndex"
	GENERALCOMMANDTYPE_SET_SUBTITLE_STREAM_INDEX GeneralCommandType = "SetSubtitleStreamIndex"
	GENERALCOMMANDTYPE_TOGGLE_FULLSCREEN GeneralCommandType = "ToggleFullscreen"
	GENERALCOMMANDTYPE_DISPLAY_CONTENT GeneralCommandType = "DisplayContent"
	GENERALCOMMANDTYPE_GO_TO_SEARCH GeneralCommandType = "GoToSearch"
	GENERALCOMMANDTYPE_DISPLAY_MESSAGE GeneralCommandType = "DisplayMessage"
	GENERALCOMMANDTYPE_SET_REPEAT_MODE GeneralCommandType = "SetRepeatMode"
	GENERALCOMMANDTYPE_CHANNEL_UP GeneralCommandType = "ChannelUp"
	GENERALCOMMANDTYPE_CHANNEL_DOWN GeneralCommandType = "ChannelDown"
	GENERALCOMMANDTYPE_GUIDE GeneralCommandType = "Guide"
	GENERALCOMMANDTYPE_TOGGLE_STATS GeneralCommandType = "ToggleStats"
	GENERALCOMMANDTYPE_PLAY_MEDIA_SOURCE GeneralCommandType = "PlayMediaSource"
	GENERALCOMMANDTYPE_PLAY_TRAILERS GeneralCommandType = "PlayTrailers"
	GENERALCOMMANDTYPE_SET_SHUFFLE_QUEUE GeneralCommandType = "SetShuffleQueue"
	GENERALCOMMANDTYPE_PLAY_STATE GeneralCommandType = "PlayState"
	GENERALCOMMANDTYPE_PLAY_NEXT GeneralCommandType = "PlayNext"
	GENERALCOMMANDTYPE_TOGGLE_OSD_MENU GeneralCommandType = "ToggleOsdMenu"
	GENERALCOMMANDTYPE_PLAY GeneralCommandType = "Play"
	GENERALCOMMANDTYPE_SET_MAX_STREAMING_BITRATE GeneralCommandType = "SetMaxStreamingBitrate"
	GENERALCOMMANDTYPE_SET_PLAYBACK_ORDER GeneralCommandType = "SetPlaybackOrder"
)

// All allowed values of GeneralCommandType enum
var AllowedGeneralCommandTypeEnumValues = []GeneralCommandType{
	"MoveUp",
	"MoveDown",
	"MoveLeft",
	"MoveRight",
	"PageUp",
	"PageDown",
	"PreviousLetter",
	"NextLetter",
	"ToggleOsd",
	"ToggleContextMenu",
	"Select",
	"Back",
	"TakeScreenshot",
	"SendKey",
	"SendString",
	"GoHome",
	"GoToSettings",
	"VolumeUp",
	"VolumeDown",
	"Mute",
	"Unmute",
	"ToggleMute",
	"SetVolume",
	"SetAudioStreamIndex",
	"SetSubtitleStreamIndex",
	"ToggleFullscreen",
	"DisplayContent",
	"GoToSearch",
	"DisplayMessage",
	"SetRepeatMode",
	"ChannelUp",
	"ChannelDown",
	"Guide",
	"ToggleStats",
	"PlayMediaSource",
	"PlayTrailers",
	"SetShuffleQueue",
	"PlayState",
	"PlayNext",
	"ToggleOsdMenu",
	"Play",
	"SetMaxStreamingBitrate",
	"SetPlaybackOrder",
}

func (v *GeneralCommandType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GeneralCommandType(value)
	for _, existing := range AllowedGeneralCommandTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GeneralCommandType", value)
}

// NewGeneralCommandTypeFromValue returns a pointer to a valid GeneralCommandType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGeneralCommandTypeFromValue(v string) (*GeneralCommandType, error) {
	ev := GeneralCommandType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GeneralCommandType: valid values are %v", v, AllowedGeneralCommandTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GeneralCommandType) IsValid() bool {
	for _, existing := range AllowedGeneralCommandTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GeneralCommandType value
func (v GeneralCommandType) Ptr() *GeneralCommandType {
	return &v
}

type NullableGeneralCommandType struct {
	value *GeneralCommandType
	isSet bool
}

func (v NullableGeneralCommandType) Get() *GeneralCommandType {
	return v.value
}

func (v *NullableGeneralCommandType) Set(val *GeneralCommandType) {
	v.value = val
	v.isSet = true
}

func (v NullableGeneralCommandType) IsSet() bool {
	return v.isSet
}

func (v *NullableGeneralCommandType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGeneralCommandType(val *GeneralCommandType) *NullableGeneralCommandType {
	return &NullableGeneralCommandType{value: val, isSet: true}
}

func (v NullableGeneralCommandType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGeneralCommandType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
