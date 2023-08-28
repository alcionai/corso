package test_test

import (
	"context"
	"fmt"
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
	"github.com/alcionai/corso/src/internal/events"
	evmock "github.com/alcionai/corso/src/internal/events/mock"
	"github.com/alcionai/corso/src/internal/kopia"
	"github.com/alcionai/corso/src/internal/m365"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/internal/m365/resource"
	exchMock "github.com/alcionai/corso/src/internal/m365/service/exchange/mock"
	odConsts "github.com/alcionai/corso/src/internal/m365/service/onedrive/consts"
	"github.com/alcionai/corso/src/internal/model"
	"github.com/alcionai/corso/src/internal/operations"
	"github.com/alcionai/corso/src/internal/operations/inject"
	"github.com/alcionai/corso/src/internal/streamstore"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/backup"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/backup/identity"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/control/repository"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/extensions"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
	"github.com/alcionai/corso/src/pkg/services/m365/api/mock"
	"github.com/alcionai/corso/src/pkg/storage"
	storeTD "github.com/alcionai/corso/src/pkg/storage/testdata"
	"github.com/alcionai/corso/src/pkg/store"
)

// Does not use the tester.DefaultTestRestoreDestination syntax as some of these
// items are created directly, not as a result of restoration, and we want to ensure
// they get clearly selected without accidental overlap.
const incrementalsDestContainerPrefix = "incrementals_ci_"

type backupOpDependencies struct {
	acct account.Account
	ctrl *m365.Controller
	kms  *kopia.ModelStore
	kw   *kopia.Wrapper
	sel  selectors.Selector
	sss  streamstore.Streamer
	st   storage.Storage
	sw   store.BackupStorer

	closer func()
}

func (bod *backupOpDependencies) close(
	t *testing.T,
	ctx context.Context, //revive:disable-line:context-as-argument
) {
	bod.closer()

	if bod.kw != nil {
		err := bod.kw.Close(ctx)
		assert.NoErrorf(t, err, "kw close: %+v", clues.ToCore(err))
	}

	if bod.kms != nil {
		err := bod.kw.Close(ctx)
		assert.NoErrorf(t, err, "kms close: %+v", clues.ToCore(err))
	}
}

// prepNewTestBackupOp generates all clients required to run a backup operation,
// returning both a backup operation created with those clients, as well as
// the clients themselves.
func prepNewTestBackupOp(
	t *testing.T,
	ctx context.Context, //revive:disable-line:context-as-argument
	bus events.Eventer,
	sel selectors.Selector,
	opts control.Options,
	backupVersion int,
) (
	operations.BackupOperation,
	*backupOpDependencies,
) {
	bod := &backupOpDependencies{
		acct: tconfig.NewM365Account(t),
		st:   storeTD.NewPrefixedS3Storage(t),
	}

	k := kopia.NewConn(bod.st)

	err := k.Initialize(ctx, repository.Options{}, repository.Retention{})
	require.NoError(t, err, clues.ToCore(err))

	defer func() {
		if err != nil {
			bod.close(t, ctx)
			t.FailNow()
		}
	}()

	// kopiaRef comes with a count of 1 and Wrapper bumps it again
	// we're so safe to close here.
	bod.closer = func() {
		err := k.Close(ctx)
		assert.NoErrorf(t, err, "k close: %+v", clues.ToCore(err))
	}

	bod.kw, err = kopia.NewWrapper(k)
	if !assert.NoError(t, err, clues.ToCore(err)) {
		return operations.BackupOperation{}, nil
	}

	bod.kms, err = kopia.NewModelStore(k)
	if !assert.NoError(t, err, clues.ToCore(err)) {
		return operations.BackupOperation{}, nil
	}

	bod.sw = store.NewWrapper(bod.kms)

	connectorResource := resource.Users

	switch sel.Service {
	case selectors.ServiceSharePoint:
		connectorResource = resource.Sites
	case selectors.ServiceGroups:
		connectorResource = resource.Groups
	}

	bod.ctrl, bod.sel = ControllerWithSelector(
		t,
		ctx,
		bod.acct,
		connectorResource,
		sel,
		nil,
		bod.close)

	bo := newTestBackupOp(
		t,
		ctx,
		bod,
		bus,
		opts)

	bod.sss = streamstore.NewStreamer(
		bod.kw,
		bod.acct.ID(),
		bod.sel.PathService())

	return bo, bod
}

// newTestBackupOp accepts the clients required to compose a backup operation, plus
// any other metadata, and uses them to generate a new backup operation.  This
// allows backup chains to utilize the same temp directory and configuration
// details.
func newTestBackupOp(
	t *testing.T,
	ctx context.Context, //revive:disable-line:context-as-argument
	bod *backupOpDependencies,
	bus events.Eventer,
	opts control.Options,
) operations.BackupOperation {
	bod.ctrl.IDNameLookup = idname.NewCache(map[string]string{bod.sel.ID(): bod.sel.Name()})

	bo, err := operations.NewBackupOperation(
		ctx,
		opts,
		bod.kw,
		bod.sw,
		bod.ctrl,
		bod.acct,
		bod.sel,
		bod.sel,
		bus)
	if !assert.NoError(t, err, clues.ToCore(err)) {
		bod.close(t, ctx)
		t.FailNow()
	}

	return bo
}

func runAndCheckBackup(
	t *testing.T,
	ctx context.Context, //revive:disable-line:context-as-argument
	bo *operations.BackupOperation,
	mb *evmock.Bus,
	acceptNoData bool,
) {
	err := bo.Run(ctx)
	require.NoError(t, err, clues.ToCore(err))
	require.NotEmpty(t, bo.Results, "the backup had non-zero results")
	require.NotEmpty(t, bo.Results.BackupID, "the backup generated an ID")

	expectStatus := []operations.OpStatus{operations.Completed}
	if acceptNoData {
		expectStatus = append(expectStatus, operations.NoData)
	}

	require.Contains(
		t,
		expectStatus,
		bo.Status,
		"backup doesn't match expectation, wanted any of %v, got %s",
		expectStatus,
		bo.Status)

	require.NotZero(t, bo.Results.ItemsWritten)
	assert.NotZero(t, bo.Results.ItemsRead, "count of items read")
	assert.NotZero(t, bo.Results.BytesRead, "bytes read")
	assert.NotZero(t, bo.Results.BytesUploaded, "bytes uploaded")
	assert.Equal(t, 1, bo.Results.ResourceOwners, "count of resource owners")
	assert.NoError(t, bo.Errors.Failure(), "incremental non-recoverable error", clues.ToCore(bo.Errors.Failure()))
	assert.Empty(t, bo.Errors.Recovered(), "incremental recoverable/iteration errors")
	assert.Equal(t, 1, mb.TimesCalled[events.BackupStart], "backup-start events")
	assert.Equal(t, 1, mb.TimesCalled[events.BackupEnd], "backup-end events")
	assert.Equal(t,
		mb.CalledWith[events.BackupStart][0][events.BackupID],
		bo.Results.BackupID, "backupID pre-declaration")
}

func checkBackupIsInManifests(
	t *testing.T,
	ctx context.Context, //revive:disable-line:context-as-argument
	kw *kopia.Wrapper,
	sw store.BackupStorer,
	bo *operations.BackupOperation,
	sel selectors.Selector,
	resourceOwner string,
	categories ...path.CategoryType,
) {
	for _, category := range categories {
		t.Run(category.String(), func(t *testing.T) {
			var (
				r     = kopia.NewReason("", resourceOwner, sel.PathService(), category)
				tags  = map[string]string{kopia.TagBackupCategory: ""}
				found bool
			)

			bf, err := kw.NewBaseFinder(sw)
			require.NoError(t, err, clues.ToCore(err))

			mans := bf.FindBases(ctx, []identity.Reasoner{r}, tags)
			for _, man := range mans.MergeBases() {
				bID, ok := man.GetTag(kopia.TagBackupID)
				if !assert.Truef(t, ok, "snapshot manifest %s missing backup ID tag", man.ID) {
					continue
				}

				if bID == string(bo.Results.BackupID) {
					found = true
					break
				}
			}

			assert.True(t, found, "backup retrieved by previous snapshot manifest")
		})
	}
}

func checkMetadataFilesExist(
	t *testing.T,
	ctx context.Context, //revive:disable-line:context-as-argument
	backupID model.StableID,
	kw *kopia.Wrapper,
	ms *kopia.ModelStore,
	tenant, resourceOwner string,
	service path.ServiceType,
	filesByCat map[path.CategoryType][]string,
) {
	for category, files := range filesByCat {
		t.Run(category.String(), func(t *testing.T) {
			bup := &backup.Backup{}

			err := ms.Get(ctx, model.BackupSchema, backupID, bup)
			if !assert.NoError(t, err, clues.ToCore(err)) {
				return
			}

			paths := []path.RestorePaths{}
			pathsByRef := map[string][]string{}

			for _, fName := range files {
				p, err := path.Builder{}.
					Append(fName).
					ToServiceCategoryMetadataPath(tenant, resourceOwner, service, category, true)
				if !assert.NoError(t, err, "bad metadata path", clues.ToCore(err)) {
					continue
				}

				dir, err := p.Dir()
				if !assert.NoError(t, err, "parent path", clues.ToCore(err)) {
					continue
				}

				paths = append(
					paths,
					path.RestorePaths{StoragePath: p, RestorePath: dir})
				pathsByRef[dir.ShortRef()] = append(pathsByRef[dir.ShortRef()], fName)
			}

			cols, err := kw.ProduceRestoreCollections(
				ctx,
				bup.SnapshotID,
				paths,
				nil,
				fault.New(true))
			assert.NoError(t, err, clues.ToCore(err))

			for _, col := range cols {
				itemNames := []string{}

				for item := range col.Items(ctx, fault.New(true)) {
					assert.Implements(t, (*data.ItemSize)(nil), item)

					s := item.(data.ItemSize)
					assert.Greaterf(
						t,
						s.Size(),
						int64(0),
						"empty metadata file: %s/%s",
						col.FullPath(),
						item.ID(),
					)

					itemNames = append(itemNames, item.ID())
				}

				assert.ElementsMatchf(
					t,
					pathsByRef[col.FullPath().ShortRef()],
					itemNames,
					"collection %s missing expected files",
					col.FullPath(),
				)
			}
		})
	}
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
	tenantID, resourceOwner, driveID, destFldr string,
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

	pathFolders := []string{destFldr}

	switch service {
	case path.OneDriveService, path.SharePointService:
		pathFolders = []string{odConsts.DrivesPathDir, driveID, odConsts.RootPathDir, destFldr}
	}

	collections := []incrementalCollection{{
		pathFolders: pathFolders,
		category:    cat,
		items:       items,
	}}

	restoreCfg := control.DefaultRestoreConfig(dttm.SafeForTesting)
	restoreCfg.Location = destFldr
	restoreCfg.IncludePermissions = true

	dataColls := buildCollections(
		t,
		service,
		tenantID, resourceOwner,
		restoreCfg,
		collections)

	opts := control.DefaultOptions()

	rcc := inject.RestoreConsumerConfig{
		BackupVersion:     backupVersion,
		Options:           opts,
		ProtectedResource: sel,
		RestoreConfig:     restoreCfg,
		Selector:          sel,
	}

	deets, err := ctrl.ConsumeRestoreCollections(
		ctx,
		rcc,
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
	restoreCfg control.RestoreConfig,
	colls []incrementalCollection,
) []data.RestoreCollection {
	t.Helper()

	collections := make([]data.RestoreCollection, 0, len(colls))

	for _, c := range colls {
		pth := toDataLayerPath(
			t,
			service,
			tenant,
			user,
			c.category,
			c.pathFolders,
			false)

		mc := exchMock.NewCollection(pth, pth, len(c.items))

		for i := 0; i < len(c.items); i++ {
			mc.Names[i] = c.items[i].name
			mc.Data[i] = c.items[i].data
		}

		collections = append(collections, data.NoFetchRestoreCollection{Collection: mc})
	}

	return collections
}

func toDataLayerPath(
	t *testing.T,
	service path.ServiceType,
	tenant, resourceOwner string,
	category path.CategoryType,
	elements []string,
	isItem bool,
) path.Path {
	t.Helper()

	var (
		pb  = path.Builder{}.Append(elements...)
		p   path.Path
		err error
	)

	switch service {
	case path.ExchangeService:
		p, err = pb.ToDataLayerExchangePathForCategory(tenant, resourceOwner, category, isItem)
	case path.OneDriveService:
		p, err = pb.ToDataLayerOneDrivePath(tenant, resourceOwner, isItem)
	case path.SharePointService:
		p, err = pb.ToDataLayerSharePointPath(tenant, resourceOwner, category, isItem)
	case path.GroupsService:
		p, err = pb.ToDataLayerPath(tenant, resourceOwner, service, category, false)
	default:
		err = clues.New(fmt.Sprintf("unknown service: %s", service))
	}

	require.NoError(t, err, clues.ToCore(err))

	return p
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
	ctrl, err := m365.NewController(ctx, acct, cr, sel.PathService(), control.DefaultOptions())
	if !assert.NoError(t, err, clues.ToCore(err)) {
		if onFail != nil {
			onFail(t, ctx)
		}

		t.FailNow()
	}

	id, name, err := ctrl.PopulateProtectedResourceIDAndName(ctx, sel.DiscreteOwner, ins)
	if !assert.NoError(t, err, clues.ToCore(err)) {
		if onFail != nil {
			onFail(t, ctx)
		}

		t.FailNow()
	}

	sel = sel.SetDiscreteOwnerIDName(id, name)

	return ctrl, sel
}

// ---------------------------------------------------------------------------
// Suite Setup
// ---------------------------------------------------------------------------

type ids struct {
	ID                string
	DriveID           string
	DriveRootFolderID string
}

type gids struct {
	ID                        string
	RootSiteID                string
	RootSiteDriveID           string
	RootSiteDriveRootFolderID string
}

type intgTesterSetup struct {
	ac            api.Client
	gockAC        api.Client
	user          ids
	secondaryUser ids
	site          ids
	secondarySite ids
	group         gids
}

func newIntegrationTesterSetup(t *testing.T) intgTesterSetup {
	its := intgTesterSetup{}

	ctx, flush := tester.NewContext(t)
	defer flush()

	graph.InitializeConcurrencyLimiter(ctx, true, 4)

	a := tconfig.NewM365Account(t)
	creds, err := a.M365Config()
	require.NoError(t, err, clues.ToCore(err))

	its.ac, err = api.NewClient(creds, control.DefaultOptions())
	require.NoError(t, err, clues.ToCore(err))

	its.gockAC, err = mock.NewClient(creds)
	require.NoError(t, err, clues.ToCore(err))

	its.user = userIDs(t, tconfig.M365UserID(t), its.ac)
	its.secondaryUser = userIDs(t, tconfig.SecondaryM365UserID(t), its.ac)
	its.site = siteIDs(t, tconfig.M365SiteID(t), its.ac)
	its.secondarySite = siteIDs(t, tconfig.SecondaryM365SiteID(t), its.ac)
	// teamID is used here intentionally.  We want the group
	// to have access to teams data
	its.group = groupIDs(t, tconfig.M365TeamID(t), its.ac)

	return its
}

func userIDs(t *testing.T, id string, ac api.Client) ids {
	ctx, flush := tester.NewContext(t)
	defer flush()

	r := ids{ID: id}

	drive, err := ac.Users().GetDefaultDrive(ctx, id)
	require.NoError(t, err, clues.ToCore(err))

	r.DriveID = ptr.Val(drive.GetId())

	driveRootFolder, err := ac.Drives().GetRootFolder(ctx, r.DriveID)
	require.NoError(t, err, clues.ToCore(err))

	r.DriveRootFolderID = ptr.Val(driveRootFolder.GetId())

	return r
}

func siteIDs(t *testing.T, id string, ac api.Client) ids {
	ctx, flush := tester.NewContext(t)
	defer flush()

	r := ids{ID: id}

	drive, err := ac.Sites().GetDefaultDrive(ctx, id)
	require.NoError(t, err, clues.ToCore(err))

	r.DriveID = ptr.Val(drive.GetId())

	driveRootFolder, err := ac.Drives().GetRootFolder(ctx, r.DriveID)
	require.NoError(t, err, clues.ToCore(err))

	r.DriveRootFolderID = ptr.Val(driveRootFolder.GetId())

	return r
}

func groupIDs(t *testing.T, id string, ac api.Client) gids {
	ctx, flush := tester.NewContext(t)
	defer flush()

	r := gids{ID: id}

	site, err := ac.Groups().GetRootSite(ctx, id)
	require.NoError(t, err, clues.ToCore(err))

	r.RootSiteID = ptr.Val(site.GetId())

	drive, err := ac.Sites().GetDefaultDrive(ctx, r.RootSiteID)
	require.NoError(t, err, clues.ToCore(err))

	r.RootSiteDriveID = ptr.Val(drive.GetId())

	driveRootFolder, err := ac.Drives().GetRootFolder(ctx, r.RootSiteDriveID)
	require.NoError(t, err, clues.ToCore(err))

	r.RootSiteDriveRootFolderID = ptr.Val(driveRootFolder.GetId())

	return r
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
