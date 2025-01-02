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

// SessionMessageType The different kinds of messages that are used in the WebSocket api.
type SessionMessageType string

// List of SessionMessageType
const (
	SESSIONMESSAGETYPE_FORCE_KEEP_ALIVE SessionMessageType = "ForceKeepAlive"
	SESSIONMESSAGETYPE_GENERAL_COMMAND SessionMessageType = "GeneralCommand"
	SESSIONMESSAGETYPE_USER_DATA_CHANGED SessionMessageType = "UserDataChanged"
	SESSIONMESSAGETYPE_SESSIONS SessionMessageType = "Sessions"
	SESSIONMESSAGETYPE_PLAY SessionMessageType = "Play"
	SESSIONMESSAGETYPE_SYNC_PLAY_COMMAND SessionMessageType = "SyncPlayCommand"
	SESSIONMESSAGETYPE_SYNC_PLAY_GROUP_UPDATE SessionMessageType = "SyncPlayGroupUpdate"
	SESSIONMESSAGETYPE_PLAYSTATE SessionMessageType = "Playstate"
	SESSIONMESSAGETYPE_RESTART_REQUIRED SessionMessageType = "RestartRequired"
	SESSIONMESSAGETYPE_SERVER_SHUTTING_DOWN SessionMessageType = "ServerShuttingDown"
	SESSIONMESSAGETYPE_SERVER_RESTARTING SessionMessageType = "ServerRestarting"
	SESSIONMESSAGETYPE_LIBRARY_CHANGED SessionMessageType = "LibraryChanged"
	SESSIONMESSAGETYPE_USER_DELETED SessionMessageType = "UserDeleted"
	SESSIONMESSAGETYPE_USER_UPDATED SessionMessageType = "UserUpdated"
	SESSIONMESSAGETYPE_SERIES_TIMER_CREATED SessionMessageType = "SeriesTimerCreated"
	SESSIONMESSAGETYPE_TIMER_CREATED SessionMessageType = "TimerCreated"
	SESSIONMESSAGETYPE_SERIES_TIMER_CANCELLED SessionMessageType = "SeriesTimerCancelled"
	SESSIONMESSAGETYPE_TIMER_CANCELLED SessionMessageType = "TimerCancelled"
	SESSIONMESSAGETYPE_REFRESH_PROGRESS SessionMessageType = "RefreshProgress"
	SESSIONMESSAGETYPE_SCHEDULED_TASK_ENDED SessionMessageType = "ScheduledTaskEnded"
	SESSIONMESSAGETYPE_PACKAGE_INSTALLATION_CANCELLED SessionMessageType = "PackageInstallationCancelled"
	SESSIONMESSAGETYPE_PACKAGE_INSTALLATION_FAILED SessionMessageType = "PackageInstallationFailed"
	SESSIONMESSAGETYPE_PACKAGE_INSTALLATION_COMPLETED SessionMessageType = "PackageInstallationCompleted"
	SESSIONMESSAGETYPE_PACKAGE_INSTALLING SessionMessageType = "PackageInstalling"
	SESSIONMESSAGETYPE_PACKAGE_UNINSTALLED SessionMessageType = "PackageUninstalled"
	SESSIONMESSAGETYPE_ACTIVITY_LOG_ENTRY SessionMessageType = "ActivityLogEntry"
	SESSIONMESSAGETYPE_SCHEDULED_TASKS_INFO SessionMessageType = "ScheduledTasksInfo"
	SESSIONMESSAGETYPE_ACTIVITY_LOG_ENTRY_START SessionMessageType = "ActivityLogEntryStart"
	SESSIONMESSAGETYPE_ACTIVITY_LOG_ENTRY_STOP SessionMessageType = "ActivityLogEntryStop"
	SESSIONMESSAGETYPE_SESSIONS_START SessionMessageType = "SessionsStart"
	SESSIONMESSAGETYPE_SESSIONS_STOP SessionMessageType = "SessionsStop"
	SESSIONMESSAGETYPE_SCHEDULED_TASKS_INFO_START SessionMessageType = "ScheduledTasksInfoStart"
	SESSIONMESSAGETYPE_SCHEDULED_TASKS_INFO_STOP SessionMessageType = "ScheduledTasksInfoStop"
	SESSIONMESSAGETYPE_KEEP_ALIVE SessionMessageType = "KeepAlive"
)

// All allowed values of SessionMessageType enum
var AllowedSessionMessageTypeEnumValues = []SessionMessageType{
	"ForceKeepAlive",
	"GeneralCommand",
	"UserDataChanged",
	"Sessions",
	"Play",
	"SyncPlayCommand",
	"SyncPlayGroupUpdate",
	"Playstate",
	"RestartRequired",
	"ServerShuttingDown",
	"ServerRestarting",
	"LibraryChanged",
	"UserDeleted",
	"UserUpdated",
	"SeriesTimerCreated",
	"TimerCreated",
	"SeriesTimerCancelled",
	"TimerCancelled",
	"RefreshProgress",
	"ScheduledTaskEnded",
	"PackageInstallationCancelled",
	"PackageInstallationFailed",
	"PackageInstallationCompleted",
	"PackageInstalling",
	"PackageUninstalled",
	"ActivityLogEntry",
	"ScheduledTasksInfo",
	"ActivityLogEntryStart",
	"ActivityLogEntryStop",
	"SessionsStart",
	"SessionsStop",
	"ScheduledTasksInfoStart",
	"ScheduledTasksInfoStop",
	"KeepAlive",
}

func (v *SessionMessageType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SessionMessageType(value)
	for _, existing := range AllowedSessionMessageTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SessionMessageType", value)
}

// NewSessionMessageTypeFromValue returns a pointer to a valid SessionMessageType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSessionMessageTypeFromValue(v string) (*SessionMessageType, error) {
	ev := SessionMessageType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SessionMessageType: valid values are %v", v, AllowedSessionMessageTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SessionMessageType) IsValid() bool {
	for _, existing := range AllowedSessionMessageTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SessionMessageType value
func (v SessionMessageType) Ptr() *SessionMessageType {
	return &v
}

type NullableSessionMessageType struct {
	value *SessionMessageType
	isSet bool
}

func (v NullableSessionMessageType) Get() *SessionMessageType {
	return v.value
}

func (v *NullableSessionMessageType) Set(val *SessionMessageType) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionMessageType) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionMessageType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionMessageType(val *SessionMessageType) *NullableSessionMessageType {
	return &NullableSessionMessageType{value: val, isSet: true}
}

func (v NullableSessionMessageType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionMessageType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
