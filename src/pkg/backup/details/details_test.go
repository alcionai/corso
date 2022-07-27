package details_test

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
	"github.com/alcionai/corso/pkg/backup/details"
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
	deets = details.Details{
		DetailsModel: details.DetailsModel{
			BaseModel: model.BaseModel{
				ID:           model.StableID(detailsID),
				ModelStoreID: manifest.ID(uuid.NewString()),
			},
		},
	}
)

type StoreDetailsUnitSuite struct {
	suite.Suite
}

func TestStoreDetailsUnitSuite(t *testing.T) {
	suite.Run(t, new(StoreDetailsUnitSuite))
}

func (suite *StoreDetailsUnitSuite) TestGetDetails() {
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
			sm := &store.Wrapper{Storer: test.mock}
			result, err := sm.GetDetails(ctx, manifest.ID(uuid.NewString()))
			test.expect(t, err)
			if err != nil {
				return
			}
			assert.Equal(t, deets.ID, result.ID)
		})
	}
}

func (suite *StoreDetailsUnitSuite) TestGetDetailsFromBackupID() {
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
			store := &store.Wrapper{Storer: test.mock}
			dResult, bResult, err := store.GetDetailsFromBackupID(ctx, model.StableID(uuid.NewString()))
			test.expect(t, err)
			if err != nil {
				return
			}
			assert.Equal(t, deets.ID, dResult.ID)
			assert.Equal(t, bu.ID, bResult.ID)
		})
	}
}
