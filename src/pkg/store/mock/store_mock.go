package mock

import (
	"context"
	"encoding/json"

	"github.com/kopia/kopia/repo/manifest"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/internal/model"
	"github.com/alcionai/corso/pkg/backup"
	"github.com/alcionai/corso/pkg/backup/details"
)

// ------------------------------------------------------------
// model wrapper model store
// ------------------------------------------------------------

type MockModelStore struct {
	backup  []byte
	details []byte
	err     error
}

func NewMock(b *backup.Backup, d *details.Details, err error) *MockModelStore {
	return &MockModelStore{
		backup:  marshal(b),
		details: marshal(d),
		err:     err,
	}
}

func marshal(a any) []byte {
	bs, _ := json.Marshal(a)
	return bs
}

func unmarshal(b []byte, a any) {
	//nolint:errcheck
	json.Unmarshal(b, a)
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
		unmarshal(mms.backup, data)
	case model.BackupDetailsSchema:
		unmarshal(mms.details, data)
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
		b := backup.Backup{}
		unmarshal(mms.backup, &b)
		return []*model.BaseModel{&b.BaseModel}, nil
	case model.BackupDetailsSchema:
		d := details.Details{}
		unmarshal(mms.backup, &d)
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
		unmarshal(mms.backup, data)
	case model.BackupDetailsSchema:
		unmarshal(mms.details, data)
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
		mms.backup = marshal(m)
	case model.BackupDetailsSchema:
		mms.details = marshal(m)
	default:
		return errors.Errorf("schema %s not supported by mock Put", s)
	}
	return mms.err
}

func (mms *MockModelStore) Update(ctx context.Context, s model.Schema, m model.Model) error {
	switch s {
	case model.BackupSchema:
		mms.backup = marshal(m)
	case model.BackupDetailsSchema:
		mms.details = marshal(m)
	default:
		return errors.Errorf("schema %s not supported by mock Update", s)
	}
	return mms.err
}
