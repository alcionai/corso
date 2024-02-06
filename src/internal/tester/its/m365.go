package its

import (
	"strings"
	"sync"
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
	DisplayName       string
	DriveID           string
	DriveRootFolderID string
	TestContainerID   string
	WebURL            string
	RootSite          struct {
		Provider          idname.Provider
		ID                string
		DisplayName       string
		DriveID           string
		DriveRootFolderID string
		WebURL            string
	}
}

// M365IntgTestSetup provides all the common references used in an m365 integration
// test suite.  Call `its.GetM365()` to get the singleton for your test suite.
// If you're looking for unit testing setup, use `uts.GetM365()` instead.
type M365IntgTestSetup struct {
	Acct     account.Account
	Creds    account.M365Config
	TenantID string

	AC     api.Client
	GockAC api.Client

	Site          IDs
	SecondarySite IDs

	Group          IDs
	SecondaryGroup IDs

	User          IDs
	SecondaryUser IDs
	TertiaryUser  IDs
}

var (
	singleton *M365IntgTestSetup
	mu        sync.Mutex
)

// GetM365 returns the populated its.m365 singleton.
func GetM365(t *testing.T) M365IntgTestSetup {
	mu.Lock()
	defer mu.Unlock()

	if singleton != nil {
		return *singleton
	}

	var err error

	setup := M365IntgTestSetup{}

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

	fillSite(t, setup.AC, tconfig.M365SiteID(t), &setup.Site)
	fillSite(t, setup.AC, tconfig.SecondaryM365SiteID(t), &setup.SecondarySite)

	// team

	fillTeam(t, setup.AC, tconfig.M365TeamID(t), &setup.Group)
	fillTeam(t, setup.AC, tconfig.SecondaryM365TeamID(t), &setup.SecondaryGroup)

	singleton = &setup

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
	ids.DisplayName = ptr.Val(user.GetDisplayName())

	drive, err := ac.Users().GetDefaultDrive(ctx, ids.ID)
	require.NoError(t, err, clues.ToCore(err))

	ids.DriveID = ptr.Val(drive.GetId())

	rootFolder, err := ac.Drives().GetRootFolder(ctx, ids.DriveID)
	require.NoError(t, err, clues.ToCore(err))

	ids.DriveRootFolderID = ptr.Val(rootFolder.GetId())
}

func fillSite(
	t *testing.T,
	ac api.Client,
	sid string,
	ids *IDs,
) {
	ctx, flush := tester.NewContext(t)
	defer flush()

	site, err := ac.Sites().GetByID(ctx, sid, api.CallConfig{})
	require.NoError(t, err, clues.ToCore(err))

	ids.ID = ptr.Val(site.GetId())
	ids.WebURL = ptr.Val(site.GetWebUrl())
	ids.Provider = idname.NewProvider(ids.ID, ids.WebURL)
	ids.DisplayName = ptr.Val(site.GetDisplayName())

	drive, err := ac.Sites().GetDefaultDrive(ctx, ids.ID)
	require.NoError(t, err, clues.ToCore(err))

	ids.DriveID = ptr.Val(drive.GetId())

	rootFolder, err := ac.Drives().GetRootFolder(ctx, ids.DriveID)
	require.NoError(t, err, clues.ToCore(err))

	ids.DriveRootFolderID = ptr.Val(rootFolder.GetId())
}

func fillTeam(
	t *testing.T,
	ac api.Client,
	gid string,
	ids *IDs,
) {
	ctx, flush := tester.NewContext(t)
	defer flush()

	team, err := ac.Groups().GetByID(ctx, gid, api.CallConfig{})
	require.NoError(t, err, clues.ToCore(err))

	ids.ID = ptr.Val(team.GetId())
	ids.Email = ptr.Val(team.GetMail())
	ids.Provider = idname.NewProvider(ids.ID, ids.Email)
	ids.DisplayName = ptr.Val(team.GetDisplayName())

	channel, err := ac.Channels().
		GetChannelByName(
			ctx,
			ids.ID,
			"Test")
	require.NoError(t, err, clues.ToCore(err))
	require.Equal(t, "Test", ptr.Val(channel.GetDisplayName()))

	ids.TestContainerID = ptr.Val(channel.GetId())

	site, err := ac.Groups().GetRootSite(ctx, gid)
	require.NoError(t, err, clues.ToCore(err))

	ids.RootSite.ID = ptr.Val(site.GetId())
	ids.RootSite.WebURL = ptr.Val(site.GetWebUrl())
	ids.RootSite.DisplayName = ptr.Val(site.GetDisplayName())
	ids.RootSite.Provider = idname.NewProvider(ids.RootSite.ID, ids.RootSite.WebURL)

	drive, err := ac.Sites().GetDefaultDrive(ctx, ids.RootSite.ID)
	require.NoError(t, err, clues.ToCore(err))

	ids.RootSite.DriveID = ptr.Val(drive.GetId())

	rootFolder, err := ac.Drives().GetRootFolder(ctx, ids.RootSite.DriveID)
	require.NoError(t, err, clues.ToCore(err))

	ids.RootSite.DriveRootFolderID = ptr.Val(rootFolder.GetId())
}
