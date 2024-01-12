# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased] (beta)

### Fixed
- Retry transient 400 "invalidRequest" errors during onedrive & sharepoint backup.
- Backup attachments associated with group mailbox items.
- Groups and Teams backups no longer fail when a resource has no display name.

### Changed
- When running `backup details` on an empty backup returns a more helpful error message.

### Known issues
- Backing up a group mailbox item may fail if it has a very large number of attachments (500+).

## [v0.18.0] (beta) - 2024-01-02

### Fixed
- Handle the case where an email cannot be retrieved from Exchange due to an `ErrorInvalidRecipients` error. In
this case, Corso will skip over the item but report this in the backup summary.
- Fix `ErrorItemNotFound` errors when restoring emails with multiple attachments.
- Avoid Graph SDK `Requests must contain extension changes exclusively.` errors by removing server-populated field from restored event items.
- Improve Group mailbox(conversations) backup performance by only downloading new items or items with modified content.
- Handle cases where Exchange backup stored invalid JSON blobs if there were special characters in the user content. These would result in errors during restore.

### Known issues
- Restoring OneDrive, SharePoint, or Teams & Groups items shared with external users while the tenant or site is configured to not allow sharing with external users will not restore permissions.

### Added
- Contacts can now be exported from Exchange backups as .vcf files

## [v0.17.0] (beta) - 2023-12-11

### Changed
- Memory optimizations for large scale OneDrive and Sharepoint backups.

### Fixed
- Resolved a possible deadlock when backing up Teams Channel Messages.
- Fixed an attachment download failure(ErrorTooManyObjectsOpened) during exchange backup.

## [v0.16.0] (beta) - 2023-11-28

### Added
- Export support for emails in exchange backups as `.eml` files.
- More colorful and informational cli display.

### Changed
- The file extension in Teams messages exports has switched to json to match the content type.
- SDK consumption of the /services/m365 package has shifted from independent functions to a client-based api.
- SDK consumers can now configure the /services/m365 graph api client configuration when constructing a new m365 client.
- Dynamic api rate limiting allows small-scale Exchange backups to complete more quickly.
- Kopia's local config files now uses unique filenames that match Corso configurations.  This can protect concurrent Corso operations from mistakenly clobbering storage configs during runtime.

### Fixed
- Handle OneDrive folders being deleted and recreated midway through a backup.
- Automatically re-run a full delta query on incremental if the prior backup is found to have malformed prior-state information.
- Retry drive item permission downloads during long-running backups after the jwt token expires and refreshes.
- Retry item downloads during connection timeouts.

## [v0.15.0] (beta) - 2023-10-31

### Added
- Added `corso repo update-passphrase` command to update the passphrase of an existing Corso repository
- Added Subject and Message preview to channel messages detail entries

### Fixed
- SharePoint backup would fail if any site had an empty display name
- Fix a bug with exports hanging post completion
- Handle 503 errors in nested OneDrive packages

### Changed
- Item Details formatting in Groups and Teams backups

## [v0.14.2] (beta) - 2023-10-17

### Added
- Skips graph calls for expired item download URLs.
- Export operation now shows the stats at the end of the run

### Fixed
- Catch and report cases where a protected resource is locked out of access.  SDK consumers have a new errs sentinel that allows them to check for this case.
- Fix a case where missing item LastModifiedTimes could cause incremental backups to fail.
- Email size metadata was incorrectly set to the size of the last attachment.  Emails will now correctly report the size of the mail content plus the size of all attachments.
- Improves the filtering capabilities for Groups restore and backup
- Improve check to skip OneNote files that cannot be downloaded.
- Fix Groups backup for non Team groups

### Changed
- Groups restore now expects the site whose backup we should restore

## [v0.14.0] (beta) - 2023-10-09

### Added
- Enables local or network-attached storage for Corso repositories.
- Reduce backup runtime for OneDrive and SharePoint incremental backups that have no file changes.
- Increase Exchange backup performance by lazily fetching data only for items whose content changed.
- Added `--backups` flag to delete multiple backups in `corso backup delete` command.
- Backup now includes all sites that belongs to a team, not just the root site.

### Fixed
- Teams Channels that cannot support delta tokens (those without messages) fall back to non-delta enumeration and no longer fail a backup.

### Known issues
- Restoring the data into a different Group from the one it was backed up from is not currently supported

### Other
- Groups and Teams service support is still in feature preview

## [v0.13.0] (beta) - 2023-09-18

### Added
- Groups and Teams service support available as a feature preview!  Channel messages and Files are now available for backup and restore in the CLI: `corso backup create groups --group '*'`
  - The cli commands for "groups" and "teams" can be used interchangeably, and will operate on the same backup data.
  - New permissions are required to backup Channel messages.  See the [Corso Documentation](https://corsobackup.io/docs/setup/m365-access/#configure-required-permissions) for complete details.
  Even though Channel message restoration is not available, message write permissions are included to cover future integration.
  - This is a feature preview, and may be subject to breaking changes based on feedback and testing.

### Changed
- Switched to Go 1.21
- SharePoint exported libraries are now exported with a `Libraries` prefix.

### Fixed
- Contacts backups no longer slices root-folder data if outlook is set to languages other than english.
- Failed backups if the --disable-incrementals flag was passed when there was a valid merge base under some conditions.

## [v0.12.0] (beta) - 2023-08-29

### Added
- Added `export` command to export data from OneDrive and SharePoint backups as individual files or as a single zip file.
- Restore commands now accept an optional resource override with the `--to-resource` flag.  This allows restores to recreate backup data within different mailboxes, sites, and users.  
- Improve `--mask-sensitive-data` logging mode.
- Reliability: Handle connection cancellation and resets observed when backing up or restoring large data sets.
- Reliability: Recover from Graph SDK panics when the Graph API returns incomplete responses.
- Performance: Improve backup delete performance by batching multiple storage operations into a single operation.

### Fixed
- SharePoint document libraries deleted after the last backup can now be restored.
- Restore requires the protected resource to have access to the service being restored.
- SharePoint data from multiple document libraries are not merged in exports
- `corso backup delete` was not removing the backup details data associated with that snapshot
- Fix OneDrive restores could fail with a concurrent map write error
- Fix backup list displaying backups that had errors
- Fix OneDrive backup could fail if item was deleted during backup
- Exchange backups would fail attempting to use delta tokens even if the user was over quota


## [v0.11.1] (beta) - 2023-07-20

### Fixed
- Allow repo connect to succeed when a `corso.toml` file was not provided but configuration is specified using environment variables and flags.

## [v0.11.0] (beta) - 2023-07-18

### Added
- Drive items backup and restore link shares
- Restore commands now accept an optional top-level restore destination with the `--destination` flag.  Setting the destination to '/' will restore items back into their original location.  
- Restore commands can specify item collision behavior.  Options are Skip (default), Replace, and Copy.
- Introduced repository maintenance commands to help optimize the repository as well as unreferenced data.

### Fixed
- Return a ServiceNotEnabled error when a tenant has no active SharePoint license.
- Added retries for http/2 stream connection failures when downloading large item content.
- SharePoint document libraries that were deleted after the last backup can now be restored.

### Known issues
- If a link share is created for an item with inheritance disabled
  (via the Graph API), the link shares restored in that item will
  not be inheritable by children
- Link shares with password protection can't be restored

## [v0.10.0] (beta) - 2023-06-26

### Added
- Exceptions and cancellations for recurring events are now backed up and restored
- Introduced a URL cache for OneDrive that helps reduce Graph API calls for long running (>1hr) backups
- Improve incremental backup behavior by leveraging information from incomplete backups
- Improve restore performance and memory use for Exchange and OneDrive

### Fixed
- Handle OLE conversion errors when trying to fetch attachments
- Fix uploading large attachments for emails and calendar
- Fixed high memory use in OneDrive backup related to logging
- Return a ServiceNotEnabled error when a tenant has no active SharePoint license.

### Changed
- Switched to Go 1.20
  
## [v0.9.0] (beta) - 2023-06-05

### Added
- Added ProtectedResourceName to the backup list json output.  ProtectedResourceName holds either a UPN or a WebURL, depending on the resource type.
- Rework base selection logic for incremental backups so it's more likely to find a valid base.
- Improve OneDrive restore performance by paralleling item restores

### Fixed
- Fix Exchange folder cache population error when parent folder isn't found.
- Fix Exchange backup issue caused by incorrect json serialization
- Fix issues with details model containing duplicate entry for api consumers

### Changed
- Do not display all the items that we restored at the end if there are more than 15. You can override this with `--verbose`.

## [v0.8.0] (beta) - 2023-05-15

### Added
- Released the --mask-sensitive-data flag, which will automatically obscure private data in logs.
- Added `--disable-delta` flag to disable delta based backups for Exchange
- Permission support for SharePoint libraries.

### Fixed
- Graph requests now automatically retry in case of a Bad Gateway or Gateway Timeout.
- POST Retries following certain status codes (500, 502, 504) will re-use the post body instead of retrying with a no-content request.
- Fix nil pointer exception when running an incremental backup on SharePoint where the base backup used an older index data format.
- --user and --mailbox flags have been removed from CLI examples for details and restore commands (they were already not supported, this only updates the docs).
- Improve restore time on large restores by optimizing how items are loaded from the remote repository.
- Remove exchange item filtering based on m365 item ID via the CLI.
- OneDrive backups no longer include a user's non-default drives.
- OneDrive and SharePoint file downloads will properly redirect from 3xx responses.
- Refined oneDrive rate limiter controls to reduce throttling errors.
- Fix handling of duplicate folders at the same hierarchy level in Exchange. Duplicate folders will be merged during restore operations.
- Fix backup for mailboxes that has used up all their storage quota
- Restored folders no longer appear in the Restore results. Only restored items will be displayed.

### Known Issues
- Restore operations will merge duplicate Exchange folders at the same hierarchy level into a single folder.
- Sharepoint SiteGroup permissions are not restored.
- SharePoint document library data can't be restored after the library has been deleted.

## [v0.7.0] (beta) - 2023-05-02

### Added
- Permissions backup for OneDrive is now out of experimental (By default, only newly backed up items will have their permissions backed up. You will have to run a full backup to ensure all items have their permissions backed up.)
- LocationRef is now populated for all services and data types. It should be used in place of RepoRef if a location for an item is required.
- User selection for Exchange and OneDrive can accept either a user PrincipalName or the user's canonical ID.
- Add path information to items that were skipped during backup because they were flagged as malware.

### Fixed
- Fixed permissions restore in latest backup version.
- Incremental OneDrive backups could panic if the delta token expired and a folder was seen and deleted in the course of item enumeration for the backup.
- Incorrectly moving subfolder hierarchy from a deleted folder to a new folder at the same path during OneDrive incremental backup.
- Handle calendar events with no body.
- Items not being deleted if they were created and deleted during item enumeration of a OneDrive backup.
- Enable compression for all data uploaded by kopia.
- SharePoint --folder selectors correctly return items.
- Fix Exchange cli args for filtering items
- Skip OneNote items bigger than 2GB (Graph API prevents us from downloading them)
- ParentPath of json output for Exchange calendar now shows names instead of IDs.
- Fixed failure when downloading huge amount of attachments
- Graph API requests that return an ECONNRESET error are now retried.
- Fixed edge case in incremental backups where moving a subfolder, deleting and recreating the subfolder's original parent folder, and moving the subfolder back to where it started would skip backing up unchanged items in the subfolder.
- SharePoint now correctly displays site urls on `backup list`, instead of the site id.
- Drives with a directory containing a folder named 'folder' will now restore without error.
- The CORSO_LOG_FILE env is appropriately utilized if no --log-file flag is provided.
- Fixed Exchange events progress output to show calendar names instead of IDs.
- Fixed reporting no items match if restoring or listing details on an older Exchange backup and filtering by folder.
- Fix backup for mailboxes that has used up all their storage quota

### Known Issues
- Restoring a OneDrive or SharePoint file with the same name as a file with that name as its M365 ID may restore both items.
- Exchange event restores will display calendar IDs instead of names in the progress output.

## [v0.6.1] (beta) - 2023-03-21

### Added
- Sharepoint library (document files) support: backup, list, details, and restore.
- OneDrive item downloads that return 404 during backup (normally due to external deletion while Corso processes) are now skipped instead of quietly dropped.  These items will appear in the skipped list alongside other skipped cases such as malware detection.
- Listing a single backup by id will also list the skipped and failed items that occurred during the backup.  These can be filtered out with the flags `--failed-items hide`, `--skipped-items hide`, and `--recovered-errors hide`.
- Enable incremental backups for OneDrive if permissions aren't being backed up.
- Show progressbar while files for user are enumerated
- Hidden flag to control parallelism for fetching Exchange items (`--fetch-parallelism`). May help reduce `ApplicationThrottled` errors but will slow down backup.

### Fixed
- Fix repo connect not working without a config file
- Fix item re-download on expired links silently being skipped
- Improved permissions backup and restore for OneDrive

### Known Issues
- Owner (Full control) or empty (Restricted View) roles cannot be restored for OneDrive
- OneDrive will not do an incremental backup if permissions are being backed up.
- SharePoint --folder selection in details and restore always return "no items match the specified selectors".
- Event instance exceptions (ie: changes to a single event within a recurring series) are not backed up.

## [v0.5.0] (beta) - 2023-03-13

### Added
- Show owner information when doing backup list in json format
- Permissions for groups can now be backed up and restored
- Onedrive files that are flagged as malware get skipped during backup.  Skipped files are listed in the backup results as part of the status, including a reference to their categorization, eg: "Completed (0 errors, 1 skipped: 1 malware)".

### Fixed
- Corso-generated .meta files and permissions no longer appear in the backup details.
- Panic and recovery if a user didn't exist in the tenant.

### Known Issues
- Folders and Calendars containing zero items or subfolders are not included in the backup.
- OneDrive files ending in `.meta` or `.dirmeta` are omitted from details and restores.
- Backups generated prior to this version will show `0 errors` when listed, even if error count was originally non-zero.

## [v0.4.0] (beta) - 2023-02-20

### Fixed
- Support for item.Attachment:Mail restore
- Errors from duplicate names in Exchange Calendars
- Resolved an issue where progress bar displays could fail to exit, causing unbounded CPU consumption.
- Fix Corso panic within Docker images
- Debugging with the CORSO_URL_LOGGING env variable no longer causes accidental request failures.
- Don't discover all users when backing up each user in a multi-user backup

### Changed
- When using Restore and Details on Exchange Calendars, the `--event-calendar` flag can now identify calendars by either a Display Name or a Microsoft 365 ID.
- Exchange Calendars storage entries now construct their paths using container IDs instead of display names.  This fixes cases where duplicate display names caused system failures.

### Known Issues
- Nested attachments are currently not restored due to an [issue](https://github.com/microsoft/kiota-serialization-json-go/issues/61) discovered in the Graph APIs
- Breaking changes to Exchange Calendar backups.
- The debugging env variable CORSO_URL_LOGGING causes exchange get requests to fail.
- Onedrive files that are flagged as Malware consistently fail during backup.

## [v0.3.0] (alpha) - 2023-02-07

### Added

- Document Corso's fault-tolerance and restartability features
- Add retries on timeouts and status code 500 for Exchange
- Increase page size preference for delta requests for Exchange to reduce number of roundtrips
- OneDrive file/folder permissions can now be backed up and restored
- Add `--restore-permissions` flag to toggle restoration of OneDrive permissions
- Add versions to backups so that we can understand/handle older backup formats

### Fixed

- Added additional backoff-retry to all OneDrive queries.
- Users with `null` userType values are no longer excluded from user queries.
- Fix bug when backing up a calendar that has the same name as the default calendar

### Known Issues

- When the same user has permissions to a file and the containing
  folder, we only restore folder level permissions for the user and no
  separate file only permission is restored.
- Link shares are not restored

## [v0.2.0] (alpha) - 2023-01-29

### Fixed

- Check if the user specified for an exchange backup operation has a mailbox.

### Changed
- Item.Attachments are disabled from being restored for the patching of ([#2353](https://github.com/alcionai/corso/issues/2353))
- BetaClient introduced. Enables Corso to be able to interact with SharePoint Page objects. Package located `/internal/connector/graph/betasdk`
- Handle case where user's drive has not been initialized
- Inline attachments (e.g. copy/paste ) are discovered and backed up correctly ([#2163](https://github.com/alcionai/corso/issues/2163))
- Guest and External users (for cloud accounts) and non-on-premise users (for systems that use on-prem AD syncs) are now excluded from backup and restore operations.
- Remove the M365 license guid check in OneDrive backup which wasn't reliable.
- Reduced extra socket consumption while downloading multiple drive files.
- Extended timeout boundaries for exchange attachment downloads, reducing risk of cancellation on large files.
- Identify all drives associated with a user or SharePoint site instead of just the results on the first page returned by Graph API.

## [v0.1.0] (alpha) - 2023-01-13

### Added

- Folder entries in backup details now indicate whether an item in the hierarchy was updated
- Incremental backup support for exchange is now enabled by default.

### Changed

- The selectors Reduce() process will only include details that match the DiscreteOwner, if one is specified.
- New selector constructors will automatically set the DiscreteOwner if given a single-item slice.
- Write logs to disk by default ([#2082](https://github.com/alcionai/corso/pull/2082))

### Fixed

- Issue where repository connect progress bar was clobbering backup/restore operation output.
- Issue where a `backup create exchange` produced one backup record per data type.
- Specifying multiple users in a onedrive backup (ex: `--user a,b,c`) now properly delimits the input along the commas.
- Updated the list of M365 SKUs used to check if a user has a OneDrive license.

### Known Issues

- `backup list` will not display a resource owner for backups created prior to this release.

## [v0.0.4] (alpha) - 2022-12-23

### Added

- Incremental backup support for Exchange ([#1777](https://github.com/alcionai/corso/issues/1777)). This is currently enabled by specifying the `--enable-incrementals`  
  with the `backup create` command. This functionality will be enabled by default in an upcoming release.
- Folder entries in backup details now include size and modified time for the hierarchy ([#1896](https://github.com/alcionai/corso/issues/1896))

### Changed

- **Breaking Change**:
  Changed how backup details are stored in the repository to
  improve memory usage ([#1735](https://github.com/alcionai/corso/issues/1735))
- Improve OneDrive backup speed ([#1842](https://github.com/alcionai/corso/issues/1842))
- Upgrade MS Graph SDK libraries ([#1856](https://github.com/alcionai/corso/issues/1856))
- Docs: Add Algolia docsearch to Corso docs ([#1844](https://github.com/alcionai/corso/pull/1844))
- Add an `updated` flag to backup details ([#1813](https://github.com/alcionai/corso/pull/1813))
- Docs: Speed up Windows Powershell download ([#1798](https://github.com/alcionai/corso/pull/1798))
- Switch to Go 1.19 ([#1632](https://github.com/alcionai/corso/pull/1632))

### Fixed

- Fixed retry logic in the Graph SDK that would result in an `400 Empty Payload` error when the request was retried ([1778](https://github.com/alcionai/corso/issues/1778))([msgraph-sdk-go #341](https://github.com/microsoftgraph/msgraph-sdk-go/issues/341))
- Don't error out if a folder was deleted during an exchange backup operation ([#1849](https://github.com/alcionai/corso/pull/1849))
- Docs: Fix CLI auto-generated docs headers ([#1845](https://github.com/alcionai/corso/pull/1845))

## [v0.0.3] (alpha) - 2022-12-05

### Added

- Display backup size in backup list command (#1648) from [meain](https://github.com/meain)
- Improve OneDrive backup performance (#1607) from [meain](https://github.com/meain)
- Improve Exchange backup performance (#1608) from [meain](https://github.com/meain)
- Add flag to retain all progress bars (#1582) from [ryanfkeepers](https://github.com/ryanfkeepers)
- Fix resource owner display on backup list (#1580) from [ryanfkeepers](https://github.com/ryanfkeepers)

### Changed

- Improve logging (#1642) from [ryanfkeepers](https://github.com/ryanfkeepers)
- Generate separate backup for each resource owner (#1609) from [ashmrtn](https://github.com/ashmrtn)
- Print version info to stdout instead of stderr (#1503) from [meain](https://github.com/meain)

## [v0.0.2] (alpha) - 2022-11-14

### Added

- Added AWS X-Ray support for better observability (#1111) from [ryanfkeepers](https://github.com/ryanfkeepers)
- Allow disabling TLS and TLS verification (#1415) from [vkamra](https://github.com/vkamra)
- Add filtering based on path prefix/contains (#1224) from [ryanfkeepers](https://github.com/ryanfkeepers)
- Add info about doc owner for OneDrive files (#1366) from [meain](https://github.com/meain)
- Add end time for Exchange events from (#1366) [meain](https://github.com/meain)

### Changed

- Export `RepoAlreadyExists` error for sdk users (#1136)from [ryanfkeepers](https://github.com/ryanfkeepers)
- RudderStack logger now respects corso logger settings (#1324) from [ryanfkeepers](https://github.com/ryanfkeepers)

## [v0.0.1] (alpha) - 2022-10-24

### New features

- Supported M365 Services

  - Exchange - email, events, contacts ([RM-8](https://github.com/alcionai/corso-roadmap/issues/28))
  - OneDrive - files ([RM-12](https://github.com/alcionai/corso-roadmap/issues/28))

- Backup workflows

  - Create a full backup ([RM-19](https://github.com/alcionai/corso-roadmap/issues/19))
  - Create a backup for a specific service and all or some data types ([RM-19](https://github.com/alcionai/corso-roadmap/issues/19))
  - Create a backup for all or a specific user ([RM-20](https://github.com/alcionai/corso-roadmap/issues/20))
  - Delete a backup manually ([RM-24](https://github.com/alcionai/corso-roadmap/issues/24))

- Restore workflows

  - List, filter, and view backup content details ([RM-23](https://github.com/alcionai/corso-roadmap/issues/23))
  - Restore one or more items or folders from backup ([RM-28](https://github.com/alcionai/corso-roadmap/issues/28), [RM-29](https://github.com/alcionai/corso-roadmap/issues/29))
  - Non-destructive restore to a new folder/calendar in the same account ([RM-30](https://github.com/alcionai/corso-roadmap/issues/30))

- Backup storage

  - Zero knowledge encrypted backups with user conrolled passphrase ([RM-6](https://github.com/alcionai/corso-roadmap/issues/6))
  - Initialize and connect to an S3-compliant backup repository ([RM-5](https://github.com/alcionai/corso-roadmap/issues/5))

- Miscellaneous
  - Optional usage statistics reporting ([RM-35](https://github.com/alcionai/corso-roadmap/issues/35))

[Unreleased]: https://github.com/alcionai/corso/compare/v0.18.0...HEAD
[v0.18.0]: https://github.com/alcionai/corso/compare/v0.17.0...v0.18.0
[v0.17.0]: https://github.com/alcionai/corso/compare/v0.16.0...v0.17.0
[v0.16.0]: https://github.com/alcionai/corso/compare/v0.15.0...v0.16.0
[v0.15.0]: https://github.com/alcionai/corso/compare/v0.14.0...v0.15.0
[v0.14.0]: https://github.com/alcionai/corso/compare/v0.13.0...v0.14.0
[v0.13.0]: https://github.com/alcionai/corso/compare/v0.12.0...v0.13.0
[v0.12.0]: https://github.com/alcionai/corso/compare/v0.11.1...v0.12.0
[v0.11.1]: https://github.com/alcionai/corso/compare/v0.11.0...v0.11.1
[v0.11.0]: https://github.com/alcionai/corso/compare/v0.10.0...v0.11.0
[v0.10.0]: https://github.com/alcionai/corso/compare/v0.9.0...v0.10.0
[v0.9.0]: https://github.com/alcionai/corso/compare/v0.8.1...v0.9.0
[v0.8.0]: https://github.com/alcionai/corso/compare/v0.7.1...v0.8.0
[v0.7.0]: https://github.com/alcionai/corso/compare/v0.6.1...v0.7.0
[v0.6.1]: https://github.com/alcionai/corso/compare/v0.5.0...v0.6.1
[v0.5.0]: https://github.com/alcionai/corso/compare/v0.4.0...v0.5.0
[v0.4.0]: https://github.com/alcionai/corso/compare/v0.3.0...v0.4.0
[v0.3.0]: https://github.com/alcionai/corso/compare/v0.2.0...v0.3.0
[v0.2.0]: https://github.com/alcionai/corso/compare/v0.1.0...v0.2.0
[v0.1.0]: https://github.com/alcionai/corso/compare/v0.0.4...v0.1.0
[v0.0.4]: https://github.com/alcionai/corso/compare/v0.0.3...v0.0.4
[v0.0.3]: https://github.com/alcionai/corso/compare/v0.0.2...v0.0.3
[v0.0.2]: https://github.com/alcionai/corso/compare/v0.0.1...v0.0.2
[v0.0.1]: https://github.com/alcionai/corso/tag/v0.0.1
