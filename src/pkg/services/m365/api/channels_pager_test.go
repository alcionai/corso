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

	cc := CallConfig{
		CanMakeDeltaQueries: true,
	}

	aar, err := ac.GetChannelMessageIDs(
		ctx,
		suite.its.group.id,
		suite.its.group.testContainerID,
		"",
		cc)
	require.NoError(t, err, clues.ToCore(err))
	require.NotEmpty(t, aar.Added)
	require.NotZero(t, aar.DU.URL, "delta link")
	require.True(t, aar.DU.Reset, "reset due to empty prev delta link")

	aar, err = ac.GetChannelMessageIDs(
		ctx,
		suite.its.group.id,
		suite.its.group.testContainerID,
		aar.DU.URL,
		cc)
	require.NoError(t, err, clues.ToCore(err))
	require.Empty(t, aar.Added, "should have no new messages from delta")
	require.Empty(t, aar.Removed, "should have no deleted messages from delta")
	require.NotZero(t, aar.DU.URL, "delta link")
	require.False(t, aar.DU.Reset, "prev delta link should be valid")

	for id := range aar.Added {
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
		lastReply   models.ChatMessageable
		lastReplyAt time.Time
		replyIDs    = map[string]struct{}{}
	)

	for _, r := range replies {
		cdt := ptr.Val(r.GetCreatedDateTime())
		if cdt.After(lastReplyAt) {
			lastReply = r
			lastReplyAt = cdt
		}

		replyIDs[ptr.Val(r.GetId())] = struct{}{}
	}

	assert.Equal(t, messageID, ptr.Val(msg.GetId()))
	assert.Equal(t, channelID, ptr.Val(msg.GetChannelIdentity().GetChannelId()))
	assert.Equal(t, groupID, ptr.Val(msg.GetChannelIdentity().GetTeamId()))
	// message
	assert.Equal(t, len(msg.GetAttachments()), len(info.Message.AttachmentNames))
	assert.Equal(t, len(replies), info.Message.ReplyCount)
	assert.Equal(t, lastReplyAt, info.Message.CreatedAt)
	assert.Equal(t, msg.GetFrom().GetUser().GetDisplayName(), info.Message.Creator)
	assert.Equal(t, str.Preview(ptr.Val(msg.GetBody().GetContent()), 128), info.Message.Preview)
	assert.Equal(t, len(ptr.Val(msg.GetBody().GetContent())), info.Message.Size)
	// last reply
	assert.Equal(t, len(lastReply.GetAttachments()), len(info.LastReply.AttachmentNames))
	assert.Zero(t, info.LastReply.ReplyCount)
	assert.Equal(t, lastReplyAt, info.LastReply.CreatedAt)
	assert.Equal(t, lastReply.GetFrom().GetUser().GetDisplayName(), info.LastReply.Creator)
	assert.Equal(t, str.Preview(ptr.Val(lastReply.GetBody().GetContent()), 128), info.LastReply.Preview)
	assert.Equal(t, len(ptr.Val(lastReply.GetBody().GetContent())), info.LastReply.Size)

	msgReplyIDs := map[string]struct{}{}

	for _, reply := range msg.GetReplies() {
		msgReplyIDs[ptr.Val(reply.GetId())] = struct{}{}
	}

	assert.Equal(t, replyIDs, msgReplyIDs)
}

func (suite *ChannelsPagerIntgSuite) TestIsSystemMessage() {
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
			test.expect(suite.T(), IsNotSystemMessage(test.cm))
		})
	}
}
