# Known issues

Below is a list of known Corso issues and limitations:

* Only supports Exchange (email, calendars, contact) and OneDrive (files) M365 data. Additional
  data types and services will be added in subsequent releases.

* Restores are non-destructive to a dedicated restore folder in the original Exchange mailbox or OneDrive account.
  Advanced restore options such as in-place restore, or restore to a specific folder or to a different account aren't
  yet supported.

* Provides no guarantees about whether data moved, added, or deleted in M365
  while a backup is being created will be included in the running backup.
  Future backups run when the data isn't modified will include the data.
