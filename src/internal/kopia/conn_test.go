package kopia

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	ctesting "github.com/alcionai/corso/internal/testing"
)

func openKopiaRepo(t *testing.T, ctx context.Context) (*conn, error) {
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

// ---------------
// unit tests
// ---------------
type WrapperUnitSuite struct {
	suite.Suite
}

func TestWrapperUnitSuite(t *testing.T) {
	suite.Run(t, new(WrapperUnitSuite))
}

func (suite *WrapperUnitSuite) TestCloseWithoutOpenDoesNotCrash() {
	ctx := context.Background()

	k := conn{}
	assert.NotPanics(suite.T(), func() {
		k.Close(ctx)
	})
}

// ---------------
// integration tests that use kopia
// ---------------
type WrapperIntegrationSuite struct {
	suite.Suite
}

func TestWrapperIntegrationSuite(t *testing.T) {
	if err := ctesting.RunOnAny(
		ctesting.CorsoCITests,
		ctesting.CorsoKopiaWrapperTests,
	); err != nil {
		t.Skip()
	}

	suite.Run(t, new(WrapperIntegrationSuite))
}

func (suite *WrapperIntegrationSuite) SetupSuite() {
	_, err := ctesting.GetRequiredEnvVars(ctesting.AWSStorageCredEnvs...)
	require.NoError(suite.T(), err)
}

func (suite *WrapperIntegrationSuite) TestCloseTwiceDoesNotCrash() {
	ctx := context.Background()
	t := suite.T()

	k, err := openKopiaRepo(t, ctx)
	require.NoError(t, err)
	assert.NoError(t, k.Close(ctx))
	assert.Nil(t, k.Repository)
	assert.NoError(t, k.Close(ctx))
}

func (suite *WrapperIntegrationSuite) TestCloseAfterWrap() {
	ctx := context.Background()
	t := suite.T()

	k, err := openKopiaRepo(t, ctx)
	require.NoError(t, err)

	k.wrap()

	assert.Equal(t, 2, k.refCount)

	require.NoError(t, k.Close(ctx))
	assert.NotNil(t, k.Repository)
	assert.Equal(t, 1, k.refCount)

	require.NoError(t, k.Close(ctx))
	assert.Nil(t, k.Repository)
	assert.Equal(t, 0, k.refCount)
}

func (suite *WrapperIntegrationSuite) TestOpenAfterClose() {
	ctx := context.Background()
	t := suite.T()

	k, err := openKopiaRepo(t, ctx)
	require.NoError(t, err)

	assert.NoError(t, k.Close(ctx))
	assert.Error(t, k.wrap())
}
