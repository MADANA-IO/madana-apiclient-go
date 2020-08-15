# \SystemServiceApi

All URIs are relative to *http://api.madana.io/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllObjects**](SystemServiceApi.md#GetAllObjects) | **Get** /system/health | 
[**GetApplication**](SystemServiceApi.md#GetApplication) | **Get** /system/usage | Return the current application usage.



## GetAllObjects

> *os.File GetAllObjects(ctx, )



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


## GetApplication

> map[string]map[string]interface{} GetApplication(ctx, )

Return the current application usage.

Return the current application usage

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

