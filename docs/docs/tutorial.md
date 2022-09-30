# Tutorial

In this tutorial you will perform your first backup followed by a restore.

## Prerequisites

* Install Docker
* Pull the Corso container (see [Installation](/install))
* Configure connection to your M365 Tenant (see [M365 Access](/configuration/m365_access))
* Initialize a Corso backup repository (see [Repositories](/configuration/repos))

## Your first backup

Corso can do much more, but you can start by creating a backup of your Exchange mailbox.

To do this, you can run the following command:

```bash
$ docker run -e CORSO_PASSPHRASE \
    --env-file ~/.corso/corso.env \
    -v ~/.corso/config:/app/config \
    -v ~/.corso/logs:/app/logs corso/corso:latest \
    backup create exchange --user <your exchange email address>
```

:::note
Your first backup may take some time if your mailbox is large.
:::

<!-- vale proselint.Annotations = NO -->
**TODO:** Update ^^^ after the finalization of Corso output from operations.
<!-- vale proselint.Annotations = YES -->

## Restore an email

Now lets explore how you can restore data from one of your backups.

You can see all Exchange backups available with the following command:

```bash
$ docker run -e CORSO_PASSPHRASE \
    --env-file ~/.corso/corso.env \
    -v ~/.corso/config:/app/config \
    -v ~/.corso/logs:/app/logs corso/corso:latest \
    backup list exchange --user <your exchange email address>
```
<!-- vale proselint.Annotations = NO -->
**TODO:** Update ^^^ after the finalization of Corso output from operations.
<!-- vale proselint.Annotations = YES -->

Select one of the available backups and search through its contents.

```bash
$ docker run -e CORSO_PASSPHRASE \
    --env-file ~/.corso/corso.env \
    -v ~/.corso/config:/app/config \
    -v ~/.corso/logs:/app/logs corso/corso:latest \
    backup details exchange \
    --backup <id of your selected backup> \
    --user <your exchange email address> \
    --email-subject <portion of subject of email you want to recover>
```

The output from the command above should display a list of any matching emails. Note the ID
of the one to use for testing restore.

When you are ready to restore, use the following command:

```bash
$ docker run -e CORSO_PASSPHRASE \
    --env-file ~/.corso/corso.env \
    -v ~/.corso/config:/app/config \
    -v ~/.corso/logs:/app/logs corso/corso:latest \
    backup details exchange \
    --backup <id of your selected backup> \
    --user <your exchange email address> \
    --email <id of your selected email>
```

You can now find the recovered email in a mailbox folder named `Corso_Restore_DD-MMM-YYYY_HH:MM:SS`.

You are now ready to explore the [Command Line Reference](cli/corso) and try everything that Corso can do.
