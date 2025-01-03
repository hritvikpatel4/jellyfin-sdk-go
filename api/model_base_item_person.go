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

// checks if the BaseItemPerson type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaseItemPerson{}

// BaseItemPerson This is used by the api to get information about a Person within a BaseItem.
type BaseItemPerson struct {
	// Gets or sets the name.
	Name NullableString `json:"Name,omitempty"`
	// Gets or sets the identifier.
	Id *string `json:"Id,omitempty"`
	// Gets or sets the role.
	Role NullableString `json:"Role,omitempty"`
	// The person kind.
	Type *PersonKind `json:"Type,omitempty"`
	// Gets or sets the primary image tag.
	PrimaryImageTag NullableString                        `json:"PrimaryImageTag,omitempty"`
	ImageBlurHashes NullableBaseItemPersonImageBlurHashes `json:"ImageBlurHashes,omitempty"`
}

// NewBaseItemPerson instantiates a new BaseItemPerson object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseItemPerson() *BaseItemPerson {
	this := BaseItemPerson{}
	return &this
}

// NewBaseItemPersonWithDefaults instantiates a new BaseItemPerson object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseItemPersonWithDefaults() *BaseItemPerson {
	this := BaseItemPerson{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BaseItemPerson) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BaseItemPerson) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *BaseItemPerson) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *BaseItemPerson) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *BaseItemPerson) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *BaseItemPerson) UnsetName() {
	o.Name.Unset()
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BaseItemPerson) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseItemPerson) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BaseItemPerson) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BaseItemPerson) SetId(v string) {
	o.Id = &v
}

// GetRole returns the Role field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BaseItemPerson) GetRole() string {
	if o == nil || IsNil(o.Role.Get()) {
		var ret string
		return ret
	}
	return *o.Role.Get()
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BaseItemPerson) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Role.Get(), o.Role.IsSet()
}

// HasRole returns a boolean if a field has been set.
func (o *BaseItemPerson) HasRole() bool {
	if o != nil && o.Role.IsSet() {
		return true
	}

	return false
}

// SetRole gets a reference to the given NullableString and assigns it to the Role field.
func (o *BaseItemPerson) SetRole(v string) {
	o.Role.Set(&v)
}

// SetRoleNil sets the value for Role to be an explicit nil
func (o *BaseItemPerson) SetRoleNil() {
	o.Role.Set(nil)
}

// UnsetRole ensures that no value is present for Role, not even an explicit nil
func (o *BaseItemPerson) UnsetRole() {
	o.Role.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BaseItemPerson) GetType() PersonKind {
	if o == nil || IsNil(o.Type) {
		var ret PersonKind
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseItemPerson) GetTypeOk() (*PersonKind, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BaseItemPerson) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given PersonKind and assigns it to the Type field.
func (o *BaseItemPerson) SetType(v PersonKind) {
	o.Type = &v
}

// GetPrimaryImageTag returns the PrimaryImageTag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BaseItemPerson) GetPrimaryImageTag() string {
	if o == nil || IsNil(o.PrimaryImageTag.Get()) {
		var ret string
		return ret
	}
	return *o.PrimaryImageTag.Get()
}

// GetPrimaryImageTagOk returns a tuple with the PrimaryImageTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BaseItemPerson) GetPrimaryImageTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrimaryImageTag.Get(), o.PrimaryImageTag.IsSet()
}

// HasPrimaryImageTag returns a boolean if a field has been set.
func (o *BaseItemPerson) HasPrimaryImageTag() bool {
	if o != nil && o.PrimaryImageTag.IsSet() {
		return true
	}

	return false
}

// SetPrimaryImageTag gets a reference to the given NullableString and assigns it to the PrimaryImageTag field.
func (o *BaseItemPerson) SetPrimaryImageTag(v string) {
	o.PrimaryImageTag.Set(&v)
}

// SetPrimaryImageTagNil sets the value for PrimaryImageTag to be an explicit nil
func (o *BaseItemPerson) SetPrimaryImageTagNil() {
	o.PrimaryImageTag.Set(nil)
}

// UnsetPrimaryImageTag ensures that no value is present for PrimaryImageTag, not even an explicit nil
func (o *BaseItemPerson) UnsetPrimaryImageTag() {
	o.PrimaryImageTag.Unset()
}

// GetImageBlurHashes returns the ImageBlurHashes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BaseItemPerson) GetImageBlurHashes() BaseItemPersonImageBlurHashes {
	if o == nil || IsNil(o.ImageBlurHashes.Get()) {
		var ret BaseItemPersonImageBlurHashes
		return ret
	}
	return *o.ImageBlurHashes.Get()
}

// GetImageBlurHashesOk returns a tuple with the ImageBlurHashes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BaseItemPerson) GetImageBlurHashesOk() (*BaseItemPersonImageBlurHashes, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImageBlurHashes.Get(), o.ImageBlurHashes.IsSet()
}

// HasImageBlurHashes returns a boolean if a field has been set.
func (o *BaseItemPerson) HasImageBlurHashes() bool {
	if o != nil && o.ImageBlurHashes.IsSet() {
		return true
	}

	return false
}

// SetImageBlurHashes gets a reference to the given NullableBaseItemPersonImageBlurHashes and assigns it to the ImageBlurHashes field.
func (o *BaseItemPerson) SetImageBlurHashes(v BaseItemPersonImageBlurHashes) {
	o.ImageBlurHashes.Set(&v)
}

// SetImageBlurHashesNil sets the value for ImageBlurHashes to be an explicit nil
func (o *BaseItemPerson) SetImageBlurHashesNil() {
	o.ImageBlurHashes.Set(nil)
}

// UnsetImageBlurHashes ensures that no value is present for ImageBlurHashes, not even an explicit nil
func (o *BaseItemPerson) UnsetImageBlurHashes() {
	o.ImageBlurHashes.Unset()
}

func (o BaseItemPerson) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BaseItemPerson) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["Name"] = o.Name.Get()
	}
	if !IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if o.Role.IsSet() {
		toSerialize["Role"] = o.Role.Get()
	}
	if !IsNil(o.Type) {
		toSerialize["Type"] = o.Type
	}
	if o.PrimaryImageTag.IsSet() {
		toSerialize["PrimaryImageTag"] = o.PrimaryImageTag.Get()
	}
	if o.ImageBlurHashes.IsSet() {
		toSerialize["ImageBlurHashes"] = o.ImageBlurHashes.Get()
	}
	return toSerialize, nil
}

type NullableBaseItemPerson struct {
	value *BaseItemPerson
	isSet bool
}

func (v NullableBaseItemPerson) Get() *BaseItemPerson {
	return v.value
}

func (v *NullableBaseItemPerson) Set(val *BaseItemPerson) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseItemPerson) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseItemPerson) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseItemPerson(val *BaseItemPerson) *NullableBaseItemPerson {
	return &NullableBaseItemPerson{value: val, isSet: true}
}

func (v NullableBaseItemPerson) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseItemPerson) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
