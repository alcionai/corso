# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]


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

[Unreleased]: https://github.com/alcionai/corso/compare/v0.0.2...HEAD
[v0.0.2]: https://github.com/alcionai/corso/compare/v0.0.1...v0.0.2
[v0.0.1]: https://github.com/alcionai/corso/tag/v0.0.1
