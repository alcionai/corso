# CLI Commands
## Status

Revision: v0.0.1

-----


## Commands

Standard format:  
`corso <command> [<subcommand>] [<service|repository>] [<flag>...]` 

| Cmd |  |  | Flags | Notes |
| --- | --- | --- | --- | --- |
| version |  |  |  | Same as `corso --version` |
|  |  |  | —version | Outputs Corso version details. |
| help |  |  |  | Same as `corso —-help` |
| * | * | help |  | Same as `{command} -—help` |
| * | * |  | —help | Same as `{command} help` |
|  |  |  |  |  |
|  |  |  |  |  |
| repo | * |  |  | Same as `repo [*] --help`. |
| repo | init | \<repository> |  | Initialize a Corso repository. |
| repo | init | \<repository> | —tenant {tenant_id} | Provides the account’s tenant ID. |
| repo | init | \<repository> | —client {client_id} | Provides the account’s client ID. |
| repo | connect | \<repository> |  | Connects to the specified repo. |
| repo | configure | \<repository> |  | Sets mutable config properties to the provided values. |
| repo | * | * | —config {cfg_file_path} | Specify a repo configuration file.  Values may also be provided via individual flags and env vars. |
| repo | * | * | —{config-prop} | Blanket commitment to support config via flags. |
| repo | * | * | —credentials {creds_file_path} | Specify a file containing credentials or secrets.  Values may also be provided via env vars. |
|  |  |  |  |  |
|  |  |  |  |  |
| backup | * |  |  | Same as backup [*] -—help |
| backup | list | \<service> |  | List all backups in the repository for the specified service. |
| backup | create | \<service> |  | Backup the specified service. |
| backup | * | \<service> | —token {token} | Provides a security key for permission to perform backup. |
| backup | * | \<service> | —\<entity> {<entity_id>}... | Only involve the target entity(s).  Entities are things like users, groups, sites, etc.  Entity flag support is service-specific. |
|  |  |  |  |  |
|  |  |  |  |  |
| restore |  |  |  | Same as `restore -—help` |
| restore | \<service> |  |  | Complete service restoration using the latest versioned backup. |
| restore | \<service> |  | —backup {backup_id} | Restore data from only the targeted backup(s). |
| restore | \<service> |  | —\<entity> {<entity_id>}... | Only involve the target entity(s).  Entities are things like users, groups, sites, etc.  Entity flag support is service-specific. |
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
---


## Change Log
### v0.0.1
Initial release.