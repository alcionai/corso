# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased] (alpha)

### Added

- OneDrive file/folder permissions can now be backed up and restored
- Add `--restore-permissions` flag to toggle restoration of OneDrive permissions
- Add a hidden `--disable-permissions-backup` to completely disable backing up permissions
- Add versions to backups so that we can understand/handle older backup formats

### Known Issues

- When the same user has permissions to a file and the containing
  folder, we only restore folder level permissions for the user and no
  separate file only permission is restored.


## [v0.2.0] (alpha) - 2023-1-29

### Fixed

- Check if the user specified for an exchange backup operation has a mailbox.
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

[Unreleased]: https://github.com/alcionai/corso/compare/v0.2.0...HEAD
[v0.2.0]: https://github.com/alcionai/corso/compare/v0.1.0...v0.2.0
[v0.1.0]: https://github.com/alcionai/corso/compare/v0.0.4...v0.1.0
[v0.0.4]: https://github.com/alcionai/corso/compare/v0.0.3...v0.0.4
[v0.0.3]: https://github.com/alcionai/corso/compare/v0.0.2...v0.0.3
[v0.0.2]: https://github.com/alcionai/corso/compare/v0.0.1...v0.0.2
[v0.0.1]: https://github.com/alcionai/corso/tag/v0.0.1