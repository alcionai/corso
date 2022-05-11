# Key Data Structures for MSGraph API

Microsoft Graph API is a mechanism for retrieving data from Microsoft 365
applications. When data is exporting into this application, the data is
transported in custom objects. This documentation will describe the structure
of these objects.

## Overview
Each Microsoft object is located in the msgraph
[repository](https://github.com/microsoftgraph/msgraph-sdk-go) within the model
directory. The structures described contain the highest level of granularity
for an object. For example, there exists a `MessageCollectionResponse` object;
however, the msgraph-sdk implements a pagination mechanism that allows for only
one message to be retrieved from the server at a time. Thus, the `Message`
structure is included in this section. The data structures for mail and
calendars are included.

## Mail Data Structure

Graph API states that all items of an individual’s emails are of type
`Message`. A message object "inherits" several other objects of note:
OutlookItem, Entity, and additional data. These are important because important
information and functionality are obfuscated in these structures. Each object
has an interface associated with them with an `*able.go` suffix (e.g
`entityable.go`). Again all of these objects are located in the `models`
directory.

Message:

```go
type Message struct {
    OutlookItem
    // The fileAttachment and itemAttachment attachments for the message.
    attachments []Attachmentable
    // The Bcc: recipients for the message.
    bccRecipients []Recipientable
    // The body of the message. It can be in HTML or text format. Find out about safe HTML in a message body.
    body ItemBodyable
    // The first 255 characters of the message body. It is in text format.
    bodyPreview *string
    // The Cc: recipients for the message.
    ccRecipients []Recipientable
    // The ID of the conversation the email belongs to.
    conversationId *string
    // Indicates the position of the message within the conversation.
    conversationIndex []byte
    // The collection of open extensions defined for the message. Nullable.
    extensions []Extensionable
    // The flag value that indicates the status, start date, due date, or completion date for the message.
    flag FollowupFlagable
    // The owner of the mailbox from which the message is sent. In most cases, this value is the same as the sender property, except for sharing or delegation scenarios. The value must correspond to the actual mailbox used. Find out more about setting the from and sender properties of a message.
    from Recipientable
    // Indicates whether the message has attachments. This property doesn't include inline attachments, so if a message contains only inline attachments, this property is false. To verify the existence of inline attachments, parse the body property to look for a src attribute, such as <IMG src='cid:image001.jpg@01D26CD8.6C05F070'>.
    hasAttachments *bool
    // The importance property
    importance *Importance
    // The inferenceClassification property
    inferenceClassification *InferenceClassificationType
    // The internetMessageHeaders property
    internetMessageHeaders []InternetMessageHeaderable
    // The internetMessageId property
    internetMessageId *string
    // The isDeliveryReceiptRequested property
    isDeliveryReceiptRequested *bool
    // The isDraft property
    isDraft *bool
    // The isRead property
    isRead *bool
    // The isReadReceiptRequested property
    isReadReceiptRequested *bool
    // The collection of multi-value extended properties defined for the message. Nullable.
    multiValueExtendedProperties []MultiValueLegacyExtendedPropertyable
    // The parentFolderId property
    parentFolderId *string
    // The receivedDateTime property
    receivedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The replyTo property
    replyTo []Recipientable
    // The sender property
    sender Recipientable
    // The sentDateTime property
    sentDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The collection of single-value extended properties defined for the message. Nullable.
    singleValueExtendedProperties []SingleValueLegacyExtendedPropertyable
    // The subject property
    subject *string
    // The toRecipients property
    toRecipients []Recipientable
    // The uniqueBody property
    uniqueBody ItemBodyable
    // The webLink property
    webLink *string
}
```

Specifically, `OutlookItem` allows for the tracking of changes to items as well
as the category that it associated with the object. **NOTE:** `parentFolderId`
is associated with the folder such as `inbox` rather than an individual user or
tenant. Additional details to be found in [Transport](msgraphTransport.md)
section.

## Calendar data structure

The calendar object follows a similar pattern as the Message. A chief
difference is the `OutlookItem` is not associated with this object which tracks
changes. Calendar events are resent when the objects are changed.

```go
type Calendar struct {
    Entity
    // Represent the online meeting service providers that can be used to create online meetings in this calendar. Possible values are: unknown, skypeForBusiness, skypeForConsumer, teamsForBusiness.
    allowedOnlineMeetingProviders []OnlineMeetingProviderType
    // The permissions of the users with whom the calendar is shared.
    calendarPermissions []CalendarPermissionable
    // The calendar view for the calendar. Navigation property. Read-only.
    calendarView []Eventable
    // true if the user can write to the calendar, false otherwise. This property is true for the user who created the calendar. This property is also true for a user who has been shared a calendar and granted write access.
    canEdit *bool
    // true if the user has the permission to share the calendar, false otherwise. Only the user who created the calendar can share it.
    canShare *bool
    // true if the user can read calendar items that have been marked private, false otherwise.
    canViewPrivateItems *bool
    // Identifies the version of the calendar object. Every time the calendar is changed, changeKey changes as well. This allows Exchange to apply changes to the correct version of the object. Read-only.
    changeKey *string
    // Specifies the color theme to distinguish the calendar from other calendars in a UI. The property values are: auto, lightBlue, lightGreen, lightOrange, lightGray, lightYellow, lightTeal, lightPink, lightBrown, lightRed, maxColor.
    color *CalendarColor
    // The default online meeting provider for meetings sent from this calendar. Possible values are: unknown, skypeForBusiness, skypeForConsumer, teamsForBusiness.
    defaultOnlineMeetingProvider *OnlineMeetingProviderType
    // The events in the calendar. Navigation property. Read-only.
    events []Eventable
    // The calendar color, expressed in a hex color code of three hexadecimal values, each ranging from 00 to FF and representing the red, green, or blue components of the color in the RGB color space. If the user has never explicitly set a color for the calendar, this property is empty. Read-only.
    hexColor *string
    // true if this is the default calendar where new events are created by default, false otherwise.
    isDefaultCalendar *bool
    // Indicates whether this user calendar can be deleted from the user mailbox.
    isRemovable *bool
    // Indicates whether this user calendar supports tracking of meeting responses. Only meeting invites sent from users' primary calendars support tracking of meeting responses.
    isTallyingResponses *bool
    // The collection of multi-value extended properties defined for the calendar. Read-only. Nullable.
    multiValueExtendedProperties []MultiValueLegacyExtendedPropertyable
    // The calendar name.
    name *string
    // If set, this represents the user who created or added the calendar. For a calendar that the user created or added, the owner property is set to the user. For a calendar shared with the user, the owner property is set to the person who shared that calendar with the user.
    owner EmailAddressable
    // The collection of single-value extended properties defined for the calendar. Read-only. Nullable.
    singleValueExtendedProperties []SingleValueLegacyExtendedPropertyable
}
```

## Working with Attachments

There is an open
[issue](https://github.com/microsoftgraph/msgraph-sdk-go-core/issues/12) that
states the Graph API is unable to upload large data sets. The largest parts of
a message or event are often attachments. MSFT
[docs](https://docs.microsoft.com/en-us/graph/sdks/large-file-upload?tabs=csharp)
show how to this feature is supposed to work for OneDrive and Messages. The API
does support the mechanism of attaching files. However, it is unclear which
part of the API is missing/broken. Further investigation required.

Attachments
```go
// Attachment
type Attachment struct {
    Entity
    // The MIME type.
    contentType *string
    // true if the attachment is an inline attachment; otherwise, false.
    isInline *bool
    // The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The attachment's file name.
    name *string
    // The length of the attachment in bytes.
    size *int32
```
