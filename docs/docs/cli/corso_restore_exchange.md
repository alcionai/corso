---
title: corso restore exchange
hide_title: true
---
## corso restore exchange

Restore M365 Exchange service data

```bash
corso restore exchange [flags]
```

### Flags

|Flag|Short|Default|Help|
|:----|:-----|:-------|:----|
|`--backup`|||ID of the backup to restore|
|`--user`||`[]`|Restore all data by user ID; accepts * to select all users|
|`--email`||`[]`|Restore emails by ID; accepts * to select all emails|
|`--email-folder`||`[]`|Restore all emails by folder ID; accepts * to select all email folders|
|`--email-subject`|||Restore mail where the email subject lines contain this value|
|`--email-sender`|||Restore mail where the email sender matches this user id|
|`--email-received-after`|||Restore mail where the email was received after this datetime|
|`--email-received-before`|||Restore mail where the email was received before this datetime|
|`--event`||`[]`|Restore events by ID; accepts * to select all events|
|`--event-calendar`||`[]`|Restore events by calendar ID; accepts * to select all event calendars|
|`--event-subject`|||Restore events where the event subject contains this value|
|`--event-organizer`|||Restore events where the event organizer user id contains this value|
|`--event-recurs`|||Restore events if the event recurs. Use `--event-recurs false` to select non-recurring events|
|`--event-starts-after`|||Restore events where the event starts after this datetime|
|`--event-starts-before`|||Restore events where the event starts before this datetime|
|`--contact`||`[]`|Restore contacts by ID; accepts * to select all contacts|
|`--contact-folder`||`[]`|Restore all contacts within the folder ID; accepts * to select all contact folders|
|`--contact-name`|||Restore contacts where the contact name contains this value|
|`--help`|`-h`|`false`|help for exchange|

### Global and inherited flags

|Flag|Short|Default|Help|
|:----|:-----|:-------|:----|
|`--config-file`||`$HOME/.corso.toml`|config file location (default is $HOME/.corso.toml)|
|`--json`||`false`|output data in JSON format|
|`--log-level`||`info`|set the log level to debug|info|warn|error|
|`--no-stats`||`false`|disable anonymous usage statistics gathering|
