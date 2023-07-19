# Restore options

import CodeBlock from '@theme/CodeBlock';
import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';
import {Version} from '@site/src/corsoEnv';

The default restore command is an great way to restore data in a non-destructive
manner to a new folder. When you need more control over the results you can
use the advanced configuration options to change where and how your data
gets restored.

## Restore to target folder

The `--destination` flag lets you select the top-level folder where Corso will
write all of the restored data.

### The default destination

<CodeBlock language="bash">{
    `corso restore onedrive --backup a422895c-c20c-4b06-883d-b866db9f86ef`
}</CodeBlock>

If the flag isn't provided, Corso will create a new folder with a standard name:
`Corso_Restore_<current-date-time>`. The default restore folder is always newly
created, and can't interfere with any existing items. If you're concerned about
data integrity then this is always the safest option.

### An alternate destination

<CodeBlock language="bash">{
    `corso restore onedrive --destination /my-latest-restore --backup a422895c-c20c-4b06-883d-b866db9f86ef`
}</CodeBlock>

When a destination is manually specified, all restored will appear in that top-level
folder. In the example above, Corso will restore everything into `my-latest-restore`.
If that folder doesn't already exist, Corso will automatically create it. If it does
exist, the restore will use the existing folder, allowing you to restore to the same
folder multiple times.

### The original location

<CodeBlock language="bash">{
    `corso restore onedrive --destination / --backup a422895c-c20c-4b06-883d-b866db9f86ef`
}</CodeBlock>

You can restore items back to their original location by setting the destination
to `/`. This skips the creation of a top-level folder, and all restored items will
appear back in their location at the time of backup.

### Limitations

* Destination won't create N-depth folder structures. `--destination a/b/c`
doesn't create three folders; it creates a single, top-level folder named `a/b/c`.

* Exchange Calendars don't support folder hierarchy. If your backup contains the
calendars `MyCalendar` and `Birthdays`, and you restore to `--destination Restored`,
all of the restored calendar events will appear in the `Restored` calendar. However,
if you restore events in-place (`--destination /`) then all events will return to
their original calendars.

* When restoring Exchange Calendar Events to a destination folder, Events that were
safe in different calendars may collide with each other in the destination calendar.

## Item collision handling

When restoring data into an existing folder, the items restored may conflict
with existing data. When this happens, Corso resolves the conflict using its
collision configuration.

Collision detection differs between each service and type of data. The general
comparison always follows the same pattern: "within the current folder, if the
restore item looks identical to an existing item, it collides."

The comparison uses item metadata (names, subjects, titles, etc), not item content.
If the current `reports.txt` has different contents than the backup `reports.txt`,
it still collides.

Collisions can be handled with three different configurations: `Skip`, `Copy`,
and `Replace`.

### Skip (default)

<CodeBlock language="bash">{
    `corso restore onedrive --collisions skip --destination / --backup a422895c-c20c-4b06-883d-b866db9f86ef`
}</CodeBlock>

When a collision is identified, the item is skipped and
no restore is attempted.

### Copy

<CodeBlock language="bash">{
    `corso restore onedrive --collisions copy --destination /my-latest-restore --backup a422895c-c20c-4b06-883d-b866db9f86ef`
}</CodeBlock>

Item collisions create a copy of the item in the backup. The copy holds the backup
version of the item, leaving the current version unchanged. If necessary, changes
item properties (such as filenames) to avoid additional collisions. Eg:
the copy of`reports.txt` is named `reports 1.txt`.

### Replace

<CodeBlock language="bash">{
    `corso restore onedrive --collisions replace --destination / --backup a422895c-c20c-4b06-883d-b866db9f86ef`
}</CodeBlock>

Collisions will entirely replace the current version of the item with the backup
version. If multiple existing items collide with the backup item, only one of the
existing items is replaced.

## Restore to target resource

The `--to-resource` flag lets you select which resource will receive the restored data.
A resource can be a mailbox, user, sharepoint site, or other owner of data.

When restoring to a target resource, all other restore configuration behaves normally.
Data is restored into the default folder: `Corso_Restore_<current-date-time>` (unless a
`--destination` flag is added).  When restoring in-place, collision policies are followed.

<CodeBlock language="bash">{
    `corso restore onedrive --to-resource adelev@alcion.ai --backup a422895c-c20c-4b06-883d-b866db9f86ef`
}</CodeBlock>

### Limitations

* The resource must exist.  Corso will not create new mailboxes, users, or sites.
* The resource must have access to the service being restored.  No restore will be
performed for an unlicensed resource.