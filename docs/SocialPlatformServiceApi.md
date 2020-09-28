# \SocialPlatformServiceApi

All URIs are relative to *http://api.madana.io/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPlatforms**](SocialPlatformServiceApi.md#GetPlatforms) | **Get** /platforms | Used to Handle Incoming Webhooks from Facebook.
[**ListenTwitterWebhook**](SocialPlatformServiceApi.md#ListenTwitterWebhook) | **Post** /platforms/twitter | Used to Handle Incoming Webhooks from Facebook.
[**RegisterTwitterWebhook**](SocialPlatformServiceApi.md#RegisterTwitterWebhook) | **Get** /platforms/twitter | Used to Handle Incoming Webhooks from Twitter.



## GetPlatforms

> *os.File GetPlatforms(ctx).Body(body).Execute()

Used to Handle Incoming Webhooks from Facebook.



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
    body := "body_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SocialPlatformServiceApi.GetPlatforms(context.Background(), ).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SocialPlatformServiceApi.GetPlatforms``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPlatforms`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `SocialPlatformServiceApi.GetPlatforms`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPlatformsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **string** |  | 

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


## ListenTwitterWebhook

> *os.File ListenTwitterWebhook(ctx).Body(body).Execute()

Used to Handle Incoming Webhooks from Facebook.



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
    body := "body_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SocialPlatformServiceApi.ListenTwitterWebhook(context.Background(), ).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SocialPlatformServiceApi.ListenTwitterWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListenTwitterWebhook`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `SocialPlatformServiceApi.ListenTwitterWebhook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListenTwitterWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **string** |  | 

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


## RegisterTwitterWebhook

> *os.File RegisterTwitterWebhook(ctx).CrcToken(crcToken).Execute()

Used to Handle Incoming Webhooks from Twitter.



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
    crcToken := "crcToken_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SocialPlatformServiceApi.RegisterTwitterWebhook(context.Background(), ).CrcToken(crcToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SocialPlatformServiceApi.RegisterTwitterWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegisterTwitterWebhook`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `SocialPlatformServiceApi.RegisterTwitterWebhook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterTwitterWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **crcToken** | **string** |  | 

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

