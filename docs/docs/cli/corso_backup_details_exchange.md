---
title: "corso backup details exchange"
hide_title: true
---
## corso backup details exchange

Shows the details of a M365 Exchange service backup

```
corso backup details exchange [flags]
```

### Options

```
      --backup string                  ID of the backup containing the details to be shown
      --contact strings                Select backup details by contact ID; accepts * to select all contacts
      --contact-folder strings         Select backup details by contact folder ID; accepts * to select all contact folders
      --contact-name string            Select backup details where the contact name contains this value
      --email strings                  Select backup details by emails ID; accepts * to select all emails
      --email-folder strings           Select backup details by email folder ID; accepts * to select all email folders
      --email-received-after string    Restore mail where the email was received after this datetime
      --email-received-before string   Restore mail where the email was received before this datetime
      --email-sender string            Restore mail where the email sender matches this user id
      --email-subject string           Restore mail where the email subject lines contain this value
      --event strings                  Select backup details by event ID; accepts * to select all events
      --event-calendar strings         Select backup details by event calendar ID; accepts * to select all events
      --event-organizer string         Select backup details where the event organizer user id contains this value
      --event-recurs string            Select backup details if the event recurs.  Use --event-recurs false to select where the event does not recur.
      --event-starts-after string      Select backup details where the event starts after this datetime
      --event-starts-before string     Select backup details where the event starts before this datetime
      --event-subject string           Select backup details where the event subject contains this value
  -h, --help                           help for exchange
      --user strings                   Select backup details by user ID; accepts * to select all users
```

### Options inherited from parent commands

```
      --config-file string   config file (default is $HOME/.corso) (default "/home/runner/.corso.toml")
      --json                 output data in JSON format
      --log-level string     set the log level to debug|info|warn|error (default "info")
```

### SEE ALSO

* [corso backup details](corso_backup_details.md)	 - Shows the details of a backup for a service

