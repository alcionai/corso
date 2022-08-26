# Installation

Corso releases are currently available using the following options:

import TOCInline from '@theme/TOCInline';

<TOCInline toc={toc} maxHeadingLevel={2}/><br/>

:::note

To maximize portability accross platforms, Corso is currently distributed as a container image. In the future,
releases may also be available as operating system specific pre-built binaries.

In the meantime, if you are interested in running Corso as a binary, refer to the
[instructions on how to build from source](developers/build).

:::

## Docker Image

To use Corso as a Docker image, you need to have [Docker installed](https://docs.docker.com/engine/install/)
on your machine.

### Docker Command

To run the Corso container, it is recommended that you:

* Export [Corso key configuration environment variables](cli/corso_env) and add their names to an
[environemnt variables file](https://docs.docker.com/engine/reference/commandline/run/#set-environment-variables--e---env---env-file)
* Map a local config directory to `/app/config` where Corso will look for or create the `corso.toml` config file.
This will preserve configuration accross container runs.
* Map a local log directory to `/app/logs` where Corso will write logs

To create the environment variables file, you can run the following.

```bash
# create an env vars file
$ cat <<EOF ~/.corso/corso.env 
CORSO_PASSPHRASE
AZURE_TENANT_ID
AZURE_CLIENT_ID
AZURE_CLIENT_SECRET
AWS_ACCESS_KEY_ID
AWS_SECRET_ACCESS_KEY
AWS_SESSION_TOKEN
EOF
```

:::note

Depending on your preferred method for passing AWS credentials, you may need to adapt the `AWS_` variables in the file.
See [AWS Credentials Setup](/configuration/repos##s3-creds-setup) for more details.

:::

The following command will list the Corso Exchange backups. You can adapt the folder mappings, container tag, and the command
as needed.

```bash
$ docker run --env-file ~/.corso/corso.env \
    -v ~/.corso/config:/app/config \ 
    -v ~/.corso/logs:/app/logs \
    corso/corso:latest backup list exchange 
```

### Available Variants

The Corso image is available on DockerHub for the following architectures:

* Linux x86-64 - `amd64`
* ARM 64-bit - `arm64`

:::tip

For Windows, you can run the `amd64` container in
[Linux Mode](https://docs.microsoft.com/en-us/virtualization/windowscontainers/quick-start/quick-start-windows-10-linux)

:::

Additionally, the following tags are available:

* `:latest` - The most recent stable release
* `:x.y.z` - A specific release build
* `:pre-release` - The most recent pre-release (beta or RC) if newer that the latest stable release
* `:nightly` - The most recent unstable nightly developer build
* `:SHA` - A specific nightly build
