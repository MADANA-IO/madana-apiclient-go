# \CertificateServiceApi

All URIs are relative to *http://api.madana.io/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthenticateCertificate**](CertificateServiceApi.md#AuthenticateCertificate) | **Post** /certificates | Issues certificates for logged-in users.
[**GetCertificateByFingerprint**](CertificateServiceApi.md#GetCertificateByFingerprint) | **Get** /certificates/{fingerprint} | 
[**GetRootCertificate**](CertificateServiceApi.md#GetRootCertificate) | **Get** /certificates/root | 



## AuthenticateCertificate

> JsonMdnCertificate AuthenticateCertificate(ctx, optional)

Issues certificates for logged-in users.

Issues certificates for logged-in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AuthenticateCertificateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AuthenticateCertificateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of JsonMdnData**](JsonMdnData.md)|  | 

### Return type

[**JsonMdnCertificate**](json_MDN_Certificate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCertificateByFingerprint

> *os.File GetCertificateByFingerprint(ctx, fingerprint)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fingerprint** | **string**|  | 

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


## GetRootCertificate

> *os.File GetRootCertificate(ctx, )



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

