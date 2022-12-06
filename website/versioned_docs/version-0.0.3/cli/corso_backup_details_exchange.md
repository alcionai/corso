---
title: corso backup details exchange
hide_title: true
---
## corso backup details exchange

Shows the details of a M365 Exchange service backup

```bash
corso backup details exchange --backup <backupId> [flags]
```

### Examples

```bash
# Explore Alice's items in backup 1234abcd-12ab-cd34-56de-1234abcd 
corso backup details exchange --backup 1234abcd-12ab-cd34-56de-1234abcd --user alice@example.com

# Explore Alice's emails with subject containing "Hello world" in folder "Inbox" from a specific backup 
corso backup details exchange --backup 1234abcd-12ab-cd34-56de-1234abcd \
      --user alice@example.com --email-subject "Hello world" --email-folder Inbox

# Explore Bobs's events occurring after start of 2022 from a specific backup
corso backup details exchange --backup 1234abcd-12ab-cd34-56de-1234abcd \
      --user bob@example.com --event-starts-after 2022-01-01T00:00:00

# Explore Alice's contacts with name containing Andy from a specific backup
corso backup details exchange --backup 1234abcd-12ab-cd34-56de-1234abcd \
      --user alice@example.com --contact-name Andy
```

### Flags

|Flag|Short|Default|Help|
|:----|:-----|:-------|:----|
|`--backup`|||ID of the backup to explore. <div class='required'>Required</div>|
|`--user`||``|Select backup details by user ID; accepts '*' to select all users.|
|`--email`||``|Select backup details for emails by email ID; accepts '*' to select all emails.|
|`--email-folder`||``|Select backup details for emails within a folder; accepts '*' to select all email folders.|
|`--email-subject`|||Select backup details for emails with a subject containing this value.|
|`--email-sender`|||Select backup details for emails from a specific sender.|
|`--email-received-after`|||Select backup details for emails received after this datetime.|
|`--email-received-before`|||Select backup details for emails received before this datetime.|
|`--event`||``|Select backup details for events by event ID; accepts '*' to select all events.|
|`--event-calendar`||``|Select backup details for events under a calendar; accepts '*' to select all events.|
|`--event-subject`|||Select backup details for events with a subject containing this value.|
|`--event-organizer`|||Select backup details for events from a specific organizer.|
|`--event-recurs`|||Select backup details for recurring events. Use `--event-recurs false` to select non-recurring events.|
|`--event-starts-after`|||Select backup details for events starting after this datetime.|
|`--event-starts-before`|||Select backup details for events starting before this datetime.|
|`--contact`||``|Select backup details for contacts by contact ID; accepts '*' to select all contacts.|
|`--contact-folder`||``|Select backup details for contacts within a folder; accepts '*' to select all contact folders.|
|`--contact-name`|||Select backup details for contacts whose contact name contains this value.|
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
