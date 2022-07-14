package store_test

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/alcionai/corso/internal/kopia"
	"github.com/alcionai/corso/internal/model"
	"github.com/alcionai/corso/pkg/backup"
	"github.com/alcionai/corso/pkg/store"
	"github.com/google/uuid"
	"github.com/kopia/kopia/repo/manifest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// ------------------------------------------------------------
// getter mock
// ------------------------------------------------------------

type mockModelStoreGetter struct {
	backup  []byte
	details []byte
	err     error
}

func NewMock(b *backup.Backup, d *backup.Details, err error) mockModelStoreGetter {
	return mockModelStoreGetter{
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
	//nolint
	json.Unmarshal(b, a)
}

func (m mockModelStoreGetter) Get(
	ctx context.Context,
	t kopia.ModelType,
	id model.ID,
	data model.Model,
) error {
	if m.err != nil {
		return m.err
	}
	if t == kopia.BackupModel {
		unmarshal(m.backup, data)
	} else {
		unmarshal(m.details, data)
	}
	return nil
}

func (m mockModelStoreGetter) GetIDsForType(
	ctx context.Context,
	t kopia.ModelType,
	tags map[string]string,
) ([]*model.BaseModel, error) {
	if m.err != nil {
		return nil, m.err
	}
	if t == kopia.BackupModel {
		b := backup.Backup{}
		unmarshal(m.backup, &b)
		return []*model.BaseModel{&b.BaseModel}, nil
	}
	d := backup.Details{}
	unmarshal(m.backup, &d)
	return []*model.BaseModel{&d.BaseModel}, nil
}

func (m mockModelStoreGetter) GetWithModelStoreID(
	ctx context.Context,
	t kopia.ModelType,
	id manifest.ID,
	data model.Model,
) error {
	if m.err != nil {
		return m.err
	}
	if t == kopia.BackupModel {
		unmarshal(m.backup, data)
	} else {
		unmarshal(m.details, data)
	}
	return nil
}

// ------------------------------------------------------------
// unit tests
// ------------------------------------------------------------

var (
	detailsID = uuid.NewString()
	bu        = backup.Backup{
		BaseModel: model.BaseModel{
			StableID:     model.ID(uuid.NewString()),
			ModelStoreID: manifest.ID(uuid.NewString()),
		},
		CreationTime: time.Now(),
		SnapshotID:   uuid.NewString(),
		DetailsID:    detailsID,
	}
	deets = backup.Details{
		DetailsModel: backup.DetailsModel{
			BaseModel: model.BaseModel{
				StableID:     model.ID(detailsID),
				ModelStoreID: manifest.ID(uuid.NewString()),
			},
		},
	}
)

type StoreBackupUnitSuite struct {
	suite.Suite
}

func TestStoreBackupUnitSuite(t *testing.T) {
	suite.Run(t, new(StoreBackupUnitSuite))
}

func (suite *StoreBackupUnitSuite) TestGetBackup() {
	ctx := context.Background()

	table := []struct {
		name   string
		mock   mockModelStoreGetter
		expect assert.ErrorAssertionFunc
	}{
		{
			name:   "gets backup",
			mock:   NewMock(&bu, nil, nil),
			expect: assert.NoError,
		},
		{
			name:   "errors",
			mock:   NewMock(&bu, nil, assert.AnError),
			expect: assert.Error,
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			result, err := store.GetBackup(ctx, test.mock, model.ID(uuid.NewString()))
			test.expect(t, err)
			if err != nil {
				return
			}
			assert.Equal(t, bu.StableID, result.StableID)
		})
	}
}

func (suite *StoreBackupUnitSuite) TestGetBackups() {
	ctx := context.Background()

	table := []struct {
		name   string
		mock   mockModelStoreGetter
		expect assert.ErrorAssertionFunc
	}{
		{
			name:   "gets backups",
			mock:   NewMock(&bu, nil, nil),
			expect: assert.NoError,
		},
		{
			name:   "errors",
			mock:   NewMock(&bu, nil, assert.AnError),
			expect: assert.Error,
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			result, err := store.GetBackups(ctx, test.mock)
			test.expect(t, err)
			if err != nil {
				return
			}
			assert.Equal(t, 1, len(result))
			assert.Equal(t, bu.StableID, result[0].StableID)
		})
	}
}

func (suite *StoreBackupUnitSuite) TestGetDetails() {
	ctx := context.Background()

	table := []struct {
		name   string
		mock   mockModelStoreGetter
		expect assert.ErrorAssertionFunc
	}{
		{
			name:   "gets details",
			mock:   NewMock(nil, &deets, nil),
			expect: assert.NoError,
		},
		{
			name:   "errors",
			mock:   NewMock(nil, &deets, assert.AnError),
			expect: assert.Error,
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			result, err := store.GetDetails(ctx, test.mock, manifest.ID(uuid.NewString()))
			test.expect(t, err)
			if err != nil {
				return
			}
			assert.Equal(t, deets.StableID, result.StableID)
		})
	}
}

func (suite *StoreBackupUnitSuite) TestGetDetailsFromBackupID() {
	ctx := context.Background()

	table := []struct {
		name   string
		mock   mockModelStoreGetter
		expect assert.ErrorAssertionFunc
	}{
		{
			name:   "gets details from backup id",
			mock:   NewMock(&bu, &deets, nil),
			expect: assert.NoError,
		},
		{
			name:   "errors",
			mock:   NewMock(&bu, &deets, assert.AnError),
			expect: assert.Error,
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			dResult, bResult, err := store.GetDetailsFromBackupID(ctx, test.mock, model.ID(uuid.NewString()))
			test.expect(t, err)
			if err != nil {
				return
			}
			assert.Equal(t, deets.StableID, dResult.StableID)
			assert.Equal(t, bu.StableID, bResult.StableID)
		})
	}
}
