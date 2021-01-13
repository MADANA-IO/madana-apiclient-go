# \SubscriptionServiceApi

All URIs are relative to *http://api.madana.io/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddFreeSubscription**](SubscriptionServiceApi.md#AddFreeSubscription) | **Post** /subscriptions/free | 
[**GetApplication**](SubscriptionServiceApi.md#GetApplication) | **Get** /subscriptions/active | 



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

