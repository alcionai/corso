---
title: corso repo connect s3
hide_title: true
---
## corso repo connect s3

Ensures a connection to an existing S3 repository.

```bash
corso repo connect s3 --bucket <bucket> [flags]
```

### Flags

|Flag|Short|Default|Help|
|:----|:-----|:-------|:----|
|`--bucket`|||Name of S3 bucket for repo. <div class='required'>Required</div>|
|`--prefix`|||Repo prefix within bucket.|
|`--endpoint`||`s3.amazonaws.com`|S3 service endpoint.|
|`--help`|`-h`|`false`|help for s3|

### Global and inherited flags

|Flag|Short|Default|Help|
|:----|:-----|:-------|:----|
|`--config-file`||`$HOME/.corso.toml`|config file location (default is $HOME/.corso.toml)|
|`--json`||`false`|output data in JSON format|
|`--log-level`||`info`|set the log level to debug|info|warn|error|
|`--no-stats`||`false`|disable anonymous usage statistics gathering|
