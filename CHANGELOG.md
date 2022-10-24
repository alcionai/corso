# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

[Unreleased]: https://github.com/https://github.com/alcionai/corso/compare/...HEAD

## v0.0.1 (alpha)

Release date: 2022-10-24

### New features

* Supported M365 Services
  * Exchange - email, events, contacts (RM-8)
  * OneDrive - files (RM-12)

* Backup workflows
  * Create a full backup (RM-19)
  * Create a backup for a specific service and all or some data types (RM-19)
  * Create a backup for all or a specific user (RM-20)
  * Delete a backup manually (RM-24)

* Restore workflows
  * List, filter, and view backup content details (RM-23)
  * Restore one or more items or folders from backup (RM-28, RM-29)
  * Non-destructive restore to a new folder/calendar in the same account (RM-30)

* Backup storage
  * Zero knowledge encrypted backups with user conrolled passphrase (RM-6)
  * Initialize and connect to an S3-compliant backup repository (RM-5)

* Miscelaneous
  * Optional usage statistics reporting (RM-35)
