# Bugs and new features

import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';

You can learn more about the Corso roadmap and how to interpret it [here](https://github.com/alcionai/corso-roadmap).

If you run into a bug or have feature requests, please file a [GitHub issue](https://github.com/alcionai/corso/issues/)
and attach the `bug` or `enhancement` label to the issue. When filing bugs, please run Corso with
`--log-level debug --hide-progress --mask-sensitive-data` and add the logs to the bug report. You can find more
information about where logs are stored in the [log files](../../setup/configuration/#log-files) section in setup docs.

## Memory Issues

Corso's memory usage depends on the size of the resource being backed up. The maximum memory usage occurs during full
backups (usually the first backup) vs. later incremental backups. If you believe Corso is using unexpected amounts of
memory, please run Corso with the following options:

- Prefix the Corso run with `GODEBUG=gctrace=1` to get GC (Garbage Collection) logs
- Add `--log-level debug --hide-progress --mask-sensitive-data`
- Redirect output to a new log file (for example, `corso-gc.log`)

<Tabs groupId="os">
<TabItem value="win" label="Powershell">

  ```powershell
  # Connect to the Corso Repository
  GODEBUG=gctrace=1 .\corso <command> -hide-progress --log-level debug --mask-sensitive-data <command-options> > corso-gc.log 2>&1
  ```

</TabItem>
<TabItem value="unix" label="Linux/macOS">

  ```bash
  # Connect to the Corso Repository
  GODEBUG=gctrace=1 ./corso <command> -hide-progress --log-level debug --mask-sensitive-data <command-options> > corso-gc.log 2>&1
  ```

</TabItem>
</Tabs>

Next, file a [GitHub issue](https://github.com/alcionai/corso/issues/) with the two log files
([default log file](../../setup/configuration/#log-files) and `corso-gc.log`, the Corso GC log file, from above) and
information on the size of the Exchange mailbox, OneDrive location, or SharePoint site that you are having an issue
with.
