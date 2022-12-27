package operations

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/connector/exchange"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/events"
	evmock "github.com/alcionai/corso/src/internal/events/mock"
	"github.com/alcionai/corso/src/internal/kopia"
	"github.com/alcionai/corso/src/internal/model"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/backup"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/store"
)

// ---------------------------------------------------------------------------
// helpers
// ---------------------------------------------------------------------------

// prepNewBackupOp generates all clients required to run a backup operation,
// returning both a backup operation created with those clients, as well as
// the clients themselves.
//
//revive:disable:context-as-argument
func prepNewBackupOp(
	t *testing.T,
	ctx context.Context,
	bus events.Eventer,
	sel selectors.Selector,
	featureFlags control.FeatureFlags,
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

	bo := newBackupOp(t, ctx, kw, ms, acct, sel, bus, featureFlags, closer)

	return bo, acct, kw, ms, closer
}

// newBackupOp accepts the clients required to compose a backup operation, plus
// any other metadata, and uses them to generate a new backup operation.  This
// allows backup chains to utilize the same temp directory and configuration
// details.
//
//revive:disable:context-as-argument
func newBackupOp(
	t *testing.T,
	ctx context.Context,
	kw *kopia.Wrapper,
	ms *kopia.ModelStore,
	acct account.Account,
	sel selectors.Selector,
	bus events.Eventer,
	featureFlags control.FeatureFlags,
	closer func(),
) BackupOperation {
	//revive:enable:context-as-argument
	var (
		sw   = store.NewKopiaStore(ms)
		opts = control.Options{}
	)

	opts.EnabledFeatures = featureFlags

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
		bo.Status,
	)
	require.Less(t, 0, bo.Results.ItemsWritten)

	assert.Less(t, 0, bo.Results.ItemsRead, "count of items read")
	assert.Less(t, int64(0), bo.Results.BytesRead, "bytes read")
	assert.Less(t, int64(0), bo.Results.BytesUploaded, "bytes uploaded")
	assert.Equal(t, 1, bo.Results.ResourceOwners, "count of resource owners")
	assert.NoError(t, bo.Results.ReadErrors, "errors reading data")
	assert.NoError(t, bo.Results.WriteErrors, "errors writing data")
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
	category path.CategoryType,
) {
	//revive:enable:context-as-argument
	var (
		sck, scv = kopia.MakeServiceCat(sel.PathService(), category)
		oc       = &kopia.OwnersCats{
			ResourceOwners: map[string]struct{}{resourceOwner: {}},
			ServiceCats:    map[string]kopia.ServiceCat{sck: scv},
		}
		tags  = map[string]string{kopia.TagBackupCategory: ""}
		found bool
	)

	mans, err := kw.FetchPrevSnapshotManifests(ctx, oc, tags)
	require.NoError(t, err)

	for _, man := range mans {
		tk, _ := kopia.MakeTagKV(kopia.TagBackupID)
		if man.Tags[tk] == string(bo.Results.BackupID) {
			found = true
			break
		}
	}

	assert.True(t, found, "backup retrieved by previous snapshot manifest")
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
	category path.CategoryType,
	files []string,
) {
	//revive:enable:context-as-argument
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

	cols, err := kw.RestoreMultipleItems(ctx, bup.SnapshotID, paths, nil)
	assert.NoError(t, err)

	for _, col := range cols {
		itemNames := []string{}

		for item := range col.Items() {
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
}

// ---------------------------------------------------------------------------
// integration tests
// ---------------------------------------------------------------------------

type BackupOpIntegrationSuite struct {
	suite.Suite
	user, site string
}

func TestBackupOpIntegrationSuite(t *testing.T) {
	if err := tester.RunOnAny(
		tester.CorsoCITests,
		tester.CorsoOperationTests,
		tester.CorsoOperationBackupTests,
	); err != nil {
		t.Skip(err)
	}

	suite.Run(t, new(BackupOpIntegrationSuite))
}

func (suite *BackupOpIntegrationSuite) SetupSuite() {
	_, err := tester.GetRequiredEnvSls(
		tester.AWSStorageCredEnvs,
		tester.M365AcctCredEnvs)
	require.NoError(suite.T(), err)

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
		suite.T().Run(test.name, func(t *testing.T) {
			ctx, flush := tester.NewContext()
			defer flush()

			_, err := NewBackupOperation(
				ctx,
				test.opts,
				test.kw,
				test.sw,
				test.acct,
				selectors.Selector{},
				evmock.NewBus())
			test.errCheck(t, err)
		})
	}
}

// TestBackup_Run ensures that Integration Testing works
// for the following scopes: Contacts, Events, and Mail
func (suite *BackupOpIntegrationSuite) TestBackup_Run_exchange() {
	ctx, flush := tester.NewContext()
	defer flush()

	users := []string{suite.user}

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
				sel := selectors.NewExchangeBackup(users)
				sel.Include(sel.MailFolders(users, []string{exchange.DefaultMailFolder}, selectors.PrefixMatch()))
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
				sel := selectors.NewExchangeBackup(users)
				sel.Include(sel.ContactFolders(
					users,
					[]string{exchange.DefaultContactFolder},
					selectors.PrefixMatch()))

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
				sel := selectors.NewExchangeBackup(users)
				sel.Include(sel.EventCalendars(users, []string{exchange.DefaultCalendar}, selectors.PrefixMatch()))
				return sel
			},
			resourceOwner: suite.user,
			category:      path.EventsCategory,
			metadataFiles: exchange.MetadataFileNames(path.EventsCategory),
		},
	}
	for _, test := range tests {
		suite.T().Run(test.name, func(t *testing.T) {
			var (
				mb  = evmock.NewBus()
				sel = test.selector().Selector
				ffs = control.FeatureFlags{ExchangeIncrementals: test.runIncremental}
			)

			bo, acct, kw, ms, closer := prepNewBackupOp(t, ctx, mb, sel, ffs)
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
				test.category,
				test.metadataFiles,
			)

			if !test.runIncremental {
				return
			}

			// Basic, happy path incremental test.  No changes are dictated or expected.
			// This only tests that an incremental backup is runnable at all, and that it
			// produces fewer results than the last backup.
			var (
				incMB = evmock.NewBus()
				incBO = newBackupOp(t, ctx, kw, ms, acct, sel, incMB, ffs, closer)
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
				test.category,
				test.metadataFiles,
			)

			// do some additional checks to ensure the incremental dealt with fewer items.
			assert.Greater(t, bo.Results.ItemsWritten, incBO.Results.ItemsWritten, "incremental items written")
			assert.Greater(t, bo.Results.ItemsRead, incBO.Results.ItemsRead, "incremental items read")
			assert.Greater(t, bo.Results.BytesRead, incBO.Results.BytesRead, "incremental bytes read")
			assert.Greater(t, bo.Results.BytesUploaded, incBO.Results.BytesUploaded, "incremental bytes uploaded")
			assert.Equal(t, bo.Results.ResourceOwners, incBO.Results.ResourceOwners, "incremental backup resource owner")
			assert.NoError(t, incBO.Results.ReadErrors, "incremental read errors")
			assert.NoError(t, incBO.Results.WriteErrors, "incremental write errors")
			assert.Equal(t, 1, incMB.TimesCalled[events.BackupStart], "incremental backup-start events")
			assert.Equal(t, 1, incMB.TimesCalled[events.BackupEnd], "incremental backup-end events")
			assert.Equal(t,
				incMB.CalledWith[events.BackupStart][0][events.BackupID],
				incBO.Results.BackupID, "incremental backupID pre-declaration")
		})
	}
}

func (suite *BackupOpIntegrationSuite) TestBackup_Run_oneDrive() {
	ctx, flush := tester.NewContext()
	defer flush()

	var (
		t          = suite.T()
		mb         = evmock.NewBus()
		m365UserID = tester.SecondaryM365UserID(t)
		sel        = selectors.NewOneDriveBackup([]string{m365UserID})
	)

	sel.Include(sel.Users([]string{m365UserID}))

	bo, _, _, _, closer := prepNewBackupOp(t, ctx, mb, sel.Selector, control.FeatureFlags{})
	defer closer()

	runAndCheckBackup(t, ctx, &bo, mb)
}

func (suite *BackupOpIntegrationSuite) TestBackup_Run_sharePoint() {
	ctx, flush := tester.NewContext()
	defer flush()

	var (
		t   = suite.T()
		mb  = evmock.NewBus()
		sel = selectors.NewSharePointBackup([]string{suite.site})
	)

	sel.Include(sel.Sites([]string{suite.site}))

	bo, _, _, _, closer := prepNewBackupOp(t, ctx, mb, sel.Selector, control.FeatureFlags{})
	defer closer()

	runAndCheckBackup(t, ctx, &bo, mb)
}
