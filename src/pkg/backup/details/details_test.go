package details_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/pkg/backup/details"
)

// ------------------------------------------------------------
// unit tests
// ------------------------------------------------------------

type DetailsUnitSuite struct {
	suite.Suite
}

func TestDetailsUnitSuite(t *testing.T) {
	suite.Run(t, new(DetailsUnitSuite))
}

func (suite *DetailsUnitSuite) TestDetailsEntry_HeadersValues() {
	nowStr := common.FormatNow(common.TabularOutputTimeFormat)
	now, err := common.ParseTime(nowStr)
	require.NoError(suite.T(), err)

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
			name: "exchange event info",
			entry: details.DetailsEntry{
				RepoRef: "reporef",
				ItemInfo: details.ItemInfo{
					Exchange: &details.ExchangeInfo{
						ItemType:    details.ExchangeEvent,
						EventStart:  now,
						Organizer:   "organizer",
						EventRecurs: true,
						Subject:     "subject",
					},
				},
			},
			expectHs: []string{"Repo Ref", "Organizer", "Subject", "Starts", "Recurring"},
			expectVs: []string{"reporef", "organizer", "subject", nowStr, "true"},
		},
		{
			name: "exchange contact info",
			entry: details.DetailsEntry{
				RepoRef: "reporef",
				ItemInfo: details.ItemInfo{
					Exchange: &details.ExchangeInfo{
						ItemType:    details.ExchangeContact,
						ContactName: "contactName",
					},
				},
			},
			expectHs: []string{"Repo Ref", "Contact Name"},
			expectVs: []string{"reporef", "contactName"},
		},
		{
			name: "exchange mail info",
			entry: details.DetailsEntry{
				RepoRef: "reporef",
				ItemInfo: details.ItemInfo{
					Exchange: &details.ExchangeInfo{
						ItemType: details.ExchangeMail,
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
		{
			name: "oneDrive info",
			entry: details.DetailsEntry{
				RepoRef: "reporef",
				ItemInfo: details.ItemInfo{
					OneDrive: &details.OneDriveInfo{
						ItemName:   "itemName",
						ParentPath: "parentPath",
					},
				},
			},
			expectHs: []string{"Repo Ref", "ItemName", "ParentPath"},
			expectVs: []string{"reporef", "itemName", "parentPath"},
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
