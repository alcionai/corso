---
title: corso backup delete onedrive
hide_title: true
---
## corso backup delete onedrive

Delete backed-up M365 OneDrive service data

```bash
corso backup delete onedrive --backup <backupId> [flags]
```

### Examples

```bash
# Delete OneDrive backup with ID 1234abcd-12ab-cd34-56de-1234abcd
corso backup delete onedrive --backup 1234abcd-12ab-cd34-56de-1234abcd
```

### Flags

|Flag|Short|Default|Help|
|:----|:-----|:-------|:----|
|`--backup`|||ID of the backup to delete. <div class='required'>Required</div>|
|`--help`|`-h`|`false`|help for onedrive|

### Global and inherited flags

|Flag|Short|Default|Help|
|:----|:-----|:-------|:----|
|`--config-file`||`$HOME/.corso.toml`|config file location (default is $HOME/.corso.toml)|
|`--hide-progress`||`false`|turn off the progress bar displays|
|`--json`||`false`|output data in JSON format|
|`--log-level`||`info`|set the log level to debug|info|warn|error|
|`--no-stats`||`false`|disable anonymous usage statistics gathering|
|`--retain-progress`||`false`|retain the progress bar displays after completion|
