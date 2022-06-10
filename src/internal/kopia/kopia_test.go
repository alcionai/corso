package kopia

import (
	"context"
	"testing"

	"github.com/kopia/kopia/fs"
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
	if err != nil {
		suite.T().Fatal(err)
	}
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
