# Fault tolerance

Given the millions of objects found in a typical Microsoft 365 tenant,
Corso is optimized for high-performance processing, hardened to
tolerate transient failures and, most importantly, able to restart backups.

Corso’s fault-tolerance architecture is motivated by Microsoft’s Graph
API variable performance and throttling. Corso follows Microsoft’s
recommend best practices (for example, [correctly decorating API
traffic](https://learn.microsoft.com/en-us/sharepoint/dev/general-development/how-to-avoid-getting-throttled-or-blocked-in-sharepoint-online#how-to-decorate-your-http-traffic))
and, in addition, implements a number of optimizations to improve
backup and restore reliability.

## Recovery from transient failures

Corso, at the HTTP layer, will retry requests (after a HTTP timeout,
for example) and will respect Graph API’s directives such as the
`retry-after` header to backoff when needed. This allows backups to
succeed in the face of transient or temporary failures.

## Restarting from permanent API failures

The Graph API can, for internal reasons, exhibit extended periods of
failures for particular Graph objects. In this scenario, bounded retries
will be ineffective. Unless invoked with the
on-error recover option, Corso will skip over these failing objects. For
backups, it will move forward with backing up other objects belonging
to the user and, for restores, it will continue with trying to restore
any remaining objects. If a multi-user backed is in progress (via `*`
or by specifying multiple users with the `—user` argument), Corso will
also continue processing backups for the remaining users. In both
cases, Corso will exit with a non-zero exit code to reflect incomplete
backups or restores.

On subsequent backup attempts, Corso will try to
minimize the work involved. If the previous backup was successful and
Corso’s stored state tokens haven’t expired, it will use [delta
queries](https://learn.microsoft.com/en-us/graph/delta-query-overview),
wherever supported, to perform incremental backups.

If the previous backup for a user had resulted in a failure, Corso
uses a variety of fallback mechanisms to reduce the amount of data
downloaded and reduce the number of objects enumerated. For example, with
OneDrive, Corso won't redo downloads of data from Microsoft 365 or
uploads of data to the Corso repository if it had successfully backed
up that OneDrive file as a part of a previously incomplete and failed
backup. Even if the Graph API might not allow Corso to skip
downloading data, Corso can still skip another upload it to the repository.
