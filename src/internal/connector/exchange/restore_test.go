package exchange

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/internal/connector/exchange/api"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/mockconnector"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/path"
)

type ExchangeRestoreSuite struct {
	suite.Suite
	gs          graph.Servicer
	credentials account.M365Config
	ac          api.Client
}

func TestExchangeRestoreSuite(t *testing.T) {
	tester.RunOnAny(
		t,
		tester.CorsoCITests,
		tester.CorsoConnectorRestoreExchangeCollectionTests)

	suite.Run(t, new(ExchangeRestoreSuite))
}

func (suite *ExchangeRestoreSuite) SetupSuite() {
	t := suite.T()
	tester.MustGetEnvSets(t, tester.AWSStorageCredEnvs, tester.M365AcctCredEnvs)

	a := tester.NewM365Account(t)
	m365, err := a.M365Config()
	require.NoError(t, err)

	suite.credentials = m365
	suite.ac, err = api.NewClient(m365)
	require.NoError(t, err)

	require.NoError(suite.T(), err)
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

	aFolder, err := suite.ac.CreateContactFolder(ctx, userID, folderName)
	require.NoError(t, err)

	folderID := *aFolder.GetId()

	defer func() {
		// Remove the folder containing contact prior to exiting test
		err = suite.ac.DeleteContactFolder(ctx, userID, folderID)
		assert.NoError(t, err)
	}()

	info, err := RestoreExchangeContact(ctx,
		mockconnector.GetMockContactBytes("Corso TestContact"),
		suite.gs,
		control.Copy,
		folderID,
		userID)
	assert.NoError(t, err, support.ConnectorStackErrorTrace(err))
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

	calendar, err := suite.ac.CreateCalendar(ctx, userID, name)
	require.NoError(t, err)

	calendarID := *calendar.GetId()

	defer func() {
		// Removes calendar containing events created during the test
		err = suite.ac.DeleteCalendar(ctx, userID, calendarID)
		assert.NoError(t, err)
	}()

	info, err := RestoreExchangeEvent(ctx,
		mockconnector.GetMockEventWithAttendeesBytes(name),
		suite.gs,
		control.Copy,
		calendarID,
		userID)
	assert.NoError(t, err, support.ConnectorStackErrorTrace(err))
	assert.NotNil(t, info, "event item info")
}

// TestRestoreExchangeObject verifies path.Category usage for restored objects
func (suite *ExchangeRestoreSuite) TestRestoreExchangeObject() {
	a := tester.NewM365Account(suite.T())
	m365, err := a.M365Config()
	require.NoError(suite.T(), err)

	service, err := createService(m365)
	require.NoError(suite.T(), err)

	userID := tester.M365UserID(suite.T())
	now := time.Now()
	tests := []struct {
		name        string
		bytes       []byte
		category    path.CategoryType
		cleanupFunc func(context.Context, string, string) error
		destination func(*testing.T, context.Context) string
	}{
		{
			name:        "Test Mail",
			bytes:       mockconnector.GetMockMessageBytes("Restore Exchange Object"),
			category:    path.EmailCategory,
			cleanupFunc: suite.ac.DeleteMailFolder,
			destination: func(t *testing.T, ctx context.Context) string {
				folderName := "TestRestoreMailObject: " + common.FormatSimpleDateTime(now)
				folder, err := suite.ac.CreateMailFolder(ctx, userID, folderName)
				require.NoError(t, err)

				return *folder.GetId()
			},
		},
		{
			name:        "Test Mail: One Direct Attachment",
			bytes:       mockconnector.GetMockMessageWithDirectAttachment("Restore 1 Attachment"),
			category:    path.EmailCategory,
			cleanupFunc: suite.ac.DeleteMailFolder,
			destination: func(t *testing.T, ctx context.Context) string {
				folderName := "TestRestoreMailwithAttachment: " + common.FormatSimpleDateTime(now)
				folder, err := suite.ac.CreateMailFolder(ctx, userID, folderName)
				require.NoError(t, err)

				return *folder.GetId()
			},
		},
		{
			name:        "Test Mail: One Large Attachment",
			bytes:       mockconnector.GetMockMessageWithLargeAttachment("Restore Large Attachment"),
			category:    path.EmailCategory,
			cleanupFunc: suite.ac.DeleteMailFolder,
			destination: func(t *testing.T, ctx context.Context) string {
				folderName := "TestRestoreMailwithLargeAttachment: " + common.FormatSimpleDateTime(now)
				folder, err := suite.ac.CreateMailFolder(ctx, userID, folderName)
				require.NoError(t, err)

				return *folder.GetId()
			},
		},
		{
			name:        "Test Mail: Two Attachments",
			bytes:       mockconnector.GetMockMessageWithTwoAttachments("Restore 2 Attachments"),
			category:    path.EmailCategory,
			cleanupFunc: suite.ac.DeleteMailFolder,
			destination: func(t *testing.T, ctx context.Context) string {
				folderName := "TestRestoreMailwithAttachments: " + common.FormatSimpleDateTime(now)
				folder, err := suite.ac.CreateMailFolder(ctx, userID, folderName)
				require.NoError(t, err)

				return *folder.GetId()
			},
		},
		{
			name:        "Test Mail: Reference(OneDrive) Attachment",
			bytes:       mockconnector.GetMessageWithOneDriveAttachment("Restore Reference(OneDrive) Attachment"),
			category:    path.EmailCategory,
			cleanupFunc: suite.ac.DeleteMailFolder,
			destination: func(t *testing.T, ctx context.Context) string {
				folderName := "TestRestoreMailwithReferenceAttachment: " + common.FormatSimpleDateTime(now)
				folder, err := suite.ac.CreateMailFolder(ctx, userID, folderName)
				require.NoError(t, err)

				return *folder.GetId()
			},
		},
		// TODO: #884 - reinstate when able to specify root folder by name
		{
			name:        "Test Contact",
			bytes:       mockconnector.GetMockContactBytes("Test_Omega"),
			category:    path.ContactsCategory,
			cleanupFunc: suite.ac.DeleteContactFolder,
			destination: func(t *testing.T, ctx context.Context) string {
				folderName := "TestRestoreContactObject: " + common.FormatSimpleDateTime(now)
				folder, err := suite.ac.CreateContactFolder(ctx, userID, folderName)
				require.NoError(t, err)

				return *folder.GetId()
			},
		},
		{
			name:        "Test Events",
			bytes:       mockconnector.GetDefaultMockEventBytes("Restored Event Object"),
			category:    path.EventsCategory,
			cleanupFunc: suite.ac.DeleteCalendar,
			destination: func(t *testing.T, ctx context.Context) string {
				calendarName := "TestRestoreEventObject: " + common.FormatSimpleDateTime(now)
				calendar, err := suite.ac.CreateCalendar(ctx, userID, calendarName)
				require.NoError(t, err)

				return *calendar.GetId()
			},
		},
		{
			name:        "Test Event with Attachment",
			bytes:       mockconnector.GetMockEventWithAttachment("Restored Event Attachment"),
			category:    path.EventsCategory,
			cleanupFunc: suite.ac.DeleteCalendar,
			destination: func(t *testing.T, ctx context.Context) string {
				calendarName := "TestRestoreEventObject_" + common.FormatSimpleDateTime(now)
				calendar, err := suite.ac.CreateCalendar(ctx, userID, calendarName)
				require.NoError(t, err)

				return *calendar.GetId()
			},
		},
	}

	for _, test := range tests {
		suite.T().Run(test.name, func(t *testing.T) {
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
			)
			assert.NoError(t, err, support.ConnectorStackErrorTrace(err))
			assert.NotNil(t, info, "item info is populated")

			cleanupError := test.cleanupFunc(ctx, userID, destination)
			assert.NoError(t, cleanupError)
		})
	}
}
