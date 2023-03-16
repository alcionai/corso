package exchange

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/clues"
	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/connector/exchange/api"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/mockconnector"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
)

type ExchangeRestoreSuite struct {
	tester.Suite
	gs          graph.Servicer
	credentials account.M365Config
	ac          api.Client
}

func TestExchangeRestoreSuite(t *testing.T) {
	suite.Run(t, &ExchangeRestoreSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tester.M365AcctCredEnvs},
		),
	})
}

func (suite *ExchangeRestoreSuite) SetupSuite() {
	t := suite.T()

	a := tester.NewM365Account(t)
	m365, err := a.M365Config()
	require.NoError(t, err, clues.ToCore(err))

	suite.credentials = m365
	suite.ac, err = api.NewClient(m365)
	require.NoError(t, err, clues.ToCore(err))

	adpt, err := graph.CreateAdapter(m365.AzureTenantID, m365.AzureClientID, m365.AzureClientSecret)
	require.NoError(t, err, clues.ToCore(err))

	suite.gs = graph.NewService(adpt)
}

// TestRestoreContact ensures contact object can be created, placed into
// the Corso Folder. The function handles test clean-up.
func (suite *ExchangeRestoreSuite) TestRestoreContact() {
	ctx, flush := tester.NewContext()
	defer flush()

	var (
		t          = suite.T()
		userID     = tester.M365UserID(t)
		now        = time.Now()
		folderName = "TestRestoreContact: " + common.FormatSimpleDateTime(now)
	)

	aFolder, err := suite.ac.Contacts().CreateContactFolder(ctx, userID, folderName)
	require.NoError(t, err, clues.ToCore(err))

	folderID := ptr.Val(aFolder.GetId())

	defer func() {
		// Remove the folder containing contact prior to exiting test
		err = suite.ac.Contacts().DeleteContainer(ctx, userID, folderID)
		assert.NoError(t, err, clues.ToCore(err))
	}()

	info, err := RestoreExchangeContact(
		ctx,
		mockconnector.GetMockContactBytes("Corso TestContact"),
		suite.gs,
		control.Copy,
		folderID,
		userID)
	assert.NoError(t, err, clues.ToCore(err))
	assert.NotNil(t, info, "contact item info")
}

// TestRestoreEvent verifies that event object is able to created
// and sent into the test account of the Corso user in the newly created Corso Calendar
func (suite *ExchangeRestoreSuite) TestRestoreEvent() {
	ctx, flush := tester.NewContext()
	defer flush()

	var (
		t      = suite.T()
		userID = tester.M365UserID(t)
		name   = "TestRestoreEvent: " + common.FormatSimpleDateTime(time.Now())
	)

	calendar, err := suite.ac.Events().CreateCalendar(ctx, userID, name)
	require.NoError(t, err, clues.ToCore(err))

	calendarID := ptr.Val(calendar.GetId())

	defer func() {
		// Removes calendar containing events created during the test
		err = suite.ac.Events().DeleteContainer(ctx, userID, calendarID)
		assert.NoError(t, err, clues.ToCore(err))
	}()

	info, err := RestoreExchangeEvent(ctx,
		mockconnector.GetMockEventWithAttendeesBytes(name),
		suite.gs,
		control.Copy,
		calendarID,
		userID,
		fault.New(true))
	assert.NoError(t, err, clues.ToCore(err))
	assert.NotNil(t, info, "event item info")
}

type containerDeleter interface {
	DeleteContainer(context.Context, string, string) error
}

// TestRestoreExchangeObject verifies path.Category usage for restored objects
func (suite *ExchangeRestoreSuite) TestRestoreExchangeObject() {
	t := suite.T()
	a := tester.NewM365Account(t)
	m365, err := a.M365Config()
	require.NoError(t, err, clues.ToCore(err))

	service, err := createService(m365)
	require.NoError(t, err, clues.ToCore(err))

	deleters := map[path.CategoryType]containerDeleter{
		path.EmailCategory:    suite.ac.Mail(),
		path.ContactsCategory: suite.ac.Contacts(),
		path.EventsCategory:   suite.ac.Events(),
	}

	userID := tester.M365UserID(suite.T())
	now := time.Now()
	tests := []struct {
		name        string
		bytes       []byte
		category    path.CategoryType
		destination func(*testing.T, context.Context) string
	}{
		{
			name:     "Test Mail",
			bytes:    mockconnector.GetMockMessageBytes("Restore Exchange Object"),
			category: path.EmailCategory,
			destination: func(t *testing.T, ctx context.Context) string {
				folderName := "TestRestoreMailObject: " + common.FormatSimpleDateTime(now)
				folder, err := suite.ac.Mail().CreateMailFolder(ctx, userID, folderName)
				require.NoError(t, err, clues.ToCore(err))

				return ptr.Val(folder.GetId())
			},
		},
		{
			name:     "Test Mail: One Direct Attachment",
			bytes:    mockconnector.GetMockMessageWithDirectAttachment("Restore 1 Attachment"),
			category: path.EmailCategory,
			destination: func(t *testing.T, ctx context.Context) string {
				folderName := "TestRestoreMailwithAttachment: " + common.FormatSimpleDateTime(now)
				folder, err := suite.ac.Mail().CreateMailFolder(ctx, userID, folderName)
				require.NoError(t, err, clues.ToCore(err))

				return ptr.Val(folder.GetId())
			},
		},
		{
			name:     "Test Mail: Item Attachment_Event",
			bytes:    mockconnector.GetMockMessageWithItemAttachmentEvent("Event Item Attachment"),
			category: path.EmailCategory,
			destination: func(t *testing.T, ctx context.Context) string {
				folderName := "TestRestoreEventItemAttachment: " + common.FormatSimpleDateTime(now)
				folder, err := suite.ac.Mail().CreateMailFolder(ctx, userID, folderName)
				require.NoError(t, err, clues.ToCore(err))

				return ptr.Val(folder.GetId())
			},
		},
		{
			name:     "Test Mail: Item Attachment_Mail",
			bytes:    mockconnector.GetMockMessageWithItemAttachmentMail("Mail Item Attachment"),
			category: path.EmailCategory,
			destination: func(t *testing.T, ctx context.Context) string {
				folderName := "TestRestoreMailItemAttachment: " + common.FormatSimpleDateTime(now)
				folder, err := suite.ac.Mail().CreateMailFolder(ctx, userID, folderName)
				require.NoError(t, err, clues.ToCore(err))

				return ptr.Val(folder.GetId())
			},
		},
		{
			name: "Test Mail: Hydrated Item Attachment Mail",
			bytes: mockconnector.GetMockMessageWithNestedItemAttachmentMail(t,
				mockconnector.GetMockMessageBytes("Basic Item Attachment"),
				"Mail Item Attachment",
			),
			category: path.EmailCategory,
			destination: func(t *testing.T, ctx context.Context) string {
				folderName := "TestRestoreMailBasicItemAttachment: " + common.FormatSimpleDateTime(now)
				folder, err := suite.ac.Mail().CreateMailFolder(ctx, userID, folderName)
				require.NoError(t, err, clues.ToCore(err))

				return ptr.Val(folder.GetId())
			},
		},
		{
			name: "Test Mail: Hydrated Item Attachment Mail One Attach",
			bytes: mockconnector.GetMockMessageWithNestedItemAttachmentMail(t,
				mockconnector.GetMockMessageWithDirectAttachment("Item Attachment Included"),
				"Mail Item Attachment",
			),
			category: path.EmailCategory,
			destination: func(t *testing.T, ctx context.Context) string {
				folderName := "ItemMailAttachmentwAttachment " + common.FormatSimpleDateTime(now)
				folder, err := suite.ac.Mail().CreateMailFolder(ctx, userID, folderName)
				require.NoError(t, err, clues.ToCore(err))

				return ptr.Val(folder.GetId())
			},
		},
		{
			name: "Test Mail: Item Attachment_Contact",
			bytes: mockconnector.GetMockMessageWithNestedItemAttachmentContact(t,
				mockconnector.GetMockContactBytes("Victor"),
				"Contact Item Attachment",
			),
			category: path.EmailCategory,
			destination: func(t *testing.T, ctx context.Context) string {
				folderName := "ItemMailAttachment_Contact " + common.FormatSimpleDateTime(now)
				folder, err := suite.ac.Mail().CreateMailFolder(ctx, userID, folderName)
				require.NoError(t, err, clues.ToCore(err))

				return ptr.Val(folder.GetId())
			},
		},
		{ // Restore will upload the Message without uploading the attachment
			name:     "Test Mail: Item Attachment_NestedEvent",
			bytes:    mockconnector.GetMockMessageWithNestedItemAttachmentEvent("Nested Item Attachment"),
			category: path.EmailCategory,
			destination: func(t *testing.T, ctx context.Context) string {
				folderName := "TestRestoreNestedEventItemAttachment: " + common.FormatSimpleDateTime(now)
				folder, err := suite.ac.Mail().CreateMailFolder(ctx, userID, folderName)
				require.NoError(t, err, clues.ToCore(err))

				return ptr.Val(folder.GetId())
			},
		},
		{
			name:     "Test Mail: One Large Attachment",
			bytes:    mockconnector.GetMockMessageWithLargeAttachment("Restore Large Attachment"),
			category: path.EmailCategory,
			destination: func(t *testing.T, ctx context.Context) string {
				folderName := "TestRestoreMailwithLargeAttachment: " + common.FormatSimpleDateTime(now)
				folder, err := suite.ac.Mail().CreateMailFolder(ctx, userID, folderName)
				require.NoError(t, err, clues.ToCore(err))

				return ptr.Val(folder.GetId())
			},
		},
		{
			name:     "Test Mail: Two Attachments",
			bytes:    mockconnector.GetMockMessageWithTwoAttachments("Restore 2 Attachments"),
			category: path.EmailCategory,
			destination: func(t *testing.T, ctx context.Context) string {
				folderName := "TestRestoreMailwithAttachments: " + common.FormatSimpleDateTime(now)
				folder, err := suite.ac.Mail().CreateMailFolder(ctx, userID, folderName)
				require.NoError(t, err, clues.ToCore(err))

				return ptr.Val(folder.GetId())
			},
		},
		{
			name:     "Test Mail: Reference(OneDrive) Attachment",
			bytes:    mockconnector.GetMessageWithOneDriveAttachment("Restore Reference(OneDrive) Attachment"),
			category: path.EmailCategory,
			destination: func(t *testing.T, ctx context.Context) string {
				folderName := "TestRestoreMailwithReferenceAttachment: " + common.FormatSimpleDateTime(now)
				folder, err := suite.ac.Mail().CreateMailFolder(ctx, userID, folderName)
				require.NoError(t, err, clues.ToCore(err))

				return ptr.Val(folder.GetId())
			},
		},
		// TODO: #884 - reinstate when able to specify root folder by name
		{
			name:     "Test Contact",
			bytes:    mockconnector.GetMockContactBytes("Test_Omega"),
			category: path.ContactsCategory,
			destination: func(t *testing.T, ctx context.Context) string {
				folderName := "TestRestoreContactObject: " + common.FormatSimpleDateTime(now)
				folder, err := suite.ac.Contacts().CreateContactFolder(ctx, userID, folderName)
				require.NoError(t, err, clues.ToCore(err))

				return ptr.Val(folder.GetId())
			},
		},
		{
			name:     "Test Events",
			bytes:    mockconnector.GetDefaultMockEventBytes("Restored Event Object"),
			category: path.EventsCategory,
			destination: func(t *testing.T, ctx context.Context) string {
				calendarName := "TestRestoreEventObject: " + common.FormatSimpleDateTime(now)
				calendar, err := suite.ac.Events().CreateCalendar(ctx, userID, calendarName)
				require.NoError(t, err, clues.ToCore(err))

				return ptr.Val(calendar.GetId())
			},
		},
		{
			name:     "Test Event with Attachment",
			bytes:    mockconnector.GetMockEventWithAttachment("Restored Event Attachment"),
			category: path.EventsCategory,
			destination: func(t *testing.T, ctx context.Context) string {
				calendarName := "TestRestoreEventObject_" + common.FormatSimpleDateTime(now)
				calendar, err := suite.ac.Events().CreateCalendar(ctx, userID, calendarName)
				require.NoError(t, err, clues.ToCore(err))

				return ptr.Val(calendar.GetId())
			},
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext()
			defer flush()

			destination := test.destination(t, ctx)
			info, err := RestoreExchangeObject(
				ctx,
				test.bytes,
				test.category,
				control.Copy,
				service,
				destination,
				userID,
				fault.New(true))
			assert.NoError(t, err, clues.ToCore(err))
			assert.NotNil(t, info, "item info was not populated")
			assert.NotNil(t, deleters)

			err = deleters[test.category].DeleteContainer(ctx, userID, destination)
			assert.NoError(t, err, clues.ToCore(err))
		})
	}
}
