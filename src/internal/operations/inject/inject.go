package inject

import (
	"context"

	"github.com/kopia/kopia/repo/manifest"

	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/internal/common/prefixmatcher"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/kopia/inject"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/backup/identity"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/control/repository"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/export"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
)

type (
	BackupProducer interface {
		ProduceBackupCollections(
			ctx context.Context,
			bpc BackupProducerConfig,
			counter *count.Bus,
			errs *fault.Bus,
		) ([]data.BackupCollection, prefixmatcher.StringSetReader, bool, error)

		IsServiceEnableder

		// GetMetadataPaths returns a list of paths that form metadata
		// collections. In case of service that have just a single
		// underlying service like OneDrive or SharePoint, it will mostly
		// just have a single collection per manifest reason, but in the
		// case of groups, it will contain a collection each for the
		// underlying service, for example one per SharePoint site.
		GetMetadataPaths(
			ctx context.Context,
			r inject.RestoreProducer,
			base ReasonAndSnapshotIDer,
			errs *fault.Bus,
		) ([]path.RestorePaths, error)

		Wait() *data.CollectionStats

		// SetRateLimiter selects a rate limiter type for the service being
		// backed up and binds it to the context.
		SetRateLimiter(
			ctx context.Context,
			service path.ServiceType,
			options control.Options,
		) context.Context
	}

	RestoreConsumer interface {
		ConsumeRestoreCollections(
			ctx context.Context,
			rcc RestoreConsumerConfig,
			dcs []data.RestoreCollection,
			errs *fault.Bus,
			ctr *count.Bus,
		) (*details.Details, *data.CollectionStats, error)

		IsServiceEnableder

		Wait() *data.CollectionStats

		CacheItemInfoer
		PopulateProtectedResourceIDAndNamer
	}

	IsServiceEnableder interface {
		// IsServiceEnabled checks if the service is enabled for backup/restore
		// for the provided resource owner.
		IsServiceEnabled(
			ctx context.Context,
			service path.ServiceType,
			resourceOwner string,
		) (bool, error)
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
			exportCfg control.ExportConfig,
			dcs []data.RestoreCollection,
			stats *data.ExportStats,
			errs *fault.Bus,
		) ([]export.Collectioner, error)

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
		) (idname.Provider, error)
	}

	RepoMaintenancer interface {
		RepoMaintenance(ctx context.Context, opts repository.Maintenance) error
	}

	// ServiceHandler contains the set of functions required to implement all
	// service-specific functionality for backups, restores, and exports.
	ServiceHandler interface {
		ExportConsumer
	}

	ToServiceHandler interface {
		NewServiceHandler(
			opts control.Options,
			service path.ServiceType,
		) (ServiceHandler, error)
	}

	ReasonAndSnapshotIDer interface {
		GetReasons() []identity.Reasoner
		GetSnapshotID() manifest.ID
	}
)
