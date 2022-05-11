# Restoring Data from MSGraph

The process of restoring information back to Microsoft applications greatly
differs upon the application. If objects are migrated back into an application
as described within the [Transport](msgraphTransport.md) document. This document will describe how
items are migrated back into their individual applications.

# Microsoft Exchange

## Mailbox Messages
As of this time, it does not appear msgraph possesses the ability to "restore"
a message that was previously deleted from the Exchange server. While there are
ways to migrate messages from the `trash can` back to the inbox, the `restore`
capability within msgraph does not appear to exist.

## Calendar Events

Calendar events are not  modifiable. Therefore, restoring calendar data could
be done by creating an
[event](https://docs.microsoft.com/en-us/graph/api/calendar-post-events?view=graph-rest-1.0&tabs=go)
with the same details of the original event. However, this may not an
acceptable work around.

## Summary
There are limitations for the msgraph platform as it comes to placing data back
into applications. While there are many applications for the Exchange
application that would allow for data to be exported, the import function
appears well guarded. A larger discussion after we have verified which style of
application "restores" data in a way that bests supports our product
assertions.
