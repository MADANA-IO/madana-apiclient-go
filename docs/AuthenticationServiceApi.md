# \AuthenticationServiceApi

All URIs are relative to *http://api.madana.io/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthenticateApplication**](AuthenticationServiceApi.md#AuthenticateApplication) | **Post** /authentication/application | Authenticates a new application and returns the token.
[**AuthenticateEthereumWallet**](AuthenticationServiceApi.md#AuthenticateEthereumWallet) | **Post** /authentication/ethereum/{wallet} | 
[**AuthenticateUser**](AuthenticationServiceApi.md#AuthenticateUser) | **Post** /authentication | Authenticates a new user and returns the token (  forbidden if the credentials cannot be validated ).
[**AuthenticateWithEthereumChallenge**](AuthenticationServiceApi.md#AuthenticateWithEthereumChallenge) | **Post** /authentication/ethereum/{wallet}/challenge | 
[**GetFractalAuthenticationURL**](AuthenticationServiceApi.md#GetFractalAuthenticationURL) | **Get** /authentication/fractal | Returns the AUthorization URL to verify a Twitter Accounts.
[**GetNonceForEthereumWallet**](AuthenticationServiceApi.md#GetNonceForEthereumWallet) | **Get** /authentication/ethereum/{wallet} | Returns a nonce for the client which is used as content for the to be created signature.
[**GetObject**](AuthenticationServiceApi.md#GetObject) | **Get** /authentication | Used to validate the active connection with the API.
[**GetTwitterAuthenticationURL**](AuthenticationServiceApi.md#GetTwitterAuthenticationURL) | **Get** /authentication/twitter | Returns the AUthorization URL to verify a Twitter Accounts.
[**SetFacebookUID**](AuthenticationServiceApi.md#SetFacebookUID) | **Post** /authentication/facebook | Used as Callback URL when users have successfully authorized their facbeook account.
[**SetFractalUID**](AuthenticationServiceApi.md#SetFractalUID) | **Post** /authentication/fractal | 
[**SetTwitterUID**](AuthenticationServiceApi.md#SetTwitterUID) | **Post** /authentication/twitter | 



## AuthenticateApplication

> JsonMdnToken AuthenticateApplication(ctx, optional)

Authenticates a new application and returns the token.

Authenticates a new application and returns the token

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AuthenticateApplicationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AuthenticateApplicationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of JsonMdnCertificate**](JsonMdnCertificate.md)| the credentials used to validate the user | 

### Return type

[**JsonMdnToken**](json_MDN_Token.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthenticateEthereumWallet

> *os.File AuthenticateEthereumWallet(ctx, wallet, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wallet** | **string**| the wallet which should be authenticated | 
 **optional** | ***AuthenticateEthereumWalletOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AuthenticateEthereumWalletOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of JsonMdnOAuthToken**](JsonMdnOAuthToken.md)| Token containing nonce and signate | 

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


## AuthenticateUser

> JsonMdnToken AuthenticateUser(ctx, optional)

Authenticates a new user and returns the token (  forbidden if the credentials cannot be validated ).

Authenticates a new user and returns the token (  forbidden if the credentials cannot be validated )

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AuthenticateUserOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AuthenticateUserOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of JsonMdnUserCredentials**](JsonMdnUserCredentials.md)| the credentials used to validate the user | 

### Return type

[**JsonMdnToken**](json_MDN_Token.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthenticateWithEthereumChallenge

> *os.File AuthenticateWithEthereumChallenge(ctx, wallet, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wallet** | **string**| the wallet which should be authenticated | 
 **optional** | ***AuthenticateWithEthereumChallengeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AuthenticateWithEthereumChallengeOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of JsonMdnOAuthToken**](JsonMdnOAuthToken.md)| Token containing nonce and signate | 

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


## GetFractalAuthenticationURL

> *os.File GetFractalAuthenticationURL(ctx, )

Returns the AUthorization URL to verify a Twitter Accounts.

Returns the AUthorization URL to verify a Twitter Accounts

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


## GetNonceForEthereumWallet

> JsonMdnToken GetNonceForEthereumWallet(ctx, wallet, optional)

Returns a nonce for the client which is used as content for the to be created signature.

Returns a nonce for the client which is used as content for the to be created signature

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wallet** | **string**| - wallet address as String * @HTTP 417 If the address is not valid | 
 **optional** | ***GetNonceForEthereumWalletOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetNonceForEthereumWalletOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **optional.String**| Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c | 

### Return type

[**JsonMdnToken**](json_MDN_Token.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetObject

> map[string]map[string]interface{} GetObject(ctx, )

Used to validate the active connection with the API.

Used to validate the active connection with the API

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


## GetTwitterAuthenticationURL

> *os.File GetTwitterAuthenticationURL(ctx, )

Returns the AUthorization URL to verify a Twitter Accounts.

Returns the AUthorization URL to verify a Twitter Accounts

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


## SetFacebookUID

> *os.File SetFacebookUID(ctx, optional)

Used as Callback URL when users have successfully authorized their facbeook account.

Used as Callback URL when users have successfully authorized their facbeook account

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SetFacebookUIDOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SetFacebookUIDOpts struct


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


## SetFractalUID

> *os.File SetFractalUID(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SetFractalUIDOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SetFractalUIDOpts struct


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


## SetTwitterUID

> *os.File SetTwitterUID(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SetTwitterUIDOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SetTwitterUIDOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of JsonMdnOAuthToken**](JsonMdnOAuthToken.md)|  | 

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

