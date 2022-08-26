# Tutorial

In this tutorial you will perform your first backup followed by a restore.

## Prerequisites

* Docker is installed and Corso container is available (see [Installation](/install))
* Corso is connected to your M365 Tenant (see [M365 Access](/configuration/m365_access))
* Corso has initilized a backup repository (see [Repositories](/configuration/repos))

## Your first backup

Corso can do much more, but let's start by creating a backup of your Exchange mailbox.

To accomplish this, you can execute the following command:

```bash
$ docker run -e CORSO_PASSPHRASE \
    --env-file ~/.corso/corso.env \
    -v ~/.corso/config:/app/config \
    -v ~/.corso/logs:/app/logs corso/corso:latest \
    backup create exchange --user <your exchange email address>
```

:::note
Your first backup may take some time if your mailbox has a large number of items so please be patient.
:::

**TODO:** Update ^^^ after Corso output from operations is finalized.

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

**TODO:** Update after Corso output from operations is finalized.

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

The above should give you a list of any matching emails. Note the ID of the one you would like to
use for testing restore.

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

The email would now be recovered in a folder named "Corso_Restore_DD-MMM-YYYY_HH:MM:SS" in your mailbox.

You are now ready to explore the [Command Line Reference](cli) and try everything that Corso can do for you.
