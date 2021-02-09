/*
 * madana-api
 *
 * <h1>Using the madana-api</h1>        <p>This documentation contains a Quickstart Guide, relating client functionality and information about the available         endpoints and used datamodels.   </p>       <p> The madana-api and its implementations are still in heavy development. This means that there may be problems in our protocols, or there may be mistakes in our implementations. We take security vulnerabilities very seriously. If you discover a security issue, please bring it to our attention right away! If you find a vulnerability that may affect live deployments -- for example, by exposing a remote execution exploit -- please send your report privately to info@madana.io. Please DO NOT file a public issue. If the issue is a protocol weakness that cannot be immediately exploited or something not yet deployed, just discuss it openly   </p>   <br>   <p> Note: Not all functionality might be acessible without having accquired and api-license token. For more information visit <a href=\"https://www.madana.io\">www.madana.io</a> </p>       <br>
 *
 * API version: 0.5.0-master.52
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package madanaapiclient

import (
	"encoding/json"
)

// JsonEnclaveProcess 
type JsonEnclaveProcess struct {
	// 
	AttestationServer *string `json:"attestationServer,omitempty"`
	// 
	PortMapping *map[string]string `json:"portMapping,omitempty"`
	// 
	Ports *[]JsonEnclavePort `json:"ports,omitempty"`
	// 
	StartupTime *string `json:"startupTime,omitempty"`
	// 
	Status *string `json:"status,omitempty"`
	// 
	InternalIdent *string `json:"internalIdent,omitempty"`
	// 
	StartupCMD *string `json:"startupCMD,omitempty"`
	// 
	EnclaveIdent *string `json:"enclaveIdent,omitempty"`
	Environment *JsonEnvironment `json:"environment,omitempty"`
	// 
	PublicIdent *string `json:"publicIdent,omitempty"`
	// 
	ConsoleOutput *string `json:"consoleOutput,omitempty"`
	// 
	InternalWireguardServer *string `json:"internalWireguardServer,omitempty"`
	// 
	InternalAttesationServer *string `json:"internalAttesationServer,omitempty"`
	Process *JsonProcess `json:"process,omitempty"`
	// 
	InternalRemoteControlServer *string `json:"internalRemoteControlServer,omitempty"`
	// 
	SignerIdent *string `json:"signerIdent,omitempty"`
	WgInterface *JsonWireguardInterface `json:"wgInterface,omitempty"`
	// 
	WireguardServer *string `json:"wireguardServer,omitempty"`
	// 
	RemoteControlServer *string `json:"remoteControlServer,omitempty"`
	// 
	WireguardPublicKey *string `json:"wireguardPublicKey,omitempty"`
	// 
	EndingTime *string `json:"endingTime,omitempty"`
	KubernetesEnclave *JsonKubernetesEnclave `json:"kubernetesEnclave,omitempty"`
	// 
	EnclaveInputstream *map[string]interface{} `json:"enclaveInputstream,omitempty"`
}

// NewJsonEnclaveProcess instantiates a new JsonEnclaveProcess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJsonEnclaveProcess() *JsonEnclaveProcess {
	this := JsonEnclaveProcess{}
	return &this
}

// NewJsonEnclaveProcessWithDefaults instantiates a new JsonEnclaveProcess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJsonEnclaveProcessWithDefaults() *JsonEnclaveProcess {
	this := JsonEnclaveProcess{}
	return &this
}

// GetAttestationServer returns the AttestationServer field value if set, zero value otherwise.
func (o *JsonEnclaveProcess) GetAttestationServer() string {
	if o == nil || o.AttestationServer == nil {
		var ret string
		return ret
	}
	return *o.AttestationServer
}

// GetAttestationServerOk returns a tuple with the AttestationServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonEnclaveProcess) GetAttestationServerOk() (*string, bool) {
	if o == nil || o.AttestationServer == nil {
		return nil, false
	}
	return o.AttestationServer, true
}

// HasAttestationServer returns a boolean if a field has been set.
func (o *JsonEnclaveProcess) HasAttestationServer() bool {
	if o != nil && o.AttestationServer != nil {
		return true
	}

	return false
}

// SetAttestationServer gets a reference to the given string and assigns it to the AttestationServer field.
func (o *JsonEnclaveProcess) SetAttestationServer(v string) {
	o.AttestationServer = &v
}

// GetPortMapping returns the PortMapping field value if set, zero value otherwise.
func (o *JsonEnclaveProcess) GetPortMapping() map[string]string {
	if o == nil || o.PortMapping == nil {
		var ret map[string]string
		return ret
	}
	return *o.PortMapping
}

// GetPortMappingOk returns a tuple with the PortMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonEnclaveProcess) GetPortMappingOk() (*map[string]string, bool) {
	if o == nil || o.PortMapping == nil {
		return nil, false
	}
	return o.PortMapping, true
}

// HasPortMapping returns a boolean if a field has been set.
func (o *JsonEnclaveProcess) HasPortMapping() bool {
	if o != nil && o.PortMapping != nil {
		return true
	}

	return false
}

// SetPortMapping gets a reference to the given map[string]string and assigns it to the PortMapping field.
func (o *JsonEnclaveProcess) SetPortMapping(v map[string]string) {
	o.PortMapping = &v
}

// GetPorts returns the Ports field value if set, zero value otherwise.
func (o *JsonEnclaveProcess) GetPorts() []JsonEnclavePort {
	if o == nil || o.Ports == nil {
		var ret []JsonEnclavePort
		return ret
	}
	return *o.Ports
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonEnclaveProcess) GetPortsOk() (*[]JsonEnclavePort, bool) {
	if o == nil || o.Ports == nil {
		return nil, false
	}
	return o.Ports, true
}

// HasPorts returns a boolean if a field has been set.
func (o *JsonEnclaveProcess) HasPorts() bool {
	if o != nil && o.Ports != nil {
		return true
	}

	return false
}

// SetPorts gets a reference to the given []JsonEnclavePort and assigns it to the Ports field.
func (o *JsonEnclaveProcess) SetPorts(v []JsonEnclavePort) {
	o.Ports = &v
}

// GetStartupTime returns the StartupTime field value if set, zero value otherwise.
func (o *JsonEnclaveProcess) GetStartupTime() string {
	if o == nil || o.StartupTime == nil {
		var ret string
		return ret
	}
	return *o.StartupTime
}

// GetStartupTimeOk returns a tuple with the StartupTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonEnclaveProcess) GetStartupTimeOk() (*string, bool) {
	if o == nil || o.StartupTime == nil {
		return nil, false
	}
	return o.StartupTime, true
}

// HasStartupTime returns a boolean if a field has been set.
func (o *JsonEnclaveProcess) HasStartupTime() bool {
	if o != nil && o.StartupTime != nil {
		return true
	}

	return false
}

// SetStartupTime gets a reference to the given string and assigns it to the StartupTime field.
func (o *JsonEnclaveProcess) SetStartupTime(v string) {
	o.StartupTime = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *JsonEnclaveProcess) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonEnclaveProcess) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *JsonEnclaveProcess) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *JsonEnclaveProcess) SetStatus(v string) {
	o.Status = &v
}

// GetInternalIdent returns the InternalIdent field value if set, zero value otherwise.
func (o *JsonEnclaveProcess) GetInternalIdent() string {
	if o == nil || o.InternalIdent == nil {
		var ret string
		return ret
	}
	return *o.InternalIdent
}

// GetInternalIdentOk returns a tuple with the InternalIdent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonEnclaveProcess) GetInternalIdentOk() (*string, bool) {
	if o == nil || o.InternalIdent == nil {
		return nil, false
	}
	return o.InternalIdent, true
}

// HasInternalIdent returns a boolean if a field has been set.
func (o *JsonEnclaveProcess) HasInternalIdent() bool {
	if o != nil && o.InternalIdent != nil {
		return true
	}

	return false
}

// SetInternalIdent gets a reference to the given string and assigns it to the InternalIdent field.
func (o *JsonEnclaveProcess) SetInternalIdent(v string) {
	o.InternalIdent = &v
}

// GetStartupCMD returns the StartupCMD field value if set, zero value otherwise.
func (o *JsonEnclaveProcess) GetStartupCMD() string {
	if o == nil || o.StartupCMD == nil {
		var ret string
		return ret
	}
	return *o.StartupCMD
}

// GetStartupCMDOk returns a tuple with the StartupCMD field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonEnclaveProcess) GetStartupCMDOk() (*string, bool) {
	if o == nil || o.StartupCMD == nil {
		return nil, false
	}
	return o.StartupCMD, true
}

// HasStartupCMD returns a boolean if a field has been set.
func (o *JsonEnclaveProcess) HasStartupCMD() bool {
	if o != nil && o.StartupCMD != nil {
		return true
	}

	return false
}

// SetStartupCMD gets a reference to the given string and assigns it to the StartupCMD field.
func (o *JsonEnclaveProcess) SetStartupCMD(v string) {
	o.StartupCMD = &v
}

// GetEnclaveIdent returns the EnclaveIdent field value if set, zero value otherwise.
func (o *JsonEnclaveProcess) GetEnclaveIdent() string {
	if o == nil || o.EnclaveIdent == nil {
		var ret string
		return ret
	}
	return *o.EnclaveIdent
}

// GetEnclaveIdentOk returns a tuple with the EnclaveIdent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonEnclaveProcess) GetEnclaveIdentOk() (*string, bool) {
	if o == nil || o.EnclaveIdent == nil {
		return nil, false
	}
	return o.EnclaveIdent, true
}

// HasEnclaveIdent returns a boolean if a field has been set.
func (o *JsonEnclaveProcess) HasEnclaveIdent() bool {
	if o != nil && o.EnclaveIdent != nil {
		return true
	}

	return false
}

// SetEnclaveIdent gets a reference to the given string and assigns it to the EnclaveIdent field.
func (o *JsonEnclaveProcess) SetEnclaveIdent(v string) {
	o.EnclaveIdent = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *JsonEnclaveProcess) GetEnvironment() JsonEnvironment {
	if o == nil || o.Environment == nil {
		var ret JsonEnvironment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonEnclaveProcess) GetEnvironmentOk() (*JsonEnvironment, bool) {
	if o == nil || o.Environment == nil {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *JsonEnclaveProcess) HasEnvironment() bool {
	if o != nil && o.Environment != nil {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given JsonEnvironment and assigns it to the Environment field.
func (o *JsonEnclaveProcess) SetEnvironment(v JsonEnvironment) {
	o.Environment = &v
}

// GetPublicIdent returns the PublicIdent field value if set, zero value otherwise.
func (o *JsonEnclaveProcess) GetPublicIdent() string {
	if o == nil || o.PublicIdent == nil {
		var ret string
		return ret
	}
	return *o.PublicIdent
}

// GetPublicIdentOk returns a tuple with the PublicIdent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonEnclaveProcess) GetPublicIdentOk() (*string, bool) {
	if o == nil || o.PublicIdent == nil {
		return nil, false
	}
	return o.PublicIdent, true
}

// HasPublicIdent returns a boolean if a field has been set.
func (o *JsonEnclaveProcess) HasPublicIdent() bool {
	if o != nil && o.PublicIdent != nil {
		return true
	}

	return false
}

// SetPublicIdent gets a reference to the given string and assigns it to the PublicIdent field.
func (o *JsonEnclaveProcess) SetPublicIdent(v string) {
	o.PublicIdent = &v
}

// GetConsoleOutput returns the ConsoleOutput field value if set, zero value otherwise.
func (o *JsonEnclaveProcess) GetConsoleOutput() string {
	if o == nil || o.ConsoleOutput == nil {
		var ret string
		return ret
	}
	return *o.ConsoleOutput
}

// GetConsoleOutputOk returns a tuple with the ConsoleOutput field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonEnclaveProcess) GetConsoleOutputOk() (*string, bool) {
	if o == nil || o.ConsoleOutput == nil {
		return nil, false
	}
	return o.ConsoleOutput, true
}

// HasConsoleOutput returns a boolean if a field has been set.
func (o *JsonEnclaveProcess) HasConsoleOutput() bool {
	if o != nil && o.ConsoleOutput != nil {
		return true
	}

	return false
}

// SetConsoleOutput gets a reference to the given string and assigns it to the ConsoleOutput field.
func (o *JsonEnclaveProcess) SetConsoleOutput(v string) {
	o.ConsoleOutput = &v
}

// GetInternalWireguardServer returns the InternalWireguardServer field value if set, zero value otherwise.
func (o *JsonEnclaveProcess) GetInternalWireguardServer() string {
	if o == nil || o.InternalWireguardServer == nil {
		var ret string
		return ret
	}
	return *o.InternalWireguardServer
}

// GetInternalWireguardServerOk returns a tuple with the InternalWireguardServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonEnclaveProcess) GetInternalWireguardServerOk() (*string, bool) {
	if o == nil || o.InternalWireguardServer == nil {
		return nil, false
	}
	return o.InternalWireguardServer, true
}

// HasInternalWireguardServer returns a boolean if a field has been set.
func (o *JsonEnclaveProcess) HasInternalWireguardServer() bool {
	if o != nil && o.InternalWireguardServer != nil {
		return true
	}

	return false
}

// SetInternalWireguardServer gets a reference to the given string and assigns it to the InternalWireguardServer field.
func (o *JsonEnclaveProcess) SetInternalWireguardServer(v string) {
	o.InternalWireguardServer = &v
}

// GetInternalAttesationServer returns the InternalAttesationServer field value if set, zero value otherwise.
func (o *JsonEnclaveProcess) GetInternalAttesationServer() string {
	if o == nil || o.InternalAttesationServer == nil {
		var ret string
		return ret
	}
	return *o.InternalAttesationServer
}

// GetInternalAttesationServerOk returns a tuple with the InternalAttesationServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonEnclaveProcess) GetInternalAttesationServerOk() (*string, bool) {
	if o == nil || o.InternalAttesationServer == nil {
		return nil, false
	}
	return o.InternalAttesationServer, true
}

// HasInternalAttesationServer returns a boolean if a field has been set.
func (o *JsonEnclaveProcess) HasInternalAttesationServer() bool {
	if o != nil && o.InternalAttesationServer != nil {
		return true
	}

	return false
}

// SetInternalAttesationServer gets a reference to the given string and assigns it to the InternalAttesationServer field.
func (o *JsonEnclaveProcess) SetInternalAttesationServer(v string) {
	o.InternalAttesationServer = &v
}

// GetProcess returns the Process field value if set, zero value otherwise.
func (o *JsonEnclaveProcess) GetProcess() JsonProcess {
	if o == nil || o.Process == nil {
		var ret JsonProcess
		return ret
	}
	return *o.Process
}

// GetProcessOk returns a tuple with the Process field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonEnclaveProcess) GetProcessOk() (*JsonProcess, bool) {
	if o == nil || o.Process == nil {
		return nil, false
	}
	return o.Process, true
}

// HasProcess returns a boolean if a field has been set.
func (o *JsonEnclaveProcess) HasProcess() bool {
	if o != nil && o.Process != nil {
		return true
	}

	return false
}

// SetProcess gets a reference to the given JsonProcess and assigns it to the Process field.
func (o *JsonEnclaveProcess) SetProcess(v JsonProcess) {
	o.Process = &v
}

// GetInternalRemoteControlServer returns the InternalRemoteControlServer field value if set, zero value otherwise.
func (o *JsonEnclaveProcess) GetInternalRemoteControlServer() string {
	if o == nil || o.InternalRemoteControlServer == nil {
		var ret string
		return ret
	}
	return *o.InternalRemoteControlServer
}

// GetInternalRemoteControlServerOk returns a tuple with the InternalRemoteControlServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonEnclaveProcess) GetInternalRemoteControlServerOk() (*string, bool) {
	if o == nil || o.InternalRemoteControlServer == nil {
		return nil, false
	}
	return o.InternalRemoteControlServer, true
}

// HasInternalRemoteControlServer returns a boolean if a field has been set.
func (o *JsonEnclaveProcess) HasInternalRemoteControlServer() bool {
	if o != nil && o.InternalRemoteControlServer != nil {
		return true
	}

	return false
}

// SetInternalRemoteControlServer gets a reference to the given string and assigns it to the InternalRemoteControlServer field.
func (o *JsonEnclaveProcess) SetInternalRemoteControlServer(v string) {
	o.InternalRemoteControlServer = &v
}

// GetSignerIdent returns the SignerIdent field value if set, zero value otherwise.
func (o *JsonEnclaveProcess) GetSignerIdent() string {
	if o == nil || o.SignerIdent == nil {
		var ret string
		return ret
	}
	return *o.SignerIdent
}

// GetSignerIdentOk returns a tuple with the SignerIdent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonEnclaveProcess) GetSignerIdentOk() (*string, bool) {
	if o == nil || o.SignerIdent == nil {
		return nil, false
	}
	return o.SignerIdent, true
}

// HasSignerIdent returns a boolean if a field has been set.
func (o *JsonEnclaveProcess) HasSignerIdent() bool {
	if o != nil && o.SignerIdent != nil {
		return true
	}

	return false
}

// SetSignerIdent gets a reference to the given string and assigns it to the SignerIdent field.
func (o *JsonEnclaveProcess) SetSignerIdent(v string) {
	o.SignerIdent = &v
}

// GetWgInterface returns the WgInterface field value if set, zero value otherwise.
func (o *JsonEnclaveProcess) GetWgInterface() JsonWireguardInterface {
	if o == nil || o.WgInterface == nil {
		var ret JsonWireguardInterface
		return ret
	}
	return *o.WgInterface
}

// GetWgInterfaceOk returns a tuple with the WgInterface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonEnclaveProcess) GetWgInterfaceOk() (*JsonWireguardInterface, bool) {
	if o == nil || o.WgInterface == nil {
		return nil, false
	}
	return o.WgInterface, true
}

// HasWgInterface returns a boolean if a field has been set.
func (o *JsonEnclaveProcess) HasWgInterface() bool {
	if o != nil && o.WgInterface != nil {
		return true
	}

	return false
}

// SetWgInterface gets a reference to the given JsonWireguardInterface and assigns it to the WgInterface field.
func (o *JsonEnclaveProcess) SetWgInterface(v JsonWireguardInterface) {
	o.WgInterface = &v
}

// GetWireguardServer returns the WireguardServer field value if set, zero value otherwise.
func (o *JsonEnclaveProcess) GetWireguardServer() string {
	if o == nil || o.WireguardServer == nil {
		var ret string
		return ret
	}
	return *o.WireguardServer
}

// GetWireguardServerOk returns a tuple with the WireguardServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonEnclaveProcess) GetWireguardServerOk() (*string, bool) {
	if o == nil || o.WireguardServer == nil {
		return nil, false
	}
	return o.WireguardServer, true
}

// HasWireguardServer returns a boolean if a field has been set.
func (o *JsonEnclaveProcess) HasWireguardServer() bool {
	if o != nil && o.WireguardServer != nil {
		return true
	}

	return false
}

// SetWireguardServer gets a reference to the given string and assigns it to the WireguardServer field.
func (o *JsonEnclaveProcess) SetWireguardServer(v string) {
	o.WireguardServer = &v
}

// GetRemoteControlServer returns the RemoteControlServer field value if set, zero value otherwise.
func (o *JsonEnclaveProcess) GetRemoteControlServer() string {
	if o == nil || o.RemoteControlServer == nil {
		var ret string
		return ret
	}
	return *o.RemoteControlServer
}

// GetRemoteControlServerOk returns a tuple with the RemoteControlServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonEnclaveProcess) GetRemoteControlServerOk() (*string, bool) {
	if o == nil || o.RemoteControlServer == nil {
		return nil, false
	}
	return o.RemoteControlServer, true
}

// HasRemoteControlServer returns a boolean if a field has been set.
func (o *JsonEnclaveProcess) HasRemoteControlServer() bool {
	if o != nil && o.RemoteControlServer != nil {
		return true
	}

	return false
}

// SetRemoteControlServer gets a reference to the given string and assigns it to the RemoteControlServer field.
func (o *JsonEnclaveProcess) SetRemoteControlServer(v string) {
	o.RemoteControlServer = &v
}

// GetWireguardPublicKey returns the WireguardPublicKey field value if set, zero value otherwise.
func (o *JsonEnclaveProcess) GetWireguardPublicKey() string {
	if o == nil || o.WireguardPublicKey == nil {
		var ret string
		return ret
	}
	return *o.WireguardPublicKey
}

// GetWireguardPublicKeyOk returns a tuple with the WireguardPublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonEnclaveProcess) GetWireguardPublicKeyOk() (*string, bool) {
	if o == nil || o.WireguardPublicKey == nil {
		return nil, false
	}
	return o.WireguardPublicKey, true
}

// HasWireguardPublicKey returns a boolean if a field has been set.
func (o *JsonEnclaveProcess) HasWireguardPublicKey() bool {
	if o != nil && o.WireguardPublicKey != nil {
		return true
	}

	return false
}

// SetWireguardPublicKey gets a reference to the given string and assigns it to the WireguardPublicKey field.
func (o *JsonEnclaveProcess) SetWireguardPublicKey(v string) {
	o.WireguardPublicKey = &v
}

// GetEndingTime returns the EndingTime field value if set, zero value otherwise.
func (o *JsonEnclaveProcess) GetEndingTime() string {
	if o == nil || o.EndingTime == nil {
		var ret string
		return ret
	}
	return *o.EndingTime
}

// GetEndingTimeOk returns a tuple with the EndingTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonEnclaveProcess) GetEndingTimeOk() (*string, bool) {
	if o == nil || o.EndingTime == nil {
		return nil, false
	}
	return o.EndingTime, true
}

// HasEndingTime returns a boolean if a field has been set.
func (o *JsonEnclaveProcess) HasEndingTime() bool {
	if o != nil && o.EndingTime != nil {
		return true
	}

	return false
}

// SetEndingTime gets a reference to the given string and assigns it to the EndingTime field.
func (o *JsonEnclaveProcess) SetEndingTime(v string) {
	o.EndingTime = &v
}

// GetKubernetesEnclave returns the KubernetesEnclave field value if set, zero value otherwise.
func (o *JsonEnclaveProcess) GetKubernetesEnclave() JsonKubernetesEnclave {
	if o == nil || o.KubernetesEnclave == nil {
		var ret JsonKubernetesEnclave
		return ret
	}
	return *o.KubernetesEnclave
}

// GetKubernetesEnclaveOk returns a tuple with the KubernetesEnclave field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonEnclaveProcess) GetKubernetesEnclaveOk() (*JsonKubernetesEnclave, bool) {
	if o == nil || o.KubernetesEnclave == nil {
		return nil, false
	}
	return o.KubernetesEnclave, true
}

// HasKubernetesEnclave returns a boolean if a field has been set.
func (o *JsonEnclaveProcess) HasKubernetesEnclave() bool {
	if o != nil && o.KubernetesEnclave != nil {
		return true
	}

	return false
}

// SetKubernetesEnclave gets a reference to the given JsonKubernetesEnclave and assigns it to the KubernetesEnclave field.
func (o *JsonEnclaveProcess) SetKubernetesEnclave(v JsonKubernetesEnclave) {
	o.KubernetesEnclave = &v
}

// GetEnclaveInputstream returns the EnclaveInputstream field value if set, zero value otherwise.
func (o *JsonEnclaveProcess) GetEnclaveInputstream() map[string]interface{} {
	if o == nil || o.EnclaveInputstream == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.EnclaveInputstream
}

// GetEnclaveInputstreamOk returns a tuple with the EnclaveInputstream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonEnclaveProcess) GetEnclaveInputstreamOk() (*map[string]interface{}, bool) {
	if o == nil || o.EnclaveInputstream == nil {
		return nil, false
	}
	return o.EnclaveInputstream, true
}

// HasEnclaveInputstream returns a boolean if a field has been set.
func (o *JsonEnclaveProcess) HasEnclaveInputstream() bool {
	if o != nil && o.EnclaveInputstream != nil {
		return true
	}

	return false
}

// SetEnclaveInputstream gets a reference to the given map[string]interface{} and assigns it to the EnclaveInputstream field.
func (o *JsonEnclaveProcess) SetEnclaveInputstream(v map[string]interface{}) {
	o.EnclaveInputstream = &v
}

func (o JsonEnclaveProcess) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AttestationServer != nil {
		toSerialize["attestationServer"] = o.AttestationServer
	}
	if o.PortMapping != nil {
		toSerialize["portMapping"] = o.PortMapping
	}
	if o.Ports != nil {
		toSerialize["ports"] = o.Ports
	}
	if o.StartupTime != nil {
		toSerialize["startupTime"] = o.StartupTime
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.InternalIdent != nil {
		toSerialize["internalIdent"] = o.InternalIdent
	}
	if o.StartupCMD != nil {
		toSerialize["startupCMD"] = o.StartupCMD
	}
	if o.EnclaveIdent != nil {
		toSerialize["enclaveIdent"] = o.EnclaveIdent
	}
	if o.Environment != nil {
		toSerialize["environment"] = o.Environment
	}
	if o.PublicIdent != nil {
		toSerialize["publicIdent"] = o.PublicIdent
	}
	if o.ConsoleOutput != nil {
		toSerialize["consoleOutput"] = o.ConsoleOutput
	}
	if o.InternalWireguardServer != nil {
		toSerialize["internalWireguardServer"] = o.InternalWireguardServer
	}
	if o.InternalAttesationServer != nil {
		toSerialize["internalAttesationServer"] = o.InternalAttesationServer
	}
	if o.Process != nil {
		toSerialize["process"] = o.Process
	}
	if o.InternalRemoteControlServer != nil {
		toSerialize["internalRemoteControlServer"] = o.InternalRemoteControlServer
	}
	if o.SignerIdent != nil {
		toSerialize["signerIdent"] = o.SignerIdent
	}
	if o.WgInterface != nil {
		toSerialize["wgInterface"] = o.WgInterface
	}
	if o.WireguardServer != nil {
		toSerialize["wireguardServer"] = o.WireguardServer
	}
	if o.RemoteControlServer != nil {
		toSerialize["remoteControlServer"] = o.RemoteControlServer
	}
	if o.WireguardPublicKey != nil {
		toSerialize["wireguardPublicKey"] = o.WireguardPublicKey
	}
	if o.EndingTime != nil {
		toSerialize["endingTime"] = o.EndingTime
	}
	if o.KubernetesEnclave != nil {
		toSerialize["kubernetesEnclave"] = o.KubernetesEnclave
	}
	if o.EnclaveInputstream != nil {
		toSerialize["enclaveInputstream"] = o.EnclaveInputstream
	}
	return json.Marshal(toSerialize)
}

type NullableJsonEnclaveProcess struct {
	value *JsonEnclaveProcess
	isSet bool
}

func (v NullableJsonEnclaveProcess) Get() *JsonEnclaveProcess {
	return v.value
}

func (v *NullableJsonEnclaveProcess) Set(val *JsonEnclaveProcess) {
	v.value = val
	v.isSet = true
}

func (v NullableJsonEnclaveProcess) IsSet() bool {
	return v.isSet
}

func (v *NullableJsonEnclaveProcess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJsonEnclaveProcess(val *JsonEnclaveProcess) *NullableJsonEnclaveProcess {
	return &NullableJsonEnclaveProcess{value: val, isSet: true}
}

func (v NullableJsonEnclaveProcess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJsonEnclaveProcess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


