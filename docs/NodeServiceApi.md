# \NodeServiceApi

All URIs are relative to *http://api.madana.io/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNodes2**](NodeServiceApi.md#GetNodes2) | **Get** /nodes | 
[**PostNodeInfo**](NodeServiceApi.md#PostNodeInfo) | **Post** /nodes | 



## GetNodes2

> *os.File GetNodes2(ctx, )



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


## PostNodeInfo

> *os.File PostNodeInfo(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PostNodeInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PostNodeInfoOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of JsonNodeInfo**](JsonNodeInfo.md)|  | 

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

