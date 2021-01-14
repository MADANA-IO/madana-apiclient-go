# \InvoiceServiceApi

All URIs are relative to *http://api.madana.io/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetActiveSaaSSubscriptions**](InvoiceServiceApi.md#GetActiveSaaSSubscriptions) | **Get** /invoices | 



## GetActiveSaaSSubscriptions

> *os.File GetActiveSaaSSubscriptions(ctx).Dayssince(dayssince).Execute()



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
    dayssince := "dayssince_example" // string |  (optional) (default to "366")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InvoiceServiceApi.GetActiveSaaSSubscriptions(context.Background()).Dayssince(dayssince).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoiceServiceApi.GetActiveSaaSSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetActiveSaaSSubscriptions`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `InvoiceServiceApi.GetActiveSaaSSubscriptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetActiveSaaSSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dayssince** | **string** |  | [default to &quot;366&quot;]

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

