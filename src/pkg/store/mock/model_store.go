package mock

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/kopia/kopia/repo/manifest"

	"github.com/alcionai/corso/src/internal/model"
	"github.com/alcionai/corso/src/pkg/backup"
)

// ------------------------------------------------------------
// model wrapper model store
// ------------------------------------------------------------

type ModelStore struct {
	backup *backup.Backup
	err    error
}

func NewModelStoreMock(b *backup.Backup, err error) *ModelStore {
	return &ModelStore{
		backup: b,
		err:    err,
	}
}

// ------------------------------------------------------------
// deleter iface
// ------------------------------------------------------------

func (mms *ModelStore) Delete(ctx context.Context, s model.Schema, id model.StableID) error {
	return mms.err
}

func (mms *ModelStore) DeleteWithModelStoreID(ctx context.Context, ids ...manifest.ID) error {
	return mms.err
}

// ------------------------------------------------------------
// getter iface
// ------------------------------------------------------------

func (mms *ModelStore) Get(
	ctx context.Context,
	s model.Schema,
	id model.StableID,
	data model.Model,
) error {
	if mms.err != nil {
		return mms.err
	}

	switch s {
	case model.BackupSchema:
		bm := data.(*backup.Backup)
		*bm = *mms.backup

	default:
		return clues.New("schema not supported by mock Get").With("schema", s)
	}

	return nil
}

func (mms *ModelStore) GetIDsForType(
	ctx context.Context,
	s model.Schema,
	tags map[string]string,
) ([]*model.BaseModel, error) {
	if mms.err != nil {
		return nil, mms.err
	}

	switch s {
	case model.BackupSchema:
		b := *mms.backup
		return []*model.BaseModel{&b.BaseModel}, nil
	}

	return nil, clues.New("schema not supported by mock GetIDsForType").With("schema", s)
}

func (mms *ModelStore) GetWithModelStoreID(
	ctx context.Context,
	s model.Schema,
	id manifest.ID,
	data model.Model,
) error {
	if mms.err != nil {
		return mms.err
	}

	switch s {
	case model.BackupSchema:
		bm := data.(*backup.Backup)
		*bm = *mms.backup

	default:
		return clues.New("schema not supported by mock GetWithModelStoreID").With("schema", s)
	}

	return nil
}

// ------------------------------------------------------------
// updater iface
// ------------------------------------------------------------

func (mms *ModelStore) Put(ctx context.Context, s model.Schema, m model.Model) error {
	switch s {
	case model.BackupSchema:
		bm := m.(*backup.Backup)
		mms.backup = bm

	default:
		return clues.New("schema not supported by mock Put").With("schema", s)
	}

	return mms.err
}

func (mms *ModelStore) Update(ctx context.Context, s model.Schema, m model.Model) error {
	switch s {
	case model.BackupSchema:
		bm := m.(*backup.Backup)
		mms.backup = bm

	default:
		return clues.New("schema not supported by mock Update").With("schema", s)
	}

	return mms.err
}
