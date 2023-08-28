package inject

import (
	"context"

	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/internal/common/prefixmatcher"
	"github.com/alcionai/corso/src/internal/data"
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
			bpc BackupProducerConfig,
			errs *fault.Bus,
		) ([]data.BackupCollection, prefixmatcher.StringSetReader, bool, error)

		IsServiceEnabled(
			ctx context.Context,
			service path.ServiceType,
			resourceOwner string,
		) (bool, error)

		Wait() *data.CollectionStats
	}

	RestoreConsumer interface {
		ConsumeRestoreCollections(
			ctx context.Context,
			rcc RestoreConsumerConfig,
			dcs []data.RestoreCollection,
			errs *fault.Bus,
			ctr *count.Bus,
		) (*details.Details, error)

		IsServiceEnabled(
			ctx context.Context,
			service path.ServiceType,
			resourceOwner string,
		) (bool, error)

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

	ExportConsumer interface {
		ProduceExportCollections(
			ctx context.Context,
			backupVersion int,
			selector selectors.Selector,
			exportCfg control.ExportConfig,
			opts control.Options,
			dcs []data.RestoreCollection,
			errs *fault.Bus,
		) ([]export.Collection, error)

		Wait() *data.CollectionStats

		CacheItemInfoer
	}

	PopulateProtectedResourceIDAndNamer interface {
		// PopulateProtectedResourceIDAndName takes the provided owner identifier and produces
		// the owner's name and ID from that value.  Returns an error if the owner is
		// not recognized by the current tenant.
		//
		// The id-name cacher should be optional.  Some processes will look up all owners in
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
)
