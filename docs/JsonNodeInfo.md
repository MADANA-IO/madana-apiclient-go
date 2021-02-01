# JsonNodeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Processors** | Pointer to **[]string** |  | [optional] 
**Memory** | Pointer to **string** |  | [optional] 
**OperatingSystemUptime** | Pointer to **float32** |  | [optional] 
**CpuLogicalCount** | Pointer to **int32** |  | [optional] 
**SgxInfo** | Pointer to [**JsonSGXInfo**](json_SGXInfo.md) |  | [optional] 
**IpfsInfo** | Pointer to [**JsonIPFSSystemInfo**](json_IPFSSystemInfo.md) |  | [optional] 
**OperatingSystem** | Pointer to **string** |  | [optional] 
**CpuModel** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**CpuFamily** | Pointer to **string** |  | [optional] 
**CpuFrequency** | Pointer to **string** |  | [optional] 
**ConnectionURL** | Pointer to **string** |  | [optional] 
**PublicKey** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**CpuPhysicalCores** | Pointer to **int32** |  | [optional] 
**HardwareFirmware** | Pointer to **string** |  | [optional] 
**HardwareBaseboard** | Pointer to **string** |  | [optional] 

## Methods

### NewJsonNodeInfo

`func NewJsonNodeInfo() *JsonNodeInfo`

NewJsonNodeInfo instantiates a new JsonNodeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJsonNodeInfoWithDefaults

`func NewJsonNodeInfoWithDefaults() *JsonNodeInfo`

NewJsonNodeInfoWithDefaults instantiates a new JsonNodeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProcessors

`func (o *JsonNodeInfo) GetProcessors() []string`

GetProcessors returns the Processors field if non-nil, zero value otherwise.

### GetProcessorsOk

`func (o *JsonNodeInfo) GetProcessorsOk() (*[]string, bool)`

GetProcessorsOk returns a tuple with the Processors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessors

`func (o *JsonNodeInfo) SetProcessors(v []string)`

SetProcessors sets Processors field to given value.

### HasProcessors

`func (o *JsonNodeInfo) HasProcessors() bool`

HasProcessors returns a boolean if a field has been set.

### GetMemory

`func (o *JsonNodeInfo) GetMemory() string`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *JsonNodeInfo) GetMemoryOk() (*string, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *JsonNodeInfo) SetMemory(v string)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *JsonNodeInfo) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetOperatingSystemUptime

`func (o *JsonNodeInfo) GetOperatingSystemUptime() float32`

GetOperatingSystemUptime returns the OperatingSystemUptime field if non-nil, zero value otherwise.

### GetOperatingSystemUptimeOk

`func (o *JsonNodeInfo) GetOperatingSystemUptimeOk() (*float32, bool)`

GetOperatingSystemUptimeOk returns a tuple with the OperatingSystemUptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystemUptime

`func (o *JsonNodeInfo) SetOperatingSystemUptime(v float32)`

SetOperatingSystemUptime sets OperatingSystemUptime field to given value.

### HasOperatingSystemUptime

`func (o *JsonNodeInfo) HasOperatingSystemUptime() bool`

HasOperatingSystemUptime returns a boolean if a field has been set.

### GetCpuLogicalCount

`func (o *JsonNodeInfo) GetCpuLogicalCount() int32`

GetCpuLogicalCount returns the CpuLogicalCount field if non-nil, zero value otherwise.

### GetCpuLogicalCountOk

`func (o *JsonNodeInfo) GetCpuLogicalCountOk() (*int32, bool)`

GetCpuLogicalCountOk returns a tuple with the CpuLogicalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuLogicalCount

`func (o *JsonNodeInfo) SetCpuLogicalCount(v int32)`

SetCpuLogicalCount sets CpuLogicalCount field to given value.

### HasCpuLogicalCount

`func (o *JsonNodeInfo) HasCpuLogicalCount() bool`

HasCpuLogicalCount returns a boolean if a field has been set.

### GetSgxInfo

`func (o *JsonNodeInfo) GetSgxInfo() JsonSGXInfo`

GetSgxInfo returns the SgxInfo field if non-nil, zero value otherwise.

### GetSgxInfoOk

`func (o *JsonNodeInfo) GetSgxInfoOk() (*JsonSGXInfo, bool)`

GetSgxInfoOk returns a tuple with the SgxInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSgxInfo

`func (o *JsonNodeInfo) SetSgxInfo(v JsonSGXInfo)`

SetSgxInfo sets SgxInfo field to given value.

### HasSgxInfo

`func (o *JsonNodeInfo) HasSgxInfo() bool`

HasSgxInfo returns a boolean if a field has been set.

### GetIpfsInfo

`func (o *JsonNodeInfo) GetIpfsInfo() JsonIPFSSystemInfo`

GetIpfsInfo returns the IpfsInfo field if non-nil, zero value otherwise.

### GetIpfsInfoOk

`func (o *JsonNodeInfo) GetIpfsInfoOk() (*JsonIPFSSystemInfo, bool)`

GetIpfsInfoOk returns a tuple with the IpfsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpfsInfo

`func (o *JsonNodeInfo) SetIpfsInfo(v JsonIPFSSystemInfo)`

SetIpfsInfo sets IpfsInfo field to given value.

### HasIpfsInfo

`func (o *JsonNodeInfo) HasIpfsInfo() bool`

HasIpfsInfo returns a boolean if a field has been set.

### GetOperatingSystem

`func (o *JsonNodeInfo) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *JsonNodeInfo) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *JsonNodeInfo) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *JsonNodeInfo) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### GetCpuModel

`func (o *JsonNodeInfo) GetCpuModel() string`

GetCpuModel returns the CpuModel field if non-nil, zero value otherwise.

### GetCpuModelOk

`func (o *JsonNodeInfo) GetCpuModelOk() (*string, bool)`

GetCpuModelOk returns a tuple with the CpuModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuModel

`func (o *JsonNodeInfo) SetCpuModel(v string)`

SetCpuModel sets CpuModel field to given value.

### HasCpuModel

`func (o *JsonNodeInfo) HasCpuModel() bool`

HasCpuModel returns a boolean if a field has been set.

### GetStatus

`func (o *JsonNodeInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *JsonNodeInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *JsonNodeInfo) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *JsonNodeInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCpuFamily

`func (o *JsonNodeInfo) GetCpuFamily() string`

GetCpuFamily returns the CpuFamily field if non-nil, zero value otherwise.

### GetCpuFamilyOk

`func (o *JsonNodeInfo) GetCpuFamilyOk() (*string, bool)`

GetCpuFamilyOk returns a tuple with the CpuFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuFamily

`func (o *JsonNodeInfo) SetCpuFamily(v string)`

SetCpuFamily sets CpuFamily field to given value.

### HasCpuFamily

`func (o *JsonNodeInfo) HasCpuFamily() bool`

HasCpuFamily returns a boolean if a field has been set.

### GetCpuFrequency

`func (o *JsonNodeInfo) GetCpuFrequency() string`

GetCpuFrequency returns the CpuFrequency field if non-nil, zero value otherwise.

### GetCpuFrequencyOk

`func (o *JsonNodeInfo) GetCpuFrequencyOk() (*string, bool)`

GetCpuFrequencyOk returns a tuple with the CpuFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuFrequency

`func (o *JsonNodeInfo) SetCpuFrequency(v string)`

SetCpuFrequency sets CpuFrequency field to given value.

### HasCpuFrequency

`func (o *JsonNodeInfo) HasCpuFrequency() bool`

HasCpuFrequency returns a boolean if a field has been set.

### GetConnectionURL

`func (o *JsonNodeInfo) GetConnectionURL() string`

GetConnectionURL returns the ConnectionURL field if non-nil, zero value otherwise.

### GetConnectionURLOk

`func (o *JsonNodeInfo) GetConnectionURLOk() (*string, bool)`

GetConnectionURLOk returns a tuple with the ConnectionURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionURL

`func (o *JsonNodeInfo) SetConnectionURL(v string)`

SetConnectionURL sets ConnectionURL field to given value.

### HasConnectionURL

`func (o *JsonNodeInfo) HasConnectionURL() bool`

HasConnectionURL returns a boolean if a field has been set.

### GetPublicKey

`func (o *JsonNodeInfo) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *JsonNodeInfo) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *JsonNodeInfo) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *JsonNodeInfo) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetOwner

`func (o *JsonNodeInfo) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *JsonNodeInfo) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *JsonNodeInfo) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *JsonNodeInfo) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetCpuPhysicalCores

`func (o *JsonNodeInfo) GetCpuPhysicalCores() int32`

GetCpuPhysicalCores returns the CpuPhysicalCores field if non-nil, zero value otherwise.

### GetCpuPhysicalCoresOk

`func (o *JsonNodeInfo) GetCpuPhysicalCoresOk() (*int32, bool)`

GetCpuPhysicalCoresOk returns a tuple with the CpuPhysicalCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuPhysicalCores

`func (o *JsonNodeInfo) SetCpuPhysicalCores(v int32)`

SetCpuPhysicalCores sets CpuPhysicalCores field to given value.

### HasCpuPhysicalCores

`func (o *JsonNodeInfo) HasCpuPhysicalCores() bool`

HasCpuPhysicalCores returns a boolean if a field has been set.

### GetHardwareFirmware

`func (o *JsonNodeInfo) GetHardwareFirmware() string`

GetHardwareFirmware returns the HardwareFirmware field if non-nil, zero value otherwise.

### GetHardwareFirmwareOk

`func (o *JsonNodeInfo) GetHardwareFirmwareOk() (*string, bool)`

GetHardwareFirmwareOk returns a tuple with the HardwareFirmware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareFirmware

`func (o *JsonNodeInfo) SetHardwareFirmware(v string)`

SetHardwareFirmware sets HardwareFirmware field to given value.

### HasHardwareFirmware

`func (o *JsonNodeInfo) HasHardwareFirmware() bool`

HasHardwareFirmware returns a boolean if a field has been set.

### GetHardwareBaseboard

`func (o *JsonNodeInfo) GetHardwareBaseboard() string`

GetHardwareBaseboard returns the HardwareBaseboard field if non-nil, zero value otherwise.

### GetHardwareBaseboardOk

`func (o *JsonNodeInfo) GetHardwareBaseboardOk() (*string, bool)`

GetHardwareBaseboardOk returns a tuple with the HardwareBaseboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareBaseboard

`func (o *JsonNodeInfo) SetHardwareBaseboard(v string)`

SetHardwareBaseboard sets HardwareBaseboard field to given value.

### HasHardwareBaseboard

`func (o *JsonNodeInfo) HasHardwareBaseboard() bool`

HasHardwareBaseboard returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


