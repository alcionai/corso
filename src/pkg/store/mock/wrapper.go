package mock

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/kopia/kopia/repo/manifest"

	"github.com/alcionai/corso/src/internal/model"
	"github.com/alcionai/corso/src/pkg/backup"
	"github.com/alcionai/corso/src/pkg/store"
)

type BackupWrapper struct {
	Backup    *backup.Backup
	GetErr    error
	DeleteErr error
}

func (bw BackupWrapper) GetBackup(
	ctx context.Context,
	backupID model.StableID,
) (*backup.Backup, error) {
	bw.Backup.SnapshotID = bw.Backup.ID.String()

	return bw.Backup, clues.Stack(bw.GetErr).OrNil()
}

func (bw BackupWrapper) DeleteBackup(
	ctx context.Context,
	backupID model.StableID,
) error {
	return clues.Stack(bw.DeleteErr).OrNil()
}

func (bw BackupWrapper) GetBackups(
	ctx context.Context,
	filters ...store.FilterOption,
) ([]*backup.Backup, error) {
	return nil, clues.New("GetBackups mock not implemented yet")
}

func (bw BackupWrapper) DeleteWithModelStoreIDs(
	ctx context.Context,
	ids ...manifest.ID,
) error {
	return clues.Stack(bw.DeleteErr).OrNil()
}
