package operations

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/alcionai/clues"
	"github.com/google/uuid"
	"github.com/microsoftgraph/msgraph-sdk-go/drive"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/users"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"golang.org/x/exp/maps"

	"github.com/alcionai/corso/src/internal/common/dttm"
	"github.com/alcionai/corso/src/internal/common/idname"
	inMock "github.com/alcionai/corso/src/internal/common/idname/mock"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/connector"
	"github.com/alcionai/corso/src/internal/connector/exchange"
	exapi "github.com/alcionai/corso/src/internal/connector/exchange/api"
	exchMock "github.com/alcionai/corso/src/internal/connector/exchange/mock"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/mock"
	"github.com/alcionai/corso/src/internal/connector/onedrive"
	odapi "github.com/alcionai/corso/src/internal/connector/onedrive/api"
	"github.com/alcionai/corso/src/internal/connector/onedrive/metadata"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/events"
	evmock "github.com/alcionai/corso/src/internal/events/mock"
	"github.com/alcionai/corso/src/internal/kopia"
	"github.com/alcionai/corso/src/internal/model"
	"github.com/alcionai/corso/src/internal/operations/inject"
	"github.com/alcionai/corso/src/internal/streamstore"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/version"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/backup"
	"github.com/alcionai/corso/src/pkg/backup/details"
	deeTD "github.com/alcionai/corso/src/pkg/backup/details/testdata"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/control/repository"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	selTD "github.com/alcionai/corso/src/pkg/selectors/testdata"
	"github.com/alcionai/corso/src/pkg/store"
)

// Does not use the tester.DefaultTestRestoreDestination syntax as some of these
// items are created directly, not as a result of restoration, and we want to ensure
// they get clearly selected without accidental overlap.
const incrementalsDestContainerPrefix = "incrementals_ci_"

// ---------------------------------------------------------------------------
// helpers
// ---------------------------------------------------------------------------

// prepNewTestBackupOp generates all clients required to run a backup operation,
// returning both a backup operation created with those clients, as well as
// the clients themselves.
func prepNewTestBackupOp(
	t *testing.T,
	ctx context.Context, //revive:disable-line:context-as-argument
	bus events.Eventer,
	sel selectors.Selector,
	featureToggles control.Toggles,
	backupVersion int,
) (
	BackupOperation,
	account.Account,
	*kopia.Wrapper,
	*kopia.ModelStore,
	streamstore.Streamer,
	*connector.GraphConnector,
	selectors.Selector,
	func(),
) {
	var (
		acct = tester.NewM365Account(t)
		// need to initialize the repository before we can test connecting to it.
		st = tester.NewPrefixedS3Storage(t)
		k  = kopia.NewConn(st)
	)

	err := k.Initialize(ctx, repository.Options{})
	require.NoError(t, err, clues.ToCore(err))

	// kopiaRef comes with a count of 1 and Wrapper bumps it again so safe
	// to close here.
	closer := func() { k.Close(ctx) }

	kw, err := kopia.NewWrapper(k)
	if !assert.NoError(t, err, clues.ToCore(err)) {
		closer()
		t.FailNow()
	}

	closer = func() {
		k.Close(ctx)
		kw.Close(ctx)
	}

	ms, err := kopia.NewModelStore(k)
	if !assert.NoError(t, err, clues.ToCore(err)) {
		closer()
		t.FailNow()
	}

	closer = func() {
		k.Close(ctx)
		kw.Close(ctx)
		ms.Close(ctx)
	}

	connectorResource := connector.Users
	if sel.Service == selectors.ServiceSharePoint {
		connectorResource = connector.Sites
	}

	gc, sel := GCWithSelector(t, ctx, acct, connectorResource, sel, nil, closer)
	bo := newTestBackupOp(t, ctx, kw, ms, gc, acct, sel, bus, featureToggles, closer)

	ss := streamstore.NewStreamer(kw, acct.ID(), sel.PathService())

	return bo, acct, kw, ms, ss, gc, sel, closer
}

// newTestBackupOp accepts the clients required to compose a backup operation, plus
// any other metadata, and uses them to generate a new backup operation.  This
// allows backup chains to utilize the same temp directory and configuration
// details.
func newTestBackupOp(
	t *testing.T,
	ctx context.Context, //revive:disable-line:context-as-argument
	kw *kopia.Wrapper,
	ms *kopia.ModelStore,
	gc *connector.GraphConnector,
	acct account.Account,
	sel selectors.Selector,
	bus events.Eventer,
	featureToggles control.Toggles,
	closer func(),
) BackupOperation {
	var (
		sw   = store.NewKopiaStore(ms)
		opts = control.Defaults()
	)

	opts.ToggleFeatures = featureToggles
	gc.IDNameLookup = idname.NewCache(map[string]string{sel.ID(): sel.Name()})

	bo, err := NewBackupOperation(ctx, opts, kw, sw, gc, acct, sel, sel, bus)
	if !assert.NoError(t, err, clues.ToCore(err)) {
		closer()
		t.FailNow()
	}

	return bo
}

func runAndCheckBackup(
	t *testing.T,
	ctx context.Context, //revive:disable-line:context-as-argument
	bo *BackupOperation,
	mb *evmock.Bus,
	acceptNoData bool,
) {
	err := bo.Run(ctx)
	require.NoError(t, err, clues.ToCore(err))
	require.NotEmpty(t, bo.Results, "the backup had non-zero results")
	require.NotEmpty(t, bo.Results.BackupID, "the backup generated an ID")

	expectStatus := []opStatus{Completed}
	if acceptNoData {
		expectStatus = append(expectStatus, NoData)
	}

	require.Contains(
		t,
		expectStatus,
		bo.Status,
		"backup doesn't match expectation, wanted any of %v, got %s",
		expectStatus,
		bo.Status)

	require.Less(t, 0, bo.Results.ItemsWritten)
	assert.Less(t, 0, bo.Results.ItemsRead, "count of items read")
	assert.Less(t, int64(0), bo.Results.BytesRead, "bytes read")
	assert.Less(t, int64(0), bo.Results.BytesUploaded, "bytes uploaded")
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
	bo *BackupOperation,
	sel selectors.Selector,
	resourceOwner string,
	categories ...path.CategoryType,
) {
	for _, category := range categories {
		t.Run(category.String(), func(t *testing.T) {
			var (
				reasons = []kopia.Reason{
					{
						ResourceOwner: resourceOwner,
						Service:       sel.PathService(),
						Category:      category,
					},
				}
				tags  = map[string]string{kopia.TagBackupCategory: ""}
				found bool
			)

			mans, err := kw.FetchPrevSnapshotManifests(ctx, reasons, tags)
			require.NoError(t, err, clues.ToCore(err))

			for _, man := range mans {
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
					assert.Implements(t, (*data.StreamSize)(nil), item)

					s := item.(data.StreamSize)
					assert.Greaterf(
						t,
						s.Size(),
						int64(0),
						"empty metadata file: %s/%s",
						col.FullPath(),
						item.UUID(),
					)

					itemNames = append(itemNames, item.UUID())
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
	gc *connector.GraphConnector,
	service path.ServiceType,
	acct account.Account,
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
		pathFolders = []string{"drives", driveID, "root:", destFldr}
	}

	collections := []incrementalCollection{{
		pathFolders: pathFolders,
		category:    cat,
		items:       items,
	}}

	dest := control.DefaultRestoreDestination(dttm.SafeForTesting)
	dest.ContainerName = destFldr

	dataColls := buildCollections(
		t,
		service,
		tenantID, resourceOwner,
		dest,
		collections)

	deets, err := gc.ConsumeRestoreCollections(
		ctx,
		backupVersion,
		acct,
		sel,
		dest,
		control.Options{RestorePermissions: true},
		dataColls,
		fault.New(true))
	require.NoError(t, err, clues.ToCore(err))

	// have to wait here, both to ensure the process
	// finishes, and also to clean up the gc status
	gc.Wait()

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
	dest control.RestoreDestination,
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

		collections = append(collections, data.NotFoundRestoreCollection{Collection: mc})
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
	default:
		err = clues.New(fmt.Sprintf("unknown service: %s", service))
	}

	require.NoError(t, err, clues.ToCore(err))

	return p
}

// ---------------------------------------------------------------------------
// integration tests
// ---------------------------------------------------------------------------

type BackupOpIntegrationSuite struct {
	tester.Suite
	user, site string
}

func TestBackupOpIntegrationSuite(t *testing.T) {
	suite.Run(t, &BackupOpIntegrationSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tester.AWSStorageCredEnvs, tester.M365AcctCredEnvs}),
	})
}

func (suite *BackupOpIntegrationSuite) SetupSuite() {
	suite.user = tester.M365UserID(suite.T())
	suite.site = tester.M365SiteID(suite.T())
}

func (suite *BackupOpIntegrationSuite) TestNewBackupOperation() {
	var (
		kw   = &kopia.Wrapper{}
		sw   = &store.Wrapper{}
		gc   = &mock.GraphConnector{}
		acct = tester.NewM365Account(suite.T())
		opts = control.Defaults()
	)

	table := []struct {
		name     string
		kw       *kopia.Wrapper
		sw       *store.Wrapper
		bp       inject.BackupProducer
		acct     account.Account
		targets  []string
		errCheck assert.ErrorAssertionFunc
	}{
		{"good", kw, sw, gc, acct, nil, assert.NoError},
		{"missing kopia", nil, sw, gc, acct, nil, assert.Error},
		{"missing modelstore", kw, nil, gc, acct, nil, assert.Error},
		{"missing backup producer", kw, sw, nil, acct, nil, assert.Error},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			ctx, flush := tester.NewContext()
			defer flush()

			sel := selectors.Selector{DiscreteOwner: "test"}

			_, err := NewBackupOperation(
				ctx,
				opts,
				test.kw,
				test.sw,
				test.bp,
				test.acct,
				sel,
				sel,
				evmock.NewBus())
			test.errCheck(suite.T(), err, clues.ToCore(err))
		})
	}
}

// ---------------------------------------------------------------------------
// Exchange
// ---------------------------------------------------------------------------

// TestBackup_Run ensures that Integration Testing works
// for the following scopes: Contacts, Events, and Mail
func (suite *BackupOpIntegrationSuite) TestBackup_Run_exchange() {
	ctx, flush := tester.NewContext()
	defer flush()

	tests := []struct {
		name          string
		selector      func() *selectors.ExchangeBackup
		category      path.CategoryType
		metadataFiles []string
	}{
		{
			name: "Mail",
			selector: func() *selectors.ExchangeBackup {
				sel := selectors.NewExchangeBackup([]string{suite.user})
				sel.Include(sel.MailFolders([]string{exchange.DefaultMailFolder}, selectors.PrefixMatch()))
				sel.DiscreteOwner = suite.user

				return sel
			},
			category:      path.EmailCategory,
			metadataFiles: exchange.MetadataFileNames(path.EmailCategory),
		},
		{
			name: "Contacts",
			selector: func() *selectors.ExchangeBackup {
				sel := selectors.NewExchangeBackup([]string{suite.user})
				sel.Include(sel.ContactFolders([]string{exchange.DefaultContactFolder}, selectors.PrefixMatch()))
				return sel
			},
			category:      path.ContactsCategory,
			metadataFiles: exchange.MetadataFileNames(path.ContactsCategory),
		},
		{
			name: "Calendar Events",
			selector: func() *selectors.ExchangeBackup {
				sel := selectors.NewExchangeBackup([]string{suite.user})
				sel.Include(sel.EventCalendars([]string{exchange.DefaultCalendar}, selectors.PrefixMatch()))
				return sel
			},
			category:      path.EventsCategory,
			metadataFiles: exchange.MetadataFileNames(path.EventsCategory),
		},
	}
	for _, test := range tests {
		suite.Run(test.name, func() {
			var (
				t       = suite.T()
				mb      = evmock.NewBus()
				sel     = test.selector().Selector
				ffs     = control.Toggles{}
				whatSet = deeTD.CategoryFromRepoRef
			)

			bo, acct, kw, ms, ss, gc, sel, closer := prepNewTestBackupOp(t, ctx, mb, sel, ffs, version.Backup)
			defer closer()

			userID := sel.ID()

			m365, err := acct.M365Config()
			require.NoError(t, err, clues.ToCore(err))

			// run the tests
			runAndCheckBackup(t, ctx, &bo, mb, false)
			checkBackupIsInManifests(t, ctx, kw, &bo, sel, userID, test.category)
			checkMetadataFilesExist(
				t,
				ctx,
				bo.Results.BackupID,
				kw,
				ms,
				m365.AzureTenantID,
				userID,
				path.ExchangeService,
				map[path.CategoryType][]string{test.category: test.metadataFiles})

			_, expectDeets := deeTD.GetDeetsInBackup(
				t,
				ctx,
				bo.Results.BackupID,
				acct.ID(),
				userID,
				path.ExchangeService,
				whatSet,
				ms,
				ss)
			deeTD.CheckBackupDetails(t, ctx, bo.Results.BackupID, whatSet, ms, ss, expectDeets, false)

			// Basic, happy path incremental test.  No changes are dictated or expected.
			// This only tests that an incremental backup is runnable at all, and that it
			// produces fewer results than the last backup.
			var (
				incMB = evmock.NewBus()
				incBO = newTestBackupOp(t, ctx, kw, ms, gc, acct, sel, incMB, ffs, closer)
			)

			runAndCheckBackup(t, ctx, &incBO, incMB, true)
			checkBackupIsInManifests(t, ctx, kw, &incBO, sel, userID, test.category)
			checkMetadataFilesExist(
				t,
				ctx,
				incBO.Results.BackupID,
				kw,
				ms,
				m365.AzureTenantID,
				userID,
				path.ExchangeService,
				map[path.CategoryType][]string{test.category: test.metadataFiles})
			deeTD.CheckBackupDetails(
				t,
				ctx,
				incBO.Results.BackupID,
				whatSet,
				ms,
				ss,
				expectDeets,
				false)

			// do some additional checks to ensure the incremental dealt with fewer items.
			assert.Greater(t, bo.Results.ItemsWritten, incBO.Results.ItemsWritten, "incremental items written")
			assert.Greater(t, bo.Results.ItemsRead, incBO.Results.ItemsRead, "incremental items read")
			assert.Greater(t, bo.Results.BytesRead, incBO.Results.BytesRead, "incremental bytes read")
			assert.Greater(t, bo.Results.BytesUploaded, incBO.Results.BytesUploaded, "incremental bytes uploaded")
			assert.Equal(t, bo.Results.ResourceOwners, incBO.Results.ResourceOwners, "incremental backup resource owner")
			assert.NoError(t, incBO.Errors.Failure(), "incremental non-recoverable error", clues.ToCore(bo.Errors.Failure()))
			assert.Empty(t, incBO.Errors.Recovered(), "count incremental recoverable/iteration errors")
			assert.Equal(t, 1, incMB.TimesCalled[events.BackupStart], "incremental backup-start events")
			assert.Equal(t, 1, incMB.TimesCalled[events.BackupEnd], "incremental backup-end events")
			assert.Equal(t,
				incMB.CalledWith[events.BackupStart][0][events.BackupID],
				incBO.Results.BackupID, "incremental backupID pre-declaration")
		})
	}
}

// TestBackup_Run ensures that Integration Testing works
// for the following scopes: Contacts, Events, and Mail
func (suite *BackupOpIntegrationSuite) TestBackup_Run_incrementalExchange() {
	ctx, flush := tester.NewContext()
	defer flush()

	tester.LogTimeOfTest(suite.T())

	var (
		t          = suite.T()
		acct       = tester.NewM365Account(t)
		ffs        = control.Toggles{}
		mb         = evmock.NewBus()
		now        = dttm.Now()
		service    = path.ExchangeService
		categories = map[path.CategoryType][]string{
			path.EmailCategory:    exchange.MetadataFileNames(path.EmailCategory),
			path.ContactsCategory: exchange.MetadataFileNames(path.ContactsCategory),
			// TODO: not currently functioning; cannot retrieve generated calendars
			// path.EventsCategory:   exchange.MetadataFileNames(path.EventsCategory),
		}
		container1      = fmt.Sprintf("%s%d_%s", incrementalsDestContainerPrefix, 1, now)
		container2      = fmt.Sprintf("%s%d_%s", incrementalsDestContainerPrefix, 2, now)
		container3      = fmt.Sprintf("%s%d_%s", incrementalsDestContainerPrefix, 3, now)
		containerRename = fmt.Sprintf("%s%d_%s", incrementalsDestContainerPrefix, 4, now)

		// container3 and containerRename don't exist yet.  Those will get created
		// later on during the tests.  Putting their identifiers into the selector
		// at this point is harmless.
		containers = []string{container1, container2, container3, containerRename}
		sel        = selectors.NewExchangeBackup([]string{suite.user})
		whatSet    = deeTD.CategoryFromRepoRef
	)

	gc, sels := GCWithSelector(t, ctx, acct, connector.Users, sel.Selector, nil, nil)
	sel.DiscreteOwner = sels.ID()
	sel.DiscreteOwnerName = sels.Name()

	uidn := inMock.NewProvider(sels.ID(), sels.Name())

	sel.Include(
		sel.MailFolders(containers, selectors.PrefixMatch()),
		sel.ContactFolders(containers, selectors.PrefixMatch()))

	m365, err := acct.M365Config()
	require.NoError(t, err, clues.ToCore(err))

	ac, err := exapi.NewClient(m365)
	require.NoError(t, err, clues.ToCore(err))

	// generate 3 new folders with two items each.
	// Only the first two folders will be part of the initial backup and
	// incrementals.  The third folder will be introduced partway through
	// the changes.
	// This should be enough to cover most delta actions, since moving one
	// container into another generates a delta for both addition and deletion.
	type contDeets struct {
		containerID string
		locRef      string
		itemRefs    []string // cached for populating expected deets, otherwise not used
	}

	mailDBF := func(id, timeStamp, subject, body string) []byte {
		return exchMock.MessageWith(
			suite.user, suite.user, suite.user,
			subject, body, body,
			now, now, now, now)
	}

	contactDBF := func(id, timeStamp, subject, body string) []byte {
		given, mid, sur := id[:8], id[9:13], id[len(id)-12:]

		return exchMock.ContactBytesWith(
			given+" "+sur,
			sur+", "+given,
			given, mid, sur,
			"123-456-7890")
	}

	eventDBF := func(id, timeStamp, subject, body string) []byte {
		return exchMock.EventWith(
			suite.user, subject, body, body,
			now, now, exchMock.NoRecurrence, exchMock.NoAttendees, false)
	}

	// test data set
	dataset := map[path.CategoryType]struct {
		dbf   dataBuilderFunc
		dests map[string]contDeets
	}{
		path.EmailCategory: {
			dbf: mailDBF,
			dests: map[string]contDeets{
				container1: {},
				container2: {},
			},
		},
		path.ContactsCategory: {
			dbf: contactDBF,
			dests: map[string]contDeets{
				container1: {},
				container2: {},
			},
		},
		// TODO: not currently functioning; cannot retrieve generated calendars
		// path.EventsCategory: {
		// 	dbf: eventDBF,
		// 	dests: map[string]contDeets{
		// 		container1: {},
		// 		container2: {},
		// 	},
		// },
	}

	// populate initial test data
	for category, gen := range dataset {
		for destName := range gen.dests {
			// TODO: the details.Builder returned by restore can contain entries with
			// incorrect information.  non-representative repo-refs and the like.  Until
			// that gets fixed, we can't consume that info for testing.
			deets := generateContainerOfItems(
				t,
				ctx,
				gc,
				service,
				acct,
				category,
				selectors.NewExchangeRestore([]string{uidn.ID()}).Selector,
				m365.AzureTenantID, uidn.ID(), "", destName,
				2,
				version.Backup,
				gen.dbf)

			itemRefs := []string{}

			for _, ent := range deets.Entries {
				if ent.Exchange == nil || ent.Folder != nil {
					continue
				}

				if len(ent.ItemRef) > 0 {
					itemRefs = append(itemRefs, ent.ItemRef)
				}
			}

			// save the item ids for building expectedDeets later on
			cd := dataset[category].dests[destName]
			cd.itemRefs = itemRefs
			dataset[category].dests[destName] = cd
		}
	}

	bo, acct, kw, ms, ss, gc, sels, closer := prepNewTestBackupOp(t, ctx, mb, sel.Selector, ffs, version.Backup)
	defer closer()

	// run the initial backup
	runAndCheckBackup(t, ctx, &bo, mb, false)

	rrPfx, err := path.ServicePrefix(acct.ID(), uidn.ID(), service, path.EmailCategory)
	require.NoError(t, err, clues.ToCore(err))

	// strip the category from the prefix; we primarily want the tenant and resource owner.
	expectDeets := deeTD.NewInDeets(rrPfx.ToBuilder().Dir().String())
	bupDeets, _ := deeTD.GetDeetsInBackup(t, ctx, bo.Results.BackupID, acct.ID(), uidn.ID(), service, whatSet, ms, ss)

	// update the datasets with their location refs
	for category, gen := range dataset {
		for destName, cd := range gen.dests {
			var longestLR string

			for _, ent := range bupDeets.Entries {
				// generated destinations should always contain items
				if ent.Folder != nil {
					continue
				}

				p, err := path.FromDataLayerPath(ent.RepoRef, false)
				require.NoError(t, err, clues.ToCore(err))

				// category must match, and the owning folder must be this destination
				if p.Category() != category || strings.HasSuffix(ent.LocationRef, destName) {
					continue
				}

				// emails, due to folder nesting and our design for populating data via restore,
				// will duplicate the dest folder as both the restore destination, and the "old parent
				// folder".  we'll get both a prefix/destName and a prefix/destName/destName folder.
				// since we want future comparison to only use the leaf dir, we select for the longest match.
				if len(ent.LocationRef) > len(longestLR) {
					longestLR = ent.LocationRef
				}
			}

			require.NotEmptyf(t, longestLR, "must find an expected details entry matching the generated folder: %s", destName)

			cd.locRef = longestLR

			dataset[category].dests[destName] = cd
			expectDeets.AddLocation(category.String(), cd.locRef)

			for _, i := range dataset[category].dests[destName].itemRefs {
				expectDeets.AddItem(category.String(), cd.locRef, i)
			}
		}
	}

	// verify test data was populated, and track it for comparisons
	// TODO: this can be swapped out for InDeets checks if we add itemRefs to folder ents.
	for category, gen := range dataset {
		qp := graph.QueryParams{
			Category:      category,
			ResourceOwner: uidn,
			Credentials:   m365,
		}

		cr, err := exchange.PopulateExchangeContainerResolver(ctx, qp, fault.New(true))
		require.NoError(t, err, "populating container resolver", category, clues.ToCore(err))

		for destName, dest := range gen.dests {
			id, ok := cr.LocationInCache(dest.locRef)
			require.True(t, ok, "dir %s found in %s cache", dest.locRef, category)

			dest.containerID = id
			dataset[category].dests[destName] = dest
		}
	}

	// precheck to ensure the expectedDeets are correct.
	// if we fail here, the expectedDeets were populated incorrectly.
	deeTD.CheckBackupDetails(t, ctx, bo.Results.BackupID, whatSet, ms, ss, expectDeets, true)

	// Although established as a table, these tests are no isolated from each other.
	// Assume that every test's side effects cascade to all following test cases.
	// The changes are split across the table so that we can monitor the deltas
	// in isolation, rather than debugging one change from the rest of a series.
	table := []struct {
		name string
		// performs the incremental update required for the test.
		updateUserData func(t *testing.T)
		itemsRead      int
		itemsWritten   int
	}{
		{
			name:           "clean incremental, no changes",
			updateUserData: func(t *testing.T) {},
			itemsRead:      0,
			itemsWritten:   0,
		},
		{
			name: "move an email folder to a subfolder",
			updateUserData: func(t *testing.T) {
				cat := path.EmailCategory

				// contacts and events cannot be sufoldered; this is an email-only change
				from := dataset[cat].dests[container2]
				to := dataset[cat].dests[container1]

				body := users.NewItemMailFoldersItemMovePostRequestBody()
				body.SetDestinationId(ptr.To(to.containerID))

				_, err := gc.Service.
					Client().
					UsersById(uidn.ID()).
					MailFoldersById(from.containerID).
					Move().
					Post(ctx, body, nil)
				require.NoError(t, err, clues.ToCore(err))

				newLoc := expectDeets.MoveLocation(cat.String(), from.locRef, to.locRef)
				from.locRef = newLoc
			},
			itemsRead:    0, // zero because we don't count container reads
			itemsWritten: 2,
		},
		{
			name: "delete a folder",
			updateUserData: func(t *testing.T) {
				for category, d := range dataset {
					containerID := d.dests[container2].containerID

					switch category {
					case path.EmailCategory:
						err := ac.Mail().DeleteContainer(ctx, uidn.ID(), containerID)
						require.NoError(t, err, "deleting an email folder", clues.ToCore(err))
					case path.ContactsCategory:
						err := ac.Contacts().DeleteContainer(ctx, uidn.ID(), containerID)
						require.NoError(t, err, "deleting a contacts folder", clues.ToCore(err))
					case path.EventsCategory:
						err := ac.Events().DeleteContainer(ctx, uidn.ID(), containerID)
						require.NoError(t, err, "deleting a calendar", clues.ToCore(err))
					}

					expectDeets.RemoveLocation(category.String(), d.dests[container2].locRef)
				}
			},
			itemsRead:    0,
			itemsWritten: 0, // deletions are not counted as "writes"
		},
		{
			name: "add a new folder",
			updateUserData: func(t *testing.T) {
				for category, gen := range dataset {
					deets := generateContainerOfItems(
						t,
						ctx,
						gc,
						service,
						acct,
						category,
						selectors.NewExchangeRestore([]string{uidn.ID()}).Selector,
						m365.AzureTenantID, suite.user, "", container3,
						2,
						version.Backup,
						gen.dbf)

					qp := graph.QueryParams{
						Category:      category,
						ResourceOwner: uidn,
						Credentials:   m365,
					}

					expectedLocRef := container3
					if category == path.EmailCategory {
						expectedLocRef = path.Builder{}.Append(container3, container3).String()
					}

					cr, err := exchange.PopulateExchangeContainerResolver(ctx, qp, fault.New(true))
					require.NoError(t, err, "populating container resolver", category, clues.ToCore(err))

					id, ok := cr.LocationInCache(expectedLocRef)
					require.Truef(t, ok, "dir %s found in %s cache", expectedLocRef, category)

					dataset[category].dests[container3] = contDeets{
						containerID: id,
						locRef:      expectedLocRef,
						itemRefs:    nil, // not needed at this point
					}

					for _, ent := range deets.Entries {
						if ent.Folder == nil {
							expectDeets.AddItem(category.String(), expectedLocRef, ent.ItemRef)
						}
					}
				}
			},
			itemsRead:    4,
			itemsWritten: 4,
		},
		{
			name: "rename a folder",
			updateUserData: func(t *testing.T) {
				for category, d := range dataset {
					cli := gc.Service.Client().UsersById(uidn.ID())
					containerID := d.dests[container3].containerID
					newLoc := containerRename

					if category == path.EmailCategory {
						newLoc = path.Builder{}.Append(container3, containerRename).String()
					}

					d.dests[containerRename] = contDeets{
						containerID: containerID,
						locRef:      newLoc,
					}

					expectDeets.RenameLocation(
						category.String(),
						d.dests[container3].containerID,
						newLoc)

					switch category {
					case path.EmailCategory:
						cmf := cli.MailFoldersById(containerID)

						body, err := cmf.Get(ctx, nil)
						require.NoError(t, err, "getting mail folder", clues.ToCore(err))

						body.SetDisplayName(&containerRename)
						_, err = cmf.Patch(ctx, body, nil)
						require.NoError(t, err, "updating mail folder name", clues.ToCore(err))

					case path.ContactsCategory:
						ccf := cli.ContactFoldersById(containerID)

						body, err := ccf.Get(ctx, nil)
						require.NoError(t, err, "getting contact folder", clues.ToCore(err))

						body.SetDisplayName(&containerRename)
						_, err = ccf.Patch(ctx, body, nil)
						require.NoError(t, err, "updating contact folder name", clues.ToCore(err))

					case path.EventsCategory:
						cbi := cli.CalendarsById(containerID)

						body, err := cbi.Get(ctx, nil)
						require.NoError(t, err, "getting calendar", clues.ToCore(err))

						body.SetName(&containerRename)
						_, err = cbi.Patch(ctx, body, nil)
						require.NoError(t, err, "updating calendar name", clues.ToCore(err))
					}
				}
			},
			itemsRead:    0, // containers are not counted as reads
			itemsWritten: 4, // two items per category
		},
		{
			name: "add a new item",
			updateUserData: func(t *testing.T) {
				for category, d := range dataset {
					containerID := d.dests[container1].containerID
					cli := gc.Service.Client().UsersById(uidn.ID())

					switch category {
					case path.EmailCategory:
						_, itemData := generateItemData(t, category, uidn.ID(), mailDBF)
						body, err := support.CreateMessageFromBytes(itemData)
						require.NoError(t, err, "transforming mail bytes to messageable", clues.ToCore(err))

						itm, err := cli.MailFoldersById(containerID).Messages().Post(ctx, body, nil)
						require.NoError(t, err, "posting email item", clues.ToCore(err))

						expectDeets.AddItem(
							category.String(),
							d.dests[category.String()].locRef,
							ptr.Val(itm.GetId()))

					case path.ContactsCategory:
						_, itemData := generateItemData(t, category, uidn.ID(), contactDBF)
						body, err := support.CreateContactFromBytes(itemData)
						require.NoError(t, err, "transforming contact bytes to contactable", clues.ToCore(err))

						itm, err := cli.ContactFoldersById(containerID).Contacts().Post(ctx, body, nil)
						require.NoError(t, err, "posting contact item", clues.ToCore(err))

						expectDeets.AddItem(
							category.String(),
							d.dests[category.String()].locRef,
							ptr.Val(itm.GetId()))

					case path.EventsCategory:
						_, itemData := generateItemData(t, category, uidn.ID(), eventDBF)
						body, err := support.CreateEventFromBytes(itemData)
						require.NoError(t, err, "transforming event bytes to eventable", clues.ToCore(err))

						itm, err := cli.CalendarsById(containerID).Events().Post(ctx, body, nil)
						require.NoError(t, err, "posting events item", clues.ToCore(err))

						expectDeets.AddItem(
							category.String(),
							d.dests[category.String()].locRef,
							ptr.Val(itm.GetId()))
					}
				}
			},
			itemsRead:    2,
			itemsWritten: 2,
		},
		{
			name: "delete an existing item",
			updateUserData: func(t *testing.T) {
				for category, d := range dataset {
					containerID := d.dests[container1].containerID
					cli := gc.Service.Client().UsersById(uidn.ID())

					switch category {
					case path.EmailCategory:
						ids, _, _, err := ac.Mail().GetAddedAndRemovedItemIDs(ctx, uidn.ID(), containerID, "", false)
						require.NoError(t, err, "getting message ids", clues.ToCore(err))
						require.NotEmpty(t, ids, "message ids in folder")

						err = cli.MessagesById(ids[0]).Delete(ctx, nil)
						require.NoError(t, err, "deleting email item", clues.ToCore(err))

						expectDeets.RemoveItem(
							category.String(),
							d.dests[category.String()].locRef,
							ids[0])

					case path.ContactsCategory:
						ids, _, _, err := ac.Contacts().GetAddedAndRemovedItemIDs(ctx, uidn.ID(), containerID, "", false)
						require.NoError(t, err, "getting contact ids", clues.ToCore(err))
						require.NotEmpty(t, ids, "contact ids in folder")

						err = cli.ContactsById(ids[0]).Delete(ctx, nil)
						require.NoError(t, err, "deleting contact item", clues.ToCore(err))

						expectDeets.RemoveItem(
							category.String(),
							d.dests[category.String()].locRef,
							ids[0])

					case path.EventsCategory:
						ids, _, _, err := ac.Events().GetAddedAndRemovedItemIDs(ctx, uidn.ID(), containerID, "", false)
						require.NoError(t, err, "getting event ids", clues.ToCore(err))
						require.NotEmpty(t, ids, "event ids in folder")

						err = cli.CalendarsById(ids[0]).Delete(ctx, nil)
						require.NoError(t, err, "deleting calendar", clues.ToCore(err))

						expectDeets.RemoveItem(
							category.String(),
							d.dests[category.String()].locRef,
							ids[0])
					}
				}
			},
			itemsRead:    2,
			itemsWritten: 0, // deletes are not counted as "writes"
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			var (
				t     = suite.T()
				incMB = evmock.NewBus()
				incBO = newTestBackupOp(t, ctx, kw, ms, gc, acct, sels, incMB, ffs, closer)
				atid  = m365.AzureTenantID
			)

			test.updateUserData(t)

			err := incBO.Run(ctx)
			require.NoError(t, err, clues.ToCore(err))

			bupID := incBO.Results.BackupID

			checkBackupIsInManifests(t, ctx, kw, &incBO, sels, uidn.ID(), maps.Keys(categories)...)
			checkMetadataFilesExist(t, ctx, bupID, kw, ms, atid, uidn.ID(), service, categories)
			deeTD.CheckBackupDetails(t, ctx, bupID, whatSet, ms, ss, expectDeets, true)

			// do some additional checks to ensure the incremental dealt with fewer items.
			// +4 on read/writes to account for metadata: 1 delta and 1 path for each type.
			assert.Equal(t, test.itemsWritten+4, incBO.Results.ItemsWritten, "incremental items written")
			assert.Equal(t, test.itemsRead+4, incBO.Results.ItemsRead, "incremental items read")
			assert.NoError(t, incBO.Errors.Failure(), "incremental non-recoverable error", clues.ToCore(incBO.Errors.Failure()))
			assert.Empty(t, incBO.Errors.Recovered(), "incremental recoverable/iteration errors")
			assert.Equal(t, 1, incMB.TimesCalled[events.BackupStart], "incremental backup-start events")
			assert.Equal(t, 1, incMB.TimesCalled[events.BackupEnd], "incremental backup-end events")
			assert.Equal(t,
				incMB.CalledWith[events.BackupStart][0][events.BackupID],
				bupID, "incremental backupID pre-declaration")
		})
	}
}

// ---------------------------------------------------------------------------
// OneDrive
// ---------------------------------------------------------------------------

func (suite *BackupOpIntegrationSuite) TestBackup_Run_oneDrive() {
	ctx, flush := tester.NewContext()
	defer flush()

	var (
		t      = suite.T()
		tenID  = tester.M365TenantID(t)
		mb     = evmock.NewBus()
		userID = tester.SecondaryM365UserID(t)
		osel   = selectors.NewOneDriveBackup([]string{userID})
		ws     = deeTD.DriveIDFromRepoRef
		svc    = path.OneDriveService
	)

	osel.Include(selTD.OneDriveBackupFolderScope(osel))

	bo, _, _, ms, ss, _, sel, closer := prepNewTestBackupOp(t, ctx, mb, osel.Selector, control.Toggles{}, version.Backup)
	defer closer()

	runAndCheckBackup(t, ctx, &bo, mb, false)

	bID := bo.Results.BackupID

	_, expectDeets := deeTD.GetDeetsInBackup(t, ctx, bID, tenID, sel.ID(), svc, ws, ms, ss)
	deeTD.CheckBackupDetails(t, ctx, bID, ws, ms, ss, expectDeets, false)
}

func (suite *BackupOpIntegrationSuite) TestBackup_Run_incrementalOneDrive() {
	sel := selectors.NewOneDriveRestore([]string{suite.user})

	ic := func(cs []string) selectors.Selector {
		sel.Include(sel.Folders(cs, selectors.PrefixMatch()))
		return sel.Selector
	}

	gtdi := func(
		t *testing.T,
		ctx context.Context,
		gs graph.Servicer,
	) string {
		d, err := odapi.GetUsersDrive(ctx, gs, suite.user)
		if err != nil {
			err = graph.Wrap(ctx, err, "retrieving default user drive").
				With("user", suite.user)
		}

		require.NoError(t, err, clues.ToCore(err))

		id := ptr.Val(d.GetId())
		require.NotEmpty(t, id, "drive ID")

		return id
	}

	runDriveIncrementalTest(
		suite,
		suite.user,
		suite.user,
		connector.Users,
		path.OneDriveService,
		path.FilesCategory,
		ic,
		gtdi,
		false)
}

func (suite *BackupOpIntegrationSuite) TestBackup_Run_incrementalSharePoint() {
	sel := selectors.NewSharePointRestore([]string{suite.site})

	ic := func(cs []string) selectors.Selector {
		sel.Include(sel.LibraryFolders(cs, selectors.PrefixMatch()))
		return sel.Selector
	}

	gtdi := func(
		t *testing.T,
		ctx context.Context,
		gs graph.Servicer,
	) string {
		d, err := odapi.GetSitesDefaultDrive(ctx, gs, suite.site)
		if err != nil {
			err = graph.Wrap(ctx, err, "retrieving default site drive").
				With("site", suite.site)
		}

		require.NoError(t, err, clues.ToCore(err))

		id := ptr.Val(d.GetId())
		require.NotEmpty(t, id, "drive ID")

		return id
	}

	runDriveIncrementalTest(
		suite,
		suite.site,
		suite.user,
		connector.Sites,
		path.SharePointService,
		path.LibrariesCategory,
		ic,
		gtdi,
		true)
}

// Good reference point for testing incremental backups
func runDriveIncrementalTest(
	suite *BackupOpIntegrationSuite,
	owner, permissionsUser string,
	resource connector.Resource,
	service path.ServiceType,
	category path.CategoryType,
	includeContainers func([]string) selectors.Selector,
	getTestDriveID func(*testing.T, context.Context, graph.Servicer) string,
	skipPermissionsTests bool,
) {
	ctx, flush := tester.NewContext()
	defer flush()

	var (
		t    = suite.T()
		acct = tester.NewM365Account(t)
		ffs  = control.Toggles{}
		mb   = evmock.NewBus()
		ws   = deeTD.DriveIDFromRepoRef

		// `now` has to be formatted with SimpleDateTimeTesting as
		// some drives cannot have `:` in file/folder names
		now = dttm.FormatNow(dttm.SafeForTesting)

		categories = map[path.CategoryType][]string{
			category: {graph.DeltaURLsFileName, graph.PreviousPathFileName},
		}
		container1      = fmt.Sprintf("%s%d_%s", incrementalsDestContainerPrefix, 1, now)
		container2      = fmt.Sprintf("%s%d_%s", incrementalsDestContainerPrefix, 2, now)
		container3      = fmt.Sprintf("%s%d_%s", incrementalsDestContainerPrefix, 3, now)
		containerRename = "renamed_folder"

		genDests = []string{container1, container2}

		// container3 does not exist yet. It will get created later on
		// during the tests.
		containers = []string{container1, container2, container3}
	)

	sel := includeContainers(containers)

	creds, err := acct.M365Config()
	require.NoError(t, err, clues.ToCore(err))

	gc, sel := GCWithSelector(t, ctx, acct, resource, sel, nil, nil)

	roidn := inMock.NewProvider(sel.ID(), sel.Name())

	var (
		atid    = creds.AzureTenantID
		driveID = getTestDriveID(t, ctx, gc.Service)
		fileDBF = func(id, timeStamp, subject, body string) []byte {
			return []byte(id + subject)
		}
		makeLocRef = func(flds ...string) string {
			elems := append([]string{driveID, "root:"}, flds...)
			return path.Builder{}.Append(elems...).String()
		}
	)

	rrPfx, err := path.ServicePrefix(atid, roidn.ID(), service, category)
	require.NoError(t, err, clues.ToCore(err))

	// strip the category from the prefix; we primarily want the tenant and resource owner.
	expectDeets := deeTD.NewInDeets(rrPfx.ToBuilder().Dir().String())

	// Populate initial test data.
	// Generate 2 new folders with two items each. Only the first two
	// folders will be part of the initial backup and
	// incrementals. The third folder will be introduced partway
	// through the changes. This should be enough to cover most delta
	// actions.
	for _, destName := range genDests {
		deets := generateContainerOfItems(
			t,
			ctx,
			gc,
			service,
			acct,
			category,
			sel,
			atid, roidn.ID(), driveID, destName,
			2,
			// Use an old backup version so we don't need metadata files.
			0,
			fileDBF)

		for _, ent := range deets.Entries {
			if ent.Folder != nil {
				continue
			}

			expectDeets.AddItem(driveID, makeLocRef(destName), ent.ItemRef)
		}
	}

	containerIDs := map[string]string{}

	// verify test data was populated, and track it for comparisons
	for _, destName := range genDests {
		// Use path-based indexing to get the folder's ID. This is sourced from the
		// onedrive package `getFolder` function.
		itemURL := fmt.Sprintf("https://graph.microsoft.com/v1.0/drives/%s/root:/%s", driveID, destName)
		resp, err := drive.
			NewItemsDriveItemItemRequestBuilder(itemURL, gc.Service.Adapter()).
			Get(ctx, nil)
		require.NoError(t, err, "getting drive folder ID", "folder name", destName, clues.ToCore(err))

		containerIDs[destName] = ptr.Val(resp.GetId())
	}

	bo, _, kw, ms, ss, gc, _, closer := prepNewTestBackupOp(t, ctx, mb, sel, ffs, version.Backup)
	defer closer()

	// run the initial backup
	runAndCheckBackup(t, ctx, &bo, mb, false)

	// precheck to ensure the expectedDeets are correct.
	// if we fail here, the expectedDeets were populated incorrectly.
	deeTD.CheckBackupDetails(t, ctx, bo.Results.BackupID, ws, ms, ss, expectDeets, true)

	var (
		newFile     models.DriveItemable
		newFileName = "new_file.txt"
		newFileID   string

		permissionIDMappings = map[string]string{}
		writePerm            = metadata.Permission{
			ID:       "perm-id",
			Roles:    []string{"write"},
			EntityID: permissionsUser,
		}
	)

	// Although established as a table, these tests are not isolated from each other.
	// Assume that every test's side effects cascade to all following test cases.
	// The changes are split across the table so that we can monitor the deltas
	// in isolation, rather than debugging one change from the rest of a series.
	table := []struct {
		name string
		// performs the incremental update required for the test.
		updateFiles  func(t *testing.T)
		itemsRead    int
		itemsWritten int
		skip         bool
	}{
		{
			name:         "clean incremental, no changes",
			updateFiles:  func(t *testing.T) {},
			itemsRead:    0,
			itemsWritten: 0,
		},
		{
			name: "create a new file",
			updateFiles: func(t *testing.T) {
				targetContainer := containerIDs[container1]
				driveItem := models.NewDriveItem()
				driveItem.SetName(&newFileName)
				driveItem.SetFile(models.NewFile())
				newFile, err = onedrive.CreateItem(
					ctx,
					gc.Service,
					driveID,
					targetContainer,
					driveItem)
				require.NoErrorf(t, err, "creating new file %v", clues.ToCore(err))

				newFileID = ptr.Val(newFile.GetId())

				expectDeets.AddItem(driveID, makeLocRef(container1), newFileID)
			},
			itemsRead:    1, // .data file for newitem
			itemsWritten: 3, // .data and .meta for newitem, .dirmeta for parent
		},
		{
			name: "add permission to new file",
			skip: skipPermissionsTests,
			updateFiles: func(t *testing.T) {
				driveItem := models.NewDriveItem()
				driveItem.SetName(&newFileName)
				driveItem.SetFile(models.NewFile())
				err = onedrive.UpdatePermissions(
					ctx,
					creds,
					gc.Service,
					driveID,
					*newFile.GetId(),
					[]metadata.Permission{writePerm},
					[]metadata.Permission{},
					permissionIDMappings,
				)
				require.NoErrorf(t, err, "adding permission to file %v", clues.ToCore(err))
				// no expectedDeets: metadata isn't tracked
			},
			itemsRead:    1, // .data file for newitem
			itemsWritten: 2, // .meta for newitem, .dirmeta for parent (.data is not written as it is not updated)
		},
		{
			name: "remove permission from new file",
			skip: skipPermissionsTests,
			updateFiles: func(t *testing.T) {
				driveItem := models.NewDriveItem()
				driveItem.SetName(&newFileName)
				driveItem.SetFile(models.NewFile())
				err = onedrive.UpdatePermissions(
					ctx,
					creds,
					gc.Service,
					driveID,
					*newFile.GetId(),
					[]metadata.Permission{},
					[]metadata.Permission{writePerm},
					permissionIDMappings,
				)
				require.NoErrorf(t, err, "adding permission to file %v", clues.ToCore(err))
				// no expectedDeets: metadata isn't tracked
			},
			itemsRead:    1, // .data file for newitem
			itemsWritten: 2, // .meta for newitem, .dirmeta for parent (.data is not written as it is not updated)
		},
		{
			name: "add permission to container",
			skip: skipPermissionsTests,
			updateFiles: func(t *testing.T) {
				targetContainer := containerIDs[container1]
				driveItem := models.NewDriveItem()
				driveItem.SetName(&newFileName)
				driveItem.SetFile(models.NewFile())
				err = onedrive.UpdatePermissions(
					ctx,
					creds,
					gc.Service,
					driveID,
					targetContainer,
					[]metadata.Permission{writePerm},
					[]metadata.Permission{},
					permissionIDMappings,
				)
				require.NoErrorf(t, err, "adding permission to file %v", clues.ToCore(err))
				// no expectedDeets: metadata isn't tracked5tgb
			},
			itemsRead:    0,
			itemsWritten: 1, // .dirmeta for collection
		},
		{
			name: "remove permission from container",
			skip: skipPermissionsTests,
			updateFiles: func(t *testing.T) {
				targetContainer := containerIDs[container1]
				driveItem := models.NewDriveItem()
				driveItem.SetName(&newFileName)
				driveItem.SetFile(models.NewFile())
				err = onedrive.UpdatePermissions(
					ctx,
					creds,
					gc.Service,
					driveID,
					targetContainer,
					[]metadata.Permission{},
					[]metadata.Permission{writePerm},
					permissionIDMappings,
				)
				require.NoErrorf(t, err, "adding permission to file %v", clues.ToCore(err))
				// no expectedDeets: metadata isn't tracked
			},
			itemsRead:    0,
			itemsWritten: 1, // .dirmeta for collection
		},
		{
			name: "update contents of a file",
			updateFiles: func(t *testing.T) {
				err := gc.Service.
					Client().
					DrivesById(driveID).
					ItemsById(ptr.Val(newFile.GetId())).
					Content().
					Put(ctx, []byte("new content"), nil)
				require.NoErrorf(t, err, "updating file contents: %v", clues.ToCore(err))
				// no expectedDeets: neither file id nor location changed
			},
			itemsRead:    1, // .data file for newitem
			itemsWritten: 3, // .data and .meta for newitem, .dirmeta for parent
		},
		{
			name: "rename a file",
			updateFiles: func(t *testing.T) {
				container := containerIDs[container1]

				driveItem := models.NewDriveItem()
				name := "renamed_new_file.txt"
				driveItem.SetName(&name)
				parentRef := models.NewItemReference()
				parentRef.SetId(&container)
				driveItem.SetParentReference(parentRef)

				_, err := gc.Service.
					Client().
					DrivesById(driveID).
					ItemsById(ptr.Val(newFile.GetId())).
					Patch(ctx, driveItem, nil)
				require.NoError(t, err, "renaming file %v", clues.ToCore(err))
			},
			itemsRead:    1, // .data file for newitem
			itemsWritten: 3, // .data and .meta for newitem, .dirmeta for parent
			// no expectedDeets: neither file id nor location changed
		},
		{
			name: "move a file between folders",
			updateFiles: func(t *testing.T) {
				dest := containerIDs[container2]

				driveItem := models.NewDriveItem()
				driveItem.SetName(&newFileName)
				parentRef := models.NewItemReference()
				parentRef.SetId(&dest)
				driveItem.SetParentReference(parentRef)

				_, err := gc.Service.
					Client().
					DrivesById(driveID).
					ItemsById(ptr.Val(newFile.GetId())).
					Patch(ctx, driveItem, nil)
				require.NoErrorf(t, err, "moving file between folders %v", clues.ToCore(err))

				expectDeets.MoveItem(
					driveID,
					makeLocRef(container1),
					makeLocRef(container2),
					ptr.Val(newFile.GetId()))
			},
			itemsRead:    1, // .data file for newitem
			itemsWritten: 3, // .data and .meta for newitem, .dirmeta for parent
		},
		{
			name: "delete file",
			updateFiles: func(t *testing.T) {
				// deletes require unique http clients
				// https://github.com/alcionai/corso/issues/2707
				err = newDeleteServicer(t).
					Client().
					DrivesById(driveID).
					ItemsById(ptr.Val(newFile.GetId())).
					Delete(ctx, nil)
				require.NoErrorf(t, err, "deleting file %v", clues.ToCore(err))

				expectDeets.RemoveItem(driveID, makeLocRef(container2), ptr.Val(newFile.GetId()))
			},
			itemsRead:    0,
			itemsWritten: 0,
		},
		{
			name: "move a folder to a subfolder",
			updateFiles: func(t *testing.T) {
				parent := containerIDs[container1]
				child := containerIDs[container2]

				driveItem := models.NewDriveItem()
				driveItem.SetName(&container2)
				parentRef := models.NewItemReference()
				parentRef.SetId(&parent)
				driveItem.SetParentReference(parentRef)

				_, err := gc.Service.
					Client().
					DrivesById(driveID).
					ItemsById(child).
					Patch(ctx, driveItem, nil)
				require.NoError(t, err, "moving folder", clues.ToCore(err))

				expectDeets.MoveLocation(
					driveID,
					makeLocRef(container2),
					makeLocRef(container1))
			},
			itemsRead:    0,
			itemsWritten: 7, // 2*2(data and meta of 2 files) + 3 (dirmeta of two moved folders and target)
		},
		{
			name: "rename a folder",
			updateFiles: func(t *testing.T) {
				parent := containerIDs[container1]
				child := containerIDs[container2]

				driveItem := models.NewDriveItem()
				driveItem.SetName(&containerRename)
				parentRef := models.NewItemReference()
				parentRef.SetId(&parent)
				driveItem.SetParentReference(parentRef)

				_, err := gc.Service.
					Client().
					DrivesById(driveID).
					ItemsById(child).
					Patch(ctx, driveItem, nil)
				require.NoError(t, err, "renaming folder", clues.ToCore(err))

				containerIDs[containerRename] = containerIDs[container2]

				expectDeets.RenameLocation(
					driveID,
					makeLocRef(container1, container2),
					makeLocRef(container1, containerRename))
			},
			itemsRead:    0,
			itemsWritten: 7, // 2*2(data and meta of 2 files) + 3 (dirmeta of two moved folders and target)
		},
		{
			name: "delete a folder",
			updateFiles: func(t *testing.T) {
				container := containerIDs[containerRename]
				// deletes require unique http clients
				// https://github.com/alcionai/corso/issues/2707
				err = newDeleteServicer(t).
					Client().
					DrivesById(driveID).
					ItemsById(container).
					Delete(ctx, nil)
				require.NoError(t, err, "deleting folder", clues.ToCore(err))

				expectDeets.RemoveLocation(driveID, makeLocRef(container1, containerRename))
			},
			itemsRead:    0,
			itemsWritten: 0,
		},
		{
			name: "add a new folder",
			updateFiles: func(t *testing.T) {
				generateContainerOfItems(
					t,
					ctx,
					gc,
					service,
					acct,
					category,
					sel,
					atid, roidn.ID(), driveID, container3,
					2,
					0,
					fileDBF)

				// Validate creation
				itemURL := fmt.Sprintf(
					"https://graph.microsoft.com/v1.0/drives/%s/root:/%s",
					driveID,
					container3)
				resp, err := drive.NewItemsDriveItemItemRequestBuilder(itemURL, gc.Service.Adapter()).
					Get(ctx, nil)
				require.NoError(t, err, "getting drive folder ID", "folder name", container3, clues.ToCore(err))

				containerIDs[container3] = ptr.Val(resp.GetId())

				expectDeets.AddLocation(driveID, container3)
			},
			itemsRead:    2, // 2 .data for 2 files
			itemsWritten: 6, // read items + 2 directory meta
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			if test.skip {
				suite.T().Skip("flagged to skip")
			}

			cleanGC, err := connector.NewGraphConnector(ctx, acct, resource)
			require.NoError(t, err, clues.ToCore(err))

			var (
				t     = suite.T()
				incMB = evmock.NewBus()
				incBO = newTestBackupOp(t, ctx, kw, ms, cleanGC, acct, sel, incMB, ffs, closer)
			)

			tester.LogTimeOfTest(suite.T())

			test.updateFiles(t)

			err = incBO.Run(ctx)
			require.NoError(t, err, clues.ToCore(err))

			bupID := incBO.Results.BackupID

			checkBackupIsInManifests(t, ctx, kw, &incBO, sel, roidn.ID(), maps.Keys(categories)...)
			checkMetadataFilesExist(t, ctx, bupID, kw, ms, atid, roidn.ID(), service, categories)
			deeTD.CheckBackupDetails(t, ctx, bupID, ws, ms, ss, expectDeets, true)

			// do some additional checks to ensure the incremental dealt with fewer items.
			// +2 on read/writes to account for metadata: 1 delta and 1 path.
			assert.Equal(t, test.itemsWritten+2, incBO.Results.ItemsWritten, "incremental items written")
			assert.Equal(t, test.itemsRead+2, incBO.Results.ItemsRead, "incremental items read")
			assert.NoError(t, incBO.Errors.Failure(), "incremental non-recoverable error", clues.ToCore(incBO.Errors.Failure()))
			assert.Empty(t, incBO.Errors.Recovered(), "incremental recoverable/iteration errors")
			assert.Equal(t, 1, incMB.TimesCalled[events.BackupStart], "incremental backup-start events")
			assert.Equal(t, 1, incMB.TimesCalled[events.BackupEnd], "incremental backup-end events")
			assert.Equal(t,
				incMB.CalledWith[events.BackupStart][0][events.BackupID],
				bupID, "incremental backupID pre-declaration")
		})
	}
}

func (suite *BackupOpIntegrationSuite) TestBackup_Run_oneDriveOwnerMigration() {
	ctx, flush := tester.NewContext()
	defer flush()

	var (
		t    = suite.T()
		acct = tester.NewM365Account(t)
		ffs  = control.Toggles{}
		mb   = evmock.NewBus()

		categories = map[path.CategoryType][]string{
			path.FilesCategory: {graph.DeltaURLsFileName, graph.PreviousPathFileName},
		}
	)

	creds, err := acct.M365Config()
	require.NoError(t, err, clues.ToCore(err))

	gc, err := connector.NewGraphConnector(
		ctx,
		acct,
		connector.Users)
	require.NoError(t, err, clues.ToCore(err))

	userable, err := gc.Discovery.Users().GetByID(ctx, suite.user)
	require.NoError(t, err, clues.ToCore(err))

	uid := ptr.Val(userable.GetId())
	uname := ptr.Val(userable.GetUserPrincipalName())

	oldsel := selectors.NewOneDriveBackup([]string{uname})
	oldsel.Include(selTD.OneDriveBackupFolderScope(oldsel))

	bo, _, kw, ms, _, gc, sel, closer := prepNewTestBackupOp(t, ctx, mb, oldsel.Selector, ffs, 0)
	defer closer()

	// ensure the initial owner uses name in both cases
	bo.ResourceOwner = sel.SetDiscreteOwnerIDName(uname, uname)
	// required, otherwise we don't run the migration
	bo.backupVersion = version.All8MigrateUserPNToID - 1

	require.Equalf(
		t,
		bo.ResourceOwner.Name(),
		bo.ResourceOwner.ID(),
		"historical representation of user id [%s] should match pn [%s]",
		bo.ResourceOwner.ID(),
		bo.ResourceOwner.Name())

	// run the initial backup
	runAndCheckBackup(t, ctx, &bo, mb, false)

	newsel := selectors.NewOneDriveBackup([]string{uid})
	newsel.Include(selTD.OneDriveBackupFolderScope(newsel))
	sel = newsel.SetDiscreteOwnerIDName(uid, uname)

	var (
		incMB = evmock.NewBus()
		// the incremental backup op should have a proper user ID for the id.
		incBO = newTestBackupOp(t, ctx, kw, ms, gc, acct, sel, incMB, ffs, closer)
	)

	require.NotEqualf(
		t,
		incBO.ResourceOwner.Name(),
		incBO.ResourceOwner.ID(),
		"current representation of user: id [%s] should differ from PN [%s]",
		incBO.ResourceOwner.ID(),
		incBO.ResourceOwner.Name())

	err = incBO.Run(ctx)
	require.NoError(t, err, clues.ToCore(err))
	checkBackupIsInManifests(t, ctx, kw, &incBO, sel, uid, maps.Keys(categories)...)
	checkMetadataFilesExist(
		t,
		ctx,
		incBO.Results.BackupID,
		kw,
		ms,
		creds.AzureTenantID,
		uid,
		path.OneDriveService,
		categories)

	// 2 on read/writes to account for metadata: 1 delta and 1 path.
	assert.LessOrEqual(t, 2, incBO.Results.ItemsWritten, "items written")
	assert.LessOrEqual(t, 2, incBO.Results.ItemsRead, "items read")
	assert.NoError(t, incBO.Errors.Failure(), "non-recoverable error", clues.ToCore(incBO.Errors.Failure()))
	assert.Empty(t, incBO.Errors.Recovered(), "recoverable/iteration errors")
	assert.Equal(t, 1, incMB.TimesCalled[events.BackupStart], "backup-start events")
	assert.Equal(t, 1, incMB.TimesCalled[events.BackupEnd], "backup-end events")
	assert.Equal(t,
		incMB.CalledWith[events.BackupStart][0][events.BackupID],
		incBO.Results.BackupID, "backupID pre-declaration")

	bid := incBO.Results.BackupID
	bup := &backup.Backup{}

	err = ms.Get(ctx, model.BackupSchema, bid, bup)
	require.NoError(t, err, clues.ToCore(err))

	var (
		ssid  = bup.StreamStoreID
		deets details.Details
		ss    = streamstore.NewStreamer(kw, creds.AzureTenantID, path.OneDriveService)
	)

	err = ss.Read(ctx, ssid, streamstore.DetailsReader(details.UnmarshalTo(&deets)), fault.New(true))
	require.NoError(t, err, clues.ToCore(err))

	for _, ent := range deets.Entries {
		// 46 is the tenant uuid + "onedrive" + two slashes
		if len(ent.RepoRef) > 46 {
			assert.Contains(t, ent.RepoRef, uid)
		}
	}
}

// ---------------------------------------------------------------------------
// SharePoint
// ---------------------------------------------------------------------------

func (suite *BackupOpIntegrationSuite) TestBackup_Run_sharePoint() {
	ctx, flush := tester.NewContext()
	defer flush()

	var (
		t   = suite.T()
		mb  = evmock.NewBus()
		sel = selectors.NewSharePointBackup([]string{suite.site})
	)

	sel.Include(selTD.SharePointBackupFolderScope(sel))

	bo, _, kw, _, _, _, sels, closer := prepNewTestBackupOp(t, ctx, mb, sel.Selector, control.Toggles{}, version.Backup)
	defer closer()

	runAndCheckBackup(t, ctx, &bo, mb, false)
	checkBackupIsInManifests(t, ctx, kw, &bo, sels, suite.site, path.LibrariesCategory)
}

// ---------------------------------------------------------------------------
// helpers
// ---------------------------------------------------------------------------

func newDeleteServicer(t *testing.T) graph.Servicer {
	acct := tester.NewM365Account(t)

	m365, err := acct.M365Config()
	require.NoError(t, err, clues.ToCore(err))

	a, err := graph.CreateAdapter(acct.ID(), m365.AzureClientID, m365.AzureClientSecret)
	require.NoError(t, err, clues.ToCore(err))

	return graph.NewService(a)
}
