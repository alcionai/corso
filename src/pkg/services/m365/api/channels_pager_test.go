package api_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/alcionai/clues"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
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

	teamID := tconfig.M365TeamID(t)
	channelID := tconfig.M365ChannelID(t)
	pager := suite.its.ac.Channels().NewMessagePager(teamID, channelID, []string{})
	a, err := pager.GetPage(ctx)
	assert.NoError(t, err, clues.ToCore(err))
	assert.NotNil(t, a)
}

func (suite *ChannelPagerIntgSuite) TestChannels_CreateGetAndDelete() {
	t := suite.T()
	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		yy, mm, dd    = time.Now().Date()
		hh            = time.Now().Hour()
		min           = time.Now().Minute()
		ss            = time.Now().Second()
		containerName = fmt.Sprintf("testChannel%d%d%d%d%d%d", yy, mm, dd, hh, min, ss)
		teamID        = tconfig.M365TeamID(t)
		credentials   = suite.its.ac.Credentials
		chanClient    = suite.its.ac.Channels()
	)

	// GET channel - should be not found
	_, err := suite.its.ac.Channels().GetChannelByName(ctx, teamID, containerName)
	assert.Error(t, err, clues.ToCore(err))

	// POST channel
	channelPost, err := suite.its.ac.Channels().CreateChannel(ctx, teamID, containerName)
	assert.NoError(t, err, clues.ToCore(err))

	postChannelID := ptr.Val(channelPost.GetId())

	// DELETE channel
	defer func() {
		_, err := chanClient.GetChannelByID(ctx, teamID, postChannelID)

		if err != nil {
			fmt.Println("could not find channel: ", err)
		} else {
			deleteChannel(ctx, credentials, teamID, postChannelID)
		}
	}()

	// GET channel -should be found
	channel, err := chanClient.GetChannelByName(ctx, teamID, containerName)
	assert.NoError(t, err, clues.ToCore(err))
	assert.Equal(t, ptr.Val(channel.GetDisplayName()), containerName)

	// PATCH channel
	patchBody := models.NewChannel()
	patchName := fmt.Sprintf("othername%d%d%d%d%d%d", yy, mm, dd, hh, min, ss)
	patchBody.SetDisplayName(ptr.To(patchName))
	err = chanClient.PatchChannel(ctx, teamID, postChannelID, patchBody)
	assert.NoError(t, err, clues.ToCore(err))
	assert.Equal(t, ptr.Val(channel.GetDisplayName()), containerName)

	// GET channel -should not be found with old name
	_, err = chanClient.GetChannelByName(ctx, teamID, containerName)
	assert.Error(t, err, clues.ToCore(err))

	// GET channel -should be found with new name
	channel, err = chanClient.GetChannelByName(ctx, teamID, patchName)
	assert.NoError(t, err, clues.ToCore(err))
	assert.Equal(t, ptr.Val(channel.GetDisplayName()), patchName)
	assert.Equal(t, ptr.Val(channel.GetId()), postChannelID)

	// GET channel -should not be found with old name
	err = chanClient.DeleteChannel(ctx, teamID, postChannelID)
	assert.NoError(t, err, clues.ToCore(err))

	// GET channel -should not be found anymore
	_, err = chanClient.GetChannel(ctx, teamID, postChannelID)
	assert.Error(t, err, clues.ToCore(err))
}

func deleteChannel(ctx context.Context, credentials account.M365Config, teamID, postChannelID string) {
	srv, err := api.NewService(credentials)
	if err != nil {
		fmt.Println("Error found in getting creds")
	}

	if err != nil {
		fmt.Println("Error found in getting creds")
	}

	err = srv.Client().
		Teams().
		ByTeamId(teamID).
		Channels().
		ByChannelId(postChannelID).
		Delete(ctx, nil)
	if err != nil {
		fmt.Println("channel could not be delete in defer")
	}
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
