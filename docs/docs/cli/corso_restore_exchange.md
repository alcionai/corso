---
title: corso restore exchange
hide_title: true
---
## corso restore exchange

Restore M365 Exchange service data

```bash
corso restore exchange --backup <backupId> [flags]
```

### Flags

|Flag|Short|Default|Help|
|:----|:-----|:-------|:----|
|`--backup`|||ID of the backup to restore. <div class='required'>Required</div>|
|`--user`||``|Restore data by user ID; accepts * to select all users.|
|`--email`||``|Restore emails by ID; accepts * to select all emails.|
|`--email-folder`||``|Restore emails within a folder; accepts * to select all email folders.|
|`--email-subject`|||Restore emails with a subject containing this value.|
|`--email-sender`|||Restore emails from a specific sender.|
|`--email-received-after`|||Restore emails received after this datetime.|
|`--email-received-before`|||Restore emails received before this datetime.|
|`--event`||``|Restore events by event ID; accepts * to select all events.|
|`--event-calendar`||``|Restore events under a calendar; accepts * to select all event calendars.|
|`--event-subject`|||Restore events with a subject containing this value.|
|`--event-organizer`|||Restore events from a specific organizer.|
|`--event-recurs`|||Restore recurring events. Use `--event-recurs false` to restore non-recurring events.|
|`--event-starts-after`|||Restore events starting after this datetime.|
|`--event-starts-before`|||Restore events starting before this datetime.|
|`--contact`||``|Restore contacts by contact ID; accepts * to select all contacts.|
|`--contact-folder`||``|Restore contacts within a folder; accepts * to select all contact folders.|
|`--contact-name`|||Restore contacts whose contact name contains this value.|
|`--help`|`-h`|`false`|help for exchange|

### Global and inherited flags

|Flag|Short|Default|Help|
|:----|:-----|:-------|:----|
|`--config-file`||`$HOME/.corso.toml`|config file location (default is $HOME/.corso.toml)|
|`--json`||`false`|output data in JSON format|
|`--log-level`||`info`|set the log level to debug|info|warn|error|
|`--no-stats`||`false`|disable anonymous usage statistics gathering|
