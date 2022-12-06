---
title: corso repo connect s3
hide_title: true
---
## corso repo connect s3

Ensures a connection to an existing S3 repository.

```bash
corso repo connect s3 --bucket <bucket> [flags]
```

### Examples

```bash
# Connect to a Corso repo in AWS S3 bucket named "my-bucket"
corso repo connect s3 --bucket my-bucket

# Connect to a Corso repo in AWS S3 bucket named "my-bucket" using a prefix
corso repo connect s3 --bucket my-bucket --prefix my-prefix

# Connect to a Corso repo in an S3 compliant storage provider
corso repo connect s3 --bucket my-bucket --endpoint https://my-s3-server-endpoint
```

### Flags

|Flag|Short|Default|Help|
|:----|:-----|:-------|:----|
|`--bucket`|||Name of S3 bucket for repo. <div class='required'>Required</div>|
|`--prefix`|||Repo prefix within bucket.|
|`--endpoint`||`s3.amazonaws.com`|S3 service endpoint.|
|`--disable-tls`||`false`|Disable TLS (HTTPS)|
|`--disable-tls-verification`||`false`|Disable TLS (HTTPS) certificate verification.|
|`--help`|`-h`|`false`|help for s3|

### Global and inherited flags

|Flag|Short|Default|Help|
|:----|:-----|:-------|:----|
|`--config-file`||`$HOME/.corso.toml`|config file location (default is $HOME/.corso.toml)|
|`--hide-progress`||`false`|turn off the progress bar displays|
|`--json`||`false`|output data in JSON format|
|`--log-level`||`info`|set the log level to debug|info|warn|error|
|`--no-stats`||`false`|disable anonymous usage statistics gathering|
|`--retain-progress`||`false`|retain the progress bar displays after completion|
