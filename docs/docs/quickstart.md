# Quick start

import CodeBlock from '@theme/CodeBlock';
import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';
import TOCInline from '@theme/TOCInline';
import {Version} from '@site/src/corsoEnv';
import DownloadBinaries from './setup/_download_binary.md';

This quick start guide runs through the steps you can follow to create your first Microsoft 365 backup and restore:

<TOCInline toc={toc} maxHeadingLevel={2}/>

## Download Corso

Download a Corso binary for your platform. If you would like to use a Docker container, the first `docker run` will
pull the image.

<DownloadBinaries />

## Connect to Microsoft 365

Obtaining credentials from Microsoft 365 to allow Corso to run is a one-time operation. Follow the instructions
[here](../setup/m365_access) to obtain the necessary credentials and then make them available to Corso.

<Tabs groupId="os">
<TabItem value="win" label="Powershell">

  ```powershell
  $Env:AZURE_CLIENT_ID = "<Application (../client) ID for configured app>"
  $Env:AZURE_TENANT_ID = "<Directory (tenant) ID for configured app>"
  $Env:AZURE_CLIENT_SECRET = "<Client secret value>"
  ```

</TabItem>
<TabItem value="unix" label="Linux/macOS">

   ```bash
   export AZURE_CLIENT_ID=<Application (../client) ID for configured app>
   export AZURE_TENANT_ID=<Directory (tenant) ID for configured app>
   export AZURE_CLIENT_SECRET=<Client secret value>
   ```

</TabItem>
<TabItem value="docker" label="Docker">

   ```bash
   export AZURE_CLIENT_ID=<Application (../client) ID for configured app>
   export AZURE_TENANT_ID=<Directory (tenant) ID for configured app>
   export AZURE_CLIENT_SECRET=<Client secret value>
   ```

</TabItem>
</Tabs>

## Create a Corso repository

To create a secure backup location for Corso, you will first need to [download Corso](../setup/download).
Use the container or native executable to initialize the Corso repository using an
[encryption passphrase](../setup/configuration#environment-variables) and a pre-created S3 bucket (Corso doesn't create
the bucket if it doesn't exist). The steps below use `corso-test` as the bucket name but, if you are using AWS, you
will need a different unique name for the bucket.

The following commands assume that in addition to the configuration values from the previous step, `AWS_ACCESS_KEY_ID`
and `AWS_SECRET_ACCESS_KEY` (and `AWS_SESSION_TOKEN` if you are using temporary credentials) are available to the
Corso binary or container.

<Tabs groupId="os">
<TabItem value="win" label="Powershell">

  ```powershell
  # Initialize the Corso Repository
  $Env:CORSO_PASSPHRASE = "CHANGE-ME-THIS-IS-INSECURE"
  .\corso repo init s3 --bucket corso-test
  ```

</TabItem>
<TabItem value="unix" label="Linux/macOS">

  ```bash
  # Initialize the Corso Repository
  export CORSO_PASSPHRASE="CHANGE-ME-THIS-IS-INSECURE"
  ./corso repo init s3 --bucket corso-test
  ```

</TabItem>
<TabItem value="docker" label="Docker">

<!-- vale Vale.Spelling = NO -->
<!-- markdownlint-disable MD022 MD025 MD033 -->

<CodeBlock language="bash">{
`# Create an environment variables file
mkdir -p $HOME/.corso
cat <<EOF > $HOME/.corso/corso.env
AWS_ACCESS_KEY_ID
AWS_SECRET_ACCESS_KEY
AWS_SESSION_TOKEN
AZURE_CLIENT_ID
AZURE_TENANT_ID
AZURE_CLIENT_SECRET
CORSO_PASSPHRASE
EOF
  
# Initialize the Corso Repository
export CORSO_PASSPHRASE="CHANGE-ME-THIS-IS-INSECURE"
docker run --env-file $HOME/.corso/corso.env \\
  --volume $HOME/.corso:/app/corso ghcr.io/alcionai/corso:${Version()} \\
  repo init s3 --bucket corso-test`
}</CodeBlock>

<!-- markdownlint-enable MD022 MD025 MD033 -->
<!-- vale Vale.Spelling = YES -->

</TabItem>
</Tabs>

## Create your first backup

Corso can do much more, but you can start by creating a backup of your Exchange mailbox. If it has been a while since
you initialized the Corso repository, you might need to [connect to it again](../setup/repos#connect-to-a-repository).

<Tabs groupId="os">
<TabItem value="win" label="Powershell">

  ```powershell
  # Backup your inbox
  .\corso backup create exchange --user <your exchange email address>
  ```

</TabItem>
<TabItem value="unix" label="Linux/macOS">

  ```bash
  # Backup your inbox
  ./corso backup create exchange --user <your exchange email address>
  ```

</TabItem>
<TabItem value="docker" label="Docker">

<CodeBlock language="bash">{
`# Backup your inbox
docker run --env-file $HOME/.corso/corso.env \\
  --volume $HOME/.corso:/app/corso ghcr.io/alcionai/corso:${Version()} \\
  backup create exchange --user <your exchange email address>`
}</CodeBlock>

</TabItem>
</Tabs>

:::note
Your first backup may take some time if your mailbox is large.
:::

There will be progress indicators as the backup and, on completion, you should see output similar to:

```text
  Started At            ID                                    Status                Selectors
  2022-10-20T18:28:53Z  d8cd833a-fc63-4872-8981-de5c08e0661b  Completed (0 errors)  alice@contoso.com
```

## Restore an email

Now, lets explore how you can restore data from one of your backups. You can see all Exchange backups available with
the following command:

<Tabs groupId="os">
<TabItem value="win" label="Powershell">

  ```powershell
  # List all Exchange backups
  .\corso backup list exchange
  ```

</TabItem>
<TabItem value="unix" label="Linux/macOS">

  ```bash
  # List all Exchange backups
  ./corso backup list exchange
  ```

</TabItem>
<TabItem value="docker" label="Docker">

<CodeBlock language="bash">{
`# List all Exchange backups
docker run --env-file $HOME/.corso/corso.env \\
  --volume $HOME/.corso:/app/corso ghcr.io/alcionai/corso:${Version()} \\
  backup list exchange`
}</CodeBlock>

</TabItem>
</Tabs>

```text
  Started At            ID                                    Status                Selectors
  2022-10-20T18:28:53Z  d8cd833a-fc63-4872-8981-de5c08e0661b  Completed (0 errors)  alice@contoso.com
  2022-10-20T18:40:45Z  391ceeb3-b44d-4365-9a8e-8a8e1315b565  Completed (0 errors)  alice@contoso.com
  ...
```

Next, select one of the available backups and list all backed up emails. See
[here](../cli/corso_backup_details_exchange) for more advanced filtering options.

<Tabs groupId="os">
<TabItem value="win" label="Powershell">

  ```powershell
  # List emails in a selected backup
  .\corso backup details exchange --backup <id of your selected backup> --email "*" | Select-Object -First 5
  ```

</TabItem>
<TabItem value="unix" label="Linux/macOS">

  ```bash
  # List emails in a selected backup
  ./corso backup details exchange --backup <id of your selected backup> --email "*" | head
  ```

</TabItem>
<TabItem value="docker" label="Docker">

<CodeBlock language="bash">{
`# List emails in a selected backup
docker run --env-file $HOME/.corso/corso.env \\
  --volume $HOME/.corso:/app/corso ghcr.io/alcionai/corso:${Version()} \\
  backup details exchange --backup <id of your selected backup> --email "*" | head`
}</CodeBlock>

</TabItem>
</Tabs>

The output from the command above should display a list of any matching emails. Note the ID
of the email you would like to use for testing restore.

```text
  ID            Sender                 Subject                                  Received
  360bf6840396  phish@contoso.info     Re: Request for Apple/Amazon gift cards  2022-10-18T02:27:47Z
  84dbad89b9f5  ravi@cohovineyard.com  Come join us!                            2022-10-19T06:12:08Z
  ...
```

To restore the selected email, use the following command.

<Tabs groupId="os">
<TabItem value="win" label="Powershell">

  ```powershell
  # Restore a selected email
  .\corso restore exchange --backup <id of your selected backup> --email <email ID>
  ```

</TabItem>
<TabItem value="unix" label="Linux/macOS">

  ```bash
  # Restore a selected email
  ./corso restore exchange --backup <id of your selected backup> --email <email ID>
  ```

</TabItem>
<TabItem value="docker" label="Docker">

<CodeBlock language="bash">{
`# Restore a selected email
docker run --env-file $HOME/.corso/corso.env \\
  --volume $HOME/.corso:/app/corso ghcr.io/alcionai/corso:${Version()} \\
  restore exchange --backup <id of your selected backup> --email <email ID>`
}</CodeBlock>

</TabItem>
</Tabs>

A confirmation of the recovered email will be shown and the email will appear in a new mailbox folder named `Corso_Restore_DD-MMM-YYYY_HH:MM:SS`.

```text
  ID            Sender                 Subject                                  Received
  360bf6840396  phish@contoso.info     Re: Request for Apple/Amazon gift cards  2022-10-18T02:27:47Z
```

## Read more

The above tutorial only scratches the surface for Corso's capabilities. We encourage you to dig deeper by:

* Learning about [Corso concepts and setup](../setup/concepts)
* Explore Corso backup and restore options for Exchange and Onedrive in the [Command Line Reference](../cli/corso)
