# \UserServiceApi

All URIs are relative to *http://api.madana.io/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateObject**](UserServiceApi.md#CreateObject) | **Post** /users | Creates a new user object.
[**DeleteObject**](UserServiceApi.md#DeleteObject) | **Delete** /users/{username} | Deletes an User based on the provided id and securitycontext.
[**DeleteObject_0**](UserServiceApi.md#DeleteObject_0) | **Delete** /users/{username}/social/{platform}/{ident} | Deletes linked account from the user and securitycontext.
[**GetAvatars**](UserServiceApi.md#GetAvatars) | **Get** /users/{username}/avatars | 
[**GetCertificates**](UserServiceApi.md#GetCertificates) | **Get** /users/{username}/certificates | 
[**GetEnclaveHistory**](UserServiceApi.md#GetEnclaveHistory) | **Get** /users/{username}/enclavehistory | 
[**GetObject2**](UserServiceApi.md#GetObject2) | **Get** /users/{username} | 
[**SetAvatar**](UserServiceApi.md#SetAvatar) | **Post** /users/{username}/avatars | 
[**SetSettings**](UserServiceApi.md#SetSettings) | **Post** /users/{username}/settings | 
[**UpdateObject**](UserServiceApi.md#UpdateObject) | **Put** /users/{username} | Updates Userproperties based on the provided user object.



## CreateObject

> *os.File CreateObject(ctx).Referrer(referrer).Body(body).Execute()

Creates a new user object.



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
    referrer := "referrer_example" // string |  (optional)
    body := openapiclient.json_MDN_User{UserName: "UserName_example", Created: "Created_example", LastActive: "LastActive_example", Activated: "Activated_example", Image: "Image_example", SocialAccounts: []JsonMDNSocialUserObject{openapiclient.json_MDN_SocialUserObject{Image: "Image_example", Platform: "Platform_example", Ident: "Ident_example"}), FirstName: "FirstName_example", Guid: "Guid_example", LastName: "LastName_example", Credentials: openapiclient.json_MDN_UserCredentials{Password: "Password_example", Username: "Username_example"}, Settings: []JsonMDNUserSetting{openapiclient.json_MDN_UserSetting{Name: "Name_example", Id: "Id_example", Description: "Description_example", Value: "Value_example"}), Mail: "Mail_example"} // JsonMDNUser | provided user object inheriting properties and credentials (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserServiceApi.CreateObject(context.Background(), ).Referrer(referrer).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserServiceApi.CreateObject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateObject`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `UserServiceApi.CreateObject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **referrer** | **string** |  | 
 **body** | [**JsonMDNUser**](JsonMDNUser.md) | provided user object inheriting properties and credentials | 

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


## DeleteObject

> *os.File DeleteObject(ctx, username).Execute()

Deletes an User based on the provided id and securitycontext.



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
    resp, r, err := api_client.UserServiceApi.DeleteObject(context.Background(), username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserServiceApi.DeleteObject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteObject`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `UserServiceApi.DeleteObject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteObject_0

> *os.File DeleteObject_0(ctx, ident, platform, username).Execute()

Deletes linked account from the user and securitycontext.



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
    ident := "ident_example" // string | 
    platform := "platform_example" // string | 
    username := "username_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserServiceApi.DeleteObject_0(context.Background(), ident, platform, username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserServiceApi.DeleteObject_0``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteObject_0`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `UserServiceApi.DeleteObject_0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ident** | **string** |  | 
**platform** | **string** |  | 
**username** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteObject_1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAvatars

> *os.File GetAvatars(ctx, username).Execute()



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
    resp, r, err := api_client.UserServiceApi.GetAvatars(context.Background(), username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserServiceApi.GetAvatars``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAvatars`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `UserServiceApi.GetAvatars`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAvatarsRequest struct via the builder pattern


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


## GetCertificates

> *os.File GetCertificates(ctx, username).Execute()



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
    resp, r, err := api_client.UserServiceApi.GetCertificates(context.Background(), username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserServiceApi.GetCertificates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCertificates`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `UserServiceApi.GetCertificates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCertificatesRequest struct via the builder pattern


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


## GetEnclaveHistory

> *os.File GetEnclaveHistory(ctx, username).Limit(limit).Offset(offset).Execute()



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
    limit := "limit_example" // string |  (optional) (default to "30")
    offset := "offset_example" // string |  (optional) (default to "0")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserServiceApi.GetEnclaveHistory(context.Background(), username).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserServiceApi.GetEnclaveHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEnclaveHistory`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `UserServiceApi.GetEnclaveHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnclaveHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **string** |  | [default to &quot;30&quot;]
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


## GetObject2

> *os.File GetObject2(ctx, username).Execute()



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
    resp, r, err := api_client.UserServiceApi.GetObject2(context.Background(), username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserServiceApi.GetObject2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetObject2`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `UserServiceApi.GetObject2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetObject2Request struct via the builder pattern


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


## SetAvatar

> *os.File SetAvatar(ctx, username).Body(body).Execute()



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
    body := openapiclient.json_MDN_UserProfileImage{Id: "Id_example", Image: "Image_example"} // JsonMDNUserProfileImage |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserServiceApi.SetAvatar(context.Background(), username).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserServiceApi.SetAvatar``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetAvatar`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `UserServiceApi.SetAvatar`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetAvatarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**JsonMDNUserProfileImage**](JsonMDNUserProfileImage.md) |  | 

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


## SetSettings

> *os.File SetSettings(ctx, username).Body(body).Execute()



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
    body := openapiclient.json_MDN_UserSetting{Name: "Name_example", Id: "Id_example", Description: "Description_example", Value: "Value_example"} // JsonMDNUserSetting |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserServiceApi.SetSettings(context.Background(), username).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserServiceApi.SetSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetSettings`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `UserServiceApi.SetSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**JsonMDNUserSetting**](JsonMDNUserSetting.md) |  | 

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


## UpdateObject

> *os.File UpdateObject(ctx, username).Body(body).Execute()

Updates Userproperties based on the provided user object.



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
    body := openapiclient.json_MDN_User{UserName: "UserName_example", Created: "Created_example", LastActive: "LastActive_example", Activated: "Activated_example", Image: "Image_example", SocialAccounts: []JsonMDNSocialUserObject{openapiclient.json_MDN_SocialUserObject{Image: "Image_example", Platform: "Platform_example", Ident: "Ident_example"}), FirstName: "FirstName_example", Guid: "Guid_example", LastName: "LastName_example", Credentials: openapiclient.json_MDN_UserCredentials{Password: "Password_example", Username: "Username_example"}, Settings: []JsonMDNUserSetting{), Mail: "Mail_example"} // JsonMDNUser | the new user object inherting all properties that should be change (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserServiceApi.UpdateObject(context.Background(), username).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserServiceApi.UpdateObject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateObject`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `UserServiceApi.UpdateObject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**JsonMDNUser**](JsonMDNUser.md) | the new user object inherting all properties that should be change | 

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

