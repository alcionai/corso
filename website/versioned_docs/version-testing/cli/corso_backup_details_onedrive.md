---
title: corso backup details onedrive
hide_title: true
---
## corso backup details onedrive

Shows the details of a M365 OneDrive service backup

```bash
corso backup details onedrive --backup <backupId> [flags]
```

### Examples

```bash
# Explore Alice's files from backup 1234abcd-12ab-cd34-56de-1234abcd 
corso backup details onedrive --backup 1234abcd-12ab-cd34-56de-1234abcd --user alice@example.com

# Explore Alice or Bob's files with name containing "Fiscal 22" in folder "Reports"
corso backup details onedrive --backup 1234abcd-12ab-cd34-56de-1234abcd \
      --user alice@example.com,bob@example.com  --file-name "Fiscal 22" --folder "Reports"

# Explore Alice's files created before end of 2015 from a specific backup
corso backup details onedrive --backup 1234abcd-12ab-cd34-56de-1234abcd \
      --user alice@example.com --file-created-before 2015-01-01T00:00:00
```

### Flags

|Flag|Short|Default|Help|
|:----|:-----|:-------|:----|
|`--backup`|||ID of the backup to explore. <div class='required'>Required</div>|
|`--folder`||``|Select backup details by OneDrive folder; defaults to root.|
|`--file`||``|Select backup details by file name or ID.|
|`--file-created-after`|||Select backup details for files created after this datetime.|
|`--file-created-before`|||Select backup details for files created before this datetime.|
|`--file-modified-after`|||Select backup details for files modified after this datetime.|
|`--file-modified-before`|||Select backup details for files modified before this datetime.|
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
