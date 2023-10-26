package tsetup

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/require"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/internal/m365/graph/mock"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

// ---------------------------------------------------------------------------
// Gockable client
// ---------------------------------------------------------------------------

// GockClient produces a new exchange api client that can be
// mocked using gock.
func gockClient(creds account.M365Config) (api.Client, error) {
	s, err := mock.NewService(creds)
	if err != nil {
		return api.Client{}, err
	}

	li, err := mock.NewService(creds, graph.NoTimeout())
	if err != nil {
		return api.Client{}, err
	}

	return api.Client{
		Credentials: creds,
		Stable:      s,
		LargeItem:   li,
	}, nil
}

// ---------------------------------------------------------------------------
// Suite Setup
// ---------------------------------------------------------------------------

type m365IDs struct {
	ID                string
	Email             string
	DriveID           string
	DriveRootFolderID string
	TestContainerID   string
}

type M365 struct {
	AC           api.Client
	GockAC       api.Client
	User         m365IDs
	Site         m365IDs
	Group        m365IDs
	NonTeamGroup m365IDs // group which does not have an associated team
}

func NewM365IntegrationTester(t *testing.T) M365 {
	mit := M365{}

	ctx, flush := tester.NewContext(t)
	defer flush()

	graph.InitializeConcurrencyLimiter(ctx, true, 4)

	a := tconfig.NewM365Account(t)
	creds, err := a.M365Config()
	require.NoError(t, err, clues.ToCore(err))

	mit.AC, err = api.NewClient(creds, control.DefaultOptions())
	require.NoError(t, err, clues.ToCore(err))

	mit.GockAC, err = gockClient(creds)
	require.NoError(t, err, clues.ToCore(err))

	// user drive

	mit.User.ID = tconfig.M365UserID(t)

	userDrive, err := mit.AC.Users().GetDefaultDrive(ctx, mit.User.ID)
	require.NoError(t, err, clues.ToCore(err))

	mit.User.DriveID = ptr.Val(userDrive.GetId())

	userDriveRootFolder, err := mit.AC.Drives().GetRootFolder(ctx, mit.User.DriveID)
	require.NoError(t, err, clues.ToCore(err))

	mit.User.DriveRootFolderID = ptr.Val(userDriveRootFolder.GetId())

	// site

	mit.Site.ID = tconfig.M365SiteID(t)

	siteDrive, err := mit.AC.Sites().GetDefaultDrive(ctx, mit.Site.ID)
	require.NoError(t, err, clues.ToCore(err))

	mit.Site.DriveID = ptr.Val(siteDrive.GetId())

	siteDriveRootFolder, err := mit.AC.Drives().GetRootFolder(ctx, mit.Site.DriveID)
	require.NoError(t, err, clues.ToCore(err))

	mit.Site.DriveRootFolderID = ptr.Val(siteDriveRootFolder.GetId())

	// groups/teams

	// use of the TeamID is intentional here, so that we are assured
	// the group has full usage of the teams api.
	mit.Group.ID = tconfig.M365TeamID(t)
	mit.Group.Email = tconfig.M365TeamEmail(t)

	mit.NonTeamGroup.ID = tconfig.M365GroupID(t)

	channel, err := mit.AC.Channels().
		GetChannelByName(
			ctx,
			mit.Group.ID,
			"Test")
	require.NoError(t, err, clues.ToCore(err))
	require.Equal(t, "Test", ptr.Val(channel.GetDisplayName()))

	mit.Group.TestContainerID = ptr.Val(channel.GetId())

	return mit
}
