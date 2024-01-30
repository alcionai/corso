package eml

import (
	"bytes"
	"regexp"
	"strings"
	"testing"
	"time"

	ical "github.com/arran4/golang-ical"
	"github.com/jhillyerd/enmime"
	kjson "github.com/microsoft/kiota-serialization-json-go"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/converters/eml/testdata"
	"github.com/alcionai/corso/src/internal/converters/ics"
	"github.com/alcionai/corso/src/internal/m365/collection/groups/metadata"
	stub "github.com/alcionai/corso/src/internal/m365/service/groups/mock"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

type EMLUnitSuite struct {
	tester.Suite
}

func TestEMLUnitSuite(t *testing.T) {
	suite.Run(t, &EMLUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *EMLUnitSuite) TestFormatAddress() {
	t := suite.T()

	tests := []struct {
		tname    string
		name     string
		email    string
		expected string
	}{
		{
			tname:    "different name and email",
			name:     "John Doe",
			email:    "johndoe@provider.com",
			expected: `"John Doe" <johndoe@provider.com>`,
		},
		{
			tname:    "same name and email",
			name:     "johndoe@provider.com",
			email:    "johndoe@provider.com",
			expected: "johndoe@provider.com",
		},
		{
			tname:    "only email",
			name:     "",
			email:    "johndoe@provider.com",
			expected: "johndoe@provider.com",
		},
		{
			tname:    "only name",
			name:     "john doe",
			email:    "",
			expected: `"john doe"`,
		},
		{
			tname:    "neither mail or name",
			name:     "",
			email:    "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.tname, func(t *testing.T) {
			entity := models.NewEmailAddress()
			if len(tt.name) != 0 {
				entity.SetName(ptr.To(tt.name))
			}
			if len(tt.email) != 0 {
				entity.SetAddress(ptr.To(tt.email))
			}

			assert.Equal(t, tt.expected, formatAddress(entity))
		})
	}
}

func (suite *EMLUnitSuite) TestConvert_messageble_to_eml() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	body := []byte(testdata.EmailWithAttachments)

	out, err := FromJSON(ctx, body)
	assert.NoError(t, err, "converting to eml")

	msg, err := api.BytesToMessageable(body)
	require.NoError(t, err, "creating message")

	eml, err := enmime.ReadEnvelope(strings.NewReader(out))
	require.NoError(t, err, "reading created eml")

	assert.Equal(t, ptr.Val(msg.GetSubject()), eml.GetHeader("Subject"))
	assert.Equal(t, msg.GetSentDateTime().Format(time.RFC1123Z), eml.GetHeader("Date"))

	assert.Equal(t, formatAddress(msg.GetFrom().GetEmailAddress()), eml.GetHeader("From"))

	ccs := strings.Split(eml.GetHeader("Cc"), ", ")
	for _, cc := range msg.GetCcRecipients() {
		assert.Contains(t, ccs, formatAddress(cc.GetEmailAddress()))
	}

	bccs := strings.Split(eml.GetHeader("Bcc"), ", ")
	for _, bcc := range msg.GetBccRecipients() {
		assert.Contains(t, bccs, formatAddress(bcc.GetEmailAddress()))
	}

	tos := strings.Split(eml.GetHeader("To"), ", ")
	for _, to := range msg.GetToRecipients() {
		assert.Contains(t, tos, formatAddress(to.GetEmailAddress()))
	}

	source := strings.ReplaceAll(eml.HTML, "\n", "")
	target := strings.ReplaceAll(ptr.Val(msg.GetBody().GetContent()), "\n", "")

	// replace the cid with a constant value to make the comparison
	re := regexp.MustCompile(`src="cid:[^"]*"`)
	source = re.ReplaceAllString(source, `src="cid:replaced"`)
	target = re.ReplaceAllString(target, `src="cid:replaced"`)

	assert.Equal(t, source, target)
}

func (suite *EMLUnitSuite) TestConvert_edge_cases() {
	tests := []struct {
		name      string
		transform func(models.Messageable)
	}{
		{
			name: "just a name",
			transform: func(msg models.Messageable) {
				msg.GetFrom().GetEmailAddress().SetName(ptr.To("alphabob"))
				msg.GetFrom().GetEmailAddress().SetAddress(nil)
			},
		},
		{
			name: "incorrect address",
			transform: func(msg models.Messageable) {
				msg.GetFrom().GetEmailAddress().SetAddress(ptr.To("invalid"))
			},
		},
		{
			name: "empty attachment",
			transform: func(msg models.Messageable) {
				attachments := msg.GetAttachments()
				err := attachments[0].GetBackingStore().Set("contentBytes", []uint8{})
				require.NoError(suite.T(), err, "setting attachment content")
			},
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			body := []byte(testdata.EmailWithAttachments)

			msg, err := api.BytesToMessageable(body)
			require.NoError(t, err, "creating message")

			test.transform(msg)

			writer := kjson.NewJsonSerializationWriter()

			defer writer.Close()

			err = writer.WriteObjectValue("", msg)
			require.NoError(t, err, "serializing message")

			nbody, err := writer.GetSerializedContent()
			require.NoError(t, err, "getting serialized content")

			_, err = FromJSON(ctx, nbody)
			assert.NoError(t, err, "converting to eml")
		})
	}
}

func (suite *EMLUnitSuite) TestConvert_eml_ics() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	body := []byte(testdata.EmailWithEventInfo)

	out, err := FromJSON(ctx, body)
	assert.NoError(t, err, "converting to eml")

	rmsg, err := api.BytesToMessageable(body)
	require.NoError(t, err, "creating message")

	msg := rmsg.(*models.EventMessageRequest)

	eml, err := enmime.ReadEnvelope(strings.NewReader(out))
	require.NoError(t, err, "reading created eml")
	require.NotNil(t, eml, "eml should not be nil")

	require.Equal(t, 1, len(eml.OtherParts), "eml should have 1 attachment")
	require.Equal(t, "text/calendar", eml.OtherParts[0].ContentType, "eml attachment should be a calendar")

	catt := *eml.OtherParts[0]
	cal, err := ical.ParseCalendar(bytes.NewReader(catt.Content))
	require.NoError(t, err, "parsing calendar")

	event := cal.Events()[0]

	assert.Equal(t, ptr.Val(msg.GetId()), event.Id())
	assert.Equal(t, ptr.Val(msg.GetSubject()), event.GetProperty(ical.ComponentPropertySummary).Value)

	assert.Equal(
		t,
		msg.GetCreatedDateTime().Format(ics.ICalDateTimeFormat),
		event.GetProperty(ical.ComponentPropertyCreated).Value)
	assert.Equal(
		t,
		msg.GetLastModifiedDateTime().Format(ics.ICalDateTimeFormat),
		event.GetProperty(ical.ComponentPropertyLastModified).Value)

	st, err := ics.GetUTCTime(
		ptr.Val(msg.GetStartDateTime().GetDateTime()),
		ptr.Val(msg.GetStartDateTime().GetTimeZone()))
	require.NoError(t, err, "getting start time")

	et, err := ics.GetUTCTime(
		ptr.Val(msg.GetEndDateTime().GetDateTime()),
		ptr.Val(msg.GetEndDateTime().GetTimeZone()))
	require.NoError(t, err, "getting end time")

	assert.Equal(
		t,
		st.Format(ics.ICalDateTimeFormat),
		event.GetProperty(ical.ComponentPropertyDtStart).Value)
	assert.Equal(
		t,
		et.Format(ics.ICalDateTimeFormat),
		event.GetProperty(ical.ComponentPropertyDtEnd).Value)

	tos := msg.GetToRecipients()
	ccs := msg.GetCcRecipients()
	att := event.Attendees()

	assert.Equal(t, len(tos)+len(ccs), len(att))

	for _, to := range tos {
		found := false

		for _, attendee := range att {
			if "mailto:"+ptr.Val(to.GetEmailAddress().GetAddress()) == attendee.Value {
				found = true

				assert.Equal(t, "REQ-PARTICIPANT", attendee.ICalParameters["ROLE"][0])

				break
			}
		}

		assert.True(t, found, "to recipient not found in attendees")
	}

	for _, cc := range ccs {
		found := false

		for _, attendee := range att {
			if "mailto:"+ptr.Val(cc.GetEmailAddress().GetAddress()) == attendee.Value {
				found = true

				assert.Equal(t, "OPT-PARTICIPANT", attendee.ICalParameters["ROLE"][0])

				break
			}
		}

		assert.True(t, found, "cc recipient not found in attendees")
	}
}

func (suite *EMLUnitSuite) TestConvert_eml_ics_from_event_obj() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	body := []byte(testdata.EmailWithEventObject)

	out, err := FromJSON(ctx, body)
	assert.NoError(t, err, "converting to eml")

	rmsg, err := api.BytesToMessageable(body)
	require.NoError(t, err, "creating message")

	msg := rmsg.(*models.EventMessageRequest)
	evt := msg.GetEvent()

	eml, err := enmime.ReadEnvelope(strings.NewReader(out))
	require.NoError(t, err, "reading created eml")
	require.NotNil(t, eml, "eml should not be nil")

	require.Equal(t, 1, len(eml.OtherParts), "eml should have 1 attachment")
	require.Equal(t, "text/calendar", eml.OtherParts[0].ContentType, "eml attachment should be a calendar")

	catt := *eml.OtherParts[0]
	cal, err := ical.ParseCalendar(bytes.NewReader(catt.Content))
	require.NoError(t, err, "parsing calendar")

	event := cal.Events()[0]

	assert.Equal(t, ptr.Val(evt.GetId()), event.Id())
	assert.NotEqual(t, ptr.Val(msg.GetSubject()), event.GetProperty(ical.ComponentPropertySummary).Value)
	assert.Equal(t, ptr.Val(evt.GetSubject()), event.GetProperty(ical.ComponentPropertySummary).Value)
}

//-------------------------------------------------------------
// Postable -> EML tests
//-------------------------------------------------------------

func (suite *EMLUnitSuite) TestConvert_postable_to_eml() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	body := []byte(stub.PostWithAttachments)

	postMetadata := metadata.ConversationPostMetadata{
		Recipients: []string{"group@example.com"},
		Topic:      "test subject",
	}

	out, err := FromJSONPostToEML(ctx, body, postMetadata)
	assert.NoError(t, err, "converting to eml")

	post, err := api.BytesToPostable(body)
	require.NoError(t, err, "creating post")

	eml, err := enmime.ReadEnvelope(strings.NewReader(out))
	require.NoError(t, err, "reading created eml")

	assert.Equal(t, postMetadata.Topic, eml.GetHeader("Subject"))
	assert.Equal(t, post.GetCreatedDateTime().Format(time.RFC1123Z), eml.GetHeader("Date"))

	assert.Equal(t, formatAddress(post.GetFrom().GetEmailAddress()), eml.GetHeader("From"))

	// Test recipients. The post metadata should contain the group email address.

	tos := strings.Split(eml.GetHeader("To"), ", ")
	for _, sourceTo := range postMetadata.Recipients {
		assert.Contains(t, tos, sourceTo)
	}

	// Assert cc, bcc to be empty since they are not supported for posts right now.
	assert.Equal(t, "", eml.GetHeader("Cc"))
	assert.Equal(t, "", eml.GetHeader("Bcc"))

	// Test attachments using PostWithAttachments data as a reference.
	// This data has 1 direct attachment and 1 inline attachment.
	assert.Equal(t, 1, len(eml.Attachments), "direct attachment count")
	assert.Equal(t, 1, len(eml.Inlines), "inline attachment count")

	for _, sourceAttachment := range post.GetAttachments() {
		targetContent := eml.Attachments[0].Content
		if ptr.Val(sourceAttachment.GetIsInline()) {
			targetContent = eml.Inlines[0].Content
		}

		sourceContent, err := sourceAttachment.GetBackingStore().Get("contentBytes")
		assert.NoError(t, err, "getting source attachment content")

		assert.Equal(t, sourceContent, targetContent)
	}

	// Test body
	source := strings.ReplaceAll(eml.HTML, "\n", "")
	target := strings.ReplaceAll(ptr.Val(post.GetBody().GetContent()), "\n", "")

	// replace the cid with a constant value to make the comparison
	re := regexp.MustCompile(`(?:src|originalSrc)="cid:[^"]*"`)
	source = re.ReplaceAllString(source, `src="cid:replaced"`)
	target = re.ReplaceAllString(target, `src="cid:replaced"`)

	assert.Equal(t, source, target)
}
