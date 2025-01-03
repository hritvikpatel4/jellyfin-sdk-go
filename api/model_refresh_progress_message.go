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

// checks if the RefreshProgressMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RefreshProgressMessage{}

// RefreshProgressMessage Refresh progress message.
type RefreshProgressMessage struct {
	// Gets or sets the data.
	Data map[string]string `json:"Data,omitempty"`
	// Gets or sets the message id.
	MessageId *string `json:"MessageId,omitempty"`
	// The different kinds of messages that are used in the WebSocket api.
	MessageType *SessionMessageType `json:"MessageType,omitempty"`
}

// NewRefreshProgressMessage instantiates a new RefreshProgressMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRefreshProgressMessage() *RefreshProgressMessage {
	this := RefreshProgressMessage{}
	return &this
}

// NewRefreshProgressMessageWithDefaults instantiates a new RefreshProgressMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRefreshProgressMessageWithDefaults() *RefreshProgressMessage {
	this := RefreshProgressMessage{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RefreshProgressMessage) GetData() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RefreshProgressMessage) GetDataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return &o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *RefreshProgressMessage) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]string and assigns it to the Data field.
func (o *RefreshProgressMessage) SetData(v map[string]string) {
	o.Data = v
}

// GetMessageId returns the MessageId field value if set, zero value otherwise.
func (o *RefreshProgressMessage) GetMessageId() string {
	if o == nil || IsNil(o.MessageId) {
		var ret string
		return ret
	}
	return *o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefreshProgressMessage) GetMessageIdOk() (*string, bool) {
	if o == nil || IsNil(o.MessageId) {
		return nil, false
	}
	return o.MessageId, true
}

// HasMessageId returns a boolean if a field has been set.
func (o *RefreshProgressMessage) HasMessageId() bool {
	if o != nil && !IsNil(o.MessageId) {
		return true
	}

	return false
}

// SetMessageId gets a reference to the given string and assigns it to the MessageId field.
func (o *RefreshProgressMessage) SetMessageId(v string) {
	o.MessageId = &v
}

// GetMessageType returns the MessageType field value if set, zero value otherwise.
func (o *RefreshProgressMessage) GetMessageType() SessionMessageType {
	if o == nil || IsNil(o.MessageType) {
		var ret SessionMessageType
		return ret
	}
	return *o.MessageType
}

// GetMessageTypeOk returns a tuple with the MessageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefreshProgressMessage) GetMessageTypeOk() (*SessionMessageType, bool) {
	if o == nil || IsNil(o.MessageType) {
		return nil, false
	}
	return o.MessageType, true
}

// HasMessageType returns a boolean if a field has been set.
func (o *RefreshProgressMessage) HasMessageType() bool {
	if o != nil && !IsNil(o.MessageType) {
		return true
	}

	return false
}

// SetMessageType gets a reference to the given SessionMessageType and assigns it to the MessageType field.
func (o *RefreshProgressMessage) SetMessageType(v SessionMessageType) {
	o.MessageType = &v
}

func (o RefreshProgressMessage) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RefreshProgressMessage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["Data"] = o.Data
	}
	if !IsNil(o.MessageId) {
		toSerialize["MessageId"] = o.MessageId
	}
	if !IsNil(o.MessageType) {
		toSerialize["MessageType"] = o.MessageType
	}
	return toSerialize, nil
}

type NullableRefreshProgressMessage struct {
	value *RefreshProgressMessage
	isSet bool
}

func (v NullableRefreshProgressMessage) Get() *RefreshProgressMessage {
	return v.value
}

func (v *NullableRefreshProgressMessage) Set(val *RefreshProgressMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableRefreshProgressMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableRefreshProgressMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRefreshProgressMessage(val *RefreshProgressMessage) *NullableRefreshProgressMessage {
	return &NullableRefreshProgressMessage{value: val, isSet: true}
}

func (v NullableRefreshProgressMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRefreshProgressMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
