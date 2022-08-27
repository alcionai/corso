---
description: "Connect to M365 Tenant"
---

# M365 access

To perform backup and restore operations, Corso requires access to your [M365 tenant](/concepts#m365-concepts)
through a properly configured [Azure AD application](/concepts#m365-concepts).

## Create Azure AD application

For the official documentation for adding an Azure AD Application and Service Principal using the Azure Portal see
[here](https://docs.microsoft.com/en-us/azure/active-directory/develop/howto-create-service-principal-portal).

The following steps outline a simplified procedure for creating an Azure Ad application suitable for use with Corso.

1. **Create a new application**

   Select **Azure Active Directory &#8594; App Registrations &#8594; New Registration**
   <img src="/img/m365app_create_new.png" className="guideImages"/>

2. **Configure basic settings**

   * Give the application a name
   * Select **Accounts in this organizational directory only**
   * Skip the **Redirect URI** option

   <br/><img src="/img/m365app_configure.png" className="guideImages"/>

3. **Configure required permissions**

   Select **App Permission** from the app management panel.

   <img src="/img/m365app_permissions.png" className="guideImages"/>

   The required permissions are as follows:

   **TODO: Complete list of permissions** 

   | API / Permissions Name | Type |
   |--|--|
   | Permission 1 | Application |
   | Permission 2 | Delegated |

4. **Grant admin consent**

   <img src="/img/m365app_consent.png" className="guideImages"/>

## Export application credentials

Now that the Corso Azure AD application is configured, you need to capture as environment variables the information that
Corso will need to connect to the application.

### Tenant ID and client ID

To obtain these, select Overview from the app management panel and export the corresponding environment variables.

```bash
export AZURE_TENANT_ID=<Directory (tenent) ID for configured app>
export AZURE_CLIENT_ID=<Application (client) ID for configured app>
```

<img src="/img/m365app_ids.png" className="guideImages"/>

### Azure client secret

Lastly, you need to configure a client secret associated with the app using **Certificates & Secrets** from the app
management panel.

Click **New Client Secret** and follow the instructions to create a secret. After creating the secret, copy the secret
value right away because it's only available after creation. In the end, export the secret value as an environment
variable.

```bash
export AZURE_CLIENT_SECRET=<client secret value>
```

<img src="/img/m365app_secret.png" className="guideImages"/>
