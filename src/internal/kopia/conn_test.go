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

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/aw"
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
			tester.CorsoKopiaWrapperTests,
		),
	})
}

func (suite *WrapperIntegrationSuite) TestRepoExistsError() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()

	st := tester.NewPrefixedS3Storage(t)
	k := NewConn(st)
	aw.MustNoErr(t, k.Initialize(ctx))

	aw.MustNoErr(t, k.Close(ctx))

	err := k.Initialize(ctx)
	aw.Err(t, err)
	aw.ErrIs(t, err, ErrorRepoAlreadyExists)
}

func (suite *WrapperIntegrationSuite) TestBadProviderErrors() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()

	st := tester.NewPrefixedS3Storage(t)
	st.Provider = storage.ProviderUnknown

	k := NewConn(st)
	aw.Err(t, k.Initialize(ctx))
}

func (suite *WrapperIntegrationSuite) TestConnectWithoutInitErrors() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()

	st := tester.NewPrefixedS3Storage(t)
	k := NewConn(st)
	aw.Err(t, k.Connect(ctx))
}

func (suite *WrapperIntegrationSuite) TestCloseTwiceDoesNotCrash() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()

	k, err := openKopiaRepo(t, ctx)
	aw.MustNoErr(t, err)
	aw.NoErr(t, k.Close(ctx))
	assert.Nil(t, k.Repository)
	aw.NoErr(t, k.Close(ctx))
}

func (suite *WrapperIntegrationSuite) TestCloseAfterWrap() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()

	k, err := openKopiaRepo(t, ctx)
	aw.MustNoErr(t, err)

	aw.MustNoErr(t, k.wrap())

	assert.Equal(t, 2, k.refCount)

	aw.MustNoErr(t, k.Close(ctx))
	assert.NotNil(t, k.Repository)
	assert.Equal(t, 1, k.refCount)

	aw.MustNoErr(t, k.Close(ctx))
	assert.Nil(t, k.Repository)
	assert.Equal(t, 0, k.refCount)
}

func (suite *WrapperIntegrationSuite) TestOpenAfterClose() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()

	k, err := openKopiaRepo(t, ctx)
	aw.MustNoErr(t, err)

	aw.NoErr(t, k.Close(ctx))
	aw.Err(t, k.wrap())
}

func (suite *WrapperIntegrationSuite) TestBadCompressorType() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()

	k, err := openKopiaRepo(t, ctx)
	aw.MustNoErr(t, err)

	defer func() {
		aw.NoErr(t, k.Close(ctx))
	}()

	aw.Err(t, k.Compression(ctx, "not-a-compressor"))
}

func (suite *WrapperIntegrationSuite) TestGetPolicyOrDefault_GetsDefault() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()

	k, err := openKopiaRepo(t, ctx)
	aw.MustNoErr(t, err)

	defer func() {
		aw.NoErr(t, k.Close(ctx))
	}()

	si := snapshot.SourceInfo{
		Host:     corsoHost,
		UserName: corsoUser,
		Path:     "test-path-root",
	}

	p, err := k.getPolicyOrEmpty(ctx, si)
	aw.MustNoErr(t, err)

	assert.Equal(t, policy.Policy{}, *p)
}

func (suite *WrapperIntegrationSuite) TestSetCompressor() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()
	compressor := "pgzip"

	k, err := openKopiaRepo(t, ctx)
	aw.MustNoErr(t, err)

	defer func() {
		aw.NoErr(t, k.Close(ctx))
	}()

	aw.NoErr(t, k.Compression(ctx, compressor))

	// Check the policy was actually created and has the right compressor.
	p, err := k.getPolicyOrEmpty(ctx, policy.GlobalPolicySourceInfo)
	aw.MustNoErr(t, err)

	assert.Equal(t, compressor, string(p.CompressionPolicy.CompressorName))

	// Check the global policy will be the effective policy in future snapshots
	// for some source info.
	si := snapshot.SourceInfo{
		Host:     corsoHost,
		UserName: corsoUser,
		Path:     "test-path-root",
	}

	policyTree, err := policy.TreeForSource(ctx, k, si)
	aw.MustNoErr(t, err)

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
				assert.Equal(
					t,
					math.MaxInt,
					p.RetentionPolicy.EffectiveKeepLatest().OrDefault(42),
				)
			},
			mutator: func(innerCtx context.Context, p *policy.Policy) error {
				newRetentionDaily := policy.OptionalInt(42)
				newRetention := policy.RetentionPolicy{KeepDaily: &newRetentionDaily}
				updateRetentionOnPolicy(newRetention, p)

				return nil
			},
		},
		{
			name: "Scheduling",
			checkFunc: func(t *testing.T, p *policy.Policy) {
				t.Helper()
				require.Equal(t, time.Second*0, p.SchedulingPolicy.Interval())
			},
			mutator: func(innerCtx context.Context, p *policy.Policy) error {
				updateSchedulingOnPolicy(time.Second*42, p)

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
			aw.MustNoErr(t, err)

			p, err := k.getPolicyOrEmpty(ctx, policy.GlobalPolicySourceInfo)
			aw.MustNoErr(t, err)

			test.checkFunc(t, p)

			aw.MustNoErr(t, test.mutator(ctx, p))
			aw.MustNoErr(t, k.writeGlobalPolicy(ctx, "TestDefaultPolicyConfigSet", p))
			aw.MustNoErr(t, k.Close(ctx))

			aw.MustNoErr(t, k.Connect(ctx))

			defer func() {
				aw.NoErr(t, k.Close(ctx))
			}()

			p, err = k.getPolicyOrEmpty(ctx, policy.GlobalPolicySourceInfo)
			aw.MustNoErr(t, err)

			test.checkFunc(t, p)
		})
	}
}

func (suite *WrapperIntegrationSuite) TestInitAndConnWithTempDirectory() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()

	k, err := openKopiaRepo(t, ctx)
	aw.MustNoErr(t, err)
	aw.MustNoErr(t, k.Close(ctx))

	// Re-open with Connect.
	aw.MustNoErr(t, k.Connect(ctx))
	aw.NoErr(t, k.Close(ctx))
}
