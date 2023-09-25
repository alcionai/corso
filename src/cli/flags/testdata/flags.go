package testdata

import (
	"strings"

	"github.com/spf13/cobra"
)

func FlgInputs(in []string) string { return strings.Join(in, ",") }

var (
	BackupInput = "backup-id"

	GroupsInput  = []string{"team1", "group2"}
	MailboxInput = []string{"mailbox1", "mailbox2"}
	UsersInput   = []string{"users1", "users2"}
	SiteIDInput  = []string{"siteID1", "siteID2"}
	WebURLInput  = []string{"webURL1", "webURL2"}

	ExchangeCategoryDataInput   = []string{"email", "events", "contacts"}
	SharepointCategoryDataInput = []string{"files", "lists", "pages"}
	GroupsCategoryDataInput     = []string{"files", "lists", "pages", "messages"}

	ContactInput     = []string{"contact1", "contact2"}
	ContactFldInput  = []string{"contactFld1", "contactFld2"}
	ContactNameInput = "contactName"

	EmailInput               = []string{"mail1", "mail2"}
	EmailFldInput            = []string{"mailFld1", "mailFld2"}
	EmailReceivedAfterInput  = "mailReceivedAfter"
	EmailReceivedBeforeInput = "mailReceivedBefore"
	EmailSenderInput         = "mailSender"
	EmailSubjectInput        = "mailSubject"

	EventInput             = []string{"event1", "event2"}
	EventCalInput          = []string{"eventCal1", "eventCal2"}
	EventOrganizerInput    = "eventOrganizer"
	EventRecursInput       = "eventRecurs"
	EventStartsAfterInput  = "eventStartsAfter"
	EventStartsBeforeInput = "eventStartsBefore"
	EventSubjectInput      = "eventSubject"

	LibraryInput            = "library"
	FileNameInput           = []string{"fileName1", "fileName2"}
	FolderPathInput         = []string{"folderPath1", "folderPath2"}
	FileCreatedAfterInput   = "fileCreatedAfter"
	FileCreatedBeforeInput  = "fileCreatedBefore"
	FileModifiedAfterInput  = "fileModifiedAfter"
	FileModifiedBeforeInput = "fileModifiedBefore"

	ListFolderInput = []string{"listFolder1", "listFolder2"}
	ListItemInput   = []string{"listItem1", "listItem2"}

	PageFolderInput = []string{"pageFolder1", "pageFolder2"}
	PageInput       = []string{"page1", "page2"}

	Collisions      = "collisions"
	Destination     = "destination"
	ToResource      = "toResource"
	SkipPermissions = false

	DeltaPageSize = "7"

	Archive    = true
	FormatType = "json"

	AzureClientID     = "testAzureClientId"
	AzureTenantID     = "testAzureTenantId"
	AzureClientSecret = "testAzureClientSecret"

	AWSAccessKeyID     = "testAWSAccessKeyID"
	AWSSecretAccessKey = "testAWSSecretAccessKey"
	AWSSessionToken    = "testAWSSessionToken"

	CorsoPassphrase = "testCorsoPassphrase"

	RestoreDestination = "test-restore-destination"

	FetchParallelism = "3"

	FailFast                  = true
	DisableIncrementals       = true
	ForceItemDataDownload     = true
	DisableDelta              = true
	EnableImmutableID         = true
	DisableConcurrencyLimiter = true
)

func WithFlags(
	cc *cobra.Command,
	command string,
	flagSets ...[]string,
) {
	args := []string{command}

	for _, sl := range flagSets {
		args = append(args, sl...)
	}

	cc.SetArgs(args)
}
