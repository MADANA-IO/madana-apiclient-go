# \SubscriptionServiceApi

All URIs are relative to *http://api.madana.io/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddFreeSubscription**](SubscriptionServiceApi.md#AddFreeSubscription) | **Post** /subscriptions/saas/free | 
[**AddPassTrialSubscription**](SubscriptionServiceApi.md#AddPassTrialSubscription) | **Post** /subscriptions/paas/trial | 
[**GetApplication**](SubscriptionServiceApi.md#GetApplication) | **Get** /subscriptions/active | 
[**GetCheckoutSession**](SubscriptionServiceApi.md#GetCheckoutSession) | **Get** /subscriptions/{productname}/checkout | 
[**GetCheckoutSession2**](SubscriptionServiceApi.md#GetCheckoutSession2) | **Post** /subscriptions/{productname}/{newplan} | 



## AddFreeSubscription

> *os.File AddFreeSubscription(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubscriptionServiceApi.AddFreeSubscription(context.Background()).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionServiceApi.AddFreeSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddFreeSubscription`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionServiceApi.AddFreeSubscription`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAddFreeSubscriptionRequest struct via the builder pattern


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


## AddPassTrialSubscription

> *os.File AddPassTrialSubscription(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubscriptionServiceApi.AddPassTrialSubscription(context.Background()).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionServiceApi.AddPassTrialSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddPassTrialSubscription`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionServiceApi.AddPassTrialSubscription`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAddPassTrialSubscriptionRequest struct via the builder pattern


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

> *os.File GetApplication(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubscriptionServiceApi.GetApplication(context.Background()).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionServiceApi.GetApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplication`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionServiceApi.GetApplication`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationRequest struct via the builder pattern


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


## GetCheckoutSession

> *os.File GetCheckoutSession(ctx, productname).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    productname := "productname_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubscriptionServiceApi.GetCheckoutSession(context.Background(), productname).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionServiceApi.GetCheckoutSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCheckoutSession`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionServiceApi.GetCheckoutSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productname** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCheckoutSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetCheckoutSession2

> *os.File GetCheckoutSession2(ctx, newplan, productname).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    newplan := "newplan_example" // string | 
    productname := "productname_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubscriptionServiceApi.GetCheckoutSession2(context.Background(), newplan, productname).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionServiceApi.GetCheckoutSession2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCheckoutSession2`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionServiceApi.GetCheckoutSession2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**newplan** | **string** |  | 
**productname** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCheckoutSession2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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

