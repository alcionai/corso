# Known issues

Below is a list of known Corso issues and limitations:

* Corso supports Exchange (email, calendars, contacts), OneDrive (files) and SharePoint (Document Libraries) M365 data.
  Additional data types and services will be added in subsequent releases.

* Backups of Exchange email may not include changes to the read status of an email if no other changes
  to the email have been made since the previous backup.

* Restoration of Nested attachments within Exchange Mail or Calendars aren't yet supported.

* Folders and Calendars containing zero items or subfolders aren't included in the backup.

* Provides no guarantees the inclusion of data that is moved, added, or deleted
  from M365 while a backup creation is running.
  The next backup creation will correct any missing data.

* SharePoint document library data can't be restored after the library has been deleted.

* Sharing information of items in OneDrive/SharePoint using sharing links aren't backed up and restored.

* Permissions/Access given to a site group can't be restored.

* If a link share is created for an item with inheritance disabled
  (via the Graph API), the link shares restored in that item will
  not be inheritable by children.

* Link shares with password protection can't be restored.
