package exchange

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/ptr"
	exchMock "github.com/alcionai/corso/src/internal/m365/exchange/mock"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

type TransformUnitTest struct {
	tester.Suite
}

func TestSupportTestSuite(t *testing.T) {
	suite.Run(t, &TransformUnitTest{Suite: tester.NewUnitSuite(t)})
}

func (suite *TransformUnitTest) TestToMessage() {
	t := suite.T()

	bytes := exchMock.MessageBytes("m365 mail support test")
	message, err := api.BytesToMessageable(bytes)
	require.NoError(suite.T(), err, clues.ToCore(err))

	clone := toMessage(message)
	assert.Equal(t, message.GetBccRecipients(), clone.GetBccRecipients())
	assert.Equal(t, message.GetSubject(), clone.GetSubject())
	assert.Equal(t, message.GetSender(), clone.GetSender())
	assert.Equal(t, message.GetSentDateTime(), clone.GetSentDateTime())
	assert.NotEqual(t, message.GetId(), clone.GetId())
}

func (suite *TransformUnitTest) TestToEventSimplified_attendees() {
	t := suite.T()
	bytes := exchMock.EventWithAttendeesBytes("M365 Event Support Test")
	event, err := api.BytesToEventable(bytes)
	require.NoError(t, err, clues.ToCore(err))

	attendees := event.GetAttendees()
	newEvent := toEventSimplified(event)

	assert.Empty(t, newEvent.GetHideAttendees())
	assert.Equal(t, ptr.Val(event.GetBody().GetContentType()), ptr.Val(newEvent.GetBody().GetContentType()))
	assert.Equal(t, event.GetBody().GetAdditionalData(), newEvent.GetBody().GetAdditionalData())
	assert.Contains(t, ptr.Val(event.GetBody().GetContent()), "Required:")

	for _, member := range attendees {
		assert.Contains(t, ptr.Val(event.GetBody().GetContent()), ptr.Val(member.GetEmailAddress().GetName()))
		assert.Contains(t, ptr.Val(event.GetBody().GetContent()), ptr.Val(member.GetEmailAddress().GetAddress()))
	}
}

func (suite *TransformUnitTest) TestToEventSimplified_recurrence() {
	var (
		t       = suite.T()
		subject = "M365 Event Support Test"
	)

	tests := []struct {
		name           string
		event          func() models.Eventable
		validateOutput func(e models.Eventable) bool
	}{
		{
			name: "Test recurrence: Unspecified",
			event: func() models.Eventable {
				bytes := exchMock.EventWithSubjectBytes(subject)
				e, err := api.BytesToEventable(bytes)
				require.NoError(t, err, clues.ToCore(err))
				return e
			},

			validateOutput: func(e models.Eventable) bool {
				return e.GetRecurrence() == nil
			},
		},
		{
			name: "Test recurrenceTimeZone: Unspecified",
			event: func() models.Eventable {
				bytes := exchMock.EventWithRecurrenceBytes(subject, `null`)
				e, err := api.BytesToEventable(bytes)
				require.NoError(t, err, clues.ToCore(err))
				return e
			},

			validateOutput: func(e models.Eventable) bool {
				return e.GetRecurrence().GetRange().GetRecurrenceTimeZone() == nil
			},
		},
		{
			name: "Test recurrenceTimeZone: Empty",
			event: func() models.Eventable {
				bytes := exchMock.EventWithRecurrenceBytes(subject, `""`)
				event, err := api.BytesToEventable(bytes)
				require.NoError(t, err, clues.ToCore(err))
				return event
			},

			validateOutput: func(e models.Eventable) bool {
				return e.GetRecurrence().GetRange().GetRecurrenceTimeZone() == nil
			},
		},
		{
			name: "Test recurrenceTimeZone: Valid",
			event: func() models.Eventable {
				bytes := exchMock.EventWithRecurrenceBytes(subject, `"Pacific Standard Time"`)
				event, err := api.BytesToEventable(bytes)
				require.NoError(t, err, clues.ToCore(err))
				return event
			},

			validateOutput: func(e models.Eventable) bool {
				return ptr.Val(e.GetRecurrence().GetRange().GetRecurrenceTimeZone()) == "Pacific Standard Time"
			},
		},
		{
			name: "Test cancelledOccurrences",
			event: func() models.Eventable {
				bytes := exchMock.EventWithRecurrenceAndCancellationBytes(subject)
				event, err := api.BytesToEventable(bytes)
				require.NoError(t, err, clues.ToCore(err))
				return event
			},

			validateOutput: func(e models.Eventable) bool {
				return e.GetAdditionalData()["cancelledOccurrences"] == nil
			},
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			event := test.event()
			newEvent := toEventSimplified(event)
			assert.True(t, test.validateOutput(newEvent), test.name)
		})
	}
}

type mockContenter struct {
	content     *string
	contentType *models.BodyType
}

func (mc mockContenter) GetContent() *string {
	return mc.content
}

func (mc mockContenter) GetContentType() *models.BodyType {
	return mc.contentType
}

func makeMockContent(c string, ct models.BodyType) mockContenter {
	return mockContenter{&c, &ct}
}

func (suite *TransformUnitTest) TestInsertStringToBody() {
	nilTextContent := makeMockContent("", models.TEXT_BODYTYPE)
	nilTextContent.content = nil
	nilHTMLContent := makeMockContent("", models.HTML_BODYTYPE)
	nilHTMLContent.content = nil
	nilContentType := makeMockContent("brawnhilda", models.TEXT_BODYTYPE)
	nilContentType.contentType = nil

	table := []struct {
		name    string
		input   mockContenter
		content string
		expect  string
	}{
		{
			name:    "nil text content",
			input:   nilTextContent,
			content: "nil",
			expect:  "",
		},
		{
			name:    "nil html content",
			input:   nilHTMLContent,
			content: "nil",
			expect:  "",
		},
		{
			name:    "nil content type",
			input:   nilContentType,
			content: "nil",
			expect:  "",
		},
		{
			name:    "text",
			input:   makeMockContent("_text", models.TEXT_BODYTYPE),
			content: "new",
			expect:  "new_text",
		},
		{
			name:    "empty text",
			input:   makeMockContent("", models.TEXT_BODYTYPE),
			content: "new",
			expect:  "",
		},
		{
			name:    "expected html",
			input:   makeMockContent("_<body><div>_text</div></body>_", models.HTML_BODYTYPE),
			content: "foo",
			expect:  "_<body><div>foo_text</div></body>_",
		},
		{
			name:    "no div html",
			input:   makeMockContent("_<body>_text</body>_", models.HTML_BODYTYPE),
			content: "bar",
			expect:  "_<body>bar_text</body>_",
		},
		{
			name:    "no body html",
			input:   makeMockContent("_text", models.HTML_BODYTYPE),
			content: "baz",
			expect:  "baz_text",
		},
		{
			name:    "empty html",
			input:   makeMockContent("", models.HTML_BODYTYPE),
			content: "fnords",
			expect:  "",
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			result := insertStringToBody(test.input, test.content)
			assert.Equal(suite.T(), test.expect, result)
		})
	}
}
