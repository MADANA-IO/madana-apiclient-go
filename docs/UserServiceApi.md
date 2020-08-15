# \UserServiceApi

All URIs are relative to *http://api.madana.io/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateObject2**](UserServiceApi.md#CreateObject2) | **Post** /users | Creates a new user object.
[**DeleteObject**](UserServiceApi.md#DeleteObject) | **Delete** /users/{username} | Deletes an User based on the provided id and securitycontext.
[**DeleteObject_0**](UserServiceApi.md#DeleteObject_0) | **Delete** /users/{username}/social/{platform}/{ident} | Deletes linked account from the user and securitycontext.
[**GetAvatars**](UserServiceApi.md#GetAvatars) | **Get** /users/{username}/avatars | 
[**GetCertificates**](UserServiceApi.md#GetCertificates) | **Get** /users/{username}/certificates | 
[**GetObject2**](UserServiceApi.md#GetObject2) | **Get** /users/{username} | 
[**SetAvatar**](UserServiceApi.md#SetAvatar) | **Post** /users/{username}/avatars | 
[**SetSettings**](UserServiceApi.md#SetSettings) | **Post** /users/{username}/settings | 
[**UpdateObject2**](UserServiceApi.md#UpdateObject2) | **Put** /users/{username} | Updates Userproperties based on the provided user object.



## CreateObject2

> *os.File CreateObject2(ctx, optional)

Creates a new user object.

Creates a new user object

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateObject2Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateObject2Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **referrer** | **optional.String**|  | 
 **body** | [**optional.Interface of JsonMdnUser**](JsonMdnUser.md)| provided user object inheriting properties and credentials | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteObject

> *os.File DeleteObject(ctx, username)

Deletes an User based on the provided id and securitycontext.

Deletes an User based on the provided id and securitycontext

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
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteObject_0

> *os.File DeleteObject_0(ctx, ident, platform, username)

Deletes linked account from the user and securitycontext.

Deletes linked account from the user and securitycontext

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ident** | **string**|  | 
**platform** | **string**|  | 
**username** | **string**|  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAvatars

> *os.File GetAvatars(ctx, username)



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


## GetCertificates

> *os.File GetCertificates(ctx, username)



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


## GetObject2

> *os.File GetObject2(ctx, username)



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


## SetAvatar

> *os.File SetAvatar(ctx, username, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string**|  | 
 **optional** | ***SetAvatarOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SetAvatarOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of JsonMdnUserProfileImage**](JsonMdnUserProfileImage.md)|  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetSettings

> *os.File SetSettings(ctx, username, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string**|  | 
 **optional** | ***SetSettingsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SetSettingsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of JsonMdnUserSetting**](JsonMdnUserSetting.md)|  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateObject2

> *os.File UpdateObject2(ctx, username, optional)

Updates Userproperties based on the provided user object.

Updates Userproperties based on the provided user object

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string**|  | 
 **optional** | ***UpdateObject2Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateObject2Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of JsonMdnUser**](JsonMdnUser.md)| the new user object inherting all properties that should be change | 

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

