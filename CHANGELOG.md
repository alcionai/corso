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
  * Exchange - email, events, contacts
  * OneDrive - files

* Backup workflows
  * Create a full backup
  * Create a backup for a specific service and all or some data types
  * Create a backup for all or a specific user
  * Delete a backup manually

* Restore workflows
  * List, filter, and view backup content details
  * Restore one or more items or folders from backup
  * Non-destructive restore to a new folder/calendar in the same account

* Backup storage
  * Zero knowledge encrypted backups with user conrolled passphrase
  * Initialize and connect to an S3-compliant backup repository

* Miscelaneous
  * Optional usage statistics reporting
