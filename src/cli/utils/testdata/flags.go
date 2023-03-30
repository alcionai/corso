package testdata

import "strings"

func FlgInpts(in []string) string { return strings.Join(in, ",") }

var (
	BackupInpt = "backup-id"

	UsersInpt  = []string{"users1", "users2"}
	SiteIDInpt = []string{"siteID1", "siteID2"}
	WebURLInpt = []string{"webURL1", "webURL2"}

	ContactInpt     = []string{"contact1", "contact2"}
	ContactFldInpt  = []string{"contactFld1", "contactFld2"}
	ContactNameInpt = "contactName"

	EmailInpt               = []string{"mail1", "mail2"}
	EmailFldInpt            = []string{"mailFld1", "mailFld2"}
	EmailReceivedAfterInpt  = "mailReceivedAfter"
	EmailReceivedBeforeInpt = "mailReceivedBefore"
	EmailSenderInpt         = "mailSender"
	EmailSubjectInpt        = "mailSubjet"

	EventInpt             = []string{"event1", "event2"}
	EventCalInpt          = []string{"eventCal1", "eventCal2"}
	EventOrganizerInpt    = "eventOrganizer"
	EventRecursInpt       = "eventRecurs"
	EventStartsAfterInpt  = "eventStartsAfter"
	EventStartsBeforeInpt = "eventStartsBefore"
	EventSubjectInpt      = "eventSubject"

	LibraryInpt            = "library"
	FileNamesInpt          = []string{"fileName1", "fileName2"}
	FolderPathsInpt        = []string{"folderPath1", "folderPath2"}
	FileCreatedAfterInpt   = "fileCreatedAfter"
	FileCreatedBeforeInpt  = "fileCreatedBefore"
	FileModifiedAfterInpt  = "fileModifiedAfter"
	FileModifiedBeforeInpt = "fileModifiedBefore"

	ListFolderInpt = []string{"listFolder1", "listFolder2"}
	ListItemInpt   = []string{"listItem1", "listItem2"}

	PageFolderInpt = []string{"pageFolder1", "pageFolder2"}
	PageInpt       = []string{"page1", "page2"}
)
