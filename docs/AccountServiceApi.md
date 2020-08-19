# \AccountServiceApi

All URIs are relative to *http://api.madana.io/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateUser**](AccountServiceApi.md#ActivateUser) | **Get** /account/activation/{token} | 
[**CreatePasswordReset**](AccountServiceApi.md#CreatePasswordReset) | **Post** /account/password | Sends an Password reset mail to the given MailAddress.
[**RequestVerificationMail**](AccountServiceApi.md#RequestVerificationMail) | **Get** /account/verifymail | Used to request a new  activation-mail for the user.
[**UpdatePassword**](AccountServiceApi.md#UpdatePassword) | **Put** /account/password | Receives the Password reset and tries to set the provided password for the user.



## ActivateUser

> *os.File ActivateUser(ctx, token)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string**|  | 

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


## CreatePasswordReset

> *os.File CreatePasswordReset(ctx, optional)

Sends an Password reset mail to the given MailAddress.

Sends an Password reset mail to the given MailAddress

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreatePasswordResetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreatePasswordResetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of JsonMdnMailAddress**](JsonMdnMailAddress.md)| - the MaiAddress under which the user has signed up | 

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


## RequestVerificationMail

> map[string]map[string]interface{} RequestVerificationMail(ctx, )

Used to request a new  activation-mail for the user.

Used to request a new  activation-mail for the user

### Required Parameters

This endpoint does not need any parameter.

### Return type

**map[string]map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePassword

> *os.File UpdatePassword(ctx, optional)

Receives the Password reset and tries to set the provided password for the user.

Receives the Password reset and tries to set the provided password for the user. The Password will only be set if a valid token is provided

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UpdatePasswordOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdatePasswordOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of JsonMdnPasswordReset**](JsonMdnPasswordReset.md)| - the MDN_PasswordReset Object | 

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

