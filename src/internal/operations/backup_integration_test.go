package operations

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	msdrive "github.com/microsoftgraph/msgraph-sdk-go/drive"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/users"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"golang.org/x/exp/maps"

	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/connector"
	"github.com/alcionai/corso/src/internal/connector/exchange"
	"github.com/alcionai/corso/src/internal/connector/exchange/api"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/mockconnector"
	"github.com/alcionai/corso/src/internal/connector/onedrive"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/events"
	evmock "github.com/alcionai/corso/src/internal/events/mock"
	"github.com/alcionai/corso/src/internal/kopia"
	"github.com/alcionai/corso/src/internal/model"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/version"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/backup"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/store"
)

const incrementalsDestContainerPrefix = "incrementals_ci_"

// ---------------------------------------------------------------------------
// helpers
// ---------------------------------------------------------------------------

// prepNewTestBackupOp generates all clients required to run a backup operation,
// returning both a backup operation created with those clients, as well as
// the clients themselves.
//
//revive:disable:context-as-argument
func prepNewTestBackupOp(
	t *testing.T,
	ctx context.Context,
	bus events.Eventer,
	sel selectors.Selector,
	featureToggles control.Toggles,
) (BackupOperation, account.Account, *kopia.Wrapper, *kopia.ModelStore, func()) {
	//revive:enable:context-as-argument
	acct := tester.NewM365Account(t)

	// need to initialize the repository before we can test connecting to it.
	st := tester.NewPrefixedS3Storage(t)

	k := kopia.NewConn(st)
	require.NoError(t, k.Initialize(ctx))

	// kopiaRef comes with a count of 1 and Wrapper bumps it again so safe
	// to close here.
	closer := func() { k.Close(ctx) }

	kw, err := kopia.NewWrapper(k)
	if !assert.NoError(t, err) {
		closer()
		t.FailNow()
	}

	closer = func() {
		k.Close(ctx)
		kw.Close(ctx)
	}

	ms, err := kopia.NewModelStore(k)
	if !assert.NoError(t, err) {
		closer()
		t.FailNow()
	}

	closer = func() {
		k.Close(ctx)
		kw.Close(ctx)
		ms.Close(ctx)
	}

	bo := newTestBackupOp(t, ctx, kw, ms, acct, sel, bus, featureToggles, closer)

	return bo, acct, kw, ms, closer
}

// newTestBackupOp accepts the clients required to compose a backup operation, plus
// any other metadata, and uses them to generate a new backup operation.  This
// allows backup chains to utilize the same temp directory and configuration
// details.
//
//revive:disable:context-as-argument
func newTestBackupOp(
	t *testing.T,
	ctx context.Context,
	kw *kopia.Wrapper,
	ms *kopia.ModelStore,
	acct account.Account,
	sel selectors.Selector,
	bus events.Eventer,
	featureToggles control.Toggles,
	closer func(),
) BackupOperation {
	//revive:enable:context-as-argument
	var (
		sw   = store.NewKopiaStore(ms)
		opts = control.Options{}
	)

	opts.ToggleFeatures = featureToggles

	bo, err := NewBackupOperation(ctx, opts, kw, sw, acct, sel, bus)
	if !assert.NoError(t, err) {
		closer()
		t.FailNow()
	}

	return bo
}

//revive:disable:context-as-argument
func runAndCheckBackup(
	t *testing.T,
	ctx context.Context,
	bo *BackupOperation,
	mb *evmock.Bus,
) {
	//revive:enable:context-as-argument
	require.NoError(t, bo.Run(ctx))
	require.NotEmpty(t, bo.Results, "the backup had non-zero results")
	require.NotEmpty(t, bo.Results.BackupID, "the backup generated an ID")
	require.Equalf(
		t,
		Completed,
		bo.Status,
		"backup status should be Completed, got %s",
		bo.Status)
	require.Less(t, 0, bo.Results.ItemsWritten)

	assert.Less(t, 0, bo.Results.ItemsRead, "count of items read")
	assert.Less(t, int64(0), bo.Results.BytesRead, "bytes read")
	assert.Less(t, int64(0), bo.Results.BytesUploaded, "bytes uploaded")
	assert.Equal(t, 1, bo.Results.ResourceOwners, "count of resource owners")
	assert.NoError(t, bo.Errors.Failure(), "incremental non-recoverable error")
	assert.Empty(t, bo.Errors.Recovered(), "incremental recoverable/iteration errors")
	assert.Equal(t, 1, mb.TimesCalled[events.BackupStart], "backup-start events")
	assert.Equal(t, 1, mb.TimesCalled[events.BackupEnd], "backup-end events")
	assert.Equal(t,
		mb.CalledWith[events.BackupStart][0][events.BackupID],
		bo.Results.BackupID, "backupID pre-declaration")
}

//revive:disable:context-as-argument
func checkBackupIsInManifests(
	t *testing.T,
	ctx context.Context,
	kw *kopia.Wrapper,
	bo *BackupOperation,
	sel selectors.Selector,
	resourceOwner string,
	categories ...path.CategoryType,
) {
	//revive:enable:context-as-argument
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
			require.NoError(t, err)

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

//revive:disable:context-as-argument
func checkMetadataFilesExist(
	t *testing.T,
	ctx context.Context,
	backupID model.StableID,
	kw *kopia.Wrapper,
	ms *kopia.ModelStore,
	tenant, user string,
	service path.ServiceType,
	filesByCat map[path.CategoryType][]string,
) {
	//revive:enable:context-as-argument
	for category, files := range filesByCat {
		t.Run(category.String(), func(t *testing.T) {
			bup := &backup.Backup{}

			err := ms.Get(ctx, model.BackupSchema, backupID, bup)
			if !assert.NoError(t, err) {
				return
			}

			paths := []path.Path{}
			pathsByRef := map[string][]string{}

			for _, fName := range files {
				p, err := path.Builder{}.
					Append(fName).
					ToServiceCategoryMetadataPath(tenant, user, service, category, true)
				if !assert.NoError(t, err, "bad metadata path") {
					continue
				}

				dir, err := p.Dir()
				if !assert.NoError(t, err, "parent path") {
					continue
				}

				paths = append(paths, p)
				pathsByRef[dir.ShortRef()] = append(pathsByRef[dir.ShortRef()], fName)
			}

			cols, err := kw.RestoreMultipleItems(ctx, bup.SnapshotID, paths, nil, fault.New(true))
			assert.NoError(t, err)

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

//revive:disable:context-as-argument
func generateContainerOfItems(
	t *testing.T,
	ctx context.Context,
	gc *connector.GraphConnector,
	service path.ServiceType,
	acct account.Account,
	cat path.CategoryType,
	sel selectors.Selector,
	tenantID, userID, driveID, destFldr string,
	howManyItems int,
	backupVersion int,
	dbf dataBuilderFunc,
) *details.Details {
	//revive:enable:context-as-argument
	t.Helper()

	items := make([]incrementalItem, 0, howManyItems)

	for i := 0; i < howManyItems; i++ {
		id, d := generateItemData(t, cat, userID, dbf)

		items = append(items, incrementalItem{
			name: id,
			data: d,
		})
	}

	pathFolders := []string{destFldr}
	if service == path.OneDriveService {
		pathFolders = []string{"drives", driveID, "root:", destFldr}
	}

	collections := []incrementalCollection{{
		pathFolders: pathFolders,
		category:    cat,
		items:       items,
	}}

	dest := control.DefaultRestoreDestination(common.SimpleTimeTesting)
	dest.ContainerName = destFldr

	dataColls := buildCollections(
		t,
		service,
		tenantID, userID,
		dest,
		collections)

	deets, err := gc.RestoreDataCollections(
		ctx,
		backupVersion,
		acct,
		sel,
		dest,
		control.Options{RestorePermissions: true},
		dataColls,
		fault.New(true))
	require.NoError(t, err)

	return deets
}

func generateItemData(
	t *testing.T,
	category path.CategoryType,
	resourceOwner string,
	dbf dataBuilderFunc,
) (string, []byte) {
	var (
		now       = common.Now()
		nowLegacy = common.FormatLegacyTime(time.Now())
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

		mc := mockconnector.NewMockExchangeCollection(pth, pth, len(c.items))

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
	tenant, user string,
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
		p, err = pb.ToDataLayerExchangePathForCategory(tenant, user, category, isItem)
	case path.OneDriveService:
		p, err = pb.ToDataLayerOneDrivePath(tenant, user, isItem)
	default:
		err = errors.Errorf("unknown service %s", service.String())
	}

	require.NoError(t, err)

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
			[][]string{tester.AWSStorageCredEnvs, tester.M365AcctCredEnvs},
			tester.CorsoOperationTests,
			tester.CorsoOperationBackupTests),
	})
}

func (suite *BackupOpIntegrationSuite) SetupSuite() {
	suite.user = tester.M365UserID(suite.T())
	suite.site = tester.M365SiteID(suite.T())
}

func (suite *BackupOpIntegrationSuite) TestNewBackupOperation() {
	kw := &kopia.Wrapper{}
	sw := &store.Wrapper{}
	acct := tester.NewM365Account(suite.T())

	table := []struct {
		name     string
		opts     control.Options
		kw       *kopia.Wrapper
		sw       *store.Wrapper
		acct     account.Account
		targets  []string
		errCheck assert.ErrorAssertionFunc
	}{
		{"good", control.Options{}, kw, sw, acct, nil, assert.NoError},
		{"missing kopia", control.Options{}, nil, sw, acct, nil, assert.Error},
		{"missing modelstore", control.Options{}, kw, nil, acct, nil, assert.Error},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			ctx, flush := tester.NewContext()
			defer flush()

			_, err := NewBackupOperation(
				ctx,
				test.opts,
				test.kw,
				test.sw,
				test.acct,
				selectors.Selector{DiscreteOwner: "test"},
				evmock.NewBus())
			test.errCheck(suite.T(), err)
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

	owners := []string{suite.user}

	tests := []struct {
		name           string
		selector       func() *selectors.ExchangeBackup
		resourceOwner  string
		category       path.CategoryType
		metadataFiles  []string
		runIncremental bool
	}{
		{
			name: "Mail",
			selector: func() *selectors.ExchangeBackup {
				sel := selectors.NewExchangeBackup(owners)
				sel.Include(sel.MailFolders([]string{exchange.DefaultMailFolder}, selectors.PrefixMatch()))
				sel.DiscreteOwner = suite.user

				return sel
			},
			resourceOwner:  suite.user,
			category:       path.EmailCategory,
			metadataFiles:  exchange.MetadataFileNames(path.EmailCategory),
			runIncremental: true,
		},
		{
			name: "Contacts",
			selector: func() *selectors.ExchangeBackup {
				sel := selectors.NewExchangeBackup(owners)
				sel.Include(sel.ContactFolders([]string{exchange.DefaultContactFolder}, selectors.PrefixMatch()))
				return sel
			},
			resourceOwner:  suite.user,
			category:       path.ContactsCategory,
			metadataFiles:  exchange.MetadataFileNames(path.ContactsCategory),
			runIncremental: true,
		},
		{
			name: "Calendar Events",
			selector: func() *selectors.ExchangeBackup {
				sel := selectors.NewExchangeBackup(owners)
				sel.Include(sel.EventCalendars([]string{exchange.DefaultCalendar}, selectors.PrefixMatch()))
				return sel
			},
			resourceOwner: suite.user,
			category:      path.EventsCategory,
			metadataFiles: exchange.MetadataFileNames(path.EventsCategory),
		},
	}
	for _, test := range tests {
		suite.Run(test.name, func() {
			var (
				t   = suite.T()
				mb  = evmock.NewBus()
				sel = test.selector().Selector
				ffs = control.Toggles{}
			)

			bo, acct, kw, ms, closer := prepNewTestBackupOp(t, ctx, mb, sel, ffs)
			defer closer()

			m365, err := acct.M365Config()
			require.NoError(t, err)

			// run the tests
			runAndCheckBackup(t, ctx, &bo, mb)
			checkBackupIsInManifests(t, ctx, kw, &bo, sel, test.resourceOwner, test.category)
			checkMetadataFilesExist(
				t,
				ctx,
				bo.Results.BackupID,
				kw,
				ms,
				m365.AzureTenantID,
				test.resourceOwner,
				path.ExchangeService,
				map[path.CategoryType][]string{test.category: test.metadataFiles},
			)

			if !test.runIncremental {
				return
			}

			// Basic, happy path incremental test.  No changes are dictated or expected.
			// This only tests that an incremental backup is runnable at all, and that it
			// produces fewer results than the last backup.
			var (
				incMB = evmock.NewBus()
				incBO = newTestBackupOp(t, ctx, kw, ms, acct, sel, incMB, ffs, closer)
			)

			runAndCheckBackup(t, ctx, &incBO, incMB)
			checkBackupIsInManifests(t, ctx, kw, &incBO, sel, test.resourceOwner, test.category)
			checkMetadataFilesExist(
				t,
				ctx,
				incBO.Results.BackupID,
				kw,
				ms,
				m365.AzureTenantID,
				test.resourceOwner,
				path.ExchangeService,
				map[path.CategoryType][]string{test.category: test.metadataFiles},
			)

			// do some additional checks to ensure the incremental dealt with fewer items.
			assert.Greater(t, bo.Results.ItemsWritten, incBO.Results.ItemsWritten, "incremental items written")
			assert.Greater(t, bo.Results.ItemsRead, incBO.Results.ItemsRead, "incremental items read")
			assert.Greater(t, bo.Results.BytesRead, incBO.Results.BytesRead, "incremental bytes read")
			assert.Greater(t, bo.Results.BytesUploaded, incBO.Results.BytesUploaded, "incremental bytes uploaded")
			assert.Equal(t, bo.Results.ResourceOwners, incBO.Results.ResourceOwners, "incremental backup resource owner")
			assert.NoError(t, incBO.Errors.Failure(), "incremental non-recoverable error")
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
func (suite *BackupOpIntegrationSuite) TestBackup_Run_exchangeIncrementals() {
	ctx, flush := tester.NewContext()
	defer flush()

	tester.LogTimeOfTest(suite.T())

	var (
		t          = suite.T()
		acct       = tester.NewM365Account(t)
		ffs        = control.Toggles{}
		mb         = evmock.NewBus()
		now        = common.Now()
		owners     = []string{suite.user}
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
	)

	m365, err := acct.M365Config()
	require.NoError(t, err)

	gc, err := connector.NewGraphConnector(
		ctx,
		graph.HTTPClient(graph.NoTimeout()),
		acct,
		connector.Users,
		fault.New(true))
	require.NoError(t, err)

	ac, err := api.NewClient(m365)
	require.NoError(t, err)

	// generate 3 new folders with two items each.
	// Only the first two folders will be part of the initial backup and
	// incrementals.  The third folder will be introduced partway through
	// the changes.
	// This should be enough to cover most delta actions, since moving one
	// container into another generates a delta for both addition and deletion.
	type contDeets struct {
		containerID string
		deets       *details.Details
	}

	mailDBF := func(id, timeStamp, subject, body string) []byte {
		return mockconnector.GetMockMessageWith(
			suite.user, suite.user, suite.user,
			subject, body, body,
			now, now, now, now)
	}

	contactDBF := func(id, timeStamp, subject, body string) []byte {
		given, mid, sur := id[:8], id[9:13], id[len(id)-12:]

		return mockconnector.GetMockContactBytesWith(
			given+" "+sur,
			sur+", "+given,
			given, mid, sur,
			"123-456-7890",
		)
	}

	eventDBF := func(id, timeStamp, subject, body string) []byte {
		return mockconnector.GetMockEventWith(
			suite.user, subject, body, body,
			now, now, false)
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
			deets := generateContainerOfItems(
				t,
				ctx,
				gc,
				path.ExchangeService,
				acct,
				category,
				selectors.NewExchangeRestore(owners).Selector,
				m365.AzureTenantID, suite.user, "", destName,
				2,
				version.Backup,
				gen.dbf)

			dataset[category].dests[destName] = contDeets{"", deets}
		}
	}

	// verify test data was populated, and track it for comparisons
	for category, gen := range dataset {
		qp := graph.QueryParams{
			Category:      category,
			ResourceOwner: suite.user,
			Credentials:   m365,
		}
		cr, err := exchange.PopulateExchangeContainerResolver(ctx, qp, fault.New(true))
		require.NoError(t, err, "populating %s container resolver", category)

		for destName, dest := range gen.dests {
			p, err := path.FromDataLayerPath(dest.deets.Entries[0].RepoRef, true)
			require.NoError(t, err)

			id, ok := cr.PathInCache(p.Folder(false))
			require.True(t, ok, "dir %s found in %s cache", p.Folder(false), category)

			d := dataset[category].dests[destName]
			d.containerID = id
			dataset[category].dests[destName] = d
		}
	}

	// container3 and containerRename don't exist yet.  Those will get created
	// later on during the tests.  Putting their identifiers into the selector
	// at this point is harmless.
	containers := []string{container1, container2, container3, containerRename}
	sel := selectors.NewExchangeBackup(owners)
	sel.Include(
		sel.MailFolders(containers, selectors.PrefixMatch()),
		sel.ContactFolders(containers, selectors.PrefixMatch()),
	)

	bo, _, kw, ms, closer := prepNewTestBackupOp(t, ctx, mb, sel.Selector, ffs)
	defer closer()

	// run the initial backup
	runAndCheckBackup(t, ctx, &bo, mb)

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
				// contacts and events cannot be sufoldered; this is an email-only change
				toContainer := dataset[path.EmailCategory].dests[container1].containerID
				fromContainer := dataset[path.EmailCategory].dests[container2].containerID

				body := users.NewItemMailFoldersItemMovePostRequestBody()
				body.SetDestinationId(&toContainer)

				_, err := gc.Service.
					Client().
					UsersById(suite.user).
					MailFoldersById(fromContainer).
					Move().
					Post(ctx, body, nil)
				require.NoError(t, err)
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
						require.NoError(
							t,
							ac.Mail().DeleteContainer(ctx, suite.user, containerID),
							"deleting an email folder")
					case path.ContactsCategory:
						require.NoError(
							t,
							ac.Contacts().DeleteContainer(ctx, suite.user, containerID),
							"deleting a contacts folder")
					case path.EventsCategory:
						require.NoError(
							t,
							ac.Events().DeleteContainer(ctx, suite.user, containerID),
							"deleting a calendar")
					}
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
						path.ExchangeService,
						acct,
						category,
						selectors.NewExchangeRestore(owners).Selector,
						m365.AzureTenantID, suite.user, "", container3,
						2,
						version.Backup,
						gen.dbf)

					qp := graph.QueryParams{
						Category:      category,
						ResourceOwner: suite.user,
						Credentials:   m365,
					}
					cr, err := exchange.PopulateExchangeContainerResolver(ctx, qp, fault.New(true))
					require.NoError(t, err, "populating %s container resolver", category)

					p, err := path.FromDataLayerPath(deets.Entries[0].RepoRef, true)
					require.NoError(t, err)

					id, ok := cr.PathInCache(p.Folder(false))
					require.True(t, ok, "dir %s found in %s cache", p.Folder(false), category)

					dataset[category].dests[container3] = contDeets{id, deets}
				}
			},
			itemsRead:    4,
			itemsWritten: 4,
		},
		{
			name: "rename a folder",
			updateUserData: func(t *testing.T) {
				for category, d := range dataset {
					containerID := d.dests[container3].containerID
					cli := gc.Service.Client().UsersById(suite.user)

					// copy the container info, since both names should
					// reference the same container by id.  Though the
					// details refs won't line up, so those get deleted.
					d.dests[containerRename] = contDeets{
						containerID: d.dests[container3].containerID,
						deets:       nil,
					}

					switch category {
					case path.EmailCategory:
						cmf := cli.MailFoldersById(containerID)

						body, err := cmf.Get(ctx, nil)
						require.NoError(t, err, "getting mail folder")

						body.SetDisplayName(&containerRename)
						_, err = cmf.Patch(ctx, body, nil)
						require.NoError(t, err, "updating mail folder name")

					case path.ContactsCategory:
						ccf := cli.ContactFoldersById(containerID)

						body, err := ccf.Get(ctx, nil)
						require.NoError(t, err, "getting contact folder")

						body.SetDisplayName(&containerRename)
						_, err = ccf.Patch(ctx, body, nil)
						require.NoError(t, err, "updating contact folder name")

					case path.EventsCategory:
						cbi := cli.CalendarsById(containerID)

						body, err := cbi.Get(ctx, nil)
						require.NoError(t, err, "getting calendar")

						body.SetName(&containerRename)
						_, err = cbi.Patch(ctx, body, nil)
						require.NoError(t, err, "updating calendar name")
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
					cli := gc.Service.Client().UsersById(suite.user)

					switch category {
					case path.EmailCategory:
						_, itemData := generateItemData(t, category, suite.user, mailDBF)
						body, err := support.CreateMessageFromBytes(itemData)
						require.NoError(t, err, "transforming mail bytes to messageable")

						_, err = cli.MailFoldersById(containerID).Messages().Post(ctx, body, nil)
						require.NoError(t, err, "posting email item")

					case path.ContactsCategory:
						_, itemData := generateItemData(t, category, suite.user, contactDBF)
						body, err := support.CreateContactFromBytes(itemData)
						require.NoError(t, err, "transforming contact bytes to contactable")

						_, err = cli.ContactFoldersById(containerID).Contacts().Post(ctx, body, nil)
						require.NoError(t, err, "posting contact item")

					case path.EventsCategory:
						_, itemData := generateItemData(t, category, suite.user, eventDBF)
						body, err := support.CreateEventFromBytes(itemData)
						require.NoError(t, err, "transforming event bytes to eventable")

						_, err = cli.CalendarsById(containerID).Events().Post(ctx, body, nil)
						require.NoError(t, err, "posting events item")
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
					cli := gc.Service.Client().UsersById(suite.user)

					switch category {
					case path.EmailCategory:
						ids, _, _, err := ac.Mail().GetAddedAndRemovedItemIDs(ctx, suite.user, containerID, "")
						require.NoError(t, err, "getting message ids")
						require.NotEmpty(t, ids, "message ids in folder")

						err = cli.MessagesById(ids[0]).Delete(ctx, nil)
						require.NoError(t, err, "deleting email item")

					case path.ContactsCategory:
						ids, _, _, err := ac.Contacts().GetAddedAndRemovedItemIDs(ctx, suite.user, containerID, "")
						require.NoError(t, err, "getting contact ids")
						require.NotEmpty(t, ids, "contact ids in folder")

						err = cli.ContactsById(ids[0]).Delete(ctx, nil)
						require.NoError(t, err, "deleting contact item")

					case path.EventsCategory:
						ids, _, _, err := ac.Events().GetAddedAndRemovedItemIDs(ctx, suite.user, containerID, "")
						require.NoError(t, err, "getting event ids")
						require.NotEmpty(t, ids, "event ids in folder")

						err = cli.CalendarsById(ids[0]).Delete(ctx, nil)
						require.NoError(t, err, "deleting calendar")
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
				incBO = newTestBackupOp(t, ctx, kw, ms, acct, sel.Selector, incMB, ffs, closer)
			)

			test.updateUserData(t)
			require.NoError(t, incBO.Run(ctx))
			checkBackupIsInManifests(t, ctx, kw, &incBO, sel.Selector, suite.user, maps.Keys(categories)...)
			checkMetadataFilesExist(
				t,
				ctx,
				incBO.Results.BackupID,
				kw,
				ms,
				m365.AzureTenantID,
				suite.user,
				path.ExchangeService,
				categories,
			)

			// do some additional checks to ensure the incremental dealt with fewer items.
			// +4 on read/writes to account for metadata: 1 delta and 1 path for each type.
			assert.Equal(t, test.itemsWritten+4, incBO.Results.ItemsWritten, "incremental items written")
			assert.Equal(t, test.itemsRead+4, incBO.Results.ItemsRead, "incremental items read")
			assert.NoError(t, incBO.Errors.Failure(), "incremental non-recoverable error")
			assert.Empty(t, incBO.Errors.Recovered(), "incremental recoverable/iteration errors")
			assert.Equal(t, 1, incMB.TimesCalled[events.BackupStart], "incremental backup-start events")
			assert.Equal(t, 1, incMB.TimesCalled[events.BackupEnd], "incremental backup-end events")
			assert.Equal(t,
				incMB.CalledWith[events.BackupStart][0][events.BackupID],
				incBO.Results.BackupID, "incremental backupID pre-declaration")
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
		t          = suite.T()
		mb         = evmock.NewBus()
		m365UserID = tester.SecondaryM365UserID(t)
		sel        = selectors.NewOneDriveBackup([]string{m365UserID})
	)

	sel.Include(sel.AllData())

	bo, _, _, _, closer := prepNewTestBackupOp(t, ctx, mb, sel.Selector, control.Toggles{EnablePermissionsBackup: true})
	defer closer()

	runAndCheckBackup(t, ctx, &bo, mb)
}

func mustGetDefaultDriveID(
	t *testing.T,
	ctx context.Context, //revive:disable-line:context-as-argument
	service graph.Servicer,
	userID string,
) string {
	d, err := service.Client().UsersById(userID).Drive().Get(ctx, nil)
	if err != nil {
		err = graph.Wrap(
			ctx,
			err,
			"failed to retrieve default user drive").
			With("user", userID)
	}

	require.NoError(t, err)

	id, ok := ptr.ValOK(d.GetId())
	require.True(t, ok, "drive ID not set")
	require.NotEmpty(t, id)

	return id
}

// TestBackup_Run ensures that Integration Testing works for OneDrive
func (suite *BackupOpIntegrationSuite) TestBackup_Run_oneDriveIncrementals() {
	// TODO: Enable once we have https://github.com/alcionai/corso/pull/2642
	suite.T().Skip("Enable once OneDrive incrementals is available")

	ctx, flush := tester.NewContext()
	defer flush()

	var (
		t    = suite.T()
		acct = tester.NewM365Account(t)
		ffs  = control.Toggles{}
		mb   = evmock.NewBus()

		// `now` has to be formatted with SimpleDateTimeOneDrive as
		// some onedrive cannot have `:` in file/folder names
		now = common.FormatNow(common.SimpleTimeTesting)

		owners = []string{suite.user}

		categories = map[path.CategoryType][]string{
			path.FilesCategory: {graph.DeltaURLsFileName, graph.PreviousPathFileName},
		}
		container1 = fmt.Sprintf("%s%d_%s", incrementalsDestContainerPrefix, 1, now)
		container2 = fmt.Sprintf("%s%d_%s", incrementalsDestContainerPrefix, 2, now)
		container3 = fmt.Sprintf("%s%d_%s", incrementalsDestContainerPrefix, 3, now)

		genDests = []string{container1, container2}
	)

	m365, err := acct.M365Config()
	require.NoError(t, err)

	gc, err := connector.NewGraphConnector(
		ctx,
		graph.HTTPClient(graph.NoTimeout()),
		acct,
		connector.Users,
		fault.New(true))
	require.NoError(t, err)

	driveID := mustGetDefaultDriveID(t, ctx, gc.Service, suite.user)

	fileDBF := func(id, timeStamp, subject, body string) []byte {
		return []byte(id + subject)
	}

	// Populate initial test data.
	// Generate 2 new folders with two items each. Only the first two
	// folders will be part of the initial backup and
	// incrementals. The third folder will be introduced partway
	// through the changes. This should be enough to cover most delta
	// actions.
	for _, destName := range genDests {
		generateContainerOfItems(
			t,
			ctx,
			gc,
			path.OneDriveService,
			acct,
			path.FilesCategory,
			selectors.NewOneDriveRestore(owners).Selector,
			m365.AzureTenantID, suite.user, driveID, destName,
			2,
			// Use an old backup version so we don't need metadata files.
			0,
			fileDBF)
	}

	containerIDs := map[string]string{}

	// verify test data was populated, and track it for comparisons
	for _, destName := range genDests {
		// Use path-based indexing to get the folder's ID. This is sourced from the
		// onedrive package `getFolder` function.
		itemURL := fmt.Sprintf(
			"https://graph.microsoft.com/v1.0/drives/%s/root:/%s",
			driveID,
			destName)
		resp, err := msdrive.NewItemsDriveItemItemRequestBuilder(itemURL, gc.Service.Adapter()).
			Get(ctx, nil)
		require.NoErrorf(t, err, "getting drive folder ID", "folder name: %s", destName)

		containerIDs[destName] = ptr.Val(resp.GetId())
	}

	// container3 does not exist yet. It will get created later on
	// during the tests.
	containers := []string{container1, container2, container3}
	sel := selectors.NewOneDriveBackup(owners)
	sel.Include(
		sel.Folders(containers, selectors.PrefixMatch()),
	)

	bo, _, kw, ms, closer := prepNewTestBackupOp(t, ctx, mb, sel.Selector, ffs)
	defer closer()

	// run the initial backup
	runAndCheckBackup(t, ctx, &bo, mb)

	var (
		newFile     models.DriveItemable
		newFileName = "new_file.txt"
	)

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
			name: "create a new file",
			updateUserData: func(t *testing.T) {
				targetContainer := containerIDs[container1]
				driveItem := models.NewDriveItem()
				driveItem.SetName(&newFileName)
				driveItem.SetFile(models.NewFile())
				newFile, err = onedrive.CreateItem(
					ctx,
					gc.Service,
					driveID,
					targetContainer,
					driveItem,
				)
				require.NoError(t, err, "creating new file")
			},
			itemsRead:    1, // .data file for newitem
			itemsWritten: 3, // .data and .meta for newitem, .dirmeta for parent
		},
		{
			name: "update contents a file",
			updateUserData: func(t *testing.T) {
				err := gc.Service.
					Client().
					DrivesById(driveID).
					ItemsById(ptr.Val(newFile.GetId())).
					Content().
					Put(ctx, []byte("new content"), nil)
				require.NoError(t, err, "updating file content")
			},
			itemsRead:    1, // .data file for newitem
			itemsWritten: 3, // .data and .meta for newitem, .dirmeta for parent
		},
		{
			name: "rename a file",
			updateUserData: func(t *testing.T) {
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
				require.NoError(t, err, "renaming file")
			},
			itemsRead:    1, // .data file for newitem
			itemsWritten: 3, // .data and .meta for newitem, .dirmeta for parent
		},
		{
			name: "move a file between folders",
			updateUserData: func(t *testing.T) {
				dest := containerIDs[container1]

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
				require.NoError(t, err, "moving file between folders")
			},
			itemsRead:    1, // .data file for newitem
			itemsWritten: 3, // .data and .meta for newitem, .dirmeta for parent
		},
		{
			name: "delete file",
			updateUserData: func(t *testing.T) {
				err = gc.Service.
					Client().
					DrivesById(driveID).
					ItemsById(ptr.Val(newFile.GetId())).
					Delete(ctx, nil)
				require.NoError(t, err, "deleting file")
			},
			itemsRead:    0,
			itemsWritten: 0,
		},
		{
			name: "move a folder to a subfolder",
			updateUserData: func(t *testing.T) {
				dest := containerIDs[container1]
				source := containerIDs[container2]

				driveItem := models.NewDriveItem()
				driveItem.SetName(&container2)
				parentRef := models.NewItemReference()
				parentRef.SetId(&dest)
				driveItem.SetParentReference(parentRef)

				_, err := gc.Service.
					Client().
					DrivesById(driveID).
					ItemsById(source).
					Patch(ctx, driveItem, nil)
				require.NoError(t, err, "moving folder")
			},
			itemsRead:    0,
			itemsWritten: 7, // 2*2(data and meta of 2 files) + 3 (dirmeta of two moved folders and target)
		},
		{
			name: "rename a folder",
			updateUserData: func(t *testing.T) {
				parent := containerIDs[container1]
				child := containerIDs[container2]

				driveItem := models.NewDriveItem()
				name := "renamed_folder"
				driveItem.SetName(&name)
				parentRef := models.NewItemReference()
				parentRef.SetId(&parent)
				driveItem.SetParentReference(parentRef)

				_, err := gc.Service.
					Client().
					DrivesById(driveID).
					ItemsById(child).
					Patch(ctx, driveItem, nil)
				require.NoError(t, err, "renaming folder")
			},
			itemsRead:    0,
			itemsWritten: 7, // 2*2(data and meta of 2 files) + 3 (dirmeta of two moved folders and target)
		},
		{
			name: "delete a folder",
			updateUserData: func(t *testing.T) {
				container := containerIDs[container2]
				err := gc.Service.
					Client().
					DrivesById(driveID).
					ItemsById(container).
					Delete(ctx, nil)
				require.NoError(t, err, "deleting folder")
			},
			itemsRead:    0,
			itemsWritten: 0,
		},
		{
			name: "add a new folder",
			updateUserData: func(t *testing.T) {
				generateContainerOfItems(
					t,
					ctx,
					gc,
					path.OneDriveService,
					acct,
					path.FilesCategory,
					selectors.NewOneDriveRestore(owners).Selector,
					m365.AzureTenantID, suite.user, driveID, container3,
					2,
					0,
					fileDBF)

				// Validate creation
				itemURL := fmt.Sprintf(
					"https://graph.microsoft.com/v1.0/drives/%s/root:/%s",
					driveID,
					container3)
				resp, err := msdrive.NewItemsDriveItemItemRequestBuilder(itemURL, gc.Service.Adapter()).
					Get(ctx, nil)
				require.NoErrorf(t, err, "getting drive folder ID", "folder name: %s", container3)

				containerIDs[container3] = ptr.Val(resp.GetId())
			},
			itemsRead:    4, // 2*2 (.data and .meta for 2 files)
			itemsWritten: 6, // read items + 2 directory meta
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			var (
				t     = suite.T()
				incMB = evmock.NewBus()
				incBO = newTestBackupOp(t, ctx, kw, ms, acct, sel.Selector, incMB, ffs, closer)
			)

			tester.LogTimeOfTest(suite.T())

			test.updateUserData(t)
			require.NoError(t, incBO.Run(ctx))
			checkBackupIsInManifests(t, ctx, kw, &incBO, sel.Selector, suite.user, maps.Keys(categories)...)
			checkMetadataFilesExist(
				t,
				ctx,
				incBO.Results.BackupID,
				kw,
				ms,
				m365.AzureTenantID,
				suite.user,
				path.OneDriveService,
				categories,
			)

			// do some additional checks to ensure the incremental dealt with fewer items.
			// +2 on read/writes to account for metadata: 1 delta and 1 path.
			assert.Equal(t, test.itemsWritten+2, incBO.Results.ItemsWritten, "incremental items written")
			assert.Equal(t, test.itemsRead+2, incBO.Results.ItemsRead, "incremental items read")
			assert.NoError(t, incBO.Errors.Failure(), "incremental non-recoverable error")
			assert.Empty(t, incBO.Errors.Recovered(), "incremental recoverable/iteration errors")
			assert.Equal(t, 1, incMB.TimesCalled[events.BackupStart], "incremental backup-start events")
			assert.Equal(t, 1, incMB.TimesCalled[events.BackupEnd], "incremental backup-end events")
			assert.Equal(t,
				incMB.CalledWith[events.BackupStart][0][events.BackupID],
				incBO.Results.BackupID, "incremental backupID pre-declaration")
		})
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

	sel.Include(sel.Libraries(selectors.Any()))

	bo, _, kw, _, closer := prepNewTestBackupOp(t, ctx, mb, sel.Selector, control.Toggles{})
	defer closer()

	runAndCheckBackup(t, ctx, &bo, mb)
	checkBackupIsInManifests(t, ctx, kw, &bo, sel.Selector, suite.site, path.LibrariesCategory)
}
