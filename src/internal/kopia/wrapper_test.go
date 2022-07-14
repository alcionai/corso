package kopia

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"path"
	"testing"

	"github.com/google/uuid"
	"github.com/kopia/kopia/fs"
	"github.com/kopia/kopia/fs/virtualfs"
	"github.com/kopia/kopia/repo/manifest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/internal/connector"
	"github.com/alcionai/corso/internal/connector/mockconnector"
	"github.com/alcionai/corso/internal/kopia/mockkopia"
	ctesting "github.com/alcionai/corso/internal/testing"
	"github.com/alcionai/corso/pkg/backup"
)

const (
	testTenant     = "a-tenant"
	testUser       = "user1"
	testEmailDir   = "mail"
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
	testPath      = []string{testTenant, testUser, testEmailDir, testInboxDir}
	testPath2     = []string{testTenant, testUser, testEmailDir, testArchiveDir}
	testFileData  = []byte("abcdefghijklmnopqrstuvwxyz")
	testFileData2 = []byte("zyxwvutsrqponmlkjihgfedcba")
	testFileData3 = []byte("foo")
	testFileData4 = []byte("bar")
	testFileData5 = []byte("baz")
	// Intentional duplicate to make sure all files are scanned during recovery
	// (contrast to behavior of snapshotfs.TreeWalker).
	testFileData6 = testFileData
)

func entriesToNames(entries []fs.Entry) []string {
	res := make([]string, 0, len(entries))

	for _, e := range entries {
		res = append(res, e.Name())
	}

	return res
}

func testForFiles(
	t *testing.T,
	expected map[string][]byte,
	collections []connector.DataCollection,
) {
	count := 0
	for _, c := range collections {
		for s := range c.Items() {
			count++

			fullPath := path.Join(append(c.FullPath(), s.UUID())...)

			expected, ok := expected[fullPath]
			require.True(t, ok, "unexpected file with path %q", fullPath)

			buf, err := ioutil.ReadAll(s.ToReader())
			require.NoError(t, err, "reading collection item: %s", fullPath)

			assert.Equal(t, expected, buf, "comparing collection item: %s", fullPath)
		}
	}

	assert.Equal(t, len(expected), count)
}

// ---------------
// unit tests
// ---------------
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
	ctesting.LogTimeOfTest(suite.T())

	ctx := context.Background()
	tenant := "a-tenant"
	user1 := "user1"
	user2 := "user2"
	emails := "emails"

	expectedFileCount := map[string]int{
		user1: 5,
		user2: 42,
	}

	details := &backup.Details{}

	collections := []connector.DataCollection{
		mockconnector.NewMockExchangeDataCollection(
			[]string{tenant, user1, emails},
			expectedFileCount[user1],
		),
		mockconnector.NewMockExchangeDataCollection(
			[]string{tenant, user2, emails},
			expectedFileCount[user2],
		),
	}

	// Returned directory structure should look like:
	// - a-tenant
	//   - user1
	//     - emails
	//       - 5 separate files
	//   - user2
	//     - emails
	//       - 42 separate files
	dirTree, err := inflateDirTree(ctx, collections, details)
	require.NoError(suite.T(), err)
	assert.Equal(suite.T(), dirTree.Name(), tenant)

	entries, err := fs.GetAllEntries(ctx, dirTree)
	require.NoError(suite.T(), err)
	names := entriesToNames(entries)
	assert.Len(suite.T(), names, 2)
	assert.Contains(suite.T(), names, user1)
	assert.Contains(suite.T(), names, user2)

	for _, entry := range entries {
		dir, ok := entry.(fs.Directory)
		require.True(suite.T(), ok)

		subEntries, err := fs.GetAllEntries(ctx, dir)
		require.NoError(suite.T(), err)
		require.Len(suite.T(), subEntries, 1)
		assert.Contains(suite.T(), subEntries[0].Name(), emails)

		subDir := subEntries[0].(fs.Directory)
		emailFiles, err := fs.GetAllEntries(ctx, subDir)
		require.NoError(suite.T(), err)
		assert.Len(suite.T(), emailFiles, expectedFileCount[entry.Name()])
	}

	totalFileCount := 0
	for _, c := range expectedFileCount {
		totalFileCount += c
	}

	assert.Len(suite.T(), details.Entries, totalFileCount)
}

func (suite *KopiaUnitSuite) TestBuildDirectoryTree_NoAncestorDirs() {
	ctesting.LogTimeOfTest(suite.T())

	ctx := context.Background()
	emails := "emails"

	expectedFileCount := 42

	details := &backup.Details{}
	collections := []connector.DataCollection{
		mockconnector.NewMockExchangeDataCollection(
			[]string{emails},
			expectedFileCount,
		),
	}

	// Returned directory structure should look like:
	// - emails
	//   - 42 separate files
	dirTree, err := inflateDirTree(ctx, collections, details)
	require.NoError(suite.T(), err)
	assert.Equal(suite.T(), dirTree.Name(), emails)

	entries, err := fs.GetAllEntries(ctx, dirTree)
	require.NoError(suite.T(), err)
	assert.Len(suite.T(), entries, 42)
}

func (suite *KopiaUnitSuite) TestBuildDirectoryTree_Fails() {
	table := []struct {
		name   string
		layout []connector.DataCollection
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
			[]connector.DataCollection{
				mockconnector.NewMockExchangeDataCollection(
					[]string{"user1", "emails"},
					5,
				),
				mockconnector.NewMockExchangeDataCollection(
					[]string{"user2", "emails"},
					42,
				),
			},
		},
		{
			"NoCollectionPath",
			[]connector.DataCollection{
				mockconnector.NewMockExchangeDataCollection(
					nil,
					5,
				),
			},
		},
		{
			"MixedDirectory",
			// Directory structure would look like (but should return error):
			// - a-tenant
			//   - user1
			//     - emails
			//       - 5 separate files
			//     - 42 separate files
			[]connector.DataCollection{
				mockconnector.NewMockExchangeDataCollection(
					[]string{"a-tenant", "user1", "emails"},
					5,
				),
				mockconnector.NewMockExchangeDataCollection(
					[]string{"a-tenant", "user1"},
					42,
				),
			},
		},
	}

	for _, test := range table {
		ctx := context.Background()

		suite.T().Run(test.name, func(t *testing.T) {
			details := &backup.Details{}
			_, err := inflateDirTree(ctx, test.layout, details)
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

func (suite *KopiaUnitSuite) TestRestoreDirectory_FailGettingReader() {
	ctx := context.Background()
	t := suite.T()

	expectedStreamData := map[string][]byte{
		path.Join(testInboxDir, testFileName):  testFileData,
		path.Join(testInboxDir, testFileName3): testFileData3,
	}

	dirs := virtualfs.NewStaticDirectory(testInboxDir, []fs.Entry{
		&mockkopia.MockFile{
			Entry: &mockkopia.MockEntry{
				EntryName: testFileName,
				EntryMode: mockkopia.DefaultPermissions,
			},
			Data: testFileData,
		},
		&mockkopia.MockFile{
			Entry: &mockkopia.MockEntry{
				EntryName: testFileName2,
				EntryMode: mockkopia.DefaultPermissions,
			},
			OpenErr: assert.AnError,
		},
		&mockkopia.MockFile{
			Entry: &mockkopia.MockEntry{
				EntryName: testFileName3,
				EntryMode: mockkopia.DefaultPermissions,
			},
			Data: testFileData3,
		},
	})

	collections, err := restoreSubtree(ctx, dirs, nil)
	assert.Error(t, err)

	assert.Len(t, collections, 1)
	testForFiles(t, expectedStreamData, collections)
}

func (suite *KopiaUnitSuite) TestRestoreDirectory_FailWrongItemType() {
	ctx := context.Background()
	t := suite.T()

	expectedStreamData := map[string][]byte{
		path.Join(testEmailDir, testInboxDir, testFileName):    testFileData,
		path.Join(testEmailDir, testArchiveDir, testFileName3): testFileData3,
	}

	dirs := virtualfs.NewStaticDirectory(testEmailDir, []fs.Entry{
		virtualfs.NewStaticDirectory(testInboxDir, []fs.Entry{
			&mockkopia.MockFile{
				Entry: &mockkopia.MockEntry{
					EntryName: testFileName,
					EntryMode: mockkopia.DefaultPermissions,
				},
				Data: testFileData,
			},
		}),
		virtualfs.NewStaticDirectory("foo", []fs.Entry{
			virtualfs.StreamingFileFromReader(
				testFileName2, bytes.NewReader(testFileData2)),
		}),
		virtualfs.NewStaticDirectory(testArchiveDir, []fs.Entry{
			&mockkopia.MockFile{
				Entry: &mockkopia.MockEntry{
					EntryName: testFileName3,
					EntryMode: mockkopia.DefaultPermissions,
				},
				Data: testFileData3,
			},
		}),
	})

	collections, err := restoreSubtree(ctx, dirs, nil)
	assert.Error(t, err)

	assert.Len(t, collections, 2)
	testForFiles(t, expectedStreamData, collections)
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
	if err := ctesting.RunOnAny(
		ctesting.CorsoCITests,
		ctesting.CorsoKopiaWrapperTests,
	); err != nil {
		t.Skip()
	}

	suite.Run(t, new(KopiaIntegrationSuite))
}

func (suite *KopiaIntegrationSuite) SetupSuite() {
	_, err := ctesting.GetRequiredEnvVars(ctesting.AWSStorageCredEnvs...)
	require.NoError(suite.T(), err)
}

func (suite *KopiaIntegrationSuite) SetupTest() {
	suite.ctx = context.Background()
	c, err := openKopiaRepo(suite.T(), suite.ctx)
	require.NoError(suite.T(), err)
	suite.w = &Wrapper{c}
}

func (suite *KopiaIntegrationSuite) TearDownTest() {
	assert.NoError(suite.T(), suite.w.Close(suite.ctx))
}

func (suite *KopiaIntegrationSuite) TestBackupCollections() {
	t := suite.T()

	collections := []connector.DataCollection{
		mockconnector.NewMockExchangeDataCollection(
			[]string{"a-tenant", "user1", "emails"},
			5,
		),
		mockconnector.NewMockExchangeDataCollection(
			[]string{"a-tenant", "user2", "emails"},
			42,
		),
	}

	stats, rp, err := suite.w.BackupCollections(suite.ctx, collections)
	assert.NoError(t, err)
	assert.Equal(t, stats.TotalFileCount, 47)
	assert.Equal(t, stats.TotalDirectoryCount, 5)
	assert.Equal(t, stats.IgnoredErrorCount, 0)
	assert.Equal(t, stats.ErrorCount, 0)
	assert.False(t, stats.Incomplete)
	assert.Len(t, rp.Entries, 47)
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
	if err := ctesting.RunOnAny(
		ctesting.CorsoCITests,
		ctesting.CorsoKopiaWrapperTests,
	); err != nil {
		t.Skip()
	}

	suite.Run(t, new(KopiaSimpleRepoIntegrationSuite))
}

func (suite *KopiaSimpleRepoIntegrationSuite) SetupSuite() {
	_, err := ctesting.GetRequiredEnvVars(ctesting.AWSStorageCredEnvs...)
	require.NoError(suite.T(), err)
}

func (suite *KopiaSimpleRepoIntegrationSuite) SetupTest() {
	t := suite.T()
	suite.ctx = context.Background()
	c, err := openKopiaRepo(t, suite.ctx)
	require.NoError(t, err)

	suite.w = &Wrapper{c}

	collections := []connector.DataCollection{
		&kopiaDataCollection{
			path: testPath,
			streams: []connector.DataStream{
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
			path: testPath2,
			streams: []connector.DataStream{
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
	require.Equal(t, stats.TotalDirectoryCount, 5)
	require.Equal(t, stats.IgnoredErrorCount, 0)
	require.False(t, stats.Incomplete)
	assert.Len(t, rp.Entries, 6)

	suite.snapshotID = manifest.ID(stats.SnapshotID)

	// path.Join doesn't like (testPath..., testFileName).
	suite.inboxExpectedFiles = map[string][]byte{
		path.Join(append(testPath, testFileName)...):  testFileData,
		path.Join(append(testPath, testFileName2)...): testFileData2,
	}
	suite.archiveExpectedFiles = map[string][]byte{
		path.Join(append(testPath2, testFileName3)...): testFileData3,
		path.Join(append(testPath2, testFileName4)...): testFileData4,
		path.Join(append(testPath2, testFileName5)...): testFileData5,
		path.Join(append(testPath2, testFileName6)...): testFileData6,
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

func (suite *KopiaSimpleRepoIntegrationSuite) TestBackupRestoreDirectory() {
	table := []struct {
		name          string
		dirPath       []string
		expectedFiles map[string][]byte
	}{
		{
			"RecoverUser",
			[]string{testTenant, testUser},
			suite.allExpectedFiles,
		},
		{
			"RecoverMail",
			[]string{testTenant, testUser, testEmailDir},
			suite.allExpectedFiles,
		},
		{
			"RecoverInbox",
			[]string{testTenant, testUser, testEmailDir, testInboxDir},
			suite.inboxExpectedFiles,
		},
		{
			"RecoverArchive",
			[]string{testTenant, testUser, testEmailDir, testArchiveDir},
			suite.archiveExpectedFiles,
		},
	}

	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			collections, err := suite.w.RestoreDirectory(
				suite.ctx, string(suite.snapshotID), test.dirPath)
			require.NoError(t, err)

			testForFiles(t, test.expectedFiles, collections)
		})
	}
}

func (suite *KopiaSimpleRepoIntegrationSuite) TestBackupRestoreDirectory_Errors() {
	table := []struct {
		name       string
		snapshotID string
		dirPath    []string
	}{
		{
			"EmptyPath",
			string(suite.snapshotID),
			[]string{},
		},
		{
			"BadSnapshotID",
			"foo",
			[]string{testTenant, testUser, testEmailDir},
		},
		{
			"NotADirectory",
			string(suite.snapshotID),
			append(testPath, testFileName),
		},
		{
			"NonExistantDirectory",
			string(suite.snapshotID),
			[]string{testTenant, testUser, testEmailDir, "subdir"},
		},
	}

	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			_, err := suite.w.RestoreDirectory(
				suite.ctx, test.snapshotID, test.dirPath)
			assert.Error(t, err)
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
	p1 := []string{tid, "uid", "emails", "fid"}
	p2 := []string{tid, "uid2", "emails", "fid"}
	dc1 := mockconnector.NewMockExchangeDataCollection(p1, 1)
	dc2 := mockconnector.NewMockExchangeDataCollection(p2, 1)
	fp1 := append(p1, dc1.Names[0])
	fp2 := append(p2, dc2.Names[0])

	stats, _, err := w.BackupCollections(ctx, []connector.DataCollection{dc1, dc2})
	require.NoError(t, err)

	expected := map[string][]byte{
		path.Join(fp1...): dc1.Data[0],
		path.Join(fp2...): dc2.Data[0],
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
