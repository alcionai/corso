package api

import (
	"strings"
	"testing"

	"github.com/alcionai/clues"
	"github.com/h2non/gock"
	"github.com/stretchr/testify/require"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
)

// ---------------------------------------------------------------------------
// Gockable client
// ---------------------------------------------------------------------------

// GockClient produces a new exchange api client that can be
// mocked using gock.
func gockClient(creds account.M365Config, counter *count.Bus) (Client, error) {
	s, err := graph.NewGockService(creds, counter)
	if err != nil {
		return Client{}, err
	}

	li, err := graph.NewGockService(creds, counter, graph.NoTimeout())
	if err != nil {
		return Client{}, err
	}

	return Client{
		Credentials: creds,
		Stable:      s,
		LargeItem:   li,
		options:     control.DefaultOptions(),
	}, nil
}

// ---------------------------------------------------------------------------
// Intercepting calls with Gock
// ---------------------------------------------------------------------------

const graphAPIHostURL = "https://graph.microsoft.com"

func v1APIURLPath(parts ...string) string {
	return strings.Join(append([]string{"/v1.0"}, parts...), "/")
}

func interceptV1Path(pathParts ...string) *gock.Request {
	return gock.New(graphAPIHostURL).Get(v1APIURLPath(pathParts...))
}

// ---------------------------------------------------------------------------
// Suite Setup
// ---------------------------------------------------------------------------

type ids struct {
	id                string
	email             string
	driveID           string
	driveRootFolderID string
	testContainerID   string
}

type intgTesterSetup struct {
	ac           Client
	gockAC       Client
	user         ids
	site         ids
	group        ids
	nonTeamGroup ids // group which does not have an associated team
}

func newIntegrationTesterSetup(t *testing.T) intgTesterSetup {
	its := intgTesterSetup{}

	ctx, flush := tester.NewContext(t)
	defer flush()

	graph.InitializeConcurrencyLimiter(ctx, true, 4)

	a := tconfig.NewM365Account(t)
	creds, err := a.M365Config()
	require.NoError(t, err, clues.ToCore(err))

	its.ac, err = NewClient(creds, control.DefaultOptions(), count.New())
	require.NoError(t, err, clues.ToCore(err))

	its.gockAC, err = gockClient(creds, count.New())
	require.NoError(t, err, clues.ToCore(err))

	// user drive

	its.user.id = tconfig.M365UserID(t)

	userDrive, err := its.ac.Users().GetDefaultDrive(ctx, its.user.id)
	require.NoError(t, err, clues.ToCore(err))

	its.user.driveID = ptr.Val(userDrive.GetId())

	userDriveRootFolder, err := its.ac.Drives().GetRootFolder(ctx, its.user.driveID)
	require.NoError(t, err, clues.ToCore(err))

	its.user.driveRootFolderID = ptr.Val(userDriveRootFolder.GetId())

	// site

	its.site.id = tconfig.M365SiteID(t)

	siteDrive, err := its.ac.Sites().GetDefaultDrive(ctx, its.site.id)
	require.NoError(t, err, clues.ToCore(err))

	its.site.driveID = ptr.Val(siteDrive.GetId())

	siteDriveRootFolder, err := its.ac.Drives().GetRootFolder(ctx, its.site.driveID)
	require.NoError(t, err, clues.ToCore(err))

	its.site.driveRootFolderID = ptr.Val(siteDriveRootFolder.GetId())

	// groups/teams

	// use of the TeamID is intentional here, so that we are assured
	// the group has full usage of the teams api.
	its.group.id = tconfig.M365TeamID(t)
	its.group.email = tconfig.M365TeamEmail(t)

	its.nonTeamGroup.id = tconfig.M365GroupID(t)

	channel, err := its.ac.Channels().
		GetChannelByName(
			ctx,
			its.group.id,
			"Test")
	require.NoError(t, err, clues.ToCore(err))
	require.Equal(t, "Test", ptr.Val(channel.GetDisplayName()))

	its.group.testContainerID = ptr.Val(channel.GetId())

	return its
}
