# Known issues

Below is a list of known Corso issues and limitations:

* Corso supports Exchange (email, calendars, contacts), OneDrive (files) and SharePoint (Document Libraries) M365 data.
  Additional data types and services will be added in subsequent releases.

* Backups of Exchange email may not include changes to the read status of an email if no other changes
  to the email have been made since the previous backup.

* Restoration of Nested attachments within Exchange Mail or Calendars aren't yet supported.

* Folders and Calendars containing zero items or subfolders aren't included in the backup.

* Provides no guarantees the inclusion of data that's moved, added, or deleted
  from M365 while a backup creation is running.
  The next backup creation will correct any missing data.

* Sharing information of items in OneDrive/SharePoint using sharing links aren't backed up and restored.

* Permissions/Access given to a site group can't be restored.

* If a link share is created for an item with inheritance disabled
  (via the Graph API), the link shares restored in that item will
  not be inheritable by children.

* Link shares with password protection can't be restored.

* Teams conversation replies are only backed up if the parent message is available at the time of backup.

* Groups SharePoint files don't support Export. This limitation will be addressed in a follow-up release

* Teams messages don't support Restore due to limited Graph API support for message creation.

* Groups and Teams support is available in an early-access status, and may be subject to breaking changes.

* Restoring the data into a different Group from the one it was backed up from isn't currently supported

* CLI completions can't autocomplete multiple values for flags
