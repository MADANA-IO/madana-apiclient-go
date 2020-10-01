# \EnclaveServiceApi

All URIs are relative to *http://api.madana.io/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddHistory**](EnclaveServiceApi.md#AddHistory) | **Post** /enclaves/{uuid}/history | 
[**ApproveEnclave**](EnclaveServiceApi.md#ApproveEnclave) | **Post** /enclaves/{uuid}/approval | 
[**AssignEnclaveAgent**](EnclaveServiceApi.md#AssignEnclaveAgent) | **Post** /enclaves/{uuid}/assign | 
[**AttestateEnclave**](EnclaveServiceApi.md#AttestateEnclave) | **Post** /enclaves/{uuid}/attestation | 
[**CreateEnclaveRunRequest**](EnclaveServiceApi.md#CreateEnclaveRunRequest) | **Post** /enclaves | 
[**GetEnclave**](EnclaveServiceApi.md#GetEnclave) | **Get** /enclaves/{uuid} | 
[**GetEnclaveTypes**](EnclaveServiceApi.md#GetEnclaveTypes) | **Get** /enclaves/types | 
[**GetEnclaves**](EnclaveServiceApi.md#GetEnclaves) | **Get** /enclaves | Returns UUIDs of existing analyses.
[**KillEnclave**](EnclaveServiceApi.md#KillEnclave) | **Post** /enclaves/{uuid}/kill | 



## AddHistory

> *os.File AddHistory(ctx, uuid).Body(body).Execute()



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
    uuid := "uuid_example" // string | 
    body := openapiclient.json_SignedData{Signature: "Signature_example", Fingerpint: "Fingerpint_example", Data: "Data_example"} // JsonSignedData |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EnclaveServiceApi.AddHistory(context.Background(), uuid).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnclaveServiceApi.AddHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddHistory`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `EnclaveServiceApi.AddHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**JsonSignedData**](JsonSignedData.md) |  | 

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


## ApproveEnclave

> *os.File ApproveEnclave(ctx, uuid).Body(body).Execute()



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
    uuid := "uuid_example" // string | 
    body := openapiclient.json_EnclaveRunningAttestationApproval{NodeInfo: openapiclient.json_NodeInfo{OperatingSystemUptime: 123, IpfsInfo: openapiclient.json_IPFSSystemInfo{AgentVersion: "AgentVersion_example", SwarmConnection: "SwarmConnection_example", Id: "Id_example", PublicKey: "PublicKey_example", ProtocolVersion: "ProtocolVersion_example"}, CpuFamily: "CpuFamily_example", Processors: []string{"Processors_example"), CpuFrequency: "CpuFrequency_example", OperatingSystem: "OperatingSystem_example", Owner: "Owner_example", Status: "Status_example", CpuPhysicalCores: 123, CpuModel: "CpuModel_example", Memory: "Memory_example", HardwareFirmware: "HardwareFirmware_example", PublicKey: "PublicKey_example", ConnectionURL: "ConnectionURL_example", HardwareBaseboard: "HardwareBaseboard_example", CpuLogicalCount: 123}, EnclaveProcess: openapiclient.json_EnclaveProcess{InternalWireguardServer: "InternalWireguardServer_example", StartupCMD: "StartupCMD_example", InternalAttesationServer: "InternalAttesationServer_example", EnclaveIdent: "EnclaveIdent_example", WgInterface: openapiclient.json_WireguardInterface{Name: "Name_example", Address: "Address_example", Port: "Port_example"}, SignerIdent: "SignerIdent_example", Status: "Status_example", Process: openapiclient.json_Process{InputStream: 123, Alive: false, OutputStream: 123, ErrorStream: 123}, WireguardServer: "WireguardServer_example", EndingTime: "EndingTime_example", InternalIdent: "InternalIdent_example", EnclaveInputstream: 123, WireguardPublicKey: "WireguardPublicKey_example", PublicIdent: "PublicIdent_example", ConsoleOutput: "ConsoleOutput_example", AttestationServer: "AttestationServer_example", RemoteControlServer: "RemoteControlServer_example", StartupTime: "StartupTime_example", InternalRemoteControlServer: "InternalRemoteControlServer_example", Environment: openapiclient.json_Environment{Content: []string{"Content_example"), Published: false, Roothash: "Roothash_example", Name: "Name_example", Description: "Description_example", IpfsHash: "IpfsHash_example", Packages: []string{"Packages_example"), RootHashOffset: "RootHashOffset_example", DefaultRunConfiguration: openapiclient.json_RunConfig{Environment: map[string]string{ "Key" = "Value" }, DiskConfig: []JsonDiskConfig{openapiclient.json_Disk_config{Disk: "Disk_example", RoothashOffset: 123, Readonly: false, Roothash: "Roothash_example"}), Run: "Run_example", Args: []string{"Args_example")}, Uuid: "Uuid_example", Size: "Size_example"}}, Approved: "Approved_example"} // JsonEnclaveRunningAttestationApproval |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EnclaveServiceApi.ApproveEnclave(context.Background(), uuid).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnclaveServiceApi.ApproveEnclave``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApproveEnclave`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `EnclaveServiceApi.ApproveEnclave`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApproveEnclaveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**JsonEnclaveRunningAttestationApproval**](JsonEnclaveRunningAttestationApproval.md) |  | 

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


## AssignEnclaveAgent

> *os.File AssignEnclaveAgent(ctx, uuid).Body(body).Execute()



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
    uuid := "uuid_example" // string | 
    body := openapiclient.json_NodeInfo{OperatingSystemUptime: 123, IpfsInfo: openapiclient.json_IPFSSystemInfo{AgentVersion: "AgentVersion_example", SwarmConnection: "SwarmConnection_example", Id: "Id_example", PublicKey: "PublicKey_example", ProtocolVersion: "ProtocolVersion_example"}, CpuFamily: "CpuFamily_example", Processors: []string{"Processors_example"), CpuFrequency: "CpuFrequency_example", OperatingSystem: "OperatingSystem_example", Owner: "Owner_example", Status: "Status_example", CpuPhysicalCores: 123, CpuModel: "CpuModel_example", Memory: "Memory_example", HardwareFirmware: "HardwareFirmware_example", PublicKey: "PublicKey_example", ConnectionURL: "ConnectionURL_example", HardwareBaseboard: "HardwareBaseboard_example", CpuLogicalCount: 123} // JsonNodeInfo |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EnclaveServiceApi.AssignEnclaveAgent(context.Background(), uuid).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnclaveServiceApi.AssignEnclaveAgent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssignEnclaveAgent`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `EnclaveServiceApi.AssignEnclaveAgent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignEnclaveAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**JsonNodeInfo**](JsonNodeInfo.md) |  | 

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


## AttestateEnclave

> *os.File AttestateEnclave(ctx, uuid).Body(body).Execute()



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
    uuid := "uuid_example" // string | 
    body := openapiclient.json_EnclaveRunningAttestation{NodeInfo: , EnclaveProcess: openapiclient.json_EnclaveProcess{InternalWireguardServer: "InternalWireguardServer_example", StartupCMD: "StartupCMD_example", InternalAttesationServer: "InternalAttesationServer_example", EnclaveIdent: "EnclaveIdent_example", WgInterface: openapiclient.json_WireguardInterface{Name: "Name_example", Address: "Address_example", Port: "Port_example"}, SignerIdent: "SignerIdent_example", Status: "Status_example", Process: openapiclient.json_Process{InputStream: 123, Alive: false, OutputStream: 123, ErrorStream: 123}, WireguardServer: "WireguardServer_example", EndingTime: "EndingTime_example", InternalIdent: "InternalIdent_example", EnclaveInputstream: 123, WireguardPublicKey: "WireguardPublicKey_example", PublicIdent: "PublicIdent_example", ConsoleOutput: "ConsoleOutput_example", AttestationServer: "AttestationServer_example", RemoteControlServer: "RemoteControlServer_example", StartupTime: "StartupTime_example", InternalRemoteControlServer: "InternalRemoteControlServer_example", Environment: openapiclient.json_Environment{Content: []string{"Content_example"), Published: false, Roothash: "Roothash_example", Name: "Name_example", Description: "Description_example", IpfsHash: "IpfsHash_example", Packages: []string{"Packages_example"), RootHashOffset: "RootHashOffset_example", DefaultRunConfiguration: openapiclient.json_RunConfig{Environment: map[string]string{ "Key" = "Value" }, DiskConfig: []JsonDiskConfig{openapiclient.json_Disk_config{Disk: "Disk_example", RoothashOffset: 123, Readonly: false, Roothash: "Roothash_example"}), Run: "Run_example", Args: []string{"Args_example")}, Uuid: "Uuid_example", Size: "Size_example"}}} // JsonEnclaveRunningAttestation |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EnclaveServiceApi.AttestateEnclave(context.Background(), uuid).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnclaveServiceApi.AttestateEnclave``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AttestateEnclave`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `EnclaveServiceApi.AttestateEnclave`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttestateEnclaveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**JsonEnclaveRunningAttestation**](JsonEnclaveRunningAttestation.md) |  | 

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


## CreateEnclaveRunRequest

> *os.File CreateEnclaveRunRequest(ctx).Body(body).Execute()



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
    body := openapiclient.json_EnclaveRunRequest{EnclaveExecutionType: "EnclaveExecutionType_example", UsingDefaultRunConfig: false, EnvironmentUUID: "EnvironmentUUID_example", WireguardPublicKey: "WireguardPublicKey_example"} // JsonEnclaveRunRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EnclaveServiceApi.CreateEnclaveRunRequest(context.Background(), ).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnclaveServiceApi.CreateEnclaveRunRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEnclaveRunRequest`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `EnclaveServiceApi.CreateEnclaveRunRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEnclaveRunRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**JsonEnclaveRunRequest**](JsonEnclaveRunRequest.md) |  | 

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


## GetEnclave

> *os.File GetEnclave(ctx, uuid).Execute()



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
    uuid := "uuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EnclaveServiceApi.GetEnclave(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnclaveServiceApi.GetEnclave``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEnclave`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `EnclaveServiceApi.GetEnclave`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnclaveRequest struct via the builder pattern


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


## GetEnclaveTypes

> *os.File GetEnclaveTypes(ctx).Execute()



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
    resp, r, err := api_client.EnclaveServiceApi.GetEnclaveTypes(context.Background(), ).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnclaveServiceApi.GetEnclaveTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEnclaveTypes`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `EnclaveServiceApi.GetEnclaveTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnclaveTypesRequest struct via the builder pattern


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


## GetEnclaves

> *os.File GetEnclaves(ctx).Authorization(authorization).Created(created).Limit(limit).Offset(offset).Status(status).Execute()

Returns UUIDs of existing analyses.



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
    authorization := "authorization_example" // string | Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c (optional)
    created := "created_example" // string | - if Queryparam \"created=true\" only the UUIDs of own Requests are shown (optional) (default to "true")
    limit := "limit_example" // string | Used for offset pagination. Limit/Offset Paging would look like GET /request?limit=20&offset=100. This query would return the 20 rows starting with the 100th row (optional) (default to "30")
    offset := "offset_example" // string | Used for offset pagination. Limit/Offset Paging would look like GET /request?limit=20&offset=100. This query would return the 20 rows starting with the 100th row (optional) (default to "0")
    status := "status_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EnclaveServiceApi.GetEnclaves(context.Background(), ).Authorization(authorization).Created(created).Limit(limit).Offset(offset).Status(status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnclaveServiceApi.GetEnclaves``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEnclaves`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `EnclaveServiceApi.GetEnclaves`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEnclavesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c | 
 **created** | **string** | - if Queryparam \&quot;created&#x3D;true\&quot; only the UUIDs of own Requests are shown | [default to &quot;true&quot;]
 **limit** | **string** | Used for offset pagination. Limit/Offset Paging would look like GET /request?limit&#x3D;20&amp;offset&#x3D;100. This query would return the 20 rows starting with the 100th row | [default to &quot;30&quot;]
 **offset** | **string** | Used for offset pagination. Limit/Offset Paging would look like GET /request?limit&#x3D;20&amp;offset&#x3D;100. This query would return the 20 rows starting with the 100th row | [default to &quot;0&quot;]
 **status** | **string** |  | 

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


## KillEnclave

> *os.File KillEnclave(ctx, uuid).Execute()



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
    uuid := "uuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EnclaveServiceApi.KillEnclave(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnclaveServiceApi.KillEnclave``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KillEnclave`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `EnclaveServiceApi.KillEnclave`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKillEnclaveRequest struct via the builder pattern


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

