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

type DetailsUnitSuite struct {
	suite.Suite
}

func TestDetailsUnitSuite(t *testing.T) {
	suite.Run(t, new(DetailsUnitSuite))
}

func (suite *DetailsUnitSuite) TestGetDetails() {
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

func (suite *DetailsUnitSuite) TestGetDetailsFromBackupID() {
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

func (suite *DetailsUnitSuite) TestDetailsEntry_HeadersValues() {
	now := time.Now()
	nowStr := now.Format(time.RFC3339Nano)

	table := []struct {
		name     string
		entry    details.DetailsEntry
		expectHs []string
		expectVs []string
	}{
		{
			name: "no info",
			entry: details.DetailsEntry{
				RepoRef: "reporef",
			},
			expectHs: []string{"Repo Ref"},
			expectVs: []string{"reporef"},
		},
		{
			name: "exhange info",
			entry: details.DetailsEntry{
				RepoRef: "reporef",
				ItemInfo: details.ItemInfo{
					Exchange: &details.ExchangeInfo{
						Sender:   "sender",
						Subject:  "subject",
						Received: now,
					},
				},
			},
			expectHs: []string{"Repo Ref", "Sender", "Subject", "Received"},
			expectVs: []string{"reporef", "sender", "subject", nowStr},
		},
		{
			name: "sharepoint info",
			entry: details.DetailsEntry{
				RepoRef: "reporef",
				ItemInfo: details.ItemInfo{
					Sharepoint: &details.SharepointInfo{},
				},
			},
			expectHs: []string{"Repo Ref"},
			expectVs: []string{"reporef"},
		},
	}

	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			hs := test.entry.Headers()
			assert.Equal(t, test.expectHs, hs)
			vs := test.entry.Values()
			assert.Equal(t, test.expectVs, vs)
		})
	}
}

func (suite *DetailsUnitSuite) TestDetailsModel_Path() {
	table := []struct {
		name   string
		ents   []details.DetailsEntry
		expect []string
	}{
		{
			name:   "nil entries",
			ents:   nil,
			expect: []string{},
		},
		{
			name: "single entry",
			ents: []details.DetailsEntry{
				{RepoRef: "abcde"},
			},
			expect: []string{"abcde"},
		},
		{
			name: "multiple entries",
			ents: []details.DetailsEntry{
				{RepoRef: "abcde"},
				{RepoRef: "12345"},
			},
			expect: []string{"abcde", "12345"},
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			d := details.Details{
				DetailsModel: details.DetailsModel{
					Entries: test.ents,
				},
			}
			assert.Equal(t, test.expect, d.Paths())
		})
	}
}
