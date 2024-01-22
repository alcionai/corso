package kopia

import (
	"context"
	"math"
	"testing"
	"time"

	"github.com/alcionai/clues"
	"github.com/kopia/kopia/repo"
	"github.com/kopia/kopia/repo/blob"
	"github.com/kopia/kopia/repo/format"
	"github.com/kopia/kopia/snapshot"
	"github.com/kopia/kopia/snapshot/policy"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/ptr"
	strTD "github.com/alcionai/corso/src/internal/common/str/testdata"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/control/repository"
	"github.com/alcionai/corso/src/pkg/storage"
	storeTD "github.com/alcionai/corso/src/pkg/storage/testdata"
)

func openLocalKopiaRepo(
	t tester.TestT,
	ctx context.Context, //revive:disable-line:context-as-argument
) (*conn, error) {
	st := storeTD.NewFilesystemStorage(t)
	repoNameHash := strTD.NewHashForRepoConfigName()

	k := NewConn(st)
	if err := k.Initialize(ctx, repository.Options{}, repository.Retention{}, repoNameHash); err != nil {
		return nil, err
	}

	return k, nil
}

func openKopiaRepo(
	t tester.TestT,
	ctx context.Context, //revive:disable-line:context-as-argument
) (*conn, error) {
	st := storeTD.NewPrefixedS3Storage(t)
	repoNameHash := strTD.NewHashForRepoConfigName()

	k := NewConn(st)
	if err := k.Initialize(ctx, repository.Options{}, repository.Retention{}, repoNameHash); err != nil {
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
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	k := conn{}

	assert.NotPanics(t, func() {
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
			[][]string{storeTD.AWSStorageCredEnvs}),
	})
}

func (suite *WrapperIntegrationSuite) TestRepoExistsError() {
	t := suite.T()
	repoNameHash := strTD.NewHashForRepoConfigName()

	ctx, flush := tester.NewContext(t)
	defer flush()

	st := storeTD.NewFilesystemStorage(t)
	k := NewConn(st)

	err := k.Initialize(ctx, repository.Options{}, repository.Retention{}, repoNameHash)
	require.NoError(t, err, clues.ToCore(err))

	err = k.Close(ctx)
	require.NoError(t, err, clues.ToCore(err))

	err = k.Initialize(ctx, repository.Options{}, repository.Retention{}, repoNameHash)
	assert.Error(t, err, clues.ToCore(err))
	assert.ErrorIs(t, err, ErrorRepoAlreadyExists)
}

func (suite *WrapperIntegrationSuite) TestBadProviderErrors() {
	t := suite.T()
	repoNameHash := strTD.NewHashForRepoConfigName()

	ctx, flush := tester.NewContext(t)
	defer flush()

	st := storeTD.NewFilesystemStorage(t)
	st.Provider = storage.ProviderUnknown
	k := NewConn(st)

	err := k.Initialize(ctx, repository.Options{}, repository.Retention{}, repoNameHash)
	assert.Error(t, err, clues.ToCore(err))
}

func (suite *WrapperIntegrationSuite) TestConnectWithoutInitErrors() {
	t := suite.T()
	repoNameHash := strTD.NewHashForRepoConfigName()

	ctx, flush := tester.NewContext(t)
	defer flush()

	st := storeTD.NewFilesystemStorage(t)
	k := NewConn(st)

	err := k.Connect(ctx, repository.Options{}, repoNameHash)
	assert.Error(t, err, clues.ToCore(err))
}

func (suite *WrapperIntegrationSuite) TestCloseTwiceDoesNotCrash() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	k, err := openKopiaRepo(t, ctx)
	require.NoError(t, err, clues.ToCore(err))

	err = k.Close(ctx)
	assert.NoError(t, err, clues.ToCore(err))
	assert.Nil(t, k.Repository)

	err = k.Close(ctx)
	assert.NoError(t, err, clues.ToCore(err))
}

func (suite *WrapperIntegrationSuite) TestCloseAfterWrap() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

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
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	k, err := openKopiaRepo(t, ctx)
	require.NoError(t, err, clues.ToCore(err))

	err = k.Close(ctx)
	assert.NoError(t, err, clues.ToCore(err))

	err = k.wrap()
	assert.Error(t, err, clues.ToCore(err))
}

func (suite *WrapperIntegrationSuite) TestBadCompressorType() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

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
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	k, err := openKopiaRepo(t, ctx)
	require.NoError(t, err, clues.ToCore(err))

	defer func() {
		err := k.Close(ctx)
		assert.NoError(t, err, clues.ToCore(err))
	}()

	si := snapshot.SourceInfo{
		Host:     "exchangeemail",
		UserName: "tenantID-resourceID",
		Path:     "test-path-root",
	}

	p, err := k.getPolicyOrEmpty(ctx, si)
	require.NoError(t, err, clues.ToCore(err))
	assert.Equal(t, policy.Policy{}, *p)
}

func (suite *WrapperIntegrationSuite) TestSetCompressor() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

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
		Host:     "exchangeemail",
		UserName: "tenantID-resourceID",
		Path:     "test-path-root",
	}

	policyTree, err := policy.TreeForSource(ctx, k, si)
	require.NoError(t, err, clues.ToCore(err))
	assert.Equal(
		t,
		compressor,
		string(policyTree.EffectivePolicy().CompressionPolicy.CompressorName))
}

func (suite *WrapperIntegrationSuite) TestConfigPersistentConfigOnInitAndNotOnConnect() {
	repoNameHash := strTD.NewHashForRepoConfigName()

	table := []struct {
		name         string
		mutateParams repository.PersistentConfig
		checkFunc    func(
			t *testing.T,
			wanted repository.PersistentConfig,
			mutableParams format.MutableParameters,
			blobConfig format.BlobStorageConfiguration,
		)
	}{
		{
			name: "MinEpochDuration",
			mutateParams: repository.PersistentConfig{
				MinEpochDuration: ptr.To(defaultMinEpochDuration + time.Minute),
			},
			checkFunc: func(
				t *testing.T,
				wanted repository.PersistentConfig,
				mutableParams format.MutableParameters,
				blobConfig format.BlobStorageConfiguration,
			) {
				assert.Equal(t, *wanted.MinEpochDuration, mutableParams.EpochParameters.MinEpochDuration)
			},
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			k, err := openLocalKopiaRepo(t, ctx)
			require.NoError(t, err, clues.ToCore(err))

			// Close is safe to call even if the repo is already closed.
			t.Cleanup(func() {
				k.Close(ctx)
			})

			// Need to disconnect and connect again because in-memory state in kopia
			// isn't updated immediately.
			err = k.Close(ctx)
			require.NoError(t, err, clues.ToCore(err))

			err = k.Connect(ctx, repository.Options{}, repoNameHash)
			require.NoError(t, err, clues.ToCore(err))

			mutable, blob, err := k.getPersistentConfig(ctx)
			require.NoError(t, err, clues.ToCore(err))

			defaultParams := repository.PersistentConfig{
				MinEpochDuration: ptr.To(defaultMinEpochDuration),
			}

			test.checkFunc(t, defaultParams, mutable, blob)

			err = k.updatePersistentConfig(ctx, test.mutateParams)
			require.NoError(t, err, clues.ToCore(err))

			err = k.Close(ctx)
			require.NoError(t, err, clues.ToCore(err))

			err = k.Connect(ctx, repository.Options{}, repoNameHash)
			require.NoError(t, err, clues.ToCore(err))

			mutable, blob, err = k.getPersistentConfig(ctx)
			require.NoError(t, err, clues.ToCore(err))
			test.checkFunc(t, test.mutateParams, mutable, blob)

			err = k.Close(ctx)
			require.NoError(t, err, clues.ToCore(err))
		})
	}
}

func (suite *WrapperIntegrationSuite) TestConfigPolicyDefaultsSetOnInitAndNotOnConnect() {
	newCompressor := "pgzip"
	newRetentionDaily := policy.OptionalInt(42)
	newRetention := policy.RetentionPolicy{KeepDaily: &newRetentionDaily}
	newSchedInterval := time.Second * 42
	repoNameHash := strTD.NewHashForRepoConfigName()

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
					p.RetentionPolicy)
				assert.Equal(
					t,
					math.MaxInt,
					p.RetentionPolicy.EffectiveKeepLatest().OrDefault(42))
			},
			checkFunc: func(t *testing.T, p *policy.Policy) {
				t.Helper()
				require.Equal(
					t,
					newRetention,
					p.RetentionPolicy)
				assert.Equal(
					t,
					42,
					p.RetentionPolicy.EffectiveKeepLatest().OrDefault(42))
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
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

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

			err = k.Connect(ctx, repository.Options{}, repoNameHash)
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
	t := suite.T()
	repoNameHash := strTD.NewHashForRepoConfigName()

	ctx, flush := tester.NewContext(t)
	defer flush()

	k, err := openKopiaRepo(t, ctx)
	require.NoError(t, err, clues.ToCore(err))

	err = k.Close(ctx)
	require.NoError(t, err, clues.ToCore(err))

	// Re-open with Connect.
	err = k.Connect(ctx, repository.Options{}, repoNameHash)
	require.NoError(t, err, clues.ToCore(err))

	err = k.Close(ctx)
	assert.NoError(t, err, clues.ToCore(err))
}

func (suite *WrapperIntegrationSuite) TestSetUserAndHost() {
	t := suite.T()
	repoNameHash := strTD.NewHashForRepoConfigName()

	ctx, flush := tester.NewContext(t)
	defer flush()

	opts := repository.Options{
		User: "foo",
		Host: "bar",
	}

	st := storeTD.NewFilesystemStorage(t)
	k := NewConn(st)

	err := k.Initialize(ctx, opts, repository.Retention{}, repoNameHash)
	require.NoError(t, err, clues.ToCore(err))

	kopiaOpts := k.ClientOptions()
	require.Equal(t, opts.User, kopiaOpts.Username)
	require.Equal(t, opts.Host, kopiaOpts.Hostname)

	err = k.Close(ctx)
	require.NoError(t, err, clues.ToCore(err))

	// Re-open with Connect and a different user/hostname.
	opts.User = "hello"
	opts.Host = "world"

	err = k.Connect(ctx, opts, repoNameHash)
	require.NoError(t, err, clues.ToCore(err))

	kopiaOpts = k.ClientOptions()
	require.Equal(t, opts.User, kopiaOpts.Username)
	require.Equal(t, opts.Host, kopiaOpts.Hostname)

	err = k.Close(ctx)
	require.NoError(t, err, clues.ToCore(err))

	// Make sure not setting the values uses the kopia defaults.
	opts.User = ""
	opts.Host = ""

	err = k.Connect(ctx, opts, repoNameHash)
	require.NoError(t, err, clues.ToCore(err))

	kopiaOpts = k.ClientOptions()
	assert.NotEmpty(t, kopiaOpts.Username)
	assert.NotEqual(t, "hello", kopiaOpts.Username)
	assert.NotEmpty(t, kopiaOpts.Hostname)
	assert.NotEqual(t, "world", kopiaOpts.Hostname)

	err = k.Close(ctx)
	assert.NoError(t, err, clues.ToCore(err))
}

func (suite *WrapperIntegrationSuite) TestUpdatePersistentConfig() {
	table := []struct {
		name string
		opts func(
			startParams format.MutableParameters,
			startBlobConfig format.BlobStorageConfiguration,
		) repository.PersistentConfig
		expectErr    assert.ErrorAssertionFunc
		expectConfig func(
			startParams format.MutableParameters,
			startBlobConfig format.BlobStorageConfiguration,
		) (format.MutableParameters, format.BlobStorageConfiguration)
	}{
		{
			name: "NoOptionsSet NoChange",
			opts: func(
				startParams format.MutableParameters,
				startBlobConfig format.BlobStorageConfiguration,
			) repository.PersistentConfig {
				return repository.PersistentConfig{}
			},
			expectErr: assert.NoError,
			expectConfig: func(
				startParams format.MutableParameters,
				startBlobConfig format.BlobStorageConfiguration,
			) (format.MutableParameters, format.BlobStorageConfiguration) {
				return startParams, startBlobConfig
			},
		},
		{
			name: "NoValueChange NoChange",
			opts: func(
				startParams format.MutableParameters,
				startBlobConfig format.BlobStorageConfiguration,
			) repository.PersistentConfig {
				return repository.PersistentConfig{
					MinEpochDuration: ptr.To(startParams.EpochParameters.MinEpochDuration),
				}
			},
			expectErr: assert.NoError,
			expectConfig: func(
				startParams format.MutableParameters,
				startBlobConfig format.BlobStorageConfiguration,
			) (format.MutableParameters, format.BlobStorageConfiguration) {
				return startParams, startBlobConfig
			},
		},
		{
			name: "MinEpochLessThanLowerBound Errors",
			opts: func(
				startParams format.MutableParameters,
				startBlobConfig format.BlobStorageConfiguration,
			) repository.PersistentConfig {
				return repository.PersistentConfig{
					MinEpochDuration: ptr.To(minEpochDurationLowerBound - time.Second),
				}
			},
			expectErr: assert.Error,
			expectConfig: func(
				startParams format.MutableParameters,
				startBlobConfig format.BlobStorageConfiguration,
			) (format.MutableParameters, format.BlobStorageConfiguration) {
				return startParams, startBlobConfig
			},
		},
		{
			name: "MinEpochGreaterThanUpperBound Errors",
			opts: func(
				startParams format.MutableParameters,
				startBlobConfig format.BlobStorageConfiguration,
			) repository.PersistentConfig {
				return repository.PersistentConfig{
					MinEpochDuration: ptr.To(minEpochDurationUpperBound + time.Second),
				}
			},
			expectErr: assert.Error,
			expectConfig: func(
				startParams format.MutableParameters,
				startBlobConfig format.BlobStorageConfiguration,
			) (format.MutableParameters, format.BlobStorageConfiguration) {
				return startParams, startBlobConfig
			},
		},
		{
			name: "UpdateMinEpoch Succeeds",
			opts: func(
				startParams format.MutableParameters,
				startBlobConfig format.BlobStorageConfiguration,
			) repository.PersistentConfig {
				return repository.PersistentConfig{
					MinEpochDuration: ptr.To(minEpochDurationLowerBound),
				}
			},
			expectErr: assert.NoError,
			expectConfig: func(
				startParams format.MutableParameters,
				startBlobConfig format.BlobStorageConfiguration,
			) (format.MutableParameters, format.BlobStorageConfiguration) {
				startParams.EpochParameters.MinEpochDuration = minEpochDurationLowerBound
				return startParams, startBlobConfig
			},
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()
			repoNameHash := strTD.NewHashForRepoConfigName()

			ctx, flush := tester.NewContext(t)
			defer flush()

			st1 := storeTD.NewPrefixedS3Storage(t)

			connection := NewConn(st1)
			err := connection.Initialize(ctx, repository.Options{}, repository.Retention{}, repoNameHash)
			require.NoError(t, err, "initializing repo: %v", clues.ToCore(err))

			startParams, startBlobConfig, err := connection.getPersistentConfig(ctx)
			require.NoError(t, err, clues.ToCore(err))

			opts := test.opts(startParams, startBlobConfig)

			err = connection.updatePersistentConfig(ctx, opts)
			test.expectErr(t, err, clues.ToCore(err))

			// Need to close and reopen the repo since the format manager will cache
			// the old value for some amount of time.
			err = connection.Close(ctx)
			require.NoError(t, err, clues.ToCore(err))

			// Open another connection to the repo to verify params are as expected
			// across repo connect calls.
			err = connection.Connect(ctx, repository.Options{}, repoNameHash)
			require.NoError(t, err, clues.ToCore(err))

			gotParams, gotBlobConfig, err := connection.getPersistentConfig(ctx)
			require.NoError(t, err, clues.ToCore(err))

			expectParams, expectBlobConfig := test.expectConfig(startParams, startBlobConfig)

			assert.Equal(t, expectParams, gotParams)
			assert.Equal(t, expectBlobConfig, gotBlobConfig)
		})
	}
}

// ---------------
// integration tests that require object locking to be enabled on the bucket.
// ---------------
type ConnRetentionIntegrationSuite struct {
	tester.Suite
}

func TestConnRetentionIntegrationSuite(t *testing.T) {
	suite.Run(t, &ConnRetentionIntegrationSuite{
		Suite: tester.NewRetentionSuite(
			t,
			[][]string{storeTD.AWSStorageCredEnvs}),
	})
}

// Test that providing retention doesn't change anything but retention values
// from the default values that kopia uses.
func (suite *ConnRetentionIntegrationSuite) TestInitWithAndWithoutRetention() {
	t := suite.T()
	repoNameHash := strTD.NewHashForRepoConfigName()

	ctx, flush := tester.NewContext(t)
	defer flush()

	st1 := storeTD.NewPrefixedS3Storage(t)

	k1 := NewConn(st1)
	err := k1.Initialize(ctx, repository.Options{}, repository.Retention{}, repoNameHash)
	require.NoError(t, err, "initializing repo 1: %v", clues.ToCore(err))

	st2 := storeTD.NewPrefixedS3Storage(t)

	k2 := NewConn(st2)
	err = k2.Initialize(
		ctx,
		repository.Options{},
		repository.Retention{
			Mode:     ptr.To(repository.GovernanceRetention),
			Duration: ptr.To(time.Hour * 48),
			Extend:   ptr.To(true),
		},
		repoNameHash)
	require.NoError(t, err, "initializing repo 2: %v", clues.ToCore(err))

	dr1, ok := k1.Repository.(repo.DirectRepository)
	require.True(t, ok, "getting direct repo 1")

	dr2, ok := k2.Repository.(repo.DirectRepository)
	require.True(t, ok, "getting direct repo 2")

	format1 := dr1.FormatManager().ScrubbedContentFormat()
	format2 := dr2.FormatManager().ScrubbedContentFormat()

	assert.Equal(t, format1, format2)

	blobCfg1, err := dr1.FormatManager().BlobCfgBlob()
	require.NoError(t, err, "getting blob config 1: %v", clues.ToCore(err))

	blobCfg2, err := dr2.FormatManager().BlobCfgBlob()
	require.NoError(t, err, "getting retention config 2: %v", clues.ToCore(err))

	assert.NotEqual(t, blobCfg1, blobCfg2)

	// Check to make sure retention not enabled unexpectedly.
	checkRetentionParams(t, ctx, k1, blob.RetentionMode(""), 0, assert.False)

	// Some checks to make sure retention was fully initialized as expected.
	checkRetentionParams(t, ctx, k2, blob.Governance, time.Hour*48, assert.True)
}
