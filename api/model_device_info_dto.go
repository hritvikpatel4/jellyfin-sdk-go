/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.10.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// checks if the DeviceInfoDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceInfoDto{}

// DeviceInfoDto A DTO representing device information.
type DeviceInfoDto struct {
	// Gets or sets the name.
	Name NullableString `json:"Name,omitempty"`
	// Gets or sets the custom name.
	CustomName NullableString `json:"CustomName,omitempty"`
	// Gets or sets the access token.
	AccessToken NullableString `json:"AccessToken,omitempty"`
	// Gets or sets the identifier.
	Id NullableString `json:"Id,omitempty"`
	// Gets or sets the last name of the user.
	LastUserName NullableString `json:"LastUserName,omitempty"`
	// Gets or sets the name of the application.
	AppName NullableString `json:"AppName,omitempty"`
	// Gets or sets the application version.
	AppVersion NullableString `json:"AppVersion,omitempty"`
	// Gets or sets the last user identifier.
	LastUserId NullableString `json:"LastUserId,omitempty"`
	// Gets or sets the date last modified.
	DateLastActivity NullableTime `json:"DateLastActivity,omitempty"`
	// Gets or sets the capabilities.
	Capabilities *ClientCapabilitiesDto `json:"Capabilities,omitempty"`
	// Gets or sets the icon URL.
	IconUrl NullableString `json:"IconUrl,omitempty"`
}

// NewDeviceInfoDto instantiates a new DeviceInfoDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceInfoDto() *DeviceInfoDto {
	this := DeviceInfoDto{}
	return &this
}

// NewDeviceInfoDtoWithDefaults instantiates a new DeviceInfoDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceInfoDtoWithDefaults() *DeviceInfoDto {
	this := DeviceInfoDto{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceInfoDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceInfoDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *DeviceInfoDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *DeviceInfoDto) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *DeviceInfoDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *DeviceInfoDto) UnsetName() {
	o.Name.Unset()
}

// GetCustomName returns the CustomName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceInfoDto) GetCustomName() string {
	if o == nil || IsNil(o.CustomName.Get()) {
		var ret string
		return ret
	}
	return *o.CustomName.Get()
}

// GetCustomNameOk returns a tuple with the CustomName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceInfoDto) GetCustomNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomName.Get(), o.CustomName.IsSet()
}

// HasCustomName returns a boolean if a field has been set.
func (o *DeviceInfoDto) HasCustomName() bool {
	if o != nil && o.CustomName.IsSet() {
		return true
	}

	return false
}

// SetCustomName gets a reference to the given NullableString and assigns it to the CustomName field.
func (o *DeviceInfoDto) SetCustomName(v string) {
	o.CustomName.Set(&v)
}

// SetCustomNameNil sets the value for CustomName to be an explicit nil
func (o *DeviceInfoDto) SetCustomNameNil() {
	o.CustomName.Set(nil)
}

// UnsetCustomName ensures that no value is present for CustomName, not even an explicit nil
func (o *DeviceInfoDto) UnsetCustomName() {
	o.CustomName.Unset()
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceInfoDto) GetAccessToken() string {
	if o == nil || IsNil(o.AccessToken.Get()) {
		var ret string
		return ret
	}
	return *o.AccessToken.Get()
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceInfoDto) GetAccessTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccessToken.Get(), o.AccessToken.IsSet()
}

// HasAccessToken returns a boolean if a field has been set.
func (o *DeviceInfoDto) HasAccessToken() bool {
	if o != nil && o.AccessToken.IsSet() {
		return true
	}

	return false
}

// SetAccessToken gets a reference to the given NullableString and assigns it to the AccessToken field.
func (o *DeviceInfoDto) SetAccessToken(v string) {
	o.AccessToken.Set(&v)
}

// SetAccessTokenNil sets the value for AccessToken to be an explicit nil
func (o *DeviceInfoDto) SetAccessTokenNil() {
	o.AccessToken.Set(nil)
}

// UnsetAccessToken ensures that no value is present for AccessToken, not even an explicit nil
func (o *DeviceInfoDto) UnsetAccessToken() {
	o.AccessToken.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceInfoDto) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceInfoDto) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *DeviceInfoDto) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *DeviceInfoDto) SetId(v string) {
	o.Id.Set(&v)
}

// SetIdNil sets the value for Id to be an explicit nil
func (o *DeviceInfoDto) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *DeviceInfoDto) UnsetId() {
	o.Id.Unset()
}

// GetLastUserName returns the LastUserName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceInfoDto) GetLastUserName() string {
	if o == nil || IsNil(o.LastUserName.Get()) {
		var ret string
		return ret
	}
	return *o.LastUserName.Get()
}

// GetLastUserNameOk returns a tuple with the LastUserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceInfoDto) GetLastUserNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUserName.Get(), o.LastUserName.IsSet()
}

// HasLastUserName returns a boolean if a field has been set.
func (o *DeviceInfoDto) HasLastUserName() bool {
	if o != nil && o.LastUserName.IsSet() {
		return true
	}

	return false
}

// SetLastUserName gets a reference to the given NullableString and assigns it to the LastUserName field.
func (o *DeviceInfoDto) SetLastUserName(v string) {
	o.LastUserName.Set(&v)
}

// SetLastUserNameNil sets the value for LastUserName to be an explicit nil
func (o *DeviceInfoDto) SetLastUserNameNil() {
	o.LastUserName.Set(nil)
}

// UnsetLastUserName ensures that no value is present for LastUserName, not even an explicit nil
func (o *DeviceInfoDto) UnsetLastUserName() {
	o.LastUserName.Unset()
}

// GetAppName returns the AppName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceInfoDto) GetAppName() string {
	if o == nil || IsNil(o.AppName.Get()) {
		var ret string
		return ret
	}
	return *o.AppName.Get()
}

// GetAppNameOk returns a tuple with the AppName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceInfoDto) GetAppNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AppName.Get(), o.AppName.IsSet()
}

// HasAppName returns a boolean if a field has been set.
func (o *DeviceInfoDto) HasAppName() bool {
	if o != nil && o.AppName.IsSet() {
		return true
	}

	return false
}

// SetAppName gets a reference to the given NullableString and assigns it to the AppName field.
func (o *DeviceInfoDto) SetAppName(v string) {
	o.AppName.Set(&v)
}

// SetAppNameNil sets the value for AppName to be an explicit nil
func (o *DeviceInfoDto) SetAppNameNil() {
	o.AppName.Set(nil)
}

// UnsetAppName ensures that no value is present for AppName, not even an explicit nil
func (o *DeviceInfoDto) UnsetAppName() {
	o.AppName.Unset()
}

// GetAppVersion returns the AppVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceInfoDto) GetAppVersion() string {
	if o == nil || IsNil(o.AppVersion.Get()) {
		var ret string
		return ret
	}
	return *o.AppVersion.Get()
}

// GetAppVersionOk returns a tuple with the AppVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceInfoDto) GetAppVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AppVersion.Get(), o.AppVersion.IsSet()
}

// HasAppVersion returns a boolean if a field has been set.
func (o *DeviceInfoDto) HasAppVersion() bool {
	if o != nil && o.AppVersion.IsSet() {
		return true
	}

	return false
}

// SetAppVersion gets a reference to the given NullableString and assigns it to the AppVersion field.
func (o *DeviceInfoDto) SetAppVersion(v string) {
	o.AppVersion.Set(&v)
}

// SetAppVersionNil sets the value for AppVersion to be an explicit nil
func (o *DeviceInfoDto) SetAppVersionNil() {
	o.AppVersion.Set(nil)
}

// UnsetAppVersion ensures that no value is present for AppVersion, not even an explicit nil
func (o *DeviceInfoDto) UnsetAppVersion() {
	o.AppVersion.Unset()
}

// GetLastUserId returns the LastUserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceInfoDto) GetLastUserId() string {
	if o == nil || IsNil(o.LastUserId.Get()) {
		var ret string
		return ret
	}
	return *o.LastUserId.Get()
}

// GetLastUserIdOk returns a tuple with the LastUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceInfoDto) GetLastUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUserId.Get(), o.LastUserId.IsSet()
}

// HasLastUserId returns a boolean if a field has been set.
func (o *DeviceInfoDto) HasLastUserId() bool {
	if o != nil && o.LastUserId.IsSet() {
		return true
	}

	return false
}

// SetLastUserId gets a reference to the given NullableString and assigns it to the LastUserId field.
func (o *DeviceInfoDto) SetLastUserId(v string) {
	o.LastUserId.Set(&v)
}

// SetLastUserIdNil sets the value for LastUserId to be an explicit nil
func (o *DeviceInfoDto) SetLastUserIdNil() {
	o.LastUserId.Set(nil)
}

// UnsetLastUserId ensures that no value is present for LastUserId, not even an explicit nil
func (o *DeviceInfoDto) UnsetLastUserId() {
	o.LastUserId.Unset()
}

// GetDateLastActivity returns the DateLastActivity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceInfoDto) GetDateLastActivity() time.Time {
	if o == nil || IsNil(o.DateLastActivity.Get()) {
		var ret time.Time
		return ret
	}
	return *o.DateLastActivity.Get()
}

// GetDateLastActivityOk returns a tuple with the DateLastActivity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceInfoDto) GetDateLastActivityOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DateLastActivity.Get(), o.DateLastActivity.IsSet()
}

// HasDateLastActivity returns a boolean if a field has been set.
func (o *DeviceInfoDto) HasDateLastActivity() bool {
	if o != nil && o.DateLastActivity.IsSet() {
		return true
	}

	return false
}

// SetDateLastActivity gets a reference to the given NullableTime and assigns it to the DateLastActivity field.
func (o *DeviceInfoDto) SetDateLastActivity(v time.Time) {
	o.DateLastActivity.Set(&v)
}

// SetDateLastActivityNil sets the value for DateLastActivity to be an explicit nil
func (o *DeviceInfoDto) SetDateLastActivityNil() {
	o.DateLastActivity.Set(nil)
}

// UnsetDateLastActivity ensures that no value is present for DateLastActivity, not even an explicit nil
func (o *DeviceInfoDto) UnsetDateLastActivity() {
	o.DateLastActivity.Unset()
}

// GetCapabilities returns the Capabilities field value if set, zero value otherwise.
func (o *DeviceInfoDto) GetCapabilities() ClientCapabilitiesDto {
	if o == nil || IsNil(o.Capabilities) {
		var ret ClientCapabilitiesDto
		return ret
	}
	return *o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceInfoDto) GetCapabilitiesOk() (*ClientCapabilitiesDto, bool) {
	if o == nil || IsNil(o.Capabilities) {
		return nil, false
	}
	return o.Capabilities, true
}

// HasCapabilities returns a boolean if a field has been set.
func (o *DeviceInfoDto) HasCapabilities() bool {
	if o != nil && !IsNil(o.Capabilities) {
		return true
	}

	return false
}

// SetCapabilities gets a reference to the given ClientCapabilitiesDto and assigns it to the Capabilities field.
func (o *DeviceInfoDto) SetCapabilities(v ClientCapabilitiesDto) {
	o.Capabilities = &v
}

// GetIconUrl returns the IconUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceInfoDto) GetIconUrl() string {
	if o == nil || IsNil(o.IconUrl.Get()) {
		var ret string
		return ret
	}
	return *o.IconUrl.Get()
}

// GetIconUrlOk returns a tuple with the IconUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceInfoDto) GetIconUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IconUrl.Get(), o.IconUrl.IsSet()
}

// HasIconUrl returns a boolean if a field has been set.
func (o *DeviceInfoDto) HasIconUrl() bool {
	if o != nil && o.IconUrl.IsSet() {
		return true
	}

	return false
}

// SetIconUrl gets a reference to the given NullableString and assigns it to the IconUrl field.
func (o *DeviceInfoDto) SetIconUrl(v string) {
	o.IconUrl.Set(&v)
}

// SetIconUrlNil sets the value for IconUrl to be an explicit nil
func (o *DeviceInfoDto) SetIconUrlNil() {
	o.IconUrl.Set(nil)
}

// UnsetIconUrl ensures that no value is present for IconUrl, not even an explicit nil
func (o *DeviceInfoDto) UnsetIconUrl() {
	o.IconUrl.Unset()
}

func (o DeviceInfoDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceInfoDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["Name"] = o.Name.Get()
	}
	if o.CustomName.IsSet() {
		toSerialize["CustomName"] = o.CustomName.Get()
	}
	if o.AccessToken.IsSet() {
		toSerialize["AccessToken"] = o.AccessToken.Get()
	}
	if o.Id.IsSet() {
		toSerialize["Id"] = o.Id.Get()
	}
	if o.LastUserName.IsSet() {
		toSerialize["LastUserName"] = o.LastUserName.Get()
	}
	if o.AppName.IsSet() {
		toSerialize["AppName"] = o.AppName.Get()
	}
	if o.AppVersion.IsSet() {
		toSerialize["AppVersion"] = o.AppVersion.Get()
	}
	if o.LastUserId.IsSet() {
		toSerialize["LastUserId"] = o.LastUserId.Get()
	}
	if o.DateLastActivity.IsSet() {
		toSerialize["DateLastActivity"] = o.DateLastActivity.Get()
	}
	if !IsNil(o.Capabilities) {
		toSerialize["Capabilities"] = o.Capabilities
	}
	if o.IconUrl.IsSet() {
		toSerialize["IconUrl"] = o.IconUrl.Get()
	}
	return toSerialize, nil
}

type NullableDeviceInfoDto struct {
	value *DeviceInfoDto
	isSet bool
}

func (v NullableDeviceInfoDto) Get() *DeviceInfoDto {
	return v.value
}

func (v *NullableDeviceInfoDto) Set(val *DeviceInfoDto) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceInfoDto) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceInfoDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceInfoDto(val *DeviceInfoDto) *NullableDeviceInfoDto {
	return &NullableDeviceInfoDto{value: val, isSet: true}
}

func (v NullableDeviceInfoDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceInfoDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
