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

// ItemSortBy These represent sort orders.
type ItemSortBy string

// List of ItemSortBy
const (
	ITEMSORTBY_DEFAULT ItemSortBy = "Default"
	ITEMSORTBY_AIRED_EPISODE_ORDER ItemSortBy = "AiredEpisodeOrder"
	ITEMSORTBY_ALBUM ItemSortBy = "Album"
	ITEMSORTBY_ALBUM_ARTIST ItemSortBy = "AlbumArtist"
	ITEMSORTBY_ARTIST ItemSortBy = "Artist"
	ITEMSORTBY_DATE_CREATED ItemSortBy = "DateCreated"
	ITEMSORTBY_OFFICIAL_RATING ItemSortBy = "OfficialRating"
	ITEMSORTBY_DATE_PLAYED ItemSortBy = "DatePlayed"
	ITEMSORTBY_PREMIERE_DATE ItemSortBy = "PremiereDate"
	ITEMSORTBY_START_DATE ItemSortBy = "StartDate"
	ITEMSORTBY_SORT_NAME ItemSortBy = "SortName"
	ITEMSORTBY_NAME ItemSortBy = "Name"
	ITEMSORTBY_RANDOM ItemSortBy = "Random"
	ITEMSORTBY_RUNTIME ItemSortBy = "Runtime"
	ITEMSORTBY_COMMUNITY_RATING ItemSortBy = "CommunityRating"
	ITEMSORTBY_PRODUCTION_YEAR ItemSortBy = "ProductionYear"
	ITEMSORTBY_PLAY_COUNT ItemSortBy = "PlayCount"
	ITEMSORTBY_CRITIC_RATING ItemSortBy = "CriticRating"
	ITEMSORTBY_IS_FOLDER ItemSortBy = "IsFolder"
	ITEMSORTBY_IS_UNPLAYED ItemSortBy = "IsUnplayed"
	ITEMSORTBY_IS_PLAYED ItemSortBy = "IsPlayed"
	ITEMSORTBY_SERIES_SORT_NAME ItemSortBy = "SeriesSortName"
	ITEMSORTBY_VIDEO_BIT_RATE ItemSortBy = "VideoBitRate"
	ITEMSORTBY_AIR_TIME ItemSortBy = "AirTime"
	ITEMSORTBY_STUDIO ItemSortBy = "Studio"
	ITEMSORTBY_IS_FAVORITE_OR_LIKED ItemSortBy = "IsFavoriteOrLiked"
	ITEMSORTBY_DATE_LAST_CONTENT_ADDED ItemSortBy = "DateLastContentAdded"
	ITEMSORTBY_SERIES_DATE_PLAYED ItemSortBy = "SeriesDatePlayed"
	ITEMSORTBY_PARENT_INDEX_NUMBER ItemSortBy = "ParentIndexNumber"
	ITEMSORTBY_INDEX_NUMBER ItemSortBy = "IndexNumber"
	ITEMSORTBY_SIMILARITY_SCORE ItemSortBy = "SimilarityScore"
	ITEMSORTBY_SEARCH_SCORE ItemSortBy = "SearchScore"
)

// All allowed values of ItemSortBy enum
var AllowedItemSortByEnumValues = []ItemSortBy{
	"Default",
	"AiredEpisodeOrder",
	"Album",
	"AlbumArtist",
	"Artist",
	"DateCreated",
	"OfficialRating",
	"DatePlayed",
	"PremiereDate",
	"StartDate",
	"SortName",
	"Name",
	"Random",
	"Runtime",
	"CommunityRating",
	"ProductionYear",
	"PlayCount",
	"CriticRating",
	"IsFolder",
	"IsUnplayed",
	"IsPlayed",
	"SeriesSortName",
	"VideoBitRate",
	"AirTime",
	"Studio",
	"IsFavoriteOrLiked",
	"DateLastContentAdded",
	"SeriesDatePlayed",
	"ParentIndexNumber",
	"IndexNumber",
	"SimilarityScore",
	"SearchScore",
}

func (v *ItemSortBy) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ItemSortBy(value)
	for _, existing := range AllowedItemSortByEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ItemSortBy", value)
}

// NewItemSortByFromValue returns a pointer to a valid ItemSortBy
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewItemSortByFromValue(v string) (*ItemSortBy, error) {
	ev := ItemSortBy(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ItemSortBy: valid values are %v", v, AllowedItemSortByEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ItemSortBy) IsValid() bool {
	for _, existing := range AllowedItemSortByEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ItemSortBy value
func (v ItemSortBy) Ptr() *ItemSortBy {
	return &v
}

type NullableItemSortBy struct {
	value *ItemSortBy
	isSet bool
}

func (v NullableItemSortBy) Get() *ItemSortBy {
	return v.value
}

func (v *NullableItemSortBy) Set(val *ItemSortBy) {
	v.value = val
	v.isSet = true
}

func (v NullableItemSortBy) IsSet() bool {
	return v.isSet
}

func (v *NullableItemSortBy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemSortBy(val *ItemSortBy) *NullableItemSortBy {
	return &NullableItemSortBy{value: val, isSet: true}
}

func (v NullableItemSortBy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItemSortBy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

