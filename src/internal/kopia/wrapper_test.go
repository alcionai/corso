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
	"golang.org/x/exp/maps"

	"github.com/alcionai/corso/src/internal/connector/mockconnector"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/aw"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
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

//revive:disable:context-as-argument
func testForFiles(
	t *testing.T,
	ctx context.Context,
	expected map[string][]byte,
	collections []data.RestoreCollection,
) {
	//revive:enable:context-as-argument
	t.Helper()

	count := 0

	for _, c := range collections {
		for s := range c.Items(ctx, fault.New(true)) {
			count++

			fullPath, err := c.FullPath().Append(s.UUID(), true)
			aw.MustNoErr(t, err)

			expected, ok := expected[fullPath.String()]
			require.True(t, ok, "unexpected file with path %q", fullPath)

			buf, err := io.ReadAll(s.ToReader())
			aw.MustNoErr(t, err, "reading collection item: %s", fullPath)

			assert.Equal(t, expected, buf, "comparing collection item: %s", fullPath)

			require.Implements(t, (*data.StreamSize)(nil), s)
			ss := s.(data.StreamSize)
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
	aw.MustNoErr(t, err)
	assert.Equal(t, expectedTags, man.Tags)
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
			testInboxDir,
		),
		false,
	)
	aw.MustNoErr(suite.T(), err)

	suite.testPath = tmp
}

func TestKopiaUnitSuite(t *testing.T) {
	suite.Run(t, &KopiaUnitSuite{Suite: tester.NewUnitSuite(t)})
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
			[][]string{tester.AWSStorageCredEnvs},
			tester.CorsoKopiaWrapperTests,
		),
	})
}

func (suite *KopiaIntegrationSuite) SetupSuite() {
	tmp, err := path.Builder{}.Append(testInboxDir).ToDataLayerExchangePathForCategory(
		testTenant,
		testUser,
		path.EmailCategory,
		false)
	aw.MustNoErr(suite.T(), err)

	suite.storePath1 = tmp
	suite.locPath1 = tmp

	tmp, err = path.Builder{}.Append(testArchiveDir).ToDataLayerExchangePathForCategory(
		testTenant,
		testUser,
		path.EmailCategory,
		false)
	aw.MustNoErr(suite.T(), err)

	suite.storePath2 = tmp
	suite.locPath2 = tmp
}

func (suite *KopiaIntegrationSuite) SetupTest() {
	t := suite.T()
	suite.ctx, suite.flush = tester.NewContext()

	c, err := openKopiaRepo(t, suite.ctx)
	aw.MustNoErr(t, err)

	suite.w = &Wrapper{c}
}

func (suite *KopiaIntegrationSuite) TearDownTest() {
	defer suite.flush()
	aw.NoErr(suite.T(), suite.w.Close(suite.ctx))
}

func (suite *KopiaIntegrationSuite) TestBackupCollections() {
	collections := []data.BackupCollection{
		mockconnector.NewMockExchangeCollection(
			suite.storePath1,
			suite.locPath1,
			5),
		mockconnector.NewMockExchangeCollection(
			suite.storePath2,
			suite.locPath2,
			42),
	}

	// tags that are supplied by the caller. This includes basic tags to support
	// lookups and extra tags the caller may want to apply.
	tags := map[string]string{
		"fnords":    "smarf",
		"brunhilda": "",
	}

	reasons := []Reason{
		{
			ResourceOwner: suite.storePath1.ResourceOwner(),
			Service:       suite.storePath1.Service(),
			Category:      suite.storePath1.Category(),
		},
		{
			ResourceOwner: suite.storePath2.ResourceOwner(),
			Service:       suite.storePath2.Service(),
			Category:      suite.storePath2.Category(),
		},
	}

	for _, r := range reasons {
		for _, k := range r.TagKeys() {
			tags[k] = ""
		}
	}

	expectedTags := map[string]string{}

	maps.Copy(expectedTags, normalizeTagKVs(tags))

	table := []struct {
		name                  string
		expectedUploadedFiles int
		expectedCachedFiles   int
		// Whether entries in the resulting details should be marked as updated.
		deetsUpdated bool
	}{
		{
			name:                  "Uncached",
			expectedUploadedFiles: 47,
			expectedCachedFiles:   0,
			deetsUpdated:          true,
		},
		{
			name:                  "Cached",
			expectedUploadedFiles: 0,
			expectedCachedFiles:   47,
			deetsUpdated:          false,
		},
	}

	prevSnaps := []IncrementalBase{}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			stats, deets, _, err := suite.w.BackupCollections(
				suite.ctx,
				prevSnaps,
				collections,
				nil,
				tags,
				true,
				fault.New(true))
			aw.NoErr(t, err)

			assert.Equal(t, test.expectedUploadedFiles, stats.TotalFileCount, "total files")
			assert.Equal(t, test.expectedUploadedFiles, stats.UncachedFileCount, "uncached files")
			assert.Equal(t, test.expectedCachedFiles, stats.CachedFileCount, "cached files")
			assert.Equal(t, 6, stats.TotalDirectoryCount)
			assert.Equal(t, 0, stats.IgnoredErrorCount)
			assert.Equal(t, 0, stats.ErrorCount)
			assert.False(t, stats.Incomplete)

			// 47 file and 6 folder entries.
			details := deets.Details().Entries
			assert.Len(
				t,
				details,
				test.expectedUploadedFiles+test.expectedCachedFiles+6,
			)

			for _, entry := range details {
				assert.Equal(t, test.deetsUpdated, entry.Updated)
			}

			checkSnapshotTags(
				t,
				suite.ctx,
				suite.w.c,
				expectedTags,
				stats.SnapshotID,
			)

			snap, err := snapshot.LoadSnapshot(
				suite.ctx,
				suite.w.c,
				manifest.ID(stats.SnapshotID),
			)
			aw.MustNoErr(t, err)

			prevSnaps = append(prevSnaps, IncrementalBase{
				Manifest: snap,
				SubtreePaths: []*path.Builder{
					suite.storePath1.ToBuilder().Dir(),
				},
			})
		})
	}
}

func (suite *KopiaIntegrationSuite) TestRestoreAfterCompressionChange() {
	t := suite.T()
	ctx, flush := tester.NewContext()

	defer flush()

	k, err := openKopiaRepo(t, ctx)
	aw.MustNoErr(t, err)

	aw.MustNoErr(t, k.Compression(ctx, "s2-default"))

	w := &Wrapper{k}

	tags := map[string]string{}
	reason := Reason{
		ResourceOwner: testUser,
		Service:       path.ExchangeService,
		Category:      path.EmailCategory,
	}

	for _, k := range reason.TagKeys() {
		tags[k] = ""
	}

	dc1 := mockconnector.NewMockExchangeCollection(suite.storePath1, suite.locPath1, 1)
	dc2 := mockconnector.NewMockExchangeCollection(suite.storePath2, suite.locPath2, 1)

	fp1, err := suite.storePath1.Append(dc1.Names[0], true)
	aw.MustNoErr(t, err)

	fp2, err := suite.storePath2.Append(dc2.Names[0], true)
	aw.MustNoErr(t, err)

	stats, _, _, err := w.BackupCollections(
		ctx,
		nil,
		[]data.BackupCollection{dc1, dc2},
		nil,
		tags,
		true,
		fault.New(true))
	aw.MustNoErr(t, err)

	aw.MustNoErr(t, k.Compression(ctx, "gzip"))

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
		nil,
		fault.New(true))
	aw.MustNoErr(t, err)
	assert.Equal(t, 2, len(result))

	testForFiles(t, ctx, expected, result)
}

type mockBackupCollection struct {
	path    path.Path
	streams []data.Stream
}

func (c *mockBackupCollection) Items(context.Context, *fault.Bus) <-chan data.Stream {
	res := make(chan data.Stream)

	go func() {
		defer close(res)

		for _, s := range c.streams {
			res <- s
		}
	}()

	return res
}

func (c mockBackupCollection) FullPath() path.Path {
	return c.path
}

func (c mockBackupCollection) PreviousPath() path.Path {
	return nil
}

func (c mockBackupCollection) State() data.CollectionState {
	return data.NewState
}

func (c mockBackupCollection) DoNotMergeItems() bool {
	return false
}

func (suite *KopiaIntegrationSuite) TestBackupCollections_ReaderError() {
	t := suite.T()

	tags := map[string]string{}
	reason := Reason{
		ResourceOwner: testUser,
		Service:       path.ExchangeService,
		Category:      path.EmailCategory,
	}

	for _, k := range reason.TagKeys() {
		tags[k] = ""
	}

	collections := []data.BackupCollection{
		&mockBackupCollection{
			path: suite.storePath1,
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
		&mockBackupCollection{
			path: suite.storePath2,
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

	stats, deets, _, err := suite.w.BackupCollections(
		suite.ctx,
		nil,
		collections,
		nil,
		tags,
		true,
		fault.New(true))
	aw.MustErr(t, err)

	assert.Equal(t, 0, stats.ErrorCount)
	assert.Equal(t, 5, stats.TotalFileCount)
	assert.Equal(t, 6, stats.TotalDirectoryCount)
	assert.Equal(t, 1, stats.IgnoredErrorCount)
	assert.False(t, stats.Incomplete)
	// 5 file and 6 folder entries.
	assert.Len(t, deets.Details().Entries, 5+6)

	failedPath, err := suite.storePath2.Append(testFileName4, true)
	aw.MustNoErr(t, err)

	ic := i64counter{}

	_, err = suite.w.RestoreMultipleItems(
		suite.ctx,
		string(stats.SnapshotID),
		[]path.Path{failedPath},
		&ic,
		fault.New(true))
	// Files that had an error shouldn't make a dir entry in kopia. If they do we
	// may run into kopia-assisted incrementals issues because only mod time and
	// not file size is checked for StreamingFiles.
	aw.ErrIs(t, err, data.ErrNotFound, "errored file is restorable")
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

			ctx, flush := tester.NewContext()
			defer flush()

			s, d, _, err := suite.w.BackupCollections(
				ctx,
				nil,
				test.collections,
				nil,
				nil,
				true,
				fault.New(true))
			aw.MustNoErr(t, err)

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
			[][]string{tester.AWSStorageCredEnvs},
			tester.CorsoKopiaWrapperTests,
		),
	})
}

func (suite *KopiaSimpleRepoIntegrationSuite) SetupSuite() {
	tmp, err := path.Builder{}.Append(testInboxDir).ToDataLayerExchangePathForCategory(
		testTenant,
		testUser,
		path.EmailCategory,
		false,
	)
	aw.MustNoErr(suite.T(), err)

	suite.testPath1 = tmp

	tmp, err = path.Builder{}.Append(testArchiveDir).ToDataLayerExchangePathForCategory(
		testTenant,
		testUser,
		path.EmailCategory,
		false,
	)
	aw.MustNoErr(suite.T(), err)

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
		aw.MustNoErr(suite.T(), err)

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
	aw.MustNoErr(t, err)

	suite.w = &Wrapper{c}

	collections := []data.BackupCollection{}

	for _, parent := range []path.Path{suite.testPath1, suite.testPath2} {
		collection := &mockBackupCollection{path: parent}

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

	tags := map[string]string{}
	reason := Reason{
		ResourceOwner: testUser,
		Service:       path.ExchangeService,
		Category:      path.EmailCategory,
	}

	for _, k := range reason.TagKeys() {
		tags[k] = ""
	}

	stats, deets, _, err := suite.w.BackupCollections(
		suite.ctx,
		nil,
		collections,
		nil,
		tags,
		false,
		fault.New(true))
	aw.MustNoErr(t, err)
	require.Equal(t, stats.ErrorCount, 0)
	require.Equal(t, stats.TotalFileCount, expectedFiles)
	require.Equal(t, stats.TotalDirectoryCount, expectedDirs)
	require.Equal(t, stats.IgnoredErrorCount, 0)
	require.False(t, stats.Incomplete)
	// 6 file and 6 folder entries.
	assert.Len(t, deets.Details().Entries, expectedFiles+expectedDirs)

	suite.snapshotID = manifest.ID(stats.SnapshotID)
}

func (suite *KopiaSimpleRepoIntegrationSuite) TearDownTest() {
	aw.NoErr(suite.T(), suite.w.Close(suite.ctx))
	logger.Flush(suite.ctx)
}

type i64counter struct {
	i int64
}

func (c *i64counter) Count(i int64) {
	c.i += i
}

func (suite *KopiaSimpleRepoIntegrationSuite) TestBackupExcludeItem() {
	reason := Reason{
		ResourceOwner: testUser,
		Service:       path.ExchangeService,
		Category:      path.EmailCategory,
	}

	subtreePathTmp, err := path.Builder{}.Append("tmp").ToDataLayerExchangePathForCategory(
		testTenant,
		testUser,
		path.EmailCategory,
		false,
	)
	aw.MustNoErr(suite.T(), err)

	subtreePath := subtreePathTmp.ToBuilder().Dir()

	manifests, err := suite.w.FetchPrevSnapshotManifests(
		suite.ctx,
		[]Reason{reason},
		nil,
	)
	aw.MustNoErr(suite.T(), err)
	require.Len(suite.T(), manifests, 1)
	require.Equal(suite.T(), suite.snapshotID, manifests[0].ID)

	tags := map[string]string{}

	for _, k := range reason.TagKeys() {
		tags[k] = ""
	}

	table := []struct {
		name                  string
		excludeItem           bool
		expectedCachedItems   int
		expectedUncachedItems int
		cols                  func() []data.BackupCollection
		backupIDCheck         require.ValueAssertionFunc
		restoreCheck          assert.ErrorAssertionFunc
	}{
		{
			name:                  "ExcludeItem",
			excludeItem:           true,
			expectedCachedItems:   len(suite.filesByPath) - 1,
			expectedUncachedItems: 0,
			cols: func() []data.BackupCollection {
				return nil
			},
			backupIDCheck: require.NotEmpty,
			restoreCheck:  aw.Err,
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
				c := mockconnector.NewMockExchangeCollection(
					suite.testPath1,
					suite.testPath1,
					1)
				c.ColState = data.NotMovedState

				return []data.BackupCollection{c}
			},
			backupIDCheck: require.NotEmpty,
			restoreCheck:  aw.NoErr,
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			var excluded map[string]struct{}
			if test.excludeItem {
				excluded = map[string]struct{}{
					suite.files[suite.testPath1.String()][0].itemPath.Item(): {},
				}
			}

			stats, _, _, err := suite.w.BackupCollections(
				suite.ctx,
				[]IncrementalBase{
					{
						Manifest: manifests[0].Manifest,
						SubtreePaths: []*path.Builder{
							subtreePath,
						},
					},
				},
				test.cols(),
				excluded,
				tags,
				true,
				fault.New(true))
			aw.MustNoErr(t, err)
			assert.Equal(t, test.expectedCachedItems, stats.CachedFileCount)
			assert.Equal(t, test.expectedUncachedItems, stats.UncachedFileCount)

			test.backupIDCheck(t, stats.SnapshotID)

			if len(stats.SnapshotID) == 0 {
				return
			}

			ic := i64counter{}

			_, err = suite.w.RestoreMultipleItems(
				suite.ctx,
				string(stats.SnapshotID),
				[]path.Path{
					suite.files[suite.testPath1.String()][0].itemPath,
				},
				&ic,
				fault.New(true))
			test.restoreCheck(t, err)
		})
	}
}

func (suite *KopiaSimpleRepoIntegrationSuite) TestRestoreMultipleItems() {
	doesntExist, err := path.Builder{}.Append("subdir", "foo").ToDataLayerExchangePathForCategory(
		testTenant,
		testUser,
		path.EmailCategory,
		true,
	)
	aw.MustNoErr(suite.T(), err)

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
			expectedErr:         aw.NoErr,
		},
		{
			name: "MultipleItemsSameCollection",
			inputPaths: []path.Path{
				suite.files[suite.testPath1.String()][0].itemPath,
				suite.files[suite.testPath1.String()][1].itemPath,
			},
			expectedCollections: 1,
			expectedErr:         aw.NoErr,
		},
		{
			name: "MultipleItemsDifferentCollections",
			inputPaths: []path.Path{
				suite.files[suite.testPath1.String()][0].itemPath,
				suite.files[suite.testPath2.String()][0].itemPath,
			},
			expectedCollections: 2,
			expectedErr:         aw.NoErr,
		},
		{
			name: "TargetNotAFile",
			inputPaths: []path.Path{
				suite.files[suite.testPath1.String()][0].itemPath,
				suite.testPath1,
				suite.files[suite.testPath2.String()][0].itemPath,
			},
			expectedCollections: 0,
			expectedErr:         aw.Err,
		},
		{
			name: "NonExistentFile",
			inputPaths: []path.Path{
				suite.files[suite.testPath1.String()][0].itemPath,
				doesntExist,
				suite.files[suite.testPath2.String()][0].itemPath,
			},
			expectedCollections: 0,
			expectedErr:         aw.Err,
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			ctx, flush := tester.NewContext()
			defer flush()

			t := suite.T()

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
				&ic,
				fault.New(true))
			test.expectedErr(t, err)

			if err != nil {
				return
			}

			assert.Len(t, result, test.expectedCollections)
			assert.Less(t, int64(0), ic.i)
			testForFiles(t, ctx, expected, result)
		})
	}
}

func (suite *KopiaSimpleRepoIntegrationSuite) TestRestoreMultipleItems_Errors() {
	itemPath, err := suite.testPath1.Append(testFileName, true)
	aw.MustNoErr(suite.T(), err)

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
		suite.Run(test.name, func() {
			t := suite.T()

			c, err := suite.w.RestoreMultipleItems(
				suite.ctx,
				test.snapshotID,
				test.paths,
				nil,
				fault.New(true))
			aw.Err(t, err)
			assert.Empty(t, c)
		})
	}
}

func (suite *KopiaSimpleRepoIntegrationSuite) TestDeleteSnapshot() {
	t := suite.T()

	aw.NoErr(t, suite.w.DeleteSnapshot(suite.ctx, string(suite.snapshotID)))

	// assert the deletion worked
	itemPath := suite.files[suite.testPath1.String()][0].itemPath
	ic := i64counter{}

	c, err := suite.w.RestoreMultipleItems(
		suite.ctx,
		string(suite.snapshotID),
		[]path.Path{itemPath},
		&ic,
		fault.New(true))
	aw.Err(t, err, "snapshot should be deleted")
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
			expect:     aw.Err,
		},
		{
			name:       "unknown id",
			snapshotID: uuid.NewString(),
			expect:     aw.NoErr,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			test.expect(t, suite.w.DeleteSnapshot(suite.ctx, test.snapshotID))
		})
	}
}
