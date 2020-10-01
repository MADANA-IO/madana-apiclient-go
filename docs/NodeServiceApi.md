# \NodeServiceApi

All URIs are relative to *http://api.madana.io/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBootstrap**](NodeServiceApi.md#GetBootstrap) | **Get** /nodes/bootstrap | 
[**GetNodes2**](NodeServiceApi.md#GetNodes2) | **Get** /nodes | 
[**PostNodeInfo**](NodeServiceApi.md#PostNodeInfo) | **Post** /nodes | 



## GetBootstrap

> *os.File GetBootstrap(ctx).Execute()



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
    resp, r, err := api_client.NodeServiceApi.GetBootstrap(context.Background(), ).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodeServiceApi.GetBootstrap``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBootstrap`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `NodeServiceApi.GetBootstrap`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBootstrapRequest struct via the builder pattern


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


## GetNodes2

> *os.File GetNodes2(ctx).Owner(owner).Execute()



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
    owner := "owner_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NodeServiceApi.GetNodes2(context.Background(), ).Owner(owner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodeServiceApi.GetNodes2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNodes2`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `NodeServiceApi.GetNodes2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNodes2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** |  | 

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


## PostNodeInfo

> *os.File PostNodeInfo(ctx).Body(body).Execute()



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
    body := openapiclient.json_NodeInfo{OperatingSystemUptime: 123, IpfsInfo: openapiclient.json_IPFSSystemInfo{AgentVersion: "AgentVersion_example", SwarmConnection: "SwarmConnection_example", Id: "Id_example", PublicKey: "PublicKey_example", ProtocolVersion: "ProtocolVersion_example"}, CpuFamily: "CpuFamily_example", Processors: []string{"Processors_example"), CpuFrequency: "CpuFrequency_example", OperatingSystem: "OperatingSystem_example", Owner: "Owner_example", Status: "Status_example", CpuPhysicalCores: 123, CpuModel: "CpuModel_example", Memory: "Memory_example", HardwareFirmware: "HardwareFirmware_example", PublicKey: "PublicKey_example", ConnectionURL: "ConnectionURL_example", HardwareBaseboard: "HardwareBaseboard_example", CpuLogicalCount: 123} // JsonNodeInfo |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NodeServiceApi.PostNodeInfo(context.Background(), ).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodeServiceApi.PostNodeInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostNodeInfo`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `NodeServiceApi.PostNodeInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostNodeInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**JsonNodeInfo**](JsonNodeInfo.md) |  | 

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

