# FAQ

<details>
    <summary>Where does Corso store local configuration?</summary>

Corso's local configuration is stored in a file named `.corso.toml` in your home directory. Corso can be pointed at
separate configuration files using the `--config-file` option.

</details>

<details>
    <summary>Does Corso report usage telemetry?</summary>

In order to better understand how people use Corso and to guide feature development, Corso supports reporting telemetry
metadata for basic information about installed versions and usage in a privacy-preserving manner. This includes a
generic description of most-commonly used backup operations and statistics on the duration and size of backups. No user
data is stored or transmitted during this process.

Telemetry reporting can be turned off by using the `--no-stats` flag. See the [Command Line Reference](../../cli/corso)
section for more information.

</details>

<details>
    <summary>Does Corso work with Microsoft 365 Personal or Family subscriptions?</summary>

Unfortunately, Corso leverages the Microsoft Graph API and that's only accessible to organizations with paid
subscriptions to Microsoft 365 (for example, Business, Enterprise, or Education plans).

</details>
