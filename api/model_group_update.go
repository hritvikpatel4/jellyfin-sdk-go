/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.10.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"gopkg.in/validator.v2"
	"fmt"
)

// GroupUpdate - Group update without data.
type GroupUpdate struct {
	GroupInfoDtoGroupUpdate *GroupInfoDtoGroupUpdate
	GroupStateUpdateGroupUpdate *GroupStateUpdateGroupUpdate
	PlayQueueUpdateGroupUpdate *PlayQueueUpdateGroupUpdate
	StringGroupUpdate *StringGroupUpdate
}

// GroupInfoDtoGroupUpdateAsGroupUpdate is a convenience function that returns GroupInfoDtoGroupUpdate wrapped in GroupUpdate
func GroupInfoDtoGroupUpdateAsGroupUpdate(v *GroupInfoDtoGroupUpdate) GroupUpdate {
	return GroupUpdate{
		GroupInfoDtoGroupUpdate: v,
	}
}

// GroupStateUpdateGroupUpdateAsGroupUpdate is a convenience function that returns GroupStateUpdateGroupUpdate wrapped in GroupUpdate
func GroupStateUpdateGroupUpdateAsGroupUpdate(v *GroupStateUpdateGroupUpdate) GroupUpdate {
	return GroupUpdate{
		GroupStateUpdateGroupUpdate: v,
	}
}

// PlayQueueUpdateGroupUpdateAsGroupUpdate is a convenience function that returns PlayQueueUpdateGroupUpdate wrapped in GroupUpdate
func PlayQueueUpdateGroupUpdateAsGroupUpdate(v *PlayQueueUpdateGroupUpdate) GroupUpdate {
	return GroupUpdate{
		PlayQueueUpdateGroupUpdate: v,
	}
}

// StringGroupUpdateAsGroupUpdate is a convenience function that returns StringGroupUpdate wrapped in GroupUpdate
func StringGroupUpdateAsGroupUpdate(v *StringGroupUpdate) GroupUpdate {
	return GroupUpdate{
		StringGroupUpdate: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *GroupUpdate) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into GroupInfoDtoGroupUpdate
	err = newStrictDecoder(data).Decode(&dst.GroupInfoDtoGroupUpdate)
	if err == nil {
		jsonGroupInfoDtoGroupUpdate, _ := json.Marshal(dst.GroupInfoDtoGroupUpdate)
		if string(jsonGroupInfoDtoGroupUpdate) == "{}" { // empty struct
			dst.GroupInfoDtoGroupUpdate = nil
		} else {
			if err = validator.Validate(dst.GroupInfoDtoGroupUpdate); err != nil {
				dst.GroupInfoDtoGroupUpdate = nil
			} else {
				match++
			}
		}
	} else {
		dst.GroupInfoDtoGroupUpdate = nil
	}

	// try to unmarshal data into GroupStateUpdateGroupUpdate
	err = newStrictDecoder(data).Decode(&dst.GroupStateUpdateGroupUpdate)
	if err == nil {
		jsonGroupStateUpdateGroupUpdate, _ := json.Marshal(dst.GroupStateUpdateGroupUpdate)
		if string(jsonGroupStateUpdateGroupUpdate) == "{}" { // empty struct
			dst.GroupStateUpdateGroupUpdate = nil
		} else {
			if err = validator.Validate(dst.GroupStateUpdateGroupUpdate); err != nil {
				dst.GroupStateUpdateGroupUpdate = nil
			} else {
				match++
			}
		}
	} else {
		dst.GroupStateUpdateGroupUpdate = nil
	}

	// try to unmarshal data into PlayQueueUpdateGroupUpdate
	err = newStrictDecoder(data).Decode(&dst.PlayQueueUpdateGroupUpdate)
	if err == nil {
		jsonPlayQueueUpdateGroupUpdate, _ := json.Marshal(dst.PlayQueueUpdateGroupUpdate)
		if string(jsonPlayQueueUpdateGroupUpdate) == "{}" { // empty struct
			dst.PlayQueueUpdateGroupUpdate = nil
		} else {
			if err = validator.Validate(dst.PlayQueueUpdateGroupUpdate); err != nil {
				dst.PlayQueueUpdateGroupUpdate = nil
			} else {
				match++
			}
		}
	} else {
		dst.PlayQueueUpdateGroupUpdate = nil
	}

	// try to unmarshal data into StringGroupUpdate
	err = newStrictDecoder(data).Decode(&dst.StringGroupUpdate)
	if err == nil {
		jsonStringGroupUpdate, _ := json.Marshal(dst.StringGroupUpdate)
		if string(jsonStringGroupUpdate) == "{}" { // empty struct
			dst.StringGroupUpdate = nil
		} else {
			if err = validator.Validate(dst.StringGroupUpdate); err != nil {
				dst.StringGroupUpdate = nil
			} else {
				match++
			}
		}
	} else {
		dst.StringGroupUpdate = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.GroupInfoDtoGroupUpdate = nil
		dst.GroupStateUpdateGroupUpdate = nil
		dst.PlayQueueUpdateGroupUpdate = nil
		dst.StringGroupUpdate = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GroupUpdate)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GroupUpdate)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GroupUpdate) MarshalJSON() ([]byte, error) {
	if src.GroupInfoDtoGroupUpdate != nil {
		return json.Marshal(&src.GroupInfoDtoGroupUpdate)
	}

	if src.GroupStateUpdateGroupUpdate != nil {
		return json.Marshal(&src.GroupStateUpdateGroupUpdate)
	}

	if src.PlayQueueUpdateGroupUpdate != nil {
		return json.Marshal(&src.PlayQueueUpdateGroupUpdate)
	}

	if src.StringGroupUpdate != nil {
		return json.Marshal(&src.StringGroupUpdate)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GroupUpdate) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.GroupInfoDtoGroupUpdate != nil {
		return obj.GroupInfoDtoGroupUpdate
	}

	if obj.GroupStateUpdateGroupUpdate != nil {
		return obj.GroupStateUpdateGroupUpdate
	}

	if obj.PlayQueueUpdateGroupUpdate != nil {
		return obj.PlayQueueUpdateGroupUpdate
	}

	if obj.StringGroupUpdate != nil {
		return obj.StringGroupUpdate
	}

	// all schemas are nil
	return nil
}

type NullableGroupUpdate struct {
	value *GroupUpdate
	isSet bool
}

func (v NullableGroupUpdate) Get() *GroupUpdate {
	return v.value
}

func (v *NullableGroupUpdate) Set(val *GroupUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupUpdate(val *GroupUpdate) *NullableGroupUpdate {
	return &NullableGroupUpdate{value: val, isSet: true}
}

func (v NullableGroupUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

