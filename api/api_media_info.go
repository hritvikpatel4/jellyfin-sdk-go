/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.10.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type MediaInfoAPI interface {

	/*
		CloseLiveStream Closes a media source.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiCloseLiveStreamRequest
	*/
	CloseLiveStream(ctx context.Context) ApiCloseLiveStreamRequest

	// CloseLiveStreamExecute executes the request
	CloseLiveStreamExecute(r ApiCloseLiveStreamRequest) (*http.Response, error)

	/*
		GetBitrateTestBytes Tests the network with a request with the size of the bitrate.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiGetBitrateTestBytesRequest
	*/
	GetBitrateTestBytes(ctx context.Context) ApiGetBitrateTestBytesRequest

	// GetBitrateTestBytesExecute executes the request
	//  @return *os.File
	GetBitrateTestBytesExecute(r ApiGetBitrateTestBytesRequest) (*os.File, *http.Response, error)

	/*
		GetPlaybackInfo Gets live playback media info for an item.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param itemId The item id.
		@return ApiGetPlaybackInfoRequest
	*/
	GetPlaybackInfo(ctx context.Context, itemId string) ApiGetPlaybackInfoRequest

	// GetPlaybackInfoExecute executes the request
	//  @return PlaybackInfoResponse
	GetPlaybackInfoExecute(r ApiGetPlaybackInfoRequest) (*PlaybackInfoResponse, *http.Response, error)

	/*
		GetPostedPlaybackInfo Gets live playback media info for an item.

		For backwards compatibility parameters can be sent via Query or Body, with Query having higher precedence.
	Query parameters are obsolete.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param itemId The item id.
		@return ApiGetPostedPlaybackInfoRequest
	*/
	GetPostedPlaybackInfo(ctx context.Context, itemId string) ApiGetPostedPlaybackInfoRequest

	// GetPostedPlaybackInfoExecute executes the request
	//  @return PlaybackInfoResponse
	GetPostedPlaybackInfoExecute(r ApiGetPostedPlaybackInfoRequest) (*PlaybackInfoResponse, *http.Response, error)

	/*
		OpenLiveStream Opens a media source.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiOpenLiveStreamRequest
	*/
	OpenLiveStream(ctx context.Context) ApiOpenLiveStreamRequest

	// OpenLiveStreamExecute executes the request
	//  @return LiveStreamResponse
	OpenLiveStreamExecute(r ApiOpenLiveStreamRequest) (*LiveStreamResponse, *http.Response, error)
}

// MediaInfoAPIService MediaInfoAPI service
type MediaInfoAPIService service

type ApiCloseLiveStreamRequest struct {
	ctx          context.Context
	ApiService   MediaInfoAPI
	liveStreamId *string
}

// The livestream id.
func (r ApiCloseLiveStreamRequest) LiveStreamId(liveStreamId string) ApiCloseLiveStreamRequest {
	r.liveStreamId = &liveStreamId
	return r
}

func (r ApiCloseLiveStreamRequest) Execute() (*http.Response, error) {
	return r.ApiService.CloseLiveStreamExecute(r)
}

/*
CloseLiveStream Closes a media source.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCloseLiveStreamRequest
*/
func (a *MediaInfoAPIService) CloseLiveStream(ctx context.Context) ApiCloseLiveStreamRequest {
	return ApiCloseLiveStreamRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *MediaInfoAPIService) CloseLiveStreamExecute(r ApiCloseLiveStreamRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MediaInfoAPIService.CloseLiveStream")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/LiveStreams/Close"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.liveStreamId == nil {
		return nil, reportError("liveStreamId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "liveStreamId", r.liveStreamId, "form", "")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["CustomAuthentication"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetBitrateTestBytesRequest struct {
	ctx        context.Context
	ApiService MediaInfoAPI
	size       *int32
}

// The bitrate. Defaults to 102400.
func (r ApiGetBitrateTestBytesRequest) Size(size int32) ApiGetBitrateTestBytesRequest {
	r.size = &size
	return r
}

func (r ApiGetBitrateTestBytesRequest) Execute() (*os.File, *http.Response, error) {
	return r.ApiService.GetBitrateTestBytesExecute(r)
}

/*
GetBitrateTestBytes Tests the network with a request with the size of the bitrate.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetBitrateTestBytesRequest
*/
func (a *MediaInfoAPIService) GetBitrateTestBytes(ctx context.Context) ApiGetBitrateTestBytesRequest {
	return ApiGetBitrateTestBytesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return *os.File
func (a *MediaInfoAPIService) GetBitrateTestBytesExecute(r ApiGetBitrateTestBytesRequest) (*os.File, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *os.File
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MediaInfoAPIService.GetBitrateTestBytes")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Playback/BitrateTest"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.size != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "size", r.size, "form", "")
	} else {
		var defaultValue int32 = 102400
		r.size = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/octet-stream"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["CustomAuthentication"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetPlaybackInfoRequest struct {
	ctx        context.Context
	ApiService MediaInfoAPI
	itemId     string
	userId     *string
}

// The user id.
func (r ApiGetPlaybackInfoRequest) UserId(userId string) ApiGetPlaybackInfoRequest {
	r.userId = &userId
	return r
}

func (r ApiGetPlaybackInfoRequest) Execute() (*PlaybackInfoResponse, *http.Response, error) {
	return r.ApiService.GetPlaybackInfoExecute(r)
}

/*
GetPlaybackInfo Gets live playback media info for an item.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param itemId The item id.
	@return ApiGetPlaybackInfoRequest
*/
func (a *MediaInfoAPIService) GetPlaybackInfo(ctx context.Context, itemId string) ApiGetPlaybackInfoRequest {
	return ApiGetPlaybackInfoRequest{
		ApiService: a,
		ctx:        ctx,
		itemId:     itemId,
	}
}

// Execute executes the request
//
//	@return PlaybackInfoResponse
func (a *MediaInfoAPIService) GetPlaybackInfoExecute(r ApiGetPlaybackInfoRequest) (*PlaybackInfoResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PlaybackInfoResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MediaInfoAPIService.GetPlaybackInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Items/{itemId}/PlaybackInfo"
	localVarPath = strings.Replace(localVarPath, "{"+"itemId"+"}", url.PathEscape(parameterValueToString(r.itemId, "itemId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.userId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "userId", r.userId, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/json; profile=CamelCase", "application/json; profile=PascalCase"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["CustomAuthentication"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetPostedPlaybackInfoRequest struct {
	ctx                  context.Context
	ApiService           MediaInfoAPI
	itemId               string
	userId               *string
	maxStreamingBitrate  *int32
	startTimeTicks       *int64
	audioStreamIndex     *int32
	subtitleStreamIndex  *int32
	maxAudioChannels     *int32
	mediaSourceId        *string
	liveStreamId         *string
	autoOpenLiveStream   *bool
	enableDirectPlay     *bool
	enableDirectStream   *bool
	enableTranscoding    *bool
	allowVideoStreamCopy *bool
	allowAudioStreamCopy *bool
	playbackInfoDto      *PlaybackInfoDto
}

// The user id.
// Deprecated
func (r ApiGetPostedPlaybackInfoRequest) UserId(userId string) ApiGetPostedPlaybackInfoRequest {
	r.userId = &userId
	return r
}

// The maximum streaming bitrate.
// Deprecated
func (r ApiGetPostedPlaybackInfoRequest) MaxStreamingBitrate(maxStreamingBitrate int32) ApiGetPostedPlaybackInfoRequest {
	r.maxStreamingBitrate = &maxStreamingBitrate
	return r
}

// The start time in ticks.
// Deprecated
func (r ApiGetPostedPlaybackInfoRequest) StartTimeTicks(startTimeTicks int64) ApiGetPostedPlaybackInfoRequest {
	r.startTimeTicks = &startTimeTicks
	return r
}

// The audio stream index.
// Deprecated
func (r ApiGetPostedPlaybackInfoRequest) AudioStreamIndex(audioStreamIndex int32) ApiGetPostedPlaybackInfoRequest {
	r.audioStreamIndex = &audioStreamIndex
	return r
}

// The subtitle stream index.
// Deprecated
func (r ApiGetPostedPlaybackInfoRequest) SubtitleStreamIndex(subtitleStreamIndex int32) ApiGetPostedPlaybackInfoRequest {
	r.subtitleStreamIndex = &subtitleStreamIndex
	return r
}

// The maximum number of audio channels.
// Deprecated
func (r ApiGetPostedPlaybackInfoRequest) MaxAudioChannels(maxAudioChannels int32) ApiGetPostedPlaybackInfoRequest {
	r.maxAudioChannels = &maxAudioChannels
	return r
}

// The media source id.
// Deprecated
func (r ApiGetPostedPlaybackInfoRequest) MediaSourceId(mediaSourceId string) ApiGetPostedPlaybackInfoRequest {
	r.mediaSourceId = &mediaSourceId
	return r
}

// The livestream id.
// Deprecated
func (r ApiGetPostedPlaybackInfoRequest) LiveStreamId(liveStreamId string) ApiGetPostedPlaybackInfoRequest {
	r.liveStreamId = &liveStreamId
	return r
}

// Whether to auto open the livestream.
// Deprecated
func (r ApiGetPostedPlaybackInfoRequest) AutoOpenLiveStream(autoOpenLiveStream bool) ApiGetPostedPlaybackInfoRequest {
	r.autoOpenLiveStream = &autoOpenLiveStream
	return r
}

// Whether to enable direct play. Default: true.
// Deprecated
func (r ApiGetPostedPlaybackInfoRequest) EnableDirectPlay(enableDirectPlay bool) ApiGetPostedPlaybackInfoRequest {
	r.enableDirectPlay = &enableDirectPlay
	return r
}

// Whether to enable direct stream. Default: true.
// Deprecated
func (r ApiGetPostedPlaybackInfoRequest) EnableDirectStream(enableDirectStream bool) ApiGetPostedPlaybackInfoRequest {
	r.enableDirectStream = &enableDirectStream
	return r
}

// Whether to enable transcoding. Default: true.
// Deprecated
func (r ApiGetPostedPlaybackInfoRequest) EnableTranscoding(enableTranscoding bool) ApiGetPostedPlaybackInfoRequest {
	r.enableTranscoding = &enableTranscoding
	return r
}

// Whether to allow to copy the video stream. Default: true.
// Deprecated
func (r ApiGetPostedPlaybackInfoRequest) AllowVideoStreamCopy(allowVideoStreamCopy bool) ApiGetPostedPlaybackInfoRequest {
	r.allowVideoStreamCopy = &allowVideoStreamCopy
	return r
}

// Whether to allow to copy the audio stream. Default: true.
// Deprecated
func (r ApiGetPostedPlaybackInfoRequest) AllowAudioStreamCopy(allowAudioStreamCopy bool) ApiGetPostedPlaybackInfoRequest {
	r.allowAudioStreamCopy = &allowAudioStreamCopy
	return r
}

// The playback info.
func (r ApiGetPostedPlaybackInfoRequest) PlaybackInfoDto(playbackInfoDto PlaybackInfoDto) ApiGetPostedPlaybackInfoRequest {
	r.playbackInfoDto = &playbackInfoDto
	return r
}

func (r ApiGetPostedPlaybackInfoRequest) Execute() (*PlaybackInfoResponse, *http.Response, error) {
	return r.ApiService.GetPostedPlaybackInfoExecute(r)
}

/*
GetPostedPlaybackInfo Gets live playback media info for an item.

For backwards compatibility parameters can be sent via Query or Body, with Query having higher precedence.
Query parameters are obsolete.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param itemId The item id.
 @return ApiGetPostedPlaybackInfoRequest
*/
func (a *MediaInfoAPIService) GetPostedPlaybackInfo(ctx context.Context, itemId string) ApiGetPostedPlaybackInfoRequest {
	return ApiGetPostedPlaybackInfoRequest{
		ApiService: a,
		ctx:        ctx,
		itemId:     itemId,
	}
}

// Execute executes the request
//
//	@return PlaybackInfoResponse
func (a *MediaInfoAPIService) GetPostedPlaybackInfoExecute(r ApiGetPostedPlaybackInfoRequest) (*PlaybackInfoResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PlaybackInfoResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MediaInfoAPIService.GetPostedPlaybackInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Items/{itemId}/PlaybackInfo"
	localVarPath = strings.Replace(localVarPath, "{"+"itemId"+"}", url.PathEscape(parameterValueToString(r.itemId, "itemId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.userId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "userId", r.userId, "form", "")
	}
	if r.maxStreamingBitrate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "maxStreamingBitrate", r.maxStreamingBitrate, "form", "")
	}
	if r.startTimeTicks != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startTimeTicks", r.startTimeTicks, "form", "")
	}
	if r.audioStreamIndex != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "audioStreamIndex", r.audioStreamIndex, "form", "")
	}
	if r.subtitleStreamIndex != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "subtitleStreamIndex", r.subtitleStreamIndex, "form", "")
	}
	if r.maxAudioChannels != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "maxAudioChannels", r.maxAudioChannels, "form", "")
	}
	if r.mediaSourceId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "mediaSourceId", r.mediaSourceId, "form", "")
	}
	if r.liveStreamId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "liveStreamId", r.liveStreamId, "form", "")
	}
	if r.autoOpenLiveStream != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "autoOpenLiveStream", r.autoOpenLiveStream, "form", "")
	}
	if r.enableDirectPlay != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "enableDirectPlay", r.enableDirectPlay, "form", "")
	}
	if r.enableDirectStream != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "enableDirectStream", r.enableDirectStream, "form", "")
	}
	if r.enableTranscoding != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "enableTranscoding", r.enableTranscoding, "form", "")
	}
	if r.allowVideoStreamCopy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "allowVideoStreamCopy", r.allowVideoStreamCopy, "form", "")
	}
	if r.allowAudioStreamCopy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "allowAudioStreamCopy", r.allowAudioStreamCopy, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "text/json", "application/*+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/json; profile=CamelCase", "application/json; profile=PascalCase"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.playbackInfoDto
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["CustomAuthentication"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiOpenLiveStreamRequest struct {
	ctx                                 context.Context
	ApiService                          MediaInfoAPI
	openToken                           *string
	userId                              *string
	playSessionId                       *string
	maxStreamingBitrate                 *int32
	startTimeTicks                      *int64
	audioStreamIndex                    *int32
	subtitleStreamIndex                 *int32
	maxAudioChannels                    *int32
	itemId                              *string
	enableDirectPlay                    *bool
	enableDirectStream                  *bool
	alwaysBurnInSubtitleWhenTranscoding *bool
	openLiveStreamDto                   *OpenLiveStreamDto
}

// The open token.
func (r ApiOpenLiveStreamRequest) OpenToken(openToken string) ApiOpenLiveStreamRequest {
	r.openToken = &openToken
	return r
}

// The user id.
func (r ApiOpenLiveStreamRequest) UserId(userId string) ApiOpenLiveStreamRequest {
	r.userId = &userId
	return r
}

// The play session id.
func (r ApiOpenLiveStreamRequest) PlaySessionId(playSessionId string) ApiOpenLiveStreamRequest {
	r.playSessionId = &playSessionId
	return r
}

// The maximum streaming bitrate.
func (r ApiOpenLiveStreamRequest) MaxStreamingBitrate(maxStreamingBitrate int32) ApiOpenLiveStreamRequest {
	r.maxStreamingBitrate = &maxStreamingBitrate
	return r
}

// The start time in ticks.
func (r ApiOpenLiveStreamRequest) StartTimeTicks(startTimeTicks int64) ApiOpenLiveStreamRequest {
	r.startTimeTicks = &startTimeTicks
	return r
}

// The audio stream index.
func (r ApiOpenLiveStreamRequest) AudioStreamIndex(audioStreamIndex int32) ApiOpenLiveStreamRequest {
	r.audioStreamIndex = &audioStreamIndex
	return r
}

// The subtitle stream index.
func (r ApiOpenLiveStreamRequest) SubtitleStreamIndex(subtitleStreamIndex int32) ApiOpenLiveStreamRequest {
	r.subtitleStreamIndex = &subtitleStreamIndex
	return r
}

// The maximum number of audio channels.
func (r ApiOpenLiveStreamRequest) MaxAudioChannels(maxAudioChannels int32) ApiOpenLiveStreamRequest {
	r.maxAudioChannels = &maxAudioChannels
	return r
}

// The item id.
func (r ApiOpenLiveStreamRequest) ItemId(itemId string) ApiOpenLiveStreamRequest {
	r.itemId = &itemId
	return r
}

// Whether to enable direct play. Default: true.
func (r ApiOpenLiveStreamRequest) EnableDirectPlay(enableDirectPlay bool) ApiOpenLiveStreamRequest {
	r.enableDirectPlay = &enableDirectPlay
	return r
}

// Whether to enable direct stream. Default: true.
func (r ApiOpenLiveStreamRequest) EnableDirectStream(enableDirectStream bool) ApiOpenLiveStreamRequest {
	r.enableDirectStream = &enableDirectStream
	return r
}

// Always burn-in subtitle when transcoding.
func (r ApiOpenLiveStreamRequest) AlwaysBurnInSubtitleWhenTranscoding(alwaysBurnInSubtitleWhenTranscoding bool) ApiOpenLiveStreamRequest {
	r.alwaysBurnInSubtitleWhenTranscoding = &alwaysBurnInSubtitleWhenTranscoding
	return r
}

// The open live stream dto.
func (r ApiOpenLiveStreamRequest) OpenLiveStreamDto(openLiveStreamDto OpenLiveStreamDto) ApiOpenLiveStreamRequest {
	r.openLiveStreamDto = &openLiveStreamDto
	return r
}

func (r ApiOpenLiveStreamRequest) Execute() (*LiveStreamResponse, *http.Response, error) {
	return r.ApiService.OpenLiveStreamExecute(r)
}

/*
OpenLiveStream Opens a media source.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenLiveStreamRequest
*/
func (a *MediaInfoAPIService) OpenLiveStream(ctx context.Context) ApiOpenLiveStreamRequest {
	return ApiOpenLiveStreamRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return LiveStreamResponse
func (a *MediaInfoAPIService) OpenLiveStreamExecute(r ApiOpenLiveStreamRequest) (*LiveStreamResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *LiveStreamResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MediaInfoAPIService.OpenLiveStream")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/LiveStreams/Open"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.openToken != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "openToken", r.openToken, "form", "")
	}
	if r.userId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "userId", r.userId, "form", "")
	}
	if r.playSessionId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "playSessionId", r.playSessionId, "form", "")
	}
	if r.maxStreamingBitrate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "maxStreamingBitrate", r.maxStreamingBitrate, "form", "")
	}
	if r.startTimeTicks != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startTimeTicks", r.startTimeTicks, "form", "")
	}
	if r.audioStreamIndex != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "audioStreamIndex", r.audioStreamIndex, "form", "")
	}
	if r.subtitleStreamIndex != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "subtitleStreamIndex", r.subtitleStreamIndex, "form", "")
	}
	if r.maxAudioChannels != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "maxAudioChannels", r.maxAudioChannels, "form", "")
	}
	if r.itemId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "itemId", r.itemId, "form", "")
	}
	if r.enableDirectPlay != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "enableDirectPlay", r.enableDirectPlay, "form", "")
	}
	if r.enableDirectStream != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "enableDirectStream", r.enableDirectStream, "form", "")
	}
	if r.alwaysBurnInSubtitleWhenTranscoding != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "alwaysBurnInSubtitleWhenTranscoding", r.alwaysBurnInSubtitleWhenTranscoding, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "text/json", "application/*+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/json; profile=CamelCase", "application/json; profile=PascalCase"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.openLiveStreamDto
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["CustomAuthentication"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
