# MSGraph Authentication

There are several types of authentication provider option which are not
supported at this time by msgraph-sdk-go. A full discussion on the supported providers are found [here](https://docs.microsoft.com/en-us/graph/sdks/choose-authentication-providers?tabs=CS). Additional reading on OAuth Credential [flow](https://docs.microsoft.com/en-us/azure/active-directory/develop/v2-oauth2-client-creds-grant-flow). The repository has the following methods checked-in at this time.


**Authentication Methods (Checked-Items have verified Go solutions)**

- [x]  Interactive provider
- [ ]  Authorization code provider
- [ ]  On-behalf-of provider
- [ ]  Device code provider
- [x]  Username/password provider
- [x]  Client credentials provider

### **Interactive Authentication Provider**

Shows a UI popup provided by the Auth Server. Individuals can input their credentials via the UI.

### **UsernamePassword Authentication Provider**

Passes username/password over the wire to the AuthServer. No interaction
required; however, this authentication cannot be used when 2FA enabled.

```go
"error": "invalid_grant",
  "error_description": "AADSTS50076: Due to a configuration change made by your administrator, or because you moved to a new location, you must use multi-factor authentication to access '00000003-0000-0000-c000-000000000000'.\r\nTrace ID: 5e178ca4-7a77-491a-a607-1f0a8187d500\r\nCorrelation ID: a3dce328-f8d3-4932-bf95-d2a89693f7e5\r\nTimestamp: 2022-04-11 15:39:35Z",
  "error_codes": [
    50076
  ],
  "timestamp": "2022-04-11 15:39:35Z",
  "trace_id": "5e178ca4-7a77-491a-a607-1f0a8187d500",
  "correlation_id": "a3dce328-f8d3-4932-bf95-d2a89693f7e5",
  "error_uri": "https://login.microsoftonline.com/error?code=50076",
  "suberror": "basic_action"
}
```

### **IntegratedWindows Authentication Provider**

Grabs credentials from the OS (Not supported in Go as of yet)

### **DeviceCode Authentication Provider**

Outputs a code that needs to be manually entered in a web form at a designated
URL.

### ClientCredential Authentication Provider
No interaction required during authentication and works with 2FA authentication. A client secret is required to be created on behalf of the app to utilize this method. The msgraph application uses this method currently.