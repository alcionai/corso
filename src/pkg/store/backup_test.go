package store_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/kopia/kopia/repo/manifest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/model"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/backup"
	"github.com/alcionai/corso/src/pkg/store"
	storeMock "github.com/alcionai/corso/src/pkg/store/mock"
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
	tester.Suite
}

func TestStoreBackupUnitSuite(t *testing.T) {
	suite.Run(t, &StoreBackupUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *StoreBackupUnitSuite) TestGetBackup() {
	ctx, flush := tester.NewContext()
	defer flush()

	table := []struct {
		name   string
		mock   *storeMock.MockModelStore
		expect assert.ErrorAssertionFunc
	}{
		{
			name:   "gets backup",
			mock:   storeMock.NewMock(&bu, nil),
			expect: assert.NoError,
		},
		{
			name:   "errors",
			mock:   storeMock.NewMock(&bu, assert.AnError),
			expect: assert.Error,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

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
	ctx, flush := tester.NewContext()
	defer flush()

	table := []struct {
		name   string
		mock   *storeMock.MockModelStore
		expect assert.ErrorAssertionFunc
	}{
		{
			name:   "gets backups",
			mock:   storeMock.NewMock(&bu, nil),
			expect: assert.NoError,
		},
		{
			name:   "errors",
			mock:   storeMock.NewMock(&bu, assert.AnError),
			expect: assert.Error,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

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

func (suite *StoreBackupUnitSuite) TestDeleteBackup() {
	ctx, flush := tester.NewContext()
	defer flush()

	table := []struct {
		name   string
		mock   *storeMock.MockModelStore
		expect assert.ErrorAssertionFunc
	}{
		{
			name:   "deletes backup",
			mock:   storeMock.NewMock(&bu, nil),
			expect: assert.NoError,
		},
		{
			name:   "errors",
			mock:   storeMock.NewMock(&bu, assert.AnError),
			expect: assert.Error,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			sm := &store.Wrapper{Storer: test.mock}
			err := sm.DeleteBackup(ctx, model.StableID(uuid.NewString()))
			test.expect(t, err)
		})
	}
}

func (suite *StoreBackupUnitSuite) TestGetDetailsIDFromBackupID() {
	ctx, flush := tester.NewContext()
	defer flush()

	table := []struct {
		name   string
		mock   *storeMock.MockModelStore
		expect assert.ErrorAssertionFunc
	}{
		{
			name:   "gets details from backup id",
			mock:   storeMock.NewMock(&bu, nil),
			expect: assert.NoError,
		},
		{
			name:   "errors",
			mock:   storeMock.NewMock(&bu, assert.AnError),
			expect: assert.Error,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			store := &store.Wrapper{Storer: test.mock}
			dResult, bResult, err := store.GetDetailsIDFromBackupID(ctx, model.StableID(uuid.NewString()))
			test.expect(t, err)
			if err != nil {
				return
			}
			assert.Equal(t, bu.DetailsID, dResult)
			assert.Equal(t, bu.ID, bResult.ID)
		})
	}
}
