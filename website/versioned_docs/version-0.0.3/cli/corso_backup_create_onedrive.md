---
title: corso backup create onedrive
hide_title: true
---
## corso backup create onedrive

Backup M365 OneDrive service data

```bash
corso backup create onedrive --user <userId or email> | '*' [flags]
```

### Examples

```bash
# Backup OneDrive data for Alice
corso backup create onedrive --user alice@example.com

# Backup OneDrive for Alice and Bob
corso backup create onedrive --user alice@example.com,bob@example.com

# Backup all OneDrive data for all M365 users 
corso backup create onedrive --user '*'
```

### Flags

|Flag|Short|Default|Help|
|:----|:-----|:-------|:----|
|`--user`||``|Backup OneDrive data by user ID; accepts '*' to select all users. <div class='required'>Required</div>|
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
