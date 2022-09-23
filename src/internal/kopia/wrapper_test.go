package kopia

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	stdpath "path"
	"testing"

	"github.com/google/uuid"
	"github.com/kopia/kopia/fs"
	"github.com/kopia/kopia/repo/manifest"
	"github.com/kopia/kopia/snapshot/snapshotfs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/connector/mockconnector"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/path"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/backup/details"
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

			buf, err := ioutil.ReadAll(s.ToReader())
			require.NoError(t, err, "reading collection item: %s", fullPath)

			assert.Equal(t, expected, buf, "comparing collection item: %s", fullPath)
		}
	}

	assert.Equal(t, len(expected), count)
}

func expectDirs(
	t *testing.T,
	entries []fs.Entry,
	dirs []string,
	exactly bool,
) {
	t.Helper()

	if exactly {
		require.Len(t, entries, len(dirs))
	}

	names := make([]string, 0, len(entries))
	for _, e := range entries {
		names = append(names, e.Name())
	}

	assert.Subset(t, names, dirs)
}

//revive:disable:context-as-argument
func getDirEntriesForEntry(
	t *testing.T,
	ctx context.Context,
	entry fs.Entry,
) []fs.Entry {
	//revive:enable:context-as-argument
	d, ok := entry.(fs.Directory)
	require.True(t, ok, "returned entry is not a directory")

	entries, err := fs.GetAllEntries(ctx, d)
	require.NoError(t, err)

	return entries
}

// ---------------
// unit tests
// ---------------
type CorsoProgressUnitSuite struct {
	suite.Suite
	targetFilePath path.Path
	targetFileName string
}

func TestCorsoProgressUnitSuite(t *testing.T) {
	suite.Run(t, new(CorsoProgressUnitSuite))
}

func (suite *CorsoProgressUnitSuite) SetupSuite() {
	p, err := path.Builder{}.Append(
		testInboxDir,
		"testFile",
	).ToDataLayerExchangePathForCategory(
		testTenant,
		testUser,
		path.EmailCategory,
		true,
	)
	require.NoError(suite.T(), err)

	suite.targetFilePath = p
	suite.targetFileName = suite.targetFilePath.ToBuilder().Dir().String()
}

type testInfo struct {
	info       *itemDetails
	err        error
	totalBytes int64
}

var finishedFileTable = []struct {
	name               string
	cachedItems        func(fname string, fpath path.Path) map[string]testInfo
	expectedBytes      int64
	expectedNumEntries int
	err                error
}{
	{
		name: "DetailsExist",
		cachedItems: func(fname string, fpath path.Path) map[string]testInfo {
			return map[string]testInfo{
				fname: {
					info:       &itemDetails{details.ItemInfo{}, fpath},
					err:        nil,
					totalBytes: 100,
				},
			}
		},
		expectedBytes: 100,
		// 1 file and 5 folders.
		expectedNumEntries: 6,
	},
	{
		name: "PendingNoDetails",
		cachedItems: func(fname string, fpath path.Path) map[string]testInfo {
			return map[string]testInfo{
				fname: {
					info: nil,
					err:  nil,
				},
			}
		},
		expectedNumEntries: 0,
	},
	{
		name: "HadError",
		cachedItems: func(fname string, fpath path.Path) map[string]testInfo {
			return map[string]testInfo{
				fname: {
					info: &itemDetails{details.ItemInfo{}, fpath},
					err:  assert.AnError,
				},
			}
		},
		expectedNumEntries: 0,
	},
	{
		name: "NotPending",
		cachedItems: func(fname string, fpath path.Path) map[string]testInfo {
			return nil
		},
		expectedNumEntries: 0,
	},
}

func (suite *CorsoProgressUnitSuite) TestFinishedFile() {
	for _, test := range finishedFileTable {
		suite.T().Run(test.name, func(t *testing.T) {
			bd := &details.Details{}
			cp := corsoProgress{
				UploadProgress: &snapshotfs.NullUploadProgress{},
				deets:          bd,
				pending:        map[string]*itemDetails{},
			}

			ci := test.cachedItems(suite.targetFileName, suite.targetFilePath)

			for k, v := range ci {
				cp.put(k, v.info)
			}

			require.Len(t, cp.pending, len(ci))

			for k, v := range ci {
				cp.FinishedFile(k, v.err)
			}

			assert.Empty(t, cp.pending)
			assert.Len(t, bd.Entries, test.expectedNumEntries)
		})
	}
}

func (suite *CorsoProgressUnitSuite) TestFinishedFileBuildsHierarchy() {
	t := suite.T()
	// Order of folders in hierarchy from root to leaf (excluding the item).
	expectedFolderOrder := suite.targetFilePath.ToBuilder().Dir().Elements()

	// Setup stuff.
	bd := &details.Details{}
	cp := corsoProgress{
		UploadProgress: &snapshotfs.NullUploadProgress{},
		deets:          bd,
		pending:        map[string]*itemDetails{},
	}

	deets := &itemDetails{details.ItemInfo{}, suite.targetFilePath}
	cp.put(suite.targetFileName, deets)
	require.Len(t, cp.pending, 1)

	cp.FinishedFile(suite.targetFileName, nil)

	// Gather information about the current state.
	var (
		curRef     *details.DetailsEntry
		refToEntry = map[string]*details.DetailsEntry{}
	)

	for i := 0; i < len(bd.Entries); i++ {
		e := &bd.Entries[i]
		if e.Folder == nil {
			continue
		}

		refToEntry[e.ShortRef] = e

		if e.Folder.DisplayName == expectedFolderOrder[len(expectedFolderOrder)-1] {
			curRef = e
		}
	}

	// Actual tests start here.
	var rootRef *details.DetailsEntry

	// Traverse the details entries from leaf to root, following the ParentRef
	// fields. At the end rootRef should point to the root of the path.
	for i := len(expectedFolderOrder) - 1; i >= 0; i-- {
		name := expectedFolderOrder[i]

		require.NotNil(t, curRef)
		assert.Equal(t, name, curRef.Folder.DisplayName)

		rootRef = curRef
		curRef = refToEntry[curRef.ParentRef]
	}

	// Hierarchy root's ParentRef = "" and map will return nil.
	assert.Nil(t, curRef)
	require.NotNil(t, rootRef)
	assert.Empty(t, rootRef.ParentRef)
}

func (suite *CorsoProgressUnitSuite) TestFinishedHashingFile() {
	for _, test := range finishedFileTable {
		suite.T().Run(test.name, func(t *testing.T) {
			bd := &details.Details{}
			cp := corsoProgress{
				UploadProgress: &snapshotfs.NullUploadProgress{},
				deets:          bd,
				pending:        map[string]*itemDetails{},
			}

			ci := test.cachedItems(suite.targetFileName, suite.targetFilePath)

			for k, v := range ci {
				cp.FinishedHashingFile(k, v.totalBytes)
			}

			assert.Empty(t, cp.pending)
			assert.Equal(t, test.expectedBytes, cp.totalBytes)
		})
	}
}

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
		w := &Wrapper{}
		w.Close(context.Background())
	})
}

func (suite *KopiaUnitSuite) TestBuildDirectoryTree() {
	tester.LogTimeOfTest(suite.T())

	t := suite.T()
	ctx := context.Background()
	tenant := "a-tenant"
	user1 := testUser
	user1Encoded := encodeAsPath(user1)
	user2 := "user2"
	user2Encoded := encodeAsPath(user2)

	p2, err := path.FromDataLayerPath(
		stdpath.Join(
			tenant,
			service,
			user2,
			category,
			testInboxDir,
		),
		false,
	)
	require.NoError(t, err)

	// Encode user names here so we don't have to decode things later.
	expectedFileCount := map[string]int{
		user1Encoded: 5,
		user2Encoded: 42,
	}

	progress := &corsoProgress{pending: map[string]*itemDetails{}}

	collections := []data.Collection{
		mockconnector.NewMockExchangeCollection(
			suite.testPath,
			expectedFileCount[user1Encoded],
		),
		mockconnector.NewMockExchangeCollection(
			p2,
			expectedFileCount[user2Encoded],
		),
	}

	// Returned directory structure should look like:
	// - a-tenant
	//   - exchange
	//     - user1
	//       - emails
	//         - Inbox
	//           - 5 separate files
	//     - user2
	//       - emails
	//         - Inbox
	//           - 42 separate files
	dirTree, err := inflateDirTree(ctx, collections, progress)
	require.NoError(t, err)
	assert.Equal(t, encodeAsPath(testTenant), dirTree.Name())

	entries, err := fs.GetAllEntries(ctx, dirTree)
	require.NoError(t, err)

	expectDirs(t, entries, encodeElements(service), true)

	entries = getDirEntriesForEntry(t, ctx, entries[0])
	expectDirs(t, entries, encodeElements(user1, user2), true)

	for _, entry := range entries {
		userName := entry.Name()

		entries = getDirEntriesForEntry(t, ctx, entry)
		expectDirs(t, entries, encodeElements(category), true)

		entries = getDirEntriesForEntry(t, ctx, entries[0])
		expectDirs(t, entries, encodeElements(testInboxDir), true)

		entries = getDirEntriesForEntry(t, ctx, entries[0])
		assert.Len(t, entries, expectedFileCount[userName])
	}

	totalFileCount := 0
	for _, c := range expectedFileCount {
		totalFileCount += c
	}

	assert.Len(t, progress.pending, totalFileCount)
}

func (suite *KopiaUnitSuite) TestBuildDirectoryTree_MixedDirectory() {
	ctx := context.Background()
	subdir := "subfolder"

	p2, err := suite.testPath.Append(subdir, false)
	require.NoError(suite.T(), err)

	// Test multiple orders of items because right now order can matter. Both
	// orders result in a directory structure like:
	// - a-tenant
	//   - exchange
	//     - user1
	//       - emails
	//         - Inbox
	//           - subfolder
	//             - 5 separate files
	//           - 42 separate files
	table := []struct {
		name   string
		layout []data.Collection
	}{
		{
			name: "SubdirFirst",
			layout: []data.Collection{
				mockconnector.NewMockExchangeCollection(
					p2,
					5,
				),
				mockconnector.NewMockExchangeCollection(
					suite.testPath,
					42,
				),
			},
		},
		{
			name: "SubdirLast",
			layout: []data.Collection{
				mockconnector.NewMockExchangeCollection(
					suite.testPath,
					42,
				),
				mockconnector.NewMockExchangeCollection(
					p2,
					5,
				),
			},
		},
	}

	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			progress := &corsoProgress{pending: map[string]*itemDetails{}}

			dirTree, err := inflateDirTree(ctx, test.layout, progress)
			require.NoError(t, err)
			assert.Equal(t, encodeAsPath(testTenant), dirTree.Name())

			entries, err := fs.GetAllEntries(ctx, dirTree)
			require.NoError(t, err)

			expectDirs(t, entries, encodeElements(service), true)

			entries = getDirEntriesForEntry(t, ctx, entries[0])
			expectDirs(t, entries, encodeElements(testUser), true)

			entries = getDirEntriesForEntry(t, ctx, entries[0])
			expectDirs(t, entries, encodeElements(category), true)

			entries = getDirEntriesForEntry(t, ctx, entries[0])
			expectDirs(t, entries, encodeElements(testInboxDir), true)

			entries = getDirEntriesForEntry(t, ctx, entries[0])
			// 42 files and 1 subdirectory.
			assert.Len(t, entries, 43)

			// One of these entries should be a subdirectory with items in it.
			subDirs := []fs.Directory(nil)
			for _, e := range entries {
				d, ok := e.(fs.Directory)
				if !ok {
					continue
				}

				subDirs = append(subDirs, d)
				assert.Equal(t, encodeAsPath(subdir), d.Name())
			}

			require.Len(t, subDirs, 1)

			entries = getDirEntriesForEntry(t, ctx, entries[0])
			assert.Len(t, entries, 5)
		})
	}
}

func (suite *KopiaUnitSuite) TestBuildDirectoryTree_Fails() {
	p2, err := path.Builder{}.Append(testInboxDir).ToDataLayerExchangePathForCategory(
		"tenant2",
		"user2",
		path.EmailCategory,
		false,
	)
	require.NoError(suite.T(), err)

	table := []struct {
		name   string
		layout []data.Collection
	}{
		{
			"MultipleRoots",
			// Directory structure would look like:
			// - tenant1
			//   - exchange
			//     - user1
			//       - emails
			//         - Inbox
			//           - 5 separate files
			// - tenant2
			//   - exchange
			//     - user2
			//       - emails
			//         - Inbox
			//           - 42 separate files
			[]data.Collection{
				mockconnector.NewMockExchangeCollection(
					suite.testPath,
					5,
				),
				mockconnector.NewMockExchangeCollection(
					p2,
					42,
				),
			},
		},
		{
			"NoCollectionPath",
			[]data.Collection{
				mockconnector.NewMockExchangeCollection(
					nil,
					5,
				),
			},
		},
	}

	for _, test := range table {
		ctx := context.Background()

		suite.T().Run(test.name, func(t *testing.T) {
			_, err := inflateDirTree(ctx, test.layout, nil)
			assert.Error(t, err)
		})
	}
}

// ---------------
// integration tests that use kopia
// ---------------
type KopiaIntegrationSuite struct {
	suite.Suite
	w   *Wrapper
	ctx context.Context

	testPath1 path.Path
	testPath2 path.Path
}

func TestKopiaIntegrationSuite(t *testing.T) {
	if err := tester.RunOnAny(
		tester.CorsoCITests,
		tester.CorsoKopiaWrapperTests,
	); err != nil {
		t.Skip()
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
	suite.ctx = context.Background()

	c, err := openKopiaRepo(t, suite.ctx)
	require.NoError(t, err)

	suite.w = &Wrapper{c}
}

func (suite *KopiaIntegrationSuite) TearDownTest() {
	assert.NoError(suite.T(), suite.w.Close(suite.ctx))
}

func (suite *KopiaIntegrationSuite) TestBackupCollections() {
	t := suite.T()

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

	stats, rp, err := suite.w.BackupCollections(suite.ctx, collections)
	assert.NoError(t, err)
	assert.Equal(t, stats.TotalFileCount, 47)
	assert.Equal(t, stats.TotalDirectoryCount, 6)
	assert.Equal(t, stats.IgnoredErrorCount, 0)
	assert.Equal(t, stats.ErrorCount, 0)
	assert.False(t, stats.Incomplete)
	// 47 file and 6 folder entries.
	assert.Len(t, rp.Entries, 47+6)
}

func (suite *KopiaIntegrationSuite) TestRestoreAfterCompressionChange() {
	t := suite.T()
	ctx := context.Background()

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

	stats, _, err := w.BackupCollections(ctx, []data.Collection{dc1, dc2})
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
		})

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

	stats, rp, err := suite.w.BackupCollections(suite.ctx, collections)
	require.NoError(t, err)

	assert.Equal(t, 0, stats.ErrorCount)
	assert.Equal(t, 5, stats.TotalFileCount)
	assert.Equal(t, 6, stats.TotalDirectoryCount)
	assert.Equal(t, 1, stats.IgnoredErrorCount)
	assert.False(t, stats.Incomplete)
	// 5 file and 6 folder entries.
	assert.Len(t, rp.Entries, 5+6)
}

type backedupFile struct {
	parentPath path.Path
	itemPath   path.Path
	data       []byte
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
		t.Skip()
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
	suite.ctx = context.Background()
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

	stats, rp, err := suite.w.BackupCollections(suite.ctx, collections)
	require.NoError(t, err)
	require.Equal(t, stats.ErrorCount, 0)
	require.Equal(t, stats.TotalFileCount, expectedFiles)
	require.Equal(t, stats.TotalDirectoryCount, expectedDirs)
	require.Equal(t, stats.IgnoredErrorCount, 0)
	require.False(t, stats.Incomplete)
	// 6 file and 6 folder entries.
	assert.Len(t, rp.Entries, expectedFiles+expectedDirs)

	suite.snapshotID = manifest.ID(stats.SnapshotID)
}

func (suite *KopiaSimpleRepoIntegrationSuite) TearDownTest() {
	assert.NoError(suite.T(), suite.w.Close(suite.ctx))
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

			result, err := suite.w.RestoreMultipleItems(
				suite.ctx,
				string(suite.snapshotID),
				test.inputPaths,
			)
			test.expectedErr(t, err)

			assert.Len(t, result, test.expectedCollections)
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
			)
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
	c, err := suite.w.RestoreMultipleItems(
		suite.ctx,
		string(suite.snapshotID),
		[]path.Path{itemPath},
	)
	assert.Error(t, err, "snapshot should be deleted")
	assert.Empty(t, c)
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
