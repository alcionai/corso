package inject

import (
	"context"

	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/internal/common/prefixmatcher"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/model"
	"github.com/alcionai/corso/src/pkg/backup"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/control/repository"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/export"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
)

type (
	BackupProducer interface {
		ProduceBackupCollections(
			ctx context.Context,
			resourceOwner idname.Provider,
			sels selectors.Selector,
			metadata []data.RestoreCollection,
			lastBackupVersion int,
			ctrlOpts control.Options,
			errs *fault.Bus,
		) ([]data.BackupCollection, prefixmatcher.StringSetReader, bool, error)
		IsBackupRunnable(ctx context.Context, service path.ServiceType, resourceOwner string) (bool, error)

		Wait() *data.CollectionStats
	}

	RestoreConsumer interface {
		ConsumeRestoreCollections(
			ctx context.Context,
			backupVersion int,
			selector selectors.Selector,
			restoreCfg control.RestoreConfig,
			opts control.Options,
			dcs []data.RestoreCollection,
			errs *fault.Bus,
			ctr *count.Bus,
		) (*details.Details, error)

		Wait() *data.CollectionStats

		CacheItemInfoer
	}

	CacheItemInfoer interface {
		// CacheItemInfo is used by the consumer to cache metadata that is
		// sourced from per-item info, but may be valuable to the restore at
		// large.
		// Ex: pairing drive ids with drive names as they appeared at the time
		// of backup.
		CacheItemInfo(v details.ItemInfo)
	}

	ExportConsumer interface {
		ExportRestoreCollections(
			ctx context.Context,
			backupVersion int,
			selector selectors.Selector,
			exportCfg control.ExportConfig,
			opts control.Options,
			dcs []data.RestoreCollection,
			errs *fault.Bus,
		) ([]export.Collection, error)

		Wait() *data.CollectionStats
	}

	RepoMaintenancer interface {
		RepoMaintenance(ctx context.Context, opts repository.Maintenance) error
	}

	GetBackuper interface {
		GetBackup(
			ctx context.Context,
			backupID model.StableID,
		) (*backup.Backup, error)
	}
)
