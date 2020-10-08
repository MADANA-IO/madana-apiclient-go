# \AuthenticationServiceApi

All URIs are relative to *http://api.madana.io/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthenticateApplication**](AuthenticationServiceApi.md#AuthenticateApplication) | **Post** /authentication/application | Authenticates a new application and returns the token.
[**AuthenticateEthereumWallet**](AuthenticationServiceApi.md#AuthenticateEthereumWallet) | **Post** /authentication/ethereum/{wallet} | 
[**AuthenticateUser**](AuthenticationServiceApi.md#AuthenticateUser) | **Post** /authentication | Authenticates a new user and returns the token (  forbidden if the credentials cannot be validated ).
[**AuthenticateWithEthereumChallenge**](AuthenticationServiceApi.md#AuthenticateWithEthereumChallenge) | **Post** /authentication/ethereum/{wallet}/challenge | 
[**GetFractalAuthenticationURL**](AuthenticationServiceApi.md#GetFractalAuthenticationURL) | **Get** /authentication/fractal | Returns the AUthorization URL to verify a Twitter Accounts.
[**GetNonceForEthereumWallet**](AuthenticationServiceApi.md#GetNonceForEthereumWallet) | **Get** /authentication/ethereum/{wallet} | Returns a nonce for the client which is used as content for the to be created signature.
[**GetObject**](AuthenticationServiceApi.md#GetObject) | **Get** /authentication | Used to validate the active connection with the API.
[**GetTwitterAuthenticationURL**](AuthenticationServiceApi.md#GetTwitterAuthenticationURL) | **Get** /authentication/twitter | Returns the AUthorization URL to verify a Twitter Accounts.
[**SetFacebookUID**](AuthenticationServiceApi.md#SetFacebookUID) | **Post** /authentication/facebook | Used as Callback URL when users have successfully authorized their facbeook account.
[**SetFractalUID**](AuthenticationServiceApi.md#SetFractalUID) | **Post** /authentication/fractal | 
[**SetTwitterUID**](AuthenticationServiceApi.md#SetTwitterUID) | **Post** /authentication/twitter | 



## AuthenticateApplication

> JsonMDNToken AuthenticateApplication(ctx).Body(body).Execute()

Authenticates a new application and returns the token.



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
    body := openapiclient.json_MDN_Certificate{Pem: "Pem_example"} // JsonMDNCertificate | the credentials used to validate the user (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticationServiceApi.AuthenticateApplication(context.Background(), ).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationServiceApi.AuthenticateApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthenticateApplication`: JsonMDNToken
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationServiceApi.AuthenticateApplication`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticateApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**JsonMDNCertificate**](JsonMDNCertificate.md) | the credentials used to validate the user | 

### Return type

[**JsonMDNToken**](json_MDN_Token.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthenticateEthereumWallet

> *os.File AuthenticateEthereumWallet(ctx, wallet).Body(body).Execute()



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
    wallet := "wallet_example" // string | the wallet which should be authenticated
    body := openapiclient.json_MDN_OAuthToken{Token: "Token_example", Verifier: "Verifier_example"} // JsonMDNOAuthToken | Token containing nonce and signate (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticationServiceApi.AuthenticateEthereumWallet(context.Background(), wallet).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationServiceApi.AuthenticateEthereumWallet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthenticateEthereumWallet`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationServiceApi.AuthenticateEthereumWallet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wallet** | **string** | the wallet which should be authenticated | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticateEthereumWalletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**JsonMDNOAuthToken**](JsonMDNOAuthToken.md) | Token containing nonce and signate | 

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


## AuthenticateUser

> JsonMDNToken AuthenticateUser(ctx).Body(body).Execute()

Authenticates a new user and returns the token (  forbidden if the credentials cannot be validated ).



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
    body := openapiclient.json_MDN_UserCredentials{Password: "Password_example", Username: "Username_example"} // JsonMDNUserCredentials | the credentials used to validate the user (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticationServiceApi.AuthenticateUser(context.Background(), ).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationServiceApi.AuthenticateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthenticateUser`: JsonMDNToken
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationServiceApi.AuthenticateUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**JsonMDNUserCredentials**](JsonMDNUserCredentials.md) | the credentials used to validate the user | 

### Return type

[**JsonMDNToken**](json_MDN_Token.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthenticateWithEthereumChallenge

> *os.File AuthenticateWithEthereumChallenge(ctx, wallet).Body(body).Execute()



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
    wallet := "wallet_example" // string | the wallet which should be authenticated
    body := openapiclient.json_MDN_OAuthToken{Token: "Token_example", Verifier: "Verifier_example"} // JsonMDNOAuthToken | Token containing nonce and signate (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticationServiceApi.AuthenticateWithEthereumChallenge(context.Background(), wallet).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationServiceApi.AuthenticateWithEthereumChallenge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthenticateWithEthereumChallenge`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationServiceApi.AuthenticateWithEthereumChallenge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wallet** | **string** | the wallet which should be authenticated | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticateWithEthereumChallengeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**JsonMDNOAuthToken**](JsonMDNOAuthToken.md) | Token containing nonce and signate | 

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


## GetFractalAuthenticationURL

> *os.File GetFractalAuthenticationURL(ctx).Execute()

Returns the AUthorization URL to verify a Twitter Accounts.



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
    resp, r, err := api_client.AuthenticationServiceApi.GetFractalAuthenticationURL(context.Background(), ).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationServiceApi.GetFractalAuthenticationURL``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFractalAuthenticationURL`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationServiceApi.GetFractalAuthenticationURL`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFractalAuthenticationURLRequest struct via the builder pattern


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


## GetNonceForEthereumWallet

> JsonMDNToken GetNonceForEthereumWallet(ctx, wallet).Authorization(authorization).Execute()

Returns a nonce for the client which is used as content for the to be created signature.



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
    wallet := "wallet_example" // string | - wallet address as String * @HTTP 417 If the address is not valid
    authorization := "authorization_example" // string | Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticationServiceApi.GetNonceForEthereumWallet(context.Background(), wallet).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationServiceApi.GetNonceForEthereumWallet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNonceForEthereumWallet`: JsonMDNToken
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationServiceApi.GetNonceForEthereumWallet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wallet** | **string** | - wallet address as String * @HTTP 417 If the address is not valid | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNonceForEthereumWalletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c | 

### Return type

[**JsonMDNToken**](json_MDN_Token.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetObject

> map[string]map[string]interface{} GetObject(ctx).Execute()

Used to validate the active connection with the API.



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
    resp, r, err := api_client.AuthenticationServiceApi.GetObject(context.Background(), ).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationServiceApi.GetObject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetObject`: map[string]map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationServiceApi.GetObject`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetObjectRequest struct via the builder pattern


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


## GetTwitterAuthenticationURL

> *os.File GetTwitterAuthenticationURL(ctx).Execute()

Returns the AUthorization URL to verify a Twitter Accounts.



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
    resp, r, err := api_client.AuthenticationServiceApi.GetTwitterAuthenticationURL(context.Background(), ).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationServiceApi.GetTwitterAuthenticationURL``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTwitterAuthenticationURL`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationServiceApi.GetTwitterAuthenticationURL`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTwitterAuthenticationURLRequest struct via the builder pattern


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


## SetFacebookUID

> *os.File SetFacebookUID(ctx).Body(body).Execute()

Used as Callback URL when users have successfully authorized their facbeook account.



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
    resp, r, err := api_client.AuthenticationServiceApi.SetFacebookUID(context.Background(), ).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationServiceApi.SetFacebookUID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetFacebookUID`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationServiceApi.SetFacebookUID`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetFacebookUIDRequest struct via the builder pattern


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


## SetFractalUID

> *os.File SetFractalUID(ctx).Body(body).Execute()



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
    resp, r, err := api_client.AuthenticationServiceApi.SetFractalUID(context.Background(), ).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationServiceApi.SetFractalUID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetFractalUID`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationServiceApi.SetFractalUID`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetFractalUIDRequest struct via the builder pattern


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


## SetTwitterUID

> *os.File SetTwitterUID(ctx).Body(body).Execute()



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
    body :=  // JsonMDNOAuthToken |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticationServiceApi.SetTwitterUID(context.Background(), ).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationServiceApi.SetTwitterUID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetTwitterUID`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationServiceApi.SetTwitterUID`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetTwitterUIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**JsonMDNOAuthToken**](JsonMDNOAuthToken.md) |  | 

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

