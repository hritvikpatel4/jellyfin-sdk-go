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

// checks if the WakeOnLanInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WakeOnLanInfo{}

// WakeOnLanInfo Provides the MAC address and port for wake-on-LAN functionality.
type WakeOnLanInfo struct {
	// Gets the MAC address of the device.
	MacAddress NullableString `json:"MacAddress,omitempty"`
	// Gets or sets the wake-on-LAN port.
	Port *int32 `json:"Port,omitempty"`
}

// NewWakeOnLanInfo instantiates a new WakeOnLanInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWakeOnLanInfo() *WakeOnLanInfo {
	this := WakeOnLanInfo{}
	return &this
}

// NewWakeOnLanInfoWithDefaults instantiates a new WakeOnLanInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWakeOnLanInfoWithDefaults() *WakeOnLanInfo {
	this := WakeOnLanInfo{}
	return &this
}

// GetMacAddress returns the MacAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WakeOnLanInfo) GetMacAddress() string {
	if o == nil || IsNil(o.MacAddress.Get()) {
		var ret string
		return ret
	}
	return *o.MacAddress.Get()
}

// GetMacAddressOk returns a tuple with the MacAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WakeOnLanInfo) GetMacAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MacAddress.Get(), o.MacAddress.IsSet()
}

// HasMacAddress returns a boolean if a field has been set.
func (o *WakeOnLanInfo) HasMacAddress() bool {
	if o != nil && o.MacAddress.IsSet() {
		return true
	}

	return false
}

// SetMacAddress gets a reference to the given NullableString and assigns it to the MacAddress field.
func (o *WakeOnLanInfo) SetMacAddress(v string) {
	o.MacAddress.Set(&v)
}

// SetMacAddressNil sets the value for MacAddress to be an explicit nil
func (o *WakeOnLanInfo) SetMacAddressNil() {
	o.MacAddress.Set(nil)
}

// UnsetMacAddress ensures that no value is present for MacAddress, not even an explicit nil
func (o *WakeOnLanInfo) UnsetMacAddress() {
	o.MacAddress.Unset()
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *WakeOnLanInfo) GetPort() int32 {
	if o == nil || IsNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WakeOnLanInfo) GetPortOk() (*int32, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *WakeOnLanInfo) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *WakeOnLanInfo) SetPort(v int32) {
	o.Port = &v
}

func (o WakeOnLanInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WakeOnLanInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.MacAddress.IsSet() {
		toSerialize["MacAddress"] = o.MacAddress.Get()
	}
	if !IsNil(o.Port) {
		toSerialize["Port"] = o.Port
	}
	return toSerialize, nil
}

type NullableWakeOnLanInfo struct {
	value *WakeOnLanInfo
	isSet bool
}

func (v NullableWakeOnLanInfo) Get() *WakeOnLanInfo {
	return v.value
}

func (v *NullableWakeOnLanInfo) Set(val *WakeOnLanInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableWakeOnLanInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableWakeOnLanInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWakeOnLanInfo(val *WakeOnLanInfo) *NullableWakeOnLanInfo {
	return &NullableWakeOnLanInfo{value: val, isSet: true}
}

func (v NullableWakeOnLanInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWakeOnLanInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
