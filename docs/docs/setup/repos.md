---
description: "Configure backup repository"
---

# Repositories

A Corso [repository](concepts#corso-concepts) stores encrypted copies of your backup data. Repositories are
supported on the following object storage systems:

import TOCInline from '@theme/TOCInline';

<TOCInline toc={toc} maxHeadingLevel={2}/><br/>

:::note
Depending on community interest, Corso may support other object storage backends in the future.
:::

## Amazon S3

### Prerequisites

Before setting you your Corso S3 repository, the following prerequisites must be met:

* S3 bucket for the repository already exists. Corso won't create it for you.
* You have access to credentials for a user or an IAM role that represent the following permissions

<!-- vale proselint.Annotations = NO -->
**TODO: Verify if these permissions are correct? What about multi-part upload permissions?**
<!-- vale proselint.Annotations = YES -->

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
                "s3:AbortMultipartUpload", 
                "s3:ListMultipartUploadParts",
                "s3:ListBucketMultipartUploads"
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

* **Environment variables** - set and export `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY`. If using temporary
  credentials derived by assuming an IAM Role, you will also need `AWS_SESSION_TOKEN`.

* **Credentials file** - ensure that the credentials file is available to Corso (for example, may need to map it if
  using Corso as a container). You may also want to set and export `AWS_PROFILE`, if not using the default profile, and
  `AWS_SHARED_CREDENTIALS_FILE`, if not using the default file location. You can learn more about the AWS CLI
  environment variables [here](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-envvars.html).

### Initialize repository

Before first use, you need to initialize a Corso repository with `corso repo init s3`. See command details
[here](/cli/corso_repo_init_s3).

If a repository already exists, you can connect to it with `corso repo connect s3`. See command details
[here](/cli/corso_repo_connect_s3).
