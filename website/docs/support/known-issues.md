# Known issues and limitations

## Issues

The following are the issues that we're currently aware of within the project.

## General

- Corso supports Exchange (email, calendars, contacts), OneDrive (files) and SharePoint (Document Libraries) M365 data.
  Additional data types and services will be added in subsequent releases.

- Provides no guarantees the inclusion of data that's moved, added, or deleted
  from M365 while a backup creation is running.
  The next backup creation will correct any missing data.

### Drive items (OneDrive/SharePoint)

- Permissions/Access given to a site group can't be restored.

- If a link share is created for an item with inheritance disabled
  (via the Graph API), the link shares restored in that item will
  not be inheritable by children.

### Exchange

- Backups of Exchange email may not include changes to the read status of an email if no other changes
  to the email have been made since the previous backup.

- Restoration of Nested attachments within Exchange Mail or Calendars aren't yet supported.

- Calendars containing zero items or subfolders aren't included in the backup.

- Backing up a group mailbox item may fail if it has a large number of attachments (500+).

### Teams

- Teams conversation replies are only backed up if the parent message is available at the time of backup.

- Groups and Teams support is available in an early-access status, and may be subject to breaking changes.

- Restoring the data into a different Group from the one it was backed up from isn't currently supported.

## Limitations

Following are unexpected behaviors or inherent limitations of the project.

<!-- markdownlint-disable-next-line no-duplicate-heading -->
### Drive items (OneDrive/SharePoint)

- Link shares with password protection can't be restored.

- Restored link shares always generate different links from the original.

- Anonymous link shares (link shares which aren't associated with any user) aren't restored.

<!-- markdownlint-disable-next-line no-duplicate-heading -->
### Exchange

- Exports of multipart emails containing both HTML and text versions will only produce the HTML version of the email.

<!-- markdownlint-disable-next-line no-duplicate-heading -->
### Teams

- Teams messages don't support restore due to limited Graph API support for message creation.
