/*
 * madana-api
 *
 * <h1>Using the madana-api</h1>        <p>This documentation contains a Quickstart Guide, relating client functionality and information about the available         endpoints and used datamodels.   </p>       <p> The madana-api and its implementations are still in heavy development. This means that there may be problems in our protocols, or there may be mistakes in our implementations. We take security vulnerabilities very seriously. If you discover a security issue, please bring it to our attention right away! If you find a vulnerability that may affect live deployments -- for example, by exposing a remote execution exploit -- please send your report privately to info@madana.io. Please DO NOT file a public issue. If the issue is a protocol weakness that cannot be immediately exploited or something not yet deployed, just discuss it openly   </p>   <br>   <p> Note: Not all functionality might be acessible without having accquired and api-license token. For more information visit <a href=\"https://www.madana.io\">www.madana.io</a> </p>       <br>
 *
 * API version: 0.4.14-master.18
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package madanaapiclient
// XmlNs0EnclaveProcess 
type XmlNs0EnclaveProcess struct {
	// 
	AttestationServer string `json:"attestationServer,omitempty"`
	// 
	ConsoleOutput string `json:"consoleOutput,omitempty"`
	// 
	EnclaveIdent string `json:"enclaveIdent,omitempty"`
	EnclaveInputstream XmlNs0InputStream `json:"enclaveInputstream,omitempty"`
	// 
	EndingTime string `json:"endingTime,omitempty"`
	Environment XmlNs0Environment `json:"environment,omitempty"`
	// 
	InternalAttesationServer string `json:"internalAttesationServer,omitempty"`
	// 
	InternalIdent string `json:"internalIdent,omitempty"`
	// 
	InternalRemoteControlServer string `json:"internalRemoteControlServer,omitempty"`
	Process XmlNs0Process `json:"process,omitempty"`
	// 
	RemoteControlServer string `json:"remoteControlServer,omitempty"`
	// 
	SignerIdent string `json:"signerIdent,omitempty"`
	// 
	StartupCMD string `json:"startupCMD,omitempty"`
	// 
	StartupTime string `json:"startupTime,omitempty"`
	// 
	Status string `json:"status,omitempty"`
	WgInterface XmlNs0WireguardInterface `json:"wgInterface,omitempty"`
	// 
	WireguardPublicKey string `json:"wireguardPublicKey,omitempty"`
}
