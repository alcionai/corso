package inject

import (
	"context"

	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/internal/common/prefixmatcher"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/model"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/backup"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/control/repository"
	"github.com/alcionai/corso/src/pkg/fault"
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
		) ([]data.BackupCollection, prefixmatcher.StringSetReader, error)

		Wait() *data.CollectionStats
	}

	RestoreConsumer interface {
		ConsumeRestoreCollections(
			ctx context.Context,
			backupVersion int,
			acct account.Account,
			selector selectors.Selector,
			dest control.RestoreDestination,
			opts control.Options,
			dcs []data.RestoreCollection,
			errs *fault.Bus,
		) (*details.Details, error)

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
