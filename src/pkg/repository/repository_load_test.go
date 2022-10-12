package repository_test

import (
	"context"
	"runtime/pprof"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/connector/exchange"
	"github.com/alcionai/corso/src/internal/operations"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/backup"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/repository"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/storage"
)

var users = []string{
	"AdeleV@8qzvrj.onmicrosoft.com",
	"AlexW@8qzvrj.onmicrosoft.com",
	"ashmarks@8qzvrj.onmicrosoft.com",
	"DiegoS@8qzvrj.onmicrosoft.com",
	"dustina@8qzvrj.onmicrosoft.com",
	"george.martinez@8qzvrj.onmicrosoft.com",
	"GradyA@8qzvrj.onmicrosoft.com",
	"HenriettaM@8qzvrj.onmicrosoft.com",
	"IsaiahL@8qzvrj.onmicrosoft.com",
	"JohannaL@8qzvrj.onmicrosoft.com",
	"JoniS@8qzvrj.onmicrosoft.com",
	"LeeG@8qzvrj.onmicrosoft.com",
	"LidiaH@8qzvrj.onmicrosoft.com",
	"LynneR@8qzvrj.onmicrosoft.com",
	"MeganB@8qzvrj.onmicrosoft.com",
	"MiriamG@8qzvrj.onmicrosoft.com",
	"NestorW@8qzvrj.onmicrosoft.com",
	"ntoja@8qzvrj.onmicrosoft.com",
	"PattiF@8qzvrj.onmicrosoft.com",
	"PradeepG@8qzvrj.onmicrosoft.com",
	"Rfinders@8qzvrj.onmicrosoft.com",
	"vkarma@8qzvrj.onmicrosoft.com",
	"greg.sanders@8qzvrj.onmicrosoft.com",
}

func initM365Repo(t *testing.T) (
	context.Context, repository.Repository, account.Account, storage.Storage,
) {
	_, err := tester.GetRequiredEnvSls(
		tester.AWSStorageCredEnvs,
		tester.M365AcctCredEnvs,
	)
	require.NoError(t, err)

	ctx, flush := tester.NewContext()
	defer flush()

	st := tester.NewPrefixedS3Storage(t)
	ac := tester.NewM365Account(t)
	opts := control.Options{
		DisableMetrics: true,
		FailFast:       true,
	}

	repo, err := repository.Initialize(ctx, ac, st, opts)
	require.NoError(t, err)

	return ctx, repo, ac, st
}

//revive:disable:context-as-argument
func runBackupLoadTest(
	t *testing.T,
	ctx context.Context,
	b *operations.BackupOperation,
	name string,
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

		require.NoError(t, err, "running backup")
		require.NotEmpty(t, b.Results, "has results after run")
		assert.NotEmpty(t, b.Results.BackupID, "has an ID after run")
		assert.Equal(t, b.Status, operations.Completed, "backup status")
		assert.Less(t, 0, b.Results.ItemsRead, "items read")
		assert.Less(t, 0, b.Results.ItemsWritten, "items written")
		assert.Less(t, int64(0), b.Results.BytesUploaded, "bytes uploaded")
		assert.Equal(t, len(users), b.Results.ResourceOwners, "resource owners")
		assert.Zero(t, b.Results.ReadErrors, "read errors")
		assert.Zero(t, b.Results.WriteErrors, "write errors")
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
			bs     []backup.Backup
			labels = pprof.Labels("list_load_test", name)
		)

		pprof.Do(ctx, labels, func(ctx context.Context) {
			bs, err = r.Backups(ctx)
		})

		require.NoError(t, err, "retrieving backups")
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
) {
	//revive:enable:context-as-argument
	require.NotEmpty(t, backupID, "backup ID to retrieve deails")

	t.Run("backup_details_"+name, func(t *testing.T) {
		var (
			err    error
			b      *backup.Backup
			ds     *details.Details
			labels = pprof.Labels("details_load_test", name)
		)

		pprof.Do(ctx, labels, func(ctx context.Context) {
			ds, b, err = r.BackupDetails(ctx, backupID)
		})

		require.NoError(t, err, "retrieving details in backup "+backupID)
		require.NotNil(t, ds, "backup details must exist")
		require.NotNil(t, b, "backup must exist")

		assert.Equal(t,
			b.ItemsWritten, len(noFolders(t, ds.Entries)),
			"items written to backup must match the count of entries, minus folder entries")

		ensureAllUsersInDetails(t, ds, "backup", name)
	})
}

//revive:disable:context-as-argument
func runRestoreLoadTest(
	t *testing.T,
	ctx context.Context,
	r operations.RestoreOperation,
	name string,
	expectItemCount int,
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

		require.NoError(t, err, "running restore")
		require.NotEmpty(t, r.Results, "has results after run")
		require.NotNil(t, ds, "has restored details")
		assert.Equal(t, r.Status, operations.Completed, "restore status")
		assert.Equal(t, r.Results.ItemsWritten, len(ds.Entries), "count of items written matches restored entries in details")
		assert.Less(t, 0, r.Results.ItemsRead, "items read")
		assert.Less(t, 0, r.Results.ItemsWritten, "items written")
		assert.Less(t, 0, r.Results.ResourceOwners, "resource owners")
		assert.Zero(t, r.Results.ReadErrors, "read errors")
		assert.Zero(t, r.Results.WriteErrors, "write errors")
		assert.Equal(t, expectItemCount, r.Results.ItemsWritten, "backup and restore wrote the same count of items")

		ensureAllUsersInDetails(t, ds, "restore", name)
	})
}

// noFolders removes all "folder" category details entries
func noFolders(t *testing.T, des []details.DetailsEntry) []details.DetailsEntry {
	t.Helper()

	sansfldr := []details.DetailsEntry{}

	for _, ent := range des {
		if ent.Folder == nil {
			sansfldr = append(sansfldr, ent)
		}
	}

	return sansfldr
}

func ensureAllUsersInDetails(
	t *testing.T,
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
			foundUsers[u] = false
			userCategories[u] = map[string]struct{}{}
		}

		for _, ent := range noFolders(t, ds.Entries) {
			p, err := path.FromDataLayerPath(ent.RepoRef, true)
			require.NoError(t, err, "converting to path: "+ent.RepoRef)

			ro := p.ResourceOwner()
			assert.NotEmpty(t, ro, "resource owner in path: "+ent.RepoRef)

			ct := p.Category()
			assert.NotEmpty(t, ro, "category type of path: "+ent.RepoRef)

			foundUsers[ro] = true
			foundCategories[ct.String()] = struct{}{}

			if _, ok := userCategories[ro]; !ok {
				userCategories[ro] = map[string]struct{}{}
			}

			userCategories[ro][ct.String()] = struct{}{}
		}

		for u, cats := range userCategories {
			t.Run(u, func(t *testing.T) {
				assert.True(t, foundUsers[u], "user was involved in operation")
				assert.Equal(t, len(foundCategories), len(cats), "all app categories involved in operation")
			})
		}
	})
}

// ------------------------------------------------------------------------------------------------
// Exchange
// ------------------------------------------------------------------------------------------------

type RepositoryLoadTestExchangeSuite struct {
	suite.Suite
	ctx  context.Context
	repo repository.Repository
	acct account.Account
	st   storage.Storage
}

func TestRepositoryLoadTestExchangeSuite(t *testing.T) {
	if err := tester.RunOnAny(tester.CorsoLoadTests); err != nil {
		t.Skip(err)
	}

	suite.Run(t, new(RepositoryLoadTestExchangeSuite))
}

func (suite *RepositoryLoadTestExchangeSuite) SetupSuite() {
	t := suite.T()
	t.Parallel()
	suite.ctx, suite.repo, suite.acct, suite.st = initM365Repo(t)
}

func (suite *RepositoryLoadTestExchangeSuite) TeardownSuite() {
	suite.repo.Close(suite.ctx)
}

func (suite *RepositoryLoadTestExchangeSuite) TestExchange() {
	ctx, flush := tester.NewContext()
	defer flush()

	var (
		t       = suite.T()
		r       = suite.repo
		service = "exchange"
	)

	// backup
	bsel := selectors.NewExchangeBackup()
	bsel.Include(bsel.MailFolders(users, []string{exchange.DefaultMailFolder}))
	bsel.Include(bsel.ContactFolders(users, []string{exchange.DefaultContactFolder}))
	bsel.Include(bsel.EventCalendars(users, []string{exchange.DefaultCalendar}))

	b, err := r.NewBackup(ctx, bsel.Selector)
	require.NoError(t, err)

	runBackupLoadTest(t, ctx, &b, service)
	bid := string(b.Results.BackupID)

	runBackupListLoadTest(t, ctx, r, service, bid)
	runBackupDetailsLoadTest(t, ctx, r, service, bid)

	// restore
	rsel, err := bsel.ToExchangeRestore()
	require.NoError(t, err)

	dest := tester.DefaultTestRestoreDestination()

	rst, err := r.NewRestore(ctx, bid, rsel.Selector, dest)
	require.NoError(t, err)

	runRestoreLoadTest(t, ctx, rst, service, b.Results.ItemsWritten)
}

// ------------------------------------------------------------------------------------------------
// OneDrive
// ------------------------------------------------------------------------------------------------

type RepositoryLoadTestOneDriveSuite struct {
	suite.Suite
	ctx  context.Context
	repo repository.Repository
	acct account.Account
	st   storage.Storage
}

func TestRepositoryLoadTestOneDriveSuite(t *testing.T) {
	if err := tester.RunOnAny(tester.CorsoLoadTests); err != nil {
		t.Skip(err)
	}

	suite.Run(t, new(RepositoryLoadTestOneDriveSuite))
}

func (suite *RepositoryLoadTestOneDriveSuite) SetupSuite() {
	t := suite.T()
	t.Skip("temp issue-902-live")
	t.Parallel()
	suite.ctx, suite.repo, suite.acct, suite.st = initM365Repo(t)
}

func (suite *RepositoryLoadTestOneDriveSuite) TeardownSuite() {
	suite.repo.Close(suite.ctx)
}

func (suite *RepositoryLoadTestOneDriveSuite) TestOneDrive() {
	ctx, flush := tester.NewContext()
	defer flush()

	var (
		t       = suite.T()
		r       = suite.repo
		service = "one_drive"
	)

	// backup
	bsel := selectors.NewOneDriveBackup()
	bsel.Include(bsel.Users(users))

	b, err := r.NewBackup(ctx, bsel.Selector)
	require.NoError(t, err)

	runBackupLoadTest(t, ctx, &b, service)
	bid := string(b.Results.BackupID)

	runBackupListLoadTest(t, ctx, r, service, bid)
	runBackupDetailsLoadTest(t, ctx, r, service, bid)

	// restore
	rsel, err := bsel.ToOneDriveRestore()
	require.NoError(t, err)

	dest := tester.DefaultTestRestoreDestination()

	rst, err := r.NewRestore(ctx, bid, rsel.Selector, dest)
	require.NoError(t, err)

	runRestoreLoadTest(t, ctx, rst, service, b.Results.ItemsWritten)
}
