package exchange

// exchange_vars.go is package level collection of interfaces and
// constants that are used within the exchange.

// Legacy Value Tags and constants are used to override certain values within
// M365 objects.
// Master Property Value Document:
//  https://interoperability.blob.core.windows.net/files/MS-OXPROPS/%5bMS-OXPROPS%5d.pdf
const (
	// MailRestorePropertyTag inhibits exchange.Mail.Message from being "resent" through the server.
	// DEFINED: Section 2.791 PidTagMessageFlags
	MailRestorePropertyTag = "Integer 0x0E07"

	// RestoreCanonicalEnableValue marks message as sent via RopSubmitMessage
	// Defined: https://interoperability.blob.core.windows.net/files/MS-OXCMSG/%5bMS-OXCMSG%5d.pdf
	// Section: 2.2.1.6 PidTagMessageFlags Property
	// Additional Information: https://docs.microsoft.com/en-us/office/client-developer/outlook/mapi/pidtagmessageflags-canonical-property
	RestoreCanonicalEnableValue = "4"

	// MailSendTimeOverrideProperty allows for send time to be updated.
	// Section: 2.635 PidTagClientSubmitTime
	MailSendDateTimeOverrideProperty = "SystemTime 0x0039"

	// MailReceiveDateTimeOverrideProperty allows receive date time to be updated.
	// Section: 2.789 PidTagMessageDeliveryTime
	MailReceiveDateTimeOverriveProperty = "SystemTime 0x0E06"
)

// descendable represents objects that implement msgraph-sdk-go/models.entityable
// and have the concept of a "parent folder".
type descendable interface {
	GetId() *string
	GetParentFolderId() *string
}

// displayable represents objects that implement msgraph-sdk-fo/models.entityable
// and have the concept of a display name.
type displayable interface {
	GetId() *string
	GetDisplayName() *string
}
