# Go API client for madanaapiclient

<h1>Using the madana-api</h1>
       <p>This documentation contains a Quickstart Guide, relating client functionality and information about the available 
       endpoints and used datamodels.   </p>    

 <p> The madana-api and its implementations are still in heavy development. This means that there may be problems in our protocols, or there may be mistakes in our implementations. We take security vulnerabilities very seriously. If you discover a security issue, please bring it to our attention right away! If you find a vulnerability that may affect live deployments -- for example, by exposing a remote execution exploit -- please send your report privately to info@madana.io. Please DO NOT file a public issue. If the issue is a protocol weakness that cannot be immediately exploited or something not yet deployed, just discuss it openly   </p> 
 <br>
  <p> Note: Not all functionality might be acessible without having accquired and api-license token. For more information visit <a href=\"https://www.madana.io\">www.madana.io</a> </p>    
  <br>

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 0.5.0-master.56
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import sw "./madanaapiclient"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identifield by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), sw.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), sw.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://api.madana.io/rest*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AccountServiceApi* | [**ActivateUser**](docs/AccountServiceApi.md#activateuser) | **Get** /account/activation/{token} | 
*AccountServiceApi* | [**CreatePasswordReset**](docs/AccountServiceApi.md#createpasswordreset) | **Post** /account/password | Sends an Password reset mail to the given MailAddress.
*AccountServiceApi* | [**RequestVerificationMail**](docs/AccountServiceApi.md#requestverificationmail) | **Get** /account/verifymail | Used to request a new  activation-mail for the user.
*AccountServiceApi* | [**UpdatePassword**](docs/AccountServiceApi.md#updatepassword) | **Put** /account/password | Receives the Password reset and tries to set the provided password for the user.
*AuthenticationServiceApi* | [**AuthenticateApplication**](docs/AuthenticationServiceApi.md#authenticateapplication) | **Post** /authentication/application | Authenticates a new application and returns the token.
*AuthenticationServiceApi* | [**AuthenticateEthereumWallet**](docs/AuthenticationServiceApi.md#authenticateethereumwallet) | **Post** /authentication/ethereum/{wallet} | 
*AuthenticationServiceApi* | [**AuthenticateUser**](docs/AuthenticationServiceApi.md#authenticateuser) | **Post** /authentication | Authenticates a new user and returns the token (  forbidden if the credentials cannot be validated ).
*AuthenticationServiceApi* | [**AuthenticateWithEthereumChallenge**](docs/AuthenticationServiceApi.md#authenticatewithethereumchallenge) | **Post** /authentication/ethereum/{wallet}/challenge | 
*AuthenticationServiceApi* | [**GetFractalAuthenticationURL**](docs/AuthenticationServiceApi.md#getfractalauthenticationurl) | **Get** /authentication/fractal | Returns the AUthorization URL to verify a Twitter Accounts.
*AuthenticationServiceApi* | [**GetNonceForEthereumWallet**](docs/AuthenticationServiceApi.md#getnonceforethereumwallet) | **Get** /authentication/ethereum/{wallet} | Returns a nonce for the client which is used as content for the to be created signature.
*AuthenticationServiceApi* | [**GetObject**](docs/AuthenticationServiceApi.md#getobject) | **Get** /authentication | Used to validate the active connection with the API.
*AuthenticationServiceApi* | [**GetTwitterAuthenticationURL**](docs/AuthenticationServiceApi.md#gettwitterauthenticationurl) | **Get** /authentication/twitter | Returns the AUthorization URL to verify a Twitter Accounts.
*AuthenticationServiceApi* | [**SetFacebookUID**](docs/AuthenticationServiceApi.md#setfacebookuid) | **Post** /authentication/facebook | Used as Callback URL when users have successfully authorized their facbeook account.
*AuthenticationServiceApi* | [**SetFractalUID**](docs/AuthenticationServiceApi.md#setfractaluid) | **Post** /authentication/fractal | 
*AuthenticationServiceApi* | [**SetTwitterUID**](docs/AuthenticationServiceApi.md#settwitteruid) | **Post** /authentication/twitter | 
*CertificateServiceApi* | [**AuthenticateCertificate**](docs/CertificateServiceApi.md#authenticatecertificate) | **Post** /certificates | Issues certificates for logged-in users.
*CertificateServiceApi* | [**GetCertificateByFingerprint**](docs/CertificateServiceApi.md#getcertificatebyfingerprint) | **Get** /certificates/{fingerprint} | 
*CertificateServiceApi* | [**GetRootCertificate**](docs/CertificateServiceApi.md#getrootcertificate) | **Get** /certificates/root | 
*DataCollectionServiceApi* | [**GetMethodsForType**](docs/DataCollectionServiceApi.md#getmethodsfortype) | **Get** /datacollection/types/{name}/methods | 
*DataCollectionServiceApi* | [**GetNodes**](docs/DataCollectionServiceApi.md#getnodes) | **Get** /datacollection/methods | 
*DataCollectionServiceApi* | [**GetTypes**](docs/DataCollectionServiceApi.md#gettypes) | **Get** /datacollection/types | 
*EnclaveServiceApi* | [**AddHistory**](docs/EnclaveServiceApi.md#addhistory) | **Post** /enclaves/{uuid}/history | 
*EnclaveServiceApi* | [**ApproveEnclave**](docs/EnclaveServiceApi.md#approveenclave) | **Post** /enclaves/{uuid}/approval | 
*EnclaveServiceApi* | [**AssignEnclaveAgent**](docs/EnclaveServiceApi.md#assignenclaveagent) | **Post** /enclaves/{uuid}/assign | 
*EnclaveServiceApi* | [**AttestateEnclave**](docs/EnclaveServiceApi.md#attestateenclave) | **Post** /enclaves/{uuid}/attestation | 
*EnclaveServiceApi* | [**CreateEnclaveRunRequest**](docs/EnclaveServiceApi.md#createenclaverunrequest) | **Post** /enclaves | 
*EnclaveServiceApi* | [**GetEnclave**](docs/EnclaveServiceApi.md#getenclave) | **Get** /enclaves/{uuid} | 
*EnclaveServiceApi* | [**GetEnclaveTypes**](docs/EnclaveServiceApi.md#getenclavetypes) | **Get** /enclaves/types | 
*EnclaveServiceApi* | [**GetEnclaves**](docs/EnclaveServiceApi.md#getenclaves) | **Get** /enclaves | Returns UUIDs of existing analyses.
*EnclaveServiceApi* | [**GetStats**](docs/EnclaveServiceApi.md#getstats) | **Get** /enclaves/stats | 
*EnclaveServiceApi* | [**KillEnclave**](docs/EnclaveServiceApi.md#killenclave) | **Post** /enclaves/{uuid}/kill | 
*EnvironmentServiceApi* | [**DeleteEnvironment**](docs/EnvironmentServiceApi.md#deleteenvironment) | **Delete** /environments/{uuid} | 
*EnvironmentServiceApi* | [**DeleteEnvironmentSubscription**](docs/EnvironmentServiceApi.md#deleteenvironmentsubscription) | **Delete** /environments/{uuid}/subscribe | 
*EnvironmentServiceApi* | [**GetEnvironment**](docs/EnvironmentServiceApi.md#getenvironment) | **Get** /environments/{uuid} | 
*EnvironmentServiceApi* | [**GetEnvironments**](docs/EnvironmentServiceApi.md#getenvironments) | **Get** /environments | Returns UUIDs of existing analyses.
*EnvironmentServiceApi* | [**GetPublishedEnvironments**](docs/EnvironmentServiceApi.md#getpublishedenvironments) | **Get** /environments/published | 
*EnvironmentServiceApi* | [**GetSubscribedEnvironments**](docs/EnvironmentServiceApi.md#getsubscribedenvironments) | **Get** /environments/subscriptions | 
*EnvironmentServiceApi* | [**PublishEnvironment**](docs/EnvironmentServiceApi.md#publishenvironment) | **Post** /environments | 
*EnvironmentServiceApi* | [**SubscribeEnvironment**](docs/EnvironmentServiceApi.md#subscribeenvironment) | **Post** /environments/{uuid}/subscribe | 
*EnvironmentServiceApi* | [**UpdateEnvironment**](docs/EnvironmentServiceApi.md#updateenvironment) | **Put** /environments/{uuid} | 
*InvoiceServiceApi* | [**GetBillingPortalURL**](docs/InvoiceServiceApi.md#getbillingportalurl) | **Get** /invoices/portal | 
*InvoiceServiceApi* | [**GetInvoices**](docs/InvoiceServiceApi.md#getinvoices) | **Get** /invoices | 
*NodeServiceApi* | [**CreateNode**](docs/NodeServiceApi.md#createnode) | **Post** /nodes/v2 | 
*NodeServiceApi* | [**GetBootstrap**](docs/NodeServiceApi.md#getbootstrap) | **Get** /nodes/bootstrap | 
*NodeServiceApi* | [**GetNodeLicenses**](docs/NodeServiceApi.md#getnodelicenses) | **Get** /nodes/licenses | 
*NodeServiceApi* | [**GetNodeV2**](docs/NodeServiceApi.md#getnodev2) | **Get** /nodes/v2/{ident} | 
*NodeServiceApi* | [**GetNodes2**](docs/NodeServiceApi.md#getnodes2) | **Get** /nodes | 
*NodeServiceApi* | [**GetNodesV2**](docs/NodeServiceApi.md#getnodesv2) | **Get** /nodes/v2 | Returns UUIDs of existing analyses.
*NodeServiceApi* | [**KillNode**](docs/NodeServiceApi.md#killnode) | **Post** /nodes/v2/{ident}/kill | 
*NodeServiceApi* | [**PostNodeInfo**](docs/NodeServiceApi.md#postnodeinfo) | **Post** /nodes | 
*NodeServiceApi* | [**PostNodeInfo_0**](docs/NodeServiceApi.md#postnodeinfo_0) | **Post** /nodes/create | 
*OrganizationServiceApi* | [**GetNodes3**](docs/OrganizationServiceApi.md#getnodes3) | **Get** /organizations | 
*RequestServiceApi* | [**AddData**](docs/RequestServiceApi.md#adddata) | **Post** /requests/{uuid}/data | Is used to upload and park the data till the AnalysisRequest gets processed.
*RequestServiceApi* | [**CancelProcessing**](docs/RequestServiceApi.md#cancelprocessing) | **Post** /requests/{uuid}/cancel | Endpoint is called from the Analysis Processing entity to submit the result.
*RequestServiceApi* | [**CreateNewRequest**](docs/RequestServiceApi.md#createnewrequest) | **Post** /requests | Endpoint used to create a new Analysis Request.
*RequestServiceApi* | [**GetActions**](docs/RequestServiceApi.md#getactions) | **Get** /requests/actions | 
*RequestServiceApi* | [**GetAgent**](docs/RequestServiceApi.md#getagent) | **Get** /requests/{uuid}/agent | Is called from the APE to request all parked datasets.
*RequestServiceApi* | [**GetAllRequests**](docs/RequestServiceApi.md#getallrequests) | **Get** /requests | Returns UUIDs of existing analyses.
*RequestServiceApi* | [**GetData**](docs/RequestServiceApi.md#getdata) | **Get** /requests/{uuid}/data | Is called from the APE to request all parked datasets.
*RequestServiceApi* | [**GetRequest**](docs/RequestServiceApi.md#getrequest) | **Get** /requests/{uuid} | Returns the details for certain Request.
*RequestServiceApi* | [**GetResult**](docs/RequestServiceApi.md#getresult) | **Get** /requests/{uuid}/result | Can be called from creator to request the AnalysisResult.
*RequestServiceApi* | [**GetStatus**](docs/RequestServiceApi.md#getstatus) | **Get** /requests/stats | 
*RequestServiceApi* | [**GiveConsent**](docs/RequestServiceApi.md#giveconsent) | **Post** /requests/{uuid}/consent | Used to give consent for request.
*RequestServiceApi* | [**InitRequestParameters**](docs/RequestServiceApi.md#initrequestparameters) | **Post** /requests/{uuid} | Endpoint used initialized addition datacollection parameters for requester.
*RequestServiceApi* | [**SetAgent**](docs/RequestServiceApi.md#setagent) | **Post** /requests/{uuid}/agent | Is called from the APE to request all parked datasets.
*RequestServiceApi* | [**SetResult**](docs/RequestServiceApi.md#setresult) | **Post** /requests/{uuid}/result | Endpoint is called from the Analysis Processing entity to submit the result.
*SocialPlatformServiceApi* | [**GetPlatforms**](docs/SocialPlatformServiceApi.md#getplatforms) | **Get** /platforms | Used to Handle Incoming Webhooks from Facebook.
*SocialPlatformServiceApi* | [**ListenTwitterWebhook**](docs/SocialPlatformServiceApi.md#listentwitterwebhook) | **Post** /platforms/twitter | Used to Handle Incoming Webhooks from Facebook.
*SocialPlatformServiceApi* | [**RegisterTwitterWebhook**](docs/SocialPlatformServiceApi.md#registertwitterwebhook) | **Get** /platforms/twitter | Used to Handle Incoming Webhooks from Twitter.
*SocialServiceApi* | [**GetMyProfile**](docs/SocialServiceApi.md#getmyprofile) | **Get** /social/profiles/me | 
*SocialServiceApi* | [**GetPlatforms2**](docs/SocialServiceApi.md#getplatforms2) | **Get** /social | Returns all Platforms / Systems that can be Connected to the MADANA Service.
*SocialServiceApi* | [**GetRanking**](docs/SocialServiceApi.md#getranking) | **Get** /social/ranking | Returns the Ranking by PTS within the System.
*SocialServiceApi* | [**GetSocialPlatformFeed**](docs/SocialServiceApi.md#getsocialplatformfeed) | **Get** /social/feed/{platform} | 
*SocialServiceApi* | [**GetUserProfile**](docs/SocialServiceApi.md#getuserprofile) | **Get** /social/profiles/{username} | 
*SocialServiceApi* | [**GetUserProfile_0**](docs/SocialServiceApi.md#getuserprofile_0) | **Get** /social/profiles/{username}/simple | 
*SubscriptionServiceApi* | [**AddFreeSubscription**](docs/SubscriptionServiceApi.md#addfreesubscription) | **Post** /subscriptions/saas/free | 
*SubscriptionServiceApi* | [**AddPassTrialSubscription**](docs/SubscriptionServiceApi.md#addpasstrialsubscription) | **Post** /subscriptions/paas/trial | 
*SubscriptionServiceApi* | [**GetApplication**](docs/SubscriptionServiceApi.md#getapplication) | **Get** /subscriptions/active | 
*SubscriptionServiceApi* | [**GetCheckoutSession**](docs/SubscriptionServiceApi.md#getcheckoutsession) | **Get** /subscriptions/{productname}/checkout | 
*SubscriptionServiceApi* | [**GetCheckoutSession2**](docs/SubscriptionServiceApi.md#getcheckoutsession2) | **Post** /subscriptions/{productname}/{newplan} | 
*SystemServiceApi* | [**GetAllObjects**](docs/SystemServiceApi.md#getallobjects) | **Get** /system/health | 
*SystemServiceApi* | [**GetApplication2**](docs/SystemServiceApi.md#getapplication2) | **Get** /system/usage | Return the current application usage.
*UserServiceApi* | [**CancelSubscription**](docs/UserServiceApi.md#cancelsubscription) | **Post** /users/{username}/subscriptions/{planname}/cancel | 
*UserServiceApi* | [**CreateObject**](docs/UserServiceApi.md#createobject) | **Post** /users | Creates a new user object.
*UserServiceApi* | [**DeleteObject**](docs/UserServiceApi.md#deleteobject) | **Delete** /users/{username} | Deletes an User based on the provided id and securitycontext.
*UserServiceApi* | [**DeleteObject_0**](docs/UserServiceApi.md#deleteobject_0) | **Delete** /users/{username}/social/{platform}/{ident} | Deletes linked account from the user and securitycontext.
*UserServiceApi* | [**GetAvatars**](docs/UserServiceApi.md#getavatars) | **Get** /users/{username}/avatars | 
*UserServiceApi* | [**GetCertificates**](docs/UserServiceApi.md#getcertificates) | **Get** /users/{username}/certificates | 
*UserServiceApi* | [**GetEnclaveHistory**](docs/UserServiceApi.md#getenclavehistory) | **Get** /users/{username}/enclavehistory | 
*UserServiceApi* | [**GetObject2**](docs/UserServiceApi.md#getobject2) | **Get** /users/{username} | 
*UserServiceApi* | [**SetAvatar**](docs/UserServiceApi.md#setavatar) | **Post** /users/{username}/avatars | 
*UserServiceApi* | [**SetSettings**](docs/UserServiceApi.md#setsettings) | **Post** /users/{username}/settings | 
*UserServiceApi* | [**UpdateObject**](docs/UserServiceApi.md#updateobject) | **Put** /users/{username} | Updates Userproperties based on the provided user object.


## Documentation For Models

 - [JsonDiskConfig](docs/JsonDiskConfig.md)
 - [JsonEnclavePort](docs/JsonEnclavePort.md)
 - [JsonEnclaveProcess](docs/JsonEnclaveProcess.md)
 - [JsonEnclaveRunRequest](docs/JsonEnclaveRunRequest.md)
 - [JsonEnclaveRunningAttestation](docs/JsonEnclaveRunningAttestation.md)
 - [JsonEnclaveRunningAttestationApproval](docs/JsonEnclaveRunningAttestationApproval.md)
 - [JsonEnclaveRunningAttestationApprovalAllOf](docs/JsonEnclaveRunningAttestationApprovalAllOf.md)
 - [JsonEnvironment](docs/JsonEnvironment.md)
 - [JsonEnvironmentPublishingRequest](docs/JsonEnvironmentPublishingRequest.md)
 - [JsonIPFSSystemInfo](docs/JsonIPFSSystemInfo.md)
 - [JsonKubernetesEnclave](docs/JsonKubernetesEnclave.md)
 - [JsonKubernetesEnclaveAllOf](docs/JsonKubernetesEnclaveAllOf.md)
 - [JsonMDNAUserObject](docs/JsonMDNAUserObject.md)
 - [JsonMDNCertificate](docs/JsonMDNCertificate.md)
 - [JsonMDNData](docs/JsonMDNData.md)
 - [JsonMDNMailAddress](docs/JsonMDNMailAddress.md)
 - [JsonMDNOAuthToken](docs/JsonMDNOAuthToken.md)
 - [JsonMDNPasswordReset](docs/JsonMDNPasswordReset.md)
 - [JsonMDNSetting](docs/JsonMDNSetting.md)
 - [JsonMDNSocialUserObject](docs/JsonMDNSocialUserObject.md)
 - [JsonMDNToken](docs/JsonMDNToken.md)
 - [JsonMDNUser](docs/JsonMDNUser.md)
 - [JsonMDNUserAllOf](docs/JsonMDNUserAllOf.md)
 - [JsonMDNUserCredentials](docs/JsonMDNUserCredentials.md)
 - [JsonMDNUserProfileImage](docs/JsonMDNUserProfileImage.md)
 - [JsonMDNUserSetting](docs/JsonMDNUserSetting.md)
 - [JsonMDNUserSettingAllOf](docs/JsonMDNUserSettingAllOf.md)
 - [JsonNetworkInterface](docs/JsonNetworkInterface.md)
 - [JsonNodeInfo](docs/JsonNodeInfo.md)
 - [JsonNodeRunRequest](docs/JsonNodeRunRequest.md)
 - [JsonProcess](docs/JsonProcess.md)
 - [JsonRunConfig](docs/JsonRunConfig.md)
 - [JsonSGXInfo](docs/JsonSGXInfo.md)
 - [JsonSignedData](docs/JsonSignedData.md)
 - [JsonV1Event](docs/JsonV1Event.md)
 - [JsonV1EventList](docs/JsonV1EventList.md)
 - [JsonV1EventSeries](docs/JsonV1EventSeries.md)
 - [JsonV1EventSource](docs/JsonV1EventSource.md)
 - [JsonV1ListMeta](docs/JsonV1ListMeta.md)
 - [JsonV1ManagedFieldsEntry](docs/JsonV1ManagedFieldsEntry.md)
 - [JsonV1ObjectMeta](docs/JsonV1ObjectMeta.md)
 - [JsonV1ObjectReference](docs/JsonV1ObjectReference.md)
 - [JsonV1OwnerReference](docs/JsonV1OwnerReference.md)
 - [JsonWireguardInterface](docs/JsonWireguardInterface.md)
 - [JsonWireguardInterfaceAllOf](docs/JsonWireguardInterfaceAllOf.md)
 - [XmlNs0DiskConfig](docs/XmlNs0DiskConfig.md)
 - [XmlNs0DiskConfigAllOf](docs/XmlNs0DiskConfigAllOf.md)
 - [XmlNs0EnclavePort](docs/XmlNs0EnclavePort.md)
 - [XmlNs0EnclavePortAllOf](docs/XmlNs0EnclavePortAllOf.md)
 - [XmlNs0EnclaveProcess](docs/XmlNs0EnclaveProcess.md)
 - [XmlNs0EnclaveProcessAllOf](docs/XmlNs0EnclaveProcessAllOf.md)
 - [XmlNs0EnclaveRunningAttestation](docs/XmlNs0EnclaveRunningAttestation.md)
 - [XmlNs0EnclaveRunningAttestationAllOf](docs/XmlNs0EnclaveRunningAttestationAllOf.md)
 - [XmlNs0EnclaveRunningAttestationApproval](docs/XmlNs0EnclaveRunningAttestationApproval.md)
 - [XmlNs0EnclaveRunningAttestationApprovalAllOf](docs/XmlNs0EnclaveRunningAttestationApprovalAllOf.md)
 - [XmlNs0Environment](docs/XmlNs0Environment.md)
 - [XmlNs0EnvironmentAllOf](docs/XmlNs0EnvironmentAllOf.md)
 - [XmlNs0IPFSSystemInfo](docs/XmlNs0IPFSSystemInfo.md)
 - [XmlNs0IPFSSystemInfoAllOf](docs/XmlNs0IPFSSystemInfoAllOf.md)
 - [XmlNs0InputStream](docs/XmlNs0InputStream.md)
 - [XmlNs0KubernetesEnclave](docs/XmlNs0KubernetesEnclave.md)
 - [XmlNs0KubernetesEnclaveAllOf](docs/XmlNs0KubernetesEnclaveAllOf.md)
 - [XmlNs0MDNSetting](docs/XmlNs0MDNSetting.md)
 - [XmlNs0MDNSettingAllOf](docs/XmlNs0MDNSettingAllOf.md)
 - [XmlNs0MDNUserProfileImage](docs/XmlNs0MDNUserProfileImage.md)
 - [XmlNs0MDNUserProfileImageAllOf](docs/XmlNs0MDNUserProfileImageAllOf.md)
 - [XmlNs0MDNUserSetting](docs/XmlNs0MDNUserSetting.md)
 - [XmlNs0MDNUserSettingAllOf](docs/XmlNs0MDNUserSettingAllOf.md)
 - [XmlNs0NetworkInterface](docs/XmlNs0NetworkInterface.md)
 - [XmlNs0NetworkInterfaceAllOf](docs/XmlNs0NetworkInterfaceAllOf.md)
 - [XmlNs0NodeInfo](docs/XmlNs0NodeInfo.md)
 - [XmlNs0NodeInfoAllOf](docs/XmlNs0NodeInfoAllOf.md)
 - [XmlNs0Process](docs/XmlNs0Process.md)
 - [XmlNs0RunConfig](docs/XmlNs0RunConfig.md)
 - [XmlNs0RunConfigAllOf](docs/XmlNs0RunConfigAllOf.md)
 - [XmlNs0SGXInfo](docs/XmlNs0SGXInfo.md)
 - [XmlNs0SGXInfoAllOf](docs/XmlNs0SGXInfoAllOf.md)
 - [XmlNs0SignedData](docs/XmlNs0SignedData.md)
 - [XmlNs0SignedDataAllOf](docs/XmlNs0SignedDataAllOf.md)
 - [XmlNs0WireguardInterface](docs/XmlNs0WireguardInterface.md)
 - [XmlNs0WireguardInterfaceAllOf](docs/XmlNs0WireguardInterfaceAllOf.md)


## Documentation For Authorization

 Endpoints do not require authorization.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



