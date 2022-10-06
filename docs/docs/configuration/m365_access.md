---
description: "Connect to a Microsft 365 tenant"
---

# Microsoft 365 access

To perform backup and restore operations, Corso requires access to your [M365 tenant](concepts#m365-concepts)
through an [Azure AD application](concepts#m365-concepts) with appropriate permissions.

## Create an Azure AD application

For the official documentation for adding an Azure AD Application and Service Principal using the Azure Portal see
[here](https://docs.microsoft.com/en-us/azure/active-directory/develop/howto-create-service-principal-portal).

The following steps outline a simplified procedure for creating an Azure Ad application suitable for use with Corso.

1. **Create a new application**

 Select **Azure Active Directory &#8594; App Registrations &#8594; New Registration**
 <img src="/img/m365app_create_new.png" className="guideImages"/>

1. **Configure basic settings**

   * Give the application a name
   * Select **Accounts in this organizational directory only**
   * Skip the **Redirect URI** option

   <br/><img src="/img/m365app_configure.png" className="guideImages"/>

1. **Configure required permissions**

   Select **App Permission** from the app management panel.

   <img src="/img/m365app_permissions.png" className="guideImages"/>

   The required permissions are as follows:

   | API / Permissions Name | Type | Description
   |:--|:--|:--|
   | Calendars.ReadWrite | Application | Read and write calendars in all mailboxes |
   | Contacts.ReadWrite | Application | Read and write contacts in all mailboxes |
   | Files.ReadWrite.All | Application | Read and write files in all site collections |
   | Mail.ReadWrite | Application | Read and write mail in all mailboxes |
   | User.Read.All | Application | Read all users' full profiles |

1. **Grant admin consent**

   <img src="/img/m365app_consent.png" className="guideImages"/>

## Export application credentials

After configuring the Corso Azure AD application, store the information needed by Corso to connect to the application
as environment variables.

### Tenant ID and client ID

To extract the tenant and client ID, select Overview from the app management panel and export the corresponding
environment variables.

```bash
export AZURE_TENANT_ID=<Directory (tenent) ID for configured app>
export AZURE_CLIENT_ID=<Application (client) ID for configured app>
```

<img src="/img/m365app_ids.png" className="guideImages"/>

### Azure client secret

Lastly, you need to configure a client secret associated with the app using **Certificates & Secrets** from the app
management panel.

Click **New Client Secret** and follow the instructions to create a secret. After creating the secret, copy the secret
value right away because it won't be available later and export it as an environment variable.

```bash
export AZURE_CLIENT_SECRET=<client secret value>
```

<img src="/img/m365app_secret.png" className="guideImages"/>
