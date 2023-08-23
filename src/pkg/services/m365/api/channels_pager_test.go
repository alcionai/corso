package api_test

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ChannelPagerIntgSuite struct {
	tester.Suite
	its intgTesterSetup
}

func TestChannelPagerIntgSuite(t *testing.T) {
	suite.Run(t, &ChannelPagerIntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs}),
	})
}

func (suite *ChannelPagerIntgSuite) SetupSuite() {
	suite.its = newIntegrationTesterSetup(suite.T())
}

func (suite *ChannelPagerIntgSuite) TestChannels_GetPage() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	teamID := tconfig.M365TeamsID(t)
	channelID := tconfig.M365ChannelID(t)
	pager := suite.its.ac.Channels().NewMessagePager(teamID, channelID, []string{})
	a, err := pager.GetPage(ctx)
	assert.NoError(t, err, clues.ToCore(err))
	assert.NotNil(t, a)
}

func (suite *ChannelPagerIntgSuite) TestChannels_Get() {
	t := suite.T()
	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		containerName = "General"
		teamID        = tconfig.M365TeamsID(t)
		chanClient    = suite.its.ac.Channels()
	)

	// GET channel -should be found
	channel, err := chanClient.GetChannelByName(ctx, teamID, containerName)
	assert.NoError(t, err, clues.ToCore(err))
	assert.Equal(t, ptr.Val(channel.GetDisplayName()), containerName)

	// GET channel -should not be found anymore
	_, err = chanClient.GetChannel(ctx, teamID, ptr.Val(channel.GetId()))
	assert.Error(t, err, clues.ToCore(err))
}

// func (suite *ChannelPagerIntgSuite) TestMessages_CreateGetAndDelete() {
// 	t := suite.T()
// 	ctx, flush := tester.NewContext(t)
// 	defer flush()

// 	var (
// 		teamID      = tconfig.M365TeamsID(t)
// 		channelID   = tconfig.M365ChannelID(t)
// 		credentials = suite.its.ac.Credentials
// 		chanClient  = suite.its.ac.Channels()
// 	)

// 	// GET channel - should be not found
// 	message, _, err := chanClient.GetMessage(ctx, teamID, channelID, "", "")
// 	assert.Error(t, err, clues.ToCore(err))

// 	// POST channel
// 	// patchBody := models.NewChatMessage()
// 	// body := models.NewItemBody()
// 	// content := "Hello World"
// 	// body.SetContent(&content)
// 	// patchBody.SetBody(body)

// 	// _,  := suite.its.ac.Channels().PostMessage(ctx, teamID, channelID, patchBody)
// 	// assert.NoError(t, err, clues.ToCore(err))

// }
