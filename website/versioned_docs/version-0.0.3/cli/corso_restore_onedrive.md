---
title: corso restore onedrive
hide_title: true
---
## corso restore onedrive

Restore M365 OneDrive service data

```bash
corso restore onedrive --backup <backupId> [flags]
```

### Examples

```bash
# Restore file with ID 98765abcdef
corso restore onedrive --backup 1234abcd-12ab-cd34-56de-1234abcd --file 98765abcdef

# Restore Alice's file named "FY2021 Planning.xlsx in "Documents/Finance Reports" from a specific backup
corso restore onedrive --backup 1234abcd-12ab-cd34-56de-1234abcd \
      --user alice@example.com --file "FY2021 Planning.xlsx" --folder "Documents/Finance Reports"

# Restore all files from Bob's folder that were created before 2020 when captured in a specific backup
corso restore onedrive --backup 1234abcd-12ab-cd34-56de-1234abcd 
      --user bob@example.com --folder "Documents/Finance Reports" --file-created-before 2020-01-01T00:00:00
```

### Flags

|Flag|Short|Default|Help|
|:----|:-----|:-------|:----|
|`--backup`|||ID of the backup to restore. <div class='required'>Required</div>|
|`--user`||``|Restore data by user ID; accepts '*' to select all users.|
|`--folder`||``|Restore items by OneDrive folder; defaults to root|
|`--file`||``|Restore items by file name or ID|
|`--file-created-after`|||Restore files created after this datetime|
|`--file-created-before`|||Restore files created before this datetime|
|`--file-modified-after`|||Restore files modified after this datetime|
|`--file-modified-before`|||Restore files modified before this datetime|
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
