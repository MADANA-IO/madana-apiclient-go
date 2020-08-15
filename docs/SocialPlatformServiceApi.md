# \SocialPlatformServiceApi

All URIs are relative to *http://api.madana.io/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPlatforms**](SocialPlatformServiceApi.md#GetPlatforms) | **Get** /platforms | Used to Handle Incoming Webhooks from Facebook.
[**ListenTwitterWebhook**](SocialPlatformServiceApi.md#ListenTwitterWebhook) | **Post** /platforms/twitter | Used to Handle Incoming Webhooks from Facebook.
[**RegisterTwitterWebhook**](SocialPlatformServiceApi.md#RegisterTwitterWebhook) | **Get** /platforms/twitter | Used to Handle Incoming Webhooks from Twitter.



## GetPlatforms

> *os.File GetPlatforms(ctx, optional)

Used to Handle Incoming Webhooks from Facebook.

Used to Handle Incoming Webhooks from Facebook

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetPlatformsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetPlatformsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **optional.String**|  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListenTwitterWebhook

> *os.File ListenTwitterWebhook(ctx, optional)

Used to Handle Incoming Webhooks from Facebook.

Used to Handle Incoming Webhooks from Facebook

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListenTwitterWebhookOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListenTwitterWebhookOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **optional.String**|  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterTwitterWebhook

> *os.File RegisterTwitterWebhook(ctx, optional)

Used to Handle Incoming Webhooks from Twitter.

Used to Handle Incoming Webhooks from Twitter

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RegisterTwitterWebhookOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RegisterTwitterWebhookOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **crcToken** | **optional.String**|  | 

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

