package exchange

// exchange_vars.go is package level collection of interfaces and
// constants that are used within the exchange.

<<<<<<< Updated upstream
const(
    	// RestorePropertyTag defined:
	// https://docs.microsoft.com/en-us/office/client-developer/outlook/mapi/pidtagmessageflags-canonical-property
	RestorePropertyTag          = "Integer 0x0E07"
	RestoreCanonicalEnableValue = "4"
    )
=======
const (
	// RestorePropertyTag defined:
	// https://docs.microsoft.com/en-us/office/client-developer/outlook/mapi/pidtagmessageflags-canonical-property
	RestorePropertyTag          = "Integer 0x0E07"
	RestoreCanonicalEnableValue = "4"
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
>>>>>>> Stashed changes
