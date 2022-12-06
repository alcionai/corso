---
description: "Connect to a Microsft 365 tenant"
---

# Microsoft 365 access

import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';

To perform backup and restore operations, Corso requires access to your [M365 tenant](../concepts#m365-concepts)
by creating an [Azure AD application](../concepts#m365-concepts) with appropriate permissions.

The following steps outline a simplified procedure for creating an Azure Ad application suitable for use with Corso.
For more details, please refer to the
[official documentation](https://docs.microsoft.com/en-us/azure/active-directory/develop/howto-create-service-principal-portal)
for adding an Azure AD Application and Service Principal using the Azure Portal.

## Create an Azure AD application

Sign in into the [Azure Portal](https://portal.azure.com/) with a user that has sufficient permissions to create an
AD application.

### Register a new application

From the list of [Azure services](https://portal.azure.com/#allservices), select
**Azure Active Directory &#8594; App Registrations &#8594; New Registration**

![Registering a new application](/img/m365app_create_new.png)

### Configure basic settings

Next, configure the following:

* Give the application a name
* Select **Accounts in this organizational directory only**
* Skip the **Redirect URI** option
* Click **Register** at the bottom of the screen

![Configuring the application](/img/m365app_configure.png)

### Configure required permissions

Within the new application (`CorsoApp` in the below diagram), select **API Permissions &#8594; Add a permission** from
the management panel.

![Adding application permissions](/img/m365app_permissions.png)

Select the following permissions from **Microsoft API &#8594; Microsoft Graph &#8594; Application Permissions** and
then click **Add permissions**.

<!-- vale Microsoft.Spacing = NO -->
| API / Permissions Name | Type | Description
|:--|:--|:--|
| Calendars.ReadWrite | Application | Read and write calendars in all mailboxes |
| Contacts.ReadWrite | Application | Read and write contacts in all mailboxes |
| Files.ReadWrite.All | Application | Read and write files in all site collections |
| Mail.ReadWrite | Application | Read and write mail in all mailboxes |
| User.Read.All | Application | Read all users' full profiles |
<!-- vale Microsoft.Spacing = YES -->

### Grant admin consent

Finally, grant admin consent to this application. This step is required even if the user that created the application
is an Microsoft 365 admin.

![Granting administrator consent](/img/m365app_consent.png)

## Export application credentials

After configuring the Corso Azure AD application, store the information needed by Corso to connect to the application
as environment variables.

### Tenant ID and client ID

To view the tenant and client ID, select Overview from the app management panel.

![Obtaining Tenant and Client IDs](/img/m365app_ids.png)

Copy the client and tenant IDs and export them into the following environment variables.

<Tabs groupId="os">
<TabItem value="win" label="Powershell">

  ```powershell
  $Env:AZURE_CLIENT_ID = "<Application (client) ID for configured app>"
  $Env:AZURE_TENANT_ID = "<Directory (tenant) ID for configured app>"
  ```

</TabItem>
<TabItem value="unix" label="Linux/macOS">

   ```bash
   export AZURE_TENANT_ID=<Directory (tenant) ID for configured app>
   export AZURE_CLIENT_ID=<Application (client) ID for configured app>
   ```

</TabItem>
<TabItem value="docker" label="Docker">

   ```bash
   export AZURE_TENANT_ID=<Directory (tenant) ID for configured app>
   export AZURE_CLIENT_ID=<Application (client) ID for configured app>
   ```

</TabItem>
</Tabs>

### Azure client secret

Finally, you need to obtain a client secret associated with the app using **Certificates & Secrets** from the app
management panel.

Click **New Client Secret** under **Client secrets** and follow the instructions to create a secret.

![Obtaining the Azure client secrete](/img/m365app_secret.png)

After creating the secret, immediately copy the secret **Value** because it won't be available later. Export it as an
environment variable.

<Tabs groupId="os">
<TabItem value="win" label="Powershell">

  ```powershell
  $Env:AZURE_CLIENT_SECRET = "<Client secret value>"
  ```

</TabItem>
<TabItem value="unix" label="Linux/macOS">

   ```bash
   export AZURE_CLIENT_SECRET=<Client secret value>
   ```

</TabItem>
<TabItem value="docker" label="Docker">

   ```bash
   export AZURE_CLIENT_SECRET=<Client secret value>
   ```

</TabItem>
</Tabs>
