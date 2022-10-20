# Configuration

import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';

Corso is available as a [Docker](https://docs.docker.com/engine/install/) image (Linux `x86_64` and `arm64`) and
as an `x86_64` and `arm64` executable for Windows, Linux and macOS.

Two things are needed to run Corso:

* Environment variables containing configuration information
* A directory for Corso to store its configuration file

## Environment Variables

Three distinct pieces of configuration are required by Corso:

* S3 object storage configuration to store backups. See [AWS Credentials Setup](/setup/repos##s3-creds-setup) for
alternate ways to pass AWS credentials.
  * `AWS_ACCESS_KEY_ID`: Access key for an IAM user or role for accessing an S3 bucket
  * `AWS_SECRET_ACCESS_KEY`: Secret key associated with the access key
  * (Optional) `AWS_SESSION_TOKEN`: Session token required when using temporary credentials

* Microsoft 365 Configuration
  * `AZURE_CLIENT_ID`: Client ID for your Azure AD application used to access your M365 tenant
  * `AZURE_CLIENT_SECRET`: Azure secret for your Azure AD application used to access your M365 tenant
  * `AZURE_TENANT_ID`: ID for the M365 tenant where the Azure AD application is registered

* Corso Security Passphrase
  * `CORSO_PASSPHRASE`: Passphrase to protect encrypted repository contents

<Tabs groupId="os">
<TabItem value="win" label="Powershell">

Ensure that all of the above environment variables are available in your Powershell environment.

</TabItem>
<TabItem value="docker" label="Docker">

For ease of use with Docker, we recommend adding the names of the required environment variables (but not their
values!) to a [Docker environment variables file](https://docs.docker.com/engine/reference/commandline/run/#set-environment-variables--e---env---env-file).
To create the environment variables file, you can run the following command:

  ```bash
  # Create an environment variables file
  cat <<EOF ~/.corso/corso.env
  CORSO_PASSPHRASE
  AZURE_TENANT_ID
  AZURE_CLIENT_ID
  AZURE_CLIENT_SECRET
  AWS_ACCESS_KEY_ID
  AWS_SECRET_ACCESS_KEY
  AWS_SESSION_TOKEN
  EOF
  ```

</TabItem>
</Tabs>

## Configuration File

<Tabs groupId="os">
<TabItem value="win" label="Powershell">

By default, Corso store its configuration file (`.corso.toml`) in the directory where the binary is executed.
The location of the configuration file can be specified using the `--config-file` option.

</TabItem>
<TabItem value="docker" label="Docker">

To preserve configuration across container runs, Corso requires access to a directory outside of its Docker container
to read or create its configuration file (`.corso.toml`). This directory must be mapped, by Docker, to the `/app/corso`
directory within the container.

```bash
$ docker run --env-file ~/.corso/corso.env \
    --volume ~/.corso/corso:/app/corso \
    corso/corso <command> <command options>
```

</TabItem>
</Tabs>
