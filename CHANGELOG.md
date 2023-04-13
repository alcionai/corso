# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased] (beta)

### Added
- Permissions backup for OneDrive is now out of experimental (By default, only newly backed up items will have their permissions backed up. You will have to run a full backup to ensure all items have their permissions backed up.)

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

[Unreleased]: https://github.com/alcionai/corso/compare/v0.6.1...HEAD
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
