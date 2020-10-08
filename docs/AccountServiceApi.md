# \AccountServiceApi

All URIs are relative to *http://api.madana.io/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateUser**](AccountServiceApi.md#ActivateUser) | **Get** /account/activation/{token} | 
[**CreatePasswordReset**](AccountServiceApi.md#CreatePasswordReset) | **Post** /account/password | Sends an Password reset mail to the given MailAddress.
[**RequestVerificationMail**](AccountServiceApi.md#RequestVerificationMail) | **Get** /account/verifymail | Used to request a new  activation-mail for the user.
[**UpdatePassword**](AccountServiceApi.md#UpdatePassword) | **Put** /account/password | Receives the Password reset and tries to set the provided password for the user.



## ActivateUser

> *os.File ActivateUser(ctx, token).Execute()



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
    token := "token_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountServiceApi.ActivateUser(context.Background(), token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountServiceApi.ActivateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActivateUser`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `AccountServiceApi.ActivateUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateUserRequest struct via the builder pattern


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


## CreatePasswordReset

> *os.File CreatePasswordReset(ctx).Body(body).Execute()

Sends an Password reset mail to the given MailAddress.



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
    body := openapiclient.json_MDN_MailAddress{Mail: "Mail_example"} // JsonMDNMailAddress | - the MaiAddress under which the user has signed up (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountServiceApi.CreatePasswordReset(context.Background(), ).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountServiceApi.CreatePasswordReset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePasswordReset`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `AccountServiceApi.CreatePasswordReset`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePasswordResetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**JsonMDNMailAddress**](JsonMDNMailAddress.md) | - the MaiAddress under which the user has signed up | 

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


## RequestVerificationMail

> map[string]map[string]interface{} RequestVerificationMail(ctx).Execute()

Used to request a new  activation-mail for the user.



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
    resp, r, err := api_client.AccountServiceApi.RequestVerificationMail(context.Background(), ).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountServiceApi.RequestVerificationMail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RequestVerificationMail`: map[string]map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AccountServiceApi.RequestVerificationMail`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRequestVerificationMailRequest struct via the builder pattern


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


## UpdatePassword

> *os.File UpdatePassword(ctx).Body(body).Execute()

Receives the Password reset and tries to set the provided password for the user.



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
    body := openapiclient.json_MDN_PasswordReset{Password: "Password_example", Mail: "Mail_example", Token: "Token_example"} // JsonMDNPasswordReset | - the MDN_PasswordReset Object (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountServiceApi.UpdatePassword(context.Background(), ).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountServiceApi.UpdatePassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePassword`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `AccountServiceApi.UpdatePassword`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**JsonMDNPasswordReset**](JsonMDNPasswordReset.md) | - the MDN_PasswordReset Object | 

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

