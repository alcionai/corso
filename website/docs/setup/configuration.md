# Configuration

import CodeBlock from '@theme/CodeBlock';
import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';
import {Version} from '@site/src/corsoEnv';

Two things are needed to configure Corso:

* Environment variables containing configuration information
* A directory for Corso to store its configuration file

## Environment variables

Three distinct pieces of configuration are required by Corso:

* S3 object storage configuration to store backups. See [AWS Credentials Setup](../repos#s3-creds-setup) for
alternate ways to pass AWS credentials.
  * `AWS_ACCESS_KEY_ID`: Access key for an IAM user or role for accessing an S3 bucket
  * `AWS_SECRET_ACCESS_KEY`: Secret key associated with the access key
  * (Optional) `AWS_SESSION_TOKEN`: Session token required when using temporary credentials

* Microsoft 365 Configuration
  * `AZURE_CLIENT_ID`: Client ID for your Azure AD application used to access your M365 tenant
  * `AZURE_TENANT_ID`: ID for the M365 tenant where the Azure AD application is registered
  * `AZURE_CLIENT_SECRET`: Azure secret for your Azure AD application used to access your M365 tenant

* Corso Security Passphrase
  * `CORSO_PASSPHRASE`: Passphrase to protect encrypted repository contents

<Tabs groupId="os">
<TabItem value="win" label="Powershell">

Ensure that all of the above environment variables are defined in your Powershell environment.

  ```powershell
  $Env:AWS_ACCESS_KEY_ID = "..."
  $Env:AWS_SECRET_ACCESS_KEY = "..."
  $Env:AWS_SESSION_TOKEN = ""

  $Env:AZURE_CLIENT_ID = "..."
  $Env:AZURE_TENANT_ID = "..."
  $Env:AZURE_CLIENT_SECRET = "..."

  $Env:CORSO_PASSPHRASE = "CHANGE-ME-THIS-IS-INSECURE"
  ```

</TabItem>
<TabItem value="unix" label="Linux/macOS">

Ensure that all of the above environment variables are defined in your shell environment.

  ```bash
  export AWS_ACCESS_KEY_ID=...
  export AWS_SECRET_ACCESS_KEY=...
  export AWS_SESSION_TOKEN=...

  export AZURE_CLIENT_ID=...
  export AZURE_TENANT_ID=...
  export AZURE_CLIENT_SECRET=...

  export CORSO_PASSPHRASE=CHANGE-ME-THIS-IS-INSECURE
  ```

</TabItem>
<TabItem value="docker" label="Docker">

For ease of use with Docker, we recommend adding the names of the required environment variables (but not their
values!) to a [Docker environment variables file](https://docs.docker.com/engine/reference/commandline/run/#set-environment-variables--e---env---env-file).
To create the environment variables file, you can run the following command:

  ```bash
  # Create an environment variables file
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

  # Export required variables
  export AWS_ACCESS_KEY_ID=...
  export AWS_SECRET_ACCESS_KEY=...
  export AWS_SESSION_TOKEN=...

  export AZURE_CLIENT_ID=...
  export AZURE_TENANT_ID=...
  export AZURE_CLIENT_SECRET=...

  export CORSO_PASSPHRASE=CHANGE-ME-THIS-IS-INSECURE
  ```

</TabItem>
</Tabs>

## Configuration File

<Tabs groupId="os">
<TabItem value="win" label="Powershell">

By default, Corso stores its configuration file (`.corso.toml`) in the user's home directory.
The location of the configuration file can be specified using the `--config-file` option.

</TabItem>
<TabItem value="unix" label="Linux/macOS">

By default, Corso stores its configuration file (`.corso.toml`) in the user's home directory.
The location of the configuration file can be specified using the `--config-file` option.

</TabItem>
<TabItem value="docker" label="Docker">

To preserve configuration across container runs, Corso requires access to a directory outside of its Docker container
to read or create its configuration file (`.corso.toml`). This directory must be mapped, by Docker, to the `/app/corso`
directory within the container.

<CodeBlock language="bash">{
`docker run --env-file $HOME/.corso/corso.env \\
  --volume $HOME/.corso:/app/corso ghcr.io/alcionai/corso:${Version()} \\
  <command> <command options>`
}</CodeBlock>

</TabItem>
</Tabs>

## Log Files

The default location of Corso's log file is shown below but the location can be overridden by using the `--log-file` flag.
You can also use `stdout` or `stderr` as the `--log-file` location to redirect the logs to "stdout" and "stderr" respectively.

<Tabs groupId="os">
<TabItem value="win" label="Windows">

  ```powershell
  %LocalAppData%\corso\logs\<timestamp>.log
  ```

</TabItem>
<TabItem value="unix" label="Linux">

  ```bash
  $HOME/.cache/corso/logs/<timestamp>.log
  ```

</TabItem>
<TabItem value="macos" label="macOS">

  ```bash
  $HOME/Library/Logs/corso/logs/<timestamp>.log
  ```

</TabItem>
</Tabs>
