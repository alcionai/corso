package api_test

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
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

// This will be added once 'pager' is implemented
// func (suite *ChannelPagerIntgSuite) TestChannels_GetPage() {
// 	t := suite.T()

// 	ctx, flush := tester.NewContext(t)
// 	defer flush()

// 	teamID := tconfig.M365TeamID(t)
// 	channelID := tconfig.M365ChannelID(t)
// 	pager := suite.its.ac.Channels().NewMessagePager(teamID, channelID, []string{})
// 	a, err := pager.GetPage(ctx)
// 	assert.NoError(t, err, clues.ToCore(err))
// 	assert.NotNil(t, a)
// }

func (suite *ChannelPagerIntgSuite) TestChannels_Get() {
	t := suite.T()
	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		containerName = "General"
		teamID        = tconfig.M365TeamID(t)
		chanClient    = suite.its.ac.Channels()
	)

	// GET channel -should be found
	channel, err := chanClient.GetChannelByName(ctx, teamID, containerName)
	assert.NoError(t, err, clues.ToCore(err))
	assert.Equal(t, ptr.Val(channel.GetDisplayName()), containerName)

	// GET channel -should be found
	_, err = chanClient.GetChannel(ctx, teamID, ptr.Val(channel.GetId()))
	assert.NoError(t, err, clues.ToCore(err))
}
