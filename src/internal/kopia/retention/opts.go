package retention

import (
	"context"
	"time"

	"github.com/alcionai/clues"
	"github.com/kopia/kopia/repo/blob"
	"github.com/kopia/kopia/repo/format"
	"github.com/kopia/kopia/repo/maintenance"

	"github.com/alcionai/corso/src/pkg/control/repository"
)

type Opts struct {
	blobCfg format.BlobStorageConfiguration
	params  maintenance.Params

	blobChanged   bool
	paramsChanged bool
}

func NewOpts() *Opts {
	return &Opts{}
}

func OptsFromConfigs(
	blobCfg format.BlobStorageConfiguration,
	params maintenance.Params,
) *Opts {
	return &Opts{
		blobCfg: blobCfg,
		params:  params,
	}
}

func (r *Opts) AsConfigs(
	ctx context.Context,
) (format.BlobStorageConfiguration, maintenance.Params, error) {
	// Check the new config is valid.
	if r.blobCfg.IsRetentionEnabled() {
		if err := maintenance.CheckExtendRetention(ctx, r.blobCfg, &r.params); err != nil {
			return format.BlobStorageConfiguration{}, maintenance.Params{}, clues.Wrap(
				err,
				"invalid retention config",
			).WithClues(ctx)
		}
	}

	return r.blobCfg, r.params, nil
}

func (r *Opts) BlobChanged() bool {
	return r.blobChanged
}

func (r *Opts) ParamsChanged() bool {
	return r.paramsChanged
}

func (r *Opts) Set(opts repository.Retention) error {
	r.setMaintenanceParams(opts.Extend)

	return clues.Wrap(
		r.setBlobConfigParams(opts.Mode, opts.Duration),
		"setting mode or duration",
	).OrNil()
}

func (r *Opts) setMaintenanceParams(extend *bool) {
	if extend != nil && r.params.ExtendObjectLocks != *extend {
		r.params.ExtendObjectLocks = *extend
		r.paramsChanged = true
	}
}

func (r *Opts) setBlobConfigParams(
	mode *repository.RetentionMode,
	duration *time.Duration,
) error {
	err := r.setBlobConfigMode(mode)
	if err != nil {
		return clues.Stack(err)
	}

	r.setBlobConfigDuration(duration)

	return nil
}

func (r *Opts) setBlobConfigDuration(duration *time.Duration) {
	if duration != nil && r.blobCfg.RetentionPeriod != *duration {
		r.blobCfg.RetentionPeriod = *duration
		r.blobChanged = true
	}
}

func (r *Opts) setBlobConfigMode(
	mode *repository.RetentionMode,
) error {
	if mode == nil {
		return nil
	}

	startMode := r.blobCfg.RetentionMode

	switch *mode {
	case repository.NoRetention:
		if !r.blobCfg.IsRetentionEnabled() {
			return nil
		}

		r.blobCfg.RetentionMode = ""
		r.blobCfg.RetentionPeriod = 0

	case repository.GovernanceRetention:
		r.blobCfg.RetentionMode = blob.Governance

	case repository.ComplianceRetention:
		r.blobCfg.RetentionMode = blob.Compliance

	default:
		return clues.New("unknown retention mode").
			With("provided_retention_mode", mode.String())
	}

	// Only check if the retention mode is not empty. IsValid errors out if it's
	// empty.
	if len(r.blobCfg.RetentionMode) > 0 && !r.blobCfg.RetentionMode.IsValid() {
		return clues.New("invalid retention mode").
			With("retention_mode", r.blobCfg.RetentionMode)
	}

	// Take into account previous operations on r that could have already updated
	// blobChanged.
	r.blobChanged = r.blobChanged || startMode != r.blobCfg.RetentionMode

	return nil
}
