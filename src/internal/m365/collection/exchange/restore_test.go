package exchange

import (
	"context"
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/canario/src/internal/common/ptr"
	exchMock "github.com/alcionai/canario/src/internal/m365/service/exchange/mock"
	"github.com/alcionai/canario/src/internal/tester"
	"github.com/alcionai/canario/src/internal/tester/its"
	"github.com/alcionai/canario/src/internal/tester/tconfig"
	"github.com/alcionai/canario/src/pkg/control"
	"github.com/alcionai/canario/src/pkg/control/testdata"
	"github.com/alcionai/canario/src/pkg/count"
	"github.com/alcionai/canario/src/pkg/fault"
	"github.com/alcionai/canario/src/pkg/path"
	"github.com/alcionai/canario/src/pkg/services/m365/api"
)

type RestoreIntgSuite struct {
	tester.Suite
	m365 its.M365IntgTestSetup
}

func TestRestoreIntgSuite(t *testing.T) {
	suite.Run(t, &RestoreIntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs}),
	})
}

func (suite *RestoreIntgSuite) SetupSuite() {
	suite.m365 = its.GetM365(suite.T())
}

// TestRestoreContact ensures contact object can be created, placed into
// the Corso Folder. The function handles test clean-up.
func (suite *RestoreIntgSuite) TestRestoreContact() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		folderName = testdata.DefaultRestoreConfig("contact").Location
		handler    = newContactRestoreHandler(suite.m365.AC)
	)

	aFolder, err := handler.ac.CreateContainer(ctx, suite.m365.User.ID, "", folderName)
	require.NoError(t, err, clues.ToCore(err))

	folderID := ptr.Val(aFolder.GetId())

	defer func() {
		// Remove the folder containing contact prior to exiting test
		err = suite.m365.AC.Contacts().DeleteContainer(ctx, suite.m365.User.ID, folderID)
		assert.NoError(t, err, clues.ToCore(err))
	}()

	info, err := handler.restore(
		ctx,
		exchMock.ContactBytes("Corso TestContact"),
		suite.m365.User.ID,
		folderID,
		nil,
		control.Copy,
		fault.New(true),
		count.New())
	assert.NoError(t, err, clues.ToCore(err))
	assert.NotNil(t, info, "contact item info")
}

// TestRestoreEvent verifies that event object is able to created
// and sent into the test account of the Corso user in the newly created Corso Calendar
func (suite *RestoreIntgSuite) TestRestoreEvent() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		subject = testdata.DefaultRestoreConfig("event").Location
		handler = newEventRestoreHandler(suite.m365.AC)
	)

	calendar, err := handler.ac.CreateContainer(ctx, suite.m365.User.ID, "", subject)
	require.NoError(t, err, clues.ToCore(err))

	calendarID := ptr.Val(calendar.GetId())

	defer func() {
		// Removes calendar containing events created during the test
		err = suite.m365.AC.Events().DeleteContainer(ctx, suite.m365.User.ID, calendarID)
		assert.NoError(t, err, clues.ToCore(err))
	}()

	tests := []struct {
		name  string
		bytes []byte
	}{
		{
			name:  "Test Event With Attendees",
			bytes: exchMock.EventWithAttendeesBytes(subject),
		},
		{
			name:  "Test recurrenceTimeZone: Empty",
			bytes: exchMock.EventWithRecurrenceBytes(subject, `""`),
		},
		{
			name:  "Test cancelledOccurrences",
			bytes: exchMock.EventWithRecurrenceAndCancellationBytes(subject),
		},
		{
			name:  "Test exceptionOccurrences",
			bytes: exchMock.EventWithRecurrenceAndExceptionBytes(subject),
		},
		{
			name:  "Test exceptionOccurrences with different attachments",
			bytes: exchMock.EventWithRecurrenceAndExceptionAndAttachmentBytes(subject),
		},
	}

	for _, test := range tests {
		// Skip till https://github.com/alcionai/canario/issues/3675 is fixed
		if test.name == "Test exceptionOccurrences" {
			t.Skip("Bug 3675")
		}

		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			info, err := handler.restore(
				ctx,
				test.bytes,
				suite.m365.User.ID,
				calendarID,
				nil,
				control.Copy,
				fault.New(true),
				count.New())
			assert.NoError(t, err, clues.ToCore(err))
			assert.NotNil(t, info, "event item info")
		})
	}
}

// TestRestoreExchangeObject verifies path.Category usage for restored objects
func (suite *RestoreIntgSuite) TestRestoreExchangeObject() {
	t := suite.T()
	handlers := RestoreHandlers(suite.m365.AC)

	tests := []struct {
		name        string
		bytes       []byte
		category    path.CategoryType
		destination func(*testing.T, context.Context) string
	}{
		{
			name:     "Test Mail",
			bytes:    exchMock.MessageBytes("Restore Exchange Object"),
			category: path.EmailCategory,
			destination: func(t *testing.T, ctx context.Context) string {
				folderName := testdata.DefaultRestoreConfig("mailobj").Location
				folder, err := handlers[path.EmailCategory].
					CreateContainer(ctx, suite.m365.User.ID, "", folderName)
				require.NoError(t, err, clues.ToCore(err))

				return ptr.Val(folder.GetId())
			},
		},
		{
			name:     "Test Mail: One Direct Attachment",
			bytes:    exchMock.MessageWithDirectAttachment("Restore 1 Attachment"),
			category: path.EmailCategory,
			destination: func(t *testing.T, ctx context.Context) string {
				folderName := testdata.DefaultRestoreConfig("mailwattch").Location
				folder, err := handlers[path.EmailCategory].
					CreateContainer(ctx, suite.m365.User.ID, "", folderName)
				require.NoError(t, err, clues.ToCore(err))

				return ptr.Val(folder.GetId())
			},
		},
		{
			name:     "Test Mail: Item Attachment_Event",
			bytes:    exchMock.MessageWithItemAttachmentEvent("Event Item Attachment"),
			category: path.EmailCategory,
			destination: func(t *testing.T, ctx context.Context) string {
				folderName := testdata.DefaultRestoreConfig("eventwattch").Location
				folder, err := handlers[path.EmailCategory].
					CreateContainer(ctx, suite.m365.User.ID, "", folderName)
				require.NoError(t, err, clues.ToCore(err))

				return ptr.Val(folder.GetId())
			},
		},
		{
			name:     "Test Mail: Item Attachment_Mail",
			bytes:    exchMock.MessageWithItemAttachmentMail("Mail Item Attachment"),
			category: path.EmailCategory,
			destination: func(t *testing.T, ctx context.Context) string {
				folderName := testdata.DefaultRestoreConfig("mailitemattch").Location
				folder, err := handlers[path.EmailCategory].
					CreateContainer(ctx, suite.m365.User.ID, "", folderName)
				require.NoError(t, err, clues.ToCore(err))

				return ptr.Val(folder.GetId())
			},
		},
		{
			name: "Test Mail: Hydrated Item Attachment Mail",
			bytes: exchMock.MessageWithNestedItemAttachmentMail(t,
				exchMock.MessageBytes("Basic Item Attachment"),
				"Mail Item Attachment"),
			category: path.EmailCategory,
			destination: func(t *testing.T, ctx context.Context) string {
				folderName := testdata.DefaultRestoreConfig("mailbasicattch").Location
				folder, err := handlers[path.EmailCategory].
					CreateContainer(ctx, suite.m365.User.ID, "", folderName)
				require.NoError(t, err, clues.ToCore(err))

				return ptr.Val(folder.GetId())
			},
		},
		{
			name: "Test Mail: Hydrated Item Attachment Mail One Attach",
			bytes: exchMock.MessageWithNestedItemAttachmentMail(t,
				exchMock.MessageWithDirectAttachment("Item Attachment Included"),
				"Mail Item Attachment"),
			category: path.EmailCategory,
			destination: func(t *testing.T, ctx context.Context) string {
				folderName := testdata.DefaultRestoreConfig("mailnestattch").Location
				folder, err := handlers[path.EmailCategory].
					CreateContainer(ctx, suite.m365.User.ID, "", folderName)
				require.NoError(t, err, clues.ToCore(err))

				return ptr.Val(folder.GetId())
			},
		},
		{
			name: "Test Mail: Item Attachment_Contact",
			bytes: exchMock.MessageWithNestedItemAttachmentContact(t,
				exchMock.ContactBytes("Victor"),
				"Contact Item Attachment"),
			category: path.EmailCategory,
			destination: func(t *testing.T, ctx context.Context) string {
				folderName := testdata.DefaultRestoreConfig("mailcontactattch").Location
				folder, err := handlers[path.EmailCategory].
					CreateContainer(ctx, suite.m365.User.ID, "", folderName)
				require.NoError(t, err, clues.ToCore(err))

				return ptr.Val(folder.GetId())
			},
		},
		{ // Restore will upload the Message without uploading the attachment
			name:     "Test Mail: Item Attachment_NestedEvent",
			bytes:    exchMock.MessageWithNestedItemAttachmentEvent("Nested Item Attachment"),
			category: path.EmailCategory,
			destination: func(t *testing.T, ctx context.Context) string {
				folderName := testdata.DefaultRestoreConfig("nestedattch").Location
				folder, err := handlers[path.EmailCategory].
					CreateContainer(ctx, suite.m365.User.ID, "", folderName)
				require.NoError(t, err, clues.ToCore(err))

				return ptr.Val(folder.GetId())
			},
		},
		{
			name:     "Test Mail: One Large Attachment",
			bytes:    exchMock.MessageWithLargeAttachment("Restore Large Attachment"),
			category: path.EmailCategory,
			destination: func(t *testing.T, ctx context.Context) string {
				folderName := testdata.DefaultRestoreConfig("maillargeattch").Location
				folder, err := handlers[path.EmailCategory].
					CreateContainer(ctx, suite.m365.User.ID, "", folderName)
				require.NoError(t, err, clues.ToCore(err))

				return ptr.Val(folder.GetId())
			},
		},
		{
			name:     "Test Mail: Two Attachments",
			bytes:    exchMock.MessageWithTwoAttachments("Restore 2 Attachments"),
			category: path.EmailCategory,
			destination: func(t *testing.T, ctx context.Context) string {
				folderName := testdata.DefaultRestoreConfig("mailtwoattch").Location
				folder, err := handlers[path.EmailCategory].
					CreateContainer(ctx, suite.m365.User.ID, "", folderName)
				require.NoError(t, err, clues.ToCore(err))

				return ptr.Val(folder.GetId())
			},
		},
		{
			name:     "Test Mail: Reference(OneDrive) Attachment",
			bytes:    exchMock.MessageWithOneDriveAttachment("Restore Reference(OneDrive) Attachment"),
			category: path.EmailCategory,
			destination: func(t *testing.T, ctx context.Context) string {
				folderName := testdata.DefaultRestoreConfig("mailrefattch").Location
				folder, err := handlers[path.EmailCategory].
					CreateContainer(ctx, suite.m365.User.ID, "", folderName)
				require.NoError(t, err, clues.ToCore(err))

				return ptr.Val(folder.GetId())
			},
		},
		{
			name:     "Test Contact",
			bytes:    exchMock.ContactBytes("Test_Omega"),
			category: path.ContactsCategory,
			destination: func(t *testing.T, ctx context.Context) string {
				folderName := testdata.DefaultRestoreConfig("contact").Location
				folder, err := handlers[path.ContactsCategory].
					CreateContainer(ctx, suite.m365.User.ID, "", folderName)
				require.NoError(t, err, clues.ToCore(err))

				return ptr.Val(folder.GetId())
			},
		},
		{
			name:     "Test Events",
			bytes:    exchMock.EventBytes("Restored Event Object"),
			category: path.EventsCategory,
			destination: func(t *testing.T, ctx context.Context) string {
				folderName := testdata.DefaultRestoreConfig("event").Location
				calendar, err := handlers[path.EventsCategory].
					CreateContainer(ctx, suite.m365.User.ID, "", folderName)
				require.NoError(t, err, clues.ToCore(err))

				return ptr.Val(calendar.GetId())
			},
		},
		{
			name:     "Test Event with Attachment",
			bytes:    exchMock.EventWithAttachment("Restored Event Attachment"),
			category: path.EventsCategory,
			destination: func(t *testing.T, ctx context.Context) string {
				folderName := testdata.DefaultRestoreConfig("eventobj").Location
				calendar, err := handlers[path.EventsCategory].
					CreateContainer(ctx, suite.m365.User.ID, "", folderName)
				require.NoError(t, err, clues.ToCore(err))

				return ptr.Val(calendar.GetId())
			},
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			destination := test.destination(t, ctx)
			info, err := handlers[test.category].restore(
				ctx,
				test.bytes,
				suite.m365.User.ID,
				destination,
				nil,
				control.Copy,
				fault.New(true),
				count.New())
			assert.NoError(t, err, clues.ToCore(err))
			assert.NotNil(t, info, "item info was not populated")
		})
	}
}

func (suite *RestoreIntgSuite) TestRestoreAndBackupEvent_recurringInstancesWithAttachments() {
	t := suite.T()

	t.Skip("Bug 3675")

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		subject = testdata.DefaultRestoreConfig("event").Location
		handler = newEventRestoreHandler(suite.m365.AC)
	)

	calendar, err := handler.ac.CreateContainer(ctx, suite.m365.User.ID, "", subject)
	require.NoError(t, err, clues.ToCore(err))

	calendarID := ptr.Val(calendar.GetId())

	bytes := exchMock.EventWithRecurrenceAndExceptionAndAttachmentBytes("Reoccurring event restore and backup test")
	info, err := handler.restore(
		ctx,
		bytes,
		suite.m365.User.ID,
		calendarID,
		nil,
		control.Copy,
		fault.New(true),
		count.New())
	require.NoError(t, err, clues.ToCore(err))
	assert.NotNil(t, info, "event item info")

	ec, err := handler.ac.Stable.
		Client().
		Users().
		ByUserId(suite.m365.User.ID).
		Calendars().
		ByCalendarId(calendarID).
		Events().
		Get(ctx, nil)
	require.NoError(t, err, clues.ToCore(err))

	evts := ec.GetValue()
	assert.Len(t, evts, 1, "count of events")

	sp, info, err := suite.m365.AC.Events().GetItem(
		ctx,
		suite.m365.User.ID,
		ptr.Val(evts[0].GetId()),
		fault.New(true))
	require.NoError(t, err, clues.ToCore(err))
	assert.NotNil(t, info, "event item info")

	body, err := suite.m365.AC.Events().Serialize(
		ctx,
		sp,
		suite.m365.User.ID,
		ptr.Val(evts[0].GetId()))
	require.NoError(t, err, clues.ToCore(err))

	event, err := api.BytesToEventable(body)
	require.NoError(t, err, clues.ToCore(err))
	assert.NotNil(t, event.GetRecurrence(), "recurrence")

	eo := event.GetAdditionalData()["exceptionOccurrences"]
	assert.NotNil(t, eo, "exceptionOccurrences")

	assert.NotEqual(
		t,
		ptr.Val(event.GetSubject()),
		ptr.Val(eo.([]any)[0].(map[string]any)["subject"].(*string)),
		"name equal")

	atts := eo.([]any)[0].(map[string]any)["attachments"]
	assert.NotEqual(
		t,
		ptr.Val(event.GetAttachments()[0].GetName()),
		ptr.Val(atts.([]any)[0].(map[string]any)["name"].(*string)),
		"attachment name equal")
}
