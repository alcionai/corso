---
slug: large-exchange-backups
title: "Backing up large Exchange servers with Corso"
description: "A guide to using Corso to back up very large installations"
authors: nica
tags: [corso, microsoft 365, backups, S3]
date: 2023-2-10
image: ./images/heavy-mover.jpg
---

![heavy earth mover By Lowcarb23 - Own work, CC BY-SA 4.0, https://commons.wikimedia.org/w/index.php?curid=114344394](./images/heavy-mover.jpg)

Over the last few months it’s been amazing sharing Corso with more and more users. One pleasant surprise has been users
who are operating in large, often multi-tenant deployments of Microsoft 365 who want to use Corso to back up all their
data. In our discussions on the [Corso User Discord](https://discord.gg/Gd63GRFvcb), we’ve found some best practices for
backing up large Exchange servers with Corso.
<!-- truncate -->
1. make sure you’re using the latest version of Corso

Our recent work has included support for incremental backups, which you’ll definitely need for larger data sets.

1. Don’t be afraid to re-start your backups

Fundamentally Corso is a client of the Microsoft Graph API, which like all large object API’s, isn’t 100% predictable.
In the event of a failed backup, Corso will often have stored multiple objects in the course of a backup. These stored
objects will often prevent your re-download and re-uploading of objects, meaning your next backup isn’t starting from
zero. A second attempt is likely to run faster with a better chance of completing successfully.

1. Batch your users

If many of your users have large file attachments (or if you have more than a few hundred users), you’ll want to batch
your users for your first backup. A tool like [Microsoft365dsc](https://microsoft365dsc.com/) can help you get a list
of all user emails ready for parsing. After that you can back up a single user at a time with the Corso command
`corso backup create exchange --user alice@example.com`

Why can’t you just run them all in one go with `--user '*'` ? Again we’re limited by the Microsoft’s Graph API which
often has timeouts, 5xx errors, and throttles its clients.

While the tool will accept an array of users, it probably makes sense just to run your first backups for a single user
at a time. The good news is that with Corso’s robust ability to do incremental backups, after your first backup, you can
absolutely batch up your users, as all future backups to the same repository will run **much** faster.

1. Use multiple repositories for different tenants

If you’re a managed service provider or otherwise running a multi-tennant architecture, you should use multiple separate
repositories with Corso. Two ways to pursue this:

- point to separate buckets
- place other repositories in subfolders of the same bucket with the `prefix` option

In both cases, the best way to keep these settings tidy is by using multiple `.corso.toml` configuration files. Use the
`-config-file` option to point to separate config files

Corso is under active development, and we expect our support for this type of use case to improve rapidly.
If you have feedback for us please do [join our discord](https://discord.gg/Gd63GRFvcb) and talk directly with the team!
