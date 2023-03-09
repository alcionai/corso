---
slug: tutorial-corso-videos
title: "Backup Basics: A Video Tour of Using Free Open-Source Corso to Protect your Microsoft 365 Data"
description: "By watching my
four short videos, you’ll be set up and on your way to backing up all your M365
data in less than 20 minutes. That’s shorter than most meetings!"
authors: nica
tags: [corso, microsoft 365, backups]
date: 2023-3-4
image: ./images/vhs.jpg
---

![a VHS tape being loaded into a player](./images/vhs.jpg)

Each day, you and your colleagues put hours of work into creating, sending, and
receiving all kinds of critical data: hundreds of emails, Word documents,
spreadsheets, and more. You and your company want to protect your
data from everything that could go wrong, such as server outages, cyberattacks,
accidental deletions and anything that could cause you to lose valuable work,
time, and money. Even small instances of data loss cost businesses an
[average of between $18,000 and $35,000](https://invenioit.com/continuity/cost-of-data-loss/)
– and that’s with losses of fewer than 100 files! That number can grow into the
millions for large-scale losses and breaches.

My question to you is, how can you protect your data without spending countless hours and thousands of dollars?

<!-- truncate -->

The answer? Corso, a free open-source backup tool that helps you safeguard your
teams’ Microsoft 365 data, and it’s really quick to get started! By watching my
four short videos, you’ll be set up and on your way to backing up all your M365
data in less than 20 minutes. That’s shorter than most meetings!

## Get started with Corso for Microsoft 365 backups

As you install and set up Corso, take a look at my
[instructional video](https://youtu.be/mlwfEbPqD94) and
[Quick Start](https://corsobackup.io/docs/quickstart/) guide. Together, we walk
you through downloading Corso, setting up an AWS S3 bucket with user access
permissions, establishing Microsoft 365 access, and initiating a repository.
From there, you can create your first backup and be ready to interact with it
within minutes.

<iframe width="560" height="315" src="https://www.youtube.com/embed/mlwfEbPqD94"
title="YouTube video player" frameborder="0" allow="accelerometer; autoplay;
clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share"
allowfullscreen></iframe>

## Backup all your exchange data with Corso

Now that you’ve got a handle on Corso basics, you’re ready to tackle more
involved data backups. [Watch](https://youtu.be/R1AOc2xz2Rg) how you can back up
data from multiple users and specify desired data types after connecting to an
existing S3 bucket. The status and user associated with each backup will
automatically display, and backups for each user have individual IDs. With
Corso, you can back up specific users’ data or all user data, delete data for
users that are no longer needed, and request a list of all past backups.

<iframe width="560" height="315" src="https://www.youtube.com/embed/R1AOc2xz2Rg"
title="YouTube video player" frameborder="0" allow="accelerometer; autoplay;
clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share"
allowfullscreen></iframe>

## Explore the details of your exchange server backup
<!-- vale Microsoft.Contractions = NO -->
After completing your first backup, dig into the details of your data through a
variety of commands and filters. See how quick it is to get information about
different data types, locations, and other attributes such as event type or
email subject line. The steps in this [video](https://youtu.be/mweAUDhUE7I) are
great for restoring only the components of a backup you’re interested in, like
non-recurring events for a specific user.
<!-- vale Microsoft.Contractions = YES -->

<iframe width="560" height="315" src="https://www.youtube.com/embed/mweAUDhUE7I"
title="YouTube video player" frameborder="0" allow="accelerometer; autoplay;
clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share"
allowfullscreen></iframe>

## Create local Microsoft 365 backups using MinIO

For quick, low latency backup testing, you can use MinIO to run a Corso M365
backup through a local S3 bucket. While you don’t want to rely on this as your
primary backup location, it's a cost effective option for testing that gives
you full control. See my [video](https://youtu.be/ABIiVufyOkM) on local M365
backups for a step-by-step guide on this Corso feature.

<iframe width="560" height="315" src="https://www.youtube.com/embed/ABIiVufyOkM"
title="YouTube video player" frameborder="0" allow="accelerometer; autoplay;
clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share"
allowfullscreen></iframe>

## Restoring from your Backups

We can hope this day will never come. Backups ideally are stored for a period of
time, then rotated out, and never used. But should you deal with a corruption of
data, ransomware attack, or other unforeseen circumstance, it’ll be time to use
Corso to restore from your backups.

In this last video I demonstrate both total restorations for a user or group of
users, and restoring single records with the `corso restore` command.

<iframe width="560" height="315" src="https://www.youtube.com/embed/ABIiVufyOkM"
title="YouTube video player" frameborder="0" allow="accelerometer; autoplay;
clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share"
allowfullscreen></iframe>

Remember, don’t leave your data vulnerable to cyberattacks, accidental
deletion, server outages, or any of the countless other sources of data loss.
Try out [Corso](https://corsobackup.io/) for your company’s M365 backup needs,
and be sure to check out my Corso video
[playlist](https://youtube.com/playlist?list=PLSukexZlj1V0D0xGV2ON-MWRmPpLWi6hK)
on YouTube for more tips and tricks on protecting your data! Please come find me
on [Discord](https://discord.gg/63DTTSnuhT), join our discussions, and give us
your feedback. See you there!
