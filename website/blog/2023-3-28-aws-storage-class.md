---
slug: aws-storage-class
title: "Choosing the Best AWS S3 Storage Class for Corso Backups"
description: "By watching my
four short videos, you’ll be set up and on your way to backing up all your Microsoft 365
data in less than 20 minutes. That’s shorter than most meetings!"
authors: nica
tags: [corso, microsoft 365, AWS, backups]
date: 2023-3-28
image: ./images/box_sizes.jpeg
---

![multiple box sizes](./images/box_sizes.jpeg)
Recently when writing
about the storage options for Corso, I found myself going pretty far in the
weeds on storage classes in S3. I thought I’d make a list of all the storage
options and why they might, or might not, work for backups.
<!-- truncate -->

## First, some assumptions

If we’re talking about backups, we're assuming far more writes than read, and
that most objects that are written will never be read.

Increasing age of an object increases the chances that it will never be read.

And we can’t afford to lose anything! One-zone options that carry a small
chance of data loss like OneZone_IA won't work.  

Finally, there will be index data and metadata that may well be overwritten
frequently. For more detail on this, and an exploration of non-AWS alternatives
to S3, see our past article on
[where to store your Corso data](https://corsobackup.io/blog/where-to-store-corso/).

If your backup solution breaks one of these expectations, for example if you’re
restoring from backups every single day, the advice in this article may not be
applicable to you.

## Best practices no matter your storage class

Using a tool for backups rather than a naive file copy process is the first step
towards an efficient backup process. Before you drag that folder over to that
network drive icon, consider the following requirements:

- Compression - don’t use more network bandwidth than you have to
- De-duplication - backing up a team’s email shouldn’t mean storing 50 identical copies of ‘Presentation_FINAL.pptx’
- Incremental Backups - Ideally, your second backup should only include updated objects
- Bundling - creating millions of 2kb objects each backup is going to add to costs and hurt performance

## Storage Classes, considered

The AWS Storage classes are STANDARD | REDUCED_REDUNDANCY | STANDARD_IA |
ONEZONE_IA | INTELLIGENT_TIERING | GLACIER | DEEP_ARCHIVE | OUTPOSTS |
GLACIER_IR

of which we won’t consider REDUCED_REDUNDANCY (it’s outdated and Standard is now
cheaper) and OUTPOSTS (if you need on-prem S3, it’s not for cost or efficiency).

### STANDARD

The S3 Standard storage should work for all backup implementations, as long as
you’re not using something that can’t really work with object storage with
network latency (for example your backup application is trying to do fine-grained
low-latency database-style queries using indices stored in S3).

For Corso, Standard is a great place to start testing your setup, letting you
perform regular backups, restores, and deletions. We also recommend storing all
your non-blob data in Standard, how to do this automatically is covered at the
end of this list.

### STANDARD_IA and ONEZONE_IA

These are the storage classes AWS recommends for backups! But it’s likely that
Glacier Instant Retrieval will be cheaper. Also, Infrequent Access charges a
minimum storage size of 128KB and a minimum storage time of 30 days. If your
backups are creating many small objects, or if you have incremental backups
constantly updating most objects, Infrequent Access may come out more expensive
than standard.

For Corso, it’s not likely that this storage class will make the most sense.
Maybe a case where periodic restores are expected with some frequency would
benefit from this class, but that would have to be so frequent I’m not sure
‘backup’ is the right term. If you found this was the best class for you please
join our Discord and tell us about it.
<!-- vale Vale.Spelling = NO -->
### INTELLIGENT_TIERING

Intelligent Tiering is the most appealing of AWS’s new S3 offerings for backups.
As objects age they’ll move down to cheaper and cheaper, finally dropping to the
same storage costs per GB as Glacier Instant Retrieval.

Two considerations should give you pause when using Intelligent Tiering for backups: first
there’s a small compute cost to Intelligent Tiering, and second you probably
do know the usage pattern of these backups: almost all will never be touched.

With Intelligent Tiering you’ll pay for your backups to be on a more expensive
tier for 60 days before you get the pricing that you probably could have picked
out for yourself in the first place.

Intelligent Tiering probably only makes sense if you’re using backups in a
nonstandard way, for example restoring from backups every morning. If you’re not sure
*how* your data will be used, Intelligent Tiering is a safe bet.
<!-- vale Vale.Spelling = YES -->

### GLACIER and DEEP_ARCHIVE

Glacier (not Instant retrieval, which is discussed below) is a great way to
archive data, which is a slightly different idea than backups. If you have a
reason to store and not touch data (for example, for compliance) and can tolerate
extremely high latencies (hours to days) for recovery you may want to use
Glacier Archive. However, high-performance backup tools, like Corso, usually
contain smart optimizations like incremental backups and backup indexes that
won’t work if the latency for all requests is in minutes. Further, for cost and
efficiency, deduplicating object stores will often compact data as the primary
data source churns. Using default Glacier or Glacier Deep Archive is a poor fit
for that workload.

### GLACIER_IR

Likely to be your best option for backups Glacier IR is cheaper than any
non-glacier option for storage costs, with a low request latency. Corso’s
de-duplication, bundling, and compression will help ensure that you’re paying as
little as possible for storage.

## Glacier Instant Retrieval is the best choice for Corso backups

With these considerations, and with the best practices mentioned above, you
should be able to build reliable backups with a minimal cost impact. If you’re
ready to give Corso a try, check out our
[Quickstart Guide](https://corsobackup.io/docs/quickstart/), or take a look at a
recent article on backing up
[large Exchange instances](https://corsobackup.io/blog/large-microsoft-365-exchange-backups/)
with Corso.
