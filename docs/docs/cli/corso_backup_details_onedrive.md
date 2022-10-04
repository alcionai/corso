---
title: corso backup details onedrive
hide_title: true
---
## corso backup details onedrive

Shows the details of a M365 OneDrive service backup

```bash
corso backup details onedrive --backup <backupId> [flags]
```

### Flags

|Flag|Short|Default|Help|
|:----|:-----|:-------|:----|
|`--backup`|||ID of the backup to explore. <div class='required'>Required</div>|
|`--folder`||``|Select backup details by OneDrive folder; defaults to root|
|`--file-name`||``|Select backup details by OneDrive file name|
|`--file-created-after`|||Select files created after this datetime|
|`--file-created-before`|||Select files created before this datetime|
|`--file-modified-after`|||Select files modified after this datetime|
|`--file-modified-before`|||Select files modified before this datetime|
|`--help`|`-h`|`false`|help for onedrive|

### Global and inherited flags

|Flag|Short|Default|Help|
|:----|:-----|:-------|:----|
|`--config-file`||`$HOME/.corso.toml`|config file location (default is $HOME/.corso.toml)|
|`--json`||`false`|output data in JSON format|
|`--log-level`||`info`|set the log level to debug|info|warn|error|
|`--no-stats`||`false`|disable anonymous usage statistics gathering|
