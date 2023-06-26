# Known issues

Below is a list of known Corso issues and limitations:

* Corso supports Exchange (email, calendars, contacts), OneDrive (files) and SharePoint (Document Libraries) M365 data.
  Additional data types and services will be added in subsequent releases.

* Backups of Exchange email may not include changes to the read status of an email if no other changes
  to the email have been made since the previous backup.

* Restores are non-destructive to a dedicated restore folder in the original Exchange mailbox or OneDrive account.
  Advanced restore options such as in-place restore, or restore to a specific folder or to a different account aren't
  yet supported.

* Restoration of Nested attachments within Exchange Mail or Calendars aren't yet supported.

* Folders and Calendars containing zero items or subfolders aren't included in the backup.

* Provides no guarantees about whether data moved, added, or deleted in M365
  while a backup is being created will be included in the running backup.
  Future backups run when the data isn't modified will include the data.

* Exchange Calender Event instance exceptions (changes to a single event within a recurring series) aren't
included in backup and restore.

* SharePoint document library data can't be restored after the library has been deleted.

* Sharing information of items in OneDrive/SharePoint using sharing links aren't backed up and restored.

* Partial inheritance of link shares from a parent folder can't be restored.
