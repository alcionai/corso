# Running Corso

Corso is available as a [Docker](https://docs.docker.com/engine/install/) image (Linux `x86_64` and `arm64`) and
as an `x86_64` and `arm64` executable for Windows, Linux and OS X.

:::note

We highly recommend the use of Corso's Docker container when getting started. For Windows, you can run the `amd64`
container in [Linux Mode](https://docs.microsoft.com/en-us/virtualization/windowscontainers/quick-start/quick-start-windows-10-linux).

While our documentation focuses on the Docker command-line, converting the examples to work with native executables
should be straightforward.

:::

As shown in the following example command and described further below, two things are needed to run Corso commands in
its Docker container:

* Environment variables containing configuration information
* A local directory (a `volume` in Docker terminology) for Corso to store configuration and logs

```bash
$ docker run --env-file ~/.corso/corso.env \
    --volume ~/.corso/corso:/app/corso \
    corso/corso <command> <command options>
```

## Environment Variables

Three distinct pieces of configuration are required by Corso:

* S3 object storage configuration to store backups. See [AWS Credentials Setup](/configuration/repos##s3-creds-setup) for
alternate ways to pass AWS credentials.
  * `AWS_ACCESS_KEY_ID`:
  * `AWS_SECRET_ACCESS_KEY`:
  * (Optional) `AWS_SESSION_TOKEN`:

* Microsoft 365 Configuration
  * `AZURE_TENANT_ID`:
  * `AZURE_CLIENT_ID`:
  * `AZURE_CLIENT_SECRET`:

* Corso Security Passphrase
  * `CORSO_PASSPHRASE`:

For ease of use with Docker, we recommend adding the names of the required environment variables (but not their
values!) to a [Docker environment variables file](https://docs.docker.com/engine/reference/commandline/run/#set-environment-variables--e---env---env-file).
To create the environment variables file, you can run the following command:

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

## Local directory

To preserve configuration across container runs and for logs, if enabled, Corso requires access to a directory outside
of its Docker container to read or create its configuration file (`corso.toml`). This directory is mapped by Docker to
the `/app/corso` directory within the container.
