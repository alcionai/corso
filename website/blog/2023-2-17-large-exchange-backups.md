---
slug: large-microsoft-365-exchange-backups
title: "Backing up large Microsoft 365 Exchange mailboxes with Corso"
description: "A guide to using Corso to back up very large Exchange mailboxes in Microsoft 365"
authors: nica
tags: [corso, microsoft 365, backups, S3]
date: 2023-2-17
image: ./images/heavy-mover.jpg
---

![heavy earth mover By Lowcarb23 - Own work, CC BY-SA 4.0, https://commons.wikimedia.org/w/index.php?curid=114344394](./images/heavy-mover.jpg)

Over the last few months it’s been amazing sharing Corso with more and more users. One pleasant surprise has been users
who are operating in large, often multi-tenant deployments of Microsoft 365 who want to use Corso to back up all their
data. In our discussions on the [Corso User Discord](https://discord.gg/63DTTSnuhT), we’ve found some best practices for
backing up large Exchange mailboxes with Corso.

<!-- truncate -->

### Make sure you’re using the latest version of Corso

We have recently done a lot of work  to harden Corso against transient network outages and Graph API timeouts. This
hardening work makes the most impact during large backups as their long runtime increase the probability of running
into transient errors.

Our recent work has also included support for incremental backups, which you’ll definitely need for larger data sets.
This means that while your first backup of a user with a large mailbox can take some time, all subsequenet backups
will be blazingly fast as Corso will only capture the incremental changes while still constructing a full backup.

### Don’t be afraid to restart your backups

Fundamentally, Corso is a consumer of the Microsoft Graph API, which like all complex API’s, isn’t 100% predictable.
Even in the event of a failed backup, Corso will often have stored multiple objects in the course of a backup. Corso
will work hard to reuse these stored objects in the next backup, meaning your next backup isn’t starting from
zero. A second attempt is likely to run faster with a better chance of completing successfully.

### Batch your users

If many of your users have large file attachments (or if you have more than a few hundred users), you’ll want to batch
your users for your first backup. A tool like [Microsoft365dsc](https://microsoft365dsc.com/) can help you get a list
of all user emails ready for parsing. After that you can back up a few users or even a single user at a time with the
Corso command `corso backup create exchange --user "alice@example.com,bob@example.com"`

Why can’t you just run them all in one go with `--user '*'` ? Again we’re limited by the Microsoft’s Graph API which
often has timeouts, 5xx errors, and throttles its clients.

The good news is that with Corso’s robust ability to do incremental backups, after your first backup, you can
absolutely use larger batches of users, as all future backups to the same repository will run **much** faster.

### Use multiple repositories for different tenants

If you’re a managed service provider or otherwise running a multi-tennant architecture, you should use multiple separate
repositories with Corso. Two ways to pursue this:

- Point to separate buckets
- Place other repositories in subfolders of the same bucket with the `prefix` option

In both cases, the best way to keep these settings tidy is by using multiple `.corso.toml`
[configuration files](../../docs/setup/configuration/#configuration-file). Use the
`-config-file` option to point to separate config files

### Have questions?

Corso is under active development, and we expect our support for this type of use case to improve rapidly.
If you have feedback for us please do [join our discord](https://discord.gg/63DTTSnuhT) and talk directly with the team!
