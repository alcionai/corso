---
title: "corso repo connect s3"
hide_title: true
---
## corso repo connect s3

Connect to a S3 repository

### Synopsis

Ensures a connection to an existing S3 repository.

```
corso repo connect s3 [flags]
```

### Options

```
      --bucket string     Name of the S3 bucket (required).
      --endpoint string   Server endpoint for S3 communication. (default "s3.amazonaws.com")
  -h, --help              help for s3
      --prefix string     Prefix applied to objects in the bucket.
```

### Options inherited from parent commands

```
      --config-file string   config file (default is $HOME/.corso) (default "/home/runner/.corso.toml")
      --json                 output data in JSON format
      --log-level string     set the log level to debug|info|warn|error (default "info")
```

### SEE ALSO

* [corso repo connect](corso_repo_connect.md)	 - Connect to a repository.

