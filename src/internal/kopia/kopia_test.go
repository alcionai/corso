package kopia

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"testing"

	"github.com/kopia/kopia/fs"
	"github.com/kopia/kopia/fs/virtualfs"
	"github.com/kopia/kopia/repo"
	"github.com/kopia/kopia/snapshot"
	"github.com/kopia/kopia/snapshot/snapshotfs"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/internal/connector"
	"github.com/alcionai/corso/internal/connector/mockconnector"
	ctesting "github.com/alcionai/corso/internal/testing"
)

func openKopiaRepo(ctx context.Context, prefix string) (*KopiaWrapper, error) {
	storage, err := ctesting.NewS3Storage(prefix)
	if err != nil {
		return nil, err
	}

	k := New(storage)
	if err = k.Initialize(ctx); err != nil {
		return nil, err
	}

	return k, nil
}

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

func (suite *KopiaUnitSuite) TestCloseWithoutOpenDoesNotCrash() {
	ctx := context.Background()
	ctesting.LogTimeOfTest(suite.T())

	k := KopiaWrapper{}
	assert.NotPanics(suite.T(), func() {
		k.Close(ctx)
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

		subDir, ok := subEntries[0].(fs.Directory)
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
	_, err := ctesting.GetRequiredEnvVars(ctesting.AWSCredentialEnvs...)
	require.NoError(suite.T(), err)
}

func (suite *KopiaIntegrationSuite) TestCloseTwiceDoesNotCrash() {
	ctx := context.Background()
	timeOfTest := ctesting.LogTimeOfTest(suite.T())

	k, err := openKopiaRepo(ctx, "init-s3-"+timeOfTest)
	require.NoError(suite.T(), err)
	assert.NoError(suite.T(), k.Close(ctx))
	assert.Nil(suite.T(), k.rep)
	assert.NoError(suite.T(), k.Close(ctx))
}

func (suite *KopiaIntegrationSuite) TestBackupCollections() {
	ctx := context.Background()
	timeOfTest := ctesting.LogTimeOfTest(suite.T())

	k, err := openKopiaRepo(ctx, "init-s3-"+timeOfTest)
	require.NoError(suite.T(), err)
	defer func() {
		assert.NoError(suite.T(), k.Close(ctx))
	}()

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

	stats, err := k.BackupCollections(ctx, collections)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), stats.TotalFileCount, 47)
	assert.Equal(suite.T(), stats.TotalDirectoryCount, 5)
	assert.Equal(suite.T(), stats.IgnoredErrorCount, 0)
	assert.Equal(suite.T(), stats.ErrorCount, 0)
	assert.False(suite.T(), stats.Incomplete)
}

// TODO(ashmrtn): Update this once we have a helper for getting the snapshot
// root.
func getSnapshotRoot(
	t *testing.T,
	ctx context.Context,
	rep repo.Repository,
	rootName string,
) fs.Entry {
	si := snapshot.SourceInfo{
		Host:     kTestHost,
		UserName: kTestUser,
		Path:     rootName,
	}

	manifests, err := snapshot.ListSnapshots(ctx, rep, si)
	require.NoError(t, err)
	require.Len(t, manifests, 1)

	rootDirEntry, err := snapshotfs.SnapshotRoot(rep, manifests[0])
	require.NoError(t, err)

	rootDir, ok := rootDirEntry.(fs.Directory)
	require.True(t, ok)

	return rootDir
}

func (suite *KopiaIntegrationSuite) TestBackupAndRestoreSingleItem() {
	ctx := context.Background()
	timeOfTest := ctesting.LogTimeOfTest(suite.T())
	t := suite.T()

	k, err := openKopiaRepo(ctx, "backup-restore-single-item-"+timeOfTest)
	require.NoError(t, err)
	defer func() {
		assert.NoError(t, k.Close(ctx))
	}()

	path := []string{"a-tenant", "user1", "emails"}
	fileUUID := "a-file"
	data := []byte("abcdefghijklmnopqrstuvwxyz")
	collections := []connector.DataCollection{
		&singleItemCollection{
			path: path,
			stream: &kopiaDataStream{
				uuid:   fileUUID,
				reader: io.NopCloser(bytes.NewReader(data)),
			},
		},
	}

	stats, err := k.BackupCollections(ctx, collections)
	require.NoError(t, err)
	require.Equal(t, stats.TotalFileCount, 1)
	require.Equal(t, stats.TotalDirectoryCount, 3)
	require.Equal(t, stats.IgnoredErrorCount, 0)
	require.Equal(t, stats.ErrorCount, 0)
	require.False(t, stats.Incomplete)

	rootDir := getSnapshotRoot(t, ctx, k.rep, "a-tenant")

	c, err := k.restoreSingleItem(ctx, rootDir, append(path[1:], fileUUID))
	require.NoError(t, err)

	assert.Equal(t, c.FullPath(), path)

	resultStream, err := c.NextItem()
	require.NoError(t, err)

	_, err = c.NextItem()
	assert.ErrorIs(t, err, io.EOF)

	buf, err := ioutil.ReadAll(resultStream.ToReader())
	require.NoError(t, err)
	assert.Equal(t, buf, data)
}

func (suite *KopiaIntegrationSuite) TestBackupAndRestoreSingleItem_FileAsRoot() {
	ctx := context.Background()
	timeOfTest := ctesting.LogTimeOfTest(suite.T())
	t := suite.T()

	k, err := openKopiaRepo(ctx, "restore-single-file-as-root-"+timeOfTest)
	require.NoError(t, err)
	defer func() {
		assert.NoError(t, k.Close(ctx))
	}()

	path := []string{"a-tenant", "user1", "emails"}
	fileUUID := "a-file"
	data := []byte("abcdefghijklmnopqrstuvwxyz")

	_, err = k.restoreSingleItem(
		ctx,
		virtualfs.StreamingFileFromReader("a-file", bytes.NewReader(data)),
		append(path[1:], fileUUID),
	)
	require.Error(t, err)
}

func (suite *KopiaIntegrationSuite) TestBackupAndRestoreSingleItem_NoRootDir() {
	ctx := context.Background()
	timeOfTest := ctesting.LogTimeOfTest(suite.T())
	t := suite.T()

	k, err := openKopiaRepo(ctx, "restore-single-no-root-"+timeOfTest)
	require.NoError(t, err)
	defer func() {
		assert.NoError(t, k.Close(ctx))
	}()

	path := []string{"a-tenant", "user1", "emails"}
	fileUUID := "a-file"
	data := []byte("abcdefghijklmnopqrstuvwxyz")
	collections := []connector.DataCollection{
		&singleItemCollection{
			path: path,
			stream: &kopiaDataStream{
				uuid:   fileUUID,
				reader: io.NopCloser(bytes.NewReader(data)),
			},
		},
	}

	stats, err := k.BackupCollections(ctx, collections)
	require.NoError(t, err)
	require.Equal(t, stats.TotalFileCount, 1)
	require.Equal(t, stats.TotalDirectoryCount, 3)
	require.Equal(t, stats.IgnoredErrorCount, 0)
	require.Equal(t, stats.ErrorCount, 0)
	require.False(t, stats.Incomplete)

	_, err = k.restoreSingleItem(ctx, nil, append(path[1:], fileUUID))
	require.Error(t, err)
}

func (suite *KopiaIntegrationSuite) TestBackupAndRestoreSingleItem_TargetNotAFile() {
	ctx := context.Background()
	timeOfTest := ctesting.LogTimeOfTest(suite.T())
	t := suite.T()

	k, err := openKopiaRepo(ctx, "restore-single-not-a-file-"+timeOfTest)
	require.NoError(t, err)
	defer func() {
		assert.NoError(t, k.Close(ctx))
	}()

	path := []string{"a-tenant", "user1", "emails"}
	fileUUID := "a-file"
	data := []byte("abcdefghijklmnopqrstuvwxyz")
	collections := []connector.DataCollection{
		&singleItemCollection{
			path: path,
			stream: &kopiaDataStream{
				uuid:   fileUUID,
				reader: io.NopCloser(bytes.NewReader(data)),
			},
		},
	}

	stats, err := k.BackupCollections(ctx, collections)
	require.NoError(t, err)
	require.Equal(t, stats.TotalFileCount, 1)
	require.Equal(t, stats.TotalDirectoryCount, 3)
	require.Equal(t, stats.IgnoredErrorCount, 0)
	require.Equal(t, stats.ErrorCount, 0)
	require.False(t, stats.Incomplete)

	rootDir := getSnapshotRoot(t, ctx, k.rep, "a-tenant")

	_, err = k.restoreSingleItem(ctx, rootDir, []string{path[1]})
	require.Error(t, err)
}

func (suite *KopiaIntegrationSuite) TestBackupAndRestoreSingleItem_NonExistentFile() {
	ctx := context.Background()
	timeOfTest := ctesting.LogTimeOfTest(suite.T())
	t := suite.T()

	k, err := openKopiaRepo(ctx, "restore-single-does-not-exist-"+timeOfTest)
	require.NoError(t, err)
	defer func() {
		assert.NoError(t, k.Close(ctx))
	}()

	path := []string{"a-tenant", "user1", "emails"}
	fileUUID := "a-file"
	data := []byte("abcdefghijklmnopqrstuvwxyz")
	collections := []connector.DataCollection{
		&singleItemCollection{
			path: path,
			stream: &kopiaDataStream{
				uuid:   fileUUID,
				reader: io.NopCloser(bytes.NewReader(data)),
			},
		},
	}

	stats, err := k.BackupCollections(ctx, collections)
	require.NoError(t, err)
	require.Equal(t, stats.TotalFileCount, 1)
	require.Equal(t, stats.TotalDirectoryCount, 3)
	require.Equal(t, stats.IgnoredErrorCount, 0)
	require.Equal(t, stats.ErrorCount, 0)
	require.False(t, stats.Incomplete)

	rootDir := getSnapshotRoot(t, ctx, k.rep, "a-tenant")

	_, err = k.restoreSingleItem(ctx, rootDir, append(path[1:], "foo"))
	require.Error(t, err)
}
