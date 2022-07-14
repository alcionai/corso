package store

import (
	"context"

	"github.com/kopia/kopia/repo/manifest"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/internal/kopia"
	"github.com/alcionai/corso/internal/model"
	"github.com/alcionai/corso/pkg/backup"
)

type modelStoreGetter interface {
	Get(ctx context.Context, t kopia.ModelType, id model.ID, data model.Model) error
	GetIDsForType(ctx context.Context, t kopia.ModelType, tags map[string]string) ([]*model.BaseModel, error)
	GetWithModelStoreID(ctx context.Context, t kopia.ModelType, id manifest.ID, data model.Model) error
}

var _ modelStoreGetter = &kopia.ModelStore{}

// GetBackup gets a single backup by id.
func GetBackup(ctx context.Context, ms modelStoreGetter, backupID model.ID) (*backup.Backup, error) {
	b := backup.Backup{}
	err := ms.Get(ctx, kopia.BackupModel, backupID, &b)
	if err != nil {
		return nil, errors.Wrap(err, "getting backup")
	}
	return &b, nil
}

// GetDetailsFromBackupID retrieves all backups in the model store.
func GetBackups(ctx context.Context, ms modelStoreGetter) ([]*backup.Backup, error) {
	bms, err := ms.GetIDsForType(ctx, kopia.BackupModel, nil)
	if err != nil {
		return nil, err
	}
	bs := make([]*backup.Backup, len(bms))
	for i, bm := range bms {
		b := backup.Backup{}
		err := ms.GetWithModelStoreID(ctx, kopia.BackupModel, bm.ModelStoreID, &b)
		if err != nil {
			return nil, err
		}
		bs[i] = &b
	}
	return bs, nil
}

// GetDetails gets the backup details by ID.
func GetDetails(ctx context.Context, ms modelStoreGetter, detailsID manifest.ID) (*backup.Details, error) {
	d := backup.Details{}
	err := ms.GetWithModelStoreID(ctx, kopia.BackupDetailsModel, detailsID, &d)
	if err != nil {
		return nil, errors.Wrap(err, "getting details")
	}
	return &d, nil
}

// GetDetailsFromBackupID retrieves the backup.Details within the specified backup.
func GetDetailsFromBackupID(
	ctx context.Context,
	ms modelStoreGetter,
	backupID model.ID,
) (*backup.Details, *backup.Backup, error) {
	b, err := GetBackup(ctx, ms, backupID)
	if err != nil {
		return nil, nil, err
	}

	d, err := GetDetails(ctx, ms, manifest.ID(b.DetailsID))
	if err != nil {
		return nil, nil, err
	}

	return d, b, nil
}
