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
	"github.com/alcionai/corso/src/internal/kopia/mockkopia"
	"github.com/alcionai/corso/src/internal/path"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/backup/details"
)

const (
	testTenant     = "a-tenant"
	testUser       = "user1"
	testInboxDir   = "inbox"
	testArchiveDir = "archive"
	testFileName   = "file1"
	testFileName2  = "file2"
	testFileName3  = "file3"
	testFileName4  = "file4"
	testFileName5  = "file5"
	testFileName6  = "file6"
)

var (
	service  = path.ExchangeService.String()
	category = path.EmailCategory.String()
	testPath = []string{
		testTenant,
		service,
		testUser,
		category,
		testInboxDir,
	}
	testPath2 = []string{
		testTenant,
		service,
		testUser,
		category,
		testArchiveDir,
	}
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
	count := 0

	for _, c := range collections {
		for s := range c.Items() {
			count++

			fullPath := stdpath.Join(append(c.FullPath(), s.UUID())...)

			expected, ok := expected[fullPath]
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
}

func TestCorsoProgressUnitSuite(t *testing.T) {
	suite.Run(t, new(CorsoProgressUnitSuite))
}

func (suite *CorsoProgressUnitSuite) TestFinishedFile() {
	type testInfo struct {
		info *itemDetails
		err  error
	}

	targetFileName := "testFile"
	deets := &itemDetails{details.ItemInfo{}, targetFileName}

	table := []struct {
		name        string
		cachedItems map[string]testInfo
		expectedLen int
		err         error
	}{
		{
			name: "DetailsExist",
			cachedItems: map[string]testInfo{
				targetFileName: {
					info: deets,
					err:  nil,
				},
			},
			expectedLen: 1,
		},
		{
			name: "PendingNoDetails",
			cachedItems: map[string]testInfo{
				targetFileName: {
					info: nil,
					err:  nil,
				},
			},
			expectedLen: 0,
		},
		{
			name: "HadError",
			cachedItems: map[string]testInfo{
				targetFileName: {
					info: deets,
					err:  assert.AnError,
				},
			},
			expectedLen: 0,
		},
		{
			name:        "NotPending",
			expectedLen: 0,
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			bd := &details.Details{}
			cp := corsoProgress{
				UploadProgress: &snapshotfs.NullUploadProgress{},
				deets:          bd,
				pending:        map[string]*itemDetails{},
			}

			for k, v := range test.cachedItems {
				cp.put(k, v.info)
			}

			require.Len(t, cp.pending, len(test.cachedItems))

			for k, v := range test.cachedItems {
				cp.FinishedFile(k, v.err)
			}

			assert.Empty(t, cp.pending)
			assert.Len(t, bd.Entries, test.expectedLen)
		})
	}
}

type KopiaUnitSuite struct {
	suite.Suite
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
	user2 := "user2"

	expectedFileCount := map[string]int{
		user1: 5,
		user2: 42,
	}

	progress := &corsoProgress{pending: map[string]*itemDetails{}}

	collections := []data.Collection{
		mockconnector.NewMockExchangeCollection(
			[]string{tenant, service, user1, category, testInboxDir},
			expectedFileCount[user1],
		),
		mockconnector.NewMockExchangeCollection(
			[]string{tenant, service, user2, category, testInboxDir},
			expectedFileCount[user2],
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
	assert.Equal(t, testTenant, dirTree.Name())

	entries, err := fs.GetAllEntries(ctx, dirTree)
	require.NoError(t, err)

	expectDirs(t, entries, []string{service}, true)

	entries = getDirEntriesForEntry(t, ctx, entries[0])
	expectDirs(t, entries, []string{user1, user2}, true)

	for _, entry := range entries {
		userName := entry.Name()

		entries = getDirEntriesForEntry(t, ctx, entry)
		expectDirs(t, entries, []string{category}, true)

		entries = getDirEntriesForEntry(t, ctx, entries[0])
		expectDirs(t, entries, []string{testInboxDir}, true)

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
					[]string{testTenant, service, testUser, category, testInboxDir, subdir},
					5,
				),
				mockconnector.NewMockExchangeCollection(
					[]string{testTenant, service, testUser, category, testInboxDir},
					42,
				),
			},
		},
		{
			name: "SubdirLast",
			layout: []data.Collection{
				mockconnector.NewMockExchangeCollection(
					[]string{testTenant, service, testUser, category, testInboxDir},
					42,
				),
				mockconnector.NewMockExchangeCollection(
					[]string{testTenant, service, testUser, category, testInboxDir, subdir},
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
			assert.Equal(t, testTenant, dirTree.Name())

			entries, err := fs.GetAllEntries(ctx, dirTree)
			require.NoError(t, err)

			expectDirs(t, entries, []string{service}, true)

			entries = getDirEntriesForEntry(t, ctx, entries[0])
			expectDirs(t, entries, []string{testUser}, true)

			entries = getDirEntriesForEntry(t, ctx, entries[0])
			expectDirs(t, entries, []string{category}, true)

			entries = getDirEntriesForEntry(t, ctx, entries[0])
			expectDirs(t, entries, []string{testInboxDir}, true)

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
				assert.Equal(t, "subfolder", d.Name())
			}

			require.Len(t, subDirs, 1)

			entries = getDirEntriesForEntry(t, ctx, entries[0])
			assert.Len(t, entries, 5)
		})
	}
}

func (suite *KopiaUnitSuite) TestBuildDirectoryTree_Fails() {
	table := []struct {
		name   string
		layout []data.Collection
	}{
		{
			"MultipleRoots",
			// Directory structure would look like:
			// - user1
			//   - emails
			//     - 5 separate files
			// - user2
			//   - emails
			//     - 42 separate files
			[]data.Collection{
				mockconnector.NewMockExchangeCollection(
					[]string{"user1", "emails"},
					5,
				),
				mockconnector.NewMockExchangeCollection(
					[]string{"user2", "emails"},
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

func (suite *KopiaUnitSuite) TestRestoreItem() {
	ctx := context.Background()

	file := &mockkopia.MockFile{
		Entry: &mockkopia.MockEntry{
			EntryName: testFileName2,
			EntryMode: mockkopia.DefaultPermissions,
		},
		OpenErr: assert.AnError,
	}

	_, err := restoreSingleItem(ctx, file, nil)
	assert.Error(suite.T(), err)
}

// ---------------
// integration tests that use kopia
// ---------------
type KopiaIntegrationSuite struct {
	suite.Suite
	w   *Wrapper
	ctx context.Context
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
			[]string{"a-tenant", service, "user1", category, testInboxDir},
			5,
		),
		mockconnector.NewMockExchangeCollection(
			[]string{"a-tenant", service, "user2", category, testInboxDir},
			42,
		),
	}

	stats, rp, err := suite.w.BackupCollections(suite.ctx, collections)
	assert.NoError(t, err)
	assert.Equal(t, stats.TotalFileCount, 47)
	assert.Equal(t, stats.TotalDirectoryCount, 8)
	assert.Equal(t, stats.IgnoredErrorCount, 0)
	assert.Equal(t, stats.ErrorCount, 0)
	assert.False(t, stats.Incomplete)
	assert.Len(t, rp.Entries, 47)
}

func (suite *KopiaIntegrationSuite) TestRestoreAfterCompressionChange() {
	t := suite.T()
	ctx := context.Background()

	k, err := openKopiaRepo(t, ctx)
	require.NoError(t, err)

	require.NoError(t, k.Compression(ctx, "s2-default"))

	w := &Wrapper{k}

	tid := uuid.NewString()
	p1 := []string{
		tid,
		service,
		"uid",
		category,
		"fid",
	}
	p2 := []string{
		tid,
		service,
		"uid2",
		category,
		"fid",
	}
	dc1 := mockconnector.NewMockExchangeCollection(p1, 1)
	dc2 := mockconnector.NewMockExchangeCollection(p2, 1)

	fp1 := append(p1, dc1.Names[0])
	fp2 := append(p2, dc2.Names[0])

	stats, _, err := w.BackupCollections(ctx, []data.Collection{dc1, dc2})
	require.NoError(t, err)

	require.NoError(t, k.Compression(ctx, "gzip"))

	expected := map[string][]byte{
		stdpath.Join(fp1...): dc1.Data[0],
		stdpath.Join(fp2...): dc2.Data[0],
	}

	result, err := w.RestoreMultipleItems(
		ctx,
		string(stats.SnapshotID),
		[][]string{fp1, fp2})

	require.NoError(t, err)
	assert.Equal(t, 2, len(result))

	testForFiles(t, expected, result)
}

func (suite *KopiaIntegrationSuite) TestBackupCollections_ReaderError() {
	t := suite.T()
	tmpBuilder := path.Builder{}.Append(testInboxDir)

	p1, err := tmpBuilder.ToDataLayerExchangePathForCategory(
		testTenant,
		testUser,
		path.EmailCategory,
		false,
	)
	require.NoError(t, err)

	tmpBuilder = path.Builder{}.Append(testArchiveDir)

	p2, err := tmpBuilder.ToDataLayerExchangePathForCategory(
		testTenant,
		testUser,
		path.EmailCategory,
		false,
	)
	require.NoError(t, err)

	collections := []data.Collection{
		&kopiaDataCollection{
			path: p1,
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
			path: p2,
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
	assert.Len(t, rp.Entries, 5)
}

type KopiaSimpleRepoIntegrationSuite struct {
	suite.Suite
	w                    *Wrapper
	ctx                  context.Context
	snapshotID           manifest.ID
	inboxExpectedFiles   map[string][]byte
	archiveExpectedFiles map[string][]byte
	allExpectedFiles     map[string][]byte
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
}

func (suite *KopiaSimpleRepoIntegrationSuite) SetupTest() {
	t := suite.T()
	suite.ctx = context.Background()
	c, err := openKopiaRepo(t, suite.ctx)
	require.NoError(t, err)

	suite.w = &Wrapper{c}
	tmpBuilder := path.Builder{}.Append(testInboxDir)

	p1, err := tmpBuilder.ToDataLayerExchangePathForCategory(
		testTenant,
		testUser,
		path.EmailCategory,
		false,
	)
	require.NoError(t, err)

	tmpBuilder = path.Builder{}.Append(testArchiveDir)

	p2, err := tmpBuilder.ToDataLayerExchangePathForCategory(
		testTenant,
		testUser,
		path.EmailCategory,
		false,
	)
	require.NoError(t, err)

	collections := []data.Collection{
		&kopiaDataCollection{
			path: p1,
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
			path: p2,
			streams: []data.Stream{
				&mockconnector.MockExchangeData{
					ID:     testFileName3,
					Reader: io.NopCloser(bytes.NewReader(testFileData3)),
				},
				&mockconnector.MockExchangeData{
					ID:     testFileName4,
					Reader: io.NopCloser(bytes.NewReader(testFileData4)),
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
	require.Equal(t, stats.ErrorCount, 0)
	require.Equal(t, stats.TotalFileCount, 6)
	require.Equal(t, stats.TotalDirectoryCount, 6)
	require.Equal(t, stats.IgnoredErrorCount, 0)
	require.False(t, stats.Incomplete)
	assert.Len(t, rp.Entries, 6)

	suite.snapshotID = manifest.ID(stats.SnapshotID)

	// path.Join doesn't like (testPath..., testFileName).
	suite.inboxExpectedFiles = map[string][]byte{
		stdpath.Join(append(testPath, testFileName)...):  testFileData,
		stdpath.Join(append(testPath, testFileName2)...): testFileData2,
	}
	suite.archiveExpectedFiles = map[string][]byte{
		stdpath.Join(append(testPath2, testFileName3)...): testFileData3,
		stdpath.Join(append(testPath2, testFileName4)...): testFileData4,
		stdpath.Join(append(testPath2, testFileName5)...): testFileData5,
		stdpath.Join(append(testPath2, testFileName6)...): testFileData6,
	}

	suite.allExpectedFiles = map[string][]byte{}
	for k, v := range suite.inboxExpectedFiles {
		suite.allExpectedFiles[k] = v
	}

	for k, v := range suite.archiveExpectedFiles {
		suite.allExpectedFiles[k] = v
	}
}

func (suite *KopiaSimpleRepoIntegrationSuite) TearDownTest() {
	assert.NoError(suite.T(), suite.w.Close(suite.ctx))
}

func (suite *KopiaSimpleRepoIntegrationSuite) TestBackupAndRestoreSingleItem() {
	t := suite.T()

	c, err := suite.w.RestoreSingleItem(
		suite.ctx,
		string(suite.snapshotID),
		append(testPath, testFileName),
	)
	require.NoError(t, err)

	assert.Equal(t, c.FullPath(), testPath)

	count := 0

	for resultStream := range c.Items() {
		buf, err := ioutil.ReadAll(resultStream.ToReader())
		require.NoError(t, err)
		assert.Equal(t, buf, testFileData)

		count++
	}

	assert.Equal(t, 1, count)
}

// TestBackupAndRestoreSingleItem_Errors exercises the public RestoreSingleItem
// function.
func (suite *KopiaSimpleRepoIntegrationSuite) TestBackupAndRestoreSingleItem_Errors() {
	table := []struct {
		name       string
		snapshotID string
		path       []string
	}{
		{
			"EmptyPath",
			string(suite.snapshotID),
			[]string{},
		},
		{
			"NoSnapshot",
			"foo",
			append(testPath, testFileName),
		},
		{
			"TargetNotAFile",
			string(suite.snapshotID),
			testPath[:2],
		},
		{
			"NonExistentFile",
			string(suite.snapshotID),
			append(testPath, "subdir", "foo"),
		},
	}

	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			_, err := suite.w.RestoreSingleItem(
				suite.ctx,
				test.snapshotID,
				test.path,
			)
			require.Error(t, err)
		})
	}
}

func (suite *KopiaSimpleRepoIntegrationSuite) TestRestoreMultipleItems() {
	t := suite.T()
	ctx := context.Background()

	k, err := openKopiaRepo(t, ctx)
	require.NoError(t, err)

	w := &Wrapper{k}

	tid := uuid.NewString()
	p1 := []string{
		tid,
		service,
		"uid",
		category,
		"fid",
	}
	p2 := []string{
		tid,
		service,
		"uid2",
		category,
		"fid",
	}
	dc1 := mockconnector.NewMockExchangeCollection(p1, 1)
	dc2 := mockconnector.NewMockExchangeCollection(p2, 1)

	fp1 := append(p1, dc1.Names[0])
	fp2 := append(p2, dc2.Names[0])

	stats, _, err := w.BackupCollections(ctx, []data.Collection{dc1, dc2})
	require.NoError(t, err)

	expected := map[string][]byte{
		stdpath.Join(fp1...): dc1.Data[0],
		stdpath.Join(fp2...): dc2.Data[0],
	}

	result, err := w.RestoreMultipleItems(
		ctx,
		string(stats.SnapshotID),
		[][]string{fp1, fp2})

	require.NoError(t, err)
	assert.Equal(t, 2, len(result))

	testForFiles(t, expected, result)
}

func (suite *KopiaSimpleRepoIntegrationSuite) TestRestoreMultipleItems_Errors() {
	table := []struct {
		name       string
		snapshotID string
		paths      [][]string
	}{
		{
			"EmptyPaths",
			string(suite.snapshotID),
			[][]string{{}},
		},
		{
			"NoSnapshot",
			"foo",
			[][]string{append(testPath, testFileName)},
		},
		{
			"TargetNotAFile",
			string(suite.snapshotID),
			[][]string{testPath[:2]},
		},
		{
			"NonExistentFile",
			string(suite.snapshotID),
			[][]string{append(testPath, "subdir", "foo")},
		},
	}

	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			_, err := suite.w.RestoreMultipleItems(
				suite.ctx,
				test.snapshotID,
				test.paths,
			)
			require.Error(t, err)
		})
	}
}

func (suite *KopiaIntegrationSuite) TestDeleteSnapshot() {
	t := suite.T()

	dc1 := mockconnector.NewMockExchangeCollection(
		[]string{"a-tenant", service, "user1", category, testInboxDir},
		5,
	)
	collections := []data.Collection{
		dc1,
		mockconnector.NewMockExchangeCollection(
			[]string{"a-tenant", service, "user2", category, testInboxDir},
			42,
		),
	}

	bs, _, err := suite.w.BackupCollections(suite.ctx, collections)
	require.NoError(t, err)

	snapshotID := bs.SnapshotID
	assert.NoError(t, suite.w.DeleteSnapshot(suite.ctx, snapshotID))

	// assert the deletion worked
	itemPath := []string{"a-tenant", "user1", "emails", dc1.Names[0]}
	_, err = suite.w.RestoreSingleItem(suite.ctx, snapshotID, itemPath)
	assert.Error(t, err, "snapshot should be deleted")
}

func (suite *KopiaIntegrationSuite) TestDeleteSnapshot_BadIDs() {
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
