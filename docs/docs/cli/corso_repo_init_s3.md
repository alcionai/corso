---
title: corso repo init s3
hide_title: true
---
## corso repo init s3

Bootstraps a new S3 repository and connects it to your m356 account.

```bash
corso repo init s3 [flags]
```

### Flags

|Flag|Short|Default|Help|
|:----|:-----|:-------|:----|
|`--bucket`|||Name of the S3 bucket (required).|
|`--prefix`|||Prefix applied to objects in the bucket.|
|`--endpoint`||`s3.amazonaws.com`|Server endpoint for S3 communication.|
|`--help`|`-h`|`false`|help for s3|

### Global and inherited flags

|Flag|Short|Default|Help|
|:----|:-----|:-------|:----|
|`--config-file`||`$HOME/.corso.toml`|config file location (default is $HOME/.corso.toml)|
|`--json`||`false`|output data in JSON format|
|`--log-level`||`info`|set the log level to debug|info|warn|error|
|`--no-stats`||`false`|disable anonymous usage statistics gathering|
