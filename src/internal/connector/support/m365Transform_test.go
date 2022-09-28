package support

import (
	"testing"

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

	newEvent := ToEventSimplified(event)

	assert.Empty(t, newEvent.GetHideAttendees())
	assert.Equal(t, *event.GetBody().GetContentType(), *newEvent.GetBody().GetContentType())
	assert.Equal(t, event.GetBody().GetAdditionalData(), newEvent.GetBody().GetAdditionalData())
	assert.Contains(t, *event.GetBody().GetContent(), "Required:")
}
