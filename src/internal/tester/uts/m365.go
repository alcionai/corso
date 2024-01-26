package uts

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/require"

	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
)

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

// m365UnitTestSetup provides all the common references used in an m365 unit
// test suite.  Call `uts.GetM365()` to get the values for your test suite.
// If you're looking for integration testing settup, use `its.GetM365()` instead.
type M365UnitTestSetup struct {
	Acct     account.Account
	Creds    account.M365Config
	TenantID string

	Site  IDs
	Group IDs

	User          IDs
	SecondaryUser IDs
	TertiaryUser  IDs
}

// GetM365 returns the populated its.m365 singleton.
func GetM365(t *testing.T) M365UnitTestSetup {
	setup := M365UnitTestSetup{}

	ctx, flush := tester.NewContext(t)
	defer flush()

	graph.InitializeConcurrencyLimiter(ctx, true, 4)

	setup.Acct = tconfig.NewM365Account(t)
	creds, err := setup.Acct.M365Config()
	require.NoError(t, err, clues.ToCore(err))

	setup.Creds = creds
	setup.TenantID = creds.AzureTenantID

	// users

	fillUser(t, tconfig.M365UserID(t), &setup.User)
	fillUser(t, tconfig.SecondaryM365UserID(t), &setup.SecondaryUser)
	fillUser(t, tconfig.TertiaryM365UserID(t), &setup.TertiaryUser)

	// site

	setup.Site.ID = tconfig.M365SiteID(t)
	setup.Site.WebURL = setup.Site.ID
	setup.Site.Provider = idname.NewProvider(setup.Site.ID, setup.Site.WebURL)
	setup.Site.DriveID = "drive-id-" + setup.Site.ID
	setup.Site.DriveRootFolderID = "root-folder-" + setup.Site.ID

	// team group

	setup.Group.Email = tconfig.M365TeamEmail(t)
	setup.Group.ID = tconfig.M365TeamID(t)
	setup.Group.Provider = idname.NewProvider(setup.Group.ID, setup.Group.Email)
	setup.Group.TestContainerID = "channel-id-" + setup.Group.ID

	return setup
}

func fillUser(
	t *testing.T,
	uid string,
	ids *IDs,
) {
	ids.ID = tconfig.M365UserID(t)
	ids.Email = ids.ID
	ids.Provider = idname.NewProvider(ids.ID, ids.Email)
	ids.DriveID = "drive-id-" + ids.ID
	ids.DriveRootFolderID = "root-folder-" + ids.ID
}
