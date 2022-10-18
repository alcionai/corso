package mock

import (
	"context"

	"github.com/kopia/kopia/repo/manifest"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/model"
	"github.com/alcionai/corso/src/pkg/backup"
	"github.com/alcionai/corso/src/pkg/backup/details"
)

// ------------------------------------------------------------
// model wrapper model store
// ------------------------------------------------------------

type MockModelStore struct {
	backup  *backup.Backup
	details *details.Details
	err     error
}

func NewMock(b *backup.Backup, d *details.Details, err error) *MockModelStore {
	return &MockModelStore{
		backup:  b,
		details: d,
		err:     err,
	}
}

// ------------------------------------------------------------
// deleter iface
// ------------------------------------------------------------

func (mms *MockModelStore) Delete(ctx context.Context, s model.Schema, id model.StableID) error {
	return mms.err
}

func (mms *MockModelStore) DeleteWithModelStoreID(ctx context.Context, id manifest.ID) error {
	return mms.err
}

// ------------------------------------------------------------
// getter iface
// ------------------------------------------------------------

func (mms *MockModelStore) Get(
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

	case model.BackupDetailsSchema:
		dm := data.(*details.Details)
		dm.DetailsModel = mms.details.DetailsModel

	default:
		return errors.Errorf("schema %s not supported by mock Get", s)
	}

	return nil
}

func (mms *MockModelStore) GetIDsForType(
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

	case model.BackupDetailsSchema:
		d := details.Details{}
		d.DetailsModel = mms.details.DetailsModel

		return []*model.BaseModel{&d.BaseModel}, nil
	}

	return nil, errors.Errorf("schema %s not supported by mock GetIDsForType", s)
}

func (mms *MockModelStore) GetWithModelStoreID(
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

	case model.BackupDetailsSchema:
		dm := data.(*details.Details)
		dm.DetailsModel = mms.details.DetailsModel

	default:
		return errors.Errorf("schema %s not supported by mock GetWithModelStoreID", s)
	}

	return nil
}

// ------------------------------------------------------------
// updater iface
// ------------------------------------------------------------

func (mms *MockModelStore) Put(ctx context.Context, s model.Schema, m model.Model) error {
	switch s {
	case model.BackupSchema:
		bm := m.(*backup.Backup)
		mms.backup = bm

	case model.BackupDetailsSchema:
		dm := m.(*details.Details)
		mms.details = dm

	default:
		return errors.Errorf("schema %s not supported by mock Put", s)
	}

	return mms.err
}

func (mms *MockModelStore) Update(ctx context.Context, s model.Schema, m model.Model) error {
	switch s {
	case model.BackupSchema:
		bm := m.(*backup.Backup)
		mms.backup = bm

	case model.BackupDetailsSchema:
		dm := m.(*details.Details)
		mms.details = dm

	default:
		return errors.Errorf("schema %s not supported by mock Update", s)
	}

	return mms.err
}
