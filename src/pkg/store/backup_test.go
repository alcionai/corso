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
			store := &store.Wrapper{test.mock}
			result, err := store.GetBackup(ctx, model.ID(uuid.NewString()))
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
			sm := &store.Wrapper{test.mock}
			result, err := sm.GetBackups(ctx)
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
		mock   *storeMock.MockModelStore
		expect assert.ErrorAssertionFunc
	}{
		{
			name:   "gets details",
			mock:   storeMock.NewMock(nil, &deets, nil),
			expect: assert.NoError,
		},
		{
			name:   "errors",
			mock:   storeMock.NewMock(nil, &deets, assert.AnError),
			expect: assert.Error,
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			sm := &store.Wrapper{test.mock}
			result, err := sm.GetDetails(ctx, manifest.ID(uuid.NewString()))
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
		mock   *storeMock.MockModelStore
		expect assert.ErrorAssertionFunc
	}{
		{
			name:   "gets details from backup id",
			mock:   storeMock.NewMock(&bu, &deets, nil),
			expect: assert.NoError,
		},
		{
			name:   "errors",
			mock:   storeMock.NewMock(&bu, &deets, assert.AnError),
			expect: assert.Error,
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			store := &store.Wrapper{test.mock}
			dResult, bResult, err := store.GetDetailsFromBackupID(ctx, model.ID(uuid.NewString()))
			test.expect(t, err)
			if err != nil {
				return
			}
			assert.Equal(t, deets.StableID, dResult.StableID)
			assert.Equal(t, bu.StableID, bResult.StableID)
		})
	}
}
