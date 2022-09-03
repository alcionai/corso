---
title: "corso restore exchange"
hide_title: true
---
## corso restore exchange

Restore M365 Exchange service data

```
corso restore exchange [flags]
```

### Options

```
      --backup string                  ID of the backup to restore
      --contact strings                Restore contacts by ID; accepts * to select all contacts
      --contact-folder strings         Restore all contacts within the folder ID; accepts * to select all contact folders
      --contact-name string            Restore contacts where the contact name contains this value
      --email strings                  Restore emails by ID; accepts * to select all emails
      --email-folder strings           Restore all emails by folder ID; accepts * to select all email folders
      --email-received-after string    Restore mail where the email was received after this datetime
      --email-received-before string   Restore mail where the email was received before this datetime
      --email-sender string            Restore mail where the email sender matches this user id
      --email-subject string           Restore mail where the email subject lines contain this value
      --event strings                  Restore events by ID; accepts * to select all events
      --event-calendar strings         Restore events by calendar ID; accepts * to select all event calendars
      --event-organizer string         Restore events where the event organizer user id contains this value
      --event-recurs string            Restore events if the event recurs.  Use --event-recurs false to select where the event does not recur.
      --event-starts-after string      Restore events where the event starts after this datetime
      --event-starts-before string     Restore events where the event starts before this datetime
      --event-subject string           Restore events where the event subject contains this value
  -h, --help                           help for exchange
      --user strings                   Restore all data by user ID; accepts * to select all users
```

### Options inherited from parent commands

```
      --config-file string   config file (default is $HOME/.corso) (default "/home/runner/.corso.toml")
      --json                 output data in JSON format
      --log-level string     set the log level to debug|info|warn|error (default "info")
```

### SEE ALSO

* [corso restore](corso_restore.md)	 - Restore your service data

