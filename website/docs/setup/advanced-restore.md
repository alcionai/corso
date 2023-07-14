# Advanced Restorations

import CodeBlock from '@theme/CodeBlock';
import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';
import {Version} from '@site/src/corsoEnv';

The basic restore command is an easy way to return all of your data.
But when you need more control over the results you can use the advanced
configuration options to change where and how the data gets restored

## Destination

The `--destination` flag lets you select the top-level folder where corso will
write all of the restored data.

<CodeBlock language="bash">{
    `corso restore onedrive --backup abcd --destination my-latest-restore`
}</CodeBlock>

If the flag is not provided, Corso will use its standard restore folder name:
`Corso_Restore_<current-date-time>`.  If you're concerned about data integrity
then this is always the safest option, 

In the above example, the onedrive restore will appear in a top-level folder
named `my-latest-restore`.  If that folder does not already exist, Corso will
automatically create it.  If it does exist, the restore will use the existing
folder and restore the items again. This allows you to restore to the same folder
multiple times.

You can restore items back to their original location by setting the destination
to `/`: `--destination /`.  This skips the creation of a top-level folder, and
all restored items will appear back in their location at the time of backup.

### Known Issues

* Destination will not create N-depth folder structures.  `--destination a/b/c`
does not create three folders; it creates a single, top-level folder named `a/b/c`.

* Exchange Calendars do not support folder hierarchy.  If your backup contains the
calendars `MyCalendar` and `Birthdays`, and you restore to `--destination Restored`,
all of the restored calendar events will appear in the `Restored` calendar.  However,
if you restore events in-place (`--destination /`) then all events will return to
their original calendars.

*  When restoring Exchange Calendar Events to a destination folder, Events that were
previously isolated in different calendars may collide with each other in the destination
calendar.

## Item Collision Handling

When restoring data into an existing folder, the items from the backup may conflict
with existing data.  When this happpens, Corso resolves the conflict using its
collision configuration.

<CodeBlock language="bash">{
    `corso restore onedrive --backup abcd --collisions skip`
}</CodeBlock>

Collision detection differs between each service and type of data.  The general
comparison always follows the same pattern: "within the current folder, if the
restore item looks identical to an existing item, it collides".

The comparison uses item metadata (names, subjects, titles, etc), not item content.
So if the current `reports.txt` has different contents than the backup `reports.txt`,
it still collides.

Collisions can be handled with three different configurations: `Skip`, `Copy`,
and `Replace`

## Skip

The default behavior.  When a collision is identified, no restore is attempted.

## Copy

Item collisions create a copy of the item in the backup.  The copy holds the backup
version of the item, leaving the current version unchanged.  If necessary, changes
item properties (such as filenames) to avoid additional collisions.

## Replace

Collisions will entirely replace the current version of the item with the backup
version.  If multiple existing items collide with the backup item, only one of the
existing items is replaced.

