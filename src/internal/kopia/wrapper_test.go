package kopia

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"testing"

	"github.com/kopia/kopia/fs"
	"github.com/kopia/kopia/fs/virtualfs"
	"github.com/kopia/kopia/repo/manifest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/internal/connector"
	"github.com/alcionai/corso/internal/connector/mockconnector"
	ctesting "github.com/alcionai/corso/internal/testing"
)

const (
	testTenant   = "a-tenant"
	testUser     = "user1"
	testEmailDir = "mail"
	testFileUUID = "a-file"
)

var (
	testPath     = []string{testTenant, testUser, testEmailDir}
	testFileData = []byte("abcdefghijklmnopqrstuvwxyz")
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
	dirTree, err := inflateDirTree(ctx, collections)
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
}

func (suite *KopiaUnitSuite) TestBuildDirectoryTree_NoAncestorDirs() {
	ctesting.LogTimeOfTest(suite.T())

	ctx := context.Background()
	emails := "emails"

	expectedFileCount := 42

	collections := []connector.DataCollection{
		mockconnector.NewMockExchangeDataCollection(
			[]string{emails},
			expectedFileCount,
		),
	}

	// Returned directory structure should look like:
	// - emails
	//   - 42 separate files
	dirTree, err := inflateDirTree(ctx, collections)
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
			_, err := inflateDirTree(ctx, test.layout)
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

	stats, err := suite.w.BackupCollections(suite.ctx, collections)
	assert.NoError(t, err)
	assert.Equal(t, stats.TotalFileCount, 47)
	assert.Equal(t, stats.TotalDirectoryCount, 5)
	assert.Equal(t, stats.IgnoredErrorCount, 0)
	assert.Equal(t, stats.ErrorCount, 0)
	assert.False(t, stats.Incomplete)
}

type KopiaSimpleRepoIntegrationSuite struct {
	suite.Suite
	w          *Wrapper
	ctx        context.Context
	snapshotID manifest.ID
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
				&kopiaDataStream{
					uuid:   testFileUUID,
					reader: io.NopCloser(bytes.NewReader(testFileData)),
				},
			},
		},
	}

	stats, err := suite.w.BackupCollections(suite.ctx, collections)
	require.NoError(t, err)
	require.Equal(t, stats.TotalFileCount, 1)
	require.Equal(t, stats.TotalDirectoryCount, 3)
	require.Equal(t, stats.IgnoredErrorCount, 0)
	require.Equal(t, stats.ErrorCount, 0)
	require.False(t, stats.Incomplete)

	suite.snapshotID = manifest.ID(stats.SnapshotID)
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
		name           string
		snapshotIDFunc func(manifest.ID) manifest.ID
		path           []string
	}{
		{
			"NoSnapshot",
			func(manifest.ID) manifest.ID {
				return manifest.ID("foo")
			},
			append(testPath, testFileUUID),
		},
		{
			"TargetNotAFile",
			func(m manifest.ID) manifest.ID {
				return m
			},
			testPath[:2],
		},
		{
			"NonExistentFile",
			func(m manifest.ID) manifest.ID {
				return m
			},
			append(testPath, "foo"),
		},
	}

	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			_, err := suite.w.RestoreSingleItem(
				suite.ctx,
				string(test.snapshotIDFunc(suite.snapshotID)),
				test.path,
			)
			require.Error(t, err)
		})
	}
}

// TestBackupAndRestoreSingleItem_Errors2 exercises some edge cases in the
// package-private restoreSingleItem function. It helps ensure kopia behaves the
// way we expect.
func (suite *KopiaSimpleRepoIntegrationSuite) TestBackupAndRestoreSingleItem_Errors2() {
	table := []struct {
		name        string
		rootDirFunc func(*testing.T, context.Context, *Wrapper) fs.Entry
		path        []string
	}{
		{
			"FileAsRoot",
			func(t *testing.T, ctx context.Context, w *Wrapper) fs.Entry {
				return virtualfs.StreamingFileFromReader(testFileUUID, bytes.NewReader(testFileData))
			},
			append(testPath[1:], testFileUUID),
		},
		{
			"NoRootDir",
			func(t *testing.T, ctx context.Context, w *Wrapper) fs.Entry {
				return nil
			},
			append(testPath[1:], testFileUUID),
		},
	}

	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			_, err := suite.w.restoreSingleItem(
				suite.ctx,
				test.rootDirFunc(t, suite.ctx, suite.w),
				test.path,
			)
			require.Error(t, err)
		})
	}
}
