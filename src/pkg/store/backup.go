package store

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/kopia/kopia/repo/manifest"

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

type (
	BackupWrapper interface {
		BackupGetterDeleter
		GetBackups(
			ctx context.Context,
			filters ...FilterOption,
		) ([]*backup.Backup, error)
	}

	BackupGetterDeleter interface {
		BackupGetter
		BackupDeleter
	}

	BackupGetter interface {
		GetBackup(ctx context.Context, backupID model.StableID) (*backup.Backup, error)
	}

	BackupDeleter interface {
		DeleteBackup(ctx context.Context, backupID model.StableID) error
	}

	Storer interface {
		Delete(ctx context.Context, s model.Schema, id model.StableID) error
		DeleteWithModelStoreIDs(ctx context.Context, ids ...manifest.ID) error
		Get(ctx context.Context, s model.Schema, id model.StableID, data model.Model) error
		GetIDsForType(ctx context.Context, s model.Schema, tags map[string]string) ([]*model.BaseModel, error)
		GetWithModelStoreID(ctx context.Context, s model.Schema, id manifest.ID, data model.Model) error
		Put(ctx context.Context, s model.Schema, m model.Model) error
		Update(ctx context.Context, s model.Schema, m model.Model) error
	}

	BackupStorer interface {
		Storer
		BackupWrapper
	}
)

type wrapper struct {
	Storer
}

func NewWrapper(s Storer) *wrapper {
	return &wrapper{Storer: s}
}

// GetBackup gets a single backup by id.
func (w wrapper) GetBackup(ctx context.Context, backupID model.StableID) (*backup.Backup, error) {
	b := backup.Backup{}

	err := w.Get(ctx, model.BackupSchema, backupID, &b)
	if err != nil {
		return nil, clues.Wrap(err, "getting backup")
	}

	return &b, nil
}

// GetDetailsFromBackupID retrieves all backups in the model store.
func (w wrapper) GetBackups(
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
func (w wrapper) DeleteBackup(ctx context.Context, backupID model.StableID) error {
	return w.Delete(ctx, model.BackupSchema, backupID)
}
