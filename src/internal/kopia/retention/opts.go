package retention

import (
  "time"

  "github.com/alcionai/clues"
  "github.com/kopia/kopia/repo/blob"
  "github.com/kopia/kopia/repo/format"

  "github.com/alcionai/corso/src/pkg/control/repository"
)

func setBlobConfigParams(
	mode *repository.RetentionMode,
	duration *time.Duration,
	blobCfg *format.BlobStorageConfiguration,
) (bool, error) {
	changed, err := setBlobConfigMode(mode, blobCfg)
	if err != nil {
		return false, clues.Stack(err)
	}

	tmp := setBlobConfigDuration(duration, blobCfg)
	changed = changed || tmp

	return changed, nil
}

func setBlobConfigDuration(
	duration *time.Duration,
	blobCfg *format.BlobStorageConfiguration,
) bool {
	var changed bool

	if duration != nil && blobCfg.RetentionPeriod != *duration {
		blobCfg.RetentionPeriod = *duration
		changed = true
	}

	return changed
}

func setBlobConfigMode(
	mode *repository.RetentionMode,
	blobCfg *format.BlobStorageConfiguration,
) (bool, error) {
	if mode == nil {
		return false, nil
	}

	startMode := blobCfg.RetentionMode

	switch *mode {
	case repository.NoRetention:
		if !blobCfg.IsRetentionEnabled() {
			return false, nil
		}

		blobCfg.RetentionMode = ""
		blobCfg.RetentionPeriod = 0

	case repository.GovernanceRetention:
		blobCfg.RetentionMode = blob.Governance

	case repository.ComplianceRetention:
		blobCfg.RetentionMode = blob.Compliance

	default:
		return false, clues.New("unknown retention mode").
			With("provided_retention_mode", mode.String())
	}

	// Only check if the retention mode is not empty. IsValid errors out if it's
	// empty.
	if len(blobCfg.RetentionMode) > 0 && !blobCfg.RetentionMode.IsValid() {
		return false, clues.New("invalid retention mode").
			With("retention_mode", blobCfg.RetentionMode)
	}

	return startMode != blobCfg.RetentionMode, nil
}

