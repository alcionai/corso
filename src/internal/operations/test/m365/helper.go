package m365

import (
	"context"
	"testing"
	"time"

	"github.com/alcionai/clues"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/kopia"
	"github.com/alcionai/corso/src/internal/m365"
	exchMock "github.com/alcionai/corso/src/internal/m365/service/exchange/mock"
	odConsts "github.com/alcionai/corso/src/internal/m365/service/onedrive/consts"
	"github.com/alcionai/corso/src/internal/model"
	"github.com/alcionai/corso/src/internal/operations/inject"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/backup"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/dttm"
	"github.com/alcionai/corso/src/pkg/extensions"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
)

// ---------------------------------------------------------------------------
// Gockable client
// ---------------------------------------------------------------------------

// GockClient produces a new exchange api client that can be
// mocked using gock.
func GockClient(creds account.M365Config, counter *count.Bus) (api.Client, error) {
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

// Does not use the tester.DefaultTestRestoreDestination syntax as some of these
// items are created directly, not as a result of restoration, and we want to ensure
// they get clearly selected without accidental overlap.
const IncrementalsDestContainerPrefix = "incrementals_ci_"

func CheckMetadataFilesExist(
	t *testing.T,
	ctx context.Context, //revive:disable-line:context-as-argument
	backupID model.StableID,
	kw *kopia.Wrapper,
	ms *kopia.ModelStore,
	tenant, resourceOwner string,
	service path.ServiceType,
	filesByCat map[path.CategoryType][][]string,
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
				p, err := path.BuildMetadata(tenant, resourceOwner, service, category, true, fName...)
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
				pathsByRef[dir.ShortRef()] = append(pathsByRef[dir.ShortRef()], fName[len(fName)-1])
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
						item.ID())

					itemNames = append(itemNames, item.ID())
				}

				assert.ElementsMatchf(
					t,
					pathsByRef[col.FullPath().ShortRef()],
					itemNames,
					"collection %s missing expected files",
					col.FullPath())
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
type DataBuilderFunc func(id, timeStamp, subject, body string) []byte

func GenerateContainerOfItems(
	t *testing.T,
	ctx context.Context, //revive:disable-line:context-as-argument
	ctrl *m365.Controller,
	service path.ServiceType,
	cat path.CategoryType,
	sel selectors.Selector,
	tenantID, resourceOwner, siteID, driveID, destFldr string,
	howManyItems int,
	backupVersion int,
	dbf DataBuilderFunc,
) *details.Details {
	t.Helper()

	items := make([]IncrementalItem, 0, howManyItems)

	for i := 0; i < howManyItems; i++ {
		id, d := GenerateItemData(t, cat, resourceOwner, dbf)

		items = append(items, IncrementalItem{
			name: id,
			data: d,
		})
	}

	pathFolders := []string{destFldr}

	switch service {
	case path.OneDriveService, path.SharePointService:
		pathFolders = []string{odConsts.DrivesPathDir, driveID, odConsts.RootPathDir, destFldr}
	case path.GroupsService:
		pathFolders = []string{odConsts.SitesPathDir, siteID, odConsts.DrivesPathDir, driveID, odConsts.RootPathDir, destFldr}
	}

	collections := []IncrementalCollection{{
		pathFolders: pathFolders,
		category:    cat,
		items:       items,
	}}

	restoreCfg := control.DefaultRestoreConfig(dttm.SafeForTesting)
	restoreCfg.Location = destFldr
	restoreCfg.IncludePermissions = true

	dataColls := BuildCollections(
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

	handler, err := ctrl.NewServiceHandler(service)
	require.NoError(t, err, clues.ToCore(err))

	deets, _, err := handler.ConsumeRestoreCollections(
		ctx,
		rcc,
		dataColls,
		fault.New(true),
		count.New())
	require.NoError(t, err, clues.ToCore(err))

	return deets
}

func GenerateItemData(
	t *testing.T,
	category path.CategoryType,
	resourceOwner string,
	dbf DataBuilderFunc,
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

type IncrementalItem struct {
	name string
	data []byte
}

type IncrementalCollection struct {
	pathFolders []string
	category    path.CategoryType
	items       []IncrementalItem
}

func BuildCollections(
	t *testing.T,
	service path.ServiceType,
	tenant, user string,
	restoreCfg control.RestoreConfig,
	colls []IncrementalCollection,
) []data.RestoreCollection {
	t.Helper()

	collections := make([]data.RestoreCollection, 0, len(colls))

	for _, c := range colls {
		pth := ToDataLayerPath(
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

func ToDataLayerPath(
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

	p, err = pb.ToDataLayerPath(tenant, resourceOwner, service, category, isItem)
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
	sel selectors.Selector,
	ins idname.Cacher,
	onFail func(*testing.T, context.Context),
	counter *count.Bus,
) (*m365.Controller, selectors.Selector) {
	ctx = clues.Add(ctx, "controller_selector", sel)

	ctrl, err := m365.NewController(
		ctx,
		acct,
		sel.PathService(),
		control.DefaultOptions(),
		counter)
	if !assert.NoError(t, err, clues.ToCore(err)) {
		if onFail != nil {
			onFail(t, ctx)
		}

		t.FailNow()
	}

	resource, err := ctrl.PopulateProtectedResourceIDAndName(ctx, sel.DiscreteOwner, ins)
	if !assert.NoError(t, err, clues.ToCore(err)) {
		if onFail != nil {
			onFail(t, ctx)
		}

		t.FailNow()
	}

	sel = sel.SetDiscreteOwnerIDName(resource.ID(), resource.Name())

	return ctrl, sel
}

// ---------------------------------------------------------------------------
// Suite Setup
// ---------------------------------------------------------------------------

func GetTestExtensionFactories() []extensions.CreateItemExtensioner {
	return []extensions.CreateItemExtensioner{
		&extensions.MockItemExtensionFactory{},
	}
}

func VerifyExtensionData(
	t *testing.T,
	itemInfo details.ItemInfo,
	p path.ServiceType,
) {
	require.NotNil(t, itemInfo.Extension, "nil extension")
	assert.NotNil(t, itemInfo.Extension.Data[extensions.KNumBytes], "key not found in extension")

	var (
		detailsSize   int64
		extensionSize = int64(itemInfo.Extension.Data[extensions.KNumBytes].(float64))
	)

	switch p {
	case path.SharePointService:
		detailsSize = itemInfo.SharePoint.Size
	case path.OneDriveService:
		detailsSize = itemInfo.OneDrive.Size
	case path.GroupsService:
		// FIXME: needs update for message.
		detailsSize = itemInfo.Groups.Size
	default:
		assert.Fail(t, "unrecognized data type")
	}

	assert.Equal(t, extensionSize, detailsSize, "incorrect size in extension")
}
