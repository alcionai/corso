package inject

import (
	"context"

	"github.com/alcionai/canario/src/internal/common/prefixmatcher"
	"github.com/alcionai/canario/src/internal/data"
	"github.com/alcionai/canario/src/internal/kopia"
	"github.com/alcionai/canario/src/pkg/backup/details"
	"github.com/alcionai/canario/src/pkg/backup/identity"
	"github.com/alcionai/canario/src/pkg/count"
	"github.com/alcionai/canario/src/pkg/fault"
	"github.com/alcionai/canario/src/pkg/path"
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
			counter *count.Bus,
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
