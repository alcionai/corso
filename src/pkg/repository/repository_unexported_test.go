package repository

import (
	"context"
	"testing"

	"github.com/alcionai/clues"
	"github.com/google/uuid"
	"github.com/kopia/kopia/repo/manifest"
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
// Mocks
// ---------------------------------------------------------------------------

type mockBackupList struct {
	backups []*backup.Backup
	err     error
	check   func(fs []store.FilterOption)
}

func (mbl mockBackupList) GetBackup(
	ctx context.Context,
	backupID model.StableID,
) (*backup.Backup, error) {
	return nil, clues.New("not implemented")
}

func (mbl mockBackupList) DeleteBackup(
	ctx context.Context,
	backupID model.StableID,
) error {
	return clues.New("not implemented")
}

func (mbl mockBackupList) GetBackups(
	ctx context.Context,
	filters ...store.FilterOption,
) ([]*backup.Backup, error) {
	if mbl.check != nil {
		mbl.check(filters)
	}

	return mbl.backups, mbl.err
}

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

func (suite *RepositoryBackupsUnitSuite) TestBackupsByTag() {
	unlabeled1 := &backup.Backup{
		BaseModel: model.BaseModel{
			ID: model.StableID(uuid.NewString()),
		},
	}
	unlabeled2 := &backup.Backup{
		BaseModel: model.BaseModel{
			ID: model.StableID(uuid.NewString()),
		},
	}

	merge1 := &backup.Backup{
		BaseModel: model.BaseModel{
			ID: model.StableID(uuid.NewString()),
			Tags: map[string]string{
				model.BackupTypeTag: model.MergeBackup,
			},
		},
	}
	merge2 := &backup.Backup{
		BaseModel: model.BaseModel{
			ID: model.StableID(uuid.NewString()),
			Tags: map[string]string{
				model.BackupTypeTag: model.MergeBackup,
			},
		},
	}

	assist1 := &backup.Backup{
		BaseModel: model.BaseModel{
			ID: model.StableID(uuid.NewString()),
			Tags: map[string]string{
				model.BackupTypeTag: model.AssistBackup,
			},
		},
	}
	assist2 := &backup.Backup{
		BaseModel: model.BaseModel{
			ID: model.StableID(uuid.NewString()),
			Tags: map[string]string{
				model.BackupTypeTag: model.AssistBackup,
			},
		},
	}

	table := []struct {
		name       string
		getBackups []*backup.Backup
		filters    []store.FilterOption
		listErr    error
		expectErr  assert.ErrorAssertionFunc
		expect     []*backup.Backup
	}{
		{
			name: "UnlabeledOnly",
			getBackups: []*backup.Backup{
				unlabeled1,
				unlabeled2,
			},
			expectErr: assert.NoError,
			expect: []*backup.Backup{
				unlabeled1,
				unlabeled2,
			},
		},
		{
			name: "MergeOnly",
			getBackups: []*backup.Backup{
				merge1,
				merge2,
			},
			expectErr: assert.NoError,
			expect: []*backup.Backup{
				merge1,
				merge2,
			},
		},
		{
			name: "AssistOnly",
			getBackups: []*backup.Backup{
				assist1,
				assist2,
			},
			expectErr: assert.NoError,
		},
		{
			name: "UnlabledAndMerge",
			getBackups: []*backup.Backup{
				merge1,
				unlabeled1,
				merge2,
				unlabeled2,
			},
			expectErr: assert.NoError,
			expect: []*backup.Backup{
				merge1,
				merge2,
				unlabeled1,
				unlabeled2,
			},
		},
		{
			name: "UnlabeledAndAssist",
			getBackups: []*backup.Backup{
				unlabeled1,
				assist1,
				unlabeled2,
				assist2,
			},
			expectErr: assert.NoError,
			expect: []*backup.Backup{
				unlabeled1,
				unlabeled2,
			},
		},
		{
			name: "MergeAndAssist",
			getBackups: []*backup.Backup{
				merge1,
				assist1,
				merge2,
				assist2,
			},
			expectErr: assert.NoError,
			expect: []*backup.Backup{
				merge1,
				merge2,
			},
		},
		{
			name: "UnlabeledAndMergeAndAssist",
			getBackups: []*backup.Backup{
				unlabeled1,
				merge1,
				assist1,
				merge2,
				unlabeled2,
				assist2,
			},
			expectErr: assert.NoError,
			expect: []*backup.Backup{
				merge1,
				merge2,
				unlabeled1,
				unlabeled2,
			},
		},
		{
			name: "LookupError",
			getBackups: []*backup.Backup{
				unlabeled1,
				merge1,
				assist1,
				merge2,
				unlabeled2,
				assist2,
			},
			listErr:   assert.AnError,
			expectErr: assert.Error,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			mbl := mockBackupList{
				backups: test.getBackups,
				err:     test.listErr,
				check: func(fs []store.FilterOption) {
					assert.ElementsMatch(t, test.filters, fs)
				},
			}

			bs, err := backupsByTag(ctx, mbl, test.filters)
			test.expectErr(t, err, clues.ToCore(err))

			assert.ElementsMatch(t, test.expect, bs)
		})
	}
}

type getRes struct {
	bup *backup.Backup
	err error
}

type mockBackupGetterModelDeleter struct {
	t *testing.T

	gets       []getRes
	deleteErrs []error

	expectGets []model.StableID
	expectDels [][]string

	getCount int
	delCount int
}

func (m *mockBackupGetterModelDeleter) GetBackup(
	_ context.Context,
	id model.StableID,
) (*backup.Backup, error) {
	defer func() {
		m.getCount++
	}()

	assert.Equal(m.t, m.expectGets[m.getCount], id)

	return m.gets[m.getCount].bup, clues.Stack(m.gets[m.getCount].err).OrNil()
}

func (m *mockBackupGetterModelDeleter) DeleteWithModelStoreIDs(
	_ context.Context,
	ids ...manifest.ID,
) error {
	defer func() {
		m.delCount++
	}()

	converted := make([]string, 0, len(ids))
	for _, id := range ids {
		converted = append(converted, string(id))
	}

	assert.ElementsMatch(m.t, m.expectDels[m.delCount], converted)

	return clues.Stack(m.deleteErrs[m.delCount]).OrNil()
}

func (suite *RepositoryBackupsUnitSuite) TestDeleteBackups() {
	bup := &backup.Backup{
		BaseModel: model.BaseModel{
			ID:           model.StableID("current-bup-id"),
			ModelStoreID: manifest.ID("current-bup-msid"),
		},
		SnapshotID:    "current-bup-dsid",
		StreamStoreID: "current-bup-ssid",
	}

	bupLegacy := &backup.Backup{
		BaseModel: model.BaseModel{
			ID:           model.StableID("legacy-bup-id"),
			ModelStoreID: manifest.ID("legacy-bup-msid"),
		},
		SnapshotID: "legacy-bup-dsid",
		DetailsID:  "legacy-bup-did",
	}

	bupNoSnapshot := &backup.Backup{
		BaseModel: model.BaseModel{
			ID:           model.StableID("ns-bup-id"),
			ModelStoreID: manifest.ID("ns-bup-id-msid"),
		},
		StreamStoreID: "ns-bup-ssid",
	}

	bupNoDetails := &backup.Backup{
		BaseModel: model.BaseModel{
			ID:           model.StableID("nssid-bup-id"),
			ModelStoreID: manifest.ID("nssid-bup-msid"),
		},
		SnapshotID: "nssid-bup-dsid",
	}

	table := []struct {
		name       string
		inputIDs   []model.StableID
		gets       []getRes
		expectGets []model.StableID
		dels       []error
		expectDels [][]string
		expectErr  func(t *testing.T, result error)
	}{
		{
			name: "SingleBackup NoError",
			inputIDs: []model.StableID{
				bup.ID,
			},
			gets: []getRes{
				{bup: bup},
			},
			expectGets: []model.StableID{
				bup.ID,
			},
			dels: []error{
				nil,
			},
			expectDels: [][]string{
				{
					string(bup.ModelStoreID),
					bup.SnapshotID,
					bup.StreamStoreID,
				},
			},
			expectErr: func(t *testing.T, result error) {
				assert.NoError(t, result, clues.ToCore(result))
			},
		},
		{
			name: "SingleBackup GetError",
			inputIDs: []model.StableID{
				bup.ID,
			},
			gets: []getRes{
				{err: data.ErrNotFound},
			},
			expectGets: []model.StableID{
				bup.ID,
			},
			expectErr: func(t *testing.T, result error) {
				assert.ErrorIs(t, result, data.ErrNotFound, clues.ToCore(result))
				assert.ErrorIs(t, result, ErrorBackupNotFound, clues.ToCore(result))
			},
		},
		{
			name: "SingleBackup DeleteError",
			inputIDs: []model.StableID{
				bup.ID,
			},
			gets: []getRes{
				{bup: bup},
			},
			expectGets: []model.StableID{
				bup.ID,
			},
			dels: []error{assert.AnError},
			expectDels: [][]string{
				{
					string(bup.ModelStoreID),
					bup.SnapshotID,
					bup.StreamStoreID,
				},
			},
			expectErr: func(t *testing.T, result error) {
				assert.ErrorIs(t, result, assert.AnError, clues.ToCore(result))
			},
		},
		{
			name: "SingleBackup NoSnapshot",
			inputIDs: []model.StableID{
				bupNoSnapshot.ID,
			},
			gets: []getRes{
				{bup: bupNoSnapshot},
			},
			expectGets: []model.StableID{
				bupNoSnapshot.ID,
			},
			dels: []error{nil},
			expectDels: [][]string{
				{
					string(bupNoSnapshot.ModelStoreID),
					bupNoSnapshot.StreamStoreID,
				},
			},
			expectErr: func(t *testing.T, result error) {
				assert.NoError(t, result, clues.ToCore(result))
			},
		},
		{
			name: "SingleBackup NoDetails",
			inputIDs: []model.StableID{
				bupNoDetails.ID,
			},
			gets: []getRes{
				{bup: bupNoDetails},
			},
			expectGets: []model.StableID{
				bupNoDetails.ID,
			},
			dels: []error{nil},
			expectDels: [][]string{
				{
					string(bupNoDetails.ModelStoreID),
					bupNoDetails.SnapshotID,
				},
			},
			expectErr: func(t *testing.T, result error) {
				assert.NoError(t, result, clues.ToCore(result))
			},
		},
		{
			name: "SingleBackup OldDetailsID",
			inputIDs: []model.StableID{
				bupLegacy.ID,
			},
			gets: []getRes{
				{bup: bupLegacy},
			},
			expectGets: []model.StableID{
				bupLegacy.ID,
			},
			dels: []error{nil},
			expectDels: [][]string{
				{
					string(bupLegacy.ModelStoreID),
					bupLegacy.SnapshotID,
					bupLegacy.DetailsID,
				},
			},
			expectErr: func(t *testing.T, result error) {
				assert.NoError(t, result, clues.ToCore(result))
			},
		},
		{
			name: "MultipleBackups NoError",
			inputIDs: []model.StableID{
				bup.ID,
				bupLegacy.ID,
				bupNoSnapshot.ID,
				bupNoDetails.ID,
			},
			gets: []getRes{
				{bup: bup},
				{bup: bupLegacy},
				{bup: bupNoSnapshot},
				{bup: bupNoDetails},
			},
			expectGets: []model.StableID{
				bup.ID,
				bupLegacy.ID,
				bupNoSnapshot.ID,
				bupNoDetails.ID,
			},
			dels: []error{nil},
			expectDels: [][]string{
				{
					string(bup.ModelStoreID),
					bup.SnapshotID,
					bup.StreamStoreID,
					string(bupLegacy.ModelStoreID),
					bupLegacy.SnapshotID,
					bupLegacy.DetailsID,
					string(bupNoSnapshot.ModelStoreID),
					bupNoSnapshot.StreamStoreID,
					string(bupNoDetails.ModelStoreID),
					bupNoDetails.SnapshotID,
				},
			},
			expectErr: func(t *testing.T, result error) {
				assert.NoError(t, result, clues.ToCore(result))
			},
		},
		{
			name: "MultipleBackups GetError",
			inputIDs: []model.StableID{
				bup.ID,
				bupLegacy.ID,
				bupNoSnapshot.ID,
				bupNoDetails.ID,
			},
			gets: []getRes{
				{bup: bup},
				{bup: bupLegacy},
				{bup: bupNoSnapshot},
				{err: data.ErrNotFound},
			},
			expectGets: []model.StableID{
				bup.ID,
				bupLegacy.ID,
				bupNoSnapshot.ID,
				bupNoDetails.ID,
			},
			expectErr: func(t *testing.T, result error) {
				assert.ErrorIs(t, result, data.ErrNotFound, clues.ToCore(result))
				assert.ErrorIs(t, result, ErrorBackupNotFound, clues.ToCore(result))
			},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()
			m := &mockBackupGetterModelDeleter{
				t: t,

				gets:       test.gets,
				deleteErrs: test.dels,

				expectGets: test.expectGets,
				expectDels: test.expectDels,
			}

			ctx, flush := tester.NewContext(t)
			defer flush()

			strIDs := make([]string, 0, len(test.inputIDs))
			for _, id := range test.inputIDs {
				strIDs = append(strIDs, string(id))
			}

			err := deleteBackups(ctx, m, strIDs...)
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
	sw          store.BackupStorer
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

	err = k.Initialize(ctx, rep.Options{}, rep.Retention{})
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

	suite.sw = store.NewWrapper(suite.ms)
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

	err := k.Initialize(ctx, rep.Options{}, rep.Retention{})
	require.NoError(t, err, "initializing repo: %v", clues.ToCore(err))

	err = k.Connect(ctx, rep.Options{})
	require.NoError(t, err, "connecting to repo: %v", clues.ToCore(err))

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
	sw store.BackupStorer,
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
	require.NoError(suite.T(), builder.Add(repoPath, loc, info))

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
	require.NoError(suite.T(), builder.Add(repoPath, loc, info))

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
