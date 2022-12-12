---
slug: corso-announcement-free-backup-for-microsoft-365
title: "A Home to Call Your Own: The Need for Your Own Backups in IT"
description: "Announcing Corso, a free, open-source, and secure backup tool for Microsoft 365."
authors: nica
tags: [corso, microsoft 365]
image: ./images/office_desk.jpg
---

![Office desk](./images/office_desk.jpg)

Have you had it with Google sheets? Me too! Excel is my home. It’s where I write all my best formulae. And what
about PowerPoint? The way it just finds stock photos for you? The automatic ‘alternative designs for this slide’
button? It’s too good. I can’t give up Microsoft 365.

If you did some work today, there’s a good chance you opened a Microsoft tool. M365 is used by
[more than a million](https://www.statista.com/statistics/983321/worldwide-office-365-user-numbers-by-country/)
companies worldwide, and nearly 880,000 companies in the U.S. use the software suite. But with that widespread usage
comes risk, business-critical data is at risk of loss or corruption, if not securely backed up and protected.

<!-- truncate -->

## The problem with built-in backups

A couple of years back I took the time to get the AWS ‘baby cert,’ their first certification. The focus of the
material surprised me, along with learning about the most popular AWS products and their benefits, you had to learn
cold all the things that AWS won't do for you. AWS won’t alert you to poor performance on your applications. It
won’t automatically scale down your instances. And while AWS and the other public cloud providers completely meet their
promised SLA, they don't promise to deliver the backups that you expect them to deliver.

“I accidentally deleted the customer DB,” isn’t a situation that public cloud companies are built to prevent or
ameliorate. Fundamentally, on all public clouds, backups are a shared responsibility between administrator and service.
For another example, see this from Microsoft's Diana Kelley:
["Driving data security is a shared responsibility, here’s how you can protect yourself"](https://www.statista.com/statistics/983321/worldwide-office-365-user-numbers-by-country/)

## Data Loss Happens

Let’s talk about the stats on data loss:

Data loss can result from accidental or intentional deletion, cyber-attacks and malware, a poorly executed migration,
or the cancellation of a software license, among other reasons. For example,
[2 out of 5 servers](https://www.veeam.com/blog/data-loss-2022.html) had at least one or more outages over the
past 12 months. And cybercrime is on a continual rise; the average number of data breaches and cyberattacks were
[up 15.1% in 2021](https://www.forbes.com/sites/chuckbrooks/2022/06/03/alarming-cyber-statistics-for-mid-year-2022-that-you-need-to-know/?sh=642204357864),
compared with the previous year. As of 2022, the average cost of a data breach in the U.S. was $9.44 million.

## The limitations of backup tools

While many backup and recovery solutions exist for M365 data, many lack comprehensive workflows, are hard to use and
implement, or are cost prohibitive. Yet reliable and regular backups are critical for ensuring your data is protected
and always available, in case of a breach, unexpected server failure or failed data migration.

And that’s why we built Corso, a free and secure, 100% open-source solution that protects M365 data by securely and
efficiently backing up all business-critical data to object storage.

## Why Corso?

<!-- vale alex.Condescending = NO -->

Corso is purpose-built for protection of your M365 organization account (this tool doesn’t work with consumer accounts)
with easy-to-use comprehensive backup and restore workflows that reduce backup time and administrative overhead,
improve time-to-recovery, and replace unreliable scripts or workarounds. It enables high-throughput, high-tolerance
backups that feature end-to-end encryption, deduplication, and compression. Plus, it’s compatible with any S3-compliant
object storage system: AWS S3 (including Glacier Instant Access), Google Cloud Storage and Backblaze. (Azure Blob
support is coming soon).

<!-- vale alex.Condescending = YES -->

Corso’s secure backup protects against accidental data loss, service provider downtime and malicious threats, including
ransomware attacks. Plus, a robust user community provides a venue for admins to share and learn about data protection
and find best practices for how to securely configure their M365 environments. As a member of the community, you’ll
have access to blogs, forums, and discussion, as well as updates on public and feedback-driven development.
[Join the Corso community on Discord](https://discord.gg/63DTTSnuhT).

## Low-Cost and Highly Secure

Corso's source code is licensed under the Apache v2 open-source license. It’s open source, and it’s free, which makes
it the perfect solution for cost-conscious teams. And that’s not where the cost savings end, Corso’s flexible retention
policies and ability to compress and deduplicate data efficiently before sending it to storage, helps reduce storage
costs, as well.

## Interested in Trying Corso?

<!-- vale Microsoft.Contractions = NO -->

Corso, currently in alpha, provides a CLI-based tool for backups of your M365 data.
[Follow the quickstart guide](../../docs/quickstart) to start protecting your business-critical M365 data in
just a few minutes. Because Corso is currently in alpha, it should **NOT** be used in production.

<!-- vale Microsoft.Contractions = YES -->

Corso supports Microsoft 365 Exchange and OneDrive, with SharePoint and Teams support in active development. Coverage
for more services, beyond M365, will expand based on the interests and needs of the community.

Your feedback is critical for our work on this tool! Please
[tell us what you think of Corso](https://discord.gg/63DTTSnuhT).
