# \SocialServiceApi

All URIs are relative to *http://api.madana.io/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMyProfile**](SocialServiceApi.md#GetMyProfile) | **Get** /social/profiles/me | 
[**GetPlatforms2**](SocialServiceApi.md#GetPlatforms2) | **Get** /social | Returns all Platforms / Systems that can be Connected to the MADANA Service.
[**GetRanking**](SocialServiceApi.md#GetRanking) | **Get** /social/ranking | Returns the Ranking by PTS within the System.
[**GetSocialPlatformFeed**](SocialServiceApi.md#GetSocialPlatformFeed) | **Get** /social/feed/{platform} | 
[**GetUserProfile**](SocialServiceApi.md#GetUserProfile) | **Get** /social/profiles/{username} | 
[**GetUserProfile_0**](SocialServiceApi.md#GetUserProfile_0) | **Get** /social/profiles/{username}/simple | 



## GetMyProfile

> *os.File GetMyProfile(ctx, )



### Required Parameters

This endpoint does not need any parameter.

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlatforms2

> *os.File GetPlatforms2(ctx, )

Returns all Platforms / Systems that can be Connected to the MADANA Service.

Returns all Platforms / Systems that can be Connected to the MADANA Service

### Required Parameters

This endpoint does not need any parameter.

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRanking

> *os.File GetRanking(ctx, )

Returns the Ranking by PTS within the System.

Returns the Ranking by PTS within the System

### Required Parameters

This endpoint does not need any parameter.

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSocialPlatformFeed

> *os.File GetSocialPlatformFeed(ctx, platform)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**platform** | **string**|  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserProfile

> *os.File GetUserProfile(ctx, username, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string**|  | 
 **optional** | ***GetUserProfileOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetUserProfileOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **simple** | **optional.String**|  | [default to false]

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserProfile_0

> *os.File GetUserProfile_0(ctx, username)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string**|  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

