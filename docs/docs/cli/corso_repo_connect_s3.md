---
title: corso repo connect s3
hide_title: true
---
## corso repo connect s3

Ensures a connection to an existing S3 repository.

```bash
corso repo connect s3 [flags]
```

### Flags

|Flag|Short|Default|Help|
|:----|:-----|:-------|:----|
|`--bucket`|||Name of the S3 bucket (required).|
|`--endpoint`||`s3.amazonaws.com`|Server endpoint for S3 communication.|
|`--help`|`-h`|`false`|help for s3|
|`--prefix`|||Prefix applied to objects in the bucket.|

### Global and inherited flags

|Flag|Short|Default|Help|
|:----|:-----|:-------|:----|
|`--config-file`||`$HOME/.corso.toml`|config file location (default is $HOME/.corso.toml)|
|`--json`||`false`|output data in JSON format|
|`--log-level`||`info`|set the log level to debug|info|warn|error|
|`--no-stats`||`false`|disable anonymous usage statistics gathering|
