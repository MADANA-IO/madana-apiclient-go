# \SocialServiceApi

All URIs are relative to *http://api.madana.io/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMyProfile**](SocialServiceApi.md#GetMyProfile) | **Get** /social/profiles/me | 
[**GetPlatforms2**](SocialServiceApi.md#GetPlatforms2) | **Get** /social | Returns all Platforms / Systems that can be Connected to the MADANA Service.
[**GetRanking**](SocialServiceApi.md#GetRanking) | **Get** /social/ranking | Returns the Ranking by PTS within the System.
[**GetSocialPlatformFeed**](SocialServiceApi.md#GetSocialPlatformFeed) | **Get** /social/feed/{platform} | 
[**GetUserProfile**](SocialServiceApi.md#GetUserProfile) | **Get** /social/profiles/{username} | 
[**GetUserProfile_0**](SocialServiceApi.md#GetUserProfile_0) | **Get** /social/profiles/{username}/simple | 



## GetMyProfile

> *os.File GetMyProfile(ctx).Execute()



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
    resp, r, err := api_client.SocialServiceApi.GetMyProfile(context.Background()).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `SocialServiceApi.GetMyProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMyProfile`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `SocialServiceApi.GetMyProfile`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMyProfileRequest struct via the builder pattern


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


## GetPlatforms2

> *os.File GetPlatforms2(ctx).Execute()

Returns all Platforms / Systems that can be Connected to the MADANA Service.



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
    resp, r, err := api_client.SocialServiceApi.GetPlatforms2(context.Background()).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `SocialServiceApi.GetPlatforms2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPlatforms2`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `SocialServiceApi.GetPlatforms2`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlatforms2Request struct via the builder pattern


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


## GetRanking

> *os.File GetRanking(ctx).Limit(limit).Offset(offset).Execute()

Returns the Ranking by PTS within the System.



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
    limit := "limit_example" // string |  (optional) (default to "99")
    offset := "offset_example" // string |  (optional) (default to "0")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SocialServiceApi.GetRanking(context.Background()).Limit(limit).Offset(offset).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `SocialServiceApi.GetRanking``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRanking`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `SocialServiceApi.GetRanking`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRankingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **string** |  | [default to &quot;99&quot;]
 **offset** | **string** |  | [default to &quot;0&quot;]

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


## GetSocialPlatformFeed

> *os.File GetSocialPlatformFeed(ctx, platform).Execute()



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
    platform := "platform_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SocialServiceApi.GetSocialPlatformFeed(context.Background(), platform).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `SocialServiceApi.GetSocialPlatformFeed``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSocialPlatformFeed`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `SocialServiceApi.GetSocialPlatformFeed`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**platform** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSocialPlatformFeedRequest struct via the builder pattern


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


## GetUserProfile

> *os.File GetUserProfile(ctx, username).Simple(simple).Execute()



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
    username := "username_example" // string | 
    simple := "simple_example" // string |  (optional) (default to "false")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SocialServiceApi.GetUserProfile(context.Background(), username).Simple(simple).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `SocialServiceApi.GetUserProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserProfile`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `SocialServiceApi.GetUserProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **simple** | **string** |  | [default to &quot;false&quot;]

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


## GetUserProfile_0

> *os.File GetUserProfile_0(ctx, username).Execute()



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
    username := "username_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SocialServiceApi.GetUserProfile_0(context.Background(), username).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `SocialServiceApi.GetUserProfile_0``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserProfile_0`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `SocialServiceApi.GetUserProfile_0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserProfile_1Request struct via the builder pattern


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

