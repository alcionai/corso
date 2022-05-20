# CLI Commands
## Status

Revision: v0.0.1

-----


This is a proposal for Corso cli commands extrapolated from the Functional Requirements product documentation.  Open questions are listed in the `Details & Discussion` section.  The command set includes some p1/p2 actions for completeness.  This proposal only intends to describe the available commands themselves and does not evaluate functionality or feature design beyond that goal.

# CLI Goals

- Ease (and enjoyment) of Use, more than minimal functionality.
- Intended for use by Humans, not Computers.
- Outputs should be either interactive/progressive (for ongoing work) or easily greppable/parseable.

## Todo/Undefined:

- Interactivity and sub-selection/helpful action completion within command operation.
- Quality-of-life and niceties such as interactive/output display, formatting and presentation, or maximum minimization of user effort to run Corso.

-----
## Commands

Standard format:  
`corso {command} [{subcommand}] [{service|repository}] [{flag}...]` 

| Cmd |  |  | Flags | Notes |
| --- | --- | --- | --- | --- |
| version |  |  |  | Same as `corso --version` |
|  |  |  | —version | Outputs Corso version details. |
| help |  |  |  | Same as `corso —-help` |
| * | * | help |  | Same as `{command} -—help` |
| * | * |  | —help | Same as `{command} help` |

| Cmd |  |  | Flags | Notes |
| --- | --- | --- | --- | --- |
| repo | * |  |  | Same as `repo [*] --help`. |
| repo | init | {repository} |  | Initialize a Corso repository. |
| repo | init | {repository} | —tenant {tenant_id} | Provides the account’s tenant ID. |
| repo | init | {repository} | —client {client_id} | Provides the account’s client ID. |
| repo | connect | {repository} |  | Connects to the specified repo. |
| repo | configure | {repository} |  | Sets mutable config properties to the provided values. |
| repo | * | * | —config {cfg_file_path} | Specify a repo configuration file.  Values may also be provided via individual flags and env vars. |
| repo | * | * | —{config-prop} | Blanket commitment to support config via flags. |
| repo | * | * | —credentials {creds_file_path} | Specify a file containing credentials or secrets.  Values may also be provided via env vars. |

| Cmd |  |  | Flags | Notes |
| --- | --- | --- | --- | --- |
| backup | * |  |  | Same as backup [*] -—help |
| backup | list | {service} |  | List all backups in the repository for the specified service. |
| backup | create | {service} |  | Backup the specified service. |
| backup | * | {service} | —token {token} | Provides a security key for permission to perform backup. |
| backup | * | {service} | —{entity} {entity_id}... | Only involve the target entity(s).  Entities are things like users, groups, sites, etc.  Entity flag support is service-specific. |

| Cmd |  |  | Flags | Notes |
| --- | --- | --- | --- | --- |
| restore |  |  |  | Same as `restore -—help` |
| restore | {service} |  |  | Complete service restoration using the latest versioned backup. |
| restore | {service} |  | —backup {backup_id} | Restore data from only the targeted backup(s). |
| restore | {service} |  | —{entity} {entity_id}... | Only involve the target entity(s).  Entities are things like users, groups, sites, etc.  Entity flag support is service-specific. |
---


## Examples
### Basic Usage

**First Run**

```bash
$ export O365_SECRET=my_0365_secret
$ export AWS_SECRET_ACCESS_KEY=my_s3_secret
$ corso repo init s3 --bucket my_s3_bucket --access-key my_s3_key \
		--tenant my_m365_acct --clientid my_m365_client_id
$ corso backup express
```

**Follow-up Actions**

```bash
$ corso repo connect s3 --bucket my_s3_bucket --access-key my_s3_key
$ corso backup express
$ corso backup list express
```
-----

# Details & Discussion

## UC0 - CLI User Interface

Base command: `corso`

Standard format: `corso {command} [{subcommand}] [{service}] [{flag}...]`

Examples:

- `corso help`
- `corso repo init --repository s3 --tenant t_1`
- `corso backup create teams`
- `corso restore teams --backup b_1`

## UC1 - Initialization and Connection

**Account Handling**

M365 accounts are paired with repo initialization, resulting in a single-tenancy storage.  Any `repo` action applies the same behavior to the account as well.  That is, `init` will handle all initialization steps for both the repository and the account, and both must succeed for the command to complete successfully, including all necessary validation checks.  Likewise, `connect` will validate and establish a connection (or, at least, the ability to communicate) with both the account and the repository.

**Init**

`corso repo init {repository} --config {cfg} --credentials {creds}`

Initializes a repository, bootstrapping resources as necessary and storing configuration details within Corso.  Repo is the name of the repository provider, eg: ‘s3’.  Cfg and creds, in this example, point to json (or alternatively yaml?) files containing the details required to establish the connection.  Configuration options, when known, will get support for flag-based declaration.  Similarly, env vars will be supported as needed.

**Connection**

`corso repo connect {repository} --credentials {creds}` 

[https://docs.flexera.com/flexera/EN/SaaSManager/M365CCIntegration.htm#integrations_3059193938_1840275](https://docs.flexera.com/flexera/EN/SaaSManager/M365CCIntegration.htm#integrations_3059193938_1840275)

Connects to an existing (ie, initialized) repository.

Corso is expected to gracefully handle transient disconnections during backup/restore runtimes (and otherwise, as needed).

**Deletion**

`corso repo delete {repository}`

(Included here for discussion, but not being added to the CLI command set at this time.)

Removes a repository from Corso.  More exploration is needed here to explore cascading effects (or lack thereof) from the command.  At minimum, expect additional user involvement to confirm that the deletion is wanted, and not erroneous.

## UC1.1 - Version

`corso --version` outputs the current version details such as: commit id and datetime, maybe semver (complete release version details to be decided).
Further versioning controls are not currently covered in this proposal.

## UC2 - Configuration

`corso repo configure --reposiory {repo} --config {cfg}`

Updates the configuration details for an existing repository.

Configuration is divided between mutable and immutable properties.  Generally, initialization-specific configurations (those that identify the storage repository, it’s connection, and its fundamental behavior), among other properties, are considered immutable and cannot be reconfigured.  As a result, `repo configure` will not be able to rectify a misconfigured init; some other user flow will be needed to resolve that issue.

Configure allows mutation of config properties that can be safely and transiently applied.  For example: backup retention and expiration policies. A complete list of how each property is classified is forthcoming as we build that list of properties.

## UC3 - On-Demand Backup

`corso backup` is reserved as a non-actionable command, rather than have it kick off a backup action.  This is to ensure users don’t accidentally kick off a migration in the process of exploring the api.  `corso backup` produces the same output as `corso backup --help`.

**Full Service Backup**

- `corso backup create {service}`

**Selective Backup**

- `corso backup create {service} --{entity} {entity_id}...`

Entities are service-applicable objects that match up to m365 objects.  Users, groups, sites, mailboxes, etc.  Entity flags are available on a per-service basis.  For example, —site is available for the sharepoint service, and —mailbox for express, but not the reverse.  A full list of system-entity mappings is coming in the future.

**Examples**

- `corso backup` → displays the help output.
- `corso backup create teams` → generates a full backup of the teams service.
- `corso backup create express --group g_1` → backs up the g_1 group within express.

## UC3.2 - Security Token

(This section is incomplete: further design details are needed about security expression.)  Some commands, such as Backup/Restore require a security key declaration to verify that the caller has permission to perform the command.

`corso * * --token {token}`

## UC5 - Backup Ops

`corso backup list {service}` 

Produces a list of the backups which currently exist in the repository.

`corso backup list {service} --{entity} {entity_id}...`

The list can be filtered to contain backups relevant to the specified entities.  A possible user flow for restoration is for the user to use this to discover which backups match their needs, and then apply those backups in a restore operation.

**Expiration Control**

Will appear in a future revision.

## UC6 - Restore

Similar to backup, `corso restore` is reserved as a non-actionable command to serve up the same output as `corso restore —help`.

### UC6.1

**Full Service Restore**

- `corso restore {service} [--backup {backup_id}...]`

If no backups are specified, this defaults to the most recent backup of the specified service.

**Selective Restore**

- `corso restore {service} [--backup {backup_id}...] [--{entity} {entity_id}...]`

Entities are service-applicable objects that match up to m365 objects.  Users, groups, sites, mailboxes, etc.  Entity flags are available on a per-service basis.  For example, —site is available for the sharepoint service, and —mailbox for express, but not the reverse.  A full list of system-entity mappings is coming in the future.

**Examples**

- `corso restore` → displays the help output.
- `corso restore teams` → restores all data in the teams service.
- `corso restore sharepoint --backup b_1` → restores the sharepoint data in the b_1 backup.
- `corso restore express --group g_1` → restores the g_1 group within sharepoint.

## UC6.2 - disaster recovery

Multi-service backup/restoration is still under review.
