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

// DayOfWeek the model 'DayOfWeek'
type DayOfWeek string

// List of DayOfWeek
const (
	DAYOFWEEK_SUNDAY DayOfWeek = "Sunday"
	DAYOFWEEK_MONDAY DayOfWeek = "Monday"
	DAYOFWEEK_TUESDAY DayOfWeek = "Tuesday"
	DAYOFWEEK_WEDNESDAY DayOfWeek = "Wednesday"
	DAYOFWEEK_THURSDAY DayOfWeek = "Thursday"
	DAYOFWEEK_FRIDAY DayOfWeek = "Friday"
	DAYOFWEEK_SATURDAY DayOfWeek = "Saturday"
)

// All allowed values of DayOfWeek enum
var AllowedDayOfWeekEnumValues = []DayOfWeek{
	"Sunday",
	"Monday",
	"Tuesday",
	"Wednesday",
	"Thursday",
	"Friday",
	"Saturday",
}

func (v *DayOfWeek) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DayOfWeek(value)
	for _, existing := range AllowedDayOfWeekEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DayOfWeek", value)
}

// NewDayOfWeekFromValue returns a pointer to a valid DayOfWeek
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDayOfWeekFromValue(v string) (*DayOfWeek, error) {
	ev := DayOfWeek(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DayOfWeek: valid values are %v", v, AllowedDayOfWeekEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DayOfWeek) IsValid() bool {
	for _, existing := range AllowedDayOfWeekEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DayOfWeek value
func (v DayOfWeek) Ptr() *DayOfWeek {
	return &v
}

type NullableDayOfWeek struct {
	value *DayOfWeek
	isSet bool
}

func (v NullableDayOfWeek) Get() *DayOfWeek {
	return v.value
}

func (v *NullableDayOfWeek) Set(val *DayOfWeek) {
	v.value = val
	v.isSet = true
}

func (v NullableDayOfWeek) IsSet() bool {
	return v.isSet
}

func (v *NullableDayOfWeek) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDayOfWeek(val *DayOfWeek) *NullableDayOfWeek {
	return &NullableDayOfWeek{value: val, isSet: true}
}

func (v NullableDayOfWeek) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDayOfWeek) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
