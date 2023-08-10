package inject

import (
	"context"

	"github.com/alcionai/corso/src/internal/common/prefixmatcher"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/kopia"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/backup/identity"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
)

type (
	BackupConsumer interface {
		ConsumeBackupCollections(
			ctx context.Context,
			backupReasons []identity.Reasoner,
			bases kopia.BackupBases,
			cs []data.BackupCollection,
			pmr prefixmatcher.StringSetReader,
			tags map[string]string,
			buildTreeWithBase bool,
			errs *fault.Bus,
		) (*kopia.BackupStats, *details.Builder, kopia.DetailsMergeInfoer, error)
	}

	RestoreProducer interface {
		ProduceRestoreCollections(
			ctx context.Context,
			snapshotID string,
			paths []path.RestorePaths,
			bc kopia.ByteCounter,
			errs *fault.Bus,
		) ([]data.RestoreCollection, error)
	}

	BaseFinder interface {
		FindBases(
			ctx context.Context,
			reasons []identity.Reasoner,
			tags map[string]string,
		) kopia.BackupBases
	}
)
