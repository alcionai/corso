package its

import (
	"strings"
	"testing"

	"github.com/alcionai/clues"
	"github.com/h2non/gock"
	"github.com/stretchr/testify/require"

	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
)

// ---------------------------------------------------------------------------
// Gockable client
// ---------------------------------------------------------------------------

// GockClient produces a new exchange api client that can be
// mocked using gock.
func gockClient(creds account.M365Config, counter *count.Bus) (api.Client, error) {
	s, err := graph.NewGockService(creds, counter)
	if err != nil {
		return api.Client{}, err
	}

	li, err := graph.NewGockService(creds, counter, graph.NoTimeout())
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
// Intercepting calls with Gock
// ---------------------------------------------------------------------------

const graphAPIHostURL = "https://graph.microsoft.com"

func V1APIURLPath(parts ...string) string {
	return strings.Join(append([]string{"/v1.0"}, parts...), "/")
}

func InterceptV1Path(pathParts ...string) *gock.Request {
	return gock.New(graphAPIHostURL).Get(V1APIURLPath(pathParts...))
}

// ---------------------------------------------------------------------------
// Suite Setup
// ---------------------------------------------------------------------------

type IDs struct {
	Provider          idname.Provider
	ID                string
	Email             string
	DriveID           string
	DriveRootFolderID string
	TestContainerID   string
	WebURL            string
}

// m365IntgTestSetup provides all the common references used in an m365 integration
// test suite.  Call `its.GetM365()` to get the singleton for your test suite.
// If you're looking for unit testing setup, use `uts.GetM365()` instead.
type m365IntgTestSetup struct {
	Acct     account.Account
	Creds    account.M365Config
	TenantID string

	AC     api.Client
	GockAC api.Client

	Site  IDs
	Group IDs

	User          IDs
	SecondaryUser IDs
	TertiaryUser  IDs
}

// GetM365 returns the populated its.m365 singleton.
func GetM365(t *testing.T) m365IntgTestSetup {
	var err error

	setup := m365IntgTestSetup{}

	ctx, flush := tester.NewContext(t)
	defer flush()

	graph.InitializeConcurrencyLimiter(ctx, true, 4)

	setup.Acct = tconfig.NewM365Account(t)
	setup.Creds, err = setup.Acct.M365Config()
	require.NoError(t, err, clues.ToCore(err))
	setup.TenantID = setup.Creds.AzureTenantID

	setup.AC, err = api.NewClient(setup.Creds, control.DefaultOptions(), count.New())
	require.NoError(t, err, clues.ToCore(err))

	setup.GockAC, err = gockClient(setup.Creds, count.New())
	require.NoError(t, err, clues.ToCore(err))

	// users

	fillUser(t, setup.AC, tconfig.M365UserID(t), &setup.User)
	fillUser(t, setup.AC, tconfig.SecondaryM365UserID(t), &setup.SecondaryUser)
	fillUser(t, setup.AC, tconfig.TertiaryM365UserID(t), &setup.TertiaryUser)

	// site

	setup.Site.ID = tconfig.M365SiteID(t)

	site, err := setup.AC.Sites().GetByID(ctx, setup.Site.ID, api.CallConfig{})
	require.NoError(t, err, clues.ToCore(err))

	setup.Site.ID = ptr.Val(site.GetId())
	setup.Site.WebURL = ptr.Val(site.GetWebUrl())
	setup.Site.Provider = idname.NewProvider(setup.Site.ID, setup.Site.WebURL)

	siteDrive, err := setup.AC.Sites().GetDefaultDrive(ctx, setup.Site.ID)
	require.NoError(t, err, clues.ToCore(err))

	setup.Site.DriveID = ptr.Val(siteDrive.GetId())

	siteDriveRootFolder, err := setup.AC.Drives().GetRootFolder(ctx, setup.Site.DriveID)
	require.NoError(t, err, clues.ToCore(err))

	setup.Site.DriveRootFolderID = ptr.Val(siteDriveRootFolder.GetId())

	// team group

	setup.Group.Email = tconfig.M365TeamEmail(t)

	// use of the TeamID is intentional here, so that we are assured
	// the team has full usage of the teams api.
	team, err := setup.AC.Groups().GetByID(ctx, tconfig.M365TeamID(t), api.CallConfig{})
	require.NoError(t, err, clues.ToCore(err))

	setup.Group.ID = ptr.Val(team.GetId())
	setup.Group.Provider = idname.NewProvider(setup.Group.ID, setup.Group.Email)

	channel, err := setup.AC.Channels().
		GetChannelByName(
			ctx,
			setup.Group.ID,
			"Test")
	require.NoError(t, err, clues.ToCore(err))
	require.Equal(t, "Test", ptr.Val(channel.GetDisplayName()))

	setup.Group.TestContainerID = ptr.Val(channel.GetId())

	return setup
}

func fillUser(
	t *testing.T,
	ac api.Client,
	uid string,
	ids *IDs,
) {
	ctx, flush := tester.NewContext(t)
	defer flush()

	ids.ID = tconfig.M365UserID(t)

	user, err := ac.Users().GetByID(ctx, uid, api.CallConfig{})
	require.NoError(t, err, clues.ToCore(err))

	ids.ID = ptr.Val(user.GetId())
	ids.Email = ptr.Val(user.GetUserPrincipalName())
	ids.Provider = idname.NewProvider(ids.ID, ids.Email)

	drive, err := ac.Users().GetDefaultDrive(ctx, ids.ID)
	require.NoError(t, err, clues.ToCore(err))

	ids.DriveID = ptr.Val(drive.GetId())

	rootFolder, err := ac.Drives().GetRootFolder(ctx, ids.DriveID)
	require.NoError(t, err, clues.ToCore(err))

	ids.DriveRootFolderID = ptr.Val(rootFolder.GetId())
}
