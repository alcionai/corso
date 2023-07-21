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
		PopulateProtectedResourceIDAndNamer
	}

	CacheItemInfoer interface {
		// CacheItemInfo is used by the consumer to cache metadata that is
		// sourced from per-item info, but may be valuable to the restore at
		// large.
		// Ex: pairing drive ids with drive names as they appeared at the time
		// of backup.
		CacheItemInfo(v details.ItemInfo)
	}

	PopulateProtectedResourceIDAndNamer interface {
		// PopulateProtectedResourceIDAndName takes the provided owner identifier and produces
		// the owner's name and ID from that value.  Returns an error if the owner is
		// not recognized by the current tenant.
		//
		// The id-name swapper should be optional.  Some processes will look up all owners in
		// the tenant before reaching this step.  In that case, the data gets handed
		// down for this func to consume instead of performing further queries.  The
		// data gets stored inside the controller instance for later re-use.
		PopulateProtectedResourceIDAndName(
			ctx context.Context,
			owner string, // input value, can be either id or name
			ins idname.Cacher,
		) (
			id, name string,
			err error,
		)
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
