package repository

import (
	"context"
	"testing"

	"github.com/alcionai/clues"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/kopia"
	"github.com/alcionai/corso/src/internal/model"
	"github.com/alcionai/corso/src/internal/operations"
	"github.com/alcionai/corso/src/internal/stats"
	"github.com/alcionai/corso/src/internal/streamstore"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/internal/version"
	"github.com/alcionai/corso/src/pkg/backup"
	"github.com/alcionai/corso/src/pkg/backup/details"
	rep "github.com/alcionai/corso/src/pkg/control/repository"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	storeTD "github.com/alcionai/corso/src/pkg/storage/testdata"
	"github.com/alcionai/corso/src/pkg/store"
	"github.com/alcionai/corso/src/pkg/store/mock"
)

// ---------------------------------------------------------------------------
// Unit
// ---------------------------------------------------------------------------

type RepositoryBackupsUnitSuite struct {
	tester.Suite
}

func TestRepositoryBackupsUnitSuite(t *testing.T) {
	suite.Run(t, &RepositoryBackupsUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *RepositoryBackupsUnitSuite) TestGetBackup() {
	bup := &backup.Backup{
		BaseModel: model.BaseModel{
			ID: model.StableID(uuid.NewString()),
		},
	}

	table := []struct {
		name      string
		sw        mock.BackupWrapper
		expectErr func(t *testing.T, result error)
		expectID  model.StableID
	}{
		{
			name: "no error",
			sw: mock.BackupWrapper{
				Backup:    bup,
				GetErr:    nil,
				DeleteErr: nil,
			},
			expectErr: func(t *testing.T, result error) {
				assert.NoError(t, result, clues.ToCore(result))
			},
			expectID: bup.ID,
		},
		{
			name: "get error",
			sw: mock.BackupWrapper{
				Backup:    bup,
				GetErr:    data.ErrNotFound,
				DeleteErr: nil,
			},
			expectErr: func(t *testing.T, result error) {
				assert.ErrorIs(t, result, data.ErrNotFound, clues.ToCore(result))
				assert.ErrorIs(t, result, ErrorBackupNotFound, clues.ToCore(result))
			},
			expectID: bup.ID,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			b, err := getBackup(ctx, string(bup.ID), test.sw)
			test.expectErr(t, err)

			if err != nil {
				return
			}

			assert.Equal(t, test.expectID, b.ID)
		})
	}
}

type mockSSDeleter struct {
	err error
}

func (sd mockSSDeleter) DeleteSnapshot(_ context.Context, _ string) error {
	return sd.err
}

func (suite *RepositoryBackupsUnitSuite) TestDeleteBackup() {
	bup := &backup.Backup{
		BaseModel: model.BaseModel{
			ID: model.StableID(uuid.NewString()),
		},
	}

	bupNoSnapshot := &backup.Backup{
		BaseModel: model.BaseModel{},
	}

	table := []struct {
		name      string
		sw        mock.BackupWrapper
		kw        mockSSDeleter
		expectErr func(t *testing.T, result error)
		expectID  model.StableID
	}{
		{
			name: "no error",
			sw: mock.BackupWrapper{
				Backup:    bup,
				GetErr:    nil,
				DeleteErr: nil,
			},
			kw: mockSSDeleter{},
			expectErr: func(t *testing.T, result error) {
				assert.NoError(t, result, clues.ToCore(result))
			},
			expectID: bup.ID,
		},
		{
			name: "get error",
			sw: mock.BackupWrapper{
				Backup:    bup,
				GetErr:    data.ErrNotFound,
				DeleteErr: nil,
			},
			kw: mockSSDeleter{},
			expectErr: func(t *testing.T, result error) {
				assert.ErrorIs(t, result, data.ErrNotFound, clues.ToCore(result))
				assert.ErrorIs(t, result, ErrorBackupNotFound, clues.ToCore(result))
			},
			expectID: bup.ID,
		},
		{
			name: "delete error",
			sw: mock.BackupWrapper{
				Backup:    bup,
				GetErr:    nil,
				DeleteErr: assert.AnError,
			},
			kw: mockSSDeleter{},
			expectErr: func(t *testing.T, result error) {
				assert.ErrorIs(t, result, assert.AnError, clues.ToCore(result))
			},
			expectID: bup.ID,
		},
		{
			name: "snapshot delete error",
			sw: mock.BackupWrapper{
				Backup:    bup,
				GetErr:    nil,
				DeleteErr: nil,
			},
			kw: mockSSDeleter{assert.AnError},
			expectErr: func(t *testing.T, result error) {
				assert.ErrorIs(t, result, assert.AnError, clues.ToCore(result))
			},
			expectID: bup.ID,
		},
		{
			name: "no snapshot present",
			sw: mock.BackupWrapper{
				Backup:    bupNoSnapshot,
				GetErr:    nil,
				DeleteErr: nil,
			},
			kw: mockSSDeleter{assert.AnError},
			expectErr: func(t *testing.T, result error) {
				assert.NoError(t, result, clues.ToCore(result))
			},
			expectID: bupNoSnapshot.ID,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			err := deleteBackup(ctx, string(test.sw.Backup.ID), test.kw, test.sw)
			test.expectErr(t, err)
		})
	}
}

// ---------------------------------------------------------------------------
// integration
// ---------------------------------------------------------------------------

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
			[][]string{storeTD.AWSStorageCredEnvs, tconfig.M365AcctCredEnvs}),
	})
}

func (suite *RepositoryModelIntgSuite) SetupSuite() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		s   = storeTD.NewPrefixedS3Storage(t)
		k   = kopia.NewConn(s)
		err error
	)

	require.NotNil(t, k)

	err = k.Initialize(ctx, rep.Options{})
	require.NoError(t, err, clues.ToCore(err))

	err = k.Connect(ctx, rep.Options{})
	require.NoError(t, err, clues.ToCore(err))

	suite.kopiaCloser = func(ctx context.Context) {
		k.Close(ctx)
	}

	suite.kw, err = kopia.NewWrapper(k)
	require.NoError(t, err, clues.ToCore(err))

	suite.ms, err = kopia.NewModelStore(k)
	require.NoError(t, err, clues.ToCore(err))

	suite.sw = store.NewKopiaStore(suite.ms)
}

func (suite *RepositoryModelIntgSuite) TearDownSuite() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	if suite.ms != nil {
		err := suite.ms.Close(ctx)
		assert.NoError(t, err, clues.ToCore(err))
	}

	if suite.kw != nil {
		err := suite.kw.Close(ctx)
		assert.NoError(t, err, clues.ToCore(err))
	}

	if suite.kopiaCloser != nil {
		suite.kopiaCloser(ctx)
	}
}

func (suite *RepositoryModelIntgSuite) TestGetRepositoryModel() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		s = storeTD.NewPrefixedS3Storage(t)
		k = kopia.NewConn(s)
	)

	require.NoError(t, k.Initialize(ctx, rep.Options{}))
	require.NoError(t, k.Connect(ctx, rep.Options{}))

	defer k.Close(ctx)

	ms, err := kopia.NewModelStore(k)
	require.NoError(t, err)

	defer ms.Close(ctx)

	err = newRepoModel(ctx, ms, "fnords")
	require.NoError(t, err, clues.ToCore(err))

	got, err := getRepoModel(ctx, ms)
	require.NoError(t, err, clues.ToCore(err))
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
	ownerID, ownerName string,
	deets *details.Details,
	fe *fault.Errors,
	errs *fault.Bus,
) *backup.Backup {
	var (
		serv   = sel.PathService()
		sstore = streamstore.NewStreamer(kw, tID, serv)
	)

	err := sstore.Collect(ctx, streamstore.DetailsCollector(deets))
	require.NoError(t, err, "collecting details in streamstore")

	err = sstore.Collect(ctx, streamstore.FaultErrorsCollector(fe))
	require.NoError(t, err, "collecting errors in streamstore")

	ssid, err := sstore.Write(ctx, errs)
	require.NoError(t, err, "writing to streamstore")

	tags := map[string]string{
		model.ServiceTag: sel.PathService().String(),
	}

	b := backup.New(
		snapID, ssid,
		operations.Completed.String(),
		version.Backup,
		model.StableID(backupID),
		sel,
		ownerID, ownerName,
		stats.ReadWrites{},
		stats.StartAndEndTime{},
		fe,
		tags)

	err = sw.Put(ctx, model.BackupSchema, b)
	require.NoError(t, err)

	return b
}

func (suite *RepositoryModelIntgSuite) TestGetBackupDetails() {
	const (
		brunhilda = "brunhilda"
		tenantID  = "tenant"
	)

	info := details.ItemInfo{
		Exchange: &details.ExchangeInfo{
			ItemType: details.ExchangeMail,
		},
	}

	repoPath, err := path.FromDataLayerPath(tenantID+"/exchange/user-id/email/test/foo", true)
	require.NoError(suite.T(), err, clues.ToCore(err))

	loc := path.Builder{}.Append(repoPath.Folders()...)

	builder := &details.Builder{}
	require.NoError(suite.T(), builder.Add(repoPath, loc, false, info))

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
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			b := writeBackup(
				t,
				ctx,
				suite.kw,
				suite.sw,
				tenantID, "snapID", test.writeBupID,
				selectors.NewExchangeBackup([]string{brunhilda}).Selector,
				brunhilda, brunhilda,
				test.deets,
				&fault.Errors{},
				fault.New(true))

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
		tenantID  = "tenant"
		failFast  = true
		brunhilda = "brunhilda"
	)

	var (
		err  = clues.Wrap(assert.AnError, "wrap")
		cec  = err.Core()
		item = fault.FileErr(err, "ns", "file-id", "file-name", map[string]any{"foo": "bar"})
		skip = fault.FileSkip(fault.SkipMalware, "ns", "s-file-id", "s-file-name", map[string]any{"foo": "bar"})
		info = details.ItemInfo{
			Exchange: &details.ExchangeInfo{
				ItemType: details.ExchangeMail,
			},
		}
	)

	repoPath, err2 := path.FromDataLayerPath(tenantID+"/exchange/user-id/email/test/foo", true)
	require.NoError(suite.T(), err2, clues.ToCore(err2))

	loc := path.Builder{}.Append(repoPath.Folders()...)

	builder := &details.Builder{}
	require.NoError(suite.T(), builder.Add(repoPath, loc, false, info))

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
			expectErrors: &fault.Errors{},
			expectErr:    require.NoError,
		},
		{
			name:       "good",
			writeBupID: "error_squirrels",
			readBupID:  "error_squirrels",
			deets:      builder.Details(),
			errors: &fault.Errors{
				Failure:   cec,
				Recovered: []*clues.ErrCore{cec},
				Items:     []fault.Item{*item},
				Skipped:   []fault.Skipped{*skip},
				FailFast:  failFast,
			},
			expectErrors: &fault.Errors{
				Failure:   cec,
				Recovered: []*clues.ErrCore{cec},
				Items:     []fault.Item{*item},
				Skipped:   []fault.Skipped{*skip},
				FailFast:  failFast,
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
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			b := writeBackup(
				t,
				ctx,
				suite.kw,
				suite.sw,
				tenantID, "snapID", test.writeBupID,
				selectors.NewExchangeBackup([]string{brunhilda}).Selector,
				brunhilda, brunhilda,
				test.deets,
				test.errors,
				fault.New(failFast))

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
