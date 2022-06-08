package kopia

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

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

// ---------------
// unit tests
// ---------------
type KopiaUnitSuite struct {
	suite.Suite
}

func (suite *KopiaUnitSuite) TestCloseWithoutOpenDoesNotCrash() {
	ctx := context.Background()
	ctesting.LogTimeOfTest(suite.T())

	k := KopiaWrapper{}
	assert.NotPanics(suite.T(), func() {
		k.Close(ctx)
	})
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
	require.NoError(suite.T(), ctesting.CheckS3EnvVars())
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

	stats, err := k.BackupCollections(ctx, nil)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), stats.TotalFileCount, 0)
	assert.Equal(suite.T(), stats.TotalDirectoryCount, 1)
	assert.Equal(suite.T(), stats.IgnoredErrorCount, 0)
	assert.Equal(suite.T(), stats.ErrorCount, 0)
	assert.False(suite.T(), stats.Incomplete)
}
