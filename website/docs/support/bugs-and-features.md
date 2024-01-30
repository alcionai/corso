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
  GODEBUG=gctrace=1 .\corso <command> -hide-progress --log-level debug --mask-sensitive-data `
    <command-options> > corso-gc.log 2>&1
  ```

</TabItem>
<TabItem value="unix" label="Linux/macOS">

  ```bash
  # Connect to the Corso Repository
  GODEBUG=gctrace=1 ./corso <command> -hide-progress --log-level debug --mask-sensitive-data \
    <command-options> > corso-gc.log 2>&1
  ```

</TabItem>
</Tabs>

Next, file a [GitHub issue](https://github.com/alcionai/corso/issues/) with the two log files
([default log file](../../setup/configuration/#log-files) and `corso-gc.log`, the Corso GC log file, from above) and
information on the size of the Exchange mailbox, OneDrive location, or SharePoint site that you are having an issue
with.

## Sharepoint List Anomalies

Some columns in sharepoint list aren't recognizable from the GRAPH API response.
More about this issue [here](https://github.com/alcionai/corso/issues/5166).

Therefore while `restore` of these columns, we default them to as text fields.
The value they hold are therefore not reinstated to the way the originally were.

<Tabs groupId="columns">
<TabItem value="hyp" label="Hyperlink">

### Originally created hyperlink column in Site

![diagram of list with Hyperlink column in a site](../../blog/images/Hyperlink-Column.png)

### Restored hyperlink column

![diagram of restored list with Hyperlink column in a site](../../blog/images/Restored-Hyperlink-Column.png)

### Issue tracker for hyperlink column support

To track progress, [visit](https://github.com/microsoftgraph/msgraph-sdk-go/issues/640).

</TabItem>
<TabItem value="loc" label="Location">

### Originally created location column in Site

![diagram of list with Location column in a site](../../blog/images/Location-Column.png)

### Restored location column

![diagram of restored list with Location column in a site](../../blog/images/Restored-Location-Column.png)

### Issue tracker for location column support

To track progress, [visit](https://github.com/microsoftgraph/msgraph-sdk-go/issues/638).
</TabItem>
</Tabs>