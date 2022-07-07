package kopia

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"path"
	"testing"

	"github.com/kopia/kopia/fs"
	"github.com/kopia/kopia/repo/manifest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/internal/connector"
	"github.com/alcionai/corso/internal/connector/mockconnector"
	ctesting "github.com/alcionai/corso/internal/testing"
	"github.com/alcionai/corso/pkg/backup"
)

const (
	testTenant     = "a-tenant"
	testUser       = "user1"
	testEmailDir   = "mail"
	testInboxDir   = "inbox"
	testArchiveDir = "archive"
	testFileUUID   = "file1"
	testFileUUID2  = "file2"
	testFileUUID3  = "file3"
	testFileUUID4  = "file4"
	testFileUUID5  = "file5"
	testFileUUID6  = "file6"
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
					ID:     testFileUUID,
					Reader: io.NopCloser(bytes.NewReader(testFileData)),
				},
				&mockconnector.MockExchangeData{
					ID:     testFileUUID2,
					Reader: io.NopCloser(bytes.NewReader(testFileData2)),
				},
			},
		},
		&kopiaDataCollection{
			path: testPath2,
			streams: []connector.DataStream{
				&mockconnector.MockExchangeData{
					ID:     testFileUUID3,
					Reader: io.NopCloser(bytes.NewReader(testFileData3)),
				},
				&mockconnector.MockExchangeData{
					ID:     testFileUUID4,
					Reader: io.NopCloser(bytes.NewReader(testFileData4)),
				},
				&mockconnector.MockExchangeData{
					ID:     testFileUUID5,
					Reader: io.NopCloser(bytes.NewReader(testFileData5)),
				},
				&mockconnector.MockExchangeData{
					ID:     testFileUUID6,
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

	// path.Join doesn't like (testPath..., testFileUUID).
	suite.inboxExpectedFiles = map[string][]byte{
		path.Join(append(testPath, testFileUUID)...):  testFileData,
		path.Join(append(testPath, testFileUUID2)...): testFileData2,
	}
	suite.archiveExpectedFiles = map[string][]byte{
		path.Join(append(testPath2, testFileUUID3)...): testFileData3,
		path.Join(append(testPath2, testFileUUID4)...): testFileData4,
		path.Join(append(testPath2, testFileUUID5)...): testFileData5,
		path.Join(append(testPath2, testFileUUID6)...): testFileData6,
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
		append(testPath, testFileUUID),
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
			append(testPath, testFileUUID),
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

			found := 0
			for _, c := range collections {
				for ds := range c.Items() {
					found++

					fullPath := path.Join(append(c.FullPath(), ds.UUID())...)
					expected, ok := test.expectedFiles[fullPath]
					require.True(t, ok, "unexpected path %q", fullPath)

					buf, err := ioutil.ReadAll(ds.ToReader())
					require.NoError(t, err)

					assert.Equal(t, expected, buf)
				}
			}

			assert.Equal(t, len(test.expectedFiles), found)
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
			append(testPath, testFileUUID),
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
