package kopia

import (
	"bytes"
	"context"
	"io"
	stdpath "path"
	"strings"
	"testing"
	"time"

	"github.com/alcionai/clues"
	"github.com/kopia/kopia/repo"
	"github.com/kopia/kopia/repo/blob"
	"github.com/kopia/kopia/repo/format"
	"github.com/kopia/kopia/repo/maintenance"
	"github.com/kopia/kopia/repo/manifest"
	"github.com/kopia/kopia/snapshot"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"golang.org/x/exp/maps"

	pmMock "github.com/alcionai/corso/src/internal/common/prefixmatcher/mock"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/data"
	dataMock "github.com/alcionai/corso/src/internal/data/mock"
	"github.com/alcionai/corso/src/internal/m365/collection/drive/metadata"
	m365Mock "github.com/alcionai/corso/src/internal/m365/mock"
	exchMock "github.com/alcionai/corso/src/internal/m365/service/exchange/mock"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/backup/identity"
	"github.com/alcionai/corso/src/pkg/control/repository"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
	storeTD "github.com/alcionai/corso/src/pkg/storage/testdata"
)

const (
	testTenant     = "a-tenant"
	testUser       = "user1"
	testInboxID    = "Inbox_ID"
	testInboxDir   = "Inbox"
	testArchiveID  = "Archive_ID"
	testArchiveDir = "Archive"
	testFileName   = "file1"
	testFileName2  = "file2"
	testFileName3  = "file3"
	testFileName4  = "file4"
	testFileName5  = "file5"
	testFileName6  = "file6"
)

var (
	service       = path.ExchangeService.String()
	category      = path.EmailCategory.String()
	testFileData  = []byte("abcdefghijklmnopqrstuvwxyz")
	testFileData2 = []byte("zyxwvutsrqponmlkjihgfedcba")
	testFileData3 = []byte("foo")
	testFileData4 = []byte("bar")
	testFileData5 = []byte("baz")
	// Intentional duplicate to make sure all files are scanned during recovery
	// (contrast to behavior of snapshotfs.TreeWalker).
	testFileData6 = testFileData
)

func testForFiles(
	t *testing.T,
	ctx context.Context, //revive:disable-line:context-as-argument
	expected map[string][]byte,
	collections []data.RestoreCollection,
) {
	t.Helper()

	count := 0

	for _, c := range collections {
		for s := range c.Items(ctx, fault.New(true)) {
			count++

			fullPath, err := c.FullPath().AppendItem(s.ID())
			require.NoError(t, err, clues.ToCore(err))

			expected, ok := expected[fullPath.String()]
			require.True(t, ok, "unexpected file with path %q", fullPath)

			buf, err := io.ReadAll(s.ToReader())
			require.NoError(t, err, "reading collection item", fullPath, clues.ToCore(err))
			assert.Equal(t, expected, buf, "comparing collection item", fullPath)

			require.Implements(t, (*data.ItemSize)(nil), s)

			ss := s.(data.ItemSize)
			assert.Equal(t, len(buf), int(ss.Size()))
		}
	}

	assert.Equal(t, len(expected), count)
}

func checkSnapshotTags(
	t *testing.T,
	ctx context.Context, //revive:disable-line:context-as-argument
	rep repo.Repository,
	expectedTags map[string]string,
	snapshotID string,
) {
	man, err := snapshot.LoadSnapshot(ctx, rep, manifest.ID(snapshotID))
	require.NoError(t, err, clues.ToCore(err))
	assert.Equal(t, expectedTags, man.Tags)
}

func toRestorePaths(t *testing.T, paths ...path.Path) []path.RestorePaths {
	res := make([]path.RestorePaths, 0, len(paths))

	for _, p := range paths {
		dir, err := p.Dir()
		require.NoError(t, err, clues.ToCore(err))

		res = append(res, path.RestorePaths{StoragePath: p, RestorePath: dir})
	}

	return res
}

// ---------------
// unit tests
// ---------------
type KopiaUnitSuite struct {
	tester.Suite
	testPath path.Path
}

func (suite *KopiaUnitSuite) SetupSuite() {
	tmp, err := path.FromDataLayerPath(
		stdpath.Join(
			testTenant,
			path.ExchangeService.String(),
			testUser,
			path.EmailCategory.String(),
			testInboxDir),
		false)
	require.NoError(suite.T(), err, clues.ToCore(err))

	suite.testPath = tmp
}

func TestKopiaUnitSuite(t *testing.T) {
	suite.Run(t, &KopiaUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *KopiaUnitSuite) TestCloseWithoutInitDoesNotPanic() {
	assert.NotPanics(suite.T(), func() {
		ctx, flush := tester.NewContext(suite.T())
		defer flush()

		w := &Wrapper{}
		w.Close(ctx)
	})
}

// ---------------
// integration tests that use kopia.
// ---------------
type BasicKopiaIntegrationSuite struct {
	tester.Suite
}

func TestBasicKopiaIntegrationSuite(t *testing.T) {
	suite.Run(t, &BasicKopiaIntegrationSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{storeTD.AWSStorageCredEnvs}),
	})
}

// TestMaintenance checks that different username/hostname pairs will or won't
// cause maintenance to run. It treats kopia maintenance as a black box and
// only checks the returned error.
func (suite *BasicKopiaIntegrationSuite) TestMaintenance_FirstRun_NoChanges() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	k, err := openKopiaRepo(t, ctx)
	require.NoError(t, err, clues.ToCore(err))

	w := &Wrapper{k}

	opts := repository.Maintenance{
		Safety: repository.FullMaintenanceSafety,
		Type:   repository.MetadataMaintenance,
	}

	err = w.RepoMaintenance(ctx, opts)
	require.NoError(t, err, clues.ToCore(err))
}

func (suite *BasicKopiaIntegrationSuite) TestMaintenance_WrongUser_NoForce_Fails() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	k, err := openKopiaRepo(t, ctx)
	require.NoError(t, err, clues.ToCore(err))

	w := &Wrapper{k}

	mOpts := repository.Maintenance{
		Safety: repository.FullMaintenanceSafety,
		Type:   repository.MetadataMaintenance,
	}

	// This will set the user.
	err = w.RepoMaintenance(ctx, mOpts)
	require.NoError(t, err, clues.ToCore(err))

	err = k.Close(ctx)
	require.NoError(t, err, clues.ToCore(err))

	opts := repository.Options{
		User: "foo",
		Host: "bar",
	}

	err = k.Connect(ctx, opts)
	require.NoError(t, err, clues.ToCore(err))

	var notOwnedErr maintenance.NotOwnedError

	err = w.RepoMaintenance(ctx, mOpts)
	assert.ErrorAs(t, err, &notOwnedErr, clues.ToCore(err))
}

func (suite *BasicKopiaIntegrationSuite) TestMaintenance_WrongUser_Force_Succeeds() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	k, err := openKopiaRepo(t, ctx)
	require.NoError(t, err, clues.ToCore(err))

	w := &Wrapper{k}

	mOpts := repository.Maintenance{
		Safety: repository.FullMaintenanceSafety,
		Type:   repository.MetadataMaintenance,
	}

	// This will set the user.
	err = w.RepoMaintenance(ctx, mOpts)
	require.NoError(t, err, clues.ToCore(err))

	err = k.Close(ctx)
	require.NoError(t, err, clues.ToCore(err))

	opts := repository.Options{
		User: "foo",
		Host: "bar",
	}

	err = k.Connect(ctx, opts)
	require.NoError(t, err, clues.ToCore(err))

	mOpts.Force = true

	// This will set the user.
	err = w.RepoMaintenance(ctx, mOpts)
	require.NoError(t, err, clues.ToCore(err))

	mOpts.Force = false

	// Running without force should succeed now.
	err = w.RepoMaintenance(ctx, mOpts)
	require.NoError(t, err, clues.ToCore(err))
}

// Test that failing to put the storage blob will skip updating the maintenance
// manifest too. It's still possible to end up halfway updating the repo config
// blobs as there's several of them, but at least this gives us something.
func (suite *BasicKopiaIntegrationSuite) TestSetRetentionParameters_NoChangesOnFailure() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	k, err := openKopiaRepo(t, ctx)
	require.NoError(t, err, clues.ToCore(err))

	w := &Wrapper{k}

	// Enable retention.
	err = w.SetRetentionParameters(ctx, repository.Retention{
		Mode:     ptr.To(repository.GovernanceRetention),
		Duration: ptr.To(time.Hour * 48),
		Extend:   ptr.To(true),
	})
	require.Error(t, err, clues.ToCore(err))

	checkRetentionParams(
		t,
		ctx,
		k,
		blob.RetentionMode(""),
		0,
		assert.False)

	// Close and reopen the repo to make sure it's the same.
	err = w.Close(ctx)
	require.NoError(t, err, clues.ToCore(err))

	k.Close(ctx)
	require.NoError(t, err, clues.ToCore(err))

	err = k.Connect(ctx, repository.Options{})
	require.NoError(t, err, clues.ToCore(err))

	defer k.Close(ctx)

	checkRetentionParams(
		t,
		ctx,
		k,
		blob.RetentionMode(""),
		0,
		assert.False)
}

// ---------------
// integration tests that require object locking to be enabled on the bucket.
// ---------------
func mustGetBlobConfig(t *testing.T, c *conn) format.BlobStorageConfiguration {
	require.Implements(t, (*repo.DirectRepository)(nil), c.Repository)
	dr := c.Repository.(repo.DirectRepository)

	blobCfg, err := dr.FormatManager().BlobCfgBlob()
	require.NoError(t, err, "getting repo config blob")

	return blobCfg
}

func checkRetentionParams(
	t *testing.T,
	ctx context.Context, //revive:disable-line:context-as-argument
	c *conn,
	expectMode blob.RetentionMode,
	expectDuration time.Duration,
	expectExtend assert.BoolAssertionFunc,
) {
	blobCfg := mustGetBlobConfig(t, c)

	assert.Equal(t, expectMode, blobCfg.RetentionMode, "retention mode")
	// Empty mode isn't considered valid so only check if it's non-empty.
	if len(blobCfg.RetentionMode) > 0 {
		assert.True(t, blobCfg.RetentionMode.IsValid(), "valid retention mode")
	}

	assert.Equal(t, expectDuration, blobCfg.RetentionPeriod, "retention duration")

	params, err := maintenance.GetParams(ctx, c)
	require.NoError(t, err, "getting maintenance config")

	expectExtend(t, params.ExtendObjectLocks, "extend object locks")
}

// mustReopen closes and reopens the connection that w uses. Assumes no other
// structs besides w are holding a reference to the conn that w has.
//
//revive:disable-next-line:context-as-argument
func mustReopen(t *testing.T, ctx context.Context, w *Wrapper) {
	k := w.c

	err := w.Close(ctx)
	require.NoError(t, err, "closing wrapper: %v", clues.ToCore(err))

	err = k.Close(ctx)
	require.NoError(t, err, "closing conn: %v", clues.ToCore(err))

	err = k.Connect(ctx, repository.Options{})
	require.NoError(t, err, "reconnecting conn: %v", clues.ToCore(err))

	w.c = k
}

type RetentionIntegrationSuite struct {
	tester.Suite
}

func TestRetentionIntegrationSuite(t *testing.T) {
	suite.Run(t, &RetentionIntegrationSuite{
		Suite: tester.NewRetentionSuite(
			t,
			[][]string{storeTD.AWSStorageCredEnvs}),
	})
}

func (suite *RetentionIntegrationSuite) TestSetRetentionParameters() {
	table := []struct {
		name           string
		opts           repository.Retention
		expectErr      assert.ErrorAssertionFunc
		expectMode     blob.RetentionMode
		expectDuration time.Duration
		expectExtend   assert.BoolAssertionFunc
	}{
		{
			name:         "NoChanges",
			opts:         repository.Retention{},
			expectErr:    assert.NoError,
			expectExtend: assert.False,
		},
		{
			name: "UpdateMode",
			opts: repository.Retention{
				Mode: ptr.To(repository.GovernanceRetention),
			},
			expectErr:    assert.Error,
			expectExtend: assert.False,
		},
		{
			name: "UpdateDuration",
			opts: repository.Retention{
				Duration: ptr.To(time.Hour * 48),
			},
			expectErr:    assert.Error,
			expectExtend: assert.False,
		},
		{
			name: "UpdateExtend",
			opts: repository.Retention{
				Extend: ptr.To(true),
			},
			expectErr:    assert.NoError,
			expectExtend: assert.True,
		},
		{
			name: "UpdateModeAndDuration_Governance",
			opts: repository.Retention{
				Mode:     ptr.To(repository.GovernanceRetention),
				Duration: ptr.To(time.Hour * 48),
			},
			expectErr:      assert.NoError,
			expectMode:     blob.Governance,
			expectDuration: time.Hour * 48,
			expectExtend:   assert.False,
		},
		// Skip for now since compliance mode won't let us delete the blobs at all
		// until they expire.
		//{
		//  name: "UpdateModeAndDuration_Compliance",
		//  opts: repository.Retention{
		//    Mode: ptr.To(repository.ComplianceRetention),
		//    Duration: ptr.To(time.Hour * 48),
		//  },
		//  expectErr: assert.NoError,
		//  expectMode: blob.Compliance,
		//  expectDuration: time.Hour * 48,
		//  expectExtend: assert.False,
		//},
		{
			name: "UpdateModeAndDuration_Invalid",
			opts: repository.Retention{
				Mode:     ptr.To(repository.RetentionMode(-1)),
				Duration: ptr.To(time.Hour * 48),
			},
			expectErr:    assert.Error,
			expectExtend: assert.False,
		},
		{
			name: "UpdateMode_NoRetention",
			opts: repository.Retention{
				Mode: ptr.To(repository.NoRetention),
			},
			expectErr:    assert.NoError,
			expectExtend: assert.False,
		},
		{
			name: "UpdateModeAndDuration_NoRetention",
			opts: repository.Retention{
				Mode:     ptr.To(repository.NoRetention),
				Duration: ptr.To(time.Hour * 48),
			},
			expectErr:    assert.Error,
			expectExtend: assert.False,
		},
		{
			name: "UpdateModeAndDurationAndExtend",
			opts: repository.Retention{
				Mode:     ptr.To(repository.GovernanceRetention),
				Duration: ptr.To(time.Hour * 48),
				Extend:   ptr.To(true),
			},
			expectErr:      assert.NoError,
			expectMode:     blob.Governance,
			expectDuration: time.Hour * 48,
			expectExtend:   assert.True,
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			k, err := openKopiaRepo(t, ctx)
			require.NoError(t, err, clues.ToCore(err))

			w := &Wrapper{k}

			err = w.SetRetentionParameters(ctx, test.opts)
			test.expectErr(t, err, clues.ToCore(err))

			checkRetentionParams(
				t,
				ctx,
				k,
				test.expectMode,
				test.expectDuration,
				test.expectExtend)
		})
	}
}

func (suite *RetentionIntegrationSuite) TestSetRetentionParameters_And_Maintenance() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	k, err := openKopiaRepo(t, ctx)
	require.NoError(t, err, clues.ToCore(err))

	w := &Wrapper{k}

	mOpts := repository.Maintenance{
		Safety: repository.FullMaintenanceSafety,
		Type:   repository.MetadataMaintenance,
	}

	// This will set common maintenance config parameters. There's some interplay
	// between the maintenance schedule and retention period that we want to check
	// below.
	err = w.RepoMaintenance(ctx, mOpts)
	require.NoError(t, err, clues.ToCore(err))

	// Enable retention.
	err = w.SetRetentionParameters(ctx, repository.Retention{
		Mode:     ptr.To(repository.GovernanceRetention),
		Duration: ptr.To(time.Hour * 48),
		Extend:   ptr.To(true),
	})
	require.NoError(t, err, clues.ToCore(err))

	checkRetentionParams(
		t,
		ctx,
		k,
		blob.Governance,
		time.Hour*48,
		assert.True)

	// Change retention duration without updating mode.
	err = w.SetRetentionParameters(ctx, repository.Retention{
		Duration: ptr.To(time.Hour * 49),
	})
	require.NoError(t, err, clues.ToCore(err))

	checkRetentionParams(
		t,
		ctx,
		k,
		blob.Governance,
		time.Hour*49,
		assert.True)

	// Disable retention.
	err = w.SetRetentionParameters(ctx, repository.Retention{
		Mode: ptr.To(repository.NoRetention),
	})
	require.NoError(t, err, clues.ToCore(err))

	checkRetentionParams(
		t,
		ctx,
		k,
		blob.RetentionMode(""),
		0,
		assert.True)

	// Disable object lock extension.
	err = w.SetRetentionParameters(ctx, repository.Retention{
		Extend: ptr.To(false),
	})
	require.NoError(t, err, clues.ToCore(err))

	checkRetentionParams(
		t,
		ctx,
		k,
		blob.RetentionMode(""),
		0,
		assert.False)
}

func (suite *RetentionIntegrationSuite) TestSetAndUpdateRetentionParameters_RunMaintenance() {
	table := []struct {
		name   string
		reopen bool
	}{
		{
			// Check that in the same connection we can create a repo, set and then
			// update the retention period, and run full maintenance to extend object
			// locks.
			name: "SameConnection",
		},
		{
			// Test that even if the retention configuration change is done from a
			// different repo connection that we still can extend the object locking
			// duration and run maintenance successfully.
			name:   "ReopenToReconfigure",
			reopen: true,
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			k, err := openKopiaRepo(t, ctx)
			require.NoError(t, err, clues.ToCore(err))

			w := &Wrapper{k}

			mOpts := repository.Maintenance{
				Safety: repository.FullMaintenanceSafety,
				Type:   repository.CompleteMaintenance,
			}

			// This will set common maintenance config parameters. There's some interplay
			// between the maintenance schedule and retention period that we want to check
			// below.
			err = w.RepoMaintenance(ctx, mOpts)
			require.NoError(t, err, clues.ToCore(err))

			// Enable retention.
			err = w.SetRetentionParameters(ctx, repository.Retention{
				Mode:     ptr.To(repository.GovernanceRetention),
				Duration: ptr.To(time.Hour * 48),
				Extend:   ptr.To(true),
			})
			require.NoError(t, err, clues.ToCore(err))

			checkRetentionParams(
				t,
				ctx,
				k,
				blob.Governance,
				time.Hour*48,
				assert.True)

			if test.reopen {
				mustReopen(t, ctx, w)
			}

			// Change retention duration without updating mode.
			err = w.SetRetentionParameters(ctx, repository.Retention{
				Duration: ptr.To(time.Hour * 96),
			})
			require.NoError(t, err, clues.ToCore(err))

			checkRetentionParams(
				t,
				ctx,
				k,
				blob.Governance,
				time.Hour*96,
				assert.True)

			// Run full maintenance again. This should extend object locks for things if
			// they exist.
			err = w.RepoMaintenance(ctx, mOpts)
			require.NoError(t, err, clues.ToCore(err))
		})
	}
}

// ---------------
// integration tests that use kopia and initialize a repo
// ---------------
type KopiaIntegrationSuite struct {
	tester.Suite
	w     *Wrapper
	ctx   context.Context
	flush func()

	storePath1 path.Path
	storePath2 path.Path
	locPath1   path.Path
	locPath2   path.Path
}

func TestKopiaIntegrationSuite(t *testing.T) {
	suite.Run(t, &KopiaIntegrationSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{storeTD.AWSStorageCredEnvs}),
	})
}

func (suite *KopiaIntegrationSuite) SetupSuite() {
	tmp, err := path.Build(
		testTenant,
		testUser,
		path.ExchangeService,
		path.EmailCategory,
		false,
		testInboxDir)
	require.NoError(suite.T(), err, clues.ToCore(err))

	suite.storePath1 = tmp
	suite.locPath1 = tmp

	tmp, err = path.Build(
		testTenant,
		testUser,
		path.ExchangeService,
		path.EmailCategory,
		false,
		testArchiveDir)
	require.NoError(suite.T(), err, clues.ToCore(err))

	suite.storePath2 = tmp
	suite.locPath2 = tmp
}

func (suite *KopiaIntegrationSuite) SetupTest() {
	t := suite.T()
	suite.ctx, suite.flush = tester.NewContext(t)

	c, err := openKopiaRepo(t, suite.ctx)
	require.NoError(t, err, clues.ToCore(err))

	suite.w = &Wrapper{c}
}

func (suite *KopiaIntegrationSuite) TearDownTest() {
	defer suite.flush()

	err := suite.w.Close(suite.ctx)
	assert.NoError(suite.T(), err, clues.ToCore(err))
}

func (suite *KopiaIntegrationSuite) TestBackupCollections() {
	c1 := exchMock.NewCollection(
		suite.storePath1,
		suite.locPath1,
		5)
	// Add a 4k chunk of data that should be compressible. This helps check
	// compression is enabled because we do some testing on the number of bytes
	// uploaded during the first backup.
	c1.Data[0] = []byte(strings.Repeat("abcdefgh", 512))

	collections := []data.BackupCollection{
		c1,
		exchMock.NewCollection(
			suite.storePath2,
			suite.locPath2,
			42),
	}

	c1 = exchMock.NewCollection(
		suite.storePath1,
		suite.locPath1,
		0)
	c1.ColState = data.NotMovedState
	c1.PrevPath = suite.storePath1

	c2 := exchMock.NewCollection(
		suite.storePath2,
		suite.locPath2,
		0)
	c2.ColState = data.NotMovedState
	c2.PrevPath = suite.storePath2

	// Make empty collections at the same locations to force a backup with no
	// changes. Needed to ensure we force a backup even if nothing has changed.
	emptyCollections := []data.BackupCollection{c1, c2}

	// tags that are supplied by the caller. This includes basic tags to support
	// lookups and extra tags the caller may want to apply.
	tags := map[string]string{
		"fnords":    "smarf",
		"brunhilda": "",
	}

	reasons := []identity.Reasoner{
		NewReason(
			testTenant,
			suite.storePath1.ResourceOwner(),
			suite.storePath1.Service(),
			suite.storePath1.Category()),
		NewReason(
			testTenant,
			suite.storePath2.ResourceOwner(),
			suite.storePath2.Service(),
			suite.storePath2.Category()),
	}

	expectedTags := map[string]string{}

	maps.Copy(expectedTags, tags)

	for _, r := range reasons {
		for _, k := range tagKeys(r) {
			expectedTags[k] = ""
		}
	}

	expectedTags = normalizeTagKVs(expectedTags)

	type testCase struct {
		name                  string
		baseBackups           func(base ManifestEntry) BackupBases
		collections           []data.BackupCollection
		expectedUploadedFiles int
		expectedCachedFiles   int
		// We're either going to get details entries or entries in the details
		// merger. Details is populated when there's entries in the collection. The
		// details merger is populated for cached entries. The details merger
		// doesn't count folders, only items.
		//
		// Setting this to true looks for details merger entries. Setting it to
		// false looks for details entries.
		expectMerge bool
		// Whether entries in the resulting details should be marked as updated.
		deetsUpdated     assert.BoolAssertionFunc
		hashedBytesCheck assert.ValueAssertionFunc
		// Range of bytes (inclusive) to expect as uploaded. A little fragile, but
		// allows us to differentiate between content that wasn't uploaded due to
		// being cached/deduped/skipped due to existing dir entries and stuff that
		// was actually pushed to S3.
		uploadedBytes []int64
	}

	// Initial backup. All files should be considered new by kopia.
	baseBackupCase := testCase{
		name: "Uncached",
		baseBackups: func(ManifestEntry) BackupBases {
			return NewMockBackupBases()
		},
		collections:           collections,
		expectedUploadedFiles: 47,
		expectedCachedFiles:   0,
		deetsUpdated:          assert.True,
		hashedBytesCheck:      assert.NotZero,
		uploadedBytes:         []int64{8000, 10000},
	}

	runAndTestBackup := func(test testCase, base ManifestEntry) ManifestEntry {
		var res ManifestEntry

		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			bbs := test.baseBackups(base)

			stats, deets, deetsMerger, err := suite.w.ConsumeBackupCollections(
				ctx,
				reasons,
				bbs,
				test.collections,
				nil,
				tags,
				true,
				fault.New(true))
			require.NoError(t, err, clues.ToCore(err))

			assert.Equal(t, test.expectedUploadedFiles, stats.TotalFileCount, "total files")
			assert.Equal(t, test.expectedUploadedFiles, stats.UncachedFileCount, "uncached files")
			assert.Equal(t, test.expectedCachedFiles, stats.CachedFileCount, "cached files")
			assert.Equal(t, 4+len(test.collections), stats.TotalDirectoryCount, "directory count")
			assert.Equal(t, 0, stats.IgnoredErrorCount)
			assert.Equal(t, 0, stats.ErrorCount)
			assert.False(t, stats.Incomplete)
			test.hashedBytesCheck(t, stats.TotalHashedBytes, "hashed bytes")
			assert.LessOrEqual(
				t,
				test.uploadedBytes[0],
				stats.TotalUploadedBytes,
				"low end of uploaded bytes")
			assert.GreaterOrEqual(
				t,
				test.uploadedBytes[1],
				stats.TotalUploadedBytes,
				"high end of uploaded bytes")

			if test.expectMerge {
				assert.Empty(t, deets.Details().Entries, "details entries")
				assert.Equal(
					t,
					test.expectedUploadedFiles+test.expectedCachedFiles,
					deetsMerger.ItemsToMerge(),
					"details merger entries")
			} else {
				assert.Zero(t, deetsMerger.ItemsToMerge(), "details merger entries")

				details := deets.Details().Entries
				assert.Len(
					t,
					details,
					// 47 file and 2 folder entries.
					test.expectedUploadedFiles+test.expectedCachedFiles+2)
			}

			checkSnapshotTags(
				t,
				ctx,
				suite.w.c,
				expectedTags,
				stats.SnapshotID)

			snap, err := snapshot.LoadSnapshot(
				ctx,
				suite.w.c,
				manifest.ID(stats.SnapshotID))
			require.NoError(t, err, clues.ToCore(err))

			res = ManifestEntry{
				Manifest: snap,
				Reasons:  reasons,
			}
		})

		return res
	}

	base := runAndTestBackup(baseBackupCase, ManifestEntry{})

	table := []testCase{
		{
			name: "Kopia Assist And Merge All Files Changed",
			baseBackups: func(base ManifestEntry) BackupBases {
				return NewMockBackupBases().WithMergeBases(base)
			},
			collections:           collections,
			expectedUploadedFiles: 0,
			expectedCachedFiles:   47,
			// Entries go to details merger since cached files are merged too.
			expectMerge:      true,
			deetsUpdated:     assert.False,
			hashedBytesCheck: assert.Zero,
			uploadedBytes:    []int64{4000, 6000},
		},
		{
			name: "Kopia Assist And Merge No Files Changed",
			baseBackups: func(base ManifestEntry) BackupBases {
				return NewMockBackupBases().WithMergeBases(base)
			},
			// Pass in empty collections to force a backup. Otherwise we'll skip
			// actually trying to do anything because we'll see there's nothing that
			// changed. The real goal is to get it to deal with the merged collections
			// again though.
			collections: emptyCollections,
			// Should hit cached check prior to dir entry check so we see them as
			// cached.
			expectedUploadedFiles: 0,
			expectedCachedFiles:   47,
			// Entries go into the details merger because we never materialize details
			// info for the items since they're from the base.
			expectMerge: true,
			// Not used since there's no details entries.
			deetsUpdated:     assert.False,
			hashedBytesCheck: assert.Zero,
			uploadedBytes:    []int64{4000, 6000},
		},
		{
			name: "Kopia Assist Only",
			baseBackups: func(base ManifestEntry) BackupBases {
				return NewMockBackupBases().WithAssistBases(base)
			},
			collections:           collections,
			expectedUploadedFiles: 0,
			expectedCachedFiles:   47,
			expectMerge:           true,
			deetsUpdated:          assert.False,
			hashedBytesCheck:      assert.Zero,
			uploadedBytes:         []int64{4000, 6000},
		},
		{
			name: "Merge Only",
			baseBackups: func(base ManifestEntry) BackupBases {
				return NewMockBackupBases().WithMergeBases(base).MockDisableAssistBases()
			},
			// Pass in empty collections to force a backup. Otherwise we'll skip
			// actually trying to do anything because we'll see there's nothing that
			// changed. The real goal is to get it to deal with the merged collections
			// again though.
			collections:           emptyCollections,
			expectedUploadedFiles: 47,
			expectedCachedFiles:   0,
			expectMerge:           true,
			// Not used since there's no details entries.
			deetsUpdated: assert.False,
			// Kopia still counts these bytes as "hashed" even though it shouldn't
			// read the file data since they already have dir entries it can reuse.
			hashedBytesCheck: assert.NotZero,
			uploadedBytes:    []int64{4000, 6000},
		},
		{
			name: "Content Hash Only",
			baseBackups: func(base ManifestEntry) BackupBases {
				return NewMockBackupBases()
			},
			collections:           collections,
			expectedUploadedFiles: 47,
			expectedCachedFiles:   0,
			// Marked as updated because we still fall into the uploadFile handler in
			// kopia instead of the cachedFile handler.
			deetsUpdated:     assert.True,
			hashedBytesCheck: assert.NotZero,
			uploadedBytes:    []int64{4000, 6000},
		},
	}

	for _, test := range table {
		runAndTestBackup(test, base)
	}
}

// TODO(ashmrtn): This should really be moved to an e2e test that just checks
// details for certain things.
func (suite *KopiaIntegrationSuite) TestBackupCollections_NoDetailsForMeta() {
	tmp, err := path.Build(
		testTenant,
		testUser,
		path.OneDriveService,
		path.FilesCategory,
		false,
		testInboxDir)
	require.NoError(suite.T(), err, clues.ToCore(err))

	storePath := tmp
	locPath := path.Builder{}.Append(tmp.Folders()...)

	baseOneDriveItemInfo := details.OneDriveInfo{
		ItemType:  details.OneDriveItem,
		DriveID:   "drive-id",
		DriveName: "drive-name",
		ItemName:  "item",
	}

	// tags that are supplied by the caller. This includes basic tags to support
	// lookups and extra tags the caller may want to apply.
	tags := map[string]string{
		"fnords":    "smarf",
		"brunhilda": "",
	}

	reasons := []identity.Reasoner{
		NewReason(
			testTenant,
			storePath.ResourceOwner(),
			storePath.Service(),
			storePath.Category()),
	}

	expectedTags := map[string]string{}

	maps.Copy(expectedTags, tags)

	for _, r := range reasons {
		for _, k := range tagKeys(r) {
			expectedTags[k] = ""
		}
	}

	expectedTags = normalizeTagKVs(expectedTags)

	table := []struct {
		name                  string
		expectedUploadedFiles int
		expectedCachedFiles   int
		numDeetsEntries       int
		hasMetaDeets          bool
		cols                  func() []data.BackupCollection
	}{
		{
			name:                  "Uncached",
			expectedUploadedFiles: 3,
			expectedCachedFiles:   0,
			// MockStream implements item info even though OneDrive doesn't.
			numDeetsEntries: 3,
			hasMetaDeets:    true,
			cols: func() []data.BackupCollection {
				streams := []data.Item{}
				fileNames := []string{
					testFileName,
					testFileName + metadata.MetaFileSuffix,
					metadata.DirMetaFileSuffix,
				}

				for _, name := range fileNames {
					info := baseOneDriveItemInfo
					info.ItemName = name

					ms := &dataMock.Item{
						ItemID:   name,
						Reader:   io.NopCloser(&bytes.Buffer{}),
						ItemSize: 0,
						ItemInfo: details.ItemInfo{OneDrive: &info},
					}

					streams = append(streams, ms)
				}

				mc := &m365Mock.BackupCollection{
					Path:    storePath,
					Loc:     locPath,
					Streams: streams,
				}

				return []data.BackupCollection{mc}
			},
		},
		{
			name:                  "Cached",
			expectedUploadedFiles: 1,
			expectedCachedFiles:   2,
			// Meta entries are filtered out.
			numDeetsEntries: 1,
			hasMetaDeets:    false,
			cols: func() []data.BackupCollection {
				info := baseOneDriveItemInfo
				info.ItemName = testFileName

				ms := &dataMock.Item{
					ItemID:   testFileName,
					Reader:   io.NopCloser(&bytes.Buffer{}),
					ItemSize: 0,
					ItemInfo: details.ItemInfo{OneDrive: &info},
				}

				mc := &m365Mock.BackupCollection{
					Path:    storePath,
					Loc:     locPath,
					Streams: []data.Item{ms},
					CState:  data.NotMovedState,
				}

				return []data.BackupCollection{mc}
			},
		},
	}

	prevSnaps := NewMockBackupBases()

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()
			collections := test.cols()

			stats, deets, prevShortRefs, err := suite.w.ConsumeBackupCollections(
				suite.ctx,
				reasons,
				prevSnaps,
				collections,
				nil,
				tags,
				true,
				fault.New(true))
			assert.NoError(t, err, clues.ToCore(err))

			assert.Equal(t, test.expectedUploadedFiles, stats.TotalFileCount, "total files")
			assert.Equal(t, test.expectedUploadedFiles, stats.UncachedFileCount, "uncached files")
			assert.Equal(t, test.expectedCachedFiles, stats.CachedFileCount, "cached files")
			assert.Equal(t, 5, stats.TotalDirectoryCount)
			assert.Equal(t, 0, stats.IgnoredErrorCount)
			assert.Equal(t, 0, stats.ErrorCount)
			assert.False(t, stats.Incomplete)

			// 47 file and 1 folder entries.
			details := deets.Details().Entries
			assert.Len(
				t,
				details,
				test.numDeetsEntries+1)

			for _, entry := range details {
				if test.hasMetaDeets {
					continue
				}

				assert.False(t, metadata.HasMetaSuffix(entry.RepoRef), "metadata entry in details")
			}

			// Shouldn't have any items to merge because the cached files are metadata
			// files.
			assert.Equal(t, 0, prevShortRefs.ItemsToMerge(), "merge items")

			checkSnapshotTags(
				t,
				suite.ctx,
				suite.w.c,
				expectedTags,
				stats.SnapshotID)

			snap, err := snapshot.LoadSnapshot(
				suite.ctx,
				suite.w.c,
				manifest.ID(stats.SnapshotID))
			require.NoError(t, err, clues.ToCore(err))

			prevSnaps.WithMergeBases(
				ManifestEntry{
					Manifest: snap,
					Reasons:  reasons,
				})
		})
	}
}

func (suite *KopiaIntegrationSuite) TestRestoreAfterCompressionChange() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	k, err := openKopiaRepo(t, ctx)
	require.NoError(t, err, clues.ToCore(err))

	err = k.Compression(ctx, "s2-default")
	require.NoError(t, err, clues.ToCore(err))

	w := &Wrapper{k}

	r := NewReason(testTenant, testUser, path.ExchangeService, path.EmailCategory)

	dc1 := exchMock.NewCollection(suite.storePath1, suite.locPath1, 1)
	dc2 := exchMock.NewCollection(suite.storePath2, suite.locPath2, 1)

	fp1, err := suite.storePath1.AppendItem(dc1.Names[0])
	require.NoError(t, err, clues.ToCore(err))

	fp2, err := suite.storePath2.AppendItem(dc2.Names[0])
	require.NoError(t, err, clues.ToCore(err))

	stats, _, _, err := w.ConsumeBackupCollections(
		ctx,
		[]identity.Reasoner{r},
		nil,
		[]data.BackupCollection{dc1, dc2},
		nil,
		nil,
		true,
		fault.New(true))
	require.NoError(t, err, clues.ToCore(err))

	err = k.Compression(ctx, "gzip")
	require.NoError(t, err, clues.ToCore(err))

	expected := map[string][]byte{
		fp1.String(): dc1.Data[0],
		fp2.String(): dc2.Data[0],
	}

	result, err := w.ProduceRestoreCollections(
		ctx,
		string(stats.SnapshotID),
		toRestorePaths(t, fp1, fp2),
		nil,
		fault.New(true))
	require.NoError(t, err, clues.ToCore(err))
	assert.Equal(t, 2, len(result))

	testForFiles(t, ctx, expected, result)
}

func (suite *KopiaIntegrationSuite) TestBackupCollections_ReaderError() {
	t := suite.T()

	loc1 := path.Builder{}.Append(suite.storePath1.Folders()...)
	loc2 := path.Builder{}.Append(suite.storePath2.Folders()...)
	r := NewReason(testTenant, testUser, path.ExchangeService, path.EmailCategory)

	collections := []data.BackupCollection{
		&m365Mock.BackupCollection{
			Path: suite.storePath1,
			Loc:  loc1,
			Streams: []data.Item{
				&dataMock.Item{
					ItemID:   testFileName,
					Reader:   io.NopCloser(bytes.NewReader(testFileData)),
					ItemInfo: exchMock.StubMailInfo(),
				},
				&dataMock.Item{
					ItemID:   testFileName2,
					Reader:   io.NopCloser(bytes.NewReader(testFileData2)),
					ItemInfo: exchMock.StubMailInfo(),
				},
			},
		},
		&m365Mock.BackupCollection{
			Path: suite.storePath2,
			Loc:  loc2,
			Streams: []data.Item{
				&dataMock.Item{
					ItemID:   testFileName3,
					Reader:   io.NopCloser(bytes.NewReader(testFileData3)),
					ItemInfo: exchMock.StubMailInfo(),
				},
				&dataMock.Item{
					ItemID:   testFileName4,
					ReadErr:  assert.AnError,
					ItemInfo: exchMock.StubMailInfo(),
				},
				&dataMock.Item{
					ItemID:   testFileName5,
					Reader:   io.NopCloser(bytes.NewReader(testFileData5)),
					ItemInfo: exchMock.StubMailInfo(),
				},
				&dataMock.Item{
					ItemID:   testFileName6,
					Reader:   io.NopCloser(bytes.NewReader(testFileData6)),
					ItemInfo: exchMock.StubMailInfo(),
				},
			},
		},
	}

	stats, deets, _, err := suite.w.ConsumeBackupCollections(
		suite.ctx,
		[]identity.Reasoner{r},
		nil,
		collections,
		nil,
		nil,
		true,
		fault.New(true))
	require.Error(t, err, clues.ToCore(err))
	assert.Equal(t, 0, stats.ErrorCount)
	assert.Equal(t, 5, stats.TotalFileCount)
	assert.Equal(t, 6, stats.TotalDirectoryCount)
	assert.Equal(t, 1, stats.IgnoredErrorCount)
	assert.False(t, stats.Incomplete)
	// 5 file and 2 folder entries.
	assert.Len(t, deets.Details().Entries, 5+2)

	failedPath, err := suite.storePath2.AppendItem(testFileName4)
	require.NoError(t, err, clues.ToCore(err))

	ic := i64counter{}

	dcs, err := suite.w.ProduceRestoreCollections(
		suite.ctx,
		string(stats.SnapshotID),
		toRestorePaths(t, failedPath),
		&ic,
		fault.New(true))
	assert.NoError(t, err, "error producing restore collections")

	require.Len(t, dcs, 1, "number of restore collections")

	errs := fault.New(true)
	items := dcs[0].Items(suite.ctx, errs)

	// Get all the items from channel
	//nolint:revive
	for range items {
	}

	// Files that had an error shouldn't make a dir entry in kopia. If they do we
	// may run into kopia-assisted incrementals issues because only mod time and
	// not file size is checked for StreamingFiles.
	assert.ErrorIs(t, errs.Failure(), data.ErrNotFound, "errored file is restorable", clues.ToCore(err))
}

type backedupFile struct {
	parentPath path.Path
	itemPath   path.Path
	data       []byte
}

func (suite *KopiaIntegrationSuite) TestBackupCollectionsHandlesNoCollections() {
	table := []struct {
		name        string
		collections []data.BackupCollection
	}{
		{
			name:        "NilCollections",
			collections: nil,
		},
		{
			name:        "EmptyCollections",
			collections: []data.BackupCollection{},
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			s, d, _, err := suite.w.ConsumeBackupCollections(
				ctx,
				nil,
				nil,
				test.collections,
				nil,
				nil,
				true,
				fault.New(true))
			require.NoError(t, err, clues.ToCore(err))

			assert.Equal(t, BackupStats{}, *s)
			assert.Empty(t, d.Details().Entries)
		})
	}
}

type KopiaSimpleRepoIntegrationSuite struct {
	tester.Suite
	w          *Wrapper
	ctx        context.Context
	snapshotID manifest.ID

	testPath1 path.Path
	testPath2 path.Path

	// List of files per parent directory.
	files map[string][]*backedupFile
	// Set of files by file path.
	filesByPath map[string]*backedupFile
}

func TestKopiaSimpleRepoIntegrationSuite(t *testing.T) {
	suite.Run(t, &KopiaSimpleRepoIntegrationSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{storeTD.AWSStorageCredEnvs}),
	})
}

func (suite *KopiaSimpleRepoIntegrationSuite) SetupSuite() {
	tmp, err := path.Build(
		testTenant,
		testUser,
		path.ExchangeService,
		path.EmailCategory,
		false,
		testInboxDir)
	require.NoError(suite.T(), err, clues.ToCore(err))

	suite.testPath1 = tmp

	tmp, err = path.Build(
		testTenant,
		testUser,
		path.ExchangeService,
		path.EmailCategory,
		false,
		testArchiveDir)
	require.NoError(suite.T(), err, clues.ToCore(err))

	suite.testPath2 = tmp

	suite.files = map[string][]*backedupFile{}
	suite.filesByPath = map[string]*backedupFile{}

	filesInfo := []struct {
		parentPath path.Path
		name       string
		data       []byte
	}{
		{
			parentPath: suite.testPath1,
			name:       testFileName,
			data:       testFileData,
		},
		{
			parentPath: suite.testPath1,
			name:       testFileName2,
			data:       testFileData2,
		},
		{
			parentPath: suite.testPath2,
			name:       testFileName3,
			data:       testFileData3,
		},
		{
			parentPath: suite.testPath2,
			name:       testFileName4,
			data:       testFileData4,
		},
		{
			parentPath: suite.testPath2,
			name:       testFileName5,
			data:       testFileData5,
		},
		{
			parentPath: suite.testPath2,
			name:       testFileName6,
			data:       testFileData6,
		},
	}

	for _, item := range filesInfo {
		pth, err := item.parentPath.AppendItem(item.name)
		require.NoError(suite.T(), err, clues.ToCore(err))

		mapKey := item.parentPath.String()
		f := &backedupFile{
			parentPath: item.parentPath,
			itemPath:   pth,
			data:       item.data,
		}

		suite.files[mapKey] = append(suite.files[mapKey], f)
		suite.filesByPath[pth.String()] = f
	}
}

func (suite *KopiaSimpleRepoIntegrationSuite) SetupTest() {
	t := suite.T()
	expectedDirs := 6
	expectedFiles := len(suite.filesByPath)

	ls := logger.Settings{
		Level:  logger.LLDebug,
		Format: logger.LFText,
	}
	//nolint:forbidigo
	suite.ctx, _ = logger.CtxOrSeed(context.Background(), ls)

	c, err := openKopiaRepo(t, suite.ctx)
	require.NoError(t, err, clues.ToCore(err))

	suite.w = &Wrapper{c}

	collections := []data.BackupCollection{}

	for _, parent := range []path.Path{suite.testPath1, suite.testPath2} {
		loc := path.Builder{}.Append(parent.Folders()...)
		collection := &m365Mock.BackupCollection{Path: parent, Loc: loc}

		for _, item := range suite.files[parent.String()] {
			collection.Streams = append(
				collection.Streams,
				&dataMock.Item{
					ItemID:   item.itemPath.Item(),
					Reader:   io.NopCloser(bytes.NewReader(item.data)),
					ItemInfo: exchMock.StubMailInfo(),
				})
		}

		collections = append(collections, collection)
	}

	r := NewReason(testTenant, testUser, path.ExchangeService, path.EmailCategory)

	stats, deets, _, err := suite.w.ConsumeBackupCollections(
		suite.ctx,
		[]identity.Reasoner{r},
		nil,
		collections,
		nil,
		nil,
		false,
		fault.New(true))
	require.NoError(t, err, clues.ToCore(err))
	require.Equal(t, stats.ErrorCount, 0)
	require.Equal(t, stats.TotalFileCount, expectedFiles)
	require.Equal(t, stats.TotalDirectoryCount, expectedDirs)
	require.Equal(t, stats.IgnoredErrorCount, 0)
	require.False(t, stats.Incomplete)
	// 6 file and 2 folder entries.
	assert.Len(t, deets.Details().Entries, expectedFiles+2)

	suite.snapshotID = manifest.ID(stats.SnapshotID)
}

func (suite *KopiaSimpleRepoIntegrationSuite) TearDownTest() {
	err := suite.w.Close(suite.ctx)
	assert.NoError(suite.T(), err, clues.ToCore(err))
	logger.Flush(suite.ctx)
}

type i64counter struct {
	i int64
}

func (c *i64counter) Count(i int64) {
	c.i += i
}

func (suite *KopiaSimpleRepoIntegrationSuite) TestBackupExcludeItem() {
	r := NewReason(testTenant, testUser, path.ExchangeService, path.EmailCategory)

	man, err := suite.w.c.LoadSnapshot(suite.ctx, suite.snapshotID)
	require.NoError(suite.T(), err, "getting base snapshot: %v", clues.ToCore(err))

	table := []struct {
		name                  string
		excludeItem           bool
		excludePrefix         bool
		expectedCachedItems   int
		expectedUncachedItems int
		cols                  func() []data.BackupCollection
		backupIDCheck         require.ValueAssertionFunc
		restoreCheck          assert.ErrorAssertionFunc
	}{
		{
			name:                  "ExcludeItem_NoPrefix",
			excludeItem:           true,
			expectedCachedItems:   len(suite.filesByPath) - 1,
			expectedUncachedItems: 0,
			cols: func() []data.BackupCollection {
				return nil
			},
			backupIDCheck: require.NotEmpty,
			restoreCheck:  assert.Error,
		},
		{
			name:                  "ExcludeItem_WithPrefix",
			excludeItem:           true,
			excludePrefix:         true,
			expectedCachedItems:   len(suite.filesByPath) - 1,
			expectedUncachedItems: 0,
			cols: func() []data.BackupCollection {
				return nil
			},
			backupIDCheck: require.NotEmpty,
			restoreCheck:  assert.Error,
		},
		{
			name: "NoExcludeItemNoChanges",
			// No snapshot should be made since there were no changes.
			expectedCachedItems:   0,
			expectedUncachedItems: 0,
			cols: func() []data.BackupCollection {
				return nil
			},
			// Backup doesn't run.
			backupIDCheck: require.Empty,
		},
		{
			name:                  "NoExcludeItemWithChanges",
			expectedCachedItems:   len(suite.filesByPath),
			expectedUncachedItems: 1,
			cols: func() []data.BackupCollection {
				c := exchMock.NewCollection(
					suite.testPath1,
					suite.testPath1,
					1)
				c.ColState = data.NotMovedState
				c.PrevPath = suite.testPath1

				return []data.BackupCollection{c}
			},
			backupIDCheck: require.NotEmpty,
			restoreCheck:  assert.NoError,
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			var (
				prefix   string
				itemPath = suite.files[suite.testPath1.String()][0].itemPath
			)

			if !test.excludePrefix {
				prefix = itemPath.ToBuilder().Dir().Dir().String()
			}

			excluded := pmMock.NewPrefixMap(nil)
			if test.excludeItem {
				excluded = pmMock.NewPrefixMap(map[string]map[string]struct{}{
					// Add a prefix if needed.
					prefix: {
						itemPath.Item(): {},
					},
				})
			}

			stats, _, _, err := suite.w.ConsumeBackupCollections(
				suite.ctx,
				[]identity.Reasoner{r},
				NewMockBackupBases().WithMergeBases(
					ManifestEntry{
						Manifest: man,
						Reasons:  []identity.Reasoner{r},
					}),
				test.cols(),
				excluded,
				nil,
				true,
				fault.New(true))
			require.NoError(t, err, clues.ToCore(err))
			assert.Equal(t, test.expectedCachedItems, stats.CachedFileCount)
			assert.Equal(t, test.expectedUncachedItems, stats.UncachedFileCount)

			test.backupIDCheck(t, stats.SnapshotID)

			if len(stats.SnapshotID) == 0 {
				return
			}

			ic := i64counter{}

			dcs, err := suite.w.ProduceRestoreCollections(
				suite.ctx,
				string(stats.SnapshotID),
				toRestorePaths(t, suite.files[suite.testPath1.String()][0].itemPath),
				&ic,
				fault.New(true))

			assert.NoError(t, err, "errors producing collection", clues.ToCore(err))
			require.Len(t, dcs, 1, "unexpected number of restore collections")

			errs := fault.New(true)
			items := dcs[0].Items(suite.ctx, errs)

			// Get all the items from channel
			//nolint:revive
			for range items {
			}

			test.restoreCheck(t, errs.Failure(), errs)
		})
	}
}

func (suite *KopiaSimpleRepoIntegrationSuite) TestProduceRestoreCollections() {
	doesntExist, err := path.Build(
		testTenant,
		testUser,
		path.ExchangeService,
		path.EmailCategory,
		true,
		"subdir", "foo")
	require.NoError(suite.T(), err)

	// Expected items is generated during the test by looking up paths in the
	// suite's map of files. Files that are not in the suite's map are assumed to
	// generate errors and not be in the output.
	table := []struct {
		name                  string
		inputPaths            []path.Path
		expectedCollections   int
		expectedErr           assert.ErrorAssertionFunc
		expectedCollectionErr assert.ErrorAssertionFunc
	}{
		{
			name: "SingleItem",
			inputPaths: []path.Path{
				suite.files[suite.testPath1.String()][0].itemPath,
			},
			expectedCollections:   1,
			expectedErr:           assert.NoError,
			expectedCollectionErr: assert.NoError,
		},
		{
			name: "MultipleItemsSameCollection",
			inputPaths: []path.Path{
				suite.files[suite.testPath1.String()][0].itemPath,
				suite.files[suite.testPath1.String()][1].itemPath,
			},
			expectedCollections:   1,
			expectedErr:           assert.NoError,
			expectedCollectionErr: assert.NoError,
		},
		{
			name: "MultipleItemsDifferentCollections",
			inputPaths: []path.Path{
				suite.files[suite.testPath1.String()][0].itemPath,
				suite.files[suite.testPath2.String()][0].itemPath,
			},
			expectedCollections:   2,
			expectedErr:           assert.NoError,
			expectedCollectionErr: assert.NoError,
		},
		{
			name: "TargetNotAFile",
			inputPaths: []path.Path{
				suite.files[suite.testPath1.String()][0].itemPath,
				suite.testPath1,
				suite.files[suite.testPath2.String()][0].itemPath,
			},
			expectedCollections:   0,
			expectedErr:           assert.Error,
			expectedCollectionErr: assert.NoError,
		},
		{
			name: "NonExistentFile",
			inputPaths: []path.Path{
				suite.files[suite.testPath1.String()][0].itemPath,
				doesntExist,
				suite.files[suite.testPath2.String()][0].itemPath,
			},
			expectedCollections:   0,
			expectedErr:           assert.NoError,
			expectedCollectionErr: assert.Error, // folder for doesntExist does not exist
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			// May slightly overallocate as only items that are actually in our map
			// are expected. The rest are errors, but best-effort says it should carry
			// on even then.
			expected := make(map[string][]byte, len(test.inputPaths))

			for _, pth := range test.inputPaths {
				item, ok := suite.filesByPath[pth.String()]
				if !ok {
					continue
				}

				expected[pth.String()] = item.data
			}

			ic := i64counter{}

			result, err := suite.w.ProduceRestoreCollections(
				suite.ctx,
				string(suite.snapshotID),
				toRestorePaths(t, test.inputPaths...),
				&ic,
				fault.New(true))
			test.expectedCollectionErr(t, err, clues.ToCore(err), "producing collections")

			if err != nil {
				return
			}

			errs := fault.New(true)

			for _, dc := range result {
				// Get all the items from channel
				items := dc.Items(suite.ctx, errs)
				//nolint:revive
				for range items {
				}
			}

			test.expectedErr(t, errs.Failure(), errs.Failure(), "getting items")

			if errs.Failure() != nil {
				return
			}

			assert.Len(t, result, test.expectedCollections)
			assert.Less(t, int64(0), ic.i)
			testForFiles(t, ctx, expected, result)
		})
	}
}

// TestProduceRestoreCollections_PathChanges tests that having different
// Restore and Storage paths works properly. Having the same Restore and Storage
// paths is tested by TestProduceRestoreCollections.
func (suite *KopiaSimpleRepoIntegrationSuite) TestProduceRestoreCollections_PathChanges() {
	rp1, err := path.Build(
		testTenant,
		testUser,
		path.ExchangeService,
		path.EmailCategory,
		false,
		"corso_restore", "Inbox")
	require.NoError(suite.T(), err)

	rp2, err := path.Build(
		testTenant,
		testUser,
		path.ExchangeService,
		path.EmailCategory,
		false,
		"corso_restore", "Archive")
	require.NoError(suite.T(), err)

	// Expected items is generated during the test by looking up paths in the
	// suite's map of files.
	table := []struct {
		name                string
		inputPaths          []path.RestorePaths
		expectedCollections int
	}{
		{
			name: "SingleItem",
			inputPaths: []path.RestorePaths{
				{
					StoragePath: suite.files[suite.testPath1.String()][0].itemPath,
					RestorePath: rp1,
				},
			},
			expectedCollections: 1,
		},
		{
			name: "MultipleItemsSameCollection",
			inputPaths: []path.RestorePaths{
				{
					StoragePath: suite.files[suite.testPath1.String()][0].itemPath,
					RestorePath: rp1,
				},
				{
					StoragePath: suite.files[suite.testPath1.String()][1].itemPath,
					RestorePath: rp1,
				},
			},
			expectedCollections: 1,
		},
		{
			name: "MultipleItemsDifferentCollections",
			inputPaths: []path.RestorePaths{
				{
					StoragePath: suite.files[suite.testPath1.String()][0].itemPath,
					RestorePath: rp1,
				},
				{
					StoragePath: suite.files[suite.testPath2.String()][0].itemPath,
					RestorePath: rp2,
				},
			},
			expectedCollections: 2,
		},
		{
			name: "Multiple Items From Different Collections To Same Collection",
			inputPaths: []path.RestorePaths{
				{
					StoragePath: suite.files[suite.testPath1.String()][0].itemPath,
					RestorePath: rp1,
				},
				{
					StoragePath: suite.files[suite.testPath2.String()][0].itemPath,
					RestorePath: rp1,
				},
			},
			expectedCollections: 1,
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			expected := make(map[string][]byte, len(test.inputPaths))

			for _, pth := range test.inputPaths {
				item, ok := suite.filesByPath[pth.StoragePath.String()]
				require.True(t, ok, "getting expected file data")

				itemPath, err := pth.RestorePath.AppendItem(pth.StoragePath.Item())
				require.NoError(t, err, "getting expected item path")

				expected[itemPath.String()] = item.data
			}

			ic := i64counter{}

			result, err := suite.w.ProduceRestoreCollections(
				suite.ctx,
				string(suite.snapshotID),
				test.inputPaths,
				&ic,
				fault.New(true))
			require.NoError(t, err, clues.ToCore(err))

			assert.Len(t, result, test.expectedCollections)
			testForFiles(t, ctx, expected, result)
		})
	}
}

// TestProduceRestoreCollections_Fetch tests that the Fetch function still works
// properly even with different Restore and Storage paths and items from
// different kopia directories.
func (suite *KopiaSimpleRepoIntegrationSuite) TestProduceRestoreCollections_FetchItemByName() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	rp1, err := path.Build(
		testTenant,
		testUser,
		path.ExchangeService,
		path.EmailCategory,
		false,
		"corso_restore", "Inbox")
	require.NoError(suite.T(), err)

	inputPaths := []path.RestorePaths{
		{
			StoragePath: suite.files[suite.testPath1.String()][0].itemPath,
			RestorePath: rp1,
		},
		{
			StoragePath: suite.files[suite.testPath2.String()][0].itemPath,
			RestorePath: rp1,
		},
	}

	// Really only interested in getting the collection so we can call fetch on
	// it.
	ic := i64counter{}

	result, err := suite.w.ProduceRestoreCollections(
		suite.ctx,
		string(suite.snapshotID),
		inputPaths,
		&ic,
		fault.New(true))
	require.NoError(t, err, "getting collection", clues.ToCore(err))
	require.Len(t, result, 1)

	// Item from first kopia directory.
	f := suite.files[suite.testPath1.String()][0]

	item, err := result[0].FetchItemByName(ctx, f.itemPath.Item())
	require.NoError(t, err, "fetching file", clues.ToCore(err))

	r := item.ToReader()

	buf, err := io.ReadAll(r)
	require.NoError(t, err, "reading file data", clues.ToCore(err))

	assert.Equal(t, f.data, buf)

	// Item from second kopia directory.
	f = suite.files[suite.testPath2.String()][0]

	item, err = result[0].FetchItemByName(ctx, f.itemPath.Item())
	require.NoError(t, err, "fetching file", clues.ToCore(err))

	r = item.ToReader()

	buf, err = io.ReadAll(r)
	require.NoError(t, err, "reading file data", clues.ToCore(err))

	assert.Equal(t, f.data, buf)
}

func (suite *KopiaSimpleRepoIntegrationSuite) TestProduceRestoreCollections_Errors() {
	itemPath, err := suite.testPath1.AppendItem(testFileName)
	require.NoError(suite.T(), err, clues.ToCore(err))

	table := []struct {
		name       string
		snapshotID string
		paths      []path.RestorePaths
	}{
		{
			"NilPaths",
			string(suite.snapshotID),
			nil,
		},
		{
			"EmptyPaths",
			string(suite.snapshotID),
			[]path.RestorePaths{},
		},
		{
			"NoSnapshot",
			"foo",
			toRestorePaths(suite.T(), itemPath),
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			c, err := suite.w.ProduceRestoreCollections(
				suite.ctx,
				test.snapshotID,
				test.paths,
				nil,
				fault.New(true))
			assert.Error(t, err, clues.ToCore(err))
			assert.Empty(t, c)
		})
	}
}
