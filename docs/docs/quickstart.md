# Quick start

import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';
import TOCInline from '@theme/TOCInline';

This quick start guide runs through the steps you can follow to create your first Microsoft 365 backup and restore:

<TOCInline toc={toc} maxHeadingLevel={2}/>

## Connecting to Microsoft 365

Obtaining credentials from Microsoft 365 to allow Corso to run is a moderately involved one-time operation.
Follow the instructions [here](setup/m365_access) to obtain the necessary credentials and then make them available to
Corso.

<Tabs groupId="os">
<TabItem value="win" label="Powershell">

  ```powershell
  $Env:AZURE_CLIENT_ID = "<Directory (tenant) ID for configured app>"
  $Env:AZURE_TENANT_ID = "<Application (client) ID for configured app>"
  $Env:AZURE_CLIENT_SECRET = "<Client secret value>"
  ```

</TabItem>
<TabItem value="unix" label="Linux/macOS">

   ```bash
   export AZURE_TENANT_ID=<Directory (tenant) ID for configured app>
   export AZURE_CLIENT_ID=<Application (client) ID for configured app>
   export AZURE_CLIENT_SECRET=<Client secret value>
   ```

</TabItem>
<TabItem value="docker" label="Docker">

   ```bash
   export AZURE_TENANT_ID=<Directory (tenant) ID for configured app>
   export AZURE_CLIENT_ID=<Application (client) ID for configured app>
   export AZURE_CLIENT_SECRET=<Client secret value>
   ```

</TabItem>
</Tabs>

## Repository creation

To create a secure backup location for Corso, we will create a bucket (`corso-test` is used as an example) in AWS S3.
The following commands assume that all configuration values from the previous step, `AWS_ACCESS_KEY_ID`, and
`AWS_SECRET_ACCESS_KEY` are available to the Corso binary or container.

<Tabs groupId="os">
<TabItem value="win" label="Powershell">

  ```powershell
  # Create the AWS S3 Bucket
  aws s3api create-bucket --bucket corso-test
  ```

</TabItem>
<TabItem value="unix" label="Linux/macOS">

  ```bash
  # Create the AWS S3 Bucket
  aws s3api create-bucket --bucket corso-test
  ```

</TabItem>
<TabItem value="docker" label="Docker">

  ```bash
  # Create the AWS S3 Bucket
  aws s3api create-bucket --bucket corso-test
  ```

</TabItem>
</Tabs>

Next, let's initialize the Corso repository using an encryption passphrase.
<Tabs groupId="os">
<TabItem value="win" label="Powershell">

  ```powershell
  # Initialize the Corso Repository
  $Env:CORSO_PASSPHRASE = "CHANGE-ME-THIS-IS-INSECURE"
  .\corso.exe repo init s3 --bucket corso-test
  ```

</TabItem>
<TabItem value="unix" label="Linux/macOS">

  ```bash
  # Initialize the Corso Repository
  export CORSO_PASSPHRASE="CHANGE-ME-THIS-IS-INSECURE"
  corso repo init s3 --bucket corso-test
  ```

</TabItem>
<TabItem value="docker" label="Docker">

  ```bash
  # Initialize the Corso Repository
  export CORSO_PASSPHRASE="CHANGE-ME-THIS-IS-INSECURE"
  docker run --env-file ~/.corso/corso.env \
    --volume $HOME/.corso:/app/corso corso/corso:latest \
    repo init s3 --bucket corso-test
  ```

</TabItem>
</Tabs>

## Your first backup

Corso can do much more, but you can start by creating a backup of your Exchange mailbox.

To do this, you can run the following command:

```bash
$ docker run -e CORSO_PASSPHRASE \
    --env-file ~/.corso/corso.env \
    -v ~/.corso:/app/corso corso/corso:<release tag> \
    backup create exchange --user <your exchange email address>

  Started At            ID                                    Status                Selectors
  2022-10-10T19:46:43Z  41e93db7-650d-44ce-b721-ae2e8071c728  Completed (0 errors)  alice@example.com
```

:::note
Your first backup may take some time if your mailbox is large.
:::

## Restore an email

Now lets explore how you can restore data from one of your backups.

You can see all Exchange backups available with the following command:

```bash
$ docker run -e CORSO_PASSPHRASE \
    --env-file ~/.corso/corso.env \
    -v ~/.corso:/app/corso corso/corso:<release tag> \
    backup list exchange 

  Started At            ID                                    Status                Selectors
  2022-09-09T42:27:16Z  72d12ef6-420a-15bd-c862-fd7c9023a014  Completed (0 errors)  alice@example.com
  2022-10-10T19:46:43Z  41e93db7-650d-44ce-b721-ae2e8071c728  Completed (0 errors)  alice@example.com
```

Select one of the available backups and search through its contents.

```bash
$ docker run -e CORSO_PASSPHRASE \
    --env-file ~/.corso/corso.env \
    -v ~/.corso:/app/corso corso/corso:<release tag> \
    backup details exchange \
    --backup <id of your selected backup> \
    --user <your exchange email address> \
    --email-subject <portion of subject of email you want to recover>
```

The output from the command above should display a list of any matching emails. Note the ID
of the one to use for testing restore.

When you are ready to restore, use the following command:

```bash
$ docker run -e CORSO_PASSPHRASE \
    --env-file ~/.corso/corso.env \
    -v ~/.corso:/app/corso corso/corso:<release tag> \
    backup details exchange \
    --backup <id of your selected backup> \
    --user <your exchange email address> \
    --email <id of your selected email>
```

You can now find the recovered email in a mailbox folder named `Corso_Restore_DD-MMM-YYYY_HH:MM:SS`.

You are now ready to explore the [Command Line Reference](cli/corso) and try everything that Corso can do.
