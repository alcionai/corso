package inject

import (
	"context"

	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/kopia"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
)

type (
	BackupProducer interface {
		ProduceBackupCollections(
			ctx context.Context,
			resourceOwner common.IDNamer,
			sels selectors.Selector,
			metadata []data.RestoreCollection,
			lastBackupVersion int,
			ctrlOpts control.Options,
			errs *fault.Bus,
		) ([]data.BackupCollection, map[string]map[string]struct{}, error)

		Wait() *data.CollectionStats
	}

	BackupConsumer interface {
		ConsumeBackupCollections(
			ctx context.Context,
			bases []kopia.IncrementalBase,
			cs []data.BackupCollection,
			excluded map[string]map[string]struct{},
			tags map[string]string,
			buildTreeWithBase bool,
			errs *fault.Bus,
		) (*kopia.BackupStats, *details.Builder, kopia.DetailsMergeInfoer, error)
	}

	RestoreProducer interface {
		ProduceRestoreCollections(
			ctx context.Context,
			snapshotID string,
			paths []path.Path,
			bc kopia.ByteCounter,
			errs *fault.Bus,
		) ([]data.RestoreCollection, error)
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
)
