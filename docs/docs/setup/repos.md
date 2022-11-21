---
description: "Configure backup repository"
---

# Repositories

import CodeBlock from '@theme/CodeBlock';
import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';
import TOCInline from '@theme/TOCInline';
import {Version} from '@site/src/corsoEnv';

A Corso [repository](concepts#corso-concepts) stores encrypted copies of your backup data. Repositories are
supported on the following object storage systems:

<TOCInline toc={toc} maxHeadingLevel={2}/><br/>

:::note
Depending on community interest, Corso will add support for other object storage backends in the future.
:::

## Amazon S3

### Prerequisites

Before setting up your Corso S3 repository, the following prerequisites must be met:

* The S3 bucket for the repository already exists. Corso won't create it for you.
* You have access to credentials for a user or an IAM role that has the following permissions:

```json
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "s3:PutObject",
                "s3:GetObject",
                "s3:ListBucket",
                "s3:DeleteObject",
                "s3:GetBucketLocation",
            ],
            "Resource": [
                "arn:aws:s3:::<YOUR_BUCKET_NAME>",
                "arn:aws:s3:::<YOUR_BUCKET_NAME>/*"
            ]
        }
    ]
}
```

### Credential setup {#s3-creds-setup}

Corso supports the credential options offered by the AWS Go SDK. For Full details, see the *Specifying Credentials*
section of the [official documentation](https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html).
The two most commonly-used options are:

* **Environment variables** - set and export `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY`. If using temporary
  credentials derived by assuming an IAM Role, you will also need `AWS_SESSION_TOKEN`.

* **Credentials file** - ensure that the credentials file is available to Corso (for example, may need to map it if
  using Corso as a container). You may also want to set and export `AWS_PROFILE`, if not using the default profile, and
  `AWS_SHARED_CREDENTIALS_FILE`, if not using the default file location. You can learn more about the AWS CLI
  environment variables [here](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-envvars.html).

### Initialize repository

Before first use, you need to initialize a Corso repository with `corso repo init s3`. See the command details
[here](../cli/corso_repo_init_s3).

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

<CodeBlock language="bash">{
`# Initialize the Corso Repository
export CORSO_PASSPHRASE="CHANGE-ME-THIS-IS-INSECURE"
docker run --env-file $HOME/.corso/corso.env \\
  --volume $HOME/.corso:/app/corso ghcr.io/alcionai/corso:${Version()} \\
  repo init s3 --bucket corso-test`
}</CodeBlock>

</TabItem>
</Tabs>

### Connect to a repository

If a repository already exists, you can connect to it with `corso repo connect s3`. See the command details
[here](../cli/corso_repo_connect_s3).

<Tabs groupId="os">
<TabItem value="win" label="Powershell">

  ```powershell
  # Connect to the Corso Repository
  .\corso repo connect s3 --bucket corso-test
  ```

</TabItem>
<TabItem value="unix" label="Linux/macOS">

  ```bash
  # Connect to the Corso Repository
  ./corso repo connect s3 --bucket corso-test
  ```

</TabItem>
<TabItem value="docker" label="Docker">

<CodeBlock language="bash">{
`# Connect to the Corso Repository
docker run --env-file $HOME/.corso/corso.env \\
  --volume $HOME/.corso:/app/corso ghcr.io/alcionai/corso:${Version()} \\
  repo connect s3 --bucket corso-test`
}</CodeBlock>

</TabItem>
</Tabs>

## S3-compatible object storage

Configuring Corso to use object storage systems compatible with the AWS S3 API (for example, Google Cloud Storage,
Backblaze B2, MinIO, etc.) is almost identical to the Amazon S3 instructions above with the exception that you will
need to use the following flag with the initial Corso `repo init` command:

```bash
  --endpoint <domain.example.com>
```

### Testing with insecure TLS configurations

Corso also supports the use of object storage systems with no TLS certificate or with self-signed
TLS certificates with the `--disable-tls` or `--disable-tls-verification` flags.
[These flags](../cli/corso_repo_init_s3) should only be used for testing.
