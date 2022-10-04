---
title: corso backup create exchange
hide_title: true
---
## corso backup create exchange

Backup M365 Exchange service data

```bash
corso backup create exchange --all | --user <userId or email> [flags]
```

### Flags

|Flag|Short|Default|Help|
|:----|:-----|:-------|:----|
|`--all`||`false`|Backup all Exchange data for all users|
|`--user`||``|Backup Exchange data by user ID; accepts * to select all users|
|`--data`||``|Select one or more types of data to backup: email, contacts, or events|
|`--help`|`-h`|`false`|help for exchange|

### Global and inherited flags

|Flag|Short|Default|Help|
|:----|:-----|:-------|:----|
|`--config-file`||`$HOME/.corso.toml`|config file location (default is $HOME/.corso.toml)|
|`--json`||`false`|output data in JSON format|
|`--log-level`||`info`|set the log level to debug|info|warn|error|
|`--no-stats`||`false`|disable anonymous usage statistics gathering|
