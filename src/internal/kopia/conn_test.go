package kopia

import (
	"context"
	"testing"

	"github.com/kopia/kopia/snapshot"
	"github.com/kopia/kopia/snapshot/policy"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/internal/tester"
)

//revive:disable:context-as-argument
func openKopiaRepo(t *testing.T, ctx context.Context) (*conn, error) {
	storage := tester.NewPrefixedS3Storage(t)

	k := NewConn(storage)
	if err := k.Initialize(ctx); err != nil {
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
	if err := tester.RunOnAny(
		tester.CorsoCITests,
		tester.CorsoKopiaWrapperTests,
	); err != nil {
		t.Skip()
	}

	suite.Run(t, new(WrapperIntegrationSuite))
}

func (suite *WrapperIntegrationSuite) SetupSuite() {
	_, err := tester.GetRequiredEnvVars(tester.AWSStorageCredEnvs...)
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

	require.NoError(t, k.wrap())

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

func (suite *WrapperIntegrationSuite) TestBadCompressorType() {
	ctx := context.Background()
	t := suite.T()

	k, err := openKopiaRepo(t, ctx)
	require.NoError(t, err)

	defer func() {
		assert.NoError(t, k.Close(ctx))
	}()

	assert.Error(t, k.Compression(ctx, "not-a-compressor"))
}

func (suite *WrapperIntegrationSuite) TestSetCompressor() {
	ctx := context.Background()
	t := suite.T()
	compressor := "s2-default"

	k, err := openKopiaRepo(t, ctx)
	require.NoError(t, err)

	defer func() {
		assert.NoError(t, k.Close(ctx))
	}()

	assert.NoError(t, k.Compression(ctx, compressor))

	// Check the policy was actually created and has the right compressor.
	p, err := policy.GetDefinedPolicy(ctx, k.Repository, policy.GlobalPolicySourceInfo)
	require.NoError(t, err)

	assert.Equal(t, compressor, string(p.CompressionPolicy.CompressorName))

	// Check the global policy will be the effective policy in future snapshots
	// for some source info.
	si := snapshot.SourceInfo{
		Host:     corsoHost,
		UserName: corsoUser,
		Path:     "test-path-root",
	}

	policyTree, err := policy.TreeForSource(ctx, k, si)
	require.NoError(t, err)

	assert.Equal(
		t,
		compressor,
		string(policyTree.EffectivePolicy().CompressionPolicy.CompressorName),
	)
}
