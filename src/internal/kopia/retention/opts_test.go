package retention_test

import (
	"testing"
	"time"

	"github.com/alcionai/clues"
	"github.com/kopia/kopia/repo/blob"
	"github.com/kopia/kopia/repo/format"
	"github.com/kopia/kopia/repo/maintenance"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/kopia/retention"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/control/repository"
)

type OptsUnitSuite struct {
	tester.Suite
}

func TestOptsUnitSuite(t *testing.T) {
	suite.Run(t, &OptsUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *OptsUnitSuite) TestOptsFromConfigs() {
	var (
		t = suite.T()

		mode     = blob.Governance
		duration = time.Hour * 48
		extend   = true

		blobCfgInput = format.BlobStorageConfiguration{
			RetentionMode:   mode,
			RetentionPeriod: duration,
		}
		paramsInput = maintenance.Params{ExtendObjectLocks: extend}
	)

	ctx, flush := tester.NewContext(t)
	defer flush()

	opts := retention.OptsFromConfigs(blobCfgInput, paramsInput)

	assert.False(t, opts.BlobChanged(), "BlobChanged")
	assert.False(t, opts.ParamsChanged(), "ParamsChanged")

	blobCfg, params, err := opts.AsConfigs(ctx)
	require.NoError(t, err, "AsConfigs: %v", clues.ToCore(err))
	assert.Equal(t, blobCfgInput, blobCfg)
	assert.Equal(t, paramsInput, params)
}

func (suite *OptsUnitSuite) TestSet() {
	var (
		kopiaMode = blob.Governance
		mode      = repository.GovernanceRetention
		duration  = time.Hour * 48
	)

	table := []struct {
		name                string
		inputBlob           format.BlobStorageConfiguration
		inputParams         maintenance.Params
		ctrlOpts            repository.Retention
		setErr              require.ErrorAssertionFunc
		expectMode          blob.RetentionMode
		expectDuration      time.Duration
		expectExtend        bool
		expectBlobChanged   bool
		expectParamsChanged bool
	}{
		{
			name:   "All Nils",
			setErr: require.NoError,
		},
		{
			name: "All Off",
			ctrlOpts: repository.Retention{
				Mode:     ptr.To(repository.NoRetention),
				Duration: ptr.To(time.Duration(0)),
				Extend:   ptr.To(false),
			},
			setErr: require.NoError,
		},
		{
			name: "UnknownRetention",
			ctrlOpts: repository.Retention{
				Mode:     ptr.To(repository.UnknownRetention),
				Duration: ptr.To(duration),
			},
			setErr: require.Error,
		},
		{
			name: "Invalid Retention Mode",
			ctrlOpts: repository.Retention{
				Mode:     ptr.To(repository.RetentionMode(-1)),
				Duration: ptr.To(duration),
			},
			setErr: require.Error,
		},
		{
			name: "Valid Set All",
			ctrlOpts: repository.Retention{
				Mode:     ptr.To(mode),
				Duration: ptr.To(duration),
				Extend:   ptr.To(true),
			},
			setErr:              require.NoError,
			expectMode:          kopiaMode,
			expectDuration:      duration,
			expectExtend:        true,
			expectBlobChanged:   true,
			expectParamsChanged: true,
		},
		{
			name: "Valid Set BlobConfig",
			ctrlOpts: repository.Retention{
				Mode:     ptr.To(mode),
				Duration: ptr.To(duration),
			},
			setErr:            require.NoError,
			expectMode:        kopiaMode,
			expectDuration:    duration,
			expectBlobChanged: true,
		},
		{
			name: "Valid Set Params",
			ctrlOpts: repository.Retention{
				Extend: ptr.To(true),
			},
			setErr:              require.NoError,
			expectExtend:        true,
			expectParamsChanged: true,
		},
		{
			name: "Partial BlobConfig Change",
			inputBlob: format.BlobStorageConfiguration{
				RetentionMode:   kopiaMode,
				RetentionPeriod: duration,
			},
			ctrlOpts: repository.Retention{
				Duration: ptr.To(duration + time.Hour),
			},
			setErr:            require.NoError,
			expectMode:        kopiaMode,
			expectDuration:    duration + time.Hour,
			expectBlobChanged: true,
		},
		{
			name: "No BlobConfig Change",
			inputBlob: format.BlobStorageConfiguration{
				RetentionMode:   kopiaMode,
				RetentionPeriod: duration,
			},
			ctrlOpts: repository.Retention{
				Mode:     ptr.To(mode),
				Duration: ptr.To(duration),
			},
			setErr:         require.NoError,
			expectMode:     kopiaMode,
			expectDuration: duration,
		},
		{
			name:        "No Params Change",
			inputParams: maintenance.Params{ExtendObjectLocks: true},
			ctrlOpts: repository.Retention{
				Extend: ptr.To(true),
			},
			setErr:       require.NoError,
			expectExtend: true,
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			opts := retention.OptsFromConfigs(test.inputBlob, test.inputParams)
			err := opts.Set(test.ctrlOpts)
			test.setErr(t, err, "setting params: %v", clues.ToCore(err))

			if err != nil {
				return
			}

			blobCfg, params, err := opts.AsConfigs(ctx)
			require.NoError(t, err, "getting configs: %v", clues.ToCore(err))

			assert.Equal(t, test.expectMode, blobCfg.RetentionMode, "mode")
			assert.Equal(t, test.expectDuration, blobCfg.RetentionPeriod, "duration")
			assert.Equal(t, test.expectExtend, params.ExtendObjectLocks, "extend locks")
			assert.Equal(t, test.expectBlobChanged, opts.BlobChanged(), "blob changed")
			assert.Equal(t, test.expectParamsChanged, opts.ParamsChanged(), "params changed")
		})
	}
}
