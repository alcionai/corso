---
title: corso backup create exchange
hide_title: true
---
## corso backup create exchange

Backup M365 Exchange service data

```bash
corso backup create exchange --user <userId or email> | '*' [flags]
```

### Examples

```bash
# Backup all Exchange data for Alice
corso backup create exchange --user alice@example.com

# Backup only Exchange contacts for Alice and Bob
corso backup create exchange --user alice@example.com,bob@example.com --data contacts

# Backup all Exchange data for all M365 users 
corso backup create exchange --user '*'
```

### Flags

|Flag|Short|Default|Help|
|:----|:-----|:-------|:----|
|`--user`||``|Backup Exchange data by user ID; accepts '*' to select all users|
|`--data`||``|Select one or more types of data to backup: email, contacts, or events|
|`--help`|`-h`|`false`|help for exchange|

### Global and inherited flags

|Flag|Short|Default|Help|
|:----|:-----|:-------|:----|
|`--config-file`||`$HOME/.corso.toml`|config file location (default is $HOME/.corso.toml)|
|`--hide-progress`||`false`|turn off the progress bar displays|
|`--json`||`false`|output data in JSON format|
|`--log-level`||`info`|set the log level to debug|info|warn|error|
|`--no-stats`||`false`|disable anonymous usage statistics gathering|
|`--retain-progress`||`false`|retain the progress bar displays after completion|
