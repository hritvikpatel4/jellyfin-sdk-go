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

// PluginStatus Plugin load status.
type PluginStatus string

// List of PluginStatus
const (
	PLUGINSTATUS_ACTIVE PluginStatus = "Active"
	PLUGINSTATUS_RESTART PluginStatus = "Restart"
	PLUGINSTATUS_DELETED PluginStatus = "Deleted"
	PLUGINSTATUS_SUPERCEDED PluginStatus = "Superceded"
	PLUGINSTATUS_MALFUNCTIONED PluginStatus = "Malfunctioned"
	PLUGINSTATUS_NOT_SUPPORTED PluginStatus = "NotSupported"
	PLUGINSTATUS_DISABLED PluginStatus = "Disabled"
)

// All allowed values of PluginStatus enum
var AllowedPluginStatusEnumValues = []PluginStatus{
	"Active",
	"Restart",
	"Deleted",
	"Superceded",
	"Malfunctioned",
	"NotSupported",
	"Disabled",
}

func (v *PluginStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PluginStatus(value)
	for _, existing := range AllowedPluginStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PluginStatus", value)
}

// NewPluginStatusFromValue returns a pointer to a valid PluginStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPluginStatusFromValue(v string) (*PluginStatus, error) {
	ev := PluginStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PluginStatus: valid values are %v", v, AllowedPluginStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PluginStatus) IsValid() bool {
	for _, existing := range AllowedPluginStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PluginStatus value
func (v PluginStatus) Ptr() *PluginStatus {
	return &v
}

type NullablePluginStatus struct {
	value *PluginStatus
	isSet bool
}

func (v NullablePluginStatus) Get() *PluginStatus {
	return v.value
}

func (v *NullablePluginStatus) Set(val *PluginStatus) {
	v.value = val
	v.isSet = true
}

func (v NullablePluginStatus) IsSet() bool {
	return v.isSet
}

func (v *NullablePluginStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePluginStatus(val *PluginStatus) *NullablePluginStatus {
	return &NullablePluginStatus{value: val, isSet: true}
}

func (v NullablePluginStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePluginStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
