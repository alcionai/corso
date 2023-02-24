---
slug: product-hunt-welcome
title: "Welcome Product Hunt community! We launched Corso 3 months ago, here’s what we learned"
description: "Corso is a free and open-source tool to back up your Microsoft 365 data. For Microsoft Office 365 administrators, it backs up user data to any S3 compliant storage. Leaving you in control of when your data is backed up, and how it’s restored."
authors: nica
tags: [corso, microsoft 365, backups]
date: 2023-2-25
image: ./images/rocket.png
---

![SpaceX Falcon Heavy Launch, https://unsplash.com/photos/OHOU-5UVIYQ](./images/rocket.png)
Corso is a free and open-source tool to back up your Microsoft 365 or Office 365 data. For
Microsoft Office 365 administrators, it backs up user data to any S3-compliant
storage. Leaving you in control of when your data is backed up, and how it’s
restored.

<!-- truncate -->

### We love our users

Since our release in late 2022, we’ve been really happy to see just how many people
have given Corso a try. Almost every day we chat on
[Corso discord](https://discord.gg/63DTTSnuhT), with users trying out Corso.
While small teams have used Corso for their Microsoft 365 backups, we’ve also been
pleasantly surprised to see some larger groups evaluating Corso. Mainly Managed
Service Providers (MSPs) who want to back up Microsoft 365 instances with thousands of
users!

### Simplicity is key

One of the most gratifying things working on Corso, especially working in
Developer Relations as I do, is having a tool be straightforward enough for people to just
pick it up and use it. To see that most of the people who look at our
[Quick Start](https://corsobackup.io/docs/quickstart/) guide end up successfully
backing up their Microsoft 365 data and hearing extremely positive usability feedback is
hugely gratifying. I’ve been able to make
[tutorial videos](https://www.youtube.com/watch?v=mlwfEbPqD94&list=PLSukexZlj1V0D0xGV2ON-MWRmPpLWi6hK)
that go a bit deeper into specific use cases, rather than having to try and show
every new user how to get the tool working.

## Key Lessons

Bigger Microsoft 365 domains brought use cases that we weren’t expecting. And
early users helped find these problems before they affected more users. If you
look at our
[releases page on GitHub](https://github.com/alcionai/corso/releases/), most of
the bug fixes we’ve released in the last two months have been things that were
initially reported by our users. That’s so cool! Corso is rapidly approaching a v1.0 release and
that has been enabled by the community of system administrators and Microsoft 365 admins who were
willing to test out early versions of the tool. These users taught us great things and we will
always be indebted to them. The key lessons, described below, that we learned from their
experiences should make your Corso usage much smoother.

### Microsoft 365 domains come in a variety of shapes

Inherently, a backup tool for Microsoft Office 365 is going to be directly
dependent on the Microsoft Graph API. While the graph API is uniform, Microsoft
365 has gone through a number of iterations over the years resulting in bespoke
configuration. A recent bug we investigated involved users that had two
calendars with identical names. Normally Microsoft doesn’t allow identical calendar names,
but this could happen as part of a tenant migration. This was an edge case that we
weren’t expecting! The community helped us troubleshoot these bespoke
configurations that would otherwise be impossible to recreate for testing.

### Early adopters are pushing Corso to significant scale

While we push the Corso in testing to non-trivial scale with synthetic data,
real world usage by early adopters pushed us to front-load
scalability work. We now have users who have used Corso with 100GB+ mailboxes
and 1TB+ OneDrive deployments. This led us to improve the robustness of retries
and, as almost always happens when finding new problems; we tweaked the logs to
make finding future issues easier.

## How you can try out Corso and get a gift

If you’d like to give Corso a try, start at our
[Quick Start guide](https://corsobackup.io/docs/quickstart/) that will get your
first backup running in just a few minutes. If you have any trouble, or just
want to chat with the tool’s creators,
[join our Discord](https://discord.gg/63DTTSnuhT) to get support and give
feedback.

Finally, we’d like to offer something to everyone who gives Corso a try, so once
you’ve used the tool
[fill out our feedback](https://forms.microsoft.com/r/mRVNKqeKDp) form to get
some swag. For the next month we’ll also be selecting one response from our
feedback form to send a Microsoft Zune (yes, THAT Zune, we got one restored).

Thank you again to the Product Hunt community for checking out
Corso. And a big thanks to all our existing users who have helped us take the
tool so far, so fast. We appreciate you!
