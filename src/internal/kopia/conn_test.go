package kopia

import (
	"context"
	"math"
	"testing"
	"time"

	"github.com/kopia/kopia/snapshot"
	"github.com/kopia/kopia/snapshot/policy"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/clues"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/storage"
)

//revive:disable:context-as-argument
func openKopiaRepo(t *testing.T, ctx context.Context) (*conn, error) {
	//revive:enable:context-as-argument
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
	tester.Suite
}

func TestWrapperUnitSuite(t *testing.T) {
	suite.Run(t, &WrapperUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *WrapperUnitSuite) TestCloseWithoutOpenDoesNotCrash() {
	ctx, flush := tester.NewContext()
	defer flush()

	k := conn{}

	assert.NotPanics(suite.T(), func() {
		k.Close(ctx)
	})
}

// ---------------
// integration tests that use kopia
// ---------------
type WrapperIntegrationSuite struct {
	tester.Suite
}

func TestWrapperIntegrationSuite(t *testing.T) {
	suite.Run(t, &WrapperIntegrationSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tester.AWSStorageCredEnvs},
		),
	})
}

func (suite *WrapperIntegrationSuite) TestRepoExistsError() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()

	st := tester.NewPrefixedS3Storage(t)
	k := NewConn(st)

	err := k.Initialize(ctx)
	require.NoError(t, err, clues.ToCore(err))

	err = k.Close(ctx)
	require.NoError(t, err, clues.ToCore(err))

	err = k.Initialize(ctx)
	assert.Error(t, err, clues.ToCore(err))
	assert.ErrorIs(t, err, ErrorRepoAlreadyExists)
}

func (suite *WrapperIntegrationSuite) TestBadProviderErrors() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()
	st := tester.NewPrefixedS3Storage(t)
	st.Provider = storage.ProviderUnknown
	k := NewConn(st)

	err := k.Initialize(ctx)
	assert.Error(t, err, clues.ToCore(err))
}

func (suite *WrapperIntegrationSuite) TestConnectWithoutInitErrors() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()
	st := tester.NewPrefixedS3Storage(t)
	k := NewConn(st)

	err := k.Connect(ctx)
	assert.Error(t, err, clues.ToCore(err))
}

func (suite *WrapperIntegrationSuite) TestCloseTwiceDoesNotCrash() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()

	k, err := openKopiaRepo(t, ctx)
	require.NoError(t, err, clues.ToCore(err))

	err = k.Close(ctx)
	assert.NoError(t, err, clues.ToCore(err))
	assert.Nil(t, k.Repository)

	err = k.Close(ctx)
	assert.NoError(t, err, clues.ToCore(err))
}

func (suite *WrapperIntegrationSuite) TestCloseAfterWrap() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()

	k, err := openKopiaRepo(t, ctx)
	require.NoError(t, err, clues.ToCore(err))

	err = k.wrap()
	require.NoError(t, err, clues.ToCore(err))

	assert.Equal(t, 2, k.refCount)

	err = k.Close(ctx)
	require.NoError(t, err, clues.ToCore(err))
	assert.NotNil(t, k.Repository)
	assert.Equal(t, 1, k.refCount)

	err = k.Close(ctx)
	require.NoError(t, err, clues.ToCore(err))
	assert.Nil(t, k.Repository)
	assert.Equal(t, 0, k.refCount)
}

func (suite *WrapperIntegrationSuite) TestOpenAfterClose() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()

	k, err := openKopiaRepo(t, ctx)
	require.NoError(t, err, clues.ToCore(err))

	err = k.Close(ctx)
	assert.NoError(t, err, clues.ToCore(err))

	err = k.wrap()
	assert.Error(t, err, clues.ToCore(err))
}

func (suite *WrapperIntegrationSuite) TestBadCompressorType() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()

	k, err := openKopiaRepo(t, ctx)
	require.NoError(t, err, clues.ToCore(err))

	defer func() {
		err := k.Close(ctx)
		assert.NoError(t, err, clues.ToCore(err))
	}()

	err = k.Compression(ctx, "not-a-compressor")
	assert.Error(t, err, clues.ToCore(err))
}

func (suite *WrapperIntegrationSuite) TestGetPolicyOrDefault_GetsDefault() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()

	k, err := openKopiaRepo(t, ctx)
	require.NoError(t, err, clues.ToCore(err))

	defer func() {
		err := k.Close(ctx)
		assert.NoError(t, err, clues.ToCore(err))
	}()

	si := snapshot.SourceInfo{
		Host:     corsoHost,
		UserName: corsoUser,
		Path:     "test-path-root",
	}

	p, err := k.getPolicyOrEmpty(ctx, si)
	require.NoError(t, err, clues.ToCore(err))
	assert.Equal(t, policy.Policy{}, *p)
}

func (suite *WrapperIntegrationSuite) TestSetCompressor() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()
	compressor := "pgzip"

	k, err := openKopiaRepo(t, ctx)
	require.NoError(t, err, clues.ToCore(err))

	defer func() {
		err := k.Close(ctx)
		assert.NoError(t, err, clues.ToCore(err))
	}()

	err = k.Compression(ctx, compressor)
	assert.NoError(t, err, clues.ToCore(err))

	// Check the policy was actually created and has the right compressor.
	p, err := k.getPolicyOrEmpty(ctx, policy.GlobalPolicySourceInfo)
	require.NoError(t, err, clues.ToCore(err))
	assert.Equal(t, compressor, string(p.CompressionPolicy.CompressorName))

	// Check the global policy will be the effective policy in future snapshots
	// for some source info.
	si := snapshot.SourceInfo{
		Host:     corsoHost,
		UserName: corsoUser,
		Path:     "test-path-root",
	}

	policyTree, err := policy.TreeForSource(ctx, k, si)
	require.NoError(t, err, clues.ToCore(err))
	assert.Equal(
		t,
		compressor,
		string(policyTree.EffectivePolicy().CompressionPolicy.CompressorName))
}

func (suite *WrapperIntegrationSuite) TestConfigDefaultsSetOnInitAndNotOnConnect() {
	newCompressor := "pgzip"
	newRetentionDaily := policy.OptionalInt(42)
	newRetention := policy.RetentionPolicy{KeepDaily: &newRetentionDaily}
	newSchedInterval := time.Second * 42

	table := []struct {
		name          string
		checkInitFunc func(*testing.T, *policy.Policy)
		checkFunc     func(*testing.T, *policy.Policy)
		mutator       func(context.Context, *policy.Policy) error
	}{
		{
			name: "Compression",
			checkInitFunc: func(t *testing.T, p *policy.Policy) {
				t.Helper()
				require.Equal(t, defaultCompressor, string(p.CompressionPolicy.CompressorName))
			},
			checkFunc: func(t *testing.T, p *policy.Policy) {
				t.Helper()
				require.Equal(t, newCompressor, string(p.CompressionPolicy.CompressorName))
			},
			mutator: func(innerCtx context.Context, p *policy.Policy) error {
				_, res := updateCompressionOnPolicy(newCompressor, p)
				return res
			},
		},
		{
			name: "Retention",
			checkInitFunc: func(t *testing.T, p *policy.Policy) {
				t.Helper()
				require.Equal(
					t,
					defaultRetention,
					p.RetentionPolicy,
				)
				assert.Equal(
					t,
					math.MaxInt,
					p.RetentionPolicy.EffectiveKeepLatest().OrDefault(42),
				)
			},
			checkFunc: func(t *testing.T, p *policy.Policy) {
				t.Helper()
				require.Equal(
					t,
					newRetention,
					p.RetentionPolicy,
				)
				assert.Equal(
					t,
					42,
					p.RetentionPolicy.EffectiveKeepLatest().OrDefault(42),
				)
			},
			mutator: func(innerCtx context.Context, p *policy.Policy) error {
				updateRetentionOnPolicy(newRetention, p)

				return nil
			},
		},
		{
			name: "Scheduling",
			checkInitFunc: func(t *testing.T, p *policy.Policy) {
				t.Helper()
				require.Equal(t, defaultSchedulingInterval, p.SchedulingPolicy.Interval())
			},
			checkFunc: func(t *testing.T, p *policy.Policy) {
				t.Helper()
				require.Equal(t, newSchedInterval, p.SchedulingPolicy.Interval())
			},
			mutator: func(innerCtx context.Context, p *policy.Policy) error {
				updateSchedulingOnPolicy(newSchedInterval, p)

				return nil
			},
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			ctx, flush := tester.NewContext()
			defer flush()

			t := suite.T()

			k, err := openKopiaRepo(t, ctx)
			require.NoError(t, err, clues.ToCore(err))

			p, err := k.getPolicyOrEmpty(ctx, policy.GlobalPolicySourceInfo)
			require.NoError(t, err, clues.ToCore(err))

			test.checkInitFunc(t, p)

			err = test.mutator(ctx, p)
			require.NoError(t, err, clues.ToCore(err))

			err = k.writeGlobalPolicy(ctx, "TestDefaultPolicyConfigSet", p)
			require.NoError(t, err, clues.ToCore(err))

			err = k.Close(ctx)
			require.NoError(t, err, clues.ToCore(err))

			err = k.Connect(ctx)
			require.NoError(t, err, clues.ToCore(err))

			defer func() {
				err := k.Close(ctx)
				assert.NoError(t, err, clues.ToCore(err))
			}()

			p, err = k.getPolicyOrEmpty(ctx, policy.GlobalPolicySourceInfo)
			require.NoError(t, err, clues.ToCore(err))
			test.checkFunc(t, p)
		})
	}
}

func (suite *WrapperIntegrationSuite) TestInitAndConnWithTempDirectory() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()

	k, err := openKopiaRepo(t, ctx)
	require.NoError(t, err, clues.ToCore(err))

	err = k.Close(ctx)
	require.NoError(t, err, clues.ToCore(err))

	// Re-open with Connect.
	err = k.Connect(ctx)
	require.NoError(t, err, clues.ToCore(err))

	err = k.Close(ctx)
	assert.NoError(t, err, clues.ToCore(err))
}
