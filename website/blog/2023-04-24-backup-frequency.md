---
slug: how-often-should-you-run-microsoft-365-backups
title: "How often should you run Microsoft 365 backups?"
description: "On the ideal cadence for backups"
authors: nica
tags: [corso, microsoft 365, backups, best practices]
date: 2023-04-24
image: ./images/astro-clock.jpg
---
<!-- vale Vale.Spelling = NO -->
![a closeup of the Prague Astronomical Clock By EWilson (Volunteer) - Own work, CC BY-SA 4.0, https://commons.wikimedia.org/w/index.php?curid=115416822](./images/astro-clock.jpg)
<!-- vale Vale.Spelling = YES -->

I was inspired by some recent conversations with Corso users on Discord, and
this
[Reddit thread](https://www.reddit.com/r/Office365/comments/127rt5q/what_is_your_backup_schedule/),
to talk about the ideal cadence for backups.

## Why do we need backups again?

I know you’re here at the blog for Corso, a Microsoft 365 backup tool, so you
probably don’t need to be sold on the necessity of backups. But just as a
reminder, the
[Microsoft Shared Responsibility Model](https://www.veeam.com/blog/office365-shared-responsibility-model.html),
similar to that of all public cloud providers, means there’s a place where their
responsibility to help you with recovery stops.
<!-- truncate -->
The most common reasons people need a backup (based on the last few months’ discussion among Microsoft 365 admins) are:

- Malware, ransomware, or a similar attack
- Data lost in migration (for example employee leaving the org or changing roles)
- Accidental deletion

In all of these scenarios, Microsoft will take zero responsibility for restoring your data.

### What about the recycle bin?

If you've been pondering the same question, you're probably already aware that
Microsoft offers a few different recycle bin options, which can prove helpful in
the event of short-term, limited data loss. Even though this solution can
provide limited backup capabilities, it's far from perfect. Data in the recycle bin
gets automatically purged after a few days and malicious users can also force
early deletion of data residing in the recycle bin.

Further, the recycle bin can't provide the in-depth data control over important
business data that you need. To guarantee complete access and control of
important data, a comprehensive backup and disaster recovery plan is required.
This includes both short-term and long-term retention, and the ability to
recover in bulk, granularly, or from a particular point in time.

## How frequently should you back up?

Let’s start by defining your team’s *Recovery Point Objective (RPO).* RPO
generally refers to calculating how much
[data loss](https://www.acronis.com/products/cloud/cyber-protect/data-loss-prevention/)
a company can experience within a period most relevant to its business before
significant harm occurs, from the point of a disruptive event to the last data
backup.

RPO helps determine how much data a company can tolerate losing during an unforeseen event.

The ideal frequency of backups should be a business-level decision - what RPO
are you aiming for, any technical considerations will probably be secondary.

### Shouldn’t you back up continuously?

There have been a number of expensive backup tools in the past that offer
something like ‘continuous backups,’ where every single file change is reflected
in the backup almost instantly. This is a cool-sounding idea with some
significant drawbacks, namely:

- Without item versioning and/or preservation of older full backups, this model
  drastically increases the chances that your backups will be worthless: if data
  is accidentally corrupted, an extremely rapid backup will overwrite good
  backed up data with junk almost right away.
- If you want item versioning and extensive retention policies, the cost overheads for super-frequent backups can be prohibitive.

While backup frequency will vary with each business, it’s generally not the case
that a backup interval of “nearly 0ms” will make sense.

## Technical Considerations: Microsoft Graph State Tokens

One of the reasons to back up fairly frequently is the use of Microsoft Graph
state tokens to show what has changed about your data. For example Corso only
captures incremental changes during backups, only needing to store the items
that have been updated or added since the last backup. It does this using
[state tokens](https://learn.microsoft.com/en-us/graph/delta-query-overview#state-tokens)
that it stores within the Microsoft 365 infrastructure to checkpoint the end of
a backup. This token is used by the next backup invocation to see what has
changed, including deletions, within your data.

The exact expiry of state tokens isn’t published by Microsoft, but our
observations show that if you are only backing up every few days, these tokens
can expire. This will force a new full backup each time which is both
unnecessary and costly (in terms of time and bandwidth but not
storage because of Corso’s deduplication).

You can therefore reduce data transmission overhead, improve backup performance, and reduce RPO, by backing up more frequently.

## Cost Considerations: Storage Costs

With the threat of ransomware and other malicious data corruption, it’s a great
idea to store full backups with some frequency. This means that, if you want to
have frequent backups ****with**** retention of older versions, you’re going to
need a lot of storage.  

Tools like Corso that use S3 object storage will be cheaper than most others.
Further Corso not only deduplicates, packs, and compresses data, but it also has
a ton of smarts to only capture incremental changes between backups but always
present them as a full backup (topic for another blog post!). This ensures that
even if you have 1000s of backups per user or SharePoint site, you will always
see fast restores, minimal storage overhead, and always-full backups.

For other tools, you should evaluate if it uses storage efficiently:

- Since there’s a per-object storage cost with most S3 tiers, backups should bundle small items together
- Backups should include compression and de-duplication to be as small as possible

Take a look at some of our recent writing on selecting the best S3 storage tier
(spoiler warning it’s probably Glacier IR) for your S3 backups.

### You still haven’t answered my question: How often should you back up?

Independent of whether it's Microsoft 365 or other systems, at least once a
day. Probably about once every 8 hours. It will ensure your backups are
incremental and that you don’t lose too much work in the event of an incident.
Higher frequencies will be necessary for higher RPO goals.
