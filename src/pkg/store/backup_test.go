package store_test

import (
	"testing"
	"time"

	"github.com/alcionai/clues"
	"github.com/google/uuid"
	"github.com/kopia/kopia/repo/manifest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/model"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/backup"
	"github.com/alcionai/corso/src/pkg/store"
	"github.com/alcionai/corso/src/pkg/store/mock"
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
	table := []struct {
		name   string
		mock   *mock.ModelStore
		expect assert.ErrorAssertionFunc
	}{
		{
			name:   "gets backup",
			mock:   mock.NewModelStoreMock(&bu, nil),
			expect: assert.NoError,
		},
		{
			name:   "errors",
			mock:   mock.NewModelStoreMock(&bu, assert.AnError),
			expect: assert.Error,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			sm := store.NewWrapper(test.mock)

			result, err := sm.GetBackup(ctx, model.StableID(uuid.NewString()))
			test.expect(t, err, clues.ToCore(err))

			if err != nil {
				return
			}

			assert.Equal(t, bu.ID, result.ID)
		})
	}
}

func (suite *StoreBackupUnitSuite) TestGetBackups() {
	table := []struct {
		name   string
		mock   *mock.ModelStore
		expect assert.ErrorAssertionFunc
	}{
		{
			name:   "gets backups",
			mock:   mock.NewModelStoreMock(&bu, nil),
			expect: assert.NoError,
		},
		{
			name:   "errors",
			mock:   mock.NewModelStoreMock(&bu, assert.AnError),
			expect: assert.Error,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			sm := store.NewWrapper(test.mock)

			result, err := sm.GetBackups(ctx)
			test.expect(t, err, clues.ToCore(err))

			if err != nil {
				return
			}

			assert.Equal(t, 1, len(result))
			assert.Equal(t, bu.ID, result[0].ID)
		})
	}
}

func (suite *StoreBackupUnitSuite) TestDeleteBackup() {
	table := []struct {
		name   string
		mock   *mock.ModelStore
		expect assert.ErrorAssertionFunc
	}{
		{
			name:   "deletes backup",
			mock:   mock.NewModelStoreMock(&bu, nil),
			expect: assert.NoError,
		},
		{
			name:   "errors",
			mock:   mock.NewModelStoreMock(&bu, assert.AnError),
			expect: assert.Error,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			sm := store.NewWrapper(test.mock)

			err := sm.DeleteBackup(ctx, model.StableID(uuid.NewString()))
			test.expect(t, err, clues.ToCore(err))
		})
	}
}
