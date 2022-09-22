package repository_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	evmock "github.com/alcionai/corso/src/internal/events/mock"
	"github.com/alcionai/corso/src/internal/operations"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/repository"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/storage"
)

func initM365Repo(t *testing.T) (
	context.Context, repository.Repository, account.Account, storage.Storage,
) {
	_, err := tester.GetRequiredEnvSls(
		tester.AWSStorageCredEnvs,
		tester.M365AcctCredEnvs,
	)
	require.NoError(t, err)

	ctx := tester.NewContext()
	st := tester.NewPrefixedS3Storage(t)
	ac := tester.NewM365Account(t)
	opts := control.Options{
		DisableMetrics: true,
		FailFast:       true,
	}

	repo, err := repository.Initialize(ctx, ac, st, opts)
	require.NoError(t, err)

	repo.Bus = evmock.NewBus()

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
		require.NoError(t, b.Run(ctx), "running backup")
		require.NotEmpty(t, b.Results, "has results after run")
		assert.NotEmpty(t, b.Results.BackupID, "has an ID after run")
		assert.Equal(t, b.Status, operations.Completed, "backup status")
		assert.Less(t, 0, b.Results.ItemsRead, "items read")
		assert.Less(t, 0, b.Results.ItemsWritten, "items written")
		assert.Less(t, int64(0), b.Results.BytesWritten, "bytes written")
		assert.Less(t, 0, b.Results.ResourceOwners, "resource owners")
		assert.Zero(t, b.Results.ReadErrors, "read errors")
		assert.Zero(t, b.Results.WriteErrors, "write errors")
	})
}

//revive:disable:context-as-argument
func runBackupListLoadTest(
	t *testing.T,
	ctx context.Context,
	r *repository.Repository,
	name, expectID string,
) {
	//revive:enable:context-as-argument
	t.Run("backup_list_"+name, func(t *testing.T) {
		bs, err := r.Backups(ctx)
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
	r *repository.Repository,
	name, backupID string,
) {
	//revive:enable:context-as-argument
	require.NotEmpty(t, backupID, "backup ID to retrieve deails")

	t.Run("backup_details_"+name, func(t *testing.T) {
		ds, b, err := r.BackupDetails(ctx, backupID)
		require.NoError(t, err, "retrieving details in backup "+backupID)
		require.NotNil(t, ds, "backup details")
		require.NotNil(t, b, "backup")
		assert.Equal(t, b.ItemsWritten, len(ds.Entries))
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
		require.NoError(t, r.Run(ctx), "running restore")
		require.NotEmpty(t, r.Results, "has results after run")
		assert.Equal(t, r.Status, operations.Completed, "restore status")
		assert.Less(t, 0, r.Results.ItemsRead, "items read")
		assert.Less(t, 0, r.Results.ItemsWritten, "items written")
		assert.Less(t, 0, r.Results.ResourceOwners, "resource owners")
		assert.Zero(t, r.Results.ReadErrors, "read errors")
		assert.Zero(t, r.Results.WriteErrors, "write errors")
		assert.Equal(t, expectItemCount, r.Results.ItemsWritten, "backup and restore wrote the same count of items")
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
	suite.ctx, suite.repo, suite.acct, suite.st = initM365Repo(t)
}

func (suite *RepositoryLoadTestExchangeSuite) TeardownSuite() {
	suite.repo.Close(suite.ctx)
}

func (suite *RepositoryLoadTestExchangeSuite) SetupTest() {
	suite.ctx, _ = logger.SeedLevel(context.Background(), logger.Development)
}

func (suite *RepositoryLoadTestExchangeSuite) TeardownTest() {
	logger.Flush(suite.ctx)
}

func (suite *RepositoryLoadTestExchangeSuite) TestExchange() {
	var (
		t    = suite.T()
		ctx  = context.Background()
		r    = suite.repo
		serv = "exchange"
	)

	t.Parallel()

	// backup
	bsel := selectors.NewExchangeBackup()
	bsel.Include(bsel.Users(selectors.Any()))

	b, err := r.NewBackup(ctx, bsel.Selector)
	require.NoError(t, err)

	runBackupLoadTest(t, ctx, &b, serv)
	bid := string(b.Results.BackupID)

	runBackupListLoadTest(t, ctx, r, serv, bid)
	runBackupDetailsLoadTest(t, ctx, r, serv, bid)

	// restore
	rsel, err := bsel.ToExchangeRestore()
	require.NoError(t, err)

	rst, err := r.NewRestore(ctx, bid, rsel.Selector)
	require.NoError(t, err)

	runRestoreLoadTest(t, ctx, rst, serv, b.Results.ItemsWritten)
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
	suite.ctx, suite.repo, suite.acct, suite.st = initM365Repo(t)
}

func (suite *RepositoryLoadTestOneDriveSuite) TeardownSuite() {
	suite.repo.Close(suite.ctx)
}

func (suite *RepositoryLoadTestOneDriveSuite) SetupTest() {
	suite.ctx, _ = logger.SeedLevel(context.Background(), logger.Development)
}

func (suite *RepositoryLoadTestOneDriveSuite) TeardownTest() {
	logger.Flush(suite.ctx)
}

func (suite *RepositoryLoadTestOneDriveSuite) TestOneDrive() {
	var (
		t    = suite.T()
		ctx  = context.Background()
		r    = suite.repo
		serv = "one_drive"
	)

	t.Parallel()

	// backup
	bsel := selectors.NewOneDriveBackup()
	bsel.Include(bsel.Users(selectors.Any()))

	b, err := r.NewBackup(ctx, bsel.Selector)
	require.NoError(t, err)

	runBackupLoadTest(t, ctx, &b, serv)
	bid := string(b.Results.BackupID)

	runBackupListLoadTest(t, ctx, r, serv, bid)
	runBackupDetailsLoadTest(t, ctx, r, serv, bid)

	// restore
	rsel, err := bsel.ToOneDriveRestore()
	require.NoError(t, err)

	rst, err := r.NewRestore(ctx, bid, rsel.Selector)
	require.NoError(t, err)

	runRestoreLoadTest(t, ctx, rst, serv, b.Results.ItemsWritten)
}
