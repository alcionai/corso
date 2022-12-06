---
title: corso backup details sharepoint
hide_title: true
---
## corso backup details sharepoint

Shows the details of a M365 SharePoint service backup

```bash
corso backup details sharepoint --backup <backupId> [flags]
```

### Examples

```bash
# Explore <site>'s files from backup 1234abcd-12ab-cd34-56de-1234abcd

corso backup details sharepoint --backup 1234abcd-12ab-cd34-56de-1234abcd --site <site_id>
```

### Flags

|Flag|Short|Default|Help|
|:----|:-----|:-------|:----|
|`--backup`|||ID of the backup to retrieve.|
|`--library`||``|Select backup details by Library name.|
|`--library-item`||``|Select backup details by library item name or ID.|
|`--help`|`-h`|`false`|help for sharepoint|

### Global and inherited flags

|Flag|Short|Default|Help|
|:----|:-----|:-------|:----|
|`--config-file`||`$HOME/.corso.toml`|config file location (default is $HOME/.corso.toml)|
|`--hide-progress`||`false`|turn off the progress bar displays|
|`--json`||`false`|output data in JSON format|
|`--log-level`||`info`|set the log level to debug|info|warn|error|
|`--no-stats`||`false`|disable anonymous usage statistics gathering|
|`--retain-progress`||`false`|retain the progress bar displays after completion|
