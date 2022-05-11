# Microsoft Graph SDK for Go

Notes on how to create an introductory Golang application utilizing the
Microsoft Graph SDK for Go. The application requests user to log via a web
browser and prints the user’s display name. Microsoft SDK GitHub repository is
located here.

## Installation

**Locally install modules**

1. Navigate to empty directory
3. Install modules

```bash
go get
```

## Registration

To register application, navigate to this
[link](https://docs.microsoft.com/en-gb/graph/auth-register-app-v2) and log in
with your developer account information. Link will direct to the Azure Portal.

- Sign-in with Alcion dev account (note this is not your regular Alcion
    account).
  - If your account gives you access to more than one tenant, select your
    account in the top right corner, and set your portal session to the Azure
    AD tenant that you want. Select developer account if applicable.
- In the left-hand navigation pane, select the **Azure Active Directory**
  service, and then select **App registrations → New registration**
- When the **Register an application** page appears, enter your application's
  registration information
  - **Name** - Enter a meaningful application name that will be displayed to
    users of the app.
  - **Supported account types** - Select which accounts you would like your
    application to support → Accounts in **THIS** organizational directory only
    (Select for prototyping)
  - **Redirect URI (optional)**— This can be changed later.
    - Select Platform →  Mobile or Desktop App
    - URI → [`http://localhost:31544`](http://localhost:31544/) [or your
      desired port]

Click the  *Register* button

## Permissions

Applications can have permissions added one of two ways. In code, or in
[portal.azure.com](http://portal.azure.com). The later is described here.
First navigate to **Azure Active Directory → App registrations** on the left nav
bar. Once there, click your app and then scroll down the ‘Overview’ page and
click, “View API Permissions”.

**Add a Permission → Microsoft Graph → Delegated permissions → \<make
selections\>**

Finally, click **ADD permissions**

## Creating a Program

TYPE :: Golang.

**Data Flow:**   Select Authentication type, request permissions, client
creation, data query, exit.

Program creates a new browser tab. User will have to authenticate through
Microsoft for the code to complete.

Updated: 4/11/2022

```go
package main

import (
        "context"
        "fmt"

        azidentity "github.com/Azure/azure-sdk-for-go/sdk/azidentity"
        a "github.com/microsoft/kiota-authentication-azure-go"
        msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
    )

func main(){
    cred, err := az.NewInteractiveBrowserCredential(&az.InteractiveBrowserCredentialOptions {
        TenantID:    "Tenant_ID",
        ClientID:    "Client_ID",
        RedirectURL: "Redirect_URL",
    })

    if err != nil {
        fmt.Printf("Error creating credentials: %v\n", err)
        return
    }

    auth, err := a.NewAzureIdentityAuthenticationProviderWithScopes(cred, []string{"User.Read"})

    if err != nil {
        fmt.Printf("Error authentication provider: %v\n", err)
        return
    }

    adapter, err := msgraphsdk.NewGraphRequestAdapter(auth)
    if err != nil {
        fmt.Printf("Error creating adapter: %v\n", err)
        return
    }
    client := msgraphsdk.NewGraphServiceClient(adapter)
    fmt.Printf("Runtime Passes: %v", client)
}
```

## Complications

### Library Errors Clients

Supporting libraries are unstable at present date. The Microsoft Graph SDK
documentation was updated in March 2022, yet file directories and the
repository were changed by April 7, 2022. The [change
log](https://github.com/microsoftgraph/msgraph-sdk-go/blob/main/CHANGELOG.md)
will speak to the difference in the new release. They do not always contain
instructions on how to navigate those changes.

```go
"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
"github.com/microsoft/kiota-authentication-azure-go"
```

Original documentation for client initialization
<[link](https://docs.microsoft.com/en-gb/graph/sdks/create-client?tabs=Go)>
creates following errors. Error relatively easy to correct.

```bash
./clientTest.go:33:55: cannot use auth (variable of type *microsoft_kiota_authentication_azure.AzureIdentityAuthenticationProvider) as type "github.com/microsoft/kiota-abstractions-go/authentication".AuthenticationProvider in argument to msgraphsdk.NewGraphRequestAdapter:
		*microsoft_kiota_authentication_azure.AzureIdentityAuthenticationProvider does not implement "github.com/microsoft/kiota-abstractions-go/authentication".AuthenticationProvider (wrong type for AuthenticateRequest method)
    		have AuthenticateRequest(request "github.com/microsoft/kiota/abstractions/go".RequestInformation) error
        	want AuthenticateRequest(request *"github.com/microsoft/kiota-abstractions-go".RequestInformation) error
```

