package operations

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/connector/exchange"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/events"
	evmock "github.com/alcionai/corso/src/internal/events/mock"
	"github.com/alcionai/corso/src/internal/kopia"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/store"
)

// ---------------------------------------------------------------------------
// unit
// ---------------------------------------------------------------------------

type BackupOpSuite struct {
	suite.Suite
}

func TestBackupOpSuite(t *testing.T) {
	suite.Run(t, new(BackupOpSuite))
}

func (suite *BackupOpSuite) TestBackupOperation_PersistResults() {
	ctx, flush := tester.NewContext()
	defer flush()

	var (
		kw   = &kopia.Wrapper{}
		sw   = &store.Wrapper{}
		acct = account.Account{}
		now  = time.Now()
	)

	table := []struct {
		expectStatus opStatus
		expectErr    assert.ErrorAssertionFunc
		stats        backupStats
	}{
		{
			expectStatus: Completed,
			expectErr:    assert.NoError,
			stats: backupStats{
				started:       true,
				resourceCount: 1,
				k: &kopia.BackupStats{
					TotalFileCount:     1,
					TotalHashedBytes:   1,
					TotalUploadedBytes: 1,
				},
				gc: &support.ConnectorOperationStatus{
					Successful: 1,
				},
			},
		},
		{
			expectStatus: Failed,
			expectErr:    assert.Error,
			stats: backupStats{
				started: false,
				k:       &kopia.BackupStats{},
				gc:      &support.ConnectorOperationStatus{},
			},
		},
		{
			expectStatus: NoData,
			expectErr:    assert.NoError,
			stats: backupStats{
				started: true,
				k:       &kopia.BackupStats{},
				gc:      &support.ConnectorOperationStatus{},
			},
		},
	}
	for _, test := range table {
		suite.T().Run(test.expectStatus.String(), func(t *testing.T) {
			op, err := NewBackupOperation(
				ctx,
				control.Options{},
				kw,
				sw,
				acct,
				selectors.Selector{},
				evmock.NewBus())
			require.NoError(t, err)
			test.expectErr(t, op.persistResults(now, &test.stats))

			assert.Equal(t, test.expectStatus.String(), op.Status.String(), "status")
			assert.Equal(t, test.stats.gc.Successful, op.Results.ItemsRead, "items read")
			assert.Equal(t, test.stats.readErr, op.Results.ReadErrors, "read errors")
			assert.Equal(t, test.stats.k.TotalFileCount, op.Results.ItemsWritten, "items written")
			assert.Equal(t, test.stats.k.TotalHashedBytes, op.Results.BytesRead, "bytes read")
			assert.Equal(t, test.stats.k.TotalUploadedBytes, op.Results.BytesUploaded, "bytes written")
			assert.Equal(t, test.stats.resourceCount, op.Results.ResourceOwners, "resource owners")
			assert.Equal(t, test.stats.writeErr, op.Results.WriteErrors, "write errors")
			assert.Equal(t, now, op.Results.StartedAt, "started at")
			assert.Less(t, now, op.Results.CompletedAt, "completed at")
		})
	}
}

// ---------------------------------------------------------------------------
// integration
// ---------------------------------------------------------------------------

type BackupOpIntegrationSuite struct {
	suite.Suite
}

func TestBackupOpIntegrationSuite(t *testing.T) {
	if err := tester.RunOnAny(
		tester.CorsoCITests,
		tester.CorsoOperationTests,
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
func (suite *BackupOpIntegrationSuite) TestBackup_Run() {
	ctx, flush := tester.NewContext()
	defer flush()

	m365UserID := tester.M365UserID(suite.T())
	acct := tester.NewM365Account(suite.T())

	tests := []struct {
		name       string
		selectFunc func() *selectors.Selector
	}{
		{
			name: "Integration Exchange.Mail",
			selectFunc: func() *selectors.Selector {
				sel := selectors.NewExchangeBackup()
				sel.Include(sel.MailFolders([]string{m365UserID}, []string{exchange.DefaultMailFolder}, selectors.PrefixMatch()))
				return &sel.Selector
			},
		},
		{
			name: "Integration Exchange.Contacts",
			selectFunc: func() *selectors.Selector {
				sel := selectors.NewExchangeBackup()
				sel.Include(sel.ContactFolders(
					[]string{m365UserID},
					[]string{exchange.DefaultContactFolder},
					selectors.PrefixMatch()))
				return &sel.Selector
			},
		},
		{
			name: "Integration Exchange.Events",
			selectFunc: func() *selectors.Selector {
				sel := selectors.NewExchangeBackup()
				sel.Include(sel.EventCalendars([]string{m365UserID}, []string{exchange.DefaultCalendar}, selectors.PrefixMatch()))
				return &sel.Selector
			},
		},
	}
	for _, test := range tests {
		suite.T().Run(test.name, func(t *testing.T) {
			// need to initialize the repository before we can test connecting to it.
			st := tester.NewPrefixedS3Storage(t)
			k := kopia.NewConn(st)
			require.NoError(t, k.Initialize(ctx))

			// kopiaRef comes with a count of 1 and Wrapper bumps it again so safe
			// to close here.
			defer k.Close(ctx)

			kw, err := kopia.NewWrapper(k)
			require.NoError(t, err)
			defer kw.Close(ctx)

			ms, err := kopia.NewModelStore(k)
			require.NoError(t, err)
			defer ms.Close(ctx)

			mb := evmock.NewBus()

			sw := store.NewKopiaStore(ms)
			selected := test.selectFunc()
			bo, err := NewBackupOperation(
				ctx,
				control.Options{},
				kw,
				sw,
				acct,
				*selected,
				mb)
			require.NoError(t, err)

			require.NoError(t, bo.Run(ctx))
			require.NotEmpty(t, bo.Results)
			require.NotEmpty(t, bo.Results.BackupID)
			assert.Equalf(t, Completed, bo.Status, "backup status %s is not Completed", bo.Status)
			assert.Less(t, 0, bo.Results.ItemsRead)
			assert.Less(t, 0, bo.Results.ItemsWritten)
			assert.Less(t, int64(0), bo.Results.BytesRead, "bytes read")
			assert.Less(t, int64(0), bo.Results.BytesUploaded, "bytes uploaded")
			assert.Equal(t, 1, bo.Results.ResourceOwners)
			assert.NoError(t, bo.Results.ReadErrors)
			assert.NoError(t, bo.Results.WriteErrors)
			assert.Equal(t, 1, mb.TimesCalled[events.BackupStart], "backup-start events")
			assert.Equal(t, 1, mb.TimesCalled[events.BackupEnd], "backup-end events")
			assert.Equal(t,
				mb.CalledWith[events.BackupStart][0][events.BackupID],
				bo.Results.BackupID, "backupID pre-declaration")
		})
	}
}

func (suite *BackupOpIntegrationSuite) TestBackupOneDrive_Run() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()

	m365UserID := tester.SecondaryM365UserID(t)
	acct := tester.NewM365Account(t)

	// need to initialize the repository before we can test connecting to it.
	st := tester.NewPrefixedS3Storage(t)

	k := kopia.NewConn(st)
	require.NoError(t, k.Initialize(ctx))

	// kopiaRef comes with a count of 1 and Wrapper bumps it again so safe
	// to close here.
	defer k.Close(ctx)

	kw, err := kopia.NewWrapper(k)
	require.NoError(t, err)

	defer kw.Close(ctx)

	ms, err := kopia.NewModelStore(k)
	require.NoError(t, err)

	defer ms.Close(ctx)

	sw := store.NewKopiaStore(ms)

	mb := evmock.NewBus()

	sel := selectors.NewOneDriveBackup()
	sel.Include(sel.Users([]string{m365UserID}))

	bo, err := NewBackupOperation(
		ctx,
		control.Options{},
		kw,
		sw,
		acct,
		sel.Selector,
		mb)
	require.NoError(t, err)

	require.NoError(t, bo.Run(ctx))
	require.NotEmpty(t, bo.Results)
	require.NotEmpty(t, bo.Results.BackupID)
	assert.Equalf(t, Completed, bo.Status, "backup status %s is not Completed", bo.Status)
	assert.Equal(t, bo.Results.ItemsRead, bo.Results.ItemsWritten)
	assert.Less(t, int64(0), bo.Results.BytesRead, "bytes read")
	assert.Less(t, int64(0), bo.Results.BytesUploaded, "bytes uploaded")
	assert.Equal(t, 1, bo.Results.ResourceOwners)
	assert.NoError(t, bo.Results.ReadErrors)
	assert.NoError(t, bo.Results.WriteErrors)
	assert.Equal(t, 1, mb.TimesCalled[events.BackupStart], "backup-start events")
	assert.Equal(t, 1, mb.TimesCalled[events.BackupEnd], "backup-end events")
	assert.Equal(t,
		mb.CalledWith[events.BackupStart][0][events.BackupID],
		bo.Results.BackupID, "backupID pre-declaration")
}
