package kopia

import (
	"context"
	"testing"

	"github.com/kopia/kopia/snapshot"
	"github.com/kopia/kopia/snapshot/policy"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
)

//revive:disable:context-as-argument
func openKopiaRepo(t *testing.T, ctx context.Context) (*conn, error) {
	st := tester.NewPrefixedS3Storage(t)

	k := NewConn(st)
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

func (suite *WrapperIntegrationSuite) TestGetPolicyOrDefault_GetsDefault() {
	ctx := context.Background()
	t := suite.T()

	k, err := openKopiaRepo(t, ctx)
	require.NoError(t, err)

	defer func() {
		assert.NoError(t, k.Close(ctx))
	}()

	si := snapshot.SourceInfo{
		Host:     corsoHost,
		UserName: corsoUser,
		Path:     "test-path-root",
	}

	p, err := k.getPolicyOrEmpty(ctx, si)
	require.NoError(t, err)

	assert.Equal(t, policy.Policy{}, *p)
}

func (suite *WrapperIntegrationSuite) TestSetCompressor() {
	ctx := context.Background()
	t := suite.T()
	compressor := "pgzip"

	k, err := openKopiaRepo(t, ctx)
	require.NoError(t, err)

	defer func() {
		assert.NoError(t, k.Close(ctx))
	}()

	assert.NoError(t, k.Compression(ctx, compressor))

	// Check the policy was actually created and has the right compressor.
	p, err := k.getPolicyOrEmpty(ctx, policy.GlobalPolicySourceInfo)
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

func (suite *WrapperIntegrationSuite) TestConfigDefaultsSetOnInitAndConnect() {
	table := []struct {
		name      string
		checkFunc func(*testing.T, *policy.Policy)
		mutator   func(context.Context, *policy.Policy) error
	}{
		{
			name: "Compression",
			checkFunc: func(t *testing.T, p *policy.Policy) {
				t.Helper()
				require.Equal(t, defaultCompressor, string(p.CompressionPolicy.CompressorName))
			},
			mutator: func(innerCtx context.Context, p *policy.Policy) error {
				_, res := updateCompressionOnPolicy("pgzip", p)
				return res
			},
		},
		{
			name: "Retention",
			checkFunc: func(t *testing.T, p *policy.Policy) {
				t.Helper()
				require.Equal(
					t,
					defaultRetention,
					p.RetentionPolicy,
				)
			},
			mutator: func(innerCtx context.Context, p *policy.Policy) error {
				newRetentionDaily := policy.OptionalInt(42)
				newRetention := policy.RetentionPolicy{KeepDaily: &newRetentionDaily}
				updateRetentionOnPolicy(newRetention, p)

				return nil
			},
		},
	}

	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			ctx := context.Background()

			k, err := openKopiaRepo(t, ctx)
			require.NoError(t, err)

			p, err := k.getPolicyOrEmpty(ctx, policy.GlobalPolicySourceInfo)
			require.NoError(t, err)

			test.checkFunc(t, p)

			require.NoError(t, test.mutator(ctx, p))
			require.NoError(t, k.writeGlobalPolicy(ctx, "TestDefaultPolicyConfigSet", p))
			require.NoError(t, k.Close(ctx))

			require.NoError(t, k.Connect(ctx))

			defer func() {
				assert.NoError(t, k.Close(ctx))
			}()

			p, err = k.getPolicyOrEmpty(ctx, policy.GlobalPolicySourceInfo)
			require.NoError(t, err)

			test.checkFunc(t, p)
		})
	}
}

func (suite *WrapperIntegrationSuite) TestInitAndConnWithTempDirectory() {
	ctx := context.Background()
	t := suite.T()

	k, err := openKopiaRepo(t, ctx)
	require.NoError(t, err)
	require.NoError(t, k.Close(ctx))

	// Re-open with Connect.
	require.NoError(t, k.Connect(ctx))
	assert.NoError(t, k.Close(ctx))
}
