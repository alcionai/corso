package store_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/kopia/kopia/repo/manifest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/internal/model"
	"github.com/alcionai/corso/pkg/backup"
	"github.com/alcionai/corso/pkg/store"
	storeMock "github.com/alcionai/corso/pkg/store/mock"
)

// ------------------------------------------------------------
// unit tests
// ------------------------------------------------------------

var (
	detailsID = uuid.NewString()
	bu        = backup.Backup{
		BaseModel: model.BaseModel{
			ID:           model.StableID(uuid.NewString()),
			ModelStoreID: manifest.ID(uuid.NewString()),
		},
		CreationTime: time.Now(),
		SnapshotID:   uuid.NewString(),
		DetailsID:    detailsID,
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
		mock   *storeMock.MockModelStore
		expect assert.ErrorAssertionFunc
	}{
		{
			name:   "gets backup",
			mock:   storeMock.NewMock(&bu, nil, nil),
			expect: assert.NoError,
		},
		{
			name:   "errors",
			mock:   storeMock.NewMock(&bu, nil, assert.AnError),
			expect: assert.Error,
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			sm := &store.Wrapper{Storer: test.mock}
			result, err := sm.GetBackup(ctx, model.StableID(uuid.NewString()))
			test.expect(t, err)
			if err != nil {
				return
			}
			assert.Equal(t, bu.ID, result.ID)
		})
	}
}

func (suite *StoreBackupUnitSuite) TestGetBackups() {
	ctx := context.Background()

	table := []struct {
		name   string
		mock   *storeMock.MockModelStore
		expect assert.ErrorAssertionFunc
	}{
		{
			name:   "gets backups",
			mock:   storeMock.NewMock(&bu, nil, nil),
			expect: assert.NoError,
		},
		{
			name:   "errors",
			mock:   storeMock.NewMock(&bu, nil, assert.AnError),
			expect: assert.Error,
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			sm := &store.Wrapper{Storer: test.mock}
			result, err := sm.GetBackups(ctx)
			test.expect(t, err)
			if err != nil {
				return
			}
			assert.Equal(t, 1, len(result))
			assert.Equal(t, bu.ID, result[0].ID)
		})
	}
}
