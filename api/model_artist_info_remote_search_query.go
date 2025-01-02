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

// checks if the ArtistInfoRemoteSearchQuery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ArtistInfoRemoteSearchQuery{}

// ArtistInfoRemoteSearchQuery struct for ArtistInfoRemoteSearchQuery
type ArtistInfoRemoteSearchQuery struct {
	SearchInfo NullableArtistInfo `json:"SearchInfo,omitempty"`
	ItemId *string `json:"ItemId,omitempty"`
	// Gets or sets the provider name to search within if set.
	SearchProviderName NullableString `json:"SearchProviderName,omitempty"`
	// Gets or sets a value indicating whether disabled providers should be included.
	IncludeDisabledProviders *bool `json:"IncludeDisabledProviders,omitempty"`
}

// NewArtistInfoRemoteSearchQuery instantiates a new ArtistInfoRemoteSearchQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArtistInfoRemoteSearchQuery() *ArtistInfoRemoteSearchQuery {
	this := ArtistInfoRemoteSearchQuery{}
	return &this
}

// NewArtistInfoRemoteSearchQueryWithDefaults instantiates a new ArtistInfoRemoteSearchQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArtistInfoRemoteSearchQueryWithDefaults() *ArtistInfoRemoteSearchQuery {
	this := ArtistInfoRemoteSearchQuery{}
	return &this
}

// GetSearchInfo returns the SearchInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ArtistInfoRemoteSearchQuery) GetSearchInfo() ArtistInfo {
	if o == nil || IsNil(o.SearchInfo.Get()) {
		var ret ArtistInfo
		return ret
	}
	return *o.SearchInfo.Get()
}

// GetSearchInfoOk returns a tuple with the SearchInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ArtistInfoRemoteSearchQuery) GetSearchInfoOk() (*ArtistInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.SearchInfo.Get(), o.SearchInfo.IsSet()
}

// HasSearchInfo returns a boolean if a field has been set.
func (o *ArtistInfoRemoteSearchQuery) HasSearchInfo() bool {
	if o != nil && o.SearchInfo.IsSet() {
		return true
	}

	return false
}

// SetSearchInfo gets a reference to the given NullableArtistInfo and assigns it to the SearchInfo field.
func (o *ArtistInfoRemoteSearchQuery) SetSearchInfo(v ArtistInfo) {
	o.SearchInfo.Set(&v)
}
// SetSearchInfoNil sets the value for SearchInfo to be an explicit nil
func (o *ArtistInfoRemoteSearchQuery) SetSearchInfoNil() {
	o.SearchInfo.Set(nil)
}

// UnsetSearchInfo ensures that no value is present for SearchInfo, not even an explicit nil
func (o *ArtistInfoRemoteSearchQuery) UnsetSearchInfo() {
	o.SearchInfo.Unset()
}

// GetItemId returns the ItemId field value if set, zero value otherwise.
func (o *ArtistInfoRemoteSearchQuery) GetItemId() string {
	if o == nil || IsNil(o.ItemId) {
		var ret string
		return ret
	}
	return *o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtistInfoRemoteSearchQuery) GetItemIdOk() (*string, bool) {
	if o == nil || IsNil(o.ItemId) {
		return nil, false
	}
	return o.ItemId, true
}

// HasItemId returns a boolean if a field has been set.
func (o *ArtistInfoRemoteSearchQuery) HasItemId() bool {
	if o != nil && !IsNil(o.ItemId) {
		return true
	}

	return false
}

// SetItemId gets a reference to the given string and assigns it to the ItemId field.
func (o *ArtistInfoRemoteSearchQuery) SetItemId(v string) {
	o.ItemId = &v
}

// GetSearchProviderName returns the SearchProviderName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ArtistInfoRemoteSearchQuery) GetSearchProviderName() string {
	if o == nil || IsNil(o.SearchProviderName.Get()) {
		var ret string
		return ret
	}
	return *o.SearchProviderName.Get()
}

// GetSearchProviderNameOk returns a tuple with the SearchProviderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ArtistInfoRemoteSearchQuery) GetSearchProviderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SearchProviderName.Get(), o.SearchProviderName.IsSet()
}

// HasSearchProviderName returns a boolean if a field has been set.
func (o *ArtistInfoRemoteSearchQuery) HasSearchProviderName() bool {
	if o != nil && o.SearchProviderName.IsSet() {
		return true
	}

	return false
}

// SetSearchProviderName gets a reference to the given NullableString and assigns it to the SearchProviderName field.
func (o *ArtistInfoRemoteSearchQuery) SetSearchProviderName(v string) {
	o.SearchProviderName.Set(&v)
}
// SetSearchProviderNameNil sets the value for SearchProviderName to be an explicit nil
func (o *ArtistInfoRemoteSearchQuery) SetSearchProviderNameNil() {
	o.SearchProviderName.Set(nil)
}

// UnsetSearchProviderName ensures that no value is present for SearchProviderName, not even an explicit nil
func (o *ArtistInfoRemoteSearchQuery) UnsetSearchProviderName() {
	o.SearchProviderName.Unset()
}

// GetIncludeDisabledProviders returns the IncludeDisabledProviders field value if set, zero value otherwise.
func (o *ArtistInfoRemoteSearchQuery) GetIncludeDisabledProviders() bool {
	if o == nil || IsNil(o.IncludeDisabledProviders) {
		var ret bool
		return ret
	}
	return *o.IncludeDisabledProviders
}

// GetIncludeDisabledProvidersOk returns a tuple with the IncludeDisabledProviders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtistInfoRemoteSearchQuery) GetIncludeDisabledProvidersOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeDisabledProviders) {
		return nil, false
	}
	return o.IncludeDisabledProviders, true
}

// HasIncludeDisabledProviders returns a boolean if a field has been set.
func (o *ArtistInfoRemoteSearchQuery) HasIncludeDisabledProviders() bool {
	if o != nil && !IsNil(o.IncludeDisabledProviders) {
		return true
	}

	return false
}

// SetIncludeDisabledProviders gets a reference to the given bool and assigns it to the IncludeDisabledProviders field.
func (o *ArtistInfoRemoteSearchQuery) SetIncludeDisabledProviders(v bool) {
	o.IncludeDisabledProviders = &v
}

func (o ArtistInfoRemoteSearchQuery) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ArtistInfoRemoteSearchQuery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.SearchInfo.IsSet() {
		toSerialize["SearchInfo"] = o.SearchInfo.Get()
	}
	if !IsNil(o.ItemId) {
		toSerialize["ItemId"] = o.ItemId
	}
	if o.SearchProviderName.IsSet() {
		toSerialize["SearchProviderName"] = o.SearchProviderName.Get()
	}
	if !IsNil(o.IncludeDisabledProviders) {
		toSerialize["IncludeDisabledProviders"] = o.IncludeDisabledProviders
	}
	return toSerialize, nil
}

type NullableArtistInfoRemoteSearchQuery struct {
	value *ArtistInfoRemoteSearchQuery
	isSet bool
}

func (v NullableArtistInfoRemoteSearchQuery) Get() *ArtistInfoRemoteSearchQuery {
	return v.value
}

func (v *NullableArtistInfoRemoteSearchQuery) Set(val *ArtistInfoRemoteSearchQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableArtistInfoRemoteSearchQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableArtistInfoRemoteSearchQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArtistInfoRemoteSearchQuery(val *ArtistInfoRemoteSearchQuery) *NullableArtistInfoRemoteSearchQuery {
	return &NullableArtistInfoRemoteSearchQuery{value: val, isSet: true}
}

func (v NullableArtistInfoRemoteSearchQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArtistInfoRemoteSearchQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

