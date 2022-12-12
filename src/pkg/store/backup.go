package store

import (
	"context"

	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/model"
	"github.com/alcionai/corso/src/pkg/backup"
	"github.com/alcionai/corso/src/pkg/path"
)

type queryFilters struct {
	tags map[string]string
}

type FilterOption func(*queryFilters)

func (q *queryFilters) populate(qf ...FilterOption) {
	if len(qf) == 0 {
		return
	}

	q.tags = map[string]string{}

	for _, fn := range qf {
		fn(q)
	}
}

// Service ensures the retrieved backups only match
// the specified service.
func Service(pst path.ServiceType) FilterOption {
	return func(qf *queryFilters) {
		qf.tags[model.ServiceTag] = pst.String()
	}
}

// GetBackup gets a single backup by id.
func (w Wrapper) GetBackup(ctx context.Context, backupID model.StableID) (*backup.Backup, error) {
	b := backup.Backup{}

	err := w.Get(ctx, model.BackupSchema, backupID, &b)
	if err != nil {
		return nil, errors.Wrap(err, "getting backup")
	}

	return &b, nil
}

// GetDetailsFromBackupID retrieves all backups in the model store.
func (w Wrapper) GetBackups(
	ctx context.Context,
	filters ...FilterOption,
) ([]*backup.Backup, error) {
	q := &queryFilters{}
	q.populate(filters...)

	bms, err := w.GetIDsForType(ctx, model.BackupSchema, q.tags)
	if err != nil {
		return nil, err
	}

	bs := make([]*backup.Backup, len(bms))

	for i, bm := range bms {
		b := &backup.Backup{}

		err := w.GetWithModelStoreID(ctx, model.BackupSchema, bm.ModelStoreID, b)
		if err != nil {
			return nil, err
		}

		bs[i] = b
	}

	return bs, nil
}

// DeleteBackup deletes the backup and its details entry from the model store.
func (w Wrapper) DeleteBackup(ctx context.Context, backupID model.StableID) error {
	return w.Delete(ctx, model.BackupSchema, backupID)
}

// GetDetailsFromBackupID retrieves the backup.Details within the specified backup.
func (w Wrapper) GetDetailsIDFromBackupID(
	ctx context.Context,
	backupID model.StableID,
) (string, *backup.Backup, error) {
	b, err := w.GetBackup(ctx, backupID)
	if err != nil {
		return "", nil, err
	}

	return b.DetailsID, b, nil
}
