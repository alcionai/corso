package api

import (
	"testing"
	"time"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/common/str"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
)

type ChannelsPagerIntgSuite struct {
	tester.Suite
	its intgTesterSetup
}

func TestChannelPagerIntgSuite(t *testing.T) {
	suite.Run(t, &ChannelsPagerIntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs}),
	})
}

func (suite *ChannelsPagerIntgSuite) SetupSuite() {
	suite.its = newIntegrationTesterSetup(suite.T())
}

func (suite *ChannelsPagerIntgSuite) TestEnumerateChannels() {
	var (
		t  = suite.T()
		ac = suite.its.ac.Channels()
	)

	ctx, flush := tester.NewContext(t)
	defer flush()

	chans, err := ac.GetChannels(ctx, suite.its.group.id)
	require.NoError(t, err, clues.ToCore(err))
	require.NotEmpty(t, chans)
}

func (suite *ChannelsPagerIntgSuite) TestEnumerateChannelMessages() {
	var (
		t  = suite.T()
		ac = suite.its.ac.Channels()
	)

	ctx, flush := tester.NewContext(t)
	defer flush()

	addedIDs, _, _, du, err := ac.GetChannelMessageIDs(
		ctx,
		suite.its.group.id,
		suite.its.group.testContainerID,
		"",
		true)
	require.NoError(t, err, clues.ToCore(err))
	require.NotEmpty(t, addedIDs)
	require.NotZero(t, du.URL, "delta link")
	require.True(t, du.Reset, "reset due to empty prev delta link")

	addedIDs, _, deletedIDs, du, err := ac.GetChannelMessageIDs(
		ctx,
		suite.its.group.id,
		suite.its.group.testContainerID,
		du.URL,
		true)
	require.NoError(t, err, clues.ToCore(err))
	require.Empty(t, addedIDs, "should have no new messages from delta")
	require.Empty(t, deletedIDs, "should have no deleted messages from delta")
	require.NotZero(t, du.URL, "delta link")
	require.False(t, du.Reset, "prev delta link should be valid")

	for id := range addedIDs {
		suite.Run(id+"-replies", func() {
			testEnumerateChannelMessageReplies(
				suite.T(),
				suite.its.ac.Channels(),
				suite.its.group.id,
				suite.its.group.testContainerID,
				id)
		})
	}
}

func testEnumerateChannelMessageReplies(
	t *testing.T,
	ac Channels,
	groupID, channelID, messageID string,
) {
	ctx, flush := tester.NewContext(t)
	defer flush()

	msg, info, err := ac.GetChannelMessage(ctx, groupID, channelID, messageID)
	require.NoError(t, err, clues.ToCore(err))

	replies, err := ac.GetChannelMessageReplies(ctx, groupID, channelID, messageID)
	require.NoError(t, err, clues.ToCore(err))

	var (
		lastReply time.Time
		replyIDs  = map[string]struct{}{}
	)

	for _, r := range replies {
		cdt := ptr.Val(r.GetCreatedDateTime())
		if cdt.After(lastReply) {
			lastReply = cdt
		}

		replyIDs[ptr.Val(r.GetId())] = struct{}{}
	}

	assert.Equal(t, messageID, ptr.Val(msg.GetId()))
	assert.Equal(t, channelID, ptr.Val(msg.GetChannelIdentity().GetChannelId()))
	assert.Equal(t, groupID, ptr.Val(msg.GetChannelIdentity().GetTeamId()))
	assert.Equal(t, len(replies), info.ReplyCount)
	assert.Equal(t, msg.GetFrom().GetUser().GetDisplayName(), info.MessageCreator)
	assert.Equal(t, lastReply, info.LastReplyAt)
	assert.Equal(t, str.Preview(ptr.Val(msg.GetBody().GetContent()), 128), info.MessagePreview)

	msgReplyIDs := map[string]struct{}{}

	for _, reply := range msg.GetReplies() {
		msgReplyIDs[ptr.Val(reply.GetId())] = struct{}{}
	}

	assert.Equal(t, replyIDs, msgReplyIDs)
}

func (suite *ChannelsPagerIntgSuite) TestFilterOutSystemMessages() {
	systemMessage := models.NewChatMessage()
	systemMessage.SetMessageType(ptr.To(models.SYSTEMEVENTMESSAGE_CHATMESSAGETYPE))

	systemMessageBody := models.NewItemBody()
	systemMessageBody.SetContent(ptr.To("<systemEventMessage/>"))

	messageBody := models.NewItemBody()
	messageBody.SetContent(ptr.To("message"))

	unknownFutureSystemMessage := models.NewChatMessage()
	unknownFutureSystemMessage.SetMessageType(ptr.To(models.UNKNOWNFUTUREVALUE_CHATMESSAGETYPE))
	unknownFutureSystemMessage.SetBody(systemMessageBody)

	unknownFutureMessage := models.NewChatMessage()
	unknownFutureMessage.SetMessageType(ptr.To(models.UNKNOWNFUTUREVALUE_CHATMESSAGETYPE))
	unknownFutureMessage.SetBody(messageBody)

	regularSystemMessage := models.NewChatMessage()
	regularSystemMessage.SetMessageType(ptr.To(models.MESSAGE_CHATMESSAGETYPE))
	regularSystemMessage.SetBody(systemMessageBody)

	regularMessage := models.NewChatMessage()
	regularMessage.SetMessageType(ptr.To(models.MESSAGE_CHATMESSAGETYPE))
	regularMessage.SetBody(messageBody)

	table := []struct {
		name   string
		cm     models.ChatMessageable
		expect assert.BoolAssertionFunc
	}{
		{
			name:   "system message type",
			cm:     systemMessage,
			expect: assert.False,
		},
		{
			name:   "unknown future type system body",
			cm:     unknownFutureSystemMessage,
			expect: assert.False,
		},
		{
			name:   "unknown future type message body",
			cm:     unknownFutureMessage,
			expect: assert.True,
		},
		{
			name:   "message type system body",
			cm:     regularSystemMessage,
			expect: assert.True,
		},
		{
			name:   "message type message body",
			cm:     regularMessage,
			expect: assert.True,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			test.expect(suite.T(), filterOutSystemMessages(test.cm))
		})
	}
}
