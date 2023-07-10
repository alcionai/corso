package test_test

import (
	"context"
	"testing"
	"time"

	"github.com/alcionai/clues"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/alcionai/corso/src/internal/common/dttm"
	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/m365"
	exchMock "github.com/alcionai/corso/src/internal/m365/exchange/mock"
	"github.com/alcionai/corso/src/internal/m365/graph"
	odConsts "github.com/alcionai/corso/src/internal/m365/onedrive/consts"
	"github.com/alcionai/corso/src/internal/m365/resource"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/extensions"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
	"github.com/alcionai/corso/src/pkg/services/m365/api/mock"
)

// Does not use the tester.DefaultTestRestoreDestination syntax as some of these
// items are created directly, not as a result of restoration, and we want to ensure
// they get clearly selected without accidental overlap.
const incrementalsDestContainerPrefix = "incrementals_ci_"

// ---------------------------------------------------------------------------
// Suite Setup
// ---------------------------------------------------------------------------

type intgTesterSetup struct {
	ac                    api.Client
	gockAC                api.Client
	acct                  account.Account
	userID                string
	userDriveID           string
	userDriveRootFolderID string
	siteID                string
	siteDriveID           string
	siteDriveRootFolderID string
}

func newIntegrationTesterSetup(t *testing.T) intgTesterSetup {
	its := intgTesterSetup{}

	ctx, flush := tester.NewContext(t)
	defer flush()

	graph.InitializeConcurrencyLimiter(ctx, true, 4)

	its.acct = tconfig.NewM365Account(t)
	creds, err := its.acct.M365Config()
	require.NoError(t, err, clues.ToCore(err))

	its.ac, err = api.NewClient(creds)
	require.NoError(t, err, clues.ToCore(err))

	its.gockAC, err = mock.NewClient(creds)
	require.NoError(t, err, clues.ToCore(err))

	// user drive

	its.userID = tconfig.M365UserID(t)

	userDrive, err := its.ac.Users().GetDefaultDrive(ctx, its.userID)
	require.NoError(t, err, clues.ToCore(err))

	its.userDriveID = ptr.Val(userDrive.GetId())

	userDriveRootFolder, err := its.ac.Drives().GetRootFolder(ctx, its.userDriveID)
	require.NoError(t, err, clues.ToCore(err))

	its.userDriveRootFolderID = ptr.Val(userDriveRootFolder.GetId())

	its.siteID = tconfig.M365SiteID(t)

	// site

	siteDrive, err := its.ac.Sites().GetDefaultDrive(ctx, its.siteID)
	require.NoError(t, err, clues.ToCore(err))

	its.siteDriveID = ptr.Val(siteDrive.GetId())

	siteDriveRootFolder, err := its.ac.Drives().GetRootFolder(ctx, its.siteDriveID)
	require.NoError(t, err, clues.ToCore(err))

	its.siteDriveRootFolderID = ptr.Val(siteDriveRootFolder.GetId())

	return its
}

// ---------------------------------------------------------------------------
// Incremental Item Generators
// TODO: this is ripped from factory.go, which is ripped from other tests.
// At this point, three variation of the sameish code in three locations
// feels like something we can clean up.  But, it's not a strong need, so
// this gets to stay for now.
// ---------------------------------------------------------------------------

// the params here are what generateContainerOfItems passes into the func.
// the callback provider can use them, or not, as wanted.
type dataBuilderFunc func(id, timeStamp, subject, body string) []byte

func generateContainerOfItems(
	t *testing.T,
	ctx context.Context, //revive:disable-line:context-as-argument
	ctrl *m365.Controller,
	service path.ServiceType,
	cat path.CategoryType,
	sel selectors.Selector,
	tenantID, resourceOwner, driveID string,
	rc control.RestoreConfig,
	howManyItems int,
	backupVersion int,
	dbf dataBuilderFunc,
) *details.Details {
	t.Helper()

	items := make([]incrementalItem, 0, howManyItems)

	for i := 0; i < howManyItems; i++ {
		id, d := generateItemData(t, cat, resourceOwner, dbf)

		items = append(items, incrementalItem{
			name: id,
			data: d,
		})
	}

	pathFolders := []string{}

	switch service {
	case path.OneDriveService, path.SharePointService:
		pathFolders = []string{odConsts.DrivesPathDir, driveID, odConsts.RootPathDir}
	}

	collections := []incrementalCollection{{
		pathFolders: pathFolders,
		category:    cat,
		items:       items,
	}}

	dataColls := buildCollections(
		t,
		service,
		tenantID, resourceOwner,
		collections)

	opts := control.Defaults()
	opts.RestorePermissions = true

	deets, err := ctrl.ConsumeRestoreCollections(
		ctx,
		backupVersion,
		sel,
		rc,
		opts,
		dataColls,
		fault.New(true),
		count.New())
	require.NoError(t, err, clues.ToCore(err))

	// have to wait here, both to ensure the process
	// finishes, and also to clean up the status
	ctrl.Wait()

	return deets
}

func generateItemData(
	t *testing.T,
	category path.CategoryType,
	resourceOwner string,
	dbf dataBuilderFunc,
) (string, []byte) {
	var (
		now       = dttm.Now()
		nowLegacy = dttm.FormatToLegacy(time.Now())
		id        = uuid.NewString()
		subject   = "incr_test " + now[:16] + " - " + id[:8]
		body      = "incr_test " + category.String() + " generation for " + resourceOwner + " at " + now + " - " + id
	)

	return id, dbf(id, nowLegacy, subject, body)
}

type incrementalItem struct {
	name string
	data []byte
}

type incrementalCollection struct {
	pathFolders []string
	category    path.CategoryType
	items       []incrementalItem
}

func buildCollections(
	t *testing.T,
	service path.ServiceType,
	tenant, user string,
	colls []incrementalCollection,
) []data.RestoreCollection {
	t.Helper()

	collections := make([]data.RestoreCollection, 0, len(colls))

	for _, c := range colls {
		pth, err := path.Build(tenant, user, service, c.category, false, c.pathFolders...)
		require.NoError(t, err, clues.ToCore(err))

		mc := exchMock.NewCollection(pth, pth, len(c.items))

		for i := 0; i < len(c.items); i++ {
			mc.Names[i] = c.items[i].name
			mc.Data[i] = c.items[i].data
		}

		collections = append(collections, data.NoFetchRestoreCollection{Collection: mc})
	}

	return collections
}

// A QoL builder for live instances that updates
// the selector's owner id and name in the process
// to help avoid gotchas.
func ControllerWithSelector(
	t *testing.T,
	ctx context.Context, //revive:disable-line:context-as-argument
	acct account.Account,
	cr resource.Category,
	sel selectors.Selector,
	ins idname.Cacher,
	onFail func(*testing.T, context.Context),
) (*m365.Controller, selectors.Selector) {
	ctrl, err := m365.NewController(ctx, acct, cr, sel.PathService(), control.Defaults())
	if !assert.NoError(t, err, clues.ToCore(err)) {
		if onFail != nil {
			onFail(t, ctx)
		}

		t.FailNow()
	}

	id, name, err := ctrl.PopulateOwnerIDAndNamesFrom(ctx, sel.DiscreteOwner, ins)
	if !assert.NoError(t, err, clues.ToCore(err)) {
		if onFail != nil {
			onFail(t, ctx)
		}

		t.FailNow()
	}

	sel = sel.SetDiscreteOwnerIDName(id, name)

	return ctrl, sel
}

func getTestExtensionFactories() []extensions.CreateItemExtensioner {
	return []extensions.CreateItemExtensioner{
		&extensions.MockItemExtensionFactory{},
	}
}

func verifyExtensionData(
	t *testing.T,
	itemInfo details.ItemInfo,
	p path.ServiceType,
) {
	require.NotNil(t, itemInfo.Extension, "nil extension")
	assert.NotNil(t, itemInfo.Extension.Data[extensions.KNumBytes], "key not found in extension")
	actualSize := int64(itemInfo.Extension.Data[extensions.KNumBytes].(float64))

	if p == path.SharePointService {
		assert.Equal(t, itemInfo.SharePoint.Size, actualSize, "incorrect data in extension")
	} else {
		assert.Equal(t, itemInfo.OneDrive.Size, actualSize, "incorrect data in extension")
	}
}
