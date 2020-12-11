# \CertificateServiceApi

All URIs are relative to *http://api.madana.io/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthenticateCertificate**](CertificateServiceApi.md#AuthenticateCertificate) | **Post** /certificates | Issues certificates for logged-in users.
[**GetCertificateByFingerprint**](CertificateServiceApi.md#GetCertificateByFingerprint) | **Get** /certificates/{fingerprint} | 
[**GetRootCertificate**](CertificateServiceApi.md#GetRootCertificate) | **Get** /certificates/root | 



## AuthenticateCertificate

> JsonMDNCertificate AuthenticateCertificate(ctx).Body(body).Execute()

Issues certificates for logged-in users.



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
    body := *openapiclient.Newjson_MDN_Data() // JsonMDNData |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CertificateServiceApi.AuthenticateCertificate(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateServiceApi.AuthenticateCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthenticateCertificate`: JsonMDNCertificate
    fmt.Fprintf(os.Stdout, "Response from `CertificateServiceApi.AuthenticateCertificate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticateCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**JsonMDNData**](JsonMDNData.md) |  | 

### Return type

[**JsonMDNCertificate**](json_MDN_Certificate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCertificateByFingerprint

> *os.File GetCertificateByFingerprint(ctx, fingerprint).Execute()



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
    fingerprint := "fingerprint_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CertificateServiceApi.GetCertificateByFingerprint(context.Background(), fingerprint).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateServiceApi.GetCertificateByFingerprint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCertificateByFingerprint`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `CertificateServiceApi.GetCertificateByFingerprint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fingerprint** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCertificateByFingerprintRequest struct via the builder pattern


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


## GetRootCertificate

> *os.File GetRootCertificate(ctx).Execute()



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
    resp, r, err := api_client.CertificateServiceApi.GetRootCertificate(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateServiceApi.GetRootCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRootCertificate`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `CertificateServiceApi.GetRootCertificate`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRootCertificateRequest struct via the builder pattern


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

