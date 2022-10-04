---
title: corso restore onedrive
hide_title: true
---
## corso restore onedrive

Restore M365 OneDrive service data

```bash
corso restore onedrive --backup <backupId> [flags]
```

### Flags

|Flag|Short|Default|Help|
|:----|:-----|:-------|:----|
|`--backup`|||ID of the backup to restore. <div class='required'>Required</div>|
|`--user`||``|Restore data by user ID; accepts * to select all users.|
|`--folder`||``|Restore items by OneDrive folder; defaults to root|
|`--file-name`||``|Restore items by OneDrive file name|
|`--file-created-after`|||Restore files created after this datetime|
|`--file-created-before`|||Restore files created before this datetime|
|`--file-modified-after`|||Restore files modified after this datetime|
|`--file-modified-before`|||Restore files modified before this datetime|
|`--help`|`-h`|`false`|help for onedrive|

### Global and inherited flags

|Flag|Short|Default|Help|
|:----|:-----|:-------|:----|
|`--config-file`||`$HOME/.corso.toml`|config file location (default is $HOME/.corso.toml)|
|`--json`||`false`|output data in JSON format|
|`--log-level`||`info`|set the log level to debug|info|warn|error|
|`--no-stats`||`false`|disable anonymous usage statistics gathering|
