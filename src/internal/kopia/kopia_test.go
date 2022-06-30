package kopia

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
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
	"github.com/alcionai/corso/pkg/restorepoint"
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

func openKopiaRepo(t *testing.T, ctx context.Context) (*KopiaWrapper, error) {
	storage, err := ctesting.NewPrefixedS3Storage(t)
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

	details := &restorepoint.Details{}

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
	for _, e := range details.Entries {
		b, err := json.MarshalIndent(e, "", "  ")
		require.NoError(suite.T(), err)
		fmt.Print(string(b))
	}
}

func (suite *KopiaUnitSuite) TestBuildDirectoryTree_NoAncestorDirs() {
	ctesting.LogTimeOfTest(suite.T())

	ctx := context.Background()
	emails := "emails"

	expectedFileCount := 42

	details := &restorepoint.Details{}
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
			details := &restorepoint.Details{}
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

func (suite *KopiaIntegrationSuite) TestCloseTwiceDoesNotCrash() {
	ctx := context.Background()
	t := suite.T()

	k, err := openKopiaRepo(t, ctx)
	require.NoError(t, err)
	assert.NoError(t, k.Close(ctx))
	assert.Nil(t, k.rep)
	assert.NoError(t, k.Close(ctx))
}

func (suite *KopiaIntegrationSuite) TestBackupCollections() {
	ctx := context.Background()
	t := suite.T()

	k, err := openKopiaRepo(t, ctx)
	require.NoError(t, err)
	defer func() {
		assert.NoError(t, k.Close(ctx))
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
	assert.NoError(t, err)
	assert.Equal(t, stats.TotalFileCount, 47)
	assert.Equal(t, stats.TotalDirectoryCount, 5)
	assert.Equal(t, stats.IgnoredErrorCount, 0)
	assert.Equal(t, stats.ErrorCount, 0)
	assert.False(t, stats.Incomplete)
}

func setupSimpleRepo(t *testing.T, ctx context.Context, k *KopiaWrapper) manifest.ID {
	collections := []connector.DataCollection{
		&singleItemCollection{
			path: testPath,
			stream: &kopiaDataStream{
				uuid:   testFileUUID,
				reader: io.NopCloser(bytes.NewReader(testFileData)),
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

	return manifest.ID(stats.SnapshotID)
}

func (suite *KopiaIntegrationSuite) TestBackupAndRestoreSingleItem() {
	ctx := context.Background()
	t := suite.T()

	k, err := openKopiaRepo(t, ctx)
	require.NoError(t, err)
	defer func() {
		assert.NoError(t, k.Close(ctx))
	}()

	id := setupSimpleRepo(t, ctx, k)

	c, err := k.RestoreSingleItem(
		ctx,
		string(id),
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
func (suite *KopiaIntegrationSuite) TestBackupAndRestoreSingleItem_Errors() {
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
			ctx := context.Background()
			ctesting.LogTimeOfTest(t)

			k, err := openKopiaRepo(t, ctx)
			require.NoError(t, err)
			defer func() {
				assert.NoError(t, k.Close(ctx))
			}()

			id := setupSimpleRepo(t, ctx, k)

			_, err = k.RestoreSingleItem(
				ctx,
				string(test.snapshotIDFunc(id)),
				test.path,
			)
			require.Error(t, err)
		})
	}
}

// TestBackupAndRestoreSingleItem_Errors2 exercises some edge cases in the
// package-private restoreSingleItem function. It helps ensure kopia behaves the
// way we expect.
func (suite *KopiaIntegrationSuite) TestBackupAndRestoreSingleItem_Errors2() {
	table := []struct {
		name        string
		rootDirFunc func(*testing.T, context.Context, *KopiaWrapper) fs.Entry
		path        []string
	}{
		{
			"FileAsRoot",
			func(t *testing.T, ctx context.Context, k *KopiaWrapper) fs.Entry {
				return virtualfs.StreamingFileFromReader(testFileUUID, bytes.NewReader(testFileData))
			},
			append(testPath[1:], testFileUUID),
		},
		{
			"NoRootDir",
			func(t *testing.T, ctx context.Context, k *KopiaWrapper) fs.Entry {
				return nil
			},
			append(testPath[1:], testFileUUID),
		},
	}

	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			ctx := context.Background()

			k, err := openKopiaRepo(t, ctx)
			require.NoError(t, err)
			defer func() {
				assert.NoError(t, k.Close(ctx))
			}()

			setupSimpleRepo(t, ctx, k)

			_, err = k.restoreSingleItem(
				ctx,
				test.rootDirFunc(t, ctx, k),
				test.path,
			)
			require.Error(t, err)
		})
	}
}
