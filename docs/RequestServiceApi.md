# \RequestServiceApi

All URIs are relative to *http://api.madana.io/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddData**](RequestServiceApi.md#AddData) | **Post** /requests/{uuid}/data | Is used to upload and park the data till the AnalysisRequest gets processed.
[**CancelProcessing**](RequestServiceApi.md#CancelProcessing) | **Post** /requests/{uuid}/cancel | Endpoint is called from the Analysis Processing entity to submit the result.
[**CreateNewRequest**](RequestServiceApi.md#CreateNewRequest) | **Post** /requests | Endpoint used to create a new Analysis Request.
[**GetActions**](RequestServiceApi.md#GetActions) | **Get** /requests/actions | 
[**GetAgent**](RequestServiceApi.md#GetAgent) | **Get** /requests/{uuid}/agent | Is called from the APE to request all parked datasets.
[**GetAllRequests**](RequestServiceApi.md#GetAllRequests) | **Get** /requests | Returns UUIDs of existing analyses.
[**GetData**](RequestServiceApi.md#GetData) | **Get** /requests/{uuid}/data | Is called from the APE to request all parked datasets.
[**GetRequest**](RequestServiceApi.md#GetRequest) | **Get** /requests/{uuid} | Returns the details for certain Request.
[**GetResult**](RequestServiceApi.md#GetResult) | **Get** /requests/{uuid}/result | Can be called from creator to request the AnalysisResult.
[**GetStatus**](RequestServiceApi.md#GetStatus) | **Get** /requests/stats | 
[**GiveConsent**](RequestServiceApi.md#GiveConsent) | **Post** /requests/{uuid}/consent | Used to give consent for request.
[**InitRequestParameters**](RequestServiceApi.md#InitRequestParameters) | **Post** /requests/{uuid} | Endpoint used initialized addition datacollection parameters for requester.
[**SetAgent**](RequestServiceApi.md#SetAgent) | **Post** /requests/{uuid}/agent | Is called from the APE to request all parked datasets.
[**SetResult**](RequestServiceApi.md#SetResult) | **Post** /requests/{uuid}/result | Endpoint is called from the Analysis Processing entity to submit the result.



## AddData

> *os.File AddData(ctx, uuid, optional)

Is used to upload and park the data till the AnalysisRequest gets processed.

Is used to upload and park the data till the AnalysisRequest gets processed

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**|  | 
 **optional** | ***AddDataOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AddDataOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **optional.String**| Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c | 
 **body** | [**optional.Interface of JsonSignedDataUtils**](JsonSignedDataUtils.md)|  | 

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


## CancelProcessing

> *os.File CancelProcessing(ctx, uuid, optional)

Endpoint is called from the Analysis Processing entity to submit the result.

Endpoint is called from the Analysis Processing entity to submit the result

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**|  | 
 **optional** | ***CancelProcessingOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CancelProcessingOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **optional.String**| Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c | 
 **body** | [**optional.Interface of JsonSignedDataUtils**](JsonSignedDataUtils.md)|  | 

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


## CreateNewRequest

> string CreateNewRequest(ctx, optional)

Endpoint used to create a new Analysis Request.

Endpoint used to create a new Analysis Request

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateNewRequestOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateNewRequestOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **optional.String**| Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c | 
 **body** | [**optional.Interface of JsonSignedDataUtils**](JsonSignedDataUtils.md)|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetActions

> *os.File GetActions(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetActionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetActionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.String**|  | [default to 30]
 **offset** | **optional.String**|  | [default to 0]

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


## GetAgent

> *os.File GetAgent(ctx, uuid, optional)

Is called from the APE to request all parked datasets.

Is called from the APE to request all parked datasets. Returns the transmitted data for certain Request. When requesting the data, the status of the request is automatically set to processing.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**|  | 
 **optional** | ***GetAgentOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAgentOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **optional.String**| Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c | 

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


## GetAllRequests

> *os.File GetAllRequests(ctx, optional)

Returns UUIDs of existing analyses.

Returns UUIDs of existing analyses.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllRequestsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAllRequestsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **optional.String**| Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c | 
 **created** | **optional.String**| - if Queryparam \&quot;created&#x3D;true\&quot; only the UUIDs of own Requests are shown | [default to false]
 **history** | **optional.String**| - if queryparam \&quot;history\&quot; is set to true, endpoint returns all user actions. False per default. | [default to false]
 **limit** | **optional.String**| Used for offset pagination. Limit/Offset Paging would look like GET /request?limit&#x3D;20&amp;offset&#x3D;100. This query would return the 20 rows starting with the 100th row | [default to 30]
 **new** | **optional.String**| -  if Queryparam \&quot;new&#x3D;true\&quot; only the UUIDs of new Requests ( Requests the user has not participated in and still allow participation) are shown | [default to true]
 **offset** | **optional.String**| Used for offset pagination. Limit/Offset Paging would look like GET /request?limit&#x3D;20&amp;offset&#x3D;100. This query would return the 20 rows starting with the 100th row | [default to 0]
 **preview** | **optional.String**|  | [default to false]
 **ready** | **optional.String**|  | [default to false]

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


## GetData

> JsonSignedDataUtils GetData(ctx, uuid, optional)

Is called from the APE to request all parked datasets.

Is called from the APE to request all parked datasets. Returns the transmitted data for certain Request. When requesting the data, the status of the request is automatically set to processing.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**|  | 
 **optional** | ***GetDataOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetDataOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **optional.String**| Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c | 

### Return type

[**JsonSignedDataUtils**](json_SignedData_utils.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRequest

> *os.File GetRequest(ctx, uuid, optional)

Returns the details for certain Request.

Returns the details for certain Request. When requesting an analysis a view of the analysis is stored in the database

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**|  | 
 **optional** | ***GetRequestOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetRequestOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **optional.String**| Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c | 

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


## GetResult

> *os.File GetResult(ctx, uuid, optional)

Can be called from creator to request the AnalysisResult.

Can be called from creator to request the AnalysisResult.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**|  | 
 **optional** | ***GetResultOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetResultOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **optional.String**| Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c | 

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


## GetStatus

> *os.File GetStatus(ctx, )



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


## GiveConsent

> *os.File GiveConsent(ctx, uuid, optional)

Used to give consent for request.

Used to give consent for request. If the Endpoint is called from the creator of the Analysis, the status of the request is set to completed. If called by another is interpreted as giving consent to participate.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**|  | 
 **optional** | ***GiveConsentOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GiveConsentOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **optional.String**| Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c | 

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


## InitRequestParameters

> string InitRequestParameters(ctx, uuid, optional)

Endpoint used initialized addition datacollection parameters for requester.

Endpoint used initialized addition datacollection parameters for requester

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**|  | 
 **optional** | ***InitRequestParametersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a InitRequestParametersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **optional.String**| Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c | 
 **body** | **optional.String**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetAgent

> *os.File SetAgent(ctx, uuid, optional)

Is called from the APE to request all parked datasets.

Is called from the APE to request all parked datasets. Returns the transmitted data for certain Request. When requesting the data, the status of the request is automatically set to processing.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**|  | 
 **optional** | ***SetAgentOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SetAgentOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **optional.String**| Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c | 

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


## SetResult

> *os.File SetResult(ctx, uuid, optional)

Endpoint is called from the Analysis Processing entity to submit the result.

Endpoint is called from the Analysis Processing entity to submit the result

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**|  | 
 **optional** | ***SetResultOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SetResultOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **optional.String**| Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c | 
 **body** | [**optional.Interface of JsonSignedDataUtils**](JsonSignedDataUtils.md)|  | 

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

