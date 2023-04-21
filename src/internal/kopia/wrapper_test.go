package kopia

import (
	"bytes"
	"context"
	"io"
	stdpath "path"
	"testing"

	"github.com/alcionai/clues"
	"github.com/google/uuid"
	"github.com/kopia/kopia/repo"
	"github.com/kopia/kopia/repo/manifest"
	"github.com/kopia/kopia/snapshot"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"golang.org/x/exp/maps"

	exchMock "github.com/alcionai/corso/src/internal/connector/exchange/mock"
	"github.com/alcionai/corso/src/internal/connector/onedrive/metadata"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/data/mock"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/backup/details"
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
			require.NoError(t, err, clues.ToCore(err))

			expected, ok := expected[fullPath.String()]
			require.True(t, ok, "unexpected file with path %q", fullPath)

			buf, err := io.ReadAll(s.ToReader())
			require.NoError(t, err, "reading collection item", fullPath, clues.ToCore(err))
			assert.Equal(t, expected, buf, "comparing collection item", fullPath)

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
	require.NoError(t, err, clues.ToCore(err))
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
	require.NoError(suite.T(), err, clues.ToCore(err))

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
		),
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
	suite.ctx, suite.flush = tester.NewContext()

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
	collections := []data.BackupCollection{
		exchMock.NewCollection(
			suite.storePath1,
			suite.locPath1,
			5),
		exchMock.NewCollection(
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

			stats, deets, _, err := suite.w.ConsumeBackupCollections(
				suite.ctx,
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
			assert.Equal(t, 6, stats.TotalDirectoryCount)
			assert.Equal(t, 0, stats.IgnoredErrorCount)
			assert.Equal(t, 0, stats.ErrorCount)
			assert.False(t, stats.Incomplete)

			// 47 file and 2 folder entries.
			details := deets.Details().Entries
			assert.Len(
				t,
				details,
				test.expectedUploadedFiles+test.expectedCachedFiles+2,
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
			require.NoError(t, err, clues.ToCore(err))

			prevSnaps = append(prevSnaps, IncrementalBase{
				Manifest: snap,
				SubtreePaths: []*path.Builder{
					suite.storePath1.ToBuilder().Dir(),
				},
			})
		})
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

	reasons := []Reason{
		{
			ResourceOwner: storePath.ResourceOwner(),
			Service:       storePath.Service(),
			Category:      storePath.Category(),
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
				streams := []data.Stream{}
				fileNames := []string{
					testFileName,
					testFileName + metadata.MetaFileSuffix,
					metadata.DirMetaFileSuffix,
				}

				for _, name := range fileNames {
					info := baseOneDriveItemInfo
					info.ItemName = name

					ms := &mock.Stream{
						ID:       name,
						Reader:   io.NopCloser(&bytes.Buffer{}),
						ItemSize: 0,
						ItemInfo: details.ItemInfo{OneDrive: &info},
					}

					streams = append(streams, ms)
				}

				mc := &mockBackupCollection{
					path:    storePath,
					loc:     locPath,
					streams: streams,
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

				ms := &mock.Stream{
					ID:       testFileName,
					Reader:   io.NopCloser(&bytes.Buffer{}),
					ItemSize: 0,
					ItemInfo: details.ItemInfo{OneDrive: &info},
				}

				mc := &mockBackupCollection{
					path:    storePath,
					loc:     locPath,
					streams: []data.Stream{ms},
					state:   data.NotMovedState,
				}

				return []data.BackupCollection{mc}
			},
		},
	}

	prevSnaps := []IncrementalBase{}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()
			collections := test.cols()

			stats, deets, prevShortRefs, err := suite.w.ConsumeBackupCollections(
				suite.ctx,
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
				test.numDeetsEntries+1,
			)

			for _, entry := range details {
				assert.True(t, entry.Updated)

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
				stats.SnapshotID,
			)

			snap, err := snapshot.LoadSnapshot(
				suite.ctx,
				suite.w.c,
				manifest.ID(stats.SnapshotID))
			require.NoError(t, err, clues.ToCore(err))

			prevSnaps = append(prevSnaps, IncrementalBase{
				Manifest: snap,
				SubtreePaths: []*path.Builder{
					storePath.ToBuilder().Dir(),
				},
			})
		})
	}
}

func (suite *KopiaIntegrationSuite) TestRestoreAfterCompressionChange() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()

	k, err := openKopiaRepo(t, ctx)
	require.NoError(t, err, clues.ToCore(err))

	err = k.Compression(ctx, "s2-default")
	require.NoError(t, err, clues.ToCore(err))

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

	dc1 := exchMock.NewCollection(suite.storePath1, suite.locPath1, 1)
	dc2 := exchMock.NewCollection(suite.storePath2, suite.locPath2, 1)

	fp1, err := suite.storePath1.Append(dc1.Names[0], true)
	require.NoError(t, err, clues.ToCore(err))

	fp2, err := suite.storePath2.Append(dc2.Names[0], true)
	require.NoError(t, err, clues.ToCore(err))

	stats, _, _, err := w.ConsumeBackupCollections(
		ctx,
		nil,
		[]data.BackupCollection{dc1, dc2},
		nil,
		tags,
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
		[]path.Path{
			fp1,
			fp2,
		},
		nil,
		fault.New(true))
	require.NoError(t, err, clues.ToCore(err))
	assert.Equal(t, 2, len(result))

	testForFiles(t, ctx, expected, result)
}

type mockBackupCollection struct {
	path    path.Path
	loc     *path.Builder
	streams []data.Stream
	state   data.CollectionState
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
	return c.path
}

func (c mockBackupCollection) LocationPath() *path.Builder {
	return c.loc
}

func (c mockBackupCollection) State() data.CollectionState {
	return c.state
}

func (c mockBackupCollection) DoNotMergeItems() bool {
	return false
}

func (suite *KopiaIntegrationSuite) TestBackupCollections_ReaderError() {
	t := suite.T()

	loc1 := path.Builder{}.Append(suite.storePath1.Folders()...)
	loc2 := path.Builder{}.Append(suite.storePath2.Folders()...)
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
			loc:  loc1,
			streams: []data.Stream{
				&exchMock.Data{
					ID:     testFileName,
					Reader: io.NopCloser(bytes.NewReader(testFileData)),
				},
				&exchMock.Data{
					ID:     testFileName2,
					Reader: io.NopCloser(bytes.NewReader(testFileData2)),
				},
			},
		},
		&mockBackupCollection{
			path: suite.storePath2,
			loc:  loc2,
			streams: []data.Stream{
				&exchMock.Data{
					ID:     testFileName3,
					Reader: io.NopCloser(bytes.NewReader(testFileData3)),
				},
				&exchMock.Data{
					ID:      testFileName4,
					ReadErr: assert.AnError,
				},
				&exchMock.Data{
					ID:     testFileName5,
					Reader: io.NopCloser(bytes.NewReader(testFileData5)),
				},
				&exchMock.Data{
					ID:     testFileName6,
					Reader: io.NopCloser(bytes.NewReader(testFileData6)),
				},
			},
		},
	}

	stats, deets, _, err := suite.w.ConsumeBackupCollections(
		suite.ctx,
		nil,
		collections,
		nil,
		tags,
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

	failedPath, err := suite.storePath2.Append(testFileName4, true)
	require.NoError(t, err, clues.ToCore(err))

	ic := i64counter{}

	_, err = suite.w.ProduceRestoreCollections(
		suite.ctx,
		string(stats.SnapshotID),
		[]path.Path{failedPath},
		&ic,
		fault.New(true))
	// Files that had an error shouldn't make a dir entry in kopia. If they do we
	// may run into kopia-assisted incrementals issues because only mod time and
	// not file size is checked for StreamingFiles.
	assert.ErrorIs(t, err, data.ErrNotFound, "errored file is restorable", clues.ToCore(err))
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

			s, d, _, err := suite.w.ConsumeBackupCollections(
				ctx,
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
			[][]string{tester.AWSStorageCredEnvs},
		),
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
		pth, err := item.parentPath.Append(item.name, true)
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
	//nolint:forbidigo
	suite.ctx, _ = logger.SeedLevel(context.Background(), logger.Development)

	c, err := openKopiaRepo(t, suite.ctx)
	require.NoError(t, err, clues.ToCore(err))

	suite.w = &Wrapper{c}

	collections := []data.BackupCollection{}

	for _, parent := range []path.Path{suite.testPath1, suite.testPath2} {
		loc := path.Builder{}.Append(parent.Folders()...)
		collection := &mockBackupCollection{path: parent, loc: loc}

		for _, item := range suite.files[parent.String()] {
			collection.streams = append(
				collection.streams,
				&exchMock.Data{
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

	stats, deets, _, err := suite.w.ConsumeBackupCollections(
		suite.ctx,
		nil,
		collections,
		nil,
		tags,
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
	reason := Reason{
		ResourceOwner: testUser,
		Service:       path.ExchangeService,
		Category:      path.EmailCategory,
	}

	subtreePathTmp, err := path.Build(
		testTenant,
		testUser,
		path.ExchangeService,
		path.EmailCategory,
		false,
		"tmp")
	require.NoError(suite.T(), err, clues.ToCore(err))

	subtreePath := subtreePathTmp.ToBuilder().Dir()

	manifests, err := suite.w.FetchPrevSnapshotManifests(
		suite.ctx,
		[]Reason{reason},
		nil,
	)
	require.NoError(suite.T(), err, clues.ToCore(err))
	require.Len(suite.T(), manifests, 1)
	require.Equal(suite.T(), suite.snapshotID, manifests[0].ID)

	tags := map[string]string{}

	for _, k := range reason.TagKeys() {
		tags[k] = ""
	}

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

			var excluded map[string]map[string]struct{}
			if test.excludeItem {
				excluded = map[string]map[string]struct{}{
					// Add a prefix if needed.
					prefix: {
						itemPath.Item(): {},
					},
				}
			}

			stats, _, _, err := suite.w.ConsumeBackupCollections(
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
			require.NoError(t, err, clues.ToCore(err))
			assert.Equal(t, test.expectedCachedItems, stats.CachedFileCount)
			assert.Equal(t, test.expectedUncachedItems, stats.UncachedFileCount)

			test.backupIDCheck(t, stats.SnapshotID)

			if len(stats.SnapshotID) == 0 {
				return
			}

			ic := i64counter{}

			_, err = suite.w.ProduceRestoreCollections(
				suite.ctx,
				string(stats.SnapshotID),
				[]path.Path{
					suite.files[suite.testPath1.String()][0].itemPath,
				},
				&ic,
				fault.New(true))
			test.restoreCheck(t, err, clues.ToCore(err))
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
			expectedCollections: 0,
			expectedErr:         assert.Error,
		},
		{
			name: "NonExistentFile",
			inputPaths: []path.Path{
				suite.files[suite.testPath1.String()][0].itemPath,
				doesntExist,
				suite.files[suite.testPath2.String()][0].itemPath,
			},
			expectedCollections: 0,
			expectedErr:         assert.Error,
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

			result, err := suite.w.ProduceRestoreCollections(
				suite.ctx,
				string(suite.snapshotID),
				test.inputPaths,
				&ic,
				fault.New(true))
			test.expectedErr(t, err, clues.ToCore(err))

			if err != nil {
				return
			}

			assert.Len(t, result, test.expectedCollections)
			assert.Less(t, int64(0), ic.i)
			testForFiles(t, ctx, expected, result)
		})
	}
}

func (suite *KopiaSimpleRepoIntegrationSuite) TestProduceRestoreCollections_Errors() {
	itemPath, err := suite.testPath1.Append(testFileName, true)
	require.NoError(suite.T(), err, clues.ToCore(err))

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

func (suite *KopiaSimpleRepoIntegrationSuite) TestDeleteSnapshot() {
	t := suite.T()

	err := suite.w.DeleteSnapshot(suite.ctx, string(suite.snapshotID))
	assert.NoError(t, err, clues.ToCore(err))

	// assert the deletion worked
	itemPath := suite.files[suite.testPath1.String()][0].itemPath
	ic := i64counter{}

	c, err := suite.w.ProduceRestoreCollections(
		suite.ctx,
		string(suite.snapshotID),
		[]path.Path{itemPath},
		&ic,
		fault.New(true))
	assert.Error(t, err, "snapshot should be deleted", clues.ToCore(err))
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
		suite.Run(test.name, func() {
			t := suite.T()

			err := suite.w.DeleteSnapshot(suite.ctx, test.snapshotID)
			test.expect(t, err, clues.ToCore(err))
		})
	}
}
