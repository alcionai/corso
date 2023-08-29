package api_test

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

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
		_, _, err := ac.GetChannelMessage(
			ctx,
			suite.its.group.id,
			suite.its.group.testContainerID,
			id)
		require.NoError(t, err, clues.ToCore(err))
	}
}
