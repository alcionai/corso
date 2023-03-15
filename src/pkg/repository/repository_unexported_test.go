package repository

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/clues"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/kopia"
	"github.com/alcionai/corso/src/internal/model"
	"github.com/alcionai/corso/src/internal/operations"
	"github.com/alcionai/corso/src/internal/stats"
	"github.com/alcionai/corso/src/internal/streamstore"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/backup"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/store"
)

type RepositoryModelIntgSuite struct {
	tester.Suite
	kw          *kopia.Wrapper
	ms          *kopia.ModelStore
	sw          *store.Wrapper
	kopiaCloser func(ctx context.Context)
}

func TestRepositoryModelIntgSuite(t *testing.T) {
	suite.Run(t, &RepositoryModelIntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tester.AWSStorageCredEnvs, tester.M365AcctCredEnvs},
			tester.CorsoRepositoryTests,
		),
	})
}

func (suite *RepositoryModelIntgSuite) SetupSuite() {
	ctx, flush := tester.NewContext()
	defer flush()

	var (
		t   = suite.T()
		s   = tester.NewPrefixedS3Storage(t)
		k   = kopia.NewConn(s)
		err error
	)

	require.NotNil(t, k)
	require.NoError(t, k.Initialize(ctx))

	suite.kopiaCloser = func(ctx context.Context) {
		k.Close(ctx)
	}

	suite.kw, err = kopia.NewWrapper(k)
	require.NoError(t, err)

	suite.ms, err = kopia.NewModelStore(k)
	require.NoError(t, err)

	suite.sw = store.NewKopiaStore(suite.ms)
}

func (suite *RepositoryModelIntgSuite) TearDownSuite() {
	ctx, flush := tester.NewContext()
	defer flush()

	if suite.ms != nil {
		suite.ms.Close(ctx)
	}

	if suite.kw != nil {
		suite.kw.Close(ctx)
	}

	if suite.kopiaCloser != nil {
		suite.kopiaCloser(ctx)
	}
}

func (suite *RepositoryModelIntgSuite) TestGetRepositoryModel() {
	ctx, flush := tester.NewContext()
	defer flush()

	var (
		t = suite.T()
		s = tester.NewPrefixedS3Storage(t)
		k = kopia.NewConn(s)
	)

	require.NoError(t, k.Initialize(ctx))
	require.NoError(t, k.Connect(ctx))

	defer k.Close(ctx)

	ms, err := kopia.NewModelStore(k)
	require.NoError(t, err)

	defer ms.Close(ctx)

	require.NoError(t, newRepoModel(ctx, ms, "fnords"))

	got, err := getRepoModel(ctx, ms)
	require.NoError(t, err)
	assert.Equal(t, "fnords", string(got.ID))
}

// helper func for writing backups
func writeBackup(
	t *testing.T,
	ctx context.Context, //revive:disable-line:context-as-argument
	kw *kopia.Wrapper,
	sw *store.Wrapper,
	tID, snapID, backupID string,
	sel selectors.Selector,
	deets *details.Details,
	errors fault.Errors,
	errs *fault.Bus,
) *backup.Backup {
	var (
		serv   = sel.PathService()
		sstore = streamstore.NewStreamer(kw, tID, serv)
	)

	err := sstore.Collect(ctx, streamstore.DetailsCollector(deets))
	require.NoError(t, err, "collecting details in streamstore")

	err = sstore.Collect(ctx, streamstore.FaultErrorsCollector(errs.Errors()))
	require.NoError(t, err, "collecting errors in streamstore")

	ssid, err := sstore.Write(ctx, errs)
	require.NoError(t, err, "writing to streamstore")

	b := backup.New(
		snapID, ssid,
		operations.Completed.String(),
		model.StableID(backupID),
		sel,
		stats.ReadWrites{},
		stats.StartAndEndTime{},
		errs)

	err = sw.Put(ctx, model.BackupSchema, b)
	require.NoError(t, err)

	return b
}

func (suite *RepositoryModelIntgSuite) TestGetBackupDetails() {
	const tenantID = "tenant"

	info := details.ItemInfo{
		Folder: &details.FolderInfo{
			DisplayName: "test",
		},
	}

	builder := &details.Builder{}
	builder.Add("ref", "short", "pref", "lref", false, info)

	table := []struct {
		name       string
		writeBupID string
		readBupID  string
		deets      *details.Details
		expectErr  require.ErrorAssertionFunc
	}{
		{
			name:       "good",
			writeBupID: "squirrels",
			readBupID:  "squirrels",
			deets:      builder.Details(),
			expectErr:  require.NoError,
		},
		{
			name:       "missing backup",
			writeBupID: "chipmunks",
			readBupID:  "weasels",
			deets:      builder.Details(),
			expectErr:  require.Error,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			ctx, flush := tester.NewContext()
			defer flush()

			var (
				t = suite.T()
				b = writeBackup(
					t,
					ctx,
					suite.kw,
					suite.sw,
					tenantID, "snapID", test.writeBupID,
					selectors.NewExchangeBackup([]string{"brunhilda"}).Selector,
					test.deets,
					fault.Errors{},
					fault.New(true))
			)

			rDeets, rBup, err := getBackupDetails(ctx, test.readBupID, tenantID, suite.kw, suite.sw, fault.New(true))
			test.expectErr(t, err)

			if err != nil {
				return
			}

			assert.Equal(t, b.DetailsID, rBup.DetailsID, "returned details ID matches")
			assert.Equal(t, test.deets, rDeets, "returned details ID matches")
		})
	}
}

func (suite *RepositoryModelIntgSuite) TestGetBackupErrors() {
	const (
		tenantID = "tenant"
		failFast = true
	)

	err := clues.Wrap(assert.AnError, "wrap")

	info := details.ItemInfo{
		Folder: &details.FolderInfo{
			DisplayName: "test",
		},
	}

	builder := &details.Builder{}
	builder.Add("ref", "short", "pref", "lref", false, info)

	table := []struct {
		name         string
		writeBupID   string
		readBupID    string
		deets        *details.Details
		errors       *fault.Errors
		expectErrors *fault.Errors
		expectErr    require.ErrorAssertionFunc
	}{
		{
			name:         "nil errors",
			writeBupID:   "error_marmots",
			readBupID:    "error_marmots",
			deets:        builder.Details(),
			errors:       nil,
			expectErrors: &fault.Errors{Items: []fault.Item{}, FailFast: failFast},
			expectErr:    require.NoError,
		},
		{
			name:       "good",
			writeBupID: "error_squirrels",
			readBupID:  "error_squirrels",
			deets:      builder.Details(),
			errors:     &fault.Errors{Failure: err},
			expectErrors: &fault.Errors{
				Failure:  nil,
				Items:    []fault.Item{},
				FailFast: failFast,
			},
			expectErr: require.NoError,
		},
		{
			name:       "missing backup",
			writeBupID: "error_chipmunks",
			readBupID:  "error_weasels",
			deets:      builder.Details(),
			errors:     nil,
			expectErr:  require.Error,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			ctx, flush := tester.NewContext()
			defer flush()

			var (
				t = suite.T()
				b = writeBackup(
					t,
					ctx,
					suite.kw,
					suite.sw,
					tenantID, "snapID", test.writeBupID,
					selectors.NewExchangeBackup([]string{"brunhilda"}).Selector,
					test.deets,
					ptr.Val(test.errors),
					fault.New(true))
			)

			rErrors, rBup, err := getBackupErrors(ctx, test.readBupID, tenantID, suite.kw, suite.sw, fault.New(failFast))
			test.expectErr(t, err)

			if err != nil {
				return
			}

			assert.Equal(t, b.StreamStoreID, rBup.StreamStoreID, "returned streamstore ID matches")
			assert.Equal(t, test.expectErrors, rErrors, "retrieved errors match")
		})
	}
}
