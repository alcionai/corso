package mock

import (
	"context"

	"github.com/alcionai/clues"

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
	return bw.Backup, bw.GetErr
}

func (bw BackupWrapper) DeleteBackup(
	ctx context.Context,
	backupID model.StableID,
) error {
	return bw.DeleteErr
}

func (bw BackupWrapper) GetBackups(
	ctx context.Context,
	filters ...store.FilterOption,
) ([]*backup.Backup, error) {
	return nil, clues.New("GetBackups mock not implemented yet")
}
