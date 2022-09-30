---
title: corso backup details exchange
hide_title: true
---
## corso backup details exchange

Shows the details of a M365 Exchange service backup

```bash
corso backup details exchange [flags]
```

### Flags

|Flag|Short|Default|Help|
|:----|:-----|:-------|:----|
|`--backup`|||ID of the backup containing the details to be shown|
|`--contact`||`[]`|Select backup details by contact ID; accepts * to select all contacts|
|`--contact-folder`||`[]`|Select backup details by contact folder ID; accepts * to select all contact folders|
|`--contact-name`|||Select backup details where the contact name contains this value|
|`--email`||`[]`|Select backup details by emails ID; accepts * to select all emails|
|`--email-folder`||`[]`|Select backup details by email folder ID; accepts * to select all email folders|
|`--email-received-after`|||Restore mail where the email was received after this datetime|
|`--email-received-before`|||Restore mail where the email was received before this datetime|
|`--email-sender`|||Restore mail where the email sender matches this user id|
|`--email-subject`|||Restore mail where the email subject lines contain this value|
|`--event`||`[]`|Select backup details by event ID; accepts * to select all events|
|`--event-calendar`||`[]`|Select backup details by event calendar ID; accepts * to select all events|
|`--event-organizer`|||Select backup details where the event organizer user id contains this value|
|`--event-recurs`|||Select backup details if the event recurs. Use `--event-recurs` false to select non-recurring events|
|`--event-starts-after`|||Select backup details where the event starts after this datetime|
|`--event-starts-before`|||Select backup details where the event starts before this datetime|
|`--event-subject`|||Select backup details where the event subject contains this value|
|`--help`|`-h`|`false`|help for exchange|
|`--user`||`[]`|Select backup details by user ID; accepts * to select all users|

### Global and inherited flags

|Flag|Short|Default|Help|
|:----|:-----|:-------|:----|
|`--config-file`||`$HOME/.corso.toml`|config file location (default is $HOME/.corso.toml)|
|`--json`||`false`|output data in JSON format|
|`--log-level`||`info`|set the log level to debug|info|warn|error|
|`--no-stats`||`false`|disable anonymous usage statistics gathering|
