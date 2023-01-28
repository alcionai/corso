package support

import (
	"testing"

	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/connector/mockconnector"
)

type SupportTestSuite struct {
	suite.Suite
}

func TestSupportTestSuite(t *testing.T) {
	suite.Run(t, new(SupportTestSuite))
}

func (suite *SupportTestSuite) TestToMessage() {
	bytes := mockconnector.GetMockMessageBytes("m365 mail support test")
	message, err := CreateMessageFromBytes(bytes)
	require.NoError(suite.T(), err)

	clone := ToMessage(message)
	suite.Equal(message.GetBccRecipients(), clone.GetBccRecipients())
	suite.Equal(message.GetSubject(), clone.GetSubject())
	suite.Equal(message.GetSender(), clone.GetSender())
	suite.Equal(message.GetSentDateTime(), clone.GetSentDateTime())
	suite.NotEqual(message.GetId(), clone.GetId())
}

func (suite *SupportTestSuite) TestToEventSimplified() {
	t := suite.T()
	bytes := mockconnector.GetMockEventWithAttendeesBytes("M365 Event Support Test")
	event, err := CreateEventFromBytes(bytes)
	require.NoError(t, err)

	attendees := event.GetAttendees()
	newEvent := ToEventSimplified(event)

	assert.Empty(t, newEvent.GetHideAttendees())
	assert.Equal(t, *event.GetBody().GetContentType(), *newEvent.GetBody().GetContentType())
	assert.Equal(t, event.GetBody().GetAdditionalData(), newEvent.GetBody().GetAdditionalData())
	assert.Contains(t, *event.GetBody().GetContent(), "Required:")

	for _, member := range attendees {
		assert.Contains(t, *event.GetBody().GetContent(), *member.GetEmailAddress().GetName())
		assert.Contains(t, *event.GetBody().GetContent(), *member.GetEmailAddress().GetAddress())
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

func (suite *SupportTestSuite) TestInsertStringToBody() {
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
		suite.T().Run(test.name, func(t *testing.T) {
			result := insertStringToBody(test.input, test.content)
			assert.Equal(t, test.expect, result)
		})
	}
}
