# XmlNs0NodeInfoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionURL** | Pointer to **string** |  | [optional] 
**CpuFamily** | Pointer to **string** |  | [optional] 
**CpuFrequency** | Pointer to **string** |  | [optional] 
**CpuLogicalCount** | Pointer to **int32** |  | [optional] 
**CpuModel** | Pointer to **string** |  | [optional] 
**CpuPhysicalCores** | Pointer to **int32** |  | [optional] 
**HardwareBaseboard** | Pointer to **string** |  | [optional] 
**HardwareFirmware** | Pointer to **string** |  | [optional] 
**IpfsInfo** | Pointer to [**XmlNs0IPFSSystemInfo**](xml_ns0_IPFSSystemInfo.md) |  | [optional] 
**Memory** | Pointer to **string** |  | [optional] 
**OperatingSystem** | Pointer to **string** |  | [optional] 
**OperatingSystemUptime** | Pointer to **int64** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**Processors** | Pointer to **[]string** |  | [optional] 
**PublicKey** | Pointer to **string** |  | [optional] 
**SgxInfo** | Pointer to [**XmlNs0SGXInfo**](xml_ns0_SGXInfo.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewXmlNs0NodeInfoAllOf

`func NewXmlNs0NodeInfoAllOf() *XmlNs0NodeInfoAllOf`

NewXmlNs0NodeInfoAllOf instantiates a new XmlNs0NodeInfoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewXmlNs0NodeInfoAllOfWithDefaults

`func NewXmlNs0NodeInfoAllOfWithDefaults() *XmlNs0NodeInfoAllOf`

NewXmlNs0NodeInfoAllOfWithDefaults instantiates a new XmlNs0NodeInfoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionURL

`func (o *XmlNs0NodeInfoAllOf) GetConnectionURL() string`

GetConnectionURL returns the ConnectionURL field if non-nil, zero value otherwise.

### GetConnectionURLOk

`func (o *XmlNs0NodeInfoAllOf) GetConnectionURLOk() (*string, bool)`

GetConnectionURLOk returns a tuple with the ConnectionURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionURL

`func (o *XmlNs0NodeInfoAllOf) SetConnectionURL(v string)`

SetConnectionURL sets ConnectionURL field to given value.

### HasConnectionURL

`func (o *XmlNs0NodeInfoAllOf) HasConnectionURL() bool`

HasConnectionURL returns a boolean if a field has been set.

### GetCpuFamily

`func (o *XmlNs0NodeInfoAllOf) GetCpuFamily() string`

GetCpuFamily returns the CpuFamily field if non-nil, zero value otherwise.

### GetCpuFamilyOk

`func (o *XmlNs0NodeInfoAllOf) GetCpuFamilyOk() (*string, bool)`

GetCpuFamilyOk returns a tuple with the CpuFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuFamily

`func (o *XmlNs0NodeInfoAllOf) SetCpuFamily(v string)`

SetCpuFamily sets CpuFamily field to given value.

### HasCpuFamily

`func (o *XmlNs0NodeInfoAllOf) HasCpuFamily() bool`

HasCpuFamily returns a boolean if a field has been set.

### GetCpuFrequency

`func (o *XmlNs0NodeInfoAllOf) GetCpuFrequency() string`

GetCpuFrequency returns the CpuFrequency field if non-nil, zero value otherwise.

### GetCpuFrequencyOk

`func (o *XmlNs0NodeInfoAllOf) GetCpuFrequencyOk() (*string, bool)`

GetCpuFrequencyOk returns a tuple with the CpuFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuFrequency

`func (o *XmlNs0NodeInfoAllOf) SetCpuFrequency(v string)`

SetCpuFrequency sets CpuFrequency field to given value.

### HasCpuFrequency

`func (o *XmlNs0NodeInfoAllOf) HasCpuFrequency() bool`

HasCpuFrequency returns a boolean if a field has been set.

### GetCpuLogicalCount

`func (o *XmlNs0NodeInfoAllOf) GetCpuLogicalCount() int32`

GetCpuLogicalCount returns the CpuLogicalCount field if non-nil, zero value otherwise.

### GetCpuLogicalCountOk

`func (o *XmlNs0NodeInfoAllOf) GetCpuLogicalCountOk() (*int32, bool)`

GetCpuLogicalCountOk returns a tuple with the CpuLogicalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuLogicalCount

`func (o *XmlNs0NodeInfoAllOf) SetCpuLogicalCount(v int32)`

SetCpuLogicalCount sets CpuLogicalCount field to given value.

### HasCpuLogicalCount

`func (o *XmlNs0NodeInfoAllOf) HasCpuLogicalCount() bool`

HasCpuLogicalCount returns a boolean if a field has been set.

### GetCpuModel

`func (o *XmlNs0NodeInfoAllOf) GetCpuModel() string`

GetCpuModel returns the CpuModel field if non-nil, zero value otherwise.

### GetCpuModelOk

`func (o *XmlNs0NodeInfoAllOf) GetCpuModelOk() (*string, bool)`

GetCpuModelOk returns a tuple with the CpuModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuModel

`func (o *XmlNs0NodeInfoAllOf) SetCpuModel(v string)`

SetCpuModel sets CpuModel field to given value.

### HasCpuModel

`func (o *XmlNs0NodeInfoAllOf) HasCpuModel() bool`

HasCpuModel returns a boolean if a field has been set.

### GetCpuPhysicalCores

`func (o *XmlNs0NodeInfoAllOf) GetCpuPhysicalCores() int32`

GetCpuPhysicalCores returns the CpuPhysicalCores field if non-nil, zero value otherwise.

### GetCpuPhysicalCoresOk

`func (o *XmlNs0NodeInfoAllOf) GetCpuPhysicalCoresOk() (*int32, bool)`

GetCpuPhysicalCoresOk returns a tuple with the CpuPhysicalCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuPhysicalCores

`func (o *XmlNs0NodeInfoAllOf) SetCpuPhysicalCores(v int32)`

SetCpuPhysicalCores sets CpuPhysicalCores field to given value.

### HasCpuPhysicalCores

`func (o *XmlNs0NodeInfoAllOf) HasCpuPhysicalCores() bool`

HasCpuPhysicalCores returns a boolean if a field has been set.

### GetHardwareBaseboard

`func (o *XmlNs0NodeInfoAllOf) GetHardwareBaseboard() string`

GetHardwareBaseboard returns the HardwareBaseboard field if non-nil, zero value otherwise.

### GetHardwareBaseboardOk

`func (o *XmlNs0NodeInfoAllOf) GetHardwareBaseboardOk() (*string, bool)`

GetHardwareBaseboardOk returns a tuple with the HardwareBaseboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareBaseboard

`func (o *XmlNs0NodeInfoAllOf) SetHardwareBaseboard(v string)`

SetHardwareBaseboard sets HardwareBaseboard field to given value.

### HasHardwareBaseboard

`func (o *XmlNs0NodeInfoAllOf) HasHardwareBaseboard() bool`

HasHardwareBaseboard returns a boolean if a field has been set.

### GetHardwareFirmware

`func (o *XmlNs0NodeInfoAllOf) GetHardwareFirmware() string`

GetHardwareFirmware returns the HardwareFirmware field if non-nil, zero value otherwise.

### GetHardwareFirmwareOk

`func (o *XmlNs0NodeInfoAllOf) GetHardwareFirmwareOk() (*string, bool)`

GetHardwareFirmwareOk returns a tuple with the HardwareFirmware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareFirmware

`func (o *XmlNs0NodeInfoAllOf) SetHardwareFirmware(v string)`

SetHardwareFirmware sets HardwareFirmware field to given value.

### HasHardwareFirmware

`func (o *XmlNs0NodeInfoAllOf) HasHardwareFirmware() bool`

HasHardwareFirmware returns a boolean if a field has been set.

### GetIpfsInfo

`func (o *XmlNs0NodeInfoAllOf) GetIpfsInfo() XmlNs0IPFSSystemInfo`

GetIpfsInfo returns the IpfsInfo field if non-nil, zero value otherwise.

### GetIpfsInfoOk

`func (o *XmlNs0NodeInfoAllOf) GetIpfsInfoOk() (*XmlNs0IPFSSystemInfo, bool)`

GetIpfsInfoOk returns a tuple with the IpfsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpfsInfo

`func (o *XmlNs0NodeInfoAllOf) SetIpfsInfo(v XmlNs0IPFSSystemInfo)`

SetIpfsInfo sets IpfsInfo field to given value.

### HasIpfsInfo

`func (o *XmlNs0NodeInfoAllOf) HasIpfsInfo() bool`

HasIpfsInfo returns a boolean if a field has been set.

### GetMemory

`func (o *XmlNs0NodeInfoAllOf) GetMemory() string`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *XmlNs0NodeInfoAllOf) GetMemoryOk() (*string, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *XmlNs0NodeInfoAllOf) SetMemory(v string)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *XmlNs0NodeInfoAllOf) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetOperatingSystem

`func (o *XmlNs0NodeInfoAllOf) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *XmlNs0NodeInfoAllOf) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *XmlNs0NodeInfoAllOf) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *XmlNs0NodeInfoAllOf) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### GetOperatingSystemUptime

`func (o *XmlNs0NodeInfoAllOf) GetOperatingSystemUptime() int64`

GetOperatingSystemUptime returns the OperatingSystemUptime field if non-nil, zero value otherwise.

### GetOperatingSystemUptimeOk

`func (o *XmlNs0NodeInfoAllOf) GetOperatingSystemUptimeOk() (*int64, bool)`

GetOperatingSystemUptimeOk returns a tuple with the OperatingSystemUptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystemUptime

`func (o *XmlNs0NodeInfoAllOf) SetOperatingSystemUptime(v int64)`

SetOperatingSystemUptime sets OperatingSystemUptime field to given value.

### HasOperatingSystemUptime

`func (o *XmlNs0NodeInfoAllOf) HasOperatingSystemUptime() bool`

HasOperatingSystemUptime returns a boolean if a field has been set.

### GetOwner

`func (o *XmlNs0NodeInfoAllOf) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *XmlNs0NodeInfoAllOf) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *XmlNs0NodeInfoAllOf) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *XmlNs0NodeInfoAllOf) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetProcessors

`func (o *XmlNs0NodeInfoAllOf) GetProcessors() []string`

GetProcessors returns the Processors field if non-nil, zero value otherwise.

### GetProcessorsOk

`func (o *XmlNs0NodeInfoAllOf) GetProcessorsOk() (*[]string, bool)`

GetProcessorsOk returns a tuple with the Processors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessors

`func (o *XmlNs0NodeInfoAllOf) SetProcessors(v []string)`

SetProcessors sets Processors field to given value.

### HasProcessors

`func (o *XmlNs0NodeInfoAllOf) HasProcessors() bool`

HasProcessors returns a boolean if a field has been set.

### GetPublicKey

`func (o *XmlNs0NodeInfoAllOf) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *XmlNs0NodeInfoAllOf) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *XmlNs0NodeInfoAllOf) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *XmlNs0NodeInfoAllOf) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetSgxInfo

`func (o *XmlNs0NodeInfoAllOf) GetSgxInfo() XmlNs0SGXInfo`

GetSgxInfo returns the SgxInfo field if non-nil, zero value otherwise.

### GetSgxInfoOk

`func (o *XmlNs0NodeInfoAllOf) GetSgxInfoOk() (*XmlNs0SGXInfo, bool)`

GetSgxInfoOk returns a tuple with the SgxInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSgxInfo

`func (o *XmlNs0NodeInfoAllOf) SetSgxInfo(v XmlNs0SGXInfo)`

SetSgxInfo sets SgxInfo field to given value.

### HasSgxInfo

`func (o *XmlNs0NodeInfoAllOf) HasSgxInfo() bool`

HasSgxInfo returns a boolean if a field has been set.

### GetStatus

`func (o *XmlNs0NodeInfoAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *XmlNs0NodeInfoAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *XmlNs0NodeInfoAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *XmlNs0NodeInfoAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


