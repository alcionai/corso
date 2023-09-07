package repository_test

import (
	"context"
	"fmt"
	"os"
	"runtime/pprof"
	"sort"
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	D "github.com/alcionai/corso/src/internal/diagnostics"
	"github.com/alcionai/corso/src/internal/operations"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/backup"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	ctrlRepo "github.com/alcionai/corso/src/pkg/control/repository"
	ctrlTD "github.com/alcionai/corso/src/pkg/control/testdata"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/repository"
	"github.com/alcionai/corso/src/pkg/selectors"
	selTD "github.com/alcionai/corso/src/pkg/selectors/testdata"
	"github.com/alcionai/corso/src/pkg/storage"
	storeTD "github.com/alcionai/corso/src/pkg/storage/testdata"
)

//lint:ignore U1000 future test use
func orgSiteSet(t *testing.T) []string {
	return tconfig.LoadTestM365OrgSites(t)
}

//lint:ignore U1000 future test use
func orgUserSet(t *testing.T) []string {
	return tconfig.LoadTestM365OrgUsers(t)
}

//lint:ignore U1000 future test use
func singleSiteSet(t *testing.T) []string {
	return []string{tconfig.LoadTestM365SiteID(t)}
}

//lint:ignore U1000 future test use
func singleUserSet(t *testing.T) []string {
	return []string{tconfig.LoadTestM365UserID(t)}
}

var loadCtx context.Context

func TestMain(m *testing.M) {
	if len(os.Getenv(tester.CorsoLoadTests)) == 0 {
		return
	}

	ctx, logFlush := tester.NewContext(nil)
	loadCtx = ctx

	if err := D.InitCollector(); err != nil {
		fmt.Println("initializing load tests:", err)
		os.Exit(1)
	}

	ctx, spanFlush := D.Start(ctx, "Load_Testing_Main")
	loadCtx = ctx
	flush := func() {
		spanFlush()
		logFlush()
	}

	exitVal := m.Run()

	flush()

	os.Exit(exitVal)
}

// ------------------------------------------------------------------------------------------------
// Common
// ------------------------------------------------------------------------------------------------

func initM365Repo(t *testing.T) (
	context.Context,
	repository.Repository,
	account.Account,
	storage.Storage,
) {
	tester.MustGetEnvSets(t, storeTD.AWSStorageCredEnvs, tconfig.M365AcctCredEnvs)

	ctx, flush := tester.WithContext(t, loadCtx)
	defer flush()

	st := storeTD.NewPrefixedS3Storage(t)
	ac := tconfig.NewM365Account(t)
	opts := control.Options{
		DisableMetrics:  true,
		FailureHandling: control.FailFast,
	}

	repo, err := repository.Initialize(ctx, ac, st, opts, ctrlRepo.Retention{})
	require.NoError(t, err, clues.ToCore(err))

	return ctx, repo, ac, st
}

// ------------------------------------------------------------------------------------------------
// Common
// ------------------------------------------------------------------------------------------------

//revive:disable:context-as-argument
func runLoadTest(
	t *testing.T,
	ctx context.Context,
	r repository.Repository,
	prefix, service string,
	usersUnderTest []string,
	bupSel, restSel selectors.Selector,
	runRestore bool,
) {
	//revive:enable:context-as-argument
	t.Run(prefix+"_load_test_main", func(t *testing.T) {
		b, err := r.NewBackup(ctx, bupSel)
		require.NoError(t, err, clues.ToCore(err))

		runBackupLoadTest(t, ctx, &b, service, usersUnderTest)
		bid := string(b.Results.BackupID)

		runBackupListLoadTest(t, ctx, r, service, bid)
		runBackupDetailsLoadTest(t, ctx, r, service, bid, usersUnderTest)

		runRestoreLoadTest(t, ctx, r, prefix, service, bid, usersUnderTest, restSel, b, runRestore)
	})
}

//revive:disable:context-as-argument
func runRestoreLoadTest(
	t *testing.T,
	ctx context.Context,
	r repository.Repository,
	prefix, service, backupID string,
	usersUnderTest []string,
	restSel selectors.Selector,
	bup operations.BackupOperation,
	runRestore bool,
) {
	//revive:enable:context-as-argument
	t.Run(prefix+"_load_test_restore", func(t *testing.T) {
		if !runRestore {
			t.Skip("restore load test is toggled off")
		}

		restoreCfg := ctrlTD.DefaultRestoreConfig("")

		rst, err := r.NewRestore(ctx, backupID, restSel, restoreCfg)
		require.NoError(t, err, clues.ToCore(err))

		doRestoreLoadTest(t, ctx, rst, service, bup.Results.ItemsWritten, usersUnderTest)
	})
}

//revive:disable:context-as-argument
func runBackupLoadTest(
	t *testing.T,
	ctx context.Context,
	b *operations.BackupOperation,
	name string,
	users []string,
) {
	//revive:enable:context-as-argument
	t.Run("backup_"+name, func(t *testing.T) {
		var (
			err    error
			labels = pprof.Labels("backup_load_test", name)
		)

		pprof.Do(ctx, labels, func(ctx context.Context) {
			err = b.Run(ctx)
		})

		require.NoError(t, err, "running backup", clues.ToCore(err))
		require.NotEmpty(t, b.Results, "has results after run")
		assert.NotEmpty(t, b.Results.BackupID, "has an ID after run")
		assert.Equal(t, b.Status, operations.Completed, "backup status")
		assert.Less(t, 0, b.Results.ItemsRead, "items read")
		assert.Less(t, 0, b.Results.ItemsWritten, "items written")
		assert.Less(t, int64(0), b.Results.BytesUploaded, "bytes uploaded")
		assert.Equal(t, len(users), b.Results.ResourceOwners, "resource owners")
		assert.NoError(t, b.Errors.Failure(), "non-recoverable error", clues.ToCore(b.Errors.Failure()))
		assert.Empty(t, b.Errors.Recovered(), "recoverable errors")
	})
}

//revive:disable:context-as-argument
func runBackupListLoadTest(
	t *testing.T,
	ctx context.Context,
	r repository.Repository,
	name, expectID string,
) {
	//revive:enable:context-as-argument
	t.Run("backup_list_"+name, func(t *testing.T) {
		var (
			err    error
			bs     []*backup.Backup
			labels = pprof.Labels("list_load_test", name)
		)

		pprof.Do(ctx, labels, func(ctx context.Context) {
			bs, err = r.BackupsByTag(ctx)
		})

		require.NoError(t, err, "retrieving backups", clues.ToCore(err))
		require.Less(t, 0, len(bs), "at least one backup is recorded")

		var found bool

		for _, b := range bs {
			bid := b.ID
			assert.NotEmpty(t, bid, "iterating backup ids")

			if string(bid) == expectID {
				found = true
			}
		}

		assert.True(t, found, "expected backup id "+expectID+" found in backups")
	})
}

//revive:disable:context-as-argument
func runBackupDetailsLoadTest(
	t *testing.T,
	ctx context.Context,
	r repository.Repository,
	name, backupID string,
	users []string,
) {
	//revive:enable:context-as-argument
	require.NotEmpty(t, backupID, "backup ID to retrieve deails")

	t.Run("backup_details_"+name, func(t *testing.T) {
		var (
			errs   *fault.Bus
			b      *backup.Backup
			ds     *details.Details
			labels = pprof.Labels("details_load_test", name)
		)

		pprof.Do(ctx, labels, func(ctx context.Context) {
			ds, b, errs = r.GetBackupDetails(ctx, backupID)
		})

		require.NoError(t, errs.Failure(), "retrieving details in backup", backupID, clues.ToCore(errs.Failure()))
		require.Empty(t, errs.Recovered(), "retrieving details in backup", backupID)
		require.NotNil(t, ds, "backup details must exist")
		require.NotNil(t, b, "backup must exist")

		assert.Equal(t,
			b.ItemsWritten, len(noFolders(t, ds.Entries)),
			"items written to backup must match the count of entries, minus folder entries")

		ensureAllUsersInDetails(t, users, ds, "backup", name)
	})
}

//revive:disable:context-as-argument
func doRestoreLoadTest(
	t *testing.T,
	ctx context.Context,
	r operations.RestoreOperation,
	name string,
	expectItemCount int,
	users []string,
) {
	//revive:enable:context-as-argument
	t.Run("restore_"+name, func(t *testing.T) {
		var (
			err    error
			ds     *details.Details
			labels = pprof.Labels("restore_load_test", name)
		)

		pprof.Do(ctx, labels, func(ctx context.Context) {
			ds, err = r.Run(ctx)
		})

		require.NoError(t, err, "running restore", clues.ToCore(err))
		require.NotEmpty(t, r.Results, "has results after run")
		require.NotNil(t, ds, "has restored details")
		assert.Equal(t, r.Status, operations.Completed, "restore status")
		assert.Equal(t, r.Results.ItemsWritten, len(ds.Entries), "count of items written matches restored entries in details")
		assert.Less(t, 0, r.Results.ItemsRead, "items read")
		assert.Less(t, 0, r.Results.ItemsWritten, "items written")
		assert.Equal(t, len(users), r.Results.ResourceOwners, "resource owners")
		assert.NoError(t, r.Errors.Failure(), "non-recoverable error")
		assert.Empty(t, r.Errors.Recovered(), "recoverable errors")
		assert.Equal(t, expectItemCount, r.Results.ItemsWritten, "backup and restore wrote the same count of items")

		ensureAllUsersInDetails(t, users, ds, "restore", name)
	})
}

// noFolders removes all "folder" category details entries
func noFolders(t *testing.T, des []details.Entry) []details.Entry {
	t.Helper()

	sansfldr := []details.Entry{}

	for _, ent := range des {
		if ent.Folder == nil {
			sansfldr = append(sansfldr, ent)
		}
	}

	return sansfldr
}

func ensureAllUsersInDetails(
	t *testing.T,
	users []string,
	ds *details.Details,
	prefix, name string,
) {
	t.Run("details_"+prefix+"_"+name, func(t *testing.T) {
		// assert that all users backed up at least one item of each category.
		var (
			foundUsers      = map[string]bool{}
			foundCategories = map[string]struct{}{}
			userCategories  = map[string]map[string]struct{}{}
		)

		for _, u := range users {
			userCategories[u] = map[string]struct{}{}
		}

		for _, ent := range noFolders(t, ds.Entries) {
			e := ent
			rr := e.RepoRef

			p, err := path.FromDataLayerPath(rr, true)
			if !assert.NoError(t, err, "converting to path: "+rr) {
				continue
			}

			ro := p.ProtectedResource()
			if !assert.NotEmpty(t, ro, "resource owner in path: "+rr) {
				continue
			}

			ct := p.Category()
			if !assert.NotEmpty(t, ro, "category type of path: "+rr) {
				continue
			}

			foundUsers[ro] = true
			foundCategories[ct.String()] = struct{}{}
			userCategories[ro][ct.String()] = struct{}{}
		}

		foundCategoriesSl := normalizeCategorySet(t, foundCategories)

		for u, cats := range userCategories {
			userCategoriesSl := normalizeCategorySet(t, cats)

			t.Run(u, func(t *testing.T) {
				assert.True(t, foundUsers[u], "user was involved in operation")
				assert.EqualValues(t, foundCategoriesSl, userCategoriesSl, "user had all app categories involved in operation")
			})
		}
	})
}

// for an easier time comparing category presence/absence
func normalizeCategorySet(t *testing.T, cats map[string]struct{}) []string {
	t.Helper()

	sl := []string{}
	for k := range cats {
		sl = append(sl, k)
	}

	sort.Strings(sl)

	return sl
}

/* ================================================
*   A note on load test setup:
*   Even though most of the code here is boiler-
*   plate and could be easily compressed into a
*   test matrix, we want to keep the suites separate
*   to maximize parallelism.  Due to how testify's
*   suites work, we can only run in parallel at the
*   level of the suite, not within each test.
*  ================================================ */

// ------------------------------------------------------------------------------------------------
// Exchange
// ------------------------------------------------------------------------------------------------

// multiple users

type LoadExchangeSuite struct {
	tester.Suite
	ctx            context.Context
	repo           repository.Repository
	acct           account.Account //lint:ignore U1000 future test use
	st             storage.Storage //lint:ignore U1000 future test use
	usersUnderTest []string
}

func TestLoadExchangeSuite(t *testing.T) {
	suite.Run(t, &LoadExchangeSuite{
		Suite: tester.NewLoadSuite(
			t,
			[][]string{storeTD.AWSStorageCredEnvs, tconfig.M365AcctCredEnvs},
		),
	})
}

func (suite *LoadExchangeSuite) SetupSuite() {
	t := suite.T()
	t.Parallel()
	suite.ctx, suite.repo, suite.acct, suite.st = initM365Repo(t)
	suite.usersUnderTest = orgUserSet(t)
}

func (suite *LoadExchangeSuite) TeardownSuite() {
	suite.repo.Close(suite.ctx)
}

func (suite *LoadExchangeSuite) TestExchange() {
	ctx, flush := tester.WithContext(suite.T(), suite.ctx)
	defer flush()

	bsel := selectors.NewExchangeBackup(suite.usersUnderTest)
	bsel.Include(bsel.MailFolders(selectors.Any()))
	bsel.Include(bsel.ContactFolders(selectors.Any()))
	bsel.Include(bsel.EventCalendars(selectors.Any()))
	sel := bsel.Selector

	runLoadTest(
		suite.T(),
		ctx,
		suite.repo,
		"all_users", "exchange",
		suite.usersUnderTest,
		sel, sel, // same selection for backup and restore
		true)
}

// single user, lots of data

type IndividualLoadExchangeSuite struct {
	tester.Suite
	ctx            context.Context
	repo           repository.Repository
	acct           account.Account //lint:ignore U1000 future test use
	st             storage.Storage //lint:ignore U1000 future test use
	usersUnderTest []string
}

func TestIndividualLoadExchangeSuite(t *testing.T) {
	suite.Run(t, &IndividualLoadExchangeSuite{
		Suite: tester.NewLoadSuite(
			t,
			[][]string{storeTD.AWSStorageCredEnvs, tconfig.M365AcctCredEnvs},
		),
	})
}

func (suite *IndividualLoadExchangeSuite) SetupSuite() {
	t := suite.T()
	t.Skip("individual user exchange suite tests are on hold until token expiry gets resolved")
	t.Parallel()
	suite.ctx, suite.repo, suite.acct, suite.st = initM365Repo(t)
	suite.usersUnderTest = singleUserSet(t)
}

func (suite *IndividualLoadExchangeSuite) TeardownSuite() {
	suite.repo.Close(suite.ctx)
}

func (suite *IndividualLoadExchangeSuite) TestExchange() {
	ctx, flush := tester.WithContext(suite.T(), suite.ctx)
	defer flush()

	bsel := selectors.NewExchangeBackup(suite.usersUnderTest)
	bsel.Include(bsel.MailFolders(selectors.Any()))
	bsel.Include(bsel.ContactFolders(selectors.Any()))
	bsel.Include(bsel.EventCalendars(selectors.Any()))
	sel := bsel.Selector

	runLoadTest(
		suite.T(),
		ctx,
		suite.repo,
		"single_user", "exchange",
		suite.usersUnderTest,
		sel, sel, // same selection for backup and restore
		true)
}

// ------------------------------------------------------------------------------------------------
// OneDrive
// ------------------------------------------------------------------------------------------------

type LoadOneDriveSuite struct {
	tester.Suite
	ctx            context.Context
	repo           repository.Repository
	acct           account.Account //lint:ignore U1000 future test use
	st             storage.Storage //lint:ignore U1000 future test use
	usersUnderTest []string
}

func TestLoadOneDriveSuite(t *testing.T) {
	suite.Run(t, &LoadOneDriveSuite{
		Suite: tester.NewLoadSuite(
			t,
			[][]string{storeTD.AWSStorageCredEnvs, tconfig.M365AcctCredEnvs},
		),
	})
}

func (suite *LoadOneDriveSuite) SetupSuite() {
	t := suite.T()
	t.Skip("not running onedrive load tests atm")
	t.Parallel()
	suite.ctx, suite.repo, suite.acct, suite.st = initM365Repo(t)
	suite.usersUnderTest = orgUserSet(t)
}

func (suite *LoadOneDriveSuite) TeardownSuite() {
	suite.repo.Close(suite.ctx)
}

func (suite *LoadOneDriveSuite) TestOneDrive() {
	ctx, flush := tester.WithContext(suite.T(), suite.ctx)
	defer flush()

	bsel := selectors.NewOneDriveBackup(suite.usersUnderTest)
	bsel.Include(selTD.OneDriveBackupFolderScope(bsel))
	sel := bsel.Selector

	runLoadTest(
		suite.T(),
		ctx,
		suite.repo,
		"all_users", "one_drive",
		suite.usersUnderTest,
		sel, sel, // same selection for backup and restore
		false)
}

type IndividualLoadOneDriveSuite struct {
	tester.Suite
	ctx            context.Context
	repo           repository.Repository
	acct           account.Account //lint:ignore U1000 future test use
	st             storage.Storage //lint:ignore U1000 future test use
	usersUnderTest []string
}

func TestIndividualLoadOneDriveSuite(t *testing.T) {
	suite.Run(t, &IndividualLoadOneDriveSuite{
		Suite: tester.NewLoadSuite(
			t,
			[][]string{storeTD.AWSStorageCredEnvs, tconfig.M365AcctCredEnvs},
		),
	})
}

func (suite *IndividualLoadOneDriveSuite) SetupSuite() {
	t := suite.T()
	t.Parallel()
	suite.ctx, suite.repo, suite.acct, suite.st = initM365Repo(t)
	suite.usersUnderTest = singleUserSet(t)
}

func (suite *IndividualLoadOneDriveSuite) TeardownSuite() {
	suite.repo.Close(suite.ctx)
}

func (suite *IndividualLoadOneDriveSuite) TestOneDrive() {
	ctx, flush := tester.WithContext(suite.T(), suite.ctx)
	defer flush()

	bsel := selectors.NewOneDriveBackup(suite.usersUnderTest)
	bsel.Include(selTD.OneDriveBackupFolderScope(bsel))
	sel := bsel.Selector

	runLoadTest(
		suite.T(),
		ctx,
		suite.repo,
		"single_user", "one_drive",
		suite.usersUnderTest,
		sel, sel, // same selection for backup and restore
		false)
}

// ------------------------------------------------------------------------------------------------
// SharePoint
// ------------------------------------------------------------------------------------------------

type LoadSharePointSuite struct {
	tester.Suite
	ctx            context.Context
	repo           repository.Repository
	acct           account.Account //lint:ignore U1000 future test use
	st             storage.Storage //lint:ignore U1000 future test use
	sitesUnderTest []string
}

func TestLoadSharePointSuite(t *testing.T) {
	suite.Run(t, &LoadSharePointSuite{
		Suite: tester.NewLoadSuite(
			t,
			[][]string{storeTD.AWSStorageCredEnvs, tconfig.M365AcctCredEnvs},
		),
	})
}

func (suite *LoadSharePointSuite) SetupSuite() {
	t := suite.T()
	t.Skip("not running sharepoint load tests atm")
	t.Parallel()
	suite.ctx, suite.repo, suite.acct, suite.st = initM365Repo(t)
	suite.sitesUnderTest = orgSiteSet(t)
}

func (suite *LoadSharePointSuite) TeardownSuite() {
	suite.repo.Close(suite.ctx)
}

func (suite *LoadSharePointSuite) TestSharePoint() {
	ctx, flush := tester.WithContext(suite.T(), suite.ctx)
	defer flush()

	bsel := selectors.NewSharePointBackup(suite.sitesUnderTest)
	bsel.Include(bsel.AllData())
	sel := bsel.Selector

	runLoadTest(
		suite.T(),
		ctx,
		suite.repo,
		"all_sites", "share_point",
		suite.sitesUnderTest,
		sel, sel, // same selection for backup and restore
		false)
}

type IndividualLoadSharePointSuite struct {
	tester.Suite
	ctx            context.Context
	repo           repository.Repository
	acct           account.Account //lint:ignore U1000 future test use
	st             storage.Storage //lint:ignore U1000 future test use
	sitesUnderTest []string
}

func TestIndividualLoadSharePointSuite(t *testing.T) {
	suite.Run(t, &IndividualLoadSharePointSuite{
		Suite: tester.NewLoadSuite(
			t,
			[][]string{storeTD.AWSStorageCredEnvs, tconfig.M365AcctCredEnvs},
		),
	})
}

func (suite *IndividualLoadSharePointSuite) SetupSuite() {
	t := suite.T()
	t.Skip("not running sharepoint load tests atm")
	t.Parallel()
	suite.ctx, suite.repo, suite.acct, suite.st = initM365Repo(t)
	suite.sitesUnderTest = singleSiteSet(t)
}

func (suite *IndividualLoadSharePointSuite) TeardownSuite() {
	suite.repo.Close(suite.ctx)
}

func (suite *IndividualLoadSharePointSuite) TestSharePoint() {
	ctx, flush := tester.WithContext(suite.T(), suite.ctx)
	defer flush()

	bsel := selectors.NewSharePointBackup(suite.sitesUnderTest)
	bsel.Include(bsel.AllData())
	sel := bsel.Selector

	runLoadTest(
		suite.T(),
		ctx,
		suite.repo,
		"single_site", "share_point",
		suite.sitesUnderTest,
		sel, sel, // same selection for backup and restore
		false)
}
