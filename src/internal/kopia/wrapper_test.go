package kopia

import (
	"bytes"
	"context"
	"io"
	stdpath "path"
	"testing"

	"github.com/google/uuid"
	"github.com/kopia/kopia/repo"
	"github.com/kopia/kopia/repo/manifest"
	"github.com/kopia/kopia/snapshot"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/connector/mockconnector"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
)

const (
	testTenant     = "a-tenant"
	testUser       = "user1"
	testInboxDir   = "Inbox"
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
	expected map[string][]byte,
	collections []data.Collection,
) {
	t.Helper()

	count := 0

	for _, c := range collections {
		for s := range c.Items() {
			count++

			fullPath, err := c.FullPath().Append(s.UUID(), true)
			require.NoError(t, err)

			expected, ok := expected[fullPath.String()]
			require.True(t, ok, "unexpected file with path %q", fullPath)

			buf, err := io.ReadAll(s.ToReader())
			require.NoError(t, err, "reading collection item: %s", fullPath)

			assert.Equal(t, expected, buf, "comparing collection item: %s", fullPath)

			require.Implements(t, (*data.StreamSize)(nil), s)
			ss := s.(data.StreamSize)
			assert.Equal(t, len(buf), int(ss.Size()))
		}
	}

	assert.Equal(t, len(expected), count)
}

//revive:disable:context-as-argument
func checkSnapshotTags(
	t *testing.T,
	ctx context.Context,
	rep repo.Repository,
	expectedTags map[string]string,
	snapshotID string,
) {
	//revive:enable:context-as-argument
	man, err := snapshot.LoadSnapshot(ctx, rep, manifest.ID(snapshotID))
	require.NoError(t, err)
	assert.Equal(t, expectedTags, man.Tags)
}

// ---------------
// unit tests
// ---------------
type KopiaUnitSuite struct {
	suite.Suite
	testPath path.Path
}

func (suite *KopiaUnitSuite) SetupSuite() {
	tmp, err := path.FromDataLayerPath(
		stdpath.Join(
			testTenant,
			path.ExchangeService.String(),
			testUser,
			path.EmailCategory.String(),
			testInboxDir,
		),
		false,
	)
	require.NoError(suite.T(), err)

	suite.testPath = tmp
}

func TestKopiaUnitSuite(t *testing.T) {
	suite.Run(t, new(KopiaUnitSuite))
}

func (suite *KopiaUnitSuite) TestCloseWithoutInitDoesNotPanic() {
	assert.NotPanics(suite.T(), func() {
		ctx, flush := tester.NewContext()
		defer flush()

		w := &Wrapper{}
		w.Close(ctx)
	})
}

// ---------------
// integration tests that use kopia
// ---------------
type KopiaIntegrationSuite struct {
	suite.Suite
	w     *Wrapper
	ctx   context.Context
	flush func()

	testPath1 path.Path
	testPath2 path.Path
}

func TestKopiaIntegrationSuite(t *testing.T) {
	if err := tester.RunOnAny(
		tester.CorsoCITests,
		tester.CorsoKopiaWrapperTests,
	); err != nil {
		t.Skip(err)
	}

	suite.Run(t, new(KopiaIntegrationSuite))
}

func (suite *KopiaIntegrationSuite) SetupSuite() {
	_, err := tester.GetRequiredEnvVars(tester.AWSStorageCredEnvs...)
	require.NoError(suite.T(), err)

	tmp, err := path.Builder{}.Append(testInboxDir).ToDataLayerExchangePathForCategory(
		testTenant,
		testUser,
		path.EmailCategory,
		false,
	)
	require.NoError(suite.T(), err)

	suite.testPath1 = tmp

	tmp, err = path.Builder{}.Append(testArchiveDir).ToDataLayerExchangePathForCategory(
		testTenant,
		testUser,
		path.EmailCategory,
		false,
	)
	require.NoError(suite.T(), err)

	suite.testPath2 = tmp
}

func (suite *KopiaIntegrationSuite) SetupTest() {
	t := suite.T()
	suite.ctx, suite.flush = tester.NewContext()

	c, err := openKopiaRepo(t, suite.ctx)
	require.NoError(t, err)

	suite.w = &Wrapper{c}
}

func (suite *KopiaIntegrationSuite) TearDownTest() {
	defer suite.flush()
	assert.NoError(suite.T(), suite.w.Close(suite.ctx))
}

func (suite *KopiaIntegrationSuite) TestBackupCollections() {
	collections := []data.Collection{
		mockconnector.NewMockExchangeCollection(
			suite.testPath1,
			5,
		),
		mockconnector.NewMockExchangeCollection(
			suite.testPath2,
			42,
		),
	}

	// tags that are expected to populate as a side effect
	// of the backup process.
	baseTagKeys := []string{
		serviceCatTag(suite.testPath1),
		suite.testPath1.ResourceOwner(),
		serviceCatTag(suite.testPath2),
		suite.testPath2.ResourceOwner(),
	}

	// tags that are supplied by the caller.
	customTags := map[string]string{
		"fnords":    "smarf",
		"brunhilda": "",
	}

	expectedTags := map[string]string{}

	for _, k := range baseTagKeys {
		tk, tv := MakeTagKV(k)
		expectedTags[tk] = tv
	}

	for k, v := range normalizeTagKVs(customTags) {
		expectedTags[k] = v
	}

	table := []struct {
		name                  string
		expectedUploadedFiles int
		expectedCachedFiles   int
	}{
		{
			name:                  "Uncached",
			expectedUploadedFiles: 47,
			expectedCachedFiles:   0,
		},
		{
			name:                  "Cached",
			expectedUploadedFiles: 0,
			expectedCachedFiles:   47,
		},
	}

	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			stats, deets, err := suite.w.BackupCollections(
				suite.ctx,
				nil,
				collections,
				path.ExchangeService,
				customTags,
			)
			assert.NoError(t, err)

			assert.Equal(t, test.expectedUploadedFiles, stats.TotalFileCount, "total files")
			assert.Equal(t, test.expectedUploadedFiles, stats.UncachedFileCount, "uncached files")
			assert.Equal(t, test.expectedCachedFiles, stats.CachedFileCount, "cached files")
			assert.Equal(t, 6, stats.TotalDirectoryCount)
			assert.Equal(t, 0, stats.IgnoredErrorCount)
			assert.Equal(t, 0, stats.ErrorCount)
			assert.False(t, stats.Incomplete)
			// 47 file and 6 folder entries.
			assert.Len(
				t,
				deets.Entries,
				test.expectedUploadedFiles+test.expectedCachedFiles+6,
			)

			checkSnapshotTags(
				t,
				suite.ctx,
				suite.w.c,
				expectedTags,
				stats.SnapshotID,
			)
		})
	}
}

func (suite *KopiaIntegrationSuite) TestRestoreAfterCompressionChange() {
	t := suite.T()
	ctx, flush := tester.NewContext()

	defer flush()

	k, err := openKopiaRepo(t, ctx)
	require.NoError(t, err)

	require.NoError(t, k.Compression(ctx, "s2-default"))

	w := &Wrapper{k}

	dc1 := mockconnector.NewMockExchangeCollection(suite.testPath1, 1)
	dc2 := mockconnector.NewMockExchangeCollection(suite.testPath2, 1)

	fp1, err := suite.testPath1.Append(dc1.Names[0], true)
	require.NoError(t, err)

	fp2, err := suite.testPath2.Append(dc2.Names[0], true)
	require.NoError(t, err)

	stats, _, err := w.BackupCollections(
		ctx,
		nil,
		[]data.Collection{dc1, dc2},
		path.ExchangeService,
		nil,
	)
	require.NoError(t, err)

	require.NoError(t, k.Compression(ctx, "gzip"))

	expected := map[string][]byte{
		fp1.String(): dc1.Data[0],
		fp2.String(): dc2.Data[0],
	}

	result, err := w.RestoreMultipleItems(
		ctx,
		string(stats.SnapshotID),
		[]path.Path{
			fp1,
			fp2,
		},
		nil)

	require.NoError(t, err)
	assert.Equal(t, 2, len(result))

	testForFiles(t, expected, result)
}

func (suite *KopiaIntegrationSuite) TestBackupCollections_ReaderError() {
	t := suite.T()

	collections := []data.Collection{
		&kopiaDataCollection{
			path: suite.testPath1,
			streams: []data.Stream{
				&mockconnector.MockExchangeData{
					ID:     testFileName,
					Reader: io.NopCloser(bytes.NewReader(testFileData)),
				},
				&mockconnector.MockExchangeData{
					ID:     testFileName2,
					Reader: io.NopCloser(bytes.NewReader(testFileData2)),
				},
			},
		},
		&kopiaDataCollection{
			path: suite.testPath2,
			streams: []data.Stream{
				&mockconnector.MockExchangeData{
					ID:     testFileName3,
					Reader: io.NopCloser(bytes.NewReader(testFileData3)),
				},
				&mockconnector.MockExchangeData{
					ID:      testFileName4,
					ReadErr: assert.AnError,
				},
				&mockconnector.MockExchangeData{
					ID:     testFileName5,
					Reader: io.NopCloser(bytes.NewReader(testFileData5)),
				},
				&mockconnector.MockExchangeData{
					ID:     testFileName6,
					Reader: io.NopCloser(bytes.NewReader(testFileData6)),
				},
			},
		},
	}

	stats, deets, err := suite.w.BackupCollections(
		suite.ctx,
		nil,
		collections,
		path.ExchangeService,
		nil,
	)
	require.NoError(t, err)

	assert.Equal(t, 0, stats.ErrorCount)
	assert.Equal(t, 5, stats.TotalFileCount)
	assert.Equal(t, 6, stats.TotalDirectoryCount)
	assert.Equal(t, 1, stats.IgnoredErrorCount)
	assert.False(t, stats.Incomplete)
	// 5 file and 6 folder entries.
	assert.Len(t, deets.Entries, 5+6)
}

type backedupFile struct {
	parentPath path.Path
	itemPath   path.Path
	data       []byte
}

func (suite *KopiaIntegrationSuite) TestBackupCollectionsHandlesNoCollections() {
	table := []struct {
		name        string
		collections []data.Collection
	}{
		{
			name:        "NilCollections",
			collections: nil,
		},
		{
			name:        "EmptyCollections",
			collections: []data.Collection{},
		},
	}

	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			ctx, flush := tester.NewContext()
			defer flush()

			s, d, err := suite.w.BackupCollections(
				ctx,
				nil,
				test.collections,
				path.UnknownService,
				nil,
			)
			require.NoError(t, err)

			assert.Equal(t, BackupStats{}, *s)
			assert.Empty(t, d.Entries)
		})
	}
}

type KopiaSimpleRepoIntegrationSuite struct {
	suite.Suite
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
	if err := tester.RunOnAny(
		tester.CorsoCITests,
		tester.CorsoKopiaWrapperTests,
	); err != nil {
		t.Skip(err)
	}

	suite.Run(t, new(KopiaSimpleRepoIntegrationSuite))
}

func (suite *KopiaSimpleRepoIntegrationSuite) SetupSuite() {
	_, err := tester.GetRequiredEnvVars(tester.AWSStorageCredEnvs...)
	require.NoError(suite.T(), err)

	tmp, err := path.Builder{}.Append(testInboxDir).ToDataLayerExchangePathForCategory(
		testTenant,
		testUser,
		path.EmailCategory,
		false,
	)
	require.NoError(suite.T(), err)

	suite.testPath1 = tmp

	tmp, err = path.Builder{}.Append(testArchiveDir).ToDataLayerExchangePathForCategory(
		testTenant,
		testUser,
		path.EmailCategory,
		false,
	)
	require.NoError(suite.T(), err)

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
		pth, err := item.parentPath.Append(item.name, true)
		require.NoError(suite.T(), err)

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
	//nolint:forbidigo
	suite.ctx, _ = logger.SeedLevel(context.Background(), logger.Development)
	c, err := openKopiaRepo(t, suite.ctx)
	require.NoError(t, err)

	suite.w = &Wrapper{c}

	collections := []data.Collection{}

	for _, parent := range []path.Path{suite.testPath1, suite.testPath2} {
		collection := &kopiaDataCollection{path: parent}

		for _, item := range suite.files[parent.String()] {
			collection.streams = append(
				collection.streams,
				&mockconnector.MockExchangeData{
					ID:     item.itemPath.Item(),
					Reader: io.NopCloser(bytes.NewReader(item.data)),
				},
			)
		}

		collections = append(collections, collection)
	}

	stats, deets, err := suite.w.BackupCollections(
		suite.ctx,
		nil,
		collections,
		path.ExchangeService,
		nil,
	)
	require.NoError(t, err)
	require.Equal(t, stats.ErrorCount, 0)
	require.Equal(t, stats.TotalFileCount, expectedFiles)
	require.Equal(t, stats.TotalDirectoryCount, expectedDirs)
	require.Equal(t, stats.IgnoredErrorCount, 0)
	require.False(t, stats.Incomplete)
	// 6 file and 6 folder entries.
	assert.Len(t, deets.Entries, expectedFiles+expectedDirs)

	suite.snapshotID = manifest.ID(stats.SnapshotID)
}

func (suite *KopiaSimpleRepoIntegrationSuite) TearDownTest() {
	assert.NoError(suite.T(), suite.w.Close(suite.ctx))
	logger.Flush(suite.ctx)
}

type i64counter struct {
	i int64
}

func (c *i64counter) Count(i int64) {
	c.i += i
}

func (suite *KopiaSimpleRepoIntegrationSuite) TestRestoreMultipleItems() {
	doesntExist, err := path.Builder{}.Append("subdir", "foo").ToDataLayerExchangePathForCategory(
		testTenant,
		testUser,
		path.EmailCategory,
		true,
	)
	require.NoError(suite.T(), err)

	// Expected items is generated during the test by looking up paths in the
	// suite's map of files. Files that are not in the suite's map are assumed to
	// generate errors and not be in the output.
	table := []struct {
		name                string
		inputPaths          []path.Path
		expectedCollections int
		expectedErr         assert.ErrorAssertionFunc
	}{
		{
			name: "SingleItem",
			inputPaths: []path.Path{
				suite.files[suite.testPath1.String()][0].itemPath,
			},
			expectedCollections: 1,
			expectedErr:         assert.NoError,
		},
		{
			name: "MultipleItemsSameCollection",
			inputPaths: []path.Path{
				suite.files[suite.testPath1.String()][0].itemPath,
				suite.files[suite.testPath1.String()][1].itemPath,
			},
			expectedCollections: 1,
			expectedErr:         assert.NoError,
		},
		{
			name: "MultipleItemsDifferentCollections",
			inputPaths: []path.Path{
				suite.files[suite.testPath1.String()][0].itemPath,
				suite.files[suite.testPath2.String()][0].itemPath,
			},
			expectedCollections: 2,
			expectedErr:         assert.NoError,
		},
		{
			name: "TargetNotAFile",
			inputPaths: []path.Path{
				suite.files[suite.testPath1.String()][0].itemPath,
				suite.testPath1,
				suite.files[suite.testPath2.String()][0].itemPath,
			},
			expectedCollections: 2,
			expectedErr:         assert.Error,
		},
		{
			name: "NonExistentFile",
			inputPaths: []path.Path{
				suite.files[suite.testPath1.String()][0].itemPath,
				doesntExist,
				suite.files[suite.testPath2.String()][0].itemPath,
			},
			expectedCollections: 2,
			expectedErr:         assert.Error,
		},
	}

	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
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

			result, err := suite.w.RestoreMultipleItems(
				suite.ctx,
				string(suite.snapshotID),
				test.inputPaths,
				&ic)
			test.expectedErr(t, err)

			assert.Len(t, result, test.expectedCollections)
			assert.Less(t, int64(0), ic.i)
			testForFiles(t, expected, result)
		})
	}
}

func (suite *KopiaSimpleRepoIntegrationSuite) TestRestoreMultipleItems_Errors() {
	itemPath, err := suite.testPath1.Append(testFileName, true)
	require.NoError(suite.T(), err)

	table := []struct {
		name       string
		snapshotID string
		paths      []path.Path
	}{
		{
			"NilPaths",
			string(suite.snapshotID),
			nil,
		},
		{
			"EmptyPaths",
			string(suite.snapshotID),
			[]path.Path{},
		},
		{
			"NoSnapshot",
			"foo",
			[]path.Path{itemPath},
		},
	}

	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			c, err := suite.w.RestoreMultipleItems(
				suite.ctx,
				test.snapshotID,
				test.paths,
				nil)
			assert.Error(t, err)
			assert.Empty(t, c)
		})
	}
}

func (suite *KopiaSimpleRepoIntegrationSuite) TestDeleteSnapshot() {
	t := suite.T()

	assert.NoError(t, suite.w.DeleteSnapshot(suite.ctx, string(suite.snapshotID)))

	// assert the deletion worked
	itemPath := suite.files[suite.testPath1.String()][0].itemPath
	ic := i64counter{}

	c, err := suite.w.RestoreMultipleItems(
		suite.ctx,
		string(suite.snapshotID),
		[]path.Path{itemPath},
		&ic)
	assert.Error(t, err, "snapshot should be deleted")
	assert.Empty(t, c)
	assert.Zero(t, ic.i)
}

func (suite *KopiaSimpleRepoIntegrationSuite) TestDeleteSnapshot_BadIDs() {
	table := []struct {
		name       string
		snapshotID string
		expect     assert.ErrorAssertionFunc
	}{
		{
			name:       "no id",
			snapshotID: "",
			expect:     assert.Error,
		},
		{
			name:       "unknown id",
			snapshotID: uuid.NewString(),
			expect:     assert.NoError,
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			test.expect(t, suite.w.DeleteSnapshot(suite.ctx, test.snapshotID))
		})
	}
}
