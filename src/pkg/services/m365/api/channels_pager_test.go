package api_test

import (
	"testing"
	"time"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/common/str"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
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

	msgIDs, du, err := ac.GetChannelMessageIDsDelta(
		ctx,
		suite.its.group.id,
		suite.its.group.testContainerID,
		"")
	require.NoError(t, err, clues.ToCore(err))
	require.NotEmpty(t, msgIDs)
	require.NotZero(t, du.URL, "delta link")
	require.True(t, du.Reset, "reset due to empty prev delta link")

	msgIDs, du, err = ac.GetChannelMessageIDsDelta(
		ctx,
		suite.its.group.id,
		suite.its.group.testContainerID,
		du.URL)
	require.NoError(t, err, clues.ToCore(err))
	require.Empty(t, msgIDs, "should have no new messages from delta")
	require.NotZero(t, du.URL, "delta link")
	require.False(t, du.Reset, "prev delta link should be valid")

	for id := range msgIDs {
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
	ac api.Channels,
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
	assert.Equal(t, str.Preview(ptr.Val(msg.GetBody().GetContent()), 16), info.MessagePreview)

	msgReplyIDs := map[string]struct{}{}

	for _, reply := range msg.GetReplies() {
		msgReplyIDs[ptr.Val(reply.GetId())] = struct{}{}
	}

	assert.Equal(t, replyIDs, msgReplyIDs)
}
