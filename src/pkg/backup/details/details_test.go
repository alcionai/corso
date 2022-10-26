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
	nowStr := common.FormatNow(common.TabularOutput)
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
				RepoRef:  "reporef",
				ShortRef: "deadbeef",
			},
			expectHs: []string{"ID"},
			expectVs: []string{"deadbeef"},
		},
		{
			name: "exchange event info",
			entry: details.DetailsEntry{
				RepoRef:  "reporef",
				ShortRef: "deadbeef",
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
			expectHs: []string{"ID", "Organizer", "Subject", "Starts", "Recurring"},
			expectVs: []string{"deadbeef", "organizer", "subject", nowStr, "true"},
		},
		{
			name: "exchange contact info",
			entry: details.DetailsEntry{
				RepoRef:  "reporef",
				ShortRef: "deadbeef",
				ItemInfo: details.ItemInfo{
					Exchange: &details.ExchangeInfo{
						ItemType:    details.ExchangeContact,
						ContactName: "contactName",
					},
				},
			},
			expectHs: []string{"ID", "Contact Name"},
			expectVs: []string{"deadbeef", "contactName"},
		},
		{
			name: "exchange mail info",
			entry: details.DetailsEntry{
				RepoRef:  "reporef",
				ShortRef: "deadbeef",
				ItemInfo: details.ItemInfo{
					Exchange: &details.ExchangeInfo{
						ItemType: details.ExchangeMail,
						Sender:   "sender",
						Subject:  "subject",
						Received: now,
					},
				},
			},
			expectHs: []string{"ID", "Sender", "Subject", "Received"},
			expectVs: []string{"deadbeef", "sender", "subject", nowStr},
		},
		{
			name: "sharepoint info",
			entry: details.DetailsEntry{
				RepoRef:  "reporef",
				ShortRef: "deadbeef",
				ItemInfo: details.ItemInfo{
					Sharepoint: &details.SharepointInfo{},
				},
			},
			expectHs: []string{"ID"},
			expectVs: []string{"deadbeef"},
		},
		{
			name: "oneDrive info",
			entry: details.DetailsEntry{
				RepoRef:  "reporef",
				ShortRef: "deadbeef",
				ItemInfo: details.ItemInfo{
					OneDrive: &details.OneDriveInfo{
						ItemName:   "itemName",
						ParentPath: "parentPath",
						Size:       1000,
						Created:    now,
						Modified:   now,
					},
				},
			},
			expectHs: []string{"ID", "ItemName", "ParentPath", "Size", "Created", "Modified"},
			expectVs: []string{"deadbeef", "itemName", "parentPath", "1.0 kB", nowStr, nowStr},
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

var pathItemsTable = []struct {
	name             string
	ents             []details.DetailsEntry
	expectRefs       []string
	expectFolderRefs []string
}{
	{
		name:       "nil entries",
		ents:       nil,
		expectRefs: []string{},
	},
	{
		name: "single entry",
		ents: []details.DetailsEntry{
			{RepoRef: "abcde"},
		},
		expectRefs: []string{"abcde"},
	},
	{
		name: "multiple entries",
		ents: []details.DetailsEntry{
			{RepoRef: "abcde"},
			{RepoRef: "12345"},
		},
		expectRefs: []string{"abcde", "12345"},
	},
	{
		name: "multiple entries with folder",
		ents: []details.DetailsEntry{
			{RepoRef: "abcde"},
			{RepoRef: "12345"},
			{
				RepoRef: "deadbeef",
				ItemInfo: details.ItemInfo{
					Folder: &details.FolderInfo{
						DisplayName: "test folder",
					},
				},
			},
		},
		expectRefs:       []string{"abcde", "12345"},
		expectFolderRefs: []string{"deadbeef"},
	},
}

func (suite *DetailsUnitSuite) TestDetailsModel_Path() {
	for _, test := range pathItemsTable {
		suite.T().Run(test.name, func(t *testing.T) {
			d := details.Details{
				DetailsModel: details.DetailsModel{
					Entries: test.ents,
				},
			}
			assert.Equal(t, test.expectRefs, d.Paths())
		})
	}
}

func (suite *DetailsUnitSuite) TestDetailsModel_Items() {
	for _, test := range pathItemsTable {
		suite.T().Run(test.name, func(t *testing.T) {
			d := details.Details{
				DetailsModel: details.DetailsModel{
					Entries: test.ents,
				},
			}

			ents := d.Items()
			assert.Len(t, ents, len(test.expectRefs))

			for _, e := range ents {
				assert.Contains(t, test.expectRefs, e.RepoRef)
			}
		})
	}
}

func (suite *DetailsUnitSuite) TestDetailsModel_Folders() {
	for _, test := range pathItemsTable {
		suite.T().Run(test.name, func(t *testing.T) {
			d := details.Details{
				DetailsModel: details.DetailsModel{
					Entries: test.ents,
				},
			}

			ents := d.Folders()
			assert.Len(t, ents, len(test.expectFolderRefs))

			for _, e := range ents {
				assert.Contains(t, test.expectFolderRefs, e.RepoRef)
			}
		})
	}
}

func (suite *DetailsUnitSuite) TestDetails_AddFolders() {
	table := []struct {
		name              string
		folders           []details.FolderEntry
		expectedShortRefs []string
	}{
		{
			name: "MultipleFolders",
			folders: []details.FolderEntry{
				{
					RepoRef:   "rr1",
					ShortRef:  "sr1",
					ParentRef: "pr1",
				},
				{
					RepoRef:   "rr2",
					ShortRef:  "sr2",
					ParentRef: "pr2",
				},
			},
			expectedShortRefs: []string{"sr1", "sr2"},
		},
		{
			name: "MultipleFoldersWithRepeats",
			folders: []details.FolderEntry{
				{
					RepoRef:   "rr1",
					ShortRef:  "sr1",
					ParentRef: "pr1",
				},
				{
					RepoRef:   "rr2",
					ShortRef:  "sr2",
					ParentRef: "pr2",
				},
				{
					RepoRef:   "rr1",
					ShortRef:  "sr1",
					ParentRef: "pr1",
				},
				{
					RepoRef:   "rr3",
					ShortRef:  "sr3",
					ParentRef: "pr3",
				},
			},
			expectedShortRefs: []string{"sr1", "sr2", "sr3"},
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			deets := details.Details{}
			deets.AddFolders(test.folders)

			assert.Len(t, deets.Entries, len(test.expectedShortRefs))

			for _, e := range deets.Entries {
				assert.Contains(t, test.expectedShortRefs, e.ShortRef)
			}
		})
	}
}
